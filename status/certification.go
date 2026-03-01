package status

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//

func certifyLedger(status Status, ledger *kinds.Ledger) error {
	//
	if err := ledger.CertifyFundamental(); err != nil {
		return err
	}

	//
	if ledger.Edition.App != status.Edition.Agreement.App ||
		ledger.Edition.Ledger != status.Edition.Agreement.Ledger {
		return fmt.Errorf("REDACTED",
			status.Edition.Agreement,
			ledger.Edition,
		)
	}
	if ledger.SuccessionUUID != status.SuccessionUUID {
		return fmt.Errorf("REDACTED",
			status.SuccessionUUID,
			ledger.SuccessionUUID,
		)
	}
	if status.FinalLedgerAltitude == 0 && ledger.Altitude != status.PrimaryAltitude {
		return fmt.Errorf("REDACTED",
			ledger.Altitude, status.PrimaryAltitude)
	}
	if status.FinalLedgerAltitude > 0 && ledger.Altitude != status.FinalLedgerAltitude+1 {
		return fmt.Errorf("REDACTED",
			status.FinalLedgerAltitude+1,
			ledger.Altitude,
		)
	}
	//
	if !ledger.FinalLedgerUUID.Matches(status.FinalLedgerUUID) {
		return fmt.Errorf("REDACTED",
			status.FinalLedgerUUID,
			ledger.FinalLedgerUUID,
		)
	}

	//
	if !bytes.Equal(ledger.PlatformDigest, status.PlatformDigest) {
		return fmt.Errorf("REDACTED",
			status.PlatformDigest,
			ledger.PlatformDigest,
		)
	}
	if !bytes.Equal(ledger.AgreementDigest, status.AgreementSettings.Digest()) {
		return fmt.Errorf("REDACTED",
			status.AgreementSettings.Digest(),
			ledger.AgreementDigest,
		)
	}
	if !bytes.Equal(ledger.FinalOutcomesDigest, status.FinalOutcomesDigest) {
		return fmt.Errorf("REDACTED",
			status.FinalOutcomesDigest,
			ledger.FinalOutcomesDigest,
		)
	}
	if !bytes.Equal(ledger.AssessorsDigest, status.Assessors.Digest()) {
		return fmt.Errorf("REDACTED",
			status.Assessors.Digest(),
			ledger.AssessorsDigest,
		)
	}
	if !bytes.Equal(ledger.FollowingAssessorsDigest, status.FollowingAssessors.Digest()) {
		return fmt.Errorf("REDACTED",
			status.FollowingAssessors.Digest(),
			ledger.FollowingAssessorsDigest,
		)
	}

	//
	if ledger.Altitude == status.PrimaryAltitude {
		if len(ledger.FinalEndorse.Notations) != 0 {
			return errors.New("REDACTED")
		}
	} else {
		//
		if err := status.FinalAssessors.ValidateEndorse(
			status.SuccessionUUID, status.FinalLedgerUUID, ledger.Altitude-1, ledger.FinalEndorse); err != nil {
			return err
		}
	}

	//
	//
	//
	if len(ledger.NominatorLocation) != security.LocatorExtent {
		return fmt.Errorf("REDACTED",
			security.LocatorExtent,
			len(ledger.NominatorLocation),
		)
	}
	if !status.Assessors.OwnsLocation(ledger.NominatorLocation) {
		return fmt.Errorf("REDACTED",
			ledger.NominatorLocation,
		)
	}

	//
	switch {
	case ledger.Altitude > status.PrimaryAltitude:
		if !ledger.Moment.After(status.FinalLedgerMoment) {
			return fmt.Errorf("REDACTED",
				ledger.Moment,
				status.FinalLedgerMoment,
			)
		}

		averageMoment, err := AverageMoment(ledger.FinalEndorse, status.FinalAssessors)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		if !ledger.Moment.Equal(averageMoment) {
			return fmt.Errorf("REDACTED",
				averageMoment,
				ledger.Moment,
			)
		}

	case ledger.Altitude == status.PrimaryAltitude:
		inaugurationMoment := status.FinalLedgerMoment
		if !ledger.Moment.Equal(inaugurationMoment) {
			return fmt.Errorf("REDACTED",
				ledger.Moment,
				inaugurationMoment,
			)
		}

	default:
		return fmt.Errorf("REDACTED",
			ledger.Altitude, status.PrimaryAltitude)
	}

	//
	if max, got := status.AgreementSettings.Proof.MaximumOctets, ledger.Proof.OctetExtent(); got > max {
		return kinds.FreshFaultProofOverrun(max, got)
	}

	return nil
}
