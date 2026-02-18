package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"path/filepath"
	"time"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/kinds"
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
func Criterion(ctx context.Context, verifychain *e2e.Verifychain, criterionExtent int64) error {
	ledger, _, err := waitForLevel(ctx, verifychain, 0)
	if err != nil {
		return err
	}

	tracer.Details("REDACTED", "REDACTED", ledger.Level)
	beginAt := time.Now()

	//
	//
	pendingTime := time.Duration(criterionExtent*5) * time.Second
	terminateLevel, err := waitForAllInstances(ctx, verifychain, ledger.Level+criterionExtent, pendingTime)
	if err != nil {
		return err
	}
	dur := time.Since(beginAt)

	tracer.Details("REDACTED", "REDACTED", terminateLevel)

	//
	ledgers, err := acquireLedgerSeriesSpecimen(ctx, verifychain, criterionExtent)
	if err != nil {
		return err
	}

	//
	timePeriods := divideTowardLedgerPeriods(ledgers)
	verifychainMetrics := retrieveVerifychainMetrics(timePeriods)
	verifychainMetrics.fillTransfers(ledgers)
	verifychainMetrics.sumTime = dur
	verifychainMetrics.beginLevel = ledgers[0].Heading.Level
	verifychainMetrics.terminateLevel = ledgers[len(ledgers)-1].Heading.Level

	//
	tracer.Details(verifychainMetrics.ResultJSON(verifychain))
	return nil
}

func (t *verifychainMetrics) fillTransfers(ledgers []*kinds.LedgerMeta) {
	t.numtransfers = 0
	for _, b := range ledgers {
		t.numtransfers += int64(b.CountTrans)
	}
}

type verifychainMetrics struct {
	beginLevel int64
	terminateLevel   int64

	numtransfers   int64
	sumTime time.Duration
	//
	average time.Duration
	//
	std float64
	//
	max time.Duration
	//
	min time.Duration
}

func (t *verifychainMetrics) ResultJSON(net *e2e.Verifychain) string {
	jsn, err := json.Marshal(map[string]any{
		"REDACTED":         filepath.Base(net.Entry),
		"REDACTED": t.beginLevel,
		"REDACTED":   t.terminateLevel,
		"REDACTED":       t.terminateLevel - t.beginLevel,
		"REDACTED":       t.std,
		"REDACTED":         t.average.Seconds(),
		"REDACTED":          t.max.Seconds(),
		"REDACTED":          t.min.Seconds(),
		"REDACTED":         len(net.Instances),
		"REDACTED":         t.numtransfers,
		"REDACTED":          t.sumTime.Seconds(),
	})
	if err != nil {
		return "REDACTED"
	}

	return string(jsn)
}

func (t *verifychainMetrics) String() string {
	return fmt.Sprintf(`REDACTEDv
REDACTEDv
REDACTEDf
REDACTEDv
REDACTEDv
REDACTED`,
		t.beginLevel,
		t.terminateLevel,
		t.average,
		t.std,
		t.max,
		t.min,
	)
}

//
//
func acquireLedgerSeriesSpecimen(ctx context.Context, verifychain *e2e.Verifychain, criterionExtent int64) ([]*kinds.LedgerMeta, error) {
	var ledgers []*kinds.LedgerMeta

	//
	catalogMember := verifychain.CatalogInstances()[0]
	c, err := catalogMember.Customer()
	if err != nil {
		return nil, err
	}

	//
	s, err := c.Status(ctx)
	if err != nil {
		return nil, err
	}

	to := s.AlignDetails.NewestLedgerLevel
	from := to - criterionExtent + 1
	if from <= verifychain.PrimaryLevel {
		return nil, fmt.Errorf("REDACTED", to)
	}

	//
	for from < to {
		//
		reply, err := c.LedgerchainDetails(ctx, from, min(from+19, to))
		if err != nil {
			return nil, err
		}

		ledgerMetadata := reply.LedgerMetadata
		//
		for i := len(ledgerMetadata) - 1; i >= 0; i-- {
			if ledgerMetadata[i].Heading.Level != from {
				return nil, fmt.Errorf("REDACTED",
					from,
					ledgerMetadata[i].Heading.Level,
				)
			}
			from++
			ledgers = append(ledgers, ledgerMetadata[i])
		}
	}

	return ledgers, nil
}

func divideTowardLedgerPeriods(ledgers []*kinds.LedgerMeta) []time.Duration {
	periods := make([]time.Duration, len(ledgers)-1)
	finalTime := ledgers[0].Heading.Time
	for i, ledger := range ledgers {
		//
		if i == 0 {
			continue
		}

		periods[i-1] = ledger.Heading.Time.Sub(finalTime)
		finalTime = ledger.Heading.Time
	}
	return periods
}

func retrieveVerifychainMetrics(periods []time.Duration) verifychainMetrics {
	var (
		sum, average time.Duration
		std       float64
		max       = periods[0]
		min       = periods[0]
	)

	for _, cadence := range periods {
		sum += cadence

		if cadence > max {
			max = cadence
		}

		if cadence < min {
			min = cadence
		}
	}
	average = sum / time.Duration(len(periods))

	for _, cadence := range periods {
		vary := (cadence - average).Seconds()
		std += vary * vary
	}
	std = math.Sqrt(std / float64(len(periods)))

	return verifychainMetrics{
		average: average,
		std:  std,
		max:  max,
		min:  min,
	}
}
