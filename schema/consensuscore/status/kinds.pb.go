//
//

package status

import (
	fmt "fmt"
	kinds "github.com/valkyrieworks/iface/kinds"
	kinds1 "github.com/valkyrieworks/schema/consensuscore/kinds"
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
//
//
//
type PastIfaceReplies struct {
	DispatchTrans []*kinds.InvokeTransferOutcome `protobuf:"octets,1,rep,name=deliver_txs,json=deliverTxs,proto3" json:"dispatch_trans,omitempty"`
	TerminateLedger   *AnswerTerminateLedger     `protobuf:"octets,2,opt,name=end_block,json=endBlock,proto3" json:"terminate_ledger,omitempty"`
	InitiateLedger *AnswerInitiateLedger   `protobuf:"octets,3,opt,name=begin_block,json=beginBlock,proto3" json:"begin_ledger,omitempty"`
}

func (m *PastIfaceReplies) Restore()         { *m = PastIfaceReplies{} }
func (m *PastIfaceReplies) String() string { return proto.CompactTextString(m) }
func (*PastIfaceReplies) SchemaSignal()    {}
func (*PastIfaceReplies) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{0}
}
func (m *PastIfaceReplies) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PastIfaceReplies) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Legacyifaceoutcomes.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PastIfaceReplies) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Legacyifaceoutcomes.Merge(m, src)
}
func (m *PastIfaceReplies) XXX_Volume() int {
	return m.Volume()
}
func (m *PastIfaceReplies) XXX_Omitunclear() {
	xxx_messagedata_Legacyifaceoutcomes.DiscardUnknown(m)
}

var xxx_messagedata_Legacyifaceoutcomes proto.InternalMessageInfo

func (m *PastIfaceReplies) FetchDispatchTrans() []*kinds.InvokeTransferOutcome {
	if m != nil {
		return m.DispatchTrans
	}
	return nil
}

func (m *PastIfaceReplies) FetchTerminateLedger() *AnswerTerminateLedger {
	if m != nil {
		return m.TerminateLedger
	}
	return nil
}

func (m *PastIfaceReplies) FetchInitiateLedger() *AnswerInitiateLedger {
	if m != nil {
		return m.InitiateLedger
	}
	return nil
}

//
type AnswerInitiateLedger struct {
	Events []kinds.Event `protobuf:"octets,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *AnswerInitiateLedger) Restore()         { *m = AnswerInitiateLedger{} }
func (m *AnswerInitiateLedger) String() string { return proto.CompactTextString(m) }
func (*AnswerInitiateLedger) SchemaSignal()    {}
func (*AnswerInitiateLedger) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{1}
}
func (m *AnswerInitiateLedger) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AnswerInitiateLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Answerinitiateblock.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnswerInitiateLedger) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Answerinitiateblock.Merge(m, src)
}
func (m *AnswerInitiateLedger) XXX_Volume() int {
	return m.Volume()
}
func (m *AnswerInitiateLedger) XXX_Omitunclear() {
	xxx_messagedata_Answerinitiateblock.DiscardUnknown(m)
}

var xxx_messagedata_Answerinitiateblock proto.InternalMessageInfo

func (m *AnswerInitiateLedger) FetchEvents() []kinds.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

//
type AnswerTerminateLedger struct {
	RatifierRefreshes      []kinds.RatifierModify `protobuf:"octets,1,rep,name=validator_updates,json=validatorUpdates,proto3" json:"ratifier_refreshes"`
	AgreementArgumentRefreshes *kinds1.AgreementOptions `protobuf:"octets,2,opt,name=consensus_param_updates,json=consensusParamUpdates,proto3" json:"agreement_argument_refreshes,omitempty"`
	Events                []kinds.Event           `protobuf:"octets,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *AnswerTerminateLedger) Restore()         { *m = AnswerTerminateLedger{} }
func (m *AnswerTerminateLedger) String() string { return proto.CompactTextString(m) }
func (*AnswerTerminateLedger) SchemaSignal()    {}
func (*AnswerTerminateLedger) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{2}
}
func (m *AnswerTerminateLedger) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AnswerTerminateLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Answerterminateblock.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnswerTerminateLedger) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Answerterminateblock.Merge(m, src)
}
func (m *AnswerTerminateLedger) XXX_Volume() int {
	return m.Volume()
}
func (m *AnswerTerminateLedger) XXX_Omitunclear() {
	xxx_messagedata_Answerterminateblock.DiscardUnknown(m)
}

