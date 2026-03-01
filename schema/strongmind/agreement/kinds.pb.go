//
//

package agreement

import (
	fmt "fmt"
	digits "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/utils/digits"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
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
//
type FreshIterationPhase struct {
	Altitude                int64  `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration                 int32  `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	Phase                  uint32 `protobuf:"variableint,3,opt,name=step,proto3" json:"phase,omitempty"`
	MomentsBecauseInitiateMoment int64  `protobuf:"variableint,4,opt,name=seconds_since_start_time,json=secondsSinceStartTime,proto3" json:"moments_because_initiate_moment,omitempty"`
	FinalEndorseIteration       int32  `protobuf:"variableint,5,opt,name=last_commit_round,json=lastCommitRound,proto3" json:"final_endorse_iteration,omitempty"`
}

func (m *FreshIterationPhase) Restore()         { *m = FreshIterationPhase{} }
func (m *FreshIterationPhase) Text() string { return proto.CompactTextString(m) }
func (*FreshIterationPhase) SchemaArtifact()    {}
func (*FreshIterationPhase) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{0}
}
func (m *FreshIterationPhase) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *FreshIterationPhase) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Newcyclephase.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FreshIterationPhase) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Newcyclephase.Merge(m, src)
}
func (m *FreshIterationPhase) XXX_Extent() int {
	return m.Extent()
}
func (m *FreshIterationPhase) XXX_Dropunfamiliar() {
	xxx_signaldetails_Newcyclephase.DiscardUnknown(m)
}

var xxx_signaldetails_Newcyclephase proto.InternalMessageInfo

func (m *FreshIterationPhase) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *FreshIterationPhase) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *FreshIterationPhase) ObtainPhase() uint32 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *FreshIterationPhase) ObtainMomentsBecauseInitiateMoment() int64 {
	if m != nil {
		return m.MomentsBecauseInitiateMoment
	}
	return 0
}

func (m *FreshIterationPhase) ObtainFinalEndorseIteration() int32 {
	if m != nil {
		return m.FinalEndorseIteration
	}
	return 0
}

//
//
//
type FreshSoundLedger struct {
	Altitude             int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration              int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	LedgerFragmentAssignHeading kinds.FragmentAssignHeading `protobuf:"octets,3,opt,name=block_part_set_header,json=blockPartSetHeader,proto3" json:"ledger_fragment_assign_headline"`
	LedgerFragments         *digits.DigitSeries      `protobuf:"octets,4,opt,name=block_parts,json=blockParts,proto3" json:"ledger_fragments,omitempty"`
	EqualsEndorse           bool                `protobuf:"variableint,5,opt,name=is_commit,json=isCommit,proto3" json:"equals_endorse,omitempty"`
}

func (m *FreshSoundLedger) Restore()         { *m = FreshSoundLedger{} }
func (m *FreshSoundLedger) Text() string { return proto.CompactTextString(m) }
func (*FreshSoundLedger) SchemaArtifact()    {}
func (*FreshSoundLedger) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{1}
}
func (m *FreshSoundLedger) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *FreshSoundLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Newvalidledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FreshSoundLedger) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Newvalidledger.Merge(m, src)
}
func (m *FreshSoundLedger) XXX_Extent() int {
	return m.Extent()
}
func (m *FreshSoundLedger) XXX_Dropunfamiliar() {
	xxx_signaldetails_Newvalidledger.DiscardUnknown(m)
}

var xxx_signaldetails_Newvalidledger proto.InternalMessageInfo

func (m *FreshSoundLedger) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *FreshSoundLedger) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *FreshSoundLedger) ObtainLedgerFragmentAssignHeadline() kinds.FragmentAssignHeading {
	if m != nil {
		return m.LedgerFragmentAssignHeading
	}
	return kinds.FragmentAssignHeading{}
}

func (m *FreshSoundLedger) ObtainLedgerFragments() *digits.DigitSeries {
	if m != nil {
		return m.LedgerFragments
	}
	return nil
}

func (m *FreshSoundLedger) ObtainEqualsEndorse() bool {
	if m != nil {
		return m.EqualsEndorse
	}
	return false
}

//
type Nomination struct {
	Nomination kinds.Nomination `protobuf:"octets,1,opt,name=proposal,proto3" json:"nomination"`
}

func (m *Nomination) Restore()         { *m = Nomination{} }
func (m *Nomination) Text() string { return proto.CompactTextString(m) }
func (*Nomination) SchemaArtifact()    {}
func (*Nomination) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{2}
}
func (m *Nomination) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Nomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Nomination.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nomination) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Nomination.Merge(m, src)
}
func (m *Nomination) XXX_Extent() int {
	return m.Extent()
}
func (m *Nomination) XXX_Dropunfamiliar() {
	xxx_signaldetails_Nomination.DiscardUnknown(m)
}

