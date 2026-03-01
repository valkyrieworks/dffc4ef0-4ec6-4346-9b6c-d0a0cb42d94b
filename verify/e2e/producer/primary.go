package primary

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

const (
	unpredictableGerm int64 = 4827085738
)

var tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))

func primary() {
	FreshShell().Run()
}

//
type CLI struct {
	origin *cobra.Command
}

//
func FreshShell() *CLI {
	cli := &CLI{}
	cli.origin = &cobra.Command{
		Use:           "REDACTED",
		Short:         "REDACTED",
		SilenceUsage:  true,
		SilenceErrors: true, //
		RunE: func(cmd *cobra.Command, arguments []string) error {
			dir, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}
			cohorts, err := cmd.Flags().GetInt("REDACTED")
			if err != nil {
				return err
			}
			variedEdition, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}
			titan, err := cmd.Flags().GetBool("REDACTED")
			if err != nil {
				return err
			}
			return cli.compose(dir, cohorts, variedEdition, titan)
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
func (cli *CLI) compose(dir string, cohorts int, variedEdition string, titan bool) error {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	cfg := &composeSettings{
		arbitraryOrigin:   rand.New(rand.NewSource(unpredictableGerm)), //
		variedEdition: variedEdition,
		titan:   titan,
	}
	declarations, err := Compose(cfg)
	if err != nil {
		return err
	}
	if cohorts <= 0 {
		for i, declaration := range declarations {
			err = declaration.Persist(filepath.Join(dir, fmt.Sprintf("REDACTED", i)))
			if err != nil {
				return err
			}
		}
	} else {
		cohortExtent := int(math.Ceil(float64(len(declarations)) / float64(cohorts)))
		for g := 0; g < cohorts; g++ {
			for i := 0; i < cohortExtent && g*cohortExtent+i < len(declarations); i++ {
				declaration := declarations[g*cohortExtent+i]
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
		tracer.Failure(err.Error())
		os.Exit(1)
	}
}
