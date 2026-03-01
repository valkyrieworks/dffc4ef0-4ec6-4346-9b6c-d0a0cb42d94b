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
	Sum isattestation_Total `protobuf_oneof:"sum"`
}

func (m *Proof) Restore()         { *m = Proof{} }
func (m *Proof) Text() string { return proto.CompactTextString(m) }
func (*Proof) SchemaArtifact()    {}
func (*Proof) Definition() ([]byte, []int) {
	return filedescriptor_6825fabc78e0a168, []int{0}
}
func (m *Proof) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Proof) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Proof.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Proof.Merge(m, src)
}
func (m *Proof) XXX_Extent() int {
	return m.Extent()
}
func (m *Proof) XXX_Dropunfamiliar() {
	xxx_signaldetails_Proof.DiscardUnknown(m)
}

var xxx_signaldetails_Proof proto.InternalMessageInfo

type isattestation_Total interface {
	isattestation_Total()
	SerializeToward([]byte) (int, error)
	Extent() int
}

type Proof_Replicatedballotevidence struct {
	ReplicatedBallotProof *ReplicatedBallotProof `protobuf:"octets,1,opt,name=duplicate_vote_evidence,json=duplicateVoteEvidence,proto3,oneof" json:"replicated_ballot_proof,omitempty"`
}
type Proof_Agilecustomerattackproof struct {
	AgileCustomerOnslaughtProof *AgileCustomerOnslaughtProof `protobuf:"octets,2,opt,name=light_client_attack_evidence,json=lightClientAttackEvidence,proto3,oneof" json:"agile_customer_onslaught_proof,omitempty"`
}

func (*Proof_Replicatedballotevidence) isattestation_Total()     {}
func (*Proof_Agilecustomerattackproof) isattestation_Total() {}

func (m *Proof) ObtainTotal() isattestation_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Proof) ObtainReplicatedBallotProof() *ReplicatedBallotProof {
	if x, ok := m.ObtainTotal().(*Proof_Replicatedballotevidence); ok {
		return x.ReplicatedBallotProof
	}
	return nil
}

func (m *Proof) ObtainAgileCustomerOnslaughtProof() *AgileCustomerOnslaughtProof {
	if x, ok := m.ObtainTotal().(*Proof_Agilecustomerattackproof); ok {
		return x.AgileCustomerOnslaughtProof
	}
	return nil
}

//
func (*Proof) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Proof_Replicatedballotevidence)(nil),
		(*Proof_Agilecustomerattackproof)(nil),
	}
}

//
type ReplicatedBallotProof struct {
	BallotAN            *Ballot     `protobuf:"octets,1,opt,name=vote_a,json=voteA,proto3" json:"ballot_an,omitempty"`
	BallotBYTE            *Ballot     `protobuf:"octets,2,opt,name=vote_b,json=voteB,proto3" json:"ballot_byte,omitempty"`
	SumBallotingPotency int64     `protobuf:"variableint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_balloting_potency,omitempty"`
	AssessorPotency   int64     `protobuf:"variableint,4,opt,name=validator_power,json=validatorPower,proto3" json:"assessor_potency,omitempty"`
	Timestamp        time.Time `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *ReplicatedBallotProof) Restore()         { *m = ReplicatedBallotProof{} }
func (m *ReplicatedBallotProof) Text() string { return proto.CompactTextString(m) }
func (*ReplicatedBallotProof) SchemaArtifact()    {}
func (*ReplicatedBallotProof) Definition() ([]byte, []int) {
	return filedescriptor_6825fabc78e0a168, []int{1}
}
func (m *ReplicatedBallotProof) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplicatedBallotProof) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replicatedballotevidence.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplicatedBallotProof) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replicatedballotevidence.Merge(m, src)
}
func (m *ReplicatedBallotProof) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplicatedBallotProof) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replicatedballotevidence.DiscardUnknown(m)
}

var xxx_signaldetails_Replicatedballotevidence proto.InternalMessageInfo

func (m *ReplicatedBallotProof) ObtainBallotAN() *Ballot {
	if m != nil {
		return m.BallotAN
	}
	return nil
}

func (m *ReplicatedBallotProof) ObtainBallotBYTE() *Ballot {
	if m != nil {
		return m.BallotBYTE
	}
	return nil
}

func (m *ReplicatedBallotProof) ObtainSumBallotingPotency() int64 {
	if m != nil {
		return m.SumBallotingPotency
	}
	return 0
}

func (m *ReplicatedBallotProof) ObtainAssessorPotency() int64 {
	if m != nil {
		return m.AssessorPotency
	}
	return 0
}

func (m *ReplicatedBallotProof) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

//
type AgileCustomerOnslaughtProof struct {
	DiscordantLedger    *AgileLedger  `protobuf:"octets,1,opt,name=conflicting_block,json=conflictingBlock,proto3" json:"discordant_ledger,omitempty"`
	SharedAltitude        int64        `protobuf:"variableint,2,opt,name=common_height,json=commonHeight,proto3" json:"shared_altitude,omitempty"`
	TreacherousAssessors []*Assessor `protobuf:"octets,3,rep,name=byzantine_validators,json=byzantineValidators,proto3" json:"treacherous_assessors,omitempty"`
	SumBallotingPotency    int64        `protobuf:"variableint,4,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_balloting_potency,omitempty"`
	Timestamp           time.Time    `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *AgileCustomerOnslaughtProof) Restore()         { *m = AgileCustomerOnslaughtProof{} }
