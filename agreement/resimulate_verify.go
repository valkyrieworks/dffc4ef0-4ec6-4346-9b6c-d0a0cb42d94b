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

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/iface/kinds/simulations"
	cfg "github.com/valkyrieworks/settings"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	smemulators "github.com/valkyrieworks/status/simulations"
	"github.com/valkyrieworks/kinds"
)

func VerifyMain(m *testing.M) {
	settings = RestoreSettings("REDACTED")
	agreementResimulateSettings = RestoreSettings("REDACTED")
	settingsStatusVerify := RestoreSettings("REDACTED")
	settingsTxpoolVerify := RestoreSettings("REDACTED")
	settingsFaultyVerify := RestoreSettings("REDACTED")
	code := m.Run()
	os.RemoveAll(settings.OriginFolder)
	os.RemoveAll(agreementResimulateSettings.OriginFolder)
	os.RemoveAll(settingsStatusVerify.OriginFolder)
	os.RemoveAll(settingsTxpoolVerify.OriginFolder)
	os.RemoveAll(settingsFaultyVerify.OriginFolder)
	os.Exit(code)
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

func beginNewStatusAndWaitForLedger(
	t *testing.T,
	agreementResimulateSettings *cfg.Settings,
	ledgerStore dbm.DB,
	statusDepot sm.Depot,
) {
	tracer := log.VerifyingTracer()
	status, _ := statusDepot.ImportFromStoreOrOriginEntry(agreementResimulateSettings.OriginEntry())
	privateRatifier := importPrivateRatifier(agreementResimulateSettings)
	cs := newStatusWithSettingsAndLedgerDepot(
		agreementResimulateSettings,
		status,
		privateRatifier,
		objectdepot.NewInRamSoftware(),
		ledgerStore,
	)
	cs.AssignTracer(tracer)

	octets, _ := os.ReadFile(cs.settings.JournalEntry())
	t.Logf("REDACTED", octets)

	err := cs.Begin()
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
	newLedgerSubtract, err := cs.eventBus.Enrol(context.Background(), verifyEnrollee, kinds.EventInquireNewLedger)
	require.NoError(t, err)
	select {
	case <-newLedgerSubtract.Out():
	case <-newLedgerSubtract.Revoked():
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
			tx := objectdepot.NewTransferFromUID(i)
			if err := affirmTxpool(cs.transferAlerter).InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				if reply.Code != 0 {
					panic(fmt.Sprintf("REDACTED", reply.Code, reply.Log))
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
		label         string
		initFn       func(dbm.DB, *Status, context.Context)
		levelToHalt int64
	}{
		{
			"REDACTED",
			func(statusStore dbm.DB, cs *Status, ctx context.Context) {},
			1,
		},
		{
			"REDACTED",
			func(statusStore dbm.DB, cs *Status, ctx context.Context) {
				go transmitTrans(ctx, cs)
			},
			3,
		},
	}

	for i, tc := range verifyScenarios {
		agreementResimulateSettings := RestoreSettings(fmt.Sprintf("REDACTED", t.Name(), i))
		t.Run(tc.label, func(t *testing.T) {
			collapseWAArriveInspectVitality(t, agreementResimulateSettings, tc.initFn, tc.levelToHalt)
		})
	}
}

func collapseWAArriveInspectVitality(t *testing.T, agreementResimulateSettings *cfg.Settings,
	initFn func(dbm.DB, *Status, context.Context), levelToHalt int64,
) {
	journalAlarmed := make(chan error)
	collapsingJournal := &collapsingJournal{alarmChan: journalAlarmed, levelToHalt: levelToHalt}

	i := 1
Cycle:
	for {
		t.Logf("REDACTED", i)

		//
		tracer := log.NewNoopTracer()
		ledgerStore := dbm.NewMemoryStore()
		statusStore := ledgerStore
		statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
			DropIfaceReplies: false,
		})
		status, err := sm.CreateOriginStatusFromEntry(agreementResimulateSettings.OriginEntry())
		require.NoError(t, err)
		privateRatifier := importPrivateRatifier(agreementResimulateSettings)
		cs := newStatusWithSettingsAndLedgerDepot(
			agreementResimulateSettings,
			status,
			privateRatifier,
			objectdepot.NewInRamSoftware(),
			ledgerStore,
		)
		cs.AssignTracer(tracer)

		//
		ctx, revoke := context.WithCancel(context.Background())
		initFn(statusStore, cs, ctx)

		//
		journalEntry := cs.settings.JournalEntry()
		os.Remove(journalEntry)

		//
		csJournal, err := cs.AccessJournal(journalEntry)
		require.NoError(t, err)
		collapsingJournal.following = csJournal

		//
		collapsingJournal.messageOrdinal = 1
		cs.wal = collapsingJournal

		//
		err = cs.Begin()
		require.NoError(t, err)

		i++

		select {
		case err := <-journalAlarmed:
			t.Logf("REDACTED", err)

			//
			beginNewStatusAndWaitForLedger(t, agreementResimulateSettings, ledgerStore, statusDepot)

			//
			cs.Halt() //
			revoke()

			//
			if _, ok := err.(AttainedLevelToHaltFault); ok {
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
	alarmChan      chan error
	levelToHalt int64

	messageOrdinal                int //
	finalAlarmedForMessageOrdinal int //
}

var _ WAL = &collapsingJournal{}

//
type JournalRecordFault struct {
	msg string
}

func (e JournalRecordFault) Fault() string {
	return e.msg
}

//
//
type AttainedLevelToHaltFault struct {
	level int64
}

func (e AttainedLevelToHaltFault) Fault() string {
	return fmt.Sprintf("REDACTED", e.level)
}

//
//
func (w *collapsingJournal) Record(m JournalSignal) error {
	if terminateMessage, ok := m.(TerminateLevelSignal); ok {
		if terminateMessage.Level == w.levelToHalt {
			w.alarmChan <- AttainedLevelToHaltFault{terminateMessage.Level}
			runtime.Goexit()
			return nil
		}

		return w.following.Record(m)
	}

	if w.messageOrdinal > w.finalAlarmedForMessageOrdinal {
		w.finalAlarmedForMessageOrdinal = w.messageOrdinal
		_, entry, row, _ := runtime.Caller(1)
		w.alarmChan <- JournalRecordFault{fmt.Sprintf("REDACTED", m, entry, row)}
		runtime.Goexit()
		return nil
	}

	w.messageOrdinal++
	return w.following.Record(m)
}

func (w *collapsingJournal) RecordAlign(m JournalSignal) error {
	return w.Record(m)
}

func (w *collapsingJournal) PurgeAndAlign() error { return w.following.PurgeAndAlign() }

func (w *collapsingJournal) ScanForTerminateLevel(
	level int64,
	options *JournalScanSettings,
) (rd io.ReadCloser, located bool, err error) {
	return w.following.ScanForTerminateLevel(level, options)
}

func (w *collapsingJournal) Begin() error { return w.following.Begin() }
func (w *collapsingJournal) Halt() error  { return w.following.Halt() }
func (w *collapsingJournal) Wait()        { w.following.Wait() }

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
func configureSeriesWithModifyingRatifiers(t *testing.T, label string, nLedgers int) (*cfg.Settings, []*kinds.Ledger, []*kinds.ExpandedEndorse, sm.Status) {
	ctx, revoke := context.WithCancel(context.Background())
	defer revoke()

	nNodes := 7
	nValues := 4
	css, generatePaper, settings, sanitize := randomAgreementNetWithNodes(
		t,
		nValues,
		nNodes,
		label,
		newEmulateTimerFunction(true),
		func(_ string) iface.Software {
			return newObjectDepot()
		})
	originStatus, err := sm.CreateOriginStatus(generatePaper)
	require.NoError(t, err)
	t.Cleanup(sanitize)

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	newEpochChan := enrol(css[0].eventBus, kinds.EventInquireNewEpoch)
	nominationChan := enrol(css[0].eventBus, kinds.EventInquireFinishedNomination)

	vss := make([]*ratifierProxy, nNodes)
	for i := 0; i < nNodes; i++ {
		vss[i] = newRatifierProxy(css[i].privateRatifier, int32(i))
	}
	level, duration := css[0].Level, css[0].Cycle

	//
	beginVerifyEpoch(css[0], level, duration)
	augmentLevel(vss...)
	assureNewEpoch(newEpochChan, level, 0)
	assureNewNomination(nominationChan, level, duration)
	rs := css[0].FetchDurationStatus()
	attestAppendBallots(css[0], engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), true, vss[1:nValues]...)
	assureNewEpoch(newEpochChan, level+1, 0)

	//
	level++
	augmentLevel(vss...)
	newRatifierPublicKey1, err := css[nValues].privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	valuePublicKey1abci, err := cryptocode.PublicKeyToSchema(newRatifierPublicKey1)
	require.NoError(t, err)
	newRatifierTrans1 := objectdepot.CreateValueCollectionAlterTransfer(valuePublicKey1abci, verifyMinimumEnergy)
	err = affirmTxpool(css[0].transferAlerter).InspectTransfer(newRatifierTrans1, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	nominationLedger, err := css[0].instantiateNominationLedger(ctx) //
	require.NoError(t, err)
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: nominationLedger.Digest(), SegmentAssignHeading: nominationLedgerSegments.Heading()}

	nomination := kinds.NewNomination(vss[1].Level, duration, -1, ledgerUID)
	p := nomination.ToSchema()
	if err := vss[1].AttestNomination(verify.StandardVerifyLedgerUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Autograph = p.Autograph

	//
	if err := css[0].CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureNewNomination(nominationChan, level, duration)
	rs = css[0].FetchDurationStatus()
	attestAppendBallots(css[0], engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), true, vss[1:nValues]...)
	assureNewEpoch(newEpochChan, level+1, 0)

	//
	level++
	augmentLevel(vss...)
	modifyRatifierPublicKey1, err := css[nValues].privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	modifyPublicKey1abci, err := cryptocode.PublicKeyToSchema(modifyRatifierPublicKey1)
	require.NoError(t, err)
	modifyRatifierTrans1 := objectdepot.CreateValueCollectionAlterTransfer(modifyPublicKey1abci, 25)
	err = affirmTxpool(css[0].transferAlerter).InspectTransfer(modifyRatifierTrans1, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	nominationLedger, err = css[0].instantiateNominationLedger(ctx) //
	require.NoError(t, err)
	nominationLedgerSegments, err = nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	ledgerUID = kinds.LedgerUID{Digest: nominationLedger.Digest(), SegmentAssignHeading: nominationLedgerSegments.Heading()}

	nomination = kinds.NewNomination(vss[2].Level, duration, -1, ledgerUID)
	p = nomination.ToSchema()
	if err := vss[2].AttestNomination(verify.StandardVerifyLedgerUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Autograph = p.Autograph

	//
	if err := css[0].CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureNewNomination(nominationChan, level, duration)
	rs = css[0].FetchDurationStatus()
	attestAppendBallots(css[0], engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), true, vss[1:nValues]...)
	assureNewEpoch(newEpochChan, level+1, 0)

	//
	level++
	augmentLevel(vss...)
	newRatifierPublicKey2, err := css[nValues+1].privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	newVal2abci, err := cryptocode.PublicKeyToSchema(newRatifierPublicKey2)
	require.NoError(t, err)
	newRatifierTrans2 := objectdepot.CreateValueCollectionAlterTransfer(newVal2abci, verifyMinimumEnergy)
	err = affirmTxpool(css[0].transferAlerter).InspectTransfer(newRatifierTrans2, nil, txpool.TransferDetails{})
	require.NoError(t, err)
	newRatifierPublicKey3, err := css[nValues+2].privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	newVal3abci, err := cryptocode.PublicKeyToSchema(newRatifierPublicKey3)
	require.NoError(t, err)
	newRatifierTrans3 := objectdepot.CreateValueCollectionAlterTransfer(newVal3abci, verifyMinimumEnergy)
	err = affirmTxpool(css[0].transferAlerter).InspectTransfer(newRatifierTrans3, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	nominationLedger, err = css[0].instantiateNominationLedger(ctx) //
	require.NoError(t, err)
	nominationLedgerSegments, err = nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	ledgerUID = kinds.LedgerUID{Digest: nominationLedger.Digest(), SegmentAssignHeading: nominationLedgerSegments.Heading()}
	newVss := make([]*ratifierProxy, nValues+1)
	copy(newVss, vss[:nValues+1])
	sort.Sort(RatifierDummiesByEnergy(newVss))

	valueOrdinalFn := func(cssIdx int) int {
		for i, vs := range newVss {
			vsPublicKey, err := vs.FetchPublicKey()
			require.NoError(t, err)

			cssPublicKey, err := css[cssIdx].privateRatifier.FetchPublicKey()
			require.NoError(t, err)

			if vsPublicKey.Matches(cssPublicKey) {
				return i
			}
		}
		panic(fmt.Sprintf("REDACTED", cssIdx))
	}

	egoOrdinal := valueOrdinalFn(0)

	nomination = kinds.NewNomination(vss[3].Level, duration, -1, ledgerUID)
	p = nomination.ToSchema()
	if err := vss[3].AttestNomination(verify.StandardVerifyLedgerUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Autograph = p.Autograph

	//
	if err := css[0].CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureNewNomination(nominationChan, level, duration)

	deleteRatifierTrans2 := objectdepot.CreateValueCollectionAlterTransfer(newVal2abci, 0)
	err = affirmTxpool(css[0].transferAlerter).InspectTransfer(deleteRatifierTrans2, nil, txpool.TransferDetails{})
	assert.Nil(t, err)

	rs = css[0].FetchDurationStatus()
	for i := 0; i < nValues+1; i++ {
		if i == egoOrdinal {
			continue
		}
		attestAppendBallots(css[0], engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), true, newVss[i])
	}

	assureNewEpoch(newEpochChan, level+1, 0)

	//
	level++
	augmentLevel(vss...)
	//
	newVssIdx := valueOrdinalFn(nValues)
	newVss[newVssIdx].PollingEnergy = 25
	sort.Sort(RatifierDummiesByEnergy(newVss))
	egoOrdinal = valueOrdinalFn(0)
	assureNewNomination(nominationChan, level, duration)
	rs = css[0].FetchDurationStatus()
	for i := 0; i < nValues+1; i++ {
		if i == egoOrdinal {
			continue
		}
		attestAppendBallots(css[0], engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), true, newVss[i])
	}
	assureNewEpoch(newEpochChan, level+1, 0)

	//
	level++
	augmentLevel(vss...)
	deleteRatifierTrans3 := objectdepot.CreateValueCollectionAlterTransfer(newVal3abci, 0)
	err = affirmTxpool(css[0].transferAlerter).InspectTransfer(deleteRatifierTrans3, nil, txpool.TransferDetails{})
	assert.NoError(t, err)
	nominationLedger, err = css[0].instantiateNominationLedger(ctx) //
	require.NoError(t, err)
	nominationLedgerSegments, err = nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	ledgerUID = kinds.LedgerUID{Digest: nominationLedger.Digest(), SegmentAssignHeading: nominationLedgerSegments.Heading()}
	newVss = make([]*ratifierProxy, nValues+3)
	copy(newVss, vss[:nValues+3])
	sort.Sort(RatifierDummiesByEnergy(newVss))

	egoOrdinal = valueOrdinalFn(0)
	nomination = kinds.NewNomination(vss[1].Level, duration, -1, ledgerUID)
	p = nomination.ToSchema()
	if err := vss[1].AttestNomination(verify.StandardVerifyLedgerUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}
	nomination.Autograph = p.Autograph

	//
	if err := css[0].CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureNewNomination(nominationChan, level, duration)
	rs = css[0].FetchDurationStatus()
	for i := 0; i < nValues+3; i++ {
		if i == egoOrdinal {
			continue
		}
		attestAppendBallots(css[0], engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), true, newVss[i])
	}
	assureNewEpoch(newEpochChan, level+1, 0)

	series := []*kinds.Ledger{}
	extensionEndorses := []*kinds.ExpandedEndorse{}
	for i := 1; i <= nLedgers; i++ {
		series = append(series, css[0].ledgerDepot.ImportLedger(int64(i)))
		extensionEndorses = append(extensionEndorses, css[0].ledgerDepot.ImportLedgerExpandedEndorse(int64(i)))
	}
	return settings, series, extensionEndorses, originStatus
}

//
func VerifyGreetingResimulateAll(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, 0, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, 0, m, false)
		})
	}
}

