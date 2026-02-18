package agreement

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/go-kit/log/term"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	abciend "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	cfg "github.com/valkyrieworks/settings"
	cskinds "github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/intrinsic/verify"
	cometbytes "github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	cometos "github.com/valkyrieworks/utils/os"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	engineconnect "github.com/valkyrieworks/utils/align"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

const (
	verifyEnrollee = "REDACTED"
)

//
//
type sanitizeFunction func()

//
var (
	settings                *cfg.Settings //
	agreementResimulateSettings *cfg.Settings
	assureDeadline         = time.Millisecond * 200
)

func assureFolder(dir string, style os.FileMode) {
	if err := cometos.AssureFolder(dir, style); err != nil {
		panic(err)
	}
}

func RestoreSettings(label string) *cfg.Settings {
	return verify.RestoreVerifyOrigin(label)
}

//
//

type ratifierProxy struct {
	Ordinal  int32 //
	Level int64
	Cycle  int32
	kinds.PrivateRatifier
	PollingEnergy int64
	finalBallot    *kinds.Ballot
}

var verifyMinimumEnergy int64 = 10

func newRatifierProxy(privateRatifier kinds.PrivateRatifier, valueOrdinal int32) *ratifierProxy {
	return &ratifierProxy{
		Ordinal:         valueOrdinal,
		PrivateRatifier: privateRatifier,
		PollingEnergy:   verifyMinimumEnergy,
	}
}

func (vs *ratifierProxy) attestBallot(
	ballotKind engineproto.AttestedMessageKind,
	digest []byte,
	heading kinds.SegmentAssignHeading,
	ballotAddition []byte,
	extensionActivated bool,
) (*kinds.Ballot, error) {
	publicKey, err := vs.FetchPublicKey()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	ballot := &kinds.Ballot{
		Kind:             ballotKind,
		Level:           vs.Level,
		Cycle:            vs.Cycle,
		LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: heading},
		Timestamp:        engineclock.Now(),
		RatifierLocation: publicKey.Location(),
		RatifierOrdinal:   vs.Ordinal,
		Addition:        ballotAddition,
	}
	v := ballot.ToSchema()
	if err = vs.AttestBallot(verify.StandardVerifyLedgerUID, v); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	if attestDataIsEquivalent(vs.finalBallot, v) {
		v.Autograph = vs.finalBallot.Autograph
		v.Timestamp = vs.finalBallot.Timestamp
		v.AdditionAutograph = vs.finalBallot.AdditionAutograph
	}

	ballot.Autograph = v.Autograph
	ballot.Timestamp = v.Timestamp
	ballot.AdditionAutograph = v.AdditionAutograph

	if !extensionActivated {
		ballot.AdditionAutograph = nil
	}

	return ballot, err
}

//
func attestBallot(vs *ratifierProxy, ballotKind engineproto.AttestedMessageKind, digest []byte, heading kinds.SegmentAssignHeading, extensionActivated bool) *kinds.Ballot {
	var ext []byte
	//
	if extensionActivated {
		if ballotKind != engineproto.PreendorseKind {
			panic(fmt.Errorf("REDACTED"))
		}
		if len(digest) != 0 || !heading.IsNil() {
			ext = []byte("REDACTED")
		}
	}
	v, err := vs.attestBallot(ballotKind, digest, heading, ext, extensionActivated)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	vs.finalBallot = v

	return v
}

func attestBallots(
	ballotKind engineproto.AttestedMessageKind,
	digest []byte,
	heading kinds.SegmentAssignHeading,
	extensionActivated bool,
	vss ...*ratifierProxy,
) []*kinds.Ballot {
	ballots := make([]*kinds.Ballot, len(vss))
	for i, vs := range vss {
		ballots[i] = attestBallot(vs, ballotKind, digest, heading, extensionActivated)
	}
	return ballots
}

func augmentLevel(vss ...*ratifierProxy) {
	for _, vs := range vss {
		vs.Level++
	}
}

