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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	abcistubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds/simulations"
	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	nodestub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulate"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
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

func VerifyStatusNominatorOption0(t *testing.T) {
	cs1, vss := arbitraryStatus(4)
	altitude, iteration := cs1.Altitude, cs1.Iteration

	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)

	initiateVerifyIteration(cs1, altitude, iteration)

	//
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	//
	item := cs1.ObtainIterationStatus().Assessors.ObtainNominator()
	pv, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv.Location()
	if !bytes.Equal(item.Location, location) {
		t.Fatalf("REDACTED", 0, item.Location)
	}

	//
	assureFreshNomination(nominationChnl, altitude, iteration)

	rs := cs1.ObtainIterationStatus()
	attestAppendBallots(cs1, commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), true, vss[1:]...)

	//
	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	item = cs1.ObtainIterationStatus().Assessors.ObtainNominator()
	pv1, err := vss[1].ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	if !bytes.Equal(item.Location, location) {
		panic(fmt.Sprintf("REDACTED", 1, item.Location))
	}
}

//
func VerifyStatusNominatorOption2(t *testing.T) {
	cs1, vss := arbitraryStatus(4) //
	altitude := cs1.Altitude
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)

	//
	advanceIteration(vss[1:]...)
	advanceIteration(vss[1:]...)

	var iteration int32 = 2
	initiateVerifyIteration(cs1, altitude, iteration)

	assureFreshIteration(freshIterationChnl, altitude, iteration) //

	//
	for i := int32(0); int(i) < len(vss); i++ {
		item := cs1.ObtainIterationStatus().Assessors.ObtainNominator()
		pvk, err := vss[int(i+iteration)%len(vss)].ObtainPublicToken()
		require.NoError(t, err)
		location := pvk.Location()
		preciseNominator := location
		if !bytes.Equal(item.Location, preciseNominator) {
			panic(fmt.Sprintf(
				"REDACTED",
				int(i+2)%len(vss),
				item.Location))
		}

		rs := cs1.ObtainIterationStatus()
		attestAppendBallots(cs1, commitchema.PreendorseKind, nil, rs.NominationLedgerFragments.Heading(), true, vss[1:]...)
		assureFreshIteration(freshIterationChnl, altitude, i+iteration+1) //
		advanceIteration(vss[1:]...)
	}
}

//
func VerifyStatusJoinNominateNegativePrivateAssessor(t *testing.T) {
	cs, _ := arbitraryStatus(1)
	cs.AssignPrivateAssessor(nil)
	altitude, iteration := cs.Altitude, cs.Iteration

	//
	deadlineChnl := listen(cs.incidentChannel, kinds.IncidentInquireDeadlineNominate)

	initiateVerifyIteration(cs, altitude, iteration)

	//
	assureFreshDeadline(deadlineChnl, altitude, iteration, cs.settings.DeadlineNominate.Nanoseconds())

	if cs.ObtainIterationStatus().Nomination != nil {
		t.Error("REDACTED")
	}
}

//
func VerifyStatusJoinNominateTruePrivateAssessor(t *testing.T) {
	cs, _ := arbitraryStatus(1)
	altitude, iteration := cs.Altitude, cs.Iteration

	//

	deadlineChnl := listen(cs.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	nominationChnl := listen(cs.incidentChannel, kinds.IncidentInquireFinishNomination)

	cs.joinFreshIteration(altitude, iteration)
	cs.initiateThreads(3)

	assureFreshNomination(nominationChnl, altitude, iteration)

	//
	rs := cs.ObtainIterationStatus()
	if rs.Nomination == nil {
		t.Error("REDACTED")
	}
	if rs.NominationLedger == nil {
		t.Error("REDACTED")
	}
	if rs.NominationLedgerFragments.Sum() == 0 {
		t.Error("REDACTED")
	}

	//
	assureNegativeFreshDeadline(deadlineChnl, cs.settings.DeadlineNominate.Nanoseconds())
}

func VerifyStatusFlawedNomination(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(2)
	altitude, iteration := cs1.Altitude, cs1.Iteration
	vs2 := vss[1]

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	ballotChnl := listen(cs1.incidentChannel, kinds.IncidentInquireBallot)

	itemLedger, err := cs1.generateNominationLedger(ctx) //
	require.NoError(t, err)

	//
	iteration++
	advanceIteration(vss[1:]...)

	//
	statusDigest := itemLedger.PlatformDigest
	if len(statusDigest) == 0 {
		statusDigest = make([]byte, 32)
	}
	statusDigest[0] = (statusDigest[0] + 1) % 255
	itemLedger.PlatformDigest = statusDigest
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: itemLedger.Digest(), FragmentAssignHeading: itemLedgerFragments.Heading()}
	nomination := kinds.FreshNomination(vs2.Altitude, iteration, -1, ledgerUUID)
	p := nomination.TowardSchema()
	if err := vs2.AttestNomination(cs1.status.SuccessionUUID, p); err != nil {
		t.Fatal("REDACTED", err)
	}

	nomination.Notation = p.Notation

	//
	if err := cs1.AssignNominationAlsoLedger(nomination, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	//
	initiateVerifyIteration(cs1, altitude, iteration)

	//
	assureNomination(nominationChnl, altitude, iteration, ledgerUUID)

	//
	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], nil)

	//
	bps, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedger.Digest(), bps.Heading(), false, vs2)
	assurePreballot(ballotChnl, altitude, iteration)

	//
	assurePreendorse(ballotChnl, altitude, iteration)
	certifyPreendorse(t, cs1, iteration, -1, vss[0], nil, nil)

	bpp2, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedger.Digest(), bpp2.Heading(), true, vs2)
}

func VerifyStatusBulkyLedger(t *testing.T) {
	const maximumOctets = int64(kinds.LedgerFragmentExtentOctets)

	for _, verifyInstance := range []struct {
		alias      string
		bulky bool
	}{
		{
			alias:      "REDACTED",
			bulky: false,
		},
		{
			alias:      "REDACTED",
			bulky: true,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			cs1, vss := arbitraryStatus(2)
			cs1.status.AgreementSettings.Ledger.MaximumOctets = maximumOctets
			altitude, iteration := cs1.Altitude, cs1.Iteration
			vs2 := vss[1]

			fragmentExtent := kinds.LedgerFragmentExtentOctets

			itemLedger, itemLedgerFragments := locateLedgerExtentThreshold(t, altitude, maximumOctets, cs1, fragmentExtent, verifyInstance.bulky)

			deadlineNominateChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
			ballotChnl := listen(cs1.incidentChannel, kinds.IncidentInquireBallot)

			//
			iteration++
			advanceIteration(vss[1:]...)

			ledgerUUID := kinds.LedgerUUID{Digest: itemLedger.Digest(), FragmentAssignHeading: itemLedgerFragments.Heading()}
			nomination := kinds.FreshNomination(altitude, iteration, -1, ledgerUUID)
			p := nomination.TowardSchema()
			if err := vs2.AttestNomination(cs1.status.SuccessionUUID, p); err != nil {
				t.Fatal("REDACTED", err)
			}
			nomination.Notation = p.Notation

			sumOctets := 0
			for i := 0; i < int(itemLedgerFragments.Sum()); i++ {
				fragment := itemLedgerFragments.ObtainFragment(i)
				sumOctets += len(fragment.Octets)
			}

			maximumLedgerFragments := maximumOctets / int64(kinds.LedgerFragmentExtentOctets)
			if maximumOctets > maximumLedgerFragments*int64(kinds.LedgerFragmentExtentOctets) {
				maximumLedgerFragments++
			}
			countLedgerFragments := int64(itemLedgerFragments.Sum())

			if err := cs1.AssignNominationAlsoLedger(nomination, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
				t.Fatal(err)
			}

			//
			initiateVerifyIteration(cs1, altitude, iteration)

			t.Log("REDACTED", "REDACTED", maximumOctets, "REDACTED", sumOctets)
			t.Log("REDACTED", "REDACTED", maximumLedgerFragments, "REDACTED", countLedgerFragments)

			certifyDigest := itemLedger.Digest()
			securedIteration := int32(1)
			if verifyInstance.bulky {
				certifyDigest = nil
				securedIteration = -1
				//
				//
				assureFreshDeadline(deadlineNominateChnl, altitude, iteration, cs1.settings.Nominate(iteration).Nanoseconds())
				//
				//
			}
			assurePreballot(ballotChnl, altitude, iteration)
			certifyPreballot(t, cs1, iteration, vss[0], certifyDigest)

			//
			if countLedgerFragments > maximumLedgerFragments {
				require.Nil(t, cs1.Nomination)
			}

			bps, err := itemLedger.CreateFragmentAssign(fragmentExtent)
			require.NoError(t, err)

			attestAppendBallots(cs1, commitchema.PreballotKind, itemLedger.Digest(), bps.Heading(), false, vs2)
			assurePreballot(ballotChnl, altitude, iteration)
			assurePreendorse(ballotChnl, altitude, iteration)
			certifyPreendorse(t, cs1, iteration, securedIteration, vss[0], certifyDigest, certifyDigest)

			bpp2, err := itemLedger.CreateFragmentAssign(fragmentExtent)
			require.NoError(t, err)
			attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedger.Digest(), bpp2.Heading(), true, vs2)
		})
	}
}

