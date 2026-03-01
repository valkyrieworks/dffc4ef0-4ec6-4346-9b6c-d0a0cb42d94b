//
//

package kinds

import (
	fmt "fmt"
	security "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	edition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
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
type AttestedSignalKind int32

const (
	UnfamiliarKind AttestedSignalKind = 0
	//
	PreballotKind   AttestedSignalKind = 1
	PreendorseKind AttestedSignalKind = 2
	//
	NominationKind AttestedSignalKind = 32
)

var Notatedsignalkind_alias = map[int32]string{
	0:  "REDACTED",
	1:  "REDACTED",
	2:  "REDACTED",
	32: "REDACTED",
}

var Notatedsignalkind_datum = map[string]int32{
	"REDACTED":   0,
	"REDACTED":   1,
	"REDACTED": 2,
	"REDACTED":  32,
}

func (x AttestedSignalKind) Text() string {
	return proto.EnumName(Notatedsignalkind_alias, int32(x))
}

func (AttestedSignalKind) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{0}
}

//
type FragmentAssignHeading struct {
	Sum uint32 `protobuf:"variableint,1,opt,name=total,proto3" json:"sum,omitempty"`
	Digest  []byte `protobuf:"octets,2,opt,name=hash,proto3" json:"digest,omitempty"`
}

func (m *FragmentAssignHeading) Restore()         { *m = FragmentAssignHeading{} }
func (m *FragmentAssignHeading) Text() string { return proto.CompactTextString(m) }
func (*FragmentAssignHeading) SchemaArtifact()    {}
func (*FragmentAssignHeading) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{0}
}
func (m *FragmentAssignHeading) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *FragmentAssignHeading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Fragmentsetheadline.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FragmentAssignHeading) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Fragmentsetheadline.Merge(m, src)
}
func (m *FragmentAssignHeading) XXX_Extent() int {
	return m.Extent()
}
func (m *FragmentAssignHeading) XXX_Dropunfamiliar() {
	xxx_signaldetails_Fragmentsetheadline.DiscardUnknown(m)
}

var xxx_signaldetails_Fragmentsetheadline proto.InternalMessageInfo

func (m *FragmentAssignHeading) ObtainSum() uint32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *FragmentAssignHeading) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Fragment struct {
	Ordinal uint32       `protobuf:"variableint,1,opt,name=index,proto3" json:"ordinal,omitempty"`
	Octets []byte       `protobuf:"octets,2,opt,name=bytes,proto3" json:"octets,omitempty"`
	Attestation security.Attestation `protobuf:"octets,3,opt,name=proof,proto3" json:"attestation"`
}

func (m *Fragment) Restore()         { *m = Fragment{} }
func (m *Fragment) Text() string { return proto.CompactTextString(m) }
func (*Fragment) SchemaArtifact()    {}
func (*Fragment) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{1}
}
func (m *Fragment) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Fragment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Fragment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fragment) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Fragment.Merge(m, src)
}
func (m *Fragment) XXX_Extent() int {
	return m.Extent()
}
func (m *Fragment) XXX_Dropunfamiliar() {
	xxx_signaldetails_Fragment.DiscardUnknown(m)
}

var xxx_signaldetails_Fragment proto.InternalMessageInfo

func (m *Fragment) ObtainOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *Fragment) ObtainOctets() []byte {
	if m != nil {
		return m.Octets
	}
	return nil
}

func (m *Fragment) ObtainAttestation() security.Attestation {
	if m != nil {
		return m.Attestation
	}
	return security.Attestation{}
}

//
type LedgerUUID struct {
	Digest          []byte        `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	FragmentAssignHeading FragmentAssignHeading `protobuf:"octets,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"fragment_assign_headline"`
}

func (m *LedgerUUID) Restore()         { *m = LedgerUUID{} }
func (m *LedgerUUID) Text() string { return proto.CompactTextString(m) }
func (*LedgerUUID) SchemaArtifact()    {}
func (*LedgerUUID) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{2}
}
func (m *LedgerUUID) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *LedgerUUID) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ledgeruuid.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerUUID) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ledgeruuid.Merge(m, src)
}
func (m *LedgerUUID) XXX_Extent() int {
	return m.Extent()
}
func (m *LedgerUUID) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ledgeruuid.DiscardUnknown(m)
}

var xxx_signaldetails_Ledgeruuid proto.InternalMessageInfo

func (m *LedgerUUID) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *LedgerUUID) ObtainFragmentAssignHeadline() FragmentAssignHeading {
	if m != nil {
		return m.FragmentAssignHeading
	}
	return FragmentAssignHeading{}
}

