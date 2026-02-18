package directives

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/octets"
	ctrng "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/authkey"
	"github.com/valkyrieworks/kinds"
	cttime "github.com/valkyrieworks/kinds/moment"
)

var (
	nValidators    int
	nNonValidators int
	initialHeight  int64
	configFile     string
	outputDir      string
	nodeDirPrefix  string

	populatePersistentPeers bool
	hostnamePrefix          string
	hostnameSuffix          string
	startingIPAddress       string
	hostnames               []string
	p2pPort                 int
	randomMonikers          bool
)

const (
	nodeDirPerm = 0o755
)

func init() {
	TestnetFilesCmd.Flags().IntVar(&nValidators, "REDACTED", 4,
		"REDACTED")
	TestnetFilesCmd.Flags().StringVar(&configFile, "REDACTED", "REDACTED",
		"REDACTED")
	TestnetFilesCmd.Flags().IntVar(&nNonValidators, "REDACTED", 0,
		"REDACTED")
	TestnetFilesCmd.Flags().StringVar(&outputDir, "REDACTED", "REDACTED",
		"REDACTED")
	TestnetFilesCmd.Flags().StringVar(&nodeDirPrefix, "REDACTED", "REDACTED",
		"REDACTED")
	TestnetFilesCmd.Flags().Int64Var(&initialHeight, "REDACTED", 0,
		"REDACTED")

	TestnetFilesCmd.Flags().BoolVar(&populatePersistentPeers, "REDACTED", true,
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	TestnetFilesCmd.Flags().StringVar(&hostnamePrefix, "REDACTED", "REDACTED",
		"REDACTED")
	TestnetFilesCmd.Flags().StringVar(&hostnameSuffix, "REDACTED", "REDACTED",
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	TestnetFilesCmd.Flags().StringVar(&startingIPAddress, "REDACTED", "REDACTED",
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	TestnetFilesCmd.Flags().StringArrayVar(&hostnames, "REDACTED", []string{},
		"REDACTED")
	TestnetFilesCmd.Flags().IntVar(&p2pPort, "REDACTED", 26656,
		"REDACTED")
	TestnetFilesCmd.Flags().BoolVar(&randomMonikers, "REDACTED", false,
		"REDACTED")
}

//
var TestnetFilesCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDh
REDACTED.

REDACTED.

REDACTED.

REDACTED:

REDACTED2
REDACTED`,
	RunE: testnetFiles,
}

func testnetFiles(*cobra.Command, []string) error {
	if len(hostnames) > 0 && len(hostnames) != (nValidators+nNonValidators) {
		return fmt.Errorf(
			"REDACTED",
			nValidators+nNonValidators,
		)
	}

	config := cfg.DefaultConfig()

	//
	if configFile != "REDACTED" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		if err := viper.Unmarshal(config); err != nil {
			return err
		}
		if err := config.ValidateBasic(); err != nil {
			return err
		}
	}

	genVals := make([]kinds.GenesisValidator, nValidators)

	for i := 0; i < nValidators; i++ {
		nodeDirName := fmt.Sprintf("REDACTED", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName)
		config.SetRoot(nodeDir)

		err := os.MkdirAll(filepath.Join(nodeDir, "REDACTED"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}
		err = os.MkdirAll(filepath.Join(nodeDir, "REDACTED"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		if err := initFilesWithConfig(config); err != nil {
			return err
		}

		pvKeyFile := filepath.Join(nodeDir, config.PrivValidatorKey)
		pvStateFile := filepath.Join(nodeDir, config.PrivValidatorState)
		pv := authkey.LoadFilePV(pvKeyFile, pvStateFile)

		pubKey, err := pv.GetPubKey()
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		genVals[i] = kinds.GenesisValidator{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   1,
			Name:    nodeDirName,
		}
	}

	for i := 0; i < nNonValidators; i++ {
		nodeDir := filepath.Join(outputDir, fmt.Sprintf("REDACTED", nodeDirPrefix, i+nValidators))
		config.SetRoot(nodeDir)

		err := os.MkdirAll(filepath.Join(nodeDir, "REDACTED"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		err = os.MkdirAll(filepath.Join(nodeDir, "REDACTED"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		if err := initFilesWithConfig(config); err != nil {
			return err
		}
	}

	//
	genDoc := &kinds.GenesisDoc{
		ChainID:         "REDACTED" + ctrng.Str(6),
		ConsensusParams: kinds.DefaultConsensusParams(),
		GenesisTime:     cttime.Now(),
		InitialHeight:   initialHeight,
		Validators:      genVals,
	}

	//
	for i := 0; i < nValidators+nNonValidators; i++ {
		nodeDir := filepath.Join(outputDir, fmt.Sprintf("REDACTED", nodeDirPrefix, i))
		if err := genDoc.SaveAs(filepath.Join(nodeDir, config.Genesis)); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}
	}

	//
	var (
		persistentPeers string
		err             error
	)
	if populatePersistentPeers {
		persistentPeers, err = persistentPeersString(config)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}
	}

	//
	for i := 0; i < nValidators+nNonValidators; i++ {
		nodeDir := filepath.Join(outputDir, fmt.Sprintf("REDACTED", nodeDirPrefix, i))
		config.SetRoot(nodeDir)
		config.P2P.AddrBookStrict = false
		config.P2P.AllowDuplicateIP = true
		if populatePersistentPeers {
			config.P2P.PersistentPeers = persistentPeers
		}
		config.Moniker = moniker(i)

		cfg.WriteConfigFile(filepath.Join(nodeDir, "REDACTED", "REDACTED"), config)
	}

	fmt.Printf("REDACTED", nValidators+nNonValidators)
	return nil
}

func hostnameOrIP(i int) string {
	if len(hostnames) > 0 && i < len(hostnames) {
		return hostnames[i]
	}
	if startingIPAddress == "REDACTED" {
		return fmt.Sprintf("REDACTED", hostnamePrefix, i, hostnameSuffix)
	}
	ip := net.ParseIP(startingIPAddress)
	ip = ip.To4()
	if ip == nil {
		fmt.Printf("REDACTED", startingIPAddress)
		os.Exit(1)
	}

	for range i {
		incrementIP(ip)
	}
	return ip.String()
}

//
func incrementIP(ip net.IP) {
	//
	for idx := len(ip) - 1; idx >= 0; idx-- {
		ip[idx]++

		if ip[idx] != 0 {
			break
		}
	}
}

func persistentPeersString(config *cfg.Config) (string, error) {
	persistentPeers := make([]string, nValidators+nNonValidators)
	for i := 0; i < nValidators+nNonValidators; i++ {
		nodeDir := filepath.Join(outputDir, fmt.Sprintf("REDACTED", nodeDirPrefix, i))
		config.SetRoot(nodeDir)
		nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
		if err != nil {
			return "REDACTED", err
		}
		persistentPeers[i] = p2p.IDAddressString(nodeKey.ID(), fmt.Sprintf("REDACTED", hostnameOrIP(i), p2pPort))
	}
	return strings.Join(persistentPeers, "REDACTED"), nil
}

func moniker(i int) string {
	if randomMonikers {
		return randomMoniker()
	}
	if len(hostnames) > 0 && i < len(hostnames) {
		return hostnames[i]
	}
	if startingIPAddress == "REDACTED" {
		return fmt.Sprintf("REDACTED", hostnamePrefix, i, hostnameSuffix)
	}
	return randomMoniker()
}

func randomMoniker() string {
	return octets.HexBytes(ctrng.Bytes(8)).String()
}
