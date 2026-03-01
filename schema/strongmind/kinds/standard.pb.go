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

type StandardLedgerUUID struct {
	Digest          []byte                 `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	FragmentAssignHeading StandardFragmentAssignHeading `protobuf:"octets,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"fragment_assign_headline"`
}

func (m *StandardLedgerUUID) Restore()         { *m = StandardLedgerUUID{} }
func (m *StandardLedgerUUID) Text() string { return proto.CompactTextString(m) }
func (*StandardLedgerUUID) SchemaArtifact()    {}
func (*StandardLedgerUUID) Definition() ([]byte, []int) {
	return filedescriptor_8d1a1a84ff7267ed, []int{0}
}
func (m *StandardLedgerUUID) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *StandardLedgerUUID) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Standardledgeruuid.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardLedgerUUID) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Standardledgeruuid.Merge(m, src)
}
func (m *StandardLedgerUUID) XXX_Extent() int {
	return m.Extent()
}
func (m *StandardLedgerUUID) XXX_Dropunfamiliar() {
	xxx_signaldetails_Standardledgeruuid.DiscardUnknown(m)
}

var xxx_signaldetails_Standardledgeruuid proto.InternalMessageInfo

func (m *StandardLedgerUUID) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *StandardLedgerUUID) ObtainFragmentAssignHeadline() StandardFragmentAssignHeading {
	if m != nil {
		return m.FragmentAssignHeading
	}
	return StandardFragmentAssignHeading{}
}

type StandardFragmentAssignHeading struct {
	Sum uint32 `protobuf:"variableint,1,opt,name=total,proto3" json:"sum,omitempty"`
	Digest  []byte `protobuf:"octets,2,opt,name=hash,proto3" json:"digest,omitempty"`
}

func (m *StandardFragmentAssignHeading) Restore()         { *m = StandardFragmentAssignHeading{} }
func (m *StandardFragmentAssignHeading) Text() string { return proto.CompactTextString(m) }
func (*StandardFragmentAssignHeading) SchemaArtifact()    {}
func (*StandardFragmentAssignHeading) Definition() ([]byte, []int) {
	return filedescriptor_8d1a1a84ff7267ed, []int{1}
}
func (m *StandardFragmentAssignHeading) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *StandardFragmentAssignHeading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Standardfragmentsetheadline.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardFragmentAssignHeading) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Standardfragmentsetheadline.Merge(m, src)
}
func (m *StandardFragmentAssignHeading) XXX_Extent() int {
	return m.Extent()
}
func (m *StandardFragmentAssignHeading) XXX_Dropunfamiliar() {
	xxx_signaldetails_Standardfragmentsetheadline.DiscardUnknown(m)
}

var xxx_signaldetails_Standardfragmentsetheadline proto.InternalMessageInfo

func (m *StandardFragmentAssignHeading) ObtainSum() uint32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *StandardFragmentAssignHeading) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type StandardNomination struct {
	Kind      AttestedSignalKind     `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Altitude    int64             `protobuf:"fixed64,2,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration     int64             `protobuf:"fixed64,3,opt,name=round,proto3" json:"iteration,omitempty"`
	PolicyIteration  int64             `protobuf:"variableint,4,opt,name=pol_round,json=polRound,proto3" json:"policy_iteration,omitempty"`
	LedgerUUID   *StandardLedgerUUID `protobuf:"octets,5,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid,omitempty"`
	Timestamp time.Time         `protobuf:"octets,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	SuccessionUUID   string            `protobuf:"octets,7,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
}

func (m *StandardNomination) Restore()         { *m = StandardNomination{} }
func (m *StandardNomination) Text() string { return proto.CompactTextString(m) }
func (*StandardNomination) SchemaArtifact()    {}
func (*StandardNomination) Definition() ([]byte, []int) {
	return filedescriptor_8d1a1a84ff7267ed, []int{2}
}
func (m *StandardNomination) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *StandardNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Standardnomination.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardNomination) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Standardnomination.Merge(m, src)
}
func (m *StandardNomination) XXX_Extent() int {
	return m.Extent()
}
func (m *StandardNomination) XXX_Dropunfamiliar() {
	xxx_signaldetails_Standardnomination.DiscardUnknown(m)
}

