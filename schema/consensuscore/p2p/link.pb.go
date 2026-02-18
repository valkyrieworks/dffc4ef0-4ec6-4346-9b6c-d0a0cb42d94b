//
//

package p2p

import (
	fmt "fmt"
	vault "github.com/valkyrieworks/schema/consensuscore/vault"
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

type PackagePing struct {
}

func (m *PackagePing) Restore()         { *m = PackagePing{} }
func (m *PackagePing) String() string { return proto.CompactTextString(m) }
func (*PackagePing) SchemaSignal()    {}
func (*PackagePing) Definition() ([]byte, []int) {
	return filedefinition_hash3, []int{0}
}
func (m *PackagePing) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PackagePing) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Packageping.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PackagePing) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Packageping.Merge(m, src)
}
func (m *PackagePing) XXX_Volume() int {
	return m.Volume()
}
func (m *PackagePing) XXX_Omitunclear() {
	xxx_messagedata_Packageping.DiscardUnknown(m)
}

var xxx_messagedata_Packageping proto.InternalMessageInfo

type PackagePong struct {
}

func (m *PackagePong) Restore()         { *m = PackagePong{} }
func (m *PackagePong) String() string { return proto.CompactTextString(m) }
func (*PackagePong) SchemaSignal()    {}
func (*PackagePong) Definition() ([]byte, []int) {
	return filedefinition_hash3, []int{1}
}
func (m *PackagePong) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PackagePong) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Packagepong.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PackagePong) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Packagepong.Merge(m, src)
}
func (m *PackagePong) XXX_Volume() int {
	return m.Volume()
}
func (m *PackagePong) XXX_Omitunclear() {
	xxx_messagedata_Packagepong.DiscardUnknown(m)
}

var xxx_messagedata_Packagepong proto.InternalMessageInfo

