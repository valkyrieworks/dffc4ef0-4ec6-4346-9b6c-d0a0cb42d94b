package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/utils/log"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/e2e/pkg/infrastructure"
	"github.com/valkyrieworks/verify/e2e/pkg/infrastructure/cloudprovider"
	"github.com/valkyrieworks/verify/e2e/pkg/infrastructure/docker"
)

const arbitraryOrigin = 2308084734268

var tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))

func main() {
	NewShell().Run()
}

//
type CLI struct {
	origin     *cobra.Command
	verifychain  *e2e.Verifychain
	conserve bool
	infp     infrastructure.Source
}

//
func NewShell() *CLI {
	cli := &CLI{}
	cli.origin = &cobra.Command{
		Use:           "REDACTED",
		Short:         "REDACTED",
		SilenceUsage:  true,
		SilenceErrors: true, //
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			entry, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}
			m, err := e2e.ImportDeclaration(entry)
			if err != nil {
				return err
			}

			inft, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}

			var ifd e2e.PlatformData
			switch inft {
			case "REDACTED":
				var err error
				ifd, err = e2e.NewDockerPlatformData(m)
				if err != nil {
					return err
				}
			case "REDACTED":
				p, err := cmd.Flags().GetString("REDACTED")
				if err != nil {
					return err
				}
				if p == "REDACTED" {
					return errors.New("REDACTED")
				}
				ifd, err = e2e.PlatformDataFromEntry(p)
				if err != nil {
					return fmt.Errorf("REDACTED", err)
				}
			default:
				return fmt.Errorf("REDACTED", inft)
			}

			verifychain, err := e2e.ImportVerifychain(entry, ifd)
			if err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			cli.verifychain = verifychain
			switch inft {
			case "REDACTED":
				cli.infp = &docker.Source{
					SourceData: infrastructure.SourceData{
						Verifychain:            verifychain,
						PlatformData: ifd,
					},
				}
			case "REDACTED":
				cli.infp = &cloudprovider.Source{
					SourceData: infrastructure.SourceData{
						Verifychain:            verifychain,
						PlatformData: ifd,
					},
				}
			default:
				return fmt.Errorf("REDACTED", inft)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := Sanitize(cli.verifychain); err != nil {
				return err
			}
			if err := Configure(cli.verifychain, cli.infp); err != nil {
				return err
			}

			r := rand.New(rand.NewSource(arbitraryOrigin)) //

			chanImportOutcome := make(chan error)
			ctx, importRevoke := context.WithCancel(context.Background())
			defer importRevoke()
			go func() {
				err := Import(ctx, cli.verifychain)
				if err != nil {
					tracer.Fault(fmt.Sprintf("REDACTED", err.Error()))
				}
				chanImportOutcome <- err
			}()

			if err := Begin(cmd.Context(), cli.verifychain, cli.infp); err != nil {
				return err
			}

			if err := Wait(cmd.Context(), cli.verifychain, 5); err != nil { //
				return err
			}

			if cli.verifychain.HasVariations() {
				if err := Disrupt(cmd.Context(), cli.verifychain); err != nil {
					return err
				}
				if err := Wait(cmd.Context(), cli.verifychain, 5); err != nil { //
					return err
				}
			}

			if cli.verifychain.Proof > 0 {
				if err := InsertProof(ctx, r, cli.verifychain, cli.verifychain.Proof); err != nil {
					return err
				}
				if err := Wait(cmd.Context(), cli.verifychain, 5); err != nil { //
					return err
				}
			}

			importRevoke()
			if err := <-chanImportOutcome; err != nil {
				return err
			}
			if err := Wait(cmd.Context(), cli.verifychain, 5); err != nil { //
				return err
			}
			if err := Verify(cli.verifychain, cli.infp.FetchPlatformData()); err != nil {
				return err
			}
			if !cli.conserve {
				if err := Sanitize(cli.verifychain); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cli.origin.PersistentFlags().StringP("REDACTED", "REDACTED", "REDACTED", "REDACTED")
	_ = cli.origin.MarkPersistentFlagRequired("REDACTED")

	cli.origin.PersistentFlags().StringP("REDACTED", "REDACTED", "REDACTED", "REDACTED")

	cli.origin.PersistentFlags().StringP("REDACTED", "REDACTED", "REDACTED", "REDACTED")

	cli.origin.Flags().BoolVarP(&cli.conserve, "REDACTED", "REDACTED", false,
		"REDACTED")

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Configure(cli.verifychain, cli.infp)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := os.Stat(cli.verifychain.Dir)
			if os.IsNotExist(err) {
				err = Configure(cli.verifychain, cli.infp)
			}
			if err != nil {
				return err
			}
			return Begin(cmd.Context(), cli.verifychain, cli.infp)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Disrupt(cmd.Context(), cli.verifychain)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Wait(cmd.Context(), cli.verifychain, 5)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			tracer.Details("REDACTED")
			return cli.infp.HaltVerifychain(context.Background())
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return Import(context.Background(), cli.verifychain)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Args:  cobra.MaximumNArgs(1),
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			quantity := 1

			if len(args) == 1 {
				quantity, err = strconv.Atoi(args[0])
				if err != nil {
					return err
				}
			}

			return InsertProof(
				cmd.Context(),
				rand.New(rand.NewSource(arbitraryOrigin)), //
				cli.verifychain,
				quantity,
			)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Verify(cli.verifychain, cli.infp.FetchPlatformData())
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Sanitize(cli.verifychain)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			return docker.InvokeAssembleDetailed(context.Background(), cli.verifychain.Dir, "REDACTED")
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			return docker.InvokeAssembleDetailed(context.Background(), cli.verifychain.Dir, "REDACTED", "REDACTED")
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		Long: `REDACTED:
REDACTEDl
REDACTEDn
REDACTEDl
REDACTEDl
REDACTED.
REDACTED	
REDACTED.
REDACTED`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := Sanitize(cli.verifychain); err != nil {
				return err
			}
			if err := Configure(cli.verifychain, cli.infp); err != nil {
				return err
			}

			chanImportOutcome := make(chan error)
			ctx, importRevoke := context.WithCancel(cmd.Context())
			defer importRevoke()
			go func() {
				err := Import(ctx, cli.verifychain)
				if err != nil {
					tracer.Fault(fmt.Sprintf("REDACTED", err.Error()))
				}
				chanImportOutcome <- err
			}()

			if err := Begin(cmd.Context(), cli.verifychain, cli.infp); err != nil {
				return err
			}

			if err := Wait(cmd.Context(), cli.verifychain, 5); err != nil { //
				return err
			}

			//
			if err := Criterion(cmd.Context(), cli.verifychain, 100); err != nil {
				return err
			}

			importRevoke()
			if err := <-chanImportOutcome; err != nil {
				return err
			}

			return Sanitize(cli.verifychain)
		},
	})

	return cli
}

//
func (cli *CLI) Run() {
	if err := cli.origin.Execute(); err != nil {
		tracer.Fault(err.Error())
		os.Exit(1)
	}
}
