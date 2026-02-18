package directives

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/settings"
	ctsystem "github.com/valkyrieworks/utils/os"
	nm "github.com/valkyrieworks/instance"
)

var genesisHash []byte

//
//
func AddNodeFlags(cmd *cobra.Command) {
	//
	cmd.Flags().String("REDACTED", config.Moniker, "REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		config.PrivValidatorListenAddr,
		"REDACTED")

	//
	cmd.Flags().BytesHexVar(
		&genesisHash,
		"REDACTED",
		[]byte{},
		"REDACTED")
	cmd.Flags().Int64("REDACTED", config.Consensus.DoubleSignCheckHeight,
		"REDACTED"+
			"REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		config.ProxyApp,
		"REDACTED"+
			"REDACTED")
	cmd.Flags().String("REDACTED", config.ABCI, "REDACTED")

	//
	cmd.Flags().String("REDACTED", config.RPC.ListenAddress, "REDACTED")
	cmd.Flags().String(
		"REDACTED",
		config.RPC.GRPCListenAddress,
		"REDACTED")
	cmd.Flags().Bool("REDACTED", config.RPC.Unsafe, "REDACTED")
	cmd.Flags().String("REDACTED", config.RPC.PprofListenAddress, "REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		config.P2P.ListenAddress,
		"REDACTED")
	cmd.Flags().String("REDACTED", config.P2P.ExternalAddress, "REDACTED")
	cmd.Flags().String("REDACTED", config.P2P.Seeds, "REDACTED")
	cmd.Flags().String("REDACTED", config.P2P.PersistentPeers, "REDACTED")
	cmd.Flags().String("REDACTED",
		config.P2P.UnconditionalPeerIDs, "REDACTED")
	cmd.Flags().Bool("REDACTED", config.P2P.PexReactor, "REDACTED")
	cmd.Flags().Bool("REDACTED", config.P2P.SeedMode, "REDACTED")
	cmd.Flags().String("REDACTED", config.P2P.PrivatePeerIDs, "REDACTED")

	//
	cmd.Flags().Bool(
		"REDACTED",
		config.Consensus.CreateEmptyBlocks,
		"REDACTED")
	cmd.Flags().String(
		"REDACTED",
		config.Consensus.CreateEmptyBlocksInterval.String(),
		"REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		config.DBBackend,
		"REDACTED")
	cmd.Flags().String(
		"REDACTED",
		config.DBPath,
		"REDACTED")
}

//
//
func NewRunNodeCmd(nodeProvider nm.Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "REDACTED",
		Aliases: []string{"REDACTED", "REDACTED"},
		Short:   "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkGenesisHash(config); err != nil {
				return err
			}

			n, err := nodeProvider(config, logger)
			if err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			if err := n.Start(); err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			logger.Info("REDACTED", "REDACTED", n.Switch().NodeInfo())

			//
			ctsystem.TrapSignal(logger, func() {
				if n.IsRunning() {
					if err := n.Stop(); err != nil {
						logger.Error("REDACTED", "REDACTED", err)
					}
				}
			})

			//
			select {}
		},
	}

	AddNodeFlags(cmd)
	return cmd
}

func checkGenesisHash(config *cfg.Config) error {
	if len(genesisHash) == 0 || config.Genesis == "REDACTED" {
		return nil
	}

	//
	f, err := os.Open(config.GenesisFile())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	actualHash := h.Sum(nil)

	//
	if !bytes.Equal(genesisHash, actualHash) {
		return fmt.Errorf(
			"REDACTED",
			genesisHash, config.GenesisFile(), actualHash)
	}

	return nil
}
