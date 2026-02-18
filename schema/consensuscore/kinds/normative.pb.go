//
//

package kinds

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

//
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

//
//
//
//
const _ = proto.GoGoProtoPackageIsVersion3 //

type StandardLedgerUID struct {
	Digest          []byte                 `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	SegmentAssignHeading StandardSectionCollectionHeading `protobuf:"octets,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"section_collection_heading"`
}

func (m *StandardLedgerUID) Restore()         { *m = StandardLedgerUID{} }
func (m *StandardLedgerUID) String() string { return proto.CompactTextString(m) }
func (*StandardLedgerUID) SchemaSignal()    {}
func (*StandardLedgerUID) Definition() ([]byte, []int) {
	return filedefinition_8d1a1a84ff7267ed, []int{0}
}
func (m *StandardLedgerUID) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StandardLedgerUID) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Normativeledgeruid.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardLedgerUID) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Normativeledgeruid.Merge(m, src)
}
func (m *StandardLedgerUID) XXX_Volume() int {
	return m.Volume()
}
func (m *StandardLedgerUID) XXX_Omitunclear() {
	xxx_messagedata_Normativeledgeruid.DiscardUnknown(m)
}

var xxx_messagedata_Normativeledgeruid proto.InternalMessageInfo

func (m *StandardLedgerUID) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *StandardLedgerUID) FetchSectionCollectionHeading() StandardSectionCollectionHeading {
	if m != nil {
		return m.SegmentAssignHeading
	}
	return StandardSectionCollectionHeading{}
}

type StandardSectionCollectionHeading struct {
	Sum uint32 `protobuf:"variableint,1,opt,name=total,proto3" json:"sum,omitempty"`
	Digest  []byte `protobuf:"octets,2,opt,name=hash,proto3" json:"digest,omitempty"`
}

func (m *StandardSectionCollectionHeading) Restore()         { *m = StandardSectionCollectionHeading{} }
func (m *StandardSectionCollectionHeading) String() string { return proto.CompactTextString(m) }
func (*StandardSectionCollectionHeading) SchemaSignal()    {}
func (*StandardSectionCollectionHeading) Definition() ([]byte, []int) {
	return filedefinition_8d1a1a84ff7267ed, []int{1}
}
func (m *StandardSectionCollectionHeading) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StandardSectionCollectionHeading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Normativesectionsetheading.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardSectionCollectionHeading) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Normativesectionsetheading.Merge(m, src)
}
func (m *StandardSectionCollectionHeading) XXX_Volume() int {
	return m.Volume()
}
func (m *StandardSectionCollectionHeading) XXX_Omitunclear() {
	xxx_messagedata_Normativesectionsetheading.DiscardUnknown(m)
}

var xxx_messagedata_Normativesectionsetheading proto.InternalMessageInfo

func (m *StandardSectionCollectionHeading) FetchSum() uint32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *StandardSectionCollectionHeading) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type StandardNomination struct {
	Kind      AttestedMessageKind     `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Level    int64             `protobuf:"fixed64,2,opt,name=height,proto3" json:"level,omitempty"`
	Cycle     int64             `protobuf:"fixed64,3,opt,name=round,proto3" json:"epoch,omitempty"`
	POLDuration  int64             `protobuf:"variableint,4,opt,name=pol_round,json=polRound,proto3" json:"pol_epoch,omitempty"`
	LedgerUID   *StandardLedgerUID `protobuf:"octets,5,opt,name=block_id,json=blockId,proto3" json:"ledger_uid,omitempty"`
	Timestamp time.Time         `protobuf:"octets,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	LedgerUID   string            `protobuf:"octets,7,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
}

func (m *StandardNomination) Restore()         { *m = StandardNomination{} }
func (m *StandardNomination) String() string { return proto.CompactTextString(m) }
func (*StandardNomination) SchemaSignal()    {}
func (*StandardNomination) Definition() ([]byte, []int) {
	return filedefinition_8d1a1a84ff7267ed, []int{2}
}
func (m *StandardNomination) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StandardNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Normativeproposal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardNomination) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Normativeproposal.Merge(m, src)
}
func (m *StandardNomination) XXX_Volume() int {
	return m.Volume()
}
func (m *StandardNomination) XXX_Omitunclear() {
	xxx_messagedata_Normativeproposal.DiscardUnknown(m)
}

var xxx_messagedata_Normativeproposal proto.InternalMessageInfo

func (m *StandardNomination) FetchKind() AttestedMessageKind {
	if m != nil {
		return m.Kind
	}
	return UnclearKind
}