func augmentEpoch(vss ...*ratifierProxy) {
	for _, vs := range vss {
		vs.Cycle++
	}
}

type RatifierDummiesByEnergy []*ratifierProxy

func (vss RatifierDummiesByEnergy) Len() int {
	return len(vss)
}

func (vss RatifierDummiesByEnergy) Lower(i, j int) bool {
	vssi, err := vss[i].FetchPublicKey()
	if err != nil {
		panic(err)
	}
	vssj, err := vss[j].FetchPublicKey()
	if err != nil {
		panic(err)
	}

	if vss[i].PollingEnergy == vss[j].PollingEnergy {
		return bytes.Compare(vssi.Location(), vssj.Location()) == -1
	}
	return vss[i].PollingEnergy > vss[j].PollingEnergy
}

func (vss RatifierDummiesByEnergy) Exchange(i, j int) {
	it := vss[i]
	vss[i] = vss[j]
	vss[i].Ordinal = int32(i)
	vss[j] = it
	vss[j].Ordinal = int32(j)
}

//
//

func beginVerifyEpoch(cs *Status, level int64, duration int32) {
	cs.joinNewEpoch(level, duration)
	cs.beginProcedures(0)
}

//
func determineNomination(
	ctx context.Context,
	t *testing.T,
	cs1 *Status,
	vs *ratifierProxy,
	level int64,
	duration int32,
) (*kinds.Nomination, *kinds.Ledger) {
	cs1.mtx.Lock()
	ledger, err := cs1.instantiateNominationLedger(ctx)
	require.NoError(t, err)
	ledgerSegments, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	soundEpoch := cs1.SoundEpoch
	ledgerUID := cs1.status.LedgerUID
	cs1.mtx.Unlock()
	if ledger == nil {
		panic("REDACTED")
	}

	//
	polEpoch, nominationLedgerUID := soundEpoch, kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: ledgerSegments.Heading()}
	nomination := kinds.NewNomination(level, duration, polEpoch, nominationLedgerUID)
	p := nomination.ToSchema()
	if err := vs.AttestNomination(ledgerUID, p); err != nil {
		panic(err)
	}

	nomination.Autograph = p.Autograph

	return nomination, ledger
}

func appendBallots(to *Status, ballots ...*kinds.Ballot) {
	for _, ballot := range ballots {
		to.nodeMessageBuffer <- messageDetails{Msg: &BallotSignal{ballot}}
	}
}

func attestAppendBallots(
	to *Status,
	ballotKind engineproto.AttestedMessageKind,
	digest []byte,
	heading kinds.SegmentAssignHeading,
	extensionActivated bool,
	vss ...*ratifierProxy,
) {
	ballots := attestBallots(ballotKind, digest, heading, extensionActivated, vss...)
	appendBallots(to, ballots...)
}

func certifyPreballot(t *testing.T, cs *Status, duration int32, privateValue *ratifierProxy, ledgerDigest []byte) {
	preballots := cs.Ballots.Preballots(duration)
	publicKey, err := privateValue.FetchPublicKey()
	require.NoError(t, err)
	location := publicKey.Location()
	var ballot *kinds.Ballot
	if ballot = preballots.FetchByLocation(location); ballot == nil {
		panic("REDACTED")
	}
	if ledgerDigest == nil {
		if ballot.LedgerUID.Digest != nil {
			panic(fmt.Sprintf("REDACTED", ballot.LedgerUID.Digest))
		}
	} else {
		if !bytes.Equal(ballot.LedgerUID.Digest, ledgerDigest) {
			panic(fmt.Sprintf("REDACTED", ledgerDigest, ballot.LedgerUID.Digest))
		}
	}
}

func certifyFinalPreendorse(t *testing.T, cs *Status, privateValue *ratifierProxy, ledgerDigest []byte) {
	ballots := cs.FinalEndorse
	pv, err := privateValue.FetchPublicKey()
	require.NoError(t, err)
	location := pv.Location()
	var ballot *kinds.Ballot
	if ballot = ballots.FetchByLocation(location); ballot == nil {
		panic("REDACTED")
	}
	if !bytes.Equal(ballot.LedgerUID.Digest, ledgerDigest) {
		panic(fmt.Sprintf("REDACTED", ledgerDigest, ballot.LedgerUID.Digest))
	}
}

