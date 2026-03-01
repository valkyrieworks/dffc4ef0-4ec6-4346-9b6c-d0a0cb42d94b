//
//

package p2p

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

//
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
//
//
//
const _ = proto.GoGoProtoPackageIsVersion3 //

type NetworkLocator struct {
	ID   string `protobuf:"octets,1,opt,name=id,proto3" json:"id,omitempty"`
	IP   string `protobuf:"octets,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Channel uint32 `protobuf:"variableint,3,opt,name=port,proto3" json:"channel,omitempty"`
}

func (m *NetworkLocator) Restore()         { *m = NetworkLocator{} }
func (m *NetworkLocator) Text() string { return proto.CompactTextString(m) }
func (*NetworkLocator) SchemaArtifact()    {}
func (*NetworkLocator) Definition() ([]byte, []int) {
	return filedescriptor_c8a29e659aeca578, []int{0}
}
func (m *NetworkLocator) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *NetworkLocator) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Inetlocator.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkLocator) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Inetlocator.Merge(m, src)
}
func (m *NetworkLocator) XXX_Extent() int {
	return m.Extent()
}
func (m *NetworkLocator) XXX_Dropunfamiliar() {
	xxx_signaldetails_Inetlocator.DiscardUnknown(m)
}

var xxx_signaldetails_Inetlocator proto.InternalMessageInfo

func (m *NetworkLocator) ObtainUUID() string {
	if m != nil {
		return m.ID
	}
	return "REDACTED"
}

func (m *NetworkLocator) ObtainINET() string {
	if m != nil {
		return m.IP
	}
	return "REDACTED"
}

func (m *NetworkLocator) ObtainChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

type SchemeEdition struct {
	P2P   uint64 `protobuf:"variableint,1,opt,name=p2p,proto3" json:"p2p,omitempty"`
	Ledger uint64 `protobuf:"variableint,2,opt,name=block,proto3" json:"ledger,omitempty"`
	App   uint64 `protobuf:"variableint,3,opt,name=app,proto3" json:"app,omitempty"`
}

func (m *SchemeEdition) Restore()         { *m = SchemeEdition{} }
func (m *SchemeEdition) Text() string { return proto.CompactTextString(m) }
func (*SchemeEdition) SchemaArtifact()    {}
func (*SchemeEdition) Definition() ([]byte, []int) {
	return filedescriptor_c8a29e659aeca578, []int{1}
}
func (m *SchemeEdition) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SchemeEdition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Schemedition.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SchemeEdition) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Schemedition.Merge(m, src)
}
func (m *SchemeEdition) XXX_Extent() int {
	return m.Extent()
}
func (m *SchemeEdition) XXX_Dropunfamiliar() {
	xxx_signaldetails_Schemedition.DiscardUnknown(m)
}

var xxx_signaldetails_Schemedition proto.InternalMessageInfo

func (m *SchemeEdition) ObtainPeer2peer() uint64 {
	if m != nil {
		return m.P2P
	}
	return 0
}

func (m *SchemeEdition) ObtainLedger() uint64 {
	if m != nil {
		return m.Ledger
	}
	return 0
}

func (m *SchemeEdition) ObtainApplication() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type FallbackPeerDetails struct {
	SchemeEdition SchemeEdition      `protobuf:"octets,1,opt,name=protocol_version,json=protocolVersion,proto3" json:"scheme_edition"`
	FallbackPeerUUID   string               `protobuf:"octets,2,opt,name=default_node_id,json=defaultNodeId,proto3" json:"fallback_peer_uuid,omitempty"`
	OverhearLocation      string               `protobuf:"octets,3,opt,name=listen_addr,json=listenAddr,proto3" json:"overhear_location,omitempty"`
	Fabric         string               `protobuf:"octets,4,opt,name=network,proto3" json:"fabric,omitempty"`
	Edition         string               `protobuf:"octets,5,opt,name=version,proto3" json:"edition,omitempty"`
	Conduits        []byte               `protobuf:"octets,6,opt,name=channels,proto3" json:"conduits,omitempty"`
	Pseudonym         string               `protobuf:"octets,7,opt,name=moniker,proto3" json:"pseudonym,omitempty"`
	Another           FallbackPeerDetailsAnother `protobuf:"octets,8,opt,name=other,proto3" json:"another"`
}

func (m *FallbackPeerDetails) Restore()         { *m = FallbackPeerDetails{} }
func (m *FallbackPeerDetails) Text() string { return proto.CompactTextString(m) }
func (*FallbackPeerDetails) SchemaArtifact()    {}
func (*FallbackPeerDetails) Definition() ([]byte, []int) {
	return filedescriptor_c8a29e659aeca578, []int{2}
}
func (m *FallbackPeerDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *FallbackPeerDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Fallbacknodedetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FallbackPeerDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Fallbacknodedetails.Merge(m, src)
}
func (m *FallbackPeerDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *FallbackPeerDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Fallbacknodedetails.DiscardUnknown(m)
}

var xxx_signaldetails_Fallbacknodedetails proto.InternalMessageInfo

func (m *FallbackPeerDetails) ObtainSchemeEdition() SchemeEdition {
	if m != nil {
		return m.SchemeEdition
	}
	return SchemeEdition{}
}

func (m *FallbackPeerDetails) ObtainFallbackPeerUUID() string {
	if m != nil {
		return m.FallbackPeerUUID
	}
	return "REDACTED"
}

func (m *FallbackPeerDetails) ObtainOverhearLocation() string {
	if m != nil {
		return m.OverhearLocation
	}
	return "REDACTED"
}

func (m *FallbackPeerDetails) ObtainFabric() string {
	if m != nil {
		return m.Fabric
	}
	return "REDACTED"
}

func (m *FallbackPeerDetails) ObtainEdition() string {
	if m != nil {
		return m.Edition
	}
	return "REDACTED"
}

func (m *FallbackPeerDetails) ObtainConduits() []byte {
	if m != nil {
		return m.Conduits
	}
	return nil
}

func (m *FallbackPeerDetails) ObtainPseudonym() string {
	if m != nil {
		return m.Pseudonym
	}
	return "REDACTED"
}

func (m *FallbackPeerDetails) ObtainAnother() FallbackPeerDetailsAnother {
	if m != nil {
		return m.Another
	}
	return FallbackPeerDetailsAnother{}
}

type FallbackPeerDetailsAnother struct {
	TransferOrdinal    string `protobuf:"octets,1,opt,name=tx_index,json=txIndex,proto3" json:"transfer_position,omitempty"`
	RemoteLocator string `protobuf:"octets,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"remote_locator,omitempty"`
}