func (m *StandardNomination) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *StandardNomination) FetchDuration() int64 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *StandardNomination) FetchPOLEpoch() int64 {
	if m != nil {
		return m.POLDuration
	}
	return 0
}

func (m *StandardNomination) FetchLedgerUID() *StandardLedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return nil
}

func (m *StandardNomination) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *StandardNomination) FetchSeriesUID() string {
	if m != nil {
		return m.LedgerUID
	}
	return "REDACTED"
}

type StandardBallot struct {
	Kind      AttestedMessageKind     `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Level    int64             `protobuf:"fixed64,2,opt,name=height,proto3" json:"level,omitempty"`
	Cycle     int64             `protobuf:"fixed64,3,opt,name=round,proto3" json:"epoch,omitempty"`
	LedgerUID   *StandardLedgerUID `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uid,omitempty"`
	Timestamp time.Time         `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	LedgerUID   string            `protobuf:"octets,6,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
}

func (m *StandardBallot) Restore()         { *m = StandardBallot{} }
func (m *StandardBallot) String() string { return proto.CompactTextString(m) }
func (*StandardBallot) SchemaSignal()    {}
func (*StandardBallot) Definition() ([]byte, []int) {
	return filedefinition_8d1a1a84ff7267ed, []int{3}
}
func (m *StandardBallot) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StandardBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Normativeballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardBallot) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Normativeballot.Merge(m, src)
}
func (m *StandardBallot) XXX_Volume() int {
	return m.Volume()
}
func (m *StandardBallot) XXX_Omitunclear() {
	xxx_messagedata_Normativeballot.DiscardUnknown(m)
}

var xxx_messagedata_Normativeballot proto.InternalMessageInfo

func (m *StandardBallot) FetchKind() AttestedMessageKind {
	if m != nil {
		return m.Kind
	}
	return UnclearKind
}

func (m *StandardBallot) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *StandardBallot) FetchDuration() int64 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *StandardBallot) FetchLedgerUID() *StandardLedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return nil
}

func (m *StandardBallot) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *StandardBallot) FetchSeriesUID() string {
	if m != nil {
		return m.LedgerUID
	}
	return "REDACTED"
}