//
func VerifyGreetingResimulateSome(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, 2, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, 2, m, true)
		})
	}
}

//
func VerifyGreetingResimulateOne(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, countLedgers-1, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, countLedgers-1, m, true)
		})
	}
}

//
func VerifyGreetingResimulateVoid(t *testing.T) {
	for _, m := range styles {
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, countLedgers, m, false)
		})
		t.Run(fmt.Sprintf("REDACTED", m), func(t *testing.T) {
			verifyGreetingResimulate(t, settings, countLedgers, m, true)
		})
	}
}

func tempJournalWithData(data []byte) string {
	journalEntry, err := os.CreateTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	_, err = journalEntry.Write(data)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if err := journalEntry.Close(); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	return journalEntry.Name()
}

//
//
func verifyGreetingResimulate(t *testing.T, settings *cfg.Settings, nLedgers int, style uint, verifyRatifiersAlter bool) {
	var (
		verifySettings   *cfg.Settings
		series        []*kinds.Ledger
		extensionEndorses   []*kinds.ExpandedEndorse
		depot        *emulateLedgerDepot
		statusStore      dbm.DB
		originStatus sm.Status
		txpool      = emptyTxpool{}
		eventpool       = sm.EmptyProofDepository{}
	)

	if verifyRatifiersAlter {
		verifySettings, series, extensionEndorses, originStatus = configureSeriesWithModifyingRatifiers(t, fmt.Sprintf("REDACTED", nLedgers, style), countLedgers)
		statusStore = dbm.NewMemoryStore()
		depot = newEmulateLedgerDepot(t, settings, originStatus.AgreementOptions)
	} else {
		verifySettings = RestoreSettings(fmt.Sprintf("REDACTED", nLedgers, style))
		t.Cleanup(func() {
			_ = os.RemoveAll(verifySettings.OriginFolder)
		})
		journalContent, err := JournalWithNLedgers(t, countLedgers, verifySettings)
		require.NoError(t, err)
		journalEntry := tempJournalWithData(journalContent)
		verifySettings.Agreement.CollectionJournalEntry(journalEntry)

		wal, err := NewJournal(journalEntry)
		require.NoError(t, err)
		wal.AssignTracer(log.VerifyingTracer())
		err = wal.Begin()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := wal.Halt(); err != nil {
				t.Error(err)
			}
		})
		series, extensionEndorses, err = createLedgerchainFromJournal(wal)
		require.NoError(t, err)
		statusStore, originStatus, depot = statusAndDepot(t, verifySettings, objectdepot.ApplicationRelease)
	}

	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	t.Cleanup(func() {
		_ = statusDepot.End()
	})
	depot.series = series
	depot.extensionEndorses = extensionEndorses

	status := originStatus.Clone()
	//
	status, newestApplicationDigest := constructTMStatusFromSeries(t, verifySettings, statusDepot, txpool, eventpool, status, series, nLedgers, style, depot)

	//
	objectdepotApplication := objectdepot.NewDurableSoftware(
		filepath.Join(verifySettings.StoreFolder(), fmt.Sprintf("REDACTED", nLedgers, style)))
	t.Cleanup(func() {
		_ = objectdepotApplication.End()
	})

	customerOriginator2 := gateway.NewNativeCustomerOriginator(objectdepotApplication)
	if nLedgers > 0 {
		//
		//
		gatewayApplication := gateway.NewApplicationLinks(customerOriginator2, gateway.NoopStats())
		statusDB1 := dbm.NewMemoryStore()
		mockStatusDepot := sm.NewDepot(statusDB1, sm.DepotSettings{
			DropIfaceReplies: false,
		})
		err := mockStatusDepot.Persist(originStatus)
		require.NoError(t, err)
		constructApplicationStatusFromSeries(t, gatewayApplication, mockStatusDepot, txpool, eventpool, originStatus, series, nLedgers, style, depot)
	}

	//
	anticipateFault := false
	if style == 3 {
		trimmed, _, err := depot.TrimLedgers(2, status)
		require.NoError(t, err)
		require.EqualValues(t, 1, trimmed)
		anticipateFault = int64(nLedgers) < 2
	}

	//
	generatePaper, err := sm.CreateOriginPaperFromEntry(verifySettings.OriginEntry())
	require.NoError(t, err)
	greeter := NewGreeter(statusDepot, status, depot, generatePaper)
	gatewayApplication := gateway.NewApplicationLinks(customerOriginator2, gateway.NoopStats())
	if err := gatewayApplication.Begin(); err != nil {
		t.Fatalf("REDACTED", err)
	}

	t.Cleanup(func() {
		if err := gatewayApplication.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	err = greeter.Greeting(gatewayApplication)
	if anticipateFault {
		require.Error(t, err)
		//
		return
	}
	require.NoError(t, err)

	//
	res, err := gatewayApplication.Inquire().Details(context.Background(), gateway.QueryDetails)
	require.NoError(t, err)

	//
	require.Equal(t, depot.Level(), res.FinalLedgerLevel)

	//
	status, err = statusDepot.Import()
	require.NoError(t, err)
	require.Equal(t, status.FinalLedgerLevel, res.FinalLedgerLevel)
	require.Equal(t, int64(countLedgers), res.FinalLedgerLevel)

	//
	if !bytes.Equal(newestApplicationDigest, res.FinalLedgerApplicationDigest) {
		t.Fatalf(
			"REDACTED",
			res.FinalLedgerApplicationDigest,
			newestApplicationDigest)
	}

	anticipatedLedgersToAlign := countLedgers - nLedgers
	if nLedgers == countLedgers && style > 0 {
		anticipatedLedgersToAlign++
	} else if nLedgers > 0 && style == 1 {
		anticipatedLedgersToAlign++
	}

	if greeter.NLedgers() != anticipatedLedgersToAlign {
		t.Fatalf("REDACTED", anticipatedLedgersToAlign, greeter.NLedgers())
	}
}

func executeLedger(t *testing.T, statusDepot sm.Depot, txpool txpool.Txpool, eventpool sm.ProofDepository, st sm.Status, blk *kinds.Ledger, gatewayApplication gateway.ApplicationLinks, bs sm.LedgerDepot) sm.Status {
	verifySegmentVolume := kinds.LedgerSegmentVolumeOctets
	ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplication.Agreement(), txpool, eventpool, bs)

	bps, err := blk.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: blk.Digest(), SegmentAssignHeading: bps.Heading()}
	newStatus, err := ledgerExecute.ExecuteLedger(st, ledgerUID, blk)
	require.NoError(t, err)
	return newStatus
}

