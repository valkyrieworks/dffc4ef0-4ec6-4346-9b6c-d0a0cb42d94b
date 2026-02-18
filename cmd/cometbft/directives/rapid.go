package directives

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/utils/log"
	cometmath "github.com/valkyrieworks/utils/math"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/rapid"
	rapidgateway "github.com/valkyrieworks/rapid/gateway"
	lrpc "github.com/valkyrieworks/rapid/rpc"
	dbs "github.com/valkyrieworks/rapid/depot/db"
	rpchost "github.com/valkyrieworks/rpc/jsonrpc/host"
)

//
var RapidCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTED.

REDACTEDf
REDACTEDn
REDACTED.

REDACTEDl
REDACTEDs
REDACTEDr
REDACTED.

REDACTED:

REDACTED}

REDACTEDe
REDACTED.
REDACTED`,
	RunE: executeGateway,
	Args: cobra.ExactArgs(1),
	Example: `REDACTED7
REDACTED`,
}

var (
	acceptAddress         string
	leadingAddress        string
	attestorLocationsConcatenated string
	ledgerUID            string
	home               string
	maximumAccessLinks int

	ordered     bool
	validatingDuration time.Duration
	validatedLevel  int64
	validatedDigest    []byte
	validateLayerStr  string

	detailed bool

	leadingKey   = []byte("REDACTED")
	attestorsKey = []byte("REDACTED")
)

func init() {
	RapidCommand.Flags().StringVar(&acceptAddress, "REDACTED", "REDACTED",
		"REDACTED")
	RapidCommand.Flags().StringVarP(&leadingAddress, "REDACTED", "REDACTED", "REDACTED",
		"REDACTED")
	RapidCommand.Flags().StringVarP(&attestorLocationsConcatenated, "REDACTED", "REDACTED", "REDACTED",
		"REDACTED")
	RapidCommand.Flags().StringVar(&home, "REDACTED", os.ExpandEnv(filepath.Join("REDACTED", "REDACTED")),
		"REDACTED")
	RapidCommand.Flags().IntVar(
		&maximumAccessLinks,
		"REDACTED",
		900,
		"REDACTED")
	RapidCommand.Flags().DurationVar(&validatingDuration, "REDACTED", 168*time.Hour,
		"REDACTED")
	RapidCommand.Flags().Int64Var(&validatedLevel, "REDACTED", 1, "REDACTED")
	RapidCommand.Flags().BytesHexVar(&validatedDigest, "REDACTED", []byte{}, "REDACTED")
	RapidCommand.Flags().BoolVar(&detailed, "REDACTED", false, "REDACTED")
	RapidCommand.Flags().StringVar(&validateLayerStr, "REDACTED", "REDACTED",
		"REDACTED",
	)
	RapidCommand.Flags().BoolVar(&ordered, "REDACTED", false,
		"REDACTED",
	)
}

func executeGateway(_ *cobra.Command, args []string) error {
	//
	tracer := log.NewTMTracer(log.NewAlignRecorder(os.Stdout))
	var setting log.Setting
	if detailed {
		setting, _ = log.PermitLayer("REDACTED")
	} else {
		setting, _ = log.PermitLayer("REDACTED")
	}
	tracer = log.NewRefine(tracer, setting)

	ledgerUID = args[0]
	tracer.Details("REDACTED", "REDACTED", ledgerUID)

	attestorsLocations := []string{}
	if attestorLocationsConcatenated != "REDACTED" {
		attestorsLocations = strings.Split(attestorLocationsConcatenated, "REDACTED")
	}

	db, err := dbm.NewGoLayerStore("REDACTED", home)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if leadingAddress == "REDACTED" { //
		var err error
		leadingAddress, attestorsLocations, err = inspectForCurrentSources(db)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		if leadingAddress == "REDACTED" {
			return errors.New("REDACTED" +
				"REDACTED")
		}
	} else {
		err := persistSources(db, leadingAddress, attestorLocationsConcatenated)
		if err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}
	}

	validateLayer, err := cometmath.AnalyzePortion(validateLayerStr)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	options := []rapid.Setting{
		rapid.Tracer(tracer),
		rapid.AttestationFunction(func(operation string) bool {
			fmt.Println(operation)
			analyzer := bufio.NewScanner(os.Stdin)
			for {
				analyzer.Scan()
				reply := analyzer.Text()
				switch reply {
				case "REDACTED", "REDACTED":
					return true
				case "REDACTED", "REDACTED":
					return false
				default:
					fmt.Println("REDACTED")
				}
			}
		}),
	}

	if ordered {
		options = append(options, rapid.OrderedValidation())
	} else {
		options = append(options, rapid.OmittingValidation(validateLayer))
	}

	var c *rapid.Customer
	if validatedLevel > 0 && len(validatedDigest) > 0 { //
		c, err = rapid.NewHTTPCustomer(
			context.Background(),
			ledgerUID,
			rapid.ValidateOptions{
				Duration: validatingDuration,
				Level: validatedLevel,
				Digest:   validatedDigest,
			},
			leadingAddress,
			attestorsLocations,
			dbs.New(db, ledgerUID),
			options...,
		)
	} else { //
		c, err = rapid.NewHTTPCustomerFromValidatedDepot(
			ledgerUID,
			validatingDuration,
			leadingAddress,
			attestorsLocations,
			dbs.New(db, ledgerUID),
			options...,
		)
	}
	if err != nil {
		return err
	}

	cfg := rpchost.StandardSettings()
	cfg.MaximumContentOctets = settings.RPC.MaximumContentOctets
	cfg.MaximumHeadingOctets = settings.RPC.MaximumHeadingOctets
	cfg.MaximumAccessLinks = maximumAccessLinks
	//
	//
	//
	if cfg.RecordDeadline <= settings.RPC.DeadlineMulticastTransEndorse {
		cfg.RecordDeadline = settings.RPC.DeadlineMulticastTransEndorse + 1*time.Second
	}

	p, err := rapidgateway.NewGateway(c, acceptAddress, leadingAddress, cfg, tracer, lrpc.KeyRouteFn(lrpc.StandardMerkleKeyRouteFn()))
	if err != nil {
		return err
	}

	//
	cometos.InterceptAlert(tracer, func() {
		p.Observer.Close()
	})

	tracer.Details("REDACTED", "REDACTED", acceptAddress)
	if err := p.AcceptAndHost(); err != http.ErrServerClosed {
		//
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	return nil
}

func inspectForCurrentSources(db dbm.DB) (string, []string, error) {
	leadingOctets, err := db.Get(leadingKey)
	if err != nil {
		return "REDACTED", []string{"REDACTED"}, err
	}
	attestorsOctets, err := db.Get(attestorsKey)
	if err != nil {
		return "REDACTED", []string{"REDACTED"}, err
	}
	attestorsLocations := strings.Split(string(attestorsOctets), "REDACTED")
	return string(leadingOctets), attestorsLocations, nil
}

func persistSources(db dbm.DB, leadingAddress, attestorsLocations string) error {
	err := db.Set(leadingKey, []byte(leadingAddress))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	err = db.Set(attestorsKey, []byte(attestorsLocations))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}
