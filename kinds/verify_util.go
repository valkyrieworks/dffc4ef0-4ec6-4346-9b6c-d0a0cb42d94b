package kinds

import (
	"fmt"
	"testing"
	"time"

	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/release"
	"github.com/stretchr/testify/require"
)

func CreateExtensionEndorse(ledgerUID LedgerUID, level int64, epoch int32,
	ballotCollection *BallotCollection, ratifiers []PrivateRatifier, now time.Time, extensionActivated bool,
) (*ExpandedEndorse, error) {
	//
	for i := 0; i < len(ratifiers); i++ {
		publicKey, err := ratifiers[i].FetchPublicKey()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		ballot := &Ballot{
			RatifierLocation: publicKey.Location(),
			RatifierOrdinal:   int32(i),
			Level:           level,
			Cycle:            epoch,
			Kind:             engineproto.PreendorseKind,
			LedgerUID:          ledgerUID,
			Timestamp:        now,
		}

		_, err = attestAppendBallot(ratifiers[i], ballot, ballotCollection)
		if err != nil {
			return nil, err
		}
	}

	var activateLevel int64
	if extensionActivated {
		activateLevel = level
	}

	return ballotCollection.CreateExpandedEndorse(IfaceOptions{BallotPluginsActivateLevel: activateLevel}), nil
}

func attestAppendBallot(privateValue PrivateRatifier, ballot *Ballot, ballotCollection *BallotCollection) (bool, error) {
	if ballot.Kind != ballotCollection.attestedMessageKind {
		return false, fmt.Errorf("REDACTED", ballot.Kind, ballotCollection.attestedMessageKind)
	}
	if _, err := AttestAndInspectBallot(ballot, privateValue, ballotCollection.LedgerUID(), ballotCollection.pluginsActivated); err != nil {
		return false, err
	}
	return ballotCollection.AppendBallot(ballot)
}

func CreateBallot(
	val PrivateRatifier,
	ledgerUID string,
	valueOrdinal int32,
	level int64,
	epoch int32,
	phase engineproto.AttestedMessageKind,
	ledgerUID LedgerUID,
	moment time.Time,
) (*Ballot, error) {
	publicKey, err := val.FetchPublicKey()
	if err != nil {
		return nil, err
	}

	ballot := &Ballot{
		RatifierLocation: publicKey.Location(),
		RatifierOrdinal:   valueOrdinal,
		Level:           level,
		Cycle:            epoch,
		Kind:             phase,
		LedgerUID:          ledgerUID,
		Timestamp:        moment,
	}

	pluginsActivated := phase == engineproto.PreendorseKind
	if _, err := AttestAndInspectBallot(ballot, val, ledgerUID, pluginsActivated); err != nil {
		return nil, err
	}

	return ballot, nil
}

func CreateBallotNoFault(
	t *testing.T,
	val PrivateRatifier,
	ledgerUID string,
	valueOrdinal int32,
	level int64,
	epoch int32,
	phase engineproto.AttestedMessageKind,
	ledgerUID LedgerUID,
	moment time.Time,
) *Ballot {
	ballot, err := CreateBallot(val, ledgerUID, valueOrdinal, level, epoch, phase, ledgerUID, moment)
	require.NoError(t, err)
	return ballot
}

//
//
//
func CreateLedger(level int64, txs []Tx, finalEndorse *Endorse, proof []Proof) *Ledger {
	ledger := &Ledger{
		Heading: Heading{
			Release: cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 0},
			Level:  level,
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
