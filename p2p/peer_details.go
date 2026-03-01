package p2p

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	endorsementtexts "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/texts"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

const (
	maximumPeerDetailsExtent = 10240 //
	maximumCountConduits  = 16    //
)

//
func MaximumPeerDetailsExtent() int {
	return maximumPeerDetailsExtent
}

//

//
//
type PeerDetails interface {
	ID() ID
	peerDetailsLocator
	peerDetailsCarrier
}

type peerDetailsLocator interface {
	NetworkLocator() (*NetworkLocator, error)
}

//
//
type peerDetailsCarrier interface {
	Certify() error
	MatchedUsing(another PeerDetails) error
}

//

//
type SchemeEdition struct {
	P2P   uint64 `json:"p2p"`
	Ledger uint64 `json:"ledger"`
	App   uint64 `json:"app"`
}

//
//
var fallbackSchemeEdition = FreshSchemeEdition(
	edition.Peer2peerScheme,
	edition.LedgerScheme,
	0,
)

//
func FreshSchemeEdition(p2p, ledger, app uint64) SchemeEdition {
	return SchemeEdition{
		P2P:   p2p,
		Ledger: ledger,
		App:   app,
	}
}

//

//
var _ PeerDetails = FallbackPeerDetails{}

//
//
type FallbackPeerDetails struct {
	SchemeEdition SchemeEdition `json:"scheme_edition"`

	//
	//
	FallbackPeerUUID ID     `json:"id"`          //
	OverhearLocation    string `json:"overhear_location"` //

	//
	//
	Fabric  string            `json:"fabric"`  //
	Edition  string            `json:"edition"`  //
	Conduits tendermintoctets.HexadecimalOctets `json:"conduits"` //

	//
	Pseudonym string               `json:"pseudonym"` //
	Another   FallbackPeerDetailsAnother `json:"another"`   //
}

//
type FallbackPeerDetailsAnother struct {
	TransferOrdinal    string `json:"transfer_position"`
	RemoteLocator string `json:"remote_locator"`
}

//
func (details FallbackPeerDetails) ID() ID {
	return details.FallbackPeerUUID
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
func (details FallbackPeerDetails) Certify() error {
	//

	//
	_, err := FreshNetworkLocatorText(UUIDLocationText(details.ID(), details.OverhearLocation))
	if err != nil {
		return err
	}

	//

	//
	if len(details.Edition) > 0 &&
		(!endorsementtexts.EqualsCODETxt(details.Edition) || endorsementtexts.CODEShave(details.Edition) == "REDACTED") {

		return fmt.Errorf("REDACTED", details.Edition)
	}

	//
	if len(details.Conduits) > maximumCountConduits {
		return fmt.Errorf("REDACTED", len(details.Conduits), maximumCountConduits)
	}
	conduits := make(map[byte]struct{})
	for _, ch := range details.Conduits {
		_, ok := conduits[ch]
		if ok {
			return fmt.Errorf("REDACTED", ch)
		}
		conduits[ch] = struct{}{}
	}

	//
	if !endorsementtexts.EqualsCODETxt(details.Pseudonym) || endorsementtexts.CODEShave(details.Pseudonym) == "REDACTED" {
		return fmt.Errorf("REDACTED", details.Pseudonym)
	}

	//
	another := details.Another
	transferOrdinal := another.TransferOrdinal
	switch transferOrdinal {
	case "REDACTED", "REDACTED", "REDACTED":
	default:
		return fmt.Errorf("REDACTED", transferOrdinal)
	}
	//
	remoteLocation := another.RemoteLocator
	if len(remoteLocation) > 0 && (!endorsementtexts.EqualsCODETxt(remoteLocation) || endorsementtexts.CODEShave(remoteLocation) == "REDACTED") {
		return fmt.Errorf("REDACTED", remoteLocation)
	}

	return nil
}

//
//
//
func (details FallbackPeerDetails) MatchedUsing(anotherDetails PeerDetails) error {
	another, ok := anotherDetails.(FallbackPeerDetails)
	if !ok {
		return fmt.Errorf("REDACTED", reflect.TypeOf(anotherDetails))
	}

	if details.SchemeEdition.Ledger != another.SchemeEdition.Ledger {
		return fmt.Errorf("REDACTED",
			another.SchemeEdition.Ledger, details.SchemeEdition.Ledger)
	}

	//
	if details.Fabric != another.Fabric {
		return fmt.Errorf("REDACTED", another.Fabric, details.Fabric)
	}

	//
	if len(details.Conduits) == 0 {
		return nil
	}

	//
	detected := false
EXTERNAL_CYCLE:
	for _, ch1 := range details.Conduits {
		for _, ch2 := range another.Conduits {
			if ch1 == ch2 {
				detected = true
				break EXTERNAL_CYCLE //
			}
		}
	}
	if !detected {
		return fmt.Errorf("REDACTED", details.Conduits, another.Conduits)
	}
	return nil
}

//
//
//
//
func (details FallbackPeerDetails) NetworkLocator() (*NetworkLocator, error) {
	uuidLocation := UUIDLocationText(details.ID(), details.OverhearLocation)
	return FreshNetworkLocatorText(uuidLocation)
}

func (details FallbackPeerDetails) OwnsConduit(chnlUUID byte) bool {
	return bytes.Contains(details.Conduits, []byte{chnlUUID})
}

func (details FallbackPeerDetails) TowardSchema() *tmpfabric.FallbackPeerDetails {
	dni := new(tmpfabric.FallbackPeerDetails)
	dni.SchemeEdition = tmpfabric.SchemeEdition{
		P2P:   details.SchemeEdition.P2P,
		Ledger: details.SchemeEdition.Ledger,
		App:   details.SchemeEdition.App,
	}

	dni.FallbackPeerUUID = string(details.FallbackPeerUUID)
	dni.OverhearLocation = details.OverhearLocation
	dni.Fabric = details.Fabric
	dni.Edition = details.Edition
	dni.Conduits = details.Conduits
	dni.Pseudonym = details.Pseudonym
	dni.Another = tmpfabric.FallbackPeerDetailsAnother{
		TransferOrdinal:    details.Another.TransferOrdinal,
		RemoteLocator: details.Another.RemoteLocator,
	}

	return dni
}

func FallbackPeerDetailsOriginatingTowardSchema(pb *tmpfabric.FallbackPeerDetails) (FallbackPeerDetails, error) {
	if pb == nil {
		return FallbackPeerDetails{}, errors.New("REDACTED")
	}
	dni := FallbackPeerDetails{
		SchemeEdition: SchemeEdition{
			P2P:   pb.SchemeEdition.P2P,
			Ledger: pb.SchemeEdition.Ledger,
			App:   pb.SchemeEdition.App,
		},
		FallbackPeerUUID: ID(pb.FallbackPeerUUID),
		OverhearLocation:    pb.OverhearLocation,
		Fabric:       pb.Fabric,
		Edition:       pb.Edition,
		Conduits:      pb.Conduits,
		Pseudonym:       pb.Pseudonym,
		Another: FallbackPeerDetailsAnother{
			TransferOrdinal:    pb.Another.TransferOrdinal,
			RemoteLocator: pb.Another.RemoteLocator,
		},
	}

	return dni, nil
}
