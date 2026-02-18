//
//

package kinds

import (
	context "context"
	fmt "fmt"
	vault "github.com/valkyrieworks/schema/consensuscore/vault"
	kinds1 "github.com/valkyrieworks/schema/consensuscore/kinds"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InspectTransferKind int32

const (
	Transfercheckkind_New     InspectTransferKind = 0
	Transfercheckkind_Revalidate InspectTransferKind = 1
)

var Transfercheckkind_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
}

var Transfercheckkind_item = map[string]int32{
	"REDACTED":     0,
	"REDACTED": 1,
}

func (x InspectTransferKind) String() string {
	return proto.EnumName(Transfercheckkind_label, int32(x))
}

func (InspectTransferKind) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{0}
}

type MalpracticeKind int32

const (
	Misconductkind_UNCLEAR             MalpracticeKind = 0
	Misconductkind_REPLICATED_BALLOT      MalpracticeKind = 1
	Misconductkind_RAPID_CUSTOMER_ASSAULT MalpracticeKind = 2
)

var Misconductkind_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
}

var Misconductkind_item = map[string]int32{
	"REDACTED":             0,
	"REDACTED":      1,
	"REDACTED": 2,
}

func (x MalpracticeKind) String() string {
	return proto.EnumName(Misconductkind_label, int32(x))
}

func (MalpracticeKind) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{1}
}

type Replymirrorsnapshot_Outcome int32

const (
	Replymirrorsnapshot_UNCLEAR       Replymirrorsnapshot_Outcome = 0
	Replymirrorsnapshot_ALLOW        Replymirrorsnapshot_Outcome = 1
	Replymirrorsnapshot_CANCEL         Replymirrorsnapshot_Outcome = 2
	Replymirrorsnapshot_DECLINE        Replymirrorsnapshot_Outcome = 3
	Replymirrorsnapshot_DECLINE_LAYOUT Replymirrorsnapshot_Outcome = 4
	Replymirrorsnapshot_DECLINE_EMITTER Replymirrorsnapshot_Outcome = 5
)

var Replymirrorsnapshot_Outcome_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
	4: "REDACTED",
	5: "REDACTED",
}

var Replymirrorsnapshot_Outcome_item = map[string]int32{
	"REDACTED":       0,
	"REDACTED":        1,
	"REDACTED":         2,
	"REDACTED":        3,
	"REDACTED": 4,
	"REDACTED": 5,
}

func (x Replymirrorsnapshot_Outcome) String() string {
	return proto.EnumName(Replymirrorsnapshot_Outcome_label, int32(x))
}

func (Replymirrorsnapshot_Outcome) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{31, 0}
}

type Replyexecutemirrorsegment_Outcome int32

const (
	Replyexecutemirrorsegment_UNCLEAR         Replyexecutemirrorsegment_Outcome = 0
	Replyexecutemirrorsegment_ALLOW          Replyexecutemirrorsegment_Outcome = 1
	Replyexecutemirrorsegment_CANCEL           Replyexecutemirrorsegment_Outcome = 2
	Replyexecutemirrorsegment_REPROCESS           Replyexecutemirrorsegment_Outcome = 3
	Replyexecutemirrorsegment_REPROCESS_MIRROR  Replyexecutemirrorsegment_Outcome = 4
	Replyexecutemirrorsegment_DECLINE_MIRROR Replyexecutemirrorsegment_Outcome = 5
)

var Replyexecutemirrorsegment_Outcome_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
	4: "REDACTED",
	5: "REDACTED",
}

var Replyexecutemirrorsegment_Outcome_item = map[string]int32{
	"REDACTED":         0,
	"REDACTED":          1,
	"REDACTED":           2,
	"REDACTED":           3,
	"REDACTED":  4,
	"REDACTED": 5,
}

func (x Replyexecutemirrorsegment_Outcome) String() string {
	return proto.EnumName(Replyexecutemirrorsegment_Outcome_label, int32(x))
}

func (Replyexecutemirrorsegment_Outcome) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{33, 0}
}

type Responseprocessnomination_Nominationstate int32

const (
	Responseprocessnomination_UNCLEAR Responseprocessnomination_Nominationstate = 0
	Responseprocessnomination_ALLOW  Responseprocessnomination_Nominationstate = 1
	Responseprocessnomination_DECLINE  Responseprocessnomination_Nominationstate = 2
)

var Responseprocessnomination_Nominationstate_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
}

var Responseprocessnomination_Nominationstate_item = map[string]int32{
	"REDACTED": 0,
	"REDACTED":  1,
	"REDACTED":  2,
}

func (x Responseprocessnomination_Nominationstate) String() string {
	return proto.EnumName(Responseprocessnomination_Nominationstate_label, int32(x))
}

func (Responseprocessnomination_Nominationstate) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{35, 0}
}

type Responseverifyballotextension_Validatestatus int32

const (
	Responseverifyballotextension_UNCLEAR Responseverifyballotextension_Validatestatus = 0
	Responseverifyballotextension_ALLOW  Responseverifyballotextension_Validatestatus = 1
	//
	//
	//
	//
	Responseverifyballotextension_DECLINE Responseverifyballotextension_Validatestatus = 2
)

var Responseverifyballotextension_Validatestatus_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
}

var Responseverifyballotextension_Validatestatus_item = map[string]int32{
	"REDACTED": 0,
	"REDACTED":  1,
	"REDACTED":  2,
}

func (x Responseverifyballotextension_Validatestatus) String() string {
	return proto.EnumName(Responseverifyballotextension_Validatestatus_label, int32(x))
}

func (Responseverifyballotextension_Validatestatus) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{37, 0}
}

type Query struct {
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
	//
	//
	//
	//
	//
	//
	//
	//
	//
	Item isquery_Item `protobuf_oneof:"item"`
}

func (m *Query) Restore()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) SchemaSignal()    {}
func (*Query) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{0}
}
func (m *Query) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Query) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Query.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Query) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Query.Merge(m, src)
}
func (m *Query) XXX_Volume() int {
	return m.Volume()
}
func (m *Query) XXX_Omitunclear() {
	xxx_messagedata_Query.DiscardUnknown(m)
}

var xxx_messagedata_Query proto.InternalMessageInfo

type isquery_Item interface {
	isquery_Item()
	SerializeTo([]byte) (int, error)
	Volume() int
}

type Query_Reverberate struct {
	Replicate *QueryReverberate `protobuf:"octets,1,opt,name=echo,proto3,oneof" json:"reverberate,omitempty"`
}
type Query_Purge struct {
	Purge *QueryPurge `protobuf:"octets,2,opt,name=flush,proto3,oneof" json:"purge,omitempty"`
}
type Query_Details struct {
	Details *QueryDetails `protobuf:"octets,3,opt,name=info,proto3,oneof" json:"details,omitempty"`
}
type Query_Initiatechain struct {
	InitSeries *QueryInitSeries `protobuf:"octets,5,opt,name=init_chain,json=initChain,proto3,oneof" json:"init_series,omitempty"`
}
type Query_Inquire struct {
	Inquire *QueryInquire `protobuf:"octets,6,opt,name=query,proto3,oneof" json:"inquire,omitempty"`
}
type Query_Transfercheck struct {
	InspectTransfer *QueryInspectTransfer `protobuf:"octets,8,opt,name=check_tx,json=checkTx,proto3,oneof" json:"inspect_transfer,omitempty"`
}
type Query_Endorse struct {
	Endorse *QueryEndorse `protobuf:"octets,11,opt,name=commit,proto3,oneof" json:"endorse,omitempty"`
}
type Query_Catalogmirrors struct {
	CatalogMirrors *QueryCatalogMirrors `protobuf:"octets,12,opt,name=list_snapshots,json=listSnapshots,proto3,oneof" json:"catalog_mirrors,omitempty"`
}
type Query_Mirrorsnapshot struct {
	ProposalMirror *QueryProposalMirror `protobuf:"octets,13,opt,name=offer_snapshot,json=offerSnapshot,proto3,oneof" json:"proposal_mirror,omitempty"`
}
type Query_Loadmirrorsegment struct {
	ImportMirrorSegment *QueryImportMirrorSegment `protobuf:"octets,14,opt,name=load_snapshot_chunk,json=loadSnapshotChunk,proto3,oneof" json:"import_mirror_segment,omitempty"`
}
type Query_Executemirrorsegment struct {
	ExecuteMirrorSegment *QueryExecuteMirrorSegment `protobuf:"octets,15,opt,name=apply_snapshot_chunk,json=applySnapshotChunk,proto3,oneof" json:"execute_mirror_segment,omitempty"`
}
type Query_Arrangenomination struct {
	ArrangeNomination *QueryArrangeNomination `protobuf:"octets,16,opt,name=prepare_proposal,json=prepareProposal,proto3,oneof" json:"arrange_nomination,omitempty"`
}
type Query_Processnomination struct {
	HandleNomination *QueryHandleNomination `protobuf:"octets,17,opt,name=process_proposal,json=processProposal,proto3,oneof" json:"handle_nomination,omitempty"`
}
type Query_Ballotextend struct {
	ExpandBallot *QueryExpandBallot `protobuf:"octets,18,opt,name=extend_vote,json=extendVote,proto3,oneof" json:"expand_ballot,omitempty"`
}
type Query_Validateballotextension struct {
	ValidateBallotAddition *QueryValidateBallotAddition `protobuf:"octets,19,opt,name=verify_vote_extension,json=verifyVoteExtension,proto3,oneof" json:"validate_ballot_addition,omitempty"`
}
type Query_Terminateblock struct {
	CompleteLedger *QueryCompleteLedger `protobuf:"octets,20,opt,name=finalize_block,json=finalizeBlock,proto3,oneof" json:"complete_ledger,omitempty"`
}
type Query_Transferinsert struct {
	EmbedTransfer *QueryEmbedTransfer `protobuf:"octets,21,opt,name=insert_tx,json=insertTx,proto3,oneof" json:"append_transfer,omitempty"`
}
type Query_Reaptransfers struct {
	HarvestTrans *QueryHarvestTrans `protobuf:"octets,22,opt,name=reap_txs,json=reapTxs,proto3,oneof" json:"harvest_trans,omitempty"`
}

func (*Query_Reverberate) isquery_Item()                {}
func (*Query_Purge) isquery_Item()               {}
func (*Query_Details) isquery_Item()                {}
func (*Query_Initiatechain) isquery_Item()           {}
func (*Query_Inquire) isquery_Item()               {}
func (*Query_Transfercheck) isquery_Item()             {}
func (*Query_Endorse) isquery_Item()              {}
func (*Query_Catalogmirrors) isquery_Item()       {}
func (*Query_Mirrorsnapshot) isquery_Item()       {}
func (*Query_Loadmirrorsegment) isquery_Item()   {}
func (*Query_Executemirrorsegment) isquery_Item()  {}
func (*Query_Arrangenomination) isquery_Item()     {}
func (*Query_Processnomination) isquery_Item()     {}
func (*Query_Ballotextend) isquery_Item()          {}
func (*Query_Validateballotextension) isquery_Item() {}
func (*Query_Terminateblock) isquery_Item()       {}
func (*Query_Transferinsert) isquery_Item()            {}
func (*Query_Reaptransfers) isquery_Item()             {}

func (m *Query) FetchItem() isquery_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Query) FetchReverberate() *QueryReverberate {
	if x, ok := m.FetchItem().(*Query_Reverberate); ok {
		return x.Replicate
	}
	return nil
}

func (m *Query) FetchPurge() *QueryPurge {
	if x, ok := m.FetchItem().(*Query_Purge); ok {
		return x.Purge
	}
	return nil
}

func (m *Query) FetchDetails() *QueryDetails {
	if x, ok := m.FetchItem().(*Query_Details); ok {
		return x.Details
	}
	return nil
}

func (m *Query) FetchInitSeries() *QueryInitSeries {
	if x, ok := m.FetchItem().(*Query_Initiatechain); ok {
		return x.InitSeries
	}
	return nil
}

func (m *Query) FetchInquire() *QueryInquire {
	if x, ok := m.FetchItem().(*Query_Inquire); ok {
		return x.Inquire
	}
	return nil
}

func (m *Query) FetchInspectTransfer() *QueryInspectTransfer {
	if x, ok := m.FetchItem().(*Query_Transfercheck); ok {
		return x.InspectTransfer
	}
	return nil
}

func (m *Query) FetchEndorse() *QueryEndorse {
	if x, ok := m.FetchItem().(*Query_Endorse); ok {
		return x.Endorse
	}
	return nil
}

func (m *Query) FetchCatalogMirrors() *QueryCatalogMirrors {
	if x, ok := m.FetchItem().(*Query_Catalogmirrors); ok {
		return x.CatalogMirrors
	}
	return nil
}

func (m *Query) FetchProposalMirror() *QueryProposalMirror {
	if x, ok := m.FetchItem().(*Query_Mirrorsnapshot); ok {
		return x.ProposalMirror
	}
	return nil
}

func (m *Query) FetchImportMirrorSegment() *QueryImportMirrorSegment {
	if x, ok := m.FetchItem().(*Query_Loadmirrorsegment); ok {
		return x.ImportMirrorSegment
	}
	return nil
}

func (m *Query) FetchExecuteMirrorSegment() *QueryExecuteMirrorSegment {
	if x, ok := m.FetchItem().(*Query_Executemirrorsegment); ok {
		return x.ExecuteMirrorSegment
	}
	return nil
}

func (m *Query) FetchArrangeNomination() *QueryArrangeNomination {
	if x, ok := m.FetchItem().(*Query_Arrangenomination); ok {
		return x.ArrangeNomination
	}
	return nil
}

func (m *Query) FetchHandleNomination() *QueryHandleNomination {
	if x, ok := m.FetchItem().(*Query_Processnomination); ok {
		return x.HandleNomination
	}
	return nil
}

func (m *Query) FetchExpandBallot() *QueryExpandBallot {
	if x, ok := m.FetchItem().(*Query_Ballotextend); ok {
		return x.ExpandBallot
	}
	return nil
}

func (m *Query) FetchValidateBallotAddition() *QueryValidateBallotAddition {
	if x, ok := m.FetchItem().(*Query_Validateballotextension); ok {
		return x.ValidateBallotAddition
	}
	return nil
}

func (m *Query) FetchCompleteLedger() *QueryCompleteLedger {
	if x, ok := m.FetchItem().(*Query_Terminateblock); ok {
		return x.CompleteLedger
	}
	return nil
}

func (m *Query) FetchEmbedTransfer() *QueryEmbedTransfer {
	if x, ok := m.FetchItem().(*Query_Transferinsert); ok {
		return x.EmbedTransfer
	}
	return nil
}

func (m *Query) FetchHarvestTrans() *QueryHarvestTrans {
	if x, ok := m.FetchItem().(*Query_Reaptransfers); ok {
		return x.HarvestTrans
	}
	return nil
}

//
func (*Query) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Query_Reverberate)(nil),
		(*Query_Purge)(nil),
		(*Query_Details)(nil),
		(*Query_Initiatechain)(nil),
		(*Query_Inquire)(nil),
		(*Query_Transfercheck)(nil),
		(*Query_Endorse)(nil),
		(*Query_Catalogmirrors)(nil),
		(*Query_Mirrorsnapshot)(nil),
		(*Query_Loadmirrorsegment)(nil),
		(*Query_Executemirrorsegment)(nil),
		(*Query_Arrangenomination)(nil),
		(*Query_Processnomination)(nil),
		(*Query_Ballotextend)(nil),
		(*Query_Validateballotextension)(nil),
		(*Query_Terminateblock)(nil),
		(*Query_Transferinsert)(nil),
		(*Query_Reaptransfers)(nil),
	}
}

type QueryReverberate struct {
	Signal string `protobuf:"octets,1,opt,name=message,proto3" json:"signal,omitempty"`
}

func (m *QueryReverberate) Restore()         { *m = QueryReverberate{} }
func (m *QueryReverberate) String() string { return proto.CompactTextString(m) }
func (*QueryReverberate) SchemaSignal()    {}
func (*QueryReverberate) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{1}
}
func (m *QueryReverberate) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryReverberate) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryecho.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryReverberate) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryecho.Merge(m, src)
}
func (m *QueryReverberate) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryReverberate) XXX_Omitunclear() {
	xxx_messagedata_Queryecho.DiscardUnknown(m)
}

var xxx_messagedata_Queryecho proto.InternalMessageInfo

func (m *QueryReverberate) FetchSignal() string {
	if m != nil {
		return m.Signal
	}
	return "REDACTED"
}

type QueryPurge struct {
}

func (m *QueryPurge) Restore()         { *m = QueryPurge{} }
func (m *QueryPurge) String() string { return proto.CompactTextString(m) }
func (*QueryPurge) SchemaSignal()    {}
func (*QueryPurge) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{2}
}
func (m *QueryPurge) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryPurge) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querypurge.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPurge) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querypurge.Merge(m, src)
}
func (m *QueryPurge) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryPurge) XXX_Omitunclear() {
	xxx_messagedata_Querypurge.DiscardUnknown(m)
}

var xxx_messagedata_Querypurge proto.InternalMessageInfo

type QueryDetails struct {
	Release      string `protobuf:"octets,1,opt,name=version,proto3" json:"release,omitempty"`
	LedgerRelease uint64 `protobuf:"variableint,2,opt,name=block_version,json=blockVersion,proto3" json:"ledger_release,omitempty"`
	P2PRelease   uint64 `protobuf:"variableint,3,opt,name=p2p_version,json=p2pVersion,proto3" json:"p2p_release,omitempty"`
	IfaceRelease  string `protobuf:"octets,4,opt,name=abci_version,json=abciVersion,proto3" json:"iface_release,omitempty"`
}

func (m *QueryDetails) Restore()         { *m = QueryDetails{} }
func (m *QueryDetails) String() string { return proto.CompactTextString(m) }
func (*QueryDetails) SchemaSignal()    {}
func (*QueryDetails) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{3}
}
func (m *QueryDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querydata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querydata.Merge(m, src)
}
func (m *QueryDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryDetails) XXX_Omitunclear() {
	xxx_messagedata_Querydata.DiscardUnknown(m)
}

var xxx_messagedata_Querydata proto.InternalMessageInfo

func (m *QueryDetails) FetchRelease() string {
	if m != nil {
		return m.Release
	}
	return "REDACTED"
}

func (m *QueryDetails) FetchLedgerRelease() uint64 {
	if m != nil {
		return m.LedgerRelease
	}
	return 0
}

func (m *QueryDetails) FetchP2PRelease() uint64 {
	if m != nil {
		return m.P2PRelease
	}
	return 0
}

func (m *QueryDetails) FetchIfaceRelease() string {
	if m != nil {
		return m.IfaceRelease
	}
	return "REDACTED"
}

type QueryInitSeries struct {
	Time            time.Time               `protobuf:"octets,1,opt,name=time,proto3,stdtime" json:"moment"`
	SeriesUid         string                  `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
	AgreementOptions *kinds1.AgreementOptions `protobuf:"octets,3,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_options,omitempty"`
	Ratifiers      []RatifierModify       `protobuf:"octets,4,rep,name=validators,proto3" json:"ratifiers"`
	ApplicationStatusOctets   []byte                  `protobuf:"octets,5,opt,name=app_state_bytes,json=appStateBytes,proto3" json:"application_status_octets,omitempty"`
	PrimaryLevel   int64                   `protobuf:"variableint,6,opt,name=initial_height,json=initialHeight,proto3" json:"primary_level,omitempty"`
}

func (m *QueryInitSeries) Restore()         { *m = QueryInitSeries{} }
func (m *QueryInitSeries) String() string { return proto.CompactTextString(m) }
func (*QueryInitSeries) SchemaSignal()    {}
func (*QueryInitSeries) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{4}
}
func (m *QueryInitSeries) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryInitSeries) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryinitiatechain.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInitSeries) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryinitiatechain.Merge(m, src)
}
func (m *QueryInitSeries) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryInitSeries) XXX_Omitunclear() {
	xxx_messagedata_Queryinitiatechain.DiscardUnknown(m)
}

var xxx_messagedata_Queryinitiatechain proto.InternalMessageInfo

func (m *QueryInitSeries) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *QueryInitSeries) FetchSeriesUid() string {
	if m != nil {
		return m.SeriesUid
	}
	return "REDACTED"
}

func (m *QueryInitSeries) FetchAgreementOptions() *kinds1.AgreementOptions {
	if m != nil {
		return m.AgreementOptions
	}
	return nil
}

func (m *QueryInitSeries) FetchRatifiers() []RatifierModify {
	if m != nil {
		return m.Ratifiers
	}
	return nil
}

func (m *QueryInitSeries) FetchApplicationStatusOctets() []byte {
	if m != nil {
		return m.ApplicationStatusOctets
	}
	return nil
}

func (m *QueryInitSeries) FetchPrimaryLevel() int64 {
	if m != nil {
		return m.PrimaryLevel
	}
	return 0
}

type QueryInquire struct {
	Data   []byte `protobuf:"octets,1,opt,name=data,proto3" json:"data,omitempty"`
	Route   string `protobuf:"octets,2,opt,name=path,proto3" json:"route,omitempty"`
	Level int64  `protobuf:"variableint,3,opt,name=height,proto3" json:"level,omitempty"`
	Demonstrate  bool   `protobuf:"variableint,4,opt,name=prove,proto3" json:"demonstrate,omitempty"`
}

func (m *QueryInquire) Restore()         { *m = QueryInquire{} }
func (m *QueryInquire) String() string { return proto.CompactTextString(m) }
func (*QueryInquire) SchemaSignal()    {}
func (*QueryInquire) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{5}
}
func (m *QueryInquire) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryInquire) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryquery.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInquire) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryquery.Merge(m, src)
}
func (m *QueryInquire) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryInquire) XXX_Omitunclear() {
	xxx_messagedata_Queryquery.DiscardUnknown(m)
}

var xxx_messagedata_Queryquery proto.InternalMessageInfo

func (m *QueryInquire) FetchData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryInquire) FetchRoute() string {
	if m != nil {
		return m.Route
	}
	return "REDACTED"
}

func (m *QueryInquire) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *QueryInquire) FetchDemonstrate() bool {
	if m != nil {
		return m.Demonstrate
	}
	return false
}

type QueryInspectTransfer struct {
	Tx   []byte      `protobuf:"octets,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Kind InspectTransferKind `protobuf:"variableint,2,opt,name=type,proto3,enum=tendermint.abci.CheckTxType" json:"kind,omitempty"`
}

func (m *QueryInspectTransfer) Restore()         { *m = QueryInspectTransfer{} }
func (m *QueryInspectTransfer) String() string { return proto.CompactTextString(m) }
func (*QueryInspectTransfer) SchemaSignal()    {}
func (*QueryInspectTransfer) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{6}
}
func (m *QueryInspectTransfer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryInspectTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querytransfercheck.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInspectTransfer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querytransfercheck.Merge(m, src)
}
func (m *QueryInspectTransfer) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryInspectTransfer) XXX_Omitunclear() {
	xxx_messagedata_Querytransfercheck.DiscardUnknown(m)
}

var xxx_messagedata_Querytransfercheck proto.InternalMessageInfo

func (m *QueryInspectTransfer) FetchTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *QueryInspectTransfer) FetchKind() InspectTransferKind {
	if m != nil {
		return m.Kind
	}
	return Transfercheckkind_New
}

type QueryEmbedTransfer struct {
	Tx []byte `protobuf:"octets,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *QueryEmbedTransfer) Restore()         { *m = QueryEmbedTransfer{} }
func (m *QueryEmbedTransfer) String() string { return proto.CompactTextString(m) }
func (*QueryEmbedTransfer) SchemaSignal()    {}
func (*QueryEmbedTransfer) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{7}
}
func (m *QueryEmbedTransfer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryEmbedTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querytransferinsert.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEmbedTransfer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querytransferinsert.Merge(m, src)
}
func (m *QueryEmbedTransfer) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryEmbedTransfer) XXX_Omitunclear() {
	xxx_messagedata_Querytransferinsert.DiscardUnknown(m)
}

var xxx_messagedata_Querytransferinsert proto.InternalMessageInfo

func (m *QueryEmbedTransfer) FetchTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type QueryHarvestTrans struct {
	MaximumOctets uint64 `protobuf:"variableint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"maximum_octets,omitempty"`
	MaximumFuel   uint64 `protobuf:"variableint,2,opt,name=max_gas,json=maxGas,proto3" json:"maximum_fuel,omitempty"`
}

func (m *QueryHarvestTrans) Restore()         { *m = QueryHarvestTrans{} }
func (m *QueryHarvestTrans) String() string { return proto.CompactTextString(m) }
func (*QueryHarvestTrans) SchemaSignal()    {}
func (*QueryHarvestTrans) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{8}
}
func (m *QueryHarvestTrans) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryHarvestTrans) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryreaptransfers.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHarvestTrans) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryreaptransfers.Merge(m, src)
}
func (m *QueryHarvestTrans) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryHarvestTrans) XXX_Omitunclear() {
	xxx_messagedata_Queryreaptransfers.DiscardUnknown(m)
}

var xxx_messagedata_Queryreaptransfers proto.InternalMessageInfo

func (m *QueryHarvestTrans) FetchMaximumOctets() uint64 {
	if m != nil {
		return m.MaximumOctets
	}
	return 0
}

func (m *QueryHarvestTrans) FetchMaximumFuel() uint64 {
	if m != nil {
		return m.MaximumFuel
	}
	return 0
}

type QueryEndorse struct {
}

func (m *QueryEndorse) Restore()         { *m = QueryEndorse{} }
func (m *QueryEndorse) String() string { return proto.CompactTextString(m) }
func (*QueryEndorse) SchemaSignal()    {}
func (*QueryEndorse) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{9}
}
func (m *QueryEndorse) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryEndorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryendorse.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEndorse) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryendorse.Merge(m, src)
}
func (m *QueryEndorse) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryEndorse) XXX_Omitunclear() {
	xxx_messagedata_Queryendorse.DiscardUnknown(m)
}

