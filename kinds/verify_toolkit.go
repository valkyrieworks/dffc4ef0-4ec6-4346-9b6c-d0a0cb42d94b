package kinds

import (
	"fmt"
	"testing"
	"time"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
	"github.com/stretchr/testify/require"
)

func CreateAddnEndorse(ledgerUUID LedgerUUID, altitude int64, iteration int32,
	ballotAssign *BallotAssign, assessors []PrivateAssessor, now time.Time, addnActivated bool,
) (*ExpandedEndorse, error) {
	//
	for i := 0; i < len(assessors); i++ {
		publicToken, err := assessors[i].ObtainPublicToken()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		ballot := &Ballot{
			AssessorLocation: publicToken.Location(),
			AssessorOrdinal:   int32(i),
			Altitude:           altitude,
			Iteration:            iteration,
			Kind:             commitchema.PreendorseKind,
			LedgerUUID:          ledgerUUID,
			Timestamp:        now,
		}

		_, err = attestAppendBallot(assessors[i], ballot, ballotAssign)
		if err != nil {
			return nil, err
		}
	}

	var activateAltitude int64
	if addnActivated {
		activateAltitude = altitude
	}

	return ballotAssign.CreateExpandedEndorse(IfaceParameters{BallotAdditionsActivateAltitude: activateAltitude}), nil
}

func attestAppendBallot(privateItem PrivateAssessor, ballot *Ballot, ballotAssign *BallotAssign) (bool, error) {
	if ballot.Kind != ballotAssign.notatedSignalKind {
		return false, fmt.Errorf("REDACTED", ballot.Kind, ballotAssign.notatedSignalKind)
	}
	if _, err := AttestAlsoInspectBallot(ballot, privateItem, ballotAssign.SuccessionUUID(), ballotAssign.additionsActivated); err != nil {
		return false, err
	}
	return ballotAssign.AppendBallot(ballot)
}

func CreateBallot(
	val PrivateAssessor,
	successionUUID string,
	itemOrdinal int32,
	altitude int64,
	iteration int32,
	phase commitchema.AttestedSignalKind,
	ledgerUUID LedgerUUID,
	moment time.Time,
) (*Ballot, error) {
	publicToken, err := val.ObtainPublicToken()
	if err != nil {
		return nil, err
	}

	ballot := &Ballot{
		AssessorLocation: publicToken.Location(),
		AssessorOrdinal:   itemOrdinal,
		Altitude:           altitude,
		Iteration:            iteration,
		Kind:             phase,
		LedgerUUID:          ledgerUUID,
		Timestamp:        moment,
	}

	additionsActivated := phase == commitchema.PreendorseKind
	if _, err := AttestAlsoInspectBallot(ballot, val, successionUUID, additionsActivated); err != nil {
		return nil, err
	}

	return ballot, nil
}

func CreateBallotNegativeFailure(
	t *testing.T,
	val PrivateAssessor,
	successionUUID string,
	itemOrdinal int32,
	altitude int64,
	iteration int32,
	phase commitchema.AttestedSignalKind,
	ledgerUUID LedgerUUID,
	moment time.Time,
) *Ballot {
	ballot, err := CreateBallot(val, successionUUID, itemOrdinal, altitude, iteration, phase, ledgerUUID, moment)
	require.NoError(t, err)
	return ballot
}

//
//
//
func CreateLedger(altitude int64, txs []Tx, finalEndorse *Endorse, proof []Proof) *Ledger {
	ledger := &Ledger{
		Heading: Heading{
			Edition: strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 0},
			Altitude:  altitude,
		},
		Data: Data{
			Txs: txs,
		},
		Proof:   ProofData{Proof: proof},
		FinalEndorse: finalEndorse,
	}
	ledger.populateHeading()
	return ledger
}
