package primary

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/netp2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	ApplicationLocatorTcpsocket  = "REDACTED"
	ApplicationLocatorPosix = "REDACTED"

	PrivatevalueLocatorTcpsocket     = "REDACTED"
	PrivatevalueLocatorPosix    = "REDACTED"
	PrivatevalueTokenRecord        = "REDACTED"
	PrivatevalueStatusRecord      = "REDACTED"
	PrivatevaluePlaceholderTokenRecord   = "REDACTED"
	PrivatevaluePlaceholderStatusRecord = "REDACTED"
)

//
func Configure(simnet *e2e.Simnet, infpnt platform.Supplier) error {
	tracer.Details("REDACTED", "REDACTED", simnet.Dir)

	if err := os.MkdirAll(simnet.Dir, os.ModePerm); err != nil {
		return err
	}

	if err := infpnt.Configure(); err != nil {
		return err
	}

	inauguration, err := CreateInauguration(simnet)
	if err != nil {
		return err
	}

	for _, peer := range simnet.Peers {
		peerPath := filepath.Join(simnet.Dir, peer.Alias)

		folders := []string{
			filepath.Join(peerPath, "REDACTED"),
			filepath.Join(peerPath, "REDACTED"),
			filepath.Join(peerPath, "REDACTED", "REDACTED"),
		}
		for _, dir := range folders {
			//
			if peer.Style == e2e.StyleAgile && strings.Contains(dir, "REDACTED") {
				continue
			}
			err := os.MkdirAll(dir, 0o755)
			if err != nil {
				return err
			}
		}

		cfg, err := CreateSettings(peer)
		if err != nil {
			return err
		}
		settings.PersistSettingsRecord(filepath.Join(peerPath, "REDACTED", "REDACTED"), cfg) //

		applicationConfig, err := CreateApplicationSettings(peer)
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(peerPath, "REDACTED", "REDACTED"), applicationConfig, 0o644) //
		if err != nil {
			return err
		}

		if peer.Style == e2e.StyleAgile {
			//
			continue
		}

		err = inauguration.PersistLike(filepath.Join(peerPath, "REDACTED", "REDACTED"))
		if err != nil {
			return err
		}

		err = (&p2p.PeerToken{PrivateToken: peer.PeerToken}).PersistLike(filepath.Join(peerPath, "REDACTED", "REDACTED"))
		if err != nil {
			return err
		}

		(privatevalue.FreshRecordPRV(peer.PrivatevalueToken,
			filepath.Join(peerPath, PrivatevalueTokenRecord),
			filepath.Join(peerPath, PrivatevalueStatusRecord),
		)).Persist()

		//
		//
		(privatevalue.FreshRecordPRV(edwards25519.ProducePrivateToken(),
			filepath.Join(peerPath, PrivatevaluePlaceholderTokenRecord),
			filepath.Join(peerPath, PrivatevaluePlaceholderStatusRecord),
		)).Persist()
	}

	if simnet.Titan {
		if err := simnet.PersistTitanSettings(); err != nil {
			return err
		}
	}

	return nil
}