//
//
type StandardBallotAddition struct {
	Addition []byte `protobuf:"octets,1,opt,name=extension,proto3" json:"addition,omitempty"`
	Level    int64  `protobuf:"fixed64,2,opt,name=height,proto3" json:"level,omitempty"`
	Cycle     int64  `protobuf:"fixed64,3,opt,name=round,proto3" json:"epoch,omitempty"`
	SeriesUid   string `protobuf:"octets,4,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
}

func (m *StandardBallotAddition) Restore()         { *m = StandardBallotAddition{} }
func (m *StandardBallotAddition) String() string { return proto.CompactTextString(m) }
func (*StandardBallotAddition) SchemaSignal()    {}
func (*StandardBallotAddition) Definition() ([]byte, []int) {
	return filedefinition_8d1a1a84ff7267ed, []int{4}
}
func (m *StandardBallotAddition) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StandardBallotAddition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Normativeballotextension.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardBallotAddition) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Normativeballotextension.Merge(m, src)
}
func (m *StandardBallotAddition) XXX_Volume() int {
	return m.Volume()
}
func (m *StandardBallotAddition) XXX_Omitunclear() {
	xxx_messagedata_Normativeballotextension.DiscardUnknown(m)
}

var xxx_messagedata_Normativeballotextension proto.InternalMessageInfo

func (m *StandardBallotAddition) FetchAddition() []byte {
	if m != nil {
		return m.Addition
	}
	return nil
}

func (m *StandardBallotAddition) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *StandardBallotAddition) FetchDuration() int64 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *StandardBallotAddition) FetchSeriesUid() string {
	if m != nil {
		return m.SeriesUid
	}
	return "REDACTED"
}

func init() {
	proto.RegisterType((*StandardLedgerUID)(nil), "REDACTED")
	proto.RegisterType((*StandardSectionCollectionHeading)(nil), "REDACTED")
	proto.RegisterType((*StandardNomination)(nil), "REDACTED")
	proto.RegisterType((*StandardBallot)(nil), "REDACTED")
	proto.RegisterType((*StandardBallotAddition)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_8d1a1a84ff7267ed) }

var filedefinition_8d1a1a84ff7267ed = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x6e, 0x9b, 0x40,
	0x10, 0x35, 0x0e, 0xb6, 0x61, 0x13, 0xb7, 0xee, 0x2a, 0x8a, 0xa8, 0x15, 0x01, 0xe2, 0x50, 0xd1,
	0x0b, 0x48, 0xf1, 0x1f, 0x90, 0x56, 0xaa, 0xab, 0x46, 0x8d, 0x48, 0x94, 0x43, 0x2f, 0xd6, 0x02,
	0x1b, 0x40, 0x05, 0x16, 0xc1, 0x5a, 0x6a, 0x2e, 0xed, 0x2f, 0xe4, 0x3b, 0xfa, 0x25, 0x39, 0xe6,
	0xd8, 0x5e, 0xdc, 0x0a, 0xff, 0x48, 0xb5, 0x0b, 0x06, 0x2b, 0xa9, 0x2c, 0x55, 0xad, 0x7a, 0x41,
	0x33, 0x6f, 0xde, 0xce, 0x3c, 0xbd, 0x61, 0x17, 0xe8, 0x14, 0x67, 0x01, 0x2e, 0xd2, 0x38, 0xa3,
	0x36, 0xbd, 0xc9, 0x71, 0x69, 0xfb, 0x28, 0x23, 0x59, 0xec, 0xa3, 0xc4, 0xca, 0x0b, 0x42, 0x09,
	0x9c, 0x74, 0x0c, 0x8b, 0x33, 0xa6, 0x87, 0x21, 0x09, 0x09, 0x2f, 0xda, 0x2c, 0xaa, 0x79, 0xd3,
	0xe3, 0x47, 0x9d, 0xf8, 0xb7, 0xa9, 0x6a, 0x21, 0x21, 0x61, 0x82, 0x6d, 0x9e, 0x79, 0xcb, 0x6b,
	0x9b, 0xc6, 0x29, 0x2e, 0x29, 0x4a, 0xf3, 0x9a, 0x60, 0x7c, 0x06, 0x93, 0xd3, 0xcd, 0x64, 0x27,
	0x21, 0xfe, 0xc7, 0xf9, 0x2b, 0x08, 0x81, 0x18, 0xa1, 0x32, 0x52, 0x04, 0x5d, 0x30, 0x0f, 0x5c,
	0x1e, 0xc3, 0x2b, 0xf0, 0x34, 0x47, 0x05, 0x5d, 0x94, 0x98, 0x2e, 0x22, 0x8c, 0x02, 0x5c, 0x28,
	0x7d, 0x5d, 0x30, 0xf7, 0x4f, 0x4c, 0xeb, 0xa1, 0x50, 0xab, 0x6d, 0x78, 0x8e, 0x0a, 0x7a, 0x81,
	0xe9, 0x1b, 0xce, 0x77, 0xc4, 0xbb, 0x95, 0xd6, 0x73, 0xc7, 0xf9, 0x36, 0x68, 0x38, 0xe0, 0xe8,
	0xf7, 0x74, 0x78, 0x08, 0x06, 0x94, 0x50, 0x94, 0x70, 0x19, 0x63, 0xb7, 0x4e, 0x5a, 0x6d, 0xfd,
	0x4e, 0x9b, 0xf1, 0xbd, 0x0f, 0x9e, 0x75, 0x4d, 0x0a, 0x92, 0x93, 0x12, 0x25, 0x70, 0x06, 0x44,
	0x26, 0x87, 0x1f, 0x7f, 0x72, 0xa2, 0x3d, 0x96, 0x79, 0x11, 0x87, 0x19, 0x0e, 0xce, 0xca, 0xf0,
	0xf2, 0x26, 0xc7, 0x2e, 0x27, 0xc3, 0x23, 0x30, 0x8c, 0x70, 0x1c, 0x46, 0x94, 0x0f, 0x98, 0xb8,
	0x4d, 0xc6, 0xc4, 0x14, 0x64, 0x99, 0x05, 0xca, 0x1e, 0x87, 0xeb, 0x04, 0xbe, 0x04, 0x72, 0x4e,
	0x92, 0x45, 0x5d, 0x11, 0x75, 0xc1, 0xdc, 0x73, 0x0e, 0xaa, 0x95, 0x26, 0x9d, 0xbf, 0x7f, 0xe7,
	0x32, 0xcc, 0x95, 0x72, 0x92, 0xf0, 0x08, 0xbe, 0x05, 0x92, 0xc7, 0xec, 0x5d, 0xc4, 0x81, 0x32,
	0xe0, 0xc6, 0x19, 0x3b, 0x8c, 0x6b, 0x36, 0xe1, 0xec, 0x57, 0x2b, 0x6d, 0xd4, 0x24, 0xee, 0x88,
	0x37, 0x98, 0x07, 0xd0, 0x01, 0x72, 0xbb, 0x46, 0x65, 0xc8, 0x9b, 0x4d, 0xad, 0x7a, 0xd1, 0xd6,
	0x66, 0xd1, 0xd6, 0xe5, 0x86, 0xe1, 0x48, 0xcc, 0xf7, 0xdb, 0x1f, 0x9a, 0xe0, 0x76, 0xc7, 0xe0,
	0x0b, 0x20, 0xf9, 0x11, 0x8a, 0x33, 0xa6, 0x67, 0xa4, 0x0b, 0xa6, 0x5c, 0xcf, 0x3a, 0x65, 0x18,
	0x9b, 0xc5, 0x8b, 0xf3, 0xc0, 0xf8, 0xda, 0x07, 0xe3, 0x56, 0xd6, 0x15, 0xa1, 0xf8, 0x7f, 0xf8,
	0xba, 0x6d, 0x96, 0xf8, 0x2f, 0xcd, 0x1a, 0xfc, 0xbd, 0x59, 0xc3, 0x1d, 0x66, 0x7d, 0xd9, 0xfa,
	0x99, 0x99, 0x57, 0xaf, 0x3f, 0x51, 0x9c, 0x95, 0x31, 0xc9, 0xe0, 0x31, 0x90, 0xf1, 0x26, 0x69,
	0xee, 0x55, 0x07, 0xfc, 0xa1, 0x3b, 0xcf, 0xb7, 0xd4, 0x30, 0x77, 0xe4, 0x56, 0x80, 0x73, 0x76,
	0x57, 0xa9, 0xc2, 0x7d, 0xa5, 0x0a, 0x3f, 0x2b, 0x55, 0xb8, 0x5d, 0xab, 0xbd, 0xfb, 0xb5, 0xda,
	0xfb, 0xb6, 0x56, 0x7b, 0x1f, 0x66, 0x61, 0x4c, 0xa3, 0xa5, 0x67, 0xf9, 0x24, 0xb5, 0x7d, 0x92,
	0x62, 0xea, 0x5d, 0xd3, 0x2e, 0xa8, 0x5f, 0x95, 0x87, 0x2f, 0x89, 0x37, 0xe4, 0xf8, 0xec, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x22, 0x5b, 0x0b, 0xae, 0x04, 0x00, 0x00,
}

func (m *StandardLedgerUID) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardLedgerUID) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StandardLedgerUID) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.SegmentAssignHeading.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = encodeVariableintNormative(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVariableintNormative(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StandardSectionCollectionHeading) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardSectionCollectionHeading) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StandardSectionCollectionHeading) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVariableintNormative(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sum != 0 {
		i = encodeVariableintNormative(dAtA, i, uint64(m.Sum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StandardNomination) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardNomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StandardNomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LedgerUID) > 0 {
		i -= len(m.LedgerUID)
		copy(dAtA[i:], m.LedgerUID)
		i = encodeVariableintNormative(dAtA, i, uint64(len(m.LedgerUID)))
		i--
		dAtA[i] = 0x3a
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVariableintNormative(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	if m.LedgerUID != nil {
		{
			volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintNormative(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.POLDuration != 0 {
		i = encodeVariableintNormative(dAtA, i, uint64(m.POLDuration))
		i--
		dAtA[i] = 0x20
	}
	if m.Cycle != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Cycle))
		i--
		dAtA[i] = 0x19
	}
	if m.Level != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Level))
		i--
		dAtA[i] = 0x11
	}
	if m.Kind != 0 {
		i = encodeVariableintNormative(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StandardBallot) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardBallot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StandardBallot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LedgerUID) > 0 {
		i -= len(m.LedgerUID)
		copy(dAtA[i:], m.LedgerUID)
		i = encodeVariableintNormative(dAtA, i, uint64(len(m.LedgerUID)))
		i--
		dAtA[i] = 0x32
	}
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVariableintNormative(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x2a
	if m.LedgerUID != nil {
		{
			volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintNormative(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Cycle != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Cycle))
		i--
		dAtA[i] = 0x19
	}
	if m.Level != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Level))
		i--
		dAtA[i] = 0x11
	}
	if m.Kind != 0 {
		i = encodeVariableintNormative(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StandardBallotAddition) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandardBallotAddition) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StandardBallotAddition) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SeriesUid) > 0 {
		i -= len(m.SeriesUid)
		copy(dAtA[i:], m.SeriesUid)
		i = encodeVariableintNormative(dAtA, i, uint64(len(m.SeriesUid)))
		i--
		dAtA[i] = 0x22
	}
	if m.Cycle != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Cycle))
		i--
		dAtA[i] = 0x19
	}
	if m.Level != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Level))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Addition) > 0 {
		i -= len(m.Addition)
		copy(dAtA[i:], m.Addition)
		i = encodeVariableintNormative(dAtA, i, uint64(len(m.Addition)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVariableintNormative(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovNormative(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *StandardLedgerUID) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovNormative(uint64(l))
	}
	l = m.SegmentAssignHeading.Volume()
	n += 1 + l + sovNormative(uint64(l))
	return n
}

func (m *StandardSectionCollectionHeading) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != 0 {
		n += 1 + sovNormative(uint64(m.Sum))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovNormative(uint64(l))
	}
	return n
}

func (m *StandardNomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovNormative(uint64(m.Kind))
	}
	if m.Level != 0 {
		n += 9
	}
	if m.Cycle != 0 {
		n += 9
	}
	if m.POLDuration != 0 {
		n += 1 + sovNormative(uint64(m.POLDuration))
	}
	if m.LedgerUID != nil {
		l = m.LedgerUID.Volume()
		n += 1 + l + sovNormative(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovNormative(uint64(l))
	l = len(m.LedgerUID)
	if l > 0 {
		n += 1 + l + sovNormative(uint64(l))
	}
	return n
}

func (m *StandardBallot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovNormative(uint64(m.Kind))
	}
	if m.Level != 0 {
		n += 9
	}
	if m.Cycle != 0 {
		n += 9
	}
	if m.LedgerUID != nil {
		l = m.LedgerUID.Volume()
		n += 1 + l + sovNormative(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovNormative(uint64(l))
	l = len(m.LedgerUID)
	if l > 0 {
		n += 1 + l + sovNormative(uint64(l))
	}
	return n
}

func (m *StandardBallotAddition) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addition)
	if l > 0 {
		n += 1 + l + sovNormative(uint64(l))
	}
	if m.Level != 0 {
		n += 9
	}
	if m.Cycle != 0 {
		n += 9
	}
	l = len(m.SeriesUid)
	if l > 0 {
		n += 1 + l + sovNormative(uint64(l))
	}
	return n
}

func sovNormative(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNormative(x uint64) (n int) {
	return sovNormative(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StandardLedgerUID) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadNormative
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
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = append(m.Digest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SegmentAssignHeading.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitNormative(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentNormative
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
func (m *StandardSectionCollectionHeading) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadNormative
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
					return ErrIntegerOverloadNormative
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Sum |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = append(m.Digest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitNormative(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentNormative
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
func (m *StandardNomination) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadNormative
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
			m.Kind = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Kind |= AttestedMessageKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Level = 0
			if (idxNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = int64(encoding_binary.LittleEndian.Uint64(dAtA[idxNdEx:]))
			idxNdEx += 8
		case 3:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			if (idxNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Cycle = int64(encoding_binary.LittleEndian.Uint64(dAtA[idxNdEx:]))
			idxNdEx += 8
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.POLDuration = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.POLDuration |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.LedgerUID == nil {
				m.LedgerUID = &StandardLedgerUID{}
			}
			if err := m.LedgerUID.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.LedgerUID = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitNormative(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentNormative
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
func (m *StandardBallot) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadNormative
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
			m.Kind = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Kind |= AttestedMessageKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Level = 0
			if (idxNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = int64(encoding_binary.LittleEndian.Uint64(dAtA[idxNdEx:]))
			idxNdEx += 8
		case 3:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			if (idxNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Cycle = int64(encoding_binary.LittleEndian.Uint64(dAtA[idxNdEx:]))
			idxNdEx += 8
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.LedgerUID == nil {
				m.LedgerUID = &StandardLedgerUID{}
			}
			if err := m.LedgerUID.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.LedgerUID = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitNormative(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentNormative
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
func (m *StandardBallotAddition) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadNormative
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
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Addition = append(m.Addition[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Addition == nil {
				m.Addition = []byte{}
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Level = 0
			if (idxNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = int64(encoding_binary.LittleEndian.Uint64(dAtA[idxNdEx:]))
			idxNdEx += 8
		case 3:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			if (idxNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Cycle = int64(encoding_binary.LittleEndian.Uint64(dAtA[idxNdEx:]))
			idxNdEx += 8
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadNormative
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
				return ErrCorruptExtentNormative
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentNormative
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.SeriesUid = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitNormative(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentNormative
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
func omitNormative(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadNormative
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
					return 0, ErrIntegerOverloadNormative
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
					return 0, ErrIntegerOverloadNormative
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
				return 0, ErrCorruptExtentNormative
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterNormative
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentNormative
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentNormative        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadNormative          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterNormative = fmt.Errorf("REDACTED")
)
