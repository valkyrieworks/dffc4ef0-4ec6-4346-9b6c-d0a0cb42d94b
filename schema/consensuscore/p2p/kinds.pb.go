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

type NetLocation struct {
	ID   string `protobuf:"octets,1,opt,name=id,proto3" json:"id,omitempty"`
	IP   string `protobuf:"octets,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port uint32 `protobuf:"variableint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *NetLocation) Restore()         { *m = NetLocation{} }
func (m *NetLocation) String() string { return proto.CompactTextString(m) }
func (*NetLocation) SchemaSignal()    {}
func (*NetLocation) Definition() ([]byte, []int) {
	return filedefinition_hash8, []int{0}
}
func (m *NetLocation) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *NetLocation) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Netlocation.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetLocation) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Netlocation.Merge(m, src)
}
func (m *NetLocation) XXX_Volume() int {
	return m.Volume()
}
func (m *NetLocation) XXX_Omitunclear() {
	xxx_messagedata_Netlocation.DiscardUnknown(m)
}

var xxx_messagedata_Netlocation proto.InternalMessageInfo

func (m *NetLocation) FetchUID() string {
	if m != nil {
		return m.ID
	}
	return "REDACTED"
}

func (m *NetLocation) FetchIP() string {
	if m != nil {
		return m.IP
	}
	return "REDACTED"
}

func (m *NetLocation) FetchPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ProtocolRelease struct {
	P2P   uint64 `protobuf:"variableint,1,opt,name=p2p,proto3" json:"p2p,omitempty"`
	Ledger uint64 `protobuf:"variableint,2,opt,name=block,proto3" json:"ledger,omitempty"`
	App   uint64 `protobuf:"variableint,3,opt,name=app,proto3" json:"app,omitempty"`
}

func (m *ProtocolRelease) Restore()         { *m = ProtocolRelease{} }
func (m *ProtocolRelease) String() string { return proto.CompactTextString(m) }
func (*ProtocolRelease) SchemaSignal()    {}
func (*ProtocolRelease) Definition() ([]byte, []int) {
	return filedefinition_hash8, []int{1}
}
func (m *ProtocolRelease) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ProtocolRelease) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Protocolrelease.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolRelease) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Protocolrelease.Merge(m, src)
}
func (m *ProtocolRelease) XXX_Volume() int {
	return m.Volume()
}
func (m *ProtocolRelease) XXX_Omitunclear() {
	xxx_messagedata_Protocolrelease.DiscardUnknown(m)
}

var xxx_messagedata_Protocolrelease proto.InternalMessageInfo

func (m *ProtocolRelease) FetchP2P() uint64 {
	if m != nil {
		return m.P2P
	}
	return 0
}

func (m *ProtocolRelease) FetchLedger() uint64 {
	if m != nil {
		return m.Ledger
	}
	return 0
}