//
func CreateInauguration(simnet *e2e.Simnet) (kinds.OriginPaper, error) {
	inauguration := kinds.OriginPaper{
		OriginMoment:     time.Now(),
		SuccessionUUID:         simnet.Alias,
		AgreementSettings: kinds.FallbackAgreementSettings(),
		PrimaryAltitude:   simnet.PrimaryAltitude,
	}
	//
	inauguration.AgreementSettings.Edition.App = 1
	inauguration.AgreementSettings.Proof.MaximumLifespanCountLedgers = e2e.ProofLifespanAltitude
	inauguration.AgreementSettings.Proof.MaximumLifespanInterval = e2e.ProofLifespanMoment
	inauguration.AgreementSettings.Assessor.PublicTokenKinds = []string{simnet.TokenKind}
	if simnet.LedgerMaximumOctets != 0 {
		inauguration.AgreementSettings.Ledger.MaximumOctets = simnet.LedgerMaximumOctets
	}
	if simnet.BallotAdditionsReviseAltitude == -1 {
		inauguration.AgreementSettings.Iface.BallotAdditionsActivateAltitude = simnet.BallotAdditionsActivateAltitude
	}
	for assessor, potency := range simnet.Assessors {
		inauguration.Assessors = append(inauguration.Assessors, kinds.OriginAssessor{
			Alias:    assessor.Alias,
			Location: assessor.PrivatevalueToken.PublicToken().Location(),
			PublicToken:  assessor.PrivatevalueToken.PublicToken(),
			Potency:   potency,
		})
	}
	//
	//
	sort.Slice(inauguration.Assessors, func(i, j int) bool {
		return strings.Compare(inauguration.Assessors[i].Alias, inauguration.Assessors[j].Alias) == -1
	})
	if len(simnet.PrimaryStatus) > 0 {
		applicationStatus, err := json.Marshal(simnet.PrimaryStatus)
		if err != nil {
			return inauguration, err
		}
		inauguration.ApplicationStatus = applicationStatus
	}
	return inauguration, inauguration.CertifyAlsoFinish()
}

//
func CreateSettings(peer *e2e.Peer) (*settings.Settings, error) {
	cfg := settings.FallbackSettings()
	cfg.Pseudonym = peer.Alias
	cfg.DelegateApplication = ApplicationLocatorTcpsocket
	cfg.RPC.OverhearLocation = "REDACTED"
	cfg.RPC.ProfilerOverhearLocation = "REDACTED"
	cfg.P2P.OutsideLocation = fmt.Sprintf("REDACTED", peer.LocatorPeer2peer(false))
	cfg.P2P.LocationRegisterPrecise = false
	cfg.DatastoreRepository = peer.Repository
	cfg.StatusChronize.ExplorationMoment = 5 * time.Second
	cfg.LedgerChronize.Edition = peer.LedgerChronizeEdition
	cfg.LedgerChronize.AggregateStyle = peer.LedgerChronizeAggregateStyle
	cfg.Txpool.ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes = int(peer.Simnet.ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes)
	cfg.Txpool.ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes = int(peer.Simnet.ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes)

	switch peer.TxpoolKind {
	case settings.TxpoolKindOverflow, settings.TxpoolKindApplication, settings.TxpoolKindNooperation:
		cfg.Txpool.Kind = peer.TxpoolKind
	case "REDACTED":
		cfg.Txpool.Kind = settings.TxpoolKindOverflow
	default:
		return nil, fmt.Errorf("REDACTED", peer.TxpoolKind)
	}

	switch peer.IfaceScheme {
	case e2e.SchemePosix:
		cfg.DelegateApplication = ApplicationLocatorPosix
	case e2e.SchemeTcpsocket:
		cfg.DelegateApplication = ApplicationLocatorTcpsocket
	case e2e.SchemeGRPS:
		cfg.DelegateApplication = ApplicationLocatorTcpsocket
		cfg.Iface = "REDACTED"
	case e2e.SchemeIntrinsic, e2e.SchemeIntrinsicLinkChronize:
		cfg.DelegateApplication = "REDACTED"
		cfg.Iface = "REDACTED"
	default:
		return nil, fmt.Errorf("REDACTED", peer.IfaceScheme)
	}

	//
	//
	//
	//
	cfg.PrivateAssessorOverhearLocation = "REDACTED"
	cfg.PrivateAssessorToken = PrivatevaluePlaceholderTokenRecord
	cfg.PrivateAssessorStatus = PrivatevaluePlaceholderStatusRecord

	switch peer.Style {
	case e2e.StyleAssessor:
		switch peer.PrivatevalueScheme {
		case e2e.SchemeRecord:
			cfg.PrivateAssessorToken = PrivatevalueTokenRecord
			cfg.PrivateAssessorStatus = PrivatevalueStatusRecord
		case e2e.SchemePosix:
			cfg.PrivateAssessorOverhearLocation = PrivatevalueLocatorPosix
		case e2e.SchemeTcpsocket:
			cfg.PrivateAssessorOverhearLocation = PrivatevalueLocatorTcpsocket
		default:
			return nil, fmt.Errorf("REDACTED", peer.PrivatevalueScheme)
		}
	case e2e.StyleGerm:
		cfg.P2P.OriginStyle = true
		cfg.P2P.PeerxHandler = true
	case e2e.StyleComplete, e2e.StyleAgile:
		//
	default:
		return nil, fmt.Errorf("REDACTED", peer.Style)
	}

	if peer.StatusChronize {
		cfg.StatusChronize.Activate = true
		cfg.StatusChronize.RemoteHosts = []string{}
		for _, node := range peer.Simnet.RepositoryPeers() {
			if node.Alias == peer.Alias {
				continue
			}
			cfg.StatusChronize.RemoteHosts = append(cfg.StatusChronize.RemoteHosts, node.LocatorRemote())
		}
		if len(cfg.StatusChronize.RemoteHosts) < 2 {
			return nil, errors.New("REDACTED")
		}
	}

	cfg.P2P.Origins = "REDACTED"
	for _, germ := range peer.Origins {
		if len(cfg.P2P.Origins) > 0 {
			cfg.P2P.Origins += "REDACTED"
		}
		cfg.P2P.Origins += germ.LocatorPeer2peer(true)
	}

	cfg.P2P.EnduringNodes = "REDACTED"
	for _, node := range peer.EnduringNodes {
		if len(cfg.P2P.EnduringNodes) > 0 {
			cfg.P2P.EnduringNodes += "REDACTED"
		}
		cfg.P2P.EnduringNodes += node.LocatorPeer2peer(true)
	}

	if peer.UtilizeLibpeer2peer {
		tracer.Details("REDACTED", "REDACTED", peer.Alias)

		//
		cfg.P2P.PeerxHandler = false
		cfg.P2P.LibraryPeer2peerSettings.Activated = true

		initiateNodes, err := CreateLibpeer2peerLocatorRegister(peer)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		cfg.P2P.LibraryPeer2peerSettings.InitiateNodes = initiateNodes
	}

	if peer.Simnet.RecordStratum != "REDACTED" {
		cfg.RecordStratum = peer.Simnet.RecordStratum
	}

	if peer.Simnet.RecordLayout != "REDACTED" {
		cfg.RecordLayout = peer.Simnet.RecordLayout
	}

	if peer.Titan {
		cfg.Telemetry.Titan = true
	}

	return cfg, nil
}

