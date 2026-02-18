package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/netpeer"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/privatekey"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/e2e/pkg/platform"
	"github.com/valkyrieworks/kinds"
)

const (
	ApplicationLocationTCP  = "REDACTED"
	ApplicationLocationUNIX = "REDACTED"

	PrivatekeyLocationTCP     = "REDACTED"
	PrivatekeyLocationUNIX    = "REDACTED"
	PrivatekeyKeyEntry        = "REDACTED"
	PrivatekeyStatusEntry      = "REDACTED"
	PrivatekeyMockKeyEntry   = "REDACTED"
	PrivatekeyMockStatusEntry = "REDACTED"
)

//
func Configure(verifychain *e2e.Verifychain, infp platform.Source) error {
	tracer.Details("REDACTED", "REDACTED", verifychain.Dir)

	if err := os.MkdirAll(verifychain.Dir, os.ModePerm); err != nil {
		return err
	}

	if err := infp.Configure(); err != nil {
		return err
	}

	origin, err := CreateOrigin(verifychain)
	if err != nil {
		return err
	}

	for _, member := range verifychain.Instances {
		memberFolder := filepath.Join(verifychain.Dir, member.Label)

		folders := []string{
			filepath.Join(memberFolder, "REDACTED"),
			filepath.Join(memberFolder, "REDACTED"),
			filepath.Join(memberFolder, "REDACTED", "REDACTED"),
		}
		for _, dir := range folders {
			//
			if member.Style == e2e.StyleRapid && strings.Contains(dir, "REDACTED") {
				continue
			}
			err := os.MkdirAll(dir, 0o755)
			if err != nil {
				return err
			}
		}

		cfg, err := CreateSettings(member)
		if err != nil {
			return err
		}
		settings.RecordSettingsEntry(filepath.Join(memberFolder, "REDACTED", "REDACTED"), cfg) //

		applicationConfig, err := CreateApplicationSettings(member)
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(memberFolder, "REDACTED", "REDACTED"), applicationConfig, 0o644) //
		if err != nil {
			return err
		}

		if member.Style == e2e.StyleRapid {
			//
			continue
		}

		err = origin.PersistAs(filepath.Join(memberFolder, "REDACTED", "REDACTED"))
		if err != nil {
			return err
		}

		err = (&p2p.MemberKey{PrivateKey: member.MemberKey}).PersistAs(filepath.Join(memberFolder, "REDACTED", "REDACTED"))
		if err != nil {
			return err
		}

		(privatekey.NewEntryPV(member.PrivatekeyKey,
			filepath.Join(memberFolder, PrivatekeyKeyEntry),
			filepath.Join(memberFolder, PrivatekeyStatusEntry),
		)).Persist()

		//
		//
		(privatekey.NewEntryPV(ed25519.GeneratePrivateKey(),
			filepath.Join(memberFolder, PrivatekeyMockKeyEntry),
			filepath.Join(memberFolder, PrivatekeyMockStatusEntry),
		)).Persist()
	}

	if verifychain.Monitorstats {
		if err := verifychain.RecordMonitorstatsSettings(); err != nil {
			return err
		}
	}

	return nil
}

