package agreement

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	abciemulators "github.com/valkyrieworks/iface/kinds/simulations"
	cskinds "github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/intrinsic/verify"
	cometbytes "github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/protoio"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	engineseed "github.com/valkyrieworks/utils/random"
	p2pemulator "github.com/valkyrieworks/p2p/emulate"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

/**

e
0
+
d
)
l
d
e
d
l
d
e
.
a
l
d
d
d
d
e
d
d
e
e
e
t

*/

//
//

func VerifyStatusRecommenderChoice0(t *testing.T) {
	cs1, vss := randomStatus(4)
	level, duration := cs1.Level, cs1.Cycle

	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)

	beginVerifyEpoch(cs1, level, duration)

	//
	assureNewEpoch(newEpochChan, level, duration)

	//
	nomination := cs1.FetchDurationStatus().Ratifiers.FetchRecommender()
	pv, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	location := pv.Location()
	if !bytes.Equal(nomination.Location, location) {
		t.Fatalf("REDACTED", 0, nomination.Location)
	}

	//
	assureNewNomination(nominationChan, level, duration)

	rs := cs1.FetchDurationStatus()
	attestAppendBallots(cs1, engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), true, vss[1:]...)

	//
	assureNewEpoch(newEpochChan, level+1, 0)

	nomination = cs1.FetchDurationStatus().Ratifiers.FetchRecommender()
	pv1, err := vss[1].FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	if !bytes.Equal(nomination.Location, address) {
		panic(fmt.Sprintf("REDACTED", 1, nomination.Location))
	}
}

//
func VerifyStatusRecommenderChoice2(t *testing.T) {
	cs1, vss := randomStatus(4) //
	level := cs1.Level
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)

	//
	augmentEpoch(vss[1:]...)
	augmentEpoch(vss[1:]...)

	var duration int32 = 2
	beginVerifyEpoch(cs1, level, duration)

	assureNewEpoch(newEpochChan, level, duration) //

	//
	for i := int32(0); int(i) < len(vss); i++ {
		nomination := cs1.FetchDurationStatus().Ratifiers.FetchRecommender()
		pvk, err := vss[int(i+duration)%len(vss)].FetchPublicKey()
		require.NoError(t, err)
		address := pvk.Location()
		accurateRecommender := address
		if !bytes.Equal(nomination.Location, accurateRecommender) {
			panic(fmt.Sprintf(
				"REDACTED",
				int(i+2)%len(vss),
				nomination.Location))
		}

		rs := cs1.FetchDurationStatus()
		attestAppendBallots(cs1, engineproto.PreendorseKind, nil, rs.NominationLedgerSegments.Heading(), true, vss[1:]...)
		assureNewEpoch(newEpochChan, level, i+duration+1) //
		augmentEpoch(vss[1:]...)
	}
}

//
func VerifyStatusJoinNominateNoPrivateRatifier(t *testing.T) {
	cs, _ := randomStatus(1)
	cs.CollectionPrivateRatifier(nil)
	level, duration := cs.Level, cs.Cycle

	//
	deadlineChan := enrol(cs.eventBus, kinds.EventInquireDeadlineNominate)

	beginVerifyEpoch(cs, level, duration)

	//
	assureNewDeadline(deadlineChan, level, duration, cs.settings.DeadlineNominate.Nanoseconds())

	if cs.FetchDurationStatus().Nomination != nil {
		t.Error("REDACTED")
	}
}

//
func VerifyStatusJoinNominateYesPrivateRatifier(t *testing.T) {
	cs, _ := randomStatus(1)
	level, duration := cs.Level, cs.Cycle

	//

	deadlineChan := enrol(cs.eventBus, kinds.EventInquireDeadlineNominate)
	nominationChan := enrol(cs.eventBus, kinds.EventInquireFinishedNomination)

	cs.joinNewEpoch(level, duration)
	cs.beginProcedures(3)

	assureNewNomination(nominationChan, level, duration)

	//
	rs := cs.FetchDurationStatus()
	if rs.Nomination == nil {
		t.Error("REDACTED")
	}
	if rs.NominationLedger == nil {
		t.Error("REDACTED")
	}
	if rs.NominationLedgerSegments.Sum() == 0 {
		t.Error("REDACTED")
	}

	//
	assureNoNewDeadline(deadlineChan, cs.settings.DeadlineNominate.Nanoseconds())
}

func VerifyStatusFlawedNomination(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(2)
	level, duration := cs1.Level, cs1.Cycle
	vs2 := vss[1]

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	ballotChan := enrol(cs1.eventBus, kinds.EventInquireBallot)

	nominationLedger, err := cs1.instantiateNominationLedger(ctx) //
	require.NoError(t, err)

	//
	duration++
	augmentEpoch(vss[1:]...)

	//
	statusDigest := nominationLedger.ApplicationDigest
	if len(statusDigest) == 0 {
		statusDigest = make([]byte, 32)
	}
	statusDigest[0] = (statusDigest[0] + 1) % 255
	nominationLedger.ApplicationDigest = statusDigest
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: nominationLedger.Digest(), SegmentAssignHeading: nominationLedgerSegments.Heading()}
	nomination := kinds.NewNomination(vs2.Level, duration, -1, ledgerUID)
	p := nomination.ToSchema()
	if err := vs2.AttestNomination(cs1.status.LedgerUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}

	nomination.Autograph = p.Autograph

	//
	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	//
	beginVerifyEpoch(cs1, level, duration)

	//
	assureNomination(nominationChan, level, duration, ledgerUID)

	//
	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nil)

	//
	bps, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedger.Digest(), bps.Heading(), false, vs2)
	assurePreballot(ballotChan, level, duration)

	//
	assurePreendorse(ballotChan, level, duration)
	certifyPreendorse(t, cs1, duration, -1, vss[0], nil, nil)

	bps2, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedger.Digest(), bps2.Heading(), true, vs2)
}

func VerifyStatusExcessiveLedger(t *testing.T) {
	const maximumOctets = int64(kinds.LedgerSegmentVolumeOctets)

	for _, verifyInstance := range []struct {
		label      string
		excessive bool
	}{
		{
			label:      "REDACTED",
			excessive: false,
		},
		{
			label:      "REDACTED",
			excessive: true,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			cs1, vss := randomStatus(2)
			cs1.status.AgreementOptions.Ledger.MaximumOctets = maximumOctets
			level, duration := cs1.Level, cs1.Cycle
			vs2 := vss[1]

			segmentVolume := kinds.LedgerSegmentVolumeOctets

			nominationLedger, nominationLedgerSegments := locateLedgerVolumeCeiling(t, level, maximumOctets, cs1, segmentVolume, verifyInstance.excessive)

			deadlineNominateChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
			ballotChan := enrol(cs1.eventBus, kinds.EventInquireBallot)

			//
			duration++
			augmentEpoch(vss[1:]...)

			ledgerUID := kinds.LedgerUID{Digest: nominationLedger.Digest(), SegmentAssignHeading: nominationLedgerSegments.Heading()}
			nomination := kinds.NewNomination(level, duration, -1, ledgerUID)
			p := nomination.ToSchema()
			if err := vs2.AttestNomination(cs1.status.LedgerUID, p); err != nil {
				t.Fatal("REDACTED", err)
			}
			nomination.Autograph = p.Autograph

			sumOctets := 0
			for i := 0; i < int(nominationLedgerSegments.Sum()); i++ {
				segment := nominationLedgerSegments.FetchSegment(i)
				sumOctets += len(segment.Octets)
			}

			maximumLedgerSegments := maximumOctets / int64(kinds.LedgerSegmentVolumeOctets)
			if maximumOctets > maximumLedgerSegments*int64(kinds.LedgerSegmentVolumeOctets) {
				maximumLedgerSegments++
			}
			countLedgerSegments := int64(nominationLedgerSegments.Sum())

			if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
				t.Fatal(err)
			}

			//
			beginVerifyEpoch(cs1, level, duration)

			t.Log("REDACTED", "REDACTED", maximumOctets, "REDACTED", sumOctets)
			t.Log("REDACTED", "REDACTED", maximumLedgerSegments, "REDACTED", countLedgerSegments)

			certifyDigest := nominationLedger.Digest()
			latchedEpoch := int32(1)
			if verifyInstance.excessive {
				certifyDigest = nil
				latchedEpoch = -1
				//
				//
				assureNewDeadline(deadlineNominateChan, level, duration, cs1.settings.Nominate(duration).Nanoseconds())
				//
				//
			}
			assurePreballot(ballotChan, level, duration)
			certifyPreballot(t, cs1, duration, vss[0], certifyDigest)

			//
			if countLedgerSegments > maximumLedgerSegments {
				require.Nil(t, cs1.Nomination)
			}

			bps, err := nominationLedger.CreateSegmentAssign(segmentVolume)
			require.NoError(t, err)

			attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedger.Digest(), bps.Heading(), false, vs2)
			assurePreballot(ballotChan, level, duration)
			assurePreendorse(ballotChan, level, duration)
			certifyPreendorse(t, cs1, duration, latchedEpoch, vss[0], certifyDigest, certifyDigest)

			bps2, err := nominationLedger.CreateSegmentAssign(segmentVolume)
			require.NoError(t, err)
			attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedger.Digest(), bps2.Heading(), true, vs2)
		})
	}
}

