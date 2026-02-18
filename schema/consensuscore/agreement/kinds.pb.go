//
//

package agreement

import (
	fmt "fmt"
	bits "github.com/valkyrieworks/schema/consensuscore/utils/bits"
	kinds "github.com/valkyrieworks/schema/consensuscore/kinds"
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
type NewDurationPhase struct {
	Level                int64  `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle                 int32  `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	Phase                  uint32 `protobuf:"variableint,3,opt,name=step,proto3" json:"phase,omitempty"`
	MomentsSinceBeginTime int64  `protobuf:"variableint,4,opt,name=seconds_since_start_time,json=secondsSinceStartTime,proto3" json:"moments_since_begin_time,omitempty"`
	FinalEndorseDuration       int32  `protobuf:"variableint,5,opt,name=last_commit_round,json=lastCommitRound,proto3" json:"final_endorse_epoch,omitempty"`
}

func (m *NewDurationPhase) Restore()         { *m = NewDurationPhase{} }
func (m *NewDurationPhase) String() string { return proto.CompactTextString(m) }
func (*NewDurationPhase) SchemaSignal()    {}
func (*NewDurationPhase) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{0}
}
func (m *NewDurationPhase) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *NewDurationPhase) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Newepochphase.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NewDurationPhase) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Newepochphase.Merge(m, src)
}
func (m *NewDurationPhase) XXX_Volume() int {
	return m.Volume()
}
func (m *NewDurationPhase) XXX_Omitunclear() {
	xxx_messagedata_Newepochphase.DiscardUnknown(m)
}

var xxx_messagedata_Newepochphase proto.InternalMessageInfo

func (m *NewDurationPhase) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *NewDurationPhase) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *NewDurationPhase) FetchPhase() uint32 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *NewDurationPhase) FetchMomentsSinceBeginTime() int64 {
	if m != nil {
		return m.MomentsSinceBeginTime
	}
	return 0
}

func (m *NewDurationPhase) FetchFinalEndorseEpoch() int32 {
	if m != nil {
		return m.FinalEndorseDuration
	}
	return 0
}

//
//
//
type NewSoundLedger struct {
	Level             int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle              int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	LedgerSegmentAssignHeading kinds.SegmentAssignHeading `protobuf:"octets,3,opt,name=block_part_set_header,json=blockPartSetHeader,proto3" json:"ledger_section_collection_heading"`
	LedgerSegments         *bits.BitList      `protobuf:"octets,4,opt,name=block_parts,json=blockParts,proto3" json:"ledger_segments,omitempty"`
	IsEndorse           bool                `protobuf:"variableint,5,opt,name=is_commit,json=isCommit,proto3" json:"is_endorse,omitempty"`
}

func (m *NewSoundLedger) Restore()         { *m = NewSoundLedger{} }
func (m *NewSoundLedger) String() string { return proto.CompactTextString(m) }
func (*NewSoundLedger) SchemaSignal()    {}
func (*NewSoundLedger) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{1}
}
func (m *NewSoundLedger) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *NewSoundLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Newvalidledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NewSoundLedger) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Newvalidledger.Merge(m, src)
}
func (m *NewSoundLedger) XXX_Volume() int {
	return m.Volume()
}
func (m *NewSoundLedger) XXX_Omitunclear() {
	xxx_messagedata_Newvalidledger.DiscardUnknown(m)
}

var xxx_messagedata_Newvalidledger proto.InternalMessageInfo

func (m *NewSoundLedger) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *NewSoundLedger) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *NewSoundLedger) FetchLedgerSectionCollectionHeading() kinds.SegmentAssignHeading {
	if m != nil {
		return m.LedgerSegmentAssignHeading
	}
	return kinds.SegmentAssignHeading{}
}

func (m *NewSoundLedger) FetchLedgerSections() *bits.BitList {
	if m != nil {
		return m.LedgerSegments
	}
	return nil
}

func (m *NewSoundLedger) FetchIsEndorse() bool {
	if m != nil {
		return m.IsEndorse
	}
	return false
}

//
type Nomination struct {
	Nomination kinds.Nomination `protobuf:"octets,1,opt,name=proposal,proto3" json:"nomination"`
}