func constructApplicationStatusFromSeries(t *testing.T, gatewayApplication gateway.ApplicationLinks, statusDepot sm.Depot, txpool txpool.Txpool, eventpool sm.ProofDepository,
	status sm.Status, series []*kinds.Ledger, nLedgers int, style uint, bs sm.LedgerDepot,
) {
	//
	if err := gatewayApplication.Begin(); err != nil {
		panic(err)
	}
	defer gatewayApplication.Halt() //

	status.Release.Agreement.App = objectdepot.ApplicationRelease //
	ratifiers := kinds.Tm2schema.RatifierRefreshes(status.Ratifiers)
	if _, err := gatewayApplication.Agreement().InitSeries(context.Background(), &iface.QueryInitSeries{
		Ratifiers: ratifiers,
	}); err != nil {
		panic(err)
	}
	if err := statusDepot.Persist(status); err != nil { //
		panic(err)
	}
	switch style {
	case 0:
		for i := 0; i < nLedgers; i++ {
			ledger := series[i]
			status = executeLedger(t, statusDepot, txpool, eventpool, status, ledger, gatewayApplication, bs)
		}
	case 1, 2, 3:
		for i := 0; i < nLedgers-1; i++ {
			ledger := series[i]
			status = executeLedger(t, statusDepot, txpool, eventpool, status, ledger, gatewayApplication, bs)
		}

		//
		//
		if style == 2 || style == 3 {
			//
			//
			//
			status = executeLedger(t, statusDepot, txpool, eventpool, status, series[nLedgers-1], gatewayApplication, bs)
		}
	default:
		panic(fmt.Sprintf("REDACTED", style))
	}
}