func certifyPreendorse(
	t *testing.T,
	cs *Status,
	thisEpoch,
	secureEpoch int32,
	privateValue *ratifierProxy,
	polledLedgerDigest,
	latchedLedgerDigest []byte,
) {
	preendorsements := cs.Ballots.Preendorsements(thisEpoch)
	pv, err := privateValue.FetchPublicKey()
	require.NoError(t, err)
	location := pv.Location()
	var ballot *kinds.Ballot
	if ballot = preendorsements.FetchByLocation(location); ballot == nil {
		panic("REDACTED")
	}

	if polledLedgerDigest == nil {
		if ballot.LedgerUID.Digest != nil {
			panic("REDACTED")
		}
	} else {
		if !bytes.Equal(ballot.LedgerUID.Digest, polledLedgerDigest) {
			panic("REDACTED")
		}
	}

	rs := cs.FetchDurationStatus()
	if latchedLedgerDigest == nil {
		if rs.LatchedEpoch != secureEpoch || rs.LatchedLedger != nil {
			panic(fmt.Sprintf(
				"REDACTED",
				secureEpoch,
				rs.LatchedEpoch,
				rs.LatchedLedger))
		}
	} else {
		if rs.LatchedEpoch != secureEpoch || !bytes.Equal(rs.LatchedLedger.Digest(), latchedLedgerDigest) {
			panic(fmt.Sprintf(
				"REDACTED",
				secureEpoch,
				rs.LatchedEpoch,
				rs.LatchedLedger.Digest(),
				latchedLedgerDigest))
		}
	}
}

func enrolToPoller(cs *Status, address []byte) <-chan cometbroadcast.Signal {
	ballotsSubtract, err := cs.eventBus.EnrolUnbuffered(context.Background(), verifyEnrollee, kinds.EventInquireBallot)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", verifyEnrollee, kinds.EventInquireBallot))
	}
	ch := make(chan cometbroadcast.Signal)
	go func() {
		for msg := range ballotsSubtract.Out() {
			ballot := msg.Data().(kinds.EventDataBallot)
			//
			if bytes.Equal(address, ballot.Ballot.RatifierLocation) {
				ch <- msg
			}
		}
	}()
	return ch
}

//
//

func newStatus(status sm.Status, pv kinds.PrivateRatifier, app iface.Software) *Status {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	return newStatusWithSettings(settings, status, pv, app)
}

func newStatusWithSettings(
	thisSettings *cfg.Settings,
	status sm.Status,
	pv kinds.PrivateRatifier,
	app iface.Software,
) *Status {
	ledgerStore := dbm.NewMemoryStore()
	return newStatusWithSettingsAndLedgerDepot(thisSettings, status, pv, app, ledgerStore)
}

func newStatusWithSettingsAndLedgerDepot(
	thisSettings *cfg.Settings,
	status sm.Status,
	pv kinds.PrivateRatifier,
	app iface.Software,
	ledgerStore dbm.DB,
) *Status {
	//
	ledgerDepot := depot.NewLedgerDepot(ledgerStore)

	//
	mtx := new(engineconnect.Lock)

	gatewayApplicationLinkConnect := gateway.NewApplicationLinkAgreement(abciend.NewNativeCustomer(mtx, app), gateway.NoopStats())
	gatewayApplicationLinkMemory := gateway.NewApplicationLinkTxpool(abciend.NewNativeCustomer(mtx, app), gateway.NoopStats())
	//
	memplStats := txpool.NoopStats()

	//
	txpool := txpool.NewCCatalogTxpool(settings.Txpool,
		gatewayApplicationLinkMemory,
		status.FinalLedgerLevel,
		txpool.WithStats(memplStats),
		txpool.WithPreInspect(sm.TransferPreInspect(status)),
		txpool.WithSubmitInspect(sm.TransferSubmitInspect(status)))

	if thisSettings.Agreement.WaitForTrans() {
		txpool.ActivateTransAccessible()
	}

	eventpool := sm.EmptyProofDepository{}

	//
	statusStore := ledgerStore
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	if err := statusDepot.Persist(status); err != nil { //
		panic(err)
	}

	ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplicationLinkConnect, txpool, eventpool, ledgerDepot)
	cs := NewStatus(thisSettings.Agreement, status, ledgerExecute, ledgerDepot, txpool, eventpool)
	cs.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	cs.CollectionPrivateRatifier(pv)

	eventBus := kinds.NewEventBus()
	eventBus.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	err := eventBus.Begin()
	if err != nil {
		panic(err)
	}
	cs.AssignEventBus(eventBus)
	return cs
}