func (m *Nomination) Restore()         { *m = Nomination{} }
func (m *Nomination) String() string { return proto.CompactTextString(m) }
func (*Nomination) SchemaSignal()    {}
func (*Nomination) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{2}
}
func (m *Nomination) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Nomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Nomination.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nomination) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Nomination.Merge(m, src)
}
func (m *Nomination) XXX_Volume() int {
	return m.Volume()
}
func (m *Nomination) XXX_Omitunclear() {
	xxx_messagedata_Nomination.DiscardUnknown(m)
}

var xxx_messagedata_Nomination proto.InternalMessageInfo

func (m *Nomination) FetchNomination() kinds.Nomination {
	if m != nil {
		return m.Nomination
	}
	return kinds.Nomination{}
}

//
type NominationPOL struct {
	Level           int64         `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	NominationPolDuration int32         `protobuf:"variableint,2,opt,name=proposal_pol_round,json=proposalPolRound,proto3" json:"nomination_pol_epoch,omitempty"`
	NominationPol      bits.BitList `protobuf:"octets,3,opt,name=proposal_pol,json=proposalPol,proto3" json:"nomination_pol"`
}

func (m *NominationPOL) Restore()         { *m = NominationPOL{} }
func (m *NominationPOL) String() string { return proto.CompactTextString(m) }
func (*NominationPOL) SchemaSignal()    {}
func (*NominationPOL) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{3}
}
func (m *NominationPOL) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *NominationPOL) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Nominationpol.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NominationPOL) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Nominationpol.Merge(m, src)
}
func (m *NominationPOL) XXX_Volume() int {
	return m.Volume()
}
func (m *NominationPOL) XXX_Omitunclear() {
	xxx_messagedata_Nominationpol.DiscardUnknown(m)
}

var xxx_messagedata_Nominationpol proto.InternalMessageInfo

func (m *NominationPOL) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *NominationPOL) FetchNominationPolEpoch() int32 {
	if m != nil {
		return m.NominationPolDuration
	}
	return 0
}

func (m *NominationPOL) FetchNominationPol() bits.BitList {
	if m != nil {
		return m.NominationPol
	}
	return bits.BitList{}
}

//
type LedgerSegment struct {
	Level int64      `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle  int32      `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	Segment   kinds.Segment `protobuf:"octets,3,opt,name=part,proto3" json:"segment"`
}

func (m *LedgerSegment) Restore()         { *m = LedgerSegment{} }
func (m *LedgerSegment) String() string { return proto.CompactTextString(m) }
func (*LedgerSegment) SchemaSignal()    {}
func (*LedgerSegment) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{4}
}
func (m *LedgerSegment) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *LedgerSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ledgersection.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerSegment) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ledgersection.Merge(m, src)
}
func (m *LedgerSegment) XXX_Volume() int {
	return m.Volume()
}
func (m *LedgerSegment) XXX_Omitunclear() {
	xxx_messagedata_Ledgersection.DiscardUnknown(m)
}

var xxx_messagedata_Ledgersection proto.InternalMessageInfo

func (m *LedgerSegment) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LedgerSegment) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *LedgerSegment) FetchSegment() kinds.Segment {
	if m != nil {
		return m.Segment
	}
	return kinds.Segment{}
}

//
type Ballot struct {
	Ballot *kinds.Ballot `protobuf:"octets,1,opt,name=vote,proto3" json:"ballot,omitempty"`
}

func (m *Ballot) Restore()         { *m = Ballot{} }
func (m *Ballot) String() string { return proto.CompactTextString(m) }
func (*Ballot) SchemaSignal()    {}
func (*Ballot) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{5}
}
func (m *Ballot) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Ballot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ballot) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ballot.Merge(m, src)
}
func (m *Ballot) XXX_Volume() int {
	return m.Volume()
}
func (m *Ballot) XXX_Omitunclear() {
	xxx_messagedata_Ballot.DiscardUnknown(m)
}

var xxx_messagedata_Ballot proto.InternalMessageInfo

func (m *Ballot) FetchBallot() *kinds.Ballot {
	if m != nil {
		return m.Ballot
	}
	return nil
}

//
type HasBallot struct {
	Level int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle  int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	Kind   kinds.AttestedMessageKind `protobuf:"variableint,3,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Ordinal  int32               `protobuf:"variableint,4,opt,name=index,proto3" json:"ordinal,omitempty"`
}