func constructTMStatusFromSeries(
	t *testing.T,
	settings *cfg.Settings,
	statusDepot sm.Depot,
	txpool txpool.Txpool,
	eventpool sm.ProofDepository,
	status sm.Status,
	series []*kinds.Ledger,
	nLedgers int,
	style uint,
	bs sm.LedgerDepot,
) (sm.Status, []byte) {
	//
	customerOriginator := gateway.NewNativeCustomerOriginator(
		objectdepot.NewDurableSoftware(
			filepath.Join(settings.StoreFolder(), fmt.Sprintf("REDACTED", nLedgers, style))))
	gatewayApplication := gateway.NewApplicationLinks(customerOriginator, gateway.NoopStats())
	if err := gatewayApplication.Begin(); err != nil {
		panic(err)
	}
	defer gatewayApplication.Halt() //

	status.Release.Agreement.App = objectdepot.ApplicationRelease //
	ratifiers := kinds.Tm2schema.RatifierRefreshes(status.Ratifiers)
	if _, err := gatewayApplication.Agreement().InitSeries(context.Background(), &iface.QueryInitSeries{
		Ratifiers: ratifiers,
	}); err != nil {
		panic(err)
	}
	if err := statusDepot.Persist(status); err != nil { //
		panic(err)
	}
	switch style {
	case 0:
		//
		for _, ledger := range series {
			status = executeLedger(t, statusDepot, txpool, eventpool, status, ledger, gatewayApplication, bs)
		}
		return status, status.ApplicationDigest

	case 1, 2, 3:
		//
		//
		for _, ledger := range series[:len(series)-1] {
			status = executeLedger(t, statusDepot, txpool, eventpool, status, ledger, gatewayApplication, bs)
		}

		mockStatusDepot := &smemulators.Depot{}
		finalLevel := int64(len(series))
		secondlastLevel := int64(len(series) - 1)
		values, _ := statusDepot.ImportRatifiers(secondlastLevel)
		mockStatusDepot.On("REDACTED", secondlastLevel).Return(values, nil)
		mockStatusDepot.On("REDACTED", mock.Anything).Return(nil)
		mockStatusDepot.On("REDACTED", finalLevel, mock.MatchedBy(func(reply *iface.ReplyCompleteLedger) bool {
			require.NoError(t, statusDepot.PersistCompleteLedgerReply(finalLevel, reply))
			return true
		})).Return(nil)

		//
		//
		s := executeLedger(t, mockStatusDepot, txpool, eventpool, status, series[len(series)-1], gatewayApplication, bs)
		return status, s.ApplicationDigest
	default:
		panic(fmt.Sprintf("REDACTED", style))
	}
}

