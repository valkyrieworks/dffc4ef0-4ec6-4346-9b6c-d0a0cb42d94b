package agreement

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

//
func validateTxpool(txn transferAlerter) txpool.Txpool {
	return txn.(txpool.Txpool)
}

func VerifyTxpoolNoAdvancementUntilTransAccessible(t *testing.T) {
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	settings.Agreement.GenerateEmptyLedgers = false
	status, privateValues := randomOriginStatus(1, false, 10, nil)
	app := objectdepot.NewInRamSoftware()
	reply, err := app.Details(context.Background(), gateway.QueryDetails)
	require.NoError(t, err)
	status.ApplicationDigest = reply.FinalLedgerApplicationDigest
	cs := newStatusWithSettings(settings, status, privateValues[0], app)
	validateTxpool(cs.transferAlerter).ActivateTransAccessible()
	level, epoch := cs.Level, cs.Cycle
	newLedgerChan := enrol(cs.eventBus, kinds.EventInquireNewLedger)
	beginVerifyEpoch(cs, level, epoch)

	assureNewEventOnStream(newLedgerChan) //
	assureNoNewEventOnStream(newLedgerChan)
	transferTransScope(t, cs, 0, 1)
	assureNewEventOnStream(newLedgerChan) //
	assureNewEventOnStream(newLedgerChan) //
	assureNoNewEventOnStream(newLedgerChan)
}

func VerifyTxpoolAdvancementAfterInstantiateEmptyLedgersCadence(t *testing.T) {
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)

	settings.Agreement.GenerateEmptyLedgersCadence = assureDeadline
	status, privateValues := randomOriginStatus(1, false, 10, nil)
	app := objectdepot.NewInRamSoftware()
	reply, err := app.Details(context.Background(), gateway.QueryDetails)
	require.NoError(t, err)
	status.ApplicationDigest = reply.FinalLedgerApplicationDigest
	cs := newStatusWithSettings(settings, status, privateValues[0], app)

	validateTxpool(cs.transferAlerter).ActivateTransAccessible()

	newLedgerChan := enrol(cs.eventBus, kinds.EventInquireNewLedger)
	beginVerifyEpoch(cs, cs.Level, cs.Cycle)

	assureNewEventOnStream(newLedgerChan)   //
	assureNoNewEventOnStream(newLedgerChan) //
	assureNewEventOnStream(newLedgerChan)   //
}

func VerifyTxpoolAdvancementInSuperiorEpoch(t *testing.T) {
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	settings.Agreement.GenerateEmptyLedgers = false
	status, privateValues := randomOriginStatus(1, false, 10, nil)
	cs := newStatusWithSettings(settings, status, privateValues[0], objectdepot.NewInRamSoftware())
	validateTxpool(cs.transferAlerter).ActivateTransAccessible()
	level, epoch := cs.Level, cs.Cycle
	newLedgerChan := enrol(cs.eventBus, kinds.EventInquireNewLedger)
	newEpochChan := enrol(cs.eventBus, kinds.EventInquireNewEpoch)
	deadlineChan := enrol(cs.eventBus, kinds.EventInquireDeadlineNominate)
	cs.assignNomination = func(nomination *kinds.Nomination) error {
		if cs.Level == 2 && cs.Cycle == 0 {
			//
			//
			cs.Tracer.Details("REDACTED")
			return nil
		}
		return cs.standardAssignNomination(nomination)
	}
	beginVerifyEpoch(cs, level, epoch)

	assureNewEpoch(newEpochChan, level, epoch) //
	assureNewEventOnStream(newLedgerChan)       //

	level++ //
	epoch = 0

	assureNewEpoch(newEpochChan, level, epoch) //
	transferTransScope(t, cs, 0, 1)              //
	assureNewDeadline(deadlineChan, level, epoch, cs.settings.DeadlineNominate.Nanoseconds())

	epoch++                                   //
	assureNewEpoch(newEpochChan, level, epoch) //
	assureNewEventOnStream(newLedgerChan)       //
}

func transferTransScope(t *testing.T, cs *Status, begin, end int) {
	//
	for i := begin; i < end; i++ {
		err := validateTxpool(cs.transferAlerter).InspectTransfer(objectdepot.NewTransfer(fmt.Sprintf("REDACTED", i), "REDACTED"), nil, txpool.TransferDetails{})
		require.NoError(t, err)
	}
}

func VerifyTxpoolTransferParallelWithEndorse(t *testing.T) {
	status, privateValues := randomOriginStatus(1, false, 10, nil)
	ledgerStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(ledgerStore, sm.DepotSettings{DropIfaceReplies: false})
	cs := newStatusWithSettingsAndLedgerDepot(settings, status, privateValues[0], objectdepot.NewInRamSoftware(), ledgerStore)
	err := statusDepot.Persist(status)
	require.NoError(t, err)
	newLedgerEventsChan := enrol(cs.eventBus, kinds.EventInquireNewLedgerEvents)

	const countTrans int64 = 3000
	go transferTransScope(t, cs, 0, int(countTrans))

	beginVerifyEpoch(cs, cs.Level, cs.Cycle)
	for n := int64(0); n < countTrans; {
		select {
		case msg := <-newLedgerEventsChan:
			event := msg.Data().(kinds.EventDataNewLedgerEvents)
			n += event.CountTrans
			t.Log("REDACTED", "REDACTED", event.CountTrans, "REDACTED", n)
		case <-time.After(30 * time.Second):
			t.Fatal("REDACTED")
		}
	}
}

func VerifyTxpoolRemoveFlawedTransfer(t *testing.T) {
	status, privateValues := randomOriginStatus(1, false, 10, nil)
	app := objectdepot.NewInRamSoftware()
	ledgerStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(ledgerStore, sm.DepotSettings{DropIfaceReplies: false})
	cs := newStatusWithSettingsAndLedgerDepot(settings, status, privateValues[0], app, ledgerStore)
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	//
	transferOctets := objectdepot.NewTransfer("REDACTED", "REDACTED")
	res, err := app.CompleteLedger(context.Background(), &iface.QueryCompleteLedger{Txs: [][]byte{transferOctets}})
	require.NoError(t, err)
	assert.False(t, res.TransOutcomes[0].IsErr())
	assert.True(t, len(res.ApplicationDigest) > 0)

	_, err = app.Endorse(context.Background(), &iface.QueryEndorse{})
	require.NoError(t, err)

	emptyTxpoolChan := make(chan struct{})
	inspectTransferReplyChan := make(chan struct{})
	go func() {
		//
		//
		//
		corruptTransfer := []byte("REDACTED")
		err := validateTxpool(cs.transferAlerter).InspectTransfer(corruptTransfer, func(r *iface.AnswerInspectTransfer) {
			if r.Code != objectdepot.CodeKindCorruptTransferLayout {
				t.Errorf("REDACTED", r)
				return
			}
			inspectTransferReplyChan <- struct{}{}
		}, txpool.TransferDetails{})
		if err != nil {
			t.Errorf("REDACTED", err)
			return
		}

		//
		for {
			txs := validateTxpool(cs.transferAlerter).HarvestMaximumOctetsMaximumFuel(int64(len(corruptTransfer)), -1)
			if len(txs) == 0 {
				emptyTxpoolChan <- struct{}{}
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	//
	timer := time.After(time.Second * 5)
	select {
	case <-inspectTransferReplyChan:
		//
	case <-timer:
		t.Errorf("REDACTED")
		return
	}

	//
	timer = time.After(time.Second * 5)
	select {
	case <-emptyTxpoolChan:
		//
	case <-timer:
		t.Errorf("REDACTED")
		return
	}
}