func (m *FallbackPeerDetailsAnother) Restore()         { *m = FallbackPeerDetailsAnother{} }
func (m *FallbackPeerDetailsAnother) Text() string { return proto.CompactTextString(m) }
func (*FallbackPeerDetailsAnother) SchemaArtifact()    {}
func (*FallbackPeerDetailsAnother) Definition() ([]byte, []int) {
	return filedescriptor_c8a29e659aeca578, []int{3}
}
func (m *FallbackPeerDetailsAnother) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *FallbackPeerDetailsAnother) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Fallbacknodedetailsanother.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FallbackPeerDetailsAnother) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Fallbacknodedetailsanother.Merge(m, src)
}
func (m *FallbackPeerDetailsAnother) XXX_Extent() int {
	return m.Extent()
}
func (m *FallbackPeerDetailsAnother) XXX_Dropunfamiliar() {
	xxx_signaldetails_Fallbacknodedetailsanother.DiscardUnknown(m)
}

var xxx_signaldetails_Fallbacknodedetailsanother proto.InternalMessageInfo

func (m *FallbackPeerDetailsAnother) ObtainTransferPosition() string {
	if m != nil {
		return m.TransferOrdinal
	}
	return "REDACTED"
}

func (m *FallbackPeerDetailsAnother) ObtainRemoteLocator() string {
	if m != nil {
		return m.RemoteLocator
	}
	return "REDACTED"
}

func initialize() {
	proto.RegisterType((*NetworkLocator)(nil), "REDACTED")
	proto.RegisterType((*SchemeEdition)(nil), "REDACTED")
	proto.RegisterType((*FallbackPeerDetails)(nil), "REDACTED")
	proto.RegisterType((*FallbackPeerDetailsAnother)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_c8a29e659aeca578) }

