package selfpool

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

//
//
//
//
//
//
//
type EfficiencyWaitstateAmplifier struct {
	minimumLaborers int
	maximumLaborers int

	limitQuantile float64
	limitWaitstate    time.Duration
	eraInterval       time.Duration

	//
	eraWaitstates []time.Duration

	weightedavgEfficiency float64

	mu     sync.Mutex
	tracer log.Tracer
}

const (
	MustRemain uint8 = iota
	MustAmplify
	MustReduce
)

func FreshEfficiencyWaitstateAmplifier(
	min, max int,
	limitQuantile float64,
	limitWaitstate time.Duration,
	eraInterval time.Duration,
	tracer log.Tracer,
) *EfficiencyWaitstateAmplifier {
	if min <= 0 {
		min = 4
	}

	if max <= 0 {
		max = min * 2
	}

	if limitQuantile < 0.0 || limitQuantile > 100.0 {
		limitQuantile = 90.0
	}

	if limitWaitstate <= 0 {
		limitWaitstate = 100 * time.Millisecond
	}

	return &EfficiencyWaitstateAmplifier{
		minimumLaborers:          min,
		maximumLaborers:          max,
		limitQuantile: limitQuantile,
		limitWaitstate:    limitWaitstate,
		eraInterval:       eraInterval,
		eraWaitstates:      []time.Duration{},
		weightedavgEfficiency:      0,
		tracer:              tracer.Using("REDACTED", "REDACTED"),
		mu:                  sync.Mutex{},
	}
}

func (s *EfficiencyWaitstateAmplifier) Text() string {
	return fmt.Sprintf(
		"REDACTED",
		s.minimumLaborers,
		s.maximumLaborers,
		s.limitQuantile,
		s.limitWaitstate.Milliseconds(),
	)
}

func (s *EfficiencyWaitstateAmplifier) EraInterval() time.Duration {
	return s.eraInterval
}

func (s *EfficiencyWaitstateAmplifier) Min() int {
	return s.minimumLaborers
}

func (s *EfficiencyWaitstateAmplifier) Max() int {
	return s.maximumLaborers
}

//
func (s *EfficiencyWaitstateAmplifier) Monitor(interval time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.eraWaitstates = append(s.eraWaitstates, interval)
}

//
func (s *EfficiencyWaitstateAmplifier) Resolve(prevailingCountLaborers, stagingLength, stagingLimit int) uint8 {
	s.mu.Lock()
	defer s.mu.Unlock()

	//
	const (
		coefficient     = 0.3
		allowance = 0.1 //
	)

	const (
		//
		stagingStressLimit = 0.6
	)

	var (
		eraEfficiency    = float64(len(s.eraWaitstates))
		eraPeriodQuantile = computeQuantile(s.eraWaitstates, s.limitQuantile)
		tracer             = s.tracer.Using(
			"REDACTED", prevailingCountLaborers,
			"REDACTED", eraEfficiency,
			"REDACTED", s.weightedavgEfficiency,
			"REDACTED", eraPeriodQuantile.Milliseconds(),
		)
	)

	if s.weightedavgEfficiency == 0 {
		s.weightedavgEfficiency = eraEfficiency
	} else {
		freshWeightedavg := coefficient*eraEfficiency + (1-coefficient)*s.weightedavgEfficiency

		s.weightedavgEfficiency = freshWeightedavg
	}

	verdict := MustRemain
	thinking := make([]string, 0, 4)

	//
	switch {
	case eraEfficiency > s.weightedavgEfficiency*(1+allowance):
		verdict = MustAmplify
		thinking = append(thinking, "REDACTED")
	case eraEfficiency < s.weightedavgEfficiency*(1-allowance):
		verdict = MustReduce
		thinking = append(thinking, "REDACTED")
	default:
		thinking = append(thinking, "REDACTED")
	}

	//
	if stagingLimit > 0 {
		stagingStress := float64(stagingLength) / float64(stagingLimit)

		if stagingStress >= stagingStressLimit && verdict != MustAmplify {
			verdict = MustAmplify
			thinking = append(thinking, "REDACTED")
		}
	}

	//
	if verdict == MustAmplify && eraPeriodQuantile >= s.limitWaitstate {
		thinking = append(thinking, "REDACTED")
		verdict = MustReduce
	}

	//
	if verdict == MustAmplify && prevailingCountLaborers >= s.maximumLaborers {
		thinking = append(thinking, "REDACTED")
		verdict = MustRemain
	}

	if verdict == MustReduce && prevailingCountLaborers <= s.minimumLaborers {
		thinking = append(thinking, "REDACTED")
		verdict = MustRemain
	}

	tracer.Diagnose("REDACTED", "REDACTED", strings.Join(thinking, "REDACTED"))

	//
	s.eraWaitstates = make([]time.Duration, 0, len(s.eraWaitstates))

	return verdict
}

func computeQuantile(intervals []time.Duration, quantile float64) time.Duration {
	//
	if quantile < 0.0 || quantile > 100.0 {
		panic("REDACTED")
	}

	if len(intervals) == 0 {
		return 0
	}

	slices.Sort(intervals)

	idx := int(float64(len(intervals)) * quantile / 100.0)
	if idx >= len(intervals) {
		idx = len(intervals) - 1
	}

	return intervals[idx]
}