var xxx_messagedata_Queryendorse proto.InternalMessageInfo

//
type QueryCatalogMirrors struct {
}

func (m *QueryCatalogMirrors) Restore()         { *m = QueryCatalogMirrors{} }
func (m *QueryCatalogMirrors) String() string { return proto.CompactTextString(m) }
func (*QueryCatalogMirrors) SchemaSignal()    {}
func (*QueryCatalogMirrors) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{10}
}
func (m *QueryCatalogMirrors) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryCatalogMirrors) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querycatalogmirrors.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCatalogMirrors) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querycatalogmirrors.Merge(m, src)
}
func (m *QueryCatalogMirrors) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryCatalogMirrors) XXX_Omitunclear() {
	xxx_messagedata_Querycatalogmirrors.DiscardUnknown(m)
}

var xxx_messagedata_Querycatalogmirrors proto.InternalMessageInfo

//
type QueryProposalMirror struct {
	Mirror *Mirror `protobuf:"octets,1,opt,name=snapshot,proto3" json:"mirror,omitempty"`
	ApplicationDigest  []byte    `protobuf:"octets,2,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *QueryProposalMirror) Restore()         { *m = QueryProposalMirror{} }
func (m *QueryProposalMirror) String() string { return proto.CompactTextString(m) }
func (*QueryProposalMirror) SchemaSignal()    {}
func (*QueryProposalMirror) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{11}
}
func (m *QueryProposalMirror) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryProposalMirror) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querymirrorsnapshot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProposalMirror) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querymirrorsnapshot.Merge(m, src)
}
func (m *QueryProposalMirror) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryProposalMirror) XXX_Omitunclear() {
	xxx_messagedata_Querymirrorsnapshot.DiscardUnknown(m)
}

var xxx_messagedata_Querymirrorsnapshot proto.InternalMessageInfo

func (m *QueryProposalMirror) FetchMirror() *Mirror {
	if m != nil {
		return m.Mirror
	}
	return nil
}

func (m *QueryProposalMirror) FetchApplicationDigest() []byte {
	if m != nil {
		return m.ApplicationDigest
	}
	return nil
}

//
type QueryImportMirrorSegment struct {
	Level uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Layout uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Segment  uint32 `protobuf:"variableint,3,opt,name=chunk,proto3" json:"segment,omitempty"`
}

func (m *QueryImportMirrorSegment) Restore()         { *m = QueryImportMirrorSegment{} }
func (m *QueryImportMirrorSegment) String() string { return proto.CompactTextString(m) }
func (*QueryImportMirrorSegment) SchemaSignal()    {}
func (*QueryImportMirrorSegment) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{12}
}
func (m *QueryImportMirrorSegment) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryImportMirrorSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryloadmirrorsegment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryImportMirrorSegment) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryloadmirrorsegment.Merge(m, src)
}
func (m *QueryImportMirrorSegment) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryImportMirrorSegment) XXX_Omitunclear() {
	xxx_messagedata_Queryloadmirrorsegment.DiscardUnknown(m)
}

var xxx_messagedata_Queryloadmirrorsegment proto.InternalMessageInfo

func (m *QueryImportMirrorSegment) FetchLevel() uint64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *QueryImportMirrorSegment) FetchLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *QueryImportMirrorSegment) FetchSegment() uint32 {
	if m != nil {
		return m.Segment
	}
	return 0
}

//
type QueryExecuteMirrorSegment struct {
	Ordinal  uint32 `protobuf:"variableint,1,opt,name=index,proto3" json:"ordinal,omitempty"`
	Segment  []byte `protobuf:"octets,2,opt,name=chunk,proto3" json:"segment,omitempty"`
	Emitter string `protobuf:"octets,3,opt,name=sender,proto3" json:"emitter,omitempty"`
}

func (m *QueryExecuteMirrorSegment) Restore()         { *m = QueryExecuteMirrorSegment{} }
func (m *QueryExecuteMirrorSegment) String() string { return proto.CompactTextString(m) }
func (*QueryExecuteMirrorSegment) SchemaSignal()    {}
func (*QueryExecuteMirrorSegment) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{13}
}
func (m *QueryExecuteMirrorSegment) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryExecuteMirrorSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryexecutemirrorsegment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryExecuteMirrorSegment) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryexecutemirrorsegment.Merge(m, src)
}
func (m *QueryExecuteMirrorSegment) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryExecuteMirrorSegment) XXX_Omitunclear() {
	xxx_messagedata_Queryexecutemirrorsegment.DiscardUnknown(m)
}

var xxx_messagedata_Queryexecutemirrorsegment proto.InternalMessageInfo

func (m *QueryExecuteMirrorSegment) FetchOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *QueryExecuteMirrorSegment) FetchSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *QueryExecuteMirrorSegment) FetchEmitter() string {
	if m != nil {
		return m.Emitter
	}
	return "REDACTED"
}

type QueryArrangeNomination struct {
	//
	MaximumTransferOctets int64 `protobuf:"variableint,1,opt,name=max_tx_bytes,json=maxTxBytes,proto3" json:"maximum_transfer_octets,omitempty"`
	//
	//
	Txs                [][]byte           `protobuf:"octets,2,rep,name=txs,proto3" json:"txs,omitempty"`
	NativeFinalEndorse    ExpandedEndorseDetails `protobuf:"octets,3,opt,name=local_last_commit,json=localLastCommit,proto3" json:"native_final_endorse"`
	Malpractice        []Malpractice      `protobuf:"octets,4,rep,name=misbehavior,proto3" json:"malpractice"`
	Level             int64              `protobuf:"variableint,5,opt,name=height,proto3" json:"level,omitempty"`
	Time               time.Time          `protobuf:"octets,6,opt,name=time,proto3,stdtime" json:"moment"`
	FollowingRatifiersDigest []byte             `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_ratifiers_digest,omitempty"`
	//
	RecommenderLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"recommender_location,omitempty"`
}

func (m *QueryArrangeNomination) Restore()         { *m = QueryArrangeNomination{} }
func (m *QueryArrangeNomination) String() string { return proto.CompactTextString(m) }
func (*QueryArrangeNomination) SchemaSignal()    {}
func (*QueryArrangeNomination) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{14}
}
func (m *QueryArrangeNomination) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryArrangeNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryarrangenomination.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryArrangeNomination) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryarrangenomination.Merge(m, src)
}
func (m *QueryArrangeNomination) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryArrangeNomination) XXX_Omitunclear() {
	xxx_messagedata_Queryarrangenomination.DiscardUnknown(m)
}

var xxx_messagedata_Queryarrangenomination proto.InternalMessageInfo

func (m *QueryArrangeNomination) FetchMaximumTransferOctets() int64 {
	if m != nil {
		return m.MaximumTransferOctets
	}
	return 0
}

func (m *QueryArrangeNomination) FetchTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *QueryArrangeNomination) FetchNativeFinalEndorse() ExpandedEndorseDetails {
	if m != nil {
		return m.NativeFinalEndorse
	}
	return ExpandedEndorseDetails{}
}

func (m *QueryArrangeNomination) FetchMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *QueryArrangeNomination) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *QueryArrangeNomination) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *QueryArrangeNomination) FetchFollowingRatifiersDigest() []byte {
	if m != nil {
		return m.FollowingRatifiersDigest
	}
	return nil
}

func (m *QueryArrangeNomination) FetchRecommenderLocation() []byte {
	if m != nil {
		return m.RecommenderLocation
	}
	return nil
}

type QueryHandleNomination struct {
	Txs                [][]byte      `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
	NominatedFinalEndorse EndorseDetails    `protobuf:"octets,2,opt,name=proposed_last_commit,json=proposedLastCommit,proto3" json:"nominated_final_endorse"`
	Malpractice        []Malpractice `protobuf:"octets,3,rep,name=misbehavior,proto3" json:"malpractice"`
	//
	Digest               []byte    `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Level             int64     `protobuf:"variableint,5,opt,name=height,proto3" json:"level,omitempty"`
	Time               time.Time `protobuf:"octets,6,opt,name=time,proto3,stdtime" json:"moment"`
	FollowingRatifiersDigest []byte    `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_ratifiers_digest,omitempty"`
	//
	RecommenderLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"recommender_location,omitempty"`
}

func (m *QueryHandleNomination) Restore()         { *m = QueryHandleNomination{} }
func (m *QueryHandleNomination) String() string { return proto.CompactTextString(m) }
func (*QueryHandleNomination) SchemaSignal()    {}
func (*QueryHandleNomination) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{15}
}
func (m *QueryHandleNomination) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryHandleNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querynominationprocess.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHandleNomination) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querynominationprocess.Merge(m, src)
}
func (m *QueryHandleNomination) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryHandleNomination) XXX_Omitunclear() {
	xxx_messagedata_Querynominationprocess.DiscardUnknown(m)
}

var xxx_messagedata_Querynominationprocess proto.InternalMessageInfo

func (m *QueryHandleNomination) FetchTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *QueryHandleNomination) FetchNominatedFinalEndorse() EndorseDetails {
	if m != nil {
		return m.NominatedFinalEndorse
	}
	return EndorseDetails{}
}

func (m *QueryHandleNomination) FetchMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *QueryHandleNomination) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *QueryHandleNomination) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *QueryHandleNomination) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *QueryHandleNomination) FetchFollowingRatifiersDigest() []byte {
	if m != nil {
		return m.FollowingRatifiersDigest
	}
	return nil
}

func (m *QueryHandleNomination) FetchRecommenderLocation() []byte {
	if m != nil {
		return m.RecommenderLocation
	}
	return nil
}

//
type QueryExpandBallot struct {
	//
	Digest []byte `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	//
	Level int64 `protobuf:"variableint,2,opt,name=height,proto3" json:"level,omitempty"`
	//
	Time               time.Time     `protobuf:"octets,3,opt,name=time,proto3,stdtime" json:"moment"`
	Txs                [][]byte      `protobuf:"octets,4,rep,name=txs,proto3" json:"txs,omitempty"`
	NominatedFinalEndorse EndorseDetails    `protobuf:"octets,5,opt,name=proposed_last_commit,json=proposedLastCommit,proto3" json:"nominated_final_endorse"`
	Malpractice        []Malpractice `protobuf:"octets,6,rep,name=misbehavior,proto3" json:"malpractice"`
	FollowingRatifiersDigest []byte        `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_ratifiers_digest,omitempty"`
	//
	RecommenderLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"recommender_location,omitempty"`
}

func (m *QueryExpandBallot) Restore()         { *m = QueryExpandBallot{} }
func (m *QueryExpandBallot) String() string { return proto.CompactTextString(m) }
func (*QueryExpandBallot) SchemaSignal()    {}
func (*QueryExpandBallot) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{16}
}
func (m *QueryExpandBallot) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryExpandBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryballotextend.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryExpandBallot) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryballotextend.Merge(m, src)
}
func (m *QueryExpandBallot) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryExpandBallot) XXX_Omitunclear() {
	xxx_messagedata_Queryballotextend.DiscardUnknown(m)
}

var xxx_messagedata_Queryballotextend proto.InternalMessageInfo

func (m *QueryExpandBallot) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *QueryExpandBallot) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *QueryExpandBallot) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *QueryExpandBallot) FetchTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *QueryExpandBallot) FetchNominatedFinalEndorse() EndorseDetails {
	if m != nil {
		return m.NominatedFinalEndorse
	}
	return EndorseDetails{}
}

func (m *QueryExpandBallot) FetchMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *QueryExpandBallot) FetchFollowingRatifiersDigest() []byte {
	if m != nil {
		return m.FollowingRatifiersDigest
	}
	return nil
}

func (m *QueryExpandBallot) FetchRecommenderLocation() []byte {
	if m != nil {
		return m.RecommenderLocation
	}
	return nil
}

//
type QueryValidateBallotAddition struct {
	//
	Digest []byte `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	//
	RatifierLocation []byte `protobuf:"octets,2,opt,name=validator_address,json=validatorAddress,proto3" json:"ratifier_location,omitempty"`
	Level           int64  `protobuf:"variableint,3,opt,name=height,proto3" json:"level,omitempty"`
	BallotAddition    []byte `protobuf:"octets,4,opt,name=vote_extension,json=voteExtension,proto3" json:"ballot_addition,omitempty"`
}

func (m *QueryValidateBallotAddition) Restore()         { *m = QueryValidateBallotAddition{} }
func (m *QueryValidateBallotAddition) String() string { return proto.CompactTextString(m) }
func (*QueryValidateBallotAddition) SchemaSignal()    {}
func (*QueryValidateBallotAddition) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{17}
}
func (m *QueryValidateBallotAddition) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryValidateBallotAddition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryvalidateballotextension.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidateBallotAddition) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryvalidateballotextension.Merge(m, src)
}
func (m *QueryValidateBallotAddition) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryValidateBallotAddition) XXX_Omitunclear() {
	xxx_messagedata_Queryvalidateballotextension.DiscardUnknown(m)
}

var xxx_messagedata_Queryvalidateballotextension proto.InternalMessageInfo

func (m *QueryValidateBallotAddition) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *QueryValidateBallotAddition) FetchRatifierLocation() []byte {
	if m != nil {
		return m.RatifierLocation
	}
	return nil
}

func (m *QueryValidateBallotAddition) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *QueryValidateBallotAddition) FetchBallotAddition() []byte {
	if m != nil {
		return m.BallotAddition
	}
	return nil
}

type QueryCompleteLedger struct {
	Txs               [][]byte      `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
	ResolvedFinalEndorse EndorseDetails    `protobuf:"octets,2,opt,name=decided_last_commit,json=decidedLastCommit,proto3" json:"resolved_final_endorse"`
	Malpractice       []Malpractice `protobuf:"octets,3,rep,name=misbehavior,proto3" json:"malpractice"`
	//
	Digest               []byte    `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Level             int64     `protobuf:"variableint,5,opt,name=height,proto3" json:"level,omitempty"`
	Time               time.Time `protobuf:"octets,6,opt,name=time,proto3,stdtime" json:"moment"`
	FollowingRatifiersDigest []byte    `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_ratifiers_digest,omitempty"`
	//
	RecommenderLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"recommender_location,omitempty"`
}

func (m *QueryCompleteLedger) Restore()         { *m = QueryCompleteLedger{} }
func (m *QueryCompleteLedger) String() string { return proto.CompactTextString(m) }
func (*QueryCompleteLedger) SchemaSignal()    {}
func (*QueryCompleteLedger) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{18}
}
func (m *QueryCompleteLedger) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryCompleteLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryterminateblock.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCompleteLedger) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryterminateblock.Merge(m, src)
}
func (m *QueryCompleteLedger) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryCompleteLedger) XXX_Omitunclear() {
	xxx_messagedata_Queryterminateblock.DiscardUnknown(m)
}

var xxx_messagedata_Queryterminateblock proto.InternalMessageInfo

func (m *QueryCompleteLedger) FetchTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *QueryCompleteLedger) FetchResolvedFinalEndorse() EndorseDetails {
	if m != nil {
		return m.ResolvedFinalEndorse
	}
	return EndorseDetails{}
}

func (m *QueryCompleteLedger) FetchMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *QueryCompleteLedger) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *QueryCompleteLedger) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *QueryCompleteLedger) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *QueryCompleteLedger) FetchFollowingRatifiersDigest() []byte {
	if m != nil {
		return m.FollowingRatifiersDigest
	}
	return nil
}

func (m *QueryCompleteLedger) FetchRecommenderLocation() []byte {
	if m != nil {
		return m.RecommenderLocation
	}
	return nil
}

type Reply struct {
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
	//
	Item isreply_Item `protobuf_oneof:"item"`
}

func (m *Reply) Restore()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) SchemaSignal()    {}
func (*Reply) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{19}
}
func (m *Reply) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Reply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Reply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reply) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Reply.Merge(m, src)
}
func (m *Reply) XXX_Volume() int {
	return m.Volume()
}
func (m *Reply) XXX_Omitunclear() {
	xxx_messagedata_Reply.DiscardUnknown(m)
}

var xxx_messagedata_Reply proto.InternalMessageInfo

type isreply_Item interface {
	isreply_Item()
	SerializeTo([]byte) (int, error)
	Volume() int
}

type Reply_Exemption struct {
	Exemption *ReplyExemption `protobuf:"octets,1,opt,name=exception,proto3,oneof" json:"exemption,omitempty"`
}
type Reply_Reverberate struct {
	Replicate *ReplyReverberate `protobuf:"octets,2,opt,name=echo,proto3,oneof" json:"reverberate,omitempty"`
}
type Reply_Purge struct {
	Purge *ReplyPurge `protobuf:"octets,3,opt,name=flush,proto3,oneof" json:"purge,omitempty"`
}
type Reply_Details struct {
	Details *ReplyDetails `protobuf:"octets,4,opt,name=info,proto3,oneof" json:"details,omitempty"`
}
type Reply_Initiatechain struct {
	InitSeries *ReplyInitSeries `protobuf:"octets,6,opt,name=init_chain,json=initChain,proto3,oneof" json:"init_series,omitempty"`
}
type Reply_Inquire struct {
	Inquire *ReplyInquire `protobuf:"octets,7,opt,name=query,proto3,oneof" json:"inquire,omitempty"`
}
type Reply_Transfercheck struct {
	InspectTransfer *ReplyInspectTransfer `protobuf:"octets,9,opt,name=check_tx,json=checkTx,proto3,oneof" json:"inspect_transfer,omitempty"`
}
type Reply_Endorse struct {
	Endorse *ReplyEndorse `protobuf:"octets,12,opt,name=commit,proto3,oneof" json:"endorse,omitempty"`
}
type Reply_Catalogmirrors struct {
	CatalogMirrors *ReplyCatalogMirrors `protobuf:"octets,13,opt,name=list_snapshots,json=listSnapshots,proto3,oneof" json:"catalog_mirrors,omitempty"`
}
type Reply_Mirrorsnapshot struct {
	ProposalMirror *ReplyProposalMirror `protobuf:"octets,14,opt,name=offer_snapshot,json=offerSnapshot,proto3,oneof" json:"proposal_mirror,omitempty"`
}
type Reply_Loadmirrorsegment struct {
	ImportMirrorSegment *ReplyImportMirrorSegment `protobuf:"octets,15,opt,name=load_snapshot_chunk,json=loadSnapshotChunk,proto3,oneof" json:"import_mirror_segment,omitempty"`
}
type Reply_Executemirrorsegment struct {
	ExecuteMirrorSegment *ReplyExecuteMirrorSegment `protobuf:"octets,16,opt,name=apply_snapshot_chunk,json=applySnapshotChunk,proto3,oneof" json:"execute_mirror_segment,omitempty"`
}
type Reply_Arrangenomination struct {
	ArrangeNomination *ReplyArrangeNomination `protobuf:"octets,17,opt,name=prepare_proposal,json=prepareProposal,proto3,oneof" json:"arrange_nomination,omitempty"`
}
type Reply_Processnomination struct {
	HandleNomination *ReplyHandleNomination `protobuf:"octets,18,opt,name=process_proposal,json=processProposal,proto3,oneof" json:"handle_nomination,omitempty"`
}
type Reply_Ballotextend struct {
	ExpandBallot *ReplyExpandBallot `protobuf:"octets,19,opt,name=extend_vote,json=extendVote,proto3,oneof" json:"expand_ballot,omitempty"`
}
type Reply_Validateballotextension struct {
	ValidateBallotAddition *ReplyValidateBallotAddition `protobuf:"octets,20,opt,name=verify_vote_extension,json=verifyVoteExtension,proto3,oneof" json:"validate_ballot_addition,omitempty"`
}
type Reply_Terminateblock struct {
	CompleteLedger *ReplyCompleteLedger `protobuf:"octets,21,opt,name=finalize_block,json=finalizeBlock,proto3,oneof" json:"complete_ledger,omitempty"`
}
type Reply_Transferinsert struct {
	EmbedTransfer *ReplyEmbedTransfer `protobuf:"octets,22,opt,name=insert_tx,json=insertTx,proto3,oneof" json:"append_transfer,omitempty"`
}
type Reply_Reaptransfers struct {
	HarvestTrans *ReplyHarvestTrans `protobuf:"octets,23,opt,name=reap_txs,json=reapTxs,proto3,oneof" json:"harvest_trans,omitempty"`
}

func (*Reply_Exemption) isreply_Item()           {}
func (*Reply_Reverberate) isreply_Item()                {}
func (*Reply_Purge) isreply_Item()               {}
func (*Reply_Details) isreply_Item()                {}
func (*Reply_Initiatechain) isreply_Item()           {}
func (*Reply_Inquire) isreply_Item()               {}
func (*Reply_Transfercheck) isreply_Item()             {}
func (*Reply_Endorse) isreply_Item()              {}
func (*Reply_Catalogmirrors) isreply_Item()       {}
func (*Reply_Mirrorsnapshot) isreply_Item()       {}
func (*Reply_Loadmirrorsegment) isreply_Item()   {}
func (*Reply_Executemirrorsegment) isreply_Item()  {}
func (*Reply_Arrangenomination) isreply_Item()     {}
func (*Reply_Processnomination) isreply_Item()     {}
func (*Reply_Ballotextend) isreply_Item()          {}
func (*Reply_Validateballotextension) isreply_Item() {}
func (*Reply_Terminateblock) isreply_Item()       {}
func (*Reply_Transferinsert) isreply_Item()            {}
func (*Reply_Reaptransfers) isreply_Item()             {}

func (m *Reply) FetchItem() isreply_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Reply) FetchExemption() *ReplyExemption {
	if x, ok := m.FetchItem().(*Reply_Exemption); ok {
		return x.Exemption
	}
	return nil
}

func (m *Reply) FetchReverberate() *ReplyReverberate {
	if x, ok := m.FetchItem().(*Reply_Reverberate); ok {
		return x.Replicate
	}
	return nil
}

func (m *Reply) FetchPurge() *ReplyPurge {
	if x, ok := m.FetchItem().(*Reply_Purge); ok {
		return x.Purge
	}
	return nil
}

func (m *Reply) FetchDetails() *ReplyDetails {
	if x, ok := m.FetchItem().(*Reply_Details); ok {
		return x.Details
	}
	return nil
}

func (m *Reply) FetchInitSeries() *ReplyInitSeries {
	if x, ok := m.FetchItem().(*Reply_Initiatechain); ok {
		return x.InitSeries
	}
	return nil
}

func (m *Reply) FetchInquire() *ReplyInquire {
	if x, ok := m.FetchItem().(*Reply_Inquire); ok {
		return x.Inquire
	}
	return nil
}

func (m *Reply) FetchInspectTransfer() *ReplyInspectTransfer {
	if x, ok := m.FetchItem().(*Reply_Transfercheck); ok {
		return x.InspectTransfer
	}
	return nil
}

func (m *Reply) FetchEndorse() *ReplyEndorse {
	if x, ok := m.FetchItem().(*Reply_Endorse); ok {
		return x.Endorse
	}
	return nil
}

func (m *Reply) FetchCatalogMirrors() *ReplyCatalogMirrors {
	if x, ok := m.FetchItem().(*Reply_Catalogmirrors); ok {
		return x.CatalogMirrors
	}
	return nil
}

func (m *Reply) FetchProposalMirror() *ReplyProposalMirror {
	if x, ok := m.FetchItem().(*Reply_Mirrorsnapshot); ok {
		return x.ProposalMirror
	}
	return nil
}

func (m *Reply) FetchImportMirrorSegment() *ReplyImportMirrorSegment {
	if x, ok := m.FetchItem().(*Reply_Loadmirrorsegment); ok {
		return x.ImportMirrorSegment
	}
	return nil
}

func (m *Reply) FetchExecuteMirrorSegment() *ReplyExecuteMirrorSegment {
	if x, ok := m.FetchItem().(*Reply_Executemirrorsegment); ok {
		return x.ExecuteMirrorSegment
	}
	return nil
}

func (m *Reply) FetchArrangeNomination() *ReplyArrangeNomination {
	if x, ok := m.FetchItem().(*Reply_Arrangenomination); ok {
		return x.ArrangeNomination
	}
	return nil
}

func (m *Reply) FetchHandleNomination() *ReplyHandleNomination {
	if x, ok := m.FetchItem().(*Reply_Processnomination); ok {
		return x.HandleNomination
	}
	return nil
}

func (m *Reply) FetchExpandBallot() *ReplyExpandBallot {
	if x, ok := m.FetchItem().(*Reply_Ballotextend); ok {
		return x.ExpandBallot
	}
	return nil
}

func (m *Reply) FetchValidateBallotAddition() *ReplyValidateBallotAddition {
	if x, ok := m.FetchItem().(*Reply_Validateballotextension); ok {
		return x.ValidateBallotAddition
	}
	return nil
}

func (m *Reply) FetchCompleteLedger() *ReplyCompleteLedger {
	if x, ok := m.FetchItem().(*Reply_Terminateblock); ok {
		return x.CompleteLedger
	}
	return nil
}

func (m *Reply) FetchEmbedTransfer() *ReplyEmbedTransfer {
	if x, ok := m.FetchItem().(*Reply_Transferinsert); ok {
		return x.EmbedTransfer
	}
	return nil
}

func (m *Reply) FetchHarvestTrans() *ReplyHarvestTrans {
	if x, ok := m.FetchItem().(*Reply_Reaptransfers); ok {
		return x.HarvestTrans
	}
	return nil
}

//
func (*Reply) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Reply_Exemption)(nil),
		(*Reply_Reverberate)(nil),
		(*Reply_Purge)(nil),
		(*Reply_Details)(nil),
		(*Reply_Initiatechain)(nil),
		(*Reply_Inquire)(nil),
		(*Reply_Transfercheck)(nil),
		(*Reply_Endorse)(nil),
		(*Reply_Catalogmirrors)(nil),
		(*Reply_Mirrorsnapshot)(nil),
		(*Reply_Loadmirrorsegment)(nil),
		(*Reply_Executemirrorsegment)(nil),
		(*Reply_Arrangenomination)(nil),
		(*Reply_Processnomination)(nil),
		(*Reply_Ballotextend)(nil),
		(*Reply_Validateballotextension)(nil),
		(*Reply_Terminateblock)(nil),
		(*Reply_Transferinsert)(nil),
		(*Reply_Reaptransfers)(nil),
	}
}

//
type ReplyExemption struct {
	Fault string `protobuf:"octets,1,opt,name=error,proto3" json:"fault,omitempty"`
}

func (m *ReplyExemption) Restore()         { *m = ReplyExemption{} }
func (m *ReplyExemption) String() string { return proto.CompactTextString(m) }
func (*ReplyExemption) SchemaSignal()    {}
func (*ReplyExemption) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{20}
}
func (m *ReplyExemption) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyExemption) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyfault.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyExemption) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyfault.Merge(m, src)
}
func (m *ReplyExemption) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyExemption) XXX_Omitunclear() {
	xxx_messagedata_Replyfault.DiscardUnknown(m)
}

var xxx_messagedata_Replyfault proto.InternalMessageInfo

func (m *ReplyExemption) FetchFault() string {
	if m != nil {
		return m.Fault
	}
	return "REDACTED"
}

type ReplyReverberate struct {
	Signal string `protobuf:"octets,1,opt,name=message,proto3" json:"signal,omitempty"`
}

func (m *ReplyReverberate) Restore()         { *m = ReplyReverberate{} }
func (m *ReplyReverberate) String() string { return proto.CompactTextString(m) }
func (*ReplyReverberate) SchemaSignal()    {}
func (*ReplyReverberate) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{21}
}
func (m *ReplyReverberate) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyReverberate) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyecho.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyReverberate) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyecho.Merge(m, src)
}
func (m *ReplyReverberate) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyReverberate) XXX_Omitunclear() {
	xxx_messagedata_Replyecho.DiscardUnknown(m)
}

var xxx_messagedata_Replyecho proto.InternalMessageInfo

func (m *ReplyReverberate) FetchSignal() string {
	if m != nil {
		return m.Signal
	}
	return "REDACTED"
}

type ReplyPurge struct {
}

func (m *ReplyPurge) Restore()         { *m = ReplyPurge{} }
func (m *ReplyPurge) String() string { return proto.CompactTextString(m) }
func (*ReplyPurge) SchemaSignal()    {}
func (*ReplyPurge) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{22}
}
func (m *ReplyPurge) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyPurge) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replypurge.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyPurge) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replypurge.Merge(m, src)
}
func (m *ReplyPurge) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyPurge) XXX_Omitunclear() {
	xxx_messagedata_Replypurge.DiscardUnknown(m)
}

var xxx_messagedata_Replypurge proto.InternalMessageInfo

type ReplyDetails struct {
	Data             string `protobuf:"octets,1,opt,name=data,proto3" json:"data,omitempty"`
	Release          string `protobuf:"octets,2,opt,name=version,proto3" json:"release,omitempty"`
	ApplicationRelease       uint64 `protobuf:"variableint,3,opt,name=app_version,json=appVersion,proto3" json:"application_release,omitempty"`
	FinalLedgerLevel  int64  `protobuf:"variableint,4,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"final_ledger_level,omitempty"`
	FinalLedgerApplicationDigest []byte `protobuf:"octets,5,opt,name=last_block_app_hash,json=lastBlockAppHash,proto3" json:"final_ledger_application_digest,omitempty"`
}

