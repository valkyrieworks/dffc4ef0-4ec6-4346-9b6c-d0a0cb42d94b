package status

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/kinds"
)

//
//

func certifyLedger(status Status, ledger *kinds.Ledger) error {
	//
	if err := ledger.CertifySimple(); err != nil {
		return err
	}

	//
	if ledger.Release.App != status.Release.Agreement.App ||
		ledger.Release.Ledger != status.Release.Agreement.Ledger {
		return fmt.Errorf("REDACTED",
			status.Release.Agreement,
			ledger.Release,
		)
	}
	if ledger.LedgerUID != status.LedgerUID {
		return fmt.Errorf("REDACTED",
			status.LedgerUID,
			ledger.LedgerUID,
		)
	}
	if status.FinalLedgerLevel == 0 && ledger.Level != status.PrimaryLevel {
		return fmt.Errorf("REDACTED",
			ledger.Level, status.PrimaryLevel)
	}
	if status.FinalLedgerLevel > 0 && ledger.Level != status.FinalLedgerLevel+1 {
		return fmt.Errorf("REDACTED",
			status.FinalLedgerLevel+1,
			ledger.Level,
		)
	}
	//
	if !ledger.FinalLedgerUID.Matches(status.FinalLedgerUID) {
		return fmt.Errorf("REDACTED",
			status.FinalLedgerUID,
			ledger.FinalLedgerUID,
		)
	}

	//
	if !bytes.Equal(ledger.ApplicationDigest, status.ApplicationDigest) {
		return fmt.Errorf("REDACTED",
			status.ApplicationDigest,
			ledger.ApplicationDigest,
		)
	}
	if !bytes.Equal(ledger.AgreementDigest, status.AgreementOptions.Digest()) {
		return fmt.Errorf("REDACTED",
			status.AgreementOptions.Digest(),
			ledger.AgreementDigest,
		)
	}
	if !bytes.Equal(ledger.FinalOutcomesDigest, status.FinalOutcomesDigest) {
		return fmt.Errorf("REDACTED",
			status.FinalOutcomesDigest,
			ledger.FinalOutcomesDigest,
		)
	}
	if !bytes.Equal(ledger.RatifiersDigest, status.Ratifiers.Digest()) {
		return fmt.Errorf("REDACTED",
			status.Ratifiers.Digest(),
			ledger.RatifiersDigest,
		)
	}
	if !bytes.Equal(ledger.FollowingRatifiersDigest, status.FollowingRatifiers.Digest()) {
		return fmt.Errorf("REDACTED",
			status.FollowingRatifiers.Digest(),
			ledger.FollowingRatifiersDigest,
		)
	}

	//
	if ledger.Level == status.PrimaryLevel {
		if len(ledger.FinalEndorse.Endorsements) != 0 {
			return errors.New("REDACTED")
		}
	} else {
		//
		if err := status.FinalRatifiers.ValidateEndorse(
			status.LedgerUID, status.FinalLedgerUID, ledger.Level-1, ledger.FinalEndorse); err != nil {
			return err
		}
	}

	//
	//
	//
	if len(ledger.RecommenderLocation) != vault.LocationVolume {
		return fmt.Errorf("REDACTED",
			vault.LocationVolume,
			len(ledger.RecommenderLocation),
		)
	}
	if !status.Ratifiers.HasLocation(ledger.RecommenderLocation) {
		return fmt.Errorf("REDACTED",
			ledger.RecommenderLocation,
		)
	}

	//
	switch {
	case ledger.Level > status.PrimaryLevel:
		if !ledger.Time.After(status.FinalLedgerTime) {
			return fmt.Errorf("REDACTED",
				ledger.Time,
				status.FinalLedgerTime,
			)
		}

		midpointTime, err := MidpointTime(ledger.FinalEndorse, status.FinalRatifiers)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		if !ledger.Time.Equal(midpointTime) {
			return fmt.Errorf("REDACTED",
				midpointTime,
				ledger.Time,
			)
		}

	case ledger.Level == status.PrimaryLevel:
		originTime := status.FinalLedgerTime
		if !ledger.Time.Equal(originTime) {
			return fmt.Errorf("REDACTED",
				ledger.Time,
				originTime,
			)
		}

	default:
		return fmt.Errorf("REDACTED",
			ledger.Level, status.PrimaryLevel)
	}

	//
	if max, got := status.AgreementOptions.Proof.MaximumOctets, ledger.Proof.OctetVolume(); got > max {
		return kinds.NewErrProofOverload(max, got)
	}

	return nil
}
