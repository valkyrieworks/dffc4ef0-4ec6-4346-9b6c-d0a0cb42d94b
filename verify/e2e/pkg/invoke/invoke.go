package invoke

import (
	"context"
	"fmt"
	"os"
	osexec "os/exec"
)

//
func Directive(ctx context.Context, args ...string) error {
	_, err := DirectiveResult(ctx, args...)
	return err
}

//
func DirectiveResult(ctx context.Context, args ...string) ([]byte, error) {
	//
	//
	cmd := osexec.CommandContext(ctx, args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	switch err := err.(type) {
	case nil:
		return out, nil
	case *osexec.ExitError:
		return nil, fmt.Errorf("REDACTED", args, string(out))
	default:
		return nil, err
	}
}

//
func DirectiveDetailed(ctx context.Context, args ...string) error {
	//
	//
	cmd := osexec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
