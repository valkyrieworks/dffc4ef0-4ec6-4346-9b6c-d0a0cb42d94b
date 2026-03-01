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

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

const (
	verifyListener = "REDACTED"
)

//
//
type sanitizeMethod func()

//
var (
	settings                *cfg.Settings //
	agreementReenactSettings *cfg.Settings
	assureDeadline         = time.Millisecond * 200
)

func assurePath(dir string, style os.FileMode) {
	if err := strongos.AssurePath(dir, style); err != nil {
		panic(err)
	}
}

func RestoreSettings(alias string) *cfg.Settings {
	return verify.RestoreVerifyOrigin(alias)
}

//
//

type assessorMock struct {
	Ordinal  int32 //
	Altitude int64
	Iteration  int32
	kinds.PrivateAssessor
	BallotingPotency int64
	finalBallot    *kinds.Ballot
}

var verifyMinimumPotency int64 = 10

func freshAssessorMock(privateAssessor kinds.PrivateAssessor, itemOrdinal int32) *assessorMock {
	return &assessorMock{
		Ordinal:         itemOrdinal,
		PrivateAssessor: privateAssessor,
		BallotingPotency:   verifyMinimumPotency,
	}
}

func (vs *assessorMock) attestBallot(
	ballotKind commitchema.AttestedSignalKind,
	digest []byte,
	heading kinds.FragmentAssignHeading,
	ballotAddition []byte,
	addnActivated bool,
) (*kinds.Ballot, error) {
	publicToken, err := vs.ObtainPublicToken()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	ballot := &kinds.Ballot{
		Kind:             ballotKind,
		Altitude:           vs.Altitude,
		Iteration:            vs.Iteration,
		LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: heading},
		Timestamp:        committime.Now(),
		AssessorLocation: publicToken.Location(),
		AssessorOrdinal:   vs.Ordinal,
		Addition:        ballotAddition,
	}
	v := ballot.TowardSchema()
	if err = vs.AttestBallot(verify.FallbackVerifySuccessionUUID, v); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	if attestDataEqualsEquivalent(vs.finalBallot, v) {
		v.Notation = vs.finalBallot.Notation
		v.Timestamp = vs.finalBallot.Timestamp
		v.AdditionNotation = vs.finalBallot.AdditionNotation
	}

	ballot.Notation = v.Notation
	ballot.Timestamp = v.Timestamp
	ballot.AdditionNotation = v.AdditionNotation

	if !addnActivated {
		ballot.AdditionNotation = nil
	}

	return ballot, err
}

//
func attestBallot(vs *assessorMock, ballotKind commitchema.AttestedSignalKind, digest []byte, heading kinds.FragmentAssignHeading, addnActivated bool) *kinds.Ballot {
	var ext []byte
	//
	if addnActivated {
		if ballotKind != commitchema.PreendorseKind {
			panic(fmt.Errorf("REDACTED"))
		}
		if len(digest) != 0 || !heading.EqualsNull() {
			ext = []byte("REDACTED")
		}
	}
	v, err := vs.attestBallot(ballotKind, digest, heading, ext, addnActivated)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	vs.finalBallot = v

	return v
}

func attestBallots(
	ballotKind commitchema.AttestedSignalKind,
	digest []byte,
	heading kinds.FragmentAssignHeading,
	addnActivated bool,
	vss ...*assessorMock,
) []*kinds.Ballot {
	ballots := make([]*kinds.Ballot, len(vss))
	for i, vs := range vss {
		ballots[i] = attestBallot(vs, ballotKind, digest, heading, addnActivated)
	}
	return ballots
}

func advanceAltitude(vss ...*assessorMock) {
	for _, vs := range vss {
		vs.Altitude++
	}
}

func advanceIteration(vss ...*assessorMock) {
	for _, vs := range vss {
		vs.Iteration++
	}
}

type AssessorMocksViaPotency []*assessorMock

func (vss AssessorMocksViaPotency) Len() int {
	return len(vss)
}

