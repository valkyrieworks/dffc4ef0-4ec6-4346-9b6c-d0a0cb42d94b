package agreement

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds/simulations"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	machinestubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyPrimary(m *testing.M) {
	settings = RestoreSettings("REDACTED")
	agreementReenactSettings = RestoreSettings("REDACTED")
	settingsStatusVerify := RestoreSettings("REDACTED")
	settingsTxpoolVerify := RestoreSettings("REDACTED")
	settingsTreacherousVerify := RestoreSettings("REDACTED")
	cipher := m.Run()
	os.RemoveAll(settings.OriginPath)
	os.RemoveAll(agreementReenactSettings.OriginPath)
	os.RemoveAll(settingsStatusVerify.OriginPath)
	os.RemoveAll(settingsTxpoolVerify.OriginPath)
	os.RemoveAll(settingsTreacherousVerify.OriginPath)
	os.Exit(cipher)
}

//
//
//
//
//
//
//

//
//

//
//
//

func initiateFreshStatusAlsoPauseForeachLedger(
	t *testing.T,
	agreementReenactSettings *cfg.Settings,
	ledgerDatastore dbm.DB,
	statusDepot sm.Depot,
) {
	tracer := log.VerifyingTracer()
	status, _ := statusDepot.FetchOriginatingDatastoreEitherInaugurationRecord(agreementReenactSettings.InaugurationRecord())
	privateAssessor := fetchPrivateAssessor(agreementReenactSettings)
	cs := freshStatusUsingSettingsAlsoLedgerDepot(
		agreementReenactSettings,
		status,
		privateAssessor,
		statedepot.FreshInsideRamPlatform(),
		ledgerDatastore,
	)
	cs.AssignTracer(tracer)

	octets, _ := os.ReadFile(cs.settings.JournalRecord())
	t.Logf("REDACTED", octets)

	err := cs.Initiate()
	require.NoError(t, err)
	defer func() {
		if err := cs.Halt(); err != nil {
			t.Error(err)
		}
	}()

	//
	//
	//
	//
	freshLedgerUnder, err := cs.incidentChannel.Listen(context.Background(), verifyListener, kinds.IncidentInquireFreshLedger)
	require.NoError(t, err)
	select {
	case <-freshLedgerUnder.Out():
	case <-freshLedgerUnder.Aborted():
		t.Fatal("REDACTED")
	case <-time.After(120 * time.Second):
		t.Fatal("REDACTED")
	}
}

func transmitTrans(ctx context.Context, cs *Status) {
	for i := 0; i < 256; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			tx := statedepot.FreshTransferOriginatingUUID(i)
			if err := attestTxpool(cs.transferObserver).InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				if reply.Cipher != 0 {
					panic(fmt.Sprintf("REDACTED", reply.Cipher, reply.Log))
				}
			}, txpool.TransferDetails{}); err != nil {
				panic(err)
			}
			i++
		}
	}
}

//
func VerifyJournalCollapse(t *testing.T) {
	verifyScenarios := []struct {
		alias         string
		initializeProc       func(dbm.DB, *Status, context.Context)
		altitudeTowardHalt int64
	}{
		{
			"REDACTED",
			func(statusDatastore dbm.DB, cs *Status, ctx context.Context) {},
			1,
		},
		{
			"REDACTED",
			func(statusDatastore dbm.DB, cs *Status, ctx context.Context) {
				go transmitTrans(ctx, cs)
			},
			3,
		},
	}

	for i, tc := range verifyScenarios {
		agreementReenactSettings := RestoreSettings(fmt.Sprintf("REDACTED", t.Name(), i))
		t.Run(tc.alias, func(t *testing.T) {
			collapseWaitableArriveInspectVitality(t, agreementReenactSettings, tc.initializeProc, tc.altitudeTowardHalt)
		})
	}
}

func collapseWaitableArriveInspectVitality(t *testing.T, agreementReenactSettings *cfg.Settings,
	initializeProc func(dbm.DB, *Status, context.Context), altitudeTowardHalt int64,
) {
	journalAlarmed := make(chan error)
	collapsingJournal := &collapsingJournal{alarmChnl: journalAlarmed, altitudeTowardHalt: altitudeTowardHalt}

	i := 1
Cycle:
	for {
		t.Logf("REDACTED", i)

		//
		tracer := log.FreshNooperationTracer()
		ledgerDatastore := dbm.FreshMemoryDatastore()
		statusDatastore := ledgerDatastore
		statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
			EjectIfaceReplies: false,
		})
		status, err := sm.CreateInaugurationStatusOriginatingRecord(agreementReenactSettings.InaugurationRecord())
		require.NoError(t, err)
		privateAssessor := fetchPrivateAssessor(agreementReenactSettings)
		cs := freshStatusUsingSettingsAlsoLedgerDepot(
			agreementReenactSettings,
			status,
			privateAssessor,
			statedepot.FreshInsideRamPlatform(),
			ledgerDatastore,
		)
		cs.AssignTracer(tracer)

		//
		ctx, abort := context.WithCancel(context.Background())
		initializeProc(statusDatastore, cs, ctx)

		//
		journalRecord := cs.settings.JournalRecord()
		os.Remove(journalRecord)

		//
		controlJournal, err := cs.UnsealJournal(journalRecord)
		require.NoError(t, err)
		collapsingJournal.following = controlJournal

		//
		collapsingJournal.signalOrdinal = 1
		cs.wal = collapsingJournal

		//
		err = cs.Initiate()
		require.NoError(t, err)

		i++

		select {
		case err := <-journalAlarmed:
			t.Logf("REDACTED", err)

			//
			initiateFreshStatusAlsoPauseForeachLedger(t, agreementReenactSettings, ledgerDatastore, statusDepot)

			//
			cs.Halt() //
			abort()

			//
			if _, ok := err.(AttainedAltitudeTowardHaltFailure); ok {
				break Cycle
			}
		case <-time.After(10 * time.Second):
			t.Fatal("REDACTED")
		}
	}
}