func importPrivateRatifier(settings *cfg.Settings) *privatekey.EntryPV {
	privateRatifierKeyEntry := settings.PrivateRatifierKeyEntry()
	assureFolder(filepath.Dir(privateRatifierKeyEntry), 0o700)
	privateRatifierStatusEntry := settings.PrivateRatifierStatusEntry()
	privateRatifier := privatekey.ImportOrGenerateEntryPV(privateRatifierKeyEntry, privateRatifierStatusEntry)
	privateRatifier.Restore()
	return privateRatifier
}

func randomStatus(nRatifiers int) (*Status, []*ratifierProxy) {
	return randomStatusWithApplication(nRatifiers, objectdepot.NewInRamSoftware())
}

func randomStatusWithApplicationWithLevel(
	nRatifiers int,
	app iface.Software,
	level int64,
) (*Status, []*ratifierProxy) {
	c := verify.AgreementOptions()
	c.Iface.BallotPluginsActivateLevel = level
	return randomStatusWithApplicationImpl(nRatifiers, app, c)
}

func randomStatusWithApplication(nRatifiers int, app iface.Software) (*Status, []*ratifierProxy) {
	c := verify.AgreementOptions()
	return randomStatusWithApplicationImpl(nRatifiers, app, c)
}

func randomStatusWithApplicationImpl(
	nRatifiers int,
	app iface.Software,
	agreementOptions *kinds.AgreementOptions,
) (*Status, []*ratifierProxy) {
	//
	status, privateValues := randomOriginStatus(nRatifiers, false, 10, agreementOptions)

	vss := make([]*ratifierProxy, nRatifiers)

	cs := newStatus(status, privateValues[0], app)

	for i := 0; i < nRatifiers; i++ {
		vss[i] = newRatifierProxy(privateValues[i], int32(i))
	}
	//
	augmentLevel(vss[1:]...)

	return cs, vss
}

//

func assureNoNewEvent(ch <-chan cometbroadcast.Signal, deadline time.Duration,
	faultSignal string,
) {
	select {
	case <-time.After(deadline):
		break
	case <-ch:
		panic(faultSignal)
	}
}

func assureNoNewEventOnConduit(ch <-chan cometbroadcast.Signal) {
	assureNoNewEvent(
		ch,
		assureDeadline*8/10, //
		"REDACTED")
}

func assureNoNewEpochPhase(phaseChan <-chan cometbroadcast.Signal) {
	assureNoNewEvent(
		phaseChan,
		assureDeadline,
		"REDACTED")
}

func assureNoNewRelease(releaseChan <-chan cometbroadcast.Signal) {
	assureNoNewEvent(
		releaseChan,
		assureDeadline,
		"REDACTED")
}

func assureNoNewDeadline(phaseChan <-chan cometbroadcast.Signal, deadline int64) {
	deadlinePeriod := time.Duration(deadline*10) * time.Nanosecond
	assureNoNewEvent(
		phaseChan,
		deadlinePeriod,
		"REDACTED")
}

