package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//
//
func RecordSettingsValues(dir string, values map[string]string) error {
	data := "REDACTED"
	for k, v := range values {
		data += fmt.Sprintf("REDACTED", k, v)
	}
	centry := filepath.Join(dir, "REDACTED")
	return os.WriteFile(centry, []byte(data), 0o600)
}

//
//
func ExecuteWithArgs(cmd Runnable, args []string, env map[string]string) error {
	oargs := os.Args
	oenv := map[string]string{}
	//
	defer func() {
		os.Args = oargs
		for k, v := range oenv {
			os.Setenv(k, v)
		}
	}()

	//
	os.Args = args
	for k, v := range env {
		//
		oenv[k] = os.Getenv(k)
		err := os.Setenv(k, v)
		if err != nil {
			return err
		}
	}

	//
	return cmd.Perform()
}

//
//
//
//
func RunSeizeWithArgs(cmd Runnable, args []string, env map[string]string) (stdout, stderr string, err error) {
	oldout, olderr := os.Stdout, os.Stderr //
	readerOut, writerOut, _ := os.Pipe()
	readerErr, writerErr, _ := os.Pipe()
	os.Stdout, os.Stderr = writerOut, writerErr
	defer func() {
		os.Stdout, os.Stderr = oldout, olderr //
	}()

	//
	cloneStd := func(scanner *os.File) *(chan string) {
		stdC := make(chan string)
		go func() {
			var buf bytes.Buffer
			//
			io.Copy(&buf, scanner) //
			stdC <- buf.String()
		}()
		return &stdC
	}
	outC := cloneStd(readerOut)
	errC := cloneStd(readerErr)

	//
	err = ExecuteWithArgs(cmd, args, env)

	//
	writerOut.Close()
	writerErr.Close()
	stdout = <-*outC
	stderr = <-*errC
	return stdout, stderr, err
}

//
//
//
func NewFinalizationCommand(originCommand *cobra.Command, concealed bool) *cobra.Command {
	markZsh := "REDACTED"
	cmd := &cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		Long: fmt.Sprintf(`REDACTED.

REDACTEDs
REDACTED:

REDACTED)

REDACTEDo
REDACTED:

REDACTED)
REDACTED`, originCommand.Use, originCommand.Use),
		RunE: func(cmd *cobra.Command, _ []string) error {
			zsh, err := cmd.Flags().GetBool(markZsh)
			if err != nil {
				return err
			}
			if zsh {
				return originCommand.GenZshCompletion(cmd.OutOrStdout())
			}
			return originCommand.GenBashCompletion(cmd.OutOrStdout())
		},
		Hidden: concealed,
		Args:   cobra.NoArgs,
	}

	cmd.Flags().Bool(markZsh, false, "REDACTED")

	return cmd
}