var xxx_messagedata_Answerterminateblock proto.InternalMessageInfo

func (m *AnswerTerminateLedger) FetchRatifierRefreshes() []kinds.RatifierModify {
	if m != nil {
		return m.RatifierRefreshes
	}
	return nil
}

func (m *AnswerTerminateLedger) FetchAgreementArgumentRefreshes() *kinds1.AgreementOptions {
	if m != nil {
		return m.AgreementArgumentRefreshes
	}
	return nil
}

func (m *AnswerTerminateLedger) FetchEvents() []kinds.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

//
type RatifiersDetails struct {
	RatifierAssign      *kinds1.RatifierAssign `protobuf:"octets,1,opt,name=validator_set,json=validatorSet,proto3" json:"ratifier_collection,omitempty"`
	FinalLevelModified int64                `protobuf:"variableint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"final_level_modified,omitempty"`
}

func (m *RatifiersDetails) Restore()         { *m = RatifiersDetails{} }
func (m *RatifiersDetails) String() string { return proto.CompactTextString(m) }
func (*RatifiersDetails) SchemaSignal()    {}
func (*RatifiersDetails) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{3}
}
func (m *RatifiersDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *RatifiersDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ratifiersdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RatifiersDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ratifiersdetails.Merge(m, src)
}
func (m *RatifiersDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *RatifiersDetails) XXX_Omitunclear() {
	xxx_messagedata_Ratifiersdetails.DiscardUnknown(m)
}

var xxx_messagedata_Ratifiersdetails proto.InternalMessageInfo

func (m *RatifiersDetails) FetchRatifierCollection() *kinds1.RatifierAssign {
	if m != nil {
		return m.RatifierAssign
	}
	return nil
}

func (m *RatifiersDetails) FetchFinalLevelModified() int64 {
	if m != nil {
		return m.FinalLevelModified
	}
	return 0
}

//
type AgreementOptionsDetails struct {
	AgreementOptions   kinds1.AgreementOptions `protobuf:"octets,1,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_options"`
	FinalLevelModified int64                  `protobuf:"variableint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"final_level_modified,omitempty"`
}

func (m *AgreementOptionsDetails) Restore()         { *m = AgreementOptionsDetails{} }
func (m *AgreementOptionsDetails) String() string { return proto.CompactTextString(m) }
func (*AgreementOptionsDetails) SchemaSignal()    {}
func (*AgreementOptionsDetails) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{4}
}
func (m *AgreementOptionsDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AgreementOptionsDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Consensusoptionsdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgreementOptionsDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Consensusoptionsdetails.Merge(m, src)
}
func (m *AgreementOptionsDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *AgreementOptionsDetails) XXX_Omitunclear() {
	xxx_messagedata_Consensusoptionsdetails.DiscardUnknown(m)
}

var xxx_messagedata_Consensusoptionsdetails proto.InternalMessageInfo

func (m *AgreementOptionsDetails) FetchAgreementOptions() kinds1.AgreementOptions {
	if m != nil {
		return m.AgreementOptions
	}
	return kinds1.AgreementOptions{}
}

func (m *AgreementOptionsDetails) FetchFinalLevelModified() int64 {
	if m != nil {
		return m.FinalLevelModified
	}
	return 0
}

type IfaceRepliesDetails struct {
	PastIfaceReplies   *PastIfaceReplies         `protobuf:"octets,1,opt,name=legacy_abci_responses,json=legacyAbciResponses,proto3" json:"past_iface_replies,omitempty"`
	Level                int64                        `protobuf:"variableint,2,opt,name=height,proto3" json:"level,omitempty"`
	ReplyCompleteLedger *kinds.ReplyCompleteLedger `protobuf:"octets,3,opt,name=response_finalize_block,json=responseFinalizeBlock,proto3" json:"reply_complete_ledger,omitempty"`
}

