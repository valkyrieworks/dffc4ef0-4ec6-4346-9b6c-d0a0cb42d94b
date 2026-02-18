package e2e

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/bls12381"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/secp256k1"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	"github.com/valkyrieworks/kinds"

	_ "embed"
)

const (
	arbitraryOrigin               int64  = 2308084734268
	gatewayPortInitial           uint32 = 5701
	monitorstatsGatewayPortInitial uint32 = 6701

	standardClusterVolume   = 2
	standardLinkages = 1
	standardTransferVolumeOctets = 1024

	nativeRelease = "REDACTED"
)

type (
	Style         string
	Protocol     string
	Variation string
)

const (
	StyleRatifier Style = "REDACTED"
	StyleComplete      Style = "REDACTED"
	StyleRapid     Style = "REDACTED"
	StyleOrigin      Style = "REDACTED"

	ProtocolIntrinsic         Protocol = "REDACTED"
	ProtocolIntrinsicLinkAlign Protocol = "REDACTED"
	ProtocolEntry            Protocol = "REDACTED"
	ProtocolGRPC            Protocol = "REDACTED"
	ProtocolTCP             Protocol = "REDACTED"
	ProtocolUNIX            Protocol = "REDACTED"

	VariationDetach Variation = "REDACTED"
	VariationTerminate       Variation = "REDACTED"
	VariationStall      Variation = "REDACTED"
	VariationReboot    Variation = "REDACTED"
	VariationEnhance    Variation = "REDACTED"

	ProofEraLevel int64         = 14
	ProofEraTime   time.Duration = 1500 * time.Millisecond
)

//
type Verifychain struct {
	Label                                                 string
	Entry                                                 string
	Dir                                                  string
	IP                                                   *net.IPNet
	PrimaryLevel                                        int64
	PrimaryStatus                                         map[string]string
	Ratifiers                                           map[*Member]int64
	RatifierRefreshes                                     map[int64]map[*Member]int64
	Instances                                                []*Member
	KeyKind                                              string
	Proof                                             int
	ImportTransferVolumeOctets                                      int
	ImportTransferClusterVolume                                      int
	ImportTransferLinkages                                    int
	ImportMaximumTrans                                           int
	IfaceProtocol                                         string
	ArrangeNominationDeferral                                 time.Duration
	HandleNominationDeferral                                 time.Duration
	InspectTransferDeferral                                         time.Duration
	BallotAdditionDeferral                                   time.Duration
	CompleteLedgerDeferral                                   time.Duration
	EnhanceRelease                                       string
	TraceLayer                                             string
	TraceLayout                                            string
	Monitorstats                                           bool
	LedgerMaximumOctets                                        int64
	BallotPluginsActivateLevel                           int64
	BallotPluginsModifyLevel                           int64
	BallotAdditionVolume                                    uint
	ExploratoryMaximumGossipLinkagesToDurableNodes    uint
	ExploratoryMaximumGossipLinkagesToNotDurableNodes uint
}

//
type Member struct {
	Label                string
	Release             string
	Verifychain             *Verifychain
	Style                Style
	PrivatekeyKey          vault.PrivateKey
	MemberKey             vault.PrivateKey
	IntrinsicIP          net.IP
	OutsideIP          net.IP
	GatewayPort           uint32
	BeginAt             int64
	LedgerAlignRelease    string
	StatusAlign           bool
	Datastore            string
	IfaceProtocol        Protocol
	PrivatekeyProtocol     Protocol
	EndureCadence     uint64
	MirrorCadence    uint64
	PreserveLedgers        uint64
	Origins               []*Member
	DurableNodes     []*Member
	Variations       []Variation
	TransmitNoImport          bool
	Monitorstats          bool
	EmployLibp2p           bool
	TxpoolKind         string
	MonitorstatsGatewayPort uint32
}

