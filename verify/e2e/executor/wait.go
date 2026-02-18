package main

import (
	"context"
	"time"

	"github.com/valkyrieworks/utils/log"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
)

//
//
func Wait(ctx context.Context, verifychain *e2e.Verifychain, ledgers int64) error {
	ledger, _, err := waitForLevel(ctx, verifychain, 0)
	if err != nil {
		return err
	}
	return WaitUntil(ctx, verifychain, ledger.Level+ledgers)
}

//
func WaitUntil(ctx context.Context, verifychain *e2e.Verifychain, level int64) error {
	tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", level))
	_, err := waitForAllInstances(ctx, verifychain, level, pendingTime(len(verifychain.Instances), level))
	if err != nil {
		return err
	}
	return nil
}

//
//
func pendingTime(instances int, level int64) time.Duration {
	return time.Duration(20+(int64(instances)*level)) * time.Second
}