//
//

//
func VerifyStatusCompleteCycle1(t *testing.T) {
	cs, vss := arbitraryStatus(1)
	altitude, iteration := cs.Altitude, cs.Iteration

	//
	//
	if err := cs.incidentChannel.Halt(); err != nil {
		t.Error(err)
	}
	incidentChannel := kinds.FreshIncidentPipelineUsingReserveVolume(0)
	incidentChannel.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	cs.AssignIncidentChannel(incidentChannel)
	if err := incidentChannel.Initiate(); err != nil {
		t.Error(err)
	}

	ballotChnl := listenNegCached(cs.incidentChannel, kinds.IncidentInquireBallot)
	itemChnl := listen(cs.incidentChannel, kinds.IncidentInquireFinishNomination)
	freshIterationChnl := listen(cs.incidentChannel, kinds.IncidentInquireFreshIteration)

	//
	initiateVerifyIteration(cs, altitude, iteration)

	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(itemChnl, altitude, iteration)
	itemLedgerDigest := cs.ObtainIterationStatus().NominationLedger.Digest()

	assurePreballot(ballotChnl, altitude, iteration) //
	certifyPreballot(t, cs, iteration, vss[0], itemLedgerDigest)

	assurePreendorse(ballotChnl, altitude, iteration) //

	//
	assureFreshIteration(freshIterationChnl, altitude+1, 0)

	certifyFinalPreendorse(t, cs, vss[0], itemLedgerDigest)
}

//
func VerifyStatusCompleteIterationVoid(t *testing.T) {
	cs, _ := arbitraryStatus(1)
	altitude, iteration := cs.Altitude, cs.Iteration

	ballotChnl := listenNegCached(cs.incidentChannel, kinds.IncidentInquireBallot)

	cs.joinPreballot(altitude, iteration)
	cs.initiateThreads(4)

	assurePreballotAlign(t, ballotChnl, altitude, iteration, nil)   //
	assurePreendorseAlign(t, ballotChnl, altitude, iteration, nil) //
}

//
//
func VerifyStatusCompleteCycle2(t *testing.T) {
	cs1, vss := arbitraryStatus(2)
	vs2 := vss[1]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	ballotChnl := listenNegCached(cs1.incidentChannel, kinds.IncidentInquireBallot)
	freshLedgerChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshLedger)

	//
	initiateVerifyIteration(cs1, altitude, iteration)

	assurePreballot(ballotChnl, altitude, iteration) //

	//
	rs := cs1.ObtainIterationStatus()
	itemLedgerDigest, itemFragmentAssignHeading := rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading()

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest, itemFragmentAssignHeading, false, vs2)
	assurePreballot(ballotChnl, altitude, iteration) //

	assurePreendorse(ballotChnl, altitude, iteration) //
	//
	certifyPreendorse(t, cs1, 0, 0, vss[0], itemLedgerDigest, itemLedgerDigest)

	//

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedgerDigest, itemFragmentAssignHeading, true, vs2)
	assurePreendorse(ballotChnl, altitude, iteration)

	//
	assureFreshLedger(freshLedgerChnl, altitude)
}

//
//

//
//
func VerifyStatusSecureNegativePolicy(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(2)
	vs2 := vss[1]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	deadlineNominateChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	ballotChnl := listenNegCached(cs1.incidentChannel, kinds.IncidentInquireBallot)
	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)

	/**
2
*/

	//
	cs1.joinFreshIteration(altitude, iteration)
	cs1.initiateThreads(0)

	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	iterationStatus := cs1.ObtainIterationStatus()
	thatLedgerDigest := iterationStatus.NominationLedger.Digest()
	thatFragmentAssignHeading := iterationStatus.NominationLedgerFragments.Heading()

	assurePreballot(ballotChnl, altitude, iteration) //

	//
	//
	attestAppendBallots(cs1, commitchema.PreballotKind, thatLedgerDigest, thatFragmentAssignHeading, false, vs2)
	assurePreballot(ballotChnl, altitude, iteration) //

	assurePreendorse(ballotChnl, altitude, iteration) //
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], thatLedgerDigest, thatLedgerDigest)

	//
	//
	digest := make([]byte, len(thatLedgerDigest))
	copy(digest, thatLedgerDigest)
	digest[0] = (digest[0] + 1) % 255
	attestAppendBallots(cs1, commitchema.PreendorseKind, digest, thatFragmentAssignHeading, true, vs2)
	assurePreendorse(ballotChnl, altitude, iteration) //

	//
	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	//

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")
	/**
2
*/

	advanceIteration(vs2)

	//
	assureFreshDeadline(deadlineNominateChnl, altitude, iteration, cs1.settings.Nominate(iteration).Nanoseconds())

	rs := cs1.ObtainIterationStatus()

	require.Nil(t, rs.NominationLedger, "REDACTED")

	//
	assurePreballot(ballotChnl, altitude, iteration)
	//
	certifyPreballot(t, cs1, iteration, vss[0], rs.SecuredLedger.Digest())

	//
	bps, err := rs.SecuredLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	attestAppendBallots(cs1, commitchema.PreballotKind, digest, bps.Heading(), false, vs2)
	assurePreballot(ballotChnl, altitude, iteration)

	//
	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preballot(iteration).Nanoseconds())

	assurePreendorse(ballotChnl, altitude, iteration) //
	//
	//
	certifyPreendorse(t, cs1, iteration, 0, vss[0], nil, thatLedgerDigest)

	//
	bpp2, err := rs.SecuredLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	attestAppendBallots(cs1, commitchema.PreendorseKind, digest, bpp2.Heading(), true, vs2)
	assurePreendorse(ballotChnl, altitude, iteration)

	//
	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")
	/**
2
*/

	advanceIteration(vs2)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs = cs1.ObtainIterationStatus()

	//
	if !bytes.Equal(rs.NominationLedger.Digest(), rs.SecuredLedger.Digest()) {
		panic(fmt.Sprintf(
			"REDACTED",
			rs.NominationLedger,
			rs.SecuredLedger))
	}

	assurePreballot(ballotChnl, altitude, iteration) //
	certifyPreballot(t, cs1, iteration, vss[0], rs.SecuredLedger.Digest())

	bpp0, err := rs.NominationLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	attestAppendBallots(cs1, commitchema.PreballotKind, digest, bpp0.Heading(), false, vs2)
	assurePreballot(ballotChnl, altitude, iteration)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preballot(iteration).Nanoseconds())
	assurePreendorse(ballotChnl, altitude, iteration) //

	certifyPreendorse(t, cs1, iteration, 0, vss[0], nil, thatLedgerDigest) //

	bpp1, err := rs.NominationLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	attestAppendBallots(
		cs1,
		commitchema.PreendorseKind,
		digest,
		bpp1.Heading(),
		true,
		vs2) //
	assurePreendorse(ballotChnl, altitude, iteration)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	cs2, _ := arbitraryStatus(2) //
	//
	item, itemLedger := resolveNomination(ctx, t, cs2, vs2, vs2.Altitude, vs2.Iteration+1)
	if item == nil || itemLedger == nil {
		t.Fatal("REDACTED")
	}

	advanceIteration(vs2)

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")
	/**
C
*/

	//
	//
	bpp3, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, bpp3, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureFreshNomination(nominationChnl, altitude, iteration)
	assurePreballot(ballotChnl, altitude, iteration) //
	//
	certifyPreballot(t, cs1, 3, vss[0], cs1.SecuredLedger.Digest())

	//
	bpp4, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedger.Digest(), bpp4.Heading(), false, vs2)
	assurePreballot(ballotChnl, altitude, iteration)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preballot(iteration).Nanoseconds())
	assurePreendorse(ballotChnl, altitude, iteration)
	certifyPreendorse(t, cs1, iteration, 0, vss[0], nil, thatLedgerDigest) //

	bpp5, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	attestAppendBallots(
		cs1,
		commitchema.PreendorseKind,
		itemLedger.Digest(),
		bpp5.Heading(),
		true,
		vs2) //
	assurePreendorse(ballotChnl, altitude, iteration)
}