func (vss AssessorMocksViaPotency) Inferior(i, j int) bool {
	vssindex, err := vss[i].ObtainPublicToken()
	if err != nil {
		panic(err)
	}
	vssjoint, err := vss[j].ObtainPublicToken()
	if err != nil {
		panic(err)
	}

	if vss[i].BallotingPotency == vss[j].BallotingPotency {
		return bytes.Compare(vssindex.Location(), vssjoint.Location()) == -1
	}
	return vss[i].BallotingPotency > vss[j].BallotingPotency
}

func (vss AssessorMocksViaPotency) Exchange(i, j int) {
	it := vss[i]
	vss[i] = vss[j]
	vss[i].Ordinal = int32(i)
	vss[j] = it
	vss[j].Ordinal = int32(j)
}

//
//

func initiateVerifyIteration(cs *Status, altitude int64, iteration int32) {
	cs.joinFreshIteration(altitude, iteration)
	cs.initiateThreads(0)
}

//
func resolveNomination(
	ctx context.Context,
	t *testing.T,
	cs1 *Status,
	vs *assessorMock,
	altitude int64,
	iteration int32,
) (*kinds.Nomination, *kinds.Ledger) {
	cs1.mtx.Lock()
	ledger, err := cs1.generateNominationLedger(ctx)
	require.NoError(t, err)
	ledgerFragments, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	soundIteration := cs1.SoundIteration
	successionUUID := cs1.status.SuccessionUUID
	cs1.mtx.Unlock()
	if ledger == nil {
		panic("REDACTED")
	}

	//
	policyIteration, itemLedgerUUID := soundIteration, kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: ledgerFragments.Heading()}
	nomination := kinds.FreshNomination(altitude, iteration, policyIteration, itemLedgerUUID)
	p := nomination.TowardSchema()
	if err := vs.AttestNomination(successionUUID, p); err != nil {
		panic(err)
	}

	nomination.Notation = p.Notation

	return nomination, ledger
}

func appendBallots(to *Status, ballots ...*kinds.Ballot) {
	for _, ballot := range ballots {
		to.nodeSignalStaging <- signalDetails{Msg: &BallotSignal{ballot}}
	}
}

func attestAppendBallots(
	to *Status,
	ballotKind commitchema.AttestedSignalKind,
	digest []byte,
	heading kinds.FragmentAssignHeading,
	addnActivated bool,
	vss ...*assessorMock,
) {
	ballots := attestBallots(ballotKind, digest, heading, addnActivated, vss...)
	appendBallots(to, ballots...)
}

func certifyPreballot(t *testing.T, cs *Status, iteration int32, privateItem *assessorMock, ledgerDigest []byte) {
	preballots := cs.Ballots.Preballots(iteration)
	publicToken, err := privateItem.ObtainPublicToken()
	require.NoError(t, err)
	location := publicToken.Location()
	var ballot *kinds.Ballot
	if ballot = preballots.ObtainViaLocation(location); ballot == nil {
		panic("REDACTED")
	}
	if ledgerDigest == nil {
		if ballot.LedgerUUID.Digest != nil {
			panic(fmt.Sprintf("REDACTED", ballot.LedgerUUID.Digest))
		}
	} else {
		if !bytes.Equal(ballot.LedgerUUID.Digest, ledgerDigest) {
			panic(fmt.Sprintf("REDACTED", ledgerDigest, ballot.LedgerUUID.Digest))
		}
	}
}

func certifyFinalPreendorse(t *testing.T, cs *Status, privateItem *assessorMock, ledgerDigest []byte) {
	ballots := cs.FinalEndorse
	pv, err := privateItem.ObtainPublicToken()
	require.NoError(t, err)
	location := pv.Location()
	var ballot *kinds.Ballot
	if ballot = ballots.ObtainViaLocation(location); ballot == nil {
		panic("REDACTED")
	}
	if !bytes.Equal(ballot.LedgerUUID.Digest, ledgerDigest) {
		panic(fmt.Sprintf("REDACTED", ledgerDigest, ballot.LedgerUUID.Digest))
	}
}

