package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/valkyrieworks/utils/log"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/e2e/pkg/platform"
)

func Begin(ctx context.Context, verifychain *e2e.Verifychain, p platform.Source) error {
	if len(verifychain.Instances) == 0 {
		return fmt.Errorf("REDACTED")
	}

	//
	//
	memberBuffer := verifychain.Instances
	sort.SliceStable(memberBuffer, func(i, j int) bool {
		a, b := memberBuffer[i], memberBuffer[j]
		switch {
		case a.Style == b.Style:
			return false
		case a.Style == e2e.StyleOrigin:
			return true
		case a.Style == e2e.StyleRatifier && b.Style == e2e.StyleComplete:
			return true
		}
		return false
	})

	sort.SliceStable(memberBuffer, func(i, j int) bool {
		return memberBuffer[i].BeginAt < memberBuffer[j].BeginAt
	})

	if memberBuffer[0].BeginAt > 0 {
		return fmt.Errorf("REDACTED")
	}

	//
	tracer.Details("REDACTED")
	instancesAtNil := make([]*e2e.Member, 0)
	for len(memberBuffer) > 0 && memberBuffer[0].BeginAt == 0 {
		instancesAtNil = append(instancesAtNil, memberBuffer[0])
		memberBuffer = memberBuffer[1:]
	}
	err := p.BeginInstances(context.Background(), instancesAtNil...)
	if err != nil {
		return err
	}
	for _, member := range instancesAtNil {
		if _, err := waitForMember(ctx, member, 0, 15*time.Second); err != nil {
			return err
		}
		if member.MonitorstatsGatewayPort > 0 {
			tracer.Details("REDACTED", "REDACTED",
				log.NewIdleFormat("REDACTED",
					member.Label,
					member.OutsideIP,
					member.GatewayPort,
					member.OutsideIP,
					member.MonitorstatsGatewayPort,
				),
			)
		} else {
			tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED",
				member.Label,
				member.OutsideIP,
				member.GatewayPort,
			))
		}
	}

	fabricLevel := verifychain.PrimaryLevel

	//
	tracer.Details("REDACTED",
		"REDACTED", fabricLevel,
		"REDACTED", len(verifychain.Instances)-len(memberBuffer),
		"REDACTED", len(memberBuffer))

	ledger, ledgerUID, err := waitForLevel(ctx, verifychain, fabricLevel)
	if err != nil {
		return err
	}

	//
	for _, member := range memberBuffer {
		if member.StatusAlign || member.Style == e2e.StyleRapid {
			err = ModifySettingsStatusAlign(member, ledger.Level, ledgerUID.Digest.Octets())
			if err != nil {
				return err
			}
		}
	}

	for _, member := range memberBuffer {
		if member.BeginAt > fabricLevel {
			//
			//
			//
			//
			//
			//

			fabricLevel = member.BeginAt

			tracer.Details("REDACTED",
				"REDACTED", member.Label,
				"REDACTED", fabricLevel)

			if _, _, err := waitForLevel(ctx, verifychain, fabricLevel); err != nil {
				return err
			}
		}

		tracer.Details("REDACTED", "REDACTED", member.Label, "REDACTED", member.BeginAt)

		err := p.BeginInstances(context.Background(), member)
		if err != nil {
			return err
		}
		state, err := waitForMember(ctx, member, member.BeginAt, 3*time.Minute)
		if err != nil {
			return err
		}
		tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED",
			member.Label, member.OutsideIP, member.GatewayPort, state.AlignDetails.NewestLedgerLevel))
	}

	return nil
}