//
type Heading struct {
	//
	Edition edition.Agreement `protobuf:"octets,1,opt,name=version,proto3" json:"edition"`
	SuccessionUUID string            `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
	Altitude  int64             `protobuf:"variableint,3,opt,name=height,proto3" json:"altitude,omitempty"`
	Moment    time.Time         `protobuf:"octets,4,opt,name=time,proto3,stdtime" json:"moment"`
	//
	FinalLedgerUuid LedgerUUID `protobuf:"octets,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"final_ledger_uuid"`
	//
	FinalEndorseDigest []byte `protobuf:"octets,6,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"final_endorse_digest,omitempty"`
	DataDigest       []byte `protobuf:"octets,7,opt,name=data_hash,json=dataHash,proto3" json:"data_digest,omitempty"`
	//
	AssessorsDigest     []byte `protobuf:"octets,8,opt,name=validators_hash,json=validatorsHash,proto3" json:"assessors_digest,omitempty"`
	FollowingAssessorsDigest []byte `protobuf:"octets,9,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_assessors_digest,omitempty"`
	AgreementDigest      []byte `protobuf:"octets,10,opt,name=consensus_hash,json=consensusHash,proto3" json:"agreement_digest,omitempty"`
	PlatformDigest            []byte `protobuf:"octets,11,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
	FinalOutcomesDigest    []byte `protobuf:"octets,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"final_outcomes_digest,omitempty"`
	//
	ProofDigest    []byte `protobuf:"octets,13,opt,name=evidence_hash,json=evidenceHash,proto3" json:"proof_digest,omitempty"`
	NominatorLocation []byte `protobuf:"octets,14,opt,name=proposer_address,json=proposerAddress,proto3" json:"nominator_location,omitempty"`
}

func (m *Heading) Restore()         { *m = Heading{} }
func (m *Heading) Text() string { return proto.CompactTextString(m) }
func (*Heading) SchemaArtifact()    {}
func (*Heading) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{3}
}
func (m *Heading) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Heading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Headline.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Heading) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Headline.Merge(m, src)
}
func (m *Heading) XXX_Extent() int {
	return m.Extent()
}
func (m *Heading) XXX_Dropunfamiliar() {
	xxx_signaldetails_Headline.DiscardUnknown(m)
}

var xxx_signaldetails_Headline proto.InternalMessageInfo

func (m *Heading) ObtainEdition() edition.Agreement {
	if m != nil {
		return m.Edition
	}
	return edition.Agreement{}
}

func (m *Heading) ObtainSuccessionUUID() string {
	if m != nil {
		return m.SuccessionUUID
	}
	return "REDACTED"
}

func (m *Heading) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Heading) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *Heading) ObtainFinalLedgerUuid() LedgerUUID {
	if m != nil {
		return m.FinalLedgerUuid
	}
	return LedgerUUID{}
}

func (m *Heading) ObtainFinalEndorseDigest() []byte {
	if m != nil {
		return m.FinalEndorseDigest
	}
	return nil
}

func (m *Heading) ObtainDataDigest() []byte {
	if m != nil {
		return m.DataDigest
	}
	return nil
}

func (m *Heading) ObtainAssessorsDigest() []byte {
	if m != nil {
		return m.AssessorsDigest
	}
	return nil
}

func (m *Heading) ObtainFollowingAssessorsDigest() []byte {
	if m != nil {
		return m.FollowingAssessorsDigest
	}
	return nil
}

func (m *Heading) ObtainAgreementDigest() []byte {
	if m != nil {
		return m.AgreementDigest
	}
	return nil
}

func (m *Heading) ObtainApplicationDigest() []byte {
	if m != nil {
		return m.PlatformDigest
	}
	return nil
}

func (m *Heading) ObtainFinalOutcomesDigest() []byte {
	if m != nil {
		return m.FinalOutcomesDigest
	}
	return nil
}

func (m *Heading) ObtainProofDigest() []byte {
	if m != nil {
		return m.ProofDigest
	}
	return nil
}

func (m *Heading) ObtainNominatorLocator() []byte {
	if m != nil {
		return m.NominatorLocation
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
func (m *Data) Text() string { return proto.CompactTextString(m) }
func (*Data) SchemaArtifact()    {}
func (*Data) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{4}
}
func (m *Data) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Data) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Data.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Data) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Data.Merge(m, src)
}
func (m *Data) XXX_Extent() int {
	return m.Extent()
}
func (m *Data) XXX_Dropunfamiliar() {
	xxx_signaldetails_Data.DiscardUnknown(m)
}

var xxx_signaldetails_Data proto.InternalMessageInfo

func (m *Data) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