var xxx_signaldetails_Nomination proto.InternalMessageInfo

func (m *Nomination) ObtainNomination() kinds.Nomination {
	if m != nil {
		return m.Nomination
	}
	return kinds.Nomination{}
}

//
type NominationPolicy struct {
	Altitude           int64         `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	NominationPolicyIteration int32         `protobuf:"variableint,2,opt,name=proposal_pol_round,json=proposalPolRound,proto3" json:"nomination_policy_iteration,omitempty"`
	NominationPolicy      digits.DigitSeries `protobuf:"octets,3,opt,name=proposal_pol,json=proposalPol,proto3" json:"nomination_policy"`
}

func (m *NominationPolicy) Restore()         { *m = NominationPolicy{} }
func (m *NominationPolicy) Text() string { return proto.CompactTextString(m) }
func (*NominationPolicy) SchemaArtifact()    {}
func (*NominationPolicy) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{3}
}
func (m *NominationPolicy) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *NominationPolicy) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Proposalpolicy.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NominationPolicy) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Proposalpolicy.Merge(m, src)
}
func (m *NominationPolicy) XXX_Extent() int {
	return m.Extent()
}
func (m *NominationPolicy) XXX_Dropunfamiliar() {
	xxx_signaldetails_Proposalpolicy.DiscardUnknown(m)
}

var xxx_signaldetails_Proposalpolicy proto.InternalMessageInfo

func (m *NominationPolicy) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *NominationPolicy) ObtainNominationPolicyIteration() int32 {
	if m != nil {
		return m.NominationPolicyIteration
	}
	return 0
}

func (m *NominationPolicy) ObtainNominationPolicy() digits.DigitSeries {
	if m != nil {
		return m.NominationPolicy
	}
	return digits.DigitSeries{}
}

//
type LedgerFragment struct {
	Altitude int64      `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration  int32      `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	Fragment   kinds.Fragment `protobuf:"octets,3,opt,name=part,proto3" json:"fragment"`
}

func (m *LedgerFragment) Restore()         { *m = LedgerFragment{} }
func (m *LedgerFragment) Text() string { return proto.CompactTextString(m) }
func (*LedgerFragment) SchemaArtifact()    {}
func (*LedgerFragment) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{4}
}
func (m *LedgerFragment) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *LedgerFragment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ledgerfragment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerFragment) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ledgerfragment.Merge(m, src)
}
func (m *LedgerFragment) XXX_Extent() int {
	return m.Extent()
}
func (m *LedgerFragment) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ledgerfragment.DiscardUnknown(m)
}

var xxx_signaldetails_Ledgerfragment proto.InternalMessageInfo

func (m *LedgerFragment) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *LedgerFragment) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *LedgerFragment) ObtainFragment() kinds.Fragment {
	if m != nil {
		return m.Fragment
	}
	return kinds.Fragment{}
}

//
type Ballot struct {
	Ballot *kinds.Ballot `protobuf:"octets,1,opt,name=vote,proto3" json:"ballot,omitempty"`
}

func (m *Ballot) Restore()         { *m = Ballot{} }
func (m *Ballot) Text() string { return proto.CompactTextString(m) }
func (*Ballot) SchemaArtifact()    {}
func (*Ballot) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{5}
}
func (m *Ballot) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Ballot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ballot) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ballot.Merge(m, src)
}
func (m *Ballot) XXX_Extent() int {
	return m.Extent()
}
func (m *Ballot) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ballot.DiscardUnknown(m)
}

var xxx_signaldetails_Ballot proto.InternalMessageInfo

func (m *Ballot) FetchBallot() *kinds.Ballot {
	if m != nil {
		return m.Ballot
	}
	return nil
}

//
type OwnsBallot struct {
	Altitude int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration  int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	Kind   kinds.AttestedSignalKind `protobuf:"variableint,3,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Ordinal  int32               `protobuf:"variableint,4,opt,name=index,proto3" json:"ordinal,omitempty"`
}

func (m *OwnsBallot) Restore()         { *m = OwnsBallot{} }
func (m *OwnsBallot) Text() string { return proto.CompactTextString(m) }
func (*OwnsBallot) SchemaArtifact()    {}
func (*OwnsBallot) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{6}
}
func (m *OwnsBallot) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *OwnsBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Hasballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnsBallot) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Hasballot.Merge(m, src)
}
func (m *OwnsBallot) XXX_Extent() int {
	return m.Extent()
}
func (m *OwnsBallot) XXX_Dropunfamiliar() {
	xxx_signaldetails_Hasballot.DiscardUnknown(m)
}

var xxx_signaldetails_Hasballot proto.InternalMessageInfo