//
//

//
func VerifyStatusCompleteEpoch1(t *testing.T) {
	cs, vss := randomStatus(1)
	level, duration := cs.Level, cs.Cycle

	//
	//
	if err := cs.eventBus.Halt(); err != nil {
		t.Error(err)
	}
	eventBus := kinds.NewEventBusWithBufferVolume(0)
	eventBus.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	cs.AssignEventBus(eventBus)
	if err := eventBus.Begin(); err != nil {
		t.Error(err)
	}

	ballotChan := enrolUnCached(cs.eventBus, kinds.EventInquireBallot)
	nominationChan := enrol(cs.eventBus, kinds.EventInquireFinishedNomination)
	newEpochChan := enrol(cs.eventBus, kinds.EventInquireNewEpoch)

	//
	beginVerifyEpoch(cs, level, duration)

	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	nominationLedgerDigest := cs.FetchDurationStatus().NominationLedger.Digest()

	assurePreballot(ballotChan, level, duration) //
	certifyPreballot(t, cs, duration, vss[0], nominationLedgerDigest)

	assurePreendorse(ballotChan, level, duration) //

	//
	assureNewEpoch(newEpochChan, level+1, 0)

	certifyFinalPreendorse(t, cs, vss[0], nominationLedgerDigest)
}

//
func VerifyStatusCompleteEpochNull(t *testing.T) {
	cs, _ := randomStatus(1)
	level, duration := cs.Level, cs.Cycle

	ballotChan := enrolUnCached(cs.eventBus, kinds.EventInquireBallot)

	cs.joinPreballot(level, duration)
	cs.beginProcedures(4)

	assurePreballotAlign(t, ballotChan, level, duration, nil)   //
	assurePreendorseAlign(t, ballotChan, level, duration, nil) //
}

//
//
func VerifyStatusCompleteEpoch2(t *testing.T) {
	cs1, vss := randomStatus(2)
	vs2 := vss[1]
	level, duration := cs1.Level, cs1.Cycle

	ballotChan := enrolUnCached(cs1.eventBus, kinds.EventInquireBallot)
	newLedgerChan := enrol(cs1.eventBus, kinds.EventInquireNewLedger)

	//
	beginVerifyEpoch(cs1, level, duration)

	assurePreballot(ballotChan, level, duration) //

	//
	rs := cs1.FetchDurationStatus()
	nominationLedgerDigest, nominationSegmentCollectionHeading := rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading()

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest, nominationSegmentCollectionHeading, false, vs2)
	assurePreballot(ballotChan, level, duration) //

	assurePreendorse(ballotChan, level, duration) //
	//
	certifyPreendorse(t, cs1, 0, 0, vss[0], nominationLedgerDigest, nominationLedgerDigest)

	//

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedgerDigest, nominationSegmentCollectionHeading, true, vs2)
	assurePreendorse(ballotChan, level, duration)

	//
	assureNewLedger(newLedgerChan, level)
}

//
//

//
//
func VerifyStatusSecureNoPOL(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(2)
	vs2 := vss[1]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	deadlineNominateChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	ballotChan := enrolUnCached(cs1.eventBus, kinds.EventInquireBallot)
	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)

	/**
2
*/

	//
	cs1.joinNewEpoch(level, duration)
	cs1.beginProcedures(0)

	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	epochStatus := cs1.FetchDurationStatus()
	theLedgerDigest := epochStatus.NominationLedger.Digest()
	theSegmentCollectionHeading := epochStatus.NominationLedgerSegments.Heading()

	assurePreballot(ballotChan, level, duration) //

	//
	//
	attestAppendBallots(cs1, engineproto.PreballotKind, theLedgerDigest, theSegmentCollectionHeading, false, vs2)
	assurePreballot(ballotChan, level, duration) //

	assurePreendorse(ballotChan, level, duration) //
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], theLedgerDigest, theLedgerDigest)

	//
	//
	digest := make([]byte, len(theLedgerDigest))
	copy(digest, theLedgerDigest)
	digest[0] = (digest[0] + 1) % 255
	attestAppendBallots(cs1, engineproto.PreendorseKind, digest, theSegmentCollectionHeading, true, vs2)
	assurePreendorse(ballotChan, level, duration) //

	//
	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	//

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")
	/**
2
*/

	augmentEpoch(vs2)

	//
	assureNewDeadline(deadlineNominateChan, level, duration, cs1.settings.Nominate(duration).Nanoseconds())

	rs := cs1.FetchDurationStatus()

	require.Nil(t, rs.NominationLedger, "REDACTED")

	//
	assurePreballot(ballotChan, level, duration)
	//
	certifyPreballot(t, cs1, duration, vss[0], rs.LatchedLedger.Digest())

	//
	bps, err := rs.LatchedLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	attestAppendBallots(cs1, engineproto.PreballotKind, digest, bps.Heading(), false, vs2)
	assurePreballot(ballotChan, level, duration)

	//
	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preballot(duration).Nanoseconds())

	assurePreendorse(ballotChan, level, duration) //
	//
	//
	certifyPreendorse(t, cs1, duration, 0, vss[0], nil, theLedgerDigest)

	//
	bps2, err := rs.LatchedLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	attestAppendBallots(cs1, engineproto.PreendorseKind, digest, bps2.Heading(), true, vs2)
	assurePreendorse(ballotChan, level, duration)

	//
	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")
	/**
2
*/

	augmentEpoch(vs2)

	assureNewNomination(nominationChan, level, duration)
	rs = cs1.FetchDurationStatus()

	//
	if !bytes.Equal(rs.NominationLedger.Digest(), rs.LatchedLedger.Digest()) {
		panic(fmt.Sprintf(
			"REDACTED",
			rs.NominationLedger,
			rs.LatchedLedger))
	}

	assurePreballot(ballotChan, level, duration) //
	certifyPreballot(t, cs1, duration, vss[0], rs.LatchedLedger.Digest())

	bps0, err := rs.NominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	attestAppendBallots(cs1, engineproto.PreballotKind, digest, bps0.Heading(), false, vs2)
	assurePreballot(ballotChan, level, duration)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preballot(duration).Nanoseconds())
	assurePreendorse(ballotChan, level, duration) //

	certifyPreendorse(t, cs1, duration, 0, vss[0], nil, theLedgerDigest) //

	bps1, err := rs.NominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	attestAppendBallots(
		cs1,
		engineproto.PreendorseKind,
		digest,
		bps1.Heading(),
		true,
		vs2) //
	assurePreendorse(ballotChan, level, duration)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	cs2, _ := randomStatus(2) //
	//
	nomination, nominationLedger := determineNomination(ctx, t, cs2, vs2, vs2.Level, vs2.Cycle+1)
	if nomination == nil || nominationLedger == nil {
		t.Fatal("REDACTED")
	}

	augmentEpoch(vs2)

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")
	/**
C
*/

	//
	//
	bps3, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, bps3, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureNewNomination(nominationChan, level, duration)
	assurePreballot(ballotChan, level, duration) //
	//
	certifyPreballot(t, cs1, 3, vss[0], cs1.LatchedLedger.Digest())

	//
	bps4, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedger.Digest(), bps4.Heading(), false, vs2)
	assurePreballot(ballotChan, level, duration)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preballot(duration).Nanoseconds())
	assurePreendorse(ballotChan, level, duration)
	certifyPreendorse(t, cs1, duration, 0, vss[0], nil, theLedgerDigest) //

	bps5, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	attestAppendBallots(
		cs1,
		engineproto.PreendorseKind,
		nominationLedger.Digest(),
		bps5.Heading(),
		true,
		vs2) //
	assurePreendorse(ballotChan, level, duration)
}