//
//
type Ballot struct {
	Kind             AttestedSignalKind `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Altitude           int64         `protobuf:"variableint,2,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration            int32         `protobuf:"variableint,3,opt,name=round,proto3" json:"iteration,omitempty"`
	LedgerUUID          LedgerUUID       `protobuf:"octets,4,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid"`
	Timestamp        time.Time     `protobuf:"octets,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	AssessorLocation []byte        `protobuf:"octets,6,opt,name=validator_address,json=validatorAddress,proto3" json:"assessor_location,omitempty"`
	AssessorOrdinal   int32         `protobuf:"variableint,7,opt,name=validator_index,json=validatorIndex,proto3" json:"assessor_position,omitempty"`
	//
	//
	Notation []byte `protobuf:"octets,8,opt,name=signature,proto3" json:"signing,omitempty"`
	//
	//
	Addition []byte `protobuf:"octets,9,opt,name=extension,proto3" json:"addition,omitempty"`
	//
	//
	//
	AdditionNotation []byte `protobuf:"octets,10,opt,name=extension_signature,json=extensionSignature,proto3" json:"addition_signing,omitempty"`
}

func (m *Ballot) Restore()         { *m = Ballot{} }
func (m *Ballot) Text() string { return proto.CompactTextString(m) }
func (*Ballot) SchemaArtifact()    {}
func (*Ballot) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{5}
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

func (m *Ballot) ObtainKind() AttestedSignalKind {
	if m != nil {
		return m.Kind
	}
	return UnfamiliarKind
}

func (m *Ballot) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Ballot) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *Ballot) ObtainLedgerUUID() LedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return LedgerUUID{}
}

func (m *Ballot) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Ballot) ObtainAssessorLocation() []byte {
	if m != nil {
		return m.AssessorLocation
	}
	return nil
}

func (m *Ballot) ObtainAssessorOrdinal() int32 {
	if m != nil {
		return m.AssessorOrdinal
	}
	return 0
}

func (m *Ballot) ObtainNotation() []byte {
	if m != nil {
		return m.Notation
	}
	return nil
}

func (m *Ballot) ObtainAddition() []byte {
	if m != nil {
		return m.Addition
	}
	return nil
}

func (m *Ballot) ObtainAdditionSigning() []byte {
	if m != nil {
		return m.AdditionNotation
	}
	return nil
}

//
type Endorse struct {
	Altitude     int64       `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration      int32       `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	LedgerUUID    LedgerUUID     `protobuf:"octets,3,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid"`
	Notations []EndorseSignature `protobuf:"octets,4,rep,name=signatures,proto3" json:"notations"`
}

func (m *Endorse) Restore()         { *m = Endorse{} }
func (m *Endorse) Text() string { return proto.CompactTextString(m) }
func (*Endorse) SchemaArtifact()    {}
func (*Endorse) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{6}
}
func (m *Endorse) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Endorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Endorse.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Endorse) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Endorse.Merge(m, src)
}
func (m *Endorse) XXX_Extent() int {
	return m.Extent()
}
func (m *Endorse) XXX_Dropunfamiliar() {
	xxx_signaldetails_Endorse.DiscardUnknown(m)
}

var xxx_signaldetails_Endorse proto.InternalMessageInfo

func (m *Endorse) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Endorse) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *Endorse) ObtainLedgerUUID() LedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return LedgerUUID{}
}

func (m *Endorse) ObtainNotations() []EndorseSignature {
	if m != nil {
		return m.Notations
	}
	return nil
}

//
type EndorseSignature struct {
	LedgerUuidMarker      LedgerUUIDMarker `protobuf:"variableint,1,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uuid_marker,omitempty"`
	AssessorLocation []byte      `protobuf:"octets,2,opt,name=validator_address,json=validatorAddress,proto3" json:"assessor_location,omitempty"`
	Timestamp        time.Time   `protobuf:"octets,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Notation        []byte      `protobuf:"octets,4,opt,name=signature,proto3" json:"signing,omitempty"`
}

func (m *EndorseSignature) Restore()         { *m = EndorseSignature{} }
func (m *EndorseSignature) Text() string { return proto.CompactTextString(m) }
func (*EndorseSignature) SchemaArtifact()    {}
func (*EndorseSignature) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{7}
}
func (m *EndorseSignature) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *EndorseSignature) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Commitsig.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndorseSignature) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Commitsig.Merge(m, src)
}
func (m *EndorseSignature) XXX_Extent() int {
	return m.Extent()
}
func (m *EndorseSignature) XXX_Dropunfamiliar() {
	xxx_signaldetails_Commitsig.DiscardUnknown(m)
}

var xxx_signaldetails_Commitsig proto.InternalMessageInfo

func (m *EndorseSignature) ObtainLedgerUuidMarker() LedgerUUIDMarker {
	if m != nil {
		return m.LedgerUuidMarker
	}
	return LedgerUUIDMarkerUnfamiliar
}

func (m *EndorseSignature) ObtainAssessorLocation() []byte {
	if m != nil {
		return m.AssessorLocation
	}
	return nil
}

func (m *EndorseSignature) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *EndorseSignature) ObtainNotation() []byte {
	if m != nil {
		return m.Notation
	}
	return nil
}

type ExpandedEndorse struct {
	Altitude             int64               `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration              int32               `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	LedgerUUID            LedgerUUID             `protobuf:"octets,3,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid"`
	ExpandedNotations []ExpandedEndorseSignature `protobuf:"octets,4,rep,name=extended_signatures,json=extendedSignatures,proto3" json:"expanded_notations"`
}

func (m *ExpandedEndorse) Restore()         { *m = ExpandedEndorse{} }
func (m *ExpandedEndorse) Text() string { return proto.CompactTextString(m) }
func (*ExpandedEndorse) SchemaArtifact()    {}
func (*ExpandedEndorse) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{8}
}
func (m *ExpandedEndorse) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ExpandedEndorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Addncommit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedEndorse) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Addncommit.Merge(m, src)
}
func (m *ExpandedEndorse) XXX_Extent() int {
	return m.Extent()
}
func (m *ExpandedEndorse) XXX_Dropunfamiliar() {
	xxx_signaldetails_Addncommit.DiscardUnknown(m)
}

var xxx_signaldetails_Addncommit proto.InternalMessageInfo

func (m *ExpandedEndorse) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *ExpandedEndorse) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *ExpandedEndorse) ObtainLedgerUUID() LedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return LedgerUUID{}
}

func (m *ExpandedEndorse) ObtainExpandedNotations() []ExpandedEndorseSignature {
	if m != nil {
		return m.ExpandedNotations
	}
	return nil
}

