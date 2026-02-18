//
//

package kinds

import (
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

type Proof struct {
	//
	//
	//
	Sum isevidence_Total `protobuf_oneof:"sum"`
}

func (m *Proof) Restore()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) SchemaSignal()    {}
func (*Proof) Definition() ([]byte, []int) {
	return filedefinition_6825fabc78e0a168, []int{0}
}
func (m *Proof) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Proof) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Proof.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Proof.Merge(m, src)
}
func (m *Proof) XXX_Volume() int {
	return m.Volume()
}
func (m *Proof) XXX_Omitunclear() {
	xxx_messagedata_Proof.DiscardUnknown(m)
}

var xxx_messagedata_Proof proto.InternalMessageInfo

type isevidence_Total interface {
	isevidence_Total()
	SerializeTo([]byte) (int, error)
	Volume() int
}

type Proof_Duplicateballotevidence struct {
	ReplicatedBallotProof *ReplicatedBallotProof `protobuf:"octets,1,opt,name=duplicate_vote_evidence,json=duplicateVoteEvidence,proto3,oneof" json:"replicated_ballot_proof,omitempty"`
}
type Proof_Rapidcustomerevidence struct {
	RapidCustomerAssaultProof *RapidCustomerAssaultProof `protobuf:"octets,2,opt,name=light_client_attack_evidence,json=lightClientAttackEvidence,proto3,oneof" json:"rapid_customer_assault_proof,omitempty"`
}

func (*Proof_Duplicateballotevidence) isevidence_Total()     {}
func (*Proof_Rapidcustomerevidence) isevidence_Total() {}

func (m *Proof) FetchTotal() isevidence_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Proof) FetchReplicatedBallotProof() *ReplicatedBallotProof {
	if x, ok := m.FetchTotal().(*Proof_Duplicateballotevidence); ok {
		return x.ReplicatedBallotProof
	}
	return nil
}

func (m *Proof) FetchRapidCustomerAssaultProof() *RapidCustomerAssaultProof {
	if x, ok := m.FetchTotal().(*Proof_Rapidcustomerevidence); ok {
		return x.RapidCustomerAssaultProof
	}
	return nil
}

//
func (*Proof) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Proof_Duplicateballotevidence)(nil),
		(*Proof_Rapidcustomerevidence)(nil),
	}
}

//
type ReplicatedBallotProof struct {
	BallotA            *Ballot     `protobuf:"octets,1,opt,name=vote_a,json=voteA,proto3" json:"ballot_a,omitempty"`
	BallotBYTE            *Ballot     `protobuf:"octets,2,opt,name=vote_b,json=voteB,proto3" json:"ballot_byte,omitempty"`
	SumPollingEnergy int64     `protobuf:"variableint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_polling_energy,omitempty"`
	RatifierEnergy   int64     `protobuf:"variableint,4,opt,name=validator_power,json=validatorPower,proto3" json:"ratifier_energy,omitempty"`
	Timestamp        time.Time `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *ReplicatedBallotProof) Restore()         { *m = ReplicatedBallotProof{} }
func (m *ReplicatedBallotProof) String() string { return proto.CompactTextString(m) }
func (*ReplicatedBallotProof) SchemaSignal()    {}
func (*ReplicatedBallotProof) Definition() ([]byte, []int) {
	return filedefinition_6825fabc78e0a168, []int{1}
}
func (m *ReplicatedBallotProof) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplicatedBallotProof) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Duplicateballotevidence.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplicatedBallotProof) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Duplicateballotevidence.Merge(m, src)
}
func (m *ReplicatedBallotProof) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplicatedBallotProof) XXX_Omitunclear() {
	xxx_messagedata_Duplicateballotevidence.DiscardUnknown(m)
}

var xxx_messagedata_Duplicateballotevidence proto.InternalMessageInfo

func (m *ReplicatedBallotProof) FetchBallotA() *Ballot {
	if m != nil {
		return m.BallotA
	}
	return nil
}

func (m *ReplicatedBallotProof) FetchBallotBYTE() *Ballot {
	if m != nil {
		return m.BallotBYTE
	}
	return nil
}

func (m *ReplicatedBallotProof) FetchSumPollingEnergy() int64 {
	if m != nil {
		return m.SumPollingEnergy
	}
	return 0
}

func (m *ReplicatedBallotProof) FetchRatifierEnergy() int64 {
	if m != nil {
		return m.RatifierEnergy
	}
	return 0
}

func (m *ReplicatedBallotProof) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

//
type RapidCustomerAssaultProof struct {
	ClashingLedger    *RapidLedger  `protobuf:"octets,1,opt,name=conflicting_block,json=conflictingBlock,proto3" json:"clashing_ledger,omitempty"`
	SharedLevel        int64        `protobuf:"variableint,2,opt,name=common_height,json=commonHeight,proto3" json:"shared_level,omitempty"`
	FaultyRatifiers []*Ratifier `protobuf:"octets,3,rep,name=byzantine_validators,json=byzantineValidators,proto3" json:"faulty_ratifiers,omitempty"`
	SumPollingEnergy    int64        `protobuf:"variableint,4,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_polling_energy,omitempty"`
	Timestamp           time.Time    `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *RapidCustomerAssaultProof) Restore()         { *m = RapidCustomerAssaultProof{} }