//
//
//
//
func VerifyStatusSecurePOLResecure(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	newLedgerChan := enrol(cs1.eventBus, kinds.EventInquireNewLedgerHeading)

	//

	/**
l

s
*/

	//
	beginVerifyEpoch(cs1, level, duration)

	assureNewEpoch(newEpochChan, level, duration)
	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	theLedgerDigest := rs.NominationLedger.Digest()
	theLedgerSegments := rs.NominationLedgerSegments.Heading()

	assurePreballot(ballotChan, level, duration) //

	attestAppendBallots(cs1, engineproto.PreballotKind, theLedgerDigest, theLedgerSegments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration) //
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], theLedgerDigest, theLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	cs2 := newStatus(cs1.status, vs2, objectdepot.NewInRamSoftware())
	nomination, nominationLedger := determineNomination(ctx, t, cs2, vs2, vs2.Level, vs2.Cycle+1)
	if nomination == nil || nominationLedger == nil {
		t.Fatal("REDACTED")
	}
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	nominationLedgerDigest := nominationLedger.Digest()
	require.NotEqual(t, nominationLedgerDigest, theLedgerDigest)

	augmentEpoch(vs2, vs3, vs4)

	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //
	//
	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")

	/**
)

!
*/

	//
	//
	assureNewNomination(nominationChan, level, duration)

	//
	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], theLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], nominationLedgerDigest, nominationLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), true, vs2, vs3)
	assureNewLedgerHeading(newLedgerChan, level, nominationLedgerDigest)

	assureNewEpoch(newEpochChan, level+1, 0)
}

//
func VerifyStatusSecurePOLRelease(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	releaseChan := enrol(cs1.eventBus, kinds.EventInquireRelease)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//

	/**
l
s
*/

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	theLedgerDigest := rs.NominationLedger.Digest()
	theLedgerSegments := rs.NominationLedgerSegments.Heading()

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], theLedgerDigest)

	attestAppendBallots(cs1, engineproto.PreballotKind, theLedgerDigest, theLedgerSegments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], theLedgerDigest, theLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs4)
	attestAppendBallots(cs1, engineproto.PreendorseKind, theLedgerDigest, theLedgerSegments, true, vs3)

	//
	nomination, nominationLedger := determineNomination(ctx, t, cs1, vs2, vs2.Level, vs2.Cycle+1)
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())
	rs = cs1.FetchDurationStatus()
	latchedLedgerDigest := rs.LatchedLedger.Digest()

	augmentEpoch(vs2, vs3, vs4)
	duration++ //

	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")
	/**
_
!
*/
	//
	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureNewNomination(nominationChan, level, duration)

	//
	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], latchedLedgerDigest)
	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, false, vs2, vs3, vs4)

	//
	assureNewRelease(releaseChan, level, duration)
	assurePreendorse(ballotChan, level, duration)

	//
	//
	certifyPreendorse(t, cs1, duration, -1, vss[0], nil, nil)

	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3)
	assureNewEpoch(newEpochChan, level, duration+1)
}

//
//
//
//
func VerifyStatusSecurePOLReleaseOnUnclearLedger(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	//

	/**
l
*/

	//
	beginVerifyEpoch(cs1, level, duration)

	assureNewEpoch(newEpochChan, level, duration)
	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	initialLedgerDigest := rs.NominationLedger.Digest()
	initialLedgerSegments := rs.NominationLedgerSegments.Heading()

	assurePreballot(ballotChan, level, duration) //

	attestAppendBallots(cs1, engineproto.PreballotKind, initialLedgerDigest, initialLedgerSegments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration) //
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], initialLedgerDigest, initialLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	cs2 := newStatus(cs1.status, vs2, objectdepot.NewInRamSoftware())
	nomination, nominationLedger := determineNomination(ctx, t, cs2, vs2, vs2.Level, vs2.Cycle+1)
	if nomination == nil || nominationLedger == nil {
		t.Fatal("REDACTED")
	}
	momentLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	momentLedgerDigest := nominationLedger.Digest()
	require.NotEqual(t, momentLedgerDigest, initialLedgerDigest)

	augmentEpoch(vs2, vs3, vs4)

	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //

	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")

	/**
)
*/

	//

	//
	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], initialLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, momentLedgerDigest, momentLedgerSegments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, -1, vss[0], nil, nil)

	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, momentLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	cs3 := newStatus(cs1.status, vs3, objectdepot.NewInRamSoftware())
	nomination, nominationLedger = determineNomination(ctx, t, cs3, vs3, vs3.Level, vs3.Cycle+1)
	if nomination == nil || nominationLedger == nil {
		t.Fatal("REDACTED")
	}
	tertiaryNominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	tertiaryNominationLedgerDigest := nominationLedger.Digest()
	require.NotEqual(t, momentLedgerDigest, tertiaryNominationLedgerDigest)

	augmentEpoch(vs2, vs3, vs4)

	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")

	/**
)
*/

	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, tertiaryNominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assurePreballot(ballotChan, level, duration)
	//
	certifyPreballot(t, cs1, duration, vss[0], tertiaryNominationLedgerDigest)

	attestAppendBallots(cs1, engineproto.PreballotKind, tertiaryNominationLedgerDigest, tertiaryNominationLedgerSegments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], tertiaryNominationLedgerDigest, tertiaryNominationLedgerDigest)
}

//
//
//
//
func VerifyStatusSecurePOLSecurity1(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	deadlineNominateChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, cs1.Level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	nominationLedger := rs.NominationLedger

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nominationLedger.Digest())

	//
	bps, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	preballots := attestBallots(engineproto.PreballotKind, nominationLedger.Digest(), bps.Heading(), false, vs2, vs3, vs4)

	t.Logf("REDACTED", fmt.Sprintf("REDACTED", nominationLedger.Digest()))

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	assurePreendorse(ballotChan, level, duration)
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	t.Log("REDACTED")

	nomination, nominationLedger := determineNomination(ctx, t, cs1, vs2, vs2.Level, vs2.Cycle+1)
	nominationLedgerDigest := nominationLedger.Digest()
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	augmentEpoch(vs2, vs3, vs4)

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)

	//
	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	/*2
k
!
*/

	assureNewNomination(nominationChan, level, duration)

	rs = cs1.FetchDurationStatus()

	if rs.LatchedLedger != nil {
		panic("REDACTED")
	}
	t.Logf("REDACTED", fmt.Sprintf("REDACTED", nominationLedgerDigest))

	//
	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nominationLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], nominationLedgerDigest, nominationLedgerDigest)

	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	augmentEpoch(vs2, vs3, vs4)
	duration++ //

	assureNewEpoch(newEpochChan, level, duration)

	t.Log("REDACTED")
	/*3
!
*/

	//
	assureNewDeadline(deadlineNominateChan, level, duration, cs1.settings.Nominate(duration).Nanoseconds())

	//
	assurePreballot(ballotChan, level, duration)
	//
	certifyPreballot(t, cs1, duration, vss[0], nominationLedgerDigest)

	newPhaseChan := enrol(cs1.eventBus, kinds.EventInquireNewEpochPhase)

	//
	//
	appendBallots(cs1, preballots...)

	t.Log("REDACTED")

	assureNoNewEpochPhase(newPhaseChan)
}