//
//
//
type collapsingJournal struct {
	following         WAL
	alarmChnl      chan error
	altitudeTowardHalt int64

	signalOrdinal                int //
	finalAlarmedForeachSignalOrdinal int //
}

var _ WAL = &collapsingJournal{}

//
type JournalPersistFailure struct {
	msg string
}

func (e JournalPersistFailure) Failure() string {
	return e.msg
}

//
//
type AttainedAltitudeTowardHaltFailure struct {
	altitude int64
}

func (e AttainedAltitudeTowardHaltFailure) Failure() string {
	return fmt.Sprintf("REDACTED", e.altitude)
}

//
//
func (w *collapsingJournal) Persist(m JournalSignal) error {
	if terminateSignal, ok := m.(TerminateAltitudeSignal); ok {
		if terminateSignal.Altitude == w.altitudeTowardHalt {
			w.alarmChnl <- AttainedAltitudeTowardHaltFailure{terminateSignal.Altitude}
			runtime.Goexit()
			return nil
		}

		return w.following.Persist(m)
	}

	if w.signalOrdinal > w.finalAlarmedForeachSignalOrdinal {
		w.finalAlarmedForeachSignalOrdinal = w.signalOrdinal
		_, record, row, _ := runtime.Caller(1)
		w.alarmChnl <- JournalPersistFailure{fmt.Sprintf("REDACTED", m, record, row)}
		runtime.Goexit()
		return nil
	}

	w.signalOrdinal++
	return w.following.Persist(m)
}

func (w *collapsingJournal) PersistChronize(m JournalSignal) error {
	return w.Persist(m)
}

func (w *collapsingJournal) PurgeAlsoChronize() error { return w.following.PurgeAlsoChronize() }

func (w *collapsingJournal) LookupForeachTerminateAltitude(
	altitude int64,
	choices *JournalLookupChoices,
) (rd io.ReadCloser, detected bool, err error) {
	return w.following.LookupForeachTerminateAltitude(altitude, choices)
}

func (w *collapsingJournal) Initiate() error { return w.following.Initiate() }
func (w *collapsingJournal) Halt() error  { return w.following.Halt() }
func (w *collapsingJournal) Await()        { w.following.Await() }

//

const countLedgers = 6

//
//

//
//
//
//
var styles = []uint{0, 1, 2, 3}

