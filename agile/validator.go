package agile

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
var FallbackRelianceStratum = strongarithmetic.Portion{Dividend: 1, Divisor: 3}

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
func ValidateUnContiguous(
	reliableHeading *kinds.NotatedHeading, //
	reliableValues *kinds.AssessorAssign, //
	unreliableHeadline *kinds.NotatedHeading, //
	unreliableValues *kinds.AssessorAssign, //
	relyingCycle time.Duration,
	now time.Time,
	maximumTimerDeviation time.Duration,
	relianceStratum strongarithmetic.Portion,
) error {
	if unreliableHeadline.Altitude == reliableHeading.Altitude+1 {
		return errors.New("REDACTED")
	}

	if HeadlineLapsed(reliableHeading, relyingCycle, now) {
		return FaultAgedHeadlineLapsed{reliableHeading.Moment.Add(relyingCycle), now}
	}

	if err := validateFreshHeadlineAlsoValues(
		unreliableHeadline, unreliableValues,
		reliableHeading,
		now, maximumTimerDeviation); err != nil {
		return FaultUnfitHeadline{err}
	}

	attestedSigningStash := kinds.FreshSigningStash()
	//
	err := reliableValues.ValidateEndorseAgileRelyingUsingStash(reliableHeading.SuccessionUUID, unreliableHeadline.Endorse, relianceStratum, attestedSigningStash)
	if err != nil {
		switch e := err.(type) {
		case kinds.FaultNegationAmpleBallotingPotencyNotated:
			return FaultFreshItemAssignCannotExistReliable{e}
		default:
			return e
		}
	}

	//
	//
	//
	//
	//
	if err := unreliableValues.ValidateEndorseAgileUsingStash(reliableHeading.SuccessionUUID, unreliableHeadline.Endorse.LedgerUUID,
		unreliableHeadline.Altitude, unreliableHeadline.Endorse, attestedSigningStash); err != nil {
		return FaultUnfitHeadline{err}
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
func ValidateContiguous(
	reliableHeading *kinds.NotatedHeading, //
	unreliableHeadline *kinds.NotatedHeading, //
	unreliableValues *kinds.AssessorAssign, //
	relyingCycle time.Duration,
	now time.Time,
	maximumTimerDeviation time.Duration,
) error {
	if unreliableHeadline.Altitude != reliableHeading.Altitude+1 {
		return errors.New("REDACTED")
	}

	if HeadlineLapsed(reliableHeading, relyingCycle, now) {
		return FaultAgedHeadlineLapsed{reliableHeading.Moment.Add(relyingCycle), now}
	}

	if err := validateFreshHeadlineAlsoValues(
		unreliableHeadline, unreliableValues,
		reliableHeading,
		now, maximumTimerDeviation); err != nil {
		return FaultUnfitHeadline{err}
	}

	//
	if !bytes.Equal(unreliableHeadline.AssessorsDigest, reliableHeading.FollowingAssessorsDigest) {
		err := fmt.Errorf("REDACTED",
			reliableHeading.FollowingAssessorsDigest,
			unreliableHeadline.AssessorsDigest,
		)
		return err
	}

	//
	if err := unreliableValues.ValidateEndorseAgile(reliableHeading.SuccessionUUID, unreliableHeadline.Endorse.LedgerUUID,
		unreliableHeadline.Altitude, unreliableHeadline.Endorse); err != nil {
		return FaultUnfitHeadline{err}
	}

	return nil
}

//
func Validate(
	reliableHeading *kinds.NotatedHeading, //
	reliableValues *kinds.AssessorAssign, //
	unreliableHeadline *kinds.NotatedHeading, //
	unreliableValues *kinds.AssessorAssign, //
	relyingCycle time.Duration,
	now time.Time,
	maximumTimerDeviation time.Duration,
	relianceStratum strongarithmetic.Portion,
) error {
	if unreliableHeadline.Altitude != reliableHeading.Altitude+1 {
		return ValidateUnContiguous(reliableHeading, reliableValues, unreliableHeadline, unreliableValues,
			relyingCycle, now, maximumTimerDeviation, relianceStratum)
	}

	return ValidateContiguous(reliableHeading, unreliableHeadline, unreliableValues, relyingCycle, now, maximumTimerDeviation)
}

func validateFreshHeadlineAlsoValues(
	unreliableHeadline *kinds.NotatedHeading,
	unreliableValues *kinds.AssessorAssign,
	reliableHeading *kinds.NotatedHeading,
	now time.Time,
	maximumTimerDeviation time.Duration,
) error {
	if err := unreliableHeadline.CertifyFundamental(reliableHeading.SuccessionUUID); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if unreliableHeadline.Altitude <= reliableHeading.Altitude {
		return fmt.Errorf("REDACTED",
			unreliableHeadline.Altitude,
			reliableHeading.Altitude)
	}

	if !unreliableHeadline.Moment.After(reliableHeading.Moment) {
		return fmt.Errorf("REDACTED",
			unreliableHeadline.Moment,
			reliableHeading.Moment)
	}

	if !unreliableHeadline.Moment.Before(now.Add(maximumTimerDeviation)) {
		return fmt.Errorf("REDACTED",
			unreliableHeadline.Moment,
			now,
			maximumTimerDeviation)
	}

	if !bytes.Equal(unreliableHeadline.AssessorsDigest, unreliableValues.Digest()) {
		return fmt.Errorf("REDACTED",
			unreliableHeadline.AssessorsDigest,
			unreliableValues.Digest(),
			unreliableHeadline.Altitude,
		)
	}

	return nil
}

//
//
//
func CertifyRelianceStratum(lvl strongarithmetic.Portion) error {
	if lvl.Dividend*3 < lvl.Divisor || //
		lvl.Dividend > lvl.Divisor || //
		lvl.Divisor == 0 {
		return fmt.Errorf("REDACTED", lvl)
	}
	return nil
}

//
func HeadlineLapsed(h *kinds.NotatedHeading, relyingCycle time.Duration, now time.Time) bool {
	maturityMoment := h.Moment.Add(relyingCycle)
	return !maturityMoment.After(now)
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
func ValidateReverse(unreliableHeadline, reliableHeading *kinds.Heading) error {
	if err := unreliableHeadline.CertifyFundamental(); err != nil {
		return FaultUnfitHeadline{err}
	}

	if unreliableHeadline.SuccessionUUID != reliableHeading.SuccessionUUID {
		return FaultUnfitHeadline{errors.New("REDACTED")}
	}

	if !unreliableHeadline.Moment.Before(reliableHeading.Moment) {
		return FaultUnfitHeadline{
			fmt.Errorf("REDACTED",
				unreliableHeadline.Moment,
				reliableHeading.Moment),
		}
	}

	if !bytes.Equal(unreliableHeadline.Digest(), reliableHeading.FinalLedgerUUID.Digest) {
		return FaultUnfitHeadline{
			fmt.Errorf("REDACTED",
				unreliableHeadline.Digest(),
				reliableHeading.FinalLedgerUUID.Digest),
		}
	}

	return nil
}