func (m *RapidCustomerAssaultProof) String() string { return proto.CompactTextString(m) }
func (*RapidCustomerAssaultProof) SchemaSignal()    {}
func (*RapidCustomerAssaultProof) Definition() ([]byte, []int) {
	return filedefinition_6825fabc78e0a168, []int{2}
}
func (m *RapidCustomerAssaultProof) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *RapidCustomerAssaultProof) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Rapidcustomerevidence.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RapidCustomerAssaultProof) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Rapidcustomerevidence.Merge(m, src)
}
func (m *RapidCustomerAssaultProof) XXX_Volume() int {
	return m.Volume()
}
func (m *RapidCustomerAssaultProof) XXX_Omitunclear() {
	xxx_messagedata_Rapidcustomerevidence.DiscardUnknown(m)
}

var xxx_messagedata_Rapidcustomerevidence proto.InternalMessageInfo

func (m *RapidCustomerAssaultProof) FetchClashingLedger() *RapidLedger {
	if m != nil {
		return m.ClashingLedger
	}
	return nil
}

func (m *RapidCustomerAssaultProof) FetchSharedLevel() int64 {
	if m != nil {
		return m.SharedLevel
	}
	return 0
}

func (m *RapidCustomerAssaultProof) FetchFaultyRatifiers() []*Ratifier {
	if m != nil {
		return m.FaultyRatifiers
	}
	return nil
}

func (m *RapidCustomerAssaultProof) FetchSumPollingEnergy() int64 {
	if m != nil {
		return m.SumPollingEnergy
	}
	return 0
}

func (m *RapidCustomerAssaultProof) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type ProofCatalog struct {
	Proof []Proof `protobuf:"octets,1,rep,name=evidence,proto3" json:"proof"`
}

func (m *ProofCatalog) Restore()         { *m = ProofCatalog{} }
func (m *ProofCatalog) String() string { return proto.CompactTextString(m) }
func (*ProofCatalog) SchemaSignal()    {}
func (*ProofCatalog) Definition() ([]byte, []int) {
	return filedefinition_6825fabc78e0a168, []int{3}
}
func (m *ProofCatalog) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ProofCatalog) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Evidencelist.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofCatalog) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Evidencelist.Merge(m, src)
}
func (m *ProofCatalog) XXX_Volume() int {
	return m.Volume()
}
func (m *ProofCatalog) XXX_Omitunclear() {
	xxx_messagedata_Evidencelist.DiscardUnknown(m)
}

var xxx_messagedata_Evidencelist proto.InternalMessageInfo

func (m *ProofCatalog) FetchProof() []Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*Proof)(nil), "REDACTED")
	proto.RegisterType((*ReplicatedBallotProof)(nil), "REDACTED")
	proto.RegisterType((*RapidCustomerAssaultProof)(nil), "REDACTED")
	proto.RegisterType((*ProofCatalog)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_6825fabc78e0a168) }