//
func configureSuccessionUsingModifyingAssessors(t *testing.T, alias string, nthLedgers int) (*cfg.Settings, []*kinds.Ledger, []*kinds.ExpandedEndorse, sm.Status) {
	ctx, abort := context.WithCancel(context.Background())
	defer abort()

	nthNodes := 7
	nthValues := 4
	css, producePaper, settings, sanitize := arbitraryAgreementNetworkUsingNodes(
		t,
		nthValues,
		nthNodes,
		alias,
		freshSimulateMetronomeMethod(true),
		func(_ string) iface.Platform {
			return freshTokvalDepot()
		})
	inaugurationStatus, err := sm.CreateInaugurationStatus(producePaper)
	require.NoError(t, err)
	t.Cleanup(sanitize)

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	freshIterationChnl := listen(css[0].incidentChannel, kinds.IncidentInquireFreshIteration)
	nominationChnl := listen(css[0].incidentChannel, kinds.IncidentInquireFinishNomination)

	vss := make([]*assessorMock, nthNodes)
	for i := 0; i < nthNodes; i++ {
		vss[i] = freshAssessorMock(css[i].privateAssessor, int32(i))
	}
	altitude, iteration := css[0].Altitude, css[0].Iteration

	//
	initiateVerifyIteration(css[0], altitude, iteration)
	advanceAltitude(vss...)
	assureFreshIteration(freshIterationChnl, altitude, 0)
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := css[0].ObtainIterationStatus()
	attestAppendBallots(css[0], commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), true, vss[1:nthValues]...)
	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	//
	altitude++
	advanceAltitude(vss...)
	freshAssessorPublicToken1, err := css[nthValues].privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	itemPublicToken1iface, err := cryptocode.PublicTokenTowardSchema(freshAssessorPublicToken1)
	require.NoError(t, err)
	freshAssessorTransfer1 := statedepot.CreateItemAssignModifyTransfer(itemPublicToken1iface, verifyMinimumPotency)
	err = attestTxpool(css[0].transferObserver).InspectTransfer(freshAssessorTransfer1, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	itemLedger, err := css[0].generateNominationLedger(ctx) //
	require.NoError(t, err)
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: itemLedger.Digest(), FragmentAssignHeading: itemLedgerFragments.Heading()}

	nomination := kinds.FreshNomination(vss[1].Altitude, iteration, -1, ledgerUUID)
	p := nomination.TowardSchema()
	if err := vss[1].AttestNomination(verify.FallbackVerifySuccessionUUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Notation = p.Notation

	//
	if err := css[0].AssignNominationAlsoLedger(nomination, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs = css[0].ObtainIterationStatus()
	attestAppendBallots(css[0], commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), true, vss[1:nthValues]...)
	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	//
	altitude++
	advanceAltitude(vss...)
	reviseAssessorPublicToken1, err := css[nthValues].privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	revisePublicToken1iface, err := cryptocode.PublicTokenTowardSchema(reviseAssessorPublicToken1)
	require.NoError(t, err)
	reviseAssessorTransfer1 := statedepot.CreateItemAssignModifyTransfer(revisePublicToken1iface, 25)
	err = attestTxpool(css[0].transferObserver).InspectTransfer(reviseAssessorTransfer1, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	itemLedger, err = css[0].generateNominationLedger(ctx) //
	require.NoError(t, err)
	itemLedgerFragments, err = itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	ledgerUUID = kinds.LedgerUUID{Digest: itemLedger.Digest(), FragmentAssignHeading: itemLedgerFragments.Heading()}

	nomination = kinds.FreshNomination(vss[2].Altitude, iteration, -1, ledgerUUID)
	p = nomination.TowardSchema()
	if err := vss[2].AttestNomination(verify.FallbackVerifySuccessionUUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Notation = p.Notation

	//
	if err := css[0].AssignNominationAlsoLedger(nomination, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs = css[0].ObtainIterationStatus()
	attestAppendBallots(css[0], commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), true, vss[1:nthValues]...)
	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	//
	altitude++
	advanceAltitude(vss...)
	freshAssessorPublicToken2, err := css[nthValues+1].privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	freshAssessor2iface, err := cryptocode.PublicTokenTowardSchema(freshAssessorPublicToken2)
	require.NoError(t, err)
	freshAssessorTransfer2 := statedepot.CreateItemAssignModifyTransfer(freshAssessor2iface, verifyMinimumPotency)
	err = attestTxpool(css[0].transferObserver).InspectTransfer(freshAssessorTransfer2, nil, txpool.TransferDetails{})
	require.NoError(t, err)
	freshAssessorPublicToken3, err := css[nthValues+2].privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	freshAssessor3iface, err := cryptocode.PublicTokenTowardSchema(freshAssessorPublicToken3)
	require.NoError(t, err)
	freshAssessorTransfer3 := statedepot.CreateItemAssignModifyTransfer(freshAssessor3iface, verifyMinimumPotency)
	err = attestTxpool(css[0].transferObserver).InspectTransfer(freshAssessorTransfer3, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	itemLedger, err = css[0].generateNominationLedger(ctx) //
	require.NoError(t, err)
	itemLedgerFragments, err = itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	ledgerUUID = kinds.LedgerUUID{Digest: itemLedger.Digest(), FragmentAssignHeading: itemLedgerFragments.Heading()}
	freshVsscheme := make([]*assessorMock, nthValues+1)
	copy(freshVsscheme, vss[:nthValues+1])
	sort.Sort(AssessorMocksViaPotency(freshVsscheme))

	itemOrdinalProc := func(styleOffset int) int {
		for i, vs := range freshVsscheme {
			versusPublicToken, err := vs.ObtainPublicToken()
			require.NoError(t, err)

			stylePublicToken, err := css[styleOffset].privateAssessor.ObtainPublicToken()
			require.NoError(t, err)

			if versusPublicToken.Matches(stylePublicToken) {
				return i
			}
		}
		panic(fmt.Sprintf("REDACTED", styleOffset))
	}

	egoOrdinal := itemOrdinalProc(0)

	nomination = kinds.FreshNomination(vss[3].Altitude, iteration, -1, ledgerUUID)
	p = nomination.TowardSchema()
	if err := vss[3].AttestNomination(verify.FallbackVerifySuccessionUUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Notation = p.Notation

	//
	if err := css[0].AssignNominationAlsoLedger(nomination, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureFreshNomination(nominationChnl, altitude, iteration)

	discardAssessorTransfer2 := statedepot.CreateItemAssignModifyTransfer(freshAssessor2iface, 0)
	err = attestTxpool(css[0].transferObserver).InspectTransfer(discardAssessorTransfer2, nil, txpool.TransferDetails{})
	assert.Nil(t, err)

	rs = css[0].ObtainIterationStatus()
	for i := 0; i < nthValues+1; i++ {
		if i == egoOrdinal {
			continue
		}
		attestAppendBallots(css[0], commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), true, freshVsscheme[i])
	}

	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	//
	altitude++
	advanceAltitude(vss...)
	//
	freshVsschemeOffset := itemOrdinalProc(nthValues)
	freshVsscheme[freshVsschemeOffset].BallotingPotency = 25
	sort.Sort(AssessorMocksViaPotency(freshVsscheme))
	egoOrdinal = itemOrdinalProc(0)
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs = css[0].ObtainIterationStatus()
	for i := 0; i < nthValues+1; i++ {
		if i == egoOrdinal {
			continue
		}
		attestAppendBallots(css[0], commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), true, freshVsscheme[i])
	}
	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	//
	altitude++
	advanceAltitude(vss...)
	discardAssessorTransfer3 := statedepot.CreateItemAssignModifyTransfer(freshAssessor3iface, 0)
	err = attestTxpool(css[0].transferObserver).InspectTransfer(discardAssessorTransfer3, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	itemLedger, err = css[0].generateNominationLedger(ctx) //
	require.NoError(t, err)
	itemLedgerFragments, err = itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	ledgerUUID = kinds.LedgerUUID{Digest: itemLedger.Digest(), FragmentAssignHeading: itemLedgerFragments.Heading()}
	freshVsscheme = make([]*assessorMock, nthValues+3)
	copy(freshVsscheme, vss[:nthValues+3])
	sort.Sort(AssessorMocksViaPotency(freshVsscheme))

	egoOrdinal = itemOrdinalProc(0)
	nomination = kinds.FreshNomination(vss[1].Altitude, iteration, -1, ledgerUUID)
	p = nomination.TowardSchema()
	if err := vss[1].AttestNomination(verify.FallbackVerifySuccessionUUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Notation = p.Notation

	//
	if err := css[0].AssignNominationAlsoLedger(nomination, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs = css[0].ObtainIterationStatus()
	for i := 0; i < nthValues+3; i++ {
		if i == egoOrdinal {
			continue
		}
		attestAppendBallots(css[0], commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), true, freshVsscheme[i])
	}
	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	succession := []*kinds.Ledger{}
	addnEndorses := []*kinds.ExpandedEndorse{}
	for i := 1; i <= nthLedgers; i++ {
		succession = append(succession, css[0].ledgerDepot.FetchLedger(int64(i)))
		addnEndorses = append(addnEndorses, css[0].ledgerDepot.FetchLedgerExpandedEndorse(int64(i)))
	}
	return settings, succession, addnEndorses, inaugurationStatus
}

//
func VerifyNegotiationReenactEvery(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, 0, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, 0, m, false)
		})
	}
}

//
func VerifyNegotiationReenactFew(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, 2, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, 2, m, true)
		})
	}
}