//
//
//
//

//
//
func VerifyStatusSecurePOLSecurity2(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	releaseChan := enrol(cs1.eventBus, kinds.EventInquireRelease)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	//
	_, nominationLedger0 := determineNomination(ctx, t, cs1, vss[0], level, duration)
	nominationLedgerDigest0 := nominationLedger0.Digest()
	nominationLedgerSegments0, err := nominationLedger0.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	nominationLedgerID0 := kinds.LedgerUID{Digest: nominationLedgerDigest0, SegmentAssignHeading: nominationLedgerSegments0.Heading()}

	//
	preballots := attestBallots(engineproto.PreballotKind, nominationLedgerDigest0, nominationLedgerSegments0.Heading(), false, vs2, vs3, vs4)

	//
	nomination1, nominationLedger1 := determineNomination(ctx, t, cs1, vs2, vs2.Level, vs2.Cycle+1)
	nominationLedgerDigest1 := nominationLedger1.Digest()
	nominationLedgerSegments1, err := nominationLedger1.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	augmentEpoch(vs2, vs3, vs4)

	duration++ //
	t.Log("REDACTED")
	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	if err := cs1.CollectionNominationAndLedger(nomination1, nominationLedger1, nominationLedgerSegments1, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureNewNomination(nominationChan, level, duration)

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nominationLedgerDigest1)

	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest1, nominationLedgerSegments1.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], nominationLedgerDigest1, nominationLedgerDigest1)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs4)
	attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedgerDigest1, nominationLedgerSegments1.Heading(), true, vs3)

	augmentEpoch(vs2, vs3, vs4)

	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //
	//
	newNomination := kinds.NewNomination(level, duration, 0, nominationLedgerID0)
	p := newNomination.ToSchema()
	if err := vs3.AttestNomination(cs1.status.LedgerUID, p); err != nil {
		t.Fatal(err)
	}

	newNomination.Autograph = p.Autograph

	if err := cs1.CollectionNominationAndLedger(newNomination, nominationLedger0, nominationLedgerSegments0, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	//
	appendBallots(cs1, preballots...)

	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")
	/*2
k
*/
	assureNewNomination(nominationChan, level, duration)

	assureNoNewRelease(releaseChan)
	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nominationLedgerDigest1)
}

//
//

//
//
func VerifyNominateSoundLedger(t *testing.T) {
	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	deadlineNominateChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	releaseChan := enrol(cs1.eventBus, kinds.EventInquireRelease)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, cs1.Level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	nominationLedger := rs.NominationLedger
	nominationLedgerDigest := nominationLedger.Digest()

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nominationLedgerDigest)

	//
	bps, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)
	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest, bps.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], nominationLedgerDigest, nominationLedgerDigest)

	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	augmentEpoch(vs2, vs3, vs4)
	duration++ //

	assureNewEpoch(newEpochChan, level, duration)

	t.Log("REDACTED")

	//
	assureNewDeadline(deadlineNominateChan, level, duration, cs1.settings.Nominate(duration).Nanoseconds())

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nominationLedgerDigest)

	attestAppendBallots(cs1, engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, false, vs2, vs3, vs4)

	assureNewRelease(releaseChan, level, duration)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, -1, vss[0], nil, nil)

	augmentEpoch(vs2, vs3, vs4)
	augmentEpoch(vs2, vs3, vs4)

	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	duration += 2 //

	assureNewEpoch(newEpochChan, level, duration)
	t.Log("REDACTED")

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //

	assureNewEpoch(newEpochChan, level, duration)

	t.Log("REDACTED")

	assureNewNomination(nominationChan, level, duration)

	rs = cs1.FetchDurationStatus()
	assert.True(t, bytes.Equal(rs.NominationLedger.Digest(), nominationLedgerDigest))
	assert.True(t, bytes.Equal(rs.NominationLedger.Digest(), rs.SoundLedger.Digest()))
	assert.True(t, rs.Nomination.POLDuration == rs.SoundEpoch)
	assert.True(t, bytes.Equal(rs.Nomination.LedgerUID.Digest, rs.SoundLedger.Digest()))
}

//
//
func VerifyCollectionSoundLedgerOnDeferredPreballot(t *testing.T) {
	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	soundLedgerChan := enrol(cs1.eventBus, kinds.EventInquireSoundLedger)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, cs1.Level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	nominationLedger := rs.NominationLedger
	nominationLedgerDigest := nominationLedger.Digest()
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nominationLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), false, vs2)

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, false, vs3)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preballot(duration).Nanoseconds())

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, -1, vss[0], nil, nil)

	rs = cs1.FetchDurationStatus()

	assert.True(t, rs.SoundLedger == nil)
	assert.True(t, rs.SoundLedgerSegments == nil)
	assert.True(t, rs.SoundEpoch == -1)

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), false, vs4)

	assureNewSoundLedger(soundLedgerChan, level, duration)

	rs = cs1.FetchDurationStatus()

	assert.True(t, bytes.Equal(rs.SoundLedger.Digest(), nominationLedgerDigest))
	assert.True(t, rs.SoundLedgerSegments.Heading().Matches(nominationLedgerSegments.Heading()))
	assert.True(t, rs.SoundEpoch == duration)
}

//
//
//
func VerifyCollectionSoundLedgerOnDeferredNomination(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	deadlineNominateChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	soundLedgerChan := enrol(cs1.eventBus, kinds.EventInquireSoundLedger)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)
	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)

	duration++ //
	augmentEpoch(vs2, vs3, vs4)

	beginVerifyEpoch(cs1, cs1.Level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewDeadline(deadlineNominateChan, level, duration, cs1.settings.Nominate(duration).Nanoseconds())

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nil)

	nomination, nominationLedger := determineNomination(ctx, t, cs1, vs2, vs2.Level, vs2.Cycle+1)
	nominationLedgerDigest := nominationLedger.Digest()
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	//
	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), false, vs2, vs3, vs4)
	assureNewSoundLedger(soundLedgerChan, level, duration)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preballot(duration).Nanoseconds())

	assurePreendorse(ballotChan, level, duration)
	certifyPreendorse(t, cs1, duration, -1, vss[0], nil, nil)

	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()

	assert.True(t, bytes.Equal(rs.SoundLedger.Digest(), nominationLedgerDigest))
	assert.True(t, rs.SoundLedgerSegments.Heading().Matches(nominationLedgerSegments.Heading()))
	assert.True(t, rs.SoundEpoch == duration)
}

func VerifyHandleNominationAllow(t *testing.T) {
	for _, verifyInstance := range []struct {
		label               string
		allow             bool
		anticipatedNullPreballot bool
	}{
		{
			label:               "REDACTED",
			allow:             true,
			anticipatedNullPreballot: false,
		},
		{
			label:               "REDACTED",
			allow:             false,
			anticipatedNullPreballot: true,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			m := abciemulators.NewSoftware(t)
			state := iface.Responseprocessnomination_DECLINE
			if verifyInstance.allow {
				state = iface.Responseprocessnomination_ALLOW
			}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Status: state}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil).Maybe()
			cs1, _ := randomStatusWithApplication(4, m)
			level, duration := cs1.Level, cs1.Cycle

			nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
			newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
			pv1, err := cs1.privateRatifier.FetchPublicKey()
			require.NoError(t, err)
			address := pv1.Location()
			ballotChan := enrolToPoller(cs1, address)

			beginVerifyEpoch(cs1, cs1.Level, duration)
			assureNewEpoch(newEpochChan, level, duration)

			assureNewNomination(nominationChan, level, duration)
			rs := cs1.FetchDurationStatus()
			var preballotDigest cometbytes.HexOctets
			if !verifyInstance.anticipatedNullPreballot {
				preballotDigest = rs.NominationLedger.Digest()
			}
			assurePreballotAlign(t, ballotChan, level, duration, preballotDigest)
		})
	}
}