func (m *ReplyDetails) Restore()         { *m = ReplyDetails{} }
func (m *ReplyDetails) String() string { return proto.CompactTextString(m) }
func (*ReplyDetails) SchemaSignal()    {}
func (*ReplyDetails) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{23}
}
func (m *ReplyDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replydata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replydata.Merge(m, src)
}
func (m *ReplyDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyDetails) XXX_Omitunclear() {
	xxx_messagedata_Replydata.DiscardUnknown(m)
}

var xxx_messagedata_Replydata proto.InternalMessageInfo

func (m *ReplyDetails) FetchData() string {
	if m != nil {
		return m.Data
	}
	return "REDACTED"
}

func (m *ReplyDetails) FetchRelease() string {
	if m != nil {
		return m.Release
	}
	return "REDACTED"
}

func (m *ReplyDetails) FetchApplicationRelease() uint64 {
	if m != nil {
		return m.ApplicationRelease
	}
	return 0
}

func (m *ReplyDetails) FetchFinalLedgerLevel() int64 {
	if m != nil {
		return m.FinalLedgerLevel
	}
	return 0
}

func (m *ReplyDetails) FetchFinalLedgerApplicationDigest() []byte {
	if m != nil {
		return m.FinalLedgerApplicationDigest
	}
	return nil
}

type ReplyInitSeries struct {
	AgreementOptions *kinds1.AgreementOptions `protobuf:"octets,1,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_options,omitempty"`
	Ratifiers      []RatifierModify       `protobuf:"octets,2,rep,name=validators,proto3" json:"ratifiers"`
	ApplicationDigest         []byte                  `protobuf:"octets,3,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *ReplyInitSeries) Restore()         { *m = ReplyInitSeries{} }
func (m *ReplyInitSeries) String() string { return proto.CompactTextString(m) }
func (*ReplyInitSeries) SchemaSignal()    {}
func (*ReplyInitSeries) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{24}
}
func (m *ReplyInitSeries) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyInitSeries) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyinitiatechain.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyInitSeries) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyinitiatechain.Merge(m, src)
}
func (m *ReplyInitSeries) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyInitSeries) XXX_Omitunclear() {
	xxx_messagedata_Replyinitiatechain.DiscardUnknown(m)
}

var xxx_messagedata_Replyinitiatechain proto.InternalMessageInfo

func (m *ReplyInitSeries) FetchAgreementOptions() *kinds1.AgreementOptions {
	if m != nil {
		return m.AgreementOptions
	}
	return nil
}

func (m *ReplyInitSeries) FetchRatifiers() []RatifierModify {
	if m != nil {
		return m.Ratifiers
	}
	return nil
}

func (m *ReplyInitSeries) FetchApplicationDigest() []byte {
	if m != nil {
		return m.ApplicationDigest
	}
	return nil
}

type ReplyInquire struct {
	Code uint32 `protobuf:"variableint,1,opt,name=code,proto3" json:"code,omitempty"`
	//
	Log       string           `protobuf:"octets,3,opt,name=log,proto3" json:"log,omitempty"`
	Details      string           `protobuf:"octets,4,opt,name=info,proto3" json:"details,omitempty"`
	Ordinal     int64            `protobuf:"variableint,5,opt,name=index,proto3" json:"ordinal,omitempty"`
	Key       []byte           `protobuf:"octets,6,opt,name=key,proto3" json:"key,omitempty"`
	Item     []byte           `protobuf:"octets,7,opt,name=value,proto3" json:"item,omitempty"`
	EvidenceActions  *vault.EvidenceActions `protobuf:"octets,8,opt,name=proof_ops,json=proofOps,proto3" json:"evidence_actions,omitempty"`
	Level    int64            `protobuf:"variableint,9,opt,name=height,proto3" json:"level,omitempty"`
	Codex string           `protobuf:"octets,10,opt,name=codespace,proto3" json:"codex,omitempty"`
}

func (m *ReplyInquire) Restore()         { *m = ReplyInquire{} }
func (m *ReplyInquire) String() string { return proto.CompactTextString(m) }
func (*ReplyInquire) SchemaSignal()    {}
func (*ReplyInquire) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{25}
}
func (m *ReplyInquire) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyInquire) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyquery.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyInquire) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyquery.Merge(m, src)
}
func (m *ReplyInquire) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyInquire) XXX_Omitunclear() {
	xxx_messagedata_Replyquery.DiscardUnknown(m)
}

var xxx_messagedata_Replyquery proto.InternalMessageInfo

func (m *ReplyInquire) FetchCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyInquire) FetchTrace() string {
	if m != nil {
		return m.Log
	}
	return "REDACTED"
}

func (m *ReplyInquire) FetchDetails() string {
	if m != nil {
		return m.Details
	}
	return "REDACTED"
}

func (m *ReplyInquire) FetchOrdinal() int64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *ReplyInquire) FetchKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ReplyInquire) FetchItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ReplyInquire) FetchEvidenceActions() *vault.EvidenceActions {
	if m != nil {
		return m.EvidenceActions
	}
	return nil
}

func (m *ReplyInquire) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *ReplyInquire) FetchCodex() string {
	if m != nil {
		return m.Codex
	}
	return "REDACTED"
}

type ReplyInspectTransfer struct {
	Code      uint32  `protobuf:"variableint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      []byte  `protobuf:"octets,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string  `protobuf:"octets,3,opt,name=log,proto3" json:"log,omitempty"`
	Details      string  `protobuf:"octets,4,opt,name=info,proto3" json:"details,omitempty"`
	FuelDesired int64   `protobuf:"variableint,5,opt,name=gas_wanted,proto3" json:"fuel_desired,omitempty"`
	FuelApplied   int64   `protobuf:"variableint,6,opt,name=gas_used,proto3" json:"fuel_applied,omitempty"`
	Events    []Event `protobuf:"octets,7,rep,name=events,proto3" json:"events,omitempty"`
	Codex string  `protobuf:"octets,8,opt,name=codespace,proto3" json:"codex,omitempty"`
}

func (m *ReplyInspectTransfer) Restore()         { *m = ReplyInspectTransfer{} }
func (m *ReplyInspectTransfer) String() string { return proto.CompactTextString(m) }
func (*ReplyInspectTransfer) SchemaSignal()    {}
func (*ReplyInspectTransfer) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{26}
}
func (m *ReplyInspectTransfer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyInspectTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replytransfercheck.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyInspectTransfer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replytransfercheck.Merge(m, src)
}
func (m *ReplyInspectTransfer) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyInspectTransfer) XXX_Omitunclear() {
	xxx_messagedata_Replytransfercheck.DiscardUnknown(m)
}

var xxx_messagedata_Replytransfercheck proto.InternalMessageInfo

func (m *ReplyInspectTransfer) FetchCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyInspectTransfer) FetchData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReplyInspectTransfer) FetchTrace() string {
	if m != nil {
		return m.Log
	}
	return "REDACTED"
}

func (m *ReplyInspectTransfer) FetchDetails() string {
	if m != nil {
		return m.Details
	}
	return "REDACTED"
}

func (m *ReplyInspectTransfer) FetchFuelDesired() int64 {
	if m != nil {
		return m.FuelDesired
	}
	return 0
}

func (m *ReplyInspectTransfer) FetchFuelApplied() int64 {
	if m != nil {
		return m.FuelApplied
	}
	return 0
}

func (m *ReplyInspectTransfer) FetchEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ReplyInspectTransfer) FetchCodex() string {
	if m != nil {
		return m.Codex
	}
	return "REDACTED"
}

type ReplyEmbedTransfer struct {
	Code uint32 `protobuf:"variableint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *ReplyEmbedTransfer) Restore()         { *m = ReplyEmbedTransfer{} }
func (m *ReplyEmbedTransfer) String() string { return proto.CompactTextString(m) }
func (*ReplyEmbedTransfer) SchemaSignal()    {}
func (*ReplyEmbedTransfer) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{27}
}
func (m *ReplyEmbedTransfer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyEmbedTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replytransferinsert.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyEmbedTransfer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replytransferinsert.Merge(m, src)
}
func (m *ReplyEmbedTransfer) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyEmbedTransfer) XXX_Omitunclear() {
	xxx_messagedata_Replytransferinsert.DiscardUnknown(m)
}

var xxx_messagedata_Replytransferinsert proto.InternalMessageInfo

func (m *ReplyEmbedTransfer) FetchCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ReplyHarvestTrans struct {
	Txs [][]byte `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *ReplyHarvestTrans) Restore()         { *m = ReplyHarvestTrans{} }
func (m *ReplyHarvestTrans) String() string { return proto.CompactTextString(m) }
func (*ReplyHarvestTrans) SchemaSignal()    {}
func (*ReplyHarvestTrans) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{28}
}
func (m *ReplyHarvestTrans) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyHarvestTrans) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyreaptransfers.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyHarvestTrans) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyreaptransfers.Merge(m, src)
}
func (m *ReplyHarvestTrans) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyHarvestTrans) XXX_Omitunclear() {
	xxx_messagedata_Replyreaptransfers.DiscardUnknown(m)
}

var xxx_messagedata_Replyreaptransfers proto.InternalMessageInfo

func (m *ReplyHarvestTrans) FetchTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ReplyEndorse struct {
	PreserveLevel int64 `protobuf:"variableint,3,opt,name=retain_height,json=retainHeight,proto3" json:"preserve_level,omitempty"`
}

func (m *ReplyEndorse) Restore()         { *m = ReplyEndorse{} }
func (m *ReplyEndorse) String() string { return proto.CompactTextString(m) }
func (*ReplyEndorse) SchemaSignal()    {}
func (*ReplyEndorse) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{29}
}
func (m *ReplyEndorse) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyEndorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyendorse.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyEndorse) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyendorse.Merge(m, src)
}
func (m *ReplyEndorse) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyEndorse) XXX_Omitunclear() {
	xxx_messagedata_Replyendorse.DiscardUnknown(m)
}

var xxx_messagedata_Replyendorse proto.InternalMessageInfo

func (m *ReplyEndorse) FetchPreserveLevel() int64 {
	if m != nil {
		return m.PreserveLevel
	}
	return 0
}

type ReplyCatalogMirrors struct {
	Mirrors []*Mirror `protobuf:"octets,1,rep,name=snapshots,proto3" json:"mirrors,omitempty"`
}

func (m *ReplyCatalogMirrors) Restore()         { *m = ReplyCatalogMirrors{} }
func (m *ReplyCatalogMirrors) String() string { return proto.CompactTextString(m) }
func (*ReplyCatalogMirrors) SchemaSignal()    {}
func (*ReplyCatalogMirrors) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{30}
}
func (m *ReplyCatalogMirrors) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyCatalogMirrors) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replycatalogmirrors.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyCatalogMirrors) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replycatalogmirrors.Merge(m, src)
}
func (m *ReplyCatalogMirrors) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyCatalogMirrors) XXX_Omitunclear() {
	xxx_messagedata_Replycatalogmirrors.DiscardUnknown(m)
}

var xxx_messagedata_Replycatalogmirrors proto.InternalMessageInfo

func (m *ReplyCatalogMirrors) FetchMirrors() []*Mirror {
	if m != nil {
		return m.Mirrors
	}
	return nil
}

type ReplyProposalMirror struct {
	Outcome Replymirrorsnapshot_Outcome `protobuf:"variableint,1,opt,name=result,proto3,enum=tendermint.abci.ResponseOfferSnapshot_Result" json:"outcome,omitempty"`
}

func (m *ReplyProposalMirror) Restore()         { *m = ReplyProposalMirror{} }
func (m *ReplyProposalMirror) String() string { return proto.CompactTextString(m) }
func (*ReplyProposalMirror) SchemaSignal()    {}
func (*ReplyProposalMirror) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{31}
}
func (m *ReplyProposalMirror) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyProposalMirror) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replymirrorsnapshot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyProposalMirror) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replymirrorsnapshot.Merge(m, src)
}
func (m *ReplyProposalMirror) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyProposalMirror) XXX_Omitunclear() {
	xxx_messagedata_Replymirrorsnapshot.DiscardUnknown(m)
}

var xxx_messagedata_Replymirrorsnapshot proto.InternalMessageInfo

func (m *ReplyProposalMirror) FetchOutcome() Replymirrorsnapshot_Outcome {
	if m != nil {
		return m.Outcome
	}
	return Replymirrorsnapshot_UNCLEAR
}

type ReplyImportMirrorSegment struct {
	Segment []byte `protobuf:"octets,1,opt,name=chunk,proto3" json:"segment,omitempty"`
}

func (m *ReplyImportMirrorSegment) Restore()         { *m = ReplyImportMirrorSegment{} }
func (m *ReplyImportMirrorSegment) String() string { return proto.CompactTextString(m) }
func (*ReplyImportMirrorSegment) SchemaSignal()    {}
func (*ReplyImportMirrorSegment) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{32}
}
func (m *ReplyImportMirrorSegment) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyImportMirrorSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyloadmirrorsegment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyImportMirrorSegment) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyloadmirrorsegment.Merge(m, src)
}
func (m *ReplyImportMirrorSegment) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyImportMirrorSegment) XXX_Omitunclear() {
	xxx_messagedata_Replyloadmirrorsegment.DiscardUnknown(m)
}

var xxx_messagedata_Replyloadmirrorsegment proto.InternalMessageInfo

func (m *ReplyImportMirrorSegment) FetchSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

type ReplyExecuteMirrorSegment struct {
	Outcome        Replyexecutemirrorsegment_Outcome `protobuf:"variableint,1,opt,name=result,proto3,enum=tendermint.abci.ResponseApplySnapshotChunk_Result" json:"outcome,omitempty"`
	ReacquireSegments []uint32                          `protobuf:"variableint,2,rep,packed,name=refetch_chunks,json=refetchChunks,proto3" json:"reacquire_segments,omitempty"`
	DeclineEmitters []string                          `protobuf:"octets,3,rep,name=reject_senders,json=rejectSenders,proto3" json:"decline_emitters,omitempty"`
}

func (m *ReplyExecuteMirrorSegment) Restore()         { *m = ReplyExecuteMirrorSegment{} }
func (m *ReplyExecuteMirrorSegment) String() string { return proto.CompactTextString(m) }
func (*ReplyExecuteMirrorSegment) SchemaSignal()    {}
func (*ReplyExecuteMirrorSegment) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{33}
}
func (m *ReplyExecuteMirrorSegment) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyExecuteMirrorSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyexecutemirrorsegment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyExecuteMirrorSegment) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyexecutemirrorsegment.Merge(m, src)
}
func (m *ReplyExecuteMirrorSegment) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyExecuteMirrorSegment) XXX_Omitunclear() {
	xxx_messagedata_Replyexecutemirrorsegment.DiscardUnknown(m)
}

var xxx_messagedata_Replyexecutemirrorsegment proto.InternalMessageInfo

func (m *ReplyExecuteMirrorSegment) FetchOutcome() Replyexecutemirrorsegment_Outcome {
	if m != nil {
		return m.Outcome
	}
	return Replyexecutemirrorsegment_UNCLEAR
}

func (m *ReplyExecuteMirrorSegment) FetchReacquireSegments() []uint32 {
	if m != nil {
		return m.ReacquireSegments
	}
	return nil
}

func (m *ReplyExecuteMirrorSegment) FetchDeclineEmitters() []string {
	if m != nil {
		return m.DeclineEmitters
	}
	return nil
}

type ReplyArrangeNomination struct {
	Txs [][]byte `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *ReplyArrangeNomination) Restore()         { *m = ReplyArrangeNomination{} }
func (m *ReplyArrangeNomination) String() string { return proto.CompactTextString(m) }
func (*ReplyArrangeNomination) SchemaSignal()    {}
func (*ReplyArrangeNomination) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{34}
}
func (m *ReplyArrangeNomination) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyArrangeNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyarrangenomination.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyArrangeNomination) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyarrangenomination.Merge(m, src)
}
func (m *ReplyArrangeNomination) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyArrangeNomination) XXX_Omitunclear() {
	xxx_messagedata_Replyarrangenomination.DiscardUnknown(m)
}

var xxx_messagedata_Replyarrangenomination proto.InternalMessageInfo

func (m *ReplyArrangeNomination) FetchTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ReplyHandleNomination struct {
	Status Responseprocessnomination_Nominationstate `protobuf:"variableint,1,opt,name=status,proto3,enum=tendermint.abci.ResponseProcessProposal_ProposalStatus" json:"state,omitempty"`
}

func (m *ReplyHandleNomination) Restore()         { *m = ReplyHandleNomination{} }
func (m *ReplyHandleNomination) String() string { return proto.CompactTextString(m) }
func (*ReplyHandleNomination) SchemaSignal()    {}
func (*ReplyHandleNomination) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{35}
}
func (m *ReplyHandleNomination) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyHandleNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Responseprocessnomination.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyHandleNomination) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Responseprocessnomination.Merge(m, src)
}
func (m *ReplyHandleNomination) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyHandleNomination) XXX_Omitunclear() {
	xxx_messagedata_Responseprocessnomination.DiscardUnknown(m)
}

var xxx_messagedata_Responseprocessnomination proto.InternalMessageInfo

func (m *ReplyHandleNomination) FetchStatus() Responseprocessnomination_Nominationstate {
	if m != nil {
		return m.Status
	}
	return Responseprocessnomination_UNCLEAR
}

type ReplyExpandBallot struct {
	BallotAddition []byte `protobuf:"octets,1,opt,name=vote_extension,json=voteExtension,proto3" json:"ballot_addition,omitempty"`
}

func (m *ReplyExpandBallot) Restore()         { *m = ReplyExpandBallot{} }
func (m *ReplyExpandBallot) String() string { return proto.CompactTextString(m) }
func (*ReplyExpandBallot) SchemaSignal()    {}
func (*ReplyExpandBallot) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{36}
}
func (m *ReplyExpandBallot) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyExpandBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyballotextend.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyExpandBallot) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyballotextend.Merge(m, src)
}
func (m *ReplyExpandBallot) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyExpandBallot) XXX_Omitunclear() {
	xxx_messagedata_Replyballotextend.DiscardUnknown(m)
}

var xxx_messagedata_Replyballotextend proto.InternalMessageInfo

func (m *ReplyExpandBallot) FetchBallotAddition() []byte {
	if m != nil {
		return m.BallotAddition
	}
	return nil
}

type ReplyValidateBallotAddition struct {
	Status Responseverifyballotextension_Validatestatus `protobuf:"variableint,1,opt,name=status,proto3,enum=tendermint.abci.ResponseVerifyVoteExtension_VerifyStatus" json:"state,omitempty"`
}

func (m *ReplyValidateBallotAddition) Restore()         { *m = ReplyValidateBallotAddition{} }
func (m *ReplyValidateBallotAddition) String() string { return proto.CompactTextString(m) }
func (*ReplyValidateBallotAddition) SchemaSignal()    {}
func (*ReplyValidateBallotAddition) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{37}
}
func (m *ReplyValidateBallotAddition) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyValidateBallotAddition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Responseverifyballotextension.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyValidateBallotAddition) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Responseverifyballotextension.Merge(m, src)
}
func (m *ReplyValidateBallotAddition) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyValidateBallotAddition) XXX_Omitunclear() {
	xxx_messagedata_Responseverifyballotextension.DiscardUnknown(m)
}

var xxx_messagedata_Responseverifyballotextension proto.InternalMessageInfo

func (m *ReplyValidateBallotAddition) FetchStatus() Responseverifyballotextension_Validatestatus {
	if m != nil {
		return m.Status
	}
	return Responseverifyballotextension_UNCLEAR
}

type ReplyCompleteLedger struct {
	//
	Events []Event `protobuf:"octets,1,rep,name=events,proto3" json:"events,omitempty"`
	//
	//
	//
	TransOutcomes []*InvokeTransferOutcome `protobuf:"octets,2,rep,name=tx_results,json=txResults,proto3" json:"transfer_outcomes,omitempty"`
	//
	RatifierRefreshes []RatifierModify `protobuf:"octets,3,rep,name=validator_updates,json=validatorUpdates,proto3" json:"ratifier_refreshes"`
	//
	AgreementArgumentRefreshes *kinds1.AgreementOptions `protobuf:"octets,4,opt,name=consensus_param_updates,json=consensusParamUpdates,proto3" json:"agreement_argument_refreshes,omitempty"`
	//
	ApplicationDigest []byte `protobuf:"octets,5,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *ReplyCompleteLedger) Restore()         { *m = ReplyCompleteLedger{} }
func (m *ReplyCompleteLedger) String() string { return proto.CompactTextString(m) }
func (*ReplyCompleteLedger) SchemaSignal()    {}
func (*ReplyCompleteLedger) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{38}
}
func (m *ReplyCompleteLedger) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReplyCompleteLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Replyterminateblock.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyCompleteLedger) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Replyterminateblock.Merge(m, src)
}
func (m *ReplyCompleteLedger) XXX_Volume() int {
	return m.Volume()
}
func (m *ReplyCompleteLedger) XXX_Omitunclear() {
	xxx_messagedata_Replyterminateblock.DiscardUnknown(m)
}

var xxx_messagedata_Replyterminateblock proto.InternalMessageInfo

func (m *ReplyCompleteLedger) FetchEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ReplyCompleteLedger) FetchTransferOutcomes() []*InvokeTransferOutcome {
	if m != nil {
		return m.TransOutcomes
	}
	return nil
}

func (m *ReplyCompleteLedger) FetchRatifierRefreshes() []RatifierModify {
	if m != nil {
		return m.RatifierRefreshes
	}
	return nil
}

func (m *ReplyCompleteLedger) FetchAgreementArgumentRefreshes() *kinds1.AgreementOptions {
	if m != nil {
		return m.AgreementArgumentRefreshes
	}
	return nil
}

func (m *ReplyCompleteLedger) FetchApplicationDigest() []byte {
	if m != nil {
		return m.ApplicationDigest
	}
	return nil
}

type EndorseDetails struct {
	Cycle int32      `protobuf:"variableint,1,opt,name=round,proto3" json:"epoch,omitempty"`
	Ballots []BallotDetails `protobuf:"octets,2,rep,name=votes,proto3" json:"ballots"`
}

