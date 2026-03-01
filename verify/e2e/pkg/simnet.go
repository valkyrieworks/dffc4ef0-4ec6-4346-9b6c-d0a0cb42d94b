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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/signature381"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"

	_ "embed"
)

const (
	unpredictableGerm               int64  = 2308084734268
	delegateChannelInitial           uint32 = 5701
	titanDelegateChannelInitial uint32 = 6701

	fallbackClusterExtent   = 2
	fallbackLinkages = 1
	fallbackTransferExtentOctets = 1024

	regionalEdition = "REDACTED"
)

type (
	Style         string
	Scheme     string
	Disruption string
)

const (
	StyleAssessor Style = "REDACTED"
	StyleComplete      Style = "REDACTED"
	StyleAgile     Style = "REDACTED"
	StyleGerm      Style = "REDACTED"

	SchemeIntrinsic         Scheme = "REDACTED"
	SchemeIntrinsicLinkChronize Scheme = "REDACTED"
	SchemeRecord            Scheme = "REDACTED"
	SchemeGRPS            Scheme = "REDACTED"
	SchemeTcpsocket             Scheme = "REDACTED"
	SchemePosix            Scheme = "REDACTED"

	DisruptionDetach Disruption = "REDACTED"
	DisruptionTerminate       Disruption = "REDACTED"
	DisruptionBreak      Disruption = "REDACTED"
	DisruptionReboot    Disruption = "REDACTED"
	DisruptionModernize    Disruption = "REDACTED"

	ProofLifespanAltitude int64         = 14
	ProofLifespanMoment   time.Duration = 1500 * time.Millisecond
)

//
type Simnet struct {
	Alias                                                 string
	Record                                                 string
	Dir                                                  string
	IP                                                   *net.IPNet
	PrimaryAltitude                                        int64
	PrimaryStatus                                         map[string]string
	Assessors                                           map[*Peer]int64
	AssessorRevisions                                     map[int64]map[*Peer]int64
	Peers                                                []*Peer
	TokenKind                                              string
	Proof                                             int
	FetchTransferExtentOctets                                      int
	FetchTransferClusterExtent                                      int
	FetchTransferLinkages                                    int
	FetchMaximumTrans                                           int
	IfaceScheme                                         string
	ArrangeNominationDeferral                                 time.Duration
	HandleNominationDeferral                                 time.Duration
	InspectTransferDeferral                                         time.Duration
	BallotAdditionDeferral                                   time.Duration
	CulminateLedgerDeferral                                   time.Duration
	ModernizeEdition                                       string
	RecordStratum                                             string
	RecordLayout                                            string
	Titan                                           bool
	LedgerMaximumOctets                                        int64
	BallotAdditionsActivateAltitude                           int64
	BallotAdditionsReviseAltitude                           int64
	BallotAdditionExtent                                    uint
	ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes    uint
	ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes uint
}

//
type Peer struct {
	Alias                  string
	Edition               string
	Simnet               *Simnet
	Style                  Style
	PrivatevalueToken            security.PrivateToken
	PeerToken               security.PrivateToken
	IntrinsicINET            net.IP
	OutsideINET            net.IP
	DelegateChannel             uint32
	InitiateLocated               int64
	LedgerChronizeEdition      string
	LedgerChronizeAggregateStyle bool
	StatusChronize             bool
	Repository              string
	IfaceScheme          Scheme
	PrivatevalueScheme       Scheme
	EndureDuration       uint64
	ImageDuration      uint64
	PreserveLedgers          uint64
	Origins                 []*Peer
	EnduringNodes       []*Peer
	Disruptions         []Disruption
	TransmitNegativeFetch            bool
	Titan            bool
	UtilizeLibpeer2peer             bool
	TxpoolKind           string
	TitanDelegateChannel   uint32
}

//
//
//
//
//
func FetchSimnet(record string, ifd FrameworkData) (*Simnet, error) {
	declaration, err := FetchDeclaration(record)
	if err != nil {
		return nil, err
	}
	return FreshSimnetOriginatingDeclaration(declaration, record, ifd)
}

