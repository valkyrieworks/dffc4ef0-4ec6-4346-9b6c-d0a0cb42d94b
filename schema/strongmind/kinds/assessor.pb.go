//
//

package kinds

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

//
type LedgerUUIDMarker int32

const (
	LedgerUUIDMarkerUnfamiliar LedgerUUIDMarker = 0
	LedgerUUIDMarkerMissing  LedgerUUIDMarker = 1
	LedgerUUIDMarkerEndorse  LedgerUUIDMarker = 2
	LedgerUUIDMarkerVoid     LedgerUUIDMarker = 3
)

var Ledgeruuidmarker_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
}

var Ledgeruuidmarker_datum = map[string]int32{
	"REDACTED": 0,
	"REDACTED":  1,
	"REDACTED":  2,
	"REDACTED":     3,
}

func (x LedgerUUIDMarker) Text() string {
	return proto.EnumName(Ledgeruuidmarker_alias, int32(x))
}

func (LedgerUUIDMarker) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_4e92274df03d3088, []int{0}
}

type AssessorAssign struct {
	Assessors       []*Assessor `protobuf:"octets,1,rep,name=validators,proto3" json:"assessors,omitempty"`
	Nominator         *Assessor   `protobuf:"octets,2,opt,name=proposer,proto3" json:"nominator,omitempty"`
	SumBallotingPotency int64        `protobuf:"variableint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_balloting_potency,omitempty"`
}

func (m *AssessorAssign) Restore()         { *m = AssessorAssign{} }
func (m *AssessorAssign) Text() string { return proto.CompactTextString(m) }
func (*AssessorAssign) SchemaArtifact()    {}
func (*AssessorAssign) Definition() ([]byte, []int) {
	return filedescriptor_4e92274df03d3088, []int{0}
}
func (m *AssessorAssign) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AssessorAssign) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Assessorgroup.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssessorAssign) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Assessorgroup.Merge(m, src)
}
func (m *AssessorAssign) XXX_Extent() int {
	return m.Extent()
}
func (m *AssessorAssign) XXX_Dropunfamiliar() {
	xxx_signaldetails_Assessorgroup.DiscardUnknown(m)
}

var xxx_signaldetails_Assessorgroup proto.InternalMessageInfo

func (m *AssessorAssign) ObtainAssessors() []*Assessor {
	if m != nil {
		return m.Assessors
	}
	return nil
}

func (m *AssessorAssign) ObtainNominator() *Assessor {
	if m != nil {
		return m.Nominator
	}
	return nil
}

func (m *AssessorAssign) ObtainSumBallotingPotency() int64 {
	if m != nil {
		return m.SumBallotingPotency
	}
	return 0
}

type Assessor struct {
	Location          []byte           `protobuf:"octets,1,opt,name=address,proto3" json:"location,omitempty"`
	PublicToken           security.CommonToken `protobuf:"octets,2,opt,name=pub_key,json=pubKey,proto3" json:"public_token"`
	BallotingPotency      int64            `protobuf:"variableint,3,opt,name=voting_power,json=votingPower,proto3" json:"balloting_potency,omitempty"`
	NominatorUrgency int64            `protobuf:"variableint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"nominator_urgency,omitempty"`
}

func (m *Assessor) Restore()         { *m = Assessor{} }
func (m *Assessor) Text() string { return proto.CompactTextString(m) }
func (*Assessor) SchemaArtifact()    {}
func (*Assessor) Definition() ([]byte, []int) {
	return filedescriptor_4e92274df03d3088, []int{1}
}
func (m *Assessor) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Assessor) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Assessor.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Assessor) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Assessor.Merge(m, src)
}
func (m *Assessor) XXX_Extent() int {
	return m.Extent()
}
func (m *Assessor) XXX_Dropunfamiliar() {
	xxx_signaldetails_Assessor.DiscardUnknown(m)
}

var xxx_signaldetails_Assessor proto.InternalMessageInfo

func (m *Assessor) ObtainLocator() []byte {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Assessor) ObtainPublicToken() security.CommonToken {
	if m != nil {
		return m.PublicToken
	}
	return security.CommonToken{}
}

func (m *Assessor) ObtainBallotingPotency() int64 {
	if m != nil {
		return m.BallotingPotency
	}
	return 0
}