func (m *AgileCustomerOnslaughtProof) Text() string { return proto.CompactTextString(m) }
func (*AgileCustomerOnslaughtProof) SchemaArtifact()    {}
func (*AgileCustomerOnslaughtProof) Definition() ([]byte, []int) {
	return filedescriptor_6825fabc78e0a168, []int{2}
}
func (m *AgileCustomerOnslaughtProof) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AgileCustomerOnslaughtProof) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Agilecustomerattackproof.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgileCustomerOnslaughtProof) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Agilecustomerattackproof.Merge(m, src)
}
func (m *AgileCustomerOnslaughtProof) XXX_Extent() int {
	return m.Extent()
}
func (m *AgileCustomerOnslaughtProof) XXX_Dropunfamiliar() {
	xxx_signaldetails_Agilecustomerattackproof.DiscardUnknown(m)
}

var xxx_signaldetails_Agilecustomerattackproof proto.InternalMessageInfo

func (m *AgileCustomerOnslaughtProof) ObtainDiscordantLedger() *AgileLedger {
	if m != nil {
		return m.DiscordantLedger
	}
	return nil
}

func (m *AgileCustomerOnslaughtProof) ObtainSharedAltitude() int64 {
	if m != nil {
		return m.SharedAltitude
	}
	return 0
}

func (m *AgileCustomerOnslaughtProof) ObtainTreacherousAssessors() []*Assessor {
	if m != nil {
		return m.TreacherousAssessors
	}
	return nil
}

func (m *AgileCustomerOnslaughtProof) ObtainSumBallotingPotency() int64 {
	if m != nil {
		return m.SumBallotingPotency
	}
	return 0
}

func (m *AgileCustomerOnslaughtProof) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type ProofCatalog struct {
	Proof []Proof `protobuf:"octets,1,rep,name=evidence,proto3" json:"proof"`
}

func (m *ProofCatalog) Restore()         { *m = ProofCatalog{} }
func (m *ProofCatalog) Text() string { return proto.CompactTextString(m) }
func (*ProofCatalog) SchemaArtifact()    {}
func (*ProofCatalog) Definition() ([]byte, []int) {
	return filedescriptor_6825fabc78e0a168, []int{3}
}
func (m *ProofCatalog) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ProofCatalog) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Attestationlist.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofCatalog) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Attestationlist.Merge(m, src)
}
func (m *ProofCatalog) XXX_Extent() int {
	return m.Extent()
}
func (m *ProofCatalog) XXX_Dropunfamiliar() {
	xxx_signaldetails_Attestationlist.DiscardUnknown(m)
}

var xxx_signaldetails_Attestationlist proto.InternalMessageInfo

func (m *ProofCatalog) ObtainProof() []Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func initialize() {
	proto.RegisterType((*Proof)(nil), "REDACTED")
	proto.RegisterType((*ReplicatedBallotProof)(nil), "REDACTED")
	proto.RegisterType((*AgileCustomerOnslaughtProof)(nil), "REDACTED")
	proto.RegisterType((*ProofCatalog)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_6825fabc78e0a168) }

var filedescriptor_6825fabc78e0a168 = []byte{
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

func (m *Proof) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Proof) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Proof) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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