//
func FreshSimnetOriginatingDeclaration(declaration Declaration, record string, ifd FrameworkData) (*Simnet, error) {
	dir := strings.TrimSuffix(record, filepath.Ext(record))

	tokenProduce := freshTokenProducer(unpredictableGerm)
	titanDelegateChannelProduce := freshChannelProducer(titanDelegateChannelInitial)
	_, inetNetwork, err := net.ParseCIDR(ifd.Fabric)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", ifd.Fabric, err)
	}

	simnet := &Simnet{
		Alias:                       filepath.Base(dir),
		Record:                       record,
		Dir:                        dir,
		IP:                         inetNetwork,
		PrimaryAltitude:              1,
		PrimaryStatus:               declaration.PrimaryStatus,
		Assessors:                 map[*Peer]int64{},
		AssessorRevisions:           map[int64]map[*Peer]int64{},
		Peers:                      []*Peer{},
		TokenKind:                    declaration.TokenKind,
		Proof:                   declaration.Proof,
		FetchTransferExtentOctets:            declaration.FetchTransferExtentOctets,
		FetchTransferClusterExtent:            declaration.FetchTransferClusterExtent,
		FetchTransferLinkages:          declaration.FetchTransferLinkages,
		FetchMaximumTrans:                 declaration.FetchMaximumTrans,
		IfaceScheme:               declaration.IfaceScheme,
		ArrangeNominationDeferral:       declaration.ArrangeNominationDeferral,
		HandleNominationDeferral:       declaration.HandleNominationDeferral,
		InspectTransferDeferral:               declaration.InspectTransferDeferral,
		BallotAdditionDeferral:         declaration.BallotAdditionDeferral,
		CulminateLedgerDeferral:         declaration.CulminateLedgerDeferral,
		ModernizeEdition:             declaration.ModernizeEdition,
		RecordStratum:                   declaration.RecordStratum,
		RecordLayout:                  declaration.RecordLayout,
		Titan:                 declaration.Titan,
		LedgerMaximumOctets:              declaration.LedgerMaximumOctets,
		BallotAdditionsActivateAltitude: declaration.BallotAdditionsActivateAltitude,
		BallotAdditionsReviseAltitude: declaration.BallotAdditionsReviseAltitude,
		BallotAdditionExtent:          declaration.BallotAdditionExtent,
		ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes:    declaration.ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes,
		ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes: declaration.ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes,
	}

	if len(declaration.TokenKind) != 0 {
		simnet.TokenKind = declaration.TokenKind
	}
	if declaration.PrimaryAltitude > 0 {
		simnet.PrimaryAltitude = declaration.PrimaryAltitude
	}
	if simnet.TokenKind == "REDACTED" {
		simnet.TokenKind = edwards25519.TokenKind
	}
	if simnet.IfaceScheme == "REDACTED" {
		simnet.IfaceScheme = string(SchemeIntrinsic)
	}
	if simnet.ModernizeEdition == "REDACTED" {
		simnet.ModernizeEdition = regionalEdition
	}
	if simnet.FetchTransferLinkages == 0 {
		simnet.FetchTransferLinkages = fallbackLinkages
	}
	if simnet.FetchTransferClusterExtent == 0 {
		simnet.FetchTransferClusterExtent = fallbackClusterExtent
	}
	if simnet.FetchTransferExtentOctets == 0 {
		simnet.FetchTransferExtentOctets = fallbackTransferExtentOctets
	}

	for _, alias := range arrangePeerIdentifiers(declaration) {
		peerDeclaration := declaration.Peers[alias]
		ind, ok := ifd.Replicates[alias]
		if !ok {
			return nil, fmt.Errorf("REDACTED", alias)
		}
		addnINET := ind.AddnINETLocator
		if len(addnINET) == 0 {
			addnINET = ind.INETLocator
		}
		v := peerDeclaration.Edition
		if v == "REDACTED" {
			v = regionalEdition
		}

		peer := &Peer{
			Alias:                  alias,
			Edition:               v,
			Simnet:               simnet,
			PrivatevalueToken:            tokenProduce.Compose(declaration.TokenKind),
			PeerToken:               tokenProduce.Compose("REDACTED"),
			IntrinsicINET:            ind.INETLocator,
			OutsideINET:            addnINET,
			DelegateChannel:             ind.Channel,
			Style:                  StyleAssessor,
			Repository:              "REDACTED",
			IfaceScheme:          Scheme(simnet.IfaceScheme),
			PrivatevalueScheme:       SchemeRecord,
			InitiateLocated:               peerDeclaration.InitiateLocated,
			LedgerChronizeEdition:      peerDeclaration.LedgerChronizeEdition,
			LedgerChronizeAggregateStyle: peerDeclaration.LedgerChronizeAggregateStyle,
			StatusChronize:             peerDeclaration.StatusChronize,
			EndureDuration:       1,
			ImageDuration:      peerDeclaration.ImageDuration,
			PreserveLedgers:          peerDeclaration.PreserveLedgers,
			Disruptions:         []Disruption{},
			TransmitNegativeFetch:            peerDeclaration.TransmitNegativeFetch,
			UtilizeLibpeer2peer:             peerDeclaration.UtilizeLibpeer2peer,
			TxpoolKind:           peerDeclaration.TxpoolKind,
			Titan:            simnet.Titan,
		}
		if peer.InitiateLocated == simnet.PrimaryAltitude {
			peer.InitiateLocated = 0 //
		}
		if peer.LedgerChronizeEdition == "REDACTED" {
			peer.LedgerChronizeEdition = "REDACTED"
		}
		if peerDeclaration.Style != "REDACTED" {
			peer.Style = Style(peerDeclaration.Style)
		}
		if peer.Style == StyleAgile {
			peer.IfaceScheme = SchemeIntrinsic
		}
		if peerDeclaration.Repository != "REDACTED" {
			peer.Repository = peerDeclaration.Repository
		}
		if peerDeclaration.PrivatevalueScheme != "REDACTED" {
			peer.PrivatevalueScheme = Scheme(peerDeclaration.PrivatevalueScheme)
		}
		if peerDeclaration.EndureDuration != nil {
			peer.EndureDuration = *peerDeclaration.EndureDuration
		}
		if peer.Titan {
			peer.TitanDelegateChannel = titanDelegateChannelProduce.Following()
		}
		for _, p := range peerDeclaration.Disrupt {
			peer.Disruptions = append(peer.Disruptions, Disruption(p))
		}
		simnet.Peers = append(simnet.Peers, peer)
	}

	//
	for _, peer := range simnet.Peers {
		peerDeclaration, ok := declaration.Peers[peer.Alias]
		if !ok {
			return nil, fmt.Errorf("REDACTED", peer.Alias)
		}
		for _, germAlias := range peerDeclaration.Origins {
			germ := simnet.SearchPeer(germAlias)
			if germ == nil {
				return nil, fmt.Errorf("REDACTED", germAlias, peer.Alias)
			}
			peer.Origins = append(peer.Origins, germ)
		}
		for _, nodeAlias := range peerDeclaration.EnduringNodes {
			node := simnet.SearchPeer(nodeAlias)
			if node == nil {
				return nil, fmt.Errorf("REDACTED", nodeAlias, peer.Alias)
			}
			peer.EnduringNodes = append(peer.EnduringNodes, node)
		}

		//
		//
		if len(peer.EnduringNodes) == 0 && len(peer.Origins) == 0 {
			for _, node := range simnet.Peers {
				if node.Alias == peer.Alias {
					continue
				}
				peer.EnduringNodes = append(peer.EnduringNodes, node)
			}
		}
	}

	//
	if declaration.Assessors != nil {
		for assessorAlias, potency := range *declaration.Assessors {
			assessor := simnet.SearchPeer(assessorAlias)
			if assessor == nil {
				return nil, fmt.Errorf("REDACTED", assessorAlias)
			}
			simnet.Assessors[assessor] = potency
		}
	} else {
		for _, peer := range simnet.Peers {
			if peer.Style == StyleAssessor {
				simnet.Assessors[peer] = 100
			}
		}
	}

	//
	for altitudeTxt, assessors := range declaration.AssessorRevisions {
		altitude, err := strconv.Atoi(altitudeTxt)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", altitude, err)
		}
		itemRevise := map[*Peer]int64{}
		for alias, potency := range assessors {
			peer := simnet.SearchPeer(alias)
			if peer == nil {
				return nil, fmt.Errorf("REDACTED", alias, altitude)
			}
			itemRevise[peer] = potency
		}
		simnet.AssessorRevisions[int64(altitude)] = itemRevise
	}

	return simnet, simnet.Certify()
}