//
func CreateOrigin(verifychain *e2e.Verifychain) (kinds.OriginPaper, error) {
	origin := kinds.OriginPaper{
		OriginMoment:     time.Now(),
		LedgerUID:         verifychain.Label,
		AgreementOptions: kinds.StandardAgreementOptions(),
		PrimaryLevel:   verifychain.PrimaryLevel,
	}
	//
	origin.AgreementOptions.Release.App = 1
	origin.AgreementOptions.Proof.MaximumDurationCountLedgers = e2e.ProofEraLevel
	origin.AgreementOptions.Proof.MaximumDurationPeriod = e2e.ProofEraTime
	origin.AgreementOptions.Ratifier.PublicKeyKinds = []string{verifychain.KeyKind}
	if verifychain.LedgerMaximumOctets != 0 {
		origin.AgreementOptions.Ledger.MaximumOctets = verifychain.LedgerMaximumOctets
	}
	if verifychain.BallotPluginsModifyLevel == -1 {
		origin.AgreementOptions.Iface.BallotPluginsActivateLevel = verifychain.BallotPluginsActivateLevel
	}
	for ratifier, energy := range verifychain.Ratifiers {
		origin.Ratifiers = append(origin.Ratifiers, kinds.OriginRatifier{
			Label:    ratifier.Label,
			Location: ratifier.PrivatekeyKey.PublicKey().Location(),
			PublicKey:  ratifier.PrivatekeyKey.PublicKey(),
			Energy:   energy,
		})
	}
	//
	//
	sort.Slice(origin.Ratifiers, func(i, j int) bool {
		return strings.Compare(origin.Ratifiers[i].Label, origin.Ratifiers[j].Label) == -1
	})
	if len(verifychain.PrimaryStatus) > 0 {
		applicationStatus, err := json.Marshal(verifychain.PrimaryStatus)
		if err != nil {
			return origin, err
		}
		origin.ApplicationStatus = applicationStatus
	}
	return origin, origin.CertifyAndFinished()
}

//
func CreateSettings(member *e2e.Member) (*settings.Settings, error) {
	cfg := settings.StandardSettings()
	cfg.Moniker = member.Label
	cfg.GatewayApplication = ApplicationLocationTCP
	cfg.RPC.AcceptLocation = "REDACTED"
	cfg.RPC.PprofAcceptLocation = "REDACTED"
	cfg.P2P.OutsideLocation = fmt.Sprintf("REDACTED", member.LocationP2P(false))
	cfg.P2P.AddressLedgerPrecise = false
	cfg.StoreOrigin = member.Datastore
	cfg.StatusAlign.DetectionTime = 5 * time.Second
	cfg.LedgerAlign.Release = member.LedgerAlignRelease
	cfg.Txpool.ExploratoryMaximumGossipLinkagesToNotDurableNodes = int(member.Verifychain.ExploratoryMaximumGossipLinkagesToNotDurableNodes)
	cfg.Txpool.ExploratoryMaximumGossipLinkagesToDurableNodes = int(member.Verifychain.ExploratoryMaximumGossipLinkagesToDurableNodes)

	switch member.TxpoolKind {
	case settings.TxpoolKindOverflow, settings.TxpoolKindApplication, settings.TxpoolKindNoop:
		cfg.Txpool.Kind = member.TxpoolKind
	case "REDACTED":
		cfg.Txpool.Kind = settings.TxpoolKindOverflow
	default:
		return nil, fmt.Errorf("REDACTED", member.TxpoolKind)
	}

	switch member.IfaceProtocol {
	case e2e.ProtocolUNIX:
		cfg.GatewayApplication = ApplicationLocationUNIX
	case e2e.ProtocolTCP:
		cfg.GatewayApplication = ApplicationLocationTCP
	case e2e.ProtocolGRPC:
		cfg.GatewayApplication = ApplicationLocationTCP
		cfg.Iface = "REDACTED"
	case e2e.ProtocolIntrinsic, e2e.ProtocolIntrinsicLinkAlign:
		cfg.GatewayApplication = "REDACTED"
		cfg.Iface = "REDACTED"
	default:
		return nil, fmt.Errorf("REDACTED", member.IfaceProtocol)
	}

	//
	//
	//
	//
	cfg.PrivateRatifierAcceptAddress = "REDACTED"
	cfg.PrivateRatifierKey = PrivatekeyMockKeyEntry
	cfg.PrivateRatifierStatus = PrivatekeyMockStatusEntry

	switch member.Style {
	case e2e.StyleRatifier:
		switch member.PrivatekeyProtocol {
		case e2e.ProtocolEntry:
			cfg.PrivateRatifierKey = PrivatekeyKeyEntry
			cfg.PrivateRatifierStatus = PrivatekeyStatusEntry
		case e2e.ProtocolUNIX:
			cfg.PrivateRatifierAcceptAddress = PrivatekeyLocationUNIX
		case e2e.ProtocolTCP:
			cfg.PrivateRatifierAcceptAddress = PrivatekeyLocationTCP
		default:
			return nil, fmt.Errorf("REDACTED", member.PrivatekeyProtocol)
		}
	case e2e.StyleOrigin:
		cfg.P2P.OriginStyle = true
		cfg.P2P.PexHandler = true
	case e2e.StyleComplete, e2e.StyleRapid:
		//
	default:
		return nil, fmt.Errorf("REDACTED", member.Style)
	}

	if member.StatusAlign {
		cfg.StatusAlign.Activate = true
		cfg.StatusAlign.RPCHosts = []string{}
		for _, node := range member.Verifychain.CatalogInstances() {
			if node.Label == member.Label {
				continue
			}
			cfg.StatusAlign.RPCHosts = append(cfg.StatusAlign.RPCHosts, node.LocationRPC())
		}
		if len(cfg.StatusAlign.RPCHosts) < 2 {
			return nil, errors.New("REDACTED")
		}
	}

	cfg.P2P.Origins = "REDACTED"
	for _, origin := range member.Origins {
		if len(cfg.P2P.Origins) > 0 {
			cfg.P2P.Origins += "REDACTED"
		}
		cfg.P2P.Origins += origin.LocationP2P(true)
	}

	cfg.P2P.DurableNodes = "REDACTED"
	for _, node := range member.DurableNodes {
		if len(cfg.P2P.DurableNodes) > 0 {
			cfg.P2P.DurableNodes += "REDACTED"
		}
		cfg.P2P.DurableNodes += node.LocationP2P(true)
	}

	if member.EmployLibp2p {
		tracer.Details("REDACTED", "REDACTED", member.Label)

		//
		cfg.P2P.PexHandler = false
		cfg.P2P.LibraryP2PSettings.Activated = true

		onboardNodes, err := CreateLibp2pLocationLedger(member)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		cfg.P2P.LibraryP2PSettings.OnboardNodes = onboardNodes
	}

	if member.Verifychain.TraceLayer != "REDACTED" {
		cfg.TraceLayer = member.Verifychain.TraceLayer
	}

	if member.Verifychain.TraceLayout != "REDACTED" {
		cfg.TraceLayout = member.Verifychain.TraceLayout
	}

	if member.Monitorstats {
		cfg.Telemetry.Monitorstats = true
	}

	return cfg, nil
}

