package directives

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	nm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/peer"
)

var inaugurationDigest []byte

//
//
func AppendPeerSwitches(cmd *cobra.Command) {
	//
	cmd.Flags().String("REDACTED", settings.Pseudonym, "REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.PrivateAssessorOverhearLocation,
		"REDACTED")

	//
	cmd.Flags().BytesHexVar(
		&inaugurationDigest,
		"REDACTED",
		[]byte{},
		"REDACTED")
	cmd.Flags().Int64("REDACTED", settings.Agreement.DuplicateAttestInspectAltitude,
		"REDACTED"+
			"REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.DelegateApplication,
		"REDACTED"+
			"REDACTED")
	cmd.Flags().String("REDACTED", settings.Iface, "REDACTED")

	//
	cmd.Flags().String("REDACTED", settings.RPC.OverhearLocation, "REDACTED")
	cmd.Flags().String(
		"REDACTED",
		settings.RPC.GRPSOverhearLocation,
		"REDACTED")
	cmd.Flags().Bool("REDACTED", settings.RPC.Insecure, "REDACTED")
	cmd.Flags().String("REDACTED", settings.RPC.ProfilerOverhearLocation, "REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.P2P.OverhearLocation,
		"REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.OutsideLocation, "REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.Origins, "REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.EnduringNodes, "REDACTED")
	cmd.Flags().String("REDACTED",
		settings.P2P.AbsoluteNodeIDXDstore, "REDACTED")
	cmd.Flags().Bool("REDACTED", settings.P2P.PeerxHandler, "REDACTED")
	cmd.Flags().Bool("REDACTED", settings.P2P.OriginStyle, "REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.SecludedNodeIDXDstore, "REDACTED")

	//
	cmd.Flags().Bool(
		"REDACTED",
		settings.Agreement.GenerateVoidLedgers,
		"REDACTED")
	cmd.Flags().String(
		"REDACTED",
		settings.Agreement.GenerateVoidLedgersDuration.String(),
		"REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.DatastoreOrigin,
		"REDACTED")
	cmd.Flags().String(
		"REDACTED",
		settings.DatastoreRoute,
		"REDACTED")
}

//
//
func FreshExecutePeerDirective(peerSupplier nm.Supplier) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "REDACTED",
		Aliases: []string{"REDACTED", "REDACTED"},
		Short:   "REDACTED",
		RunE: func(cmd *cobra.Command, arguments []string) error {
			if err := inspectInaugurationDigest(settings); err != nil {
				return err
			}

			n, err := peerSupplier(settings, tracer)
			if err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			if err := n.Initiate(); err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			tracer.Details("REDACTED", "REDACTED", n.Router().PeerDetails())

			//
			strongos.EnsnareGesture(tracer, func() {
				if n.EqualsActive() {
					if err := n.Halt(); err != nil {
						tracer.Failure("REDACTED", "REDACTED", err)
					}
				}
			})

			//
			select {}
		},
	}

	AppendPeerSwitches(cmd)
	return cmd
}

func inspectInaugurationDigest(settings *cfg.Settings) error {
	if len(inaugurationDigest) == 0 || settings.Inauguration == "REDACTED" {
		return nil
	}

	//
	f, err := os.Open(settings.InaugurationRecord())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	preciseDigest := h.Sum(nil)

	//
	if !bytes.Equal(inaugurationDigest, preciseDigest) {
		return fmt.Errorf(
			"REDACTED",
			inaugurationDigest, settings.InaugurationRecord(), preciseDigest)
	}

	return nil
}
