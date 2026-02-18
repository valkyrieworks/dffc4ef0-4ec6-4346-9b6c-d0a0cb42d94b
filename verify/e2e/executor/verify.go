package main

import (
	"context"
	"fmt"
	"os"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/e2e/pkg/invoke"
)

//
func Verify(verifychain *e2e.Verifychain, ifd *e2e.PlatformData) error {
	tracer.Details("REDACTED")

	err := os.Setenv("REDACTED", verifychain.Entry)
	if err != nil {
		return err
	}
	if p := ifd.Route; p != "REDACTED" {
		err = os.Setenv("REDACTED", p)
		if err != nil {
			return err
		}
	}
	err = os.Setenv("REDACTED", ifd.Source)
	if err != nil {
		return err
	}

	cmd := []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	detailed := os.Getenv("REDACTED")
	if detailed == "REDACTED" {
		cmd = append(cmd, "REDACTED")
	}
	cmd = append(cmd, "REDACTED")

	verifies := "REDACTED"
	runVerify := os.Getenv("REDACTED")
	if len(runVerify) != 0 {
		cmd = append(cmd, "REDACTED", runVerify)
		verifies = fmt.Sprintf("REDACTED", runVerify)
	}

	tracer.Details(fmt.Sprintf("REDACTED", verifies))

	return invoke.DirectiveDetailed(context.Background(), cmd...)
}
