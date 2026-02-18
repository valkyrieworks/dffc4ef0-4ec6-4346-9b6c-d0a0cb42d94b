//
//

package kinds

import (
	fmt "fmt"
	vault "github.com/valkyrieworks/schema/consensuscore/vault"
	release "github.com/valkyrieworks/schema/consensuscore/release"
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

//
type AttestedMessageKind int32

const (
	UnclearKind AttestedMessageKind = 0
	//
	PreballotKind   AttestedMessageKind = 1
	PreendorseKind AttestedMessageKind = 2
	//
	NominationKind AttestedMessageKind = 32
)

var Attestedsignaltype_label = map[int32]string{
	0:  "REDACTED",
	1:  "REDACTED",
	2:  "REDACTED",
	32: "REDACTED",
}

var Attestedsignaltype_item = map[string]int32{
	"REDACTED":   0,
	"REDACTED":   1,
	"REDACTED": 2,
	"REDACTED":  32,
}

func (x AttestedMessageKind) String() string {
	return proto.EnumName(Attestedsignaltype_label, int32(x))
}

func (AttestedMessageKind) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{0}
}

//
type SegmentAssignHeading struct {
	Sum uint32 `protobuf:"variableint,1,opt,name=total,proto3" json:"sum,omitempty"`
	Digest  []byte `protobuf:"octets,2,opt,name=hash,proto3" json:"digest,omitempty"`
}

func (m *SegmentAssignHeading) Restore()         { *m = SegmentAssignHeading{} }
func (m *SegmentAssignHeading) String() string { return proto.CompactTextString(m) }
func (*SegmentAssignHeading) SchemaSignal()    {}
func (*SegmentAssignHeading) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{0}
}
func (m *SegmentAssignHeading) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *SegmentAssignHeading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Sectionsetheading.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SegmentAssignHeading) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Sectionsetheading.Merge(m, src)
}
func (m *SegmentAssignHeading) XXX_Volume() int {
	return m.Volume()
}
func (m *SegmentAssignHeading) XXX_Omitunclear() {
	xxx_messagedata_Sectionsetheading.DiscardUnknown(m)
}

var xxx_messagedata_Sectionsetheading proto.InternalMessageInfo

func (m *SegmentAssignHeading) FetchSum() uint32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *SegmentAssignHeading) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Segment struct {
	Ordinal uint32       `protobuf:"variableint,1,opt,name=index,proto3" json:"ordinal,omitempty"`
	Octets []byte       `protobuf:"octets,2,opt,name=bytes,proto3" json:"octets,omitempty"`
	Attestation vault.Attestation `protobuf:"octets,3,opt,name=proof,proto3" json:"evidence"`
}

func (m *Segment) Restore()         { *m = Segment{} }
func (m *Segment) String() string { return proto.CompactTextString(m) }
func (*Segment) SchemaSignal()    {}
func (*Segment) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{1}
}
func (m *Segment) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Segment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Section.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Segment) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Section.Merge(m, src)
}
func (m *Segment) XXX_Volume() int {
	return m.Volume()
}
func (m *Segment) XXX_Omitunclear() {
	xxx_messagedata_Section.DiscardUnknown(m)
}

var xxx_messagedata_Section proto.InternalMessageInfo

func (m *Segment) FetchOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *Segment) FetchOctets() []byte {
	if m != nil {
		return m.Octets
	}
	return nil
}

func (m *Segment) FetchEvidence() vault.Attestation {
	if m != nil {
		return m.Attestation
	}
	return vault.Attestation{}
}

//
type LedgerUID struct {
	Digest          []byte        `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	SegmentAssignHeading SegmentAssignHeading `protobuf:"octets,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"section_collection_heading"`
}

func (m *LedgerUID) Restore()         { *m = LedgerUID{} }
func (m *LedgerUID) String() string { return proto.CompactTextString(m) }
func (*LedgerUID) SchemaSignal()    {}
func (*LedgerUID) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{2}
}
func (m *LedgerUID) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *LedgerUID) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ledgeruid.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerUID) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ledgeruid.Merge(m, src)
}
func (m *LedgerUID) XXX_Volume() int {
	return m.Volume()
}
func (m *LedgerUID) XXX_Omitunclear() {
	xxx_messagedata_Ledgeruid.DiscardUnknown(m)
}

var xxx_messagedata_Ledgeruid proto.InternalMessageInfo

func (m *LedgerUID) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *LedgerUID) FetchSectionCollectionHeading() SegmentAssignHeading {
	if m != nil {
		return m.SegmentAssignHeading
	}
	return SegmentAssignHeading{}
}

//
type Heading struct {
	//
	Release release.Agreement `protobuf:"octets,1,opt,name=version,proto3" json:"release"`
	LedgerUID string            `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
	Level  int64             `protobuf:"variableint,3,opt,name=height,proto3" json:"level,omitempty"`
	Time    time.Time         `protobuf:"octets,4,opt,name=time,proto3,stdtime" json:"moment"`
	//
	FinalLedgerUid LedgerUID `protobuf:"octets,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"final_ledger_uid"`
	//
	FinalEndorseDigest []byte `protobuf:"octets,6,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"final_endorse_digest,omitempty"`
	DataDigest       []byte `protobuf:"octets,7,opt,name=data_hash,json=dataHash,proto3" json:"data_digest,omitempty"`
	//
	RatifiersDigest     []byte `protobuf:"octets,8,opt,name=validators_hash,json=validatorsHash,proto3" json:"ratifiers_digest,omitempty"`
	FollowingRatifiersDigest []byte `protobuf:"octets,9,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_ratifiers_digest,omitempty"`
	AgreementDigest      []byte `protobuf:"octets,10,opt,name=consensus_hash,json=consensusHash,proto3" json:"agreement_digest,omitempty"`
	ApplicationDigest            []byte `protobuf:"octets,11,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
	FinalOutcomesDigest    []byte `protobuf:"octets,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"final_outcomes_digest,omitempty"`
	//
	ProofDigest    []byte `protobuf:"octets,13,opt,name=evidence_hash,json=evidenceHash,proto3" json:"proof_digest,omitempty"`
	RecommenderLocation []byte `protobuf:"octets,14,opt,name=proposer_address,json=proposerAddress,proto3" json:"recommender_location,omitempty"`
}

func (m *Heading) Restore()         { *m = Heading{} }
func (m *Heading) String() string { return proto.CompactTextString(m) }
func (*Heading) SchemaSignal()    {}
func (*Heading) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{3}
}
func (m *Heading) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Heading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Heading.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Heading) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Heading.Merge(m, src)
}
func (m *Heading) XXX_Volume() int {
	return m.Volume()
}
func (m *Heading) XXX_Omitunclear() {
	xxx_messagedata_Heading.DiscardUnknown(m)
}

var xxx_messagedata_Heading proto.InternalMessageInfo

func (m *Heading) FetchRelease() release.Agreement {
	if m != nil {
		return m.Release
	}
	return release.Agreement{}
}

func (m *Heading) FetchSeriesUID() string {
	if m != nil {
		return m.LedgerUID
	}
	return "REDACTED"
}

func (m *Heading) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Heading) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Heading) FetchFinalLedgerUid() LedgerUID {
	if m != nil {
		return m.FinalLedgerUid
	}
	return LedgerUID{}
}

func (m *Heading) FetchFinalEndorseDigest() []byte {
	if m != nil {
		return m.FinalEndorseDigest
	}
	return nil
}

func (m *Heading) FetchDataDigest() []byte {
	if m != nil {
		return m.DataDigest
	}
	return nil
}