var filedefinition_6825fabc78e0a168 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x38, 0xa9, 0xc2, 0xb6, 0x40, 0x58, 0x5a, 0x48, 0x43, 0xe4, 0x44, 0xe1, 0xd0,
	0x48, 0x80, 0x2d, 0xb5, 0x57, 0x2e, 0x35, 0x20, 0x15, 0x29, 0x20, 0x64, 0xa1, 0x1e, 0xb8, 0x58,
	0xeb, 0xcd, 0xc6, 0x59, 0xd5, 0xde, 0x8d, 0xe2, 0x49, 0x50, 0x79, 0x8a, 0x3c, 0x56, 0x2f, 0x48,
	0x3d, 0x72, 0x02, 0x94, 0xf0, 0x20, 0xc8, 0xeb, 0x3f, 0x89, 0xea, 0x98, 0x13, 0x97, 0xc8, 0x99,
	0xf9, 0x7d, 0x3b, 0x33, 0x9f, 0x67, 0x8d, 0xba, 0xc0, 0xc4, 0x88, 0xcd, 0x42, 0x2e, 0xc0, 0x82,
	0xeb, 0x29, 0x8b, 0x2c, 0xb6, 0xe0, 0x23, 0x26, 0x28, 0x33, 0xa7, 0x33, 0x09, 0x12, 0x37, 0x37,
	0x80, 0xa9, 0x80, 0xf6, 0xa1, 0x2f, 0x7d, 0xa9, 0x92, 0x56, 0xfc, 0x94, 0x70, 0xed, 0xae, 0x2f,
	0xa5, 0x1f, 0x30, 0x4b, 0xfd, 0xf3, 0xe6, 0x63, 0x0b, 0x78, 0xc8, 0x22, 0x20, 0xe1, 0x34, 0x05,
	0x3a, 0x85, 0x4a, 0xea, 0x37, 0xcd, 0xf6, 0x0a, 0xd9, 0x05, 0x09, 0xf8, 0x88, 0x80, 0x9c, 0x25,
	0x44, 0xff, 0x8f, 0x86, 0x1a, 0xef, 0xd2, 0xde, 0x30, 0x41, 0x4f, 0x47, 0xf3, 0x69, 0xc0, 0x29,
	0x01, 0xe6, 0x2e, 0x24, 0x30, 0x37, 0x6b, 0xbb, 0xa5, 0xf5, 0xb4, 0xc1, 0xfe, 0xe9, 0x89, 0x79,
	0xb7, 0x6f, 0xf3, 0x6d, 0x26, 0xb8, 0x94, 0xc0, 0xb2, 0x93, 0x2e, 0x2a, 0xce, 0xd1, 0x68, 0x57,
	0x02, 0x0b, 0xd4, 0x09, 0xb8, 0x3f, 0x01, 0x97, 0x06, 0x9c, 0x09, 0x70, 0x09, 0x00, 0xa1, 0x57,
	0x9b, 0x3a, 0x55, 0x55, 0xe7, 0x45, 0xb1, 0xce, 0x30, 0x56, 0xbd, 0x51, 0xa2, 0x73, 0xa5, 0xd9,
	0xaa, 0x75, 0x1c, 0x94, 0x25, 0xed, 0x3a, 0xd2, 0xa3, 0x79, 0xd8, 0x5f, 0x56, 0xd1, 0xd1, 0xce,
	0x4e, 0xf1, 0x2b, 0xb4, 0xa7, 0x26, 0x25, 0xe9, 0x88, 0x4f, 0x8a, 0xa5, 0x63, 0xde, 0xa9, 0xc7,
	0xd4, 0x79, 0x8e, 0x7b, 0x69, 0xa7, 0xff, 0xc4, 0x6d, 0xfc, 0x12, 0x61, 0x90, 0x40, 0x82, 0xd8,
	0x4d, 0x2e, 0x7c, 0x77, 0x2a, 0xbf, 0xb2, 0x59, 0x4b, 0xef, 0x69, 0x03, 0xdd, 0x69, 0xaa, 0xcc,
	0xa5, 0x4a, 0x7c, 0x8a, 0xe3, 0xf8, 0x04, 0x3d, 0xcc, 0xdf, 0x4f, 0x8a, 0xd6, 0x14, 0xfa, 0x20,
	0x0f, 0x27, 0xa0, 0x8d, 0xee, 0xe5, 0x8b, 0xd0, 0xaa, 0xab, 0x46, 0xda, 0x66, 0xb2, 0x2a, 0x66,
	0xb6, 0x2a, 0xe6, 0xe7, 0x8c, 0xb0, 0x1b, 0x37, 0x3f, 0xbb, 0x95, 0xe5, 0xaf, 0xae, 0xe6, 0x6c,
	0x64, 0xfd, 0xef, 0x55, 0x74, 0x5c, 0x6a, 0x2a, 0x7e, 0x8f, 0x1e, 0x51, 0x29, 0xc6, 0x01, 0xa7,
	0xaa, 0x6f, 0x2f, 0x90, 0xf4, 0x2a, 0x75, 0xa8, 0x53, 0xf2, 0x72, 0xec, 0x98, 0x71, 0x9a, 0x5b,
	0x32, 0x15, 0xc1, 0xcf, 0xd1, 0x7d, 0x2a, 0xc3, 0x50, 0x0a, 0x77, 0xc2, 0x62, 0x4e, 0x39, 0xa7,
	0x3b, 0x07, 0x49, 0xf0, 0x42, 0xc5, 0xf0, 0x47, 0x74, 0xe8, 0x5d, 0x7f, 0x23, 0x02, 0xb8, 0x60,
	0x6e, 0x3e, 0x6d, 0xd4, 0xd2, 0x7b, 0xfa, 0x60, 0xff, 0xf4, 0xd9, 0x0e, 0x97, 0x33, 0xc6, 0x79,
	0x9c, 0x0b, 0xf3, 0x58, 0x54, 0x62, 0x7c, 0xad, 0xc4, 0xf8, 0xff, 0xe1, 0xe7, 0x10, 0x1d, 0x64,
	0xee, 0x0d, 0x79, 0x04, 0xf8, 0x35, 0x6a, 0x6c, 0xdd, 0x1e, 0x5d, 0x1d, 0x59, 0x98, 0x22, 0xdf,
	0xd3, 0x5a, 0x7c, 0xa4, 0x93, 0x2b, 0xec, 0x0f, 0x37, 0x2b, 0x43, 0xbb, 0x5d, 0x19, 0xda, 0xef,
	0x95, 0xa1, 0x2d, 0xd7, 0x46, 0xe5, 0x76, 0x6d, 0x54, 0x7e, 0xac, 0x8d, 0xca, 0x97, 0x33, 0x9f,
	0xc3, 0x64, 0xee, 0x99, 0x54, 0x86, 0x16, 0x95, 0x21, 0x03, 0x6f, 0x0c, 0x9b, 0x87, 0xe4, 0x0b,
	0x72, 0xf7, 0xda, 0x7b, 0x7b, 0x2a, 0x7e, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xab, 0xbe, 0xb8,
	0x21, 0x99, 0x04, 0x00, 0x00,
}