func (m *Assessor) ObtainNominatorUrgency() int64 {
	if m != nil {
		return m.NominatorUrgency
	}
	return 0
}

type PlainAssessor struct {
	PublicToken      *security.CommonToken `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_token,omitempty"`
	BallotingPotency int64             `protobuf:"variableint,2,opt,name=voting_power,json=votingPower,proto3" json:"balloting_potency,omitempty"`
}

func (m *PlainAssessor) Restore()         { *m = PlainAssessor{} }
func (m *PlainAssessor) Text() string { return proto.CompactTextString(m) }
func (*PlainAssessor) SchemaArtifact()    {}
func (*PlainAssessor) Definition() ([]byte, []int) {
	return filedescriptor_4e92274df03d3088, []int{2}
}
func (m *PlainAssessor) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PlainAssessor) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Simpleassessor.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlainAssessor) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Simpleassessor.Merge(m, src)
}
func (m *PlainAssessor) XXX_Extent() int {
	return m.Extent()
}
func (m *PlainAssessor) XXX_Dropunfamiliar() {
	xxx_signaldetails_Simpleassessor.DiscardUnknown(m)
}

var xxx_signaldetails_Simpleassessor proto.InternalMessageInfo

func (m *PlainAssessor) ObtainPublicToken() *security.CommonToken {
	if m != nil {
		return m.PublicToken
	}
	return nil
}

func (m *PlainAssessor) ObtainBallotingPotency() int64 {
	if m != nil {
		return m.BallotingPotency
	}
	return 0
}

func initialize() {
	proto.RegisterEnum("REDACTED", Ledgeruuidmarker_alias, Ledgeruuidmarker_datum)
	proto.RegisterType((*AssessorAssign)(nil), "REDACTED")
	proto.RegisterType((*Assessor)(nil), "REDACTED")
	proto.RegisterType((*PlainAssessor)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_4e92274df03d3088) }