func createLedgers(n int, status sm.Status, privateValues []kinds.PrivateRatifier) ([]*kinds.Ledger, error) {
	ledgerUID := verify.CreateLedgerUID()
	ledgers := make([]*kinds.Ledger, n)

	for i := 0; i < n; i++ {
		level := status.FinalLedgerLevel + 1 + int64(i)
		finalEndorse, err := verify.CreateEndorse(ledgerUID, level-1, 0, status.FinalRatifiers, privateValues, status.LedgerUID, status.FinalLedgerTime)
		if err != nil {
			return nil, err
		}
		ledger, err := status.CreateLedger(level, verify.CreateNTrans(level, 10), finalEndorse, nil, status.FinalRatifiers.Recommender.Location)
		if err != nil {
			return nil, err
		}
		ledgers[i] = ledger
		status.FinalLedgerUID = ledgerUID
		status.FinalLedgerLevel = level
		status.FinalLedgerTime = status.FinalLedgerTime.Add(1 * time.Second)
		status.FinalRatifiers = status.Ratifiers.Clone()
		status.Ratifiers = status.FollowingRatifiers.Clone()
		status.FollowingRatifiers = status.FollowingRatifiers.CloneAugmentRecommenderUrgency(1)
		status.ApplicationDigest = verify.ArbitraryDigest()

		ledgerUID = verify.CreateLedgerUIDWithDigest(ledger.Digest())
	}

	return ledgers, nil
}

