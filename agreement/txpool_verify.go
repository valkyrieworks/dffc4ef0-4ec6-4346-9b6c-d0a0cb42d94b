package agreement

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
func validateTxpool(txn transferObserver) txpooll.Txpool {
	return txn.(txpooll.Txpool)
}

func VerifyTxpoolNegativeOnwardTillTransAccessible(t *testing.T) {
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	settings.Agreement.GenerateVoidLedgers = false
	status, privateItems := arbitraryInaugurationStatus(1, false, 10, nil)
	app := statedepot.FreshInsideRamPlatform()
	reply, err := app.Details(context.Background(), delegate.SolicitDetails)
	require.NoError(t, err)
	status.PlatformDigest = reply.FinalLedgerPlatformDigest
	cs := freshStatusUsingSettings(settings, status, privateItems[0], app)
	validateTxpool(cs.transferObserver).ActivateTransAccessible()
	altitude, iteration := cs.Altitude, cs.Iteration
	freshLedgerConduit := listen(cs.incidentPipeline, kinds.IncidentInquireFreshLedger)
	initiateVerifyIteration(cs, altitude, iteration)

	assureFreshIncidentUponConduit(freshLedgerConduit) //
	assureNegativeFreshIncidentUponConduit(freshLedgerConduit)
	dispatchTransScope(t, cs, 0, 1)
	assureFreshIncidentUponConduit(freshLedgerConduit) //
	assureFreshIncidentUponConduit(freshLedgerConduit) //
	assureNegativeFreshIncidentUponConduit(freshLedgerConduit)
}

func VerifyTxpoolOnwardSubsequentGenerateBlankLedgersDuration(t *testing.T) {
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginPath)

	settings.Agreement.GenerateVoidLedgersDuration = assureDeadline
	status, privateItems := arbitraryInaugurationStatus(1, false, 10, nil)
	app := statedepot.FreshInsideRamPlatform()
	reply, err := app.Details(context.Background(), delegate.SolicitDetails)
	require.NoError(t, err)
	status.PlatformDigest = reply.FinalLedgerPlatformDigest
	cs := freshStatusUsingSettings(settings, status, privateItems[0], app)

	validateTxpool(cs.transferObserver).ActivateTransAccessible()

	freshLedgerConduit := listen(cs.incidentPipeline, kinds.IncidentInquireFreshLedger)
	initiateVerifyIteration(cs, cs.Altitude, cs.Iteration)

	assureFreshIncidentUponConduit(freshLedgerConduit)   //
	assureNegativeFreshIncidentUponConduit(freshLedgerConduit) //
	assureFreshIncidentUponConduit(freshLedgerConduit)   //
}

func VerifyTxpoolOnwardInsideSuperiorIteration(t *testing.T) {
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	settings.Agreement.GenerateVoidLedgers = false
	status, privateItems := arbitraryInaugurationStatus(1, false, 10, nil)
	cs := freshStatusUsingSettings(settings, status, privateItems[0], statedepot.FreshInsideRamPlatform())
	validateTxpool(cs.transferObserver).ActivateTransAccessible()
	altitude, iteration := cs.Altitude, cs.Iteration
	freshLedgerConduit := listen(cs.incidentPipeline, kinds.IncidentInquireFreshLedger)
	freshIterationConduit := listen(cs.incidentPipeline, kinds.IncidentInquireFreshIteration)
	deadlineConduit := listen(cs.incidentPipeline, kinds.IncidentInquireDeadlineNominate)
	cs.assignNomination = func(nomination *kinds.Nomination) error {
		if cs.Altitude == 2 && cs.Iteration == 0 {
			//
			//
			cs.Tracer.Details("REDACTED")
			return nil
		}
		return cs.fallbackAssignNomination(nomination)
	}
	initiateVerifyIteration(cs, altitude, iteration)

	assureFreshIteration(freshIterationConduit, altitude, iteration) //
	assureFreshIncidentUponConduit(freshLedgerConduit)       //

	altitude++ //
	iteration = 0

	assureFreshIteration(freshIterationConduit, altitude, iteration) //
	dispatchTransScope(t, cs, 0, 1)              //
	assureFreshDeadline(deadlineConduit, altitude, iteration, cs.settings.DeadlineNominate.Nanoseconds())

	iteration++                                   //
	assureFreshIteration(freshIterationConduit, altitude, iteration) //
	assureFreshIncidentUponConduit(freshLedgerConduit)       //
}

