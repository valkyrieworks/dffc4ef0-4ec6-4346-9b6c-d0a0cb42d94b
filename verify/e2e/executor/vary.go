package main

import (
	"context"
	"fmt"
	"time"

	"github.com/valkyrieworks/utils/log"
	rpctypes "github.com/valkyrieworks/rpc/core/kinds"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/e2e/pkg/platform/docker"
)

//
func Vary(ctx context.Context, verifychain *e2e.Verifychain) error {
	for _, member := range verifychain.Instances {
		for _, variation := range member.Variations {
			_, err := VaryMember(ctx, member, variation)
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
func VaryMember(ctx context.Context, member *e2e.Member, variation e2e.Variation) (*rpctypes.OutcomeState, error) {
	verifychain := member.Verifychain
	out, err := docker.InvokeAssembleResult(context.Background(), verifychain.Dir, "REDACTED", "REDACTED", member.Label)
	if err != nil {
		return nil, err
	}
	label := member.Label
	enhanced := false
	if len(out) == 0 {
		label += "REDACTED"
		enhanced = true
		tracer.Details("REDACTED", "REDACTED",
			log.NewIdleFormat("REDACTED",
				member.Label, label))
	}

	switch variation {
	case e2e.VariationDetach:
		tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", member.Label))
		if err := docker.Invoke(context.Background(), "REDACTED", "REDACTED", verifychain.Label+"REDACTED"+verifychain.Label, label); err != nil {
			return nil, err
		}
		time.Sleep(10 * time.Second)
		if err := docker.Invoke(context.Background(), "REDACTED", "REDACTED", verifychain.Label+"REDACTED"+verifychain.Label, label); err != nil {
			return nil, err
		}

	case e2e.VariationTerminate:
		tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", member.Label))
		if err := docker.InvokeAssemble(context.Background(), verifychain.Dir, "REDACTED", "REDACTED", "REDACTED", label); err != nil {
			return nil, err
		}
		if err := docker.InvokeAssemble(context.Background(), verifychain.Dir, "REDACTED", label); err != nil {
			return nil, err
		}

	case e2e.VariationStall:
		tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", member.Label))
		if err := docker.InvokeAssemble(context.Background(), verifychain.Dir, "REDACTED", label); err != nil {
			return nil, err
		}
		time.Sleep(10 * time.Second)
		if err := docker.InvokeAssemble(context.Background(), verifychain.Dir, "REDACTED", label); err != nil {
			return nil, err
		}

	case e2e.VariationReboot:
		tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", member.Label))
		if err := docker.InvokeAssemble(context.Background(), verifychain.Dir, "REDACTED", label); err != nil {
			return nil, err
		}

	case e2e.VariationEnhance:
		agedV := member.Release
		newV := member.Verifychain.EnhanceRelease
		if enhanced {
			return nil, fmt.Errorf("REDACTED",
				member.Label, agedV, newV)
		}
		if agedV == newV {
			tracer.Details("REDACTED", "REDACTED",
				log.NewIdleFormat("REDACTED",
					member.Label, newV))
			break
		}
		tracer.Details("REDACTED", "REDACTED",
			log.NewIdleFormat("REDACTED",
				member.Label, agedV, newV))

		if err := docker.InvokeAssemble(context.Background(), verifychain.Dir, "REDACTED", label); err != nil {
			return nil, err
		}
		time.Sleep(10 * time.Second)
		if err := docker.InvokeAssemble(context.Background(), verifychain.Dir, "REDACTED", "REDACTED", label+"REDACTED"); err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("REDACTED", variation)
	}

	state, err := waitForMember(ctx, member, 0, 20*time.Second)
	if err != nil {
		return nil, err
	}
	tracer.Details("REDACTED",
		"REDACTED",
		log.NewIdleFormat("REDACTED", member.Label, state.AlignDetails.NewestLedgerLevel))
	return state, nil
}