//
func CreateApplicationSettings(member *e2e.Member) ([]byte, error) {
	cfg := map[string]any{
		"REDACTED":                      member.Verifychain.Label,
		"REDACTED":                           "REDACTED",
		"REDACTED":                        ApplicationLocationUNIX,
		"REDACTED":                          member.Style,
		"REDACTED":                      "REDACTED",
		"REDACTED":              member.EndureCadence,
		"REDACTED":             member.MirrorCadence,
		"REDACTED":                 member.PreserveLedgers,
		"REDACTED":                      member.PrivatekeyKey.Kind(),
		"REDACTED":        member.Verifychain.ArrangeNominationDeferral,
		"REDACTED":        member.Verifychain.HandleNominationDeferral,
		"REDACTED":                member.Verifychain.InspectTransferDeferral,
		"REDACTED":          member.Verifychain.BallotAdditionDeferral,
		"REDACTED":          member.Verifychain.CompleteLedgerDeferral,
		"REDACTED": member.Verifychain.BallotPluginsActivateLevel,
		"REDACTED": member.Verifychain.BallotPluginsModifyLevel,
		"REDACTED":           member.Verifychain.BallotAdditionVolume,
		"REDACTED":              member.TxpoolKind == settings.TxpoolKindApplication,
	}

	switch member.IfaceProtocol {
	case e2e.ProtocolUNIX:
		cfg["REDACTED"] = ApplicationLocationUNIX
	case e2e.ProtocolTCP:
		cfg["REDACTED"] = ApplicationLocationTCP
	case e2e.ProtocolGRPC:
		cfg["REDACTED"] = ApplicationLocationTCP
		cfg["REDACTED"] = "REDACTED"
	case e2e.ProtocolIntrinsic, e2e.ProtocolIntrinsicLinkAlign:
		delete(cfg, "REDACTED")
		cfg["REDACTED"] = string(member.IfaceProtocol)
	default:
		return nil, fmt.Errorf("REDACTED", member.IfaceProtocol)
	}
	if member.Style == e2e.StyleRatifier {
		switch member.PrivatekeyProtocol {
		case e2e.ProtocolEntry:
		case e2e.ProtocolTCP:
			cfg["REDACTED"] = PrivatekeyLocationTCP
			cfg["REDACTED"] = PrivatekeyKeyEntry
			cfg["REDACTED"] = PrivatekeyStatusEntry
		case e2e.ProtocolUNIX:
			cfg["REDACTED"] = PrivatekeyLocationUNIX
			cfg["REDACTED"] = PrivatekeyKeyEntry
			cfg["REDACTED"] = PrivatekeyStatusEntry
		default:
			return nil, fmt.Errorf("REDACTED", member.PrivatekeyProtocol)
		}
	}

	if len(member.Verifychain.RatifierRefreshes) > 0 {
		ratifierRefreshes := map[string]map[string]int64{}
		for level, ratifiers := range member.Verifychain.RatifierRefreshes {
			modifyValues := map[string]int64{}
			for member, energy := range ratifiers {
				modifyValues[base64.StdEncoding.EncodeToString(member.PrivatekeyKey.PublicKey().Octets())] = energy
			}
			ratifierRefreshes[fmt.Sprintf("REDACTED", level)] = modifyValues
		}
		cfg["REDACTED"] = ratifierRefreshes
	}

	var buf bytes.Buffer
	err := toml.NewEncoder(&buf).Encode(cfg)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return buf.Bytes(), nil
}