//
//
//
//
func VerifyStatusSecurePolicyResecure(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	freshLedgerChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshLedgerHeading)

	//

	/**
l

s
*/

	//
	initiateVerifyIteration(cs1, altitude, iteration)

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	thatLedgerDigest := rs.NominationLedger.Digest()
	thatLedgerFragments := rs.NominationLedgerFragments.Heading()

	assurePreballot(ballotChnl, altitude, iteration) //

	attestAppendBallots(cs1, commitchema.PreballotKind, thatLedgerDigest, thatLedgerFragments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration) //
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], thatLedgerDigest, thatLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	cs2 := freshStatus(cs1.status, vs2, statedepot.FreshInsideRamPlatform())
	item, itemLedger := resolveNomination(ctx, t, cs2, vs2, vs2.Altitude, vs2.Iteration+1)
	if item == nil || itemLedger == nil {
		t.Fatal("REDACTED")
	}
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	itemLedgerDigest := itemLedger.Digest()
	require.NotEqual(t, itemLedgerDigest, thatLedgerDigest)

	advanceIteration(vs2, vs3, vs4)

	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //
	//
	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")

	/**
)

!
*/

	//
	//
	assureFreshNomination(nominationChnl, altitude, iteration)

	//
	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], thatLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest, itemLedgerFragments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], itemLedgerDigest, itemLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedgerDigest, itemLedgerFragments.Heading(), true, vs2, vs3)
	assureFreshLedgerHeading(freshLedgerChnl, altitude, itemLedgerDigest)

	assureFreshIteration(freshIterationChnl, altitude+1, 0)
}

//
func VerifyStatusSecurePolicyRelease(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	releaseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireRelease)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//

	/**
l
s
*/

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	thatLedgerDigest := rs.NominationLedger.Digest()
	thatLedgerFragments := rs.NominationLedgerFragments.Heading()

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], thatLedgerDigest)

	attestAppendBallots(cs1, commitchema.PreballotKind, thatLedgerDigest, thatLedgerFragments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], thatLedgerDigest, thatLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs4)
	attestAppendBallots(cs1, commitchema.PreendorseKind, thatLedgerDigest, thatLedgerFragments, true, vs3)

	//
	item, itemLedger := resolveNomination(ctx, t, cs1, vs2, vs2.Altitude, vs2.Iteration+1)
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())
	rs = cs1.ObtainIterationStatus()
	securedLedgerDigest := rs.SecuredLedger.Digest()

	advanceIteration(vs2, vs3, vs4)
	iteration++ //

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")
	/**
_
!
*/
	//
	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureFreshNomination(nominationChnl, altitude, iteration)

	//
	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], securedLedgerDigest)
	//
	attestAppendBallots(cs1, commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, false, vs2, vs3, vs4)

	//
	assureFreshRelease(releaseChnl, altitude, iteration)
	assurePreendorse(ballotChnl, altitude, iteration)

	//
	//
	certifyPreendorse(t, cs1, iteration, -1, vss[0], nil, nil)

	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3)
	assureFreshIteration(freshIterationChnl, altitude, iteration+1)
}

//
//
//
//
func VerifyStatusSecurePolicyReleaseUponUnfamiliarLedger(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	//

	/**
l
*/

	//
	initiateVerifyIteration(cs1, altitude, iteration)

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	initialLedgerDigest := rs.NominationLedger.Digest()
	initialLedgerFragments := rs.NominationLedgerFragments.Heading()

	assurePreballot(ballotChnl, altitude, iteration) //

	attestAppendBallots(cs1, commitchema.PreballotKind, initialLedgerDigest, initialLedgerFragments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration) //
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], initialLedgerDigest, initialLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	cs2 := freshStatus(cs1.status, vs2, statedepot.FreshInsideRamPlatform())
	item, itemLedger := resolveNomination(ctx, t, cs2, vs2, vs2.Altitude, vs2.Iteration+1)
	if item == nil || itemLedger == nil {
		t.Fatal("REDACTED")
	}
	ordinalLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	ordinalLedgerDigest := itemLedger.Digest()
	require.NotEqual(t, ordinalLedgerDigest, initialLedgerDigest)

	advanceIteration(vs2, vs3, vs4)

	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")

	/**
)
*/

	//

	//
	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], initialLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, ordinalLedgerDigest, ordinalLedgerFragments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, -1, vss[0], nil, nil)

	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, ordinalLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	cs3 := freshStatus(cs1.status, vs3, statedepot.FreshInsideRamPlatform())
	item, itemLedger = resolveNomination(ctx, t, cs3, vs3, vs3.Altitude, vs3.Iteration+1)
	if item == nil || itemLedger == nil {
		t.Fatal("REDACTED")
	}
	tertiaryItemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	tertiaryItemLedgerDigest := itemLedger.Digest()
	require.NotEqual(t, ordinalLedgerDigest, tertiaryItemLedgerDigest)

	advanceIteration(vs2, vs3, vs4)

	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")

	/**
)
*/

	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, tertiaryItemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assurePreballot(ballotChnl, altitude, iteration)
	//
	certifyPreballot(t, cs1, iteration, vss[0], tertiaryItemLedgerDigest)

	attestAppendBallots(cs1, commitchema.PreballotKind, tertiaryItemLedgerDigest, tertiaryItemLedgerFragments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], tertiaryItemLedgerDigest, tertiaryItemLedgerDigest)
}

//
//
//
//
func VerifyStatusSecurePolicySecurity1(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	deadlineNominateChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, cs1.Altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	itemLedger := rs.NominationLedger

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], itemLedger.Digest())

	//
	bps, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	preballots := attestBallots(commitchema.PreballotKind, itemLedger.Digest(), bps.Heading(), false, vs2, vs3, vs4)

	t.Logf("REDACTED", fmt.Sprintf("REDACTED", itemLedger.Digest()))

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	//
	assurePreendorse(ballotChnl, altitude, iteration)
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	t.Log("REDACTED")

	item, itemLedger := resolveNomination(ctx, t, cs1, vs2, vs2.Altitude, vs2.Iteration+1)
	itemLedgerDigest := itemLedger.Digest()
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	advanceIteration(vs2, vs3, vs4)

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	//
	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	/*2
k
!
*/

	assureFreshNomination(nominationChnl, altitude, iteration)

	rs = cs1.ObtainIterationStatus()

	if rs.SecuredLedger != nil {
		panic("REDACTED")
	}
	t.Logf("REDACTED", fmt.Sprintf("REDACTED", itemLedgerDigest))

	//
	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], itemLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest, itemLedgerFragments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], itemLedgerDigest, itemLedgerDigest)

	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	advanceIteration(vs2, vs3, vs4)
	iteration++ //

	assureFreshIteration(freshIterationChnl, altitude, iteration)

	t.Log("REDACTED")
	/*3
!
*/

	//
	assureFreshDeadline(deadlineNominateChnl, altitude, iteration, cs1.settings.Nominate(iteration).Nanoseconds())

	//
	assurePreballot(ballotChnl, altitude, iteration)
	//
	certifyPreballot(t, cs1, iteration, vss[0], itemLedgerDigest)

	freshPhaseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIterationPhase)

	//
	//
	appendBallots(cs1, preballots...)

	t.Log("REDACTED")

	assureNegativeFreshIterationPhase(freshPhaseChnl)
}