func assureNewEvent(ch <-chan cometbroadcast.Signal, level int64, duration int32, deadline time.Duration, faultSignal string) {
	select {
	case <-time.After(deadline):
		panic(faultSignal)
	case msg := <-ch:
		epochStatusEvent, ok := msg.Data().(kinds.EventDataDurationStatus)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if epochStatusEvent.Level != level {
			panic(fmt.Sprintf("REDACTED", level, epochStatusEvent.Level))
		}
		if epochStatusEvent.Cycle != duration {
			panic(fmt.Sprintf("REDACTED", duration, epochStatusEvent.Cycle))
		}
		//
	}
}

func assureNewEpoch(epochChan <-chan cometbroadcast.Signal, level int64, duration int32) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-epochChan:
		newEpochEvent, ok := msg.Data().(kinds.EventDataNewEpoch)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if newEpochEvent.Level != level {
			panic(fmt.Sprintf("REDACTED", level, newEpochEvent.Level))
		}
		if newEpochEvent.Cycle != duration {
			panic(fmt.Sprintf("REDACTED", duration, newEpochEvent.Cycle))
		}
	}
}

func assureNewDeadline(deadlineChan <-chan cometbroadcast.Signal, level int64, duration int32, deadline int64) {
	deadlinePeriod := time.Duration(deadline*10) * time.Nanosecond
	assureNewEvent(deadlineChan, level, duration, deadlinePeriod,
		"REDACTED")
}

func assureNewNomination(nominationChan <-chan cometbroadcast.Signal, level int64, duration int32) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-nominationChan:
		nominationEvent, ok := msg.Data().(kinds.EventDataFinishedNomination)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if nominationEvent.Level != level {
			panic(fmt.Sprintf("REDACTED", level, nominationEvent.Level))
		}
		if nominationEvent.Cycle != duration {
			panic(fmt.Sprintf("REDACTED", duration, nominationEvent.Cycle))
		}
	}
}

func assureNewSoundLedger(soundLedgerChan <-chan cometbroadcast.Signal, level int64, duration int32) {
	assureNewEvent(soundLedgerChan, level, duration, assureDeadline,
		"REDACTED")
}

func assureNewLedger(ledgerChan <-chan cometbroadcast.Signal, level int64) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-ledgerChan:
		ledgerEvent, ok := msg.Data().(kinds.EventDataNewLedger)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if ledgerEvent.Ledger.Level != level {
			panic(fmt.Sprintf("REDACTED", level, ledgerEvent.Ledger.Level))
		}
	}
}

func assureNewLedgerHeading(ledgerChan <-chan cometbroadcast.Signal, level int64, ledgerDigest cometbytes.HexOctets) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-ledgerChan:
		ledgerHeadingEvent, ok := msg.Data().(kinds.EventDataNewLedgerHeading)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if ledgerHeadingEvent.Heading.Level != level {
			panic(fmt.Sprintf("REDACTED", level, ledgerHeadingEvent.Heading.Level))
		}
		if !bytes.Equal(ledgerHeadingEvent.Heading.Digest(), ledgerDigest) {
			panic(fmt.Sprintf("REDACTED", ledgerDigest, ledgerHeadingEvent.Heading.Digest()))
		}
	}
}

func assureNewRelease(releaseChan <-chan cometbroadcast.Signal, level int64, duration int32) {
	assureNewEvent(releaseChan, level, duration, assureDeadline,
		"REDACTED")
}

func assureNomination(nominationChan <-chan cometbroadcast.Signal, level int64, duration int32, nominationUID kinds.LedgerUID) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-nominationChan:
		nominationEvent, ok := msg.Data().(kinds.EventDataFinishedNomination)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if nominationEvent.Level != level {
			panic(fmt.Sprintf("REDACTED", level, nominationEvent.Level))
		}
		if nominationEvent.Cycle != duration {
			panic(fmt.Sprintf("REDACTED", duration, nominationEvent.Cycle))
		}
		if !nominationEvent.LedgerUID.Matches(nominationUID) {
			panic(fmt.Sprintf("REDACTED", nominationEvent.LedgerUID, nominationUID))
		}
	}
}

