//
//

package status

import (
	fmt "fmt"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	kinds1 "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
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
//
//
//
type HeritageIfaceReplies struct {
	DispatchTrans []*kinds.InvokeTransferOutcome `protobuf:"octets,1,rep,name=deliver_txs,json=deliverTxs,proto3" json:"dispatch_trans,omitempty"`
	TerminateLedger   *ReplyTerminateLedger     `protobuf:"octets,2,opt,name=end_block,json=endBlock,proto3" json:"terminate_ledger,omitempty"`
	InitiateLedger *ReplyInitiateLedger   `protobuf:"octets,3,opt,name=begin_block,json=beginBlock,proto3" json:"initiate_ledger,omitempty"`
}

func (m *HeritageIfaceReplies) Restore()         { *m = HeritageIfaceReplies{} }
func (m *HeritageIfaceReplies) Text() string { return proto.CompactTextString(m) }
func (*HeritageIfaceReplies) SchemaArtifact()    {}
func (*HeritageIfaceReplies) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{0}
}
func (m *HeritageIfaceReplies) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *HeritageIfaceReplies) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Heritageifacereplies.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeritageIfaceReplies) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Heritageifacereplies.Merge(m, src)
}
func (m *HeritageIfaceReplies) XXX_Extent() int {
	return m.Extent()
}
func (m *HeritageIfaceReplies) XXX_Dropunfamiliar() {
	xxx_signaldetails_Heritageifacereplies.DiscardUnknown(m)
}

var xxx_signaldetails_Heritageifacereplies proto.InternalMessageInfo

func (m *HeritageIfaceReplies) ObtainDispatchTrans() []*kinds.InvokeTransferOutcome {
	if m != nil {
		return m.DispatchTrans
	}
	return nil
}

func (m *HeritageIfaceReplies) ObtainTerminateLedger() *ReplyTerminateLedger {
	if m != nil {
		return m.TerminateLedger
	}
	return nil
}

func (m *HeritageIfaceReplies) ObtainInitiateLedger() *ReplyInitiateLedger {
	if m != nil {
		return m.InitiateLedger
	}
	return nil
}

//
type ReplyInitiateLedger struct {
	Incidents []kinds.Incident `protobuf:"octets,1,rep,name=events,proto3" json:"incidents,omitempty"`
}

func (m *ReplyInitiateLedger) Restore()         { *m = ReplyInitiateLedger{} }
func (m *ReplyInitiateLedger) Text() string { return proto.CompactTextString(m) }
func (*ReplyInitiateLedger) SchemaArtifact()    {}
func (*ReplyInitiateLedger) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{1}
}
func (m *ReplyInitiateLedger) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyInitiateLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyinitiateledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyInitiateLedger) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyinitiateledger.Merge(m, src)
}
func (m *ReplyInitiateLedger) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyInitiateLedger) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyinitiateledger.DiscardUnknown(m)
}

var xxx_signaldetails_Replyinitiateledger proto.InternalMessageInfo

func (m *ReplyInitiateLedger) ObtainIncidents() []kinds.Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

//
type ReplyTerminateLedger struct {
	AssessorRevisions      []kinds.AssessorRevise `protobuf:"octets,1,rep,name=validator_updates,json=validatorUpdates,proto3" json:"assessor_revisions"`
	AgreementArgumentRevisions *kinds1.AgreementSettings `protobuf:"octets,2,opt,name=consensus_param_updates,json=consensusParamUpdates,proto3" json:"agreement_argument_revisions,omitempty"`
	Incidents                []kinds.Incident           `protobuf:"octets,3,rep,name=events,proto3" json:"incidents,omitempty"`
}

func (m *ReplyTerminateLedger) Restore()         { *m = ReplyTerminateLedger{} }
func (m *ReplyTerminateLedger) Text() string { return proto.CompactTextString(m) }
func (*ReplyTerminateLedger) SchemaArtifact()    {}
func (*ReplyTerminateLedger) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{2}
}
func (m *ReplyTerminateLedger) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyTerminateLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyterminateledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyTerminateLedger) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyterminateledger.Merge(m, src)
}
func (m *ReplyTerminateLedger) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyTerminateLedger) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyterminateledger.DiscardUnknown(m)
}

var xxx_signaldetails_Replyterminateledger proto.InternalMessageInfo

func (m *ReplyTerminateLedger) ObtainAssessorRevisions() []kinds.AssessorRevise {
	if m != nil {
		return m.AssessorRevisions
	}
	return nil
}

func (m *ReplyTerminateLedger) ObtainAgreementArgumentRevisions() *kinds1.AgreementSettings {
	if m != nil {
		return m.AgreementArgumentRevisions
	}
	return nil
}

func (m *ReplyTerminateLedger) ObtainIncidents() []kinds.Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