//
//
//
//

//
//
func VerifyStatusSecurePolicySecurity2(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	releaseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireRelease)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	//
	_, itemLedger0 := resolveNomination(ctx, t, cs1, vss[0], altitude, iteration)
	itemLedgerDigest0 := itemLedger0.Digest()
	itemLedgerFragments0, err := itemLedger0.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	itemLedgerUuid0 := kinds.LedgerUUID{Digest: itemLedgerDigest0, FragmentAssignHeading: itemLedgerFragments0.Heading()}

	//
	preballots := attestBallots(commitchema.PreballotKind, itemLedgerDigest0, itemLedgerFragments0.Heading(), false, vs2, vs3, vs4)

	//
	item1, itemLedger1 := resolveNomination(ctx, t, cs1, vs2, vs2.Altitude, vs2.Iteration+1)
	itemLedgerDigest1 := itemLedger1.Digest()
	itemLedgerFragments1, err := itemLedger1.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	advanceIteration(vs2, vs3, vs4)

	iteration++ //
	t.Log("REDACTED")
	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	if err := cs1.AssignNominationAlsoLedger(item1, itemLedger1, itemLedgerFragments1, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureFreshNomination(nominationChnl, altitude, iteration)

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], itemLedgerDigest1)

	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest1, itemLedgerFragments1.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], itemLedgerDigest1, itemLedgerDigest1)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs4)
	attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedgerDigest1, itemLedgerFragments1.Heading(), true, vs3)

	advanceIteration(vs2, vs3, vs4)

	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //
	//
	freshItem := kinds.FreshNomination(altitude, iteration, 0, itemLedgerUuid0)
	p := freshItem.TowardSchema()
	if err := vs3.AttestNomination(cs1.status.SuccessionUUID, p); err != nil {
		t.Fatal(err)
	}

	freshItem.Notation = p.Notation

	if err := cs1.AssignNominationAlsoLedger(freshItem, itemLedger0, itemLedgerFragments0, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	//
	appendBallots(cs1, preballots...)

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")
	/*2
k
*/
	assureFreshNomination(nominationChnl, altitude, iteration)

	assureNegativeFreshRelease(releaseChnl)
	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], itemLedgerDigest1)
}

//
//

//
//
func VerifyNominateSoundLedger(t *testing.T) {
	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	deadlineNominateChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	releaseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireRelease)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, cs1.Altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	itemLedger := rs.NominationLedger
	itemLedgerDigest := itemLedger.Digest()

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], itemLedgerDigest)

	//
	bps, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)
	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest, bps.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], itemLedgerDigest, itemLedgerDigest)

	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	advanceIteration(vs2, vs3, vs4)
	iteration++ //

	assureFreshIteration(freshIterationChnl, altitude, iteration)

	t.Log("REDACTED")

	//
	assureFreshDeadline(deadlineNominateChnl, altitude, iteration, cs1.settings.Nominate(iteration).Nanoseconds())

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], itemLedgerDigest)

	attestAppendBallots(cs1, commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, false, vs2, vs3, vs4)

	assureFreshRelease(releaseChnl, altitude, iteration)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, -1, vss[0], nil, nil)

	advanceIteration(vs2, vs3, vs4)
	advanceIteration(vs2, vs3, vs4)

	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	iteration += 2 //

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	t.Log("REDACTED")

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //

	assureFreshIteration(freshIterationChnl, altitude, iteration)

	t.Log("REDACTED")

	assureFreshNomination(nominationChnl, altitude, iteration)

	rs = cs1.ObtainIterationStatus()
	assert.True(t, bytes.Equal(rs.NominationLedger.Digest(), itemLedgerDigest))
	assert.True(t, bytes.Equal(rs.NominationLedger.Digest(), rs.SoundLedger.Digest()))
	assert.True(t, rs.Nomination.PolicyIteration == rs.SoundIteration)
	assert.True(t, bytes.Equal(rs.Nomination.LedgerUUID.Digest, rs.SoundLedger.Digest()))
}

//
//
func VerifyAssignSoundLedgerUponPostponedPreballot(t *testing.T) {
	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	soundLedgerChnl := listen(cs1.incidentChannel, kinds.IncidentInquireSoundLedger)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, cs1.Altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	itemLedger := rs.NominationLedger
	itemLedgerDigest := itemLedger.Digest()
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], itemLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest, itemLedgerFragments.Heading(), false, vs2)

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, false, vs3)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preballot(iteration).Nanoseconds())

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, -1, vss[0], nil, nil)

	rs = cs1.ObtainIterationStatus()

	assert.True(t, rs.SoundLedger == nil)
	assert.True(t, rs.SoundLedgerFragments == nil)
	assert.True(t, rs.SoundIteration == -1)

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest, itemLedgerFragments.Heading(), false, vs4)

	assureFreshSoundLedger(soundLedgerChnl, altitude, iteration)

	rs = cs1.ObtainIterationStatus()

	assert.True(t, bytes.Equal(rs.SoundLedger.Digest(), itemLedgerDigest))
	assert.True(t, rs.SoundLedgerFragments.Heading().Matches(itemLedgerFragments.Heading()))
	assert.True(t, rs.SoundIteration == iteration)
}

//
//
//
func VerifyAssignSoundLedgerUponPostponedNomination(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	deadlineNominateChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	soundLedgerChnl := listen(cs1.incidentChannel, kinds.IncidentInquireSoundLedger)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)
	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)

	iteration++ //
	advanceIteration(vs2, vs3, vs4)

	initiateVerifyIteration(cs1, cs1.Altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshDeadline(deadlineNominateChnl, altitude, iteration, cs1.settings.Nominate(iteration).Nanoseconds())

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], nil)

	item, itemLedger := resolveNomination(ctx, t, cs1, vs2, vs2.Altitude, vs2.Iteration+1)
	itemLedgerDigest := itemLedger.Digest()
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	//
	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedgerDigest, itemLedgerFragments.Heading(), false, vs2, vs3, vs4)
	assureFreshSoundLedger(soundLedgerChnl, altitude, iteration)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preballot(iteration).Nanoseconds())

	assurePreendorse(ballotChnl, altitude, iteration)
	certifyPreendorse(t, cs1, iteration, -1, vss[0], nil, nil)

	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()

	assert.True(t, bytes.Equal(rs.SoundLedger.Digest(), itemLedgerDigest))
	assert.True(t, rs.SoundLedgerFragments.Heading().Matches(itemLedgerFragments.Heading()))
	assert.True(t, rs.SoundIteration == iteration)
}

func VerifyHandleNominationEmbrace(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias               string
		embrace             bool
		anticipatedVoidPreballot bool
	}{
		{
			alias:               "REDACTED",
			embrace:             true,
			anticipatedVoidPreballot: false,
		},
		{
			alias:               "REDACTED",
			embrace:             false,
			anticipatedVoidPreballot: true,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			m := abcistubs.FreshPlatform(t)
			condition := iface.Responseexecuteitem_DECLINE
			if verifyInstance.embrace {
				condition = iface.Responseexecuteitem_EMBRACE
			}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Condition: condition}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil).Maybe()
			cs1, _ := arbitraryStatusUsingApplication(4, m)
			altitude, iteration := cs1.Altitude, cs1.Iteration

			nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
			freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
			pv1, err := cs1.privateAssessor.ObtainPublicToken()
			require.NoError(t, err)
			location := pv1.Location()
			ballotChnl := listenTowardBalloter(cs1, location)

			initiateVerifyIteration(cs1, cs1.Altitude, iteration)
			assureFreshIteration(freshIterationChnl, altitude, iteration)

			assureFreshNomination(nominationChnl, altitude, iteration)
			rs := cs1.ObtainIterationStatus()
			var preballotDigest tendermintoctets.HexadecimalOctets
			if !verifyInstance.anticipatedVoidPreballot {
				preballotDigest = rs.NominationLedger.Digest()
			}
			assurePreballotAlign(t, ballotChnl, altitude, iteration, preballotDigest)
		})
	}
}

