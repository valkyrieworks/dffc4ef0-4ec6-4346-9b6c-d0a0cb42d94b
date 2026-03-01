package kinds

import (
	"time"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

//

//
const MomentLayout = time.RFC3339Nano

//
//

func NormalizeLedgerUUID(bid commitchema.LedgerUUID) *commitchema.StandardLedgerUUID {
	recid, err := LedgerUUIDOriginatingSchema(&bid)
	if err != nil {
		panic(err)
	}
	var callid *commitchema.StandardLedgerUUID
	if recid == nil || recid.EqualsNull() {
		callid = nil
	} else {
		callid = &commitchema.StandardLedgerUUID{
			Digest:          bid.Digest,
			FragmentAssignHeading: NormalizeFragmentAssignHeading(bid.FragmentAssignHeading),
		}
	}

	return callid
}

//
func NormalizeFragmentAssignHeading(psh commitchema.FragmentAssignHeading) commitchema.StandardFragmentAssignHeading {
	return commitchema.StandardFragmentAssignHeading(psh)
}

//
func NormalizeNomination(successionUUID string, nomination *commitchema.Nomination) commitchema.StandardNomination {
	return commitchema.StandardNomination{
		Kind:      commitchema.NominationKind,
		Altitude:    nomination.Altitude,       //
		Iteration:     int64(nomination.Iteration), //
		PolicyIteration:  int64(nomination.PolicyIteration),
		LedgerUUID:   NormalizeLedgerUUID(nomination.LedgerUUID),
		Timestamp: nomination.Timestamp,
		SuccessionUUID:   successionUUID,
	}
}

//
//
//
func NormalizeBallot(successionUUID string, ballot *commitchema.Ballot) commitchema.StandardBallot {
	return commitchema.StandardBallot{
		Kind:      ballot.Kind,
		Altitude:    ballot.Altitude,       //
		Iteration:     int64(ballot.Iteration), //
		LedgerUUID:   NormalizeLedgerUUID(ballot.LedgerUUID),
		Timestamp: ballot.Timestamp,
		SuccessionUUID:   successionUUID,
	}
}

//
//
//
func NormalizeBallotAddition(successionUUID string, ballot *commitchema.Ballot) commitchema.StandardBallotAddition {
	return commitchema.StandardBallotAddition{
		Addition: ballot.Addition,
		Altitude:    ballot.Altitude,
		Iteration:     int64(ballot.Iteration),
		SuccessionUuid:   successionUUID,
	}
}

//
func StandardMoment(t time.Time) string {
	//
	//
	//
	return committime.Standard(t).Format(MomentLayout)
}