var filedescriptor_c8a29e659aeca578 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x3d, 0x8f, 0xda, 0x40,
	0x10, 0xc5, 0xc6, 0x7c, 0xdc, 0x10, 0x8e, 0xcb, 0x0a, 0x45, 0x3e, 0x0a, 0x1b, 0xa1, 0x14, 0x54,
	0xa0, 0x90, 0x2a, 0x5d, 0x42, 0x68, 0x50, 0xa4, 0x8b, 0xb5, 0x8a, 0x52, 0xa4, 0x41, 0xe0, 0x5d,
	0x60, 0x85, 0xd9, 0x5d, 0xad, 0xf7, 0x12, 0xf2, 0x2f, 0xf2, 0xb3, 0xae, 0xbc, 0x32, 0x95, 0x15,
	0x99, 0x32, 0x7f, 0x22, 0xf2, 0xae, 0x2f, 0xc7, 0xa1, 0xeb, 0xe6, 0xcd, 0x9b, 0x99, 0x37, 0xf3,
	0x34, 0xd0, 0xd3, 0x94, 0x13, 0xaa, 0xf6, 0x8c, 0xeb, 0xb1, 0x9c, 0xc8, 0xb1, 0xfe, 0x29, 0x69,
	0x3a, 0x92, 0x4a, 0x68, 0x81, 0x2e, 0x1f, 0xb9, 0x91, 0x9c, 0xc8, 0x5e, 0x77, 0x23, 0x36, 0xc2,
	0x50, 0xe3, 0x22, 0xb2, 0x55, 0x83, 0x08, 0xe0, 0x86, 0xea, 0x0f, 0x84, 0x28, 0x9a, 0xa6, 0xe8,
	0x15, 0xb8, 0x8c, 0xf8, 0x4e, 0xdf, 0x19, 0x5e, 0x4c, 0xeb, 0x79, 0x16, 0xba, 0xf3, 0x19, 0x76,
	0x19, 0x31, 0x79, 0xe9, 0xbb, 0x27, 0xf9, 0x08, 0xbb, 0x4c, 0x22, 0x04, 0x9e, 0x14, 0x4a, 0xfb,
	0xd5, 0xbe, 0x33, 0x6c, 0x63, 0x13, 0x0f, 0xbe, 0x40, 0x27, 0x2a, 0x46, 0xc7, 0x22, 0xf9, 0x4a,
	0x55, 0xca, 0x04, 0x47, 0xd7, 0x50, 0x95, 0x13, 0x69, 0xe6, 0x7a, 0xd3, 0x46, 0x9e, 0x85, 0xd5,
	0x68, 0x12, 0xe1, 0x22, 0x87, 0xba, 0x50, 0x5b, 0x25, 0x22, 0xde, 0x99, 0xe1, 0x1e, 0xb6, 0x00,
	0x5d, 0x41, 0x75, 0x29, 0xa5, 0x19, 0xeb, 0xe1, 0x22, 0x1c, 0xfc, 0x75, 0xa1, 0x33, 0xa3, 0xeb,
	0xe5, 0x6d, 0xa2, 0x6f, 0x04, 0xa1, 0x73, 0xbe, 0x16, 0x28, 0x82, 0x2b, 0x59, 0x2a, 0x2d, 0xbe,
	0x5b, 0x29, 0xa3, 0xd1, 0x9a, 0x84, 0xa3, 0xa7, 0xc7, 0x8f, 0xce, 0x36, 0x9a, 0x7a, 0x77, 0x59,
	0x58, 0xc1, 0x1d, 0x79, 0xb6, 0xe8, 0x3b, 0xe8, 0x10, 0x2b, 0xb2, 0xe0, 0x82, 0xd0, 0x05, 0x23,
	0xe5, 0xd1, 0x2f, 0xf3, 0x2c, 0x6c, 0x9f, 0xea, 0xcf, 0x70, 0x9b, 0x9c, 0x40, 0x82, 0x42, 0x68,
	0x25, 0x2c, 0xd5, 0x94, 0x2f, 0x96, 0x84, 0x28, 0xb3, 0xfa, 0x05, 0x06, 0x9b, 0x2a, 0xec, 0x45,
	0x3e, 0x34, 0x38, 0xd5, 0x3f, 0x84, 0xda, 0xf9, 0x9e, 0x21, 0x1f, 0x60, 0xc1, 0x3c, 0xac, 0x5f,
	0xb3, 0x4c, 0x09, 0x51, 0x0f, 0x9a, 0xf1, 0x76, 0xc9, 0x39, 0x4d, 0x52, 0xbf, 0xde, 0x77, 0x86,
	0x2f, 0xf0, 0x7f, 0x5c, 0x74, 0xed, 0x05, 0x67, 0x3b, 0xaa, 0xfc, 0x86, 0xed, 0x2a, 0x21, 0x7a,
	0x0f, 0x35, 0xa1, 0xb7, 0x54, 0xf9, 0x4d, 0x63, 0xc6, 0xeb, 0x73, 0x33, 0xce, 0x7c, 0xfc, 0x5c,
	0xd4, 0x96, 0x8e, 0xd8, 0xc6, 0xc1, 0x0a, 0xba, 0xcf, 0x15, 0xa1, 0x6b, 0x68, 0xea, 0xc3, 0x82,
	0x71, 0x42, 0x0f, 0xf6, 0x4b, 0x70, 0x43, 0x1f, 0xe6, 0x05, 0x44, 0x63, 0x68, 0x29, 0x19, 0x9b,
	0xe3, 0x69, 0x9a, 0x96, 0xb6, 0x5d, 0xe6, 0x59, 0x08, 0x38, 0xfa, 0x58, 0xfe, 0x17, 0x06, 0x25,
	0xe3, 0x32, 0x9e, 0x7e, 0xba, 0xcb, 0x03, 0xe7, 0x3e, 0x0f, 0x9c, 0x3f, 0x79, 0xe0, 0xfc, 0x3a,
	0x06, 0x95, 0xfb, 0x63, 0x50, 0xf9, 0x7d, 0x0c, 0x2a, 0xdf, 0xde, 0x6c, 0x98, 0xde, 0xde, 0xae,
	0x46, 0xb1, 0xd8, 0x8f, 0x63, 0xb1, 0xa7, 0x7a, 0xb5, 0xd6, 0x8f, 0x81, 0x7d, 0xe1, 0xa7, 0x8f,
	0xbf, 0xaa, 0x9b, 0xec, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0xdb, 0x56, 0x6d, 0x11,
	0x03, 0x00, 0x00,
}