//
type AssessorsDetails struct {
	AssessorAssign      *kinds1.AssessorAssign `protobuf:"octets,1,opt,name=validator_set,json=validatorSet,proto3" json:"assessor_assign,omitempty"`
	FinalAltitudeAltered int64                `protobuf:"variableint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"final_altitude_altered,omitempty"`
}

func (m *AssessorsDetails) Restore()         { *m = AssessorsDetails{} }
func (m *AssessorsDetails) Text() string { return proto.CompactTextString(m) }
func (*AssessorsDetails) SchemaArtifact()    {}
func (*AssessorsDetails) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{3}
}
func (m *AssessorsDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AssessorsDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Assessorsdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssessorsDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Assessorsdetails.Merge(m, src)
}
func (m *AssessorsDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *AssessorsDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Assessorsdetails.DiscardUnknown(m)
}

var xxx_signaldetails_Assessorsdetails proto.InternalMessageInfo

func (m *AssessorsDetails) ObtainAssessorAssign() *kinds1.AssessorAssign {
	if m != nil {
		return m.AssessorAssign
	}
	return nil
}

func (m *AssessorsDetails) ObtainFinalAltitudeAltered() int64 {
	if m != nil {
		return m.FinalAltitudeAltered
	}
	return 0
}

//
type AgreementParametersDetails struct {
	AgreementSettings   kinds1.AgreementSettings `protobuf:"octets,1,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_parameters"`
	FinalAltitudeAltered int64                  `protobuf:"variableint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"final_altitude_altered,omitempty"`
}

func (m *AgreementParametersDetails) Restore()         { *m = AgreementParametersDetails{} }
func (m *AgreementParametersDetails) Text() string { return proto.CompactTextString(m) }
func (*AgreementParametersDetails) SchemaArtifact()    {}
func (*AgreementParametersDetails) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{4}
}
func (m *AgreementParametersDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AgreementParametersDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Agreementparametersdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgreementParametersDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Agreementparametersdetails.Merge(m, src)
}
func (m *AgreementParametersDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *AgreementParametersDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Agreementparametersdetails.DiscardUnknown(m)
}

var xxx_signaldetails_Agreementparametersdetails proto.InternalMessageInfo

func (m *AgreementParametersDetails) ObtainAgreementParameters() kinds1.AgreementSettings {
	if m != nil {
		return m.AgreementSettings
	}
	return kinds1.AgreementSettings{}
}

func (m *AgreementParametersDetails) ObtainFinalAltitudeAltered() int64 {
	if m != nil {
		return m.FinalAltitudeAltered
	}
	return 0
}

type IfaceRepliesDetails struct {
	HeritageIfaceReplies   *HeritageIfaceReplies         `protobuf:"octets,1,opt,name=legacy_abci_responses,json=legacyAbciResponses,proto3" json:"heritage_iface_replies,omitempty"`
	Altitude                int64                        `protobuf:"variableint,2,opt,name=height,proto3" json:"altitude,omitempty"`
	ReplyCulminateLedger *kinds.ReplyCulminateLedger `protobuf:"octets,3,opt,name=response_finalize_block,json=responseFinalizeBlock,proto3" json:"reply_culminate_ledger,omitempty"`
}

func (m *IfaceRepliesDetails) Restore()         { *m = IfaceRepliesDetails{} }
func (m *IfaceRepliesDetails) Text() string { return proto.CompactTextString(m) }
func (*IfaceRepliesDetails) SchemaArtifact()    {}
func (*IfaceRepliesDetails) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{5}
}
func (m *IfaceRepliesDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *IfaceRepliesDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ifacerepliesdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IfaceRepliesDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ifacerepliesdetails.Merge(m, src)
}
func (m *IfaceRepliesDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *IfaceRepliesDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ifacerepliesdetails.DiscardUnknown(m)
}

var xxx_signaldetails_Ifacerepliesdetails proto.InternalMessageInfo

func (m *IfaceRepliesDetails) ObtainHeritageIfaceReplies() *HeritageIfaceReplies {
	if m != nil {
		return m.HeritageIfaceReplies
	}
	return nil
}

func (m *IfaceRepliesDetails) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *IfaceRepliesDetails) ObtainReplyCulminateLedger() *kinds.ReplyCulminateLedger {
	if m != nil {
		return m.ReplyCulminateLedger
	}
	return nil
}

type Edition struct {
	Agreement edition.Agreement `protobuf:"octets,1,opt,name=consensus,proto3" json:"agreement"`
	Package  string            `protobuf:"octets,2,opt,name=software,proto3" json:"package,omitempty"`
}

func (m *Edition) Restore()         { *m = Edition{} }
func (m *Edition) Text() string { return proto.CompactTextString(m) }
func (*Edition) SchemaArtifact()    {}
func (*Edition) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{6}
}
func (m *Edition) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Edition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Edition.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Edition) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Edition.Merge(m, src)
}
func (m *Edition) XXX_Extent() int {
	return m.Extent()
}
func (m *Edition) XXX_Dropunfamiliar() {
	xxx_signaldetails_Edition.DiscardUnknown(m)
}