//
//
//
//
//
func ImportVerifychain(entry string, ifd PlatformData) (*Verifychain, error) {
	declaration, err := ImportDeclaration(entry)
	if err != nil {
		return nil, err
	}
	return NewVerifychainFromDeclaration(declaration, entry, ifd)
}

//
func NewVerifychainFromDeclaration(declaration Declaration, entry string, ifd PlatformData) (*Verifychain, error) {
	dir := strings.TrimSuffix(entry, filepath.Ext(entry))

	keyGenerate := newKeyProducer(arbitraryOrigin)
	monitorstatsGatewayPortGenerate := newPortProducer(monitorstatsGatewayPortInitial)
	_, ipNet, err := net.ParseCIDR(ifd.Fabric)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", ifd.Fabric, err)
	}

	verifychain := &Verifychain{
		Label:                       filepath.Base(dir),
		Entry:                       entry,
		Dir:                        dir,
		IP:                         ipNet,
		PrimaryLevel:              1,
		PrimaryStatus:               declaration.PrimaryStatus,
		Ratifiers:                 map[*Member]int64{},
		RatifierRefreshes:           map[int64]map[*Member]int64{},
		Instances:                      []*Member{},
		KeyKind:                    declaration.KeyKind,
		Proof:                   declaration.Proof,
		ImportTransferVolumeOctets:            declaration.ImportTransferVolumeOctets,
		ImportTransferClusterVolume:            declaration.ImportTransferClusterVolume,
		ImportTransferLinkages:          declaration.ImportTransferLinkages,
		ImportMaximumTrans:                 declaration.ImportMaximumTrans,
		IfaceProtocol:               declaration.IfaceProtocol,
		ArrangeNominationDeferral:       declaration.ArrangeNominationDeferral,
		HandleNominationDeferral:       declaration.HandleNominationDeferral,
		InspectTransferDeferral:               declaration.InspectTransferDeferral,
		BallotAdditionDeferral:         declaration.BallotAdditionDeferral,
		CompleteLedgerDeferral:         declaration.CompleteLedgerDeferral,
		EnhanceRelease:             declaration.EnhanceRelease,
		TraceLayer:                   declaration.TraceLayer,
		TraceLayout:                  declaration.TraceLayout,
		Monitorstats:                 declaration.Monitorstats,
		LedgerMaximumOctets:              declaration.LedgerMaximumOctets,
		BallotPluginsActivateLevel: declaration.BallotPluginsActivateLevel,
		BallotPluginsModifyLevel: declaration.BallotPluginsModifyLevel,
		BallotAdditionVolume:          declaration.BallotAdditionVolume,
		ExploratoryMaximumGossipLinkagesToDurableNodes:    declaration.ExploratoryMaximumGossipLinkagesToDurableNodes,
		ExploratoryMaximumGossipLinkagesToNotDurableNodes: declaration.ExploratoryMaximumGossipLinkagesToNotDurableNodes,
	}

	if len(declaration.KeyKind) != 0 {
		verifychain.KeyKind = declaration.KeyKind
	}
	if declaration.PrimaryLevel > 0 {
		verifychain.PrimaryLevel = declaration.PrimaryLevel
	}
	if verifychain.KeyKind == "REDACTED" {
		verifychain.KeyKind = ed25519.KeyKind
	}
	if verifychain.IfaceProtocol == "REDACTED" {
		verifychain.IfaceProtocol = string(ProtocolIntrinsic)
	}
	if verifychain.EnhanceRelease == "REDACTED" {
		verifychain.EnhanceRelease = nativeRelease
	}
	if verifychain.ImportTransferLinkages == 0 {
		verifychain.ImportTransferLinkages = standardLinkages
	}
	if verifychain.ImportTransferClusterVolume == 0 {
		verifychain.ImportTransferClusterVolume = standardClusterVolume
	}
	if verifychain.ImportTransferVolumeOctets == 0 {
		verifychain.ImportTransferVolumeOctets = standardTransferVolumeOctets
	}

	for _, label := range arrangeMemberLabels(declaration) {
		memberDeclaration := declaration.Instances[label]
		ind, ok := ifd.Occurrences[label]
		if !ok {
			return nil, fmt.Errorf("REDACTED", label)
		}
		extensionIP := ind.ExtensionIPLocation
		if len(extensionIP) == 0 {
			extensionIP = ind.IPLocation
		}
		v := memberDeclaration.Release
		if v == "REDACTED" {
			v = nativeRelease
		}

		member := &Member{
			Label:             label,
			Release:          v,
			Verifychain:          verifychain,
			PrivatekeyKey:       keyGenerate.Compose(declaration.KeyKind),
			MemberKey:          keyGenerate.Compose("REDACTED"),
			IntrinsicIP:       ind.IPLocation,
			OutsideIP:       extensionIP,
			GatewayPort:        ind.Port,
			Style:             StyleRatifier,
			Datastore:         "REDACTED",
			IfaceProtocol:     Protocol(verifychain.IfaceProtocol),
			PrivatekeyProtocol:  ProtocolEntry,
			BeginAt:          memberDeclaration.BeginAt,
			LedgerAlignRelease: memberDeclaration.LedgerAlignRelease,
			StatusAlign:        memberDeclaration.StatusAlign,
			EndureCadence:  1,
			MirrorCadence: memberDeclaration.MirrorCadence,
			PreserveLedgers:     memberDeclaration.PreserveLedgers,
			Variations:    []Variation{},
			TransmitNoImport:       memberDeclaration.TransmitNoImport,
			EmployLibp2p:        memberDeclaration.EmployLibp2p,
			TxpoolKind:      memberDeclaration.TxpoolKind,
			Monitorstats:       verifychain.Monitorstats,
		}
		if member.BeginAt == verifychain.PrimaryLevel {
			member.BeginAt = 0 //
		}
		if member.LedgerAlignRelease == "REDACTED" {
			member.LedgerAlignRelease = "REDACTED"
		}
		if memberDeclaration.Style != "REDACTED" {
			member.Style = Style(memberDeclaration.Style)
		}
		if member.Style == StyleRapid {
			member.IfaceProtocol = ProtocolIntrinsic
		}
		if memberDeclaration.Datastore != "REDACTED" {
			member.Datastore = memberDeclaration.Datastore
		}
		if memberDeclaration.PrivatekeyProtocol != "REDACTED" {
			member.PrivatekeyProtocol = Protocol(memberDeclaration.PrivatekeyProtocol)
		}
		if memberDeclaration.EndureCadence != nil {
			member.EndureCadence = *memberDeclaration.EndureCadence
		}
		if member.Monitorstats {
			member.MonitorstatsGatewayPort = monitorstatsGatewayPortGenerate.Following()
		}
		for _, p := range memberDeclaration.Vary {
			member.Variations = append(member.Variations, Variation(p))
		}
		verifychain.Instances = append(verifychain.Instances, member)
	}

	//
	for _, member := range verifychain.Instances {
		memberDeclaration, ok := declaration.Instances[member.Label]
		if !ok {
			return nil, fmt.Errorf("REDACTED", member.Label)
		}
		for _, originLabel := range memberDeclaration.Origins {
			origin := verifychain.SearchMember(originLabel)
			if origin == nil {
				return nil, fmt.Errorf("REDACTED", originLabel, member.Label)
			}
			member.Origins = append(member.Origins, origin)
		}
		for _, nodeLabel := range memberDeclaration.DurableNodes {
			node := verifychain.SearchMember(nodeLabel)
			if node == nil {
				return nil, fmt.Errorf("REDACTED", nodeLabel, member.Label)
			}
			member.DurableNodes = append(member.DurableNodes, node)
		}

		//
		//
		if len(member.DurableNodes) == 0 && len(member.Origins) == 0 {
			for _, node := range verifychain.Instances {
				if node.Label == member.Label {
					continue
				}
				member.DurableNodes = append(member.DurableNodes, node)
			}
		}
	}

	//
	if declaration.Ratifiers != nil {
		for ratifierLabel, energy := range *declaration.Ratifiers {
			ratifier := verifychain.SearchMember(ratifierLabel)
			if ratifier == nil {
				return nil, fmt.Errorf("REDACTED", ratifierLabel)
			}
			verifychain.Ratifiers[ratifier] = energy
		}
	} else {
		for _, member := range verifychain.Instances {
			if member.Style == StyleRatifier {
				verifychain.Ratifiers[member] = 100
			}
		}
	}

	//
	for levelStr, ratifiers := range declaration.RatifierRefreshes {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", level, err)
		}
		valueModify := map[*Member]int64{}
		for label, energy := range ratifiers {
			member := verifychain.SearchMember(label)
			if member == nil {
				return nil, fmt.Errorf("REDACTED", label, level)
			}
			valueModify[member] = energy
		}
		verifychain.RatifierRefreshes[int64(level)] = valueModify
	}

	return verifychain, verifychain.Certify()
}