//
func (t Simnet) Certify() error {
	if t.Alias == "REDACTED" {
		return errors.New("REDACTED")
	}
	if t.IP == nil {
		return errors.New("REDACTED")
	}
	if len(t.Peers) == 0 {
		return errors.New("REDACTED")
	}
	if t.LedgerMaximumOctets > kinds.MaximumLedgerExtentOctets {
		return fmt.Errorf("REDACTED", kinds.MaximumLedgerExtentOctets)
	}
	if t.BallotAdditionsReviseAltitude < -1 {
		return fmt.Errorf("REDACTED"+
			"REDACTED", t.BallotAdditionsReviseAltitude)
	}
	if t.BallotAdditionsActivateAltitude < 0 {
		return fmt.Errorf("REDACTED"+
			"REDACTED", t.BallotAdditionsActivateAltitude)
	}
	if t.BallotAdditionsReviseAltitude > 0 && t.BallotAdditionsReviseAltitude < t.PrimaryAltitude {
		return fmt.Errorf("REDACTED"+
			"REDACTED"+
			"REDACTED",
			t.BallotAdditionsReviseAltitude, t.PrimaryAltitude,
		)
	}
	if t.BallotAdditionsActivateAltitude > 0 {
		if t.BallotAdditionsActivateAltitude < t.PrimaryAltitude {
			return fmt.Errorf("REDACTED"+
				"REDACTED"+
				"REDACTED",
				t.BallotAdditionsActivateAltitude, t.PrimaryAltitude,
			)
		}
		if t.BallotAdditionsActivateAltitude <= t.BallotAdditionsReviseAltitude {
			return fmt.Errorf("REDACTED"+
				"REDACTED"+
				"REDACTED",
				t.BallotAdditionsReviseAltitude, t.BallotAdditionsActivateAltitude,
			)
		}
	}
	for _, peer := range t.Peers {
		if err := peer.Certify(t); err != nil {
			return fmt.Errorf("REDACTED", peer.Alias, err)
		}
	}
	return nil
}