//
func CreateApplicationSettings(peer *e2e.Peer) ([]byte, error) {
	cfg := map[string]any{
		"REDACTED":                      peer.Simnet.Alias,
		"REDACTED":                           "REDACTED",
		"REDACTED":                        ApplicationLocatorPosix,
		"REDACTED":                          peer.Style,
		"REDACTED":                      "REDACTED",
		"REDACTED":              peer.EndureDuration,
		"REDACTED":             peer.ImageDuration,
		"REDACTED":                 peer.PreserveLedgers,
		"REDACTED":                      peer.PrivatevalueToken.Kind(),
		"REDACTED":        peer.Simnet.ArrangeNominationDeferral,
		"REDACTED":        peer.Simnet.HandleNominationDeferral,
		"REDACTED":                peer.Simnet.InspectTransferDeferral,
		"REDACTED":          peer.Simnet.BallotAdditionDeferral,
		"REDACTED":          peer.Simnet.CulminateLedgerDeferral,
		"REDACTED": peer.Simnet.BallotAdditionsActivateAltitude,
		"REDACTED": peer.Simnet.BallotAdditionsReviseAltitude,
		"REDACTED":           peer.Simnet.BallotAdditionExtent,
		"REDACTED":              peer.TxpoolKind == settings.TxpoolKindApplication,
	}

	switch peer.IfaceScheme {
	case e2e.SchemePosix:
		cfg["REDACTED"] = ApplicationLocatorPosix
	case e2e.SchemeTcpsocket:
		cfg["REDACTED"] = ApplicationLocatorTcpsocket
	case e2e.SchemeGRPS:
		cfg["REDACTED"] = ApplicationLocatorTcpsocket
		cfg["REDACTED"] = "REDACTED"
	case e2e.SchemeIntrinsic, e2e.SchemeIntrinsicLinkChronize:
		delete(cfg, "REDACTED")
		cfg["REDACTED"] = string(peer.IfaceScheme)
	default:
		return nil, fmt.Errorf("REDACTED", peer.IfaceScheme)
	}
	if peer.Style == e2e.StyleAssessor {
		switch peer.PrivatevalueScheme {
		case e2e.SchemeRecord:
		case e2e.SchemeTcpsocket:
			cfg["REDACTED"] = PrivatevalueLocatorTcpsocket
			cfg["REDACTED"] = PrivatevalueTokenRecord
			cfg["REDACTED"] = PrivatevalueStatusRecord
		case e2e.SchemePosix:
			cfg["REDACTED"] = PrivatevalueLocatorPosix
			cfg["REDACTED"] = PrivatevalueTokenRecord
			cfg["REDACTED"] = PrivatevalueStatusRecord
		default:
			return nil, fmt.Errorf("REDACTED", peer.PrivatevalueScheme)
		}
	}

	if len(peer.Simnet.AssessorRevisions) > 0 {
		assessorRevisions := map[string]map[string]int64{}
		for altitude, assessors := range peer.Simnet.AssessorRevisions {
			reviseValues := map[string]int64{}
			for peer, potency := range assessors {
				reviseValues[base64.StdEncoding.EncodeToString(peer.PrivatevalueToken.PublicToken().Octets())] = potency
			}
			assessorRevisions[fmt.Sprintf("REDACTED", altitude)] = reviseValues
		}
		cfg["REDACTED"] = assessorRevisions
	}

	var buf bytes.Buffer
	err := toml.NewEncoder(&buf).Encode(cfg)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return buf.Bytes(), nil
}