//
func VerifyNegotiationReenactSingle(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, countLedgers-1, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, countLedgers-1, m, true)
		})
	}
}

//
func VerifyNegotiationReenactNull(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, countLedgers, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyNegotiationReenact(t, settings, countLedgers, m, true)
		})
	}
}

func transientJournalUsingData(data []byte) string {
	journalRecord, err := os.CreateTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	_, err = journalRecord.Write(data)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if err := journalRecord.Close(); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	return journalRecord.Name()
}

//
//
func verifyNegotiationReenact(t *testing.T, settings *cfg.Settings, nthLedgers int, style uint, verifyAssessorsModify bool) {
	var (
		verifySettings   *cfg.Settings
		succession        []*kinds.Ledger
		addnEndorses   []*kinds.ExpandedEndorse
		depot        *simulateLedgerDepot
		statusDatastore      dbm.DB
		inaugurationStatus sm.Status
		txpool      = blankTxpool{}
		incidentpool       = sm.VoidProofHub{}
	)

	if verifyAssessorsModify {
		verifySettings, succession, addnEndorses, inaugurationStatus = configureSuccessionUsingModifyingAssessors(t, fmt.Sprintf("REDACTED", nthLedgers, style), countLedgers)
		statusDatastore = dbm.FreshMemoryDatastore()
		depot = freshSimulateLedgerDepot(t, settings, inaugurationStatus.AgreementSettings)
	} else {
		verifySettings = RestoreSettings(fmt.Sprintf("REDACTED", nthLedgers, style))
		t.Cleanup(func() {
			_ = os.RemoveAll(verifySettings.OriginPath)
		})
		journalContent, err := JournalUsingNTHLedgers(t, countLedgers, verifySettings)
		require.NoError(t, err)
		journalRecord := transientJournalUsingData(journalContent)
		verifySettings.Agreement.AssignJournalRecord(journalRecord)

		wal, err := FreshJournal(journalRecord)
		require.NoError(t, err)
		wal.AssignTracer(log.VerifyingTracer())
		err = wal.Initiate()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := wal.Halt(); err != nil {
				t.Error(err)
			}
		})
		succession, addnEndorses, err = createLedgerchainOriginatingJournal(wal)
		require.NoError(t, err)
		statusDatastore, inaugurationStatus, depot = statusAlsoDepot(t, verifySettings, statedepot.PlatformEdition)
	}

	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	t.Cleanup(func() {
		_ = statusDepot.Shutdown()
	})
	depot.succession = succession
	depot.addnEndorses = addnEndorses

	status := inaugurationStatus.Duplicate()
	//
	status, newestApplicationDigest := constructTEMPStatusOriginatingSuccession(t, verifySettings, statusDepot, txpool, incidentpool, status, succession, nthLedgers, style, depot)

	//
	statedepotApplication := statedepot.FreshEnduringPlatform(
		filepath.Join(verifySettings.DatastorePath(), fmt.Sprintf("REDACTED", nthLedgers, style)))
	t.Cleanup(func() {
		_ = statedepotApplication.Shutdown()
	})

	customerOriginator2 := delegate.FreshRegionalCustomerOriginator(statedepotApplication)
	if nthLedgers > 0 {
		//
		//
		delegatePlatform := delegate.FreshPlatformLinks(customerOriginator2, delegate.NooperationTelemetry())
		statusDepot1 := dbm.FreshMemoryDatastore()
		placeholderStatusDepot := sm.FreshDepot(statusDepot1, sm.DepotChoices{
			EjectIfaceReplies: false,
		})
		err := placeholderStatusDepot.Persist(inaugurationStatus)
		require.NoError(t, err)
		constructApplicationStatusOriginatingSuccession(t, delegatePlatform, placeholderStatusDepot, txpool, incidentpool, inaugurationStatus, succession, nthLedgers, style, depot)
	}

	//
	anticipateFailure := false
	if style == 3 {
		trimmed, _, err := depot.TrimLedgers(2, status)
		require.NoError(t, err)
		require.EqualValues(t, 1, trimmed)
		anticipateFailure = int64(nthLedgers) < 2
	}

	//
	producePaper, err := sm.CreateInaugurationPaperOriginatingRecord(verifySettings.InaugurationRecord())
	require.NoError(t, err)
	negotiator := FreshNegotiator(statusDepot, status, depot, producePaper)
	delegatePlatform := delegate.FreshPlatformLinks(customerOriginator2, delegate.NooperationTelemetry())
	if err := delegatePlatform.Initiate(); err != nil {
		t.Fatalf("REDACTED", err)
	}

	t.Cleanup(func() {
		if err := delegatePlatform.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	err = negotiator.Negotiation(delegatePlatform)
	if anticipateFailure {
		require.Error(t, err)
		//
		return
	}
	require.NoError(t, err)

	//
	res, err := delegatePlatform.Inquire().Details(context.Background(), delegate.SolicitDetails)
	require.NoError(t, err)

	//
	require.Equal(t, depot.Altitude(), res.FinalLedgerAltitude)

	//
	status, err = statusDepot.Fetch()
	require.NoError(t, err)
	require.Equal(t, status.FinalLedgerAltitude, res.FinalLedgerAltitude)
	require.Equal(t, int64(countLedgers), res.FinalLedgerAltitude)

	//
	if !bytes.Equal(newestApplicationDigest, res.FinalLedgerPlatformDigest) {
		t.Fatalf(
			"REDACTED",
			res.FinalLedgerPlatformDigest,
			newestApplicationDigest)
	}

	anticipatedLedgersTowardChronize := countLedgers - nthLedgers
	if nthLedgers == countLedgers && style > 0 {
		anticipatedLedgersTowardChronize++
	} else if nthLedgers > 0 && style == 1 {
		anticipatedLedgersTowardChronize++
	}

	if negotiator.NTHLedgers() != anticipatedLedgersTowardChronize {
		t.Fatalf("REDACTED", anticipatedLedgersTowardChronize, negotiator.NTHLedgers())
	}
}

func executeLedger(t *testing.T, statusDepot sm.Depot, txpool txpool.Txpool, incidentpool sm.ProofHub, st sm.Status, blk *kinds.Ledger, delegatePlatform delegate.PlatformLinks, bs sm.LedgerDepot) sm.Status {
	verifyFragmentExtent := kinds.LedgerFragmentExtentOctets
	ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegatePlatform.Agreement(), txpool, incidentpool, bs)

	bps, err := blk.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)
	ldgUUID := kinds.LedgerUUID{Digest: blk.Digest(), FragmentAssignHeading: bps.Heading()}
	freshStatus, err := ledgerExecute.ExecuteLedger(st, ldgUUID, blk)
	require.NoError(t, err)
	return freshStatus
}