//
func (n Peer) Certify(simnet Simnet) error {
	if n.Alias == "REDACTED" {
		return errors.New("REDACTED")
	}
	if n.IntrinsicINET == nil {
		return errors.New("REDACTED")
	}
	if !simnet.IP.Contains(n.IntrinsicINET) {
		return fmt.Errorf("REDACTED", n.IntrinsicINET, simnet.IP)
	}
	if n.DelegateChannel == n.TitanDelegateChannel {
		return fmt.Errorf("REDACTED", n.DelegateChannel)
	}
	if n.DelegateChannel > 0 && n.DelegateChannel <= 1024 {
		return fmt.Errorf("REDACTED", n.DelegateChannel)
	}
	if n.TitanDelegateChannel > 0 && n.TitanDelegateChannel <= 1024 {
		return fmt.Errorf("REDACTED", n.TitanDelegateChannel)
	}
	for _, node := range simnet.Peers {
		if node.Alias != n.Alias && node.DelegateChannel == n.DelegateChannel && node.OutsideINET.Equal(n.OutsideINET) {
			return fmt.Errorf("REDACTED", node.Alias, n.DelegateChannel)
		}
		if n.TitanDelegateChannel > 0 {
			if node.Alias != n.Alias && node.TitanDelegateChannel == n.TitanDelegateChannel {
				return fmt.Errorf("REDACTED", node.Alias, n.TitanDelegateChannel)
			}
		}
	}
	switch n.LedgerChronizeEdition {
	case "REDACTED":
	default:
		return fmt.Errorf("REDACTED", n.LedgerChronizeEdition)
	}
	switch n.Repository {
	case "REDACTED", "REDACTED", "REDACTED", "REDACTED":
	default:
		return fmt.Errorf("REDACTED", n.Repository)
	}
	switch n.IfaceScheme {
	case SchemeIntrinsic, SchemeIntrinsicLinkChronize, SchemePosix, SchemeTcpsocket, SchemeGRPS:
	default:
		return fmt.Errorf("REDACTED", n.IfaceScheme)
	}
	if n.Style == StyleAgile && n.IfaceScheme != SchemeIntrinsic && n.IfaceScheme != SchemeIntrinsicLinkChronize {
		return errors.New("REDACTED")
	}
	switch n.PrivatevalueScheme {
	case SchemeRecord, SchemePosix, SchemeTcpsocket:
	default:
		return fmt.Errorf("REDACTED", n.PrivatevalueScheme)
	}

	if n.InitiateLocated > 0 && n.InitiateLocated < n.Simnet.PrimaryAltitude {
		return fmt.Errorf("REDACTED",
			n.InitiateLocated, n.Simnet.PrimaryAltitude)
	}
	if n.StatusChronize && n.InitiateLocated == 0 {
		return errors.New("REDACTED")
	}
	if n.PreserveLedgers != 0 && n.PreserveLedgers < uint64(ProofLifespanAltitude) {
		return fmt.Errorf("REDACTED",
			ProofLifespanAltitude)
	}
	if n.EndureDuration == 0 && n.PreserveLedgers > 0 {
		return errors.New("REDACTED")
	}
	if n.EndureDuration > 1 && n.PreserveLedgers > 0 && n.PreserveLedgers < n.EndureDuration {
		return errors.New("REDACTED")
	}
	if n.ImageDuration > 0 && n.PreserveLedgers > 0 && n.PreserveLedgers < n.ImageDuration {
		return errors.New("REDACTED")
	}

	var modernizeDetected bool
	for _, disruption := range n.Disruptions {
		switch disruption {
		case DisruptionModernize:
			if modernizeDetected {
				return fmt.Errorf("REDACTED")
			}
			modernizeDetected = true
		case DisruptionDetach, DisruptionTerminate, DisruptionBreak, DisruptionReboot:
		default:
			return fmt.Errorf("REDACTED", disruption)
		}
	}

	return nil
}

