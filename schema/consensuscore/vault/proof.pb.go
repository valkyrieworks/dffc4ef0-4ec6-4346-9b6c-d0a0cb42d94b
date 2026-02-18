//
//

package vault

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

type Attestation struct {
	Sum    int64    `protobuf:"variableint,1,opt,name=total,proto3" json:"sum,omitempty"`
	Ordinal    int64    `protobuf:"variableint,2,opt,name=index,proto3" json:"ordinal,omitempty"`
	NodeDigest []byte   `protobuf:"octets,3,opt,name=leaf_hash,json=leafHash,proto3" json:"element_digest,omitempty"`
	Kin    [][]byte `protobuf:"octets,4,rep,name=aunts,proto3" json:"kin,omitempty"`
}

func (m *Attestation) Restore()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) SchemaSignal()    {}
func (*Attestation) Definition() ([]byte, []int) {
	return filedefinition_6b60b6ba2ab5b856, []int{0}
}
func (m *Attestation) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Attestation) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Evidence.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attestation) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Evidence.Merge(m, src)
}
func (m *Attestation) XXX_Volume() int {
	return m.Volume()
}
func (m *Attestation) XXX_Omitunclear() {
	xxx_messagedata_Evidence.DiscardUnknown(m)
}

var xxx_messagedata_Evidence proto.InternalMessageInfo

func (m *Attestation) FetchSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *Attestation) FetchOrdinal() int64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *Attestation) FetchElementDigest() []byte {
	if m != nil {
		return m.NodeDigest
	}
	return nil
}

func (m *Attestation) FetchKin() [][]byte {
	if m != nil {
		return m.Kin
	}
	return nil
}

type ItemAct struct {
	//
	Key []byte `protobuf:"octets,1,opt,name=key,proto3" json:"key,omitempty"`
	//
	Attestation *Attestation `protobuf:"octets,2,opt,name=proof,proto3" json:"evidence,omitempty"`
}

func (m *ItemAct) Restore()         { *m = ItemAct{} }
func (m *ItemAct) String() string { return proto.CompactTextString(m) }
func (*ItemAct) SchemaSignal()    {}
func (*ItemAct) Definition() ([]byte, []int) {
	return filedefinition_6b60b6ba2ab5b856, []int{1}
}
func (m *ItemAct) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ItemAct) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Valueaction.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ItemAct) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Valueaction.Merge(m, src)
}
func (m *ItemAct) XXX_Volume() int {
	return m.Volume()
}
func (m *ItemAct) XXX_Omitunclear() {
	xxx_messagedata_Valueaction.DiscardUnknown(m)
}

var xxx_messagedata_Valueaction proto.InternalMessageInfo

func (m *ItemAct) FetchKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ItemAct) FetchEvidence() *Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

type DominoAct struct {
	Key    string `protobuf:"octets,1,opt,name=key,proto3" json:"key,omitempty"`
	Influx  string `protobuf:"octets,2,opt,name=input,proto3" json:"influx,omitempty"`
	Result string `protobuf:"octets,3,opt,name=output,proto3" json:"result,omitempty"`
}

func (m *DominoAct) Restore()         { *m = DominoAct{} }
func (m *DominoAct) String() string { return proto.CompactTextString(m) }
func (*DominoAct) SchemaSignal()    {}
func (*DominoAct) Definition() ([]byte, []int) {
	return filedefinition_6b60b6ba2ab5b856, []int{2}
}
func (m *DominoAct) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *DominoAct) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Dominoaction.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DominoAct) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Dominoaction.Merge(m, src)
}
func (m *DominoAct) XXX_Volume() int {
	return m.Volume()
}
func (m *DominoAct) XXX_Omitunclear() {
	xxx_messagedata_Dominoaction.DiscardUnknown(m)
}

var xxx_messagedata_Dominoaction proto.InternalMessageInfo