func constructApplicationStatusOriginatingSuccession(t *testing.T, delegatePlatform delegate.PlatformLinks, statusDepot sm.Depot, txpool txpool.Txpool, incidentpool sm.ProofHub,
	status sm.Status, succession []*kinds.Ledger, nthLedgers int, style uint, bs sm.LedgerDepot,
) {
	//
	if err := delegatePlatform.Initiate(); err != nil {
		panic(err)
	}
	defer delegatePlatform.Halt() //

	status.Edition.Agreement.App = statedepot.PlatformEdition //
	assessors := kinds.Temp2buffer.AssessorRevisions(status.Assessors)
	if _, err := delegatePlatform.Agreement().InitializeSuccession(context.Background(), &iface.SolicitInitializeSuccession{
		Assessors: assessors,
	}); err != nil {
		panic(err)
	}
	if err := statusDepot.Persist(status); err != nil { //
		panic(err)
	}
	switch style {
	case 0:
		for i := 0; i < nthLedgers; i++ {
			ledger := succession[i]
			status = executeLedger(t, statusDepot, txpool, incidentpool, status, ledger, delegatePlatform, bs)
		}
	case 1, 2, 3:
		for i := 0; i < nthLedgers-1; i++ {
			ledger := succession[i]
			status = executeLedger(t, statusDepot, txpool, incidentpool, status, ledger, delegatePlatform, bs)
		}

		//
		//
		if style == 2 || style == 3 {
			//
			//
			//
			status = executeLedger(t, statusDepot, txpool, incidentpool, status, succession[nthLedgers-1], delegatePlatform, bs)
		}
	default:
		panic(fmt.Sprintf("REDACTED", style))
	}
}