func certifyPreendorse(
	t *testing.T,
	cs *Status,
	thatIteration,
	secureIteration int32,
	privateItem *assessorMock,
	ballotedLedgerDigest,
	securedLedgerDigest []byte,
) {
	preendorsements := cs.Ballots.Preendorsements(thatIteration)
	pv, err := privateItem.ObtainPublicToken()
	require.NoError(t, err)
	location := pv.Location()
	var ballot *kinds.Ballot
	if ballot = preendorsements.ObtainViaLocation(location); ballot == nil {
		panic("REDACTED")
	}

	if ballotedLedgerDigest == nil {
		if ballot.LedgerUUID.Digest != nil {
			panic("REDACTED")
		}
	} else {
		if !bytes.Equal(ballot.LedgerUUID.Digest, ballotedLedgerDigest) {
			panic("REDACTED")
		}
	}

	rs := cs.ObtainIterationStatus()
	if securedLedgerDigest == nil {
		if rs.SecuredIteration != secureIteration || rs.SecuredLedger != nil {
			panic(fmt.Sprintf(
				"REDACTED",
				secureIteration,
				rs.SecuredIteration,
				rs.SecuredLedger))
		}
	} else {
		if rs.SecuredIteration != secureIteration || !bytes.Equal(rs.SecuredLedger.Digest(), securedLedgerDigest) {
			panic(fmt.Sprintf(
				"REDACTED",
				secureIteration,
				rs.SecuredIteration,
				rs.SecuredLedger.Digest(),
				securedLedgerDigest))
		}
	}
}

func listenTowardBalloter(cs *Status, location []byte) <-chan tendermintpubsub.Signal {
	ballotsUnder, err := cs.incidentChannel.ListenUncached(context.Background(), verifyListener, kinds.IncidentInquireBallot)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", verifyListener, kinds.IncidentInquireBallot))
	}
	ch := make(chan tendermintpubsub.Signal)
	go func() {
		for msg := range ballotsUnder.Out() {
			ballot := msg.Data().(kinds.IncidentDataBallot)
			//
			if bytes.Equal(location, ballot.Ballot.AssessorLocation) {
				ch <- msg
			}
		}
	}()
	return ch
}

//
//

func freshStatus(status sm.Status, pv kinds.PrivateAssessor, app iface.Platform) *Status {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	return freshStatusUsingSettings(settings, status, pv, app)
}

func freshStatusUsingSettings(
	thatSettings *cfg.Settings,
	status sm.Status,
	pv kinds.PrivateAssessor,
	app iface.Platform,
) *Status {
	ledgerDatastore := dbm.FreshMemoryDatastore()
	return freshStatusUsingSettingsAlsoLedgerDepot(thatSettings, status, pv, app, ledgerDatastore)
}

func freshStatusUsingSettingsAlsoLedgerDepot(
	thatSettings *cfg.Settings,
	status sm.Status,
	pv kinds.PrivateAssessor,
	app iface.Platform,
	ledgerDatastore dbm.DB,
) *Status {
	//
	ledgerDepot := depot.FreshLedgerDepot(ledgerDatastore)

	//
	mtx := new(commitchronize.Exclusion)

	delegateApplicationLinkConnection := delegate.FreshApplicationLinkAgreement(abcicustomer.FreshRegionalCustomer(mtx, app), delegate.NooperationTelemetry())
	delegateApplicationLinkMemory := delegate.FreshApplicationLinkTxpool(abcicustomer.FreshRegionalCustomer(mtx, app), delegate.NooperationTelemetry())
	//
	txpoollTelemetry := txpooll.NooperationTelemetry()

	//
	txpool := txpooll.FreshCNCatalogTxpool(settings.Txpool,
		delegateApplicationLinkMemory,
		status.FinalLedgerAltitude,
		txpooll.UsingTelemetry(txpoollTelemetry),
		txpooll.UsingPriorInspect(sm.TransferPriorInspect(status)),
		txpooll.UsingRelayInspect(sm.TransferRelayInspect(status)))

	if thatSettings.Agreement.PauseForeachTrans() {
		txpool.ActivateTransAccessible()
	}

	incidentpool := sm.VoidProofHub{}

	//
	statusDatastore := ledgerDatastore
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	if err := statusDepot.Persist(status); err != nil { //
		panic(err)
	}

	ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegateApplicationLinkConnection, txpool, incidentpool, ledgerDepot)
	cs := FreshStatus(thatSettings.Agreement, status, ledgerExecute, ledgerDepot, txpool, incidentpool)
	cs.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	cs.AssignPrivateAssessor(pv)

	incidentChannel := kinds.FreshIncidentPipeline()
	incidentChannel.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	err := incidentChannel.Initiate()
	if err != nil {
		panic(err)
	}
	cs.AssignIncidentChannel(incidentChannel)
	return cs
}

