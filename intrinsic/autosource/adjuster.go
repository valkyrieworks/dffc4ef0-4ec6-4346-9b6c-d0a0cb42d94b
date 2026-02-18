package autosource

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/valkyrieworks/utils/log"
)

//
//
//
//
//
//
//
type VelocityWaitperiodAdjuster struct {
	minimumOperators int
	maximumOperators int

	limitQuantile float64
	limitWaitperiod    time.Duration
	eraPeriod       time.Duration

	//
	eraWaitperiods []time.Duration

	ewmaVelocity float64

	mu     sync.Mutex
	tracer log.Tracer
}

const (
	MustRemain uint8 = iota
	MustSize
	MustReduce
)

func NewVelocityWaitperiodAdjuster(
	min, max int,
	limitQuantile float64,
	limitWaitperiod time.Duration,
	eraPeriod time.Duration,
	tracer log.Tracer,
) *VelocityWaitperiodAdjuster {
	if min <= 0 {
		min = 4
	}

	if max <= 0 {
		max = min * 2
	}

	if limitQuantile < 0.0 || limitQuantile > 100.0 {
		limitQuantile = 90.0
	}

	if limitWaitperiod <= 0 {
		limitWaitperiod = 100 * time.Millisecond
	}

	return &VelocityWaitperiodAdjuster{
		minimumOperators:          min,
		maximumOperators:          max,
		limitQuantile: limitQuantile,
		limitWaitperiod:    limitWaitperiod,
		eraPeriod:       eraPeriod,
		eraWaitperiods:      []time.Duration{},
		ewmaVelocity:      0,
		tracer:              tracer.With("REDACTED", "REDACTED"),
		mu:                  sync.Mutex{},
	}
}

func (s *VelocityWaitperiodAdjuster) String() string {
	return fmt.Sprintf(
		"REDACTED",
		s.minimumOperators,
		s.maximumOperators,
		s.limitQuantile,
		s.limitWaitperiod.Milliseconds(),
	)
}

func (s *VelocityWaitperiodAdjuster) EraPeriod() time.Duration {
	return s.eraPeriod
}

func (s *VelocityWaitperiodAdjuster) Min() int {
	return s.minimumOperators
}

func (s *VelocityWaitperiodAdjuster) Max() int {
	return s.maximumOperators
}

//
func (s *VelocityWaitperiodAdjuster) Monitor(period time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.eraWaitperiods = append(s.eraWaitperiods, period)
}

//
func (s *VelocityWaitperiodAdjuster) Determine(ongoingCountOperators, bufferSize, bufferLimit int) uint8 {
	s.mu.Lock()
	defer s.mu.Unlock()

	//
	const (
		alpha     = 0.3
		allowance = 0.1 //
	)

	const (
		//
		bufferStressLimit = 0.6
	)

	var (
		eraVelocity    = float64(len(s.eraWaitperiods))
		eraPeriodQuantile = computeQuantile(s.eraWaitperiods, s.limitQuantile)
		tracer             = s.tracer.With(
			"REDACTED", ongoingCountOperators,
			"REDACTED", eraVelocity,
			"REDACTED", s.ewmaVelocity,
			"REDACTED", eraPeriodQuantile.Milliseconds(),
		)
	)

	if s.ewmaVelocity == 0 {
		s.ewmaVelocity = eraVelocity
	} else {
		newEWMA := alpha*eraVelocity + (1-alpha)*s.ewmaVelocity

		s.ewmaVelocity = newEWMA
	}

	verdict := MustRemain
	rationale := make([]string, 0, 4)

	//
	switch {
	case eraVelocity > s.ewmaVelocity*(1+allowance):
		verdict = MustSize
		rationale = append(rationale, "REDACTED")
	case eraVelocity < s.ewmaVelocity*(1-allowance):
		verdict = MustReduce
		rationale = append(rationale, "REDACTED")
	default:
		rationale = append(rationale, "REDACTED")
	}

	//
	if bufferLimit > 0 {
		bufferStress := float64(bufferSize) / float64(bufferLimit)

		if bufferStress >= bufferStressLimit && verdict != MustSize {
			verdict = MustSize
			rationale = append(rationale, "REDACTED")
		}
	}

	//
	if verdict == MustSize && eraPeriodQuantile >= s.limitWaitperiod {
		rationale = append(rationale, "REDACTED")
		verdict = MustReduce
	}

	//
	if verdict == MustSize && ongoingCountOperators >= s.maximumOperators {
		rationale = append(rationale, "REDACTED")
		verdict = MustRemain
	}

	if verdict == MustReduce && ongoingCountOperators <= s.minimumOperators {
		rationale = append(rationale, "REDACTED")
		verdict = MustRemain
	}

	tracer.Diagnose("REDACTED", "REDACTED", strings.Join(rationale, "REDACTED"))

	//
	s.eraWaitperiods = make([]time.Duration, 0, len(s.eraWaitperiods))

	return verdict
}

func computeQuantile(periods []time.Duration, quantile float64) time.Duration {
	//
	if quantile < 0.0 || quantile > 100.0 {
		panic("REDACTED")
	}

	if len(periods) == 0 {
		return 0
	}

	slices.Sort(periods)

	idx := int(float64(len(periods)) * quantile / 100.0)
	if idx >= len(periods) {
		idx = len(periods) - 1
	}

	return periods[idx]
}
