//
//

package kinds

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

//
type LedgerUIDMark int32

const (
	LedgerUIDMarkUnclear LedgerUIDMark = 0
	LedgerUIDMarkMissing  LedgerUIDMark = 1
	LedgerUIDMarkEndorse  LedgerUIDMark = 2
	LedgerUIDMarkNull     LedgerUIDMark = 3
)

var Ledgeruidmark_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
}

var Ledgeruidmark_item = map[string]int32{
	"REDACTED": 0,
	"REDACTED":  1,
	"REDACTED":  2,
	"REDACTED":     3,
}

func (x LedgerUIDMark) String() string {
	return proto.EnumName(Ledgeruidmark_label, int32(x))
}

func (LedgerUIDMark) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_4e92274df03d3088, []int{0}
}

type RatifierAssign struct {
	Ratifiers       []*Ratifier `protobuf:"octets,1,rep,name=validators,proto3" json:"ratifiers,omitempty"`
	Recommender         *Ratifier   `protobuf:"octets,2,opt,name=proposer,proto3" json:"recommender,omitempty"`
	SumPollingEnergy int64        `protobuf:"variableint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_polling_energy,omitempty"`
}

func (m *RatifierAssign) Restore()         { *m = RatifierAssign{} }
func (m *RatifierAssign) String() string { return proto.CompactTextString(m) }
func (*RatifierAssign) SchemaSignal()    {}
func (*RatifierAssign) Definition() ([]byte, []int) {
	return filedefinition_4e92274df03d3088, []int{0}
}
func (m *RatifierAssign) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *RatifierAssign) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ballotset.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RatifierAssign) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ballotset.Merge(m, src)
}
func (m *RatifierAssign) XXX_Volume() int {
	return m.Volume()
}
func (m *RatifierAssign) XXX_Omitunclear() {
	xxx_messagedata_Ballotset.DiscardUnknown(m)
}

var xxx_messagedata_Ballotset proto.InternalMessageInfo

func (m *RatifierAssign) FetchRatifiers() []*Ratifier {
	if m != nil {
		return m.Ratifiers
	}
	return nil
}

func (m *RatifierAssign) FetchRecommender() *Ratifier {
	if m != nil {
		return m.Recommender
	}
	return nil
}

func (m *RatifierAssign) FetchSumPollingEnergy() int64 {
	if m != nil {
		return m.SumPollingEnergy
	}
	return 0
}

type Ratifier struct {
	Location          []byte           `protobuf:"octets,1,opt,name=address,proto3" json:"location,omitempty"`
	PublicKey           vault.PublicKey `protobuf:"octets,2,opt,name=pub_key,json=pubKey,proto3" json:"public_key"`
	PollingEnergy      int64            `protobuf:"variableint,3,opt,name=voting_power,json=votingPower,proto3" json:"polling_energy,omitempty"`
	RecommenderUrgency int64            `protobuf:"variableint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"recommender_urgency,omitempty"`
}

func (m *Ratifier) Restore()         { *m = Ratifier{} }
func (m *Ratifier) String() string { return proto.CompactTextString(m) }
func (*Ratifier) SchemaSignal()    {}
func (*Ratifier) Definition() ([]byte, []int) {
	return filedefinition_4e92274df03d3088, []int{1}
}
func (m *Ratifier) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Ratifier) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ratifier.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ratifier) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ratifier.Merge(m, src)
}
func (m *Ratifier) XXX_Volume() int {
	return m.Volume()
}
func (m *Ratifier) XXX_Omitunclear() {
	xxx_messagedata_Ratifier.DiscardUnknown(m)
}

var xxx_messagedata_Ratifier proto.InternalMessageInfo

func (m *Ratifier) FetchLocation() []byte {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Ratifier) FetchPublicKey() vault.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return vault.PublicKey{}
}

func (m *Ratifier) FetchPollingEnergy() int64 {
	if m != nil {
		return m.PollingEnergy
	}
	return 0
}

func (m *Ratifier) FetchRecommenderUrgency() int64 {
	if m != nil {
		return m.RecommenderUrgency
	}
	return 0
}

type BasicRatifier struct {
	PublicKey      *vault.PublicKey `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_key,omitempty"`
	PollingEnergy int64             `protobuf:"variableint,2,opt,name=voting_power,json=votingPower,proto3" json:"polling_energy,omitempty"`
}

func (m *BasicRatifier) Restore()         { *m = BasicRatifier{} }
func (m *BasicRatifier) String() string { return proto.CompactTextString(m) }
func (*BasicRatifier) SchemaSignal()    {}
func (*BasicRatifier) Definition() ([]byte, []int) {
	return filedefinition_4e92274df03d3088, []int{2}
}
func (m *BasicRatifier) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *BasicRatifier) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Basicballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasicRatifier) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Basicballot.Merge(m, src)
}
func (m *BasicRatifier) XXX_Volume() int {
	return m.Volume()
}
func (m *BasicRatifier) XXX_Omitunclear() {
	xxx_messagedata_Basicballot.DiscardUnknown(m)
}

var xxx_messagedata_Basicballot proto.InternalMessageInfo

