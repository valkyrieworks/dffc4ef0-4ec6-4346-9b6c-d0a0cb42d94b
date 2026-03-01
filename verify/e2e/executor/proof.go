package primary

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//
const agileCustomerProofProportion = 4

//
//
//
//
func IntroduceProof(ctx context.Context, r *rand.Rand, simnet *e2e.Simnet, quantity int) error {
	//
	var objectivePeer *e2e.Peer

	for _, idx := range r.Perm(len(simnet.Peers)) {
		objectivePeer = simnet.Peers[idx]

		if objectivePeer.Style == e2e.StyleGerm || objectivePeer.Style == e2e.StyleAgile {
			objectivePeer = nil
			continue
		}

		break
	}

	if objectivePeer == nil {
		return errors.New("REDACTED")
	}

	tracer.Details(fmt.Sprintf("REDACTED", objectivePeer.Alias, quantity))

	customer, err := objectivePeer.Customer()
	if err != nil {
		return err
	}

	//
	ledgerResult, err := customer.Ledger(ctx, nil)
	if err != nil {
		return err
	}
	proofAltitude := ledgerResult.Ledger.Altitude
	pauseAltitude := ledgerResult.Ledger.Altitude + 3

	nthAssessors := 100
	itemResult, err := customer.Assessors(ctx, &proofAltitude, nil, &nthAssessors)
	if err != nil {
		return err
	}

	itemAssign, err := kinds.AssessorAssignOriginatingCurrentAssessors(itemResult.Assessors)
	if err != nil {
		return err
	}

	//
	privateItems, err := obtainSecludedAssessorTokens(simnet)
	if err != nil {
		return err
	}

	//
	//
	_, err = pauseForeachPeer(ctx, objectivePeer, pauseAltitude, time.Minute)
	if err != nil {
		return err
	}

	var ev kinds.Proof
	for i := 0; i < quantity; i++ {
		soundOccurence := true
		if i%agileCustomerProofProportion == 0 {
			soundOccurence = i%(agileCustomerProofProportion*2) != 0 //
			ev, err = composeAgileCustomerOnslaughtProof(
				ctx, privateItems, proofAltitude, itemAssign, simnet.Alias, ledgerResult.Ledger.Moment, soundOccurence,
			)
		} else {
			var dve *kinds.ReplicatedBallotProof
			dve, err = composeReplicatedBallotProof(
				privateItems, proofAltitude, itemAssign, simnet.Alias, ledgerResult.Ledger.Moment,
			)
			if dve.BallotAN.Altitude < simnet.BallotAdditionsActivateAltitude {
				dve.BallotAN.Addition = nil
				dve.BallotAN.AdditionNotation = nil
				dve.BallotBYTE.Addition = nil
				dve.BallotBYTE.AdditionNotation = nil
			}
			ev = dve
		}
		if err != nil {
			return err
		}

		_, err := customer.MulticastProof(ctx, ev)
		if !soundOccurence {
			//
			//
			quantity++
		}
		if soundOccurence != (err == nil) {
			if err == nil {
				return errors.New("REDACTED")
			}
			return err
		}
		time.Sleep(5 * time.Second / time.Duration(quantity))
	}

	//
	//
	_, err = pauseForeachPeer(ctx, objectivePeer, ledgerResult.Ledger.Altitude+2, 30*time.Second)
	if err != nil {
		return err
	}

	tracer.Details(fmt.Sprintf("REDACTED", ledgerResult.Ledger.Altitude+2))

	return nil
}

func obtainSecludedAssessorTokens(simnet *e2e.Simnet) ([]kinds.SimulatePRV, error) {
	privateItems := []kinds.SimulatePRV{}

	for _, peer := range simnet.Peers {
		if peer.Style == e2e.StyleAssessor {
			privateTokenRoute := filepath.Join(simnet.Dir, peer.Alias, PrivatevalueTokenRecord)
			privateToken, err := retrievePrivateToken(privateTokenRoute)
			if err != nil {
				return nil, err
			}
			//
			//
			privateItems = append(privateItems, kinds.FreshSimulatePRVUsingParameters(privateToken, false, false))
		}
	}

	return privateItems, nil
}

//
//
func composeAgileCustomerOnslaughtProof(
	ctx context.Context,
	privateItems []kinds.SimulatePRV,
	altitude int64,
	values *kinds.AssessorAssign,
	successionUUID string,
	occurenceMoment time.Time,
	soundProof bool,
) (*kinds.AgileCustomerOnslaughtProof, error) {
	//
	fabricatedAltitude := altitude + 2
	fabricatedMoment := occurenceMoment.Add(1 * time.Second)
	heading := createHeadingUnpredictable(successionUUID, fabricatedAltitude)
	heading.Moment = fabricatedMoment

	//
	//
	pv, discordantValues, err := transformAssessorGroup(ctx, privateItems, values, !soundProof)
	if err != nil {
		return nil, err
	}

	heading.AssessorsDigest = discordantValues.Digest()

	//
	ledgerUUID := createLedgerUUID(heading.Digest(), 1000, []byte("REDACTED"))
	ballotAssign := kinds.FreshBallotAssign(successionUUID, fabricatedAltitude, 0, commitchema.AttestedSignalKind(2), discordantValues)
	endorse, err := verify.CreateEndorseOriginatingBallotAssign(ledgerUUID, ballotAssign, pv, fabricatedMoment)
	if err != nil {
		return nil, err
	}

	//
	if !soundProof {
		endorse.Notations[len(endorse.Notations)-1].Notation[0]++
	}

	ev := &kinds.AgileCustomerOnslaughtProof{
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: &kinds.NotatedHeading{
				Heading: heading,
				Endorse: endorse,
			},
			AssessorAssign: discordantValues,
		},
		SharedAltitude:     altitude,
		SumBallotingPotency: values.SumBallotingPotency(),
		Timestamp:        occurenceMoment,
	}
	ev.TreacherousAssessors = ev.ObtainTreacherousAssessors(values, &kinds.NotatedHeading{
		Heading: createHeadingUnpredictable(successionUUID, fabricatedAltitude),
	})
	return ev, nil
}