func (m *HasBallot) Restore()         { *m = HasBallot{} }
func (m *HasBallot) String() string { return proto.CompactTextString(m) }
func (*HasBallot) SchemaSignal()    {}
func (*HasBallot) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{6}
}
func (m *HasBallot) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *HasBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Hasballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HasBallot) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Hasballot.Merge(m, src)
}
func (m *HasBallot) XXX_Volume() int {
	return m.Volume()
}
func (m *HasBallot) XXX_Omitunclear() {
	xxx_messagedata_Hasballot.DiscardUnknown(m)
}

var xxx_messagedata_Hasballot proto.InternalMessageInfo

func (m *HasBallot) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *HasBallot) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *HasBallot) FetchKind() kinds.AttestedMessageKind {
	if m != nil {
		return m.Kind
	}
	return kinds.UnclearKind
}

func (m *HasBallot) FetchOrdinal() int32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

//
type BallotAssignMaj23 struct {
	Level  int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle   int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	Kind    kinds.AttestedMessageKind `protobuf:"variableint,3,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	LedgerUID kinds.LedgerUID       `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uid"`
}

func (m *BallotAssignMaj23) Restore()         { *m = BallotAssignMaj23{} }
func (m *BallotAssignMaj23) String() string { return proto.CompactTextString(m) }
func (*BallotAssignMaj23) SchemaSignal()    {}
func (*BallotAssignMaj23) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{7}
}
func (m *BallotAssignMaj23) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *BallotAssignMaj23) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ballotsetmaj23.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BallotAssignMaj23) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ballotsetmaj23.Merge(m, src)
}
func (m *BallotAssignMaj23) XXX_Volume() int {
	return m.Volume()
}
func (m *BallotAssignMaj23) XXX_Omitunclear() {
	xxx_messagedata_Ballotsetmaj23.DiscardUnknown(m)
}

var xxx_messagedata_Ballotsetmaj23 proto.InternalMessageInfo

func (m *BallotAssignMaj23) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *BallotAssignMaj23) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *BallotAssignMaj23) FetchKind() kinds.AttestedMessageKind {
	if m != nil {
		return m.Kind
	}
	return kinds.UnclearKind
}

func (m *BallotAssignMaj23) FetchLedgerUID() kinds.LedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return kinds.LedgerUID{}
}

//
type BallotAssignBits struct {
	Level  int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle   int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	Kind    kinds.AttestedMessageKind `protobuf:"variableint,3,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	LedgerUID kinds.LedgerUID       `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uid"`
	Ballots   bits.BitList       `protobuf:"octets,5,opt,name=votes,proto3" json:"ballots"`
}

func (m *BallotAssignBits) Restore()         { *m = BallotAssignBits{} }
func (m *BallotAssignBits) String() string { return proto.CompactTextString(m) }
func (*BallotAssignBits) SchemaSignal()    {}
func (*BallotAssignBits) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{8}
}
func (m *BallotAssignBits) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *BallotAssignBits) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ballotsetbits.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BallotAssignBits) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ballotsetbits.Merge(m, src)
}
func (m *BallotAssignBits) XXX_Volume() int {
	return m.Volume()
}
func (m *BallotAssignBits) XXX_Omitunclear() {
	xxx_messagedata_Ballotsetbits.DiscardUnknown(m)
}

var xxx_messagedata_Ballotsetbits proto.InternalMessageInfo

func (m *BallotAssignBits) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *BallotAssignBits) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *BallotAssignBits) FetchKind() kinds.AttestedMessageKind {
	if m != nil {
		return m.Kind
	}
	return kinds.UnclearKind
}

func (m *BallotAssignBits) FetchLedgerUID() kinds.LedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return kinds.LedgerUID{}
}

func (m *BallotAssignBits) FetchBallots() bits.BitList {
	if m != nil {
		return m.Ballots
	}
	return bits.BitList{}
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
	Sum ismessage_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) SchemaSignal()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedefinition_81a22d2efc008981, []int{9}
}
func (m *Signal) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Signal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Signal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signal) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Signal.Merge(m, src)
}
func (m *Signal) XXX_Volume() int {
	return m.Volume()
}
func (m *Signal) XXX_Omitunclear() {
	xxx_messagedata_Signal.DiscardUnknown(m)
}