//
//
func VerifyExpandBallotInvokedWhenActivated(t *testing.T) {
	for _, verifyInstance := range []struct {
		label    string
		activated bool
	}{
		{
			label:    "REDACTED",
			activated: true,
		},
		{
			label:    "REDACTED",
			activated: false,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			m := abciemulators.NewSoftware(t)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_ALLOW}, nil)
			if verifyInstance.activated {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyExpandBallot{
					BallotAddition: []byte("REDACTED"),
				}, nil)
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
					Status: iface.Responseverifyballotextension_ALLOW,
				}, nil)
			}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCompleteLedger{}, nil).Maybe()
			level := int64(1)
			if !verifyInstance.activated {
				level = 0
			}
			cs1, vss := randomStatusWithApplicationWithLevel(4, m, level)

			level, duration := cs1.Level, cs1.Cycle

			nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
			newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
			pv1, err := cs1.privateRatifier.FetchPublicKey()
			require.NoError(t, err)
			address := pv1.Location()
			ballotChan := enrolToPoller(cs1, address)

			beginVerifyEpoch(cs1, cs1.Level, duration)
			assureNewEpoch(newEpochChan, level, duration)
			assureNewNomination(nominationChan, level, duration)

			m.AssertNotCalled(t, "REDACTED", mock.Anything, mock.Anything)

			rs := cs1.FetchDurationStatus()

			ledgerUID := kinds.LedgerUID{
				Digest:          rs.NominationLedger.Digest(),
				SegmentAssignHeading: rs.NominationLedgerSegments.Heading(),
			}
			attestAppendBallots(cs1, engineproto.PreballotKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, false, vss[1:]...)
			assurePreballotAlign(t, ballotChan, level, duration, ledgerUID.Digest)

			assurePreendorse(ballotChan, level, duration)

			if verifyInstance.activated {
				m.AssertCalled(t, "REDACTED", context.TODO(), &iface.QueryExpandBallot{
					Level:             level,
					Digest:               ledgerUID.Digest,
					Time:               rs.NominationLedger.Time,
					Txs:                rs.NominationLedger.Txs.ToSegmentOfOctets(),
					NominatedFinalEndorse: iface.EndorseDetails{},
					Malpractice:        rs.NominationLedger.Proof.Proof.ToIface(),
					FollowingRatifiersDigest: rs.NominationLedger.FollowingRatifiersDigest,
					RecommenderLocation:    rs.NominationLedger.RecommenderLocation,
				})
			} else {
				m.AssertNotCalled(t, "REDACTED", mock.Anything, mock.Anything)
			}

			attestAppendBallots(cs1, engineproto.PreendorseKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, verifyInstance.activated, vss[1:]...)
			assureNewEpoch(newEpochChan, level+1, 0)
			m.AssertExpectations(t)

			//
			//
			for _, pv := range vss[1:3] {
				pv, err := pv.FetchPublicKey()
				require.NoError(t, err)
				address := pv.Location()
				if verifyInstance.activated {
					m.AssertCalled(t, "REDACTED", context.TODO(), &iface.QueryValidateBallotAddition{
						Digest:             ledgerUID.Digest,
						RatifierLocation: address,
						Level:           level,
						BallotAddition:    []byte("REDACTED"),
					})
				} else {
					m.AssertNotCalled(t, "REDACTED", mock.Anything, mock.Anything)
				}
			}
		})
	}
}

//
//
func VerifyValidateBallotAdditionNotInvokedOnMissingPreendorse(t *testing.T) {
	m := abciemulators.NewSoftware(t)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_ALLOW}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyExpandBallot{
		BallotAddition: []byte("REDACTED"),
	}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
		Status: iface.Responseverifyballotextension_ALLOW,
	}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCompleteLedger{}, nil).Maybe()
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
	cs1, vss := randomStatusWithApplication(4, m)
	level, duration := cs1.Level, cs1.Cycle
	cs1.status.AgreementOptions.Iface.BallotPluginsActivateLevel = cs1.Level

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	beginVerifyEpoch(cs1, cs1.Level, duration)
	assureNewEpoch(newEpochChan, level, duration)
	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()

	ledgerUID := kinds.LedgerUID{
		Digest:          rs.NominationLedger.Digest(),
		SegmentAssignHeading: rs.NominationLedgerSegments.Heading(),
	}
	attestAppendBallots(cs1, engineproto.PreballotKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, false, vss...)
	assurePreballotAlign(t, ballotChan, level, duration, ledgerUID.Digest)

	assurePreendorse(ballotChan, level, duration)

	m.AssertCalled(t, "REDACTED", context.TODO(), &iface.QueryExpandBallot{
		Level:             level,
		Digest:               ledgerUID.Digest,
		Time:               rs.NominationLedger.Time,
		Txs:                rs.NominationLedger.Txs.ToSegmentOfOctets(),
		NominatedFinalEndorse: iface.EndorseDetails{},
		Malpractice:        rs.NominationLedger.Proof.Proof.ToIface(),
		FollowingRatifiersDigest: rs.NominationLedger.FollowingRatifiersDigest,
		RecommenderLocation:    rs.NominationLedger.RecommenderLocation,
	})

	attestAppendBallots(cs1, engineproto.PreendorseKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, true, vss[2:]...)
	assureNewEpoch(newEpochChan, level+1, 0)
	m.AssertExpectations(t)

	//
	//
	pv, err := vss[1].FetchPublicKey()
	require.NoError(t, err)
	address = pv.Location()

	m.AssertNotCalled(t, "REDACTED", context.TODO(), &iface.QueryValidateBallotAddition{
		Digest:             ledgerUID.Digest,
		RatifierLocation: address,
		Level:           level,
		BallotAddition:    []byte("REDACTED"),
	})
}

//
//
//
//
//
//
func VerifyArrangeNominationAcceptsBallotPlugins(t *testing.T) {
	//
	ballotPlugins := [][]byte{
		[]byte("REDACTED"),
		[]byte("REDACTED"),
		[]byte("REDACTED"),
		[]byte("REDACTED"),
	}

	m := abciemulators.NewSoftware(t)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyExpandBallot{
		BallotAddition: ballotPlugins[0],
	}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_ALLOW}, nil)

	//
	rpp := &iface.QueryArrangeNomination{}
	m.On("REDACTED", mock.Anything, mock.MatchedBy(func(r *iface.QueryArrangeNomination) bool {
		rpp = r
		return true
	})).Return(&iface.ReplyArrangeNomination{}, nil)

	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{Status: iface.Responseverifyballotextension_ALLOW}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCompleteLedger{}, nil)

	cs1, vss := randomStatusWithApplication(4, m)
	level, duration := cs1.Level, cs1.Cycle

	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)
	assureNewNomination(nominationChan, level, duration)

	rs := cs1.FetchDurationStatus()
	ledgerUID := kinds.LedgerUID{
		Digest:          rs.NominationLedger.Digest(),
		SegmentAssignHeading: rs.NominationLedgerSegments.Heading(),
	}
	attestAppendBallots(cs1, engineproto.PreballotKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, false, vss[1:]...)

	//
	for i, vs := range vss[1:] {
		attestAppendPreendorseWithAddition(t, cs1, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, ballotPlugins[i+1], vs)
	}

	assurePreballot(ballotChan, level, duration)

	//
	assurePreendorseAlign(t, ballotChan, level, duration, ledgerUID.Digest)
	augmentLevel(vss[1:]...)

	level++
	duration = 0
	assureNewEpoch(newEpochChan, level, duration)
	augmentEpoch(vss[1:]...)
	augmentEpoch(vss[1:]...)
	augmentEpoch(vss[1:]...)
	duration = 3

	ledgerUidtwo := kinds.LedgerUID{}
	attestAppendBallots(cs1, engineproto.PreendorseKind, ledgerUidtwo.Digest, ledgerUidtwo.SegmentAssignHeading, true, vss[1:]...)
	assureNewEpoch(newEpochChan, level, duration)
	assureNewNomination(nominationChan, level, duration)

	//
	//
	require.Len(t, rpp.NativeFinalEndorse.Ballots, len(vss))
	for i := range vss {
		ballot := &rpp.NativeFinalEndorse.Ballots[i]
		require.Equal(t, ballot.BallotAddition, ballotPlugins[i])

		require.NotZero(t, len(ballot.AdditionAutograph))
		cve := engineproto.StandardBallotAddition{
			Addition: ballot.BallotAddition,
			Level:    level - 1, //
			Cycle:     int64(rpp.NativeFinalEndorse.Cycle),
			SeriesUid:   verify.StandardVerifyLedgerUID,
		}
		extensionAttestOctets, err := protoio.SerializeSeparated(&cve)
		require.NoError(t, err)
		publicKey, err := vss[i].FetchPublicKey()
		require.NoError(t, err)
		require.True(t, publicKey.ValidateAutograph(extensionAttestOctets, ballot.AdditionAutograph))
	}
}