func constructTEMPStatusOriginatingSuccession(
	t *testing.T,
	settings *cfg.Settings,
	statusDepot sm.Depot,
	txpool txpool.Txpool,
	incidentpool sm.ProofHub,
	status sm.Status,
	succession []*kinds.Ledger,
	nthLedgers int,
	style uint,
	bs sm.LedgerDepot,
) (sm.Status, []byte) {
	//
	customerOriginator := delegate.FreshRegionalCustomerOriginator(
		statedepot.FreshEnduringPlatform(
			filepath.Join(settings.DatastorePath(), fmt.Sprintf("REDACTED", nthLedgers, style))))
	delegatePlatform := delegate.FreshPlatformLinks(customerOriginator, delegate.NooperationTelemetry())
	if err := delegatePlatform.Initiate(); err != nil {
		panic(err)
	}
	defer delegatePlatform.Halt() //

	status.Edition.Agreement.App = statedepot.PlatformEdition //
	assessors := kinds.Temp2buffer.AssessorRevisions(status.Assessors)
	if _, err := delegatePlatform.Agreement().InitializeSuccession(context.Background(), &iface.SolicitInitializeSuccession{
		Assessors: assessors,
	}); err != nil {
		panic(err)
	}
	if err := statusDepot.Persist(status); err != nil { //
		panic(err)
	}
	switch style {
	case 0:
		//
		for _, ledger := range succession {
			status = executeLedger(t, statusDepot, txpool, incidentpool, status, ledger, delegatePlatform, bs)
		}
		return status, status.PlatformDigest

	case 1, 2, 3:
		//
		//
		for _, ledger := range succession[:len(succession)-1] {
			status = executeLedger(t, statusDepot, txpool, incidentpool, status, ledger, delegatePlatform, bs)
		}

		placeholderStatusDepot := &machinestubs.Depot{}
		finalAltitude := int64(len(succession))
		nexttofinalAltitude := int64(len(succession) - 1)
		values, _ := statusDepot.FetchAssessors(nexttofinalAltitude)
		placeholderStatusDepot.On("REDACTED", nexttofinalAltitude).Return(values, nil)
		placeholderStatusDepot.On("REDACTED", mock.Anything).Return(nil)
		placeholderStatusDepot.On("REDACTED", finalAltitude, mock.MatchedBy(func(reply *iface.ReplyCulminateLedger) bool {
			require.NoError(t, statusDepot.PersistCulminateLedgerReply(finalAltitude, reply))
			return true
		})).Return(nil)

		//
		//
		s := executeLedger(t, placeholderStatusDepot, txpool, incidentpool, status, succession[len(succession)-1], delegatePlatform, bs)
		return status, s.PlatformDigest
	default:
		panic(fmt.Sprintf("REDACTED", style))
	}
}

func createLedgers(n int, status sm.Status, privateItems []kinds.PrivateAssessor) ([]*kinds.Ledger, error) {
	ledgerUUID := verify.CreateLedgerUUID()
	ledgers := make([]*kinds.Ledger, n)

	for i := 0; i < n; i++ {
		altitude := status.FinalLedgerAltitude + 1 + int64(i)
		finalEndorse, err := verify.CreateEndorse(ledgerUUID, altitude-1, 0, status.FinalAssessors, privateItems, status.SuccessionUUID, status.FinalLedgerMoment)
		if err != nil {
			return nil, err
		}
		ledger, err := status.CreateLedger(altitude, verify.CreateNTHTrans(altitude, 10), finalEndorse, nil, status.FinalAssessors.Nominator.Location)
		if err != nil {
			return nil, err
		}
		ledgers[i] = ledger
		status.FinalLedgerUUID = ledgerUUID
		status.FinalLedgerAltitude = altitude
		status.FinalLedgerMoment = status.FinalLedgerMoment.Add(1 * time.Second)
		status.FinalAssessors = status.Assessors.Duplicate()
		status.Assessors = status.FollowingAssessors.Duplicate()
		status.FollowingAssessors = status.FollowingAssessors.DuplicateAdvanceNominatorUrgency(1)
		status.PlatformDigest = verify.ArbitraryDigest()

		ledgerUUID = verify.CreateLedgerUUIDUsingDigest(ledger.Digest())
	}

	return ledgers, nil
}

func VerifyNegotiationAlarmsConditionApplicationYieldsIncorrectApplicationDigest(t *testing.T) {
	//
	//
	//
	//
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	privateItem := privatevalue.FetchRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord())
	const platformEdition = 0x0
	statusDatastore, status, depot := statusAlsoDepot(t, settings, platformEdition)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	producePaper, _ := sm.CreateInaugurationPaperOriginatingRecord(settings.InaugurationRecord())
	status.FinalAssessors = status.Assessors.Duplicate()
	//
	ledgers, err := createLedgers(3, status, []kinds.PrivateAssessor{privateItem})
	require.NoError(t, err)

	depot.succession = ledgers

	//
	//
	//
	//
	{
		app := &flawedApplication{countLedgers: 3, everyDigestsExistIncorrect: true}
		customerOriginator := delegate.FreshRegionalCustomerOriginator(app)
		delegatePlatform := delegate.FreshPlatformLinks(customerOriginator, delegate.NooperationTelemetry())
		err := delegatePlatform.Initiate()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := delegatePlatform.Halt(); err != nil {
				t.Error(err)
			}
		})

		assert.Panics(t, func() {
			h := FreshNegotiator(statusDepot, status, depot, producePaper)
			if err = h.Negotiation(delegatePlatform); err != nil {
				t.Log(err)
			}
		})
	}

	//
	//
	//
	//
	{
		app := &flawedApplication{countLedgers: 3, solelyFinalDigestEqualsIncorrect: true}
		customerOriginator := delegate.FreshRegionalCustomerOriginator(app)
		delegatePlatform := delegate.FreshPlatformLinks(customerOriginator, delegate.NooperationTelemetry())
		err := delegatePlatform.Initiate()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := delegatePlatform.Halt(); err != nil {
				t.Error(err)
			}
		})

		assert.Panics(t, func() {
			h := FreshNegotiator(statusDepot, status, depot, producePaper)
			if err = h.Negotiation(delegatePlatform); err != nil {
				t.Log(err)
			}
		})
	}
}

