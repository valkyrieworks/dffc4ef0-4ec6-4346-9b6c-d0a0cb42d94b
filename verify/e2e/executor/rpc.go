package primary

import (
	"context"
	"errors"
	"fmt"
	"time"

	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
func pauseForeachAltitude(ctx context.Context, simnet *e2e.Simnet, altitude int64) (*kinds.Ledger, *kinds.LedgerUUID, error) {
	var (
		err          error
		maximumOutcome    *remoteifacetypes.OutcomeLedger
		customers      = map[string]*rpchttpsvc.Httpsvc{}
		finalAugment = time.Now()
	)

	clock := time.NewTimer(0)
	defer clock.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-clock.C:
			for _, peer := range simnet.Peers {
				if peer.Untracked() {
					continue
				}

				customer, ok := customers[peer.Alias]
				if !ok {
					customer, err = peer.Customer()
					if err != nil {
						continue
					}
					customers[peer.Alias] = customer
				}

				subcontext, abort := context.WithTimeout(ctx, 1*time.Second)
				defer abort()

				outcome, err := customer.Ledger(subcontext, nil)
				if err == context.DeadlineExceeded || err == context.Canceled {
					return nil, nil, ctx.Err()
				}
				if err != nil {
					continue
				}
				if outcome.Ledger != nil && (maximumOutcome == nil || outcome.Ledger.Altitude > maximumOutcome.Ledger.Altitude) {
					maximumOutcome = outcome
					finalAugment = time.Now()
				}
				if maximumOutcome != nil && maximumOutcome.Ledger.Altitude >= altitude {
					return maximumOutcome.Ledger, &maximumOutcome.LedgerUUID, nil
				}
			}

			if len(customers) == 0 {
				return nil, nil, errors.New("REDACTED")
			}
			if time.Since(finalAugment) >= 20*time.Second {
				if maximumOutcome == nil {
					return nil, nil, errors.New("REDACTED")
				}
				return nil, nil, fmt.Errorf("REDACTED", maximumOutcome.Ledger.Altitude)
			}
			clock.Reset(1 * time.Second)
		}
	}
}

//
func pauseForeachPeer(ctx context.Context, peer *e2e.Peer, altitude int64, deadline time.Duration) (*remoteifacetypes.OutcomeCondition, error) {
	customer, err := peer.Customer()
	if err != nil {
		return nil, err
	}

	clock := time.NewTimer(0)
	defer clock.Stop()
	var currentAltitude int64
	finalAltered := time.Now()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-clock.C:
			condition, err := customer.Condition(ctx)
			switch {
			case time.Since(finalAltered) > deadline:
				return nil, fmt.Errorf("REDACTED", peer.Alias, altitude)
			case err != nil:
			case condition.ChronizeDetails.NewestLedgerAltitude >= altitude && (altitude == 0 || !condition.ChronizeDetails.ObtainingAscend):
				return condition, nil
			case currentAltitude < condition.ChronizeDetails.NewestLedgerAltitude:
				currentAltitude = condition.ChronizeDetails.NewestLedgerAltitude
				finalAltered = time.Now()
			}

			clock.Reset(300 * time.Millisecond)
		}
	}
}

//
func pauseForeachEveryPeers(ctx context.Context, simnet *e2e.Simnet, altitude int64, deadline time.Duration) (int64, error) {
	var finalAltitude int64

	limit := time.Now().Add(deadline)

	for _, peer := range simnet.Peers {
		if peer.Style == e2e.StyleGerm {
			continue
		}

		condition, err := pauseForeachPeer(ctx, peer, altitude, time.Until(limit))
		if err != nil {
			return 0, err
		}

		if condition.ChronizeDetails.NewestLedgerAltitude > finalAltitude {
			finalAltitude = condition.ChronizeDetails.NewestLedgerAltitude
		}
	}

	return finalAltitude, nil
}