func VerifyCompleteLedgerInvoked(t *testing.T) {
	for _, verifyInstance := range []struct {
		label         string
		ballotNull      bool
		anticipateInvoked bool
	}{
		{
			label:         "REDACTED",
			ballotNull:      false,
			anticipateInvoked: true,
		},
		{
			label:         "REDACTED",
			ballotNull:      true,
			anticipateInvoked: false,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			m := abciemulators.NewSoftware(t)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{
				Status: iface.Responseprocessnomination_ALLOW,
			}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			//
			//
			if !verifyInstance.ballotNull {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyExpandBallot{}, nil)
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
					Status: iface.Responseverifyballotextension_ALLOW,
				}, nil)
			}
			r := &iface.ReplyCompleteLedger{ApplicationDigest: []byte("REDACTED")}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(r, nil).Maybe()
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()

			cs1, vss := randomStatusWithApplication(4, m)
			level, duration := cs1.Level, cs1.Cycle

			nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
			newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
			pv1, err := cs1.privateRatifier.FetchPublicKey()
			require.NoError(t, err)
			address := pv1.Location()
			ballotChan := enrolToPoller(cs1, address)

			beginVerifyEpoch(cs1, cs1.Level, duration)
			assureNewEpoch(newEpochChan, level, duration)
			assureNewNomination(nominationChan, level, duration)
			rs := cs1.FetchDurationStatus()

			ledgerUID := kinds.LedgerUID{}
			followingEpoch := duration + 1
			followingLevel := level
			if !verifyInstance.ballotNull {
				followingEpoch = 0
				followingLevel = level + 1
				ledgerUID = kinds.LedgerUID{
					Digest:          rs.NominationLedger.Digest(),
					SegmentAssignHeading: rs.NominationLedgerSegments.Heading(),
				}
			}

			attestAppendBallots(cs1, engineproto.PreballotKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, false, vss[1:]...)
			assurePreballotAlign(t, ballotChan, level, duration, rs.NominationLedger.Digest())

			attestAppendBallots(cs1, engineproto.PreendorseKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, true, vss[1:]...)
			assurePreendorse(ballotChan, level, duration)

			assureNewEpoch(newEpochChan, followingLevel, followingEpoch)
			m.AssertExpectations(t)

			if !verifyInstance.anticipateInvoked {
				m.AssertNotCalled(t, "REDACTED", context.TODO(), mock.Anything)
			} else {
				m.AssertCalled(t, "REDACTED", context.TODO(), mock.Anything)
			}
		})
	}
}

//
//
//
func VerifyBallotAdditionActivateLevel(t *testing.T) {
	for _, verifyInstance := range []struct {
		label                  string
		activateLevel          int64
		hasAddition          bool
		anticipateExpandInvoked    bool
		anticipateValidateInvoked    bool
		anticipateTriumphantEpoch bool
	}{
		{
			label:                  "REDACTED",
			hasAddition:          true,
			activateLevel:          0,
			anticipateExpandInvoked:    false,
			anticipateValidateInvoked:    false,
			anticipateTriumphantEpoch: false,
		},
		{
			label:                  "REDACTED",
			hasAddition:          false,
			activateLevel:          0,
			anticipateExpandInvoked:    false,
			anticipateValidateInvoked:    false,
			anticipateTriumphantEpoch: true,
		},
		{
			label:                  "REDACTED",
			hasAddition:          true,
			activateLevel:          1,
			anticipateExpandInvoked:    true,
			anticipateValidateInvoked:    true,
			anticipateTriumphantEpoch: true,
		},
		{
			label:                  "REDACTED",
			hasAddition:          false,
			activateLevel:          1,
			anticipateExpandInvoked:    true,
			anticipateValidateInvoked:    false,
			anticipateTriumphantEpoch: false,
		},
		{
			label:                  "REDACTED",
			hasAddition:          false,
			activateLevel:          2,
			anticipateExpandInvoked:    false,
			anticipateValidateInvoked:    false,
			anticipateTriumphantEpoch: true,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			countRatifiers := 3
			m := abciemulators.NewSoftware(t)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{
				Status: iface.Responseprocessnomination_ALLOW,
			}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			if verifyInstance.anticipateExpandInvoked {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyExpandBallot{}, nil)
			}
			if verifyInstance.anticipateValidateInvoked {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
					Status: iface.Responseverifyballotextension_ALLOW,
				}, nil).Times(countRatifiers - 1)
			}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCompleteLedger{}, nil).Maybe()
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
			cs1, vss := randomStatusWithApplicationWithLevel(countRatifiers, m, verifyInstance.activateLevel)
			cs1.status.AgreementOptions.Iface.BallotPluginsActivateLevel = verifyInstance.activateLevel
			level, duration := cs1.Level, cs1.Cycle

			deadlineChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
			nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
			newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
			pv1, err := cs1.privateRatifier.FetchPublicKey()
			require.NoError(t, err)
			address := pv1.Location()
			ballotChan := enrolToPoller(cs1, address)

			beginVerifyEpoch(cs1, cs1.Level, duration)
			assureNewEpoch(newEpochChan, level, duration)
			assureNewNomination(nominationChan, level, duration)
			rs := cs1.FetchDurationStatus()

			//
			attestAppendBallots(cs1, engineproto.PreballotKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), false, vss[1:]...)
			assurePreballotAlign(t, ballotChan, level, duration, rs.NominationLedger.Digest())

			var ext []byte
			if verifyInstance.hasAddition {
				ext = []byte("REDACTED")
			}

			for _, vs := range vss[1:] {
				ballot, err := vs.attestBallot(engineproto.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerSegments.Heading(), ext, verifyInstance.hasAddition)
				require.NoError(t, err)
				appendBallots(cs1, ballot)
			}
			if verifyInstance.anticipateTriumphantEpoch {
				assurePreendorse(ballotChan, level, duration)
				level++
				assureNewEpoch(newEpochChan, level, duration)
			} else {
				assureNoNewDeadline(deadlineChan, cs1.settings.Preendorse(duration).Nanoseconds())
			}

			m.AssertExpectations(t)
		})
	}
}

