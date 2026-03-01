package verify

import (
	"fmt"
	"time"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func CreateEndorseOriginatingBallotAssign(ledgerUUID kinds.LedgerUUID, ballotAssign *kinds.BallotAssign, assessors []kinds.PrivateAssessor, now time.Time) (*kinds.Endorse, error) {
	//
	for i := 0; i < len(assessors); i++ {
		publicToken, err := assessors[i].ObtainPublicToken()
		if err != nil {
			return nil, err
		}
		ballot := &kinds.Ballot{
			AssessorLocation: publicToken.Location(),
			AssessorOrdinal:   int32(i),
			Altitude:           ballotAssign.ObtainAltitude(),
			Iteration:            ballotAssign.ObtainIteration(),
			Kind:             commitchema.PreendorseKind,
			LedgerUUID:          ledgerUUID,
			Timestamp:        now,
		}

		v := ballot.TowardSchema()

		if err := assessors[i].AttestBallot(ballotAssign.SuccessionUUID(), v); err != nil {
			return nil, err
		}
		ballot.Notation = v.Notation
		if _, err := ballotAssign.AppendBallot(ballot); err != nil {
			return nil, err
		}
	}

	return ballotAssign.CreateExpandedEndorse(kinds.IfaceParameters{BallotAdditionsActivateAltitude: 0}).TowardEndorse(), nil
}

func CreateEndorse(ledgerUUID kinds.LedgerUUID, altitude int64, iteration int32, itemAssign *kinds.AssessorAssign, privateItems []kinds.PrivateAssessor, successionUUID string, now time.Time) (*kinds.Endorse, error) {
	signatures := make([]kinds.EndorseSignature, len(itemAssign.Assessors))
	for i := 0; i < len(itemAssign.Assessors); i++ {
		signatures[i] = kinds.FreshEndorseSignatureMissing()
	}

	for _, privateItem := range privateItems {
		pk, err := privateItem.ObtainPublicToken()
		if err != nil {
			return nil, err
		}
		location := pk.Location()

		idx, _ := itemAssign.ObtainViaLocationAlterable(location)
		if idx < 0 {
			return nil, fmt.Errorf("REDACTED", location)
		}

		ballot := &kinds.Ballot{
			AssessorLocation: location,
			AssessorOrdinal:   idx,
			Altitude:           altitude,
			Iteration:            iteration,
			Kind:             commitchema.PreendorseKind,
			LedgerUUID:          ledgerUUID,
			Timestamp:        now,
		}

		v := ballot.TowardSchema()

		if err := privateItem.AttestBallot(successionUUID, v); err != nil {
			return nil, err
		}

		signatures[idx] = kinds.EndorseSignature{
			LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
			AssessorLocation: location,
			Timestamp:        now,
			Notation:        v.Notation,
		}
	}

	return &kinds.Endorse{Altitude: altitude, Iteration: iteration, LedgerUUID: ledgerUUID, Notations: signatures}, nil
}