var xxx_messagedata_Signal proto.InternalMessageInfo

type ismessage_Total interface {
	ismessage_Total()
	SerializeTo([]byte) (int, error)
	Volume() int
}

type Signal_Newepochphase struct {
	NewDurationPhase *NewDurationPhase `protobuf:"octets,1,opt,name=new_round_step,json=newRoundStep,proto3,oneof" json:"new_epoch_phase,omitempty"`
}
type Signal_Newvalidledger struct {
	NewSoundLedger *NewSoundLedger `protobuf:"octets,2,opt,name=new_valid_block,json=newValidBlock,proto3,oneof" json:"new_sound_ledger,omitempty"`
}
type Signal_Nomination struct {
	Nomination *Nomination `protobuf:"octets,3,opt,name=proposal,proto3,oneof" json:"nomination,omitempty"`
}
type Signal_Nominationpol struct {
	NominationPol *NominationPOL `protobuf:"octets,4,opt,name=proposal_pol,json=proposalPol,proto3,oneof" json:"nomination_pol,omitempty"`
}
type Signal_Ledgersection struct {
	LedgerSegment *LedgerSegment `protobuf:"octets,5,opt,name=block_part,json=blockPart,proto3,oneof" json:"ledger_section,omitempty"`
}
type Signal_Ballot struct {
	Ballot *Ballot `protobuf:"octets,6,opt,name=vote,proto3,oneof" json:"ballot,omitempty"`
}
type Signal_Hasballot struct {
	HasBallot *HasBallot `protobuf:"octets,7,opt,name=has_vote,json=hasVote,proto3,oneof" json:"has_ballot,omitempty"`
}
type Signal_Ballotsetmaj23 struct {
	BallotAssignMaj23 *BallotAssignMaj23 `protobuf:"octets,8,opt,name=vote_set_maj23,json=voteSetMaj23,proto3,oneof" json:"ballot_collection_maj23,omitempty"`
}
type Signal_Ballotsetbits struct {
	BallotAssignBits *BallotAssignBits `protobuf:"octets,9,opt,name=vote_set_bits,json=voteSetBits,proto3,oneof" json:"ballot_collection_bits,omitempty"`
}

func (*Signal_Newepochphase) ismessage_Total()  {}
func (*Signal_Newvalidledger) ismessage_Total() {}
func (*Signal_Nomination) ismessage_Total()      {}
func (*Signal_Nominationpol) ismessage_Total()   {}
func (*Signal_Ledgersection) ismessage_Total()     {}
func (*Signal_Ballot) ismessage_Total()          {}
func (*Signal_Hasballot) ismessage_Total()       {}
func (*Signal_Ballotsetmaj23) ismessage_Total()  {}
func (*Signal_Ballotsetbits) ismessage_Total()   {}

func (m *Signal) FetchTotal() ismessage_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) FetchNewEpochPhase() *NewDurationPhase {
	if x, ok := m.FetchTotal().(*Signal_Newepochphase); ok {
		return x.NewDurationPhase
	}
	return nil
}

func (m *Signal) FetchNewSoundLedger() *NewSoundLedger {
	if x, ok := m.FetchTotal().(*Signal_Newvalidledger); ok {
		return x.NewSoundLedger
	}
	return nil
}

func (m *Signal) FetchNomination() *Nomination {
	if x, ok := m.FetchTotal().(*Signal_Nomination); ok {
		return x.Nomination
	}
	return nil
}

func (m *Signal) FetchNominationPol() *NominationPOL {
	if x, ok := m.FetchTotal().(*Signal_Nominationpol); ok {
		return x.NominationPol
	}
	return nil
}

func (m *Signal) FetchLedgerSection() *LedgerSegment {
	if x, ok := m.FetchTotal().(*Signal_Ledgersection); ok {
		return x.LedgerSegment
	}
	return nil
}

func (m *Signal) FetchBallot() *Ballot {
	if x, ok := m.FetchTotal().(*Signal_Ballot); ok {
		return x.Ballot
	}
	return nil
}