//
//
//
func VerifyStatusDoesntCollapseOnCorruptBallot(t *testing.T) {
	cs, vss := randomStatus(2)
	level, duration := cs.Level, cs.Cycle
	//
	node := p2pemulator.NewNode(nil)

	beginVerifyEpoch(cs, level, duration)

	_, nominationLedger := determineNomination(context.Background(), t, cs, vss[0], level, duration)
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	assert.NoError(t, err)

	ballot := attestBallot(vss[1], engineproto.PreendorseKind, nominationLedger.Digest(), nominationLedgerSegments.Heading(), true)

	//
	ballot.RatifierOrdinal = int32(len(vss))

	ballotSignal := &BallotSignal{ballot}
	assert.NotPanics(t, func() {
		cs.processMessage(messageDetails{ballotSignal, node.ID()})
	})

	appended, err := cs.AppendBallot(ballot, node.ID())
	assert.False(t, appended)
	assert.NoError(t, err)
	//
	//
}

//
//
//
func VerifyPendingDeadlineOnNullPolka(*testing.T) {
	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())
	assureNewEpoch(newEpochChan, level, duration+1)
}

//
//
//
func VerifyPendingDeadlineNominateOnNewEpoch(t *testing.T) {
	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assurePreballot(ballotChan, level, duration)

	augmentEpoch(vss[1:]...)
	attestAppendBallots(cs1, engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, false, vs2, vs3, vs4)

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)

	rs := cs1.FetchDurationStatus()
	assert.True(t, rs.Phase == cskinds.DurationPhaseNominate) //

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Nominate(duration).Nanoseconds())

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nil)
}

//
//
//
func VerifyEpochOmitOnNullPolkaFromSuperiorEpoch(t *testing.T) {
	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assurePreballot(ballotChan, level, duration)

	augmentEpoch(vss[1:]...)
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2, vs3, vs4)

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)

	assurePreendorse(ballotChan, level, duration)
	certifyPreendorse(t, cs1, duration, -1, vss[0], nil, nil)

	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //
	assureNewEpoch(newEpochChan, level, duration)
}

//
//
//
func VerifyWaitDeadlineNominateOnNullPolkaForTheOngoingEpoch(t *testing.T) {
	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, int32(1)

	deadlineNominateChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	augmentEpoch(vss[1:]...)
	attestAppendBallots(cs1, engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, false, vs2, vs3, vs4)

	assureNewDeadline(deadlineNominateChan, level, duration, cs1.settings.Nominate(duration).Nanoseconds())

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], nil)
}

//
//
func VerifyIssueNewSoundLedgerEventOnEndorseLackingLedger(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, int32(1)

	augmentEpoch(vs2, vs3, vs4)

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	soundLedgerChan := enrol(cs1.eventBus, kinds.EventInquireSoundLedger)

	_, nominationLedger := determineNomination(ctx, t, cs1, vs2, vs2.Level, vs2.Cycle)
	nominationLedgerDigest := nominationLedger.Digest()
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), true, vs2, vs3, vs4)
	assureNewSoundLedger(soundLedgerChan, level, duration)

	rs := cs1.FetchDurationStatus()
	assert.True(t, rs.Phase == cskinds.DurationPhaseEndorse)
	assert.True(t, rs.NominationLedger == nil)
	assert.True(t, rs.NominationLedgerSegments.Heading().Matches(nominationLedgerSegments.Heading()))
}

//
//
//
func VerifyEndorseFromPrecedingEpoch(t *testing.T) {
	ctx := t.Context()

	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, int32(1)

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	soundLedgerChan := enrol(cs1.eventBus, kinds.EventInquireSoundLedger)
	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)

	nomination, nominationLedger := determineNomination(ctx, t, cs1, vs2, vs2.Level, vs2.Cycle)
	nominationLedgerDigest := nominationLedger.Digest()
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedgerDigest, nominationLedgerSegments.Heading(), true, vs2, vs3, vs4)

	assureNewSoundLedger(soundLedgerChan, level, duration)

	rs := cs1.FetchDurationStatus()
	assert.True(t, rs.Phase == cskinds.DurationPhaseEndorse)
	assert.True(t, rs.EndorseEpoch == vs2.Cycle)
	assert.True(t, rs.NominationLedger == nil)
	assert.True(t, rs.NominationLedgerSegments.Heading().Matches(nominationLedgerSegments.Heading()))

	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureNewNomination(nominationChan, level, duration)
	assureNewEpoch(newEpochChan, level+1, 0)
}

type mockTransferAlerter struct {
	ch chan struct{}
}

func (n *mockTransferAlerter) TransAccessible() <-chan struct{} {
	return n.ch
}

func (n *mockTransferAlerter) Alert() {
	n.ch <- struct{}{}
}

//
//
//
func VerifyBeginFollowingLevelAccuratelyAfterDeadline(t *testing.T) {
	settings.Agreement.OmitDeadlineEndorse = false
	cs1, vss := randomStatus(4)
	cs1.transferAlerter = &mockTransferAlerter{ch: make(chan struct{})}

	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	deadlineNominateChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineNominate)
	preendorseDeadlineChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)

	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	newLedgerHeading := enrol(cs1.eventBus, kinds.EventInquireNewLedgerHeading)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	theLedgerDigest := rs.NominationLedger.Digest()
	theLedgerSegments := rs.NominationLedgerSegments.Heading()

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], theLedgerDigest)

	attestAppendBallots(cs1, engineproto.PreballotKind, theLedgerDigest, theLedgerSegments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], theLedgerDigest, theLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2)
	attestAppendBallots(cs1, engineproto.PreendorseKind, theLedgerDigest, theLedgerSegments, true, vs3)

	//
	assurePreendorseDeadline(preendorseDeadlineChan)

	assureNewEpoch(newEpochChan, level, duration+1)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, theLedgerDigest, theLedgerSegments, true, vs4)

	assureNewLedgerHeading(newLedgerHeading, level, theLedgerDigest)

	cs1.transferAlerter.(*mockTransferAlerter).Alert()

	assureNewDeadline(deadlineNominateChan, level+1, duration, cs1.settings.Nominate(duration).Nanoseconds())
	rs = cs1.FetchDurationStatus()
	assert.False(
		t,
		rs.ActivatedDeadlinePreendorse,
		"REDACTED")
}

func VerifyRestoreDeadlinePreendorseOnNewLevel(t *testing.T) {
	ctx := t.Context()

	settings.Agreement.OmitDeadlineEndorse = false
	cs1, vss := randomStatus(4)

	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle

	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)

	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	newLedgerHeading := enrol(cs1.eventBus, kinds.EventInquireNewLedgerHeading)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	theLedgerDigest := rs.NominationLedger.Digest()
	theLedgerSegments := rs.NominationLedgerSegments.Heading()

	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], theLedgerDigest)

	attestAppendBallots(cs1, engineproto.PreballotKind, theLedgerDigest, theLedgerSegments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	certifyPreendorse(t, cs1, duration, duration, vss[0], theLedgerDigest, theLedgerDigest)

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2)
	attestAppendBallots(cs1, engineproto.PreendorseKind, theLedgerDigest, theLedgerSegments, true, vs3)
	attestAppendBallots(cs1, engineproto.PreendorseKind, theLedgerDigest, theLedgerSegments, true, vs4)

	assureNewLedgerHeading(newLedgerHeading, level, theLedgerDigest)

	nomination, nominationLedger := determineNomination(ctx, t, cs1, vs2, level+1, 0)
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	if err := cs1.CollectionNominationAndLedger(nomination, nominationLedger, nominationLedgerSegments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureNewNomination(nominationChan, level+1, 0)

	rs = cs1.FetchDurationStatus()
	assert.False(
		t,
		rs.ActivatedDeadlinePreendorse,
		"REDACTED")
}

//
//
//

/**
{
)
]


)
)
)
)

e
)
h
h
e

)

s
t
)
5
)

h

t
)

e
)

o
}

{
)
]


)
)
)
)

e
)
h
h
e

2
)

t

s
t
)
5
)

t
)

2
)

o
}
*/

//
//

//
//