func (m *OwnsBallot) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *OwnsBallot) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *OwnsBallot) ObtainKind() kinds.AttestedSignalKind {
	if m != nil {
		return m.Kind
	}
	return kinds.UnfamiliarKind
}

func (m *OwnsBallot) ObtainOrdinal() int32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

//
type BallotAssignMajor23 struct {
	Altitude  int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration   int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	Kind    kinds.AttestedSignalKind `protobuf:"variableint,3,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	LedgerUUID kinds.LedgerUUID       `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid"`
}

func (m *BallotAssignMajor23) Restore()         { *m = BallotAssignMajor23{} }
func (m *BallotAssignMajor23) Text() string { return proto.CompactTextString(m) }
func (*BallotAssignMajor23) SchemaArtifact()    {}
func (*BallotAssignMajor23) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{7}
}
func (m *BallotAssignMajor23) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *BallotAssignMajor23) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ballotsetmaj23.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BallotAssignMajor23) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ballotsetmaj23.Merge(m, src)
}
func (m *BallotAssignMajor23) XXX_Extent() int {
	return m.Extent()
}
func (m *BallotAssignMajor23) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ballotsetmaj23.DiscardUnknown(m)
}

var xxx_signaldetails_Ballotsetmaj23 proto.InternalMessageInfo

func (m *BallotAssignMajor23) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *BallotAssignMajor23) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *BallotAssignMajor23) ObtainKind() kinds.AttestedSignalKind {
	if m != nil {
		return m.Kind
	}
	return kinds.UnfamiliarKind
}

func (m *BallotAssignMajor23) ObtainLedgerUUID() kinds.LedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return kinds.LedgerUUID{}
}

//
type BallotAssignDigits struct {
	Altitude  int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration   int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	Kind    kinds.AttestedSignalKind `protobuf:"variableint,3,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	LedgerUUID kinds.LedgerUUID       `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid"`
	Ballots   digits.DigitSeries       `protobuf:"octets,5,opt,name=votes,proto3" json:"ballots"`
}

func (m *BallotAssignDigits) Restore()         { *m = BallotAssignDigits{} }
func (m *BallotAssignDigits) Text() string { return proto.CompactTextString(m) }
func (*BallotAssignDigits) SchemaArtifact()    {}
func (*BallotAssignDigits) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{8}
}
func (m *BallotAssignDigits) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *BallotAssignDigits) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ballotsetdigits.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BallotAssignDigits) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ballotsetdigits.Merge(m, src)
}
func (m *BallotAssignDigits) XXX_Extent() int {
	return m.Extent()
}
func (m *BallotAssignDigits) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ballotsetdigits.DiscardUnknown(m)
}

var xxx_signaldetails_Ballotsetdigits proto.InternalMessageInfo

func (m *BallotAssignDigits) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *BallotAssignDigits) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *BallotAssignDigits) ObtainKind() kinds.AttestedSignalKind {
	if m != nil {
		return m.Kind
	}
	return kinds.UnfamiliarKind
}

func (m *BallotAssignDigits) ObtainLedgerUUID() kinds.LedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return kinds.LedgerUUID{}
}

func (m *BallotAssignDigits) ObtainBallots() digits.DigitSeries {
	if m != nil {
		return m.Ballots
	}
	return digits.DigitSeries{}
}

type Signal struct {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	Sum isnote_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) Text() string { return proto.CompactTextString(m) }
func (*Signal) SchemaArtifact()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedescriptor_81a22d2efc008981, []int{9}
}
func (m *Signal) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Signal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Artifact.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signal) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Artifact.Merge(m, src)
}
func (m *Signal) XXX_Extent() int {
	return m.Extent()
}
func (m *Signal) XXX_Dropunfamiliar() {
	xxx_signaldetails_Artifact.DiscardUnknown(m)
}

var xxx_signaldetails_Artifact proto.InternalMessageInfo

type isnote_Total interface {
	isnote_Total()
	SerializeToward([]byte) (int, error)
	Extent() int
}

