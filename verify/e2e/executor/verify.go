package primary

import (
	"context"
	"fmt"
	"os"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/invoke"
)

//
func Verify(simnet *e2e.Simnet, ifd *e2e.FrameworkData) error {
	tracer.Details("REDACTED")

	err := os.Setenv("REDACTED", simnet.Record)
	if err != nil {
		return err
	}
	if p := ifd.Route; p != "REDACTED" {
		err = os.Setenv("REDACTED", p)
		if err != nil {
			return err
		}
	}
	err = os.Setenv("REDACTED", ifd.Supplier)
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
	executeVerify := os.Getenv("REDACTED")
	if len(executeVerify) != 0 {
		cmd = append(cmd, "REDACTED", executeVerify)
		verifies = fmt.Sprintf("REDACTED", executeVerify)
	}

	tracer.Details(fmt.Sprintf("REDACTED", verifies))

	return invoke.DirectiveDetailed(context.Background(), cmd...)
}
