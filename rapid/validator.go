package rapid

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	cometmath "github.com/valkyrieworks/utils/math"
	"github.com/valkyrieworks/kinds"
)

//
//
var StandardRelianceLayer = cometmath.Portion{Dividend: 1, Divisor: 3}

//
//
//
//
//
//
//
//
//
//
//
//
//
func ValidateNotNeighboring(
	validatedHeading *kinds.AttestedHeading, //
	validatedValues *kinds.RatifierAssign, //
	unvalidatedHeading *kinds.AttestedHeading, //
	unvalidatedValues *kinds.RatifierAssign, //
	validatingDuration time.Duration,
	now time.Time,
	maximumTimerDeviation time.Duration,
	validateLayer cometmath.Portion,
) error {
	if unvalidatedHeading.Level == validatedHeading.Level+1 {
		return errors.New("REDACTED")
	}

	if HeadingLapsed(validatedHeading, validatingDuration, now) {
		return ErrAgedHeadingLapsed{validatedHeading.Time.Add(validatingDuration), now}
	}

	if err := validateNewHeadingAndValues(
		unvalidatedHeading, unvalidatedValues,
		validatedHeading,
		now, maximumTimerDeviation); err != nil {
		return ErrCorruptHeading{err}
	}

	certifiedAutographRepository := kinds.NewAutographRepository()
	//
	err := validatedValues.ValidateEndorseRapidValidatingWithRepository(validatedHeading.LedgerUID, unvalidatedHeading.Endorse, validateLayer, certifiedAutographRepository)
	if err != nil {
		switch e := err.(type) {
		case kinds.ErrNoSufficientPollingEnergyAttested:
			return ErrNewValueCollectionCannotBeValidated{e}
		default:
			return e
		}
	}

	//
	//
	//
	//
	//
	if err := unvalidatedValues.ValidateEndorseRapidWithRepository(validatedHeading.LedgerUID, unvalidatedHeading.Endorse.LedgerUID,
		unvalidatedHeading.Level, unvalidatedHeading.Endorse, certifiedAutographRepository); err != nil {
		return ErrCorruptHeading{err}
	}

	return nil
}

//
//
//
//
//
//
//
//
//
//
//
//
func ValidateNeighboring(
	validatedHeading *kinds.AttestedHeading, //
	unvalidatedHeading *kinds.AttestedHeading, //
	unvalidatedValues *kinds.RatifierAssign, //
	validatingDuration time.Duration,
	now time.Time,
	maximumTimerDeviation time.Duration,
) error {
	if unvalidatedHeading.Level != validatedHeading.Level+1 {
		return errors.New("REDACTED")
	}

	if HeadingLapsed(validatedHeading, validatingDuration, now) {
		return ErrAgedHeadingLapsed{validatedHeading.Time.Add(validatingDuration), now}
	}

	if err := validateNewHeadingAndValues(
		unvalidatedHeading, unvalidatedValues,
		validatedHeading,
		now, maximumTimerDeviation); err != nil {
		return ErrCorruptHeading{err}
	}

	//
	if !bytes.Equal(unvalidatedHeading.RatifiersDigest, validatedHeading.FollowingRatifiersDigest) {
		err := fmt.Errorf("REDACTED",
			validatedHeading.FollowingRatifiersDigest,
			unvalidatedHeading.RatifiersDigest,
		)
		return err
	}

	//
	if err := unvalidatedValues.ValidateEndorseRapid(validatedHeading.LedgerUID, unvalidatedHeading.Endorse.LedgerUID,
		unvalidatedHeading.Level, unvalidatedHeading.Endorse); err != nil {
		return ErrCorruptHeading{err}
	}

	return nil
}

//
func Validate(
	validatedHeading *kinds.AttestedHeading, //
	validatedValues *kinds.RatifierAssign, //
	unvalidatedHeading *kinds.AttestedHeading, //
	unvalidatedValues *kinds.RatifierAssign, //
	validatingDuration time.Duration,
	now time.Time,
	maximumTimerDeviation time.Duration,
	validateLayer cometmath.Portion,
) error {
	if unvalidatedHeading.Level != validatedHeading.Level+1 {
		return ValidateNotNeighboring(validatedHeading, validatedValues, unvalidatedHeading, unvalidatedValues,
			validatingDuration, now, maximumTimerDeviation, validateLayer)
	}

	return ValidateNeighboring(validatedHeading, unvalidatedHeading, unvalidatedValues, validatingDuration, now, maximumTimerDeviation)
}

func validateNewHeadingAndValues(
	unvalidatedHeading *kinds.AttestedHeading,
	unvalidatedValues *kinds.RatifierAssign,
	validatedHeading *kinds.AttestedHeading,
	now time.Time,
	maximumTimerDeviation time.Duration,
) error {
	if err := unvalidatedHeading.CertifySimple(validatedHeading.LedgerUID); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if unvalidatedHeading.Level <= validatedHeading.Level {
		return fmt.Errorf("REDACTED",
			unvalidatedHeading.Level,
			validatedHeading.Level)
	}

	if !unvalidatedHeading.Time.After(validatedHeading.Time) {
		return fmt.Errorf("REDACTED",
			unvalidatedHeading.Time,
			validatedHeading.Time)
	}

	if !unvalidatedHeading.Time.Before(now.Add(maximumTimerDeviation)) {
		return fmt.Errorf("REDACTED",
			unvalidatedHeading.Time,
			now,
			maximumTimerDeviation)
	}

	if !bytes.Equal(unvalidatedHeading.RatifiersDigest, unvalidatedValues.Digest()) {
		return fmt.Errorf("REDACTED",
			unvalidatedHeading.RatifiersDigest,
			unvalidatedValues.Digest(),
			unvalidatedHeading.Level,
		)
	}

	return nil
}

//
//
//
func CertifyRelianceLayer(lvl cometmath.Portion) error {
	if lvl.Dividend*3 < lvl.Divisor || //
		lvl.Dividend > lvl.Divisor || //
		lvl.Divisor == 0 {
		return fmt.Errorf("REDACTED", lvl)
	}
	return nil
}

//
func HeadingLapsed(h *kinds.AttestedHeading, validatingDuration time.Duration, now time.Time) bool {
	expiryTime := h.Time.Add(validatingDuration)
	return !expiryTime.After(now)
}

//
//
//
//
//
//
//
//
//
func ValidateReverse(unvalidatedHeading, validatedHeading *kinds.Heading) error {
	if err := unvalidatedHeading.CertifySimple(); err != nil {
		return ErrCorruptHeading{err}
	}

	if unvalidatedHeading.LedgerUID != validatedHeading.LedgerUID {
		return ErrCorruptHeading{errors.New("REDACTED")}
	}

	if !unvalidatedHeading.Time.Before(validatedHeading.Time) {
		return ErrCorruptHeading{
			fmt.Errorf("REDACTED",
				unvalidatedHeading.Time,
				validatedHeading.Time),
		}
	}

	if !bytes.Equal(unvalidatedHeading.Digest(), validatedHeading.FinalLedgerUID.Digest) {
		return ErrCorruptHeading{
			fmt.Errorf("REDACTED",
				unvalidatedHeading.Digest(),
				validatedHeading.FinalLedgerUID.Digest),
		}
	}

	return nil
}