//
func (t Verifychain) Certify() error {
	if t.Label == "REDACTED" {
		return errors.New("REDACTED")
	}
	if t.IP == nil {
		return errors.New("REDACTED")
	}
	if len(t.Instances) == 0 {
		return errors.New("REDACTED")
	}
	if t.LedgerMaximumOctets > kinds.MaximumLedgerVolumeOctets {
		return fmt.Errorf("REDACTED", kinds.MaximumLedgerVolumeOctets)
	}
	if t.BallotPluginsModifyLevel < -1 {
		return fmt.Errorf("REDACTED"+
			"REDACTED", t.BallotPluginsModifyLevel)
	}
	if t.BallotPluginsActivateLevel < 0 {
		return fmt.Errorf("REDACTED"+
			"REDACTED", t.BallotPluginsActivateLevel)
	}
	if t.BallotPluginsModifyLevel > 0 && t.BallotPluginsModifyLevel < t.PrimaryLevel {
		return fmt.Errorf("REDACTED"+
			"REDACTED"+
			"REDACTED",
			t.BallotPluginsModifyLevel, t.PrimaryLevel,
		)
	}
	if t.BallotPluginsActivateLevel > 0 {
		if t.BallotPluginsActivateLevel < t.PrimaryLevel {
			return fmt.Errorf("REDACTED"+
				"REDACTED"+
				"REDACTED",
				t.BallotPluginsActivateLevel, t.PrimaryLevel,
			)
		}
		if t.BallotPluginsActivateLevel <= t.BallotPluginsModifyLevel {
			return fmt.Errorf("REDACTED"+
				"REDACTED"+
				"REDACTED",
				t.BallotPluginsModifyLevel, t.BallotPluginsActivateLevel,
			)
		}
	}
	for _, member := range t.Instances {
		if err := member.Certify(t); err != nil {
			return fmt.Errorf("REDACTED", member.Label, err)
		}
	}
	return nil
}