type flawedApplication struct {
	iface.FoundationPlatform
	countLedgers           byte
	altitude              byte
	everyDigestsExistIncorrect   bool
	solelyFinalDigestEqualsIncorrect bool
}

func (app *flawedApplication) CulminateLedger(context.Context, *iface.SolicitCulminateLedger) (*iface.ReplyCulminateLedger, error) {
	app.altitude++
	if app.solelyFinalDigestEqualsIncorrect {
		if app.altitude == app.countLedgers {
			return &iface.ReplyCulminateLedger{PlatformDigest: commitrand.Octets(8)}, nil
		}
		return &iface.ReplyCulminateLedger{PlatformDigest: []byte{app.altitude}}, nil
	} else if app.everyDigestsExistIncorrect {
		return &iface.ReplyCulminateLedger{PlatformDigest: commitrand.Octets(8)}, nil
	}

	panic("REDACTED")
}

//
//

func createLedgerchainOriginatingJournal(wal WAL) ([]*kinds.Ledger, []*kinds.ExpandedEndorse, error) {
	var altitude int64

	//
	gr, detected, err := wal.LookupForeachTerminateAltitude(altitude, &JournalLookupChoices{})
	if err != nil {
		return nil, nil, err
	}
	if !detected {
		return nil, nil, fmt.Errorf("REDACTED", altitude)
	}
	defer gr.Close()

	//

	var (
		ledgers             []*kinds.Ledger
		addnEndorses         []*kinds.ExpandedEndorse
		thatLedgerFragments     *kinds.FragmentAssign
		thatLedgerAddnEndorse *kinds.ExpandedEndorse
	)

	dec := FreshJournalDeserializer(gr)
	for {
		msg, err := dec.Deserialize()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}

		portion := retrievePortionOriginatingJournal(msg)
		if portion == nil {
			continue
		}

		switch p := portion.(type) {
		case TerminateAltitudeSignal:
			//
			if thatLedgerFragments != nil {
				pbb := new(commitchema.Ledger)
				bz, err := io.ReadAll(thatLedgerFragments.ObtainFetcher())
				if err != nil {
					panic(err)
				}
				err = proto.Unmarshal(bz, pbb)
				if err != nil {
					panic(err)
				}
				ledger, err := kinds.LedgerOriginatingSchema(pbb)
				if err != nil {
					panic(err)
				}

				if ledger.Altitude != altitude+1 {
					panic(fmt.Sprintf("REDACTED", ledger.Altitude, altitude+1))
				}
				endorseAltitude := thatLedgerAddnEndorse.Altitude
				if endorseAltitude != altitude+1 {
					panic(fmt.Sprintf("REDACTED", endorseAltitude, altitude+1))
				}
				ledgers = append(ledgers, ledger)
				addnEndorses = append(addnEndorses, thatLedgerAddnEndorse)
				altitude++
			}
		case *kinds.FragmentAssignHeading:
			thatLedgerFragments = kinds.FreshFragmentAssignOriginatingHeading(*p)
		case *kinds.Fragment:
			_, err := thatLedgerFragments.AppendFragment(p)
			if err != nil {
				return nil, nil, err
			}
		case *kinds.Ballot:
			if p.Kind == commitchema.PreendorseKind {
				thatLedgerAddnEndorse = &kinds.ExpandedEndorse{
					Altitude:             p.Altitude,
					Iteration:              p.Iteration,
					LedgerUUID:            p.LedgerUUID,
					ExpandedNotations: []kinds.ExpandedEndorseSignature{p.ExpandedEndorseSignature()},
				}
			}
		}
	}
	//
	bz, err := io.ReadAll(thatLedgerFragments.ObtainFetcher())
	if err != nil {
		panic(err)
	}
	pbb := new(commitchema.Ledger)
	err = proto.Unmarshal(bz, pbb)
	if err != nil {
		panic(err)
	}
	ledger, err := kinds.LedgerOriginatingSchema(pbb)
	if err != nil {
		panic(err)
	}
	if ledger.Altitude != altitude+1 {
		panic(fmt.Sprintf("REDACTED", ledger.Altitude, altitude+1))
	}
	endorseAltitude := thatLedgerAddnEndorse.Altitude
	if endorseAltitude != altitude+1 {
		panic(fmt.Sprintf("REDACTED", endorseAltitude, altitude+1))
	}
	ledgers = append(ledgers, ledger)
	addnEndorses = append(addnEndorses, thatLedgerAddnEndorse)
	return ledgers, addnEndorses, nil
}

func retrievePortionOriginatingJournal(msg *ScheduledJournalSignal) any {
	//
	switch m := msg.Msg.(type) {
	case signalDetails:
		switch msg := m.Msg.(type) {
		case *NominationSignal:
			return &msg.Nomination.LedgerUUID.FragmentAssignHeading
		case *LedgerFragmentSignal:
			return msg.Fragment
		case *BallotSignal:
			return msg.Ballot
		}
	case TerminateAltitudeSignal:
		return m
	}

	return nil
}

//
func statusAlsoDepot(
	t *testing.T,
	settings *cfg.Settings,
	platformEdition uint64,
) (dbm.DB, sm.Status, *simulateLedgerDepot) {
	statusDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status, err := sm.CreateInaugurationStatusOriginatingRecord(settings.InaugurationRecord())
	require.NoError(t, err)
	status.Edition.Agreement.App = platformEdition
	depot := freshSimulateLedgerDepot(t, settings, status.AgreementSettings)
	require.NoError(t, statusDepot.Persist(status))

	return statusDatastore, status, depot
}