type Signal_Newcyclephase struct {
	FreshIterationPhase *FreshIterationPhase `protobuf:"octets,1,opt,name=new_round_step,json=newRoundStep,proto3,oneof" json:"fresh_iteration_phase,omitempty"`
}
type Signal_Newvalidledger struct {
	FreshSoundLedger *FreshSoundLedger `protobuf:"octets,2,opt,name=new_valid_block,json=newValidBlock,proto3,oneof" json:"fresh_sound_ledger,omitempty"`
}
type Signal_Nomination struct {
	Nomination *Nomination `protobuf:"octets,3,opt,name=proposal,proto3,oneof" json:"nomination,omitempty"`
}
type Signal_Proposalpolicy struct {
	NominationPolicy *NominationPolicy `protobuf:"octets,4,opt,name=proposal_pol,json=proposalPol,proto3,oneof" json:"nomination_policy,omitempty"`
}
type Signal_Ledgerfragment struct {
	LedgerFragment *LedgerFragment `protobuf:"octets,5,opt,name=block_part,json=blockPart,proto3,oneof" json:"ledger_fragment,omitempty"`
}
type Signal_Ballot struct {
	Ballot *Ballot `protobuf:"octets,6,opt,name=vote,proto3,oneof" json:"ballot,omitempty"`
}
type Signal_Hasballot struct {
	OwnsBallot *OwnsBallot `protobuf:"octets,7,opt,name=has_vote,json=hasVote,proto3,oneof" json:"owns_ballot,omitempty"`
}
type Signal_Ballotsetmaj23 struct {
	BallotAssignMajor23 *BallotAssignMajor23 `protobuf:"octets,8,opt,name=vote_set_maj23,json=voteSetMaj23,proto3,oneof" json:"ballot_assign_major23,omitempty"`
}
type Signal_Ballotsetdigits struct {
	BallotAssignDigits *BallotAssignDigits `protobuf:"octets,9,opt,name=vote_set_bits,json=voteSetBits,proto3,oneof" json:"ballot_assign_digits,omitempty"`
}

func (*Signal_Newcyclephase) isnote_Total()  {}
func (*Signal_Newvalidledger) isnote_Total() {}
func (*Signal_Nomination) isnote_Total()      {}
func (*Signal_Proposalpolicy) isnote_Total()   {}
func (*Signal_Ledgerfragment) isnote_Total()     {}
func (*Signal_Ballot) isnote_Total()          {}
func (*Signal_Hasballot) isnote_Total()       {}
func (*Signal_Ballotsetmaj23) isnote_Total()  {}
func (*Signal_Ballotsetdigits) isnote_Total()   {}

func (m *Signal) ObtainTotal() isnote_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) ObtainFreshIterationPhase() *FreshIterationPhase {
	if x, ok := m.ObtainTotal().(*Signal_Newcyclephase); ok {
		return x.FreshIterationPhase
	}
	return nil
}

func (m *Signal) ObtainFreshSoundLedger() *FreshSoundLedger {
	if x, ok := m.ObtainTotal().(*Signal_Newvalidledger); ok {
		return x.FreshSoundLedger
	}
	return nil
}

func (m *Signal) ObtainNomination() *Nomination {
	if x, ok := m.ObtainTotal().(*Signal_Nomination); ok {
		return x.Nomination
	}
	return nil
}

func (m *Signal) ObtainNominationPolicy() *NominationPolicy {
	if x, ok := m.ObtainTotal().(*Signal_Proposalpolicy); ok {
		return x.NominationPolicy
	}
	return nil
}

func (m *Signal) ObtainLedgerFragment() *LedgerFragment {
	if x, ok := m.ObtainTotal().(*Signal_Ledgerfragment); ok {
		return x.LedgerFragment
	}
	return nil
}

func (m *Signal) FetchBallot() *Ballot {
	if x, ok := m.ObtainTotal().(*Signal_Ballot); ok {
		return x.Ballot
	}
	return nil
}

func (m *Signal) ObtainOwnsBallot() *OwnsBallot {
	if x, ok := m.ObtainTotal().(*Signal_Hasballot); ok {
		return x.OwnsBallot
	}
	return nil
}

func (m *Signal) ObtainBallotAssignMajor23() *BallotAssignMajor23 {
	if x, ok := m.ObtainTotal().(*Signal_Ballotsetmaj23); ok {
		return x.BallotAssignMajor23
	}
	return nil
}

func (m *Signal) ObtainBallotAssignDigits() *BallotAssignDigits {
	if x, ok := m.ObtainTotal().(*Signal_Ballotsetdigits); ok {
		return x.BallotAssignDigits
	}
	return nil
}

//
func (*Signal) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Signal_Newcyclephase)(nil),
		(*Signal_Newvalidledger)(nil),
		(*Signal_Nomination)(nil),
		(*Signal_Proposalpolicy)(nil),
		(*Signal_Ledgerfragment)(nil),
		(*Signal_Ballot)(nil),
		(*Signal_Hasballot)(nil),
		(*Signal_Ballotsetmaj23)(nil),
		(*Signal_Ballotsetdigits)(nil),
	}
}

func initialize() {
	proto.RegisterType((*FreshIterationPhase)(nil), "REDACTED")
	proto.RegisterType((*FreshSoundLedger)(nil), "REDACTED")
	proto.RegisterType((*Nomination)(nil), "REDACTED")
	proto.RegisterType((*NominationPolicy)(nil), "REDACTED")
	proto.RegisterType((*LedgerFragment)(nil), "REDACTED")
	proto.RegisterType((*Ballot)(nil), "REDACTED")
	proto.RegisterType((*OwnsBallot)(nil), "REDACTED")
	proto.RegisterType((*BallotAssignMajor23)(nil), "REDACTED")
	proto.RegisterType((*BallotAssignDigits)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_81a22d2efc008981) }