var xxx_signaldetails_Standardnomination proto.InternalMessageInfo

func (m *StandardNomination) ObtainKind() AttestedSignalKind {
	if m != nil {
		return m.Kind
	}
	return UnfamiliarKind
}

func (m *StandardNomination) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *StandardNomination) ObtainIteration() int64 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *StandardNomination) ObtainPolicyIteration() int64 {
	if m != nil {
		return m.PolicyIteration
	}
	return 0
}

func (m *StandardNomination) ObtainLedgerUUID() *StandardLedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return nil
}

func (m *StandardNomination) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *StandardNomination) ObtainSuccessionUUID() string {
	if m != nil {
		return m.SuccessionUUID
	}
	return "REDACTED"
}

type StandardBallot struct {
	Kind      AttestedSignalKind     `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Altitude    int64             `protobuf:"fixed64,2,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration     int64             `protobuf:"fixed64,3,opt,name=round,proto3" json:"iteration,omitempty"`
	LedgerUUID   *StandardLedgerUUID `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid,omitempty"`
	Timestamp time.Time         `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	SuccessionUUID   string            `protobuf:"octets,6,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
}

func (m *StandardBallot) Restore()         { *m = StandardBallot{} }
func (m *StandardBallot) Text() string { return proto.CompactTextString(m) }
func (*StandardBallot) SchemaArtifact()    {}
func (*StandardBallot) Definition() ([]byte, []int) {
	return filedescriptor_8d1a1a84ff7267ed, []int{3}
}
func (m *StandardBallot) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *StandardBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Standardballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardBallot) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Standardballot.Merge(m, src)
}
func (m *StandardBallot) XXX_Extent() int {
	return m.Extent()
}
func (m *StandardBallot) XXX_Dropunfamiliar() {
	xxx_signaldetails_Standardballot.DiscardUnknown(m)
}

var xxx_signaldetails_Standardballot proto.InternalMessageInfo

func (m *StandardBallot) ObtainKind() AttestedSignalKind {
	if m != nil {
		return m.Kind
	}
	return UnfamiliarKind
}

func (m *StandardBallot) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *StandardBallot) ObtainIteration() int64 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *StandardBallot) ObtainLedgerUUID() *StandardLedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return nil
}

func (m *StandardBallot) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *StandardBallot) ObtainSuccessionUUID() string {
	if m != nil {
		return m.SuccessionUUID
	}
	return "REDACTED"
}

//
//
type StandardBallotAddition struct {
	Addition []byte `protobuf:"octets,1,opt,name=extension,proto3" json:"addition,omitempty"`
	Altitude    int64  `protobuf:"fixed64,2,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration     int64  `protobuf:"fixed64,3,opt,name=round,proto3" json:"iteration,omitempty"`
	SuccessionUuid   string `protobuf:"octets,4,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
}

func (m *StandardBallotAddition) Restore()         { *m = StandardBallotAddition{} }
func (m *StandardBallotAddition) Text() string { return proto.CompactTextString(m) }
func (*StandardBallotAddition) SchemaArtifact()    {}
func (*StandardBallotAddition) Definition() ([]byte, []int) {
	return filedescriptor_8d1a1a84ff7267ed, []int{4}
}
func (m *StandardBallotAddition) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *StandardBallotAddition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Standardballotaddition.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandardBallotAddition) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Standardballotaddition.Merge(m, src)
}
func (m *StandardBallotAddition) XXX_Extent() int {
	return m.Extent()
}
func (m *StandardBallotAddition) XXX_Dropunfamiliar() {
	xxx_signaldetails_Standardballotaddition.DiscardUnknown(m)
}

var xxx_signaldetails_Standardballotaddition proto.InternalMessageInfo

func (m *StandardBallotAddition) ObtainAddition() []byte {
	if m != nil {
		return m.Addition
	}
	return nil
}