//
func CreateLibpeer2peerLocatorRegister(peer *e2e.Peer) ([]settings.LibraryPeer2peerInitiateNode, error) {
	var (
		nodes = []settings.LibraryPeer2peerInitiateNode{}
		stash = make(map[string]struct{})
	)

	for _, node := range append(peer.Origins, peer.EnduringNodes...) {
		//
		if _, ok := stash[node.Alias]; ok {
			continue
		}

		nodeUUID, err := netp2p.UUIDOriginatingSecludedToken(node.PeerToken)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", node.Alias, err)
		}

		//
		//
		const (
			localmachine = "REDACTED"
			strongChannel = 26656
		)

		//
		ip := node.OutsideINET.String()
		if ip == localmachine && peer.IntrinsicINET.String() != localmachine {
			ip = node.IntrinsicINET.String()
		}

		nodes = append(nodes, settings.LibraryPeer2peerInitiateNode{
			Machine:       fmt.Sprintf("REDACTED", ip, strongChannel),
			ID:         nodeUUID.String(),
			Enduring: equalsEnduring(peer, node),
		})

		stash[node.Alias] = struct{}{}
	}

	return nodes, nil
}

//
func ReviseSettingsStatusChronize(peer *e2e.Peer, altitude int64, digest []byte) error {
	configRoute := filepath.Join(peer.Simnet.Dir, peer.Alias, "REDACTED", "REDACTED")

	//
	//
	bz, err := os.ReadFile(configRoute)
	if err != nil {
		return err
	}
	bz = regexp.MustCompile("REDACTED").ReplaceAll(bz, []byte(fmt.Sprintf("REDACTED", altitude)))
	bz = regexp.MustCompile("REDACTED").ReplaceAll(bz, []byte(fmt.Sprintf("REDACTED", digest)))
	return os.WriteFile(configRoute, bz, 0o644) //
}

func equalsEnduring(machine, node *e2e.Peer) bool {
	for _, pp := range machine.EnduringNodes {
		if pp.Alias == node.Alias {
			return true
		}
	}

	return false

}
