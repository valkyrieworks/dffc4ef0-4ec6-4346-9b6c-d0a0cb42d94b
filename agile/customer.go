package agile

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

type style byte

const (
	ordered style = iota + 1
	omitting

	fallbackThinningExtent      = 1000
	fallbackMaximumReissueEndeavors = 10
	//
	//
	//
	validateOmittingDividend   = 9
	validateOmittingDivisor = 16

	//
	//
	//
	//
	fallbackMaximumTimerDeviation = 10 * time.Second

	//
	fallbackMaximumLedgerDelay = 10 * time.Second
)

//
type Selection func(*Customer)

//
//
//
func OrderedValidation() Selection {
	return func(c *Customer) {
		c.validationStyle = ordered
	}
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
func OmittingValidation(relianceStratum strongarithmetic.Portion) Selection {
	return func(c *Customer) {
		c.validationStyle = omitting
		c.relianceStratum = relianceStratum
	}
}

//
//
//
//
func ThinningExtent(h uint16) Selection {
	return func(c *Customer) {
		c.thinningExtent = h
	}
}

//
//
//
func RatificationProcedure(fn func(deed string) bool) Selection {
	return func(c *Customer) {
		c.ratificationProc = fn
	}
}

//
func Tracer(l log.Tracer) Selection {
	return func(c *Customer) {
		c.tracer = l
	}
}

//
//
func MaximumReissueEndeavors(max uint16) Selection {
	return func(c *Customer) {
		c.maximumReissueEndeavors = max
	}
}

//
//
func MaximumTimerDeviation(d time.Duration) Selection {
	return func(c *Customer) {
		c.maximumTimerDeviation = d
	}
}

//
//
//
//
//
//
//
//
func MaximumLedgerDelay(d time.Duration) Selection {
	return func(c *Customer) {
		c.maximumLedgerDelay = d
	}
}

//
//
//
//
//
type Customer struct {
	successionUUID          string
	relyingCycle   time.Duration //
	validationStyle style
	relianceStratum       strongarithmetic.Portion
	maximumReissueEndeavors uint16 //
	maximumTimerDeviation    time.Duration
	maximumLedgerDelay      time.Duration

	//
	supplierExclusion commitchronize.Exclusion
	//
	leading supplier.Supplier
	//
	attestors []supplier.Supplier

	//
	reliableDepot depot.Depot
	//
	newestReliableLedger *kinds.AgileLedger

	//
	thinningExtent uint16
	//
	ratificationProc func(deed string) bool

	exit chan struct{}

	tracer log.Tracer
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
func FreshCustomer(
	ctx context.Context,
	successionUUID string,
	relianceChoices RelianceChoices,
	leading supplier.Supplier,
	attestors []supplier.Supplier,
	reliableDepot depot.Depot,
	choices ...Selection,
) (*Customer, error) {
	if err := relianceChoices.CertifyFundamental(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	c, err := FreshCustomerOriginatingReliableDepot(successionUUID, relianceChoices.Cycle, leading, attestors, reliableDepot, choices...)
	if err != nil {
		return nil, err
	}

	if c.newestReliableLedger != nil {
		c.tracer.Details("REDACTED")
		if err := c.inspectReliableHeadlineApplyingChoices(ctx, relianceChoices); err != nil {
			return nil, err
		}
	}

	if c.newestReliableLedger == nil || c.newestReliableLedger.Altitude < relianceChoices.Altitude {
		c.tracer.Details("REDACTED")
		if err := c.bootstrapUsingRelianceChoices(ctx, relianceChoices); err != nil {
			return nil, err
		}
	}

	return c, err
}

//
//
//
func FreshCustomerOriginatingReliableDepot(
	successionUUID string,
	relyingCycle time.Duration,
	leading supplier.Supplier,
	attestors []supplier.Supplier,
	reliableDepot depot.Depot,
	choices ...Selection,
) (*Customer, error) {
	c := &Customer{
		successionUUID:          successionUUID,
		relyingCycle:   relyingCycle,
		validationStyle: omitting,
		relianceStratum:       FallbackRelianceStratum,
		maximumReissueEndeavors: fallbackMaximumReissueEndeavors,
		maximumTimerDeviation:    fallbackMaximumTimerDeviation,
		maximumLedgerDelay:      fallbackMaximumLedgerDelay,
		leading:          leading,
		attestors:        attestors,
		reliableDepot:     reliableDepot,
		thinningExtent:      fallbackThinningExtent,
		ratificationProc:   func(deed string) bool { return true },
		exit:             make(chan struct{}),
		tracer:           log.FreshNooperationTracer(),
	}

	for _, o := range choices {
		o(c)
	}

	//
	if len(c.attestors) == 0 {
		return nil, FaultNegativeAttestors
	}

	//
	for i, w := range attestors {
		if w.SuccessionUUID() != successionUUID {
			return nil, fmt.Errorf("REDACTED",
				i, w, w.SuccessionUUID(), successionUUID)
		}
	}

	//
	if err := CertifyRelianceStratum(c.relianceStratum); err != nil {
		return nil, err
	}

	if err := c.recoverReliableAgileLedger(); err != nil {
		return nil, err
	}

	return c, nil
}

//
func (c *Customer) recoverReliableAgileLedger() error {
	finalAltitude, err := c.reliableDepot.FinalAgileLedgerAltitude()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if finalAltitude > 0 {
		reliableLedger, err := c.reliableDepot.AgileLedger(finalAltitude)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		c.newestReliableLedger = reliableLedger
		c.tracer.Details("REDACTED", "REDACTED", finalAltitude)
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
//
//
//
//
//
//
func (c *Customer) inspectReliableHeadlineApplyingChoices(ctx context.Context, choices RelianceChoices) error {
	var leadingDigest []byte
	switch {
	case choices.Altitude > c.newestReliableLedger.Altitude:
		h, err := c.agileLedgerOriginatingLeading(ctx, c.newestReliableLedger.Altitude)
		if err != nil {
			return err
		}
		leadingDigest = h.Digest()
	case choices.Altitude == c.newestReliableLedger.Altitude:
		leadingDigest = choices.Digest
	case choices.Altitude < c.newestReliableLedger.Altitude:
		c.tracer.Details("REDACTED",
			"REDACTED", choices.Altitude,
			"REDACTED", c.newestReliableLedger.Altitude,
			"REDACTED", c.newestReliableLedger.Digest())

		deed := fmt.Sprintf(
			"REDACTED",
			choices.Altitude, choices.Digest,
			c.newestReliableLedger.Altitude, c.newestReliableLedger.Digest())
		if c.ratificationProc(deed) {
			//
			err := c.sanitizeSubsequent(choices.Altitude)
			if err != nil {
				return fmt.Errorf("REDACTED", choices.Altitude, err)
			}

			c.tracer.Details("REDACTED",
				"REDACTED", choices.Altitude)
		} else {
			return nil
		}

		leadingDigest = choices.Digest
	}

	if !bytes.Equal(leadingDigest, c.newestReliableLedger.Digest()) {
		c.tracer.Details("REDACTED",
			"REDACTED", c.newestReliableLedger.Digest(), "REDACTED", leadingDigest)

		deed := fmt.Sprintf(
			"REDACTED",
			c.newestReliableLedger.Digest(), leadingDigest)
		if c.ratificationProc(deed) {
			err := c.Sanitize()
			if err != nil {
				return fmt.Errorf("REDACTED", err)
			}
		} else {
			return errors.New("REDACTED")
		}
	}

	return nil
}

//
//
func (c *Customer) bootstrapUsingRelianceChoices(ctx context.Context, choices RelianceChoices) error {
	//
	l, err := c.agileLedgerOriginatingLeading(ctx, choices.Altitude)
	if err != nil {
		return err
	}

	//
	//
	//
	if err := l.CertifyFundamental(c.successionUUID); err != nil {
		return err
	}

	if !bytes.Equal(l.Digest(), choices.Digest) {
		return fmt.Errorf("REDACTED", choices.Digest, l.Digest())
	}

	//
	err = l.AssessorAssign.ValidateEndorseAgile(c.successionUUID, l.Endorse.LedgerUUID, l.Altitude, l.Endorse)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if err := c.contrastInitialAgileLedgerUsingAttestors(ctx, l); err != nil {
		return err
	}

	//
	return c.reviseReliableAgileLedger(l)
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
func (c *Customer) ReliableAgileLedger(altitude int64) (*kinds.AgileLedger, error) {
	altitude, err := c.contrastUsingNewestAltitude(altitude)
	if err != nil {
		return nil, err
	}
	return c.reliableDepot.AgileLedger(altitude)
}

func (c *Customer) contrastUsingNewestAltitude(altitude int64) (int64, error) {
	newestAltitude, err := c.FinalReliableAltitude()
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}
	if newestAltitude == -1 {
		return 0, errors.New("REDACTED")
	}

	switch {
	case altitude > newestAltitude:
		return 0, fmt.Errorf("REDACTED", newestAltitude)
	case altitude == 0:
		return newestAltitude, nil
	case altitude < 0:
		return 0, errors.New("REDACTED")
	}

	return altitude, nil
}

//
//
//
func (c *Customer) Revise(ctx context.Context, now time.Time) (*kinds.AgileLedger, error) {
	finalReliableAltitude, err := c.FinalReliableAltitude()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	if finalReliableAltitude == -1 {
		//
		return nil, nil
	}

	newestLedger, err := c.agileLedgerOriginatingLeading(ctx, 0)
	if err != nil {
		return nil, err
	}

	if newestLedger.Altitude > finalReliableAltitude {
		err = c.validateAgileLedger(ctx, newestLedger, now)
		if err != nil {
			return nil, err
		}
		c.tracer.Details("REDACTED", "REDACTED", newestLedger.Altitude, "REDACTED", newestLedger.Digest())
		return newestLedger, nil
	}

	return nil, nil
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
func (c *Customer) ValidateAgileLedgerLocatedAltitude(ctx context.Context, altitude int64, now time.Time) (*kinds.AgileLedger, error) {
	if altitude <= 0 {
		return nil, errors.New("REDACTED")
	}

	//
	h, err := c.ReliableAgileLedger(altitude)
	if err == nil {
		c.tracer.Details("REDACTED", "REDACTED", altitude, "REDACTED", h.Digest())
		//
		return h, nil
	}

	//
	l, err := c.agileLedgerOriginatingLeading(ctx, altitude)
	if err != nil {
		return nil, err
	}

	return l, c.validateAgileLedger(ctx, l, now)
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
//
//
//
//
func (c *Customer) ValidateHeadline(ctx context.Context, freshHeadline *kinds.Heading, now time.Time) error {
	if freshHeadline == nil {
		return errors.New("REDACTED")
	}
	if freshHeadline.Altitude <= 0 {
		return errors.New("REDACTED")
	}

	//
	l, err := c.ReliableAgileLedger(freshHeadline.Altitude)
	if err == nil {
		//
		if !bytes.Equal(l.Digest(), freshHeadline.Digest()) {
			return fmt.Errorf("REDACTED", l.Digest(), freshHeadline.Digest())
		}
		c.tracer.Details("REDACTED",
			"REDACTED", freshHeadline.Altitude, "REDACTED", freshHeadline.Digest())
		return nil
	}

	//
	l, err = c.agileLedgerOriginatingLeading(ctx, freshHeadline.Altitude)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if !bytes.Equal(l.Digest(), freshHeadline.Digest()) {
		return fmt.Errorf("REDACTED", l.Digest(), freshHeadline.Digest())
	}

	return c.validateAgileLedger(ctx, l, now)
}

func (c *Customer) validateAgileLedger(ctx context.Context, freshAgileLedger *kinds.AgileLedger, now time.Time) error {
	c.tracer.Details("REDACTED", "REDACTED", freshAgileLedger.Altitude, "REDACTED", freshAgileLedger.Digest())

	var (
		validateMethod func(ctx context.Context, reliable *kinds.AgileLedger, new *kinds.AgileLedger, now time.Time) error
		err        error
	)

	switch c.validationStyle {
	case ordered:
		validateMethod = c.validateSuccessive
	case omitting:
		validateMethod = c.validateOmittingVersusLeading
	default:
		panic(fmt.Sprintf("REDACTED", c.validationStyle))
	}

	initialLedgerAltitude, err := c.InitialReliableAltitude()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	switch {
	//
	case freshAgileLedger.Altitude >= c.newestReliableLedger.Altitude:
		err = validateMethod(ctx, c.newestReliableLedger, freshAgileLedger, now)

	//
	case freshAgileLedger.Altitude < initialLedgerAltitude:
		var initialLedger *kinds.AgileLedger
		initialLedger, err = c.reliableDepot.AgileLedger(initialLedgerAltitude)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		err = c.reverse(ctx, initialLedger.Heading, freshAgileLedger.Heading)

	//
	default:
		var nearestLedger *kinds.AgileLedger
		nearestLedger, err = c.reliableDepot.AgileLedgerPrior(freshAgileLedger.Altitude)
		if err != nil {
			return fmt.Errorf("REDACTED", freshAgileLedger.Altitude, err)
		}
		err = validateMethod(ctx, nearestLedger, freshAgileLedger, now)
	}
	if err != nil {
		c.tracer.Failure("REDACTED", "REDACTED", err)
		return err
	}

	//
	return c.reviseReliableAgileLedger(freshAgileLedger)
}

//
func (c *Customer) validateSuccessive(
	ctx context.Context,
	reliableLedger *kinds.AgileLedger,
	freshAgileLedger *kinds.AgileLedger,
	now time.Time,
) error {
	var (
		attestedLedger = reliableLedger
		provisionalLedger  *kinds.AgileLedger
		err           error
		logging         = []*kinds.AgileLedger{reliableLedger}
	)

	for altitude := reliableLedger.Altitude + 1; altitude <= freshAgileLedger.Altitude; altitude++ {
		//
		if altitude == freshAgileLedger.Altitude { //
			provisionalLedger = freshAgileLedger
		} else { //
			provisionalLedger, err = c.agileLedgerOriginatingLeading(ctx, altitude)
			if err != nil {
				return FaultValidationUnsuccessful{Originating: attestedLedger.Altitude, To: altitude, Rationale: err}
			}
		}

		//
		c.tracer.Diagnose("REDACTED",
			"REDACTED", attestedLedger.Altitude,
			"REDACTED", attestedLedger.Digest(),
			"REDACTED", provisionalLedger.Altitude,
			"REDACTED", provisionalLedger.Digest())

		err = ValidateContiguous(attestedLedger.NotatedHeading, provisionalLedger.NotatedHeading, provisionalLedger.AssessorAssign,
			c.relyingCycle, now, c.maximumTimerDeviation)
		if err != nil {
			err := FaultValidationUnsuccessful{Originating: attestedLedger.Altitude, To: provisionalLedger.Altitude, Rationale: err}

			switch errors.Unwrap(err).(type) {
			case FaultUnfitHeadline:
				//
				if err.To == freshAgileLedger.Altitude {
					c.tracer.Diagnose("REDACTED", "REDACTED", err)
					return err
				}

				//
				//
				c.tracer.Failure("REDACTED", "REDACTED", err)

				substitutionLedger, discardFault := c.locateFreshLeading(ctx, freshAgileLedger.Altitude, true)
				if discardFault != nil {
					c.tracer.Diagnose("REDACTED", "REDACTED", discardFault)
					return err
				}

				if !bytes.Equal(substitutionLedger.Digest(), freshAgileLedger.Digest()) {
					c.tracer.Failure("REDACTED",
						"REDACTED", freshAgileLedger.Digest(),
						"REDACTED", substitutionLedger.Digest())
					//
					return err
				}

				//
				altitude--

				continue
			default:
				return err
			}
		}

		//
		attestedLedger = provisionalLedger

		//
		logging = append(logging, attestedLedger)
	}

	//
	//
	//
	//
	//
	return c.senseDeviation(ctx, logging, now)
}

//
//
//
//
//
//
//
func (c *Customer) validateOmitting(
	ctx context.Context,
	origin supplier.Supplier,
	reliableLedger *kinds.AgileLedger,
	freshAgileLedger *kinds.AgileLedger,
	now time.Time,
) ([]*kinds.AgileLedger, error) {
	var (
		ledgerStash = []*kinds.AgileLedger{freshAgileLedger}
		intensity      = 0

		attestedLedger = reliableLedger
		logging         = []*kinds.AgileLedger{reliableLedger}
	)

	for {
		c.tracer.Diagnose("REDACTED",
			"REDACTED", attestedLedger.Altitude,
			"REDACTED", attestedLedger.Digest(),
			"REDACTED", ledgerStash[intensity].Altitude,
			"REDACTED", ledgerStash[intensity].Digest())

		err := Validate(attestedLedger.NotatedHeading, attestedLedger.AssessorAssign, ledgerStash[intensity].NotatedHeading,
			ledgerStash[intensity].AssessorAssign, c.relyingCycle, now, c.maximumTimerDeviation, c.relianceStratum)
		switch err.(type) {
		case nil:
			//
			if intensity == 0 {
				logging = append(logging, freshAgileLedger)
				return logging, nil
			}
			//
			attestedLedger = ledgerStash[intensity]
			//
			ledgerStash = ledgerStash[:intensity]
			//
			intensity = 0
			//
			logging = append(logging, attestedLedger)

		case FaultFreshItemAssignCannotExistReliable:
			//
			if intensity == len(ledgerStash)-1 {
				fulcrumAltitude := attestedLedger.Altitude + (ledgerStash[intensity].Altitude-attestedLedger.
					Altitude)*validateOmittingDividend/validateOmittingDivisor
				provisionalLedger, supplierFault := origin.AgileLedger(ctx, fulcrumAltitude)
				switch supplierFault {
				case nil:
					ledgerStash = append(ledgerStash, provisionalLedger)

				//
				case supplier.FaultAgileLedgerNegationDetected, supplier.FaultNegativeReply, supplier.FaultAltitudeExcessivelyTall:
					return nil, err

				//
				//
				default:
					return nil, FaultValidationUnsuccessful{Originating: attestedLedger.Altitude, To: fulcrumAltitude, Rationale: supplierFault}
				}
				ledgerStash = append(ledgerStash, provisionalLedger)
			}
			intensity++

		default:
			return nil, FaultValidationUnsuccessful{Originating: attestedLedger.Altitude, To: ledgerStash[intensity].Altitude, Rationale: err}
		}
	}
}

//
//
func (c *Customer) validateOmittingVersusLeading(
	ctx context.Context,
	reliableLedger *kinds.AgileLedger,
	freshAgileLedger *kinds.AgileLedger,
	now time.Time,
) error {
	logging, err := c.validateOmitting(ctx, c.leading, reliableLedger, freshAgileLedger, now)

	switch errors.Unwrap(err).(type) {
	case FaultUnfitHeadline:
		//
		unfitHeadlineAltitude := err.(FaultValidationUnsuccessful).To
		if unfitHeadlineAltitude == freshAgileLedger.Altitude {
			c.tracer.Diagnose("REDACTED", "REDACTED", err)
			return err
		}

		//
		//
		c.tracer.Failure("REDACTED", "REDACTED", err)

		substitutionLedger, discardFault := c.locateFreshLeading(ctx, freshAgileLedger.Altitude, true)
		if discardFault != nil {
			c.tracer.Failure("REDACTED", "REDACTED", discardFault)
			return err
		}

		if !bytes.Equal(substitutionLedger.Digest(), freshAgileLedger.Digest()) {
			c.tracer.Failure("REDACTED",
				"REDACTED", freshAgileLedger.Digest(),
				"REDACTED", substitutionLedger.Digest())
			//
			return err
		}

		//
		return c.validateOmittingVersusLeading(ctx, reliableLedger, substitutionLedger, now)
	case nil:
		//
		//
		//
		//
		//
		if contrastFault := c.senseDeviation(ctx, logging, now); contrastFault != nil {
			return contrastFault
		}
	default:
		return err
	}

	return nil
}

//
//
//
//
func (c *Customer) FinalReliableAltitude() (int64, error) {
	return c.reliableDepot.FinalAgileLedgerAltitude()
}

//
//
//
//
func (c *Customer) InitialReliableAltitude() (int64, error) {
	return c.reliableDepot.InitialAgileLedgerAltitude()
}

//
//
//
func (c *Customer) SuccessionUUID() string {
	return c.successionUUID
}

//
//
//
func (c *Customer) Leading() supplier.Supplier {
	c.supplierExclusion.Lock()
	defer c.supplierExclusion.Unlock()
	return c.leading
}

//
//
//
func (c *Customer) Attestors() []supplier.Supplier {
	c.supplierExclusion.Lock()
	defer c.supplierExclusion.Unlock()
	return c.attestors
}

//
//
func (c *Customer) Sanitize() error {
	c.tracer.Details("REDACTED")
	c.newestReliableLedger = nil
	return c.reliableDepot.Trim(0)
}

//
//
func (c *Customer) sanitizeSubsequent(altitude int64) error {
	priorAltitude := c.newestReliableLedger.Altitude

	for {
		h, err := c.reliableDepot.AgileLedgerPrior(priorAltitude)
		if err == depot.FaultAgileLedgerNegationDetected || (h != nil && h.Altitude <= altitude) {
			break
		} else if err != nil {
			return fmt.Errorf("REDACTED", priorAltitude, err)
		}

		err = c.reliableDepot.EraseAgileLedger(h.Altitude)
		if err != nil {
			c.tracer.Failure("REDACTED", "REDACTED", err,
				"REDACTED", h.Altitude)
		}

		priorAltitude = h.Altitude
	}

	c.newestReliableLedger = nil
	err := c.recoverReliableAgileLedger()
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) reviseReliableAgileLedger(l *kinds.AgileLedger) error {
	c.tracer.Diagnose("REDACTED", "REDACTED", l)

	if err := c.reliableDepot.PersistAgileLedger(l); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if c.thinningExtent > 0 {
		if err := c.reliableDepot.Trim(c.thinningExtent); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	if c.newestReliableLedger == nil || l.Altitude > c.newestReliableLedger.Altitude {
		c.newestReliableLedger = l
	}

	return nil
}

//
//
//
func (c *Customer) reverse(
	ctx context.Context,
	reliableHeading *kinds.Heading,
	freshHeadline *kinds.Heading,
) error {
	var (
		attestedHeadline = reliableHeading
		provisionalHeadline  *kinds.Heading
	)

	for attestedHeadline.Altitude > freshHeadline.Altitude {
		provisionalLedger, err := c.agileLedgerOriginatingLeading(ctx, attestedHeadline.Altitude-1)
		if err != nil {
			return fmt.Errorf("REDACTED", attestedHeadline.Altitude-1, err)
		}
		provisionalHeadline = provisionalLedger.Heading
		c.tracer.Diagnose("REDACTED",
			"REDACTED", attestedHeadline.Altitude,
			"REDACTED", attestedHeadline.Digest(),
			"REDACTED", provisionalHeadline.Altitude,
			"REDACTED", provisionalHeadline.Digest())
		if err := ValidateReverse(provisionalHeadline, attestedHeadline); err != nil {
			//
			c.tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", c.leading)

			//
			freshMainnodesLedger, supplantFault := c.locateFreshLeading(ctx, freshHeadline.Altitude, true)
			if supplantFault != nil {
				c.tracer.Diagnose("REDACTED", "REDACTED", supplantFault)
				return err
			}

			//
			if !bytes.Equal(freshMainnodesLedger.Digest(), freshHeadline.Digest()) {
				c.tracer.Diagnose("REDACTED")
				//
				return err
			}

			//
			return c.reverse(ctx, attestedHeadline, freshMainnodesLedger.Heading)
		}
		attestedHeadline = provisionalHeadline
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
func (c *Customer) agileLedgerOriginatingLeading(ctx context.Context, altitude int64) (*kinds.AgileLedger, error) {
	c.supplierExclusion.Lock()
	l, err := c.leading.AgileLedger(ctx, altitude)
	c.supplierExclusion.Unlock()

	switch err {
	case nil:
		//
		return l, nil

	case context.Canceled, context.DeadlineExceeded:
		return l, err

	case supplier.FaultNegativeReply, supplier.FaultAgileLedgerNegationDetected, supplier.FaultAltitudeExcessivelyTall:
		//
		c.tracer.Details("REDACTED",
			"REDACTED", err, "REDACTED", altitude, "REDACTED", c.leading)
		return c.locateFreshLeading(ctx, altitude, false)

	default:
		//
		//
		c.tracer.Details("REDACTED",
			"REDACTED", err, "REDACTED", altitude, "REDACTED", c.leading)
		return c.locateFreshLeading(ctx, altitude, true)
	}
}

//
func (c *Customer) discardAttestors(indices []int) error {
	//
	if len(c.attestors) <= len(indices) {
		return FaultNegativeAttestors
	}

	//
	//
	sort.Ints(indices)
	for i := len(indices) - 1; i >= 0; i-- {
		c.attestors[indices[i]] = c.attestors[len(c.attestors)-1]
		c.attestors = c.attestors[:len(c.attestors)-1]
	}

	return nil
}

type attestorReply struct {
	lb           *kinds.AgileLedger
	attestorPosition int
	err          error
}

//
//
//
//
func (c *Customer) locateFreshLeading(ctx context.Context, altitude int64, discard bool) (*kinds.AgileLedger, error) {
	c.supplierExclusion.Lock()
	defer c.supplierExclusion.Unlock()

	if len(c.attestors) == 0 {
		return nil, FaultNegativeAttestors
	}

	var (
		attestorRepliesCN = make(chan attestorReply, len(c.attestors))
		attestorsTowardDiscard []int
		finalFailure         error
		wg                sync.WaitGroup
	)

	//
	subcontext, abort := context.WithCancel(ctx)
	defer abort()
	for ordinal := range c.attestors {
		wg.Add(1)
		go func(attestorPosition int, attestorRepliesCN chan attestorReply) {
			defer wg.Done()

			lb, err := c.attestors[attestorPosition].AgileLedger(subcontext, altitude)
			attestorRepliesCN <- attestorReply{lb, attestorPosition, err}
		}(ordinal, attestorRepliesCN)
	}

	//
	for i := 0; i < cap(attestorRepliesCN); i++ {
		reply := <-attestorRepliesCN
		switch reply.err {
		//
		case nil:
			abort() //

			wg.Wait() //

			//
			if !discard {
				c.attestors = append(c.attestors, c.leading)
			}

			//
			c.tracer.Diagnose("REDACTED", "REDACTED", c.attestors[reply.attestorPosition])
			c.leading = c.attestors[reply.attestorPosition]

			//
			attestorsTowardDiscard = append(attestorsTowardDiscard, reply.attestorPosition)

			//
			//
			if err := c.discardAttestors(attestorsTowardDiscard); err != nil {
				return nil, err
			}

			//
			return reply.lb, nil

		//
		case supplier.FaultNegativeReply, supplier.FaultAgileLedgerNegationDetected, supplier.FaultAltitudeExcessivelyTall:
			finalFailure = reply.err
			c.tracer.Diagnose("REDACTED",
				"REDACTED", reply.err, "REDACTED", c.attestors[reply.attestorPosition])
			continue

		//
		default:
			finalFailure = reply.err
			c.tracer.Failure("REDACTED",
				"REDACTED", reply.err, "REDACTED", c.attestors[reply.attestorPosition])
			attestorsTowardDiscard = append(attestorsTowardDiscard, reply.attestorPosition)
		}
	}

	//
	if err := c.discardAttestors(attestorsTowardDiscard); err != nil {
		c.tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", attestorsTowardDiscard)
	}

	return nil, finalFailure
}

//
//
func (c *Customer) contrastInitialAgileLedgerUsingAttestors(ctx context.Context, l *kinds.AgileLedger) error {
	contrastContext, abort := context.WithCancel(ctx)
	defer abort()

	c.supplierExclusion.Lock()
	defer c.supplierExclusion.Unlock()

	if len(c.attestors) == 0 {
		return FaultNegativeAttestors
	}

	faultchnl := make(chan error, len(c.attestors))
	for i, attestor := range c.attestors {
		go c.contrastFreshAgileLedgerUsingAttestor(contrastContext, faultchnl, l, attestor, i)
	}

	attestorsTowardDiscard := make([]int, 0, len(c.attestors))

	//
	for i := 0; i < cap(faultchnl); i++ {
		err := <-faultchnl

		switch e := err.(type) {
		case nil:
			continue
		case FaultDiscordantHeadings:
			c.tracer.Failure("REDACTED"+
				"REDACTED",
				"REDACTED", c.attestors[e.AttestorPosition], "REDACTED", err)
			return err
		case faultFlawedAttestor:
			//
			c.tracer.Details("REDACTED",
				"REDACTED", c.attestors[e.AttestorPosition],
				"REDACTED", err)
			attestorsTowardDiscard = append(attestorsTowardDiscard, e.AttestorPosition)
		case FaultNominatorUrgenciesDeviate:
			c.tracer.Failure("REDACTED"+
				"REDACTED",
				"REDACTED", c.attestors[e.AttestorPosition], "REDACTED", err)
			return err
		default: //
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return err
			}

			//
			c.tracer.Details("REDACTED",
				"REDACTED", err)
		}

	}

	//
	if err := c.discardAttestors(attestorsTowardDiscard); err != nil {
		c.tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", attestorsTowardDiscard)
	}

	return nil
}