//
//
func VerifyBroadenBallotInvokedWheneverActivated(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias    string
		activated bool
	}{
		{
			alias:    "REDACTED",
			activated: true,
		},
		{
			alias:    "REDACTED",
			activated: false,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			m := abcistubs.FreshPlatform(t)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_EMBRACE}, nil)
			if verifyInstance.activated {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyBroadenBallot{
					BallotAddition: []byte("REDACTED"),
				}, nil)
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
					Condition: iface.Responsecertifyballotaddition_EMBRACE,
				}, nil)
			}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCulminateLedger{}, nil).Maybe()
			altitude := int64(1)
			if !verifyInstance.activated {
				altitude = 0
			}
			cs1, vss := arbitraryStatusUsingApplicationUsingAltitude(4, m, altitude)

			altitude, iteration := cs1.Altitude, cs1.Iteration

			nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
			freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
			pv1, err := cs1.privateAssessor.ObtainPublicToken()
			require.NoError(t, err)
			location := pv1.Location()
			ballotChnl := listenTowardBalloter(cs1, location)

			initiateVerifyIteration(cs1, cs1.Altitude, iteration)
			assureFreshIteration(freshIterationChnl, altitude, iteration)
			assureFreshNomination(nominationChnl, altitude, iteration)

			m.AssertNotCalled(t, "REDACTED", mock.Anything, mock.Anything)

			rs := cs1.ObtainIterationStatus()

			ledgerUUID := kinds.LedgerUUID{
				Digest:          rs.NominationLedger.Digest(),
				FragmentAssignHeading: rs.NominationLedgerFragments.Heading(),
			}
			attestAppendBallots(cs1, commitchema.PreballotKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, false, vss[1:]...)
			assurePreballotAlign(t, ballotChnl, altitude, iteration, ledgerUUID.Digest)

			assurePreendorse(ballotChnl, altitude, iteration)

			if verifyInstance.activated {
				m.AssertCalled(t, "REDACTED", context.TODO(), &iface.SolicitBroadenBallot{
					Altitude:             altitude,
					Digest:               ledgerUUID.Digest,
					Moment:               rs.NominationLedger.Moment,
					Txs:                rs.NominationLedger.Txs.TowardSegmentBelongingOctets(),
					ItemizedFinalEndorse: iface.EndorseDetails{},
					Malpractice:        rs.NominationLedger.Proof.Proof.TowardIface(),
					FollowingAssessorsDigest: rs.NominationLedger.FollowingAssessorsDigest,
					NominatorLocation:    rs.NominationLedger.NominatorLocation,
				})
			} else {
				m.AssertNotCalled(t, "REDACTED", mock.Anything, mock.Anything)
			}

			attestAppendBallots(cs1, commitchema.PreendorseKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, verifyInstance.activated, vss[1:]...)
			assureFreshIteration(freshIterationChnl, altitude+1, 0)
			m.AssertExpectations(t)

			//
			//
			for _, pv := range vss[1:3] {
				pv, err := pv.ObtainPublicToken()
				require.NoError(t, err)
				location := pv.Location()
				if verifyInstance.activated {
					m.AssertCalled(t, "REDACTED", context.TODO(), &iface.SolicitValidateBallotAddition{
						Digest:             ledgerUUID.Digest,
						AssessorLocation: location,
						Altitude:           altitude,
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
func VerifyValidateBallotAdditionNegationInvokedUponMissingPreendorse(t *testing.T) {
	m := abcistubs.FreshPlatform(t)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_EMBRACE}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyBroadenBallot{
		BallotAddition: []byte("REDACTED"),
	}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
		Condition: iface.Responsecertifyballotaddition_EMBRACE,
	}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCulminateLedger{}, nil).Maybe()
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
	cs1, vss := arbitraryStatusUsingApplication(4, m)
	altitude, iteration := cs1.Altitude, cs1.Iteration
	cs1.status.AgreementSettings.Iface.BallotAdditionsActivateAltitude = cs1.Altitude

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	initiateVerifyIteration(cs1, cs1.Altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()

	ledgerUUID := kinds.LedgerUUID{
		Digest:          rs.NominationLedger.Digest(),
		FragmentAssignHeading: rs.NominationLedgerFragments.Heading(),
	}
	attestAppendBallots(cs1, commitchema.PreballotKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, false, vss...)
	assurePreballotAlign(t, ballotChnl, altitude, iteration, ledgerUUID.Digest)

	assurePreendorse(ballotChnl, altitude, iteration)

	m.AssertCalled(t, "REDACTED", context.TODO(), &iface.SolicitBroadenBallot{
		Altitude:             altitude,
		Digest:               ledgerUUID.Digest,
		Moment:               rs.NominationLedger.Moment,
		Txs:                rs.NominationLedger.Txs.TowardSegmentBelongingOctets(),
		ItemizedFinalEndorse: iface.EndorseDetails{},
		Malpractice:        rs.NominationLedger.Proof.Proof.TowardIface(),
		FollowingAssessorsDigest: rs.NominationLedger.FollowingAssessorsDigest,
		NominatorLocation:    rs.NominationLedger.NominatorLocation,
	})

	attestAppendBallots(cs1, commitchema.PreendorseKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, true, vss[2:]...)
	assureFreshIteration(freshIterationChnl, altitude+1, 0)
	m.AssertExpectations(t)

	//
	//
	pv, err := vss[1].ObtainPublicToken()
	require.NoError(t, err)
	location = pv.Location()

	m.AssertNotCalled(t, "REDACTED", context.TODO(), &iface.SolicitValidateBallotAddition{
		Digest:             ledgerUUID.Digest,
		AssessorLocation: location,
		Altitude:           altitude,
		BallotAddition:    []byte("REDACTED"),
	})
}

//
//
//
//
//
//
func VerifyArrangeNominationAcceptsBallotAdditions(t *testing.T) {
	//
	ballotAdditions := [][]byte{
		[]byte("REDACTED"),
		[]byte("REDACTED"),
		[]byte("REDACTED"),
		[]byte("REDACTED"),
	}

	m := abcistubs.FreshPlatform(t)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyBroadenBallot{
		BallotAddition: ballotAdditions[0],
	}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_EMBRACE}, nil)

	//
	rpp := &iface.SolicitArrangeNomination{}
	m.On("REDACTED", mock.Anything, mock.MatchedBy(func(r *iface.SolicitArrangeNomination) bool {
		rpp = r
		return true
	})).Return(&iface.ReplyArrangeNomination{}, nil)

	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{Condition: iface.Responsecertifyballotaddition_EMBRACE}, nil)
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
	m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCulminateLedger{}, nil)

	cs1, vss := arbitraryStatusUsingApplication(4, m)
	altitude, iteration := cs1.Altitude, cs1.Iteration

	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	assureFreshNomination(nominationChnl, altitude, iteration)

	rs := cs1.ObtainIterationStatus()
	ledgerUUID := kinds.LedgerUUID{
		Digest:          rs.NominationLedger.Digest(),
		FragmentAssignHeading: rs.NominationLedgerFragments.Heading(),
	}
	attestAppendBallots(cs1, commitchema.PreballotKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, false, vss[1:]...)

	//
	for i, vs := range vss[1:] {
		attestAppendPreendorseUsingAddition(t, cs1, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, ballotAdditions[i+1], vs)
	}

	assurePreballot(ballotChnl, altitude, iteration)

	//
	assurePreendorseAlign(t, ballotChnl, altitude, iteration, ledgerUUID.Digest)
	advanceAltitude(vss[1:]...)

	altitude++
	iteration = 0
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	advanceIteration(vss[1:]...)
	advanceIteration(vss[1:]...)
	advanceIteration(vss[1:]...)
	iteration = 3

	ledgerUuid2 := kinds.LedgerUUID{}
	attestAppendBallots(cs1, commitchema.PreendorseKind, ledgerUuid2.Digest, ledgerUuid2.FragmentAssignHeading, true, vss[1:]...)
	assureFreshIteration(freshIterationChnl, altitude, iteration)
	assureFreshNomination(nominationChnl, altitude, iteration)

	//
	//
	require.Len(t, rpp.RegionalFinalEndorse.Ballots, len(vss))
	for i := range vss {
		ballot := &rpp.RegionalFinalEndorse.Ballots[i]
		require.Equal(t, ballot.BallotAddition, ballotAdditions[i])

		require.NotZero(t, len(ballot.AdditionNotation))
		cve := commitchema.StandardBallotAddition{
			Addition: ballot.BallotAddition,
			Altitude:    altitude - 1, //
			Iteration:     int64(rpp.RegionalFinalEndorse.Iteration),
			SuccessionUuid:   verify.FallbackVerifySuccessionUUID,
		}
		addnAttestOctets, err := protocolio.SerializeSeparated(&cve)
		require.NoError(t, err)
		publicToken, err := vss[i].ObtainPublicToken()
		require.NoError(t, err)
		require.True(t, publicToken.ValidateNotation(addnAttestOctets, ballot.AdditionNotation))
	}
}