func fetchPrivateAssessor(settings *cfg.Settings) *privatevalue.RecordPRV {
	privateAssessorTokenRecord := settings.PrivateAssessorTokenRecord()
	assurePath(filepath.Dir(privateAssessorTokenRecord), 0o700)
	privateAssessorStatusRecord := settings.PrivateAssessorStatusRecord()
	privateAssessor := privatevalue.FetchEitherProduceRecordPRV(privateAssessorTokenRecord, privateAssessorStatusRecord)
	privateAssessor.Restore()
	return privateAssessor
}

func arbitraryStatus(nthAssessors int) (*Status, []*assessorMock) {
	return arbitraryStatusUsingApplication(nthAssessors, statedepot.FreshInsideRamPlatform())
}

func arbitraryStatusUsingApplicationUsingAltitude(
	nthAssessors int,
	app iface.Platform,
	altitude int64,
) (*Status, []*assessorMock) {
	c := verify.AgreementSettings()
	c.Iface.BallotAdditionsActivateAltitude = altitude
	return arbitraryStatusUsingApplicationImplementation(nthAssessors, app, c)
}

func arbitraryStatusUsingApplication(nthAssessors int, app iface.Platform) (*Status, []*assessorMock) {
	c := verify.AgreementSettings()
	return arbitraryStatusUsingApplicationImplementation(nthAssessors, app, c)
}

func arbitraryStatusUsingApplicationImplementation(
	nthAssessors int,
	app iface.Platform,
	agreementParameters *kinds.AgreementSettings,
) (*Status, []*assessorMock) {
	//
	status, privateItems := arbitraryInaugurationStatus(nthAssessors, false, 10, agreementParameters)

	vss := make([]*assessorMock, nthAssessors)

	cs := freshStatus(status, privateItems[0], app)

	for i := 0; i < nthAssessors; i++ {
		vss[i] = freshAssessorMock(privateItems[i], int32(i))
	}
	//
	advanceAltitude(vss[1:]...)

	return cs, vss
}

//

func assureNegativeFreshIncident(ch <-chan tendermintpubsub.Signal, deadline time.Duration,
	failureSignal string,
) {
	select {
	case <-time.After(deadline):
		break
	case <-ch:
		panic(failureSignal)
	}
}

func assureNegativeFreshIncidentUponConduit(ch <-chan tendermintpubsub.Signal) {
	assureNegativeFreshIncident(
		ch,
		assureDeadline*8/10, //
		"REDACTED")
}

func assureNegativeFreshIterationPhase(phaseChnl <-chan tendermintpubsub.Signal) {
	assureNegativeFreshIncident(
		phaseChnl,
		assureDeadline,
		"REDACTED")
}

func assureNegativeFreshRelease(releaseChnl <-chan tendermintpubsub.Signal) {
	assureNegativeFreshIncident(
		releaseChnl,
		assureDeadline,
		"REDACTED")
}

func assureNegativeFreshDeadline(phaseChnl <-chan tendermintpubsub.Signal, deadline int64) {
	deadlineInterval := time.Duration(deadline*10) * time.Nanosecond
	assureNegativeFreshIncident(
		phaseChnl,
		deadlineInterval,
		"REDACTED")
}