//
//

type simulateLedgerDepot struct {
	settings     *cfg.Settings
	parameters     kinds.AgreementSettings
	succession      []*kinds.Ledger
	addnEndorses []*kinds.ExpandedEndorse
	foundation       int64
	t          *testing.T
}

var _ sm.LedgerDepot = &simulateLedgerDepot{}

//
func freshSimulateLedgerDepot(t *testing.T, settings *cfg.Settings, parameters kinds.AgreementSettings) *simulateLedgerDepot {
	return &simulateLedgerDepot{
		settings: settings,
		parameters: parameters,
		t:      t,
	}
}

func (bs *simulateLedgerDepot) Altitude() int64                       { return int64(len(bs.succession)) }
func (bs *simulateLedgerDepot) Foundation() int64                         { return bs.foundation }
func (bs *simulateLedgerDepot) Extent() int64                         { return bs.Altitude() - bs.Foundation() + 1 }
func (bs *simulateLedgerDepot) FetchFoundationSummary() *kinds.LedgerSummary      { return bs.FetchLedgerSummary(bs.foundation) }
func (bs *simulateLedgerDepot) FetchLedger(altitude int64) *kinds.Ledger { return bs.succession[altitude-1] }
func (bs *simulateLedgerDepot) FetchLedgerViaDigest([]byte) *kinds.Ledger {
	return bs.succession[int64(len(bs.succession))-1]
}
func (bs *simulateLedgerDepot) FetchLedgerSummaryViaDigest([]byte) *kinds.LedgerSummary { return nil }
func (bs *simulateLedgerDepot) FetchLedgerSummary(altitude int64) *kinds.LedgerSummary {
	ledger := bs.succession[altitude-1]
	bps, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(bs.t, err)
	return &kinds.LedgerSummary{
		LedgerUUID: kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()},
		Heading:  ledger.Heading,
	}
}
func (bs *simulateLedgerDepot) FetchLedgerFragment(int64, int) *kinds.Fragment { return nil }
func (bs *simulateLedgerDepot) PersistLedgerUsingExpandedEndorse(*kinds.Ledger, *kinds.FragmentAssign, *kinds.ExpandedEndorse) {
}

func (bs *simulateLedgerDepot) PersistLedger(*kinds.Ledger, *kinds.FragmentAssign, *kinds.Endorse) {
}

func (bs *simulateLedgerDepot) FetchLedgerEndorse(altitude int64) *kinds.Endorse {
	return bs.addnEndorses[altitude-1].TowardEndorse()
}

func (bs *simulateLedgerDepot) FetchObservedEndorse(altitude int64) *kinds.Endorse {
	return bs.addnEndorses[altitude-1].TowardEndorse()
}

func (bs *simulateLedgerDepot) FetchLedgerExpandedEndorse(altitude int64) *kinds.ExpandedEndorse {
	return bs.addnEndorses[altitude-1]
}

func (bs *simulateLedgerDepot) TrimLedgers(altitude int64, _ sm.Status) (uint64, int64, error) {
	proofMark := altitude
	trimmed := uint64(0)
	for i := int64(0); i < altitude-1; i++ {
		bs.succession[i] = nil
		bs.addnEndorses[i] = nil
		trimmed++
	}
	bs.foundation = altitude
	return trimmed, proofMark, nil
}

func (bs *simulateLedgerDepot) EraseNewestLedger() error { return nil }
func (bs *simulateLedgerDepot) Shutdown() error             { return nil }

//
//

func VerifyNegotiationRevisionsAssessors(t *testing.T) {
	val, _ := kinds.ArbitraryAssessor(true, 10)
	values := kinds.FreshAssessorAssign([]*kinds.Assessor{val})
	app := &simulations.Platform{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyDetails{
		FinalLedgerAltitude: 0,
	}, nil)
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyInitializeSuccession{
		Assessors: kinds.Temp2buffer.AssessorRevisions(values),
	}, nil)
	customerOriginator := delegate.FreshRegionalCustomerOriginator(app)

	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	statusDatastore, status, depot := statusAlsoDepot(t, settings, 0x0)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	agedItemLocation := status.Assessors.Assessors[0].Location

	//
	producePaper, _ := sm.CreateInaugurationPaperOriginatingRecord(settings.InaugurationRecord())
	negotiator := FreshNegotiator(statusDepot, status, depot, producePaper)
	delegatePlatform := delegate.FreshPlatformLinks(customerOriginator, delegate.NooperationTelemetry())
	if err := delegatePlatform.Initiate(); err != nil {
		t.Fatalf("REDACTED", err)
	}
	t.Cleanup(func() {
		if err := delegatePlatform.Halt(); err != nil {
			t.Error(err)
		}
	})
	if err := negotiator.Negotiation(delegatePlatform); err != nil {
		t.Fatalf("REDACTED", err)
	}
	var err error
	//
	status, err = statusDepot.Fetch()
	require.NoError(t, err)

	freshItemLocation := status.Assessors.Assessors[0].Location
	anticipateItemLocation := val.Location
	assert.NotEqual(t, agedItemLocation, freshItemLocation)
	assert.Equal(t, freshItemLocation, anticipateItemLocation)
}