func VerifyGreetingAlarmsIfApplicationYieldsIncorrectApplicationDigest(t *testing.T) {
	//
	//
	//
	//
	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	privateValue := privatekey.ImportEntryPrivatekey(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry())
	const applicationRelease = 0x0
	statusStore, status, depot := statusAndDepot(t, settings, applicationRelease)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	generatePaper, _ := sm.CreateOriginPaperFromEntry(settings.OriginEntry())
	status.FinalRatifiers = status.Ratifiers.Clone()
	//
	ledgers, err := createLedgers(3, status, []kinds.PrivateRatifier{privateValue})
	require.NoError(t, err)

	depot.series = ledgers

	//
	//
	//
	//
	{
		app := &flawedApplication{countLedgers: 3, allDigestsAreIncorrect: true}
		customerOriginator := gateway.NewNativeCustomerOriginator(app)
		gatewayApplication := gateway.NewApplicationLinks(customerOriginator, gateway.NoopStats())
		err := gatewayApplication.Begin()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := gatewayApplication.Halt(); err != nil {
				t.Error(err)
			}
		})

		assert.Panics(t, func() {
			h := NewGreeter(statusDepot, status, depot, generatePaper)
			if err = h.Greeting(gatewayApplication); err != nil {
				t.Log(err)
			}
		})
	}

	//
	//
	//
	//
	{
		app := &flawedApplication{countLedgers: 3, solelyFinalDigestIsIncorrect: true}
		customerOriginator := gateway.NewNativeCustomerOriginator(app)
		gatewayApplication := gateway.NewApplicationLinks(customerOriginator, gateway.NoopStats())
		err := gatewayApplication.Begin()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := gatewayApplication.Halt(); err != nil {
				t.Error(err)
			}
		})

		assert.Panics(t, func() {
			h := NewGreeter(statusDepot, status, depot, generatePaper)
			if err = h.Greeting(gatewayApplication); err != nil {
				t.Log(err)
			}
		})
	}
}