func assureFreshIncident(ch <-chan tendermintpubsub.Signal, altitude int64, iteration int32, deadline time.Duration, failureSignal string) {
	select {
	case <-time.After(deadline):
		panic(failureSignal)
	case msg := <-ch:
		iterationStatusIncident, ok := msg.Data().(kinds.IncidentDataIterationStatus)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if iterationStatusIncident.Altitude != altitude {
			panic(fmt.Sprintf("REDACTED", altitude, iterationStatusIncident.Altitude))
		}
		if iterationStatusIncident.Iteration != iteration {
			panic(fmt.Sprintf("REDACTED", iteration, iterationStatusIncident.Iteration))
		}
		//
	}
}

func assureFreshIteration(iterationChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-iterationChnl:
		freshIterationIncident, ok := msg.Data().(kinds.IncidentDataFreshIteration)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if freshIterationIncident.Altitude != altitude {
			panic(fmt.Sprintf("REDACTED", altitude, freshIterationIncident.Altitude))
		}
		if freshIterationIncident.Iteration != iteration {
			panic(fmt.Sprintf("REDACTED", iteration, freshIterationIncident.Iteration))
		}
	}
}

func assureFreshDeadline(deadlineChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32, deadline int64) {
	deadlineInterval := time.Duration(deadline*10) * time.Nanosecond
	assureFreshIncident(deadlineChnl, altitude, iteration, deadlineInterval,
		"REDACTED")
}

func assureFreshNomination(nominationChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-nominationChnl:
		nominationIncident, ok := msg.Data().(kinds.IncidentDataFinishNomination)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if nominationIncident.Altitude != altitude {
			panic(fmt.Sprintf("REDACTED", altitude, nominationIncident.Altitude))
		}
		if nominationIncident.Iteration != iteration {
			panic(fmt.Sprintf("REDACTED", iteration, nominationIncident.Iteration))
		}
	}
}

func assureFreshSoundLedger(soundLedgerChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32) {
	assureFreshIncident(soundLedgerChnl, altitude, iteration, assureDeadline,
		"REDACTED")
}

func assureFreshLedger(ledgerChnl <-chan tendermintpubsub.Signal, altitude int64) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-ledgerChnl:
		ledgerIncident, ok := msg.Data().(kinds.IncidentDataFreshLedger)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if ledgerIncident.Ledger.Altitude != altitude {
			panic(fmt.Sprintf("REDACTED", altitude, ledgerIncident.Ledger.Altitude))
		}
	}
}

func assureFreshLedgerHeading(ledgerChnl <-chan tendermintpubsub.Signal, altitude int64, ledgerDigest tendermintoctets.HexadecimalOctets) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-ledgerChnl:
		ledgerHeadingIncident, ok := msg.Data().(kinds.IncidentDataFreshLedgerHeading)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if ledgerHeadingIncident.Heading.Altitude != altitude {
			panic(fmt.Sprintf("REDACTED", altitude, ledgerHeadingIncident.Heading.Altitude))
		}
		if !bytes.Equal(ledgerHeadingIncident.Heading.Digest(), ledgerDigest) {
			panic(fmt.Sprintf("REDACTED", ledgerDigest, ledgerHeadingIncident.Heading.Digest()))
		}
	}
}

func assureFreshRelease(releaseChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32) {
	assureFreshIncident(releaseChnl, altitude, iteration, assureDeadline,
		"REDACTED")
}

func assureNomination(nominationChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32, itemUUID kinds.LedgerUUID) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-nominationChnl:
		nominationIncident, ok := msg.Data().(kinds.IncidentDataFinishNomination)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		if nominationIncident.Altitude != altitude {
			panic(fmt.Sprintf("REDACTED", altitude, nominationIncident.Altitude))
		}
		if nominationIncident.Iteration != iteration {
			panic(fmt.Sprintf("REDACTED", iteration, nominationIncident.Iteration))
		}
		if !nominationIncident.LedgerUUID.Matches(itemUUID) {
			panic(fmt.Sprintf("REDACTED", nominationIncident.LedgerUUID, itemUUID))
		}
	}
}

