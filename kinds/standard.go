package kinds

import (
	"time"

	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

//

//
const TimeLayout = time.RFC3339Nano

//
//

func StandardizeLedgerUID(bid engineproto.LedgerUID) *engineproto.StandardLedgerUID {
	rbid, err := LedgerUIDFromSchema(&bid)
	if err != nil {
		panic(err)
	}
	var cbid *engineproto.StandardLedgerUID
	if rbid == nil || rbid.IsNil() {
		cbid = nil
	} else {
		cbid = &engineproto.StandardLedgerUID{
			Digest:          bid.Digest,
			SegmentAssignHeading: StandardizeSectionCollectionHeading(bid.SegmentAssignHeading),
		}
	}

	return cbid
}

//
func StandardizeSectionCollectionHeading(psh engineproto.SegmentAssignHeading) engineproto.StandardSectionCollectionHeading {
	return engineproto.StandardSectionCollectionHeading(psh)
}

//
func StandardizeNomination(ledgerUID string, nomination *engineproto.Nomination) engineproto.StandardNomination {
	return engineproto.StandardNomination{
		Kind:      engineproto.NominationKind,
		Level:    nomination.Level,       //
		Cycle:     int64(nomination.Cycle), //
		POLDuration:  int64(nomination.PolEpoch),
		LedgerUID:   StandardizeLedgerUID(nomination.LedgerUID),
		Timestamp: nomination.Timestamp,
		LedgerUID:   ledgerUID,
	}
}

//
//
//
func StandardizeBallot(ledgerUID string, ballot *engineproto.Ballot) engineproto.StandardBallot {
	return engineproto.StandardBallot{
		Kind:      ballot.Kind,
		Level:    ballot.Level,       //
		Cycle:     int64(ballot.Cycle), //
		LedgerUID:   StandardizeLedgerUID(ballot.LedgerUID),
		Timestamp: ballot.Timestamp,
		LedgerUID:   ledgerUID,
	}
}

//
//
//
func StandardizeBallotAddition(ledgerUID string, ballot *engineproto.Ballot) engineproto.StandardBallotAddition {
	return engineproto.StandardBallotAddition{
		Addition: ballot.Addition,
		Level:    ballot.Level,
		Cycle:     int64(ballot.Cycle),
		SeriesUid:   ledgerUID,
	}
}

//
func StandardTime(t time.Time) string {
	//
	//
	//
	return engineclock.Standard(t).Format(TimeLayout)
}