type flawedApplication struct {
	iface.RootSoftware
	countLedgers           byte
	level              byte
	allDigestsAreIncorrect   bool
	solelyFinalDigestIsIncorrect bool
}

func (app *flawedApplication) CompleteLedger(context.Context, *iface.QueryCompleteLedger) (*iface.ReplyCompleteLedger, error) {
	app.level++
	if app.solelyFinalDigestIsIncorrect {
		if app.level == app.countLedgers {
			return &iface.ReplyCompleteLedger{ApplicationDigest: engineseed.Octets(8)}, nil
		}
		return &iface.ReplyCompleteLedger{ApplicationDigest: []byte{app.level}}, nil
	} else if app.allDigestsAreIncorrect {
		return &iface.ReplyCompleteLedger{ApplicationDigest: engineseed.Octets(8)}, nil
	}

	panic("REDACTED")
}

//
//

func createLedgerchainFromJournal(wal WAL) ([]*kinds.Ledger, []*kinds.ExpandedEndorse, error) {
	var level int64

	//
	gr, located, err := wal.ScanForTerminateLevel(level, &JournalScanSettings{})
	if err != nil {
		return nil, nil, err
	}
	if !located {
		return nil, nil, fmt.Errorf("REDACTED", level)
	}
	defer gr.Close()

	//

	var (
		ledgers             []*kinds.Ledger
		extensionEndorses         []*kinds.ExpandedEndorse
		thisLedgerSegments     *kinds.SegmentCollection
		thisLedgerExtensionEndorse *kinds.ExpandedEndorse
	)

	dec := NewJournalParser(gr)
	for {
		msg, err := dec.Parse()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}

		fraction := scanFractionFromJournal(msg)
		if fraction == nil {
			continue
		}

		switch p := fraction.(type) {
		case TerminateLevelSignal:
			//
			if thisLedgerSegments != nil {
				pbb := new(engineproto.Ledger)
				bz, err := io.ReadAll(thisLedgerSegments.FetchScanner())
				if err != nil {
					panic(err)
				}
				err = proto.Unmarshal(bz, pbb)
				if err != nil {
					panic(err)
				}
				ledger, err := kinds.LedgerFromSchema(pbb)
				if err != nil {
					panic(err)
				}

				if ledger.Level != level+1 {
					panic(fmt.Sprintf("REDACTED", ledger.Level, level+1))
				}
				endorseLevel := thisLedgerExtensionEndorse.Level
				if endorseLevel != level+1 {
					panic(fmt.Sprintf("REDACTED", endorseLevel, level+1))
				}
				ledgers = append(ledgers, ledger)
				extensionEndorses = append(extensionEndorses, thisLedgerExtensionEndorse)
				level++
			}
		case *kinds.SegmentAssignHeading:
			thisLedgerSegments = kinds.NewSegmentCollectionFromHeading(*p)
		case *kinds.Segment:
			_, err := thisLedgerSegments.AppendSegment(p)
			if err != nil {
				return nil, nil, err
			}
		case *kinds.Ballot:
			if p.Kind == engineproto.PreendorseKind {
				thisLedgerExtensionEndorse = &kinds.ExpandedEndorse{
					Level:             p.Level,
					Cycle:              p.Cycle,
					LedgerUID:            p.LedgerUID,
					ExpandedEndorsements: []kinds.ExpandedEndorseSignature{p.ExpandedEndorseSignature()},
				}
			}
		}
	}
	//
	bz, err := io.ReadAll(thisLedgerSegments.FetchScanner())
	if err != nil {
		panic(err)
	}
	pbb := new(engineproto.Ledger)
	err = proto.Unmarshal(bz, pbb)
	if err != nil {
		panic(err)
	}
	ledger, err := kinds.LedgerFromSchema(pbb)
	if err != nil {
		panic(err)
	}
	if ledger.Level != level+1 {
		panic(fmt.Sprintf("REDACTED", ledger.Level, level+1))
	}
	endorseLevel := thisLedgerExtensionEndorse.Level
	if endorseLevel != level+1 {
		panic(fmt.Sprintf("REDACTED", endorseLevel, level+1))
	}
	ledgers = append(ledgers, ledger)
	extensionEndorses = append(extensionEndorses, thisLedgerExtensionEndorse)
	return ledgers, extensionEndorses, nil
}

func scanFractionFromJournal(msg *ScheduledJournalSignal) any {
	//
	switch m := msg.Msg.(type) {
	case messageDetails:
		switch msg := m.Msg.(type) {
		case *NominationSignal:
			return &msg.Nomination.LedgerUID.SegmentAssignHeading
		case *LedgerSegmentSignal:
			return msg.Segment
		case *BallotSignal:
			return msg.Ballot
		}
	case TerminateLevelSignal:
		return m
	}

	return nil
}

//
func statusAndDepot(
	t *testing.T,
	settings *cfg.Settings,
	applicationRelease uint64,
) (dbm.DB, sm.Status, *emulateLedgerDepot) {
	statusStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := sm.CreateOriginStatusFromEntry(settings.OriginEntry())
	require.NoError(t, err)
	status.Release.Agreement.App = applicationRelease
	depot := newEmulateLedgerDepot(t, settings, status.AgreementOptions)
	require.NoError(t, statusDepot.Persist(status))

	return statusStore, status, depot
}

