package verify

import (
	"fmt"
	"time"

	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

func CreateEndorseFromBallotCollection(ledgerUID kinds.LedgerUID, ballotCollection *kinds.BallotCollection, ratifiers []kinds.PrivateRatifier, now time.Time) (*kinds.Endorse, error) {
	//
	for i := 0; i < len(ratifiers); i++ {
		publicKey, err := ratifiers[i].FetchPublicKey()
		if err != nil {
			return nil, err
		}
		ballot := &kinds.Ballot{
			RatifierLocation: publicKey.Location(),
			RatifierOrdinal:   int32(i),
			Level:           ballotCollection.FetchLevel(),
			Cycle:            ballotCollection.FetchDuration(),
			Kind:             engineproto.PreendorseKind,
			LedgerUID:          ledgerUID,
			Timestamp:        now,
		}

		v := ballot.ToSchema()

		if err := ratifiers[i].AttestBallot(ballotCollection.LedgerUID(), v); err != nil {
			return nil, err
		}
		ballot.Autograph = v.Autograph
		if _, err := ballotCollection.AppendBallot(ballot); err != nil {
			return nil, err
		}
	}

	return ballotCollection.CreateExpandedEndorse(kinds.IfaceOptions{BallotPluginsActivateLevel: 0}).ToEndorse(), nil
}

func CreateEndorse(ledgerUID kinds.LedgerUID, level int64, epoch int32, valueCollection *kinds.RatifierAssign, privateValues []kinds.PrivateRatifier, ledgerUID string, now time.Time) (*kinds.Endorse, error) {
	autographs := make([]kinds.EndorseSignature, len(valueCollection.Ratifiers))
	for i := 0; i < len(valueCollection.Ratifiers); i++ {
		autographs[i] = kinds.NewEndorseSignatureMissing()
	}

	for _, privateValue := range privateValues {
		pk, err := privateValue.FetchPublicKey()
		if err != nil {
			return nil, err
		}
		address := pk.Location()

		idx, _ := valueCollection.FetchByLocationMut(address)
		if idx < 0 {
			return nil, fmt.Errorf("REDACTED", address)
		}

		ballot := &kinds.Ballot{
			RatifierLocation: address,
			RatifierOrdinal:   idx,
			Level:           level,
			Cycle:            epoch,
			Kind:             engineproto.PreendorseKind,
			LedgerUID:          ledgerUID,
			Timestamp:        now,
		}

		v := ballot.ToSchema()

		if err := privateValue.AttestBallot(ledgerUID, v); err != nil {
			return nil, err
		}

		autographs[idx] = kinds.EndorseSignature{
			LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
			RatifierLocation: address,
			Timestamp:        now,
			Autograph:        v.Autograph,
		}
	}

	return &kinds.Endorse{Level: level, Cycle: epoch, LedgerUID: ledgerUID, Endorsements: autographs}, nil
}