//
//
func VerifyStatusHalt1(t *testing.T) {
	cs1, vss := randomStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	level, duration := cs1.Level, cs1.Cycle
	segmentVolume := kinds.LedgerSegmentVolumeOctets

	nominationChan := enrol(cs1.eventBus, kinds.EventInquireFinishedNomination)
	deadlineWaitChan := enrol(cs1.eventBus, kinds.EventInquireDeadlineWait)
	newEpochChan := enrol(cs1.eventBus, kinds.EventInquireNewEpoch)
	newLedgerChan := enrol(cs1.eventBus, kinds.EventInquireNewLedger)
	pv1, err := cs1.privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	address := pv1.Location()
	ballotChan := enrolToPoller(cs1, address)

	//
	beginVerifyEpoch(cs1, level, duration)
	assureNewEpoch(newEpochChan, level, duration)

	assureNewNomination(nominationChan, level, duration)
	rs := cs1.FetchDurationStatus()
	nominationLedger := rs.NominationLedger
	nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
	require.NoError(t, err)

	assurePreballot(ballotChan, level, duration)

	attestAppendBallots(cs1, engineproto.PreballotKind, nominationLedger.Digest(), nominationLedgerSegments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChan, level, duration)
	//
	certifyPreendorse(t, cs1, duration, duration, vss[0], nominationLedger.Digest(), nominationLedger.Digest())

	//
	attestAppendBallots(cs1, engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, true, vs2) //
	attestAppendBallots(cs1, engineproto.PreendorseKind, nominationLedger.Digest(), nominationLedgerSegments.Heading(), true, vs3)
	//
	preendorse4 := attestBallot(vs4, engineproto.PreendorseKind, nominationLedger.Digest(), nominationLedgerSegments.Heading(), true)

	augmentEpoch(vs2, vs3, vs4)

	//
	assureNewDeadline(deadlineWaitChan, level, duration, cs1.settings.Preendorse(duration).Nanoseconds())

	duration++ //

	assureNewEpoch(newEpochChan, level, duration)
	rs = cs1.FetchDurationStatus()

	t.Log("REDACTED")
	/*2
k
!
*/

	//
	assurePreballot(ballotChan, level, duration)
	certifyPreballot(t, cs1, duration, vss[0], rs.LatchedLedger.Digest())

	//
	appendBallots(cs1, preendorse4)

	//
	assureNewLedger(newLedgerChan, level)

	assureNewEpoch(newEpochChan, level+1, 0)
}

func VerifyStatusResultsLedgerSegmentsMetrics(t *testing.T) {
	//
	cs, _ := randomStatus(1)
	node := p2pemulator.NewNode(nil)

	//
	segments := kinds.NewSegmentCollectionFromData(engineseed.Octets(100), 10)
	msg := &LedgerSegmentSignal{
		Level: 1,
		Cycle:  0,
		Segment:   segments.FetchSegment(0),
	}

	cs.NominationLedgerSegments = kinds.NewSegmentCollectionFromHeading(segments.Heading())
	cs.processMessage(messageDetails{msg, node.ID()})

	metricsSignal := <-cs.metricsMessageBuffer
	require.Equal(t, msg, metricsSignal.Msg, "REDACTED")
	require.Equal(t, node.ID(), metricsSignal.NodeUID, "REDACTED")

	//
	cs.processMessage(messageDetails{msg, "REDACTED"})

	//
	msg.Cycle = 1
	cs.processMessage(messageDetails{msg, node.ID()})

	//
	msg.Level = 0
	cs.processMessage(messageDetails{msg, node.ID()})

	//
	msg.Level = 3
	cs.processMessage(messageDetails{msg, node.ID()})

	select {
	case <-cs.metricsMessageBuffer:
		t.Errorf("REDACTED")
	case <-time.After(50 * time.Millisecond):
	}
}

func VerifyStatusResultBallotMetrics(t *testing.T) {
	cs, vss := randomStatus(2)
	//
	node := p2pemulator.NewNode(nil)

	randomOctets := engineseed.Octets(comethash.Volume)

	ballot := attestBallot(vss[1], engineproto.PreendorseKind, randomOctets, kinds.SegmentAssignHeading{}, true)

	ballotSignal := &BallotSignal{ballot}
	cs.processMessage(messageDetails{ballotSignal, node.ID()})

	metricsSignal := <-cs.metricsMessageBuffer
	require.Equal(t, ballotSignal, metricsSignal.Msg, "REDACTED")
	require.Equal(t, node.ID(), metricsSignal.NodeUID, "REDACTED")

	//
	cs.processMessage(messageDetails{&BallotSignal{ballot}, "REDACTED"})

	//
	augmentLevel(vss[1])
	ballot = attestBallot(vss[1], engineproto.PreendorseKind, randomOctets, kinds.SegmentAssignHeading{}, true)

	cs.processMessage(messageDetails{&BallotSignal{ballot}, node.ID()})

	select {
	case <-cs.metricsMessageBuffer:
		t.Errorf("REDACTED")
	case <-time.After(50 * time.Millisecond):
	}
}

func VerifyAttestIdenticalBallotTwice(t *testing.T) {
	_, vss := randomStatus(2)

	randomOctets := engineseed.Octets(comethash.Volume)

	ballot := attestBallot(vss[1],
		engineproto.PreendorseKind,
		randomOctets,
		kinds.SegmentAssignHeading{Sum: 10, Digest: randomOctets},
		true,
	)

	ballot2 := attestBallot(vss[1],
		engineproto.PreendorseKind,
		randomOctets,
		kinds.SegmentAssignHeading{Sum: 10, Digest: randomOctets},
		true,
	)

	require.Equal(t, ballot, ballot2)
}

//
func enrol(eventBus *kinds.EventBus, q cometbroadcast.Inquire) <-chan cometbroadcast.Signal {
	sub, err := eventBus.Enrol(context.Background(), verifyEnrollee, q)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", verifyEnrollee, q))
	}
	return sub.Out()
}

//
func enrolUnCached(eventBus *kinds.EventBus, q cometbroadcast.Inquire) <-chan cometbroadcast.Signal {
	sub, err := eventBus.EnrolUnbuffered(context.Background(), verifyEnrollee, q)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", verifyEnrollee, q))
	}
	return sub.Out()
}

func attestAppendPreendorseWithAddition(
	t *testing.T,
	cs *Status,
	digest []byte,
	heading kinds.SegmentAssignHeading,
	addition []byte,
	proxy *ratifierProxy,
) {
	v, err := proxy.attestBallot(engineproto.PreendorseKind, digest, heading, addition, true)
	require.NoError(t, err, "REDACTED")
	appendBallots(cs, v)
}

func locateLedgerVolumeCeiling(t *testing.T, level, maximumOctets int64, cs *Status, segmentVolume uint32, excessive bool) (*kinds.Ledger, *kinds.SegmentCollection) {
	var displacement int64
	if !excessive {
		displacement = -2
	}
	gentleMaximumDataOctets := int(kinds.MaximumDataOctets(maximumOctets, 0, 0))
	for i := gentleMaximumDataOctets; i < gentleMaximumDataOctets*2; i++ {
		nominationLedger, err := cs.status.CreateLedger(
			level,
			[]kinds.Tx{[]byte("REDACTED" + strings.Repeat("REDACTED", i-2))},
			&kinds.Endorse{},
			nil,
			cs.privateRatifierPublicKey.Location(),
		)
		require.NoError(t, err)

		nominationLedgerSegments, err := nominationLedger.CreateSegmentAssign(segmentVolume)
		require.NoError(t, err)
		if nominationLedgerSegments.OctetVolume() > maximumOctets+displacement {
			s := "REDACTED"
			if excessive {
				s = "REDACTED"
			}
			t.Log("REDACTED"+s+"REDACTED", "REDACTED", i, "REDACTED", gentleMaximumDataOctets)
			return nominationLedger, nominationLedgerSegments
		}
	}
	require.Fail(t, "REDACTED")
	return nil, nil
}
