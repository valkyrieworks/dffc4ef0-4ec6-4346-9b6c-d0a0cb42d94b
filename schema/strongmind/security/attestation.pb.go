//
//

package security

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
	NodeDigest []byte   `protobuf:"octets,3,opt,name=leaf_hash,json=leafHash,proto3" json:"terminal_digest,omitempty"`
	Kin    [][]byte `protobuf:"octets,4,rep,name=aunts,proto3" json:"kin,omitempty"`
}

func (m *Attestation) Restore()         { *m = Attestation{} }
func (m *Attestation) Text() string { return proto.CompactTextString(m) }
func (*Attestation) SchemaArtifact()    {}
func (*Attestation) Definition() ([]byte, []int) {
	return filedescriptor_6b60b6ba2ab5b856, []int{0}
}
func (m *Attestation) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Attestation) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Attestation.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attestation) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Extent() int {
	return m.Extent()
}
func (m *Attestation) XXX_Dropunfamiliar() {
	xxx_signaldetails_Attestation.DiscardUnknown(m)
}

var xxx_signaldetails_Attestation proto.InternalMessageInfo

func (m *Attestation) ObtainSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *Attestation) ObtainOrdinal() int64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *Attestation) ObtainFolioDigest() []byte {
	if m != nil {
		return m.NodeDigest
	}
	return nil
}

func (m *Attestation) ObtainKin() [][]byte {
	if m != nil {
		return m.Kin
	}
	return nil
}

type DatumAction struct {
	//
	Key []byte `protobuf:"octets,1,opt,name=key,proto3" json:"key,omitempty"`
	//
	Attestation *Attestation `protobuf:"octets,2,opt,name=proof,proto3" json:"attestation,omitempty"`
}

func (m *DatumAction) Restore()         { *m = DatumAction{} }
func (m *DatumAction) Text() string { return proto.CompactTextString(m) }
func (*DatumAction) SchemaArtifact()    {}
func (*DatumAction) Definition() ([]byte, []int) {
	return filedescriptor_6b60b6ba2ab5b856, []int{1}
}
func (m *DatumAction) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *DatumAction) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Datumaction.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DatumAction) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Datumaction.Merge(m, src)
}
func (m *DatumAction) XXX_Extent() int {
	return m.Extent()
}
func (m *DatumAction) XXX_Dropunfamiliar() {
	xxx_signaldetails_Datumaction.DiscardUnknown(m)
}

var xxx_signaldetails_Datumaction proto.InternalMessageInfo

func (m *DatumAction) ObtainToken() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *DatumAction) ObtainAttestation() *Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

type CascadeAction struct {
	Key    string `protobuf:"octets,1,opt,name=key,proto3" json:"key,omitempty"`
	Influx  string `protobuf:"octets,2,opt,name=input,proto3" json:"influx,omitempty"`
	Emission string `protobuf:"octets,3,opt,name=output,proto3" json:"emission,omitempty"`
}

func (m *CascadeAction) Restore()         { *m = CascadeAction{} }
func (m *CascadeAction) Text() string { return proto.CompactTextString(m) }
func (*CascadeAction) SchemaArtifact()    {}
func (*CascadeAction) Definition() ([]byte, []int) {
	return filedescriptor_6b60b6ba2ab5b856, []int{2}
}
func (m *CascadeAction) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *CascadeAction) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Dominoaction.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CascadeAction) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Dominoaction.Merge(m, src)
}
func (m *CascadeAction) XXX_Extent() int {
	return m.Extent()
}
func (m *CascadeAction) XXX_Dropunfamiliar() {
	xxx_signaldetails_Dominoaction.DiscardUnknown(m)
}

var xxx_signaldetails_Dominoaction proto.InternalMessageInfo

func (m *CascadeAction) ObtainToken() string {
	if m != nil {
		return m.Key
	}
	return "REDACTED"
}

func (m *CascadeAction) ObtainInflux() string {
	if m != nil {
		return m.Influx
	}
	return "REDACTED"
}

func (m *CascadeAction) ObtainEmission() string {
	if m != nil {
		return m.Emission
	}
	return "REDACTED"
}