func (m *Signal) FetchHasBallot() *HasBallot {
	if x, ok := m.FetchTotal().(*Signal_Hasballot); ok {
		return x.HasBallot
	}
	return nil
}

func (m *Signal) FetchBallotCollectionMaj23() *BallotAssignMaj23 {
	if x, ok := m.FetchTotal().(*Signal_Ballotsetmaj23); ok {
		return x.BallotAssignMaj23
	}
	return nil
}

func (m *Signal) FetchBallotCollectionBits() *BallotAssignBits {
	if x, ok := m.FetchTotal().(*Signal_Ballotsetbits); ok {
		return x.BallotAssignBits
	}
	return nil
}

//
func (*Signal) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Signal_Newepochphase)(nil),
		(*Signal_Newvalidledger)(nil),
		(*Signal_Nomination)(nil),
		(*Signal_Nominationpol)(nil),
		(*Signal_Ledgersection)(nil),
		(*Signal_Ballot)(nil),
		(*Signal_Hasballot)(nil),
		(*Signal_Ballotsetmaj23)(nil),
		(*Signal_Ballotsetbits)(nil),
	}
}

func init() {
	proto.RegisterType((*NewDurationPhase)(nil), "REDACTED")
	proto.RegisterType((*NewSoundLedger)(nil), "REDACTED")
	proto.RegisterType((*Nomination)(nil), "REDACTED")
	proto.RegisterType((*NominationPOL)(nil), "REDACTED")
	proto.RegisterType((*LedgerSegment)(nil), "REDACTED")
	proto.RegisterType((*Ballot)(nil), "REDACTED")
	proto.RegisterType((*HasBallot)(nil), "REDACTED")
	proto.RegisterType((*BallotAssignMaj23)(nil), "REDACTED")
	proto.RegisterType((*BallotAssignBits)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_81a22d2efc008981) }

var filedefinition_81a22d2efc008981 = []byte{
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

func (m *NewDurationPhase) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewDurationPhase) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *NewDurationPhase) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalEndorseDuration != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FinalEndorseDuration))
		i--
		dAtA[i] = 0x28
	}
	if m.MomentsSinceBeginTime != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.MomentsSinceBeginTime))
		i--
		dAtA[i] = 0x20
	}
	if m.Phase != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x18
	}
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NewSoundLedger) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewSoundLedger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *NewSoundLedger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsEndorse {
		i--
		if m.IsEndorse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.LedgerSegments != nil {
		{
			volume, err := m.LedgerSegments.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		volume, err := m.LedgerSegmentAssignHeading.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x1a
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nomination) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Nomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.Nomination.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *NominationPOL) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NominationPOL) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *NominationPOL) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.NominationPol.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x1a
	if m.NominationPolDuration != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.NominationPolDuration))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LedgerSegment) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LedgerSegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *LedgerSegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.Segment.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x1a
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Ballot) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ballot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Ballot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ballot != nil {
		{
			volume, err := m.Ballot.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HasBallot) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasBallot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *HasBallot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ordinal != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ordinal))
		i--
		dAtA[i] = 0x20
	}
	if m.Kind != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x18
	}
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BallotAssignMaj23) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BallotAssignMaj23) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *BallotAssignMaj23) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x22
	if m.Kind != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x18
	}
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BallotAssignBits) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BallotAssignBits) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *BallotAssignBits) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.Ballots.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x2a
	{
		volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x22
	if m.Kind != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x18
	}
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Signal) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signal) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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