func (m *EndorseDetails) Restore()         { *m = EndorseDetails{} }
func (m *EndorseDetails) String() string { return proto.CompactTextString(m) }
func (*EndorseDetails) SchemaSignal()    {}
func (*EndorseDetails) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{39}
}
func (m *EndorseDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *EndorseDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Endorsementdata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndorseDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Endorsementdata.Merge(m, src)
}
func (m *EndorseDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *EndorseDetails) XXX_Omitunclear() {
	xxx_messagedata_Endorsementdata.DiscardUnknown(m)
}

var xxx_messagedata_Endorsementdata proto.InternalMessageInfo

func (m *EndorseDetails) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *EndorseDetails) FetchBallots() []BallotDetails {
	if m != nil {
		return m.Ballots
	}
	return nil
}

//
//
//
type ExpandedEndorseDetails struct {
	//
	Cycle int32 `protobuf:"variableint,1,opt,name=round,proto3" json:"epoch,omitempty"`
	//
	//
	Ballots []ExpandedBallotDetails `protobuf:"octets,2,rep,name=votes,proto3" json:"ballots"`
}

func (m *ExpandedEndorseDetails) Restore()         { *m = ExpandedEndorseDetails{} }
func (m *ExpandedEndorseDetails) String() string { return proto.CompactTextString(m) }
func (*ExpandedEndorseDetails) SchemaSignal()    {}
func (*ExpandedEndorseDetails) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{40}
}
func (m *ExpandedEndorseDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ExpandedEndorseDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Enhancedendorsementdata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedEndorseDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Enhancedendorsementdata.Merge(m, src)
}
func (m *ExpandedEndorseDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *ExpandedEndorseDetails) XXX_Omitunclear() {
	xxx_messagedata_Enhancedendorsementdata.DiscardUnknown(m)
}

var xxx_messagedata_Enhancedendorsementdata proto.InternalMessageInfo

func (m *ExpandedEndorseDetails) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *ExpandedEndorseDetails) FetchBallots() []ExpandedBallotDetails {
	if m != nil {
		return m.Ballots
	}
	return nil
}

//
//
//
type Event struct {
	Kind       string           `protobuf:"octets,1,opt,name=type,proto3" json:"kind,omitempty"`
	Properties []EventProperty `protobuf:"octets,2,rep,name=attributes,proto3" json:"properties,omitempty"`
}

func (m *Event) Restore()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) SchemaSignal()    {}
func (*Event) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{41}
}
func (m *Event) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Event) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Event.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Event.Merge(m, src)
}
func (m *Event) XXX_Volume() int {
	return m.Volume()
}
func (m *Event) XXX_Omitunclear() {
	xxx_messagedata_Event.DiscardUnknown(m)
}

var xxx_messagedata_Event proto.InternalMessageInfo

func (m *Event) FetchKind() string {
	if m != nil {
		return m.Kind
	}
	return "REDACTED"
}

func (m *Event) FetchProperties() []EventProperty {
	if m != nil {
		return m.Properties
	}
	return nil
}

//
type EventProperty struct {
	Key   string `protobuf:"octets,1,opt,name=key,proto3" json:"key,omitempty"`
	Item string `protobuf:"octets,2,opt,name=value,proto3" json:"item,omitempty"`
	Ordinal bool   `protobuf:"variableint,3,opt,name=index,proto3" json:"ordinal,omitempty"`
}

func (m *EventProperty) Restore()         { *m = EventProperty{} }
func (m *EventProperty) String() string { return proto.CompactTextString(m) }
func (*EventProperty) SchemaSignal()    {}
func (*EventProperty) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{42}
}
func (m *EventProperty) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *EventProperty) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Eventproperty.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventProperty) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Eventproperty.Merge(m, src)
}
func (m *EventProperty) XXX_Volume() int {
	return m.Volume()
}
func (m *EventProperty) XXX_Omitunclear() {
	xxx_messagedata_Eventproperty.DiscardUnknown(m)
}

var xxx_messagedata_Eventproperty proto.InternalMessageInfo

func (m *EventProperty) FetchKey() string {
	if m != nil {
		return m.Key
	}
	return "REDACTED"
}

func (m *EventProperty) FetchItem() string {
	if m != nil {
		return m.Item
	}
	return "REDACTED"
}

func (m *EventProperty) FetchOrdinal() bool {
	if m != nil {
		return m.Ordinal
	}
	return false
}

//
//
//
type InvokeTransferOutcome struct {
	Code      uint32  `protobuf:"variableint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      []byte  `protobuf:"octets,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string  `protobuf:"octets,3,opt,name=log,proto3" json:"log,omitempty"`
	Details      string  `protobuf:"octets,4,opt,name=info,proto3" json:"details,omitempty"`
	FuelDesired int64   `protobuf:"variableint,5,opt,name=gas_wanted,proto3" json:"fuel_desired,omitempty"`
	FuelApplied   int64   `protobuf:"variableint,6,opt,name=gas_used,proto3" json:"fuel_applied,omitempty"`
	Events    []Event `protobuf:"octets,7,rep,name=events,proto3" json:"events,omitempty"`
	Codex string  `protobuf:"octets,8,opt,name=codespace,proto3" json:"codex,omitempty"`
}

func (m *InvokeTransferOutcome) Restore()         { *m = InvokeTransferOutcome{} }
func (m *InvokeTransferOutcome) String() string { return proto.CompactTextString(m) }
func (*InvokeTransferOutcome) SchemaSignal()    {}
func (*InvokeTransferOutcome) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{43}
}
func (m *InvokeTransferOutcome) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *InvokeTransferOutcome) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Invoketransferoutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InvokeTransferOutcome) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Invoketransferoutcome.Merge(m, src)
}
func (m *InvokeTransferOutcome) XXX_Volume() int {
	return m.Volume()
}
func (m *InvokeTransferOutcome) XXX_Omitunclear() {
	xxx_messagedata_Invoketransferoutcome.DiscardUnknown(m)
}

var xxx_messagedata_Invoketransferoutcome proto.InternalMessageInfo

func (m *InvokeTransferOutcome) FetchCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InvokeTransferOutcome) FetchData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeTransferOutcome) FetchTrace() string {
	if m != nil {
		return m.Log
	}
	return "REDACTED"
}

func (m *InvokeTransferOutcome) FetchDetails() string {
	if m != nil {
		return m.Details
	}
	return "REDACTED"
}

func (m *InvokeTransferOutcome) FetchFuelDesired() int64 {
	if m != nil {
		return m.FuelDesired
	}
	return 0
}

func (m *InvokeTransferOutcome) FetchFuelApplied() int64 {
	if m != nil {
		return m.FuelApplied
	}
	return 0
}

func (m *InvokeTransferOutcome) FetchEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *InvokeTransferOutcome) FetchCodex() string {
	if m != nil {
		return m.Codex
	}
	return "REDACTED"
}

//
//
//
type TransOutcome struct {
	Level int64        `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Ordinal  uint32       `protobuf:"variableint,2,opt,name=index,proto3" json:"ordinal,omitempty"`
	Tx     []byte       `protobuf:"octets,3,opt,name=tx,proto3" json:"tx,omitempty"`
	Outcome InvokeTransferOutcome `protobuf:"octets,4,opt,name=result,proto3" json:"outcome"`
}

func (m *TransOutcome) Restore()         { *m = TransOutcome{} }
func (m *TransOutcome) String() string { return proto.CompactTextString(m) }
func (*TransOutcome) SchemaSignal()    {}
func (*TransOutcome) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{44}
}
func (m *TransOutcome) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *TransOutcome) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Transferoutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransOutcome) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Transferoutcome.Merge(m, src)
}
func (m *TransOutcome) XXX_Volume() int {
	return m.Volume()
}
func (m *TransOutcome) XXX_Omitunclear() {
	xxx_messagedata_Transferoutcome.DiscardUnknown(m)
}

var xxx_messagedata_Transferoutcome proto.InternalMessageInfo

func (m *TransOutcome) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *TransOutcome) FetchOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *TransOutcome) FetchTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TransOutcome) FetchOutcome() InvokeTransferOutcome {
	if m != nil {
		return m.Outcome
	}
	return InvokeTransferOutcome{}
}

type Ratifier struct {
	Location []byte `protobuf:"octets,1,opt,name=address,proto3" json:"location,omitempty"`
	//
	Energy int64 `protobuf:"variableint,3,opt,name=power,proto3" json:"energy,omitempty"`
}

func (m *Ratifier) Restore()         { *m = Ratifier{} }
func (m *Ratifier) String() string { return proto.CompactTextString(m) }
func (*Ratifier) SchemaSignal()    {}
func (*Ratifier) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{45}
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

func (m *Ratifier) FetchEnergy() int64 {
	if m != nil {
		return m.Energy
	}
	return 0
}

type RatifierModify struct {
	PublicKey vault.PublicKey `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_key"`
	Energy  int64            `protobuf:"variableint,2,opt,name=power,proto3" json:"energy,omitempty"`
}

func (m *RatifierModify) Restore()         { *m = RatifierModify{} }
func (m *RatifierModify) String() string { return proto.CompactTextString(m) }
func (*RatifierModify) SchemaSignal()    {}
func (*RatifierModify) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{46}
}
func (m *RatifierModify) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *RatifierModify) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ratifierrefresh.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RatifierModify) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ratifierrefresh.Merge(m, src)
}
func (m *RatifierModify) XXX_Volume() int {
	return m.Volume()
}
func (m *RatifierModify) XXX_Omitunclear() {
	xxx_messagedata_Ratifierrefresh.DiscardUnknown(m)
}

var xxx_messagedata_Ratifierrefresh proto.InternalMessageInfo

func (m *RatifierModify) FetchPublicKey() vault.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return vault.PublicKey{}
}

func (m *RatifierModify) FetchEnergy() int64 {
	if m != nil {
		return m.Energy
	}
	return 0
}

type BallotDetails struct {
	Ratifier   Ratifier          `protobuf:"octets,1,opt,name=validator,proto3" json:"ratifier"`
	LedgerUidMark kinds1.LedgerUIDMark `protobuf:"variableint,3,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uid_mark,omitempty"`
}

func (m *BallotDetails) Restore()         { *m = BallotDetails{} }
func (m *BallotDetails) String() string { return proto.CompactTextString(m) }
func (*BallotDetails) SchemaSignal()    {}
func (*BallotDetails) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{47}
}
func (m *BallotDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *BallotDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ballotdata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BallotDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ballotdata.Merge(m, src)
}
func (m *BallotDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *BallotDetails) XXX_Omitunclear() {
	xxx_messagedata_Ballotdata.DiscardUnknown(m)
}

var xxx_messagedata_Ballotdata proto.InternalMessageInfo

func (m *BallotDetails) FetchRatifier() Ratifier {
	if m != nil {
		return m.Ratifier
	}
	return Ratifier{}
}

func (m *BallotDetails) FetchLedgerUidMark() kinds1.LedgerUIDMark {
	if m != nil {
		return m.LedgerUidMark
	}
	return kinds1.LedgerUIDMarkUnclear
}

type ExpandedBallotDetails struct {
	//
	Ratifier Ratifier `protobuf:"octets,1,opt,name=validator,proto3" json:"ratifier"`
	//
	BallotAddition []byte `protobuf:"octets,3,opt,name=vote_extension,json=voteExtension,proto3" json:"ballot_addition,omitempty"`
	//
	AdditionAutograph []byte `protobuf:"octets,4,opt,name=extension_signature,json=extensionSignature,proto3" json:"addition_autograph,omitempty"`
	//
	LedgerUidMark kinds1.LedgerUIDMark `protobuf:"variableint,5,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uid_mark,omitempty"`
}

func (m *ExpandedBallotDetails) Restore()         { *m = ExpandedBallotDetails{} }
func (m *ExpandedBallotDetails) String() string { return proto.CompactTextString(m) }
func (*ExpandedBallotDetails) SchemaSignal()    {}
func (*ExpandedBallotDetails) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{48}
}
func (m *ExpandedBallotDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ExpandedBallotDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Enhancedballotdata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedBallotDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Enhancedballotdata.Merge(m, src)
}
func (m *ExpandedBallotDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *ExpandedBallotDetails) XXX_Omitunclear() {
	xxx_messagedata_Enhancedballotdata.DiscardUnknown(m)
}

var xxx_messagedata_Enhancedballotdata proto.InternalMessageInfo

func (m *ExpandedBallotDetails) FetchRatifier() Ratifier {
	if m != nil {
		return m.Ratifier
	}
	return Ratifier{}
}

func (m *ExpandedBallotDetails) FetchBallotAddition() []byte {
	if m != nil {
		return m.BallotAddition
	}
	return nil
}

func (m *ExpandedBallotDetails) FetchAdditionAutograph() []byte {
	if m != nil {
		return m.AdditionAutograph
	}
	return nil
}

func (m *ExpandedBallotDetails) FetchLedgerUidMark() kinds1.LedgerUIDMark {
	if m != nil {
		return m.LedgerUidMark
	}
	return kinds1.LedgerUIDMarkUnclear
}

type Malpractice struct {
	Kind MalpracticeKind `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.abci.MisbehaviorType" json:"kind,omitempty"`
	//
	Ratifier Ratifier `protobuf:"octets,2,opt,name=validator,proto3" json:"ratifier"`
	//
	Level int64 `protobuf:"variableint,3,opt,name=height,proto3" json:"level,omitempty"`
	//
	Time time.Time `protobuf:"octets,4,opt,name=time,proto3,stdtime" json:"moment"`
	//
	//
	//
	SumPollingEnergy int64 `protobuf:"variableint,5,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_polling_energy,omitempty"`
}

func (m *Malpractice) Restore()         { *m = Malpractice{} }
func (m *Malpractice) String() string { return proto.CompactTextString(m) }
func (*Malpractice) SchemaSignal()    {}
func (*Malpractice) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{49}
}
func (m *Malpractice) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Malpractice) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Malpractice.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Malpractice) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Malpractice.Merge(m, src)
}
func (m *Malpractice) XXX_Volume() int {
	return m.Volume()
}
func (m *Malpractice) XXX_Omitunclear() {
	xxx_messagedata_Malpractice.DiscardUnknown(m)
}

var xxx_messagedata_Malpractice proto.InternalMessageInfo

func (m *Malpractice) FetchKind() MalpracticeKind {
	if m != nil {
		return m.Kind
	}
	return Misconductkind_UNCLEAR
}

func (m *Malpractice) FetchRatifier() Ratifier {
	if m != nil {
		return m.Ratifier
	}
	return Ratifier{}
}

func (m *Malpractice) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Malpractice) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Malpractice) FetchSumPollingEnergy() int64 {
	if m != nil {
		return m.SumPollingEnergy
	}
	return 0
}

