package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/valkyrieworks/utils/log"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/e2e/pkg/invoke"
	"github.com/valkyrieworks/verify/e2e/pkg/platform/docker"
)

//
func Sanitize(verifychain *e2e.Verifychain) error {
	err := sanitizeDocker()
	if err != nil {
		return err
	}
	err = sanitizeFolder(verifychain.Dir)
	if err != nil {
		return err
	}
	return nil
}

//
//
func sanitizeDocker() error {
	tracer.Details("REDACTED")

	//
	//
	xargsReader := "REDACTED"

	err := invoke.Directive(context.Background(), "REDACTED", "REDACTED", fmt.Sprintf(
		"REDACTED", xargsReader))
	if err != nil {
		return err
	}

	err = invoke.Directive(context.Background(), "REDACTED", "REDACTED", fmt.Sprintf(
		"REDACTED", xargsReader))
	if err != nil {
		return err
	}

	return nil
}

//
func sanitizeFolder(dir string) error {
	if dir == "REDACTED" {
		return errors.New("REDACTED")
	}

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", dir))

	//
	//
	//
	absoluteFolder, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	err = docker.Invoke(context.Background(), "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", fmt.Sprintf("REDACTED", absoluteFolder),
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