func assurePreendorse(ballotChan <-chan cometbroadcast.Signal, level int64, duration int32) {
	assureBallot(ballotChan, level, duration, engineproto.PreendorseKind)
}

func assurePreballot(ballotChan <-chan cometbroadcast.Signal, level int64, duration int32) {
	assureBallot(ballotChan, level, duration, engineproto.PreballotKind)
}

func assureBallot(ballotChan <-chan cometbroadcast.Signal, level int64, duration int32,
	ballotKind engineproto.AttestedMessageKind,
) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-ballotChan:
		ballotEvent, ok := msg.Data().(kinds.EventDataBallot)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		ballot := ballotEvent.Ballot
		if ballot.Level != level {
			panic(fmt.Sprintf("REDACTED", level, ballot.Level))
		}
		if ballot.Cycle != duration {
			panic(fmt.Sprintf("REDACTED", duration, ballot.Cycle))
		}
		if ballot.Kind != ballotKind {
			panic(fmt.Sprintf("REDACTED", ballotKind, ballot.Kind))
		}
	}
}

func assurePreballotAlign(t *testing.T, ballotChan <-chan cometbroadcast.Signal, level int64, duration int32, digest []byte) {
	t.Helper()
	assureBallotAlign(t, ballotChan, level, duration, digest, engineproto.PreballotKind)
}

func assurePreendorseAlign(t *testing.T, ballotChan <-chan cometbroadcast.Signal, level int64, duration int32, digest []byte) {
	t.Helper()
	assureBallotAlign(t, ballotChan, level, duration, digest, engineproto.PreendorseKind)
}

func assureBallotAlign(t *testing.T, ballotChan <-chan cometbroadcast.Signal, level int64, duration int32, digest []byte, ballotKind engineproto.AttestedMessageKind) {
	t.Helper()
	select {
	case <-time.After(assureDeadline):
		t.Fatal("REDACTED")
	case msg := <-ballotChan:
		ballotEvent, ok := msg.Data().(kinds.EventDataBallot)
		require.True(t, ok, "REDACTED",
			msg.Data())

		ballot := ballotEvent.Ballot
		assert.Equal(t, level, ballot.Level, "REDACTED", level, ballot.Level)
		assert.Equal(t, duration, ballot.Cycle, "REDACTED", duration, ballot.Cycle)
		assert.Equal(t, ballotKind, ballot.Kind, "REDACTED", ballotKind, ballot.Kind)
		if digest == nil {
			require.Nil(t, ballot.LedgerUID.Digest, "REDACTED", ballot.LedgerUID.Digest)
		} else {
			require.True(t, bytes.Equal(ballot.LedgerUID.Digest, digest), "REDACTED", digest, ballot.LedgerUID.Digest)
		}
	}
}

func assurePreendorseDeadline(ch <-chan cometbroadcast.Signal) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case <-ch:
	}
}

func assureNewEventOnConduit(ch <-chan cometbroadcast.Signal) {
	select {
	case <-time.After(assureDeadline * 12 / 10): //
		panic("REDACTED")
	case <-ch:
	}
}

//
//

//
//
func agreementTracer() log.Tracer {
	return log.VerifyingTracerWithHueFn(func(keyvalues ...any) term.FgBgColor {
		for i := 0; i < len(keyvalues)-1; i += 2 {
			if keyvalues[i] == "REDACTED" {
				return term.FgBgColor{Fg: term.Color(uint8(keyvalues[i+1].(int) + 1))}
			}
		}
		return term.FgBgColor{}
	}).With("REDACTED", "REDACTED")
}

