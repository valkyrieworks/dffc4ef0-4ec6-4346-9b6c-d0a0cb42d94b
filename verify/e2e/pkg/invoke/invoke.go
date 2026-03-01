package invoke

import (
	"context"
	"fmt"
	"os"
	osexec "os/exec"
)

//
func Directive(ctx context.Context, arguments ...string) error {
	_, err := DirectiveEmission(ctx, arguments...)
	return err
}

//
func DirectiveEmission(ctx context.Context, arguments ...string) ([]byte, error) {
	//
	//
	cmd := osexec.CommandContext(ctx, arguments[0], arguments[1:]...)
	out, err := cmd.CombinedOutput()
	switch err := err.(type) {
	case nil:
		return out, nil
	case *osexec.ExitError:
		return nil, fmt.Errorf("REDACTED", arguments, string(out))
	default:
		return nil, err
	}
}

//
func DirectiveDetailed(ctx context.Context, arguments ...string) error {
	//
	//
	cmd := osexec.CommandContext(ctx, arguments[0], arguments[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