var filedescriptor_81a22d2efc008981 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xb6, 0x19, 0x67, 0x92, 0x94, 0x93, 0x19, 0x68, 0xcd, 0xae, 0x42, 0x80, 0x24, 0x98, 0xcb,
	0x08, 0x21, 0x07, 0x65, 0x0e, 0x2b, 0xad, 0x90, 0x00, 0xf3, 0xb3, 0xde, 0xd5, 0x66, 0x37, 0x38,
	0xab, 0x15, 0xe2, 0x62, 0x39, 0x71, 0x93, 0x34, 0x1b, 0xbb, 0x2d, 0x77, 0x27, 0xc3, 0x5c, 0x79,
	0x02, 0x1e, 0x80, 0xd7, 0x40, 0xe2, 0x11, 0xe6, 0x38, 0x47, 0x4e, 0x23, 0x94, 0x79, 0x04, 0x04,
	0x67, 0xd4, 0xed, 0x4e, 0xec, 0x30, 0x9e, 0x81, 0xb9, 0x20, 0x71, 0xeb, 0x4e, 0x55, 0x7d, 0x5d,
	0xf5, 0x55, 0xd5, 0x17, 0x43, 0x8f, 0xe3, 0x38, 0xc4, 0x69, 0x44, 0x62, 0xde, 0x9f, 0xd2, 0x98,
	0xe1, 0x98, 0x2d, 0x59, 0x9f, 0x9f, 0x25, 0x98, 0xd9, 0x49, 0x4a, 0x39, 0x45, 0x47, 0xb9, 0x87,
	0xbd, 0xf5, 0x68, 0x1f, 0xcd, 0xe8, 0x8c, 0x4a, 0x87, 0xbe, 0x38, 0x65, 0xbe, 0xed, 0xb7, 0x0b,
	0x68, 0x12, 0xa3, 0x88, 0xd4, 0x2e, 0xbe, 0xb5, 0x20, 0x13, 0xd6, 0x9f, 0x10, 0xbe, 0xe3, 0x61,
	0xfd, 0xac, 0x43, 0xe3, 0x19, 0x3e, 0xf5, 0xe8, 0x32, 0x0e, 0xc7, 0x1c, 0x27, 0xe8, 0x3e, 0xec,
	0xcf, 0x31, 0x99, 0xcd, 0x79, 0x4b, 0xef, 0xe9, 0xc7, 0x7b, 0x9e, 0xba, 0xa1, 0x23, 0xa8, 0xa4,
	0xc2, 0xa9, 0xf5, 0x5a, 0x4f, 0x3f, 0xae, 0x78, 0xd9, 0x05, 0x21, 0x30, 0x18, 0xc7, 0x49, 0x6b,
	0xaf, 0xa7, 0x1f, 0x37, 0x3d, 0x79, 0x46, 0x0f, 0xa0, 0xc5, 0xf0, 0x94, 0xc6, 0x21, 0xf3, 0x19,
	0x89, 0xa7, 0xd8, 0x67, 0x3c, 0x48, 0xb9, 0xcf, 0x49, 0x84, 0x5b, 0x86, 0xc4, 0xbc, 0xa7, 0xec,
	0x63, 0x61, 0x1e, 0x0b, 0xeb, 0x0b, 0x12, 0x61, 0xf4, 0x3e, 0xbc, 0xb1, 0x08, 0x18, 0xf7, 0xa7,
	0x34, 0x8a, 0x08, 0xf7, 0xb3, 0xe7, 0x2a, 0xf2, 0xb9, 0x43, 0x61, 0xf8, 0x4c, 0xfe, 0x2e, 0x53,
	0xb5, 0xfe, 0xd0, 0xa1, 0xf9, 0x0c, 0x9f, 0xbe, 0x0c, 0x16, 0x24, 0x74, 0x16, 0x74, 0xfa, 0xea,
	0x8e, 0x89, 0x7f, 0x0d, 0xf7, 0x26, 0x22, 0xcc, 0x4f, 0x44, 0x6e, 0x0c, 0x73, 0x7f, 0x8e, 0x83,
	0x10, 0xa7, 0xb2, 0x12, 0x73, 0xd0, 0xb5, 0x0b, 0x3d, 0xc8, 0xf8, 0x1a, 0x05, 0x29, 0x1f, 0x63,
	0xee, 0x4a, 0x37, 0xc7, 0x38, 0xbf, 0xec, 0x6a, 0x1e, 0x92, 0x18, 0x3b, 0x16, 0xf4, 0x31, 0x98,
	0x39, 0x32, 0x93, 0x15, 0x9b, 0x83, 0x4e, 0x11, 0x4f, 0x74, 0xc2, 0x16, 0x9d, 0xb0, 0x1d, 0xc2,
	0x3f, 0x4d, 0xd3, 0xe0, 0xcc, 0x83, 0x2d, 0x10, 0x43, 0x6f, 0x41, 0x9d, 0x30, 0x45, 0x82, 0x2c,
	0xbf, 0xe6, 0xd5, 0x08, 0xcb, 0x8a, 0xb7, 0x5c, 0xa8, 0x8d, 0x52, 0x9a, 0x50, 0x16, 0x2c, 0xd0,
	0x47, 0x50, 0x4b, 0xd4, 0x59, 0xd6, 0x6c, 0x0e, 0xda, 0x25, 0x69, 0x2b, 0x0f, 0x95, 0xf1, 0x36,
	0xc2, 0xfa, 0x49, 0x07, 0x73, 0x63, 0x1c, 0x3d, 0x7f, 0x7a, 0x23, 0x7f, 0x1f, 0x00, 0xda, 0xc4,
	0xf8, 0x09, 0x5d, 0xf8, 0x45, 0x32, 0x5f, 0xdf, 0x58, 0x46, 0x74, 0x21, 0xfb, 0x82, 0x1e, 0x41,
	0xa3, 0xe8, 0xad, 0xe8, 0xfc, 0x87, 0xf2, 0x55, 0x6e, 0x66, 0x01, 0xcd, 0x7a, 0x05, 0x75, 0x67,
	0xc3, 0xc9, 0x1d, 0x7b, 0xfb, 0x21, 0x18, 0x82, 0x7b, 0xf5, 0xf6, 0xfd, 0xf2, 0x56, 0xaa, 0x37,
	0xa5, 0xa7, 0x35, 0x00, 0xe3, 0x25, 0xe5, 0x62, 0x02, 0x8d, 0x15, 0xe5, 0x58, 0xb1, 0x59, 0x12,
	0x29, 0xbc, 0x3c, 0xe9, 0x63, 0xfd, 0xa0, 0x43, 0xd5, 0x0d, 0x98, 0x8c, 0xbb, 0x5b, 0x7e, 0x27,
	0x60, 0x08, 0x34, 0x99, 0xdf, 0x41, 0xd9, 0xa8, 0x8d, 0xc9, 0x2c, 0xc6, 0xe1, 0x90, 0xcd, 0x5e,
	0x9c, 0x25, 0xd8, 0x93, 0xce, 0x02, 0x8a, 0xc4, 0x21, 0xfe, 0x5e, 0x0e, 0x54, 0xc5, 0xcb, 0x2e,
	0xd6, 0x2f, 0x3a, 0x34, 0x44, 0x06, 0x63, 0xcc, 0x87, 0xc1, 0x77, 0x83, 0x93, 0xff, 0x22, 0x93,
	0x2f, 0xa0, 0x96, 0x0d, 0x38, 0x09, 0xd5, 0x74, 0xbf, 0x79, 0x3d, 0x50, 0xf6, 0xee, 0xf1, 0xe7,
	0xce, 0xa1, 0x60, 0x79, 0x7d, 0xd9, 0xad, 0xaa, 0x1f, 0xbc, 0xaa, 0x8c, 0x7d, 0x1c, 0x5a, 0xbf,
	0xeb, 0x60, 0xaa, 0xd4, 0x1d, 0xc2, 0xd9, 0xff, 0x27, 0x73, 0xf4, 0x10, 0x2a, 0x62, 0x02, 0x98,
	0x5c, 0xce, 0x7f, 0x3b, 0xdc, 0x59, 0x88, 0xf5, 0xa7, 0x01, 0xd5, 0x21, 0x66, 0x2c, 0x98, 0x61,
	0xf4, 0x04, 0x0e, 0x62, 0x7c, 0x9a, 0x2d, 0x94, 0x2f, 0x65, 0x34, 0x9b, 0x3b, 0xcb, 0x2e, 0xfb,
	0x03, 0xb0, 0x8b, 0x32, 0xed, 0x6a, 0x5e, 0x23, 0x2e, 0xca, 0xf6, 0x10, 0x0e, 0x05, 0xd6, 0x4a,
	0xe8, 0xa1, 0x2f, 0x13, 0x95, 0x7c, 0x99, 0x83, 0xf7, 0x6e, 0x04, 0xcb, 0xb5, 0xd3, 0xd5, 0xbc,
	0x66, 0xbc, 0x23, 0xa6, 0x45, 0x69, 0x29, 0x59, 0xe1, 0x1c, 0x67, 0xa3, 0x20, 0x6e, 0x41, 0x5a,
	0xd0, 0x97, 0x7f, 0x13, 0x81, 0x8c, 0xeb, 0x77, 0x6f, 0x47, 0x18, 0x3d, 0x7f, 0xea, 0xee, 0x6a,
	0x00, 0xfa, 0x04, 0x20, 0x97, 0x52, 0xc5, 0x76, 0xb7, 0x1c, 0x65, 0xab, 0x15, 0xae, 0xe6, 0xd5,
	0xb7, 0x62, 0x2a, 0xa4, 0x40, 0x2e, 0xf4, 0xfe, 0x75, 0x79, 0xcc, 0x63, 0xc5, 0x14, 0xba, 0x5a,
	0xb6, 0xd6, 0xe8, 0x21, 0xd4, 0xe6, 0x01, 0xf3, 0x65, 0x54, 0x55, 0x46, 0xbd, 0x53, 0x1e, 0xa5,
	0x76, 0xdf, 0xd5, 0xbc, 0xea, 0x5c, 0xc9, 0xc0, 0x13, 0x38, 0x10, 0x71, 0xf2, 0xef, 0x24, 0x12,
	0xeb, 0xd8, 0xaa, 0xdd, 0xd6, 0xd0, 0xe2, 0xe2, 0x8a, 0x86, 0xae, 0x8a, 0x8b, 0xfc, 0x08, 0x9a,
	0x5b, 0x2c, 0x31, 0x4f, 0xad, 0xfa, 0x6d, 0x24, 0x16, 0x16, 0x49, 0x90, 0xb8, 0xca, 0xaf, 0x4e,
	0x05, 0xf6, 0xd8, 0x32, 0x72, 0xbe, 0x3a, 0x5f, 0x77, 0xf4, 0x8b, 0x75, 0x47, 0xff, 0x6d, 0xdd,
	0xd1, 0x7f, 0xbc, 0xea, 0x68, 0x17, 0x57, 0x1d, 0xed, 0xd7, 0xab, 0x8e, 0xf6, 0xcd, 0x83, 0x19,
	0xe1, 0xf3, 0xe5, 0xc4, 0x9e, 0xd2, 0xa8, 0x3f, 0xa5, 0x11, 0xe6, 0x93, 0x6f, 0x79, 0x7e, 0xc8,
	0xbe, 0x38, 0xca, 0xbe, 0x59, 0x26, 0xfb, 0xd2, 0x76, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xfe, 0x06, 0x66, 0x65, 0xd2, 0x08, 0x00, 0x00,
}