func assurePreendorse(ballotChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32) {
	assureBallot(ballotChnl, altitude, iteration, commitchema.PreendorseKind)
}

func assurePreballot(ballotChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32) {
	assureBallot(ballotChnl, altitude, iteration, commitchema.PreballotKind)
}

func assureBallot(ballotChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32,
	ballotKind commitchema.AttestedSignalKind,
) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case msg := <-ballotChnl:
		ballotIncident, ok := msg.Data().(kinds.IncidentDataBallot)
		if !ok {
			panic(fmt.Sprintf("REDACTED",
				msg.Data()))
		}
		ballot := ballotIncident.Ballot
		if ballot.Altitude != altitude {
			panic(fmt.Sprintf("REDACTED", altitude, ballot.Altitude))
		}
		if ballot.Iteration != iteration {
			panic(fmt.Sprintf("REDACTED", iteration, ballot.Iteration))
		}
		if ballot.Kind != ballotKind {
			panic(fmt.Sprintf("REDACTED", ballotKind, ballot.Kind))
		}
	}
}

func assurePreballotAlign(t *testing.T, ballotChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32, digest []byte) {
	t.Helper()
	assureBallotAlign(t, ballotChnl, altitude, iteration, digest, commitchema.PreballotKind)
}

func assurePreendorseAlign(t *testing.T, ballotChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32, digest []byte) {
	t.Helper()
	assureBallotAlign(t, ballotChnl, altitude, iteration, digest, commitchema.PreendorseKind)
}

func assureBallotAlign(t *testing.T, ballotChnl <-chan tendermintpubsub.Signal, altitude int64, iteration int32, digest []byte, ballotKind commitchema.AttestedSignalKind) {
	t.Helper()
	select {
	case <-time.After(assureDeadline):
		t.Fatal("REDACTED")
	case msg := <-ballotChnl:
		ballotIncident, ok := msg.Data().(kinds.IncidentDataBallot)
		require.True(t, ok, "REDACTED",
			msg.Data())

		ballot := ballotIncident.Ballot
		assert.Equal(t, altitude, ballot.Altitude, "REDACTED", altitude, ballot.Altitude)
		assert.Equal(t, iteration, ballot.Iteration, "REDACTED", iteration, ballot.Iteration)
		assert.Equal(t, ballotKind, ballot.Kind, "REDACTED", ballotKind, ballot.Kind)
		if digest == nil {
			require.Nil(t, ballot.LedgerUUID.Digest, "REDACTED", ballot.LedgerUUID.Digest)
		} else {
			require.True(t, bytes.Equal(ballot.LedgerUUID.Digest, digest), "REDACTED", digest, ballot.LedgerUUID.Digest)
		}
	}
}

func assurePreendorseDeadline(ch <-chan tendermintpubsub.Signal) {
	select {
	case <-time.After(assureDeadline):
		panic("REDACTED")
	case <-ch:
	}
}

func assureFreshIncidentUponConduit(ch <-chan tendermintpubsub.Signal) {
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
	return log.VerifyingTracerUsingHueProc(func(tokvals ...any) term.FgBgColor {
		for i := 0; i < len(tokvals)-1; i += 2 {
			if tokvals[i] == "REDACTED" {
				return term.FgBgColor{Fg: term.Color(uint8(tokvals[i+1].(int) + 1))}
			}
		}
		return term.FgBgColor{}
	}).Using("REDACTED", "REDACTED")
}

