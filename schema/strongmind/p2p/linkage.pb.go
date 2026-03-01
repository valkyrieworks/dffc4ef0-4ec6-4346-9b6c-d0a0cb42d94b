//
//

package p2p

import (
	fmt "fmt"
	security "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
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

type PacketPing struct {
}

func (m *PacketPing) Restore()         { *m = PacketPing{} }
func (m *PacketPing) Text() string { return proto.CompactTextString(m) }
func (*PacketPing) SchemaArtifact()    {}
func (*PacketPing) Definition() ([]byte, []int) {
	return filedescriptor_22474b5527c8fa9f, []int{0}
}
func (m *PacketPing) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PacketPing) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Pingpacket.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketPing) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Pingpacket.Merge(m, src)
}
func (m *PacketPing) XXX_Extent() int {
	return m.Extent()
}
func (m *PacketPing) XXX_Dropunfamiliar() {
	xxx_signaldetails_Pingpacket.DiscardUnknown(m)
}

var xxx_signaldetails_Pingpacket proto.InternalMessageInfo

type PacketPong struct {
}

func (m *PacketPong) Restore()         { *m = PacketPong{} }
func (m *PacketPong) Text() string { return proto.CompactTextString(m) }
func (*PacketPong) SchemaArtifact()    {}
func (*PacketPong) Definition() ([]byte, []int) {
	return filedescriptor_22474b5527c8fa9f, []int{1}
}
func (m *PacketPong) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PacketPong) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Pongpacket.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketPong) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Pongpacket.Merge(m, src)
}
func (m *PacketPong) XXX_Extent() int {
	return m.Extent()
}
func (m *PacketPong) XXX_Dropunfamiliar() {
	xxx_signaldetails_Pongpacket.DiscardUnknown(m)
}

var xxx_signaldetails_Pongpacket proto.InternalMessageInfo