func (m *FreshIterationPhase) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *FreshIterationPhase) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *FreshIterationPhase) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.FinalEndorseIteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FinalEndorseIteration))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if m.MomentsBecauseInitiateMoment != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.MomentsBecauseInitiateMoment))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.Phase != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Phase))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *FreshSoundLedger) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *FreshSoundLedger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *FreshSoundLedger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.EqualsEndorse {
		i--
		if m.EqualsEndorse {
			deltaLocatedAN[i] = 1
		} else {
			deltaLocatedAN[i] = 0
		}
		i--
		deltaLocatedAN[i] = 0x28
	}
	if m.LedgerFragments != nil {
		{
			extent, err := m.LedgerFragments.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x22
	}
	{
		extent, err := m.LedgerFragmentAssignHeading.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x1a
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Nomination) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Nomination) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Nomination) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.Nomination.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func (m *NominationPolicy) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *NominationPolicy) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *NominationPolicy) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.NominationPolicy.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x1a
	if m.NominationPolicyIteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.NominationPolicyIteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *LedgerFragment) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *LedgerFragment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *LedgerFragment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.Fragment.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x1a
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Ballot) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Ballot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Ballot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Ballot != nil {
		{
			extent, err := m.Ballot.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *OwnsBallot) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *OwnsBallot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *OwnsBallot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Ordinal != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ordinal))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.Kind != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *BallotAssignMajor23) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *BallotAssignMajor23) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *BallotAssignMajor23) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x22
	if m.Kind != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *BallotAssignDigits) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *BallotAssignDigits) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *BallotAssignDigits) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.Ballots.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x2a
	{
		extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x22
	if m.Kind != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Signal) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Signal) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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