func (m *ProtocolRelease) FetchApplication() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type StandardMemberDetails struct {
	ProtocolRelease ProtocolRelease      `protobuf:"octets,1,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_release"`
	StandardMemberUID   string               `protobuf:"octets,2,opt,name=default_node_id,json=defaultNodeId,proto3" json:"standard_member_uid,omitempty"`
	ObserveAddress      string               `protobuf:"octets,3,opt,name=listen_addr,json=listenAddr,proto3" json:"observe_address,omitempty"`
	Fabric         string               `protobuf:"octets,4,opt,name=network,proto3" json:"fabric,omitempty"`
	Release         string               `protobuf:"octets,5,opt,name=version,proto3" json:"release,omitempty"`
	Streams        []byte               `protobuf:"octets,6,opt,name=channels,proto3" json:"streams,omitempty"`
	Moniker         string               `protobuf:"octets,7,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Another           StandardMemberDetailsAnother `protobuf:"octets,8,opt,name=other,proto3" json:"another"`
}

func (m *StandardMemberDetails) Restore()         { *m = StandardMemberDetails{} }
func (m *StandardMemberDetails) String() string { return proto.CompactTextString(m) }
func (*StandardMemberDetails) SchemaSignal()    {}
func (*StandardMemberDetails) Definition() ([]byte, []int) {
	return filedefinition_hash8, []int{2}
}
func (m *StandardMemberDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StandardMemberDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Standardmemberdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardMemberDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Standardmemberdetails.Merge(m, src)
}
func (m *StandardMemberDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *StandardMemberDetails) XXX_Omitunclear() {
	xxx_messagedata_Standardmemberdetails.DiscardUnknown(m)
}

var xxx_messagedata_Standardmemberdetails proto.InternalMessageInfo

func (m *StandardMemberDetails) FetchProtocolRelease() ProtocolRelease {
	if m != nil {
		return m.ProtocolRelease
	}
	return ProtocolRelease{}
}

func (m *StandardMemberDetails) FetchStandardMemberUID() string {
	if m != nil {
		return m.StandardMemberUID
	}
	return "REDACTED"
}

func (m *StandardMemberDetails) FetchObserveAddress() string {
	if m != nil {
		return m.ObserveAddress
	}
	return "REDACTED"
}

func (m *StandardMemberDetails) FetchFabric() string {
	if m != nil {
		return m.Fabric
	}
	return "REDACTED"
}

func (m *StandardMemberDetails) FetchRelease() string {
	if m != nil {
		return m.Release
	}
	return "REDACTED"
}

func (m *StandardMemberDetails) FetchStreams() []byte {
	if m != nil {
		return m.Streams
	}
	return nil
}

func (m *StandardMemberDetails) FetchMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return "REDACTED"
}

func (m *StandardMemberDetails) FetchAnother() StandardMemberDetailsAnother {
	if m != nil {
		return m.Another
	}
	return StandardMemberDetailsAnother{}
}

type StandardMemberDetailsAnother struct {
	TransOrdinal    string `protobuf:"octets,1,opt,name=tx_index,json=txIndex,proto3" json:"transfer_ordinal,omitempty"`
	RPCLocation string `protobuf:"octets,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_location,omitempty"`
}

func (m *StandardMemberDetailsAnother) Restore()         { *m = StandardMemberDetailsAnother{} }
func (m *StandardMemberDetailsAnother) String() string { return proto.CompactTextString(m) }
func (*StandardMemberDetailsAnother) SchemaSignal()    {}
func (*StandardMemberDetailsAnother) Definition() ([]byte, []int) {
	return filedefinition_hash8, []int{3}
}
func (m *StandardMemberDetailsAnother) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StandardMemberDetailsAnother) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Standardmemberdetailsanother.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardMemberDetailsAnother) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Standardmemberdetailsanother.Merge(m, src)
}
func (m *StandardMemberDetailsAnother) XXX_Volume() int {
	return m.Volume()
}
func (m *StandardMemberDetailsAnother) XXX_Omitunclear() {
	xxx_messagedata_Standardmemberdetailsanother.DiscardUnknown(m)
}

var xxx_messagedata_Standardmemberdetailsanother proto.InternalMessageInfo

func (m *StandardMemberDetailsAnother) FetchTransferOrdinal() string {
	if m != nil {
		return m.TransOrdinal
	}
	return "REDACTED"
}

func (m *StandardMemberDetailsAnother) FetchRPCLocation() string {
	if m != nil {
		return m.RPCLocation
	}
	return "REDACTED"
}

func init() {
	proto.RegisterType((*NetLocation)(nil), "REDACTED")
	proto.RegisterType((*ProtocolRelease)(nil), "REDACTED")
	proto.RegisterType((*StandardMemberDetails)(nil), "REDACTED")
	proto.RegisterType((*StandardMemberDetailsAnother)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash8) }

var filedefinition_hash8 = []byte{
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

func (m *NetLocation) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetLocation) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *NetLocation) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Port != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x18
	}
	if len(m.IP) > 0 {
		i -= len(m.IP)
		copy(dAtA[i:], m.IP)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.IP)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolRelease) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolRelease) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ProtocolRelease) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.App))
		i--
		dAtA[i] = 0x18
	}
	if m.Ledger != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ledger))
		i--
		dAtA[i] = 0x10
	}
	if m.P2P != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.P2P))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StandardMemberDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardMemberDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StandardMemberDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.Another.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Streams) > 0 {
		i -= len(m.Streams)
		copy(dAtA[i:], m.Streams)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Streams)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Release) > 0 {
		i -= len(m.Release)
		copy(dAtA[i:], m.Release)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Release)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Fabric) > 0 {
		i -= len(m.Fabric)
		copy(dAtA[i:], m.Fabric)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Fabric)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ObserveAddress) > 0 {
		i -= len(m.ObserveAddress)
		copy(dAtA[i:], m.ObserveAddress)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ObserveAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StandardMemberUID) > 0 {
		i -= len(m.StandardMemberUID)
		copy(dAtA[i:], m.StandardMemberUID)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.StandardMemberUID)))
		i--
		dAtA[i] = 0x12
	}
	{
		volume, err := m.ProtocolRelease.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *StandardMemberDetailsAnother) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardMemberDetailsAnother) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StandardMemberDetailsAnother) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RPCLocation) > 0 {
		i -= len(m.RPCLocation)
		copy(dAtA[i:], m.RPCLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RPCLocation)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TransOrdinal) > 0 {
		i -= len(m.TransOrdinal)
		copy(dAtA[i:], m.TransOrdinal)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.TransOrdinal)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func formatVariableintKinds(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovKinds(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *NetLocation) Volume() (n int) {
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
	if m.Port != 0 {
		n += 1 + sovKinds(uint64(m.Port))
	}
	return n
}

func (m *ProtocolRelease) Volume() (n int) {
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

func (m *StandardMemberDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ProtocolRelease.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.StandardMemberUID)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.ObserveAddress)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Fabric)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Release)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Streams)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = m.Another.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *StandardMemberDetailsAnother) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TransOrdinal)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RPCLocation)
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
func (m *NetLocation) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Port = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Port |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProtocolRelease) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.P2P = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.P2P |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ledger = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Ledger |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.App = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.App |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StandardMemberDetails) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProtocolRelease.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.StandardMemberUID = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserveAddress = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Fabric = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Release = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				octetSize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if octetSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Streams = append(m.Streams[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Streams == nil {
				m.Streams = []byte{}
			}
			idxNdEx = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Another.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StandardMemberDetailsAnother) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.TransOrdinal = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.RPCLocation = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func omitKinds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadKinds
			}
			if idxNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= (uint64(b) & 0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		cableKind := int(cable & 0x7)
		switch cableKind {
		case 0:
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return 0, ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				idxNdEx++
				if dAtA[idxNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			idxNdEx += 8
		case 2:
			var extent int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return 0, ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				extent |= (int(b) & 0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if extent < 0 {
				return 0, ErrCorruptExtentKinds
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterKinds
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentKinds
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentKinds        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadKinds          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterKinds = fmt.Errorf("REDACTED")
)