var xxx_signaldetails_Edition proto.InternalMessageInfo

func (m *Edition) ObtainAgreement() edition.Agreement {
	if m != nil {
		return m.Agreement
	}
	return edition.Agreement{}
}

func (m *Edition) ObtainPackage() string {
	if m != nil {
		return m.Package
	}
	return "REDACTED"
}

type Status struct {
	Edition Edition `protobuf:"octets,1,opt,name=version,proto3" json:"edition"`
	//
	SuccessionUUID       string `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
	PrimaryAltitude int64  `protobuf:"variableint,14,opt,name=initial_height,json=initialHeight,proto3" json:"primary_altitude,omitempty"`
	//
	FinalLedgerAltitude int64          `protobuf:"variableint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"final_ledger_altitude,omitempty"`
	FinalLedgerUUID     kinds1.LedgerUUID `protobuf:"octets,4,opt,name=last_block_id,json=lastBlockId,proto3" json:"final_ledger_uuid"`
	FinalLedgerMoment   time.Time      `protobuf:"octets,5,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"final_ledger_moment"`
	//
	//
	//
	//
	//
	//
	FollowingAssessors              *kinds1.AssessorAssign `protobuf:"octets,6,opt,name=next_validators,json=nextValidators,proto3" json:"following_assessors,omitempty"`
	Assessors                  *kinds1.AssessorAssign `protobuf:"octets,7,opt,name=validators,proto3" json:"assessors,omitempty"`
	FinalAssessors              *kinds1.AssessorAssign `protobuf:"octets,8,opt,name=last_validators,json=lastValidators,proto3" json:"final_assessors,omitempty"`
	FinalAltitudeAssessorsAltered int64                `protobuf:"variableint,9,opt,name=last_height_validators_changed,json=lastHeightValidatorsChanged,proto3" json:"final_altitude_assessors_altered,omitempty"`
	//
	//
	AgreementSettings                  kinds1.AgreementSettings `protobuf:"octets,10,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_parameters"`
	FinalAltitudeAgreementParametersAltered int64                  `protobuf:"variableint,11,opt,name=last_height_consensus_params_changed,json=lastHeightConsensusParamsChanged,proto3" json:"final_altitude_agreement_parameters_altered,omitempty"`
	//
	FinalOutcomesDigest []byte `protobuf:"octets,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"final_outcomes_digest,omitempty"`
	//
	PlatformDigest []byte `protobuf:"octets,13,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *Status) Restore()         { *m = Status{} }
func (m *Status) Text() string { return proto.CompactTextString(m) }
func (*Status) SchemaArtifact()    {}
func (*Status) Definition() ([]byte, []int) {
	return filedescriptor_ccfacf933f22bf93, []int{7}
}
func (m *Status) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Status) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Status.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Status.Merge(m, src)
}
func (m *Status) XXX_Extent() int {
	return m.Extent()
}
func (m *Status) XXX_Dropunfamiliar() {
	xxx_signaldetails_Status.DiscardUnknown(m)
}

var xxx_signaldetails_Status proto.InternalMessageInfo

func (m *Status) ObtainEdition() Edition {
	if m != nil {
		return m.Edition
	}
	return Edition{}
}

func (m *Status) ObtainSuccessionUUID() string {
	if m != nil {
		return m.SuccessionUUID
	}
	return "REDACTED"
}

func (m *Status) ObtainPrimaryAltitude() int64 {
	if m != nil {
		return m.PrimaryAltitude
	}
	return 0
}

func (m *Status) ObtainFinalLedgerAltitude() int64 {
	if m != nil {
		return m.FinalLedgerAltitude
	}
	return 0
}

func (m *Status) ObtainFinalLedgerUUID() kinds1.LedgerUUID {
	if m != nil {
		return m.FinalLedgerUUID
	}
	return kinds1.LedgerUUID{}
}

func (m *Status) ObtainFinalLedgerMoment() time.Time {
	if m != nil {
		return m.FinalLedgerMoment
	}
	return time.Time{}
}

func (m *Status) ObtainFollowingAssessors() *kinds1.AssessorAssign {
	if m != nil {
		return m.FollowingAssessors
	}
	return nil
}

func (m *Status) ObtainAssessors() *kinds1.AssessorAssign {
	if m != nil {
		return m.Assessors
	}
	return nil
}

func (m *Status) ObtainFinalAssessors() *kinds1.AssessorAssign {
	if m != nil {
		return m.FinalAssessors
	}
	return nil
}

func (m *Status) ObtainFinalAltitudeAssessorsAltered() int64 {
	if m != nil {
		return m.FinalAltitudeAssessorsAltered
	}
	return 0
}

func (m *Status) ObtainAgreementParameters() kinds1.AgreementSettings {
	if m != nil {
		return m.AgreementSettings
	}
	return kinds1.AgreementSettings{}
}

func (m *Status) ObtainFinalAltitudeAgreementParametersAltered() int64 {
	if m != nil {
		return m.FinalAltitudeAgreementParametersAltered
	}
	return 0
}

func (m *Status) ObtainFinalOutcomesDigest() []byte {
	if m != nil {
		return m.FinalOutcomesDigest
	}
	return nil
}

func (m *Status) ObtainApplicationDigest() []byte {
	if m != nil {
		return m.PlatformDigest
	}
	return nil
}

func initialize() {
	proto.RegisterType((*HeritageIfaceReplies)(nil), "REDACTED")
	proto.RegisterType((*ReplyInitiateLedger)(nil), "REDACTED")
	proto.RegisterType((*ReplyTerminateLedger)(nil), "REDACTED")
	proto.RegisterType((*AssessorsDetails)(nil), "REDACTED")
	proto.RegisterType((*AgreementParametersDetails)(nil), "REDACTED")
	proto.RegisterType((*IfaceRepliesDetails)(nil), "REDACTED")
	proto.RegisterType((*Edition)(nil), "REDACTED")
	proto.RegisterType((*Status)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_ccfacf933f22bf93) }

var filedescriptor_ccfacf933f22bf93 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4b, 0x6f, 0xdb, 0x46,
	0x17, 0x35, 0xe3, 0x44, 0x8f, 0x2b, 0xcb, 0x96, 0x47, 0x9f, 0x13, 0x45, 0xf9, 0x22, 0xa9, 0x42,
	0x12, 0x18, 0x45, 0x41, 0x01, 0xc9, 0xaa, 0x9b, 0x14, 0x96, 0xec, 0xd6, 0x02, 0xdc, 0xa2, 0xa0,
	0xdd, 0x00, 0xe9, 0x22, 0xc4, 0x88, 0x1c, 0x49, 0x83, 0x4a, 0x24, 0xc1, 0x19, 0xa9, 0x72, 0xf7,
	0xdd, 0x75, 0x91, 0x6d, 0xff, 0x51, 0x96, 0x59, 0x76, 0x53, 0xb7, 0x95, 0x81, 0x2e, 0xfa, 0x2b,
	0x8a, 0x79, 0xf0, 0x25, 0xba, 0xa8, 0x8b, 0xec, 0xc8, 0xb9, 0xe7, 0x9e, 0x7b, 0xee, 0x99, 0xb9,
	0x43, 0xc2, 0xff, 0x39, 0xf1, 0x5c, 0x12, 0xce, 0xa9, 0xc7, 0x7b, 0x8c, 0x63, 0x4e, 0x7a, 0xfc,
	0x32, 0x20, 0xcc, 0x0c, 0x42, 0x9f, 0xfb, 0xa8, 0x96, 0x44, 0x4d, 0x19, 0x6d, 0xfe, 0x6f, 0xe2,
	0x4f, 0x7c, 0x19, 0xec, 0x89, 0x27, 0x85, 0x6b, 0x3e, 0x4a, 0xb1, 0xe0, 0x91, 0x43, 0xd3, 0x24,
	0xcd, 0x74, 0x09, 0xb9, 0x9e, 0x89, 0x76, 0x72, 0xd1, 0x25, 0x9e, 0x51, 0x17, 0x73, 0x3f, 0xd4,
	0x88, 0xc7, 0x39, 0x44, 0x80, 0x43, 0x3c, 0x8f, 0x08, 0x5a, 0xa9, 0xf0, 0x92, 0x84, 0x8c, 0xfa,
	0x5e, 0xa6, 0x40, 0x7b, 0xe2, 0xfb, 0x93, 0x19, 0xe9, 0xc9, 0xb7, 0xd1, 0x62, 0xdc, 0xe3, 0x74,
	0x4e, 0x18, 0xc7, 0xf3, 0x40, 0x01, 0xba, 0xbf, 0x1a, 0x50, 0x3f, 0x23, 0x13, 0xec, 0x5c, 0x1e,
	0xf5, 0x07, 0x43, 0x8b, 0xb0, 0xc0, 0xf7, 0x18, 0x61, 0xe8, 0x25, 0x54, 0x5c, 0x32, 0xa3, 0x4b,
	0x12, 0xda, 0x7c, 0xc5, 0x1a, 0x46, 0x67, 0xfb, 0xb0, 0xf2, 0xfc, 0xb1, 0x99, 0xb2, 0x44, 0xb4,
	0x6a, 0x9e, 0xac, 0x88, 0x73, 0xb1, 0xb2, 0x08, 0x5b, 0xcc, 0xb8, 0x05, 0x3a, 0xe3, 0x62, 0xc5,
	0xd0, 0x67, 0x50, 0x26, 0x9e, 0x6b, 0x8f, 0x66, 0xbe, 0xf3, 0x5d, 0xe3, 0x4e, 0xc7, 0x38, 0xac,
	0x3c, 0xef, 0x9a, 0x9b, 0x86, 0x9a, 0x51, 0xbd, 0x13, 0xcf, 0xed, 0x0b, 0xa4, 0x55, 0x22, 0xfa,
	0x09, 0x9d, 0x40, 0x65, 0x44, 0x26, 0xd4, 0xd3, 0x14, 0xdb, 0x92, 0xe2, 0xc9, 0x3f, 0x53, 0xf4,
	0x05, 0x58, 0x91, 0xc0, 0x28, 0x7e, 0xee, 0xbe, 0x01, 0x94, 0x47, 0xa0, 0x53, 0x28, 0x90, 0x25,
	0xf1, 0x78, 0xd4, 0xd8, 0xfd, 0x7c, 0x63, 0x22, 0xdc, 0x6f, 0xbc, 0xbb, 0x6a, 0x6f, 0xfd, 0x75,
	0xd5, 0xae, 0x29, 0xf4, 0x27, 0xfe, 0x9c, 0x72, 0x32, 0x0f, 0xf8, 0xa5, 0xa5, 0xf3, 0xbb, 0x3f,
	0xdd, 0x81, 0xda, 0x66, 0x17, 0xe8, 0x1c, 0xf6, 0xe3, 0x7d, 0xb4, 0x17, 0x81, 0x8b, 0x39, 0x89,
	0x2a, 0x75, 0x72, 0x95, 0x5e, 0x45, 0xc8, 0x6f, 0x24, 0xb0, 0x7f, 0x57, 0xd4, 0xb4, 0x6a, 0xcb,
	0xec, 0x32, 0x43, 0xaf, 0xe1, 0x81, 0x23, 0xaa, 0x78, 0x6c, 0xc1, 0x6c, 0x79, 0x08, 0x62, 0x6a,
	0xe5, 0xef, 0x47, 0x69, 0x6a, 0x75, 0x08, 0x06, 0x51, 0xc2, 0xd7, 0xf2, 0xd0, 0x58, 0x07, 0x4e,
	0x66, 0x21, 0xa2, 0x4e, 0xec, 0xd8, 0xfe, 0x40, 0x3b, 0x7e, 0x34, 0x60, 0x37, 0x6e, 0x88, 0x0d,
	0xbd, 0xb1, 0x8f, 0x06, 0x50, 0x4d, 0xcc, 0x60, 0x84, 0x37, 0x0c, 0xa9, 0xb6, 0x95, 0x57, 0x1b,
	0x27, 0x9e, 0x13, 0x6e, 0xed, 0x2c, 0x53, 0x6f, 0xc8, 0x84, 0xfa, 0x0c, 0x33, 0x6e, 0x4f, 0x09,
	0x9d, 0x4c, 0xb9, 0xed, 0x4c, 0xb1, 0x37, 0x21, 0xae, 0x6c, 0x7c, 0xdb, 0xda, 0x17, 0xa1, 0x53,
	0x19, 0x19, 0xa8, 0x40, 0xf7, 0x67, 0x03, 0xea, 0x1b, 0xcd, 0x4b, 0x31, 0x16, 0xd4, 0x36, 0x4c,
	0x64, 0x5a, 0xcf, 0xbf, 0xbb, 0xa7, 0x77, 0x66, 0x2f, 0xeb, 0x21, 0xfb, 0xcf, 0xda, 0xfe, 0x34,
	0x60, 0x3f, 0x33, 0x6c, 0x52, 0xd9, 0x6b, 0x38, 0x98, 0xc9, 0x39, 0xb4, 0x85, 0xe1, 0x76, 0x18,
	0x05, 0xb5, 0xbc, 0xa7, 0xf9, 0x93, 0x7f, 0xc3, 0xd8, 0x5a, 0x75, 0xc5, 0x71, 0x34, 0x72, 0x68,
	0x32, 0xcb, 0xf7, 0xa1, 0xa0, 0xb4, 0x69, 0x4d, 0xfa, 0x0d, 0xbd, 0x81, 0x07, 0x51, 0x19, 0x7b,
	0x4c, 0x3d, 0x3c, 0xa3, 0x3f, 0x90, 0xcc, 0xb8, 0x3d, 0xcb, 0x9d, 0x83, 0x88, 0xf4, 0x73, 0x0d,
	0x57, 0x03, 0x77, 0x10, 0xde, 0xb4, 0xdc, 0x9d, 0x42, 0xf1, 0x95, 0xba, 0x93, 0xd0, 0x11, 0x94,
	0x63, 0xdb, 0x74, 0x47, 0x99, 0xcb, 0x44, 0xdf, 0x5d, 0x89, 0xe5, 0xda, 0xec, 0x24, 0x0b, 0x35,
	0xa1, 0xc4, 0xfc, 0x31, 0xff, 0x1e, 0x87, 0x44, 0xf6, 0x51, 0xb6, 0xe2, 0xf7, 0xee, 0x1f, 0x05,
	0xb8, 0x77, 0x2e, 0x4c, 0x41, 0x9f, 0x42, 0x51, 0x73, 0xe9, 0x32, 0x0f, 0xf3, 0xc6, 0x69, 0x51,
	0xba, 0x44, 0x84, 0x47, 0xcf, 0xa0, 0xe4, 0x4c, 0x31, 0xf5, 0x6c, 0xaa, 0x36, 0xaf, 0xdc, 0xaf,
	0xac, 0xaf, 0xda, 0xc5, 0x81, 0x58, 0x1b, 0x1e, 0x5b, 0x45, 0x19, 0x1c, 0xba, 0xe8, 0x29, 0xec,
	0x52, 0x8f, 0x72, 0x8a, 0x67, 0x7a, 0xcb, 0x1b, 0xbb, 0xd2, 0xd6, 0xaa, 0x5e, 0x55, 0xbb, 0x8d,
	0x3e, 0x06, 0xb9, 0xf7, 0xca, 0xd0, 0x08, 0xb9, 0x2d, 0x91, 0x7b, 0x22, 0x20, 0x3d, 0xd2, 0x58,
	0x0b, 0xaa, 0x29, 0x2c, 0x75, 0x1b, 0x77, 0xf3, 0xda, 0xd5, 0x99, 0x94, 0x59, 0xc3, 0xe3, 0x7e,
	0x5d, 0x68, 0x5f, 0x5f, 0xb5, 0x2b, 0x67, 0x11, 0xd5, 0xf0, 0xd8, 0xaa, 0xc4, 0xbc, 0x43, 0x17,
	0x9d, 0xc1, 0x5e, 0x8a, 0x53, 0xdc, 0xfb, 0x8d, 0x7b, 0x92, 0xb5, 0x69, 0xaa, 0x8f, 0x82, 0x19,
	0x7d, 0x14, 0xcc, 0x8b, 0xe8, 0xa3, 0xd0, 0x2f, 0x09, 0xda, 0xb7, 0xbf, 0xb5, 0x0d, 0xab, 0x1a,
	0x73, 0x89, 0x28, 0xfa, 0x02, 0xf6, 0x3c, 0xb2, 0xe2, 0x76, 0x3c, 0x95, 0xac, 0x51, 0xb8, 0xd5,
	0x1c, 0xef, 0x8a, 0xb4, 0xe4, 0x4a, 0x40, 0x2f, 0x01, 0x52, 0x1c, 0xc5, 0x5b, 0x71, 0xa4, 0x32,
	0x84, 0x10, 0xd9, 0x56, 0x8a, 0xa4, 0x74, 0x3b, 0x21, 0x22, 0x2d, 0x25, 0x64, 0x00, 0xad, 0xf4,
	0xd8, 0x26, 0x7c, 0xf1, 0x04, 0x97, 0xe5, 0x66, 0x3d, 0x4a, 0x26, 0x38, 0xc9, 0xd6, 0xb3, 0x7c,
	0xe3, 0x7d, 0x02, 0x1f, 0x78, 0x9f, 0x7c, 0x05, 0x4f, 0x32, 0xf7, 0xc9, 0x06, 0x7f, 0x2c, 0xaf,
	0x22, 0xe5, 0x75, 0x52, 0x17, 0x4c, 0x96, 0x28, 0xd2, 0x18, 0x1d, 0xc4, 0x50, 0x7e, 0xa5, 0x99,
	0x3d, 0xc5, 0x6c, 0xda, 0xd8, 0xe9, 0x18, 0x87, 0x3b, 0xea, 0x20, 0xaa, 0xaf, 0x37, 0x3b, 0xc5,
	0x6c, 0x8a, 0x1e, 0x42, 0x09, 0x07, 0x81, 0x82, 0x54, 0x25, 0xa4, 0x88, 0x83, 0x40, 0x84, 0xfa,
	0x5f, 0xbe, 0x5b, 0xb7, 0x8c, 0xf7, 0xeb, 0x96, 0xf1, 0xfb, 0xba, 0x65, 0xbc, 0xbd, 0x6e, 0x6d,
	0xbd, 0xbf, 0x6e, 0x6d, 0xfd, 0x72, 0xdd, 0xda, 0xfa, 0xf6, 0xc5, 0x84, 0xf2, 0xe9, 0x62, 0x64,
	0x3a, 0xfe, 0xbc, 0xe7, 0xf8, 0x73, 0xc2, 0x47, 0x63, 0x9e, 0x3c, 0xa8, 0xff, 0xa5, 0xcd, 0x3f,
	0xad, 0x51, 0x41, 0xae, 0xbf, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xec, 0x26, 0xcf, 0x93, 0x84,
	0x09, 0x00, 0x00,
}

func (m *HeritageIfaceReplies) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *HeritageIfaceReplies) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *HeritageIfaceReplies) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.InitiateLedger != nil {
		{
			extent, err := m.InitiateLedger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.TerminateLedger != nil {
		{
			extent, err := m.TerminateLedger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.DispatchTrans) > 0 {
		for idxNdExc := len(m.DispatchTrans) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.DispatchTrans[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyInitiateLedger) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyInitiateLedger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyInitiateLedger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Incidents) > 0 {
		for idxNdExc := len(m.Incidents) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Incidents[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyTerminateLedger) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyTerminateLedger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyTerminateLedger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Incidents) > 0 {
		for idxNdExc := len(m.Incidents) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Incidents[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x1a
		}
	}
	if m.AgreementArgumentRevisions != nil {
		{
			extent, err := m.AgreementArgumentRevisions.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.AssessorRevisions) > 0 {
		for idxNdExc := len(m.AssessorRevisions) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.AssessorRevisions[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *AssessorsDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AssessorsDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AssessorsDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.FinalAltitudeAltered != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FinalAltitudeAltered))
		i--
		deltaLocatedAN[i] = 0x10
	}
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
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *AgreementParametersDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AgreementParametersDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AgreementParametersDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.FinalAltitudeAltered != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FinalAltitudeAltered))
		i--
		deltaLocatedAN[i] = 0x10
	}
	{
		extent, err := m.AgreementSettings.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *IfaceRepliesDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *IfaceRepliesDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *IfaceRepliesDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.ReplyCulminateLedger != nil {
		{
			extent, err := m.ReplyCulminateLedger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.HeritageIfaceReplies != nil {
		{
			extent, err := m.HeritageIfaceReplies.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *Edition) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Edition) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Edition) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Package) > 0 {
		i -= len(m.Package)
		copy(deltaLocatedAN[i:], m.Package)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Package)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	{
		extent, err := m.Agreement.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *Status) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Status) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Status) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.PrimaryAltitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.PrimaryAltitude))
		i--
		deltaLocatedAN[i] = 0x70
	}
	if len(m.PlatformDigest) > 0 {
		i -= len(m.PlatformDigest)
		copy(deltaLocatedAN[i:], m.PlatformDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.PlatformDigest)))
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
	if m.FinalAltitudeAgreementParametersAltered != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FinalAltitudeAgreementParametersAltered))
		i--
		deltaLocatedAN[i] = 0x58
	}
	{
		extent, err := m.AgreementSettings.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x52
	if m.FinalAltitudeAssessorsAltered != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FinalAltitudeAssessorsAltered))
		i--
		deltaLocatedAN[i] = 0x48
	}
	if m.FinalAssessors != nil {
		{
			extent, err := m.FinalAssessors.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x42
	}
	if m.Assessors != nil {
		{
			extent, err := m.Assessors.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x3a
	}
	if m.FollowingAssessors != nil {
		{
			extent, err := m.FollowingAssessors.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x32
	}
	n13, fault13 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.FinalLedgerMoment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FinalLedgerMoment):])
	if fault13 != nil {
		return 0, fault13
	}
	i -= n13
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n13))
	i--
	deltaLocatedAN[i] = 0x2a
	{
		extent, err := m.FinalLedgerUUID.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x22
	if m.FinalLedgerAltitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FinalLedgerAltitude))
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
func (m *HeritageIfaceReplies) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DispatchTrans) > 0 {
		for _, e := range m.DispatchTrans {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	if m.TerminateLedger != nil {
		l = m.TerminateLedger.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.InitiateLedger != nil {
		l = m.InitiateLedger.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyInitiateLedger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Incidents) > 0 {
		for _, e := range m.Incidents {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ReplyTerminateLedger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AssessorRevisions) > 0 {
		for _, e := range m.AssessorRevisions {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	if m.AgreementArgumentRevisions != nil {
		l = m.AgreementArgumentRevisions.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if len(m.Incidents) > 0 {
		for _, e := range m.Incidents {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *AssessorsDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssessorAssign != nil {
		l = m.AssessorAssign.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FinalAltitudeAltered != 0 {
		n += 1 + sovKinds(uint64(m.FinalAltitudeAltered))
	}
	return n
}

func (m *AgreementParametersDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AgreementSettings.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.FinalAltitudeAltered != 0 {
		n += 1 + sovKinds(uint64(m.FinalAltitudeAltered))
	}
	return n
}

func (m *IfaceRepliesDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HeritageIfaceReplies != nil {
		l = m.HeritageIfaceReplies.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.ReplyCulminateLedger != nil {
		l = m.ReplyCulminateLedger.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Edition) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Agreement.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Package)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Status) Extent() (n int) {
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
	if m.FinalLedgerAltitude != 0 {
		n += 1 + sovKinds(uint64(m.FinalLedgerAltitude))
	}
	l = m.FinalLedgerUUID.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FinalLedgerMoment)
	n += 1 + l + sovKinds(uint64(l))
	if m.FollowingAssessors != nil {
		l = m.FollowingAssessors.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Assessors != nil {
		l = m.Assessors.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FinalAssessors != nil {
		l = m.FinalAssessors.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FinalAltitudeAssessorsAltered != 0 {
		n += 1 + sovKinds(uint64(m.FinalAltitudeAssessorsAltered))
	}
	l = m.AgreementSettings.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.FinalAltitudeAgreementParametersAltered != 0 {
		n += 1 + sovKinds(uint64(m.FinalAltitudeAgreementParametersAltered))
	}
	l = len(m.FinalOutcomesDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.PlatformDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.PrimaryAltitude != 0 {
		n += 1 + sovKinds(uint64(m.PrimaryAltitude))
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeritageIfaceReplies) Decode(deltaLocatedAN []byte) error {
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
			m.DispatchTrans = append(m.DispatchTrans, &kinds.InvokeTransferOutcome{})
			if err := m.DispatchTrans[len(m.DispatchTrans)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.TerminateLedger == nil {
				m.TerminateLedger = &ReplyTerminateLedger{}
			}
			if err := m.TerminateLedger.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
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
			if m.InitiateLedger == nil {
				m.InitiateLedger = &ReplyInitiateLedger{}
			}
			if err := m.InitiateLedger.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *ReplyInitiateLedger) Decode(deltaLocatedAN []byte) error {
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
			m.Incidents = append(m.Incidents, kinds.Incident{})
			if err := m.Incidents[len(m.Incidents)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *ReplyTerminateLedger) Decode(deltaLocatedAN []byte) error {
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
			m.AssessorRevisions = append(m.AssessorRevisions, kinds.AssessorRevise{})
			if err := m.AssessorRevisions[len(m.AssessorRevisions)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.AgreementArgumentRevisions == nil {
				m.AgreementArgumentRevisions = &kinds1.AgreementSettings{}
			}
			if err := m.AgreementArgumentRevisions.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
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
			m.Incidents = append(m.Incidents, kinds.Incident{})
			if err := m.Incidents[len(m.Incidents)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *AssessorsDetails) Decode(deltaLocatedAN []byte) error {
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
			if m.AssessorAssign == nil {
				m.AssessorAssign = &kinds1.AssessorAssign{}
			}
			if err := m.AssessorAssign.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalAltitudeAltered = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FinalAltitudeAltered |= int64(b&0x7F) << relocate
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
func (m *AgreementParametersDetails) Decode(deltaLocatedAN []byte) error {
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
			if err := m.AgreementSettings.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalAltitudeAltered = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FinalAltitudeAltered |= int64(b&0x7F) << relocate
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
func (m *IfaceRepliesDetails) Decode(deltaLocatedAN []byte) error {
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
			if m.HeritageIfaceReplies == nil {
				m.HeritageIfaceReplies = &HeritageIfaceReplies{}
			}
			if err := m.HeritageIfaceReplies.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
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
			if m.ReplyCulminateLedger == nil {
				m.ReplyCulminateLedger = &kinds.ReplyCulminateLedger{}
			}
			if err := m.ReplyCulminateLedger.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *Edition) Decode(deltaLocatedAN []byte) error {
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
			if err := m.Agreement.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Package = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *Status) Decode(deltaLocatedAN []byte) error {
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
			m.FinalLedgerAltitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FinalLedgerAltitude |= int64(b&0x7F) << relocate
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
			if err := m.FinalLedgerUUID.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.FinalLedgerMoment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.FollowingAssessors == nil {
				m.FollowingAssessors = &kinds1.AssessorAssign{}
			}
			if err := m.FollowingAssessors.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
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
			if m.Assessors == nil {
				m.Assessors = &kinds1.AssessorAssign{}
			}
			if err := m.Assessors.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
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
			if m.FinalAssessors == nil {
				m.FinalAssessors = &kinds1.AssessorAssign{}
			}
			if err := m.FinalAssessors.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 9:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalAltitudeAssessorsAltered = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FinalAltitudeAssessorsAltered |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 10:
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
			if err := m.AgreementSettings.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 11:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalAltitudeAgreementParametersAltered = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FinalAltitudeAgreementParametersAltered |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
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
			m.PlatformDigest = append(m.PlatformDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.PlatformDigest == nil {
				m.PlatformDigest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 14:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PrimaryAltitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.PrimaryAltitude |= int64(b&0x7F) << relocate
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
