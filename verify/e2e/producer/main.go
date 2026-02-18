package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/utils/log"
)

const (
	arbitraryOrigin int64 = 4827085738
)

var tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))

func main() {
	NewShell().Run()
}

//
type CLI struct {
	origin *cobra.Command
}

//
func NewShell() *CLI {
	cli := &CLI{}
	cli.origin = &cobra.Command{
		Use:           "REDACTED",
		Short:         "REDACTED",
		SilenceUsage:  true,
		SilenceErrors: true, //
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}
			sets, err := cmd.Flags().GetInt("REDACTED")
			if err != nil {
				return err
			}
			multipleRelease, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}
			monitorstats, err := cmd.Flags().GetBool("REDACTED")
			if err != nil {
				return err
			}
			return cli.compose(dir, sets, multipleRelease, monitorstats)
		},
	}

	cli.origin.PersistentFlags().StringP("REDACTED", "REDACTED", "REDACTED", "REDACTED")
	_ = cli.origin.MarkPersistentFlagRequired("REDACTED")
	cli.origin.PersistentFlags().StringP("REDACTED", "REDACTED", "REDACTED", "REDACTED"+
		"REDACTED")
	cli.origin.PersistentFlags().IntP("REDACTED", "REDACTED", 0, "REDACTED")
	cli.origin.PersistentFlags().BoolP("REDACTED", "REDACTED", false, "REDACTED")

	return cli
}

//
func (cli *CLI) compose(dir string, sets int, multipleRelease string, monitorstats bool) error {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	cfg := &composeSettings{
		randomOrigin:   rand.New(rand.NewSource(arbitraryOrigin)), //
		multipleRelease: multipleRelease,
		monitorstats:   monitorstats,
	}
	declarations, err := Compose(cfg)
	if err != nil {
		return err
	}
	if sets <= 0 {
		for i, declaration := range declarations {
			err = declaration.Persist(filepath.Join(dir, fmt.Sprintf("REDACTED", i)))
			if err != nil {
				return err
			}
		}
	} else {
		clusterVolume := int(math.Ceil(float64(len(declarations)) / float64(sets)))
		for g := 0; g < sets; g++ {
			for i := 0; i < clusterVolume && g*clusterVolume+i < len(declarations); i++ {
				declaration := declarations[g*clusterVolume+i]
				err = declaration.Persist(filepath.Join(dir, fmt.Sprintf("REDACTED", g, i)))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

//
func (cli *CLI) Run() {
	if err := cli.origin.Execute(); err != nil {
		tracer.Fault(err.Error())
		os.Exit(1)
	}
}