func (m *Heading) FetchRatifiersDigest() []byte {
	if m != nil {
		return m.RatifiersDigest
	}
	return nil
}

func (m *Heading) FetchFollowingRatifiersDigest() []byte {
	if m != nil {
		return m.FollowingRatifiersDigest
	}
	return nil
}

func (m *Heading) FetchAgreementDigest() []byte {
	if m != nil {
		return m.AgreementDigest
	}
	return nil
}

func (m *Heading) FetchApplicationDigest() []byte {
	if m != nil {
		return m.ApplicationDigest
	}
	return nil
}

func (m *Heading) FetchFinalOutcomesDigest() []byte {
	if m != nil {
		return m.FinalOutcomesDigest
	}
	return nil
}

func (m *Heading) FetchProofDigest() []byte {
	if m != nil {
		return m.ProofDigest
	}
	return nil
}

func (m *Heading) FetchRecommenderLocation() []byte {
	if m != nil {
		return m.RecommenderLocation
	}
	return nil
}

//
type Data struct {
	//
	//
	//
	Txs [][]byte `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *Data) Restore()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) SchemaSignal()    {}
func (*Data) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{4}
}
func (m *Data) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Data) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Data.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Data) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Data.Merge(m, src)
}
func (m *Data) XXX_Volume() int {
	return m.Volume()
}
func (m *Data) XXX_Omitunclear() {
	xxx_messagedata_Data.DiscardUnknown(m)
}

var xxx_messagedata_Data proto.InternalMessageInfo

func (m *Data) FetchTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

//
//
type Ballot struct {
	Kind             AttestedMessageKind `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Level           int64         `protobuf:"variableint,2,opt,name=height,proto3" json:"level,omitempty"`
	Cycle            int32         `protobuf:"variableint,3,opt,name=round,proto3" json:"epoch,omitempty"`
	LedgerUID          LedgerUID       `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uid"`
	Timestamp        time.Time     `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	RatifierLocation []byte        `protobuf:"octets,6,opt,name=validator_address,json=validatorAddress,proto3" json:"ratifier_location,omitempty"`
	RatifierOrdinal   int32         `protobuf:"variableint,7,opt,name=validator_index,json=validatorIndex,proto3" json:"ratifier_ordinal,omitempty"`
	//
	//
	Autograph []byte `protobuf:"octets,8,opt,name=signature,proto3" json:"autograph,omitempty"`
	//
	//
	Addition []byte `protobuf:"octets,9,opt,name=extension,proto3" json:"addition,omitempty"`
	//
	//
	//
	AdditionAutograph []byte `protobuf:"octets,10,opt,name=extension_signature,json=extensionSignature,proto3" json:"addition_autograph,omitempty"`
}

func (m *Ballot) Restore()         { *m = Ballot{} }
func (m *Ballot) String() string { return proto.CompactTextString(m) }
func (*Ballot) SchemaSignal()    {}
func (*Ballot) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{5}
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

func (m *Ballot) FetchKind() AttestedMessageKind {
	if m != nil {
		return m.Kind
	}
	return UnclearKind
}

func (m *Ballot) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Ballot) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *Ballot) FetchLedgerUID() LedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return LedgerUID{}
}

func (m *Ballot) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Ballot) FetchRatifierLocation() []byte {
	if m != nil {
		return m.RatifierLocation
	}
	return nil
}

func (m *Ballot) FetchRatifierOrdinal() int32 {
	if m != nil {
		return m.RatifierOrdinal
	}
	return 0
}

func (m *Ballot) FetchAutograph() []byte {
	if m != nil {
		return m.Autograph
	}
	return nil
}

func (m *Ballot) FetchAddition() []byte {
	if m != nil {
		return m.Addition
	}
	return nil
}

func (m *Ballot) FetchAdditionAutograph() []byte {
	if m != nil {
		return m.AdditionAutograph
	}
	return nil
}

//
type Endorse struct {
	Level     int64       `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle      int32       `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	LedgerUID    LedgerUID     `protobuf:"octets,3,opt,name=block_id,json=blockId,proto3" json:"ledger_uid"`
	Endorsements []EndorseSignature `protobuf:"octets,4,rep,name=signatures,proto3" json:"endorsements"`
}

func (m *Endorse) Restore()         { *m = Endorse{} }
func (m *Endorse) String() string { return proto.CompactTextString(m) }
func (*Endorse) SchemaSignal()    {}
func (*Endorse) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{6}
}
func (m *Endorse) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Endorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Endorse.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Endorse) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Endorse.Merge(m, src)
}
func (m *Endorse) XXX_Volume() int {
	return m.Volume()
}
func (m *Endorse) XXX_Omitunclear() {
	xxx_messagedata_Endorse.DiscardUnknown(m)
}

var xxx_messagedata_Endorse proto.InternalMessageInfo

func (m *Endorse) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Endorse) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *Endorse) FetchLedgerUID() LedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return LedgerUID{}
}

func (m *Endorse) FetchEndorsements() []EndorseSignature {
	if m != nil {
		return m.Endorsements
	}
	return nil
}

//
type EndorseSignature struct {
	LedgerUidMark      LedgerUIDMark `protobuf:"variableint,1,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uid_mark,omitempty"`
	RatifierLocation []byte      `protobuf:"octets,2,opt,name=validator_address,json=validatorAddress,proto3" json:"ratifier_location,omitempty"`
	Timestamp        time.Time   `protobuf:"octets,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Autograph        []byte      `protobuf:"octets,4,opt,name=signature,proto3" json:"autograph,omitempty"`
}

func (m *EndorseSignature) Restore()         { *m = EndorseSignature{} }
func (m *EndorseSignature) String() string { return proto.CompactTextString(m) }
func (*EndorseSignature) SchemaSignal()    {}
func (*EndorseSignature) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{7}
}
func (m *EndorseSignature) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *EndorseSignature) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Endorsesignal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndorseSignature) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Endorsesignal.Merge(m, src)
}
func (m *EndorseSignature) XXX_Volume() int {
	return m.Volume()
}
func (m *EndorseSignature) XXX_Omitunclear() {
	xxx_messagedata_Endorsesignal.DiscardUnknown(m)
}

var xxx_messagedata_Endorsesignal proto.InternalMessageInfo

func (m *EndorseSignature) FetchLedgerUidMark() LedgerUIDMark {
	if m != nil {
		return m.LedgerUidMark
	}
	return LedgerUIDMarkUnclear
}

func (m *EndorseSignature) FetchRatifierLocation() []byte {
	if m != nil {
		return m.RatifierLocation
	}
	return nil
}

func (m *EndorseSignature) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *EndorseSignature) FetchAutograph() []byte {
	if m != nil {
		return m.Autograph
	}
	return nil
}

type ExpandedEndorse struct {
	Level             int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle              int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	LedgerUID            LedgerUID             `protobuf:"octets,3,opt,name=block_id,json=blockId,proto3" json:"ledger_uid"`
	ExpandedEndorsements []ExpandedEndorseSignature `protobuf:"octets,4,rep,name=extended_signatures,json=extendedSignatures,proto3" json:"expanded_endorsements"`
}

func (m *ExpandedEndorse) Restore()         { *m = ExpandedEndorse{} }
func (m *ExpandedEndorse) String() string { return proto.CompactTextString(m) }
func (*ExpandedEndorse) SchemaSignal()    {}
func (*ExpandedEndorse) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{8}
}
func (m *ExpandedEndorse) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ExpandedEndorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Expandedendorse.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedEndorse) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Expandedendorse.Merge(m, src)
}
func (m *ExpandedEndorse) XXX_Volume() int {
	return m.Volume()
}
func (m *ExpandedEndorse) XXX_Omitunclear() {
	xxx_messagedata_Expandedendorse.DiscardUnknown(m)
}