//
func (t Simnet) SearchPeer(alias string) *Peer {
	for _, peer := range t.Peers {
		if peer.Alias == alias {
			return peer
		}
	}
	return nil
}

//
//
//
func (t Simnet) RepositoryPeers() []*Peer {
	peers := []*Peer{}
	for _, peer := range t.Peers {
		if !peer.Untracked() && peer.InitiateLocated == 0 && peer.PreserveLedgers == 0 {
			peers = append(peers, peer)
		}
	}
	return peers
}

//
func (t Simnet) UnpredictablePeer() *Peer {
	for {
		peer := t.Peers[rand.Intn(len(t.Peers))] //
		if peer.Style != StyleGerm {
			return peer
		}
	}
}

//
func (t Simnet) IDXPrv6() bool {
	return t.IP.IP.To4() == nil
}

//
func (t Simnet) OwnsDisruptions() bool {
	for _, peer := range t.Peers {
		if len(peer.Disruptions) > 0 {
			return true
		}
	}
	return false
}

//
var titanYamlspecBlueprint string

func (t Simnet) titanSettingsOctets() ([]byte, error) {
	layout, err := template.New("REDACTED").Parse(titanYamlspecBlueprint)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = layout.Execute(&buf, t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t Simnet) PersistTitanSettings() error {
	octets, err := t.titanSettingsOctets()
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
func (n Peer) LocatorPeer2peer(usingUUID bool) string {
	ip := n.IntrinsicINET.String()
	if n.IntrinsicINET.To4() == nil {
		//
		ip = fmt.Sprintf("REDACTED", ip)
	}
	location := fmt.Sprintf("REDACTED", ip)
	if usingUUID {
		location = fmt.Sprintf("REDACTED", n.PeerToken.PublicToken().Location().Octets(), location)
	}
	return location
}

//
func (n Peer) LocatorRemote() string {
	ip := n.IntrinsicINET.String()
	if n.IntrinsicINET.To4() == nil {
		//
		ip = fmt.Sprintf("REDACTED", ip)
	}
	return fmt.Sprintf("REDACTED", ip)
}

//
func (n Peer) Customer() (*rpchttpsvc.Httpsvc, error) {
	return rpchttpsvc.New(fmt.Sprintf("REDACTED", n.OutsideINET, n.DelegateChannel), "REDACTED")
}

//
func (n Peer) Untracked() bool {
	return n.Style == StyleAgile || n.Style == StyleGerm
}

//
type tokenProducer struct {
	unpredictable *rand.Rand
}

func freshTokenProducer(germ int64) *tokenProducer {
	return &tokenProducer{
		unpredictable: rand.New(rand.NewSource(germ)), //
	}
}

func (g *tokenProducer) Compose(tokenKind string) security.PrivateToken {
	germ := make([]byte, edwards25519.GermExtent)

	_, err := io.ReadFull(g.unpredictable, germ)
	if err != nil {
		panic(err) //
	}
	switch tokenKind {
	case ellipticp256.TokenKind:
		return ellipticp256.ProducePrivateTokenEllipticp256(germ)
	case signature381.TokenKind:
		pk, err := signature381.ProducePrivateTokenOriginatingCredential(germ)
		if err != nil {
			panic(fmt.Sprintf("REDACTED", signature381.TokenKind, err))
		}
		return pk
	case edwards25519.TokenKind:
		return edwards25519.ProducePrivateTokenOriginatingCredential(germ)
	default:
		return edwards25519.ProducePrivateTokenOriginatingCredential(germ) //
	}
}

//
type channelProducer struct {
	followingChannel uint32
}

func freshChannelProducer(initialChannel uint32) *channelProducer {
	return &channelProducer{followingChannel: initialChannel}
}

func (g *channelProducer) Following() uint32 {
	channel := g.followingChannel
	g.followingChannel++
	if g.followingChannel == 0 {
		panic("REDACTED")
	}
	return channel
}

//
//
type inetProducer struct {
	fabric *net.IPNet
	followingINET  net.IP
}

func freshINETProducer(fabric *net.IPNet) *inetProducer {
	followingINET := make([]byte, len(fabric.IP))
	copy(followingINET, fabric.IP)
	gen := &inetProducer{fabric: fabric, followingINET: followingINET}
	//
	gen.Following()
	gen.Following()
	return gen
}

func (g *inetProducer) Fabric() *net.IPNet {
	n := &net.IPNet{
		IP:   make([]byte, len(g.fabric.IP)),
		Mask: make([]byte, len(g.fabric.Mask)),
	}
	copy(n.IP, g.fabric.IP)
	copy(n.Mask, g.fabric.Mask)
	return n
}

func (g *inetProducer) Following() net.IP {
	ip := make([]byte, len(g.followingINET))
	copy(ip, g.followingINET)
	for i := len(g.followingINET) - 1; i >= 0; i-- {
		g.followingINET[i]++
		if g.followingINET[i] != 0 {
			break
		}
	}
	return ip
}