func (m *DominoAct) FetchKey() string {
	if m != nil {
		return m.Key
	}
	return "REDACTED"
}

func (m *DominoAct) FetchInflux() string {
	if m != nil {
		return m.Influx
	}
	return "REDACTED"
}

func (m *DominoAct) FetchResult() string {
	if m != nil {
		return m.Result
	}
	return "REDACTED"
}

//
//
//
type EvidenceAct struct {
	Kind string `protobuf:"octets,1,opt,name=type,proto3" json:"kind,omitempty"`
	Key  []byte `protobuf:"octets,2,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"octets,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EvidenceAct) Restore()         { *m = EvidenceAct{} }
func (m *EvidenceAct) String() string { return proto.CompactTextString(m) }
func (*EvidenceAct) SchemaSignal()    {}
func (*EvidenceAct) Definition() ([]byte, []int) {
	return filedefinition_6b60b6ba2ab5b856, []int{3}
}
func (m *EvidenceAct) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *EvidenceAct) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Evidenceaction.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EvidenceAct) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Evidenceaction.Merge(m, src)
}
func (m *EvidenceAct) XXX_Volume() int {
	return m.Volume()
}
func (m *EvidenceAct) XXX_Omitunclear() {
	xxx_messagedata_Evidenceaction.DiscardUnknown(m)
}

var xxx_messagedata_Evidenceaction proto.InternalMessageInfo

func (m *EvidenceAct) FetchKind() string {
	if m != nil {
		return m.Kind
	}
	return "REDACTED"
}

func (m *EvidenceAct) FetchKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EvidenceAct) FetchData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

//
type EvidenceActions struct {
	Ops []EvidenceAct `protobuf:"octets,1,rep,name=ops,proto3" json:"ops"`
}

func (m *EvidenceActions) Restore()         { *m = EvidenceActions{} }
func (m *EvidenceActions) String() string { return proto.CompactTextString(m) }
func (*EvidenceActions) SchemaSignal()    {}
func (*EvidenceActions) Definition() ([]byte, []int) {
	return filedefinition_6b60b6ba2ab5b856, []int{4}
}
func (m *EvidenceActions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *EvidenceActions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Evidenceactions.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EvidenceActions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Evidenceactions.Merge(m, src)
}
func (m *EvidenceActions) XXX_Volume() int {
	return m.Volume()
}
func (m *EvidenceActions) XXX_Omitunclear() {
	xxx_messagedata_Evidenceactions.DiscardUnknown(m)
}

var xxx_messagedata_Evidenceactions proto.InternalMessageInfo

func (m *EvidenceActions) FetchActions() []EvidenceAct {
	if m != nil {
		return m.Ops
	}
	return nil
}

func init() {
	proto.RegisterType((*Attestation)(nil), "REDACTED")
	proto.RegisterType((*ItemAct)(nil), "REDACTED")
	proto.RegisterType((*DominoAct)(nil), "REDACTED")
	proto.RegisterType((*EvidenceAct)(nil), "REDACTED")
	proto.RegisterType((*EvidenceActions)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_6b60b6ba2ab5b856) }

var filedefinition_6b60b6ba2ab5b856 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xbd, 0x6a, 0xe3, 0x40,
	0x10, 0x96, 0x2c, 0xf9, 0x6f, 0xed, 0xe2, 0x6e, 0x31, 0x87, 0xf0, 0x71, 0x3a, 0xa1, 0x4a, 0x95,
	0x04, 0x4e, 0xea, 0x14, 0x4e, 0x8a, 0x90, 0x40, 0x1c, 0x54, 0xa4, 0x48, 0x13, 0xd6, 0xf6, 0xca,
	0x12, 0xb1, 0x34, 0x8b, 0x34, 0x82, 0xf8, 0x2d, 0xf2, 0x58, 0x2e, 0x5d, 0xa6, 0x0a, 0xc1, 0x7e,
	0x91, 0xb0, 0xbb, 0x0a, 0x26, 0x98, 0x74, 0xdf, 0xcf, 0xec, 0x37, 0xdf, 0x20, 0x91, 0x7f, 0xc8,
	0x8b, 0x25, 0x2f, 0xf3, 0xac, 0xc0, 0x68, 0x51, 0x6e, 0x04, 0x42, 0x24, 0x4a, 0x80, 0x24, 0x14,
	0x25, 0x20, 0xd0, 0xdf, 0x47, 0x3b, 0xd4, 0xf6, 0x78, 0xb4, 0x82, 0x15, 0x28, 0x37, 0x92, 0x48,
	0x0f, 0xfa, 0x09, 0x69, 0xdf, 0xcb, 0x77, 0x74, 0x44, 0xda, 0x08, 0xc8, 0xd6, 0x8e, 0xe9, 0x99,
	0x81, 0x15, 0x6b, 0x22, 0xd5, 0xac, 0x58, 0xf2, 0x17, 0xa7, 0xa5, 0x55, 0x45, 0xe8, 0x5f, 0xd2,
	0x5f, 0x73, 0x96, 0x3c, 0xa5, 0xac, 0x4a, 0x1d, 0xcb, 0x33, 0x83, 0x61, 0xdc, 0x93, 0xc2, 0x35,
	0xab, 0x52, 0xf9, 0x84, 0xd5, 0x05, 0x56, 0x8e, 0xed, 0x59, 0xc1, 0x30, 0xd6, 0xc4, 0xbf, 0x25,
	0xdd, 0x07, 0xb6, 0xae, 0xf9, 0x4c, 0xd0, 0x5f, 0xc4, 0x7a, 0xe6, 0x1b, 0xb5, 0x67, 0x18, 0x4b,
	0x48, 0x43, 0xd2, 0x56, 0xe5, 0xd5, 0x96, 0xc1, 0xc4, 0x09, 0x4f, 0xda, 0x87, 0xaa, 0x64, 0xac,
	0xc7, 0xfc, 0x1b, 0xd2, 0xbb, 0x82, 0x3c, 0x2b, 0xe0, 0x7b, 0x5a, 0x5f, 0xa7, 0xa9, 0xce, 0xa2,
	0x46, 0x95, 0xd6, 0x8f, 0x35, 0xa1, 0x7f, 0x48, 0x07, 0x6a, 0x94, 0xb2, 0xa5, 0xe4, 0x86, 0xf9,
	0x97, 0xa4, 0xab, 0xb2, 0x67, 0x82, 0x52, 0x62, 0xe3, 0x46, 0xf0, 0x26, 0x4b, 0xe1, 0xaf, 0xf8,
	0xd6, 0xb1, 0x2c, 0x25, 0xf6, 0x92, 0x21, 0x6b, 0xee, 0x56, 0xd8, 0xbf, 0x20, 0xbd, 0x26, 0xa4,
	0xa2, 0x13, 0x62, 0x81, 0xa8, 0x1c, 0xd3, 0xb3, 0x82, 0xc1, 0x64, 0xfc, 0xd3, 0x29, 0x33, 0x31,
	0xb5, 0xb7, 0xef, 0xff, 0x8d, 0x58, 0x0e, 0x4f, 0xef, 0xb6, 0x7b, 0xd7, 0xdc, 0xed, 0x5d, 0xf3,
	0x63, 0xef, 0x9a, 0xaf, 0x07, 0xd7, 0xd8, 0x1d, 0x5c, 0xe3, 0xed, 0xe0, 0x1a, 0x8f, 0xe7, 0xab,
	0x0c, 0xd3, 0x7a, 0x1e, 0x2e, 0x20, 0x8f, 0x16, 0x90, 0x73, 0x9c, 0x27, 0x78, 0x04, 0xfa, 0x73,
	0x9e, 0xfc, 0x0a, 0xf3, 0x8e, 0x32, 0xce, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x5a, 0xb3,
	0xb6, 0x26, 0x02, 0x00, 0x00,
}

func (m *Attestation) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attestation) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Attestation) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Kin) > 0 {
		for idxNdEx := len(m.Kin) - 1; idxNdEx >= 0; idxNdEx-- {
			i -= len(m.Kin[idxNdEx])
			copy(dAtA[i:], m.Kin[idxNdEx])
			i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Kin[idxNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.NodeDigest) > 0 {
		i -= len(m.NodeDigest)
		copy(dAtA[i:], m.NodeDigest)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.NodeDigest)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Ordinal != 0 {
		i = encodeVariableintEvidence(dAtA, i, uint64(m.Ordinal))
		i--
		dAtA[i] = 0x10
	}
	if m.Sum != 0 {
		i = encodeVariableintEvidence(dAtA, i, uint64(m.Sum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ItemAct) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ItemAct) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ItemAct) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Attestation != nil {
		{
			volume, err := m.Attestation.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintEvidence(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DominoAct) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DominoAct) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *DominoAct) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Influx) > 0 {
		i -= len(m.Influx)
		copy(dAtA[i:], m.Influx)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Influx)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EvidenceAct) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvidenceAct) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *EvidenceAct) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVariableintEvidence(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EvidenceActions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvidenceActions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *EvidenceActions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ops) > 0 {
		for idxNdEx := len(m.Ops) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Ops[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = encodeVariableintEvidence(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVariableintEvidence(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovEvidence(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *Attestation) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != 0 {
		n += 1 + sovEvidence(uint64(m.Sum))
	}
	if m.Ordinal != 0 {
		n += 1 + sovEvidence(uint64(m.Ordinal))
	}
	l = len(m.NodeDigest)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	if len(m.Kin) > 0 {
		for _, b := range m.Kin {
			l = len(b)
			n += 1 + l + sovEvidence(uint64(l))
		}
	}
	return n
}

func (m *ItemAct) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	if m.Attestation != nil {
		l = m.Attestation.Volume()
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}

func (m *DominoAct) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	l = len(m.Influx)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}

func (m *EvidenceAct) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}

func (m *EvidenceActions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ops) > 0 {
		for _, e := range m.Ops {
			l = e.Volume()
			n += 1 + l + sovEvidence(uint64(l))
		}
	}
	return n
}

func sovEvidence(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvidence(x uint64) (n int) {
	return sovEvidence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attestation) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadEvidence
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
			m.Sum = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Sum |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Ordinal |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeDigest = append(m.NodeDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.NodeDigest == nil {
				m.NodeDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Kin = append(m.Kin, make([]byte, submitOrdinal-idxNdEx))
			copy(m.Kin[len(m.Kin)-1], dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitEvidence(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentEvidence
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
func (m *ItemAct) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadEvidence
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
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attestation == nil {
				m.Attestation = &Attestation{}
			}
			if err := m.Attestation.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitEvidence(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentEvidence
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
func (m *DominoAct) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadEvidence
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
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Influx = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitEvidence(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentEvidence
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
func (m *EvidenceAct) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadEvidence
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
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
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
			skippy, err := omitEvidence(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentEvidence
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
func (m *EvidenceActions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadEvidence
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
					return ErrIntegerOverloadEvidence
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
				return ErrCorruptExtentEvidence
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvidence
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Ops = append(m.Ops, EvidenceAct{})
			if err := m.Ops[len(m.Ops)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitEvidence(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentEvidence
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
func omitEvidence(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadEvidence
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
					return 0, ErrIntegerOverloadEvidence
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
					return 0, ErrIntegerOverloadEvidence
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
				return 0, ErrCorruptExtentEvidence
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterEvidence
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentEvidence
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentEvidence        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadEvidence          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterEvidence = fmt.Errorf("REDACTED")
)
