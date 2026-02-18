package directives

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/settings"
	cometos "github.com/valkyrieworks/utils/os"
	nm "github.com/valkyrieworks/member"
)

var originDigest []byte

//
//
func AppendMemberOptions(cmd *cobra.Command) {
	//
	cmd.Flags().String("REDACTED", settings.Moniker, "REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.PrivateRatifierAcceptAddress,
		"REDACTED")

	//
	cmd.Flags().BytesHexVar(
		&originDigest,
		"REDACTED",
		[]byte{},
		"REDACTED")
	cmd.Flags().Int64("REDACTED", settings.Agreement.RepeatAttestInspectLevel,
		"REDACTED"+
			"REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.GatewayApplication,
		"REDACTED"+
			"REDACTED")
	cmd.Flags().String("REDACTED", settings.Iface, "REDACTED")

	//
	cmd.Flags().String("REDACTED", settings.RPC.AcceptLocation, "REDACTED")
	cmd.Flags().String(
		"REDACTED",
		settings.RPC.GRPCAcceptLocation,
		"REDACTED")
	cmd.Flags().Bool("REDACTED", settings.RPC.Risky, "REDACTED")
	cmd.Flags().String("REDACTED", settings.RPC.PprofAcceptLocation, "REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.P2P.AcceptLocation,
		"REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.OutsideLocation, "REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.Origins, "REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.DurableNodes, "REDACTED")
	cmd.Flags().String("REDACTED",
		settings.P2P.AbsoluteNodeIDXDatastore, "REDACTED")
	cmd.Flags().Bool("REDACTED", settings.P2P.PexHandler, "REDACTED")
	cmd.Flags().Bool("REDACTED", settings.P2P.OriginStyle, "REDACTED")
	cmd.Flags().String("REDACTED", settings.P2P.PrivateNodeIDXDatastore, "REDACTED")

	//
	cmd.Flags().Bool(
		"REDACTED",
		settings.Agreement.GenerateEmptyLedgers,
		"REDACTED")
	cmd.Flags().String(
		"REDACTED",
		settings.Agreement.GenerateEmptyLedgersCadence.String(),
		"REDACTED")

	//
	cmd.Flags().String(
		"REDACTED",
		settings.StoreOrigin,
		"REDACTED")
	cmd.Flags().String(
		"REDACTED",
		settings.StoreRoute,
		"REDACTED")
}

//
//
func NewExecuteMemberCommand(memberSource nm.Source) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "REDACTED",
		Aliases: []string{"REDACTED", "REDACTED"},
		Short:   "REDACTED",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := inspectOriginDigest(settings); err != nil {
				return err
			}

			n, err := memberSource(settings, tracer)
			if err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			if err := n.Begin(); err != nil {
				return fmt.Errorf("REDACTED", err)
			}

			tracer.Details("REDACTED", "REDACTED", n.Router().MemberDetails())

			//
			cometos.InterceptAlert(tracer, func() {
				if n.IsActive() {
					if err := n.Halt(); err != nil {
						tracer.Fault("REDACTED", "REDACTED", err)
					}
				}
			})

			//
			select {}
		},
	}

	AppendMemberOptions(cmd)
	return cmd
}

func inspectOriginDigest(settings *cfg.Settings) error {
	if len(originDigest) == 0 || settings.Origin == "REDACTED" {
		return nil
	}

	//
	f, err := os.Open(settings.OriginEntry())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	factualDigest := h.Sum(nil)

	//
	if !bytes.Equal(originDigest, factualDigest) {
		return fmt.Errorf(
			"REDACTED",
			originDigest, settings.OriginEntry(), factualDigest)
	}

	return nil
}