func arbitraryAgreementNetwork(t *testing.T, nthAssessors int, verifyAlias string, metronomeMethod func() DeadlineMetronome,
	applicationMethod func() iface.Platform, settingsOptions ...func(*cfg.Settings),
) ([]*Status, sanitizeMethod) {
	t.Helper()
	producePaper, privateItems := arbitraryOriginPaper(nthAssessors, false, 30, nil)
	css := make([]*Status, nthAssessors)
	tracer := agreementTracer()
	settingsOriginFolders := make([]string, 0, nthAssessors)
	for i := 0; i < nthAssessors; i++ {
		statusDatastore := dbm.FreshMemoryDatastore() //
		statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
			EjectIfaceReplies: false,
		})
		status, _ := statusDepot.FetchOriginatingDatastoreEitherOriginPaper(producePaper)
		thatSettings := RestoreSettings(fmt.Sprintf("REDACTED", verifyAlias, i))
		settingsOriginFolders = append(settingsOriginFolders, thatSettings.OriginPath)
		for _, opt := range settingsOptions {
			opt(thatSettings)
		}
		assurePath(filepath.Dir(thatSettings.Agreement.JournalRecord()), 0o700) //
		app := applicationMethod()
		values := kinds.Temp2buffer.AssessorRevisions(status.Assessors)
		_, err := app.InitializeSuccession(context.Background(), &iface.SolicitInitializeSuccession{Assessors: values})
		require.NoError(t, err)

		css[i] = freshStatusUsingSettingsAlsoLedgerDepot(thatSettings, status, privateItems[i], app, statusDatastore)
		css[i].AssignDeadlineMetronome(metronomeMethod())
		css[i].AssignTracer(tracer.Using("REDACTED", i, "REDACTED", "REDACTED"))
	}
	return css, func() {
		for _, dir := range settingsOriginFolders {
			os.RemoveAll(dir)
		}
	}
}

//
func arbitraryAgreementNetworkUsingNodes(
	t *testing.T,
	nthAssessors,
	nthNodes int,
	verifyAlias string,
	metronomeMethod func() DeadlineMetronome,
	applicationMethod func(string) iface.Platform,
) ([]*Status, *kinds.OriginPaper, *cfg.Settings, sanitizeMethod) {
	c := verify.AgreementSettings()
	producePaper, privateItems := arbitraryOriginPaper(nthAssessors, false, verifyMinimumPotency, c)
	css := make([]*Status, nthNodes)
	tracer := agreementTracer()
	var node0settings *cfg.Settings
	settingsOriginFolders := make([]string, 0, nthNodes)
	for i := 0; i < nthNodes; i++ {
		statusDatastore := dbm.FreshMemoryDatastore() //
		statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
			EjectIfaceReplies: false,
		})
		t.Cleanup(func() { _ = statusDepot.Shutdown() })
		status, _ := statusDepot.FetchOriginatingDatastoreEitherOriginPaper(producePaper)
		thatSettings := RestoreSettings(fmt.Sprintf("REDACTED", verifyAlias, i))
		settingsOriginFolders = append(settingsOriginFolders, thatSettings.OriginPath)
		assurePath(filepath.Dir(thatSettings.Agreement.JournalRecord()), 0o700) //
		if i == 0 {
			node0settings = thatSettings
		}
		var privateItem kinds.PrivateAssessor
		if i < nthAssessors {
			privateItem = privateItems[i]
		} else {
			transientTokenRecord, err := os.CreateTemp("REDACTED", "REDACTED")
			if err != nil {
				panic(err)
			}
			transientStatusRecord, err := os.CreateTemp("REDACTED", "REDACTED")
			if err != nil {
				panic(err)
			}

			privateItem = privatevalue.ProduceRecordPRV(transientTokenRecord.Name(), transientStatusRecord.Name())
		}

		app := applicationMethod(path.Join(settings.DatastorePath(), fmt.Sprintf("REDACTED", verifyAlias, i)))
		values := kinds.Temp2buffer.AssessorRevisions(status.Assessors)
		if _, ok := app.(*statedepot.Platform); ok {
			//
			status.Edition.Agreement.App = statedepot.PlatformEdition
		}
		_, err := app.InitializeSuccession(context.Background(), &iface.SolicitInitializeSuccession{Assessors: values})
		require.NoError(t, err)

		css[i] = freshStatusUsingSettings(thatSettings, status, privateItem, app)
		css[i].AssignDeadlineMetronome(metronomeMethod())
		css[i].AssignTracer(tracer.Using("REDACTED", i, "REDACTED", "REDACTED"))
	}
	return css, producePaper, node0settings, func() {
		for _, dir := range settingsOriginFolders {
			os.RemoveAll(dir)
		}
	}
}