func VerifyCulminateLedgerInvoked(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias         string
		ballotVoid      bool
		anticipateInvoked bool
	}{
		{
			alias:         "REDACTED",
			ballotVoid:      false,
			anticipateInvoked: true,
		},
		{
			alias:         "REDACTED",
			ballotVoid:      true,
			anticipateInvoked: false,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			m := abcistubs.FreshPlatform(t)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{
				Condition: iface.Responseexecuteitem_EMBRACE,
			}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			//
			//
			if !verifyInstance.ballotVoid {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyBroadenBallot{}, nil)
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
					Condition: iface.Responsecertifyballotaddition_EMBRACE,
				}, nil)
			}
			r := &iface.ReplyCulminateLedger{PlatformDigest: []byte("REDACTED")}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(r, nil).Maybe()
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()

			cs1, vss := arbitraryStatusUsingApplication(4, m)
			altitude, iteration := cs1.Altitude, cs1.Iteration

			nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
			freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
			pv1, err := cs1.privateAssessor.ObtainPublicToken()
			require.NoError(t, err)
			location := pv1.Location()
			ballotChnl := listenTowardBalloter(cs1, location)

			initiateVerifyIteration(cs1, cs1.Altitude, iteration)
			assureFreshIteration(freshIterationChnl, altitude, iteration)
			assureFreshNomination(nominationChnl, altitude, iteration)
			rs := cs1.ObtainIterationStatus()

			ledgerUUID := kinds.LedgerUUID{}
			followingIteration := iteration + 1
			followingAltitude := altitude
			if !verifyInstance.ballotVoid {
				followingIteration = 0
				followingAltitude = altitude + 1
				ledgerUUID = kinds.LedgerUUID{
					Digest:          rs.NominationLedger.Digest(),
					FragmentAssignHeading: rs.NominationLedgerFragments.Heading(),
				}
			}

			attestAppendBallots(cs1, commitchema.PreballotKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, false, vss[1:]...)
			assurePreballotAlign(t, ballotChnl, altitude, iteration, rs.NominationLedger.Digest())

			attestAppendBallots(cs1, commitchema.PreendorseKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, true, vss[1:]...)
			assurePreendorse(ballotChnl, altitude, iteration)

			assureFreshIteration(freshIterationChnl, followingAltitude, followingIteration)
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
func VerifyBallotAdditionActivateAltitude(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias                  string
		activateAltitude          int64
		ownsAddition          bool
		anticipateBroadenInvoked    bool
		anticipateValidateInvoked    bool
		anticipateTriumphantIteration bool
	}{
		{
			alias:                  "REDACTED",
			ownsAddition:          true,
			activateAltitude:          0,
			anticipateBroadenInvoked:    false,
			anticipateValidateInvoked:    false,
			anticipateTriumphantIteration: false,
		},
		{
			alias:                  "REDACTED",
			ownsAddition:          false,
			activateAltitude:          0,
			anticipateBroadenInvoked:    false,
			anticipateValidateInvoked:    false,
			anticipateTriumphantIteration: true,
		},
		{
			alias:                  "REDACTED",
			ownsAddition:          true,
			activateAltitude:          1,
			anticipateBroadenInvoked:    true,
			anticipateValidateInvoked:    true,
			anticipateTriumphantIteration: true,
		},
		{
			alias:                  "REDACTED",
			ownsAddition:          false,
			activateAltitude:          1,
			anticipateBroadenInvoked:    true,
			anticipateValidateInvoked:    false,
			anticipateTriumphantIteration: false,
		},
		{
			alias:                  "REDACTED",
			ownsAddition:          false,
			activateAltitude:          2,
			anticipateBroadenInvoked:    false,
			anticipateValidateInvoked:    false,
			anticipateTriumphantIteration: true,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			countAssessors := 3
			m := abcistubs.FreshPlatform(t)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{
				Condition: iface.Responseexecuteitem_EMBRACE,
			}, nil)
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			if verifyInstance.anticipateBroadenInvoked {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyBroadenBallot{}, nil)
			}
			if verifyInstance.anticipateValidateInvoked {
				m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyValidateBallotAddition{
					Condition: iface.Responsecertifyballotaddition_EMBRACE,
				}, nil).Times(countAssessors - 1)
			}
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyCulminateLedger{}, nil).Maybe()
			m.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyEndorse{}, nil).Maybe()
			cs1, vss := arbitraryStatusUsingApplicationUsingAltitude(countAssessors, m, verifyInstance.activateAltitude)
			cs1.status.AgreementSettings.Iface.BallotAdditionsActivateAltitude = verifyInstance.activateAltitude
			altitude, iteration := cs1.Altitude, cs1.Iteration

			deadlineChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
			nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
			freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
			pv1, err := cs1.privateAssessor.ObtainPublicToken()
			require.NoError(t, err)
			location := pv1.Location()
			ballotChnl := listenTowardBalloter(cs1, location)

			initiateVerifyIteration(cs1, cs1.Altitude, iteration)
			assureFreshIteration(freshIterationChnl, altitude, iteration)
			assureFreshNomination(nominationChnl, altitude, iteration)
			rs := cs1.ObtainIterationStatus()

			//
			attestAppendBallots(cs1, commitchema.PreballotKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), false, vss[1:]...)
			assurePreballotAlign(t, ballotChnl, altitude, iteration, rs.NominationLedger.Digest())

			var ext []byte
			if verifyInstance.ownsAddition {
				ext = []byte("REDACTED")
			}

			for _, vs := range vss[1:] {
				ballot, err := vs.attestBallot(commitchema.PreendorseKind, rs.NominationLedger.Digest(), rs.NominationLedgerFragments.Heading(), ext, verifyInstance.ownsAddition)
				require.NoError(t, err)
				appendBallots(cs1, ballot)
			}
			if verifyInstance.anticipateTriumphantIteration {
				assurePreendorse(ballotChnl, altitude, iteration)
				altitude++
				assureFreshIteration(freshIterationChnl, altitude, iteration)
			} else {
				assureNegativeFreshDeadline(deadlineChnl, cs1.settings.Preendorse(iteration).Nanoseconds())
			}

			m.AssertExpectations(t)
		})
	}
}

