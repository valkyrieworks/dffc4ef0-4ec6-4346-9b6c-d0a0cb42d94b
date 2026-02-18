package rapid

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/valkyrieworks/utils/log"
	cometmath "github.com/valkyrieworks/utils/math"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/rapid/source"
	"github.com/valkyrieworks/rapid/depot"
	"github.com/valkyrieworks/kinds"
)

type style byte

const (
	ordered style = iota + 1
	omitting

	standardTrimmingVolume      = 1000
	standardMaximumReprocessTries = 10
	//
	//
	//
	validateOmittingDividend   = 9
	validateOmittingDivisor = 16

	//
	//
	//
	//
	standardMaximumTimerDeviation = 10 * time.Second

	//
	standardMaximumLedgerDelay = 10 * time.Second
)

//
type Setting func(*Customer)

//
//
//
func OrderedValidation() Setting {
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
func OmittingValidation(validateLayer cometmath.Portion) Setting {
	return func(c *Customer) {
		c.validationStyle = omitting
		c.validateLayer = validateLayer
	}
}

//
//
//
//
func TrimmingVolume(h uint16) Setting {
	return func(c *Customer) {
		c.trimmingVolume = h
	}
}

//
//
//
func AttestationFunction(fn func(operation string) bool) Setting {
	return func(c *Customer) {
		c.attestationFn = fn
	}
}

//
func Tracer(l log.Tracer) Setting {
	return func(c *Customer) {
		c.tracer = l
	}
}

//
//
func MaximumReprocessTries(max uint16) Setting {
	return func(c *Customer) {
		c.maximumReprocessTries = max
	}
}

//
//
func MaximumTimerDeviation(d time.Duration) Setting {
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
func MaximumLedgerDelay(d time.Duration) Setting {
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
	ledgerUID          string
	validatingDuration   time.Duration //
	validationStyle style
	validateLayer       cometmath.Portion
	maximumReprocessTries uint16 //
	maximumTimerDeviation    time.Duration
	maximumLedgerDelay      time.Duration

	//
	sourceLock engineconnect.Lock
	//
	leading source.Source
	//
	attestors []source.Source

	//
	validatedDepot depot.Depot
	//
	newestValidatedLedger *kinds.RapidLedger

	//
	trimmingVolume uint16
	//
	attestationFn func(operation string) bool

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
func NewCustomer(
	ctx context.Context,
	ledgerUID string,
	validateOptions ValidateOptions,
	leading source.Source,
	attestors []source.Source,
	validatedDepot depot.Depot,
	options ...Setting,
) (*Customer, error) {
	if err := validateOptions.CertifySimple(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	c, err := NewCustomerFromValidatedDepot(ledgerUID, validateOptions.Duration, leading, attestors, validatedDepot, options...)
	if err != nil {
		return nil, err
	}

	if c.newestValidatedLedger != nil {
		c.tracer.Details("REDACTED")
		if err := c.inspectValidatedHeadingUtilizingSettings(ctx, validateOptions); err != nil {
			return nil, err
		}
	}

	if c.newestValidatedLedger == nil || c.newestValidatedLedger.Level < validateOptions.Level {
		c.tracer.Details("REDACTED")
		if err := c.bootstrapWithRelianceSettings(ctx, validateOptions); err != nil {
			return nil, err
		}
	}

	return c, err
}

//
//
//
func NewCustomerFromValidatedDepot(
	ledgerUID string,
	validatingDuration time.Duration,
	leading source.Source,
	attestors []source.Source,
	validatedDepot depot.Depot,
	options ...Setting,
) (*Customer, error) {
	c := &Customer{
		ledgerUID:          ledgerUID,
		validatingDuration:   validatingDuration,
		validationStyle: omitting,
		validateLayer:       StandardRelianceLayer,
		maximumReprocessTries: standardMaximumReprocessTries,
		maximumTimerDeviation:    standardMaximumTimerDeviation,
		maximumLedgerDelay:      standardMaximumLedgerDelay,
		leading:          leading,
		attestors:        attestors,
		validatedDepot:     validatedDepot,
		trimmingVolume:      standardTrimmingVolume,
		attestationFn:   func(operation string) bool { return true },
		exit:             make(chan struct{}),
		tracer:           log.NewNoopTracer(),
	}

	for _, o := range options {
		o(c)
	}

	//
	if len(c.attestors) == 0 {
		return nil, ErrNoAttestors
	}

	//
	for i, w := range attestors {
		if w.LedgerUID() != ledgerUID {
			return nil, fmt.Errorf("REDACTED",
				i, w, w.LedgerUID(), ledgerUID)
		}
	}

	//
	if err := CertifyRelianceLayer(c.validateLayer); err != nil {
		return nil, err
	}

	if err := c.recoverValidatedRapidLedger(); err != nil {
		return nil, err
	}

	return c, nil
}

//
func (c *Customer) recoverValidatedRapidLedger() error {
	finalLevel, err := c.validatedDepot.FinalRapidLedgerLevel()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if finalLevel > 0 {
		validatedLedger, err := c.validatedDepot.RapidLedger(finalLevel)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		c.newestValidatedLedger = validatedLedger
		c.tracer.Details("REDACTED", "REDACTED", finalLevel)
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
func (c *Customer) inspectValidatedHeadingUtilizingSettings(ctx context.Context, options ValidateOptions) error {
	var leadingDigest []byte
	switch {
	case options.Level > c.newestValidatedLedger.Level:
		h, err := c.rapidLedgerFromLeading(ctx, c.newestValidatedLedger.Level)
		if err != nil {
			return err
		}
		leadingDigest = h.Digest()
	case options.Level == c.newestValidatedLedger.Level:
		leadingDigest = options.Digest
	case options.Level < c.newestValidatedLedger.Level:
		c.tracer.Details("REDACTED",
			"REDACTED", options.Level,
			"REDACTED", c.newestValidatedLedger.Level,
			"REDACTED", c.newestValidatedLedger.Digest())

		operation := fmt.Sprintf(
			"REDACTED",
			options.Level, options.Digest,
			c.newestValidatedLedger.Level, c.newestValidatedLedger.Digest())
		if c.attestationFn(operation) {
			//
			err := c.sanitizeAfter(options.Level)
			if err != nil {
				return fmt.Errorf("REDACTED", options.Level, err)
			}

			c.tracer.Details("REDACTED",
				"REDACTED", options.Level)
		} else {
			return nil
		}

		leadingDigest = options.Digest
	}

	if !bytes.Equal(leadingDigest, c.newestValidatedLedger.Digest()) {
		c.tracer.Details("REDACTED",
			"REDACTED", c.newestValidatedLedger.Digest(), "REDACTED", leadingDigest)

		operation := fmt.Sprintf(
			"REDACTED",
			c.newestValidatedLedger.Digest(), leadingDigest)
		if c.attestationFn(operation) {
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
func (c *Customer) bootstrapWithRelianceSettings(ctx context.Context, options ValidateOptions) error {
	//
	l, err := c.rapidLedgerFromLeading(ctx, options.Level)
	if err != nil {
		return err
	}

	//
	//
	//
	if err := l.CertifySimple(c.ledgerUID); err != nil {
		return err
	}

	if !bytes.Equal(l.Digest(), options.Digest) {
		return fmt.Errorf("REDACTED", options.Digest, l.Digest())
	}

	//
	err = l.RatifierAssign.ValidateEndorseRapid(c.ledgerUID, l.Endorse.LedgerUID, l.Level, l.Endorse)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if err := c.contrastInitialRapidLedgerWithAttestors(ctx, l); err != nil {
		return err
	}

	//
	return c.modifyValidatedRapidLedger(l)
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
func (c *Customer) ValidatedRapidLedger(level int64) (*kinds.RapidLedger, error) {
	level, err := c.contrastWithNewestLevel(level)
	if err != nil {
		return nil, err
	}
	return c.validatedDepot.RapidLedger(level)
}

func (c *Customer) contrastWithNewestLevel(level int64) (int64, error) {
	newestLevel, err := c.FinalValidatedLevel()
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}
	if newestLevel == -1 {
		return 0, errors.New("REDACTED")
	}

	switch {
	case level > newestLevel:
		return 0, fmt.Errorf("REDACTED", newestLevel)
	case level == 0:
		return newestLevel, nil
	case level < 0:
		return 0, errors.New("REDACTED")
	}

	return level, nil
}

//
//
//
func (c *Customer) Modify(ctx context.Context, now time.Time) (*kinds.RapidLedger, error) {
	finalValidatedLevel, err := c.FinalValidatedLevel()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	if finalValidatedLevel == -1 {
		//
		return nil, nil
	}

	newestLedger, err := c.rapidLedgerFromLeading(ctx, 0)
	if err != nil {
		return nil, err
	}

	if newestLedger.Level > finalValidatedLevel {
		err = c.validateRapidLedger(ctx, newestLedger, now)
		if err != nil {
			return nil, err
		}
		c.tracer.Details("REDACTED", "REDACTED", newestLedger.Level, "REDACTED", newestLedger.Digest())
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
func (c *Customer) ValidateRapidLedgerAtLevel(ctx context.Context, level int64, now time.Time) (*kinds.RapidLedger, error) {
	if level <= 0 {
		return nil, errors.New("REDACTED")
	}

	//
	h, err := c.ValidatedRapidLedger(level)
	if err == nil {
		c.tracer.Details("REDACTED", "REDACTED", level, "REDACTED", h.Digest())
		//
		return h, nil
	}

	//
	l, err := c.rapidLedgerFromLeading(ctx, level)
	if err != nil {
		return nil, err
	}

	return l, c.validateRapidLedger(ctx, l, now)
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
func (c *Customer) ValidateHeading(ctx context.Context, newHeading *kinds.Heading, now time.Time) error {
	if newHeading == nil {
		return errors.New("REDACTED")
	}
	if newHeading.Level <= 0 {
		return errors.New("REDACTED")
	}

	//
	l, err := c.ValidatedRapidLedger(newHeading.Level)
	if err == nil {
		//
		if !bytes.Equal(l.Digest(), newHeading.Digest()) {
			return fmt.Errorf("REDACTED", l.Digest(), newHeading.Digest())
		}
		c.tracer.Details("REDACTED",
			"REDACTED", newHeading.Level, "REDACTED", newHeading.Digest())
		return nil
	}

	//
	l, err = c.rapidLedgerFromLeading(ctx, newHeading.Level)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if !bytes.Equal(l.Digest(), newHeading.Digest()) {
		return fmt.Errorf("REDACTED", l.Digest(), newHeading.Digest())
	}

	return c.validateRapidLedger(ctx, l, now)
}

func (c *Customer) validateRapidLedger(ctx context.Context, newRapidLedger *kinds.RapidLedger, now time.Time) error {
	c.tracer.Details("REDACTED", "REDACTED", newRapidLedger.Level, "REDACTED", newRapidLedger.Digest())

	var (
		validateFunction func(ctx context.Context, validated *kinds.RapidLedger, new *kinds.RapidLedger, now time.Time) error
		err        error
	)

	switch c.validationStyle {
	case ordered:
		validateFunction = c.validateOrdered
	case omitting:
		validateFunction = c.validateOmittingVersusLeading
	default:
		panic(fmt.Sprintf("REDACTED", c.validationStyle))
	}

	initialLedgerLevel, err := c.InitialValidatedLevel()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	switch {
	//
	case newRapidLedger.Level >= c.newestValidatedLedger.Level:
		err = validateFunction(ctx, c.newestValidatedLedger, newRapidLedger, now)

	//
	case newRapidLedger.Level < initialLedgerLevel:
		var initialLedger *kinds.RapidLedger
		initialLedger, err = c.validatedDepot.RapidLedger(initialLedgerLevel)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		err = c.reverse(ctx, initialLedger.Heading, newRapidLedger.Heading)

	//
	default:
		var nearestLedger *kinds.RapidLedger
		nearestLedger, err = c.validatedDepot.RapidLedgerPrior(newRapidLedger.Level)
		if err != nil {
			return fmt.Errorf("REDACTED", newRapidLedger.Level, err)
		}
		err = validateFunction(ctx, nearestLedger, newRapidLedger, now)
	}
	if err != nil {
		c.tracer.Fault("REDACTED", "REDACTED", err)
		return err
	}

	//
	return c.modifyValidatedRapidLedger(newRapidLedger)
}

//
func (c *Customer) validateOrdered(
	ctx context.Context,
	validatedLedger *kinds.RapidLedger,
	newRapidLedger *kinds.RapidLedger,
	now time.Time,
) error {
	var (
		certifiedLedger = validatedLedger
		provisionalLedger  *kinds.RapidLedger
		err           error
		track         = []*kinds.RapidLedger{validatedLedger}
	)

	for level := validatedLedger.Level + 1; level <= newRapidLedger.Level; level++ {
		//
		if level == newRapidLedger.Level { //
			provisionalLedger = newRapidLedger
		} else { //
			provisionalLedger, err = c.rapidLedgerFromLeading(ctx, level)
			if err != nil {
				return ErrValidationErrored{From: certifiedLedger.Level, To: level, Cause: err}
			}
		}

		//
		c.tracer.Diagnose("REDACTED",
			"REDACTED", certifiedLedger.Level,
			"REDACTED", certifiedLedger.Digest(),
			"REDACTED", provisionalLedger.Level,
			"REDACTED", provisionalLedger.Digest())

		err = ValidateNeighboring(certifiedLedger.AttestedHeading, provisionalLedger.AttestedHeading, provisionalLedger.RatifierAssign,
			c.validatingDuration, now, c.maximumTimerDeviation)
		if err != nil {
			err := ErrValidationErrored{From: certifiedLedger.Level, To: provisionalLedger.Level, Cause: err}

			switch errors.Unwrap(err).(type) {
			case ErrCorruptHeading:
				//
				if err.To == newRapidLedger.Level {
					c.tracer.Diagnose("REDACTED", "REDACTED", err)
					return err
				}

				//
				//
				c.tracer.Fault("REDACTED", "REDACTED", err)

				substitutionLedger, deleteErr := c.locateNewLeading(ctx, newRapidLedger.Level, true)
				if deleteErr != nil {
					c.tracer.Diagnose("REDACTED", "REDACTED", deleteErr)
					return err
				}

				if !bytes.Equal(substitutionLedger.Digest(), newRapidLedger.Digest()) {
					c.tracer.Fault("REDACTED",
						"REDACTED", newRapidLedger.Digest(),
						"REDACTED", substitutionLedger.Digest())
					//
					return err
				}

				//
				level--

				continue
			default:
				return err
			}
		}

		//
		certifiedLedger = provisionalLedger

		//
		track = append(track, certifiedLedger)
	}

	//
	//
	//
	//
	//
	return c.perceiveDeviation(ctx, track, now)
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
	origin source.Source,
	validatedLedger *kinds.RapidLedger,
	newRapidLedger *kinds.RapidLedger,
	now time.Time,
) ([]*kinds.RapidLedger, error) {
	var (
		ledgerRepository = []*kinds.RapidLedger{newRapidLedger}
		intensity      = 0

		certifiedLedger = validatedLedger
		track         = []*kinds.RapidLedger{validatedLedger}
	)

	for {
		c.tracer.Diagnose("REDACTED",
			"REDACTED", certifiedLedger.Level,
			"REDACTED", certifiedLedger.Digest(),
			"REDACTED", ledgerRepository[intensity].Level,
			"REDACTED", ledgerRepository[intensity].Digest())

		err := Validate(certifiedLedger.AttestedHeading, certifiedLedger.RatifierAssign, ledgerRepository[intensity].AttestedHeading,
			ledgerRepository[intensity].RatifierAssign, c.validatingDuration, now, c.maximumTimerDeviation, c.validateLayer)
		switch err.(type) {
		case nil:
			//
			if intensity == 0 {
				track = append(track, newRapidLedger)
				return track, nil
			}
			//
			certifiedLedger = ledgerRepository[intensity]
			//
			ledgerRepository = ledgerRepository[:intensity]
			//
			intensity = 0
			//
			track = append(track, certifiedLedger)

		case ErrNewValueCollectionCannotBeValidated:
			//
			if intensity == len(ledgerRepository)-1 {
				centerLevel := certifiedLedger.Level + (ledgerRepository[intensity].Level-certifiedLedger.
					Level)*validateOmittingDividend/validateOmittingDivisor
				provisionalLedger, sourceErr := origin.RapidLedger(ctx, centerLevel)
				switch sourceErr {
				case nil:
					ledgerRepository = append(ledgerRepository, provisionalLedger)

				//
				case source.ErrRapidLedgerNegateLocated, source.ErrNoReply, source.ErrLevelTooElevated:
					return nil, err

				//
				//
				default:
					return nil, ErrValidationErrored{From: certifiedLedger.Level, To: centerLevel, Cause: sourceErr}
				}
				ledgerRepository = append(ledgerRepository, provisionalLedger)
			}
			intensity++

		default:
			return nil, ErrValidationErrored{From: certifiedLedger.Level, To: ledgerRepository[intensity].Level, Cause: err}
		}
	}
}

//
//
func (c *Customer) validateOmittingVersusLeading(
	ctx context.Context,
	validatedLedger *kinds.RapidLedger,
	newRapidLedger *kinds.RapidLedger,
	now time.Time,
) error {
	track, err := c.validateOmitting(ctx, c.leading, validatedLedger, newRapidLedger, now)

	switch errors.Unwrap(err).(type) {
	case ErrCorruptHeading:
		//
		corruptHeadingLevel := err.(ErrValidationErrored).To
		if corruptHeadingLevel == newRapidLedger.Level {
			c.tracer.Diagnose("REDACTED", "REDACTED", err)
			return err
		}

		//
		//
		c.tracer.Fault("REDACTED", "REDACTED", err)

		substitutionLedger, deleteErr := c.locateNewLeading(ctx, newRapidLedger.Level, true)
		if deleteErr != nil {
			c.tracer.Fault("REDACTED", "REDACTED", deleteErr)
			return err
		}

		if !bytes.Equal(substitutionLedger.Digest(), newRapidLedger.Digest()) {
			c.tracer.Fault("REDACTED",
				"REDACTED", newRapidLedger.Digest(),
				"REDACTED", substitutionLedger.Digest())
			//
			return err
		}

		//
		return c.validateOmittingVersusLeading(ctx, validatedLedger, substitutionLedger, now)
	case nil:
		//
		//
		//
		//
		//
		if compareErr := c.perceiveDeviation(ctx, track, now); compareErr != nil {
			return compareErr
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
func (c *Customer) FinalValidatedLevel() (int64, error) {
	return c.validatedDepot.FinalRapidLedgerLevel()
}

//
//
//
//
func (c *Customer) InitialValidatedLevel() (int64, error) {
	return c.validatedDepot.InitialRapidLedgerLevel()
}

//
//
//
func (c *Customer) LedgerUID() string {
	return c.ledgerUID
}

//
//
//
func (c *Customer) Leading() source.Source {
	c.sourceLock.Lock()
	defer c.sourceLock.Unlock()
	return c.leading
}

//
//
//
func (c *Customer) Attestors() []source.Source {
	c.sourceLock.Lock()
	defer c.sourceLock.Unlock()
	return c.attestors
}

//
//
func (c *Customer) Sanitize() error {
	c.tracer.Details("REDACTED")
	c.newestValidatedLedger = nil
	return c.validatedDepot.Trim(0)
}

//
//
func (c *Customer) sanitizeAfter(level int64) error {
	priorLevel := c.newestValidatedLedger.Level

	for {
		h, err := c.validatedDepot.RapidLedgerPrior(priorLevel)
		if err == depot.ErrRapidLedgerNegateLocated || (h != nil && h.Level <= level) {
			break
		} else if err != nil {
			return fmt.Errorf("REDACTED", priorLevel, err)
		}

		err = c.validatedDepot.EraseRapidLedger(h.Level)
		if err != nil {
			c.tracer.Fault("REDACTED", "REDACTED", err,
				"REDACTED", h.Level)
		}

		priorLevel = h.Level
	}

	c.newestValidatedLedger = nil
	err := c.recoverValidatedRapidLedger()
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) modifyValidatedRapidLedger(l *kinds.RapidLedger) error {
	c.tracer.Diagnose("REDACTED", "REDACTED", l)

	if err := c.validatedDepot.PersistRapidLedger(l); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if c.trimmingVolume > 0 {
		if err := c.validatedDepot.Trim(c.trimmingVolume); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	if c.newestValidatedLedger == nil || l.Level > c.newestValidatedLedger.Level {
		c.newestValidatedLedger = l
	}

	return nil
}

//
//
//
func (c *Customer) reverse(
	ctx context.Context,
	validatedHeading *kinds.Heading,
	newHeading *kinds.Heading,
) error {
	var (
		certifiedHeading = validatedHeading
		provisionalHeading  *kinds.Heading
	)

	for certifiedHeading.Level > newHeading.Level {
		provisionalLedger, err := c.rapidLedgerFromLeading(ctx, certifiedHeading.Level-1)
		if err != nil {
			return fmt.Errorf("REDACTED", certifiedHeading.Level-1, err)
		}
		provisionalHeading = provisionalLedger.Heading
		c.tracer.Diagnose("REDACTED",
			"REDACTED", certifiedHeading.Level,
			"REDACTED", certifiedHeading.Digest(),
			"REDACTED", provisionalHeading.Level,
			"REDACTED", provisionalHeading.Digest())
		if err := ValidateReverse(provisionalHeading, certifiedHeading); err != nil {
			//
			c.tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", c.leading)

			//
			newLeadersLedger, overrideErr := c.locateNewLeading(ctx, newHeading.Level, true)
			if overrideErr != nil {
				c.tracer.Diagnose("REDACTED", "REDACTED", overrideErr)
				return err
			}

			//
			if !bytes.Equal(newLeadersLedger.Digest(), newHeading.Digest()) {
				c.tracer.Diagnose("REDACTED")
				//
				return err
			}

			//
			return c.reverse(ctx, certifiedHeading, newLeadersLedger.Heading)
		}
		certifiedHeading = provisionalHeading
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
func (c *Customer) rapidLedgerFromLeading(ctx context.Context, level int64) (*kinds.RapidLedger, error) {
	c.sourceLock.Lock()
	l, err := c.leading.RapidLedger(ctx, level)
	c.sourceLock.Unlock()

	switch err {
	case nil:
		//
		return l, nil

	case context.Canceled, context.DeadlineExceeded:
		return l, err

	case source.ErrNoReply, source.ErrRapidLedgerNegateLocated, source.ErrLevelTooElevated:
		//
		c.tracer.Details("REDACTED",
			"REDACTED", err, "REDACTED", level, "REDACTED", c.leading)
		return c.locateNewLeading(ctx, level, false)

	default:
		//
		//
		c.tracer.Details("REDACTED",
			"REDACTED", err, "REDACTED", level, "REDACTED", c.leading)
		return c.locateNewLeading(ctx, level, true)
	}
}

//
func (c *Customer) deleteAttestors(listings []int) error {
	//
	if len(c.attestors) <= len(listings) {
		return ErrNoAttestors
	}

	//
	//
	sort.Ints(listings)
	for i := len(listings) - 1; i >= 0; i-- {
		c.attestors[listings[i]] = c.attestors[len(c.attestors)-1]
		c.attestors = c.attestors[:len(c.attestors)-1]
	}

	return nil
}

type attestorReply struct {
	lb           *kinds.RapidLedger
	attestorOrdinal int
	err          error
}

//
//
//
//
func (c *Customer) locateNewLeading(ctx context.Context, level int64, delete bool) (*kinds.RapidLedger, error) {
	c.sourceLock.Lock()
	defer c.sourceLock.Unlock()

	if len(c.attestors) == 0 {
		return nil, ErrNoAttestors
	}

	var (
		attestorRepliesC = make(chan attestorReply, len(c.attestors))
		attestorsToDelete []int
		finalFault         error
		wg                sync.WaitGroup
	)

	//
	subcontext, revoke := context.WithCancel(ctx)
	defer revoke()
	for ordinal := range c.attestors {
		wg.Add(1)
		go func(attestorOrdinal int, attestorRepliesC chan attestorReply) {
			defer wg.Done()

			lb, err := c.attestors[attestorOrdinal].RapidLedger(subcontext, level)
			attestorRepliesC <- attestorReply{lb, attestorOrdinal, err}
		}(ordinal, attestorRepliesC)
	}

	//
	for i := 0; i < cap(attestorRepliesC); i++ {
		reply := <-attestorRepliesC
		switch reply.err {
		//
		case nil:
			revoke() //

			wg.Wait() //

			//
			if !delete {
				c.attestors = append(c.attestors, c.leading)
			}

			//
			c.tracer.Diagnose("REDACTED", "REDACTED", c.attestors[reply.attestorOrdinal])
			c.leading = c.attestors[reply.attestorOrdinal]

			//
			attestorsToDelete = append(attestorsToDelete, reply.attestorOrdinal)

			//
			//
			if err := c.deleteAttestors(attestorsToDelete); err != nil {
				return nil, err
			}

			//
			return reply.lb, nil

		//
		case source.ErrNoReply, source.ErrRapidLedgerNegateLocated, source.ErrLevelTooElevated:
			finalFault = reply.err
			c.tracer.Diagnose("REDACTED",
				"REDACTED", reply.err, "REDACTED", c.attestors[reply.attestorOrdinal])
			continue

		//
		default:
			finalFault = reply.err
			c.tracer.Fault("REDACTED",
				"REDACTED", reply.err, "REDACTED", c.attestors[reply.attestorOrdinal])
			attestorsToDelete = append(attestorsToDelete, reply.attestorOrdinal)
		}
	}

	//
	if err := c.deleteAttestors(attestorsToDelete); err != nil {
		c.tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", attestorsToDelete)
	}

	return nil, finalFault
}

//
//
func (c *Customer) contrastInitialRapidLedgerWithAttestors(ctx context.Context, l *kinds.RapidLedger) error {
	contrastCtx, revoke := context.WithCancel(ctx)
	defer revoke()

	c.sourceLock.Lock()
	defer c.sourceLock.Unlock()

	if len(c.attestors) == 0 {
		return ErrNoAttestors
	}

	faultc := make(chan error, len(c.attestors))
	for i, attestor := range c.attestors {
		go c.contrastNewRapidLedgerWithAttestor(contrastCtx, faultc, l, attestor, i)
	}

	attestorsToDelete := make([]int, 0, len(c.attestors))

	//
	for i := 0; i < cap(faultc); i++ {
		err := <-faultc

		switch e := err.(type) {
		case nil:
			continue
		case ErrClashingHeadings:
			c.tracer.Fault("REDACTED"+
				"REDACTED",
				"REDACTED", c.attestors[e.AttestorOrdinal], "REDACTED", err)
			return err
		case errFlawedAttestor:
			//
			c.tracer.Details("REDACTED",
				"REDACTED", c.attestors[e.AttestorOrdinal],
				"REDACTED", err)
			attestorsToDelete = append(attestorsToDelete, e.AttestorOrdinal)
		case ErrRecommenderUrgenciesDeviate:
			c.tracer.Fault("REDACTED"+
				"REDACTED",
				"REDACTED", c.attestors[e.AttestorOrdinal], "REDACTED", err)
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
	if err := c.deleteAttestors(attestorsToDelete); err != nil {
		c.tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", attestorsToDelete)
	}

	return nil
}
