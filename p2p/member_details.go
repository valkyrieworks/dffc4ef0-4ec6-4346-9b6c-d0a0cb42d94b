package p2p

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	cometbytes "github.com/valkyrieworks/utils/octets"
	endorsementsstrings "github.com/valkyrieworks/utils/strings"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
	"github.com/valkyrieworks/release"
)

const (
	maximumMemberDetailsVolume = 10240 //
	maximumCountStreams  = 16    //
)

//
func MaximumMemberDetailsVolume() int {
	return maximumMemberDetailsVolume
}

//

//
//
type MemberDetails interface {
	ID() ID
	memberDetailsLocation
	memberDetailsCarrier
}

type memberDetailsLocation interface {
	NetLocation() (*NetLocation, error)
}

//
//
type memberDetailsCarrier interface {
	Certify() error
	HarmoniousWith(another MemberDetails) error
}

//

//
type ProtocolRelease struct {
	P2P   uint64 `json:"p2p"`
	Ledger uint64 `json:"ledger"`
	App   uint64 `json:"app"`
}

//
//
var standardProtocolRelease = NewProtocolRelease(
	release.P2PProtocol,
	release.LedgerProtocol,
	0,
)

//
func NewProtocolRelease(p2p, ledger, app uint64) ProtocolRelease {
	return ProtocolRelease{
		P2P:   p2p,
		Ledger: ledger,
		App:   app,
	}
}

//

//
var _ MemberDetails = StandardMemberDetails{}

//
//
type StandardMemberDetails struct {
	ProtocolRelease ProtocolRelease `json:"protocol_release"`

	//
	//
	StandardMemberUID ID     `json:"id"`          //
	ObserveAddress    string `json:"observe_address"` //

	//
	//
	Fabric  string            `json:"fabric"`  //
	Release  string            `json:"release"`  //
	Streams cometbytes.HexOctets `json:"streams"` //

	//
	Moniker string               `json:"moniker"` //
	Another   StandardMemberDetailsAnother `json:"another"`   //
}

//
type StandardMemberDetailsAnother struct {
	TransOrdinal    string `json:"transfer_ordinal"`
	RPCLocation string `json:"rpc_location"`
}

//
func (details StandardMemberDetails) ID() ID {
	return details.StandardMemberUID
}

//
//
//
//
//
//
//
//
//
//
//
//
//
func (details StandardMemberDetails) Certify() error {
	//

	//
	_, err := NewNetLocationString(UIDLocationString(details.ID(), details.ObserveAddress))
	if err != nil {
		return err
	}

	//

	//
	if len(details.Release) > 0 &&
		(!endorsementsstrings.IsAsciiContent(details.Release) || endorsementsstrings.AsciiShave(details.Release) == "REDACTED") {

		return fmt.Errorf("REDACTED", details.Release)
	}

	//
	if len(details.Streams) > maximumCountStreams {
		return fmt.Errorf("REDACTED", len(details.Streams), maximumCountStreams)
	}
	streams := make(map[byte]struct{})
	for _, ch := range details.Streams {
		_, ok := streams[ch]
		if ok {
			return fmt.Errorf("REDACTED", ch)
		}
		streams[ch] = struct{}{}
	}

	//
	if !endorsementsstrings.IsAsciiContent(details.Moniker) || endorsementsstrings.AsciiShave(details.Moniker) == "REDACTED" {
		return fmt.Errorf("REDACTED", details.Moniker)
	}

	//
	another := details.Another
	transOrdinal := another.TransOrdinal
	switch transOrdinal {
	case "REDACTED", "REDACTED", "REDACTED":
	default:
		return fmt.Errorf("REDACTED", transOrdinal)
	}
	//
	rpcAddress := another.RPCLocation
	if len(rpcAddress) > 0 && (!endorsementsstrings.IsAsciiContent(rpcAddress) || endorsementsstrings.AsciiShave(rpcAddress) == "REDACTED") {
		return fmt.Errorf("REDACTED", rpcAddress)
	}

	return nil
}

//
//
//
func (details StandardMemberDetails) HarmoniousWith(anotherDetails MemberDetails) error {
	another, ok := anotherDetails.(StandardMemberDetails)
	if !ok {
		return fmt.Errorf("REDACTED", reflect.TypeOf(anotherDetails))
	}

	if details.ProtocolRelease.Ledger != another.ProtocolRelease.Ledger {
		return fmt.Errorf("REDACTED",
			another.ProtocolRelease.Ledger, details.ProtocolRelease.Ledger)
	}

	//
	if details.Fabric != another.Fabric {
		return fmt.Errorf("REDACTED", another.Fabric, details.Fabric)
	}

	//
	if len(details.Streams) == 0 {
		return nil
	}

	//
	located := false
EXTERNAL_CYCLE:
	for _, ch1 := range details.Streams {
		for _, ch2 := range another.Streams {
			if ch1 == ch2 {
				located = true
				break EXTERNAL_CYCLE //
			}
		}
	}
	if !located {
		return fmt.Errorf("REDACTED", details.Streams, another.Streams)
	}
	return nil
}

//
//
//
//
func (details StandardMemberDetails) NetLocation() (*NetLocation, error) {
	uidAddress := UIDLocationString(details.ID(), details.ObserveAddress)
	return NewNetLocationString(uidAddress)
}

func (details StandardMemberDetails) HasConduit(chanUID byte) bool {
	return bytes.Contains(details.Streams, []byte{chanUID})
}

func (details StandardMemberDetails) ToSchema() *tmp2p.StandardMemberDetails {
	dni := new(tmp2p.StandardMemberDetails)
	dni.ProtocolRelease = tmp2p.ProtocolRelease{
		P2P:   details.ProtocolRelease.P2P,
		Ledger: details.ProtocolRelease.Ledger,
		App:   details.ProtocolRelease.App,
	}

	dni.StandardMemberUID = string(details.StandardMemberUID)
	dni.ObserveAddress = details.ObserveAddress
	dni.Fabric = details.Fabric
	dni.Release = details.Release
	dni.Streams = details.Streams
	dni.Moniker = details.Moniker
	dni.Another = tmp2p.StandardMemberDetailsAnother{
		TransOrdinal:    details.Another.TransOrdinal,
		RPCLocation: details.Another.RPCLocation,
	}

	return dni
}

func StandardMemberDetailsFromToSchema(pb *tmp2p.StandardMemberDetails) (StandardMemberDetails, error) {
	if pb == nil {
		return StandardMemberDetails{}, errors.New("REDACTED")
	}
	dni := StandardMemberDetails{
		ProtocolRelease: ProtocolRelease{
			P2P:   pb.ProtocolRelease.P2P,
			Ledger: pb.ProtocolRelease.Ledger,
			App:   pb.ProtocolRelease.App,
		},
		StandardMemberUID: ID(pb.StandardMemberUID),
		ObserveAddress:    pb.ObserveAddress,
		Fabric:       pb.Fabric,
		Release:       pb.Release,
		Streams:      pb.Streams,
		Moniker:       pb.Moniker,
		Another: StandardMemberDetailsAnother{
			TransOrdinal:    pb.Another.TransOrdinal,
			RPCLocation: pb.Another.RPCLocation,
		},
	}

	return dni, nil
}