//
//
//
func VerifyStatusNotCollapseUponUnfitBallot(t *testing.T) {
	cs, vss := arbitraryStatus(2)
	altitude, iteration := cs.Altitude, cs.Iteration
	//
	node := nodestub.FreshNode(nil)

	initiateVerifyIteration(cs, altitude, iteration)

	_, itemLedger := resolveNomination(context.Background(), t, cs, vss[0], altitude, iteration)
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	assert.NoError(t, err)

	ballot := attestBallot(vss[1], commitchema.PreendorseKind, itemLedger.Digest(), itemLedgerFragments.Heading(), true)

	//
	ballot.AssessorOrdinal = int32(len(vss))

	ballotSignal := &BallotSignal{ballot}
	assert.NotPanics(t, func() {
		cs.processSignal(signalDetails{ballotSignal, node.ID()})
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
func VerifyPausingDeadlineUponVoidSpeck(*testing.T) {
	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())
	assureFreshIteration(freshIterationChnl, altitude, iteration+1)
}

//
//
//
func VerifyPausingDeadlineNominateUponFreshIteration(t *testing.T) {
	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assurePreballot(ballotChnl, altitude, iteration)

	advanceIteration(vss[1:]...)
	attestAppendBallots(cs1, commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, false, vs2, vs3, vs4)

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	rs := cs1.ObtainIterationStatus()
	assert.True(t, rs.Phase == controlkinds.IterationPhaseNominate) //

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Nominate(iteration).Nanoseconds())

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], nil)
}

//
//
//
func VerifyIterationOmitUponVoidSpeckOriginatingSuperiorIteration(t *testing.T) {
	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assurePreballot(ballotChnl, altitude, iteration)

	advanceIteration(vss[1:]...)
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2, vs3, vs4)

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assurePreendorse(ballotChnl, altitude, iteration)
	certifyPreendorse(t, cs1, iteration, -1, vss[0], nil, nil)

	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //
	assureFreshIteration(freshIterationChnl, altitude, iteration)
}

//
//
//
func VerifyPauseDeadlineNominateUponVoidSpeckForeachThatPrevailingIteration(t *testing.T) {
	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, int32(1)

	deadlineNominateChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	advanceIteration(vss[1:]...)
	attestAppendBallots(cs1, commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, false, vs2, vs3, vs4)

	assureFreshDeadline(deadlineNominateChnl, altitude, iteration, cs1.settings.Nominate(iteration).Nanoseconds())

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], nil)
}

//
//
func VerifyRelayFreshSoundLedgerIncidentUponEndorseLackingLedger(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, int32(1)

	advanceIteration(vs2, vs3, vs4)

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	soundLedgerChnl := listen(cs1.incidentChannel, kinds.IncidentInquireSoundLedger)

	_, itemLedger := resolveNomination(ctx, t, cs1, vs2, vs2.Altitude, vs2.Iteration)
	itemLedgerDigest := itemLedger.Digest()
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedgerDigest, itemLedgerFragments.Heading(), true, vs2, vs3, vs4)
	assureFreshSoundLedger(soundLedgerChnl, altitude, iteration)

	rs := cs1.ObtainIterationStatus()
	assert.True(t, rs.Phase == controlkinds.IterationPhaseEndorse)
	assert.True(t, rs.NominationLedger == nil)
	assert.True(t, rs.NominationLedgerFragments.Heading().Matches(itemLedgerFragments.Heading()))
}

//
//
//
func VerifyEndorseOriginatingPrecedingIteration(t *testing.T) {
	ctx := t.Context()

	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, int32(1)

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	soundLedgerChnl := listen(cs1.incidentChannel, kinds.IncidentInquireSoundLedger)
	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)

	item, itemLedger := resolveNomination(ctx, t, cs1, vs2, vs2.Altitude, vs2.Iteration)
	itemLedgerDigest := itemLedger.Digest()
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedgerDigest, itemLedgerFragments.Heading(), true, vs2, vs3, vs4)

	assureFreshSoundLedger(soundLedgerChnl, altitude, iteration)

	rs := cs1.ObtainIterationStatus()
	assert.True(t, rs.Phase == controlkinds.IterationPhaseEndorse)
	assert.True(t, rs.EndorseIteration == vs2.Iteration)
	assert.True(t, rs.NominationLedger == nil)
	assert.True(t, rs.NominationLedgerFragments.Heading().Matches(itemLedgerFragments.Heading()))

	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}

	assureFreshNomination(nominationChnl, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude+1, 0)
}

type mockedTransferObserver struct {
	ch chan struct{}
}

func (n *mockedTransferObserver) TransAccessible() <-chan struct{} {
	return n.ch
}

func (n *mockedTransferObserver) Alert() {
	n.ch <- struct{}{}
}

//
//
//
func VerifyInitiateFollowingAltitudeAccuratelySubsequentDeadline(t *testing.T) {
	settings.Agreement.OmitDeadlineEndorse = false
	cs1, vss := arbitraryStatus(4)
	cs1.transferObserver = &mockedTransferObserver{ch: make(chan struct{})}

	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	deadlineNominateChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlineNominate)
	preendorseDeadlineChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)

	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	freshLedgerHeading := listen(cs1.incidentChannel, kinds.IncidentInquireFreshLedgerHeading)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	thatLedgerDigest := rs.NominationLedger.Digest()
	thatLedgerFragments := rs.NominationLedgerFragments.Heading()

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], thatLedgerDigest)

	attestAppendBallots(cs1, commitchema.PreballotKind, thatLedgerDigest, thatLedgerFragments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], thatLedgerDigest, thatLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2)
	attestAppendBallots(cs1, commitchema.PreendorseKind, thatLedgerDigest, thatLedgerFragments, true, vs3)

	//
	assurePreendorseDeadline(preendorseDeadlineChnl)

	assureFreshIteration(freshIterationChnl, altitude, iteration+1)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, thatLedgerDigest, thatLedgerFragments, true, vs4)

	assureFreshLedgerHeading(freshLedgerHeading, altitude, thatLedgerDigest)

	cs1.transferObserver.(*mockedTransferObserver).Alert()

	assureFreshDeadline(deadlineNominateChnl, altitude+1, iteration, cs1.settings.Nominate(iteration).Nanoseconds())
	rs = cs1.ObtainIterationStatus()
	assert.False(
		t,
		rs.ActivatedDeadlinePreendorse,
		"REDACTED")
}