func (m *Proof) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Proof) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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

func (m *Proof_Duplicateballotevidence) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Proof_Duplicateballotevidence) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ReplicatedBallotProof != nil {
		{
			volume, err := m.ReplicatedBallotProof.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintProof(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Proof_Rapidcustomerevidence) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Proof_Rapidcustomerevidence) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RapidCustomerAssaultProof != nil {
		{
			volume, err := m.RapidCustomerAssaultProof.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintProof(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ReplicatedBallotProof) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicatedBallotProof) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplicatedBallotProof) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVariableintProof(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x2a
	if m.RatifierEnergy != 0 {
		i = encodeVariableintProof(dAtA, i, uint64(m.RatifierEnergy))
		i--
		dAtA[i] = 0x20
	}
	if m.SumPollingEnergy != 0 {
		i = encodeVariableintProof(dAtA, i, uint64(m.SumPollingEnergy))
		i--
		dAtA[i] = 0x18
	}
	if m.BallotBYTE != nil {
		{
			volume, err := m.BallotBYTE.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintProof(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BallotA != nil {
		{
			volume, err := m.BallotA.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintProof(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RapidCustomerAssaultProof) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RapidCustomerAssaultProof) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *RapidCustomerAssaultProof) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n6, err6 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVariableintProof(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x2a
	if m.SumPollingEnergy != 0 {
		i = encodeVariableintProof(dAtA, i, uint64(m.SumPollingEnergy))
		i--
		dAtA[i] = 0x20
	}
	if len(m.FaultyRatifiers) > 0 {
		for idxNdEx := len(m.FaultyRatifiers) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.FaultyRatifiers[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = encodeVariableintProof(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.SharedLevel != 0 {
		i = encodeVariableintProof(dAtA, i, uint64(m.SharedLevel))
		i--
		dAtA[i] = 0x10
	}
	if m.ClashingLedger != nil {
		{
			volume, err := m.ClashingLedger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintProof(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProofCatalog) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofCatalog) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ProofCatalog) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		for idxNdEx := len(m.Proof) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Proof[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = encodeVariableintProof(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVariableintProof(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovProof(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *Proof) Volume() (n int) {
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

func (m *Proof_Duplicateballotevidence) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReplicatedBallotProof != nil {
		l = m.ReplicatedBallotProof.Volume()
		n += 1 + l + sovProof(uint64(l))
	}
	return n
}
func (m *Proof_Rapidcustomerevidence) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RapidCustomerAssaultProof != nil {
		l = m.RapidCustomerAssaultProof.Volume()
		n += 1 + l + sovProof(uint64(l))
	}
	return n
}
func (m *ReplicatedBallotProof) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotA != nil {
		l = m.BallotA.Volume()
		n += 1 + l + sovProof(uint64(l))
	}
	if m.BallotBYTE != nil {
		l = m.BallotBYTE.Volume()
		n += 1 + l + sovProof(uint64(l))
	}
	if m.SumPollingEnergy != 0 {
		n += 1 + sovProof(uint64(m.SumPollingEnergy))
	}
	if m.RatifierEnergy != 0 {
		n += 1 + sovProof(uint64(m.RatifierEnergy))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovProof(uint64(l))
	return n
}

func (m *RapidCustomerAssaultProof) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClashingLedger != nil {
		l = m.ClashingLedger.Volume()
		n += 1 + l + sovProof(uint64(l))
	}
	if m.SharedLevel != 0 {
		n += 1 + sovProof(uint64(m.SharedLevel))
	}
	if len(m.FaultyRatifiers) > 0 {
		for _, e := range m.FaultyRatifiers {
			l = e.Volume()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	if m.SumPollingEnergy != 0 {
		n += 1 + sovProof(uint64(m.SumPollingEnergy))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovProof(uint64(l))
	return n
}

func (m *ProofCatalog) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proof) > 0 {
		for _, e := range m.Proof {
			l = e.Volume()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadProof
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
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &ReplicatedBallotProof{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Proof_Duplicateballotevidence{v}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &RapidCustomerAssaultProof{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Proof_Rapidcustomerevidence{v}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitProof(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentProof
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
func (m *ReplicatedBallotProof) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadProof
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
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.BallotA == nil {
				m.BallotA = &Ballot{}
			}
			if err := m.BallotA.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.BallotBYTE == nil {
				m.BallotBYTE = &Ballot{}
			}
			if err := m.BallotBYTE.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
					return ErrIntegerOverloadProof
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
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.RatifierEnergy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadProof
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.RatifierEnergy |= int64(b&0x7F) << displace
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
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitProof(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentProof
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
func (m *RapidCustomerAssaultProof) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadProof
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
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClashingLedger == nil {
				m.ClashingLedger = &RapidLedger{}
			}
			if err := m.ClashingLedger.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SharedLevel = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadProof
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.SharedLevel |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FaultyRatifiers = append(m.FaultyRatifiers, &Ratifier{})
			if err := m.FaultyRatifiers[len(m.FaultyRatifiers)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SumPollingEnergy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadProof
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
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitProof(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentProof
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
func (m *ProofCatalog) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadProof
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
					return ErrIntegerOverloadProof
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
				return ErrCorruptExtentProof
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof, Proof{})
			if err := m.Proof[len(m.Proof)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitProof(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentProof
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
func omitProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadProof
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
					return 0, ErrIntegerOverloadProof
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
					return 0, ErrIntegerOverloadProof
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
				return 0, ErrCorruptExtentProof
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterProof
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentProof
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentProof        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadProof          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterProof = fmt.Errorf("REDACTED")
)