//
func (n Member) Certify(verifychain Verifychain) error {
	if n.Label == "REDACTED" {
		return errors.New("REDACTED")
	}
	if n.IntrinsicIP == nil {
		return errors.New("REDACTED")
	}
	if !verifychain.IP.Contains(n.IntrinsicIP) {
		return fmt.Errorf("REDACTED", n.IntrinsicIP, verifychain.IP)
	}
	if n.GatewayPort == n.MonitorstatsGatewayPort {
		return fmt.Errorf("REDACTED", n.GatewayPort)
	}
	if n.GatewayPort > 0 && n.GatewayPort <= 1024 {
		return fmt.Errorf("REDACTED", n.GatewayPort)
	}
	if n.MonitorstatsGatewayPort > 0 && n.MonitorstatsGatewayPort <= 1024 {
		return fmt.Errorf("REDACTED", n.MonitorstatsGatewayPort)
	}
	for _, node := range verifychain.Instances {
		if node.Label != n.Label && node.GatewayPort == n.GatewayPort && node.OutsideIP.Equal(n.OutsideIP) {
			return fmt.Errorf("REDACTED", node.Label, n.GatewayPort)
		}
		if n.MonitorstatsGatewayPort > 0 {
			if node.Label != n.Label && node.MonitorstatsGatewayPort == n.MonitorstatsGatewayPort {
				return fmt.Errorf("REDACTED", node.Label, n.MonitorstatsGatewayPort)
			}
		}
	}
	switch n.LedgerAlignRelease {
	case "REDACTED":
	default:
		return fmt.Errorf("REDACTED", n.LedgerAlignRelease)
	}
	switch n.Datastore {
	case "REDACTED", "REDACTED", "REDACTED", "REDACTED":
	default:
		return fmt.Errorf("REDACTED", n.Datastore)
	}
	switch n.IfaceProtocol {
	case ProtocolIntrinsic, ProtocolIntrinsicLinkAlign, ProtocolUNIX, ProtocolTCP, ProtocolGRPC:
	default:
		return fmt.Errorf("REDACTED", n.IfaceProtocol)
	}
	if n.Style == StyleRapid && n.IfaceProtocol != ProtocolIntrinsic && n.IfaceProtocol != ProtocolIntrinsicLinkAlign {
		return errors.New("REDACTED")
	}
	switch n.PrivatekeyProtocol {
	case ProtocolEntry, ProtocolUNIX, ProtocolTCP:
	default:
		return fmt.Errorf("REDACTED", n.PrivatekeyProtocol)
	}

	if n.BeginAt > 0 && n.BeginAt < n.Verifychain.PrimaryLevel {
		return fmt.Errorf("REDACTED",
			n.BeginAt, n.Verifychain.PrimaryLevel)
	}
	if n.StatusAlign && n.BeginAt == 0 {
		return errors.New("REDACTED")
	}
	if n.PreserveLedgers != 0 && n.PreserveLedgers < uint64(ProofEraLevel) {
		return fmt.Errorf("REDACTED",
			ProofEraLevel)
	}
	if n.EndureCadence == 0 && n.PreserveLedgers > 0 {
		return errors.New("REDACTED")
	}
	if n.EndureCadence > 1 && n.PreserveLedgers > 0 && n.PreserveLedgers < n.EndureCadence {
		return errors.New("REDACTED")
	}
	if n.MirrorCadence > 0 && n.PreserveLedgers > 0 && n.PreserveLedgers < n.MirrorCadence {
		return errors.New("REDACTED")
	}

	var enhanceLocated bool
	for _, variation := range n.Variations {
		switch variation {
		case VariationEnhance:
			if enhanceLocated {
				return fmt.Errorf("REDACTED")
			}
			enhanceLocated = true
		case VariationDetach, VariationTerminate, VariationStall, VariationReboot:
		default:
			return fmt.Errorf("REDACTED", variation)
		}
	}

	return nil
}