func obtainRouterOrdinal(routers []*p2p.Router, node p2p.Node) int {
	for i, s := range routers {
		if node.PeerDetails().ID() == s.PeerDetails().ID() {
			return i
		}
	}
	panic("REDACTED")
}

//
//

func arbitraryOriginPaper(countAssessors int,
	arbitraryPotency bool,
	minimumPotency int64,
	agreementParameters *kinds.AgreementSettings,
) (*kinds.OriginPaper, []kinds.PrivateAssessor) {
	assessors := make([]kinds.OriginAssessor, countAssessors)
	privateAssessors := make([]kinds.PrivateAssessor, countAssessors)
	for i := 0; i < countAssessors; i++ {
		val, privateItem := kinds.ArbitraryAssessor(arbitraryPotency, minimumPotency)
		assessors[i] = kinds.OriginAssessor{
			PublicToken: val.PublicToken,
			Potency:  val.BallotingPotency,
		}
		privateAssessors[i] = privateItem
	}
	sort.Sort(kinds.PrivateAssessorsViaLocation(privateAssessors))

	return &kinds.OriginPaper{
		OriginMoment:     committime.Now(),
		PrimaryAltitude:   1,
		SuccessionUUID:         verify.FallbackVerifySuccessionUUID,
		Assessors:      assessors,
		AgreementSettings: agreementParameters,
	}, privateAssessors
}

func arbitraryInaugurationStatus(
	countAssessors int,
	arbitraryPotency bool,
	minimumPotency int64,
	agreementParameters *kinds.AgreementSettings,
) (sm.Status, []kinds.PrivateAssessor) {
	producePaper, privateAssessors := arbitraryOriginPaper(countAssessors, arbitraryPotency, minimumPotency, agreementParameters)
	s0, _ := sm.CreateInaugurationStatus(producePaper)
	return s0, privateAssessors
}

//
//

func freshSimulateMetronomeMethod(solelyOnetime bool) func() DeadlineMetronome {
	return func() DeadlineMetronome {
		return &simulateMetronome{
			c:        make(chan deadlineDetails, 10),
			solelyOnetime: solelyOnetime,
		}
	}
}

//
//
type simulateMetronome struct {
	c chan deadlineDetails

	mtx      sync.Mutex
	solelyOnetime bool
	relayed    bool
}

func (m *simulateMetronome) Initiate() error {
	return nil
}

func (m *simulateMetronome) Halt() error {
	return nil
}

func (m *simulateMetronome) TimelineDeadline(ti deadlineDetails) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.solelyOnetime && m.relayed {
		return
	}
	if ti.Phase == controlkinds.IterationPhaseFreshAltitude {
		m.c <- ti
		m.relayed = true
	}
}

func (m *simulateMetronome) Channel() <-chan deadlineDetails {
	return m.c
}

func (*simulateMetronome) AssignTracer(log.Tracer) {}

func freshEnduringTokvalDepot() iface.Platform {
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	return statedepot.FreshEnduringPlatform(dir)
}

func freshTokvalDepot() iface.Platform {
	return statedepot.FreshInsideRamPlatform()
}

func freshEnduringTokvalDepotUsingRoute(datastorePath string) iface.Platform {
	return statedepot.FreshEnduringPlatform(datastorePath)
}

func attestDataEqualsEquivalent(v1 *kinds.Ballot, v2 *commitchema.Ballot) bool {
	if v1 == nil || v2 == nil {
		return false
	}

	return v1.Kind == v2.Kind &&
		bytes.Equal(v1.LedgerUUID.Digest, v2.LedgerUUID.ObtainDigest()) &&
		v1.Altitude == v2.ObtainAltitude() &&
		v1.Iteration == v2.Iteration &&
		bytes.Equal(v1.AssessorLocation.Octets(), v2.ObtainAssessorLocation()) &&
		v1.AssessorOrdinal == v2.ObtainAssessorOrdinal() &&
		bytes.Equal(v1.Addition, v2.Addition)
}