func (m *StandardBallotAddition) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *StandardBallotAddition) ObtainIteration() int64 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *StandardBallotAddition) ObtainSuccessionUuid() string {
	if m != nil {
		return m.SuccessionUuid
	}
	return "REDACTED"
}

func initialize() {
	proto.RegisterType((*StandardLedgerUUID)(nil), "REDACTED")
	proto.RegisterType((*StandardFragmentAssignHeading)(nil), "REDACTED")
	proto.RegisterType((*StandardNomination)(nil), "REDACTED")
	proto.RegisterType((*StandardBallot)(nil), "REDACTED")
	proto.RegisterType((*StandardBallotAddition)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_8d1a1a84ff7267ed) }

var filedescriptor_8d1a1a84ff7267ed = []byte{
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

func (m *StandardLedgerUUID) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *StandardLedgerUUID) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *StandardLedgerUUID) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.FragmentAssignHeading.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x12
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *StandardFragmentAssignHeading) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *StandardFragmentAssignHeading) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *StandardFragmentAssignHeading) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Sum != 0 {
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(m.Sum))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *StandardNomination) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *StandardNomination) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *StandardNomination) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.SuccessionUUID) > 0 {
		i -= len(m.SuccessionUUID)
		copy(deltaLocatedAN[i:], m.SuccessionUUID)
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(len(m.SuccessionUUID)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	n2, fault2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault2 != nil {
		return 0, fault2
	}
	i -= n2
	i = encodeVariableintStandard(deltaLocatedAN, i, uint64(n2))
	i--
	deltaLocatedAN[i] = 0x32
	if m.LedgerUUID != nil {
		{
			extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintStandard(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if m.PolicyIteration != 0 {
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(m.PolicyIteration))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.Iteration != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(deltaLocatedAN[i:], uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x19
	}
	if m.Altitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(deltaLocatedAN[i:], uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x11
	}
	if m.Kind != 0 {
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *StandardBallot) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *StandardBallot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *StandardBallot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.SuccessionUUID) > 0 {
		i -= len(m.SuccessionUUID)
		copy(deltaLocatedAN[i:], m.SuccessionUUID)
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(len(m.SuccessionUUID)))
		i--
		deltaLocatedAN[i] = 0x32
	}
	n4, fault4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault4 != nil {
		return 0, fault4
	}
	i -= n4
	i = encodeVariableintStandard(deltaLocatedAN, i, uint64(n4))
	i--
	deltaLocatedAN[i] = 0x2a
	if m.LedgerUUID != nil {
		{
			extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintStandard(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Iteration != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(deltaLocatedAN[i:], uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x19
	}
	if m.Altitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(deltaLocatedAN[i:], uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x11
	}
	if m.Kind != 0 {
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *StandardBallotAddition) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *StandardBallotAddition) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *StandardBallotAddition) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.SuccessionUuid) > 0 {
		i -= len(m.SuccessionUuid)
		copy(deltaLocatedAN[i:], m.SuccessionUuid)
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(len(m.SuccessionUuid)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Iteration != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(deltaLocatedAN[i:], uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x19
	}
	if m.Altitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(deltaLocatedAN[i:], uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x11
	}
	if len(m.Addition) > 0 {
		i -= len(m.Addition)
		copy(deltaLocatedAN[i:], m.Addition)
		i = encodeVariableintStandard(deltaLocatedAN, i, uint64(len(m.Addition)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintStandard(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovStandard(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *StandardLedgerUUID) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovStandard(uint64(l))
	}
	l = m.FragmentAssignHeading.Extent()
	n += 1 + l + sovStandard(uint64(l))
	return n
}

func (m *StandardFragmentAssignHeading) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != 0 {
		n += 1 + sovStandard(uint64(m.Sum))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovStandard(uint64(l))
	}
	return n
}

func (m *StandardNomination) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovStandard(uint64(m.Kind))
	}
	if m.Altitude != 0 {
		n += 9
	}
	if m.Iteration != 0 {
		n += 9
	}
	if m.PolicyIteration != 0 {
		n += 1 + sovStandard(uint64(m.PolicyIteration))
	}
	if m.LedgerUUID != nil {
		l = m.LedgerUUID.Extent()
		n += 1 + l + sovStandard(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovStandard(uint64(l))
	l = len(m.SuccessionUUID)
	if l > 0 {
		n += 1 + l + sovStandard(uint64(l))
	}
	return n
}

func (m *StandardBallot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovStandard(uint64(m.Kind))
	}
	if m.Altitude != 0 {
		n += 9
	}
	if m.Iteration != 0 {
		n += 9
	}
	if m.LedgerUUID != nil {
		l = m.LedgerUUID.Extent()
		n += 1 + l + sovStandard(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovStandard(uint64(l))
	l = len(m.SuccessionUUID)
	if l > 0 {
		n += 1 + l + sovStandard(uint64(l))
	}
	return n
}

func (m *StandardBallotAddition) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addition)
	if l > 0 {
		n += 1 + l + sovStandard(uint64(l))
	}
	if m.Altitude != 0 {
		n += 9
	}
	if m.Iteration != 0 {
		n += 9
	}
	l = len(m.SuccessionUuid)
	if l > 0 {
		n += 1 + l + sovStandard(uint64(l))
	}
	return n
}

func sovStandard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStandard(x uint64) (n int) {
	return sovStandard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StandardLedgerUUID) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunStandard
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
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = append(m.Digest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FragmentAssignHeading.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitStandard(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeStandard
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
func (m *StandardFragmentAssignHeading) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunStandard
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
					return FaultIntegerOverrunStandard
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Sum |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = append(m.Digest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitStandard(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeStandard
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
func (m *StandardNomination) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunStandard
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
			m.Kind = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Kind |= AttestedSignalKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Altitude = 0
			if (idxNdExc + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Altitude = int64(encoding_binary.LittleEndian.Uint64(deltaLocatedAN[idxNdExc:]))
			idxNdExc += 8
		case 3:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			if (idxNdExc + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Iteration = int64(encoding_binary.LittleEndian.Uint64(deltaLocatedAN[idxNdExc:]))
			idxNdExc += 8
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PolicyIteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.PolicyIteration |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.LedgerUUID == nil {
				m.LedgerUUID = &StandardLedgerUUID{}
			}
			if err := m.LedgerUUID.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.SuccessionUUID = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitStandard(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeStandard
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
func (m *StandardBallot) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunStandard
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
			m.Kind = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Kind |= AttestedSignalKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Altitude = 0
			if (idxNdExc + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Altitude = int64(encoding_binary.LittleEndian.Uint64(deltaLocatedAN[idxNdExc:]))
			idxNdExc += 8
		case 3:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			if (idxNdExc + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Iteration = int64(encoding_binary.LittleEndian.Uint64(deltaLocatedAN[idxNdExc:]))
			idxNdExc += 8
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.LedgerUUID == nil {
				m.LedgerUUID = &StandardLedgerUUID{}
			}
			if err := m.LedgerUUID.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.SuccessionUUID = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitStandard(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeStandard
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
func (m *StandardBallotAddition) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunStandard
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
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Addition = append(m.Addition[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Addition == nil {
				m.Addition = []byte{}
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Altitude = 0
			if (idxNdExc + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Altitude = int64(encoding_binary.LittleEndian.Uint64(deltaLocatedAN[idxNdExc:]))
			idxNdExc += 8
		case 3:
			if cableKind != 1 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			if (idxNdExc + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Iteration = int64(encoding_binary.LittleEndian.Uint64(deltaLocatedAN[idxNdExc:]))
			idxNdExc += 8
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunStandard
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
				return FaultUnfitMagnitudeStandard
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeStandard
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.SuccessionUuid = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitStandard(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeStandard
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
func omitStandard(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunStandard
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
					return 0, FaultIntegerOverrunStandard
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
					return 0, FaultIntegerOverrunStandard
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
				return 0, FaultUnfitMagnitudeStandard
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionStandard
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeStandard
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeStandard        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunStandard          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionStandard = fmt.Errorf("REDACTED")
)
