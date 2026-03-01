package primary

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform"
)

func Initiate(ctx context.Context, simnet *e2e.Simnet, p platform.Supplier) error {
	if len(simnet.Peers) == 0 {
		return fmt.Errorf("REDACTED")
	}

	//
	//
	peerStaging := simnet.Peers
	sort.SliceStable(peerStaging, func(i, j int) bool {
		a, b := peerStaging[i], peerStaging[j]
		switch {
		case a.Style == b.Style:
			return false
		case a.Style == e2e.StyleGerm:
			return true
		case a.Style == e2e.StyleAssessor && b.Style == e2e.StyleComplete:
			return true
		}
		return false
	})

	sort.SliceStable(peerStaging, func(i, j int) bool {
		return peerStaging[i].InitiateLocated < peerStaging[j].InitiateLocated
	})

	if peerStaging[0].InitiateLocated > 0 {
		return fmt.Errorf("REDACTED")
	}

	//
	tracer.Details("REDACTED")
	peersLocatedNull := make([]*e2e.Peer, 0)
	for len(peerStaging) > 0 && peerStaging[0].InitiateLocated == 0 {
		peersLocatedNull = append(peersLocatedNull, peerStaging[0])
		peerStaging = peerStaging[1:]
	}
	err := p.InitiatePeers(context.Background(), peersLocatedNull...)
	if err != nil {
		return err
	}
	for _, peer := range peersLocatedNull {
		if _, err := pauseForeachPeer(ctx, peer, 0, 15*time.Second); err != nil {
			return err
		}
		if peer.TitanDelegateChannel > 0 {
			tracer.Details("REDACTED", "REDACTED",
				log.FreshIdleFormat("REDACTED",
					peer.Alias,
					peer.OutsideINET,
					peer.DelegateChannel,
					peer.OutsideINET,
					peer.TitanDelegateChannel,
				),
			)
		} else {
			tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED",
				peer.Alias,
				peer.OutsideINET,
				peer.DelegateChannel,
			))
		}
	}

	fabricAltitude := simnet.PrimaryAltitude

	//
	tracer.Details("REDACTED",
		"REDACTED", fabricAltitude,
		"REDACTED", len(simnet.Peers)-len(peerStaging),
		"REDACTED", len(peerStaging))

	ledger, ledgerUUID, err := pauseForeachAltitude(ctx, simnet, fabricAltitude)
	if err != nil {
		return err
	}

	//
	for _, peer := range peerStaging {
		if peer.StatusChronize || peer.Style == e2e.StyleAgile {
			err = ReviseSettingsStatusChronize(peer, ledger.Altitude, ledgerUUID.Digest.Octets())
			if err != nil {
				return err
			}
		}
	}

	for _, peer := range peerStaging {
		if peer.InitiateLocated > fabricAltitude {
			//
			//
			//
			//
			//
			//

			fabricAltitude = peer.InitiateLocated

			tracer.Details("REDACTED",
				"REDACTED", peer.Alias,
				"REDACTED", fabricAltitude)

			if _, _, err := pauseForeachAltitude(ctx, simnet, fabricAltitude); err != nil {
				return err
			}
		}

		tracer.Details("REDACTED", "REDACTED", peer.Alias, "REDACTED", peer.InitiateLocated)

		err := p.InitiatePeers(context.Background(), peer)
		if err != nil {
			return err
		}
		condition, err := pauseForeachPeer(ctx, peer, peer.InitiateLocated, 3*time.Minute)
		if err != nil {
			return err
		}
		tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED",
			peer.Alias, peer.OutsideINET, peer.DelegateChannel, condition.ChronizeDetails.NewestLedgerAltitude))
	}

	return nil
}