func randomAgreementNet(t *testing.T, nRatifiers int, verifyLabel string, timerFunction func() DeadlineTimer,
	applicationFunction func() iface.Software, settingsOpts ...func(*cfg.Settings),
) ([]*Status, sanitizeFunction) {
	t.Helper()
	generatePaper, privateValues := randomOriginPaper(nRatifiers, false, 30, nil)
	css := make([]*Status, nRatifiers)
	tracer := agreementTracer()
	settingsOriginFolders := make([]string, 0, nRatifiers)
	for i := 0; i < nRatifiers; i++ {
		statusStore := dbm.NewMemoryStore() //
		statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
			DropIfaceReplies: false,
		})
		status, _ := statusDepot.ImportFromStoreOrOriginPaper(generatePaper)
		thisSettings := RestoreSettings(fmt.Sprintf("REDACTED", verifyLabel, i))
		settingsOriginFolders = append(settingsOriginFolders, thisSettings.OriginFolder)
		for _, opt := range settingsOpts {
			opt(thisSettings)
		}
		assureFolder(filepath.Dir(thisSettings.Agreement.JournalEntry()), 0o700) //
		app := applicationFunction()
		values := kinds.Tm2schema.RatifierRefreshes(status.Ratifiers)
		_, err := app.InitSeries(context.Background(), &iface.QueryInitSeries{Ratifiers: values})
		require.NoError(t, err)

		css[i] = newStatusWithSettingsAndLedgerDepot(thisSettings, status, privateValues[i], app, statusStore)
		css[i].CollectionDeadlineTimer(timerFunction())
		css[i].AssignTracer(tracer.With("REDACTED", i, "REDACTED", "REDACTED"))
	}
	return css, func() {
		for _, dir := range settingsOriginFolders {
			os.RemoveAll(dir)
		}
	}
}

//
func randomAgreementNetWithNodes(
	t *testing.T,
	nRatifiers,
	nNodes int,
	verifyLabel string,
	timerFunction func() DeadlineTimer,
	applicationFunction func(string) iface.Software,
) ([]*Status, *kinds.OriginPaper, *cfg.Settings, sanitizeFunction) {
	c := verify.AgreementOptions()
	generatePaper, privateValues := randomOriginPaper(nRatifiers, false, verifyMinimumEnergy, c)
	css := make([]*Status, nNodes)
	tracer := agreementTracer()
	var peer0setting *cfg.Settings
	settingsOriginFolders := make([]string, 0, nNodes)
	for i := 0; i < nNodes; i++ {
		statusStore := dbm.NewMemoryStore() //
		statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
			DropIfaceReplies: false,
		})
		t.Cleanup(func() { _ = statusDepot.End() })
		status, _ := statusDepot.ImportFromStoreOrOriginPaper(generatePaper)
		thisSettings := RestoreSettings(fmt.Sprintf("REDACTED", verifyLabel, i))
		settingsOriginFolders = append(settingsOriginFolders, thisSettings.OriginFolder)
		assureFolder(filepath.Dir(thisSettings.Agreement.JournalEntry()), 0o700) //
		if i == 0 {
			peer0setting = thisSettings
		}
		var privateValue kinds.PrivateRatifier
		if i < nRatifiers {
			privateValue = privateValues[i]
		} else {
			temporaryKeyEntry, err := os.CreateTemp("REDACTED", "REDACTED")
			if err != nil {
				panic(err)
			}
			temporaryStatusEntry, err := os.CreateTemp("REDACTED", "REDACTED")
			if err != nil {
				panic(err)
			}

			privateValue = privatekey.GenerateEntryPrivatekey(temporaryKeyEntry.Name(), temporaryStatusEntry.Name())
		}

		app := applicationFunction(path.Join(settings.StoreFolder(), fmt.Sprintf("REDACTED", verifyLabel, i)))
		values := kinds.Tm2schema.RatifierRefreshes(status.Ratifiers)
		if _, ok := app.(*objectdepot.Software); ok {
			//
			status.Release.Agreement.App = objectdepot.ApplicationRelease
		}
		_, err := app.InitSeries(context.Background(), &iface.QueryInitSeries{Ratifiers: values})
		require.NoError(t, err)

		css[i] = newStatusWithSettings(thisSettings, status, privateValue, app)
		css[i].CollectionDeadlineTimer(timerFunction())
		css[i].AssignTracer(tracer.With("REDACTED", i, "REDACTED", "REDACTED"))
	}
	return css, generatePaper, peer0setting, func() {
		for _, dir := range settingsOriginFolders {
			os.RemoveAll(dir)
		}
	}
}