//
func (t Verifychain) SearchMember(label string) *Member {
	for _, member := range t.Instances {
		if member.Label == label {
			return member
		}
	}
	return nil
}

//
//
//
func (t Verifychain) CatalogInstances() []*Member {
	instances := []*Member{}
	for _, member := range t.Instances {
		if !member.Untracked() && member.BeginAt == 0 && member.PreserveLedgers == 0 {
			instances = append(instances, member)
		}
	}
	return instances
}

//
func (t Verifychain) ArbitraryMember() *Member {
	for {
		member := t.Instances[rand.Intn(len(t.Instances))] //
		if member.Style != StyleOrigin {
			return member
		}
	}
}

//
func (t Verifychain) IDXIpv6() bool {
	return t.IP.IP.To4() == nil
}

//
func (t Verifychain) HasVariations() bool {
	for _, member := range t.Instances {
		if len(member.Variations) > 0 {
			return true
		}
	}
	return false
}

//
var monitorstatsYamlPrototype string

func (t Verifychain) monitorstatsSettingsOctets() ([]byte, error) {
	tmpl, err := template.New("REDACTED").Parse(monitorstatsYamlPrototype)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t Verifychain) RecordMonitorstatsSettings() error {
	octets, err := t.monitorstatsSettingsOctets()
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(t.Dir, "REDACTED"), octets, 0o644) //
	if err != nil {
		return err
	}
	return nil
}