func (m *BasicRatifier) FetchPublicKey() *vault.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *BasicRatifier) FetchPollingEnergy() int64 {
	if m != nil {
		return m.PollingEnergy
	}
	return 0
}

func init() {
	proto.RegisterEnum("REDACTED", Ledgeruidmark_label, Ledgeruidmark_item)
	proto.RegisterType((*RatifierAssign)(nil), "REDACTED")
	proto.RegisterType((*Ratifier)(nil), "REDACTED")
	proto.RegisterType((*BasicRatifier)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_4e92274df03d3088) }

var filedefinition_4e92274df03d3088 = []byte{
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

func (m *RatifierAssign) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RatifierAssign) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *RatifierAssign) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SumPollingEnergy != 0 {
		i = encodeVariableintRatifier(dAtA, i, uint64(m.SumPollingEnergy))
		i--
		dAtA[i] = 0x18
	}
	if m.Recommender != nil {
		{
			volume, err := m.Recommender.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintRatifier(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Ratifiers) > 0 {
		for idxNdEx := len(m.Ratifiers) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Ratifiers[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = encodeVariableintRatifier(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Ratifier) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ratifier) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Ratifier) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RecommenderUrgency != 0 {
		i = encodeVariableintRatifier(dAtA, i, uint64(m.RecommenderUrgency))
		i--
		dAtA[i] = 0x20
	}
	if m.PollingEnergy != 0 {
		i = encodeVariableintRatifier(dAtA, i, uint64(m.PollingEnergy))
		i--
		dAtA[i] = 0x18
	}
	{
		volume, err := m.PublicKey.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = encodeVariableintRatifier(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVariableintRatifier(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BasicRatifier) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicRatifier) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *BasicRatifier) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PollingEnergy != 0 {
		i = encodeVariableintRatifier(dAtA, i, uint64(m.PollingEnergy))
		i--
		dAtA[i] = 0x10
	}
	if m.PublicKey != nil {
		{
			volume, err := m.PublicKey.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintRatifier(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVariableintRatifier(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovRatifier(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *RatifierAssign) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ratifiers) > 0 {
		for _, e := range m.Ratifiers {
			l = e.Volume()
			n += 1 + l + sovRatifier(uint64(l))
		}
	}
	if m.Recommender != nil {
		l = m.Recommender.Volume()
		n += 1 + l + sovRatifier(uint64(l))
	}
	if m.SumPollingEnergy != 0 {
		n += 1 + sovRatifier(uint64(m.SumPollingEnergy))
	}
	return n
}

func (m *Ratifier) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovRatifier(uint64(l))
	}
	l = m.PublicKey.Volume()
	n += 1 + l + sovRatifier(uint64(l))
	if m.PollingEnergy != 0 {
		n += 1 + sovRatifier(uint64(m.PollingEnergy))
	}
	if m.RecommenderUrgency != 0 {
		n += 1 + sovRatifier(uint64(m.RecommenderUrgency))
	}
	return n
}

func (m *BasicRatifier) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKey != nil {
		l = m.PublicKey.Volume()
		n += 1 + l + sovRatifier(uint64(l))
	}
	if m.PollingEnergy != 0 {
		n += 1 + sovRatifier(uint64(m.PollingEnergy))
	}
	return n
}

func sovRatifier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRatifier(x uint64) (n int) {
	return sovRatifier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RatifierAssign) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadRatifier
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
					return ErrIntegerOverloadRatifier
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
				return ErrCorruptExtentRatifier
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentRatifier
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Ratifiers = append(m.Ratifiers, &Ratifier{})
			if err := m.Ratifiers[len(m.Ratifiers)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadRatifier
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
				return ErrCorruptExtentRatifier
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentRatifier
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Recommender == nil {
				m.Recommender = &Ratifier{}
			}
			if err := m.Recommender.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SumPollingEnergy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadRatifier
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.SumPollingEnergy |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitRatifier(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentRatifier
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
func (m *Ratifier) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadRatifier
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
					return ErrIntegerOverloadRatifier
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
				return ErrCorruptExtentRatifier
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentRatifier
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = append(m.Location[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Location == nil {
				m.Location = []byte{}
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadRatifier
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
				return ErrCorruptExtentRatifier
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentRatifier
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PollingEnergy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadRatifier
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.PollingEnergy |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.RecommenderUrgency = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadRatifier
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.RecommenderUrgency |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitRatifier(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentRatifier
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
func (m *BasicRatifier) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadRatifier
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
					return ErrIntegerOverloadRatifier
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
				return ErrCorruptExtentRatifier
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentRatifier
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &vault.PublicKey{}
			}
			if err := m.PublicKey.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PollingEnergy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadRatifier
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.PollingEnergy |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitRatifier(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentRatifier
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
func omitRatifier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadRatifier
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
					return 0, ErrIntegerOverloadRatifier
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
					return 0, ErrIntegerOverloadRatifier
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
				return 0, ErrCorruptExtentRatifier
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterRatifier
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentRatifier
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentRatifier        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadRatifier          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterRatifier = fmt.Errorf("REDACTED")
)