//
//
func composeReplicatedBallotProof(
	privateItems []kinds.SimulatePRV,
	altitude int64,
	values *kinds.AssessorAssign,
	successionUUID string,
	moment time.Time,
) (*kinds.ReplicatedBallotProof, error) {
	privateItem, itemOffset, err := obtainUnpredictableAssessorPosition(privateItems, values)
	if err != nil {
		return nil, err
	}
	ballotAN, err := kinds.CreateBallot(privateItem, successionUUID, itemOffset, altitude, 0, 2, createUnpredictableLedgerUUID(), moment)
	if err != nil {
		return nil, err
	}
	ballotBYTE, err := kinds.CreateBallot(privateItem, successionUUID, itemOffset, altitude, 0, 2, createUnpredictableLedgerUUID(), moment)
	if err != nil {
		return nil, err
	}
	ev, err := kinds.FreshReplicatedBallotProof(ballotAN, ballotBYTE, moment, values)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return ev, nil
}

//
//
func obtainUnpredictableAssessorPosition(privateItems []kinds.SimulatePRV, values *kinds.AssessorAssign) (kinds.SimulatePRV, int32, error) {
	for _, idx := range rand.Perm(len(privateItems)) {
		pv := privateItems[idx]
		itemOffset, _ := values.ObtainViaLocation(pv.PrivateToken.PublicToken().Location())
		if itemOffset >= 0 {
			return pv, itemOffset, nil
		}
	}
	return kinds.SimulatePRV{}, -1, errors.New("REDACTED")
}

func retrievePrivateToken(tokenRecordRoute string) (security.PrivateToken, error) {
	tokenJSNOctets, err := os.ReadFile(tokenRecordRoute)
	if err != nil {
		return nil, err
	}
	prvToken := privatevalue.RecordPRVToken{}
	err = strongmindjson.Decode(tokenJSNOctets, &prvToken)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", tokenRecordRoute, err)
	}

	return prvToken.PrivateToken, nil
}

func createHeadingUnpredictable(successionUUID string, altitude int64) *kinds.Heading {
	return &kinds.Heading{
		Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1},
		SuccessionUUID:            successionUUID,
		Altitude:             altitude,
		Moment:               time.Now(),
		FinalLedgerUUID:        createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED")),
		FinalEndorseDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
		DataDigest:           security.CHARArbitraryOctets(tenderminthash.Extent),
		AssessorsDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
		FollowingAssessorsDigest: security.CHARArbitraryOctets(tenderminthash.Extent),
		AgreementDigest:      security.CHARArbitraryOctets(tenderminthash.Extent),
		PlatformDigest:            security.CHARArbitraryOctets(tenderminthash.Extent),
		FinalOutcomesDigest:    security.CHARArbitraryOctets(tenderminthash.Extent),
		ProofDigest:       security.CHARArbitraryOctets(tenderminthash.Extent),
		NominatorLocation:    security.CHARArbitraryOctets(security.LocatorExtent),
	}
}

func createUnpredictableLedgerUUID() kinds.LedgerUUID {
	return createLedgerUUID(security.CHARArbitraryOctets(tenderminthash.Extent), 100, security.CHARArbitraryOctets(tenderminthash.Extent))
}

func createLedgerUUID(digest []byte, fragmentAssignExtent uint32, fragmentAssignDigest []byte) kinds.LedgerUUID {
	var (
		h   = make([]byte, tenderminthash.Extent)
		psH = make([]byte, tenderminthash.Extent)
	)
	copy(h, digest)
	copy(psH, fragmentAssignDigest)
	return kinds.LedgerUUID{
		Digest: h,
		FragmentAssignHeading: kinds.FragmentAssignHeading{
			Sum: fragmentAssignExtent,
			Digest:  psH,
		},
	}
}

func transformAssessorGroup(
	ctx context.Context,
	privateItems []kinds.SimulatePRV,
	values *kinds.AssessorAssign,
	nop bool,
) ([]kinds.PrivateAssessor, *kinds.AssessorAssign, error) {
	freshItem, freshPrivateItem, err := verify.Assessor(ctx, 10)
	if err != nil {
		return nil, nil, err
	}

	var freshValues *kinds.AssessorAssign
	if nop {
		freshValues = kinds.FreshAssessorAssign(values.Duplicate().Assessors)
	} else {
		if values.Extent() > 2 {
			freshValues = kinds.FreshAssessorAssign(append(values.Duplicate().Assessors[:values.Extent()-1], freshItem))
		} else {
			freshValues = kinds.FreshAssessorAssign(append(values.Duplicate().Assessors, freshItem))
		}
	}

	//
	pv := make([]kinds.PrivateAssessor, freshValues.Extent())
	for idx, val := range freshValues.Assessors {
		detected := false
		for _, p := range append(privateItems, freshPrivateItem.(kinds.SimulatePRV)) {
			if bytes.Equal(p.PrivateToken.PublicToken().Location(), val.Location) {
				pv[idx] = p
				detected = true
				break
			}
		}
		if !detected {
			return nil, nil, fmt.Errorf("REDACTED", val.Location)
		}
	}

	return pv, freshValues, nil
}
