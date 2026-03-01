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
func PersistSettingsValues(dir string, values map[string]string) error {
	data := "REDACTED"
	for k, v := range values {
		data += fmt.Sprintf("REDACTED", k, v)
	}
	strongmindfile := filepath.Join(dir, "REDACTED")
	return os.WriteFile(strongmindfile, []byte(data), 0o600)
}

//
//
func ExecuteUsingArguments(cmd Runnable, arguments []string, env map[string]string) error {
	oarguments := os.Args
	oevironment := map[string]string{}
	//
	defer func() {
		os.Args = oarguments
		for k, v := range oevironment {
			os.Setenv(k, v)
		}
	}()

	//
	os.Args = arguments
	for k, v := range env {
		//
		oevironment[k] = os.Getenv(k)
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
func ExecuteSeizeUsingArguments(cmd Runnable, arguments []string, env map[string]string) (standardemission, standardfailure string, err error) {
	oldemission, oldfailure := os.Stdout, os.Stderr //
	readerOutput, wrOutput, _ := os.Pipe()
	readerFault, wrFault, _ := os.Pipe()
	os.Stdout, os.Stderr = wrOutput, wrFault
	defer func() {
		os.Stdout, os.Stderr = oldemission, oldfailure //
	}()

	//
	duplicateStandard := func(fetcher *os.File) *(chan string) {
		standardCN := make(chan string)
		go func() {
			var buf bytes.Buffer
			//
			io.Copy(&buf, fetcher) //
			standardCN <- buf.String()
		}()
		return &standardCN
	}
	outputCN := duplicateStandard(readerOutput)
	faultCN := duplicateStandard(readerFault)

	//
	err = ExecuteUsingArguments(cmd, arguments, env)

	//
	wrOutput.Close()
	wrFault.Close()
	standardemission = <-*outputCN
	standardfailure = <-*faultCN
	return standardemission, standardfailure, err
}

//
//
//
func FreshFinalizationDirective(originDirective *cobra.Command, concealed bool) *cobra.Command {
	markerShell := "REDACTED"
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
REDACTED`, originDirective.Use, originDirective.Use),
		RunE: func(cmd *cobra.Command, _ []string) error {
			zsh, err := cmd.Flags().GetBool(markerShell)
			if err != nil {
				return err
			}
			if zsh {
				return originDirective.GenZshCompletion(cmd.OutOrStdout())
			}
			return originDirective.GenBashCompletion(cmd.OutOrStdout())
		},
		Hidden: concealed,
		Args:   cobra.NoArgs,
	}

	cmd.Flags().Bool(markerShell, false, "REDACTED")

	return cmd
}