//
//
//
type ExpandedEndorseSignature struct {
	LedgerUuidMarker      LedgerUUIDMarker `protobuf:"variableint,1,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uuid_marker,omitempty"`
	AssessorLocation []byte      `protobuf:"octets,2,opt,name=validator_address,json=validatorAddress,proto3" json:"assessor_location,omitempty"`
	Timestamp        time.Time   `protobuf:"octets,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Notation        []byte      `protobuf:"octets,4,opt,name=signature,proto3" json:"signing,omitempty"`
	//
	Addition []byte `protobuf:"octets,5,opt,name=extension,proto3" json:"addition,omitempty"`
	//
	AdditionNotation []byte `protobuf:"octets,6,opt,name=extension_signature,json=extensionSignature,proto3" json:"addition_signing,omitempty"`
}

func (m *ExpandedEndorseSignature) Restore()         { *m = ExpandedEndorseSignature{} }
func (m *ExpandedEndorseSignature) Text() string { return proto.CompactTextString(m) }
func (*ExpandedEndorseSignature) SchemaArtifact()    {}
func (*ExpandedEndorseSignature) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{9}
}
func (m *ExpandedEndorseSignature) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ExpandedEndorseSignature) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Addncommitsig.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedEndorseSignature) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Addncommitsig.Merge(m, src)
}
func (m *ExpandedEndorseSignature) XXX_Extent() int {
	return m.Extent()
}
func (m *ExpandedEndorseSignature) XXX_Dropunfamiliar() {
	xxx_signaldetails_Addncommitsig.DiscardUnknown(m)
}

var xxx_signaldetails_Addncommitsig proto.InternalMessageInfo

func (m *ExpandedEndorseSignature) ObtainLedgerUuidMarker() LedgerUUIDMarker {
	if m != nil {
		return m.LedgerUuidMarker
	}
	return LedgerUUIDMarkerUnfamiliar
}

func (m *ExpandedEndorseSignature) ObtainAssessorLocation() []byte {
	if m != nil {
		return m.AssessorLocation
	}
	return nil
}

func (m *ExpandedEndorseSignature) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *ExpandedEndorseSignature) ObtainNotation() []byte {
	if m != nil {
		return m.Notation
	}
	return nil
}

func (m *ExpandedEndorseSignature) ObtainAddition() []byte {
	if m != nil {
		return m.Addition
	}
	return nil
}

func (m *ExpandedEndorseSignature) ObtainAdditionSigning() []byte {
	if m != nil {
		return m.AdditionNotation
	}
	return nil
}

type Nomination struct {
	Kind      AttestedSignalKind `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"kind,omitempty"`
	Altitude    int64         `protobuf:"variableint,2,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration     int32         `protobuf:"variableint,3,opt,name=round,proto3" json:"iteration,omitempty"`
	PolicyIteration  int32         `protobuf:"variableint,4,opt,name=pol_round,json=polRound,proto3" json:"policy_iteration,omitempty"`
	LedgerUUID   LedgerUUID       `protobuf:"octets,5,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid"`
	Timestamp time.Time     `protobuf:"octets,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Notation []byte        `protobuf:"octets,7,opt,name=signature,proto3" json:"signing,omitempty"`
}

func (m *Nomination) Restore()         { *m = Nomination{} }
func (m *Nomination) Text() string { return proto.CompactTextString(m) }
func (*Nomination) SchemaArtifact()    {}
func (*Nomination) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{10}
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

func (m *Nomination) ObtainKind() AttestedSignalKind {
	if m != nil {
		return m.Kind
	}
	return UnfamiliarKind
}

func (m *Nomination) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Nomination) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *Nomination) ObtainPolicyIteration() int32 {
	if m != nil {
		return m.PolicyIteration
	}
	return 0
}

func (m *Nomination) ObtainLedgerUUID() LedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return LedgerUUID{}
}

func (m *Nomination) ObtainTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Nomination) ObtainNotation() []byte {
	if m != nil {
		return m.Notation
	}
	return nil
}

type NotatedHeading struct {
	Heading *Heading `protobuf:"octets,1,opt,name=header,proto3" json:"heading,omitempty"`
	Endorse *Endorse `protobuf:"octets,2,opt,name=commit,proto3" json:"endorse,omitempty"`
}

func (m *NotatedHeading) Restore()         { *m = NotatedHeading{} }
func (m *NotatedHeading) Text() string { return proto.CompactTextString(m) }
func (*NotatedHeading) SchemaArtifact()    {}
func (*NotatedHeading) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{11}
}
func (m *NotatedHeading) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *NotatedHeading) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Notatedheadline.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotatedHeading) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Notatedheadline.Merge(m, src)
}
func (m *NotatedHeading) XXX_Extent() int {
	return m.Extent()
}
func (m *NotatedHeading) XXX_Dropunfamiliar() {
	xxx_signaldetails_Notatedheadline.DiscardUnknown(m)
}

var xxx_signaldetails_Notatedheadline proto.InternalMessageInfo

func (m *NotatedHeading) ObtainHeadline() *Heading {
	if m != nil {
		return m.Heading
	}
	return nil
}

func (m *NotatedHeading) ObtainEndorse() *Endorse {
	if m != nil {
		return m.Endorse
	}
	return nil
}

type AgileLedger struct {
	NotatedHeading *NotatedHeading `protobuf:"octets,1,opt,name=signed_header,json=signedHeader,proto3" json:"notated_heading,omitempty"`
	AssessorAssign *AssessorAssign `protobuf:"octets,2,opt,name=validator_set,json=validatorSet,proto3" json:"assessor_assign,omitempty"`
}

func (m *AgileLedger) Restore()         { *m = AgileLedger{} }
func (m *AgileLedger) Text() string { return proto.CompactTextString(m) }
func (*AgileLedger) SchemaArtifact()    {}
func (*AgileLedger) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{12}
}
func (m *AgileLedger) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AgileLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Agileledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgileLedger) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Agileledger.Merge(m, src)
}
func (m *AgileLedger) XXX_Extent() int {
	return m.Extent()
}
func (m *AgileLedger) XXX_Dropunfamiliar() {
	xxx_signaldetails_Agileledger.DiscardUnknown(m)
}

var xxx_signaldetails_Agileledger proto.InternalMessageInfo

func (m *AgileLedger) ObtainNotatedHeadline() *NotatedHeading {
	if m != nil {
		return m.NotatedHeading
	}
	return nil
}

func (m *AgileLedger) ObtainAssessorAssign() *AssessorAssign {
	if m != nil {
		return m.AssessorAssign
	}
	return nil
}

type LedgerSummary struct {
	LedgerUUID   LedgerUUID `protobuf:"octets,1,opt,name=block_id,json=blockId,proto3" json:"ledger_uuid"`
	LedgerExtent int64   `protobuf:"variableint,2,opt,name=block_size,json=blockSize,proto3" json:"ledger_extent,omitempty"`
	Heading    Heading  `protobuf:"octets,3,opt,name=header,proto3" json:"heading"`
	CountTrans    int64   `protobuf:"variableint,4,opt,name=num_txs,json=numTxs,proto3" json:"count_trans,omitempty"`
}

func (m *LedgerSummary) Restore()         { *m = LedgerSummary{} }
func (m *LedgerSummary) Text() string { return proto.CompactTextString(m) }
func (*LedgerSummary) SchemaArtifact()    {}
func (*LedgerSummary) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{13}
}
func (m *LedgerSummary) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *LedgerSummary) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ledgermeta.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerSummary) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ledgermeta.Merge(m, src)
}
func (m *LedgerSummary) XXX_Extent() int {
	return m.Extent()
}
func (m *LedgerSummary) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ledgermeta.DiscardUnknown(m)
}

var xxx_signaldetails_Ledgermeta proto.InternalMessageInfo

func (m *LedgerSummary) ObtainLedgerUUID() LedgerUUID {
	if m != nil {
		return m.LedgerUUID
	}
	return LedgerUUID{}
}

func (m *LedgerSummary) ObtainLedgerExtent() int64 {
	if m != nil {
		return m.LedgerExtent
	}
	return 0
}

func (m *LedgerSummary) ObtainHeadline() Heading {
	if m != nil {
		return m.Heading
	}
	return Heading{}
}

func (m *LedgerSummary) ObtainCountTrans() int64 {
	if m != nil {
		return m.CountTrans
	}
	return 0
}

//
type TransferAttestation struct {
	OriginDigest []byte        `protobuf:"octets,1,opt,name=root_hash,json=rootHash,proto3" json:"origin_digest,omitempty"`
	Data     []byte        `protobuf:"octets,2,opt,name=data,proto3" json:"data,omitempty"`
	Attestation    *security.Attestation `protobuf:"octets,3,opt,name=proof,proto3" json:"attestation,omitempty"`
}

func (m *TransferAttestation) Restore()         { *m = TransferAttestation{} }
func (m *TransferAttestation) Text() string { return proto.CompactTextString(m) }
func (*TransferAttestation) SchemaArtifact()    {}
func (*TransferAttestation) Definition() ([]byte, []int) {
	return filedescriptor_d3a6e55e2345de56, []int{14}
}
func (m *TransferAttestation) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *TransferAttestation) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Transattestation.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferAttestation) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Transattestation.Merge(m, src)
}
func (m *TransferAttestation) XXX_Extent() int {
	return m.Extent()
}
func (m *TransferAttestation) XXX_Dropunfamiliar() {
	xxx_signaldetails_Transattestation.DiscardUnknown(m)
}

var xxx_signaldetails_Transattestation proto.InternalMessageInfo

func (m *TransferAttestation) ObtainOriginDigest() []byte {
	if m != nil {
		return m.OriginDigest
	}
	return nil
}

func (m *TransferAttestation) ObtainData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TransferAttestation) ObtainAttestation() *security.Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

func initialize() {
	proto.RegisterEnum("REDACTED", Notatedsignalkind_alias, Notatedsignalkind_datum)
	proto.RegisterType((*FragmentAssignHeading)(nil), "REDACTED")
	proto.RegisterType((*Fragment)(nil), "REDACTED")
	proto.RegisterType((*LedgerUUID)(nil), "REDACTED")
	proto.RegisterType((*Heading)(nil), "REDACTED")
	proto.RegisterType((*Data)(nil), "REDACTED")
	proto.RegisterType((*Ballot)(nil), "REDACTED")
	proto.RegisterType((*Endorse)(nil), "REDACTED")
	proto.RegisterType((*EndorseSignature)(nil), "REDACTED")
	proto.RegisterType((*ExpandedEndorse)(nil), "REDACTED")
	proto.RegisterType((*ExpandedEndorseSignature)(nil), "REDACTED")
	proto.RegisterType((*Nomination)(nil), "REDACTED")
	proto.RegisterType((*NotatedHeading)(nil), "REDACTED")
	proto.RegisterType((*AgileLedger)(nil), "REDACTED")
	proto.RegisterType((*LedgerSummary)(nil), "REDACTED")
	proto.RegisterType((*TransferAttestation)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_d3a6e55e2345de56) }

var filedescriptor_d3a6e55e2345de56 = []byte{
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

func (m *FragmentAssignHeading) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *FragmentAssignHeading) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *FragmentAssignHeading) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Sum != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Sum))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Fragment) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Fragment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Fragment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.Attestation.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x1a
	if len(m.Octets) > 0 {
		i -= len(m.Octets)
		copy(deltaLocatedAN[i:], m.Octets)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Octets)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Ordinal != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ordinal))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *LedgerUUID) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *LedgerUUID) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *LedgerUUID) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x12
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Heading) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Heading) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Heading) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.NominatorLocation) > 0 {
		i -= len(m.NominatorLocation)
		copy(deltaLocatedAN[i:], m.NominatorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.NominatorLocation)))
		i--
		deltaLocatedAN[i] = 0x72
	}
	if len(m.ProofDigest) > 0 {
		i -= len(m.ProofDigest)
		copy(deltaLocatedAN[i:], m.ProofDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.ProofDigest)))
		i--
		deltaLocatedAN[i] = 0x6a
	}
	if len(m.FinalOutcomesDigest) > 0 {
		i -= len(m.FinalOutcomesDigest)
		copy(deltaLocatedAN[i:], m.FinalOutcomesDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FinalOutcomesDigest)))
		i--
		deltaLocatedAN[i] = 0x62
	}
	if len(m.PlatformDigest) > 0 {
		i -= len(m.PlatformDigest)
		copy(deltaLocatedAN[i:], m.PlatformDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.PlatformDigest)))
		i--
		deltaLocatedAN[i] = 0x5a
	}
	if len(m.AgreementDigest) > 0 {
		i -= len(m.AgreementDigest)
		copy(deltaLocatedAN[i:], m.AgreementDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AgreementDigest)))
		i--
		deltaLocatedAN[i] = 0x52
	}
	if len(m.FollowingAssessorsDigest) > 0 {
		i -= len(m.FollowingAssessorsDigest)
		copy(deltaLocatedAN[i:], m.FollowingAssessorsDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FollowingAssessorsDigest)))
		i--
		deltaLocatedAN[i] = 0x4a
	}
	if len(m.AssessorsDigest) > 0 {
		i -= len(m.AssessorsDigest)
		copy(deltaLocatedAN[i:], m.AssessorsDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AssessorsDigest)))
		i--
		deltaLocatedAN[i] = 0x42
	}
	if len(m.DataDigest) > 0 {
		i -= len(m.DataDigest)
		copy(deltaLocatedAN[i:], m.DataDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.DataDigest)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	if len(m.FinalEndorseDigest) > 0 {
		i -= len(m.FinalEndorseDigest)
		copy(deltaLocatedAN[i:], m.FinalEndorseDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FinalEndorseDigest)))
		i--
		deltaLocatedAN[i] = 0x32
	}
	{
		extent, err := m.FinalLedgerUuid.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x2a
	n4, fault4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault4 != nil {
		return 0, fault4
	}
	i -= n4
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n4))
	i--
	deltaLocatedAN[i] = 0x22
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if len(m.SuccessionUUID) > 0 {
		i -= len(m.SuccessionUUID)
		copy(deltaLocatedAN[i:], m.SuccessionUUID)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.SuccessionUUID)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	{
		extent, err := m.Edition.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *Data) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Data) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Data) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for idxNdExc := len(m.Txs) - 1; idxNdExc >= 0; idxNdExc-- {
			i -= len(m.Txs[idxNdExc])
			copy(deltaLocatedAN[i:], m.Txs[idxNdExc])
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Txs[idxNdExc])))
			i--
			deltaLocatedAN[i] = 0xa
		}
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
	if len(m.AdditionNotation) > 0 {
		i -= len(m.AdditionNotation)
		copy(deltaLocatedAN[i:], m.AdditionNotation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AdditionNotation)))
		i--
		deltaLocatedAN[i] = 0x52
	}
	if len(m.Addition) > 0 {
		i -= len(m.Addition)
		copy(deltaLocatedAN[i:], m.Addition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Addition)))
		i--
		deltaLocatedAN[i] = 0x4a
	}
	if len(m.Notation) > 0 {
		i -= len(m.Notation)
		copy(deltaLocatedAN[i:], m.Notation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Notation)))
		i--
		deltaLocatedAN[i] = 0x42
	}
	if m.AssessorOrdinal != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.AssessorOrdinal))
		i--
		deltaLocatedAN[i] = 0x38
	}
	if len(m.AssessorLocation) > 0 {
		i -= len(m.AssessorLocation)
		copy(deltaLocatedAN[i:], m.AssessorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AssessorLocation)))
		i--
		deltaLocatedAN[i] = 0x32
	}
	n6, fault6 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault6 != nil {
		return 0, fault6
	}
	i -= n6
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n6))
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
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Kind != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Endorse) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Endorse) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Endorse) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Notations) > 0 {
		for idxNdExc := len(m.Notations) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Notations[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x22
		}
	}
	{
		extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *EndorseSignature) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *EndorseSignature) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *EndorseSignature) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Notation) > 0 {
		i -= len(m.Notation)
		copy(deltaLocatedAN[i:], m.Notation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Notation)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	n9, fault9 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault9 != nil {
		return 0, fault9
	}
	i -= n9
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n9))
	i--
	deltaLocatedAN[i] = 0x1a
	if len(m.AssessorLocation) > 0 {
		i -= len(m.AssessorLocation)
		copy(deltaLocatedAN[i:], m.AssessorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AssessorLocation)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.LedgerUuidMarker != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.LedgerUuidMarker))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ExpandedEndorse) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ExpandedEndorse) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ExpandedEndorse) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.ExpandedNotations) > 0 {
		for idxNdExc := len(m.ExpandedNotations) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.ExpandedNotations[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x22
		}
	}
	{
		extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *ExpandedEndorseSignature) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ExpandedEndorseSignature) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ExpandedEndorseSignature) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.AdditionNotation) > 0 {
		i -= len(m.AdditionNotation)
		copy(deltaLocatedAN[i:], m.AdditionNotation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AdditionNotation)))
		i--
		deltaLocatedAN[i] = 0x32
	}
	if len(m.Addition) > 0 {
		i -= len(m.Addition)
		copy(deltaLocatedAN[i:], m.Addition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Addition)))
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if len(m.Notation) > 0 {
		i -= len(m.Notation)
		copy(deltaLocatedAN[i:], m.Notation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Notation)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	n11, fault11 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault11 != nil {
		return 0, fault11
	}
	i -= n11
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n11))
	i--
	deltaLocatedAN[i] = 0x1a
	if len(m.AssessorLocation) > 0 {
		i -= len(m.AssessorLocation)
		copy(deltaLocatedAN[i:], m.AssessorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AssessorLocation)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.LedgerUuidMarker != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.LedgerUuidMarker))
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
	if len(m.Notation) > 0 {
		i -= len(m.Notation)
		copy(deltaLocatedAN[i:], m.Notation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Notation)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	n12, fault12 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if fault12 != nil {
		return 0, fault12
	}
	i -= n12
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n12))
	i--
	deltaLocatedAN[i] = 0x32
	{
		extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x2a
	if m.PolicyIteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.PolicyIteration))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Kind != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *NotatedHeading) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *NotatedHeading) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *NotatedHeading) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Endorse != nil {
		{
			extent, err := m.Endorse.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Heading != nil {
		{
			extent, err := m.Heading.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *AgileLedger) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AgileLedger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AgileLedger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.AssessorAssign != nil {
		{
			extent, err := m.AssessorAssign.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.NotatedHeading != nil {
		{
			extent, err := m.NotatedHeading.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *LedgerSummary) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *LedgerSummary) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *LedgerSummary) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.CountTrans != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.CountTrans))
		i--
		deltaLocatedAN[i] = 0x20
	}
	{
		extent, err := m.Heading.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x1a
	if m.LedgerExtent != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.LedgerExtent))
		i--
		deltaLocatedAN[i] = 0x10
	}
	{
		extent, err := m.LedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *TransferAttestation) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *TransferAttestation) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *TransferAttestation) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(deltaLocatedAN[i:], m.Data)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Data)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.OriginDigest) > 0 {
		i -= len(m.OriginDigest)
		copy(deltaLocatedAN[i:], m.OriginDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.OriginDigest)))
		i--
		deltaLocatedAN[i] = 0xa
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
func (m *FragmentAssignHeading) Extent() (n int) {
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

func (m *Fragment) Extent() (n int) {
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
	l = m.Attestation.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *LedgerUUID) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = m.FragmentAssignHeading.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *Heading) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Edition.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.SuccessionUUID)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovKinds(uint64(l))
	l = m.FinalLedgerUuid.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FinalEndorseDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.DataDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AssessorsDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.FollowingAssessorsDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AgreementDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.PlatformDigest)
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
	l = len(m.NominatorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Data) Extent() (n int) {
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

func (m *Ballot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	l = m.LedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.AssessorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.AssessorOrdinal != 0 {
		n += 1 + sovKinds(uint64(m.AssessorOrdinal))
	}
	l = len(m.Notation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Addition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AdditionNotation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Endorse) Extent() (n int) {
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
	l = m.LedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Notations) > 0 {
		for _, e := range m.Notations {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *EndorseSignature) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerUuidMarker != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUuidMarker))
	}
	l = len(m.AssessorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Notation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ExpandedEndorse) Extent() (n int) {
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
	l = m.LedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.ExpandedNotations) > 0 {
		for _, e := range m.ExpandedNotations {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ExpandedEndorseSignature) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerUuidMarker != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUuidMarker))
	}
	l = len(m.AssessorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Notation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Addition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AdditionNotation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Nomination) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	if m.PolicyIteration != 0 {
		n += 1 + sovKinds(uint64(m.PolicyIteration))
	}
	l = m.LedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Notation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *NotatedHeading) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Heading != nil {
		l = m.Heading.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Endorse != nil {
		l = m.Endorse.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AgileLedger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotatedHeading != nil {
		l = m.NotatedHeading.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.AssessorAssign != nil {
		l = m.AssessorAssign.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *LedgerSummary) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.LedgerExtent != 0 {
		n += 1 + sovKinds(uint64(m.LedgerExtent))
	}
	l = m.Heading.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.CountTrans != 0 {
		n += 1 + sovKinds(uint64(m.CountTrans))
	}
	return n
}

func (m *TransferAttestation) Extent() (n int) {
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
		l = m.Attestation.Extent()
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
func (m *FragmentAssignHeading) Decode(deltaLocatedAN []byte) error {
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
			m.Sum = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
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
func (m *Fragment) Decode(deltaLocatedAN []byte) error {
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
				m.Ordinal |= uint32(b&0x7F) << relocate
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Octets = append(m.Octets[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Octets == nil {
				m.Octets = []byte{}
			}
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
			if err := m.Attestation.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *LedgerUUID) Decode(deltaLocatedAN []byte) error {
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
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
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
			if err := m.FragmentAssignHeading.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *Heading) Decode(deltaLocatedAN []byte) error {
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
			if err := m.Edition.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.SuccessionUUID = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 3:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Moment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if err := m.FinalLedgerUuid.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalEndorseDigest = append(m.FinalEndorseDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FinalEndorseDigest == nil {
				m.FinalEndorseDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.DataDigest = append(m.DataDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.DataDigest == nil {
				m.DataDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AssessorsDigest = append(m.AssessorsDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AssessorsDigest == nil {
				m.AssessorsDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 9:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FollowingAssessorsDigest = append(m.FollowingAssessorsDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FollowingAssessorsDigest == nil {
				m.FollowingAssessorsDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 10:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AgreementDigest = append(m.AgreementDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AgreementDigest == nil {
				m.AgreementDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 11:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.PlatformDigest = append(m.PlatformDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.PlatformDigest == nil {
				m.PlatformDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 12:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalOutcomesDigest = append(m.FinalOutcomesDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FinalOutcomesDigest == nil {
				m.FinalOutcomesDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 13:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofDigest = append(m.ProofDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.ProofDigest == nil {
				m.ProofDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 14:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.NominatorLocation = append(m.NominatorLocation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.NominatorLocation == nil {
				m.NominatorLocation = []byte{}
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
func (m *Data) Decode(deltaLocatedAN []byte) error {
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
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, make([]byte, submitOrdinal-idxNdExc))
			copy(m.Txs[len(m.Txs)-1], deltaLocatedAN[idxNdExc:submitOrdinal])
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
				m.Kind |= AttestedSignalKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AssessorLocation = append(m.AssessorLocation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AssessorLocation == nil {
				m.AssessorLocation = []byte{}
			}
			idxNdExc = submitOrdinal
		case 7:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.AssessorOrdinal = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.AssessorOrdinal |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 8:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Notation = append(m.Notation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Notation == nil {
				m.Notation = []byte{}
			}
			idxNdExc = submitOrdinal
		case 9:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Addition = append(m.Addition[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Addition == nil {
				m.Addition = []byte{}
			}
			idxNdExc = submitOrdinal
		case 10:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionNotation = append(m.AdditionNotation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AdditionNotation == nil {
				m.AdditionNotation = []byte{}
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
func (m *Endorse) Decode(deltaLocatedAN []byte) error {
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
			if err := m.LedgerUUID.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Notations = append(m.Notations, EndorseSignature{})
			if err := m.Notations[len(m.Notations)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *EndorseSignature) Decode(deltaLocatedAN []byte) error {
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
			m.LedgerUuidMarker = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.LedgerUuidMarker |= LedgerUUIDMarker(b&0x7F) << relocate
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AssessorLocation = append(m.AssessorLocation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AssessorLocation == nil {
				m.AssessorLocation = []byte{}
			}
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Notation = append(m.Notation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Notation == nil {
				m.Notation = []byte{}
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
func (m *ExpandedEndorse) Decode(deltaLocatedAN []byte) error {
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
			if err := m.LedgerUUID.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.ExpandedNotations = append(m.ExpandedNotations, ExpandedEndorseSignature{})
			if err := m.ExpandedNotations[len(m.ExpandedNotations)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *ExpandedEndorseSignature) Decode(deltaLocatedAN []byte) error {
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
			m.LedgerUuidMarker = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.LedgerUuidMarker |= LedgerUUIDMarker(b&0x7F) << relocate
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
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AssessorLocation = append(m.AssessorLocation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AssessorLocation == nil {
				m.AssessorLocation = []byte{}
			}
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Notation = append(m.Notation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Notation == nil {
				m.Notation = []byte{}
			}
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Addition = append(m.Addition[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Addition == nil {
				m.Addition = []byte{}
			}
			idxNdExc = submitOrdinal
		case 6:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionNotation = append(m.AdditionNotation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AdditionNotation == nil {
				m.AdditionNotation = []byte{}
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
				m.Kind |= AttestedSignalKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
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
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PolicyIteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.PolicyIteration |= int32(b&0x7F) << relocate
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 7:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Notation = append(m.Notation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Notation == nil {
				m.Notation = []byte{}
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
func (m *NotatedHeading) Decode(deltaLocatedAN []byte) error {
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
			if m.Heading == nil {
				m.Heading = &Heading{}
			}
			if err := m.Heading.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.Endorse == nil {
				m.Endorse = &Endorse{}
			}
			if err := m.Endorse.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *AgileLedger) Decode(deltaLocatedAN []byte) error {
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
			if m.NotatedHeading == nil {
				m.NotatedHeading = &NotatedHeading{}
			}
			if err := m.NotatedHeading.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.AssessorAssign == nil {
				m.AssessorAssign = &AssessorAssign{}
			}
			if err := m.AssessorAssign.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *LedgerSummary) Decode(deltaLocatedAN []byte) error {
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
			if err := m.LedgerUUID.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.LedgerExtent = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.LedgerExtent |= int64(b&0x7F) << relocate
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
			if err := m.Heading.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.CountTrans = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.CountTrans |= int64(b&0x7F) << relocate
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
func (m *TransferAttestation) Decode(deltaLocatedAN []byte) error {
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
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginDigest = append(m.OriginDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.OriginDigest == nil {
				m.OriginDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
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
			if m.Attestation == nil {
				m.Attestation = &security.Attestation{}
			}
			if err := m.Attestation.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