var filedescriptor_4e92274df03d3088 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x3d, 0x4d, 0xd5, 0x96, 0x49, 0x04, 0xce, 0xa8, 0x45, 0x91, 0xa9, 0x8c, 0xe9, 0x2a,
	0xfc, 0xc8, 0x16, 0x54, 0x88, 0x45, 0x57, 0x49, 0x4a, 0x51, 0x94, 0xc4, 0x89, 0x92, 0xb6, 0x48,
	0x6c, 0xac, 0x38, 0x19, 0xcc, 0xc8, 0x3f, 0x33, 0x1a, 0x4f, 0x52, 0xf9, 0x0d, 0x50, 0x56, 0xbc,
	0x40, 0x56, 0xb0, 0x60, 0xcd, 0x53, 0x74, 0xd9, 0x1d, 0xac, 0x10, 0x4a, 0x5e, 0x04, 0xc5, 0x6e,
	0x62, 0xb7, 0x01, 0x75, 0x77, 0x7d, 0xcf, 0x39, 0xf7, 0x7e, 0x1e, 0xe9, 0x42, 0x4d, 0xe0, 0x60,
	0x88, 0xb9, 0x4f, 0x02, 0x61, 0x88, 0x88, 0xe1, 0xd0, 0x18, 0xf7, 0x3d, 0x32, 0xec, 0x0b, 0xca,
	0x75, 0xc6, 0xa9, 0xa0, 0x48, 0x4e, 0x1d, 0x7a, 0xec, 0x50, 0x76, 0x1d, 0xea, 0xd0, 0x58, 0x34,
	0x16, 0x55, 0xe2, 0x53, 0xf6, 0x33, 0x93, 0x06, 0x3c, 0x62, 0x82, 0x1a, 0x2e, 0x8e, 0xc2, 0x44,
	0x3d, 0xf8, 0x01, 0x60, 0xe1, 0x7c, 0x39, 0xb9, 0x87, 0x05, 0x3a, 0x82, 0x70, 0xb5, 0x29, 0x2c,
	0x01, 0x2d, 0x57, 0xce, 0xbf, 0x7a, 0xa4, 0xdf, 0xde, 0xa5, 0xaf, 0x32, 0xdd, 0x8c, 0x1d, 0xbd,
	0x81, 0x3b, 0x8c, 0x53, 0x46, 0x43, 0xcc, 0x4b, 0x1b, 0x1a, 0xb8, 0x2b, 0xba, 0x32, 0xa3, 0x17,
	0x10, 0x09, 0x2a, 0xfa, 0x9e, 0x35, 0xa6, 0x82, 0x04, 0x8e, 0xc5, 0xe8, 0x05, 0xe6, 0xa5, 0x9c,
	0x06, 0xca, 0xb9, 0xae, 0x1c, 0x2b, 0xe7, 0xb1, 0xd0, 0x59, 0xf4, 0x17, 0xd0, 0xf7, 0x56, 0x53,
	0x50, 0x09, 0x6e, 0xf7, 0x87, 0x43, 0x8e, 0xc3, 0x05, 0x2e, 0x28, 0x17, 0xba, 0xcb, 0x4f, 0x74,
	0x04, 0xb7, 0xd9, 0xc8, 0xb6, 0x5c, 0x1c, 0x5d, 0xd3, 0xec, 0x67, 0x69, 0x92, 0xc7, 0xd0, 0x3b,
	0x23, 0xdb, 0x23, 0x83, 0x06, 0x8e, 0xaa, 0x9b, 0x97, 0xbf, 0x1f, 0x4b, 0xdd, 0x2d, 0x36, 0xb2,
	0x1b, 0x38, 0x42, 0x4f, 0x60, 0xe1, 0x1f, 0x30, 0xf9, 0x71, 0xca, 0x81, 0x9e, 0xc3, 0xe2, 0xf2,
	0x0f, 0x2c, 0xc6, 0x09, 0xe5, 0x44, 0x44, 0xa5, 0xcd, 0x04, 0x7a, 0x29, 0x74, 0xae, 0xfb, 0x07,
	0x2e, 0x7c, 0xd0, 0x23, 0x3e, 0xf3, 0x70, 0x4a, 0xfe, 0x3a, 0xe5, 0x03, 0x77, 0xf3, 0xfd, 0x97,
	0x6c, 0x63, 0x8d, 0xec, 0xd9, 0x4f, 0x00, 0xf3, 0x55, 0x8f, 0x0e, 0xdc, 0xfa, 0xf1, 0x89, 0xd7,
	0x77, 0xd0, 0x4b, 0xb8, 0x57, 0x6d, 0xb6, 0x6b, 0x0d, 0xab, 0x7e, 0x6c, 0x9d, 0x34, 0x2b, 0xef,
	0xac, 0x33, 0xb3, 0x61, 0xb6, 0xdf, 0x9b, 0xb2, 0xa4, 0x3c, 0x9c, 0x4c, 0x35, 0x94, 0xf1, 0x9e,
	0x05, 0x6e, 0x40, 0x2f, 0x02, 0x64, 0xc0, 0xdd, 0x9b, 0x91, 0x4a, 0xb5, 0xf7, 0xd6, 0x3c, 0x95,
	0x81, 0xb2, 0x37, 0x99, 0x6a, 0xc5, 0x4c, 0xa2, 0x62, 0x87, 0x38, 0x10, 0xeb, 0x81, 0x5a, 0xbb,
	0xd5, 0xaa, 0x9f, 0xca, 0x1b, 0x6b, 0x81, 0x1a, 0xf5, 0x7d, 0x22, 0xd0, 0x53, 0x58, 0xbc, 0x19,
	0x30, 0xeb, 0x4d, 0x39, 0xa7, 0xa0, 0xc9, 0x54, 0xbb, 0x9f, 0x71, 0x9b, 0xc4, 0x53, 0x76, 0x3e,
	0x7f, 0x55, 0xa5, 0xef, 0xdf, 0x54, 0x50, 0x6d, 0x5d, 0xce, 0x54, 0x70, 0x35, 0x53, 0xc1, 0x9f,
	0x99, 0x0a, 0xbe, 0xcc, 0x55, 0xe9, 0x6a, 0xae, 0x4a, 0xbf, 0xe6, 0xaa, 0xf4, 0xe1, 0xd0, 0x21,
	0xe2, 0xd3, 0xc8, 0xd6, 0x07, 0xd4, 0x37, 0x06, 0xd4, 0xc7, 0xc2, 0xfe, 0x28, 0xd2, 0x22, 0xb9,
	0x8b, 0xdb, 0x57, 0x65, 0x6f, 0xc5, 0xfd, 0xc3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x8c,
	0xe3, 0x20, 0x70, 0x03, 0x00, 0x00,
}

