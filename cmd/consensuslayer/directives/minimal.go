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
	ctalgebra "github.com/valkyrieworks/utils/algebra"
	ctsystem "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/minimal"
	clientproxy "github.com/valkyrieworks/minimal/gateway"
	crpc "github.com/valkyrieworks/minimal/rpc"
	dbs "github.com/valkyrieworks/minimal/depot/db"
	endpointserver "github.com/valkyrieworks/rpc/dataendpoint/engine"
)

//
var LightCmd = &cobra.Command{
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
	RunE: runProxy,
	Args: cobra.ExactArgs(1),
	Example: `REDACTED7
REDACTED`,
}

var (
	listenAddr         string
	primaryAddr        string
	witnessAddrsJoined string
	chainID            string
	home               string
	maxOpenConnections int

	sequential     bool
	trustingPeriod time.Duration
	trustedHeight  int64
	trustedHash    []byte
	trustLevelStr  string

	verbose bool

	primaryKey   = []byte("REDACTED")
	witnessesKey = []byte("REDACTED")
)

func init() {
	LightCmd.Flags().StringVar(&listenAddr, "REDACTED", "REDACTED",
		"REDACTED")
	LightCmd.Flags().StringVarP(&primaryAddr, "REDACTED", "REDACTED", "REDACTED",
		"REDACTED")
	LightCmd.Flags().StringVarP(&witnessAddrsJoined, "REDACTED", "REDACTED", "REDACTED",
		"REDACTED")
	LightCmd.Flags().StringVar(&home, "REDACTED", os.ExpandEnv(filepath.Join("REDACTED", "REDACTED")),
		"REDACTED")
	LightCmd.Flags().IntVar(
		&maxOpenConnections,
		"REDACTED",
		900,
		"REDACTED")
	LightCmd.Flags().DurationVar(&trustingPeriod, "REDACTED", 168*time.Hour,
		"REDACTED")
	LightCmd.Flags().Int64Var(&trustedHeight, "REDACTED", 1, "REDACTED")
	LightCmd.Flags().BytesHexVar(&trustedHash, "REDACTED", []byte{}, "REDACTED")
	LightCmd.Flags().BoolVar(&verbose, "REDACTED", false, "REDACTED")
	LightCmd.Flags().StringVar(&trustLevelStr, "REDACTED", "REDACTED",
		"REDACTED",
	)
	LightCmd.Flags().BoolVar(&sequential, "REDACTED", false,
		"REDACTED",
	)
}

func runProxy(_ *cobra.Command, args []string) error {
	//
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	var option log.Option
	if verbose {
		option, _ = log.AllowLevel("REDACTED")
	} else {
		option, _ = log.AllowLevel("REDACTED")
	}
	logger = log.NewFilter(logger, option)

	chainID = args[0]
	logger.Info("REDACTED", "REDACTED", chainID)

	witnessesAddrs := []string{}
	if witnessAddrsJoined != "REDACTED" {
		witnessesAddrs = strings.Split(witnessAddrsJoined, "REDACTED")
	}

	db, err := dbm.NewGoLevelDB("REDACTED", home)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if primaryAddr == "REDACTED" { //
		var err error
		primaryAddr, witnessesAddrs, err = checkForExistingProviders(db)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		if primaryAddr == "REDACTED" {
			return errors.New("REDACTED" +
				"REDACTED")
		}
	} else {
		err := saveProviders(db, primaryAddr, witnessAddrsJoined)
		if err != nil {
			logger.Error("REDACTED", "REDACTED", err)
		}
	}

	trustLevel, err := ctalgebra.ParseFraction(trustLevelStr)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	options := []minimal.Option{
		minimal.Logger(logger),
		minimal.ConfirmationFunction(func(action string) bool {
			fmt.Println(action)
			scanner := bufio.NewScanner(os.Stdin)
			for {
				scanner.Scan()
				response := scanner.Text()
				switch response {
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

	if sequential {
		options = append(options, minimal.SequentialVerification())
	} else {
		options = append(options, minimal.SkippingVerification(trustLevel))
	}

	var c *minimal.Client
	if trustedHeight > 0 && len(trustedHash) > 0 { //
		c, err = minimal.NewHTTPClient(
			context.Background(),
			chainID,
			minimal.TrustOptions{
				Period: trustingPeriod,
				Height: trustedHeight,
				Hash:   trustedHash,
			},
			primaryAddr,
			witnessesAddrs,
			dbs.New(db, chainID),
			options...,
		)
	} else { //
		c, err = minimal.NewHTTPClientFromTrustedStore(
			chainID,
			trustingPeriod,
			primaryAddr,
			witnessesAddrs,
			dbs.New(db, chainID),
			options...,
		)
	}
	if err != nil {
		return err
	}

	cfg := endpointserver.DefaultConfig()
	cfg.MaxBodyBytes = config.RPC.MaxBodyBytes
	cfg.MaxHeaderBytes = config.RPC.MaxHeaderBytes
	cfg.MaxOpenConnections = maxOpenConnections
	//
	//
	//
	if cfg.WriteTimeout <= config.RPC.TimeoutBroadcastTxCommit {
		cfg.WriteTimeout = config.RPC.TimeoutBroadcastTxCommit + 1*time.Second
	}

	p, err := clientproxy.NewProxy(c, listenAddr, primaryAddr, cfg, logger, crpc.KeyPathFn(crpc.DefaultMerkleKeyPathFn()))
	if err != nil {
		return err
	}

	//
	ctsystem.TrapSignal(logger, func() {
		p.Listener.Close()
	})

	logger.Info("REDACTED", "REDACTED", listenAddr)
	if err := p.ListenAndServe(); err != http.ErrServerClosed {
		//
		logger.Error("REDACTED", "REDACTED", err)
	}

	return nil
}

func checkForExistingProviders(db dbm.DB) (string, []string, error) {
	primaryBytes, err := db.Get(primaryKey)
	if err != nil {
		return "REDACTED", []string{"REDACTED"}, err
	}
	witnessesBytes, err := db.Get(witnessesKey)
	if err != nil {
		return "REDACTED", []string{"REDACTED"}, err
	}
	witnessesAddrs := strings.Split(string(witnessesBytes), "REDACTED")
	return string(primaryBytes), witnessesAddrs, nil
}

func saveProviders(db dbm.DB, primaryAddr, witnessesAddrs string) error {
	err := db.Set(primaryKey, []byte(primaryAddr))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	err = db.Set(witnessesKey, []byte(witnessesAddrs))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}
