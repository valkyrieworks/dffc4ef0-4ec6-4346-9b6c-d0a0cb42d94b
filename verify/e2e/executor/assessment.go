package primary

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"path/filepath"
	"time"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
//
//
//
//
//
func Assessment(ctx context.Context, simnet *e2e.Simnet, assessmentMagnitude int64) error {
	ledger, _, err := pauseForeachAltitude(ctx, simnet, 0)
	if err != nil {
		return err
	}

	tracer.Details("REDACTED", "REDACTED", ledger.Altitude)
	initiateLocated := time.Now()

	//
	//
	pausingMoment := time.Duration(assessmentMagnitude*5) * time.Second
	terminateAltitude, err := pauseForeachEveryPeers(ctx, simnet, ledger.Altitude+assessmentMagnitude, pausingMoment)
	if err != nil {
		return err
	}
	dur := time.Since(initiateLocated)

	tracer.Details("REDACTED", "REDACTED", terminateAltitude)

	//
	ledgers, err := acquireLedgerSuccessionSpecimen(ctx, simnet, assessmentMagnitude)
	if err != nil {
		return err
	}

	//
	momentPeriods := partitionWithinLedgerPeriods(ledgers)
	simnetMetrics := deriveSimnetMetrics(momentPeriods)
	simnetMetrics.inhabitTransfers(ledgers)
	simnetMetrics.sumMoment = dur
	simnetMetrics.initiateAltitude = ledgers[0].Heading.Altitude
	simnetMetrics.terminateAltitude = ledgers[len(ledgers)-1].Heading.Altitude

	//
	tracer.Details(simnetMetrics.EmissionJSN(simnet))
	return nil
}

func (t *simnetMetrics) inhabitTransfers(ledgers []*kinds.LedgerSummary) {
	t.transfercount = 0
	for _, b := range ledgers {
		t.transfercount += int64(b.CountTrans)
	}
}

type simnetMetrics struct {
	initiateAltitude int64
	terminateAltitude   int64

	transfercount   int64
	sumMoment time.Duration
	//
	average time.Duration
	//
	std float64
	//
	max time.Duration
	//
	min time.Duration
}

func (t *simnetMetrics) EmissionJSN(net *e2e.Simnet) string {
	jsn, err := json.Marshal(map[string]any{
		"REDACTED":         filepath.Base(net.Record),
		"REDACTED": t.initiateAltitude,
		"REDACTED":   t.terminateAltitude,
		"REDACTED":       t.terminateAltitude - t.initiateAltitude,
		"REDACTED":       t.std,
		"REDACTED":         t.average.Seconds(),
		"REDACTED":          t.max.Seconds(),
		"REDACTED":          t.min.Seconds(),
		"REDACTED":         len(net.Peers),
		"REDACTED":         t.transfercount,
		"REDACTED":          t.sumMoment.Seconds(),
	})
	if err != nil {
		return "REDACTED"
	}

	return string(jsn)
}

func (t *simnetMetrics) Text() string {
	return fmt.Sprintf(`REDACTEDv
REDACTEDv
REDACTEDf
REDACTEDv
REDACTEDv
REDACTED`,
		t.initiateAltitude,
		t.terminateAltitude,
		t.average,
		t.std,
		t.max,
		t.min,
	)
}

//
//
func acquireLedgerSuccessionSpecimen(ctx context.Context, simnet *e2e.Simnet, assessmentMagnitude int64) ([]*kinds.LedgerSummary, error) {
	var ledgers []*kinds.LedgerSummary

	//
	repositoryPeer := simnet.RepositoryPeers()[0]
	c, err := repositoryPeer.Customer()
	if err != nil {
		return nil, err
	}

	//
	s, err := c.Condition(ctx)
	if err != nil {
		return nil, err
	}

	to := s.ChronizeDetails.NewestLedgerAltitude
	originating := to - assessmentMagnitude + 1
	if originating <= simnet.PrimaryAltitude {
		return nil, fmt.Errorf("REDACTED", to)
	}

	//
	for originating < to {
		//
		reply, err := c.LedgerchainDetails(ctx, originating, min(originating+19, to))
		if err != nil {
			return nil, err
		}

		ledgerMetadata := reply.LedgerMetadata
		//
		for i := len(ledgerMetadata) - 1; i >= 0; i-- {
			if ledgerMetadata[i].Heading.Altitude != originating {
				return nil, fmt.Errorf("REDACTED",
					originating,
					ledgerMetadata[i].Heading.Altitude,
				)
			}
			originating++
			ledgers = append(ledgers, ledgerMetadata[i])
		}
	}

	return ledgers, nil
}

func partitionWithinLedgerPeriods(ledgers []*kinds.LedgerSummary) []time.Duration {
	periods := make([]time.Duration, len(ledgers)-1)
	finalMoment := ledgers[0].Heading.Moment
	for i, ledger := range ledgers {
		//
		if i == 0 {
			continue
		}

		periods[i-1] = ledger.Heading.Moment.Sub(finalMoment)
		finalMoment = ledger.Heading.Moment
	}
	return periods
}

func deriveSimnetMetrics(periods []time.Duration) simnetMetrics {
	var (
		sum, average time.Duration
		std       float64
		max       = periods[0]
		min       = periods[0]
	)

	for _, duration := range periods {
		sum += duration

		if duration > max {
			max = duration
		}

		if duration < min {
			min = duration
		}
	}
	average = sum / time.Duration(len(periods))

	for _, duration := range periods {
		variance := (duration - average).Seconds()
		std += variance * variance
	}
	std = math.Sqrt(std / float64(len(periods)))

	return simnetMetrics{
		average: average,
		std:  std,
		max:  max,
		min:  min,
	}
}