type PackageMessage struct {
	StreamUID int32  `protobuf:"variableint,1,opt,name=channel_id,json=channelId,proto3" json:"conduit_uid,omitempty"`
	EOF       bool   `protobuf:"variableint,2,opt,name=eof,proto3" json:"eof,omitempty"`
	Data      []byte `protobuf:"octets,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PackageMessage) Restore()         { *m = PackageMessage{} }
func (m *PackageMessage) String() string { return proto.CompactTextString(m) }
func (*PackageMessage) SchemaSignal()    {}
func (*PackageMessage) Definition() ([]byte, []int) {
	return filedefinition_hash3, []int{2}
}
func (m *PackageMessage) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PackageMessage) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Messagedata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PackageMessage) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Messagedata.Merge(m, src)
}
func (m *PackageMessage) XXX_Volume() int {
	return m.Volume()
}
func (m *PackageMessage) XXX_Omitunclear() {
	xxx_messagedata_Messagedata.DiscardUnknown(m)
}

var xxx_messagedata_Messagedata proto.InternalMessageInfo

func (m *PackageMessage) FetchConduitUID() int32 {
	if m != nil {
		return m.StreamUID
	}
	return 0
}

func (m *PackageMessage) FetchEOF() bool {
	if m != nil {
		return m.EOF
	}
	return false
}

func (m *PackageMessage) FetchData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Package struct {
	//
	//
	//
	//
	//
	Sum ispackage_Total `protobuf_oneof:"sum"`
}

func (m *Package) Restore()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) SchemaSignal()    {}
func (*Package) Definition() ([]byte, []int) {
	return filedefinition_hash3, []int{3}
}
func (m *Package) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Package) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Package.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Package) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Package.Merge(m, src)
}
func (m *Package) XXX_Volume() int {
	return m.Volume()
}
func (m *Package) XXX_Omitunclear() {
	xxx_messagedata_Package.DiscardUnknown(m)
}

var xxx_messagedata_Package proto.InternalMessageInfo

type ispackage_Total interface {
	ispackage_Total()
	SerializeTo([]byte) (int, error)
	Volume() int
}

type Package_Packageping struct {
	PackagePing *PackagePing `protobuf:"octets,1,opt,name=packet_ping,json=packetPing,proto3,oneof" json:"package_ping,omitempty"`
}
type Package_Packagepong struct {
	PackagePong *PackagePong `protobuf:"octets,2,opt,name=packet_pong,json=packetPong,proto3,oneof" json:"package_pong,omitempty"`
}
type Package_Messagedata struct {
	PackageMessage *PackageMessage `protobuf:"octets,3,opt,name=packet_msg,json=packetMsg,proto3,oneof" json:"package_message,omitempty"`
}

func (*Package_Packageping) ispackage_Total() {}
func (*Package_Packagepong) ispackage_Total() {}
func (*Package_Messagedata) ispackage_Total()  {}

func (m *Package) FetchTotal() ispackage_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Package) FetchPackagePing() *PackagePing {
	if x, ok := m.FetchTotal().(*Package_Packageping); ok {
		return x.PackagePing
	}
	return nil
}

func (m *Package) FetchPackagePong() *PackagePong {
	if x, ok := m.FetchTotal().(*Package_Packagepong); ok {
		return x.PackagePong
	}
	return nil
}

func (m *Package) FetchPackageMessage() *PackageMessage {
	if x, ok := m.FetchTotal().(*Package_Messagedata); ok {
		return x.PackageMessage
	}
	return nil
}

//
func (*Package) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Package_Packageping)(nil),
		(*Package_Packagepong)(nil),
		(*Package_Messagedata)(nil),
	}
}

type AuthSignatureSignal struct {
	PublicKey vault.PublicKey `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_key"`
	Sig    []byte           `protobuf:"octets,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *AuthSignatureSignal) Restore()         { *m = AuthSignatureSignal{} }
func (m *AuthSignatureSignal) String() string { return proto.CompactTextString(m) }
func (*AuthSignatureSignal) SchemaSignal()    {}
func (*AuthSignatureSignal) Definition() ([]byte, []int) {
	return filedefinition_hash3, []int{4}
}
func (m *AuthSignatureSignal) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AuthSignatureSignal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Authsignaturesmessage.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthSignatureSignal) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Authsignaturesmessage.Merge(m, src)
}
func (m *AuthSignatureSignal) XXX_Volume() int {
	return m.Volume()
}
func (m *AuthSignatureSignal) XXX_Omitunclear() {
	xxx_messagedata_Authsignaturesmessage.DiscardUnknown(m)
}

var xxx_messagedata_Authsignaturesmessage proto.InternalMessageInfo

func (m *AuthSignatureSignal) FetchPublicKey() vault.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return vault.PublicKey{}
}

func (m *AuthSignatureSignal) FetchSignature() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*PackagePing)(nil), "REDACTED")
	proto.RegisterType((*PackagePong)(nil), "REDACTED")
	proto.RegisterType((*PackageMessage)(nil), "REDACTED")
	proto.RegisterType((*Package)(nil), "REDACTED")
	proto.RegisterType((*AuthSignatureSignal)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash3) }

var filedefinition_hash3 = []byte{
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

func (m *PackagePing) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PackagePing) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PackagePing) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PackagePong) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PackagePong) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PackagePong) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PackageMessage) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PackageMessage) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PackageMessage) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVariableintLink(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.EOF {
		i--
		if m.EOF {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.StreamUID != 0 {
		i = encodeVariableintLink(dAtA, i, uint64(m.StreamUID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Package) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Package) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Package) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			volume := m.Sum.Volume()
			i -= volume
			if _, err := m.Sum.SerializeTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Package_Packageping) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Package_Packageping) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PackagePing != nil {
		{
			volume, err := m.PackagePing.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintLink(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Package_Packagepong) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Package_Packagepong) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PackagePong != nil {
		{
			volume, err := m.PackagePong.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintLink(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Package_Messagedata) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Package_Messagedata) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PackageMessage != nil {
		{
			volume, err := m.PackageMessage.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintLink(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *AuthSignatureSignal) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthSignatureSignal) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AuthSignatureSignal) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVariableintLink(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x12
	}
	{
		volume, err := m.PublicKey.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = encodeVariableintLink(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVariableintLink(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovLink(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *PackagePing) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PackagePong) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PackageMessage) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StreamUID != 0 {
		n += 1 + sovLink(uint64(m.StreamUID))
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

func (m *Package) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Volume()
	}
	return n
}

func (m *Package_Packageping) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PackagePing != nil {
		l = m.PackagePing.Volume()
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}
func (m *Package_Packagepong) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PackagePong != nil {
		l = m.PackagePong.Volume()
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}
func (m *Package_Messagedata) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PackageMessage != nil {
		l = m.PackageMessage.Volume()
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}
func (m *AuthSignatureSignal) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PublicKey.Volume()
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
func (m *PackagePing) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadLink
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
		default:
			idxNdEx = preOrdinal
			skippy, err := omitLink(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentLink
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
func (m *PackagePong) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadLink
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
		default:
			idxNdEx = preOrdinal
			skippy, err := omitLink(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentLink
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
func (m *PackageMessage) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadLink
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
			m.StreamUID = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLink
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.StreamUID |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var v int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLink
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				v |= int(b&0x7F) << displace
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
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLink
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
				return ErrCorruptExtentLink
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitLink(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentLink
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
func (m *Package) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadLink
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
					return ErrIntegerOverloadLink
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
				return ErrCorruptExtentLink
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PackagePing{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Package_Packageping{v}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLink
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
				return ErrCorruptExtentLink
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PackagePong{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Package_Packagepong{v}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLink
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
				return ErrCorruptExtentLink
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PackageMessage{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Package_Messagedata{v}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitLink(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentLink
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
func (m *AuthSignatureSignal) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadLink
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
					return ErrIntegerOverloadLink
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
				return ErrCorruptExtentLink
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLink
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
				return ErrCorruptExtentLink
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLink
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitLink(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentLink
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
func omitLink(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadLink
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
					return 0, ErrIntegerOverloadLink
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
					return 0, ErrIntegerOverloadLink
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
				return 0, ErrCorruptExtentLink
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterLink
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentLink
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentLink        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadLink          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterLink = fmt.Errorf("REDACTED")
)