func (m *Proof_Replicatedballotevidence) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Proof_Replicatedballotevidence) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ReplicatedBallotProof != nil {
		{
			extent, err := m.ReplicatedBallotProof.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintProof(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Proof_Agilecustomerattackproof) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Proof_Agilecustomerattackproof) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.AgileCustomerOnslaughtProof != nil {
		{
			extent, err := m.AgileCustomerOnslaughtProof.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintProof(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *ReplicatedBallotProof) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplicatedBallotProof) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplicatedBallotProof) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	n3, fault3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault3 != nil {
		return 0, fault3
	}
	i -= n3
	i = encodeVariableintProof(deltaLocatedAN, i, uint64(n3))
	i--
	deltaLocatedAN[i] = 0x2a
	if m.AssessorPotency != 0 {
		i = encodeVariableintProof(deltaLocatedAN, i, uint64(m.AssessorPotency))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.SumBallotingPotency != 0 {
		i = encodeVariableintProof(deltaLocatedAN, i, uint64(m.SumBallotingPotency))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.BallotBYTE != nil {
		{
			extent, err := m.BallotBYTE.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintProof(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.BallotAN != nil {
		{
			extent, err := m.BallotAN.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintProof(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *AgileCustomerOnslaughtProof) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AgileCustomerOnslaughtProof) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AgileCustomerOnslaughtProof) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	n6, fault6 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault6 != nil {
		return 0, fault6
	}
	i -= n6
	i = encodeVariableintProof(deltaLocatedAN, i, uint64(n6))
	i--
	deltaLocatedAN[i] = 0x2a
	if m.SumBallotingPotency != 0 {
		i = encodeVariableintProof(deltaLocatedAN, i, uint64(m.SumBallotingPotency))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if len(m.TreacherousAssessors) > 0 {
		for idxNdExc := len(m.TreacherousAssessors) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.TreacherousAssessors[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = encodeVariableintProof(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x1a
		}
	}
	if m.SharedAltitude != 0 {
		i = encodeVariableintProof(deltaLocatedAN, i, uint64(m.SharedAltitude))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.DiscordantLedger != nil {
		{
			extent, err := m.DiscordantLedger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintProof(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ProofCatalog) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ProofCatalog) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ProofCatalog) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		for idxNdExc := len(m.Proof) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Proof[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = encodeVariableintProof(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintProof(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovProof(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *Proof) Extent() (n int) {
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

func (m *Proof_Replicatedballotevidence) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReplicatedBallotProof != nil {
		l = m.ReplicatedBallotProof.Extent()
		n += 1 + l + sovProof(uint64(l))
	}
	return n
}
func (m *Proof_Agilecustomerattackproof) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AgileCustomerOnslaughtProof != nil {
		l = m.AgileCustomerOnslaughtProof.Extent()
		n += 1 + l + sovProof(uint64(l))
	}
	return n
}
func (m *ReplicatedBallotProof) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotAN != nil {
		l = m.BallotAN.Extent()
		n += 1 + l + sovProof(uint64(l))
	}
	if m.BallotBYTE != nil {
		l = m.BallotBYTE.Extent()
		n += 1 + l + sovProof(uint64(l))
	}
	if m.SumBallotingPotency != 0 {
		n += 1 + sovProof(uint64(m.SumBallotingPotency))
	}
	if m.AssessorPotency != 0 {
		n += 1 + sovProof(uint64(m.AssessorPotency))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovProof(uint64(l))
	return n
}

func (m *AgileCustomerOnslaughtProof) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DiscordantLedger != nil {
		l = m.DiscordantLedger.Extent()
		n += 1 + l + sovProof(uint64(l))
	}
	if m.SharedAltitude != 0 {
		n += 1 + sovProof(uint64(m.SharedAltitude))
	}
	if len(m.TreacherousAssessors) > 0 {
		for _, e := range m.TreacherousAssessors {
			l = e.Extent()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	if m.SumBallotingPotency != 0 {
		n += 1 + sovProof(uint64(m.SumBallotingPotency))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovProof(uint64(l))
	return n
}

func (m *ProofCatalog) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proof) > 0 {
		for _, e := range m.Proof {
			l = e.Extent()
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
func (m *Proof) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunProof
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
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &ReplicatedBallotProof{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Proof_Replicatedballotevidence{v}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &AgileCustomerOnslaughtProof{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Proof_Agilecustomerattackproof{v}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitProof(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeProof
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
func (m *ReplicatedBallotProof) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunProof
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
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.BallotAN == nil {
				m.BallotAN = &Ballot{}
			}
			if err := m.BallotAN.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.BallotBYTE == nil {
				m.BallotBYTE = &Ballot{}
			}
			if err := m.BallotBYTE.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
					return FaultIntegerOverrunProof
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
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.AssessorPotency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunProof
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.AssessorPotency |= int64(b&0x7F) << relocate
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
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitProof(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeProof
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
func (m *AgileCustomerOnslaughtProof) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunProof
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
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.DiscordantLedger == nil {
				m.DiscordantLedger = &AgileLedger{}
			}
			if err := m.DiscordantLedger.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SharedAltitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunProof
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.SharedAltitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.TreacherousAssessors = append(m.TreacherousAssessors, &Assessor{})
			if err := m.TreacherousAssessors[len(m.TreacherousAssessors)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SumBallotingPotency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunProof
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
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitProof(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeProof
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
func (m *ProofCatalog) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunProof
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
					return FaultIntegerOverrunProof
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
				return FaultUnfitMagnitudeProof
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeProof
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof, Proof{})
			if err := m.Proof[len(m.Proof)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitProof(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeProof
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
func omitProof(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunProof
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
					return 0, FaultIntegerOverrunProof
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
					return 0, FaultIntegerOverrunProof
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
				return 0, FaultUnfitMagnitudeProof
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionProof
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeProof
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeProof        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunProof          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionProof = fmt.Errorf("REDACTED")
)