//
func CreateLibp2pLocationLedger(member *e2e.Member) ([]settings.LibraryP2POnboardNode, error) {
	var (
		nodes = []settings.LibraryP2POnboardNode{}
		repository = make(map[string]struct{})
	)

	for _, node := range append(member.Origins, member.DurableNodes...) {
		//
		if _, ok := repository[node.Label]; ok {
			continue
		}

		nodeUID, err := netpeer.UIDFromPrivateKey(node.MemberKey)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", node.Label, err)
		}

		//
		//
		const (
			localhost = "REDACTED"
			cometPort = 26656
		)

		//
		ip := node.OutsideIP.String()
		if ip == localhost && member.IntrinsicIP.String() != localhost {
			ip = node.IntrinsicIP.String()
		}

		nodes = append(nodes, settings.LibraryP2POnboardNode{
			Machine:       fmt.Sprintf("REDACTED", ip, cometPort),
			ID:         nodeUID.String(),
			Durable: isDurable(member, node),
		})

		repository[node.Label] = struct{}{}
	}

	return nodes, nil
}

//
func ModifySettingsStatusAlign(member *e2e.Member, level int64, digest []byte) error {
	configRoute := filepath.Join(member.Verifychain.Dir, member.Label, "REDACTED", "REDACTED")

	//
	//
	bz, err := os.ReadFile(configRoute)
	if err != nil {
		return err
	}
	bz = regexp.MustCompile("REDACTED").ReplaceAll(bz, []byte(fmt.Sprintf("REDACTED", level)))
	bz = regexp.MustCompile("REDACTED").ReplaceAll(bz, []byte(fmt.Sprintf("REDACTED", digest)))
	return os.WriteFile(configRoute, bz, 0o644) //
}

func isDurable(machine, node *e2e.Member) bool {
	for _, pp := range machine.DurableNodes {
		if pp.Label == node.Label {
			return true
		}
	}

	return false

}