func dispatchTransScope(t *testing.T, cs *Status, initiate, end int) {
	//
	for i := initiate; i < end; i++ {
		err := validateTxpool(cs.transferObserver).InspectTransfer(statedepot.FreshTransfer(fmt.Sprintf("REDACTED", i), "REDACTED"), nil, txpooll.TransferDetails{})
		require.NoError(t, err)
	}
}

func VerifyTxpoolTransferParallelUsingEndorse(t *testing.T) {
	status, privateItems := arbitraryInaugurationStatus(1, false, 10, nil)
	ledgerDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(ledgerDatastore, sm.DepotChoices{EjectIfaceReplies: false})
	cs := freshStatusUsingSettingsAlsoLedgerDepot(settings, status, privateItems[0], statedepot.FreshInsideRamPlatform(), ledgerDatastore)
	err := statusDepot.Persist(status)
	require.NoError(t, err)
	freshLedgerIncidentsConduit := listen(cs.incidentPipeline, kinds.IncidentInquireFreshLedgerIncidents)

	const countTrans int64 = 3000
	go dispatchTransScope(t, cs, 0, int(countTrans))

	initiateVerifyIteration(cs, cs.Altitude, cs.Iteration)
	for n := int64(0); n < countTrans; {
		select {
		case msg := <-freshLedgerIncidentsConduit:
			incident := msg.Data().(kinds.IncidentDataFreshLedgerIncidents)
			n += incident.CountTrans
			t.Log("REDACTED", "REDACTED", incident.CountTrans, "REDACTED", n)
		case <-time.After(30 * time.Second):
			t.Fatal("REDACTED")
		}
	}
}

func VerifyTxpoolDelFlawedTransfer(t *testing.T) {
	status, privateItems := arbitraryInaugurationStatus(1, false, 10, nil)
	app := statedepot.FreshInsideRamPlatform()
	ledgerDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(ledgerDatastore, sm.DepotChoices{EjectIfaceReplies: false})
	cs := freshStatusUsingSettingsAlsoLedgerDepot(settings, status, privateItems[0], app, ledgerDatastore)
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	//
	transferOctets := statedepot.FreshTransfer("REDACTED", "REDACTED")
	res, err := app.CulminateLedger(context.Background(), &iface.SolicitCulminateLedger{Txs: [][]byte{transferOctets}})
	require.NoError(t, err)
	assert.False(t, res.TransferOutcomes[0].EqualsFault())
	assert.True(t, len(res.PlatformDigest) > 0)

	_, err = app.Endorse(context.Background(), &iface.SolicitEndorse{})
	require.NoError(t, err)

	blankTxpoolConduit := make(chan struct{})
	inspectTransferReplyConduit := make(chan struct{})
	go func() {
		//
		//
		//
		unfitTransfer := []byte("REDACTED")
		err := validateTxpool(cs.transferObserver).InspectTransfer(unfitTransfer, func(r *iface.ReplyInspectTransfer) {
			if r.Cipher != statedepot.CipherKindUnfitTransferLayout {
				t.Errorf("REDACTED", r)
				return
			}
			inspectTransferReplyConduit <- struct{}{}
		}, txpooll.TransferDetails{})
		if err != nil {
			t.Errorf("REDACTED", err)
			return
		}

		//
		for {
			txs := validateTxpool(cs.transferObserver).HarvestMaximumOctetsMaximumFuel(int64(len(unfitTransfer)), -1)
			if len(txs) == 0 {
				blankTxpoolConduit <- struct{}{}
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	//
	metronome := time.After(time.Second * 5)
	select {
	case <-inspectTransferReplyConduit:
		//
	case <-metronome:
		t.Errorf("REDACTED")
		return
	}

	//
	metronome = time.After(time.Second * 5)
	select {
	case <-blankTxpoolConduit:
		//
	case <-metronome:
		t.Errorf("REDACTED")
		return
	}
}