func (m *Signal_Newepochphase) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Newepochphase) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NewDurationPhase != nil {
		{
			volume, err := m.NewDurationPhase.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Newvalidledger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Newvalidledger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NewSoundLedger != nil {
		{
			volume, err := m.NewSoundLedger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Nomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Nomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Nomination != nil {
		{
			volume, err := m.Nomination.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Nominationpol) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Nominationpol) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NominationPol != nil {
		{
			volume, err := m.NominationPol.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Ledgersection) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Ledgersection) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LedgerSegment != nil {
		{
			volume, err := m.LedgerSegment.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Ballot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Ballot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ballot != nil {
		{
			volume, err := m.Ballot.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Hasballot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Hasballot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HasBallot != nil {
		{
			volume, err := m.HasBallot.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Ballotsetmaj23) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Ballotsetmaj23) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BallotAssignMaj23 != nil {
		{
			volume, err := m.BallotAssignMaj23.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Ballotsetbits) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Ballotsetbits) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BallotAssignBits != nil {
		{
			volume, err := m.BallotAssignBits.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func formatVariableintKinds(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovKinds(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *NewDurationPhase) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	if m.Phase != 0 {
		n += 1 + sovKinds(uint64(m.Phase))
	}
	if m.MomentsSinceBeginTime != 0 {
		n += 1 + sovKinds(uint64(m.MomentsSinceBeginTime))
	}
	if m.FinalEndorseDuration != 0 {
		n += 1 + sovKinds(uint64(m.FinalEndorseDuration))
	}
	return n
}

func (m *NewSoundLedger) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	l = m.LedgerSegmentAssignHeading.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.LedgerSegments != nil {
		l = m.LedgerSegments.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.IsEndorse {
		n += 2
	}
	return n
}

func (m *Nomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Nomination.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *NominationPOL) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.NominationPolDuration != 0 {
		n += 1 + sovKinds(uint64(m.NominationPolDuration))
	}
	l = m.NominationPol.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *LedgerSegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	l = m.Segment.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *Ballot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ballot != nil {
		l = m.Ballot.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *HasBallot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	return n
}

func (m *BallotAssignMaj23) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	l = m.LedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *BallotAssignBits) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	l = m.LedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = m.Ballots.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *Signal) Volume() (n int) {
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

func (m *Signal_Newepochphase) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NewDurationPhase != nil {
		l = m.NewDurationPhase.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Newvalidledger) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NewSoundLedger != nil {
		l = m.NewSoundLedger.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Nomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nomination != nil {
		l = m.Nomination.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Nominationpol) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NominationPol != nil {
		l = m.NominationPol.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ledgersection) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerSegment != nil {
		l = m.LedgerSegment.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ballot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ballot != nil {
		l = m.Ballot.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Hasballot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HasBallot != nil {
		l = m.HasBallot.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ballotsetmaj23) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotAssignMaj23 != nil {
		l = m.BallotAssignMaj23.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ballotsetbits) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotAssignBits != nil {
		l = m.BallotAssignBits.Volume()
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
func (m *NewDurationPhase) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Phase = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Phase |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.MomentsSinceBeginTime = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MomentsSinceBeginTime |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalEndorseDuration = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FinalEndorseDuration |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *NewSoundLedger) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LedgerSegmentAssignHeading.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.LedgerSegments == nil {
				m.LedgerSegments = &bits.BitList{}
			}
			if err := m.LedgerSegments.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var v int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
			m.IsEndorse = bool(v != 0)
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *Nomination) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Nomination.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *NominationPOL) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.NominationPolDuration = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.NominationPolDuration |= int32(b&0x7F) << displace
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NominationPol.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *LedgerSegment) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Segment.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *Ballot) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ballot == nil {
				m.Ballot = &kinds.Ballot{}
			}
			if err := m.Ballot.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *HasBallot) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Kind = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Kind |= kinds.AttestedMessageKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Ordinal |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *BallotAssignMaj23) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Kind = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Kind |= kinds.AttestedMessageKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LedgerUID.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *BallotAssignBits) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Kind = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Kind |= kinds.AttestedMessageKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ballots.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *Signal) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &NewDurationPhase{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Newepochphase{v}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &NewSoundLedger{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Newvalidledger{v}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &Nomination{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Nomination{v}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &NominationPOL{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Nominationpol{v}
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &LedgerSegment{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ledgersection{v}
			idxNdEx = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &Ballot{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ballot{v}
			idxNdEx = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &HasBallot{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Hasballot{v}
			idxNdEx = submitOrdinal
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &BallotAssignMaj23{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ballotsetmaj23{v}
			idxNdEx = submitOrdinal
		case 9:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &BallotAssignBits{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ballotsetbits{v}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func omitKinds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadKinds
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
					return 0, ErrIntegerOverloadKinds
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
					return 0, ErrIntegerOverloadKinds
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
				return 0, ErrCorruptExtentKinds
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterKinds
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentKinds
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentKinds        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadKinds          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterKinds = fmt.Errorf("REDACTED")
)