func (m *Signal_Newcyclephase) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Newcyclephase) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.FreshIterationPhase != nil {
		{
			extent, err := m.FreshIterationPhase.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Newvalidledger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Newvalidledger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.FreshSoundLedger != nil {
		{
			extent, err := m.FreshSoundLedger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Nomination) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Nomination) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Nomination != nil {
		{
			extent, err := m.Nomination.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Proposalpolicy) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Proposalpolicy) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.NominationPolicy != nil {
		{
			extent, err := m.NominationPolicy.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x22
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Ledgerfragment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Ledgerfragment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.LedgerFragment != nil {
		{
			extent, err := m.LedgerFragment.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x2a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Ballot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Ballot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Ballot != nil {
		{
			extent, err := m.Ballot.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x32
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Hasballot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Hasballot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.OwnsBallot != nil {
		{
			extent, err := m.OwnsBallot.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x3a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Ballotsetmaj23) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Ballotsetmaj23) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.BallotAssignMajor23 != nil {
		{
			extent, err := m.BallotAssignMajor23.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x42
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Signal_Ballotsetdigits) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Ballotsetdigits) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.BallotAssignDigits != nil {
		{
			extent, err := m.BallotAssignDigits.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x4a
	}
	return len(deltaLocatedAN) - i, nil
}
func serializeVariableintKinds(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovKinds(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *FreshIterationPhase) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	if m.Phase != 0 {
		n += 1 + sovKinds(uint64(m.Phase))
	}
	if m.MomentsBecauseInitiateMoment != 0 {
		n += 1 + sovKinds(uint64(m.MomentsBecauseInitiateMoment))
	}
	if m.FinalEndorseIteration != 0 {
		n += 1 + sovKinds(uint64(m.FinalEndorseIteration))
	}
	return n
}

func (m *FreshSoundLedger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	l = m.LedgerFragmentAssignHeading.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.LedgerFragments != nil {
		l = m.LedgerFragments.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.EqualsEndorse {
		n += 2
	}
	return n
}

func (m *Nomination) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Nomination.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *NominationPolicy) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.NominationPolicyIteration != 0 {
		n += 1 + sovKinds(uint64(m.NominationPolicyIteration))
	}
	l = m.NominationPolicy.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *LedgerFragment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	l = m.Fragment.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *Ballot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ballot != nil {
		l = m.Ballot.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *OwnsBallot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	return n
}

func (m *BallotAssignMajor23) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	l = m.LedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *BallotAssignDigits) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	l = m.LedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = m.Ballots.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *Signal) Extent() (n int) {
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

func (m *Signal_Newcyclephase) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FreshIterationPhase != nil {
		l = m.FreshIterationPhase.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Newvalidledger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FreshSoundLedger != nil {
		l = m.FreshSoundLedger.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Nomination) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nomination != nil {
		l = m.Nomination.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Proposalpolicy) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NominationPolicy != nil {
		l = m.NominationPolicy.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ledgerfragment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerFragment != nil {
		l = m.LedgerFragment.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ballot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ballot != nil {
		l = m.Ballot.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Hasballot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnsBallot != nil {
		l = m.OwnsBallot.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ballotsetmaj23) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotAssignMajor23 != nil {
		l = m.BallotAssignMajor23.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ballotsetdigits) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotAssignDigits != nil {
		l = m.BallotAssignDigits.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FreshIterationPhase) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Phase = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Phase |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.MomentsBecauseInitiateMoment = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MomentsBecauseInitiateMoment |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalEndorseIteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FinalEndorseIteration |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *FreshSoundLedger) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LedgerFragmentAssignHeading.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.LedgerFragments == nil {
				m.LedgerFragments = &digits.DigitSeries{}
			}
			if err := m.LedgerFragments.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var v int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
			m.EqualsEndorse = bool(v != 0)
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *Nomination) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Nomination.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *NominationPolicy) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.NominationPolicyIteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.NominationPolicyIteration |= int32(b&0x7F) << relocate
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NominationPolicy.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *LedgerFragment) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fragment.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *Ballot) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ballot == nil {
				m.Ballot = &kinds.Ballot{}
			}
			if err := m.Ballot.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *OwnsBallot) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Kind = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Kind |= kinds.AttestedSignalKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Ordinal |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *BallotAssignMajor23) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Kind = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Kind |= kinds.AttestedSignalKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LedgerUUID.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *BallotAssignDigits) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Kind = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Kind |= kinds.AttestedSignalKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ballots.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *Signal) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunKinds
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &FreshIterationPhase{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Newcyclephase{v}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &FreshSoundLedger{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Newvalidledger{v}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &Nomination{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Nomination{v}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &NominationPolicy{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Proposalpolicy{v}
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &LedgerFragment{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ledgerfragment{v}
			idxNdExc = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &Ballot{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ballot{v}
			idxNdExc = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &OwnsBallot{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Hasballot{v}
			idxNdExc = submitOrdinal
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &BallotAssignMajor23{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ballotsetmaj23{v}
			idxNdExc = submitOrdinal
		case 9:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &BallotAssignDigits{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ballotsetdigits{v}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitKinds(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeKinds
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
func omitKinds(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunKinds
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
					return 0, FaultIntegerOverrunKinds
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
					return 0, FaultIntegerOverrunKinds
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
				return 0, FaultUnfitMagnitudeKinds
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingClusterKinds
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeKinds
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeKinds        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunKinds          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingClusterKinds = fmt.Errorf("REDACTED")
)
