package primary

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform/digitalregion"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform/dock"
)

const unpredictableGerm = 2308084734268

var tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))

func primary() {
	FreshShell().Run()
}

//
type CLI struct {
	origin     *cobra.Command
	simnet  *e2e.Simnet
	maintain bool
	infpnt     platform.Supplier
}

//
func FreshShell() *CLI {
	cli := &CLI{}
	cli.origin = &cobra.Command{
		Use:           "REDACTED",
		Short:         "REDACTED",
		SilenceUsage:  true,
		SilenceErrors: true, //
		PersistentPreRunE: func(cmd *cobra.Command, arguments []string) error {
			record, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}
			m, err := e2e.FetchDeclaration(record)
			if err != nil {
				return err
			}

			infrstr, err := cmd.Flags().GetString("REDACTED")
			if err != nil {
				return err
			}

			var ifd e2e.FrameworkData
			switch infrstr {
			case "REDACTED":
				var err error
				ifd, err = e2e.FreshDockFrameworkData(m)
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
				ifd, err = e2e.FrameworkDataOriginatingRecord(p)
				if err != nil {
					return fmt.Errorf("REDACTED", err)
				}
			default:
				return fmt.Errorf("REDACTED", infrstr)
			}

			simnet, err := e2e.FetchSimnet(record, ifd)
			if err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			cli.simnet = simnet
			switch infrstr {
			case "REDACTED":
				cli.infpnt = &dock.Supplier{
					SupplierData: platform.SupplierData{
						Simnet:            simnet,
						FrameworkData: ifd,
					},
				}
			case "REDACTED":
				cli.infpnt = &digitalregion.Supplier{
					SupplierData: platform.SupplierData{
						Simnet:            simnet,
						FrameworkData: ifd,
					},
				}
			default:
				return fmt.Errorf("REDACTED", infrstr)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, arguments []string) error {
			if err := Sanitize(cli.simnet); err != nil {
				return err
			}
			if err := Configure(cli.simnet, cli.infpnt); err != nil {
				return err
			}

			r := rand.New(rand.NewSource(unpredictableGerm)) //

			chnlFetchOutcome := make(chan error)
			ctx, fetchAbort := context.WithCancel(context.Background())
			defer fetchAbort()
			go func() {
				err := Fetch(ctx, cli.simnet)
				if err != nil {
					tracer.Failure(fmt.Sprintf("REDACTED", err.Error()))
				}
				chnlFetchOutcome <- err
			}()

			if err := Initiate(cmd.Context(), cli.simnet, cli.infpnt); err != nil {
				return err
			}

			if err := Pause(cmd.Context(), cli.simnet, 5); err != nil { //
				return err
			}

			if cli.simnet.OwnsDisruptions() {
				if err := Disrupt(cmd.Context(), cli.simnet); err != nil {
					return err
				}
				if err := Pause(cmd.Context(), cli.simnet, 5); err != nil { //
					return err
				}
			}

			if cli.simnet.Proof > 0 {
				if err := IntroduceProof(ctx, r, cli.simnet, cli.simnet.Proof); err != nil {
					return err
				}
				if err := Pause(cmd.Context(), cli.simnet, 5); err != nil { //
					return err
				}
			}

			fetchAbort()
			if err := <-chnlFetchOutcome; err != nil {
				return err
			}
			if err := Pause(cmd.Context(), cli.simnet, 5); err != nil { //
				return err
			}
			if err := Verify(cli.simnet, cli.infpnt.ObtainFrameworkData()); err != nil {
				return err
			}
			if !cli.maintain {
				if err := Sanitize(cli.simnet); err != nil {
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

	cli.origin.Flags().BoolVarP(&cli.maintain, "REDACTED", "REDACTED", false,
		"REDACTED")

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			return Configure(cli.simnet, cli.infpnt)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			_, err := os.Stat(cli.simnet.Dir)
			if os.IsNotExist(err) {
				err = Configure(cli.simnet, cli.infpnt)
			}
			if err != nil {
				return err
			}
			return Initiate(cmd.Context(), cli.simnet, cli.infpnt)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			return Disrupt(cmd.Context(), cli.simnet)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			return Pause(cmd.Context(), cli.simnet, 5)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			tracer.Details("REDACTED")
			return cli.infpnt.HaltSimnet(context.Background())
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) (err error) {
			return Fetch(context.Background(), cli.simnet)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Args:  cobra.MaximumNArgs(1),
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) (err error) {
			quantity := 1

			if len(arguments) == 1 {
				quantity, err = strconv.Atoi(arguments[0])
				if err != nil {
					return err
				}
			}

			return IntroduceProof(
				cmd.Context(),
				rand.New(rand.NewSource(unpredictableGerm)), //
				cli.simnet,
				quantity,
			)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			return Verify(cli.simnet, cli.infpnt.ObtainFrameworkData())
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			return Sanitize(cli.simnet)
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			return dock.InvokeArrangeDetailed(context.Background(), cli.simnet.Dir, "REDACTED")
		},
	})

	cli.origin.AddCommand(&cobra.Command{
		Use:   "REDACTED",
		Short: "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			return dock.InvokeArrangeDetailed(context.Background(), cli.simnet.Dir, "REDACTED", "REDACTED")
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
		RunE: func(cmd *cobra.Command, arguments []string) error {
			if err := Sanitize(cli.simnet); err != nil {
				return err
			}
			if err := Configure(cli.simnet, cli.infpnt); err != nil {
				return err
			}

			chnlFetchOutcome := make(chan error)
			ctx, fetchAbort := context.WithCancel(cmd.Context())
			defer fetchAbort()
			go func() {
				err := Fetch(ctx, cli.simnet)
				if err != nil {
					tracer.Failure(fmt.Sprintf("REDACTED", err.Error()))
				}
				chnlFetchOutcome <- err
			}()

			if err := Initiate(cmd.Context(), cli.simnet, cli.infpnt); err != nil {
				return err
			}

			if err := Pause(cmd.Context(), cli.simnet, 5); err != nil { //
				return err
			}

			//
			if err := Assessment(cmd.Context(), cli.simnet, 100); err != nil {
				return err
			}

			fetchAbort()
			if err := <-chnlFetchOutcome; err != nil {
				return err
			}

			return Sanitize(cli.simnet)
		},
	})

	return cli
}

//
func (cli *CLI) Run() {
	if err := cli.origin.Execute(); err != nil {
		tracer.Failure(err.Error())
		os.Exit(1)
	}
}
