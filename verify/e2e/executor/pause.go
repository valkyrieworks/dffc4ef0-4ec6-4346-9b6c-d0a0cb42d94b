package primary

import (
	"context"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
)

//
//
func Pause(ctx context.Context, simnet *e2e.Simnet, ledgers int64) error {
	ledger, _, err := pauseForeachAltitude(ctx, simnet, 0)
	if err != nil {
		return err
	}
	return PauseTill(ctx, simnet, ledger.Altitude+ledgers)
}

//
func PauseTill(ctx context.Context, simnet *e2e.Simnet, altitude int64) error {
	tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", altitude))
	_, err := pauseForeachEveryPeers(ctx, simnet, altitude, pausingMoment(len(simnet.Peers), altitude))
	if err != nil {
		return err
	}
	return nil
}

//
//
func pausingMoment(peers int, altitude int64) time.Duration {
	return time.Duration(20+(int64(peers)*altitude)) * time.Second
}
