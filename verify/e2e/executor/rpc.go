package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	rpctypes "github.com/valkyrieworks/rpc/core/kinds"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/kinds"
)

//
//
//
func waitForLevel(ctx context.Context, verifychain *e2e.Verifychain, level int64) (*kinds.Ledger, *kinds.LedgerUID, error) {
	var (
		err          error
		maximumOutcome    *rpctypes.OutcomeLedger
		agents      = map[string]*rpchttp.HTTP{}
		finalAugment = time.Now()
	)

	clock := time.NewTimer(0)
	defer clock.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-clock.C:
			for _, member := range verifychain.Instances {
				if member.Untracked() {
					continue
				}

				customer, ok := agents[member.Label]
				if !ok {
					customer, err = member.Customer()
					if err != nil {
						continue
					}
					agents[member.Label] = customer
				}

				subcontext, revoke := context.WithTimeout(ctx, 1*time.Second)
				defer revoke()

				outcome, err := customer.Ledger(subcontext, nil)
				if err == context.DeadlineExceeded || err == context.Canceled {
					return nil, nil, ctx.Err()
				}
				if err != nil {
					continue
				}
				if outcome.Ledger != nil && (maximumOutcome == nil || outcome.Ledger.Level > maximumOutcome.Ledger.Level) {
					maximumOutcome = outcome
					finalAugment = time.Now()
				}
				if maximumOutcome != nil && maximumOutcome.Ledger.Level >= level {
					return maximumOutcome.Ledger, &maximumOutcome.LedgerUID, nil
				}
			}

			if len(agents) == 0 {
				return nil, nil, errors.New("REDACTED")
			}
			if time.Since(finalAugment) >= 20*time.Second {
				if maximumOutcome == nil {
					return nil, nil, errors.New("REDACTED")
				}
				return nil, nil, fmt.Errorf("REDACTED", maximumOutcome.Ledger.Level)
			}
			clock.Reset(1 * time.Second)
		}
	}
}

//
func waitForMember(ctx context.Context, member *e2e.Member, level int64, deadline time.Duration) (*rpctypes.OutcomeState, error) {
	customer, err := member.Customer()
	if err != nil {
		return nil, err
	}

	clock := time.NewTimer(0)
	defer clock.Stop()
	var currentLevel int64
	finalModified := time.Now()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-clock.C:
			state, err := customer.Status(ctx)
			switch {
			case time.Since(finalModified) > deadline:
				return nil, fmt.Errorf("REDACTED", member.Label, level)
			case err != nil:
			case state.AlignDetails.NewestLedgerLevel >= level && (level == 0 || !state.AlignDetails.TrappingUp):
				return state, nil
			case currentLevel < state.AlignDetails.NewestLedgerLevel:
				currentLevel = state.AlignDetails.NewestLedgerLevel
				finalModified = time.Now()
			}

			clock.Reset(300 * time.Millisecond)
		}
	}
}

//
func waitForAllInstances(ctx context.Context, verifychain *e2e.Verifychain, level int64, deadline time.Duration) (int64, error) {
	var finalLevel int64

	limit := time.Now().Add(deadline)

	for _, member := range verifychain.Instances {
		if member.Style == e2e.StyleOrigin {
			continue
		}

		state, err := waitForMember(ctx, member, level, time.Until(limit))
		if err != nil {
			return 0, err
		}

		if state.AlignDetails.NewestLedgerLevel > finalLevel {
			finalLevel = state.AlignDetails.NewestLedgerLevel
		}
	}

	return finalLevel, nil
}