func VerifyRestoreDeadlinePreendorseOnFreshAltitude(t *testing.T) {
	ctx := t.Context()

	settings.Agreement.OmitDeadlineEndorse = false
	cs1, vss := arbitraryStatus(4)

	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration

	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)

	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	freshLedgerHeading := listen(cs1.incidentChannel, kinds.IncidentInquireFreshLedgerHeading)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	thatLedgerDigest := rs.NominationLedger.Digest()
	thatLedgerFragments := rs.NominationLedgerFragments.Heading()

	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], thatLedgerDigest)

	attestAppendBallots(cs1, commitchema.PreballotKind, thatLedgerDigest, thatLedgerFragments, false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], thatLedgerDigest, thatLedgerDigest)

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2)
	attestAppendBallots(cs1, commitchema.PreendorseKind, thatLedgerDigest, thatLedgerFragments, true, vs3)
	attestAppendBallots(cs1, commitchema.PreendorseKind, thatLedgerDigest, thatLedgerFragments, true, vs4)

	assureFreshLedgerHeading(freshLedgerHeading, altitude, thatLedgerDigest)

	item, itemLedger := resolveNomination(ctx, t, cs1, vs2, altitude+1, 0)
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	if err := cs1.AssignNominationAlsoLedger(item, itemLedger, itemLedgerFragments, "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assureFreshNomination(nominationChnl, altitude+1, 0)

	rs = cs1.ObtainIterationStatus()
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
func VerifyStatusStop1(t *testing.T) {
	cs1, vss := arbitraryStatus(4)
	vs2, vs3, vs4 := vss[1], vss[2], vss[3]
	altitude, iteration := cs1.Altitude, cs1.Iteration
	fragmentExtent := kinds.LedgerFragmentExtentOctets

	nominationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFinishNomination)
	deadlinePauseChnl := listen(cs1.incidentChannel, kinds.IncidentInquireDeadlinePause)
	freshIterationChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshIteration)
	freshLedgerChnl := listen(cs1.incidentChannel, kinds.IncidentInquireFreshLedger)
	pv1, err := cs1.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	location := pv1.Location()
	ballotChnl := listenTowardBalloter(cs1, location)

	//
	initiateVerifyIteration(cs1, altitude, iteration)
	assureFreshIteration(freshIterationChnl, altitude, iteration)

	assureFreshNomination(nominationChnl, altitude, iteration)
	rs := cs1.ObtainIterationStatus()
	itemLedger := rs.NominationLedger
	itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
	require.NoError(t, err)

	assurePreballot(ballotChnl, altitude, iteration)

	attestAppendBallots(cs1, commitchema.PreballotKind, itemLedger.Digest(), itemLedgerFragments.Heading(), false, vs2, vs3, vs4)

	assurePreendorse(ballotChnl, altitude, iteration)
	//
	certifyPreendorse(t, cs1, iteration, iteration, vss[0], itemLedger.Digest(), itemLedger.Digest())

	//
	attestAppendBallots(cs1, commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, true, vs2) //
	attestAppendBallots(cs1, commitchema.PreendorseKind, itemLedger.Digest(), itemLedgerFragments.Heading(), true, vs3)
	//
	preendorse4 := attestBallot(vs4, commitchema.PreendorseKind, itemLedger.Digest(), itemLedgerFragments.Heading(), true)

	advanceIteration(vs2, vs3, vs4)

	//
	assureFreshDeadline(deadlinePauseChnl, altitude, iteration, cs1.settings.Preendorse(iteration).Nanoseconds())

	iteration++ //

	assureFreshIteration(freshIterationChnl, altitude, iteration)
	rs = cs1.ObtainIterationStatus()

	t.Log("REDACTED")
	/*2
k
!
*/

	//
	assurePreballot(ballotChnl, altitude, iteration)
	certifyPreballot(t, cs1, iteration, vss[0], rs.SecuredLedger.Digest())

	//
	appendBallots(cs1, preendorse4)

	//
	assureFreshLedger(freshLedgerChnl, altitude)

	assureFreshIteration(freshIterationChnl, altitude+1, 0)
}

func VerifyStatusYieldsLedgerFragmentsMetrics(t *testing.T) {
	//
	cs, _ := arbitraryStatus(1)
	node := nodestub.FreshNode(nil)

	//
	fragments := kinds.FreshFragmentAssignOriginatingData(commitrand.Octets(100), 10)
	msg := &LedgerFragmentSignal{
		Altitude: 1,
		Iteration:  0,
		Fragment:   fragments.ObtainFragment(0),
	}

	cs.NominationLedgerFragments = kinds.FreshFragmentAssignOriginatingHeading(fragments.Heading())
	cs.processSignal(signalDetails{msg, node.ID()})

	metricsSignal := <-cs.metricsSignalStaging
	require.Equal(t, msg, metricsSignal.Msg, "REDACTED")
	require.Equal(t, node.ID(), metricsSignal.NodeUUID, "REDACTED")

	//
	cs.processSignal(signalDetails{msg, "REDACTED"})

	//
	msg.Iteration = 1
	cs.processSignal(signalDetails{msg, node.ID()})

	//
	msg.Altitude = 0
	cs.processSignal(signalDetails{msg, node.ID()})

	//
	msg.Altitude = 3
	cs.processSignal(signalDetails{msg, node.ID()})

	select {
	case <-cs.metricsSignalStaging:
		t.Errorf("REDACTED")
	case <-time.After(50 * time.Millisecond):
	}
}

func VerifyStatusEmissionBallotMetrics(t *testing.T) {
	cs, vss := arbitraryStatus(2)
	//
	node := nodestub.FreshNode(nil)

	arbitraryOctets := commitrand.Octets(tenderminthash.Extent)

	ballot := attestBallot(vss[1], commitchema.PreendorseKind, arbitraryOctets, kinds.FragmentAssignHeading{}, true)

	ballotSignal := &BallotSignal{ballot}
	cs.processSignal(signalDetails{ballotSignal, node.ID()})

	metricsSignal := <-cs.metricsSignalStaging
	require.Equal(t, ballotSignal, metricsSignal.Msg, "REDACTED")
	require.Equal(t, node.ID(), metricsSignal.NodeUUID, "REDACTED")

	//
	cs.processSignal(signalDetails{&BallotSignal{ballot}, "REDACTED"})

	//
	advanceAltitude(vss[1])
	ballot = attestBallot(vss[1], commitchema.PreendorseKind, arbitraryOctets, kinds.FragmentAssignHeading{}, true)

	cs.processSignal(signalDetails{&BallotSignal{ballot}, node.ID()})

	select {
	case <-cs.metricsSignalStaging:
		t.Errorf("REDACTED")
	case <-time.After(50 * time.Millisecond):
	}
}

func VerifyAttestIdenticalBallotBis(t *testing.T) {
	_, vss := arbitraryStatus(2)

	arbitraryOctets := commitrand.Octets(tenderminthash.Extent)

	ballot := attestBallot(vss[1],
		commitchema.PreendorseKind,
		arbitraryOctets,
		kinds.FragmentAssignHeading{Sum: 10, Digest: arbitraryOctets},
		true,
	)

	ballot2 := attestBallot(vss[1],
		commitchema.PreendorseKind,
		arbitraryOctets,
		kinds.FragmentAssignHeading{Sum: 10, Digest: arbitraryOctets},
		true,
	)

	require.Equal(t, ballot, ballot2)
}

//
func listen(incidentChannel *kinds.IncidentChannel, q tendermintpubsub.Inquire) <-chan tendermintpubsub.Signal {
	sub, err := incidentChannel.Listen(context.Background(), verifyListener, q)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", verifyListener, q))
	}
	return sub.Out()
}

//
func listenNegCached(incidentChannel *kinds.IncidentChannel, q tendermintpubsub.Inquire) <-chan tendermintpubsub.Signal {
	sub, err := incidentChannel.ListenUncached(context.Background(), verifyListener, q)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", verifyListener, q))
	}
	return sub.Out()
}

func attestAppendPreendorseUsingAddition(
	t *testing.T,
	cs *Status,
	digest []byte,
	heading kinds.FragmentAssignHeading,
	addition []byte,
	mock *assessorMock,
) {
	v, err := mock.attestBallot(commitchema.PreendorseKind, digest, heading, addition, true)
	require.NoError(t, err, "REDACTED")
	appendBallots(cs, v)
}

func locateLedgerExtentThreshold(t *testing.T, altitude, maximumOctets int64, cs *Status, fragmentExtent uint32, bulky bool) (*kinds.Ledger, *kinds.FragmentAssign) {
	var displacement int64
	if !bulky {
		displacement = -2
	}
	gentleMaximumDataOctets := int(kinds.MaximumDataOctets(maximumOctets, 0, 0))
	for i := gentleMaximumDataOctets; i < gentleMaximumDataOctets*2; i++ {
		itemLedger, err := cs.status.CreateLedger(
			altitude,
			[]kinds.Tx{[]byte("REDACTED" + strings.Repeat("REDACTED", i-2))},
			&kinds.Endorse{},
			nil,
			cs.privateAssessorPublicToken.Location(),
		)
		require.NoError(t, err)

		itemLedgerFragments, err := itemLedger.CreateFragmentAssign(fragmentExtent)
		require.NoError(t, err)
		if itemLedgerFragments.OctetExtent() > maximumOctets+displacement {
			s := "REDACTED"
			if bulky {
				s = "REDACTED"
			}
			t.Log("REDACTED"+s+"REDACTED", "REDACTED", i, "REDACTED", gentleMaximumDataOctets)
			return itemLedger, itemLedgerFragments
		}
	}
	require.Fail(t, "REDACTED")
	return nil, nil
}