func (m *IfaceRepliesDetails) Restore()         { *m = IfaceRepliesDetails{} }
func (m *IfaceRepliesDetails) String() string { return proto.CompactTextString(m) }
func (*IfaceRepliesDetails) SchemaSignal()    {}
func (*IfaceRepliesDetails) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{5}
}
func (m *IfaceRepliesDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *IfaceRepliesDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ifaceoutcomesdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IfaceRepliesDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ifaceoutcomesdetails.Merge(m, src)
}
func (m *IfaceRepliesDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *IfaceRepliesDetails) XXX_Omitunclear() {
	xxx_messagedata_Ifaceoutcomesdetails.DiscardUnknown(m)
}

var xxx_messagedata_Ifaceoutcomesdetails proto.InternalMessageInfo

func (m *IfaceRepliesDetails) FetchPastIfaceReplies() *PastIfaceReplies {
	if m != nil {
		return m.PastIfaceReplies
	}
	return nil
}

func (m *IfaceRepliesDetails) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *IfaceRepliesDetails) FetchAnswerCompleteLedger() *kinds.ReplyCompleteLedger {
	if m != nil {
		return m.ReplyCompleteLedger
	}
	return nil
}

type Release struct {
	Agreement release.Agreement `protobuf:"octets,1,opt,name=consensus,proto3" json:"agreement"`
	Software  string            `protobuf:"octets,2,opt,name=software,proto3" json:"solution,omitempty"`
}

func (m *Release) Restore()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) SchemaSignal()    {}
func (*Release) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{6}
}
func (m *Release) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Release) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Release.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Release) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Release.Merge(m, src)
}
func (m *Release) XXX_Volume() int {
	return m.Volume()
}
func (m *Release) XXX_Omitunclear() {
	xxx_messagedata_Release.DiscardUnknown(m)
}

var xxx_messagedata_Release proto.InternalMessageInfo

func (m *Release) FetchAgreement() release.Agreement {
	if m != nil {
		return m.Agreement
	}
	return release.Agreement{}
}

func (m *Release) FetchSolution() string {
	if m != nil {
		return m.Software
	}
	return "REDACTED"
}

