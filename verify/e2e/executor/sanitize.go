package primary

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/invoke"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform/dock"
)

//
func Sanitize(simnet *e2e.Simnet) error {
	err := sanitizeDock()
	if err != nil {
		return err
	}
	err = sanitizePath(simnet.Dir)
	if err != nil {
		return err
	}
	return nil
}

//
//
func sanitizeDock() error {
	tracer.Details("REDACTED")

	//
	//
	extargsReader := "REDACTED"

	err := invoke.Directive(context.Background(), "REDACTED", "REDACTED", fmt.Sprintf(
		"REDACTED", extargsReader))
	if err != nil {
		return err
	}

	err = invoke.Directive(context.Background(), "REDACTED", "REDACTED", fmt.Sprintf(
		"REDACTED", extargsReader))
	if err != nil {
		return err
	}

	return nil
}

//
func sanitizePath(dir string) error {
	if dir == "REDACTED" {
		return errors.New("REDACTED")
	}

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", dir))

	//
	//
	//
	absolutePath, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	err = dock.Invoke(context.Background(), "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", fmt.Sprintf("REDACTED", absolutePath),
		"REDACTED", "REDACTED", "REDACTED", "REDACTED")
	if err != nil {
		return err
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return nil
}