var xxx_messagedata_Expandedendorse proto.InternalMessageInfo

func (m *ExpandedEndorse) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *ExpandedEndorse) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *ExpandedEndorse) FetchLedgerUID() LedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return LedgerUID{}
}

func (m *ExpandedEndorse) FetchExpandedEndorsements() []ExpandedEndorseSignature {
	if m != nil {
		return m.ExpandedEndorsements
	}
	return nil
}

//
//
//
type ExpandedEndorseSignature struct {
	LedgerUidMark      LedgerUIDMark `protobuf:"variableint,1,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uid_mark,omitempty"`
	RatifierLocation []byte      `protobuf:"octets,2,opt,name=validator_address,json=validatorAddress,proto3" json:"ratifier_location,omitempty"`
	Timestamp        time.Time   `protobuf:"octets,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Autograph        []byte      `protobuf:"octets,4,opt,name=signature,proto3" json:"autograph,omitempty"`
	//
	Addition []byte `protobuf:"octets,5,opt,name=extension,proto3" json:"addition,omitempty"`
	//
	AdditionAutograph []byte `protobuf:"octets,6,opt,name=extension_signature,json=extensionSignature,proto3" json:"addition_autograph,omitempty"`
}

func (m *ExpandedEndorseSignature) Restore()         { *m = ExpandedEndorseSignature{} }
func (m *ExpandedEndorseSignature) String() string { return proto.CompactTextString(m) }
func (*ExpandedEndorseSignature) SchemaSignal()    {}
func (*ExpandedEndorseSignature) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{9}
}
func (m *ExpandedEndorseSignature) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ExpandedEndorseSignature) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Expandedendorsesignal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedEndorseSignature) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Expandedendorsesignal.Merge(m, src)
}
func (m *ExpandedEndorseSignature) XXX_Volume() int {
	return m.Volume()
}
func (m *ExpandedEndorseSignature) XXX_Omitunclear() {
	xxx_messagedata_Expandedendorsesignal.DiscardUnknown(m)
}

var xxx_messagedata_Expandedendorsesignal proto.InternalMessageInfo

func (m *ExpandedEndorseSignature) FetchLedgerUidMark() LedgerUIDMark {
	if m != nil {
		return m.LedgerUidMark
	}
	return LedgerUIDMarkUnclear
}

func (m *ExpandedEndorseSignature) FetchRatifierLocation() []byte {
	if m != nil {
		return m.RatifierLocation
	}
	return nil
}

func (m *ExpandedEndorseSignature) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *ExpandedEndorseSignature) FetchAutograph() []byte {
	if m != nil {
		return m.Autograph
	}
	return nil
}

func (m *ExpandedEndorseSignature) FetchAddition() []byte {
	if m != nil {
		return m.Addition
	}
	return nil
}

func (m *ExpandedEndorseSignature) FetchAdditionAutograph() []byte {
	if m != nil {
		return m.AdditionAutograph
	}
	return nil
}

type Nomination struct {
	Kind      AttestedMessageKind `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Level    int64         `protobuf:"variableint,2,opt,name=height,proto3" json:"level,omitempty"`
	Cycle     int32         `protobuf:"variableint,3,opt,name=round,proto3" json:"epoch,omitempty"`
	PolEpoch  int32         `protobuf:"variableint,4,opt,name=pol_round,json=polRound,proto3" json:"pol_epoch,omitempty"`
	LedgerUID   LedgerUID       `protobuf:"octets,5,opt,name=block_id,json=blockId,proto3" json:"ledger_uid"`
	Timestamp time.Time     `protobuf:"octets,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Autograph []byte        `protobuf:"octets,7,opt,name=signature,proto3" json:"autograph,omitempty"`
}

func (m *Nomination) Restore()         { *m = Nomination{} }
func (m *Nomination) String() string { return proto.CompactTextString(m) }
func (*Nomination) SchemaSignal()    {}
func (*Nomination) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{10}
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

func (m *Nomination) FetchKind() AttestedMessageKind {
	if m != nil {
		return m.Kind
	}
	return UnclearKind
}

func (m *Nomination) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Nomination) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *Nomination) FetchPolEpoch() int32 {
	if m != nil {
		return m.PolEpoch
	}
	return 0
}

func (m *Nomination) FetchLedgerUID() LedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return LedgerUID{}
}

func (m *Nomination) FetchTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Nomination) FetchAutograph() []byte {
	if m != nil {
		return m.Autograph
	}
	return nil
}

type AttestedHeading struct {
	Heading *Heading `protobuf:"octets,1,opt,name=header,proto3" json:"heading,omitempty"`
	Endorse *Endorse `protobuf:"octets,2,opt,name=commit,proto3" json:"endorse,omitempty"`
}

func (m *AttestedHeading) Restore()         { *m = AttestedHeading{} }
func (m *AttestedHeading) String() string { return proto.CompactTextString(m) }
func (*AttestedHeading) SchemaSignal()    {}
func (*AttestedHeading) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{11}
}
func (m *AttestedHeading) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AttestedHeading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Attestedheading.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestedHeading) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Attestedheading.Merge(m, src)
}
func (m *AttestedHeading) XXX_Volume() int {
	return m.Volume()
}
func (m *AttestedHeading) XXX_Omitunclear() {
	xxx_messagedata_Attestedheading.DiscardUnknown(m)
}

var xxx_messagedata_Attestedheading proto.InternalMessageInfo

func (m *AttestedHeading) FetchHeading() *Heading {
	if m != nil {
		return m.Heading
	}
	return nil
}

func (m *AttestedHeading) FetchEndorse() *Endorse {
	if m != nil {
		return m.Endorse
	}
	return nil
}

type RapidLedger struct {
	AttestedHeading *AttestedHeading `protobuf:"octets,1,opt,name=signed_header,json=signedHeader,proto3" json:"attested_heading,omitempty"`
	RatifierAssign *RatifierAssign `protobuf:"octets,2,opt,name=validator_set,json=validatorSet,proto3" json:"ratifier_collection,omitempty"`
}

func (m *RapidLedger) Restore()         { *m = RapidLedger{} }
func (m *RapidLedger) String() string { return proto.CompactTextString(m) }
func (*RapidLedger) SchemaSignal()    {}
func (*RapidLedger) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{12}
}
func (m *RapidLedger) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *RapidLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Rapidledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RapidLedger) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Rapidledger.Merge(m, src)
}
func (m *RapidLedger) XXX_Volume() int {
	return m.Volume()
}
func (m *RapidLedger) XXX_Omitunclear() {
	xxx_messagedata_Rapidledger.DiscardUnknown(m)
}

var xxx_messagedata_Rapidledger proto.InternalMessageInfo

func (m *RapidLedger) FetchAttestedHeading() *AttestedHeading {
	if m != nil {
		return m.AttestedHeading
	}
	return nil
}

func (m *RapidLedger) FetchRatifierCollection() *RatifierAssign {
	if m != nil {
		return m.RatifierAssign
	}
	return nil
}

type LedgerMeta struct {
	LedgerUID   LedgerUID `protobuf:"octets,1,opt,name=block_id,json=blockId,proto3" json:"ledger_uid"`
	LedgerVolume int64   `protobuf:"variableint,2,opt,name=block_size,json=blockSize,proto3" json:"ledger_volume,omitempty"`
	Heading    Heading  `protobuf:"octets,3,opt,name=header,proto3" json:"heading"`
	CountTrans    int64   `protobuf:"variableint,4,opt,name=num_txs,json=numTxs,proto3" json:"count_trans,omitempty"`
}

func (m *LedgerMeta) Restore()         { *m = LedgerMeta{} }
func (m *LedgerMeta) String() string { return proto.CompactTextString(m) }
func (*LedgerMeta) SchemaSignal()    {}
func (*LedgerMeta) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{13}
}
func (m *LedgerMeta) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *LedgerMeta) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ledgermeta.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerMeta) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ledgermeta.Merge(m, src)
}
func (m *LedgerMeta) XXX_Volume() int {
	return m.Volume()
}
func (m *LedgerMeta) XXX_Omitunclear() {
	xxx_messagedata_Ledgermeta.DiscardUnknown(m)
}

var xxx_messagedata_Ledgermeta proto.InternalMessageInfo

func (m *LedgerMeta) FetchLedgerUID() LedgerUID {
	if m != nil {
		return m.LedgerUID
	}
	return LedgerUID{}
}

func (m *LedgerMeta) FetchLedgerVolume() int64 {
	if m != nil {
		return m.LedgerVolume
	}
	return 0
}

func (m *LedgerMeta) FetchHeading() Heading {
	if m != nil {
		return m.Heading
	}
	return Heading{}
}

func (m *LedgerMeta) FetchCountTrans() int64 {
	if m != nil {
		return m.CountTrans
	}
	return 0
}

//
type TransferEvidence struct {
	OriginDigest []byte        `protobuf:"octets,1,opt,name=root_hash,json=rootHash,proto3" json:"origin_digest,omitempty"`
	Data     []byte        `protobuf:"octets,2,opt,name=data,proto3" json:"data,omitempty"`
	Attestation    *vault.Attestation `protobuf:"octets,3,opt,name=proof,proto3" json:"evidence,omitempty"`
}

func (m *TransferEvidence) Restore()         { *m = TransferEvidence{} }
func (m *TransferEvidence) String() string { return proto.CompactTextString(m) }
func (*TransferEvidence) SchemaSignal()    {}
func (*TransferEvidence) Definition() ([]byte, []int) {
	return filedefinition_d3a6e55e2345de56, []int{14}
}
func (m *TransferEvidence) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *TransferEvidence) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Transproof.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferEvidence) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Transproof.Merge(m, src)
}
func (m *TransferEvidence) XXX_Volume() int {
	return m.Volume()
}
func (m *TransferEvidence) XXX_Omitunclear() {
	xxx_messagedata_Transproof.DiscardUnknown(m)
}

var xxx_messagedata_Transproof proto.InternalMessageInfo

func (m *TransferEvidence) FetchRootDigest() []byte {
	if m != nil {
		return m.OriginDigest
	}
	return nil
}

func (m *TransferEvidence) FetchData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TransferEvidence) FetchEvidence() *vault.Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

func init() {
	proto.RegisterEnum("REDACTED", Attestedsignaltype_label, Attestedsignaltype_item)
	proto.RegisterType((*SegmentAssignHeading)(nil), "REDACTED")
	proto.RegisterType((*Segment)(nil), "REDACTED")
	proto.RegisterType((*LedgerUID)(nil), "REDACTED")
	proto.RegisterType((*Heading)(nil), "REDACTED")
	proto.RegisterType((*Data)(nil), "REDACTED")
	proto.RegisterType((*Ballot)(nil), "REDACTED")
	proto.RegisterType((*Endorse)(nil), "REDACTED")
	proto.RegisterType((*EndorseSignature)(nil), "REDACTED")
	proto.RegisterType((*ExpandedEndorse)(nil), "REDACTED")
	proto.RegisterType((*ExpandedEndorseSignature)(nil), "REDACTED")
	proto.RegisterType((*Nomination)(nil), "REDACTED")
	proto.RegisterType((*AttestedHeading)(nil), "REDACTED")
	proto.RegisterType((*RapidLedger)(nil), "REDACTED")
	proto.RegisterType((*LedgerMeta)(nil), "REDACTED")
	proto.RegisterType((*TransferEvidence)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_d3a6e55e2345de56) }

var filedefinition_d3a6e55e2345de56 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xce, 0xda, 0xeb, 0x5f, 0xcf, 0x76, 0xe2, 0x2c, 0x11, 0x75, 0xdd, 0xc6, 0xb1, 0x5c, 0x01,
	0xa1, 0xa0, 0x4d, 0x95, 0x22, 0x04, 0x07, 0x0e, 0xf9, 0x45, 0x1b, 0x51, 0x27, 0xd6, 0xda, 0x2d,
	0xa2, 0x97, 0xd5, 0xda, 0x3b, 0xb1, 0x97, 0xda, 0x3b, 0xab, 0xdd, 0x71, 0x70, 0xfa, 0x17, 0xa0,
	0x9e, 0x7a, 0xe2, 0xd6, 0x13, 0x1c, 0xb8, 0x83, 0xc4, 0x15, 0x71, 0xea, 0xb1, 0x37, 0xb8, 0x50,
	0x20, 0x95, 0xf8, 0x3b, 0xd0, 0xbc, 0x99, 0xdd, 0xb5, 0xe3, 0x18, 0xaa, 0xa8, 0x02, 0x89, 0x8b,
	0xb5, 0xf3, 0xde, 0xf7, 0xde, 0xbc, 0x79, 0xdf, 0x37, 0xa3, 0x67, 0xb8, 0xca, 0x88, 0x6b, 0x13,
	0x7f, 0xe8, 0xb8, 0x6c, 0x83, 0x9d, 0x78, 0x24, 0x10, 0xbf, 0xba, 0xe7, 0x53, 0x46, 0xb5, 0x52,
	0xec, 0xd5, 0xd1, 0x5e, 0x59, 0xe9, 0xd1, 0x1e, 0x45, 0xe7, 0x06, 0xff, 0x12, 0xb8, 0xca, 0x5a,
	0x8f, 0xd2, 0xde, 0x80, 0x6c, 0xe0, 0xaa, 0x33, 0x3a, 0xda, 0x60, 0xce, 0x90, 0x04, 0xcc, 0x1a,
	0x7a, 0x12, 0xb0, 0x3a, 0xb1, 0x4d, 0xd7, 0x3f, 0xf1, 0x18, 0xe5, 0x58, 0x7a, 0x24, 0xdd, 0xd5,
	0x09, 0xf7, 0x31, 0xf1, 0x03, 0x87, 0xba, 0x93, 0x75, 0x54, 0x6a, 0x33, 0x55, 0x1e, 0x5b, 0x03,
	0xc7, 0xb6, 0x18, 0xf5, 0x05, 0xa2, 0xfe, 0x21, 0x14, 0x9b, 0x96, 0xcf, 0x5a, 0x84, 0xdd, 0x26,
	0x96, 0x4d, 0x7c, 0x6d, 0x05, 0x52, 0x8c, 0x32, 0x6b, 0x50, 0x56, 0x6a, 0xca, 0x7a, 0xd1, 0x10,
	0x0b, 0x4d, 0x03, 0xb5, 0x6f, 0x05, 0xfd, 0x72, 0xa2, 0xa6, 0xac, 0x17, 0x0c, 0xfc, 0xae, 0xf7,
	0x41, 0xe5, 0xa1, 0x3c, 0xc2, 0x71, 0x6d, 0x32, 0x0e, 0x23, 0x70, 0xc1, 0xad, 0x9d, 0x13, 0x46,
	0x02, 0x19, 0x22, 0x16, 0xda, 0x7b, 0x90, 0xc2, 0xfa, 0xcb, 0xc9, 0x9a, 0xb2, 0x9e, 0xdf, 0x2c,
	0xeb, 0x13, 0x8d, 0x12, 0xe7, 0xd3, 0x9b, 0xdc, 0xbf, 0xad, 0x3e, 0x7d, 0xbe, 0xb6, 0x60, 0x08,
	0x70, 0x7d, 0x00, 0x99, 0xed, 0x01, 0xed, 0x3e, 0xd8, 0xdf, 0x8d, 0x0a, 0x51, 0xe2, 0x42, 0xb4,
	0x06, 0x2c, 0x79, 0x96, 0xcf, 0xcc, 0x80, 0x30, 0xb3, 0x8f, 0xa7, 0xc0, 0x4d, 0xf3, 0x9b, 0x6b,
	0xfa, 0x59, 0x1e, 0xf4, 0xa9, 0xc3, 0xca, 0x5d, 0x8a, 0xde, 0xa4, 0xb1, 0xfe, 0xa7, 0x0a, 0x69,
	0xd9, 0x8c, 0x8f, 0x20, 0x23, 0xdb, 0x8a, 0x1b, 0xe6, 0x37, 0x57, 0x27, 0x33, 0x4a, 0x97, 0xbe,
	0x43, 0xdd, 0x80, 0xb8, 0xc1, 0x28, 0x90, 0xf9, 0xc2, 0x18, 0xed, 0x4d, 0xc8, 0x76, 0xfb, 0x96,
	0xe3, 0x9a, 0x8e, 0x8d, 0x15, 0xe5, 0xb6, 0xf3, 0xa7, 0xcf, 0xd7, 0x32, 0x3b, 0xdc, 0xb6, 0xbf,
	0x6b, 0x64, 0xd0, 0xb9, 0x6f, 0x6b, 0xaf, 0x43, 0xba, 0x4f, 0x9c, 0x5e, 0x9f, 0x61, 0x5b, 0x92,
	0x86, 0x5c, 0x69, 0x1f, 0x80, 0xca, 0x05, 0x51, 0x56, 0x71, 0xef, 0x8a, 0x2e, 0xd4, 0xa2, 0x87,
	0x6a, 0xd1, 0xdb, 0xa1, 0x5a, 0xb6, 0xb3, 0x7c, 0xe3, 0xc7, 0xbf, 0xad, 0x29, 0x06, 0x46, 0x68,
	0x3b, 0x50, 0x1c, 0x58, 0x01, 0x33, 0x3b, 0xbc, 0x6d, 0x7c, 0xfb, 0x14, 0xa6, 0xb8, 0x3c, 0xdb,
	0x10, 0xd9, 0x58, 0x59, 0x7a, 0x9e, 0x47, 0x09, 0x93, 0xad, 0xad, 0x43, 0x09, 0x93, 0x74, 0xe9,
	0x70, 0xe8, 0x30, 0x13, 0xfb, 0x9e, 0xc6, 0xbe, 0x2f, 0x72, 0xfb, 0x0e, 0x9a, 0x6f, 0x73, 0x06,
	0xae, 0x40, 0xce, 0xb6, 0x98, 0x25, 0x20, 0x19, 0x84, 0x64, 0xb9, 0x01, 0x9d, 0x6f, 0xc1, 0x52,
	0xa4, 0xba, 0x40, 0x40, 0xb2, 0x22, 0x4b, 0x6c, 0x46, 0xe0, 0x0d, 0x58, 0x71, 0xc9, 0x98, 0x99,
	0x67, 0xd1, 0x39, 0x44, 0x6b, 0xdc, 0x77, 0x6f, 0x3a, 0xe2, 0x0d, 0x58, 0xec, 0x86, 0xcd, 0x17,
	0x58, 0x40, 0x6c, 0x31, 0xb2, 0x22, 0xec, 0x32, 0x64, 0x2d, 0xcf, 0x13, 0x80, 0x3c, 0x02, 0x32,
	0x96, 0xe7, 0xa1, 0xeb, 0x3a, 0x2c, 0xe3, 0x19, 0x7d, 0x12, 0x8c, 0x06, 0x4c, 0x26, 0x29, 0x20,
	0x66, 0x89, 0x3b, 0x0c, 0x61, 0x47, 0xec, 0x35, 0x28, 0x92, 0x63, 0xc7, 0x26, 0x6e, 0x97, 0x08,
	0x5c, 0x11, 0x71, 0x85, 0xd0, 0x88, 0xa0, 0xb7, 0xa1, 0xe4, 0xf9, 0xd4, 0xa3, 0x01, 0xf1, 0x4d,
	0xcb, 0xb6, 0x7d, 0x12, 0x04, 0xe5, 0x45, 0x91, 0x2f, 0xb4, 0x6f, 0x09, 0x73, 0xbd, 0x0c, 0xea,
	0xae, 0xc5, 0x2c, 0xad, 0x04, 0x49, 0x36, 0x0e, 0xca, 0x4a, 0x2d, 0xb9, 0x5e, 0x30, 0xf8, 0x67,
	0xfd, 0x87, 0x24, 0xa8, 0xf7, 0x28, 0x23, 0xda, 0x4d, 0x50, 0x39, 0x4d, 0xa8, 0xbe, 0xc5, 0xf3,
	0xf4, 0xdc, 0x72, 0x7a, 0x2e, 0xb1, 0x1b, 0x41, 0xaf, 0x7d, 0xe2, 0x11, 0x03, 0xc1, 0x13, 0x72,
	0x4a, 0x4c, 0xc9, 0x69, 0x05, 0x52, 0x3e, 0x1d, 0xb9, 0x36, 0xaa, 0x2c, 0x65, 0x88, 0x85, 0xb6,
	0x07, 0xd9, 0x48, 0x25, 0xea, 0x3f, 0xa9, 0x64, 0x89, 0xab, 0x84, 0x6b, 0x58, 0x1a, 0x8c, 0x4c,
	0x47, 0x8a, 0x65, 0x1b, 0x72, 0xd1, 0xe3, 0x25, 0xd5, 0xf6, 0x72, 0x82, 0x8d, 0xc3, 0xb4, 0x77,
	0x60, 0x39, 0xe2, 0x3e, 0x6a, 0x9e, 0x50, 0x5c, 0x29, 0x72, 0xc8, 0xee, 0x4d, 0xc9, 0xca, 0x14,
	0x0f, 0x50, 0x06, 0xcf, 0x15, 0xcb, 0x6a, 0x1f, 0x5f, 0xa2, 0xab, 0x90, 0x0b, 0x9c, 0x9e, 0x6b,
	0xb1, 0x91, 0x4f, 0xa4, 0xf2, 0x62, 0x03, 0xf7, 0x92, 0x31, 0x23, 0x2e, 0x5e, 0x72, 0xa1, 0xb4,
	0xd8, 0xa0, 0x6d, 0xc0, 0x6b, 0xd1, 0xc2, 0x8c, 0xb3, 0x08, 0x95, 0x69, 0x91, 0xab, 0x15, 0x7a,
	0xea, 0x3f, 0x2a, 0x90, 0x16, 0x17, 0x63, 0x82, 0x06, 0xe5, 0x7c, 0x1a, 0x12, 0xf3, 0x68, 0x48,
	0x5e, 0x9c, 0x86, 0x2d, 0x80, 0xa8, 0xcc, 0xa0, 0xac, 0xd6, 0x92, 0xeb, 0xf9, 0xcd, 0x2b, 0xb3,
	0x89, 0x44, 0x89, 0x2d, 0xa7, 0x27, 0xef, 0xfd, 0x44, 0x50, 0xfd, 0x57, 0x05, 0x72, 0x91, 0x5f,
	0xdb, 0x82, 0x62, 0x58, 0x97, 0x79, 0x34, 0xb0, 0x7a, 0x52, 0x8a, 0xab, 0x73, 0x8b, 0xfb, 0x78,
	0x60, 0xf5, 0x8c, 0xbc, 0xac, 0x87, 0x2f, 0xce, 0xa7, 0x35, 0x31, 0x87, 0xd6, 0x29, 0x1d, 0x25,
	0x2f, 0xa6, 0xa3, 0x29, 0xc6, 0xd5, 0x33, 0x8c, 0xd7, 0xff, 0x50, 0x60, 0x71, 0x6f, 0x8c, 0xe5,
	0xdb, 0xff, 0x25, 0x55, 0xf7, 0xa5, 0xb6, 0x6c, 0x62, 0x9b, 0x33, 0x9c, 0x5d, 0x9b, 0xcd, 0x38,
	0x5d, 0x73, 0xcc, 0x9d, 0x16, 0x66, 0x69, 0xc5, 0x1c, 0x7e, 0x9f, 0x80, 0xe5, 0x19, 0xfc, 0xff,
	0x8f, 0xcb, 0xe9, 0xdb, 0x9b, 0x7a, 0xc9, 0xdb, 0x9b, 0x9e, 0x7b, 0x7b, 0xbf, 0x4b, 0x40, 0xb6,
	0x89, 0xaf, 0xb4, 0x35, 0xf8, 0x37, 0xde, 0xde, 0x2b, 0x90, 0xf3, 0xe8, 0xc0, 0x14, 0x1e, 0x15,
	0x3d, 0x59, 0x8f, 0x0e, 0x8c, 0x19, 0x99, 0xa5, 0x5e, 0xd1, 0xc3, 0x9c, 0x7e, 0x05, 0x24, 0x64,
	0xce, 0x5e, 0x28, 0x1f, 0x0a, 0xa2, 0x15, 0x72, 0x6a, 0xba, 0xc1, 0x7b, 0x80, 0x63, 0x98, 0x32,
	0x3b, 0xe5, 0x89, 0xb2, 0x05, 0xd2, 0x90, 0x38, 0x1e, 0x21, 0x86, 0x0c, 0x39, 0xb8, 0x95, 0xe7,
	0xbd, 0x58, 0x86, 0xc4, 0xd5, 0xbf, 0x52, 0x00, 0xee, 0xf0, 0xce, 0xe2, 0x79, 0xf9, 0xbc, 0x13,
	0x60, 0x09, 0xe6, 0xd4, 0xce, 0xd5, 0x79, 0xa4, 0xc9, 0xfd, 0x0b, 0xc1, 0x64, 0xdd, 0x3b, 0x50,
	0x8c, 0xb5, 0x1d, 0x90, 0xb0, 0x98, 0x73, 0x92, 0x44, 0x63, 0x48, 0x8b, 0x30, 0xa3, 0x70, 0x3c,
	0xb1, 0xaa, 0xff, 0xa4, 0x40, 0x0e, 0x6b, 0x6a, 0x10, 0x66, 0x4d, 0x71, 0xa8, 0x5c, 0x9c, 0xc3,
	0x55, 0x00, 0x91, 0x26, 0x70, 0x1e, 0x12, 0xa9, 0xac, 0x1c, 0x5a, 0x5a, 0xce, 0x43, 0xa2, 0xbd,
	0x1f, 0x35, 0x3c, 0xf9, 0xf7, 0x0d, 0x97, 0x2f, 0x46, 0xd8, 0xf6, 0x4b, 0x90, 0x71, 0x47, 0x43,
	0x93, 0x0f, 0x1f, 0xaa, 0x50, 0xab, 0x3b, 0x1a, 0xb6, 0xc7, 0x41, 0xfd, 0x73, 0xc8, 0xb4, 0xc7,
	0x38, 0x88, 0x73, 0x89, 0xfa, 0x94, 0xca, 0xe9, 0x4f, 0x4c, 0xdd, 0x59, 0x6e, 0xc0, 0x61, 0x47,
	0x03, 0x95, 0x8f, 0x79, 0xe1, 0xdf, 0x02, 0xfe, 0xad, 0xe9, 0x2f, 0x39, 0xe2, 0xcb, 0xe1, 0xfe,
	0xfa, 0xcf, 0x0a, 0x14, 0xa7, 0x6e, 0x92, 0xf6, 0x2e, 0x5c, 0x6a, 0xed, 0xdf, 0x3a, 0xd8, 0xdb,
	0x35, 0x1b, 0xad, 0x5b, 0x66, 0xfb, 0xb3, 0xe6, 0x9e, 0x79, 0xf7, 0xe0, 0x93, 0x83, 0xc3, 0x4f,
	0x0f, 0x4a, 0x0b, 0x95, 0xa5, 0x47, 0x4f, 0x6a, 0xf9, 0xbb, 0xee, 0x03, 0x97, 0x7e, 0xe1, 0xce,
	0x43, 0x37, 0x8d, 0xbd, 0x7b, 0x87, 0xed, 0xbd, 0x92, 0x22, 0xd0, 0x4d, 0x9f, 0x1c, 0x53, 0x46,
	0x10, 0x7d, 0x03, 0x2e, 0x9f, 0x83, 0xde, 0x39, 0x6c, 0x34, 0xf6, 0xdb, 0xa5, 0x44, 0x65, 0xf9,
	0xd1, 0x93, 0x5a, 0xb1, 0xe9, 0x13, 0xa1, 0x32, 0x8c, 0xd0, 0xa1, 0x3c, 0x1b, 0x71, 0xd8, 0x3c,
	0x6c, 0x6d, 0xdd, 0x29, 0xd5, 0x2a, 0xa5, 0x47, 0x4f, 0x6a, 0x85, 0xf0, 0xc9, 0xe0, 0xf8, 0x4a,
	0xf6, 0xcb, 0xaf, 0xab, 0x0b, 0xdf, 0x7e, 0x53, 0x55, 0xb6, 0x1b, 0x4f, 0x4f, 0xab, 0xca, 0xb3,
	0xd3, 0xaa, 0xf2, 0xfb, 0x69, 0x55, 0x79, 0xfc, 0xa2, 0xba, 0xf0, 0xec, 0x45, 0x75, 0xe1, 0x97,
	0x17, 0xd5, 0x85, 0xfb, 0x37, 0x7b, 0x0e, 0xeb, 0x8f, 0x3a, 0x7a, 0x97, 0x0e, 0x37, 0xba, 0x74,
	0x48, 0x58, 0xe7, 0x88, 0xc5, 0x1f, 0xe2, 0x6f, 0xe2, 0xd9, 0xbf, 0x6e, 0x9d, 0x34, 0xda, 0x6f,
	0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xb6, 0xa1, 0x4e, 0x7b, 0x0e, 0x00, 0x00,
}

func (m *SegmentAssignHeading) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SegmentAssignHeading) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *SegmentAssignHeading) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sum != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Sum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Segment) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Segment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Segment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.Attestation.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Octets) > 0 {
		i -= len(m.Octets)
		copy(dAtA[i:], m.Octets)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Octets)))
		i--
		dAtA[i] = 0x12
	}
	if m.Ordinal != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ordinal))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LedgerUID) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LedgerUID) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *LedgerUID) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Heading) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Heading) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Heading) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecommenderLocation) > 0 {
		i -= len(m.RecommenderLocation)
		copy(dAtA[i:], m.RecommenderLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RecommenderLocation)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.ProofDigest) > 0 {
		i -= len(m.ProofDigest)
		copy(dAtA[i:], m.ProofDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ProofDigest)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.FinalOutcomesDigest) > 0 {
		i -= len(m.FinalOutcomesDigest)
		copy(dAtA[i:], m.FinalOutcomesDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FinalOutcomesDigest)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ApplicationDigest) > 0 {
		i -= len(m.ApplicationDigest)
		copy(dAtA[i:], m.ApplicationDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ApplicationDigest)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.AgreementDigest) > 0 {
		i -= len(m.AgreementDigest)
		copy(dAtA[i:], m.AgreementDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.AgreementDigest)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.FollowingRatifiersDigest) > 0 {
		i -= len(m.FollowingRatifiersDigest)
		copy(dAtA[i:], m.FollowingRatifiersDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FollowingRatifiersDigest)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.RatifiersDigest) > 0 {
		i -= len(m.RatifiersDigest)
		copy(dAtA[i:], m.RatifiersDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RatifiersDigest)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DataDigest) > 0 {
		i -= len(m.DataDigest)
		copy(dAtA[i:], m.DataDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.DataDigest)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FinalEndorseDigest) > 0 {
		i -= len(m.FinalEndorseDigest)
		copy(dAtA[i:], m.FinalEndorseDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FinalEndorseDigest)))
		i--
		dAtA[i] = 0x32
	}
	{
		volume, err := m.FinalLedgerUid.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x2a
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = formatVariableintKinds(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x22
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if len(m.LedgerUID) > 0 {
		i -= len(m.LedgerUID)
		copy(dAtA[i:], m.LedgerUID)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.LedgerUID)))
		i--
		dAtA[i] = 0x12
	}
	{
		volume, err := m.Release.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *Data) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Data) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Data) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for idxNdEx := len(m.Txs) - 1; idxNdEx >= 0; idxNdEx-- {
			i -= len(m.Txs[idxNdEx])
			copy(dAtA[i:], m.Txs[idxNdEx])
			i = formatVariableintKinds(dAtA, i, uint64(len(m.Txs[idxNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.AdditionAutograph) > 0 {
		i -= len(m.AdditionAutograph)
		copy(dAtA[i:], m.AdditionAutograph)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.AdditionAutograph)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Addition) > 0 {
		i -= len(m.Addition)
		copy(dAtA[i:], m.Addition)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Addition)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Autograph) > 0 {
		i -= len(m.Autograph)
		copy(dAtA[i:], m.Autograph)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Autograph)))
		i--
		dAtA[i] = 0x42
	}
	if m.RatifierOrdinal != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.RatifierOrdinal))
		i--
		dAtA[i] = 0x38
	}
	if len(m.RatifierLocation) > 0 {
		i -= len(m.RatifierLocation)
		copy(dAtA[i:], m.RatifierLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RatifierLocation)))
		i--
		dAtA[i] = 0x32
	}
	n6, err6 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = formatVariableintKinds(dAtA, i, uint64(n6))
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
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x18
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	if m.Kind != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Endorse) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endorse) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Endorse) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Endorsements) > 0 {
		for idxNdEx := len(m.Endorsements) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Endorsements[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *EndorseSignature) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndorseSignature) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *EndorseSignature) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Autograph) > 0 {
		i -= len(m.Autograph)
		copy(dAtA[i:], m.Autograph)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Autograph)))
		i--
		dAtA[i] = 0x22
	}
	n9, err9 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err9 != nil {
		return 0, err9
	}
	i -= n9
	i = formatVariableintKinds(dAtA, i, uint64(n9))
	i--
	dAtA[i] = 0x1a
	if len(m.RatifierLocation) > 0 {
		i -= len(m.RatifierLocation)
		copy(dAtA[i:], m.RatifierLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RatifierLocation)))
		i--
		dAtA[i] = 0x12
	}
	if m.LedgerUidMark != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.LedgerUidMark))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExpandedEndorse) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpandedEndorse) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ExpandedEndorse) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExpandedEndorsements) > 0 {
		for idxNdEx := len(m.ExpandedEndorsements) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.ExpandedEndorsements[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *ExpandedEndorseSignature) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpandedEndorseSignature) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ExpandedEndorseSignature) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdditionAutograph) > 0 {
		i -= len(m.AdditionAutograph)
		copy(dAtA[i:], m.AdditionAutograph)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.AdditionAutograph)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Addition) > 0 {
		i -= len(m.Addition)
		copy(dAtA[i:], m.Addition)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Addition)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Autograph) > 0 {
		i -= len(m.Autograph)
		copy(dAtA[i:], m.Autograph)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Autograph)))
		i--
		dAtA[i] = 0x22
	}
	n11, err11 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err11 != nil {
		return 0, err11
	}
	i -= n11
	i = formatVariableintKinds(dAtA, i, uint64(n11))
	i--
	dAtA[i] = 0x1a
	if len(m.RatifierLocation) > 0 {
		i -= len(m.RatifierLocation)
		copy(dAtA[i:], m.RatifierLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RatifierLocation)))
		i--
		dAtA[i] = 0x12
	}
	if m.LedgerUidMark != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.LedgerUidMark))
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
	if len(m.Autograph) > 0 {
		i -= len(m.Autograph)
		copy(dAtA[i:], m.Autograph)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Autograph)))
		i--
		dAtA[i] = 0x3a
	}
	n12, err12 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err12 != nil {
		return 0, err12
	}
	i -= n12
	i = formatVariableintKinds(dAtA, i, uint64(n12))
	i--
	dAtA[i] = 0x32
	{
		volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x2a
	if m.PolEpoch != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.PolEpoch))
		i--
		dAtA[i] = 0x20
	}
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x18
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	if m.Kind != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AttestedHeading) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestedHeading) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AttestedHeading) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Endorse != nil {
		{
			volume, err := m.Endorse.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Heading != nil {
		{
			volume, err := m.Heading.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *RapidLedger) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RapidLedger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *RapidLedger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RatifierAssign != nil {
		{
			volume, err := m.RatifierAssign.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.AttestedHeading != nil {
		{
			volume, err := m.AttestedHeading.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *LedgerMeta) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LedgerMeta) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *LedgerMeta) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CountTrans != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.CountTrans))
		i--
		dAtA[i] = 0x20
	}
	{
		volume, err := m.Heading.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x1a
	if m.LedgerVolume != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.LedgerVolume))
		i--
		dAtA[i] = 0x10
	}
	{
		volume, err := m.LedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *TransferEvidence) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferEvidence) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *TransferEvidence) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OriginDigest) > 0 {
		i -= len(m.OriginDigest)
		copy(dAtA[i:], m.OriginDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.OriginDigest)))
		i--
		dAtA[i] = 0xa
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
func (m *SegmentAssignHeading) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != 0 {
		n += 1 + sovKinds(uint64(m.Sum))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Segment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	l = len(m.Octets)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = m.Attestation.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *LedgerUID) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = m.SegmentAssignHeading.Volume()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *Heading) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Release.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.LedgerUID)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovKinds(uint64(l))
	l = m.FinalLedgerUid.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FinalEndorseDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.DataDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RatifiersDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.FollowingRatifiersDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AgreementDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.ApplicationDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.FinalOutcomesDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.ProofDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RecommenderLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Data) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *Ballot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	l = m.LedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.RatifierLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.RatifierOrdinal != 0 {
		n += 1 + sovKinds(uint64(m.RatifierOrdinal))
	}
	l = len(m.Autograph)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Addition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AdditionAutograph)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Endorse) Volume() (n int) {
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
	l = m.LedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Endorsements) > 0 {
		for _, e := range m.Endorsements {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *EndorseSignature) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerUidMark != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUidMark))
	}
	l = len(m.RatifierLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Autograph)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ExpandedEndorse) Volume() (n int) {
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
	l = m.LedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.ExpandedEndorsements) > 0 {
		for _, e := range m.ExpandedEndorsements {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ExpandedEndorseSignature) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerUidMark != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUidMark))
	}
	l = len(m.RatifierLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Autograph)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Addition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AdditionAutograph)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Nomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	if m.PolEpoch != 0 {
		n += 1 + sovKinds(uint64(m.PolEpoch))
	}
	l = m.LedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Autograph)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AttestedHeading) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Heading != nil {
		l = m.Heading.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Endorse != nil {
		l = m.Endorse.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *RapidLedger) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestedHeading != nil {
		l = m.AttestedHeading.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.RatifierAssign != nil {
		l = m.RatifierAssign.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *LedgerMeta) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.LedgerVolume != 0 {
		n += 1 + sovKinds(uint64(m.LedgerVolume))
	}
	l = m.Heading.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.CountTrans != 0 {
		n += 1 + sovKinds(uint64(m.CountTrans))
	}
	return n
}

func (m *TransferEvidence) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Attestation != nil {
		l = m.Attestation.Volume()
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
func (m *SegmentAssignHeading) Unserialize(dAtA []byte) error {
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
			m.Sum = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
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
func (m *Segment) Unserialize(dAtA []byte) error {
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
				m.Ordinal |= uint32(b&0x7F) << displace
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Octets = append(m.Octets[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Octets == nil {
				m.Octets = []byte{}
			}
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
			if err := m.Attestation.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *LedgerUID) Unserialize(dAtA []byte) error {
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
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
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
			if err := m.SegmentAssignHeading.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *Heading) Unserialize(dAtA []byte) error {
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
			if err := m.Release.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.LedgerUID = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 3:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if err := m.FinalLedgerUid.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalEndorseDigest = append(m.FinalEndorseDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FinalEndorseDigest == nil {
				m.FinalEndorseDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.DataDigest = append(m.DataDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.DataDigest == nil {
				m.DataDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.RatifiersDigest = append(m.RatifiersDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.RatifiersDigest == nil {
				m.RatifiersDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 9:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FollowingRatifiersDigest = append(m.FollowingRatifiersDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FollowingRatifiersDigest == nil {
				m.FollowingRatifiersDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 10:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AgreementDigest = append(m.AgreementDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.AgreementDigest == nil {
				m.AgreementDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 11:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationDigest = append(m.ApplicationDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.ApplicationDigest == nil {
				m.ApplicationDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 12:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalOutcomesDigest = append(m.FinalOutcomesDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FinalOutcomesDigest == nil {
				m.FinalOutcomesDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 13:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofDigest = append(m.ProofDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.ProofDigest == nil {
				m.ProofDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 14:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.RecommenderLocation = append(m.RecommenderLocation[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.RecommenderLocation == nil {
				m.RecommenderLocation = []byte{}
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
func (m *Data) Unserialize(dAtA []byte) error {
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
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, make([]byte, submitOrdinal-idxNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[idxNdEx:submitOrdinal])
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
				m.Kind |= AttestedMessageKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.RatifierLocation = append(m.RatifierLocation[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.RatifierLocation == nil {
				m.RatifierLocation = []byte{}
			}
			idxNdEx = submitOrdinal
		case 7:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.RatifierOrdinal = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.RatifierOrdinal |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Autograph = append(m.Autograph[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Autograph == nil {
				m.Autograph = []byte{}
			}
			idxNdEx = submitOrdinal
		case 9:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Addition = append(m.Addition[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Addition == nil {
				m.Addition = []byte{}
			}
			idxNdEx = submitOrdinal
		case 10:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionAutograph = append(m.AdditionAutograph[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.AdditionAutograph == nil {
				m.AdditionAutograph = []byte{}
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
func (m *Endorse) Unserialize(dAtA []byte) error {
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
			if err := m.LedgerUID.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Endorsements = append(m.Endorsements, EndorseSignature{})
			if err := m.Endorsements[len(m.Endorsements)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *EndorseSignature) Unserialize(dAtA []byte) error {
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
			m.LedgerUidMark = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.LedgerUidMark |= LedgerUIDMark(b&0x7F) << displace
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.RatifierLocation = append(m.RatifierLocation[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.RatifierLocation == nil {
				m.RatifierLocation = []byte{}
			}
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Autograph = append(m.Autograph[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Autograph == nil {
				m.Autograph = []byte{}
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
func (m *ExpandedEndorse) Unserialize(dAtA []byte) error {
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
			if err := m.LedgerUID.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.ExpandedEndorsements = append(m.ExpandedEndorsements, ExpandedEndorseSignature{})
			if err := m.ExpandedEndorsements[len(m.ExpandedEndorsements)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *ExpandedEndorseSignature) Unserialize(dAtA []byte) error {
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
			m.LedgerUidMark = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.LedgerUidMark |= LedgerUIDMark(b&0x7F) << displace
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.RatifierLocation = append(m.RatifierLocation[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.RatifierLocation == nil {
				m.RatifierLocation = []byte{}
			}
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Autograph = append(m.Autograph[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Autograph == nil {
				m.Autograph = []byte{}
			}
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Addition = append(m.Addition[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Addition == nil {
				m.Addition = []byte{}
			}
			idxNdEx = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionAutograph = append(m.AdditionAutograph[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.AdditionAutograph == nil {
				m.AdditionAutograph = []byte{}
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
				m.Kind |= AttestedMessageKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
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
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PolEpoch = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.PolEpoch |= int32(b&0x7F) << displace
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Autograph = append(m.Autograph[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Autograph == nil {
				m.Autograph = []byte{}
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
func (m *AttestedHeading) Unserialize(dAtA []byte) error {
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
			if m.Heading == nil {
				m.Heading = &Heading{}
			}
			if err := m.Heading.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.Endorse == nil {
				m.Endorse = &Endorse{}
			}
			if err := m.Endorse.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *RapidLedger) Unserialize(dAtA []byte) error {
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
			if m.AttestedHeading == nil {
				m.AttestedHeading = &AttestedHeading{}
			}
			if err := m.AttestedHeading.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.RatifierAssign == nil {
				m.RatifierAssign = &RatifierAssign{}
			}
			if err := m.RatifierAssign.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *LedgerMeta) Unserialize(dAtA []byte) error {
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
			if err := m.LedgerUID.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.LedgerVolume = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.LedgerVolume |= int64(b&0x7F) << displace
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
			if err := m.Heading.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.CountTrans = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.CountTrans |= int64(b&0x7F) << displace
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
func (m *TransferEvidence) Unserialize(dAtA []byte) error {
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
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginDigest = append(m.OriginDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.OriginDigest == nil {
				m.OriginDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
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
			if m.Attestation == nil {
				m.Attestation = &vault.Attestation{}
			}
			if err := m.Attestation.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