//
func (n Member) LocationP2P(withUID bool) string {
	ip := n.IntrinsicIP.String()
	if n.IntrinsicIP.To4() == nil {
		//
		ip = fmt.Sprintf("REDACTED", ip)
	}
	address := fmt.Sprintf("REDACTED", ip)
	if withUID {
		address = fmt.Sprintf("REDACTED", n.MemberKey.PublicKey().Location().Octets(), address)
	}
	return address
}

//
func (n Member) LocationRPC() string {
	ip := n.IntrinsicIP.String()
	if n.IntrinsicIP.To4() == nil {
		//
		ip = fmt.Sprintf("REDACTED", ip)
	}
	return fmt.Sprintf("REDACTED", ip)
}

//
func (n Member) Customer() (*rpchttp.HTTP, error) {
	return rpchttp.New(fmt.Sprintf("REDACTED", n.OutsideIP, n.GatewayPort), "REDACTED")
}

//
func (n Member) Untracked() bool {
	return n.Style == StyleRapid || n.Style == StyleOrigin
}

//
type keyProducer struct {
	arbitrary *rand.Rand
}

func newKeyProducer(origin int64) *keyProducer {
	return &keyProducer{
		arbitrary: rand.New(rand.NewSource(origin)), //
	}
}

func (g *keyProducer) Compose(keyKind string) vault.PrivateKey {
	origin := make([]byte, ed25519.OriginVolume)

	_, err := io.ReadFull(g.arbitrary, origin)
	if err != nil {
		panic(err) //
	}
	switch keyKind {
	case secp256k1.KeyKind:
		return secp256k1.GeneratePrivateKeySecp256k1(origin)
	case bls12381.KeyKind:
		pk, err := bls12381.GeneratePrivateKeyFromPrivatekey(origin)
		if err != nil {
			panic(fmt.Sprintf("REDACTED", bls12381.KeyKind, err))
		}
		return pk
	case ed25519.KeyKind:
		return ed25519.GeneratePrivateKeyFromPrivatekey(origin)
	default:
		return ed25519.GeneratePrivateKeyFromPrivatekey(origin) //
	}
}

//
type portProducer struct {
	followingPort uint32
}

func newPortProducer(initialPort uint32) *portProducer {
	return &portProducer{followingPort: initialPort}
}

func (g *portProducer) Following() uint32 {
	port := g.followingPort
	g.followingPort++
	if g.followingPort == 0 {
		panic("REDACTED")
	}
	return port
}

//
//
type ipProducer struct {
	fabric *net.IPNet
	followingIP  net.IP
}

func newIPProducer(fabric *net.IPNet) *ipProducer {
	followingIP := make([]byte, len(fabric.IP))
	copy(followingIP, fabric.IP)
	gen := &ipProducer{fabric: fabric, followingIP: followingIP}
	//
	gen.Following()
	gen.Following()
	return gen
}

func (g *ipProducer) Fabric() *net.IPNet {
	n := &net.IPNet{
		IP:   make([]byte, len(g.fabric.IP)),
		Mask: make([]byte, len(g.fabric.Mask)),
	}
	copy(n.IP, g.fabric.IP)
	copy(n.Mask, g.fabric.Mask)
	return n
}

func (g *ipProducer) Following() net.IP {
	ip := make([]byte, len(g.followingIP))
	copy(ip, g.followingIP)
	for i := len(g.followingIP) - 1; i >= 0; i-- {
		g.followingIP[i]++
		if g.followingIP[i] != 0 {
			break
		}
	}
	return ip
}