type Status struct {
	Release Release `protobuf:"octets,1,opt,name=version,proto3" json:"release"`
	//
	LedgerUID       string `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
	PrimaryLevel int64  `protobuf:"variableint,14,opt,name=initial_height,json=initialHeight,proto3" json:"primary_level,omitempty"`
	//
	FinalLedgerLevel int64          `protobuf:"variableint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"final_ledger_level,omitempty"`
	FinalLedgerUID     kinds1.LedgerUID `protobuf:"octets,4,opt,name=last_block_id,json=lastBlockId,proto3" json:"final_ledger_uid"`
	FinalLedgerTime   time.Time      `protobuf:"octets,5,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"final_ledger_time"`
	//
	//
	//
	//
	//
	//
	FollowingRatifiers              *kinds1.RatifierAssign `protobuf:"octets,6,opt,name=next_validators,json=nextValidators,proto3" json:"following_ratifiers,omitempty"`
	Ratifiers                  *kinds1.RatifierAssign `protobuf:"octets,7,opt,name=validators,proto3" json:"ratifiers,omitempty"`
	FinalRatifiers              *kinds1.RatifierAssign `protobuf:"octets,8,opt,name=last_validators,json=lastValidators,proto3" json:"final_ratifiers,omitempty"`
	FinalLevelRatifiersModified int64                `protobuf:"variableint,9,opt,name=last_height_validators_changed,json=lastHeightValidatorsChanged,proto3" json:"final_level_ratifiers_modified,omitempty"`
	//
	//
	AgreementOptions                  kinds1.AgreementOptions `protobuf:"octets,10,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_options"`
	FinalLevelAgreementOptionsModified int64                  `protobuf:"variableint,11,opt,name=last_height_consensus_params_changed,json=lastHeightConsensusParamsChanged,proto3" json:"final_level_agreement_options_modified,omitempty"`
	//
	FinalOutcomesDigest []byte `protobuf:"octets,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"final_outcomes_digest,omitempty"`
	//
	ApplicationDigest []byte `protobuf:"octets,13,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *Status) Restore()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) SchemaSignal()    {}
func (*Status) Definition() ([]byte, []int) {
	return filedefinition_hash7, []int{7}
}
func (m *Status) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Status) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Status.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Status.Merge(m, src)
}
func (m *Status) XXX_Volume() int {
	return m.Volume()
}
func (m *Status) XXX_Omitunclear() {
	xxx_messagedata_Status.DiscardUnknown(m)
}

var xxx_messagedata_Status proto.InternalMessageInfo

func (m *Status) FetchRelease() Release {
	if m != nil {
		return m.Release
	}
	return Release{}
}

func (m *Status) FetchSeriesUID() string {
	if m != nil {
		return m.LedgerUID
	}
	return "REDACTED"
}

func (m *Status) FetchPrimaryLevel() int64 {
	if m != nil {
		return m.PrimaryLevel
	}
	return 0
}

func (m *Status) FetchFinalLedgerLevel() int64 {
	if m != nil {
		return m.FinalLedgerLevel
	}
	return 0
}

func (m *Status) FetchFinalLedgerUID() kinds1.LedgerUID {
	if m != nil {
		return m.FinalLedgerUID
	}
	return kinds1.LedgerUID{}
}

func (m *Status) FetchFinalLedgerTime() time.Time {
	if m != nil {
		return m.FinalLedgerTime
	}
	return time.Time{}
}

func (m *Status) FetchFollowingRatifiers() *kinds1.RatifierAssign {
	if m != nil {
		return m.FollowingRatifiers
	}
	return nil
}

func (m *Status) FetchRatifiers() *kinds1.RatifierAssign {
	if m != nil {
		return m.Ratifiers
	}
	return nil
}

func (m *Status) FetchFinalRatifiers() *kinds1.RatifierAssign {
	if m != nil {
		return m.FinalRatifiers
	}
	return nil
}

func (m *Status) FetchFinalLevelRatifiersModified() int64 {
	if m != nil {
		return m.FinalLevelRatifiersModified
	}
	return 0
}

func (m *Status) FetchAgreementOptions() kinds1.AgreementOptions {
	if m != nil {
		return m.AgreementOptions
	}
	return kinds1.AgreementOptions{}
}

func (m *Status) FetchFinalLevelAgreementOptionsModified() int64 {
	if m != nil {
		return m.FinalLevelAgreementOptionsModified
	}
	return 0
}

func (m *Status) FetchFinalOutcomesDigest() []byte {
	if m != nil {
		return m.FinalOutcomesDigest
	}
	return nil
}

func (m *Status) FetchApplicationDigest() []byte {
	if m != nil {
		return m.ApplicationDigest
	}
	return nil
}

func init() {
	proto.RegisterType((*PastIfaceReplies)(nil), "REDACTED")
	proto.RegisterType((*AnswerInitiateLedger)(nil), "REDACTED")
	proto.RegisterType((*AnswerTerminateLedger)(nil), "REDACTED")
	proto.RegisterType((*RatifiersDetails)(nil), "REDACTED")
	proto.RegisterType((*AgreementOptionsDetails)(nil), "REDACTED")
	proto.RegisterType((*IfaceRepliesDetails)(nil), "REDACTED")
	proto.RegisterType((*Release)(nil), "REDACTED")
	proto.RegisterType((*Status)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash7) }

var filedefinition_hash7 = []byte{
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

func (m *PastIfaceReplies) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PastIfaceReplies) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PastIfaceReplies) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitiateLedger != nil {
		{
			volume, err := m.InitiateLedger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TerminateLedger != nil {
		{
			volume, err := m.TerminateLedger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DispatchTrans) > 0 {
		for idxNdEx := len(m.DispatchTrans) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.DispatchTrans[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AnswerInitiateLedger) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnswerInitiateLedger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AnswerInitiateLedger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for idxNdEx := len(m.Events) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Events[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AnswerTerminateLedger) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnswerTerminateLedger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AnswerTerminateLedger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for idxNdEx := len(m.Events) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Events[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.AgreementArgumentRefreshes != nil {
		{
			volume, err := m.AgreementArgumentRefreshes.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RatifierRefreshes) > 0 {
		for idxNdEx := len(m.RatifierRefreshes) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.RatifierRefreshes[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RatifiersDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RatifiersDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *RatifiersDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalLevelModified != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FinalLevelModified))
		i--
		dAtA[i] = 0x10
	}
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AgreementOptionsDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgreementOptionsDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AgreementOptionsDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalLevelModified != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FinalLevelModified))
		i--
		dAtA[i] = 0x10
	}
	{
		volume, err := m.AgreementOptions.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *IfaceRepliesDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IfaceRepliesDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *IfaceRepliesDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReplyCompleteLedger != nil {
		{
			volume, err := m.ReplyCompleteLedger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	if m.PastIfaceReplies != nil {
		{
			volume, err := m.PastIfaceReplies.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *Release) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Release) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Release) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Software) > 0 {
		i -= len(m.Software)
		copy(dAtA[i:], m.Software)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Software)))
		i--
		dAtA[i] = 0x12
	}
	{
		volume, err := m.Agreement.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *Status) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Status) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PrimaryLevel != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.PrimaryLevel))
		i--
		dAtA[i] = 0x70
	}
	if len(m.ApplicationDigest) > 0 {
		i -= len(m.ApplicationDigest)
		copy(dAtA[i:], m.ApplicationDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ApplicationDigest)))
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
	if m.FinalLevelAgreementOptionsModified != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FinalLevelAgreementOptionsModified))
		i--
		dAtA[i] = 0x58
	}
	{
		volume, err := m.AgreementOptions.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x52
	if m.FinalLevelRatifiersModified != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FinalLevelRatifiersModified))
		i--
		dAtA[i] = 0x48
	}
	if m.FinalRatifiers != nil {
		{
			volume, err := m.FinalRatifiers.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Ratifiers != nil {
		{
			volume, err := m.Ratifiers.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.FollowingRatifiers != nil {
		{
			volume, err := m.FollowingRatifiers.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x32
	}
	n13, fault13 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.FinalLedgerTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FinalLedgerTime):])
	if fault13 != nil {
		return 0, fault13
	}
	i -= n13
	i = formatVariableintKinds(dAtA, i, uint64(n13))
	i--
	dAtA[i] = 0x2a
	{
		volume, err := m.FinalLedgerUID.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x22
	if m.FinalLedgerLevel != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FinalLedgerLevel))
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
func (m *PastIfaceReplies) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DispatchTrans) > 0 {
		for _, e := range m.DispatchTrans {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	if m.TerminateLedger != nil {
		l = m.TerminateLedger.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.InitiateLedger != nil {
		l = m.InitiateLedger.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AnswerInitiateLedger) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *AnswerTerminateLedger) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RatifierRefreshes) > 0 {
		for _, e := range m.RatifierRefreshes {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	if m.AgreementArgumentRefreshes != nil {
		l = m.AgreementArgumentRefreshes.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *RatifiersDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RatifierAssign != nil {
		l = m.RatifierAssign.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FinalLevelModified != 0 {
		n += 1 + sovKinds(uint64(m.FinalLevelModified))
	}
	return n
}

func (m *AgreementOptionsDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AgreementOptions.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.FinalLevelModified != 0 {
		n += 1 + sovKinds(uint64(m.FinalLevelModified))
	}
	return n
}

func (m *IfaceRepliesDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PastIfaceReplies != nil {
		l = m.PastIfaceReplies.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.ReplyCompleteLedger != nil {
		l = m.ReplyCompleteLedger.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Release) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Agreement.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.Software)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Status) Volume() (n int) {
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
	if m.FinalLedgerLevel != 0 {
		n += 1 + sovKinds(uint64(m.FinalLedgerLevel))
	}
	l = m.FinalLedgerUID.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FinalLedgerTime)
	n += 1 + l + sovKinds(uint64(l))
	if m.FollowingRatifiers != nil {
		l = m.FollowingRatifiers.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Ratifiers != nil {
		l = m.Ratifiers.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FinalRatifiers != nil {
		l = m.FinalRatifiers.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FinalLevelRatifiersModified != 0 {
		n += 1 + sovKinds(uint64(m.FinalLevelRatifiersModified))
	}
	l = m.AgreementOptions.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.FinalLevelAgreementOptionsModified != 0 {
		n += 1 + sovKinds(uint64(m.FinalLevelAgreementOptionsModified))
	}
	l = len(m.FinalOutcomesDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.ApplicationDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.PrimaryLevel != 0 {
		n += 1 + sovKinds(uint64(m.PrimaryLevel))
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PastIfaceReplies) Unserialize(dAtA []byte) error {
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
			m.DispatchTrans = append(m.DispatchTrans, &kinds.InvokeTransferOutcome{})
			if err := m.DispatchTrans[len(m.DispatchTrans)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.TerminateLedger == nil {
				m.TerminateLedger = &AnswerTerminateLedger{}
			}
			if err := m.TerminateLedger.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
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
			if m.InitiateLedger == nil {
				m.InitiateLedger = &AnswerInitiateLedger{}
			}
			if err := m.InitiateLedger.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *AnswerInitiateLedger) Unserialize(dAtA []byte) error {
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
			m.Events = append(m.Events, kinds.Event{})
			if err := m.Events[len(m.Events)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *AnswerTerminateLedger) Unserialize(dAtA []byte) error {
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
			m.RatifierRefreshes = append(m.RatifierRefreshes, kinds.RatifierModify{})
			if err := m.RatifierRefreshes[len(m.RatifierRefreshes)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.AgreementArgumentRefreshes == nil {
				m.AgreementArgumentRefreshes = &kinds1.AgreementOptions{}
			}
			if err := m.AgreementArgumentRefreshes.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
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
			m.Events = append(m.Events, kinds.Event{})
			if err := m.Events[len(m.Events)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *RatifiersDetails) Unserialize(dAtA []byte) error {
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
			if m.RatifierAssign == nil {
				m.RatifierAssign = &kinds1.RatifierAssign{}
			}
			if err := m.RatifierAssign.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalLevelModified = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FinalLevelModified |= int64(b&0x7F) << displace
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
func (m *AgreementOptionsDetails) Unserialize(dAtA []byte) error {
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
			if err := m.AgreementOptions.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalLevelModified = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FinalLevelModified |= int64(b&0x7F) << displace
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
func (m *IfaceRepliesDetails) Unserialize(dAtA []byte) error {
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
			if m.PastIfaceReplies == nil {
				m.PastIfaceReplies = &PastIfaceReplies{}
			}
			if err := m.PastIfaceReplies.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
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
			if m.ReplyCompleteLedger == nil {
				m.ReplyCompleteLedger = &kinds.ReplyCompleteLedger{}
			}
			if err := m.ReplyCompleteLedger.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *Release) Unserialize(dAtA []byte) error {
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
			if err := m.Agreement.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Software = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *Status) Unserialize(dAtA []byte) error {
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
			m.FinalLedgerLevel = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FinalLedgerLevel |= int64(b&0x7F) << displace
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
			if err := m.FinalLedgerUID.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.FinalLedgerTime, dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.FollowingRatifiers == nil {
				m.FollowingRatifiers = &kinds1.RatifierAssign{}
			}
			if err := m.FollowingRatifiers.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
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
			if m.Ratifiers == nil {
				m.Ratifiers = &kinds1.RatifierAssign{}
			}
			if err := m.Ratifiers.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
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
			if m.FinalRatifiers == nil {
				m.FinalRatifiers = &kinds1.RatifierAssign{}
			}
			if err := m.FinalRatifiers.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 9:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalLevelRatifiersModified = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FinalLevelRatifiersModified |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 10:
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
			if err := m.AgreementOptions.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 11:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FinalLevelAgreementOptionsModified = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FinalLevelAgreementOptionsModified |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
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
			m.ApplicationDigest = append(m.ApplicationDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.ApplicationDigest == nil {
				m.ApplicationDigest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 14:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PrimaryLevel = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.PrimaryLevel |= int64(b&0x7F) << displace
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