type Mirror struct {
	Level   uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Layout   uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Segments   uint32 `protobuf:"variableint,3,opt,name=chunks,proto3" json:"segments,omitempty"`
	Digest     []byte `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Metainfo []byte `protobuf:"octets,5,opt,name=metadata,proto3" json:"metainfo,omitempty"`
}

func (m *Mirror) Restore()         { *m = Mirror{} }
func (m *Mirror) String() string { return proto.CompactTextString(m) }
func (*Mirror) SchemaSignal()    {}
func (*Mirror) Definition() ([]byte, []int) {
	return filedefinition_252557cfdd89a31a, []int{50}
}
func (m *Mirror) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Mirror) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Mirror.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mirror) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Mirror.Merge(m, src)
}
func (m *Mirror) XXX_Volume() int {
	return m.Volume()
}
func (m *Mirror) XXX_Omitunclear() {
	xxx_messagedata_Mirror.DiscardUnknown(m)
}

var xxx_messagedata_Mirror proto.InternalMessageInfo

func (m *Mirror) FetchLevel() uint64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Mirror) FetchLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *Mirror) FetchSegments() uint32 {
	if m != nil {
		return m.Segments
	}
	return 0
}

func (m *Mirror) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Mirror) FetchMetainfo() []byte {
	if m != nil {
		return m.Metainfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("REDACTED", Transfercheckkind_label, Transfercheckkind_item)
	proto.RegisterEnum("REDACTED", Misconductkind_label, Misconductkind_item)
	proto.RegisterEnum("REDACTED", Replymirrorsnapshot_Outcome_label, Replymirrorsnapshot_Outcome_item)
	proto.RegisterEnum("REDACTED", Replyexecutemirrorsegment_Outcome_label, Replyexecutemirrorsegment_Outcome_item)
	proto.RegisterEnum("REDACTED", Responseprocessnomination_Nominationstate_label, Responseprocessnomination_Nominationstate_item)
	proto.RegisterEnum("REDACTED", Responseverifyballotextension_Validatestatus_label, Responseverifyballotextension_Validatestatus_item)
	proto.RegisterType((*Query)(nil), "REDACTED")
	proto.RegisterType((*QueryReverberate)(nil), "REDACTED")
	proto.RegisterType((*QueryPurge)(nil), "REDACTED")
	proto.RegisterType((*QueryDetails)(nil), "REDACTED")
	proto.RegisterType((*QueryInitSeries)(nil), "REDACTED")
	proto.RegisterType((*QueryInquire)(nil), "REDACTED")
	proto.RegisterType((*QueryInspectTransfer)(nil), "REDACTED")
	proto.RegisterType((*QueryEmbedTransfer)(nil), "REDACTED")
	proto.RegisterType((*QueryHarvestTrans)(nil), "REDACTED")
	proto.RegisterType((*QueryEndorse)(nil), "REDACTED")
	proto.RegisterType((*QueryCatalogMirrors)(nil), "REDACTED")
	proto.RegisterType((*QueryProposalMirror)(nil), "REDACTED")
	proto.RegisterType((*QueryImportMirrorSegment)(nil), "REDACTED")
	proto.RegisterType((*QueryExecuteMirrorSegment)(nil), "REDACTED")
	proto.RegisterType((*QueryArrangeNomination)(nil), "REDACTED")
	proto.RegisterType((*QueryHandleNomination)(nil), "REDACTED")
	proto.RegisterType((*QueryExpandBallot)(nil), "REDACTED")
	proto.RegisterType((*QueryValidateBallotAddition)(nil), "REDACTED")
	proto.RegisterType((*QueryCompleteLedger)(nil), "REDACTED")
	proto.RegisterType((*Reply)(nil), "REDACTED")
	proto.RegisterType((*ReplyExemption)(nil), "REDACTED")
	proto.RegisterType((*ReplyReverberate)(nil), "REDACTED")
	proto.RegisterType((*ReplyPurge)(nil), "REDACTED")
	proto.RegisterType((*ReplyDetails)(nil), "REDACTED")
	proto.RegisterType((*ReplyInitSeries)(nil), "REDACTED")
	proto.RegisterType((*ReplyInquire)(nil), "REDACTED")
	proto.RegisterType((*ReplyInspectTransfer)(nil), "REDACTED")
	proto.RegisterType((*ReplyEmbedTransfer)(nil), "REDACTED")
	proto.RegisterType((*ReplyHarvestTrans)(nil), "REDACTED")
	proto.RegisterType((*ReplyEndorse)(nil), "REDACTED")
	proto.RegisterType((*ReplyCatalogMirrors)(nil), "REDACTED")
	proto.RegisterType((*ReplyProposalMirror)(nil), "REDACTED")
	proto.RegisterType((*ReplyImportMirrorSegment)(nil), "REDACTED")
	proto.RegisterType((*ReplyExecuteMirrorSegment)(nil), "REDACTED")
	proto.RegisterType((*ReplyArrangeNomination)(nil), "REDACTED")
	proto.RegisterType((*ReplyHandleNomination)(nil), "REDACTED")
	proto.RegisterType((*ReplyExpandBallot)(nil), "REDACTED")
	proto.RegisterType((*ReplyValidateBallotAddition)(nil), "REDACTED")
	proto.RegisterType((*ReplyCompleteLedger)(nil), "REDACTED")
	proto.RegisterType((*EndorseDetails)(nil), "REDACTED")
	proto.RegisterType((*ExpandedEndorseDetails)(nil), "REDACTED")
	proto.RegisterType((*Event)(nil), "REDACTED")
	proto.RegisterType((*EventProperty)(nil), "REDACTED")
	proto.RegisterType((*InvokeTransferOutcome)(nil), "REDACTED")
	proto.RegisterType((*TransOutcome)(nil), "REDACTED")
	proto.RegisterType((*Ratifier)(nil), "REDACTED")
	proto.RegisterType((*RatifierModify)(nil), "REDACTED")
	proto.RegisterType((*BallotDetails)(nil), "REDACTED")
	proto.RegisterType((*ExpandedBallotDetails)(nil), "REDACTED")
	proto.RegisterType((*Malpractice)(nil), "REDACTED")
	proto.RegisterType((*Mirror)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_252557cfdd89a31a) }

var filedefinition_252557cfdd89a31a = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5a, 0xc9, 0x73, 0x1b, 0xc7,
	0xd5, 0xc7, 0x60, 0xc7, 0xc3, 0xc2, 0x61, 0x93, 0xa2, 0x20, 0x48, 0x26, 0xa9, 0x51, 0xd9, 0x96,
	0x65, 0x9b, 0xf4, 0x27, 0x7d, 0xf2, 0x52, 0xb2, 0xbf, 0x2f, 0x20, 0x04, 0x1a, 0xa4, 0x64, 0x92,
	0x1e, 0x42, 0x72, 0x39, 0x8b, 0xc7, 0x43, 0xa0, 0x49, 0x8c, 0x05, 0x60, 0xc6, 0x33, 0x0d, 0x1a,
	0xf4, 0x29, 0x15, 0x27, 0x55, 0x29, 0x9f, 0x5c, 0x95, 0x4a, 0x95, 0x0f, 0xf1, 0x21, 0x87, 0x5c,
	0xf2, 0x17, 0xe4, 0x90, 0xf2, 0x29, 0x07, 0x1f, 0x72, 0xf0, 0x31, 0x27, 0x27, 0x65, 0xdf, 0x7c,
	0xcd, 0x21, 0xd7, 0x54, 0x2f, 0xb3, 0x01, 0x18, 0x00, 0x94, 0x9d, 0x43, 0x2a, 0xb9, 0x75, 0xbf,
	0x79, 0xef, 0xf5, 0xf4, 0xeb, 0xee, 0xb7, 0xfc, 0xba, 0xe1, 0x32, 0xc1, 0xfd, 0x36, 0xb6, 0x7b,
	0x46, 0x9f, 0x6c, 0xea, 0x47, 0x2d, 0x63, 0x93, 0x9c, 0x59, 0xd8, 0xd9, 0xb0, 0x6c, 0x93, 0x98,
	0x68, 0xc1, 0xff, 0xb8, 0x41, 0x3f, 0x56, 0x9e, 0x08, 0x70, 0xb7, 0xec, 0x33, 0x8b, 0x98, 0x9b,
	0x96, 0x6d, 0x9a, 0xc7, 0x9c, 0xbf, 0x72, 0x65, 0xfc, 0xf3, 0x23, 0x7c, 0x26, 0xb4, 0x85, 0x84,
	0xd9, 0x28, 0x9b, 0x96, 0x6e, 0xeb, 0x3d, 0xf7, 0xf3, 0xfa, 0xd8, 0xe7, 0x53, 0xbd, 0x6b, 0xb4,
	0x75, 0x62, 0xda, 0x82, 0x63, 0xed, 0xc4, 0x34, 0x4f, 0xba, 0x78, 0x93, 0xf5, 0x8e, 0x06, 0xc7,
	0x9b, 0xc4, 0xe8, 0x61, 0x87, 0xe8, 0x3d, 0x4b, 0x30, 0x2c, 0x9f, 0x98, 0x27, 0x26, 0x6b, 0x6e,
	0xd2, 0x16, 0xa7, 0x2a, 0x9f, 0x03, 0x64, 0x54, 0xfc, 0xfe, 0x00, 0x3b, 0x04, 0xdd, 0x84, 0x24,
	0x6e, 0x75, 0xcc, 0xb2, 0xb4, 0x2e, 0x5d, 0xcf, 0xdf, 0xbc, 0xb2, 0x31, 0x32, 0xc1, 0x0d, 0xc1,
	0x57, 0x6f, 0x75, 0xcc, 0x46, 0x4c, 0x65, 0xbc, 0xe8, 0x36, 0xa4, 0x8e, 0xbb, 0x03, 0xa7, 0x53,
	0x8e, 0x33, 0xa1, 0x27, 0xa2, 0x84, 0xb6, 0x29, 0x53, 0x23, 0xa6, 0x72, 0x6e, 0x3a, 0x94, 0xd1,
	0x3f, 0x36, 0xcb, 0x89, 0xe9, 0x43, 0xed, 0xf4, 0x8f, 0xd9, 0x50, 0x94, 0x17, 0x6d, 0x01, 0x18,
	0x7d, 0x83, 0x68, 0xad, 0x8e, 0x6e, 0xf4, 0xcb, 0x29, 0x26, 0x79, 0x35, 0x5a, 0xd2, 0x20, 0x35,
	0xca, 0xd8, 0x88, 0xa9, 0x39, 0xc3, 0xed, 0xd0, 0xdf, 0x7d, 0x7f, 0x80, 0xed, 0xb3, 0x72, 0x7a,
	0xfa, 0xef, 0xbe, 0x49, 0x99, 0xe8, 0xef, 0x32, 0x6e, 0xf4, 0x2a, 0x64, 0x5b, 0x1d, 0xdc, 0x7a,
	0xa4, 0x91, 0x61, 0x39, 0xcb, 0x24, 0xd7, 0xa2, 0x24, 0x6b, 0x94, 0xaf, 0x39, 0x6c, 0xc4, 0xd4,
	0x4c, 0x8b, 0x37, 0xd1, 0xcb, 0x90, 0x6e, 0x99, 0xbd, 0x9e, 0x41, 0xca, 0x79, 0x26, 0xbb, 0x1a,
	0x29, 0xcb, 0xb8, 0x1a, 0x31, 0x55, 0xf0, 0xa3, 0x3d, 0x28, 0x75, 0x0d, 0x87, 0x68, 0x4e, 0x5f,
	0xb7, 0x9c, 0x8e, 0x49, 0x9c, 0x72, 0x81, 0x69, 0x78, 0x32, 0x4a, 0xc3, 0x7d, 0xc3, 0x21, 0x87,
	0x2e, 0x73, 0x23, 0xa6, 0x16, 0xbb, 0x41, 0x02, 0xd5, 0x67, 0x1e, 0x1f, 0x63, 0xdb, 0x53, 0x58,
	0x2e, 0x4e, 0xd7, 0xb7, 0x4f, 0xb9, 0x5d, 0x79, 0xaa, 0xcf, 0x0c, 0x12, 0xd0, 0x8f, 0x60, 0xa9,
	0x6b, 0xea, 0x6d, 0x4f, 0x9d, 0xd6, 0xea, 0x0c, 0xfa, 0x8f, 0xca, 0x25, 0xa6, 0xf4, 0x99, 0xc8,
	0x9f, 0x34, 0xf5, 0xb6, 0xab, 0xa2, 0x46, 0x05, 0x1a, 0x31, 0x75, 0xb1, 0x3b, 0x4a, 0x44, 0xef,
	0xc0, 0xb2, 0x6e, 0x59, 0xdd, 0xb3, 0x51, 0xed, 0x0b, 0x4c, 0xfb, 0x8d, 0x28, 0xed, 0x55, 0x2a,
	0x33, 0xaa, 0x1e, 0xe9, 0x63, 0x54, 0xd4, 0x04, 0xd9, 0xb2, 0xb1, 0xa5, 0xdb, 0x58, 0xb3, 0x6c,
	0xd3, 0x32, 0x1d, 0xbd, 0x5b, 0x96, 0x99, 0xee, 0xa7, 0xa3, 0x74, 0x1f, 0x70, 0xfe, 0x03, 0xc1,
	0xde, 0x88, 0xa9, 0x0b, 0x56, 0x98, 0xc4, 0xb5, 0x9a, 0x2d, 0xec, 0x38, 0xbe, 0xd6, 0xc5, 0x59,
	0x5a, 0x19, 0x7f, 0x58, 0x6b, 0x88, 0x84, 0xea, 0x90, 0xc7, 0x43, 0x2a, 0xae, 0x9d, 0x9a, 0x04,
	0x97, 0x11, 0x53, 0xa8, 0x44, 0x9e, 0x50, 0xc6, 0xfa, 0xd0, 0x24, 0xb8, 0x11, 0x53, 0x01, 0x7b,
	0x3d, 0xa4, 0xc3, 0x85, 0x53, 0x6c, 0x1b, 0xc7, 0x67, 0x4c, 0x8d, 0xc6, 0xbe, 0x38, 0x86, 0xd9,
	0x2f, 0x2f, 0x31, 0x85, 0xcf, 0x46, 0x29, 0x7c, 0xc8, 0x84, 0xa8, 0x8a, 0xba, 0x2b, 0xd2, 0x88,
	0xa9, 0x4b, 0xa7, 0xe3, 0x64, 0xba, 0xc5, 0x8e, 0x8d, 0xbe, 0xde, 0x35, 0x3e, 0xc4, 0xda, 0x51,
	0xd7, 0x6c, 0x3d, 0x2a, 0x2f, 0x4f, 0xdf, 0x62, 0xdb, 0x82, 0x7b, 0x8b, 0x32, 0xd3, 0x2d, 0x76,
	0x1c, 0x24, 0xa0, 0xff, 0x87, 0x9c, 0xd1, 0x77, 0xb0, 0x4d, 0xe8, 0xd9, 0xbb, 0xc0, 0x54, 0xad,
	0x47, 0x1f, 0x7a, 0xca, 0xc8, 0x0e, 0x5f, 0xd6, 0x10, 0x6d, 0x7a, 0x76, 0x6d, 0xac, 0x5b, 0x1a,
	0x19, 0x3a, 0xe5, 0x95, 0xe9, 0x67, 0x57, 0xc5, 0xba, 0xd5, 0x1c, 0xd2, 0x73, 0x93, 0xb1, 0x79,
	0x73, 0x2b, 0x03, 0xa9, 0x53, 0xbd, 0x3b, 0xc0, 0xbb, 0xc9, 0x6c, 0x52, 0x4e, 0xed, 0x26, 0xb3,
	0x19, 0x39, 0xbb, 0x9b, 0xcc, 0xe6, 0x64, 0xd8, 0x4d, 0x66, 0x41, 0xce, 0x2b, 0x4f, 0x43, 0x3e,
	0xe0, 0x17, 0x51, 0x19, 0x32, 0x3d, 0xec, 0x38, 0xfa, 0x09, 0x66, 0x6e, 0x34, 0xa7, 0xba, 0x5d,
	0xa5, 0x04, 0x85, 0xa0, 0x2f, 0x54, 0x3e, 0x91, 0x3c, 0x49, 0xea, 0xe6, 0xa8, 0xe4, 0x29, 0xb6,
	0xd9, 0x6a, 0x08, 0x49, 0xd1, 0x45, 0xd7, 0xa0, 0xc8, 0x2c, 0xa9, 0xb9, 0xdf, 0xa9, 0xaf, 0x4d,
	0xaa, 0x05, 0x46, 0x7c, 0x28, 0x98, 0xd6, 0x20, 0x6f, 0xdd, 0xb4, 0x3c, 0x96, 0x04, 0x63, 0x01,
	0xeb, 0xa6, 0xe5, 0x32, 0x5c, 0x85, 0x02, 0x9d, 0xab, 0xc7, 0x91, 0x64, 0x83, 0xe4, 0x29, 0x4d,
	0xb0, 0x28, 0x7f, 0x8e, 0x83, 0x3c, 0xea, 0x3f, 0xd1, 0xcb, 0x90, 0xa4, 0xa1, 0x44, 0x44, 0x85,
	0xca, 0x06, 0x8f, 0x33, 0x1b, 0x6e, 0x9c, 0xd9, 0x68, 0xba, 0x71, 0x66, 0x2b, 0xfb, 0xc5, 0x57,
	0x6b, 0xb1, 0x4f, 0xfe, 0xba, 0x26, 0xa9, 0x4c, 0x02, 0x5d, 0xa2, 0x5e, 0x53, 0x37, 0xfa, 0x9a,
	0xd1, 0x66, 0xbf, 0x9c, 0xa3, 0x2e, 0x51, 0x37, 0xfa, 0x3b, 0x6d, 0x74, 0x1f, 0xe4, 0x96, 0xd9,
	0x77, 0x70, 0xdf, 0x19, 0x38, 0x1a, 0x8f, 0x74, 0x22, 0x16, 0x84, 0x3c, 0x3a, 0x8f, 0xb7, 0x35,
	0x97, 0xf3, 0x80, 0x31, 0xaa, 0x0b, 0xad, 0x30, 0x01, 0x6d, 0x03, 0x78, 0xe1, 0xd0, 0x29, 0x27,
	0xd7, 0x13, 0x13, 0x37, 0xc9, 0x43, 0x97, 0xe5, 0x81, 0xd5, 0xd6, 0x09, 0xde, 0x4a, 0xd2, 0xdf,
	0x55, 0x03, 0x92, 0xe8, 0x29, 0x58, 0xd0, 0x2d, 0x4b, 0x73, 0x88, 0x4e, 0xb0, 0x76, 0x74, 0x46,
	0xb0, 0xc3, 0xc2, 0x4c, 0x41, 0x2d, 0xea, 0x96, 0x75, 0x48, 0xa9, 0x5b, 0x94, 0x88, 0x9e, 0x84,
	0x12, 0x0d, 0x29, 0x86, 0xde, 0xd5, 0x3a, 0xd8, 0x38, 0xe9, 0x10, 0x16, 0x4e, 0x12, 0x6a, 0x51,
	0x50, 0x1b, 0x8c, 0xa8, 0xb4, 0xbd, 0x15, 0x67, 0xe1, 0x04, 0x21, 0x48, 0xb6, 0x75, 0xa2, 0x33,
	0x4b, 0x16, 0x54, 0xd6, 0xa6, 0x34, 0x4b, 0x27, 0x1d, 0x61, 0x1f, 0xd6, 0x46, 0x2b, 0x90, 0x16,
	0x6a, 0x13, 0x4c, 0xad, 0xe8, 0xa1, 0x65, 0x48, 0x59, 0xb6, 0x79, 0x8a, 0xd9, 0xd2, 0x65, 0x55,
	0xde, 0x51, 0x54, 0x28, 0x85, 0x43, 0x0f, 0x2a, 0x41, 0x9c, 0x0c, 0xc5, 0x28, 0x71, 0x32, 0x44,
	0x2f, 0x40, 0x92, 0x1a, 0x92, 0x8d, 0x51, 0x9a, 0x10, 0x6c, 0x85, 0x5c, 0xf3, 0xcc, 0xc2, 0x2a,
	0xe3, 0x54, 0xae, 0xc2, 0xc2, 0xc8, 0x91, 0x1a, 0x55, 0xaa, 0x6c, 0x7b, 0xc3, 0x8a, 0x53, 0x83,
	0x2e, 0x43, 0xae, 0xa7, 0x0f, 0x85, 0xdd, 0x24, 0xb6, 0xff, 0xb2, 0x3d, 0x7d, 0xc8, 0x4d, 0x76,
	0x11, 0x32, 0xf4, 0xe3, 0x89, 0xee, 0x88, 0xdd, 0x9b, 0xee, 0xe9, 0xc3, 0xd7, 0x75, 0x47, 0x59,
	0x80, 0x62, 0x28, 0xfa, 0x29, 0x2b, 0xb0, 0x3c, 0x29, 0x98, 0x29, 0x1d, 0x8f, 0x1e, 0x0a, 0x4a,
	0xe8, 0x36, 0x64, 0xbd, 0x68, 0xc6, 0xf7, 0xe8, 0xa5, 0xb1, 0x19, 0xba, 0xcc, 0xaa, 0xc7, 0x4a,
	0x37, 0x27, 0x5d, 0xeb, 0x8e, 0x2e, 0x72, 0x97, 0x82, 0x9a, 0xd1, 0x2d, 0xab, 0xa1, 0x3b, 0x1d,
	0xe5, 0x5d, 0x28, 0x47, 0x45, 0xaa, 0xc0, 0xda, 0xf0, 0x19, 0xba, 0x6b, 0xb3, 0x02, 0xe9, 0x63,
	0xd3, 0xee, 0xe9, 0x84, 0x29, 0x2b, 0xaa, 0xa2, 0x47, 0xd7, 0x8c, 0x47, 0xad, 0x04, 0x23, 0xf3,
	0x8e, 0xa2, 0xc1, 0xa5, 0xc8, 0x68, 0x45, 0x45, 0x8c, 0x7e, 0x1b, 0x73, 0x63, 0x17, 0x55, 0xde,
	0xf1, 0x15, 0xf1, 0x9f, 0xe5, 0x1d, 0x3a, 0xac, 0xc3, 0xe6, 0xca, 0xf4, 0xe7, 0x54, 0xd1, 0x53,
	0x3e, 0x4d, 0xc0, 0xca, 0xe4, 0x98, 0x85, 0xd6, 0xa1, 0x40, 0x57, 0x82, 0x04, 0x57, 0x2a, 0xa1,
	0x42, 0x4f, 0x1f, 0x36, 0xc5, 0x5a, 0xc9, 0x90, 0xa0, 0xce, 0x32, 0xbe, 0x9e, 0xb8, 0x5e, 0x50,
	0x69, 0x13, 0x3d, 0x80, 0xc5, 0xae, 0xd9, 0xd2, 0xbb, 0x5a, 0x57, 0x77, 0x88, 0x26, 0x92, 0x19,
	0x7e, 0x5e, 0xaf, 0x8d, 0x19, 0x9b, 0x47, 0x1f, 0xdc, 0xe6, 0xeb, 0x49, 0x7d, 0x9b, 0x38, 0x6a,
	0x0b, 0x4c, 0xc7, 0x7d, 0xdd, 0x5d, 0x6a, 0x74, 0x17, 0xf2, 0x3d, 0xc3, 0x39, 0xc2, 0x1d, 0xfd,
	0xd4, 0x30, 0x6d, 0x71, 0x70, 0xc7, 0xf7, 0xe7, 0x1b, 0x3e, 0x8f, 0xd0, 0x14, 0x14, 0x0b, 0x2c,
	0x49, 0x2a, 0x74, 0x5c, 0x5c, 0xc7, 0x95, 0x3e, 0xb7, 0xe3, 0x7a, 0x01, 0x96, 0xfb, 0x78, 0x48,
	0x34, 0xdf, 0x35, 0xf0, 0x7d, 0x92, 0x61, 0xa6, 0x47, 0xf4, 0x9b, 0xe7, 0x4c, 0x1c, 0xba, 0x65,
	0xd0, 0x33, 0x2c, 0xea, 0x5b, 0xa6, 0x83, 0x6d, 0x4d, 0x6f, 0xb7, 0x6d, 0xec, 0x38, 0x2c, 0x51,
	0x2c, 0xb0, 0x50, 0xce, 0xe8, 0x55, 0x4e, 0x56, 0x7e, 0x19, 0x5c, 0x9a, 0x70, 0x94, 0x17, 0x86,
	0x97, 0x7c, 0xc3, 0x1f, 0xc2, 0xb2, 0x90, 0x6f, 0x87, 0x6c, 0xcf, 0xb3, 0xed, 0xcb, 0xe3, 0x47,
	0x79, 0xd4, 0xe6, 0xc8, 0x15, 0x8f, 0x36, 0x7b, 0xe2, 0xf1, 0xcc, 0x8e, 0x20, 0xc9, 0x8c, 0x92,
	0xe4, 0xde, 0x8c, 0xb6, 0xff, 0xdd, 0x96, 0xe2, 0xa3, 0x04, 0x2c, 0x8e, 0xa5, 0x4c, 0xde, 0xc4,
	0xa4, 0x89, 0x13, 0x8b, 0x4f, 0x9c, 0x58, 0xe2, 0xdc, 0x13, 0x13, 0x6b, 0x9d, 0x9c, 0xbd, 0xd6,
	0xa9, 0xef, 0x71, 0xad, 0xd3, 0x8f, 0xb7, 0xd6, 0xff, 0xd2, 0x55, 0xf8, 0x8d, 0x04, 0x95, 0xe8,
	0x3c, 0x73, 0xe2, 0x72, 0x3c, 0x0b, 0x8b, 0xde, 0xaf, 0x78, 0xea, 0xb9, 0x63, 0x94, 0xbd, 0x0f,
	0x42, 0x7f, 0x64, 0x38, 0x7d, 0x12, 0x4a, 0x23, 0x59, 0x30, 0xdf, 0xca, 0xc5, 0xd3, 0xe0, 0xf8,
	0xca, 0xcf, 0x13, 0x5e, 0xe0, 0x09, 0xa5, 0xaa, 0x13, 0x4e, 0xeb, 0x9b, 0xb0, 0xd4, 0xc6, 0x2d,
	0xa3, 0xfd, 0xb8, 0x87, 0x75, 0x51, 0x48, 0xff, 0xf7, 0xac, 0x8e, 0xef, 0x92, 0x5f, 0xe7, 0x21,
	0xab, 0x62, 0xc7, 0xa2, 0xa9, 0x1f, 0xda, 0x82, 0x1c, 0x1e, 0xb6, 0xb0, 0x45, 0xdc, 0x6c, 0x79,
	0x72, 0x31, 0xc4, 0xb9, 0xeb, 0x2e, 0x67, 0x23, 0xa6, 0xfa, 0x62, 0xe8, 0x96, 0x40, 0x3b, 0xa2,
	0x81, 0x0b, 0x21, 0x1e, 0x84, 0x3b, 0x5e, 0x74, 0xe1, 0x8e, 0x44, 0x64, 0x25, 0xcf, 0xa5, 0x46,
	0xf0, 0x8e, 0x5b, 0x02, 0xef, 0x48, 0xce, 0x18, 0x2c, 0x04, 0x78, 0xd4, 0x42, 0x80, 0x47, 0x7a,
	0xc6, 0x34, 0x23, 0x10, 0x8f, 0x17, 0x5d, 0xc4, 0x23, 0x33, 0xe3, 0x8f, 0x47, 0x20, 0x8f, 0xd7,
	0x02, 0x90, 0x47, 0x2e, 0xb2, 0xec, 0xe2, 0xa2, 0x13, 0x30, 0x8f, 0x57, 0x3c, 0xcc, 0xa3, 0x10,
	0x59, 0x73, 0x09, 0xe1, 0x51, 0xd0, 0x63, 0x7f, 0x0c, 0xf4, 0xe0, 0x20, 0xc5, 0x53, 0x91, 0x2a,
	0x66, 0xa0, 0x1e, 0xfb, 0x63, 0xa8, 0x47, 0x69, 0x86, 0xc2, 0x19, 0xb0, 0xc7, 0x8f, 0x27, 0xc3,
	0x1e, 0xd1, 0xc0, 0x84, 0xf8, 0xcd, 0xf9, 0x70, 0x0f, 0x2d, 0x02, 0xf7, 0x90, 0x23, 0x6b, 0x74,
	0xae, 0x7e, 0x6e, 0xe0, 0xe3, 0xc1, 0x04, 0xe0, 0x83, 0x43, 0x14, 0xd7, 0x23, 0x95, 0xcf, 0x81,
	0x7c, 0x3c, 0x98, 0x80, 0x7c, 0xa0, 0x99, 0x6a, 0x67, 0x42, 0x1f, 0xdb, 0x61, 0xe8, 0x63, 0x29,
	0x22, 0xeb, 0xf4, 0x4f, 0x7b, 0x04, 0xf6, 0x71, 0x14, 0x85, 0x7d, 0x70, 0x7c, 0xe2, 0xb9, 0x48,
	0x8d, 0xe7, 0x00, 0x3f, 0xf6, 0xc7, 0xc0, 0x8f, 0x0b, 0x33, 0x76, 0xda, 0x0c, 0xf4, 0xe3, 0x07,
	0x41, 0xf4, 0x63, 0x25, 0x12, 0xf2, 0x74, 0x3d, 0xc0, 0x04, 0xf8, 0xe3, 0xb5, 0x00, 0xfc, 0x71,
	0x71, 0xc6, 0x39, 0x9e, 0x8e, 0x7f, 0xa4, 0xe4, 0xf4, 0x6e, 0x32, 0x9b, 0x95, 0x73, 0x1c, 0xf9,
	0xd8, 0x4d, 0x66, 0xf3, 0x72, 0x41, 0x79, 0x86, 0xa6, 0x50, 0x23, 0x8e, 0x96, 0x16, 0x2b, 0xd8,
	0xb6, 0x4d, 0x5b, 0x20, 0x19, 0xbc, 0xa3, 0x5c, 0xa7, 0xf5, 0xb0, 0xef, 0x54, 0xa7, 0x60, 0x25,
	0xac, 0x28, 0x0c, 0x38, 0x52, 0xe5, 0x0f, 0x92, 0x2f, 0xcb, 0xd0, 0x92, 0x60, 0x2d, 0x9d, 0x13,
	0xb5, 0x74, 0x00, 0x41, 0x89, 0x87, 0x11, 0x94, 0x35, 0xc8, 0xd3, 0x62, 0x6f, 0x04, 0x1c, 0xd1,
	0x2d, 0x0f, 0x1c, 0xb9, 0x01, 0x8b, 0x2c, 0x62, 0x73, 0x9c, 0x45, 0xc4, 0xc5, 0x24, 0x8b, 0x8b,
	0x0b, 0xf4, 0x03, 0x5f, 0x1e, 0x1e, 0x20, 0x9f, 0x87, 0xa5, 0x00, 0xaf, 0x57, 0x44, 0x72, 0xa4,
	0x40, 0xf6, 0xb8, 0xab, 0xa2, 0x9a, 0xfc, 0x93, 0xe4, 0x5b, 0xc8, 0x47, 0x55, 0x26, 0x01, 0x20,
	0xd2, 0xf7, 0x04, 0x80, 0xc4, 0x1f, 0x1b, 0x00, 0x09, 0x16, 0xc5, 0x89, 0x70, 0x51, 0xfc, 0x0f,
	0xc9, 0x5f, 0x13, 0x0f, 0xce, 0x68, 0x99, 0x6d, 0x2c, 0xca, 0x54, 0xd6, 0xa6, 0x39, 0x51, 0xd7,
	0x3c, 0x11, 0xc5, 0x28, 0x6d, 0x52, 0x2e, 0x2f, 0xf2, 0xe5, 0x44, 0x60, 0xf3, 0x2a, 0x5c, 0x9e,
	0x79, 0x88, 0x0a, 0x57, 0x86, 0xc4, 0x23, 0xcc, 0x91, 0xf9, 0x82, 0x4a, 0x9b, 0x94, 0x8f, 0x6d,
	0x3e, 0x91, 0x41, 0xf0, 0x0e, 0x7a, 0x19, 0x72, 0xec, 0x5e, 0x45, 0x33, 0x2d, 0x47, 0xa0, 0xf1,
	0xa1, 0xdc, 0x8a, 0x5f, 0xae, 0x6c, 0x1c, 0x50, 0x9e, 0x7d, 0xcb, 0x51, 0xb3, 0x96, 0x68, 0x05,
	0x52, 0x9e, 0x5c, 0x28, 0xe5, 0xb9, 0x02, 0x39, 0xfa, 0xf7, 0x8e, 0xa5, 0xb7, 0x70, 0x19, 0xd8,
	0x8f, 0xfa, 0x04, 0xe5, 0xf7, 0x71, 0x58, 0x18, 0x89, 0x74, 0x13, 0xe7, 0xee, 0x6e, 0xc9, 0x78,
	0x00, 0xde, 0x99, 0xcf, 0x1e, 0xab, 0x00, 0x27, 0xba, 0xa3, 0x7d, 0xa0, 0xf7, 0x09, 0x6e, 0x0b,
	0xa3, 0x04, 0x28, 0xa8, 0x02, 0x59, 0xda, 0x1b, 0x38, 0xb8, 0x2d, 0x90, 0x26, 0xaf, 0x8f, 0x1a,
	0x90, 0xc6, 0xa7, 0xb8, 0x4f, 0x9c, 0x72, 0x86, 0x2d, 0xfb, 0xca, 0x78, 0x3d, 0x4e, 0x3f, 0x6f,
	0x95, 0xe9, 0x62, 0x7f, 0xfb, 0xd5, 0x9a, 0xcc, 0xb9, 0x9f, 0x33, 0x7b, 0x06, 0xc1, 0x3d, 0x8b,
	0x9c, 0xa9, 0x42, 0x3e, 0x6c, 0x85, 0xec, 0x88, 0x15, 0x18, 0xe6, 0x59, 0x70, 0xf1, 0x05, 0x6a,
	0x53, 0xc3, 0xb4, 0x0d, 0x72, 0xa6, 0x16, 0x7b, 0xb8, 0x67, 0x99, 0x66, 0x57, 0xe3, 0x67, 0xfc,
	0x29, 0x90, 0x47, 0xdd, 0xd1, 0x24, 0x63, 0x29, 0xd7, 0x7c, 0x9b, 0xba, 0xf8, 0xd1, 0x58, 0x3e,
	0xad, 0x54, 0xa1, 0x14, 0xce, 0x12, 0xd0, 0x35, 0x28, 0xda, 0x98, 0xe8, 0x46, 0x5f, 0x0b, 0xa5,
	0xf4, 0x05, 0x4e, 0xe4, 0x07, 0x74, 0x37, 0x99, 0x95, 0xe4, 0xf8, 0x6e, 0x32, 0x1b, 0x97, 0x13,
	0xca, 0x01, 0x5c, 0x98, 0x98, 0x25, 0xa0, 0x97, 0x20, 0xe7, 0x27, 0x18, 0x12, 0x33, 0xdd, 0x14,
	0xdc, 0xc8, 0xe7, 0x55, 0x3e, 0x97, 0x7c, 0x95, 0x61, 0x24, 0xaa, 0x0e, 0x69, 0x1b, 0x3b, 0x83,
	0x2e, 0xc7, 0x86, 0x4a, 0x37, 0x9f, 0x9f, 0x2f, 0xbf, 0xa0, 0xd4, 0x41, 0x97, 0xa8, 0x42, 0x58,
	0x79, 0x07, 0xd2, 0x9c, 0x82, 0xf2, 0x90, 0x79, 0xb0, 0x77, 0x6f, 0x6f, 0xff, 0xad, 0x3d, 0x39,
	0x86, 0x00, 0xd2, 0xd5, 0x5a, 0xad, 0x7e, 0xd0, 0x94, 0x25, 0x94, 0x83, 0x54, 0x75, 0x6b, 0x5f,
	0x6d, 0xca, 0x71, 0x4a, 0x56, 0xeb, 0xbb, 0xf5, 0x5a, 0x53, 0x4e, 0xa0, 0x45, 0x28, 0xf2, 0xb6,
	0xb6, 0xbd, 0xaf, 0xbe, 0x51, 0x6d, 0xca, 0xc9, 0x00, 0xe9, 0xb0, 0xbe, 0x77, 0xb7, 0xae, 0xca,
	0x29, 0xe5, 0x7f, 0xe0, 0x52, 0x64, 0x46, 0xe2, 0xc3, 0x4c, 0x52, 0x00, 0x66, 0x52, 0x3e, 0x8d,
	0xd3, 0x12, 0x2d, 0x2a, 0xcd, 0x40, 0xbb, 0x23, 0x13, 0xbf, 0x79, 0x8e, 0x1c, 0x65, 0x64, 0xf6,
	0xb4, 0x2a, 0xb3, 0xf1, 0x31, 0x26, 0xad, 0x0e, 0x4f, 0x7b, 0xb8, 0x3b, 0x2b, 0xaa, 0x45, 0x41,
	0x65, 0x42, 0x0e, 0x67, 0x7b, 0x0f, 0xb7, 0x88, 0xc6, 0x77, 0xa4, 0xc3, 0x4a, 0xa3, 0x1c, 0x65,
	0xa3, 0xd4, 0x43, 0x4e, 0x54, 0xde, 0x3d, 0x97, 0x2d, 0x73, 0x90, 0x52, 0xeb, 0x4d, 0xf5, 0x6d,
	0x39, 0x81, 0x10, 0x94, 0x58, 0x53, 0x3b, 0xdc, 0xab, 0x1e, 0x1c, 0x36, 0xf6, 0xa9, 0x2d, 0x97,
	0x60, 0xc1, 0xb5, 0xa5, 0x4b, 0x4c, 0x29, 0xcf, 0xc2, 0xc5, 0x88, 0x1c, 0x69, 0xc2, 0x86, 0xfe,
	0xad, 0x14, 0xe4, 0x0e, 0xe7, 0x39, 0xfb, 0x90, 0x76, 0x88, 0x4e, 0x06, 0x8e, 0x30, 0xe2, 0x4b,
	0xf3, 0x26, 0x4d, 0x1b, 0x6e, 0xe3, 0x90, 0x89, 0xab, 0x42, 0x8d, 0x72, 0x1b, 0x4a, 0xe1, 0x2f,
	0xd1, 0x36, 0xf0, 0x37, 0x51, 0x5c, 0xb9, 0x03, 0x68, 0x3c, 0x97, 0x9a, 0x50, 0x2c, 0x4b, 0x93,
	0x8a, 0xe5, 0xdf, 0x49, 0x70, 0x79, 0x4a, 0xde, 0x84, 0xde, 0x1c, 0x99, 0xe4, 0x2b, 0xe7, 0xc9,
	0xba, 0x36, 0x38, 0x6d, 0x64, 0x9a, 0xb7, 0xa0, 0x10, 0xa4, 0xcf, 0x37, 0xc9, 0x6f, 0xe3, 0xfe,
	0x21, 0x0e, 0x57, 0xf5, 0xbe, 0x3f, 0x95, 0xbe, 0xa3, 0x3f, 0x7d, 0x15, 0x80, 0x0c, 0x35, 0xbe,
	0xad, 0xdd, 0xa0, 0xfc, 0xc4, 0x04, 0xb4, 0x14, 0xb7, 0x9a, 0x43, 0x71, 0x08, 0x72, 0x44, 0xb4,
	0x1c, 0x74, 0x18, 0x84, 0x38, 0x06, 0x2c, 0x60, 0x3b, 0xa2, 0xfc, 0x9f, 0x37, 0xb2, 0xfb, 0x50,
	0x08, 0x27, 0x3b, 0xe8, 0x6d, 0xb8, 0x38, 0x92, 0x75, 0x78, 0xaa, 0x93, 0xf3, 0x26, 0x1f, 0x17,
	0xc2, 0xc9, 0x87, 0xab, 0x3a, 0x98, 0x3a, 0xa4, 0xc2, 0xa9, 0xc3, 0xdb, 0x00, 0x3e, 0xd4, 0x41,
	0x3d, 0x8c, 0x6d, 0x0e, 0xfa, 0x6d, 0xb6, 0x03, 0x52, 0x2a, 0xef, 0xa0, 0xdb, 0x90, 0xa2, 0x3b,
	0xc9, 0xb5, 0xd3, 0xb8, 0x2b, 0xa6, 0x3b, 0x21, 0x00, 0x95, 0x70, 0x6e, 0xc5, 0x00, 0x34, 0x0e,
	0x37, 0x47, 0x0c, 0xf1, 0x5a, 0x78, 0x88, 0xab, 0x91, 0xc0, 0xf5, 0xe4, 0xa1, 0x3e, 0x84, 0x14,
	0x5b, 0x79, 0x1a, 0xce, 0xd8, 0x75, 0x8a, 0x48, 0x3d, 0x69, 0x1b, 0xfd, 0x04, 0x40, 0x27, 0xc4,
	0x36, 0x8e, 0x06, 0xfe, 0x00, 0x6b, 0x93, 0x77, 0x4e, 0xd5, 0xe5, 0xdb, 0xba, 0x22, 0xb6, 0xd0,
	0xb2, 0x2f, 0x1a, 0xd8, 0x46, 0x01, 0x85, 0xca, 0x1e, 0x94, 0xc2, 0xb2, 0x6e, 0xb2, 0xc4, 0xff,
	0x21, 0x9c, 0x2c, 0xf1, 0xdc, 0x57, 0x24, 0x4b, 0x5e, 0xaa, 0x95, 0xe0, 0x77, 0x46, 0xac, 0xa3,
	0xfc, 0x34, 0x0e, 0x85, 0xe0, 0xc6, 0xfb, 0xcf, 0xcb, 0x67, 0x94, 0x5f, 0x48, 0x90, 0xf5, 0xa6,
	0x1f, 0xbe, 0xd5, 0x09, 0xdd, 0xb8, 0x71, 0xeb, 0xc5, 0x83, 0x57, 0x31, 0xfc, 0x2a, 0x2c, 0xe1,
	0xdd, 0xaf, 0xdd, 0xf1, 0xc2, 0x5f, 0x14, 0xbc, 0x13, 0xb4, 0xb5, 0xd8, 0x55, 0x6e, 0xb4, 0xbf,
	0x03, 0x39, 0xef, 0xf4, 0xd2, 0x0a, 0xc6, 0x85, 0xc1, 0x24, 0x71, 0x86, 0x04, 0x88, 0xb9, 0x0c,
	0x29, 0xcb, 0xfc, 0x40, 0xdc, 0xf3, 0x24, 0x54, 0xde, 0x51, 0xda, 0xb0, 0x30, 0x72, 0xf4, 0xd1,
	0x1d, 0xc8, 0x58, 0x83, 0x23, 0xcd, 0xdd, 0x1c, 0x23, 0x60, 0xa1, 0x9b, 0x1b, 0x0f, 0x8e, 0xba,
	0x46, 0xeb, 0x1e, 0x3e, 0x73, 0x7f, 0xc6, 0x1a, 0x1c, 0xdd, 0xe3, 0x7b, 0x88, 0x8f, 0x12, 0x0f,
	0x8e, 0xf2, 0x2b, 0x09, 0xb2, 0xee, 0x99, 0x40, 0xff, 0x07, 0x39, 0xcf, 0xad, 0x78, 0x77, 0xc2,
	0x91, 0xfe, 0x48, 0xe8, 0xf7, 0x45, 0x50, 0xd5, 0xbd, 0xcc, 0x36, 0xda, 0xda, 0x71, 0x57, 0xe7,
	0x7b, 0xa9, 0x14, 0xb6, 0x19, 0x77, 0x3c, 0xcc, 0x1f, 0xef, 0xdc, 0xdd, 0xee, 0xea, 0x27, 0x6a,
	0x9e, 0xc9, 0xec, 0xb4, 0x69, 0x47, 0x64, 0x76, 0x7f, 0x97, 0x40, 0x1e, 0x3d, 0xb1, 0xdf, 0xf9,
	0xef, 0xc6, 0xc3, 0x5c, 0x62, 0x42, 0x98, 0x43, 0x9b, 0xb0, 0xe4, 0x71, 0x68, 0x8e, 0x71, 0xd2,
	0xd7, 0xc9, 0xc0, 0xc6, 0x02, 0x5e, 0x45, 0xde, 0xa7, 0x43, 0xf7, 0xcb, 0xf8, 0xac, 0x53, 0x8f,
	0x39, 0xeb, 0x8f, 0xe2, 0x90, 0x0f, 0x80, 0xbd, 0xe8, 0x7f, 0x03, 0xce, 0xa8, 0x34, 0x21, 0x32,
	0x04, 0x78, 0xfd, 0xfb, 0xdd, 0xb0, 0x99, 0xe2, 0xe7, 0x37, 0x53, 0x14, 0xa4, 0xee, 0x62, 0xc7,
	0xc9, 0x73, 0x63, 0xc7, 0xcf, 0x01, 0x22, 0x26, 0xd1, 0xbb, 0xda, 0xa9, 0x49, 0x8c, 0xfe, 0x89,
	0xc6, 0xb7, 0x21, 0x77, 0x1d, 0x32, 0xfb, 0xf2, 0x90, 0x7d, 0x38, 0x60, 0x3b, 0xf2, 0x67, 0x12,
	0x64, 0xbd, 0xb4, 0xfb, 0xbc, 0x57, 0xb2, 0x2b, 0x90, 0x16, 0x99, 0x25, 0xbf, 0x93, 0x15, 0xbd,
	0x89, 0x20, 0x79, 0x05, 0xb2, 0x3d, 0x4c, 0x74, 0xe6, 0x07, 0x79, 0x54, 0xf3, 0xfa, 0x37, 0x5e,
	0x81, 0x7c, 0xe0, 0xe6, 0x9c, 0xba, 0xc6, 0xbd, 0xfa, 0x5b, 0x72, 0xac, 0x92, 0xf9, 0xf8, 0xb3,
	0xf5, 0xc4, 0x1e, 0xfe, 0x80, 0x9e, 0x66, 0xb5, 0x5e, 0x6b, 0xd4, 0x6b, 0xf7, 0x64, 0xa9, 0x92,
	0xff, 0xf8, 0xb3, 0xf5, 0x8c, 0x8a, 0x19, 0x3e, 0x7a, 0xe3, 0x1e, 0x2c, 0x8c, 0x2c, 0x4c, 0x38,
	0x6d, 0x41, 0x50, 0xba, 0xfb, 0xe0, 0xe0, 0xfe, 0x4e, 0xad, 0xda, 0xac, 0x6b, 0x0f, 0xf7, 0x9b,
	0x75, 0x59, 0x42, 0x17, 0x61, 0xe9, 0xfe, 0xce, 0xeb, 0x8d, 0xa6, 0x56, 0xbb, 0xbf, 0x53, 0xdf,
	0x6b, 0x6a, 0xd5, 0x66, 0xb3, 0x5a, 0xbb, 0x27, 0xc7, 0x6f, 0xfe, 0xb1, 0x00, 0xc9, 0xea, 0x56,
	0x6d, 0x07, 0xd5, 0x20, 0xc9, 0x70, 0x95, 0xa9, 0x2f, 0xf7, 0x2a, 0xd3, 0x91, 0x6e, 0xb4, 0x0d,
	0x29, 0x06, 0xb9, 0xa0, 0xe9, 0x4f, 0xf9, 0x2a, 0x33, 0xa0, 0x6f, 0xfa, 0x33, 0xec, 0x44, 0x4e,
	0x7d, 0xdb, 0x57, 0x99, 0x8e, 0x84, 0xa3, 0xfb, 0x90, 0x71, 0x2b, 0xee, 0x59, 0x0f, 0xee, 0x2a,
	0x33, 0xe1, 0x69, 0xb4, 0x0f, 0x59, 0xaf, 0x26, 0x9d, 0xf9, 0x86, 0xa8, 0x32, 0x1b, 0x67, 0xa3,
	0xbf, 0xe7, 0x16, 0xaf, 0xb3, 0xde, 0x14, 0x55, 0x66, 0xa2, 0x6e, 0xd4, 0xf2, 0x1c, 0x58, 0x99,
	0xfe, 0x2a, 0xb1, 0x32, 0x03, 0xc2, 0x47, 0x3b, 0x90, 0x16, 0xd5, 0xf2, 0x8c, 0x87, 0x86, 0x95,
	0x59, 0xa0, 0x3c, 0x52, 0x21, 0xe7, 0x43, 0x56, 0xb3, 0xdf, 0x5a, 0x56, 0xe6, 0xb8, 0x9d, 0x40,
	0xef, 0x40, 0x31, 0x5c, 0x89, 0xcf, 0xf7, 0x98, 0xb1, 0x32, 0x27, 0xfc, 0x4f, 0xf5, 0x87, 0xcb,
	0xf2, 0xf9, 0x1e, 0x37, 0x56, 0xe6, 0xbc, 0x0d, 0x40, 0xef, 0xc1, 0xe2, 0x78, 0xd9, 0x3c, 0xff,
	0x5b, 0xc7, 0xca, 0x39, 0xee, 0x07, 0x50, 0x0f, 0xd0, 0x84, 0x72, 0xfb, 0x1c, 0x4f, 0x1f, 0x2b,
	0xe7, 0xb9, 0x2e, 0x40, 0x6d, 0x58, 0x18, 0xad, 0x61, 0xe7, 0x7d, 0x0a, 0x59, 0x99, 0xfb, 0xea,
	0x80, 0x8f, 0x12, 0xae, 0x7d, 0xe7, 0x7d, 0x1a, 0x59, 0x99, 0xfb, 0x26, 0x01, 0x3d, 0x00, 0x08,
	0x94, 0xaf, 0x73, 0x3c, 0x95, 0xac, 0xcc, 0x73, 0xa7, 0x80, 0x2c, 0x58, 0x9a, 0x54, 0xd7, 0x9e,
	0xe7, 0xe5, 0x64, 0xe5, 0x5c, 0x57, 0x0d, 0x74, 0x3f, 0x87, 0x2b, 0xd4, 0xf9, 0x5e, 0x52, 0x56,
	0xe6, 0xbc, 0x73, 0xd8, 0xaa, 0x7e, 0xf1, 0xf5, 0xaa, 0xf4, 0xe5, 0xd7, 0xab, 0xd2, 0xdf, 0xbe,
	0x5e, 0x95, 0x3e, 0xf9, 0x66, 0x35, 0xf6, 0xe5, 0x37, 0xab, 0xb1, 0xbf, 0x7c, 0xb3, 0x1a, 0xfb,
	0xe1, 0xd3, 0x27, 0x06, 0xe9, 0x0c, 0x8e, 0x36, 0x5a, 0x66, 0x6f, 0xb3, 0x65, 0xf6, 0x30, 0x39,
	0x3a, 0x26, 0x7e, 0xc3, 0x7f, 0x10, 0x7f, 0x94, 0x66, 0x01, 0xfe, 0xd6, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x52, 0xe8, 0x15, 0x70, 0x30, 0x2f, 0x00, 0x00,
}

//
var _ context.Context
var _ grpc.ClientConn

//
//
const _ = grpc.SupportPackageIsVersion4

//
//
//
type IfaceCustomer interface {
	Replicate(ctx context.Context, in *QueryReverberate, opts ...grpc.CallOption) (*ReplyReverberate, error)
	Purge(ctx context.Context, in *QueryPurge, opts ...grpc.CallOption) (*ReplyPurge, error)
	Details(ctx context.Context, in *QueryDetails, opts ...grpc.CallOption) (*ReplyDetails, error)
	InspectTransfer(ctx context.Context, in *QueryInspectTransfer, opts ...grpc.CallOption) (*ReplyInspectTransfer, error)
	EmbedTransfer(ctx context.Context, in *QueryEmbedTransfer, opts ...grpc.CallOption) (*ReplyEmbedTransfer, error)
	HarvestTrans(ctx context.Context, in *QueryHarvestTrans, opts ...grpc.CallOption) (*ReplyHarvestTrans, error)
	Inquire(ctx context.Context, in *QueryInquire, opts ...grpc.CallOption) (*ReplyInquire, error)
	Endorse(ctx context.Context, in *QueryEndorse, opts ...grpc.CallOption) (*ReplyEndorse, error)
	InitSeries(ctx context.Context, in *QueryInitSeries, opts ...grpc.CallOption) (*ReplyInitSeries, error)
	CatalogMirrors(ctx context.Context, in *QueryCatalogMirrors, opts ...grpc.CallOption) (*ReplyCatalogMirrors, error)
	ProposalMirror(ctx context.Context, in *QueryProposalMirror, opts ...grpc.CallOption) (*ReplyProposalMirror, error)
	ImportMirrorSegment(ctx context.Context, in *QueryImportMirrorSegment, opts ...grpc.CallOption) (*ReplyImportMirrorSegment, error)
	ExecuteMirrorSegment(ctx context.Context, in *QueryExecuteMirrorSegment, opts ...grpc.CallOption) (*ReplyExecuteMirrorSegment, error)
	ArrangeNomination(ctx context.Context, in *QueryArrangeNomination, opts ...grpc.CallOption) (*ReplyArrangeNomination, error)
	HandleNomination(ctx context.Context, in *QueryHandleNomination, opts ...grpc.CallOption) (*ReplyHandleNomination, error)
	ExpandBallot(ctx context.Context, in *QueryExpandBallot, opts ...grpc.CallOption) (*ReplyExpandBallot, error)
	ValidateBallotAddition(ctx context.Context, in *QueryValidateBallotAddition, opts ...grpc.CallOption) (*ReplyValidateBallotAddition, error)
	CompleteLedger(ctx context.Context, in *QueryCompleteLedger, opts ...grpc.CallOption) (*ReplyCompleteLedger, error)
}

type aBCICustomer struct {
	cc grpc1.ClientConn
}

func NewIfaceCustomer(cc grpc1.ClientConn) IfaceCustomer {
	return &aBCICustomer{cc}
}

func (c *aBCICustomer) Replicate(ctx context.Context, in *QueryReverberate, opts ...grpc.CallOption) (*ReplyReverberate, error) {
	out := new(ReplyReverberate)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) Purge(ctx context.Context, in *QueryPurge, opts ...grpc.CallOption) (*ReplyPurge, error) {
	out := new(ReplyPurge)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) Details(ctx context.Context, in *QueryDetails, opts ...grpc.CallOption) (*ReplyDetails, error) {
	out := new(ReplyDetails)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) InspectTransfer(ctx context.Context, in *QueryInspectTransfer, opts ...grpc.CallOption) (*ReplyInspectTransfer, error) {
	out := new(ReplyInspectTransfer)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) EmbedTransfer(ctx context.Context, in *QueryEmbedTransfer, opts ...grpc.CallOption) (*ReplyEmbedTransfer, error) {
	out := new(ReplyEmbedTransfer)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) HarvestTrans(ctx context.Context, in *QueryHarvestTrans, opts ...grpc.CallOption) (*ReplyHarvestTrans, error) {
	out := new(ReplyHarvestTrans)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) Inquire(ctx context.Context, in *QueryInquire, opts ...grpc.CallOption) (*ReplyInquire, error) {
	out := new(ReplyInquire)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) Endorse(ctx context.Context, in *QueryEndorse, opts ...grpc.CallOption) (*ReplyEndorse, error) {
	out := new(ReplyEndorse)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) InitSeries(ctx context.Context, in *QueryInitSeries, opts ...grpc.CallOption) (*ReplyInitSeries, error) {
	out := new(ReplyInitSeries)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) CatalogMirrors(ctx context.Context, in *QueryCatalogMirrors, opts ...grpc.CallOption) (*ReplyCatalogMirrors, error) {
	out := new(ReplyCatalogMirrors)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) ProposalMirror(ctx context.Context, in *QueryProposalMirror, opts ...grpc.CallOption) (*ReplyProposalMirror, error) {
	out := new(ReplyProposalMirror)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) ImportMirrorSegment(ctx context.Context, in *QueryImportMirrorSegment, opts ...grpc.CallOption) (*ReplyImportMirrorSegment, error) {
	out := new(ReplyImportMirrorSegment)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) ExecuteMirrorSegment(ctx context.Context, in *QueryExecuteMirrorSegment, opts ...grpc.CallOption) (*ReplyExecuteMirrorSegment, error) {
	out := new(ReplyExecuteMirrorSegment)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) ArrangeNomination(ctx context.Context, in *QueryArrangeNomination, opts ...grpc.CallOption) (*ReplyArrangeNomination, error) {
	out := new(ReplyArrangeNomination)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) HandleNomination(ctx context.Context, in *QueryHandleNomination, opts ...grpc.CallOption) (*ReplyHandleNomination, error) {
	out := new(ReplyHandleNomination)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) ExpandBallot(ctx context.Context, in *QueryExpandBallot, opts ...grpc.CallOption) (*ReplyExpandBallot, error) {
	out := new(ReplyExpandBallot)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) ValidateBallotAddition(ctx context.Context, in *QueryValidateBallotAddition, opts ...grpc.CallOption) (*ReplyValidateBallotAddition, error) {
	out := new(ReplyValidateBallotAddition)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCICustomer) CompleteLedger(ctx context.Context, in *QueryCompleteLedger, opts ...grpc.CallOption) (*ReplyCompleteLedger, error) {
	out := new(ReplyCompleteLedger)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//
type IfaceHost interface {
	Replicate(context.Context, *QueryReverberate) (*ReplyReverberate, error)
	Purge(context.Context, *QueryPurge) (*ReplyPurge, error)
	Details(context.Context, *QueryDetails) (*ReplyDetails, error)
	InspectTransfer(context.Context, *QueryInspectTransfer) (*ReplyInspectTransfer, error)
	EmbedTransfer(context.Context, *QueryEmbedTransfer) (*ReplyEmbedTransfer, error)
	HarvestTrans(context.Context, *QueryHarvestTrans) (*ReplyHarvestTrans, error)
	Inquire(context.Context, *QueryInquire) (*ReplyInquire, error)
	Endorse(context.Context, *QueryEndorse) (*ReplyEndorse, error)
	InitSeries(context.Context, *QueryInitSeries) (*ReplyInitSeries, error)
	CatalogMirrors(context.Context, *QueryCatalogMirrors) (*ReplyCatalogMirrors, error)
	ProposalMirror(context.Context, *QueryProposalMirror) (*ReplyProposalMirror, error)
	ImportMirrorSegment(context.Context, *QueryImportMirrorSegment) (*ReplyImportMirrorSegment, error)
	ExecuteMirrorSegment(context.Context, *QueryExecuteMirrorSegment) (*ReplyExecuteMirrorSegment, error)
	ArrangeNomination(context.Context, *QueryArrangeNomination) (*ReplyArrangeNomination, error)
	HandleNomination(context.Context, *QueryHandleNomination) (*ReplyHandleNomination, error)
	ExpandBallot(context.Context, *QueryExpandBallot) (*ReplyExpandBallot, error)
	ValidateBallotAddition(context.Context, *QueryValidateBallotAddition) (*ReplyValidateBallotAddition, error)
	CompleteLedger(context.Context, *QueryCompleteLedger) (*ReplyCompleteLedger, error)
}

//
type UnexecutedIfaceHost struct {
}

func (*UnexecutedIfaceHost) Replicate(ctx context.Context, req *QueryReverberate) (*ReplyReverberate, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) Purge(ctx context.Context, req *QueryPurge) (*ReplyPurge, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) Details(ctx context.Context, req *QueryDetails) (*ReplyDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) InspectTransfer(ctx context.Context, req *QueryInspectTransfer) (*ReplyInspectTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) EmbedTransfer(ctx context.Context, req *QueryEmbedTransfer) (*ReplyEmbedTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) HarvestTrans(ctx context.Context, req *QueryHarvestTrans) (*ReplyHarvestTrans, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) Inquire(ctx context.Context, req *QueryInquire) (*ReplyInquire, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) Endorse(ctx context.Context, req *QueryEndorse) (*ReplyEndorse, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) InitSeries(ctx context.Context, req *QueryInitSeries) (*ReplyInitSeries, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) CatalogMirrors(ctx context.Context, req *QueryCatalogMirrors) (*ReplyCatalogMirrors, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) ProposalMirror(ctx context.Context, req *QueryProposalMirror) (*ReplyProposalMirror, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) ImportMirrorSegment(ctx context.Context, req *QueryImportMirrorSegment) (*ReplyImportMirrorSegment, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) ExecuteMirrorSegment(ctx context.Context, req *QueryExecuteMirrorSegment) (*ReplyExecuteMirrorSegment, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) ArrangeNomination(ctx context.Context, req *QueryArrangeNomination) (*ReplyArrangeNomination, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) HandleNomination(ctx context.Context, req *QueryHandleNomination) (*ReplyHandleNomination, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) ExpandBallot(ctx context.Context, req *QueryExpandBallot) (*ReplyExpandBallot, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) ValidateBallotAddition(ctx context.Context, req *QueryValidateBallotAddition) (*ReplyValidateBallotAddition, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedIfaceHost) CompleteLedger(ctx context.Context, req *QueryCompleteLedger) (*ReplyCompleteLedger, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}

func EnrollIfaceHost(s grpc1.Server, srv IfaceHost) {
	s.RegisterService(&_IFACE_servicedefinition, srv)
}

func _IFACE_Reverberate_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReverberate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).Replicate(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).Replicate(ctx, req.(*QueryReverberate))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Purge_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPurge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).Purge(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).Purge(ctx, req.(*QueryPurge))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Details_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).Details(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).Details(ctx, req.(*QueryDetails))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Transfercheck_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInspectTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).InspectTransfer(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).InspectTransfer(ctx, req.(*QueryInspectTransfer))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Transferinsert_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEmbedTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).EmbedTransfer(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).EmbedTransfer(ctx, req.(*QueryEmbedTransfer))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Reaptransfers_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHarvestTrans)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).HarvestTrans(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).HarvestTrans(ctx, req.(*QueryHarvestTrans))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Inquire_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInquire)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).Inquire(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).Inquire(ctx, req.(*QueryInquire))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Endorse_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEndorse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).Endorse(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).Endorse(ctx, req.(*QueryEndorse))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Initiatechain_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInitSeries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).InitSeries(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).InitSeries(ctx, req.(*QueryInitSeries))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Catalogmirrors_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCatalogMirrors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).CatalogMirrors(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).CatalogMirrors(ctx, req.(*QueryCatalogMirrors))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Mirrorsnapshot_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalMirror)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).ProposalMirror(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).ProposalMirror(ctx, req.(*QueryProposalMirror))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Loadmirrorsegment_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryImportMirrorSegment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).ImportMirrorSegment(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).ImportMirrorSegment(ctx, req.(*QueryImportMirrorSegment))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Executemirrorsegment_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExecuteMirrorSegment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).ExecuteMirrorSegment(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).ExecuteMirrorSegment(ctx, req.(*QueryExecuteMirrorSegment))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Arrangenomination_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryArrangeNomination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).ArrangeNomination(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).ArrangeNomination(ctx, req.(*QueryArrangeNomination))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Processnomination_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHandleNomination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).HandleNomination(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).HandleNomination(ctx, req.(*QueryHandleNomination))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Ballotextend_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExpandBallot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).ExpandBallot(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).ExpandBallot(ctx, req.(*QueryExpandBallot))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Validateballotextension_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidateBallotAddition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).ValidateBallotAddition(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).ValidateBallotAddition(ctx, req.(*QueryValidateBallotAddition))
	}
	return overseer(ctx, in, details, manager)
}

func _IFACE_Terminateblock_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCompleteLedger)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceHost).CompleteLedger(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceHost).CompleteLedger(ctx, req.(*QueryCompleteLedger))
	}
	return overseer(ctx, in, details, manager)
}

var IFACE_servicedefinition = _IFACE_servicedefinition
var _IFACE_servicedefinition = grpc.ServiceDesc{
	ServiceName: "REDACTED",
	HandlerType: (*IfaceHost)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Reverberate_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Purge_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Details_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Transfercheck_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Transferinsert_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Reaptransfers_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Inquire_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Endorse_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Initiatechain_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Catalogmirrors_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Mirrorsnapshot_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Loadmirrorsegment_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Executemirrorsegment_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Arrangenomination_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Processnomination_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Ballotextend_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Validateballotextension_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Terminateblock_Manager,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "REDACTED",
}

func (m *Query) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Query) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Item != nil {
		{
			volume := m.Item.Volume()
			i -= volume
			if _, err := m.Item.SerializeTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Query_Reverberate) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Reverberate) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Replicate != nil {
		{
			volume, err := m.Replicate.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Query_Purge) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Purge) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Purge != nil {
		{
			volume, err := m.Purge.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Query_Details) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Details) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Details != nil {
		{
			volume, err := m.Details.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Query_Initiatechain) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Initiatechain) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InitSeries != nil {
		{
			volume, err := m.InitSeries.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Query_Inquire) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Inquire) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Inquire != nil {
		{
			volume, err := m.Inquire.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Query_Transfercheck) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Transfercheck) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InspectTransfer != nil {
		{
			volume, err := m.InspectTransfer.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Query_Endorse) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Endorse) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
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
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *Query_Catalogmirrors) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Catalogmirrors) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CatalogMirrors != nil {
		{
			volume, err := m.CatalogMirrors.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func (m *Query_Mirrorsnapshot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Mirrorsnapshot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProposalMirror != nil {
		{
			volume, err := m.ProposalMirror.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x6a
	}
	return len(dAtA) - i, nil
}
func (m *Query_Loadmirrorsegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Loadmirrorsegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ImportMirrorSegment != nil {
		{
			volume, err := m.ImportMirrorSegment.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x72
	}
	return len(dAtA) - i, nil
}
func (m *Query_Executemirrorsegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Executemirrorsegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExecuteMirrorSegment != nil {
		{
			volume, err := m.ExecuteMirrorSegment.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x7a
	}
	return len(dAtA) - i, nil
}
func (m *Query_Arrangenomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Arrangenomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ArrangeNomination != nil {
		{
			volume, err := m.ArrangeNomination.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	return len(dAtA) - i, nil
}
func (m *Query_Processnomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Processnomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HandleNomination != nil {
		{
			volume, err := m.HandleNomination.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	return len(dAtA) - i, nil
}
func (m *Query_Ballotextend) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Ballotextend) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExpandBallot != nil {
		{
			volume, err := m.ExpandBallot.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	return len(dAtA) - i, nil
}
func (m *Query_Validateballotextension) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Validateballotextension) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ValidateBallotAddition != nil {
		{
			volume, err := m.ValidateBallotAddition.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	return len(dAtA) - i, nil
}
func (m *Query_Terminateblock) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Terminateblock) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CompleteLedger != nil {
		{
			volume, err := m.CompleteLedger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	return len(dAtA) - i, nil
}
func (m *Query_Transferinsert) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Transferinsert) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EmbedTransfer != nil {
		{
			volume, err := m.EmbedTransfer.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	return len(dAtA) - i, nil
}
func (m *Query_Reaptransfers) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Query_Reaptransfers) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HarvestTrans != nil {
		{
			volume, err := m.HarvestTrans.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb2
	}
	return len(dAtA) - i, nil
}
func (m *QueryReverberate) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryReverberate) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryReverberate) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signal) > 0 {
		i -= len(m.Signal)
		copy(dAtA[i:], m.Signal)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Signal)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPurge) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPurge) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryPurge) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IfaceRelease) > 0 {
		i -= len(m.IfaceRelease)
		copy(dAtA[i:], m.IfaceRelease)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.IfaceRelease)))
		i--
		dAtA[i] = 0x22
	}
	if m.P2PRelease != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.P2PRelease))
		i--
		dAtA[i] = 0x18
	}
	if m.LedgerRelease != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.LedgerRelease))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Release) > 0 {
		i -= len(m.Release)
		copy(dAtA[i:], m.Release)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Release)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryInitSeries) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInitSeries) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryInitSeries) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PrimaryLevel != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.PrimaryLevel))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ApplicationStatusOctets) > 0 {
		i -= len(m.ApplicationStatusOctets)
		copy(dAtA[i:], m.ApplicationStatusOctets)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ApplicationStatusOctets)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Ratifiers) > 0 {
		for idxNdEx := len(m.Ratifiers) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Ratifiers[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
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
	if m.AgreementOptions != nil {
		{
			volume, err := m.AgreementOptions.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SeriesUid) > 0 {
		i -= len(m.SeriesUid)
		copy(dAtA[i:], m.SeriesUid)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.SeriesUid)))
		i--
		dAtA[i] = 0x12
	}
	n20, fault20 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if fault20 != nil {
		return 0, fault20
	}
	i -= n20
	i = formatVariableintKinds(dAtA, i, uint64(n20))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryInquire) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInquire) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryInquire) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Demonstrate {
		i--
		if m.Demonstrate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Route) > 0 {
		i -= len(m.Route)
		copy(dAtA[i:], m.Route)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Route)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryInspectTransfer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInspectTransfer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryInspectTransfer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEmbedTransfer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEmbedTransfer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryEmbedTransfer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHarvestTrans) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHarvestTrans) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryHarvestTrans) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaximumFuel != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.MaximumFuel))
		i--
		dAtA[i] = 0x10
	}
	if m.MaximumOctets != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.MaximumOctets))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEndorse) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEndorse) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryEndorse) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCatalogMirrors) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCatalogMirrors) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryCatalogMirrors) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryProposalMirror) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProposalMirror) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryProposalMirror) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApplicationDigest) > 0 {
		i -= len(m.ApplicationDigest)
		copy(dAtA[i:], m.ApplicationDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ApplicationDigest)))
		i--
		dAtA[i] = 0x12
	}
	if m.Mirror != nil {
		{
			volume, err := m.Mirror.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *QueryImportMirrorSegment) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryImportMirrorSegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryImportMirrorSegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Segment != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Segment))
		i--
		dAtA[i] = 0x18
	}
	if m.Layout != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Layout))
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

func (m *QueryExecuteMirrorSegment) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryExecuteMirrorSegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryExecuteMirrorSegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Emitter) > 0 {
		i -= len(m.Emitter)
		copy(dAtA[i:], m.Emitter)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Emitter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Segment) > 0 {
		i -= len(m.Segment)
		copy(dAtA[i:], m.Segment)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Segment)))
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

func (m *QueryArrangeNomination) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryArrangeNomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryArrangeNomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecommenderLocation) > 0 {
		i -= len(m.RecommenderLocation)
		copy(dAtA[i:], m.RecommenderLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RecommenderLocation)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FollowingRatifiersDigest) > 0 {
		i -= len(m.FollowingRatifiersDigest)
		copy(dAtA[i:], m.FollowingRatifiersDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FollowingRatifiersDigest)))
		i--
		dAtA[i] = 0x3a
	}
	n22, fault22 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if fault22 != nil {
		return 0, fault22
	}
	i -= n22
	i = formatVariableintKinds(dAtA, i, uint64(n22))
	i--
	dAtA[i] = 0x32
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Malpractice) > 0 {
		for idxNdEx := len(m.Malpractice) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Malpractice[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
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
		volume, err := m.NativeFinalEndorse.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Txs) > 0 {
		for idxNdEx := len(m.Txs) - 1; idxNdEx >= 0; idxNdEx-- {
			i -= len(m.Txs[idxNdEx])
			copy(dAtA[i:], m.Txs[idxNdEx])
			i = formatVariableintKinds(dAtA, i, uint64(len(m.Txs[idxNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.MaximumTransferOctets != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.MaximumTransferOctets))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryHandleNomination) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHandleNomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryHandleNomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecommenderLocation) > 0 {
		i -= len(m.RecommenderLocation)
		copy(dAtA[i:], m.RecommenderLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RecommenderLocation)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FollowingRatifiersDigest) > 0 {
		i -= len(m.FollowingRatifiersDigest)
		copy(dAtA[i:], m.FollowingRatifiersDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FollowingRatifiersDigest)))
		i--
		dAtA[i] = 0x3a
	}
	n24, fault24 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if fault24 != nil {
		return 0, fault24
	}
	i -= n24
	i = formatVariableintKinds(dAtA, i, uint64(n24))
	i--
	dAtA[i] = 0x32
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Malpractice) > 0 {
		for idxNdEx := len(m.Malpractice) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Malpractice[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
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
	{
		volume, err := m.NominatedFinalEndorse.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x12
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

func (m *QueryExpandBallot) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryExpandBallot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryExpandBallot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecommenderLocation) > 0 {
		i -= len(m.RecommenderLocation)
		copy(dAtA[i:], m.RecommenderLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RecommenderLocation)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FollowingRatifiersDigest) > 0 {
		i -= len(m.FollowingRatifiersDigest)
		copy(dAtA[i:], m.FollowingRatifiersDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FollowingRatifiersDigest)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Malpractice) > 0 {
		for idxNdEx := len(m.Malpractice) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Malpractice[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		volume, err := m.NominatedFinalEndorse.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Txs) > 0 {
		for idxNdEx := len(m.Txs) - 1; idxNdEx >= 0; idxNdEx-- {
			i -= len(m.Txs[idxNdEx])
			copy(dAtA[i:], m.Txs[idxNdEx])
			i = formatVariableintKinds(dAtA, i, uint64(len(m.Txs[idxNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	n27, fault27 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if fault27 != nil {
		return 0, fault27
	}
	i -= n27
	i = formatVariableintKinds(dAtA, i, uint64(n27))
	i--
	dAtA[i] = 0x1a
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidateBallotAddition) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidateBallotAddition) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryValidateBallotAddition) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BallotAddition) > 0 {
		i -= len(m.BallotAddition)
		copy(dAtA[i:], m.BallotAddition)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.BallotAddition)))
		i--
		dAtA[i] = 0x22
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RatifierLocation) > 0 {
		i -= len(m.RatifierLocation)
		copy(dAtA[i:], m.RatifierLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RatifierLocation)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCompleteLedger) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCompleteLedger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryCompleteLedger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecommenderLocation) > 0 {
		i -= len(m.RecommenderLocation)
		copy(dAtA[i:], m.RecommenderLocation)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.RecommenderLocation)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FollowingRatifiersDigest) > 0 {
		i -= len(m.FollowingRatifiersDigest)
		copy(dAtA[i:], m.FollowingRatifiersDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FollowingRatifiersDigest)))
		i--
		dAtA[i] = 0x3a
	}
	n28, fault28 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if fault28 != nil {
		return 0, fault28
	}
	i -= n28
	i = formatVariableintKinds(dAtA, i, uint64(n28))
	i--
	dAtA[i] = 0x32
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Malpractice) > 0 {
		for idxNdEx := len(m.Malpractice) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Malpractice[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
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
	{
		volume, err := m.ResolvedFinalEndorse.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x12
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

func (m *Reply) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Item != nil {
		{
			volume := m.Item.Volume()
			i -= volume
			if _, err := m.Item.SerializeTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Reply_Exemption) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Exemption) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Exemption != nil {
		{
			volume, err := m.Exemption.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Reply_Reverberate) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Reverberate) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Replicate != nil {
		{
			volume, err := m.Replicate.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Reply_Purge) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Purge) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Purge != nil {
		{
			volume, err := m.Purge.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Reply_Details) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Details) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Details != nil {
		{
			volume, err := m.Details.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Reply_Initiatechain) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Initiatechain) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InitSeries != nil {
		{
			volume, err := m.InitSeries.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Reply_Inquire) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Inquire) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Inquire != nil {
		{
			volume, err := m.Inquire.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Reply_Transfercheck) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Transfercheck) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InspectTransfer != nil {
		{
			volume, err := m.InspectTransfer.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Reply_Endorse) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Endorse) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
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
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Catalogmirrors) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Catalogmirrors) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CatalogMirrors != nil {
		{
			volume, err := m.CatalogMirrors.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x6a
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Mirrorsnapshot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Mirrorsnapshot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProposalMirror != nil {
		{
			volume, err := m.ProposalMirror.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x72
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Loadmirrorsegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Loadmirrorsegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ImportMirrorSegment != nil {
		{
			volume, err := m.ImportMirrorSegment.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x7a
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Executemirrorsegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Executemirrorsegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExecuteMirrorSegment != nil {
		{
			volume, err := m.ExecuteMirrorSegment.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Arrangenomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Arrangenomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ArrangeNomination != nil {
		{
			volume, err := m.ArrangeNomination.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Processnomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Processnomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HandleNomination != nil {
		{
			volume, err := m.HandleNomination.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Ballotextend) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Ballotextend) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExpandBallot != nil {
		{
			volume, err := m.ExpandBallot.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Validateballotextension) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Validateballotextension) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ValidateBallotAddition != nil {
		{
			volume, err := m.ValidateBallotAddition.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Terminateblock) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Terminateblock) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CompleteLedger != nil {
		{
			volume, err := m.CompleteLedger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Transferinsert) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Transferinsert) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EmbedTransfer != nil {
		{
			volume, err := m.EmbedTransfer.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb2
	}
	return len(dAtA) - i, nil
}
func (m *Reply_Reaptransfers) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Reply_Reaptransfers) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HarvestTrans != nil {
		{
			volume, err := m.HarvestTrans.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xba
	}
	return len(dAtA) - i, nil
}
func (m *ReplyExemption) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyExemption) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyExemption) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fault) > 0 {
		i -= len(m.Fault)
		copy(dAtA[i:], m.Fault)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Fault)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReplyReverberate) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyReverberate) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyReverberate) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signal) > 0 {
		i -= len(m.Signal)
		copy(dAtA[i:], m.Signal)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Signal)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReplyPurge) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyPurge) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyPurge) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ReplyDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinalLedgerApplicationDigest) > 0 {
		i -= len(m.FinalLedgerApplicationDigest)
		copy(dAtA[i:], m.FinalLedgerApplicationDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.FinalLedgerApplicationDigest)))
		i--
		dAtA[i] = 0x2a
	}
	if m.FinalLedgerLevel != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FinalLedgerLevel))
		i--
		dAtA[i] = 0x20
	}
	if m.ApplicationRelease != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.ApplicationRelease))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Release) > 0 {
		i -= len(m.Release)
		copy(dAtA[i:], m.Release)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Release)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReplyInitSeries) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyInitSeries) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyInitSeries) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApplicationDigest) > 0 {
		i -= len(m.ApplicationDigest)
		copy(dAtA[i:], m.ApplicationDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ApplicationDigest)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Ratifiers) > 0 {
		for idxNdEx := len(m.Ratifiers) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Ratifiers[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.AgreementOptions != nil {
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
	}
	return len(dAtA) - i, nil
}

func (m *ReplyInquire) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyInquire) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyInquire) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codex) > 0 {
		i -= len(m.Codex)
		copy(dAtA[i:], m.Codex)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Codex)))
		i--
		dAtA[i] = 0x52
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x48
	}
	if m.EvidenceActions != nil {
		{
			volume, err := m.EvidenceActions.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Item) > 0 {
		i -= len(m.Item)
		copy(dAtA[i:], m.Item)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Item)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x32
	}
	if m.Ordinal != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ordinal))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplyInspectTransfer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyInspectTransfer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyInspectTransfer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codex) > 0 {
		i -= len(m.Codex)
		copy(dAtA[i:], m.Codex)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Codex)))
		i--
		dAtA[i] = 0x42
	}
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
			dAtA[i] = 0x3a
		}
	}
	if m.FuelApplied != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FuelApplied))
		i--
		dAtA[i] = 0x30
	}
	if m.FuelDesired != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FuelDesired))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Log)))
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
	if m.Code != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplyEmbedTransfer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyEmbedTransfer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyEmbedTransfer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplyHarvestTrans) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyHarvestTrans) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyHarvestTrans) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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

func (m *ReplyEndorse) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyEndorse) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyEndorse) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PreserveLevel != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.PreserveLevel))
		i--
		dAtA[i] = 0x18
	}
	return len(dAtA) - i, nil
}

func (m *ReplyCatalogMirrors) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyCatalogMirrors) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyCatalogMirrors) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Mirrors) > 0 {
		for idxNdEx := len(m.Mirrors) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Mirrors[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *ReplyProposalMirror) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyProposalMirror) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyProposalMirror) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Outcome != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Outcome))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplyImportMirrorSegment) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyImportMirrorSegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyImportMirrorSegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Segment) > 0 {
		i -= len(m.Segment)
		copy(dAtA[i:], m.Segment)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Segment)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReplyExecuteMirrorSegment) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyExecuteMirrorSegment) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyExecuteMirrorSegment) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeclineEmitters) > 0 {
		for idxNdEx := len(m.DeclineEmitters) - 1; idxNdEx >= 0; idxNdEx-- {
			i -= len(m.DeclineEmitters[idxNdEx])
			copy(dAtA[i:], m.DeclineEmitters[idxNdEx])
			i = formatVariableintKinds(dAtA, i, uint64(len(m.DeclineEmitters[idxNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ReacquireSegments) > 0 {
		dAtA52 := make([]byte, len(m.ReacquireSegments)*10)
		var j51 int
		for _, num := range m.ReacquireSegments {
			for num >= 1<<7 {
				dAtA52[j51] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j51++
			}
			dAtA52[j51] = uint8(num)
			j51++
		}
		i -= j51
		copy(dAtA[i:], dAtA52[:j51])
		i = formatVariableintKinds(dAtA, i, uint64(j51))
		i--
		dAtA[i] = 0x12
	}
	if m.Outcome != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Outcome))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplyArrangeNomination) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyArrangeNomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyArrangeNomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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

func (m *ReplyHandleNomination) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyHandleNomination) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyHandleNomination) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplyExpandBallot) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyExpandBallot) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyExpandBallot) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BallotAddition) > 0 {
		i -= len(m.BallotAddition)
		copy(dAtA[i:], m.BallotAddition)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.BallotAddition)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReplyValidateBallotAddition) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyValidateBallotAddition) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyValidateBallotAddition) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplyCompleteLedger) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyCompleteLedger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReplyCompleteLedger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApplicationDigest) > 0 {
		i -= len(m.ApplicationDigest)
		copy(dAtA[i:], m.ApplicationDigest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.ApplicationDigest)))
		i--
		dAtA[i] = 0x2a
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
		dAtA[i] = 0x22
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
			dAtA[i] = 0x1a
		}
	}
	if len(m.TransOutcomes) > 0 {
		for idxNdEx := len(m.TransOutcomes) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.TransOutcomes[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x12
		}
	}
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

func (m *EndorseDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndorseDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *EndorseDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ballots) > 0 {
		for idxNdEx := len(m.Ballots) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Ballots[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExpandedEndorseDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpandedEndorseDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ExpandedEndorseDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ballots) > 0 {
		for idxNdEx := len(m.Ballots) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Ballots[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Cycle != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Event) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Event) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Properties) > 0 {
		for idxNdEx := len(m.Properties) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Properties[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = formatVariableintKinds(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventProperty) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventProperty) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *EventProperty) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ordinal {
		i--
		if m.Ordinal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Item) > 0 {
		i -= len(m.Item)
		copy(dAtA[i:], m.Item)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Item)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InvokeTransferOutcome) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InvokeTransferOutcome) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *InvokeTransferOutcome) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codex) > 0 {
		i -= len(m.Codex)
		copy(dAtA[i:], m.Codex)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Codex)))
		i--
		dAtA[i] = 0x42
	}
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
			dAtA[i] = 0x3a
		}
	}
	if m.FuelApplied != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FuelApplied))
		i--
		dAtA[i] = 0x30
	}
	if m.FuelDesired != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.FuelDesired))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Log)))
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
	if m.Code != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TransOutcome) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransOutcome) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *TransOutcome) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		volume, err := m.Outcome.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Ordinal != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ordinal))
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
	if m.Energy != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Energy))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RatifierModify) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RatifierModify) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *RatifierModify) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Energy != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Energy))
		i--
		dAtA[i] = 0x10
	}
	{
		volume, err := m.PublicKey.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *BallotDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BallotDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *BallotDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LedgerUidMark != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.LedgerUidMark))
		i--
		dAtA[i] = 0x18
	}
	{
		volume, err := m.Ratifier.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *ExpandedBallotDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpandedBallotDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ExpandedBallotDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LedgerUidMark != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.LedgerUidMark))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AdditionAutograph) > 0 {
		i -= len(m.AdditionAutograph)
		copy(dAtA[i:], m.AdditionAutograph)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.AdditionAutograph)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BallotAddition) > 0 {
		i -= len(m.BallotAddition)
		copy(dAtA[i:], m.BallotAddition)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.BallotAddition)))
		i--
		dAtA[i] = 0x1a
	}
	{
		volume, err := m.Ratifier.SerializeToDimensionedBuffer(dAtA[:i])
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

func (m *Malpractice) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Malpractice) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Malpractice) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SumPollingEnergy != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.SumPollingEnergy))
		i--
		dAtA[i] = 0x28
	}
	n58, fault58 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if fault58 != nil {
		return 0, fault58
	}
	i -= n58
	i = formatVariableintKinds(dAtA, i, uint64(n58))
	i--
	dAtA[i] = 0x22
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	{
		volume, err := m.Ratifier.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = formatVariableintKinds(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x12
	if m.Kind != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Mirror) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mirror) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Mirror) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metainfo) > 0 {
		i -= len(m.Metainfo)
		copy(dAtA[i:], m.Metainfo)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Metainfo)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x22
	}
	if m.Segments != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Segments))
		i--
		dAtA[i] = 0x18
	}
	if m.Layout != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Layout))
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
func (m *Query) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Item != nil {
		n += m.Item.Volume()
	}
	return n
}

func (m *Query_Reverberate) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Replicate != nil {
		l = m.Replicate.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Purge) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Purge != nil {
		l = m.Purge.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Details) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Details != nil {
		l = m.Details.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Initiatechain) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitSeries != nil {
		l = m.InitSeries.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Inquire) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Inquire != nil {
		l = m.Inquire.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Transfercheck) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InspectTransfer != nil {
		l = m.InspectTransfer.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Endorse) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Endorse != nil {
		l = m.Endorse.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Catalogmirrors) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CatalogMirrors != nil {
		l = m.CatalogMirrors.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Mirrorsnapshot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalMirror != nil {
		l = m.ProposalMirror.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Loadmirrorsegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ImportMirrorSegment != nil {
		l = m.ImportMirrorSegment.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Executemirrorsegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecuteMirrorSegment != nil {
		l = m.ExecuteMirrorSegment.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Arrangenomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ArrangeNomination != nil {
		l = m.ArrangeNomination.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Processnomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HandleNomination != nil {
		l = m.HandleNomination.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Ballotextend) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpandBallot != nil {
		l = m.ExpandBallot.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Validateballotextension) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidateBallotAddition != nil {
		l = m.ValidateBallotAddition.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Terminateblock) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CompleteLedger != nil {
		l = m.CompleteLedger.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Transferinsert) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EmbedTransfer != nil {
		l = m.EmbedTransfer.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Query_Reaptransfers) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HarvestTrans != nil {
		l = m.HarvestTrans.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *QueryReverberate) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signal)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryPurge) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Release)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.LedgerRelease != 0 {
		n += 1 + sovKinds(uint64(m.LedgerRelease))
	}
	if m.P2PRelease != 0 {
		n += 1 + sovKinds(uint64(m.P2PRelease))
	}
	l = len(m.IfaceRelease)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryInitSeries) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.SeriesUid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.AgreementOptions != nil {
		l = m.AgreementOptions.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if len(m.Ratifiers) > 0 {
		for _, e := range m.Ratifiers {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.ApplicationStatusOctets)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.PrimaryLevel != 0 {
		n += 1 + sovKinds(uint64(m.PrimaryLevel))
	}
	return n
}

func (m *QueryInquire) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Route)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Demonstrate {
		n += 2
	}
	return n
}

func (m *QueryInspectTransfer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	return n
}

func (m *QueryEmbedTransfer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryHarvestTrans) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumOctets != 0 {
		n += 1 + sovKinds(uint64(m.MaximumOctets))
	}
	if m.MaximumFuel != 0 {
		n += 1 + sovKinds(uint64(m.MaximumFuel))
	}
	return n
}

func (m *QueryEndorse) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCatalogMirrors) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryProposalMirror) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mirror != nil {
		l = m.Mirror.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.ApplicationDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryImportMirrorSegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Layout != 0 {
		n += 1 + sovKinds(uint64(m.Layout))
	}
	if m.Segment != 0 {
		n += 1 + sovKinds(uint64(m.Segment))
	}
	return n
}

func (m *QueryExecuteMirrorSegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	l = len(m.Segment)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Emitter)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryArrangeNomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumTransferOctets != 0 {
		n += 1 + sovKinds(uint64(m.MaximumTransferOctets))
	}
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = m.NativeFinalEndorse.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FollowingRatifiersDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RecommenderLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryHandleNomination) Volume() (n int) {
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
	l = m.NominatedFinalEndorse.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FollowingRatifiersDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RecommenderLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryExpandBallot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = m.NominatedFinalEndorse.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.FollowingRatifiersDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RecommenderLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryValidateBallotAddition) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RatifierLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = len(m.BallotAddition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *QueryCompleteLedger) Volume() (n int) {
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
	l = m.ResolvedFinalEndorse.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FollowingRatifiersDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.RecommenderLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Reply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Item != nil {
		n += m.Item.Volume()
	}
	return n
}

func (m *Reply_Exemption) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exemption != nil {
		l = m.Exemption.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Reverberate) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Replicate != nil {
		l = m.Replicate.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Purge) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Purge != nil {
		l = m.Purge.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Details) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Details != nil {
		l = m.Details.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Initiatechain) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitSeries != nil {
		l = m.InitSeries.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Inquire) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Inquire != nil {
		l = m.Inquire.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Transfercheck) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InspectTransfer != nil {
		l = m.InspectTransfer.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Endorse) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Endorse != nil {
		l = m.Endorse.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Catalogmirrors) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CatalogMirrors != nil {
		l = m.CatalogMirrors.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Mirrorsnapshot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalMirror != nil {
		l = m.ProposalMirror.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Loadmirrorsegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ImportMirrorSegment != nil {
		l = m.ImportMirrorSegment.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Executemirrorsegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecuteMirrorSegment != nil {
		l = m.ExecuteMirrorSegment.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Arrangenomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ArrangeNomination != nil {
		l = m.ArrangeNomination.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Processnomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HandleNomination != nil {
		l = m.HandleNomination.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Ballotextend) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpandBallot != nil {
		l = m.ExpandBallot.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Validateballotextension) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidateBallotAddition != nil {
		l = m.ValidateBallotAddition.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Terminateblock) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CompleteLedger != nil {
		l = m.CompleteLedger.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Transferinsert) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EmbedTransfer != nil {
		l = m.EmbedTransfer.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Reaptransfers) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HarvestTrans != nil {
		l = m.HarvestTrans.Volume()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *ReplyExemption) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fault)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyReverberate) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signal)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyPurge) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ReplyDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Release)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.ApplicationRelease != 0 {
		n += 1 + sovKinds(uint64(m.ApplicationRelease))
	}
	if m.FinalLedgerLevel != 0 {
		n += 1 + sovKinds(uint64(m.FinalLedgerLevel))
	}
	l = len(m.FinalLedgerApplicationDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyInitSeries) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AgreementOptions != nil {
		l = m.AgreementOptions.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if len(m.Ratifiers) > 0 {
		for _, e := range m.Ratifiers {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.ApplicationDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyInquire) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovKinds(uint64(m.Code))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Item)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.EvidenceActions != nil {
		l = m.EvidenceActions.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = len(m.Codex)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyInspectTransfer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovKinds(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FuelDesired != 0 {
		n += 1 + sovKinds(uint64(m.FuelDesired))
	}
	if m.FuelApplied != 0 {
		n += 1 + sovKinds(uint64(m.FuelApplied))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Codex)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyEmbedTransfer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovKinds(uint64(m.Code))
	}
	return n
}

func (m *ReplyHarvestTrans) Volume() (n int) {
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

func (m *ReplyEndorse) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PreserveLevel != 0 {
		n += 1 + sovKinds(uint64(m.PreserveLevel))
	}
	return n
}

func (m *ReplyCatalogMirrors) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Mirrors) > 0 {
		for _, e := range m.Mirrors {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ReplyProposalMirror) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Outcome != 0 {
		n += 1 + sovKinds(uint64(m.Outcome))
	}
	return n
}

func (m *ReplyImportMirrorSegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Segment)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyExecuteMirrorSegment) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Outcome != 0 {
		n += 1 + sovKinds(uint64(m.Outcome))
	}
	if len(m.ReacquireSegments) > 0 {
		l = 0
		for _, e := range m.ReacquireSegments {
			l += sovKinds(uint64(e))
		}
		n += 1 + sovKinds(uint64(l)) + l
	}
	if len(m.DeclineEmitters) > 0 {
		for _, s := range m.DeclineEmitters {
			l = len(s)
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ReplyArrangeNomination) Volume() (n int) {
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

func (m *ReplyHandleNomination) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovKinds(uint64(m.Status))
	}
	return n
}

func (m *ReplyExpandBallot) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BallotAddition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyValidateBallotAddition) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovKinds(uint64(m.Status))
	}
	return n
}

func (m *ReplyCompleteLedger) Volume() (n int) {
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
	if len(m.TransOutcomes) > 0 {
		for _, e := range m.TransOutcomes {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
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
	l = len(m.ApplicationDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *EndorseDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	if len(m.Ballots) > 0 {
		for _, e := range m.Ballots {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ExpandedEndorseDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cycle != 0 {
		n += 1 + sovKinds(uint64(m.Cycle))
	}
	if len(m.Ballots) > 0 {
		for _, e := range m.Ballots {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *Event) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if len(m.Properties) > 0 {
		for _, e := range m.Properties {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *EventProperty) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Item)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Ordinal {
		n += 2
	}
	return n
}

func (m *InvokeTransferOutcome) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovKinds(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.FuelDesired != 0 {
		n += 1 + sovKinds(uint64(m.FuelDesired))
	}
	if m.FuelApplied != 0 {
		n += 1 + sovKinds(uint64(m.FuelApplied))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Volume()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Codex)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *TransOutcome) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = m.Outcome.Volume()
	n += 1 + l + sovKinds(uint64(l))
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
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Energy != 0 {
		n += 1 + sovKinds(uint64(m.Energy))
	}
	return n
}

func (m *RatifierModify) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PublicKey.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.Energy != 0 {
		n += 1 + sovKinds(uint64(m.Energy))
	}
	return n
}

func (m *BallotDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Ratifier.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.LedgerUidMark != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUidMark))
	}
	return n
}

func (m *ExpandedBallotDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Ratifier.Volume()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.BallotAddition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AdditionAutograph)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.LedgerUidMark != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUidMark))
	}
	return n
}

func (m *Malpractice) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	l = m.Ratifier.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovKinds(uint64(l))
	if m.SumPollingEnergy != 0 {
		n += 1 + sovKinds(uint64(m.SumPollingEnergy))
	}
	return n
}

func (m *Mirror) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Layout != 0 {
		n += 1 + sovKinds(uint64(m.Layout))
	}
	if m.Segments != 0 {
		n += 1 + sovKinds(uint64(m.Segments))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Metainfo)
	if l > 0 {
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
func (m *Query) Unserialize(dAtA []byte) error {
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
			v := &QueryReverberate{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Reverberate{v}
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
			v := &QueryPurge{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Purge{v}
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
			v := &QueryDetails{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Details{v}
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
			v := &QueryInitSeries{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Initiatechain{v}
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
			v := &QueryInquire{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Inquire{v}
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
			v := &QueryInspectTransfer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Transfercheck{v}
			idxNdEx = submitOrdinal
		case 11:
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
			v := &QueryEndorse{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Endorse{v}
			idxNdEx = submitOrdinal
		case 12:
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
			v := &QueryCatalogMirrors{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Catalogmirrors{v}
			idxNdEx = submitOrdinal
		case 13:
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
			v := &QueryProposalMirror{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Mirrorsnapshot{v}
			idxNdEx = submitOrdinal
		case 14:
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
			v := &QueryImportMirrorSegment{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Loadmirrorsegment{v}
			idxNdEx = submitOrdinal
		case 15:
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
			v := &QueryExecuteMirrorSegment{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Executemirrorsegment{v}
			idxNdEx = submitOrdinal
		case 16:
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
			v := &QueryArrangeNomination{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Arrangenomination{v}
			idxNdEx = submitOrdinal
		case 17:
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
			v := &QueryHandleNomination{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Processnomination{v}
			idxNdEx = submitOrdinal
		case 18:
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
			v := &QueryExpandBallot{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Ballotextend{v}
			idxNdEx = submitOrdinal
		case 19:
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
			v := &QueryValidateBallotAddition{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Validateballotextension{v}
			idxNdEx = submitOrdinal
		case 20:
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
			v := &QueryCompleteLedger{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Terminateblock{v}
			idxNdEx = submitOrdinal
		case 21:
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
			v := &QueryEmbedTransfer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Transferinsert{v}
			idxNdEx = submitOrdinal
		case 22:
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
			v := &QueryHarvestTrans{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Query_Reaptransfers{v}
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
func (m *QueryReverberate) Unserialize(dAtA []byte) error {
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
			m.Signal = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *QueryPurge) Unserialize(dAtA []byte) error {
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
func (m *QueryDetails) Unserialize(dAtA []byte) error {
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
			m.Release = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.LedgerRelease = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.LedgerRelease |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.P2PRelease = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.P2PRelease |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
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
			m.IfaceRelease = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *QueryInitSeries) Unserialize(dAtA []byte) error {
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.SeriesUid = string(dAtA[idxNdEx:submitOrdinal])
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
			if m.AgreementOptions == nil {
				m.AgreementOptions = &kinds1.AgreementOptions{}
			}
			if err := m.AgreementOptions.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Ratifiers = append(m.Ratifiers, RatifierModify{})
			if err := m.Ratifiers[len(m.Ratifiers)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
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
			m.ApplicationStatusOctets = append(m.ApplicationStatusOctets[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.ApplicationStatusOctets == nil {
				m.ApplicationStatusOctets = []byte{}
			}
			idxNdEx = submitOrdinal
		case 6:
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
func (m *QueryInquire) Unserialize(dAtA []byte) error {
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
			m.Data = append(m.Data[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
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
			m.Route = string(dAtA[idxNdEx:submitOrdinal])
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
			m.Demonstrate = bool(v != 0)
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
func (m *QueryInspectTransfer) Unserialize(dAtA []byte) error {
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
			m.Tx = append(m.Tx[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			idxNdEx = submitOrdinal
		case 2:
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
				m.Kind |= InspectTransferKind(b&0x7F) << displace
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
func (m *QueryEmbedTransfer) Unserialize(dAtA []byte) error {
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
			m.Tx = append(m.Tx[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Tx == nil {
				m.Tx = []byte{}
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
func (m *QueryHarvestTrans) Unserialize(dAtA []byte) error {
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
			m.MaximumOctets = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MaximumOctets |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.MaximumFuel = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MaximumFuel |= uint64(b&0x7F) << displace
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
func (m *QueryEndorse) Unserialize(dAtA []byte) error {
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
func (m *QueryCatalogMirrors) Unserialize(dAtA []byte) error {
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
func (m *QueryProposalMirror) Unserialize(dAtA []byte) error {
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
			if m.Mirror == nil {
				m.Mirror = &Mirror{}
			}
			if err := m.Mirror.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
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
			m.ApplicationDigest = append(m.ApplicationDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.ApplicationDigest == nil {
				m.ApplicationDigest = []byte{}
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
func (m *QueryImportMirrorSegment) Unserialize(dAtA []byte) error {
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
				m.Level |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Layout |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Segment = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Segment |= uint32(b&0x7F) << displace
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
func (m *QueryExecuteMirrorSegment) Unserialize(dAtA []byte) error {
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
			m.Segment = append(m.Segment[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Segment == nil {
				m.Segment = []byte{}
			}
			idxNdEx = submitOrdinal
		case 3:
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
			m.Emitter = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *QueryArrangeNomination) Unserialize(dAtA []byte) error {
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
			m.MaximumTransferOctets = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MaximumTransferOctets |= int64(b&0x7F) << displace
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
			m.Txs = append(m.Txs, make([]byte, submitOrdinal-idxNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[idxNdEx:submitOrdinal])
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
			if err := m.NativeFinalEndorse.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 5:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.FollowingRatifiersDigest = append(m.FollowingRatifiersDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FollowingRatifiersDigest == nil {
				m.FollowingRatifiersDigest = []byte{}
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
func (m *QueryHandleNomination) Unserialize(dAtA []byte) error {
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
			if err := m.NominatedFinalEndorse.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Digest = append(m.Digest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 5:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.FollowingRatifiersDigest = append(m.FollowingRatifiersDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FollowingRatifiersDigest == nil {
				m.FollowingRatifiersDigest = []byte{}
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
func (m *QueryExpandBallot) Unserialize(dAtA []byte) error {
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Txs = append(m.Txs, make([]byte, submitOrdinal-idxNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[idxNdEx:submitOrdinal])
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
			if err := m.NominatedFinalEndorse.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.FollowingRatifiersDigest = append(m.FollowingRatifiersDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FollowingRatifiersDigest == nil {
				m.FollowingRatifiersDigest = []byte{}
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
func (m *QueryValidateBallotAddition) Unserialize(dAtA []byte) error {
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
			m.BallotAddition = append(m.BallotAddition[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.BallotAddition == nil {
				m.BallotAddition = []byte{}
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
func (m *QueryCompleteLedger) Unserialize(dAtA []byte) error {
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
			if err := m.ResolvedFinalEndorse.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Digest = append(m.Digest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 5:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.FollowingRatifiersDigest = append(m.FollowingRatifiersDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FollowingRatifiersDigest == nil {
				m.FollowingRatifiersDigest = []byte{}
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
func (m *Reply) Unserialize(dAtA []byte) error {
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
			v := &ReplyExemption{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Exemption{v}
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
			v := &ReplyReverberate{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Reverberate{v}
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
			v := &ReplyPurge{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Purge{v}
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
			v := &ReplyDetails{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Details{v}
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
			v := &ReplyInitSeries{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Initiatechain{v}
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
			v := &ReplyInquire{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Inquire{v}
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
			v := &ReplyInspectTransfer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Transfercheck{v}
			idxNdEx = submitOrdinal
		case 12:
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
			v := &ReplyEndorse{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Endorse{v}
			idxNdEx = submitOrdinal
		case 13:
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
			v := &ReplyCatalogMirrors{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Catalogmirrors{v}
			idxNdEx = submitOrdinal
		case 14:
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
			v := &ReplyProposalMirror{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Mirrorsnapshot{v}
			idxNdEx = submitOrdinal
		case 15:
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
			v := &ReplyImportMirrorSegment{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Loadmirrorsegment{v}
			idxNdEx = submitOrdinal
		case 16:
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
			v := &ReplyExecuteMirrorSegment{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Executemirrorsegment{v}
			idxNdEx = submitOrdinal
		case 17:
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
			v := &ReplyArrangeNomination{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Arrangenomination{v}
			idxNdEx = submitOrdinal
		case 18:
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
			v := &ReplyHandleNomination{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Processnomination{v}
			idxNdEx = submitOrdinal
		case 19:
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
			v := &ReplyExpandBallot{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Ballotextend{v}
			idxNdEx = submitOrdinal
		case 20:
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
			v := &ReplyValidateBallotAddition{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Validateballotextension{v}
			idxNdEx = submitOrdinal
		case 21:
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
			v := &ReplyCompleteLedger{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Terminateblock{v}
			idxNdEx = submitOrdinal
		case 22:
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
			v := &ReplyEmbedTransfer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Transferinsert{v}
			idxNdEx = submitOrdinal
		case 23:
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
			v := &ReplyHarvestTrans{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Item = &Reply_Reaptransfers{v}
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
func (m *ReplyExemption) Unserialize(dAtA []byte) error {
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
			m.Fault = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *ReplyReverberate) Unserialize(dAtA []byte) error {
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
			m.Signal = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *ReplyPurge) Unserialize(dAtA []byte) error {
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
func (m *ReplyDetails) Unserialize(dAtA []byte) error {
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
			m.Data = string(dAtA[idxNdEx:submitOrdinal])
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
			m.Release = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.ApplicationRelease = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.ApplicationRelease |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
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
			m.FinalLedgerApplicationDigest = append(m.FinalLedgerApplicationDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.FinalLedgerApplicationDigest == nil {
				m.FinalLedgerApplicationDigest = []byte{}
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
func (m *ReplyInitSeries) Unserialize(dAtA []byte) error {
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
			if m.AgreementOptions == nil {
				m.AgreementOptions = &kinds1.AgreementOptions{}
			}
			if err := m.AgreementOptions.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.Ratifiers = append(m.Ratifiers, RatifierModify{})
			if err := m.Ratifiers[len(m.Ratifiers)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 3:
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
func (m *ReplyInquire) Unserialize(dAtA []byte) error {
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
			m.Code = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Code |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
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
			m.Log = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 4:
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
			m.Details = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 5:
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
				m.Ordinal |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
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
			m.Key = append(m.Key[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Key == nil {
				m.Key = []byte{}
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
			m.Item = append(m.Item[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Item == nil {
				m.Item = []byte{}
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
			if m.EvidenceActions == nil {
				m.EvidenceActions = &vault.EvidenceActions{}
			}
			if err := m.EvidenceActions.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 9:
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
		case 10:
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
			m.Codex = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *ReplyInspectTransfer) Unserialize(dAtA []byte) error {
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
			m.Code = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Code |= uint32(b&0x7F) << displace
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
			m.Data = append(m.Data[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			idxNdEx = submitOrdinal
		case 3:
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
			m.Log = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 4:
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
			m.Details = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelDesired = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FuelDesired |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 6:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelApplied = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FuelApplied |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
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
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 8:
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
			m.Codex = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *ReplyEmbedTransfer) Unserialize(dAtA []byte) error {
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
			m.Code = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Code |= uint32(b&0x7F) << displace
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
func (m *ReplyHarvestTrans) Unserialize(dAtA []byte) error {
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
func (m *ReplyEndorse) Unserialize(dAtA []byte) error {
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
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PreserveLevel = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.PreserveLevel |= int64(b&0x7F) << displace
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
func (m *ReplyCatalogMirrors) Unserialize(dAtA []byte) error {
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
			m.Mirrors = append(m.Mirrors, &Mirror{})
			if err := m.Mirrors[len(m.Mirrors)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *ReplyProposalMirror) Unserialize(dAtA []byte) error {
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
			m.Outcome = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Outcome |= Replymirrorsnapshot_Outcome(b&0x7F) << displace
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
func (m *ReplyImportMirrorSegment) Unserialize(dAtA []byte) error {
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
			m.Segment = append(m.Segment[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Segment == nil {
				m.Segment = []byte{}
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
func (m *ReplyExecuteMirrorSegment) Unserialize(dAtA []byte) error {
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
			m.Outcome = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Outcome |= Replyexecutemirrorsegment_Outcome(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind == 0 {
				var v uint32
				for displace := uint(0); ; displace += 7 {
					if displace >= 64 {
						return ErrIntegerOverloadKinds
					}
					if idxNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[idxNdEx]
					idxNdEx++
					v |= uint32(b&0x7F) << displace
					if b < 0x80 {
						break
					}
				}
				m.ReacquireSegments = append(m.ReacquireSegments, v)
			} else if cableKind == 2 {
				var compressedSize int
				for displace := uint(0); ; displace += 7 {
					if displace >= 64 {
						return ErrIntegerOverloadKinds
					}
					if idxNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[idxNdEx]
					idxNdEx++
					compressedSize |= int(b&0x7F) << displace
					if b < 0x80 {
						break
					}
				}
				if compressedSize < 0 {
					return ErrCorruptExtentKinds
				}
				submitOrdinal := idxNdEx + compressedSize
				if submitOrdinal < 0 {
					return ErrCorruptExtentKinds
				}
				if submitOrdinal > l {
					return io.ErrUnexpectedEOF
				}
				var componentCount int
				var tally int
				for _, integer := range dAtA[idxNdEx:submitOrdinal] {
					if integer < 128 {
						tally++
					}
				}
				componentCount = tally
				if componentCount != 0 && len(m.ReacquireSegments) == 0 {
					m.ReacquireSegments = make([]uint32, 0, componentCount)
				}
				for idxNdEx < submitOrdinal {
					var v uint32
					for displace := uint(0); ; displace += 7 {
						if displace >= 64 {
							return ErrIntegerOverloadKinds
						}
						if idxNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[idxNdEx]
						idxNdEx++
						v |= uint32(b&0x7F) << displace
						if b < 0x80 {
							break
						}
					}
					m.ReacquireSegments = append(m.ReacquireSegments, v)
				}
			} else {
				return fmt.Errorf("REDACTED", cableKind)
			}
		case 3:
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
			m.DeclineEmitters = append(m.DeclineEmitters, string(dAtA[idxNdEx:submitOrdinal]))
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
func (m *ReplyArrangeNomination) Unserialize(dAtA []byte) error {
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
func (m *ReplyHandleNomination) Unserialize(dAtA []byte) error {
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
			m.Status = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Status |= Responseprocessnomination_Nominationstate(b&0x7F) << displace
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
func (m *ReplyExpandBallot) Unserialize(dAtA []byte) error {
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
			m.BallotAddition = append(m.BallotAddition[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.BallotAddition == nil {
				m.BallotAddition = []byte{}
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
func (m *ReplyValidateBallotAddition) Unserialize(dAtA []byte) error {
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
			m.Status = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Status |= Responseverifyballotextension_Validatestatus(b&0x7F) << displace
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
func (m *ReplyCompleteLedger) Unserialize(dAtA []byte) error {
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
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.TransOutcomes = append(m.TransOutcomes, &InvokeTransferOutcome{})
			if err := m.TransOutcomes[len(m.TransOutcomes)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			m.RatifierRefreshes = append(m.RatifierRefreshes, RatifierModify{})
			if err := m.RatifierRefreshes[len(m.RatifierRefreshes)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.AgreementArgumentRefreshes == nil {
				m.AgreementArgumentRefreshes = &kinds1.AgreementOptions{}
			}
			if err := m.AgreementArgumentRefreshes.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
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
			m.ApplicationDigest = append(m.ApplicationDigest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.ApplicationDigest == nil {
				m.ApplicationDigest = []byte{}
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
func (m *EndorseDetails) Unserialize(dAtA []byte) error {
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
			m.Ballots = append(m.Ballots, BallotDetails{})
			if err := m.Ballots[len(m.Ballots)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *ExpandedEndorseDetails) Unserialize(dAtA []byte) error {
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
			m.Ballots = append(m.Ballots, ExpandedBallotDetails{})
			if err := m.Ballots[len(m.Ballots)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *Event) Unserialize(dAtA []byte) error {
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
			m.Kind = string(dAtA[idxNdEx:submitOrdinal])
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
			m.Properties = append(m.Properties, EventProperty{})
			if err := m.Properties[len(m.Properties)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *EventProperty) Unserialize(dAtA []byte) error {
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
			m.Key = string(dAtA[idxNdEx:submitOrdinal])
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
			m.Item = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 3:
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
			m.Ordinal = bool(v != 0)
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
func (m *InvokeTransferOutcome) Unserialize(dAtA []byte) error {
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
			m.Code = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Code |= uint32(b&0x7F) << displace
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
			m.Data = append(m.Data[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			idxNdEx = submitOrdinal
		case 3:
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
			m.Log = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 4:
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
			m.Details = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelDesired = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FuelDesired |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 6:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelApplied = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.FuelApplied |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
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
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 8:
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
			m.Codex = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *TransOutcome) Unserialize(dAtA []byte) error {
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
		case 3:
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
			m.Tx = append(m.Tx[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Tx == nil {
				m.Tx = []byte{}
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
			if err := m.Outcome.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *Ratifier) Unserialize(dAtA []byte) error {
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
			m.Location = append(m.Location[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Location == nil {
				m.Location = []byte{}
			}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Energy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Energy |= int64(b&0x7F) << displace
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
func (m *RatifierModify) Unserialize(dAtA []byte) error {
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
			if err := m.PublicKey.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Energy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Energy |= int64(b&0x7F) << displace
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
func (m *BallotDetails) Unserialize(dAtA []byte) error {
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
			if err := m.Ratifier.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 3:
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
				m.LedgerUidMark |= kinds1.LedgerUIDMark(b&0x7F) << displace
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
func (m *ExpandedBallotDetails) Unserialize(dAtA []byte) error {
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
			if err := m.Ratifier.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 3:
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
			m.BallotAddition = append(m.BallotAddition[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.BallotAddition == nil {
				m.BallotAddition = []byte{}
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
			m.AdditionAutograph = append(m.AdditionAutograph[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.AdditionAutograph == nil {
				m.AdditionAutograph = []byte{}
			}
			idxNdEx = submitOrdinal
		case 5:
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
				m.LedgerUidMark |= kinds1.LedgerUIDMark(b&0x7F) << displace
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
func (m *Malpractice) Unserialize(dAtA []byte) error {
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
				m.Kind |= MalpracticeKind(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
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
			if err := m.Ratifier.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
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
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SumPollingEnergy = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
func (m *Mirror) Unserialize(dAtA []byte) error {
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
				m.Level |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Layout |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Segments = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Segments |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
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
			m.Digest = append(m.Digest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
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
			m.Metainfo = append(m.Metainfo[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Metainfo == nil {
				m.Metainfo = []byte{}
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
