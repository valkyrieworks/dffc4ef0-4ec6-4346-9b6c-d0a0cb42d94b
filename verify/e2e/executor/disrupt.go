package primary

import (
	"context"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform/dock"
)

//
func Disrupt(ctx context.Context, simnet *e2e.Simnet) error {
	for _, peer := range simnet.Peers {
		for _, disruption := range peer.Disruptions {
			_, err := DisruptPeer(ctx, peer, disruption)
			if err != nil {
				return err
			}
			time.Sleep(3 * time.Second) //
		}
	}
	return nil
}

//
//
func DisruptPeer(ctx context.Context, peer *e2e.Peer, disruption e2e.Disruption) (*remoteifacetypes.OutcomeCondition, error) {
	simnet := peer.Simnet
	out, err := dock.InvokeArrangeEmission(context.Background(), simnet.Dir, "REDACTED", "REDACTED", peer.Alias)
	if err != nil {
		return nil, err
	}
	alias := peer.Alias
	modernized := false
	if len(out) == 0 {
		alias += "REDACTED"
		modernized = true
		tracer.Details("REDACTED", "REDACTED",
			log.FreshIdleFormat("REDACTED",
				peer.Alias, alias))
	}

	switch disruption {
	case e2e.DisruptionDetach:
		tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", peer.Alias))
		if err := dock.Invoke(context.Background(), "REDACTED", "REDACTED", simnet.Alias+"REDACTED"+simnet.Alias, alias); err != nil {
			return nil, err
		}
		time.Sleep(10 * time.Second)
		if err := dock.Invoke(context.Background(), "REDACTED", "REDACTED", simnet.Alias+"REDACTED"+simnet.Alias, alias); err != nil {
			return nil, err
		}

	case e2e.DisruptionTerminate:
		tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", peer.Alias))
		if err := dock.InvokeArrange(context.Background(), simnet.Dir, "REDACTED", "REDACTED", "REDACTED", alias); err != nil {
			return nil, err
		}
		if err := dock.InvokeArrange(context.Background(), simnet.Dir, "REDACTED", alias); err != nil {
			return nil, err
		}

	case e2e.DisruptionBreak:
		tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", peer.Alias))
		if err := dock.InvokeArrange(context.Background(), simnet.Dir, "REDACTED", alias); err != nil {
			return nil, err
		}
		time.Sleep(10 * time.Second)
		if err := dock.InvokeArrange(context.Background(), simnet.Dir, "REDACTED", alias); err != nil {
			return nil, err
		}

	case e2e.DisruptionReboot:
		tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", peer.Alias))
		if err := dock.InvokeArrange(context.Background(), simnet.Dir, "REDACTED", alias); err != nil {
			return nil, err
		}

	case e2e.DisruptionModernize:
		agedVER := peer.Edition
		freshVER := peer.Simnet.ModernizeEdition
		if modernized {
			return nil, fmt.Errorf("REDACTED",
				peer.Alias, agedVER, freshVER)
		}
		if agedVER == freshVER {
			tracer.Details("REDACTED", "REDACTED",
				log.FreshIdleFormat("REDACTED",
					peer.Alias, freshVER))
			break
		}
		tracer.Details("REDACTED", "REDACTED",
			log.FreshIdleFormat("REDACTED",
				peer.Alias, agedVER, freshVER))

		if err := dock.InvokeArrange(context.Background(), simnet.Dir, "REDACTED", alias); err != nil {
			return nil, err
		}
		time.Sleep(10 * time.Second)
		if err := dock.InvokeArrange(context.Background(), simnet.Dir, "REDACTED", "REDACTED", alias+"REDACTED"); err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("REDACTED", disruption)
	}

	condition, err := pauseForeachPeer(ctx, peer, 0, 20*time.Second)
	if err != nil {
		return nil, err
	}
	tracer.Details("REDACTED",
		"REDACTED",
		log.FreshIdleFormat("REDACTED", peer.Alias, condition.ChronizeDetails.NewestLedgerAltitude))
	return condition, nil
}