func fetchRouterOrdinal(routers []*p2p.Router, node p2p.Node) int {
	for i, s := range routers {
		if node.MemberDetails().ID() == s.MemberDetails().ID() {
			return i
		}
	}
	panic("REDACTED")
}

//
//

func randomOriginPaper(countRatifiers int,
	randomEnergy bool,
	minimumEnergy int64,
	agreementOptions *kinds.AgreementOptions,
) (*kinds.OriginPaper, []kinds.PrivateRatifier) {
	ratifiers := make([]kinds.OriginRatifier, countRatifiers)
	privateRatifiers := make([]kinds.PrivateRatifier, countRatifiers)
	for i := 0; i < countRatifiers; i++ {
		val, privateValue := kinds.RandomRatifier(randomEnergy, minimumEnergy)
		ratifiers[i] = kinds.OriginRatifier{
			PublicKey: val.PublicKey,
			Energy:  val.PollingEnergy,
		}
		privateRatifiers[i] = privateValue
	}
	sort.Sort(kinds.PrivateRatifiersByLocation(privateRatifiers))

	return &kinds.OriginPaper{
		OriginMoment:     engineclock.Now(),
		PrimaryLevel:   1,
		LedgerUID:         verify.StandardVerifyLedgerUID,
		Ratifiers:      ratifiers,
		AgreementOptions: agreementOptions,
	}, privateRatifiers
}

func randomOriginStatus(
	countRatifiers int,
	randomEnergy bool,
	minimumEnergy int64,
	agreementOptions *kinds.AgreementOptions,
) (sm.Status, []kinds.PrivateRatifier) {
	generatePaper, privateRatifiers := randomOriginPaper(countRatifiers, randomEnergy, minimumEnergy, agreementOptions)
	s0, _ := sm.CreateOriginStatus(generatePaper)
	return s0, privateRatifiers
}

//
//

func newEmulateTimerFunction(solelyOnce bool) func() DeadlineTimer {
	return func() DeadlineTimer {
		return &emulateTimer{
			c:        make(chan deadlineDetails, 10),
			solelyOnce: solelyOnce,
		}
	}
}

//
//
type emulateTimer struct {
	c chan deadlineDetails

	mtx      sync.Mutex
	solelyOnce bool
	triggered    bool
}

func (m *emulateTimer) Begin() error {
	return nil
}

func (m *emulateTimer) Halt() error {
	return nil
}

func (m *emulateTimer) SequenceDeadline(ti deadlineDetails) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.solelyOnce && m.triggered {
		return
	}
	if ti.Phase == cskinds.DurationPhaseNewLevel {
		m.c <- ti
		m.triggered = true
	}
}

func (m *emulateTimer) Chan() <-chan deadlineDetails {
	return m.c
}

func (*emulateTimer) AssignTracer(log.Tracer) {}

func newDurableObjectDepot() iface.Software {
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	return objectdepot.NewDurableSoftware(dir)
}

func newObjectDepot() iface.Software {
	return objectdepot.NewInRamSoftware()
}

func newDurableObjectDepotWithRoute(storeFolder string) iface.Software {
	return objectdepot.NewDurableSoftware(storeFolder)
}

func attestDataIsEquivalent(v1 *kinds.Ballot, v2 *engineproto.Ballot) bool {
	if v1 == nil || v2 == nil {
		return false
	}

	return v1.Kind == v2.Kind &&
		bytes.Equal(v1.LedgerUID.Digest, v2.LedgerUID.FetchDigest()) &&
		v1.Level == v2.FetchLevel() &&
		v1.Cycle == v2.Cycle &&
		bytes.Equal(v1.RatifierLocation.Octets(), v2.FetchRatifierLocation()) &&
		v1.RatifierOrdinal == v2.FetchRatifierOrdinal() &&
		bytes.Equal(v1.Addition, v2.Addition)
}