func (m *NetworkLocator) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *NetworkLocator) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *NetworkLocator) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Channel != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Channel))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if len(m.IP) > 0 {
		i -= len(m.IP)
		copy(deltaLocatedAN[i:], m.IP)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.IP)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(deltaLocatedAN[i:], m.ID)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.ID)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SchemeEdition) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SchemeEdition) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SchemeEdition) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.App))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Ledger != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ledger))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.P2P != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.P2P))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *FallbackPeerDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *FallbackPeerDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *FallbackPeerDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.Another.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x42
	if len(m.Pseudonym) > 0 {
		i -= len(m.Pseudonym)
		copy(deltaLocatedAN[i:], m.Pseudonym)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Pseudonym)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	if len(m.Conduits) > 0 {
		i -= len(m.Conduits)
		copy(deltaLocatedAN[i:], m.Conduits)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Conduits)))
		i--
		deltaLocatedAN[i] = 0x32
	}
	if len(m.Edition) > 0 {
		i -= len(m.Edition)
		copy(deltaLocatedAN[i:], m.Edition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Edition)))
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if len(m.Fabric) > 0 {
		i -= len(m.Fabric)
		copy(deltaLocatedAN[i:], m.Fabric)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Fabric)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if len(m.OverhearLocation) > 0 {
		i -= len(m.OverhearLocation)
		copy(deltaLocatedAN[i:], m.OverhearLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.OverhearLocation)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if len(m.FallbackPeerUUID) > 0 {
		i -= len(m.FallbackPeerUUID)
		copy(deltaLocatedAN[i:], m.FallbackPeerUUID)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FallbackPeerUUID)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	{
		extent, err := m.SchemeEdition.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func (m *FallbackPeerDetailsAnother) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *FallbackPeerDetailsAnother) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *FallbackPeerDetailsAnother) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.RemoteLocator) > 0 {
		i -= len(m.RemoteLocator)
		copy(deltaLocatedAN[i:], m.RemoteLocator)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.RemoteLocator)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.TransferOrdinal) > 0 {
		i -= len(m.TransferOrdinal)
		copy(deltaLocatedAN[i:], m.TransferOrdinal)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.TransferOrdinal)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func serializeVariableintKinds(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovKinds(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *NetworkLocator) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Channel != 0 {
		n += 1 + sovKinds(uint64(m.Channel))
	}
	return n
}

func (m *SchemeEdition) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.P2P != 0 {
		n += 1 + sovKinds(uint64(m.P2P))
	}
	if m.Ledger != 0 {
		n += 1 + sovKinds(uint64(m.Ledger))
	}
	if m.App != 0 {
		n += 1 + sovKinds(uint64(m.App))
	}
	return n
}

func (m *FallbackPeerDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SchemeEdition.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FallbackPeerUUID)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.OverhearLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Fabric)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Edition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Conduits)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Pseudonym)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = m.Another.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *FallbackPeerDetailsAnother) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TransferOrdinal)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RemoteLocator)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkLocator) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Channel = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Channel |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemeEdition) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.P2P = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.P2P |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ledger = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Ledger |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.App = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.App |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FallbackPeerDetails) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SchemeEdition.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FallbackPeerUUID = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.OverhearLocation = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Fabric = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Edition = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				octetSize |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if octetSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Conduits = append(m.Conduits[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Conduits == nil {
				m.Conduits = []byte{}
			}
			idxNdExc = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Pseudonym = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Another.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FallbackPeerDetailsAnother) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferOrdinal = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteLocator = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func omitKinds(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunKinds
			}
			if idxNdExc >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= (uint64(b) & 0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		cableKind := int(cable & 0x7)
		switch cableKind {
		case 0:
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return 0, FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return 0, io.ErrUnexpectedEOF
				}
				idxNdExc++
				if deltaLocatedAN[idxNdExc-1] < 0x80 {
					break
				}
			}
		case 1:
			idxNdExc += 8
		case 2:
			var magnitude int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return 0, FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				magnitude |= (int(b) & 0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if magnitude < 0 {
				return 0, FaultUnfitMagnitudeKinds
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingClusterKinds
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeKinds
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeKinds        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunKinds          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingClusterKinds = fmt.Errorf("REDACTED")
)