//
//
//
type AttestationAction struct {
	Kind string `protobuf:"octets,1,opt,name=type,proto3" json:"kind,omitempty"`
	Key  []byte `protobuf:"octets,2,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"octets,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AttestationAction) Restore()         { *m = AttestationAction{} }
func (m *AttestationAction) Text() string { return proto.CompactTextString(m) }
func (*AttestationAction) SchemaArtifact()    {}
func (*AttestationAction) Definition() ([]byte, []int) {
	return filedescriptor_6b60b6ba2ab5b856, []int{3}
}
func (m *AttestationAction) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AttestationAction) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Attestationaction.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestationAction) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Attestationaction.Merge(m, src)
}
func (m *AttestationAction) XXX_Extent() int {
	return m.Extent()
}
func (m *AttestationAction) XXX_Dropunfamiliar() {
	xxx_signaldetails_Attestationaction.DiscardUnknown(m)
}

var xxx_signaldetails_Attestationaction proto.InternalMessageInfo

func (m *AttestationAction) ObtainKind() string {
	if m != nil {
		return m.Kind
	}
	return "REDACTED"
}

func (m *AttestationAction) ObtainToken() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AttestationAction) ObtainData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

//
type AttestationActions struct {
	Ops []AttestationAction `protobuf:"octets,1,rep,name=ops,proto3" json:"ops"`
}

func (m *AttestationActions) Restore()         { *m = AttestationActions{} }
func (m *AttestationActions) Text() string { return proto.CompactTextString(m) }
func (*AttestationActions) SchemaArtifact()    {}
func (*AttestationActions) Definition() ([]byte, []int) {
	return filedescriptor_6b60b6ba2ab5b856, []int{4}
}
func (m *AttestationActions) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AttestationActions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Attestationactions.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestationActions) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Attestationactions.Merge(m, src)
}
func (m *AttestationActions) XXX_Extent() int {
	return m.Extent()
}
func (m *AttestationActions) XXX_Dropunfamiliar() {
	xxx_signaldetails_Attestationactions.DiscardUnknown(m)
}

var xxx_signaldetails_Attestationactions proto.InternalMessageInfo

func (m *AttestationActions) ObtainActions() []AttestationAction {
	if m != nil {
		return m.Ops
	}
	return nil
}

func initialize() {
	proto.RegisterType((*Attestation)(nil), "REDACTED")
	proto.RegisterType((*DatumAction)(nil), "REDACTED")
	proto.RegisterType((*CascadeAction)(nil), "REDACTED")
	proto.RegisterType((*AttestationAction)(nil), "REDACTED")
	proto.RegisterType((*AttestationActions)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_6b60b6ba2ab5b856) }

var filedescriptor_6b60b6ba2ab5b856 = []byte{
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

func (m *Attestation) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Attestation) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Attestation) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Kin) > 0 {
		for idxNdExc := len(m.Kin) - 1; idxNdExc >= 0; idxNdExc-- {
			i -= len(m.Kin[idxNdExc])
			copy(deltaLocatedAN[i:], m.Kin[idxNdExc])
			i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Kin[idxNdExc])))
			i--
			deltaLocatedAN[i] = 0x22
		}
	}
	if len(m.NodeDigest) > 0 {
		i -= len(m.NodeDigest)
		copy(deltaLocatedAN[i:], m.NodeDigest)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.NodeDigest)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.Ordinal != 0 {
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(m.Ordinal))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Sum != 0 {
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(m.Sum))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *DatumAction) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *DatumAction) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *DatumAction) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Attestation != nil {
		{
			extent, err := m.Attestation.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(deltaLocatedAN[i:], m.Key)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Key)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *CascadeAction) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *CascadeAction) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *CascadeAction) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Emission) > 0 {
		i -= len(m.Emission)
		copy(deltaLocatedAN[i:], m.Emission)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Emission)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if len(m.Influx) > 0 {
		i -= len(m.Influx)
		copy(deltaLocatedAN[i:], m.Influx)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Influx)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(deltaLocatedAN[i:], m.Key)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Key)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *AttestationAction) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AttestationAction) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AttestationAction) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(deltaLocatedAN[i:], m.Data)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Data)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(deltaLocatedAN[i:], m.Key)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Key)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(deltaLocatedAN[i:], m.Kind)
		i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(len(m.Kind)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *AttestationActions) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AttestationActions) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AttestationActions) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Ops) > 0 {
		for idxNdExc := len(m.Ops) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Ops[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = encodeVariableintAttestation(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintAttestation(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovAttestation(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *Attestation) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != 0 {
		n += 1 + sovAttestation(uint64(m.Sum))
	}
	if m.Ordinal != 0 {
		n += 1 + sovAttestation(uint64(m.Ordinal))
	}
	l = len(m.NodeDigest)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	if len(m.Kin) > 0 {
		for _, b := range m.Kin {
			l = len(b)
			n += 1 + l + sovAttestation(uint64(l))
		}
	}
	return n
}

func (m *DatumAction) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	if m.Attestation != nil {
		l = m.Attestation.Extent()
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}

func (m *CascadeAction) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = len(m.Influx)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = len(m.Emission)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}

func (m *AttestationAction) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}

func (m *AttestationActions) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ops) > 0 {
		for _, e := range m.Ops {
			l = e.Extent()
			n += 1 + l + sovAttestation(uint64(l))
		}
	}
	return n
}

func sovAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttestation(x uint64) (n int) {
	return sovAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attestation) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAttestation
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
			m.Sum = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Sum |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Ordinal |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeDigest = append(m.NodeDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.NodeDigest == nil {
				m.NodeDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Kin = append(m.Kin, make([]byte, submitOrdinal-idxNdExc))
			copy(m.Kin[len(m.Kin)-1], deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitAttestation(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAttestation
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
func (m *DatumAction) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAttestation
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
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attestation == nil {
				m.Attestation = &Attestation{}
			}
			if err := m.Attestation.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitAttestation(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAttestation
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
func (m *CascadeAction) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAttestation
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
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Influx = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Emission = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitAttestation(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAttestation
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
func (m *AttestationAction) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAttestation
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
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
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
			omitted, err := omitAttestation(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAttestation
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
func (m *AttestationActions) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAttestation
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
					return FaultIntegerOverrunAttestation
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
				return FaultUnfitMagnitudeAttestation
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAttestation
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Ops = append(m.Ops, AttestationAction{})
			if err := m.Ops[len(m.Ops)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitAttestation(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAttestation
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
func omitAttestation(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunAttestation
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
					return 0, FaultIntegerOverrunAttestation
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
					return 0, FaultIntegerOverrunAttestation
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
				return 0, FaultUnfitMagnitudeAttestation
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionAttestation
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeAttestation
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeAttestation        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunAttestation          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionAttestation = fmt.Errorf("REDACTED")
)