//
//

type emulateLedgerDepot struct {
	settings     *cfg.Settings
	options     kinds.AgreementOptions
	series      []*kinds.Ledger
	extensionEndorses []*kinds.ExpandedEndorse
	root       int64
	t          *testing.T
}

var _ sm.LedgerDepot = &emulateLedgerDepot{}

//
func newEmulateLedgerDepot(t *testing.T, settings *cfg.Settings, options kinds.AgreementOptions) *emulateLedgerDepot {
	return &emulateLedgerDepot{
		settings: settings,
		options: options,
		t:      t,
	}
}

func (bs *emulateLedgerDepot) Level() int64                       { return int64(len(bs.series)) }
func (bs *emulateLedgerDepot) Root() int64                         { return bs.root }
func (bs *emulateLedgerDepot) Volume() int64                         { return bs.Level() - bs.Root() + 1 }
func (bs *emulateLedgerDepot) ImportRootMeta() *kinds.LedgerMeta      { return bs.ImportLedgerMeta(bs.root) }
func (bs *emulateLedgerDepot) ImportLedger(level int64) *kinds.Ledger { return bs.series[level-1] }
func (bs *emulateLedgerDepot) ImportLedgerByDigest([]byte) *kinds.Ledger {
	return bs.series[int64(len(bs.series))-1]
}
func (bs *emulateLedgerDepot) ImportLedgerMetaByDigest([]byte) *kinds.LedgerMeta { return nil }
func (bs *emulateLedgerDepot) ImportLedgerMeta(level int64) *kinds.LedgerMeta {
	ledger := bs.series[level-1]
	bps, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(bs.t, err)
	return &kinds.LedgerMeta{
		LedgerUID: kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()},
		Heading:  ledger.Heading,
	}
}
func (bs *emulateLedgerDepot) ImportLedgerSegment(int64, int) *kinds.Segment { return nil }
func (bs *emulateLedgerDepot) PersistLedgerWithExpandedEndorse(*kinds.Ledger, *kinds.SegmentCollection, *kinds.ExpandedEndorse) {
}

func (bs *emulateLedgerDepot) PersistLedger(*kinds.Ledger, *kinds.SegmentCollection, *kinds.Endorse) {
}

func (bs *emulateLedgerDepot) ImportLedgerEndorse(level int64) *kinds.Endorse {
	return bs.extensionEndorses[level-1].ToEndorse()
}

func (bs *emulateLedgerDepot) ImportViewedEndorse(level int64) *kinds.Endorse {
	return bs.extensionEndorses[level-1].ToEndorse()
}

func (bs *emulateLedgerDepot) ImportLedgerExpandedEndorse(level int64) *kinds.ExpandedEndorse {
	return bs.extensionEndorses[level-1]
}

func (bs *emulateLedgerDepot) TrimLedgers(level int64, _ sm.Status) (uint64, int64, error) {
	proofSpot := level
	trimmed := uint64(0)
	for i := int64(0); i < level-1; i++ {
		bs.series[i] = nil
		bs.extensionEndorses[i] = nil
		trimmed++
	}
	bs.root = level
	return trimmed, proofSpot, nil
}

func (bs *emulateLedgerDepot) RemoveNewestLedger() error { return nil }
func (bs *emulateLedgerDepot) End() error             { return nil }

//
//

func VerifyGreetingRefreshesRatifiers(t *testing.T) {
	val, _ := kinds.RandomRatifier(true, 10)
	values := kinds.NewRatifierCollection([]*kinds.Ratifier{val})
	app := &simulations.Software{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyDetails{
		FinalLedgerLevel: 0,
	}, nil)
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyInitSeries{
		Ratifiers: kinds.Tm2schema.RatifierRefreshes(values),
	}, nil)
	customerOriginator := gateway.NewNativeCustomerOriginator(app)

	settings := RestoreSettings("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	statusStore, status, depot := statusAndDepot(t, settings, 0x0)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	agedValueAddress := status.Ratifiers.Ratifiers[0].Location

	//
	generatePaper, _ := sm.CreateOriginPaperFromEntry(settings.OriginEntry())
	greeter := NewGreeter(statusDepot, status, depot, generatePaper)
	gatewayApplication := gateway.NewApplicationLinks(customerOriginator, gateway.NoopStats())
	if err := gatewayApplication.Begin(); err != nil {
		t.Fatalf("REDACTED", err)
	}
	t.Cleanup(func() {
		if err := gatewayApplication.Halt(); err != nil {
			t.Error(err)
		}
	})
	if err := greeter.Greeting(gatewayApplication); err != nil {
		t.Fatalf("REDACTED", err)
	}
	var err error
	//
	status, err = statusDepot.Import()
	require.NoError(t, err)

	newValueAddress := status.Ratifiers.Ratifiers[0].Location
	anticipateValueAddress := val.Location
	assert.NotEqual(t, agedValueAddress, newValueAddress)
	assert.Equal(t, newValueAddress, anticipateValueAddress)
}