func (m *AssessorAssign) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AssessorAssign) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AssessorAssign) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.SumBallotingPotency != 0 {
		i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(m.SumBallotingPotency))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Nominator != nil {
		{
			extent, err := m.Nominator.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Assessors) > 0 {
		for idxNdExc := len(m.Assessors) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Assessors[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Assessor) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Assessor) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Assessor) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.NominatorUrgency != 0 {
		i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(m.NominatorUrgency))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.BallotingPotency != 0 {
		i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(m.BallotingPotency))
		i--
		deltaLocatedAN[i] = 0x18
	}
	{
		extent, err := m.PublicToken.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x12
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(deltaLocatedAN[i:], m.Location)
		i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(len(m.Location)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *PlainAssessor) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PlainAssessor) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PlainAssessor) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.BallotingPotency != 0 {
		i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(m.BallotingPotency))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.PublicToken != nil {
		{
			extent, err := m.PublicToken.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintAssessor(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintAssessor(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovAssessor(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *AssessorAssign) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Assessors) > 0 {
		for _, e := range m.Assessors {
			l = e.Extent()
			n += 1 + l + sovAssessor(uint64(l))
		}
	}
	if m.Nominator != nil {
		l = m.Nominator.Extent()
		n += 1 + l + sovAssessor(uint64(l))
	}
	if m.SumBallotingPotency != 0 {
		n += 1 + sovAssessor(uint64(m.SumBallotingPotency))
	}
	return n
}

func (m *Assessor) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovAssessor(uint64(l))
	}
	l = m.PublicToken.Extent()
	n += 1 + l + sovAssessor(uint64(l))
	if m.BallotingPotency != 0 {
		n += 1 + sovAssessor(uint64(m.BallotingPotency))
	}
	if m.NominatorUrgency != 0 {
		n += 1 + sovAssessor(uint64(m.NominatorUrgency))
	}
	return n
}

func (m *PlainAssessor) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicToken != nil {
		l = m.PublicToken.Extent()
		n += 1 + l + sovAssessor(uint64(l))
	}
	if m.BallotingPotency != 0 {
		n += 1 + sovAssessor(uint64(m.BallotingPotency))
	}
	return n
}

func sovAssessor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAssessor(x uint64) (n int) {
	return sovAssessor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssessorAssign) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAssessor
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
					return FaultIntegerOverrunAssessor
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
				return FaultUnfitMagnitudeAssessor
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAssessor
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Assessors = append(m.Assessors, &Assessor{})
			if err := m.Assessors[len(m.Assessors)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAssessor
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
				return FaultUnfitMagnitudeAssessor
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAssessor
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nominator == nil {
				m.Nominator = &Assessor{}
			}
			if err := m.Nominator.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SumBallotingPotency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAssessor
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.SumBallotingPotency |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitAssessor(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAssessor
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
func (m *Assessor) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAssessor
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
					return FaultIntegerOverrunAssessor
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
				return FaultUnfitMagnitudeAssessor
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAssessor
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = append(m.Location[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Location == nil {
				m.Location = []byte{}
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAssessor
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
				return FaultUnfitMagnitudeAssessor
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAssessor
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicToken.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.BallotingPotency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAssessor
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.BallotingPotency |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.NominatorUrgency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAssessor
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.NominatorUrgency |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitAssessor(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAssessor
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
func (m *PlainAssessor) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunAssessor
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
					return FaultIntegerOverrunAssessor
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
				return FaultUnfitMagnitudeAssessor
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeAssessor
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicToken == nil {
				m.PublicToken = &security.CommonToken{}
			}
			if err := m.PublicToken.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.BallotingPotency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunAssessor
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.BallotingPotency |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitAssessor(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeAssessor
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
func omitAssessor(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunAssessor
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
					return 0, FaultIntegerOverrunAssessor
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
					return 0, FaultIntegerOverrunAssessor
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
				return 0, FaultUnfitMagnitudeAssessor
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionAssessor
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeAssessor
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeAssessor        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunAssessor          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionAssessor = fmt.Errorf("REDACTED")
)