type PacketSignal struct {
	ConduitUUID int32  `protobuf:"variableint,1,opt,name=channel_id,json=channelId,proto3" json:"conduit_uuid,omitempty"`
	EOF       bool   `protobuf:"variableint,2,opt,name=eof,proto3" json:"eof,omitempty"`
	Data      []byte `protobuf:"octets,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PacketSignal) Restore()         { *m = PacketSignal{} }
func (m *PacketSignal) Text() string { return proto.CompactTextString(m) }
func (*PacketSignal) SchemaArtifact()    {}
func (*PacketSignal) Definition() ([]byte, []int) {
	return filedescriptor_22474b5527c8fa9f, []int{2}
}
func (m *PacketSignal) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PacketSignal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Packetsignal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketSignal) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Packetsignal.Merge(m, src)
}
func (m *PacketSignal) XXX_Extent() int {
	return m.Extent()
}
func (m *PacketSignal) XXX_Dropunfamiliar() {
	xxx_signaldetails_Packetsignal.DiscardUnknown(m)
}

var xxx_signaldetails_Packetsignal proto.InternalMessageInfo

func (m *PacketSignal) ObtainConduitUUID() int32 {
	if m != nil {
		return m.ConduitUUID
	}
	return 0
}

func (m *PacketSignal) ObtainEOF() bool {
	if m != nil {
		return m.EOF
	}
	return false
}

func (m *PacketSignal) ObtainData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Packet struct {
	//
	//
	//
	//
	//
	Sum ispacket_Total `protobuf_oneof:"sum"`
}

func (m *Packet) Restore()         { *m = Packet{} }
func (m *Packet) Text() string { return proto.CompactTextString(m) }
func (*Packet) SchemaArtifact()    {}
func (*Packet) Definition() ([]byte, []int) {
	return filedescriptor_22474b5527c8fa9f, []int{3}
}
func (m *Packet) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Packet) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Packet.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Packet.Merge(m, src)
}
func (m *Packet) XXX_Extent() int {
	return m.Extent()
}
func (m *Packet) XXX_Dropunfamiliar() {
	xxx_signaldetails_Packet.DiscardUnknown(m)
}

var xxx_signaldetails_Packet proto.InternalMessageInfo

type ispacket_Total interface {
	ispacket_Total()
	SerializeToward([]byte) (int, error)
	Extent() int
}

type Packet_Pingpacket struct {
	PacketPing *PacketPing `protobuf:"octets,1,opt,name=packet_ping,json=packetPing,proto3,oneof" json:"packet_ping,omitempty"`
}
type Packet_Pongpacket struct {
	PacketPong *PacketPong `protobuf:"octets,2,opt,name=packet_pong,json=packetPong,proto3,oneof" json:"packet_pong,omitempty"`
}
type Packet_Packetsignal struct {
	PacketSignal *PacketSignal `protobuf:"octets,3,opt,name=packet_msg,json=packetMsg,proto3,oneof" json:"packet_signal,omitempty"`
}

func (*Packet_Pingpacket) ispacket_Total() {}
func (*Packet_Pongpacket) ispacket_Total() {}
func (*Packet_Packetsignal) ispacket_Total()  {}

func (m *Packet) ObtainTotal() ispacket_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Packet) ObtainPacketPing() *PacketPing {
	if x, ok := m.ObtainTotal().(*Packet_Pingpacket); ok {
		return x.PacketPing
	}
	return nil
}

func (m *Packet) ObtainPacketPong() *PacketPong {
	if x, ok := m.ObtainTotal().(*Packet_Pongpacket); ok {
		return x.PacketPong
	}
	return nil
}

func (m *Packet) ObtainPacketSignal() *PacketSignal {
	if x, ok := m.ObtainTotal().(*Packet_Packetsignal); ok {
		return x.PacketSignal
	}
	return nil
}

//
func (*Packet) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Packet_Pingpacket)(nil),
		(*Packet_Pongpacket)(nil),
		(*Packet_Packetsignal)(nil),
	}
}

type AuthSignatureArtifact struct {
	PublicToken security.CommonToken `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_token"`
	Sig    []byte           `protobuf:"octets,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *AuthSignatureArtifact) Restore()         { *m = AuthSignatureArtifact{} }
func (m *AuthSignatureArtifact) Text() string { return proto.CompactTextString(m) }
func (*AuthSignatureArtifact) SchemaArtifact()    {}
func (*AuthSignatureArtifact) Definition() ([]byte, []int) {
	return filedescriptor_22474b5527c8fa9f, []int{4}
}
func (m *AuthSignatureArtifact) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AuthSignatureArtifact) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Authsignatureartifact.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthSignatureArtifact) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Authsignatureartifact.Merge(m, src)
}
func (m *AuthSignatureArtifact) XXX_Extent() int {
	return m.Extent()
}
func (m *AuthSignatureArtifact) XXX_Dropunfamiliar() {
	xxx_signaldetails_Authsignatureartifact.DiscardUnknown(m)
}

var xxx_signaldetails_Authsignatureartifact proto.InternalMessageInfo

func (m *AuthSignatureArtifact) ObtainPublicToken() security.CommonToken {
	if m != nil {
		return m.PublicToken
	}
	return security.CommonToken{}
}

func (m *AuthSignatureArtifact) ObtainSignature() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func initialize() {
	proto.RegisterType((*PacketPing)(nil), "REDACTED")
	proto.RegisterType((*PacketPong)(nil), "REDACTED")
	proto.RegisterType((*PacketSignal)(nil), "REDACTED")
	proto.RegisterType((*Packet)(nil), "REDACTED")
	proto.RegisterType((*AuthSignatureArtifact)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_22474b5527c8fa9f) }

var filedescriptor_22474b5527c8fa9f = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x8d, 0xc9, 0x6e, 0x97, 0x4e, 0xcb, 0x0a, 0x59, 0x1c, 0xda, 0x6a, 0x95, 0x56, 0x3d, 0xf5,
	0x80, 0x12, 0x11, 0x6e, 0x20, 0x0e, 0x84, 0x0f, 0xb1, 0xaa, 0x2a, 0xaa, 0x70, 0xe3, 0x12, 0xe5,
	0xc3, 0xeb, 0x58, 0xdd, 0xd8, 0x56, 0xed, 0x1c, 0xf2, 0x2f, 0xf8, 0x59, 0xcb, 0xad, 0x47, 0x4e,
	0x15, 0x4a, 0xff, 0x08, 0x4a, 0x5c, 0x68, 0x2a, 0xb1, 0xb7, 0xf7, 0x66, 0xfc, 0x66, 0xde, 0x93,
	0x07, 0xc6, 0x9a, 0xf0, 0x8c, 0x6c, 0x0b, 0xc6, 0xb5, 0x27, 0x7d, 0xe9, 0xa5, 0x82, 0x73, 0x57,
	0x6e, 0x85, 0x16, 0xf8, 0xfa, 0xd4, 0x72, 0xa5, 0x2f, 0x27, 0x2f, 0xa8, 0xa0, 0xa2, 0x6d, 0x79,
	0x0d, 0x32, 0xaf, 0x26, 0x37, 0x9d, 0x01, 0xe9, 0xb6, 0x92, 0x5a, 0x78, 0x1b, 0x52, 0x29, 0xd3,
	0x9d, 0x0f, 0x01, 0xd6, 0x71, 0xba, 0x21, 0x7a, 0xcd, 0x38, 0xed, 0x30, 0xc1, 0xe9, 0x3c, 0x87,
	0xbe, 0x61, 0x2b, 0x45, 0xf1, 0x4b, 0x80, 0x34, 0x8f, 0x39, 0x27, 0xf7, 0x11, 0xcb, 0x46, 0x68,
	0x86, 0x16, 0x97, 0xc1, 0xb3, 0x7a, 0x3f, 0xed, 0x7f, 0x30, 0xd5, 0xdb, 0x8f, 0x61, 0xff, 0xf8,
	0xe0, 0x36, 0xc3, 0x63, 0xb0, 0x89, 0xb8, 0x1b, 0x3d, 0x99, 0xa1, 0xc5, 0xd3, 0xe0, 0xaa, 0xde,
	0x4f, 0xed, 0x4f, 0x5f, 0x3f, 0x87, 0x4d, 0x0d, 0x63, 0xb8, 0xc8, 0x62, 0x1d, 0x8f, 0xec, 0x19,
	0x5a, 0x0c, 0xc3, 0x16, 0xcf, 0x7f, 0x22, 0xe8, 0x99, 0x55, 0xf8, 0x1d, 0x0c, 0x64, 0x8b, 0x22,
	0xc9, 0x38, 0x6d, 0x17, 0x0d, 0xfc, 0x89, 0x7b, 0x1e, 0xd5, 0x3d, 0x79, 0xfe, 0x62, 0x85, 0x20,
	0xff, 0xb1, 0xae, 0x5c, 0x70, 0xda, 0x1a, 0x78, 0x5c, 0x2e, 0xce, 0xe4, 0x82, 0x53, 0xfc, 0x06,
	0x8e, 0x2c, 0x2a, 0x14, 0x6d, 0x2d, 0x0e, 0xfc, 0xf1, 0xff, 0xd5, 0x2b, 0xd5, 0x88, 0xfb, 0xf2,
	0x2f, 0x09, 0x2e, 0xc1, 0x56, 0x65, 0x31, 0x8f, 0xe0, 0xfa, 0x7d, 0xa9, 0xf3, 0x6f, 0x8c, 0xae,
	0x88, 0x52, 0x31, 0x25, 0xf8, 0x2d, 0x5c, 0xc9, 0x32, 0x89, 0x36, 0xa4, 0x3a, 0xc6, 0xb9, 0xe9,
	0x4e, 0x34, 0x7f, 0xe2, 0xae, 0xcb, 0xe4, 0x9e, 0xa5, 0x4b, 0x52, 0x05, 0x17, 0x0f, 0xfb, 0xa9,
	0x15, 0xf6, 0x64, 0x99, 0x2c, 0x49, 0x85, 0x9f, 0x83, 0xad, 0x98, 0x09, 0x32, 0x0c, 0x1b, 0x18,
	0x2c, 0x1f, 0x6a, 0x07, 0xed, 0x6a, 0x07, 0xfd, 0xae, 0x1d, 0xf4, 0xe3, 0xe0, 0x58, 0xbb, 0x83,
	0x63, 0xfd, 0x3a, 0x38, 0xd6, 0xf7, 0x57, 0x94, 0xe9, 0xbc, 0x4c, 0xdc, 0x54, 0x14, 0x5e, 0x2a,
	0x0a, 0xa2, 0x93, 0x3b, 0x7d, 0x02, 0xe6, 0x32, 0xce, 0xcf, 0x29, 0xe9, 0xb5, 0xd5, 0xd7, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xf9, 0x75, 0xae, 0x67, 0x02, 0x00, 0x00,
}

func (m *PacketPing) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PacketPing) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PacketPing) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *PacketPong) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PacketPong) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PacketPong) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *PacketSignal) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PacketSignal) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PacketSignal) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(deltaLocatedAN[i:], m.Data)
		i = encodeVariableintLink(deltaLocatedAN, i, uint64(len(m.Data)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.EOF {
		i--
		if m.EOF {
			deltaLocatedAN[i] = 1
		} else {
			deltaLocatedAN[i] = 0
		}
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.ConduitUUID != 0 {
		i = encodeVariableintLink(deltaLocatedAN, i, uint64(m.ConduitUUID))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Packet) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Packet) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Packet) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			extent := m.Sum.Extent()
			i -= extent
			if _, err := m.Sum.SerializeToward(deltaLocatedAN[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Packet_Pingpacket) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Packet_Pingpacket) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PacketPing != nil {
		{
			extent, err := m.PacketPing.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintLink(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Packet_Pongpacket) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Packet_Pongpacket) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PacketPong != nil {
		{
			extent, err := m.PacketPong.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintLink(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Packet_Packetsignal) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Packet_Packetsignal) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PacketSignal != nil {
		{
			extent, err := m.PacketSignal.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintLink(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *AuthSignatureArtifact) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AuthSignatureArtifact) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AuthSignatureArtifact) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(deltaLocatedAN[i:], m.Sig)
		i = encodeVariableintLink(deltaLocatedAN, i, uint64(len(m.Sig)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	{
		extent, err := m.PublicToken.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = encodeVariableintLink(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintLink(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovLink(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *PacketPing) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PacketPong) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PacketSignal) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConduitUUID != 0 {
		n += 1 + sovLink(uint64(m.ConduitUUID))
	}
	if m.EOF {
		n += 2
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}

func (m *Packet) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Extent()
	}
	return n
}

func (m *Packet_Pingpacket) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketPing != nil {
		l = m.PacketPing.Extent()
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}
func (m *Packet_Pongpacket) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketPong != nil {
		l = m.PacketPong.Extent()
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}
func (m *Packet_Packetsignal) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketSignal != nil {
		l = m.PacketSignal.Extent()
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}
func (m *AuthSignatureArtifact) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PublicToken.Extent()
	n += 1 + l + sovLink(uint64(l))
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}

func sovLink(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLink(x uint64) (n int) {
	return sovLink(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PacketPing) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunLink
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
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitLink(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeLink
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
func (m *PacketPong) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunLink
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
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitLink(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeLink
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
func (m *PacketSignal) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunLink
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
			m.ConduitUUID = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLink
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.ConduitUUID |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var v int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLink
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				v |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			m.EOF = bool(v != 0)
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLink
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
				return FaultUnfitMagnitudeLink
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitLink(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeLink
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
func (m *Packet) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunLink
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
					return FaultIntegerOverrunLink
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
				return FaultUnfitMagnitudeLink
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PacketPing{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Packet_Pingpacket{v}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLink
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
				return FaultUnfitMagnitudeLink
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PacketPong{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Packet_Pongpacket{v}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLink
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
				return FaultUnfitMagnitudeLink
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PacketSignal{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Packet_Packetsignal{v}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitLink(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeLink
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
func (m *AuthSignatureArtifact) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunLink
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
					return FaultIntegerOverrunLink
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
				return FaultUnfitMagnitudeLink
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicToken.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLink
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
				return FaultUnfitMagnitudeLink
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitLink(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeLink
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
func omitLink(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunLink
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
					return 0, FaultIntegerOverrunLink
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
					return 0, FaultIntegerOverrunLink
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
				return 0, FaultUnfitMagnitudeLink
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionLink
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeLink
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeLink        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunLink          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionLink = fmt.Errorf("REDACTED")
)
