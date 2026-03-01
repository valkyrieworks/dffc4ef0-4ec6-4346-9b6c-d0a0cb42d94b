//
//

package kinds

import (
	context "context"
	fmt "fmt"
	security "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	kinds1 "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
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
	Inspecttranskind_Fresh     InspectTransferKind = 0
	Inspecttranskind_Reinspect InspectTransferKind = 1
)

var Inspecttranskind_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
}

var Inspecttranskind_datum = map[string]int32{
	"REDACTED":     0,
	"REDACTED": 1,
}

func (x InspectTransferKind) Text() string {
	return proto.EnumName(Inspecttranskind_alias, int32(x))
}

func (InspectTransferKind) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{0}
}

type MalpracticeKind int32

const (
	Malfunctionkind_UNFAMILIAR             MalpracticeKind = 0
	Malfunctionkind_REPLICATED_BALLOT      MalpracticeKind = 1
	Malfunctionkind_AGILE_CUSTOMER_ONSLAUGHT MalpracticeKind = 2
)

var Malfunctionkind_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
}

var Malfunctionkind_datum = map[string]int32{
	"REDACTED":             0,
	"REDACTED":      1,
	"REDACTED": 2,
}

func (x MalpracticeKind) Text() string {
	return proto.EnumName(Malfunctionkind_alias, int32(x))
}

func (MalpracticeKind) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{1}
}

type Replyextendimage_Outcome int32

const (
	Replyextendimage_UNFAMILIAR       Replyextendimage_Outcome = 0
	Replyextendimage_EMBRACE        Replyextendimage_Outcome = 1
	Replyextendimage_CANCEL         Replyextendimage_Outcome = 2
	Replyextendimage_DECLINE        Replyextendimage_Outcome = 3
	Replyextendimage_DECLINE_LAYOUT Replyextendimage_Outcome = 4
	Replyextendimage_DECLINE_ORIGINATOR Replyextendimage_Outcome = 5
)

var Replyextendimage_Outcome_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
	4: "REDACTED",
	5: "REDACTED",
}

var Replyextendimage_Outcome_datum = map[string]int32{
	"REDACTED":       0,
	"REDACTED":        1,
	"REDACTED":         2,
	"REDACTED":        3,
	"REDACTED": 4,
	"REDACTED": 5,
}

func (x Replyextendimage_Outcome) Text() string {
	return proto.EnumName(Replyextendimage_Outcome_alias, int32(x))
}

func (Replyextendimage_Outcome) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{31, 0}
}

type Replyapplyimagefragment_Outcome int32

const (
	Replyapplyimagefragment_UNFAMILIAR         Replyapplyimagefragment_Outcome = 0
	Replyapplyimagefragment_EMBRACE          Replyapplyimagefragment_Outcome = 1
	Replyapplyimagefragment_CANCEL           Replyapplyimagefragment_Outcome = 2
	Replyapplyimagefragment_REISSUE           Replyapplyimagefragment_Outcome = 3
	Replyapplyimagefragment_REISSUE_IMAGE  Replyapplyimagefragment_Outcome = 4
	Replyapplyimagefragment_DECLINE_IMAGE Replyapplyimagefragment_Outcome = 5
)

var Replyapplyimagefragment_Outcome_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
	4: "REDACTED",
	5: "REDACTED",
}

var Replyapplyimagefragment_Outcome_datum = map[string]int32{
	"REDACTED":         0,
	"REDACTED":          1,
	"REDACTED":           2,
	"REDACTED":           3,
	"REDACTED":  4,
	"REDACTED": 5,
}

func (x Replyapplyimagefragment_Outcome) Text() string {
	return proto.EnumName(Replyapplyimagefragment_Outcome_alias, int32(x))
}

func (Replyapplyimagefragment_Outcome) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{33, 0}
}

type Responseexecuteitem_Itemstatus int32

const (
	Responseexecuteitem_UNFAMILIAR Responseexecuteitem_Itemstatus = 0
	Responseexecuteitem_EMBRACE  Responseexecuteitem_Itemstatus = 1
	Responseexecuteitem_DECLINE  Responseexecuteitem_Itemstatus = 2
)

var Responseexecuteitem_Itemstatus_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
}

var Responseexecuteitem_Itemstatus_datum = map[string]int32{
	"REDACTED": 0,
	"REDACTED":  1,
	"REDACTED":  2,
}

func (x Responseexecuteitem_Itemstatus) Text() string {
	return proto.EnumName(Responseexecuteitem_Itemstatus_alias, int32(x))
}

func (Responseexecuteitem_Itemstatus) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{35, 0}
}

type Responsecertifyballotaddition_Verifystatus int32

const (
	Responsecertifyballotaddition_UNFAMILIAR Responsecertifyballotaddition_Verifystatus = 0
	Responsecertifyballotaddition_EMBRACE  Responsecertifyballotaddition_Verifystatus = 1
	//
	//
	//
	//
	Responsecertifyballotaddition_DECLINE Responsecertifyballotaddition_Verifystatus = 2
)

var Responsecertifyballotaddition_Verifystatus_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
}

var Responsecertifyballotaddition_Verifystatus_datum = map[string]int32{
	"REDACTED": 0,
	"REDACTED":  1,
	"REDACTED":  2,
}

func (x Responsecertifyballotaddition_Verifystatus) Text() string {
	return proto.EnumName(Responsecertifyballotaddition_Verifystatus_alias, int32(x))
}

func (Responsecertifyballotaddition_Verifystatus) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{37, 0}
}

type Solicit struct {
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
	Datum isrequesting_Datum `protobuf_oneof:"datum"`
}

func (m *Solicit) Restore()         { *m = Solicit{} }
func (m *Solicit) Text() string { return proto.CompactTextString(m) }
func (*Solicit) SchemaArtifact()    {}
func (*Solicit) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{0}
}
func (m *Solicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Solicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Solicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Solicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Solicit.Merge(m, src)
}
func (m *Solicit) XXX_Extent() int {
	return m.Extent()
}
func (m *Solicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Solicit.DiscardUnknown(m)
}

var xxx_signaldetails_Solicit proto.InternalMessageInfo

type isrequesting_Datum interface {
	isrequesting_Datum()
	SerializeToward([]byte) (int, error)
	Extent() int
}

type Solicit_Reverberate struct {
	Reverberate *SolicitReverberate `protobuf:"octets,1,opt,name=echo,proto3,oneof" json:"reverberate,omitempty"`
}
type Solicit_Purge struct {
	Purge *SolicitPurge `protobuf:"octets,2,opt,name=flush,proto3,oneof" json:"purge,omitempty"`
}
type Solicit_Details struct {
	Details *SolicitDetails `protobuf:"octets,3,opt,name=info,proto3,oneof" json:"details,omitempty"`
}
type Solicit_Initiatechain struct {
	InitializeSuccession *SolicitInitializeSuccession `protobuf:"octets,5,opt,name=init_chain,json=initChain,proto3,oneof" json:"initialize_succession,omitempty"`
}
type Solicit_Inquire struct {
	Inquire *SolicitInquire `protobuf:"octets,6,opt,name=query,proto3,oneof" json:"inquire,omitempty"`
}
type Solicit_Inspecttrans struct {
	InspectTransfer *SolicitInspectTransfer `protobuf:"octets,8,opt,name=check_tx,json=checkTx,proto3,oneof" json:"inspect_transfer,omitempty"`
}
type Solicit_Endorse struct {
	Endorse *SolicitEndorse `protobuf:"octets,11,opt,name=commit,proto3,oneof" json:"endorse,omitempty"`
}
type Solicit_Catalogimages struct {
	CollectionImages *SolicitCollectionImages `protobuf:"octets,12,opt,name=list_snapshots,json=listSnapshots,proto3,oneof" json:"catalog_images,omitempty"`
}
type Solicit_Extendimage struct {
	ExtendImage *SolicitExtendImage `protobuf:"octets,13,opt,name=offer_snapshot,json=offerSnapshot,proto3,oneof" json:"extend_image,omitempty"`
}
type Solicit_Loadimagefragment struct {
	FetchImageSegment *SolicitFetchImageSegment `protobuf:"octets,14,opt,name=load_snapshot_chunk,json=loadSnapshotChunk,proto3,oneof" json:"fetch_image_segment,omitempty"`
}
type Solicit_Executeimagefragment struct {
	ExecuteImageSegment *SolicitExecuteImageSegment `protobuf:"octets,15,opt,name=apply_snapshot_chunk,json=applySnapshotChunk,proto3,oneof" json:"execute_image_segment,omitempty"`
}
type Solicit_Prepareitem struct {
	ArrangeNomination *SolicitArrangeNomination `protobuf:"octets,16,opt,name=prepare_proposal,json=prepareProposal,proto3,oneof" json:"arrange_nomination,omitempty"`
}
type Solicit_Executeitem struct {
	HandleNomination *SolicitHandleNomination `protobuf:"octets,17,opt,name=process_proposal,json=processProposal,proto3,oneof" json:"handle_nomination,omitempty"`
}
type Solicit_Extendballot struct {
	BroadenBallot *SolicitBroadenBallot `protobuf:"octets,18,opt,name=extend_vote,json=extendVote,proto3,oneof" json:"broaden_ballot,omitempty"`
}
type Solicit_Verifyballotaddition struct {
	ValidateBallotAddition *SolicitValidateBallotAddition `protobuf:"octets,19,opt,name=verify_vote_extension,json=verifyVoteExtension,proto3,oneof" json:"validate_ballot_addition,omitempty"`
}
type Solicit_Finalizeledger struct {
	CulminateLedger *SolicitCulminateLedger `protobuf:"octets,20,opt,name=finalize_block,json=finalizeBlock,proto3,oneof" json:"culminate_ledger,omitempty"`
}
type Solicit_Appendtrans struct {
	AppendTransfer *SolicitAppendTransfer `protobuf:"octets,21,opt,name=insert_tx,json=insertTx,proto3,oneof" json:"append_transfer,omitempty"`
}
type Solicit_Harvesttrans struct {
	HarvestTrans *SolicitHarvestTrans `protobuf:"octets,22,opt,name=reap_txs,json=reapTxs,proto3,oneof" json:"harvest_trans,omitempty"`
}

func (*Solicit_Reverberate) isrequesting_Datum()                {}
func (*Solicit_Purge) isrequesting_Datum()               {}
func (*Solicit_Details) isrequesting_Datum()                {}
func (*Solicit_Initiatechain) isrequesting_Datum()           {}
func (*Solicit_Inquire) isrequesting_Datum()               {}
func (*Solicit_Inspecttrans) isrequesting_Datum()             {}
func (*Solicit_Endorse) isrequesting_Datum()              {}
func (*Solicit_Catalogimages) isrequesting_Datum()       {}
func (*Solicit_Extendimage) isrequesting_Datum()       {}
func (*Solicit_Loadimagefragment) isrequesting_Datum()   {}
func (*Solicit_Executeimagefragment) isrequesting_Datum()  {}
func (*Solicit_Prepareitem) isrequesting_Datum()     {}
func (*Solicit_Executeitem) isrequesting_Datum()     {}
func (*Solicit_Extendballot) isrequesting_Datum()          {}
func (*Solicit_Verifyballotaddition) isrequesting_Datum() {}
func (*Solicit_Finalizeledger) isrequesting_Datum()       {}
func (*Solicit_Appendtrans) isrequesting_Datum()            {}
func (*Solicit_Harvesttrans) isrequesting_Datum()             {}

func (m *Solicit) ObtainDatum() isrequesting_Datum {
	if m != nil {
		return m.Datum
	}
	return nil
}

func (m *Solicit) ObtainReverberate() *SolicitReverberate {
	if x, ok := m.ObtainDatum().(*Solicit_Reverberate); ok {
		return x.Reverberate
	}
	return nil
}

func (m *Solicit) ObtainPurge() *SolicitPurge {
	if x, ok := m.ObtainDatum().(*Solicit_Purge); ok {
		return x.Purge
	}
	return nil
}

func (m *Solicit) ObtainDetails() *SolicitDetails {
	if x, ok := m.ObtainDatum().(*Solicit_Details); ok {
		return x.Details
	}
	return nil
}

func (m *Solicit) ObtainInitializeSuccession() *SolicitInitializeSuccession {
	if x, ok := m.ObtainDatum().(*Solicit_Initiatechain); ok {
		return x.InitializeSuccession
	}
	return nil
}

func (m *Solicit) ObtainInquire() *SolicitInquire {
	if x, ok := m.ObtainDatum().(*Solicit_Inquire); ok {
		return x.Inquire
	}
	return nil
}

func (m *Solicit) ObtainInspectTransfer() *SolicitInspectTransfer {
	if x, ok := m.ObtainDatum().(*Solicit_Inspecttrans); ok {
		return x.InspectTransfer
	}
	return nil
}

func (m *Solicit) ObtainEndorse() *SolicitEndorse {
	if x, ok := m.ObtainDatum().(*Solicit_Endorse); ok {
		return x.Endorse
	}
	return nil
}

func (m *Solicit) ObtainCatalogImages() *SolicitCollectionImages {
	if x, ok := m.ObtainDatum().(*Solicit_Catalogimages); ok {
		return x.CollectionImages
	}
	return nil
}

func (m *Solicit) ObtainExtendImage() *SolicitExtendImage {
	if x, ok := m.ObtainDatum().(*Solicit_Extendimage); ok {
		return x.ExtendImage
	}
	return nil
}

func (m *Solicit) ObtainFetchImageSegment() *SolicitFetchImageSegment {
	if x, ok := m.ObtainDatum().(*Solicit_Loadimagefragment); ok {
		return x.FetchImageSegment
	}
	return nil
}

func (m *Solicit) ObtainExecuteImageSegment() *SolicitExecuteImageSegment {
	if x, ok := m.ObtainDatum().(*Solicit_Executeimagefragment); ok {
		return x.ExecuteImageSegment
	}
	return nil
}

func (m *Solicit) ObtainArrangeNomination() *SolicitArrangeNomination {
	if x, ok := m.ObtainDatum().(*Solicit_Prepareitem); ok {
		return x.ArrangeNomination
	}
	return nil
}

func (m *Solicit) ObtainHandleNomination() *SolicitHandleNomination {
	if x, ok := m.ObtainDatum().(*Solicit_Executeitem); ok {
		return x.HandleNomination
	}
	return nil
}

func (m *Solicit) ObtainBroadenBallot() *SolicitBroadenBallot {
	if x, ok := m.ObtainDatum().(*Solicit_Extendballot); ok {
		return x.BroadenBallot
	}
	return nil
}

func (m *Solicit) ObtainValidateBallotAddition() *SolicitValidateBallotAddition {
	if x, ok := m.ObtainDatum().(*Solicit_Verifyballotaddition); ok {
		return x.ValidateBallotAddition
	}
	return nil
}

func (m *Solicit) ObtainCulminateLedger() *SolicitCulminateLedger {
	if x, ok := m.ObtainDatum().(*Solicit_Finalizeledger); ok {
		return x.CulminateLedger
	}
	return nil
}

func (m *Solicit) ObtainAppendTransfer() *SolicitAppendTransfer {
	if x, ok := m.ObtainDatum().(*Solicit_Appendtrans); ok {
		return x.AppendTransfer
	}
	return nil
}

func (m *Solicit) ObtainHarvestTrans() *SolicitHarvestTrans {
	if x, ok := m.ObtainDatum().(*Solicit_Harvesttrans); ok {
		return x.HarvestTrans
	}
	return nil
}

//
func (*Solicit) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Solicit_Reverberate)(nil),
		(*Solicit_Purge)(nil),
		(*Solicit_Details)(nil),
		(*Solicit_Initiatechain)(nil),
		(*Solicit_Inquire)(nil),
		(*Solicit_Inspecttrans)(nil),
		(*Solicit_Endorse)(nil),
		(*Solicit_Catalogimages)(nil),
		(*Solicit_Extendimage)(nil),
		(*Solicit_Loadimagefragment)(nil),
		(*Solicit_Executeimagefragment)(nil),
		(*Solicit_Prepareitem)(nil),
		(*Solicit_Executeitem)(nil),
		(*Solicit_Extendballot)(nil),
		(*Solicit_Verifyballotaddition)(nil),
		(*Solicit_Finalizeledger)(nil),
		(*Solicit_Appendtrans)(nil),
		(*Solicit_Harvesttrans)(nil),
	}
}

type SolicitReverberate struct {
	Signal string `protobuf:"octets,1,opt,name=message,proto3" json:"signal,omitempty"`
}

func (m *SolicitReverberate) Restore()         { *m = SolicitReverberate{} }
func (m *SolicitReverberate) Text() string { return proto.CompactTextString(m) }
func (*SolicitReverberate) SchemaArtifact()    {}
func (*SolicitReverberate) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{1}
}
func (m *SolicitReverberate) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitReverberate) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestecho.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitReverberate) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestecho.Merge(m, src)
}
func (m *SolicitReverberate) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitReverberate) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestecho.DiscardUnknown(m)
}

var xxx_signaldetails_Requestecho proto.InternalMessageInfo

func (m *SolicitReverberate) ObtainArtifact() string {
	if m != nil {
		return m.Signal
	}
	return "REDACTED"
}

type SolicitPurge struct {
}

func (m *SolicitPurge) Restore()         { *m = SolicitPurge{} }
func (m *SolicitPurge) Text() string { return proto.CompactTextString(m) }
func (*SolicitPurge) SchemaArtifact()    {}
func (*SolicitPurge) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{2}
}
func (m *SolicitPurge) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitPurge) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestpurge.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitPurge) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestpurge.Merge(m, src)
}
func (m *SolicitPurge) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitPurge) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestpurge.DiscardUnknown(m)
}

var xxx_signaldetails_Requestpurge proto.InternalMessageInfo

type SolicitDetails struct {
	Edition      string `protobuf:"octets,1,opt,name=version,proto3" json:"edition,omitempty"`
	LedgerEdition uint64 `protobuf:"variableint,2,opt,name=block_version,json=blockVersion,proto3" json:"ledger_edition,omitempty"`
	Peer2peerEdition   uint64 `protobuf:"variableint,3,opt,name=p2p_version,json=p2pVersion,proto3" json:"peer2peer_edition,omitempty"`
	IfaceEdition  string `protobuf:"octets,4,opt,name=abci_version,json=abciVersion,proto3" json:"iface_edition,omitempty"`
}

func (m *SolicitDetails) Restore()         { *m = SolicitDetails{} }
func (m *SolicitDetails) Text() string { return proto.CompactTextString(m) }
func (*SolicitDetails) SchemaArtifact()    {}
func (*SolicitDetails) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{3}
}
func (m *SolicitDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestdetails.Merge(m, src)
}
func (m *SolicitDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestdetails.DiscardUnknown(m)
}

var xxx_signaldetails_Requestdetails proto.InternalMessageInfo

func (m *SolicitDetails) ObtainEdition() string {
	if m != nil {
		return m.Edition
	}
	return "REDACTED"
}

func (m *SolicitDetails) ObtainLedgerEdition() uint64 {
	if m != nil {
		return m.LedgerEdition
	}
	return 0
}

func (m *SolicitDetails) ObtainPeer2peerEdition() uint64 {
	if m != nil {
		return m.Peer2peerEdition
	}
	return 0
}

func (m *SolicitDetails) ObtainIfaceEdition() string {
	if m != nil {
		return m.IfaceEdition
	}
	return "REDACTED"
}

type SolicitInitializeSuccession struct {
	Moment            time.Time               `protobuf:"octets,1,opt,name=time,proto3,stdtime" json:"moment"`
	SuccessionUuid         string                  `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
	AgreementSettings *kinds1.AgreementSettings `protobuf:"octets,3,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_parameters,omitempty"`
	Assessors      []AssessorRevise       `protobuf:"octets,4,rep,name=validators,proto3" json:"assessors"`
	ApplicationStatusOctets   []byte                  `protobuf:"octets,5,opt,name=app_state_bytes,json=appStateBytes,proto3" json:"application_status_octets,omitempty"`
	PrimaryAltitude   int64                   `protobuf:"variableint,6,opt,name=initial_height,json=initialHeight,proto3" json:"primary_altitude,omitempty"`
}

func (m *SolicitInitializeSuccession) Restore()         { *m = SolicitInitializeSuccession{} }
func (m *SolicitInitializeSuccession) Text() string { return proto.CompactTextString(m) }
func (*SolicitInitializeSuccession) SchemaArtifact()    {}
func (*SolicitInitializeSuccession) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{4}
}
func (m *SolicitInitializeSuccession) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitInitializeSuccession) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestinitiatechain.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitInitializeSuccession) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestinitiatechain.Merge(m, src)
}
func (m *SolicitInitializeSuccession) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitInitializeSuccession) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestinitiatechain.DiscardUnknown(m)
}

var xxx_signaldetails_Requestinitiatechain proto.InternalMessageInfo

func (m *SolicitInitializeSuccession) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *SolicitInitializeSuccession) ObtainSuccessionUuid() string {
	if m != nil {
		return m.SuccessionUuid
	}
	return "REDACTED"
}

func (m *SolicitInitializeSuccession) ObtainAgreementParameters() *kinds1.AgreementSettings {
	if m != nil {
		return m.AgreementSettings
	}
	return nil
}

func (m *SolicitInitializeSuccession) ObtainAssessors() []AssessorRevise {
	if m != nil {
		return m.Assessors
	}
	return nil
}

func (m *SolicitInitializeSuccession) ObtainApplicationStatusOctets() []byte {
	if m != nil {
		return m.ApplicationStatusOctets
	}
	return nil
}

func (m *SolicitInitializeSuccession) ObtainPrimaryAltitude() int64 {
	if m != nil {
		return m.PrimaryAltitude
	}
	return 0
}

type SolicitInquire struct {
	Data   []byte `protobuf:"octets,1,opt,name=data,proto3" json:"data,omitempty"`
	Route   string `protobuf:"octets,2,opt,name=path,proto3" json:"route,omitempty"`
	Altitude int64  `protobuf:"variableint,3,opt,name=height,proto3" json:"altitude,omitempty"`
	Validate  bool   `protobuf:"variableint,4,opt,name=prove,proto3" json:"ascertain,omitempty"`
}

func (m *SolicitInquire) Restore()         { *m = SolicitInquire{} }
func (m *SolicitInquire) Text() string { return proto.CompactTextString(m) }
func (*SolicitInquire) SchemaArtifact()    {}
func (*SolicitInquire) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{5}
}
func (m *SolicitInquire) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitInquire) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestsearch.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitInquire) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestsearch.Merge(m, src)
}
func (m *SolicitInquire) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitInquire) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestsearch.DiscardUnknown(m)
}

var xxx_signaldetails_Requestsearch proto.InternalMessageInfo

func (m *SolicitInquire) ObtainData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SolicitInquire) ObtainRoute() string {
	if m != nil {
		return m.Route
	}
	return "REDACTED"
}

func (m *SolicitInquire) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SolicitInquire) ObtainAscertain() bool {
	if m != nil {
		return m.Validate
	}
	return false
}

type SolicitInspectTransfer struct {
	Tx   []byte      `protobuf:"octets,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Kind InspectTransferKind `protobuf:"variableint,2,opt,name=type,proto3,enum=tendermint.abci.CheckTxType" json:"kind,omitempty"`
}

func (m *SolicitInspectTransfer) Restore()         { *m = SolicitInspectTransfer{} }
func (m *SolicitInspectTransfer) Text() string { return proto.CompactTextString(m) }
func (*SolicitInspectTransfer) SchemaArtifact()    {}
func (*SolicitInspectTransfer) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{6}
}
func (m *SolicitInspectTransfer) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitInspectTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestinspecttrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitInspectTransfer) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestinspecttrans.Merge(m, src)
}
func (m *SolicitInspectTransfer) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitInspectTransfer) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestinspecttrans.DiscardUnknown(m)
}

var xxx_signaldetails_Requestinspecttrans proto.InternalMessageInfo

func (m *SolicitInspectTransfer) ObtainTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *SolicitInspectTransfer) ObtainKind() InspectTransferKind {
	if m != nil {
		return m.Kind
	}
	return Inspecttranskind_Fresh
}

type SolicitAppendTransfer struct {
	Tx []byte `protobuf:"octets,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *SolicitAppendTransfer) Restore()         { *m = SolicitAppendTransfer{} }
func (m *SolicitAppendTransfer) Text() string { return proto.CompactTextString(m) }
func (*SolicitAppendTransfer) SchemaArtifact()    {}
func (*SolicitAppendTransfer) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{7}
}
func (m *SolicitAppendTransfer) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitAppendTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestappendtrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitAppendTransfer) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestappendtrans.Merge(m, src)
}
func (m *SolicitAppendTransfer) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitAppendTransfer) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestappendtrans.DiscardUnknown(m)
}

var xxx_signaldetails_Requestappendtrans proto.InternalMessageInfo

func (m *SolicitAppendTransfer) ObtainTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type SolicitHarvestTrans struct {
	MaximumOctets uint64 `protobuf:"variableint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"maximum_octets,omitempty"`
	MaximumFuel   uint64 `protobuf:"variableint,2,opt,name=max_gas,json=maxGas,proto3" json:"maximum_fuel,omitempty"`
}

func (m *SolicitHarvestTrans) Restore()         { *m = SolicitHarvestTrans{} }
func (m *SolicitHarvestTrans) Text() string { return proto.CompactTextString(m) }
func (*SolicitHarvestTrans) SchemaArtifact()    {}
func (*SolicitHarvestTrans) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{8}
}
func (m *SolicitHarvestTrans) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitHarvestTrans) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestharvesttrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitHarvestTrans) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestharvesttrans.Merge(m, src)
}
func (m *SolicitHarvestTrans) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitHarvestTrans) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestharvesttrans.DiscardUnknown(m)
}

var xxx_signaldetails_Requestharvesttrans proto.InternalMessageInfo

func (m *SolicitHarvestTrans) ObtainMaximumOctets() uint64 {
	if m != nil {
		return m.MaximumOctets
	}
	return 0
}

func (m *SolicitHarvestTrans) ObtainMaximumFuel() uint64 {
	if m != nil {
		return m.MaximumFuel
	}
	return 0
}

type SolicitEndorse struct {
}

func (m *SolicitEndorse) Restore()         { *m = SolicitEndorse{} }
func (m *SolicitEndorse) Text() string { return proto.CompactTextString(m) }
func (*SolicitEndorse) SchemaArtifact()    {}
func (*SolicitEndorse) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{9}
}
func (m *SolicitEndorse) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitEndorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestendorse.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitEndorse) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestendorse.Merge(m, src)
}
func (m *SolicitEndorse) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitEndorse) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestendorse.DiscardUnknown(m)
}

var xxx_signaldetails_Requestendorse proto.InternalMessageInfo

//
type SolicitCollectionImages struct {
}

func (m *SolicitCollectionImages) Restore()         { *m = SolicitCollectionImages{} }
func (m *SolicitCollectionImages) Text() string { return proto.CompactTextString(m) }
func (*SolicitCollectionImages) SchemaArtifact()    {}
func (*SolicitCollectionImages) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{10}
}
func (m *SolicitCollectionImages) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitCollectionImages) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestcatalogimages.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitCollectionImages) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestcatalogimages.Merge(m, src)
}
func (m *SolicitCollectionImages) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitCollectionImages) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestcatalogimages.DiscardUnknown(m)
}

var xxx_signaldetails_Requestcatalogimages proto.InternalMessageInfo

//
type SolicitExtendImage struct {
	Image *Image `protobuf:"octets,1,opt,name=snapshot,proto3" json:"image,omitempty"`
	PlatformDigest  []byte    `protobuf:"octets,2,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *SolicitExtendImage) Restore()         { *m = SolicitExtendImage{} }
func (m *SolicitExtendImage) Text() string { return proto.CompactTextString(m) }
func (*SolicitExtendImage) SchemaArtifact()    {}
func (*SolicitExtendImage) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{11}
}
func (m *SolicitExtendImage) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitExtendImage) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestextendimage.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitExtendImage) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestextendimage.Merge(m, src)
}
func (m *SolicitExtendImage) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitExtendImage) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestextendimage.DiscardUnknown(m)
}

var xxx_signaldetails_Requestextendimage proto.InternalMessageInfo

func (m *SolicitExtendImage) ObtainImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *SolicitExtendImage) ObtainApplicationDigest() []byte {
	if m != nil {
		return m.PlatformDigest
	}
	return nil
}

//
type SolicitFetchImageSegment struct {
	Altitude uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Layout uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Segment  uint32 `protobuf:"variableint,3,opt,name=chunk,proto3" json:"segment,omitempty"`
}

func (m *SolicitFetchImageSegment) Restore()         { *m = SolicitFetchImageSegment{} }
func (m *SolicitFetchImageSegment) Text() string { return proto.CompactTextString(m) }
func (*SolicitFetchImageSegment) SchemaArtifact()    {}
func (*SolicitFetchImageSegment) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{12}
}
func (m *SolicitFetchImageSegment) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitFetchImageSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestloadimagefragment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitFetchImageSegment) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestloadimagefragment.Merge(m, src)
}
func (m *SolicitFetchImageSegment) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitFetchImageSegment) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestloadimagefragment.DiscardUnknown(m)
}

var xxx_signaldetails_Requestloadimagefragment proto.InternalMessageInfo

func (m *SolicitFetchImageSegment) ObtainAltitude() uint64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SolicitFetchImageSegment) ObtainLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *SolicitFetchImageSegment) ObtainSegment() uint32 {
	if m != nil {
		return m.Segment
	}
	return 0
}

//
type SolicitExecuteImageSegment struct {
	Ordinal  uint32 `protobuf:"variableint,1,opt,name=index,proto3" json:"ordinal,omitempty"`
	Segment  []byte `protobuf:"octets,2,opt,name=chunk,proto3" json:"segment,omitempty"`
	Originator string `protobuf:"octets,3,opt,name=sender,proto3" json:"originator,omitempty"`
}

func (m *SolicitExecuteImageSegment) Restore()         { *m = SolicitExecuteImageSegment{} }
func (m *SolicitExecuteImageSegment) Text() string { return proto.CompactTextString(m) }
func (*SolicitExecuteImageSegment) SchemaArtifact()    {}
func (*SolicitExecuteImageSegment) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{13}
}
func (m *SolicitExecuteImageSegment) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitExecuteImageSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestexecuteimagefragment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitExecuteImageSegment) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestexecuteimagefragment.Merge(m, src)
}
func (m *SolicitExecuteImageSegment) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitExecuteImageSegment) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestexecuteimagefragment.DiscardUnknown(m)
}

var xxx_signaldetails_Requestexecuteimagefragment proto.InternalMessageInfo

func (m *SolicitExecuteImageSegment) ObtainOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *SolicitExecuteImageSegment) ObtainSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *SolicitExecuteImageSegment) ObtainOriginator() string {
	if m != nil {
		return m.Originator
	}
	return "REDACTED"
}

type SolicitArrangeNomination struct {
	//
	MaximumTransferOctets int64 `protobuf:"variableint,1,opt,name=max_tx_bytes,json=maxTxBytes,proto3" json:"maximum_transfer_octets,omitempty"`
	//
	//
	Txs                [][]byte           `protobuf:"octets,2,rep,name=txs,proto3" json:"txs,omitempty"`
	RegionalFinalEndorse    ExpandedEndorseDetails `protobuf:"octets,3,opt,name=local_last_commit,json=localLastCommit,proto3" json:"regional_final_endorse"`
	Malpractice        []Malpractice      `protobuf:"octets,4,rep,name=misbehavior,proto3" json:"malpractice"`
	Altitude             int64              `protobuf:"variableint,5,opt,name=height,proto3" json:"altitude,omitempty"`
	Moment               time.Time          `protobuf:"octets,6,opt,name=time,proto3,stdtime" json:"moment"`
	FollowingAssessorsDigest []byte             `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_assessors_digest,omitempty"`
	//
	NominatorLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"nominator_location,omitempty"`
}

func (m *SolicitArrangeNomination) Restore()         { *m = SolicitArrangeNomination{} }
func (m *SolicitArrangeNomination) Text() string { return proto.CompactTextString(m) }
func (*SolicitArrangeNomination) SchemaArtifact()    {}
func (*SolicitArrangeNomination) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{14}
}
func (m *SolicitArrangeNomination) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitArrangeNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestprepareitem.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitArrangeNomination) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestprepareitem.Merge(m, src)
}
func (m *SolicitArrangeNomination) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitArrangeNomination) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestprepareitem.DiscardUnknown(m)
}

var xxx_signaldetails_Requestprepareitem proto.InternalMessageInfo

func (m *SolicitArrangeNomination) ObtainMaximumTransferOctets() int64 {
	if m != nil {
		return m.MaximumTransferOctets
	}
	return 0
}

func (m *SolicitArrangeNomination) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *SolicitArrangeNomination) ObtainRegionalFinalEndorse() ExpandedEndorseDetails {
	if m != nil {
		return m.RegionalFinalEndorse
	}
	return ExpandedEndorseDetails{}
}

func (m *SolicitArrangeNomination) ObtainMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *SolicitArrangeNomination) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SolicitArrangeNomination) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *SolicitArrangeNomination) ObtainFollowingAssessorsDigest() []byte {
	if m != nil {
		return m.FollowingAssessorsDigest
	}
	return nil
}

func (m *SolicitArrangeNomination) ObtainNominatorLocator() []byte {
	if m != nil {
		return m.NominatorLocation
	}
	return nil
}

type SolicitHandleNomination struct {
	Txs                [][]byte      `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
	ItemizedFinalEndorse EndorseDetails    `protobuf:"octets,2,opt,name=proposed_last_commit,json=proposedLastCommit,proto3" json:"itemized_final_endorse"`
	Malpractice        []Malpractice `protobuf:"octets,3,rep,name=misbehavior,proto3" json:"malpractice"`
	//
	Digest               []byte    `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Altitude             int64     `protobuf:"variableint,5,opt,name=height,proto3" json:"altitude,omitempty"`
	Moment               time.Time `protobuf:"octets,6,opt,name=time,proto3,stdtime" json:"moment"`
	FollowingAssessorsDigest []byte    `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_assessors_digest,omitempty"`
	//
	NominatorLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"nominator_location,omitempty"`
}

func (m *SolicitHandleNomination) Restore()         { *m = SolicitHandleNomination{} }
func (m *SolicitHandleNomination) Text() string { return proto.CompactTextString(m) }
func (*SolicitHandleNomination) SchemaArtifact()    {}
func (*SolicitHandleNomination) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{15}
}
func (m *SolicitHandleNomination) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitHandleNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestexecuteitem.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitHandleNomination) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestexecuteitem.Merge(m, src)
}
func (m *SolicitHandleNomination) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitHandleNomination) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestexecuteitem.DiscardUnknown(m)
}

var xxx_signaldetails_Requestexecuteitem proto.InternalMessageInfo

func (m *SolicitHandleNomination) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *SolicitHandleNomination) ObtainItemizedFinalEndorse() EndorseDetails {
	if m != nil {
		return m.ItemizedFinalEndorse
	}
	return EndorseDetails{}
}

func (m *SolicitHandleNomination) ObtainMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *SolicitHandleNomination) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *SolicitHandleNomination) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SolicitHandleNomination) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *SolicitHandleNomination) ObtainFollowingAssessorsDigest() []byte {
	if m != nil {
		return m.FollowingAssessorsDigest
	}
	return nil
}

func (m *SolicitHandleNomination) ObtainNominatorLocator() []byte {
	if m != nil {
		return m.NominatorLocation
	}
	return nil
}

//
type SolicitBroadenBallot struct {
	//
	Digest []byte `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	//
	Altitude int64 `protobuf:"variableint,2,opt,name=height,proto3" json:"altitude,omitempty"`
	//
	Moment               time.Time     `protobuf:"octets,3,opt,name=time,proto3,stdtime" json:"moment"`
	Txs                [][]byte      `protobuf:"octets,4,rep,name=txs,proto3" json:"txs,omitempty"`
	ItemizedFinalEndorse EndorseDetails    `protobuf:"octets,5,opt,name=proposed_last_commit,json=proposedLastCommit,proto3" json:"itemized_final_endorse"`
	Malpractice        []Malpractice `protobuf:"octets,6,rep,name=misbehavior,proto3" json:"malpractice"`
	FollowingAssessorsDigest []byte        `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_assessors_digest,omitempty"`
	//
	NominatorLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"nominator_location,omitempty"`
}

func (m *SolicitBroadenBallot) Restore()         { *m = SolicitBroadenBallot{} }
func (m *SolicitBroadenBallot) Text() string { return proto.CompactTextString(m) }
func (*SolicitBroadenBallot) SchemaArtifact()    {}
func (*SolicitBroadenBallot) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{16}
}
func (m *SolicitBroadenBallot) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitBroadenBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestextendballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitBroadenBallot) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestextendballot.Merge(m, src)
}
func (m *SolicitBroadenBallot) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitBroadenBallot) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestextendballot.DiscardUnknown(m)
}

var xxx_signaldetails_Requestextendballot proto.InternalMessageInfo

func (m *SolicitBroadenBallot) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *SolicitBroadenBallot) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SolicitBroadenBallot) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *SolicitBroadenBallot) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *SolicitBroadenBallot) ObtainItemizedFinalEndorse() EndorseDetails {
	if m != nil {
		return m.ItemizedFinalEndorse
	}
	return EndorseDetails{}
}

func (m *SolicitBroadenBallot) ObtainMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *SolicitBroadenBallot) ObtainFollowingAssessorsDigest() []byte {
	if m != nil {
		return m.FollowingAssessorsDigest
	}
	return nil
}

func (m *SolicitBroadenBallot) ObtainNominatorLocator() []byte {
	if m != nil {
		return m.NominatorLocation
	}
	return nil
}

//
type SolicitValidateBallotAddition struct {
	//
	Digest []byte `protobuf:"octets,1,opt,name=hash,proto3" json:"digest,omitempty"`
	//
	AssessorLocation []byte `protobuf:"octets,2,opt,name=validator_address,json=validatorAddress,proto3" json:"assessor_location,omitempty"`
	Altitude           int64  `protobuf:"variableint,3,opt,name=height,proto3" json:"altitude,omitempty"`
	BallotAddition    []byte `protobuf:"octets,4,opt,name=vote_extension,json=voteExtension,proto3" json:"ballot_addition,omitempty"`
}

func (m *SolicitValidateBallotAddition) Restore()         { *m = SolicitValidateBallotAddition{} }
func (m *SolicitValidateBallotAddition) Text() string { return proto.CompactTextString(m) }
func (*SolicitValidateBallotAddition) SchemaArtifact()    {}
func (*SolicitValidateBallotAddition) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{17}
}
func (m *SolicitValidateBallotAddition) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitValidateBallotAddition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestverifyballotaddition.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitValidateBallotAddition) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestverifyballotaddition.Merge(m, src)
}
func (m *SolicitValidateBallotAddition) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitValidateBallotAddition) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestverifyballotaddition.DiscardUnknown(m)
}

var xxx_signaldetails_Requestverifyballotaddition proto.InternalMessageInfo

func (m *SolicitValidateBallotAddition) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *SolicitValidateBallotAddition) ObtainAssessorLocation() []byte {
	if m != nil {
		return m.AssessorLocation
	}
	return nil
}

func (m *SolicitValidateBallotAddition) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SolicitValidateBallotAddition) ObtainBallotAddition() []byte {
	if m != nil {
		return m.BallotAddition
	}
	return nil
}

type SolicitCulminateLedger struct {
	Txs               [][]byte      `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
	ResolvedFinalEndorse EndorseDetails    `protobuf:"octets,2,opt,name=decided_last_commit,json=decidedLastCommit,proto3" json:"resolved_final_endorse"`
	Malpractice       []Malpractice `protobuf:"octets,3,rep,name=misbehavior,proto3" json:"malpractice"`
	//
	Digest               []byte    `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Altitude             int64     `protobuf:"variableint,5,opt,name=height,proto3" json:"altitude,omitempty"`
	Moment               time.Time `protobuf:"octets,6,opt,name=time,proto3,stdtime" json:"moment"`
	FollowingAssessorsDigest []byte    `protobuf:"octets,7,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"following_assessors_digest,omitempty"`
	//
	NominatorLocation []byte `protobuf:"octets,8,opt,name=proposer_address,json=proposerAddress,proto3" json:"nominator_location,omitempty"`
}

func (m *SolicitCulminateLedger) Restore()         { *m = SolicitCulminateLedger{} }
func (m *SolicitCulminateLedger) Text() string { return proto.CompactTextString(m) }
func (*SolicitCulminateLedger) SchemaArtifact()    {}
func (*SolicitCulminateLedger) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{18}
}
func (m *SolicitCulminateLedger) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitCulminateLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Requestfinalizeledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitCulminateLedger) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Requestfinalizeledger.Merge(m, src)
}
func (m *SolicitCulminateLedger) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitCulminateLedger) XXX_Dropunfamiliar() {
	xxx_signaldetails_Requestfinalizeledger.DiscardUnknown(m)
}

var xxx_signaldetails_Requestfinalizeledger proto.InternalMessageInfo

func (m *SolicitCulminateLedger) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *SolicitCulminateLedger) ObtainResolvedFinalEndorse() EndorseDetails {
	if m != nil {
		return m.ResolvedFinalEndorse
	}
	return EndorseDetails{}
}

func (m *SolicitCulminateLedger) ObtainMalpractice() []Malpractice {
	if m != nil {
		return m.Malpractice
	}
	return nil
}

func (m *SolicitCulminateLedger) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *SolicitCulminateLedger) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SolicitCulminateLedger) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *SolicitCulminateLedger) ObtainFollowingAssessorsDigest() []byte {
	if m != nil {
		return m.FollowingAssessorsDigest
	}
	return nil
}

func (m *SolicitCulminateLedger) ObtainNominatorLocator() []byte {
	if m != nil {
		return m.NominatorLocation
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
	Datum isreplying_Datum `protobuf_oneof:"datum"`
}

func (m *Reply) Restore()         { *m = Reply{} }
func (m *Reply) Text() string { return proto.CompactTextString(m) }
func (*Reply) SchemaArtifact()    {}
func (*Reply) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{19}
}
func (m *Reply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Reply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Reply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Reply.Merge(m, src)
}
func (m *Reply) XXX_Extent() int {
	return m.Extent()
}
func (m *Reply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Reply.DiscardUnknown(m)
}

var xxx_signaldetails_Reply proto.InternalMessageInfo

type isreplying_Datum interface {
	isreplying_Datum()
	SerializeToward([]byte) (int, error)
	Extent() int
}

type Reply_Exemption struct {
	Exemption *ReplyExemption `protobuf:"octets,1,opt,name=exception,proto3,oneof" json:"exemption,omitempty"`
}
type Reply_Reverberate struct {
	Reverberate *ReplyReverberate `protobuf:"octets,2,opt,name=echo,proto3,oneof" json:"reverberate,omitempty"`
}
type Reply_Purge struct {
	Purge *ReplyPurge `protobuf:"octets,3,opt,name=flush,proto3,oneof" json:"purge,omitempty"`
}
type Reply_Details struct {
	Details *ReplyDetails `protobuf:"octets,4,opt,name=info,proto3,oneof" json:"details,omitempty"`
}
type Reply_Initiatechain struct {
	InitializeSuccession *ReplyInitializeSuccession `protobuf:"octets,6,opt,name=init_chain,json=initChain,proto3,oneof" json:"initialize_succession,omitempty"`
}
type Reply_Inquire struct {
	Inquire *ReplyInquire `protobuf:"octets,7,opt,name=query,proto3,oneof" json:"inquire,omitempty"`
}
type Reply_Inspecttrans struct {
	InspectTransfer *ReplyInspectTransfer `protobuf:"octets,9,opt,name=check_tx,json=checkTx,proto3,oneof" json:"inspect_transfer,omitempty"`
}
type Reply_Endorse struct {
	Endorse *ReplyEndorse `protobuf:"octets,12,opt,name=commit,proto3,oneof" json:"endorse,omitempty"`
}
type Reply_Catalogimages struct {
	CollectionImages *ReplyCatalogImages `protobuf:"octets,13,opt,name=list_snapshots,json=listSnapshots,proto3,oneof" json:"catalog_images,omitempty"`
}
type Reply_Extendimage struct {
	ExtendImage *ReplyExtendImage `protobuf:"octets,14,opt,name=offer_snapshot,json=offerSnapshot,proto3,oneof" json:"extend_image,omitempty"`
}
type Reply_Loadimagefragment struct {
	FetchImageSegment *ReplyFetchImageSegment `protobuf:"octets,15,opt,name=load_snapshot_chunk,json=loadSnapshotChunk,proto3,oneof" json:"fetch_image_segment,omitempty"`
}
type Reply_Executeimagefragment struct {
	ExecuteImageSegment *ReplyExecuteImageSegment `protobuf:"octets,16,opt,name=apply_snapshot_chunk,json=applySnapshotChunk,proto3,oneof" json:"execute_image_segment,omitempty"`
}
type Reply_Prepareitem struct {
	ArrangeNomination *ReplyArrangeNomination `protobuf:"octets,17,opt,name=prepare_proposal,json=prepareProposal,proto3,oneof" json:"arrange_nomination,omitempty"`
}
type Reply_Executeitem struct {
	HandleNomination *ReplyHandleNomination `protobuf:"octets,18,opt,name=process_proposal,json=processProposal,proto3,oneof" json:"handle_nomination,omitempty"`
}
type Reply_Extendballot struct {
	BroadenBallot *ReplyBroadenBallot `protobuf:"octets,19,opt,name=extend_vote,json=extendVote,proto3,oneof" json:"broaden_ballot,omitempty"`
}
type Reply_Verifyballotaddition struct {
	ValidateBallotAddition *ReplyValidateBallotAddition `protobuf:"octets,20,opt,name=verify_vote_extension,json=verifyVoteExtension,proto3,oneof" json:"validate_ballot_addition,omitempty"`
}
type Reply_Finalizeledger struct {
	CulminateLedger *ReplyCulminateLedger `protobuf:"octets,21,opt,name=finalize_block,json=finalizeBlock,proto3,oneof" json:"culminate_ledger,omitempty"`
}
type Reply_Appendtrans struct {
	AppendTransfer *ReplyAppendTransfer `protobuf:"octets,22,opt,name=insert_tx,json=insertTx,proto3,oneof" json:"append_transfer,omitempty"`
}
type Reply_Harvesttrans struct {
	HarvestTrans *ReplyHarvestTrans `protobuf:"octets,23,opt,name=reap_txs,json=reapTxs,proto3,oneof" json:"harvest_trans,omitempty"`
}

func (*Reply_Exemption) isreplying_Datum()           {}
func (*Reply_Reverberate) isreplying_Datum()                {}
func (*Reply_Purge) isreplying_Datum()               {}
func (*Reply_Details) isreplying_Datum()                {}
func (*Reply_Initiatechain) isreplying_Datum()           {}
func (*Reply_Inquire) isreplying_Datum()               {}
func (*Reply_Inspecttrans) isreplying_Datum()             {}
func (*Reply_Endorse) isreplying_Datum()              {}
func (*Reply_Catalogimages) isreplying_Datum()       {}
func (*Reply_Extendimage) isreplying_Datum()       {}
func (*Reply_Loadimagefragment) isreplying_Datum()   {}
func (*Reply_Executeimagefragment) isreplying_Datum()  {}
func (*Reply_Prepareitem) isreplying_Datum()     {}
func (*Reply_Executeitem) isreplying_Datum()     {}
func (*Reply_Extendballot) isreplying_Datum()          {}
func (*Reply_Verifyballotaddition) isreplying_Datum() {}
func (*Reply_Finalizeledger) isreplying_Datum()       {}
func (*Reply_Appendtrans) isreplying_Datum()            {}
func (*Reply_Harvesttrans) isreplying_Datum()             {}

func (m *Reply) ObtainDatum() isreplying_Datum {
	if m != nil {
		return m.Datum
	}
	return nil
}

func (m *Reply) ObtainExemption() *ReplyExemption {
	if x, ok := m.ObtainDatum().(*Reply_Exemption); ok {
		return x.Exemption
	}
	return nil
}

func (m *Reply) ObtainReverberate() *ReplyReverberate {
	if x, ok := m.ObtainDatum().(*Reply_Reverberate); ok {
		return x.Reverberate
	}
	return nil
}

func (m *Reply) ObtainPurge() *ReplyPurge {
	if x, ok := m.ObtainDatum().(*Reply_Purge); ok {
		return x.Purge
	}
	return nil
}

func (m *Reply) ObtainDetails() *ReplyDetails {
	if x, ok := m.ObtainDatum().(*Reply_Details); ok {
		return x.Details
	}
	return nil
}

func (m *Reply) ObtainInitializeSuccession() *ReplyInitializeSuccession {
	if x, ok := m.ObtainDatum().(*Reply_Initiatechain); ok {
		return x.InitializeSuccession
	}
	return nil
}

func (m *Reply) ObtainInquire() *ReplyInquire {
	if x, ok := m.ObtainDatum().(*Reply_Inquire); ok {
		return x.Inquire
	}
	return nil
}

func (m *Reply) ObtainInspectTransfer() *ReplyInspectTransfer {
	if x, ok := m.ObtainDatum().(*Reply_Inspecttrans); ok {
		return x.InspectTransfer
	}
	return nil
}

func (m *Reply) ObtainEndorse() *ReplyEndorse {
	if x, ok := m.ObtainDatum().(*Reply_Endorse); ok {
		return x.Endorse
	}
	return nil
}

func (m *Reply) ObtainCatalogImages() *ReplyCatalogImages {
	if x, ok := m.ObtainDatum().(*Reply_Catalogimages); ok {
		return x.CollectionImages
	}
	return nil
}

func (m *Reply) ObtainExtendImage() *ReplyExtendImage {
	if x, ok := m.ObtainDatum().(*Reply_Extendimage); ok {
		return x.ExtendImage
	}
	return nil
}

func (m *Reply) ObtainFetchImageSegment() *ReplyFetchImageSegment {
	if x, ok := m.ObtainDatum().(*Reply_Loadimagefragment); ok {
		return x.FetchImageSegment
	}
	return nil
}

func (m *Reply) ObtainExecuteImageSegment() *ReplyExecuteImageSegment {
	if x, ok := m.ObtainDatum().(*Reply_Executeimagefragment); ok {
		return x.ExecuteImageSegment
	}
	return nil
}

func (m *Reply) ObtainArrangeNomination() *ReplyArrangeNomination {
	if x, ok := m.ObtainDatum().(*Reply_Prepareitem); ok {
		return x.ArrangeNomination
	}
	return nil
}

func (m *Reply) ObtainHandleNomination() *ReplyHandleNomination {
	if x, ok := m.ObtainDatum().(*Reply_Executeitem); ok {
		return x.HandleNomination
	}
	return nil
}

func (m *Reply) ObtainBroadenBallot() *ReplyBroadenBallot {
	if x, ok := m.ObtainDatum().(*Reply_Extendballot); ok {
		return x.BroadenBallot
	}
	return nil
}

func (m *Reply) ObtainValidateBallotAddition() *ReplyValidateBallotAddition {
	if x, ok := m.ObtainDatum().(*Reply_Verifyballotaddition); ok {
		return x.ValidateBallotAddition
	}
	return nil
}

func (m *Reply) ObtainCulminateLedger() *ReplyCulminateLedger {
	if x, ok := m.ObtainDatum().(*Reply_Finalizeledger); ok {
		return x.CulminateLedger
	}
	return nil
}

func (m *Reply) ObtainAppendTransfer() *ReplyAppendTransfer {
	if x, ok := m.ObtainDatum().(*Reply_Appendtrans); ok {
		return x.AppendTransfer
	}
	return nil
}

func (m *Reply) ObtainHarvestTrans() *ReplyHarvestTrans {
	if x, ok := m.ObtainDatum().(*Reply_Harvesttrans); ok {
		return x.HarvestTrans
	}
	return nil
}

//
func (*Reply) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Reply_Exemption)(nil),
		(*Reply_Reverberate)(nil),
		(*Reply_Purge)(nil),
		(*Reply_Details)(nil),
		(*Reply_Initiatechain)(nil),
		(*Reply_Inquire)(nil),
		(*Reply_Inspecttrans)(nil),
		(*Reply_Endorse)(nil),
		(*Reply_Catalogimages)(nil),
		(*Reply_Extendimage)(nil),
		(*Reply_Loadimagefragment)(nil),
		(*Reply_Executeimagefragment)(nil),
		(*Reply_Prepareitem)(nil),
		(*Reply_Executeitem)(nil),
		(*Reply_Extendballot)(nil),
		(*Reply_Verifyballotaddition)(nil),
		(*Reply_Finalizeledger)(nil),
		(*Reply_Appendtrans)(nil),
		(*Reply_Harvesttrans)(nil),
	}
}

//
type ReplyExemption struct {
	Failure string `protobuf:"octets,1,opt,name=error,proto3" json:"failure,omitempty"`
}

func (m *ReplyExemption) Restore()         { *m = ReplyExemption{} }
func (m *ReplyExemption) Text() string { return proto.CompactTextString(m) }
func (*ReplyExemption) SchemaArtifact()    {}
func (*ReplyExemption) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{20}
}
func (m *ReplyExemption) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyExemption) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyexemption.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyExemption) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyexemption.Merge(m, src)
}
func (m *ReplyExemption) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyExemption) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyexemption.DiscardUnknown(m)
}

var xxx_signaldetails_Replyexemption proto.InternalMessageInfo

func (m *ReplyExemption) ObtainFailure() string {
	if m != nil {
		return m.Failure
	}
	return "REDACTED"
}

type ReplyReverberate struct {
	Signal string `protobuf:"octets,1,opt,name=message,proto3" json:"signal,omitempty"`
}

func (m *ReplyReverberate) Restore()         { *m = ReplyReverberate{} }
func (m *ReplyReverberate) Text() string { return proto.CompactTextString(m) }
func (*ReplyReverberate) SchemaArtifact()    {}
func (*ReplyReverberate) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{21}
}
func (m *ReplyReverberate) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyReverberate) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyecho.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyReverberate) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyecho.Merge(m, src)
}
func (m *ReplyReverberate) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyReverberate) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyecho.DiscardUnknown(m)
}

var xxx_signaldetails_Replyecho proto.InternalMessageInfo

func (m *ReplyReverberate) ObtainArtifact() string {
	if m != nil {
		return m.Signal
	}
	return "REDACTED"
}

type ReplyPurge struct {
}

func (m *ReplyPurge) Restore()         { *m = ReplyPurge{} }
func (m *ReplyPurge) Text() string { return proto.CompactTextString(m) }
func (*ReplyPurge) SchemaArtifact()    {}
func (*ReplyPurge) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{22}
}
func (m *ReplyPurge) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyPurge) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replypurge.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyPurge) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replypurge.Merge(m, src)
}
func (m *ReplyPurge) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyPurge) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replypurge.DiscardUnknown(m)
}

var xxx_signaldetails_Replypurge proto.InternalMessageInfo

type ReplyDetails struct {
	Data             string `protobuf:"octets,1,opt,name=data,proto3" json:"data,omitempty"`
	Edition          string `protobuf:"octets,2,opt,name=version,proto3" json:"edition,omitempty"`
	PlatformEdition       uint64 `protobuf:"variableint,3,opt,name=app_version,json=appVersion,proto3" json:"application_edition,omitempty"`
	FinalLedgerAltitude  int64  `protobuf:"variableint,4,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"final_ledger_altitude,omitempty"`
	FinalLedgerPlatformDigest []byte `protobuf:"octets,5,opt,name=last_block_app_hash,json=lastBlockAppHash,proto3" json:"final_ledger_application_digest,omitempty"`
}

func (m *ReplyDetails) Restore()         { *m = ReplyDetails{} }
func (m *ReplyDetails) Text() string { return proto.CompactTextString(m) }
func (*ReplyDetails) SchemaArtifact()    {}
func (*ReplyDetails) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{23}
}
func (m *ReplyDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replydetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replydetails.Merge(m, src)
}
func (m *ReplyDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replydetails.DiscardUnknown(m)
}

var xxx_signaldetails_Replydetails proto.InternalMessageInfo

func (m *ReplyDetails) ObtainData() string {
	if m != nil {
		return m.Data
	}
	return "REDACTED"
}

func (m *ReplyDetails) ObtainEdition() string {
	if m != nil {
		return m.Edition
	}
	return "REDACTED"
}

func (m *ReplyDetails) ObtainApplicationEdition() uint64 {
	if m != nil {
		return m.PlatformEdition
	}
	return 0
}

func (m *ReplyDetails) ObtainFinalLedgerAltitude() int64 {
	if m != nil {
		return m.FinalLedgerAltitude
	}
	return 0
}

func (m *ReplyDetails) ObtainFinalLedgerApplicationDigest() []byte {
	if m != nil {
		return m.FinalLedgerPlatformDigest
	}
	return nil
}

type ReplyInitializeSuccession struct {
	AgreementSettings *kinds1.AgreementSettings `protobuf:"octets,1,opt,name=consensus_params,json=consensusParams,proto3" json:"agreement_parameters,omitempty"`
	Assessors      []AssessorRevise       `protobuf:"octets,2,rep,name=validators,proto3" json:"assessors"`
	PlatformDigest         []byte                  `protobuf:"octets,3,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *ReplyInitializeSuccession) Restore()         { *m = ReplyInitializeSuccession{} }
func (m *ReplyInitializeSuccession) Text() string { return proto.CompactTextString(m) }
func (*ReplyInitializeSuccession) SchemaArtifact()    {}
func (*ReplyInitializeSuccession) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{24}
}
func (m *ReplyInitializeSuccession) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyInitializeSuccession) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyinitiatechain.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyInitializeSuccession) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyinitiatechain.Merge(m, src)
}
func (m *ReplyInitializeSuccession) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyInitializeSuccession) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyinitiatechain.DiscardUnknown(m)
}

var xxx_signaldetails_Replyinitiatechain proto.InternalMessageInfo

func (m *ReplyInitializeSuccession) ObtainAgreementParameters() *kinds1.AgreementSettings {
	if m != nil {
		return m.AgreementSettings
	}
	return nil
}

func (m *ReplyInitializeSuccession) ObtainAssessors() []AssessorRevise {
	if m != nil {
		return m.Assessors
	}
	return nil
}

func (m *ReplyInitializeSuccession) ObtainApplicationDigest() []byte {
	if m != nil {
		return m.PlatformDigest
	}
	return nil
}

type ReplyInquire struct {
	Cipher uint32 `protobuf:"variableint,1,opt,name=code,proto3" json:"cipher,omitempty"`
	//
	Log       string           `protobuf:"octets,3,opt,name=log,proto3" json:"log,omitempty"`
	Details      string           `protobuf:"octets,4,opt,name=info,proto3" json:"details,omitempty"`
	Ordinal     int64            `protobuf:"variableint,5,opt,name=index,proto3" json:"ordinal,omitempty"`
	Key       []byte           `protobuf:"octets,6,opt,name=key,proto3" json:"key,omitempty"`
	Datum     []byte           `protobuf:"octets,7,opt,name=value,proto3" json:"datum,omitempty"`
	AttestationActions  *security.AttestationActions `protobuf:"octets,8,opt,name=proof_ops,json=proofOps,proto3" json:"attestation_actions,omitempty"`
	Altitude    int64            `protobuf:"variableint,9,opt,name=height,proto3" json:"altitude,omitempty"`
	Codeset string           `protobuf:"octets,10,opt,name=codespace,proto3" json:"codeset,omitempty"`
}

func (m *ReplyInquire) Restore()         { *m = ReplyInquire{} }
func (m *ReplyInquire) Text() string { return proto.CompactTextString(m) }
func (*ReplyInquire) SchemaArtifact()    {}
func (*ReplyInquire) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{25}
}
func (m *ReplyInquire) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyInquire) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replysearch.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyInquire) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replysearch.Merge(m, src)
}
func (m *ReplyInquire) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyInquire) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replysearch.DiscardUnknown(m)
}

var xxx_signaldetails_Replysearch proto.InternalMessageInfo

func (m *ReplyInquire) ObtainCipher() uint32 {
	if m != nil {
		return m.Cipher
	}
	return 0
}

func (m *ReplyInquire) ObtainRecord() string {
	if m != nil {
		return m.Log
	}
	return "REDACTED"
}

func (m *ReplyInquire) ObtainDetails() string {
	if m != nil {
		return m.Details
	}
	return "REDACTED"
}

func (m *ReplyInquire) ObtainOrdinal() int64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *ReplyInquire) ObtainToken() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ReplyInquire) ObtainDatum() []byte {
	if m != nil {
		return m.Datum
	}
	return nil
}

func (m *ReplyInquire) ObtainAttestationActions() *security.AttestationActions {
	if m != nil {
		return m.AttestationActions
	}
	return nil
}

func (m *ReplyInquire) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *ReplyInquire) ObtainCodeset() string {
	if m != nil {
		return m.Codeset
	}
	return "REDACTED"
}

type ReplyInspectTransfer struct {
	Cipher      uint32  `protobuf:"variableint,1,opt,name=code,proto3" json:"cipher,omitempty"`
	Data      []byte  `protobuf:"octets,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string  `protobuf:"octets,3,opt,name=log,proto3" json:"log,omitempty"`
	Details      string  `protobuf:"octets,4,opt,name=info,proto3" json:"details,omitempty"`
	FuelDesired int64   `protobuf:"variableint,5,opt,name=gas_wanted,proto3" json:"fuel_desired,omitempty"`
	FuelUtilized   int64   `protobuf:"variableint,6,opt,name=gas_used,proto3" json:"fuel_utilized,omitempty"`
	Incidents    []Incident `protobuf:"octets,7,rep,name=events,proto3" json:"incidents,omitempty"`
	Codeset string  `protobuf:"octets,8,opt,name=codespace,proto3" json:"codeset,omitempty"`
}

func (m *ReplyInspectTransfer) Restore()         { *m = ReplyInspectTransfer{} }
func (m *ReplyInspectTransfer) Text() string { return proto.CompactTextString(m) }
func (*ReplyInspectTransfer) SchemaArtifact()    {}
func (*ReplyInspectTransfer) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{26}
}
func (m *ReplyInspectTransfer) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyInspectTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyinspecttrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyInspectTransfer) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyinspecttrans.Merge(m, src)
}
func (m *ReplyInspectTransfer) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyInspectTransfer) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyinspecttrans.DiscardUnknown(m)
}

var xxx_signaldetails_Replyinspecttrans proto.InternalMessageInfo

func (m *ReplyInspectTransfer) ObtainCipher() uint32 {
	if m != nil {
		return m.Cipher
	}
	return 0
}

func (m *ReplyInspectTransfer) ObtainData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReplyInspectTransfer) ObtainRecord() string {
	if m != nil {
		return m.Log
	}
	return "REDACTED"
}

func (m *ReplyInspectTransfer) ObtainDetails() string {
	if m != nil {
		return m.Details
	}
	return "REDACTED"
}

func (m *ReplyInspectTransfer) ObtainFuelDesired() int64 {
	if m != nil {
		return m.FuelDesired
	}
	return 0
}

func (m *ReplyInspectTransfer) ObtainFuelUtilized() int64 {
	if m != nil {
		return m.FuelUtilized
	}
	return 0
}

func (m *ReplyInspectTransfer) ObtainIncidents() []Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

func (m *ReplyInspectTransfer) ObtainCodeset() string {
	if m != nil {
		return m.Codeset
	}
	return "REDACTED"
}

type ReplyAppendTransfer struct {
	Cipher uint32 `protobuf:"variableint,1,opt,name=code,proto3" json:"cipher,omitempty"`
}

func (m *ReplyAppendTransfer) Restore()         { *m = ReplyAppendTransfer{} }
func (m *ReplyAppendTransfer) Text() string { return proto.CompactTextString(m) }
func (*ReplyAppendTransfer) SchemaArtifact()    {}
func (*ReplyAppendTransfer) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{27}
}
func (m *ReplyAppendTransfer) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyAppendTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyappendtrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyAppendTransfer) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyappendtrans.Merge(m, src)
}
func (m *ReplyAppendTransfer) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyAppendTransfer) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyappendtrans.DiscardUnknown(m)
}

var xxx_signaldetails_Replyappendtrans proto.InternalMessageInfo

func (m *ReplyAppendTransfer) ObtainCipher() uint32 {
	if m != nil {
		return m.Cipher
	}
	return 0
}

type ReplyHarvestTrans struct {
	Txs [][]byte `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *ReplyHarvestTrans) Restore()         { *m = ReplyHarvestTrans{} }
func (m *ReplyHarvestTrans) Text() string { return proto.CompactTextString(m) }
func (*ReplyHarvestTrans) SchemaArtifact()    {}
func (*ReplyHarvestTrans) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{28}
}
func (m *ReplyHarvestTrans) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyHarvestTrans) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyharvesttrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyHarvestTrans) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyharvesttrans.Merge(m, src)
}
func (m *ReplyHarvestTrans) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyHarvestTrans) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyharvesttrans.DiscardUnknown(m)
}

var xxx_signaldetails_Replyharvesttrans proto.InternalMessageInfo

func (m *ReplyHarvestTrans) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ReplyEndorse struct {
	PreserveAltitude int64 `protobuf:"variableint,3,opt,name=retain_height,json=retainHeight,proto3" json:"preserve_altitude,omitempty"`
}

func (m *ReplyEndorse) Restore()         { *m = ReplyEndorse{} }
func (m *ReplyEndorse) Text() string { return proto.CompactTextString(m) }
func (*ReplyEndorse) SchemaArtifact()    {}
func (*ReplyEndorse) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{29}
}
func (m *ReplyEndorse) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyEndorse) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyendorse.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyEndorse) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyendorse.Merge(m, src)
}
func (m *ReplyEndorse) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyEndorse) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyendorse.DiscardUnknown(m)
}

var xxx_signaldetails_Replyendorse proto.InternalMessageInfo

func (m *ReplyEndorse) ObtainPreserveAltitude() int64 {
	if m != nil {
		return m.PreserveAltitude
	}
	return 0
}

type ReplyCatalogImages struct {
	Images []*Image `protobuf:"octets,1,rep,name=snapshots,proto3" json:"images,omitempty"`
}

func (m *ReplyCatalogImages) Restore()         { *m = ReplyCatalogImages{} }
func (m *ReplyCatalogImages) Text() string { return proto.CompactTextString(m) }
func (*ReplyCatalogImages) SchemaArtifact()    {}
func (*ReplyCatalogImages) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{30}
}
func (m *ReplyCatalogImages) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyCatalogImages) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replycatalogimages.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyCatalogImages) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replycatalogimages.Merge(m, src)
}
func (m *ReplyCatalogImages) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyCatalogImages) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replycatalogimages.DiscardUnknown(m)
}

var xxx_signaldetails_Replycatalogimages proto.InternalMessageInfo

func (m *ReplyCatalogImages) ObtainImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type ReplyExtendImage struct {
	Outcome Replyextendimage_Outcome `protobuf:"variableint,1,opt,name=result,proto3,enum=tendermint.abci.ResponseOfferSnapshot_Result" json:"outcome,omitempty"`
}

func (m *ReplyExtendImage) Restore()         { *m = ReplyExtendImage{} }
func (m *ReplyExtendImage) Text() string { return proto.CompactTextString(m) }
func (*ReplyExtendImage) SchemaArtifact()    {}
func (*ReplyExtendImage) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{31}
}
func (m *ReplyExtendImage) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyExtendImage) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyextendimage.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyExtendImage) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyextendimage.Merge(m, src)
}
func (m *ReplyExtendImage) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyExtendImage) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyextendimage.DiscardUnknown(m)
}

var xxx_signaldetails_Replyextendimage proto.InternalMessageInfo

func (m *ReplyExtendImage) ObtainOutcome() Replyextendimage_Outcome {
	if m != nil {
		return m.Outcome
	}
	return Replyextendimage_UNFAMILIAR
}

type ReplyFetchImageSegment struct {
	Segment []byte `protobuf:"octets,1,opt,name=chunk,proto3" json:"segment,omitempty"`
}

func (m *ReplyFetchImageSegment) Restore()         { *m = ReplyFetchImageSegment{} }
func (m *ReplyFetchImageSegment) Text() string { return proto.CompactTextString(m) }
func (*ReplyFetchImageSegment) SchemaArtifact()    {}
func (*ReplyFetchImageSegment) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{32}
}
func (m *ReplyFetchImageSegment) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyFetchImageSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyloadimagefragment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyFetchImageSegment) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyloadimagefragment.Merge(m, src)
}
func (m *ReplyFetchImageSegment) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyFetchImageSegment) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyloadimagefragment.DiscardUnknown(m)
}

var xxx_signaldetails_Replyloadimagefragment proto.InternalMessageInfo

func (m *ReplyFetchImageSegment) ObtainSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

type ReplyExecuteImageSegment struct {
	Outcome        Replyapplyimagefragment_Outcome `protobuf:"variableint,1,opt,name=result,proto3,enum=tendermint.abci.ResponseApplySnapshotChunk_Result" json:"outcome,omitempty"`
	RetrieveSegments []uint32                          `protobuf:"variableint,2,rep,packed,name=refetch_chunks,json=refetchChunks,proto3" json:"retrieve_segments,omitempty"`
	DeclineOriginators []string                          `protobuf:"octets,3,rep,name=reject_senders,json=rejectSenders,proto3" json:"decline_originators,omitempty"`
}

func (m *ReplyExecuteImageSegment) Restore()         { *m = ReplyExecuteImageSegment{} }
func (m *ReplyExecuteImageSegment) Text() string { return proto.CompactTextString(m) }
func (*ReplyExecuteImageSegment) SchemaArtifact()    {}
func (*ReplyExecuteImageSegment) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{33}
}
func (m *ReplyExecuteImageSegment) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyExecuteImageSegment) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyapplyimagefragment.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyExecuteImageSegment) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyapplyimagefragment.Merge(m, src)
}
func (m *ReplyExecuteImageSegment) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyExecuteImageSegment) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyapplyimagefragment.DiscardUnknown(m)
}

var xxx_signaldetails_Replyapplyimagefragment proto.InternalMessageInfo

func (m *ReplyExecuteImageSegment) ObtainOutcome() Replyapplyimagefragment_Outcome {
	if m != nil {
		return m.Outcome
	}
	return Replyapplyimagefragment_UNFAMILIAR
}

func (m *ReplyExecuteImageSegment) ObtainRetrieveSegments() []uint32 {
	if m != nil {
		return m.RetrieveSegments
	}
	return nil
}

func (m *ReplyExecuteImageSegment) ObtainDeclineOriginators() []string {
	if m != nil {
		return m.DeclineOriginators
	}
	return nil
}

type ReplyArrangeNomination struct {
	Txs [][]byte `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *ReplyArrangeNomination) Restore()         { *m = ReplyArrangeNomination{} }
func (m *ReplyArrangeNomination) Text() string { return proto.CompactTextString(m) }
func (*ReplyArrangeNomination) SchemaArtifact()    {}
func (*ReplyArrangeNomination) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{34}
}
func (m *ReplyArrangeNomination) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyArrangeNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyprepareitem.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyArrangeNomination) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyprepareitem.Merge(m, src)
}
func (m *ReplyArrangeNomination) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyArrangeNomination) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyprepareitem.DiscardUnknown(m)
}

var xxx_signaldetails_Replyprepareitem proto.InternalMessageInfo

func (m *ReplyArrangeNomination) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ReplyHandleNomination struct {
	Condition Responseexecuteitem_Itemstatus `protobuf:"variableint,1,opt,name=status,proto3,enum=tendermint.abci.ResponseProcessProposal_ProposalStatus" json:"condition,omitempty"`
}

func (m *ReplyHandleNomination) Restore()         { *m = ReplyHandleNomination{} }
func (m *ReplyHandleNomination) Text() string { return proto.CompactTextString(m) }
func (*ReplyHandleNomination) SchemaArtifact()    {}
func (*ReplyHandleNomination) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{35}
}
func (m *ReplyHandleNomination) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyHandleNomination) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Responseexecuteitem.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyHandleNomination) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Responseexecuteitem.Merge(m, src)
}
func (m *ReplyHandleNomination) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyHandleNomination) XXX_Dropunfamiliar() {
	xxx_signaldetails_Responseexecuteitem.DiscardUnknown(m)
}

var xxx_signaldetails_Responseexecuteitem proto.InternalMessageInfo

func (m *ReplyHandleNomination) ObtainCondition() Responseexecuteitem_Itemstatus {
	if m != nil {
		return m.Condition
	}
	return Responseexecuteitem_UNFAMILIAR
}

type ReplyBroadenBallot struct {
	BallotAddition []byte `protobuf:"octets,1,opt,name=vote_extension,json=voteExtension,proto3" json:"ballot_addition,omitempty"`
}

func (m *ReplyBroadenBallot) Restore()         { *m = ReplyBroadenBallot{} }
func (m *ReplyBroadenBallot) Text() string { return proto.CompactTextString(m) }
func (*ReplyBroadenBallot) SchemaArtifact()    {}
func (*ReplyBroadenBallot) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{36}
}
func (m *ReplyBroadenBallot) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyBroadenBallot) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyextendballot.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyBroadenBallot) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyextendballot.Merge(m, src)
}
func (m *ReplyBroadenBallot) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyBroadenBallot) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyextendballot.DiscardUnknown(m)
}

var xxx_signaldetails_Replyextendballot proto.InternalMessageInfo

func (m *ReplyBroadenBallot) ObtainBallotAddition() []byte {
	if m != nil {
		return m.BallotAddition
	}
	return nil
}

type ReplyValidateBallotAddition struct {
	Condition Responsecertifyballotaddition_Verifystatus `protobuf:"variableint,1,opt,name=status,proto3,enum=tendermint.abci.ResponseVerifyVoteExtension_VerifyStatus" json:"condition,omitempty"`
}

func (m *ReplyValidateBallotAddition) Restore()         { *m = ReplyValidateBallotAddition{} }
func (m *ReplyValidateBallotAddition) Text() string { return proto.CompactTextString(m) }
func (*ReplyValidateBallotAddition) SchemaArtifact()    {}
func (*ReplyValidateBallotAddition) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{37}
}
func (m *ReplyValidateBallotAddition) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyValidateBallotAddition) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Responsecertifyballotaddition.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyValidateBallotAddition) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Responsecertifyballotaddition.Merge(m, src)
}
func (m *ReplyValidateBallotAddition) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyValidateBallotAddition) XXX_Dropunfamiliar() {
	xxx_signaldetails_Responsecertifyballotaddition.DiscardUnknown(m)
}

var xxx_signaldetails_Responsecertifyballotaddition proto.InternalMessageInfo

func (m *ReplyValidateBallotAddition) ObtainCondition() Responsecertifyballotaddition_Verifystatus {
	if m != nil {
		return m.Condition
	}
	return Responsecertifyballotaddition_UNFAMILIAR
}

type ReplyCulminateLedger struct {
	//
	Incidents []Incident `protobuf:"octets,1,rep,name=events,proto3" json:"incidents,omitempty"`
	//
	//
	//
	TransferOutcomes []*InvokeTransferOutcome `protobuf:"octets,2,rep,name=tx_results,json=txResults,proto3" json:"transfer_outcomes,omitempty"`
	//
	AssessorRevisions []AssessorRevise `protobuf:"octets,3,rep,name=validator_updates,json=validatorUpdates,proto3" json:"assessor_revisions"`
	//
	AgreementArgumentRevisions *kinds1.AgreementSettings `protobuf:"octets,4,opt,name=consensus_param_updates,json=consensusParamUpdates,proto3" json:"agreement_argument_revisions,omitempty"`
	//
	PlatformDigest []byte `protobuf:"octets,5,opt,name=app_hash,json=appHash,proto3" json:"application_digest,omitempty"`
}

func (m *ReplyCulminateLedger) Restore()         { *m = ReplyCulminateLedger{} }
func (m *ReplyCulminateLedger) Text() string { return proto.CompactTextString(m) }
func (*ReplyCulminateLedger) SchemaArtifact()    {}
func (*ReplyCulminateLedger) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{38}
}
func (m *ReplyCulminateLedger) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyCulminateLedger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyfinalizeledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyCulminateLedger) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyfinalizeledger.Merge(m, src)
}
func (m *ReplyCulminateLedger) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyCulminateLedger) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyfinalizeledger.DiscardUnknown(m)
}

var xxx_signaldetails_Replyfinalizeledger proto.InternalMessageInfo

func (m *ReplyCulminateLedger) ObtainIncidents() []Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

func (m *ReplyCulminateLedger) ObtainTransferOutcomes() []*InvokeTransferOutcome {
	if m != nil {
		return m.TransferOutcomes
	}
	return nil
}

func (m *ReplyCulminateLedger) ObtainAssessorRevisions() []AssessorRevise {
	if m != nil {
		return m.AssessorRevisions
	}
	return nil
}

func (m *ReplyCulminateLedger) ObtainAgreementArgumentRevisions() *kinds1.AgreementSettings {
	if m != nil {
		return m.AgreementArgumentRevisions
	}
	return nil
}

func (m *ReplyCulminateLedger) ObtainApplicationDigest() []byte {
	if m != nil {
		return m.PlatformDigest
	}
	return nil
}

type EndorseDetails struct {
	Iteration int32      `protobuf:"variableint,1,opt,name=round,proto3" json:"iteration,omitempty"`
	Ballots []BallotDetails `protobuf:"octets,2,rep,name=votes,proto3" json:"ballots"`
}

func (m *EndorseDetails) Restore()         { *m = EndorseDetails{} }
func (m *EndorseDetails) Text() string { return proto.CompactTextString(m) }
func (*EndorseDetails) SchemaArtifact()    {}
func (*EndorseDetails) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{39}
}
func (m *EndorseDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *EndorseDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Endorseinfo.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndorseDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Endorseinfo.Merge(m, src)
}
func (m *EndorseDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *EndorseDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Endorseinfo.DiscardUnknown(m)
}

var xxx_signaldetails_Endorseinfo proto.InternalMessageInfo

func (m *EndorseDetails) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *EndorseDetails) ObtainBallots() []BallotDetails {
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
	Iteration int32 `protobuf:"variableint,1,opt,name=round,proto3" json:"iteration,omitempty"`
	//
	//
	Ballots []ExpandedBallotDetails `protobuf:"octets,2,rep,name=votes,proto3" json:"ballots"`
}

func (m *ExpandedEndorseDetails) Restore()         { *m = ExpandedEndorseDetails{} }
func (m *ExpandedEndorseDetails) Text() string { return proto.CompactTextString(m) }
func (*ExpandedEndorseDetails) SchemaArtifact()    {}
func (*ExpandedEndorseDetails) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{40}
}
func (m *ExpandedEndorseDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ExpandedEndorseDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Expandedendorseinfo.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedEndorseDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Expandedendorseinfo.Merge(m, src)
}
func (m *ExpandedEndorseDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *ExpandedEndorseDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Expandedendorseinfo.DiscardUnknown(m)
}

var xxx_signaldetails_Expandedendorseinfo proto.InternalMessageInfo

func (m *ExpandedEndorseDetails) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *ExpandedEndorseDetails) ObtainBallots() []ExpandedBallotDetails {
	if m != nil {
		return m.Ballots
	}
	return nil
}

//
//
//
type Incident struct {
	Kind       string           `protobuf:"octets,1,opt,name=type,proto3" json:"kind,omitempty"`
	Properties []IncidentProperty `protobuf:"octets,2,rep,name=attributes,proto3" json:"properties,omitempty"`
}

func (m *Incident) Restore()         { *m = Incident{} }
func (m *Incident) Text() string { return proto.CompactTextString(m) }
func (*Incident) SchemaArtifact()    {}
func (*Incident) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{41}
}
func (m *Incident) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Incident) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Incident.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Incident) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Incident.Merge(m, src)
}
func (m *Incident) XXX_Extent() int {
	return m.Extent()
}
func (m *Incident) XXX_Dropunfamiliar() {
	xxx_signaldetails_Incident.DiscardUnknown(m)
}

var xxx_signaldetails_Incident proto.InternalMessageInfo

func (m *Incident) ObtainKind() string {
	if m != nil {
		return m.Kind
	}
	return "REDACTED"
}

func (m *Incident) ObtainProperties() []IncidentProperty {
	if m != nil {
		return m.Properties
	}
	return nil
}

//
type IncidentProperty struct {
	Key   string `protobuf:"octets,1,opt,name=key,proto3" json:"key,omitempty"`
	Datum string `protobuf:"octets,2,opt,name=value,proto3" json:"datum,omitempty"`
	Ordinal bool   `protobuf:"variableint,3,opt,name=index,proto3" json:"ordinal,omitempty"`
}

func (m *IncidentProperty) Restore()         { *m = IncidentProperty{} }
func (m *IncidentProperty) Text() string { return proto.CompactTextString(m) }
func (*IncidentProperty) SchemaArtifact()    {}
func (*IncidentProperty) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{42}
}
func (m *IncidentProperty) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *IncidentProperty) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Incidentproperty.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncidentProperty) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Incidentproperty.Merge(m, src)
}
func (m *IncidentProperty) XXX_Extent() int {
	return m.Extent()
}
func (m *IncidentProperty) XXX_Dropunfamiliar() {
	xxx_signaldetails_Incidentproperty.DiscardUnknown(m)
}

var xxx_signaldetails_Incidentproperty proto.InternalMessageInfo

func (m *IncidentProperty) ObtainToken() string {
	if m != nil {
		return m.Key
	}
	return "REDACTED"
}

func (m *IncidentProperty) ObtainDatum() string {
	if m != nil {
		return m.Datum
	}
	return "REDACTED"
}

func (m *IncidentProperty) ObtainOrdinal() bool {
	if m != nil {
		return m.Ordinal
	}
	return false
}

//
//
//
type InvokeTransferOutcome struct {
	Cipher      uint32  `protobuf:"variableint,1,opt,name=code,proto3" json:"cipher,omitempty"`
	Data      []byte  `protobuf:"octets,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string  `protobuf:"octets,3,opt,name=log,proto3" json:"log,omitempty"`
	Details      string  `protobuf:"octets,4,opt,name=info,proto3" json:"details,omitempty"`
	FuelDesired int64   `protobuf:"variableint,5,opt,name=gas_wanted,proto3" json:"fuel_desired,omitempty"`
	FuelUtilized   int64   `protobuf:"variableint,6,opt,name=gas_used,proto3" json:"fuel_utilized,omitempty"`
	Incidents    []Incident `protobuf:"octets,7,rep,name=events,proto3" json:"incidents,omitempty"`
	Codeset string  `protobuf:"octets,8,opt,name=codespace,proto3" json:"codeset,omitempty"`
}

func (m *InvokeTransferOutcome) Restore()         { *m = InvokeTransferOutcome{} }
func (m *InvokeTransferOutcome) Text() string { return proto.CompactTextString(m) }
func (*InvokeTransferOutcome) SchemaArtifact()    {}
func (*InvokeTransferOutcome) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{43}
}
func (m *InvokeTransferOutcome) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *InvokeTransferOutcome) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Executecontextoutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InvokeTransferOutcome) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Executecontextoutcome.Merge(m, src)
}
func (m *InvokeTransferOutcome) XXX_Extent() int {
	return m.Extent()
}
func (m *InvokeTransferOutcome) XXX_Dropunfamiliar() {
	xxx_signaldetails_Executecontextoutcome.DiscardUnknown(m)
}

var xxx_signaldetails_Executecontextoutcome proto.InternalMessageInfo

func (m *InvokeTransferOutcome) ObtainCipher() uint32 {
	if m != nil {
		return m.Cipher
	}
	return 0
}

func (m *InvokeTransferOutcome) ObtainData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeTransferOutcome) ObtainRecord() string {
	if m != nil {
		return m.Log
	}
	return "REDACTED"
}

func (m *InvokeTransferOutcome) ObtainDetails() string {
	if m != nil {
		return m.Details
	}
	return "REDACTED"
}

func (m *InvokeTransferOutcome) ObtainFuelDesired() int64 {
	if m != nil {
		return m.FuelDesired
	}
	return 0
}

func (m *InvokeTransferOutcome) ObtainFuelUtilized() int64 {
	if m != nil {
		return m.FuelUtilized
	}
	return 0
}

func (m *InvokeTransferOutcome) ObtainIncidents() []Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

func (m *InvokeTransferOutcome) ObtainCodeset() string {
	if m != nil {
		return m.Codeset
	}
	return "REDACTED"
}

//
//
//
type TransferOutcome struct {
	Altitude int64        `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Ordinal  uint32       `protobuf:"variableint,2,opt,name=index,proto3" json:"ordinal,omitempty"`
	Tx     []byte       `protobuf:"octets,3,opt,name=tx,proto3" json:"tx,omitempty"`
	Outcome InvokeTransferOutcome `protobuf:"octets,4,opt,name=result,proto3" json:"outcome"`
}

func (m *TransferOutcome) Restore()         { *m = TransferOutcome{} }
func (m *TransferOutcome) Text() string { return proto.CompactTextString(m) }
func (*TransferOutcome) SchemaArtifact()    {}
func (*TransferOutcome) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{44}
}
func (m *TransferOutcome) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *TransferOutcome) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Transoutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferOutcome) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Transoutcome.Merge(m, src)
}
func (m *TransferOutcome) XXX_Extent() int {
	return m.Extent()
}
func (m *TransferOutcome) XXX_Dropunfamiliar() {
	xxx_signaldetails_Transoutcome.DiscardUnknown(m)
}

var xxx_signaldetails_Transoutcome proto.InternalMessageInfo

func (m *TransferOutcome) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *TransferOutcome) ObtainOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *TransferOutcome) ObtainTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TransferOutcome) ObtainOutcome() InvokeTransferOutcome {
	if m != nil {
		return m.Outcome
	}
	return InvokeTransferOutcome{}
}

type Assessor struct {
	Location []byte `protobuf:"octets,1,opt,name=address,proto3" json:"location,omitempty"`
	//
	Potency int64 `protobuf:"variableint,3,opt,name=power,proto3" json:"potency,omitempty"`
}

func (m *Assessor) Restore()         { *m = Assessor{} }
func (m *Assessor) Text() string { return proto.CompactTextString(m) }
func (*Assessor) SchemaArtifact()    {}
func (*Assessor) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{45}
}
func (m *Assessor) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Assessor) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Assessor.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Assessor) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Assessor.Merge(m, src)
}
func (m *Assessor) XXX_Extent() int {
	return m.Extent()
}
func (m *Assessor) XXX_Dropunfamiliar() {
	xxx_signaldetails_Assessor.DiscardUnknown(m)
}

var xxx_signaldetails_Assessor proto.InternalMessageInfo

func (m *Assessor) ObtainLocator() []byte {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Assessor) ObtainPotency() int64 {
	if m != nil {
		return m.Potency
	}
	return 0
}

type AssessorRevise struct {
	PublicToken security.CommonToken `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_token"`
	Potency  int64            `protobuf:"variableint,2,opt,name=power,proto3" json:"potency,omitempty"`
}

func (m *AssessorRevise) Restore()         { *m = AssessorRevise{} }
func (m *AssessorRevise) Text() string { return proto.CompactTextString(m) }
func (*AssessorRevise) SchemaArtifact()    {}
func (*AssessorRevise) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{46}
}
func (m *AssessorRevise) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AssessorRevise) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Assessorupdate.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssessorRevise) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Assessorupdate.Merge(m, src)
}
func (m *AssessorRevise) XXX_Extent() int {
	return m.Extent()
}
func (m *AssessorRevise) XXX_Dropunfamiliar() {
	xxx_signaldetails_Assessorupdate.DiscardUnknown(m)
}

var xxx_signaldetails_Assessorupdate proto.InternalMessageInfo

func (m *AssessorRevise) ObtainPublicToken() security.CommonToken {
	if m != nil {
		return m.PublicToken
	}
	return security.CommonToken{}
}

func (m *AssessorRevise) ObtainPotency() int64 {
	if m != nil {
		return m.Potency
	}
	return 0
}

type BallotDetails struct {
	Assessor   Assessor          `protobuf:"octets,1,opt,name=validator,proto3" json:"assessor"`
	LedgerUuidMarker kinds1.LedgerUUIDMarker `protobuf:"variableint,3,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uuid_marker,omitempty"`
}

func (m *BallotDetails) Restore()         { *m = BallotDetails{} }
func (m *BallotDetails) Text() string { return proto.CompactTextString(m) }
func (*BallotDetails) SchemaArtifact()    {}
func (*BallotDetails) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{47}
}
func (m *BallotDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *BallotDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ballotdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BallotDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ballotdetails.Merge(m, src)
}
func (m *BallotDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *BallotDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ballotdetails.DiscardUnknown(m)
}

var xxx_signaldetails_Ballotdetails proto.InternalMessageInfo

func (m *BallotDetails) ObtainAssessor() Assessor {
	if m != nil {
		return m.Assessor
	}
	return Assessor{}
}

func (m *BallotDetails) ObtainLedgerUuidMarker() kinds1.LedgerUUIDMarker {
	if m != nil {
		return m.LedgerUuidMarker
	}
	return kinds1.LedgerUUIDMarkerUnfamiliar
}

type ExpandedBallotDetails struct {
	//
	Assessor Assessor `protobuf:"octets,1,opt,name=validator,proto3" json:"assessor"`
	//
	BallotAddition []byte `protobuf:"octets,3,opt,name=vote_extension,json=voteExtension,proto3" json:"ballot_addition,omitempty"`
	//
	AdditionNotation []byte `protobuf:"octets,4,opt,name=extension_signature,json=extensionSignature,proto3" json:"addition_signing,omitempty"`
	//
	LedgerUuidMarker kinds1.LedgerUUIDMarker `protobuf:"variableint,5,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"ledger_uuid_marker,omitempty"`
}

func (m *ExpandedBallotDetails) Restore()         { *m = ExpandedBallotDetails{} }
func (m *ExpandedBallotDetails) Text() string { return proto.CompactTextString(m) }
func (*ExpandedBallotDetails) SchemaArtifact()    {}
func (*ExpandedBallotDetails) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{48}
}
func (m *ExpandedBallotDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ExpandedBallotDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Expandedballotdetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandedBallotDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Expandedballotdetails.Merge(m, src)
}
func (m *ExpandedBallotDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *ExpandedBallotDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Expandedballotdetails.DiscardUnknown(m)
}

var xxx_signaldetails_Expandedballotdetails proto.InternalMessageInfo

func (m *ExpandedBallotDetails) ObtainAssessor() Assessor {
	if m != nil {
		return m.Assessor
	}
	return Assessor{}
}

func (m *ExpandedBallotDetails) ObtainBallotAddition() []byte {
	if m != nil {
		return m.BallotAddition
	}
	return nil
}

func (m *ExpandedBallotDetails) ObtainAdditionSigning() []byte {
	if m != nil {
		return m.AdditionNotation
	}
	return nil
}

func (m *ExpandedBallotDetails) ObtainLedgerUuidMarker() kinds1.LedgerUUIDMarker {
	if m != nil {
		return m.LedgerUuidMarker
	}
	return kinds1.LedgerUUIDMarkerUnfamiliar
}

type Malpractice struct {
	Kind MalpracticeKind `protobuf:"variableint,1,opt,name=type,proto3,enum=tendermint.abci.MisbehaviorType" json:"kind,omitempty"`
	//
	Assessor Assessor `protobuf:"octets,2,opt,name=validator,proto3" json:"assessor"`
	//
	Altitude int64 `protobuf:"variableint,3,opt,name=height,proto3" json:"altitude,omitempty"`
	//
	Moment time.Time `protobuf:"octets,4,opt,name=time,proto3,stdtime" json:"moment"`
	//
	//
	//
	SumBallotingPotency int64 `protobuf:"variableint,5,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"sum_balloting_potency,omitempty"`
}

func (m *Malpractice) Restore()         { *m = Malpractice{} }
func (m *Malpractice) Text() string { return proto.CompactTextString(m) }
func (*Malpractice) SchemaArtifact()    {}
func (*Malpractice) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{49}
}
func (m *Malpractice) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Malpractice) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Malpractice.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Malpractice) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Malpractice.Merge(m, src)
}
func (m *Malpractice) XXX_Extent() int {
	return m.Extent()
}
func (m *Malpractice) XXX_Dropunfamiliar() {
	xxx_signaldetails_Malpractice.DiscardUnknown(m)
}

var xxx_signaldetails_Malpractice proto.InternalMessageInfo

func (m *Malpractice) ObtainKind() MalpracticeKind {
	if m != nil {
		return m.Kind
	}
	return Malfunctionkind_UNFAMILIAR
}

func (m *Malpractice) ObtainAssessor() Assessor {
	if m != nil {
		return m.Assessor
	}
	return Assessor{}
}

func (m *Malpractice) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Malpractice) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *Malpractice) ObtainSumBallotingPotency() int64 {
	if m != nil {
		return m.SumBallotingPotency
	}
	return 0
}

type Image struct {
	Altitude   uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Layout   uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Segments   uint32 `protobuf:"variableint,3,opt,name=chunks,proto3" json:"segments,omitempty"`
	Digest     []byte `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Attributes []byte `protobuf:"octets,5,opt,name=metadata,proto3" json:"attributes,omitempty"`
}

func (m *Image) Restore()         { *m = Image{} }
func (m *Image) Text() string { return proto.CompactTextString(m) }
func (*Image) SchemaArtifact()    {}
func (*Image) Definition() ([]byte, []int) {
	return filedescriptor_252557cfdd89a31a, []int{50}
}
func (m *Image) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Image) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Image.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Image) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Image.Merge(m, src)
}
func (m *Image) XXX_Extent() int {
	return m.Extent()
}
func (m *Image) XXX_Dropunfamiliar() {
	xxx_signaldetails_Image.DiscardUnknown(m)
}

var xxx_signaldetails_Image proto.InternalMessageInfo

func (m *Image) ObtainAltitude() uint64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Image) ObtainLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *Image) ObtainSegments() uint32 {
	if m != nil {
		return m.Segments
	}
	return 0
}

func (m *Image) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Image) ObtainAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func initialize() {
	proto.RegisterEnum("REDACTED", Inspecttranskind_alias, Inspecttranskind_datum)
	proto.RegisterEnum("REDACTED", Malfunctionkind_alias, Malfunctionkind_datum)
	proto.RegisterEnum("REDACTED", Replyextendimage_Outcome_alias, Replyextendimage_Outcome_datum)
	proto.RegisterEnum("REDACTED", Replyapplyimagefragment_Outcome_alias, Replyapplyimagefragment_Outcome_datum)
	proto.RegisterEnum("REDACTED", Responseexecuteitem_Itemstatus_alias, Responseexecuteitem_Itemstatus_datum)
	proto.RegisterEnum("REDACTED", Responsecertifyballotaddition_Verifystatus_alias, Responsecertifyballotaddition_Verifystatus_datum)
	proto.RegisterType((*Solicit)(nil), "REDACTED")
	proto.RegisterType((*SolicitReverberate)(nil), "REDACTED")
	proto.RegisterType((*SolicitPurge)(nil), "REDACTED")
	proto.RegisterType((*SolicitDetails)(nil), "REDACTED")
	proto.RegisterType((*SolicitInitializeSuccession)(nil), "REDACTED")
	proto.RegisterType((*SolicitInquire)(nil), "REDACTED")
	proto.RegisterType((*SolicitInspectTransfer)(nil), "REDACTED")
	proto.RegisterType((*SolicitAppendTransfer)(nil), "REDACTED")
	proto.RegisterType((*SolicitHarvestTrans)(nil), "REDACTED")
	proto.RegisterType((*SolicitEndorse)(nil), "REDACTED")
	proto.RegisterType((*SolicitCollectionImages)(nil), "REDACTED")
	proto.RegisterType((*SolicitExtendImage)(nil), "REDACTED")
	proto.RegisterType((*SolicitFetchImageSegment)(nil), "REDACTED")
	proto.RegisterType((*SolicitExecuteImageSegment)(nil), "REDACTED")
	proto.RegisterType((*SolicitArrangeNomination)(nil), "REDACTED")
	proto.RegisterType((*SolicitHandleNomination)(nil), "REDACTED")
	proto.RegisterType((*SolicitBroadenBallot)(nil), "REDACTED")
	proto.RegisterType((*SolicitValidateBallotAddition)(nil), "REDACTED")
	proto.RegisterType((*SolicitCulminateLedger)(nil), "REDACTED")
	proto.RegisterType((*Reply)(nil), "REDACTED")
	proto.RegisterType((*ReplyExemption)(nil), "REDACTED")
	proto.RegisterType((*ReplyReverberate)(nil), "REDACTED")
	proto.RegisterType((*ReplyPurge)(nil), "REDACTED")
	proto.RegisterType((*ReplyDetails)(nil), "REDACTED")
	proto.RegisterType((*ReplyInitializeSuccession)(nil), "REDACTED")
	proto.RegisterType((*ReplyInquire)(nil), "REDACTED")
	proto.RegisterType((*ReplyInspectTransfer)(nil), "REDACTED")
	proto.RegisterType((*ReplyAppendTransfer)(nil), "REDACTED")
	proto.RegisterType((*ReplyHarvestTrans)(nil), "REDACTED")
	proto.RegisterType((*ReplyEndorse)(nil), "REDACTED")
	proto.RegisterType((*ReplyCatalogImages)(nil), "REDACTED")
	proto.RegisterType((*ReplyExtendImage)(nil), "REDACTED")
	proto.RegisterType((*ReplyFetchImageSegment)(nil), "REDACTED")
	proto.RegisterType((*ReplyExecuteImageSegment)(nil), "REDACTED")
	proto.RegisterType((*ReplyArrangeNomination)(nil), "REDACTED")
	proto.RegisterType((*ReplyHandleNomination)(nil), "REDACTED")
	proto.RegisterType((*ReplyBroadenBallot)(nil), "REDACTED")
	proto.RegisterType((*ReplyValidateBallotAddition)(nil), "REDACTED")
	proto.RegisterType((*ReplyCulminateLedger)(nil), "REDACTED")
	proto.RegisterType((*EndorseDetails)(nil), "REDACTED")
	proto.RegisterType((*ExpandedEndorseDetails)(nil), "REDACTED")
	proto.RegisterType((*Incident)(nil), "REDACTED")
	proto.RegisterType((*IncidentProperty)(nil), "REDACTED")
	proto.RegisterType((*InvokeTransferOutcome)(nil), "REDACTED")
	proto.RegisterType((*TransferOutcome)(nil), "REDACTED")
	proto.RegisterType((*Assessor)(nil), "REDACTED")
	proto.RegisterType((*AssessorRevise)(nil), "REDACTED")
	proto.RegisterType((*BallotDetails)(nil), "REDACTED")
	proto.RegisterType((*ExpandedBallotDetails)(nil), "REDACTED")
	proto.RegisterType((*Malpractice)(nil), "REDACTED")
	proto.RegisterType((*Image)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_252557cfdd89a31a) }

var filedescriptor_252557cfdd89a31a = []byte{
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
	Reverberate(ctx context.Context, in *SolicitReverberate, choices ...grpc.CallOption) (*ReplyReverberate, error)
	Purge(ctx context.Context, in *SolicitPurge, choices ...grpc.CallOption) (*ReplyPurge, error)
	Details(ctx context.Context, in *SolicitDetails, choices ...grpc.CallOption) (*ReplyDetails, error)
	InspectTransfer(ctx context.Context, in *SolicitInspectTransfer, choices ...grpc.CallOption) (*ReplyInspectTransfer, error)
	AppendTransfer(ctx context.Context, in *SolicitAppendTransfer, choices ...grpc.CallOption) (*ReplyAppendTransfer, error)
	HarvestTrans(ctx context.Context, in *SolicitHarvestTrans, choices ...grpc.CallOption) (*ReplyHarvestTrans, error)
	Inquire(ctx context.Context, in *SolicitInquire, choices ...grpc.CallOption) (*ReplyInquire, error)
	Endorse(ctx context.Context, in *SolicitEndorse, choices ...grpc.CallOption) (*ReplyEndorse, error)
	InitializeSuccession(ctx context.Context, in *SolicitInitializeSuccession, choices ...grpc.CallOption) (*ReplyInitializeSuccession, error)
	CollectionImages(ctx context.Context, in *SolicitCollectionImages, choices ...grpc.CallOption) (*ReplyCatalogImages, error)
	ExtendImage(ctx context.Context, in *SolicitExtendImage, choices ...grpc.CallOption) (*ReplyExtendImage, error)
	FetchImageSegment(ctx context.Context, in *SolicitFetchImageSegment, choices ...grpc.CallOption) (*ReplyFetchImageSegment, error)
	ExecuteImageSegment(ctx context.Context, in *SolicitExecuteImageSegment, choices ...grpc.CallOption) (*ReplyExecuteImageSegment, error)
	ArrangeNomination(ctx context.Context, in *SolicitArrangeNomination, choices ...grpc.CallOption) (*ReplyArrangeNomination, error)
	HandleNomination(ctx context.Context, in *SolicitHandleNomination, choices ...grpc.CallOption) (*ReplyHandleNomination, error)
	BroadenBallot(ctx context.Context, in *SolicitBroadenBallot, choices ...grpc.CallOption) (*ReplyBroadenBallot, error)
	ValidateBallotAddition(ctx context.Context, in *SolicitValidateBallotAddition, choices ...grpc.CallOption) (*ReplyValidateBallotAddition, error)
	CulminateLedger(ctx context.Context, in *SolicitCulminateLedger, choices ...grpc.CallOption) (*ReplyCulminateLedger, error)
}

type anBCICustomer struct {
	cc grpc1.ClientConn
}

func FreshIfaceCustomer(cc grpc1.ClientConn) IfaceCustomer {
	return &anBCICustomer{cc}
}

func (c *anBCICustomer) Reverberate(ctx context.Context, in *SolicitReverberate, choices ...grpc.CallOption) (*ReplyReverberate, error) {
	out := new(ReplyReverberate)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) Purge(ctx context.Context, in *SolicitPurge, choices ...grpc.CallOption) (*ReplyPurge, error) {
	out := new(ReplyPurge)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) Details(ctx context.Context, in *SolicitDetails, choices ...grpc.CallOption) (*ReplyDetails, error) {
	out := new(ReplyDetails)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) InspectTransfer(ctx context.Context, in *SolicitInspectTransfer, choices ...grpc.CallOption) (*ReplyInspectTransfer, error) {
	out := new(ReplyInspectTransfer)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) AppendTransfer(ctx context.Context, in *SolicitAppendTransfer, choices ...grpc.CallOption) (*ReplyAppendTransfer, error) {
	out := new(ReplyAppendTransfer)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) HarvestTrans(ctx context.Context, in *SolicitHarvestTrans, choices ...grpc.CallOption) (*ReplyHarvestTrans, error) {
	out := new(ReplyHarvestTrans)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) Inquire(ctx context.Context, in *SolicitInquire, choices ...grpc.CallOption) (*ReplyInquire, error) {
	out := new(ReplyInquire)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) Endorse(ctx context.Context, in *SolicitEndorse, choices ...grpc.CallOption) (*ReplyEndorse, error) {
	out := new(ReplyEndorse)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) InitializeSuccession(ctx context.Context, in *SolicitInitializeSuccession, choices ...grpc.CallOption) (*ReplyInitializeSuccession, error) {
	out := new(ReplyInitializeSuccession)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) CollectionImages(ctx context.Context, in *SolicitCollectionImages, choices ...grpc.CallOption) (*ReplyCatalogImages, error) {
	out := new(ReplyCatalogImages)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) ExtendImage(ctx context.Context, in *SolicitExtendImage, choices ...grpc.CallOption) (*ReplyExtendImage, error) {
	out := new(ReplyExtendImage)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) FetchImageSegment(ctx context.Context, in *SolicitFetchImageSegment, choices ...grpc.CallOption) (*ReplyFetchImageSegment, error) {
	out := new(ReplyFetchImageSegment)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) ExecuteImageSegment(ctx context.Context, in *SolicitExecuteImageSegment, choices ...grpc.CallOption) (*ReplyExecuteImageSegment, error) {
	out := new(ReplyExecuteImageSegment)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) ArrangeNomination(ctx context.Context, in *SolicitArrangeNomination, choices ...grpc.CallOption) (*ReplyArrangeNomination, error) {
	out := new(ReplyArrangeNomination)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) HandleNomination(ctx context.Context, in *SolicitHandleNomination, choices ...grpc.CallOption) (*ReplyHandleNomination, error) {
	out := new(ReplyHandleNomination)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) BroadenBallot(ctx context.Context, in *SolicitBroadenBallot, choices ...grpc.CallOption) (*ReplyBroadenBallot, error) {
	out := new(ReplyBroadenBallot)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) ValidateBallotAddition(ctx context.Context, in *SolicitValidateBallotAddition, choices ...grpc.CallOption) (*ReplyValidateBallotAddition, error) {
	out := new(ReplyValidateBallotAddition)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anBCICustomer) CulminateLedger(ctx context.Context, in *SolicitCulminateLedger, choices ...grpc.CallOption) (*ReplyCulminateLedger, error) {
	out := new(ReplyCulminateLedger)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//
type IfaceDaemon interface {
	Reverberate(context.Context, *SolicitReverberate) (*ReplyReverberate, error)
	Purge(context.Context, *SolicitPurge) (*ReplyPurge, error)
	Details(context.Context, *SolicitDetails) (*ReplyDetails, error)
	InspectTransfer(context.Context, *SolicitInspectTransfer) (*ReplyInspectTransfer, error)
	AppendTransfer(context.Context, *SolicitAppendTransfer) (*ReplyAppendTransfer, error)
	HarvestTrans(context.Context, *SolicitHarvestTrans) (*ReplyHarvestTrans, error)
	Inquire(context.Context, *SolicitInquire) (*ReplyInquire, error)
	Endorse(context.Context, *SolicitEndorse) (*ReplyEndorse, error)
	InitializeSuccession(context.Context, *SolicitInitializeSuccession) (*ReplyInitializeSuccession, error)
	CollectionImages(context.Context, *SolicitCollectionImages) (*ReplyCatalogImages, error)
	ExtendImage(context.Context, *SolicitExtendImage) (*ReplyExtendImage, error)
	FetchImageSegment(context.Context, *SolicitFetchImageSegment) (*ReplyFetchImageSegment, error)
	ExecuteImageSegment(context.Context, *SolicitExecuteImageSegment) (*ReplyExecuteImageSegment, error)
	ArrangeNomination(context.Context, *SolicitArrangeNomination) (*ReplyArrangeNomination, error)
	HandleNomination(context.Context, *SolicitHandleNomination) (*ReplyHandleNomination, error)
	BroadenBallot(context.Context, *SolicitBroadenBallot) (*ReplyBroadenBallot, error)
	ValidateBallotAddition(context.Context, *SolicitValidateBallotAddition) (*ReplyValidateBallotAddition, error)
	CulminateLedger(context.Context, *SolicitCulminateLedger) (*ReplyCulminateLedger, error)
}

//
type UndevelopedIfaceDaemon struct {
}

func (*UndevelopedIfaceDaemon) Reverberate(ctx context.Context, req *SolicitReverberate) (*ReplyReverberate, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) Purge(ctx context.Context, req *SolicitPurge) (*ReplyPurge, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) Details(ctx context.Context, req *SolicitDetails) (*ReplyDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) InspectTransfer(ctx context.Context, req *SolicitInspectTransfer) (*ReplyInspectTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) AppendTransfer(ctx context.Context, req *SolicitAppendTransfer) (*ReplyAppendTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) HarvestTrans(ctx context.Context, req *SolicitHarvestTrans) (*ReplyHarvestTrans, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) Inquire(ctx context.Context, req *SolicitInquire) (*ReplyInquire, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) Endorse(ctx context.Context, req *SolicitEndorse) (*ReplyEndorse, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) InitializeSuccession(ctx context.Context, req *SolicitInitializeSuccession) (*ReplyInitializeSuccession, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) CollectionImages(ctx context.Context, req *SolicitCollectionImages) (*ReplyCatalogImages, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) ExtendImage(ctx context.Context, req *SolicitExtendImage) (*ReplyExtendImage, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) FetchImageSegment(ctx context.Context, req *SolicitFetchImageSegment) (*ReplyFetchImageSegment, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) ExecuteImageSegment(ctx context.Context, req *SolicitExecuteImageSegment) (*ReplyExecuteImageSegment, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) ArrangeNomination(ctx context.Context, req *SolicitArrangeNomination) (*ReplyArrangeNomination, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) HandleNomination(ctx context.Context, req *SolicitHandleNomination) (*ReplyHandleNomination, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) BroadenBallot(ctx context.Context, req *SolicitBroadenBallot) (*ReplyBroadenBallot, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) ValidateBallotAddition(ctx context.Context, req *SolicitValidateBallotAddition) (*ReplyValidateBallotAddition, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedIfaceDaemon) CulminateLedger(ctx context.Context, req *SolicitCulminateLedger) (*ReplyCulminateLedger, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}

func EnrollIfaceDaemon(s grpc1.Server, srv IfaceDaemon) {
	s.RegisterService(&_IFACE_servicedetails, srv)
}

func _IFACE_Reverberate_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitReverberate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).Reverberate(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).Reverberate(ctx, req.(*SolicitReverberate))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Purge_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitPurge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).Purge(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).Purge(ctx, req.(*SolicitPurge))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Details_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).Details(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).Details(ctx, req.(*SolicitDetails))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Inspecttrans_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitInspectTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).InspectTransfer(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).InspectTransfer(ctx, req.(*SolicitInspectTransfer))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Appendtrans_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitAppendTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).AppendTransfer(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).AppendTransfer(ctx, req.(*SolicitAppendTransfer))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Harvesttrans_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitHarvestTrans)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).HarvestTrans(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).HarvestTrans(ctx, req.(*SolicitHarvestTrans))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Inquire_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitInquire)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).Inquire(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).Inquire(ctx, req.(*SolicitInquire))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Endorse_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitEndorse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).Endorse(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).Endorse(ctx, req.(*SolicitEndorse))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Initiatechain_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitInitializeSuccession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).InitializeSuccession(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).InitializeSuccession(ctx, req.(*SolicitInitializeSuccession))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Catalogimages_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitCollectionImages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).CollectionImages(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).CollectionImages(ctx, req.(*SolicitCollectionImages))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Extendimage_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitExtendImage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).ExtendImage(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).ExtendImage(ctx, req.(*SolicitExtendImage))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Loadimagefragment_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitFetchImageSegment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).FetchImageSegment(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).FetchImageSegment(ctx, req.(*SolicitFetchImageSegment))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Executeimagefragment_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitExecuteImageSegment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).ExecuteImageSegment(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).ExecuteImageSegment(ctx, req.(*SolicitExecuteImageSegment))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Prepareitem_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitArrangeNomination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).ArrangeNomination(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).ArrangeNomination(ctx, req.(*SolicitArrangeNomination))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Executeitem_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitHandleNomination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).HandleNomination(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).HandleNomination(ctx, req.(*SolicitHandleNomination))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Extendballot_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitBroadenBallot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).BroadenBallot(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).BroadenBallot(ctx, req.(*SolicitBroadenBallot))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Verifyballotaddition_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitValidateBallotAddition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).ValidateBallotAddition(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).ValidateBallotAddition(ctx, req.(*SolicitValidateBallotAddition))
	}
	return overseer(ctx, in, details, processor)
}

func _IFACE_Finalizeledger_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitCulminateLedger)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(IfaceDaemon).CulminateLedger(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IfaceDaemon).CulminateLedger(ctx, req.(*SolicitCulminateLedger))
	}
	return overseer(ctx, in, details, processor)
}

var IFACE_servicedetails = _IFACE_servicedetails
var _IFACE_servicedetails = grpc.ServiceDesc{
	ServiceName: "REDACTED",
	HandlerType: (*IfaceDaemon)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Reverberate_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Purge_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Details_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Inspecttrans_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Appendtrans_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Harvesttrans_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Inquire_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Endorse_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Initiatechain_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Catalogimages_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Extendimage_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Loadimagefragment_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Executeimagefragment_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Prepareitem_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Executeitem_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Extendballot_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Verifyballotaddition_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _IFACE_Finalizeledger_Processor,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "REDACTED",
}

func (m *Solicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Solicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Datum != nil {
		{
			extent := m.Datum.Extent()
			i -= extent
			if _, err := m.Datum.SerializeToward(deltaLocatedAN[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Solicit_Reverberate) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Reverberate) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Reverberate != nil {
		{
			extent, err := m.Reverberate.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Solicit_Purge) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Purge) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Purge != nil {
		{
			extent, err := m.Purge.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Solicit_Details) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Details) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Details != nil {
		{
			extent, err := m.Details.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Solicit_Initiatechain) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Initiatechain) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.InitializeSuccession != nil {
		{
			extent, err := m.InitializeSuccession.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Solicit_Inquire) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Inquire) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Inquire != nil {
		{
			extent, err := m.Inquire.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Solicit_Inspecttrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Inspecttrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.InspectTransfer != nil {
		{
			extent, err := m.InspectTransfer.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Solicit_Endorse) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Endorse) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
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
		deltaLocatedAN[i] = 0x5a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Catalogimages) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Catalogimages) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.CollectionImages != nil {
		{
			extent, err := m.CollectionImages.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x62
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Extendimage) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Extendimage) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ExtendImage != nil {
		{
			extent, err := m.ExtendImage.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x6a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Loadimagefragment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Loadimagefragment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.FetchImageSegment != nil {
		{
			extent, err := m.FetchImageSegment.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x72
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Executeimagefragment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Executeimagefragment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ExecuteImageSegment != nil {
		{
			extent, err := m.ExecuteImageSegment.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x7a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Prepareitem) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Prepareitem) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ArrangeNomination != nil {
		{
			extent, err := m.ArrangeNomination.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x82
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Executeitem) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Executeitem) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.HandleNomination != nil {
		{
			extent, err := m.HandleNomination.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x8a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Extendballot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Extendballot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.BroadenBallot != nil {
		{
			extent, err := m.BroadenBallot.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x92
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Verifyballotaddition) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Verifyballotaddition) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ValidateBallotAddition != nil {
		{
			extent, err := m.ValidateBallotAddition.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x9a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Finalizeledger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Finalizeledger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.CulminateLedger != nil {
		{
			extent, err := m.CulminateLedger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0xa2
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Appendtrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Appendtrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.AppendTransfer != nil {
		{
			extent, err := m.AppendTransfer.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0xaa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Solicit_Harvesttrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Solicit_Harvesttrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.HarvestTrans != nil {
		{
			extent, err := m.HarvestTrans.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0xb2
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *SolicitReverberate) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitReverberate) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitReverberate) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Signal) > 0 {
		i -= len(m.Signal)
		copy(deltaLocatedAN[i:], m.Signal)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Signal)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitPurge) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitPurge) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitPurge) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.IfaceEdition) > 0 {
		i -= len(m.IfaceEdition)
		copy(deltaLocatedAN[i:], m.IfaceEdition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.IfaceEdition)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Peer2peerEdition != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Peer2peerEdition))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.LedgerEdition != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.LedgerEdition))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if len(m.Edition) > 0 {
		i -= len(m.Edition)
		copy(deltaLocatedAN[i:], m.Edition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Edition)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitInitializeSuccession) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitInitializeSuccession) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitInitializeSuccession) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.PrimaryAltitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.PrimaryAltitude))
		i--
		deltaLocatedAN[i] = 0x30
	}
	if len(m.ApplicationStatusOctets) > 0 {
		i -= len(m.ApplicationStatusOctets)
		copy(deltaLocatedAN[i:], m.ApplicationStatusOctets)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.ApplicationStatusOctets)))
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if len(m.Assessors) > 0 {
		for idxNdExc := len(m.Assessors) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Assessors[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
	if m.AgreementSettings != nil {
		{
			extent, err := m.AgreementSettings.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if len(m.SuccessionUuid) > 0 {
		i -= len(m.SuccessionUuid)
		copy(deltaLocatedAN[i:], m.SuccessionUuid)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.SuccessionUuid)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	n20, fault20 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault20 != nil {
		return 0, fault20
	}
	i -= n20
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n20))
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitInquire) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitInquire) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitInquire) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Validate {
		i--
		if m.Validate {
			deltaLocatedAN[i] = 1
		} else {
			deltaLocatedAN[i] = 0
		}
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if len(m.Route) > 0 {
		i -= len(m.Route)
		copy(deltaLocatedAN[i:], m.Route)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Route)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(deltaLocatedAN[i:], m.Data)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Data)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitInspectTransfer) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitInspectTransfer) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitInspectTransfer) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(deltaLocatedAN[i:], m.Tx)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Tx)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitAppendTransfer) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitAppendTransfer) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitAppendTransfer) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(deltaLocatedAN[i:], m.Tx)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Tx)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitHarvestTrans) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitHarvestTrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitHarvestTrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.MaximumFuel != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.MaximumFuel))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.MaximumOctets != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.MaximumOctets))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitEndorse) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitEndorse) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitEndorse) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitCollectionImages) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitCollectionImages) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitCollectionImages) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitExtendImage) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitExtendImage) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitExtendImage) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.PlatformDigest) > 0 {
		i -= len(m.PlatformDigest)
		copy(deltaLocatedAN[i:], m.PlatformDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.PlatformDigest)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Image != nil {
		{
			extent, err := m.Image.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *SolicitFetchImageSegment) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitFetchImageSegment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitFetchImageSegment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Segment != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Segment))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Layout != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Layout))
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

func (m *SolicitExecuteImageSegment) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitExecuteImageSegment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitExecuteImageSegment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Originator) > 0 {
		i -= len(m.Originator)
		copy(deltaLocatedAN[i:], m.Originator)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Originator)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if len(m.Segment) > 0 {
		i -= len(m.Segment)
		copy(deltaLocatedAN[i:], m.Segment)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Segment)))
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

func (m *SolicitArrangeNomination) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitArrangeNomination) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitArrangeNomination) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.NominatorLocation) > 0 {
		i -= len(m.NominatorLocation)
		copy(deltaLocatedAN[i:], m.NominatorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.NominatorLocation)))
		i--
		deltaLocatedAN[i] = 0x42
	}
	if len(m.FollowingAssessorsDigest) > 0 {
		i -= len(m.FollowingAssessorsDigest)
		copy(deltaLocatedAN[i:], m.FollowingAssessorsDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FollowingAssessorsDigest)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	n22, fault22 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault22 != nil {
		return 0, fault22
	}
	i -= n22
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n22))
	i--
	deltaLocatedAN[i] = 0x32
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.Malpractice) > 0 {
		for idxNdExc := len(m.Malpractice) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Malpractice[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
		extent, err := m.RegionalFinalEndorse.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x1a
	if len(m.Txs) > 0 {
		for idxNdExc := len(m.Txs) - 1; idxNdExc >= 0; idxNdExc-- {
			i -= len(m.Txs[idxNdExc])
			copy(deltaLocatedAN[i:], m.Txs[idxNdExc])
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Txs[idxNdExc])))
			i--
			deltaLocatedAN[i] = 0x12
		}
	}
	if m.MaximumTransferOctets != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.MaximumTransferOctets))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitHandleNomination) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitHandleNomination) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitHandleNomination) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.NominatorLocation) > 0 {
		i -= len(m.NominatorLocation)
		copy(deltaLocatedAN[i:], m.NominatorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.NominatorLocation)))
		i--
		deltaLocatedAN[i] = 0x42
	}
	if len(m.FollowingAssessorsDigest) > 0 {
		i -= len(m.FollowingAssessorsDigest)
		copy(deltaLocatedAN[i:], m.FollowingAssessorsDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FollowingAssessorsDigest)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	n24, fault24 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault24 != nil {
		return 0, fault24
	}
	i -= n24
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n24))
	i--
	deltaLocatedAN[i] = 0x32
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if len(m.Malpractice) > 0 {
		for idxNdExc := len(m.Malpractice) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Malpractice[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
	{
		extent, err := m.ItemizedFinalEndorse.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x12
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

func (m *SolicitBroadenBallot) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitBroadenBallot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitBroadenBallot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.NominatorLocation) > 0 {
		i -= len(m.NominatorLocation)
		copy(deltaLocatedAN[i:], m.NominatorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.NominatorLocation)))
		i--
		deltaLocatedAN[i] = 0x42
	}
	if len(m.FollowingAssessorsDigest) > 0 {
		i -= len(m.FollowingAssessorsDigest)
		copy(deltaLocatedAN[i:], m.FollowingAssessorsDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FollowingAssessorsDigest)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	if len(m.Malpractice) > 0 {
		for idxNdExc := len(m.Malpractice) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Malpractice[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x32
		}
	}
	{
		extent, err := m.ItemizedFinalEndorse.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x2a
	if len(m.Txs) > 0 {
		for idxNdExc := len(m.Txs) - 1; idxNdExc >= 0; idxNdExc-- {
			i -= len(m.Txs[idxNdExc])
			copy(deltaLocatedAN[i:], m.Txs[idxNdExc])
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Txs[idxNdExc])))
			i--
			deltaLocatedAN[i] = 0x22
		}
	}
	n27, fault27 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault27 != nil {
		return 0, fault27
	}
	i -= n27
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n27))
	i--
	deltaLocatedAN[i] = 0x1a
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitValidateBallotAddition) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitValidateBallotAddition) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitValidateBallotAddition) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.BallotAddition) > 0 {
		i -= len(m.BallotAddition)
		copy(deltaLocatedAN[i:], m.BallotAddition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.BallotAddition)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if len(m.AssessorLocation) > 0 {
		i -= len(m.AssessorLocation)
		copy(deltaLocatedAN[i:], m.AssessorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AssessorLocation)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitCulminateLedger) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitCulminateLedger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitCulminateLedger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.NominatorLocation) > 0 {
		i -= len(m.NominatorLocation)
		copy(deltaLocatedAN[i:], m.NominatorLocation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.NominatorLocation)))
		i--
		deltaLocatedAN[i] = 0x42
	}
	if len(m.FollowingAssessorsDigest) > 0 {
		i -= len(m.FollowingAssessorsDigest)
		copy(deltaLocatedAN[i:], m.FollowingAssessorsDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FollowingAssessorsDigest)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	n28, fault28 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault28 != nil {
		return 0, fault28
	}
	i -= n28
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n28))
	i--
	deltaLocatedAN[i] = 0x32
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if len(m.Malpractice) > 0 {
		for idxNdExc := len(m.Malpractice) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Malpractice[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
	{
		extent, err := m.ResolvedFinalEndorse.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x12
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

func (m *Reply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Reply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Datum != nil {
		{
			extent := m.Datum.Extent()
			i -= extent
			if _, err := m.Datum.SerializeToward(deltaLocatedAN[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Reply_Exemption) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Exemption) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Exemption != nil {
		{
			extent, err := m.Exemption.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Reply_Reverberate) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Reverberate) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Reverberate != nil {
		{
			extent, err := m.Reverberate.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Reply_Purge) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Purge) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Purge != nil {
		{
			extent, err := m.Purge.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Reply_Details) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Details) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Details != nil {
		{
			extent, err := m.Details.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Reply_Initiatechain) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Initiatechain) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.InitializeSuccession != nil {
		{
			extent, err := m.InitializeSuccession.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Reply_Inquire) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Inquire) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Inquire != nil {
		{
			extent, err := m.Inquire.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Reply_Inspecttrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Inspecttrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.InspectTransfer != nil {
		{
			extent, err := m.InspectTransfer.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Reply_Endorse) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Endorse) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
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
		deltaLocatedAN[i] = 0x62
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Catalogimages) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Catalogimages) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.CollectionImages != nil {
		{
			extent, err := m.CollectionImages.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x6a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Extendimage) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Extendimage) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ExtendImage != nil {
		{
			extent, err := m.ExtendImage.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x72
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Loadimagefragment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Loadimagefragment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.FetchImageSegment != nil {
		{
			extent, err := m.FetchImageSegment.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x7a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Executeimagefragment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Executeimagefragment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ExecuteImageSegment != nil {
		{
			extent, err := m.ExecuteImageSegment.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x82
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Prepareitem) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Prepareitem) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ArrangeNomination != nil {
		{
			extent, err := m.ArrangeNomination.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x8a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Executeitem) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Executeitem) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.HandleNomination != nil {
		{
			extent, err := m.HandleNomination.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x92
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Extendballot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Extendballot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.BroadenBallot != nil {
		{
			extent, err := m.BroadenBallot.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0x9a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Verifyballotaddition) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Verifyballotaddition) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ValidateBallotAddition != nil {
		{
			extent, err := m.ValidateBallotAddition.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0xa2
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Finalizeledger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Finalizeledger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.CulminateLedger != nil {
		{
			extent, err := m.CulminateLedger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0xaa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Appendtrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Appendtrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.AppendTransfer != nil {
		{
			extent, err := m.AppendTransfer.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0xb2
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Reply_Harvesttrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Reply_Harvesttrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.HarvestTrans != nil {
		{
			extent, err := m.HarvestTrans.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1
		i--
		deltaLocatedAN[i] = 0xba
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *ReplyExemption) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyExemption) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyExemption) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Failure) > 0 {
		i -= len(m.Failure)
		copy(deltaLocatedAN[i:], m.Failure)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Failure)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyReverberate) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyReverberate) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyReverberate) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Signal) > 0 {
		i -= len(m.Signal)
		copy(deltaLocatedAN[i:], m.Signal)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Signal)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyPurge) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyPurge) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyPurge) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.FinalLedgerPlatformDigest) > 0 {
		i -= len(m.FinalLedgerPlatformDigest)
		copy(deltaLocatedAN[i:], m.FinalLedgerPlatformDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.FinalLedgerPlatformDigest)))
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if m.FinalLedgerAltitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FinalLedgerAltitude))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.PlatformEdition != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.PlatformEdition))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if len(m.Edition) > 0 {
		i -= len(m.Edition)
		copy(deltaLocatedAN[i:], m.Edition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Edition)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(deltaLocatedAN[i:], m.Data)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Data)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyInitializeSuccession) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyInitializeSuccession) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyInitializeSuccession) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.PlatformDigest) > 0 {
		i -= len(m.PlatformDigest)
		copy(deltaLocatedAN[i:], m.PlatformDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.PlatformDigest)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if len(m.Assessors) > 0 {
		for idxNdExc := len(m.Assessors) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Assessors[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x12
		}
	}
	if m.AgreementSettings != nil {
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
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyInquire) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyInquire) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyInquire) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Codeset) > 0 {
		i -= len(m.Codeset)
		copy(deltaLocatedAN[i:], m.Codeset)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Codeset)))
		i--
		deltaLocatedAN[i] = 0x52
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x48
	}
	if m.AttestationActions != nil {
		{
			extent, err := m.AttestationActions.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x42
	}
	if len(m.Datum) > 0 {
		i -= len(m.Datum)
		copy(deltaLocatedAN[i:], m.Datum)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Datum)))
		i--
		deltaLocatedAN[i] = 0x3a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(deltaLocatedAN[i:], m.Key)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Key)))
		i--
		deltaLocatedAN[i] = 0x32
	}
	if m.Ordinal != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ordinal))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(deltaLocatedAN[i:], m.Details)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Details)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(deltaLocatedAN[i:], m.Log)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Log)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.Cipher != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Cipher))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyInspectTransfer) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyInspectTransfer) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyInspectTransfer) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Codeset) > 0 {
		i -= len(m.Codeset)
		copy(deltaLocatedAN[i:], m.Codeset)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Codeset)))
		i--
		deltaLocatedAN[i] = 0x42
	}
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
			deltaLocatedAN[i] = 0x3a
		}
	}
	if m.FuelUtilized != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FuelUtilized))
		i--
		deltaLocatedAN[i] = 0x30
	}
	if m.FuelDesired != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FuelDesired))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(deltaLocatedAN[i:], m.Details)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Details)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(deltaLocatedAN[i:], m.Log)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Log)))
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
	if m.Cipher != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Cipher))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyAppendTransfer) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyAppendTransfer) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyAppendTransfer) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Cipher != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Cipher))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyHarvestTrans) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyHarvestTrans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyHarvestTrans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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

func (m *ReplyEndorse) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyEndorse) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyEndorse) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.PreserveAltitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.PreserveAltitude))
		i--
		deltaLocatedAN[i] = 0x18
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyCatalogImages) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyCatalogImages) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyCatalogImages) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Images) > 0 {
		for idxNdExc := len(m.Images) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Images[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *ReplyExtendImage) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyExtendImage) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyExtendImage) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Outcome != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Outcome))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyFetchImageSegment) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyFetchImageSegment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyFetchImageSegment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Segment) > 0 {
		i -= len(m.Segment)
		copy(deltaLocatedAN[i:], m.Segment)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Segment)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyExecuteImageSegment) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyExecuteImageSegment) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyExecuteImageSegment) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.DeclineOriginators) > 0 {
		for idxNdExc := len(m.DeclineOriginators) - 1; idxNdExc >= 0; idxNdExc-- {
			i -= len(m.DeclineOriginators[idxNdExc])
			copy(deltaLocatedAN[i:], m.DeclineOriginators[idxNdExc])
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.DeclineOriginators[idxNdExc])))
			i--
			deltaLocatedAN[i] = 0x1a
		}
	}
	if len(m.RetrieveSegments) > 0 {
		deltaLocatedA52 := make([]byte, len(m.RetrieveSegments)*10)
		var j51 int
		for _, num := range m.RetrieveSegments {
			for num >= 1<<7 {
				deltaLocatedA52[j51] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j51++
			}
			deltaLocatedA52[j51] = uint8(num)
			j51++
		}
		i -= j51
		copy(deltaLocatedAN[i:], deltaLocatedA52[:j51])
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(j51))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Outcome != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Outcome))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyArrangeNomination) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyArrangeNomination) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyArrangeNomination) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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

func (m *ReplyHandleNomination) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyHandleNomination) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyHandleNomination) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Condition != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Condition))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyBroadenBallot) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyBroadenBallot) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyBroadenBallot) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.BallotAddition) > 0 {
		i -= len(m.BallotAddition)
		copy(deltaLocatedAN[i:], m.BallotAddition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.BallotAddition)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyValidateBallotAddition) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyValidateBallotAddition) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyValidateBallotAddition) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Condition != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Condition))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyCulminateLedger) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyCulminateLedger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyCulminateLedger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.PlatformDigest) > 0 {
		i -= len(m.PlatformDigest)
		copy(deltaLocatedAN[i:], m.PlatformDigest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.PlatformDigest)))
		i--
		deltaLocatedAN[i] = 0x2a
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
		deltaLocatedAN[i] = 0x22
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
			deltaLocatedAN[i] = 0x1a
		}
	}
	if len(m.TransferOutcomes) > 0 {
		for idxNdExc := len(m.TransferOutcomes) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.TransferOutcomes[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x12
		}
	}
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

func (m *EndorseDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *EndorseDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *EndorseDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Ballots) > 0 {
		for idxNdExc := len(m.Ballots) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Ballots[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x12
		}
	}
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ExpandedEndorseDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ExpandedEndorseDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ExpandedEndorseDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Ballots) > 0 {
		for idxNdExc := len(m.Ballots) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Ballots[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x12
		}
	}
	if m.Iteration != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Incident) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Incident) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Incident) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Properties) > 0 {
		for idxNdExc := len(m.Properties) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Properties[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
			}
			i--
			deltaLocatedAN[i] = 0x12
		}
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(deltaLocatedAN[i:], m.Kind)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Kind)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *IncidentProperty) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *IncidentProperty) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *IncidentProperty) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Ordinal {
		i--
		if m.Ordinal {
			deltaLocatedAN[i] = 1
		} else {
			deltaLocatedAN[i] = 0
		}
		i--
		deltaLocatedAN[i] = 0x18
	}
	if len(m.Datum) > 0 {
		i -= len(m.Datum)
		copy(deltaLocatedAN[i:], m.Datum)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Datum)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(deltaLocatedAN[i:], m.Key)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Key)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *InvokeTransferOutcome) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *InvokeTransferOutcome) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *InvokeTransferOutcome) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Codeset) > 0 {
		i -= len(m.Codeset)
		copy(deltaLocatedAN[i:], m.Codeset)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Codeset)))
		i--
		deltaLocatedAN[i] = 0x42
	}
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
			deltaLocatedAN[i] = 0x3a
		}
	}
	if m.FuelUtilized != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FuelUtilized))
		i--
		deltaLocatedAN[i] = 0x30
	}
	if m.FuelDesired != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.FuelDesired))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(deltaLocatedAN[i:], m.Details)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Details)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(deltaLocatedAN[i:], m.Log)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Log)))
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
	if m.Cipher != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Cipher))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *TransferOutcome) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *TransferOutcome) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *TransferOutcome) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	{
		extent, err := m.Outcome.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x22
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(deltaLocatedAN[i:], m.Tx)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Tx)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.Ordinal != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ordinal))
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

func (m *Assessor) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Assessor) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Assessor) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Potency != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Potency))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(deltaLocatedAN[i:], m.Location)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Location)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *AssessorRevise) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AssessorRevise) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AssessorRevise) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Potency != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Potency))
		i--
		deltaLocatedAN[i] = 0x10
	}
	{
		extent, err := m.PublicToken.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *BallotDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *BallotDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *BallotDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.LedgerUuidMarker != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.LedgerUuidMarker))
		i--
		deltaLocatedAN[i] = 0x18
	}
	{
		extent, err := m.Assessor.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *ExpandedBallotDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ExpandedBallotDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ExpandedBallotDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.LedgerUuidMarker != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.LedgerUuidMarker))
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.AdditionNotation) > 0 {
		i -= len(m.AdditionNotation)
		copy(deltaLocatedAN[i:], m.AdditionNotation)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.AdditionNotation)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if len(m.BallotAddition) > 0 {
		i -= len(m.BallotAddition)
		copy(deltaLocatedAN[i:], m.BallotAddition)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.BallotAddition)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	{
		extent, err := m.Assessor.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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

func (m *Malpractice) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Malpractice) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Malpractice) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.SumBallotingPotency != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.SumBallotingPotency))
		i--
		deltaLocatedAN[i] = 0x28
	}
	n58, fault58 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault58 != nil {
		return 0, fault58
	}
	i -= n58
	i = serializeVariableintKinds(deltaLocatedAN, i, uint64(n58))
	i--
	deltaLocatedAN[i] = 0x22
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x18
	}
	{
		extent, err := m.Assessor.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x12
	if m.Kind != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Kind))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Image) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Image) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Image) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		i -= len(m.Attributes)
		copy(deltaLocatedAN[i:], m.Attributes)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Attributes)))
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Segments != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Segments))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Layout != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Layout))
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
func (m *Solicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Datum != nil {
		n += m.Datum.Extent()
	}
	return n
}

func (m *Solicit_Reverberate) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Reverberate != nil {
		l = m.Reverberate.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Purge) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Purge != nil {
		l = m.Purge.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Details) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Details != nil {
		l = m.Details.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Initiatechain) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitializeSuccession != nil {
		l = m.InitializeSuccession.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Inquire) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Inquire != nil {
		l = m.Inquire.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Inspecttrans) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InspectTransfer != nil {
		l = m.InspectTransfer.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Endorse) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Endorse != nil {
		l = m.Endorse.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Catalogimages) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollectionImages != nil {
		l = m.CollectionImages.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Extendimage) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExtendImage != nil {
		l = m.ExtendImage.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Loadimagefragment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FetchImageSegment != nil {
		l = m.FetchImageSegment.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Executeimagefragment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecuteImageSegment != nil {
		l = m.ExecuteImageSegment.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Prepareitem) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ArrangeNomination != nil {
		l = m.ArrangeNomination.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Executeitem) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HandleNomination != nil {
		l = m.HandleNomination.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Extendballot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BroadenBallot != nil {
		l = m.BroadenBallot.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Verifyballotaddition) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidateBallotAddition != nil {
		l = m.ValidateBallotAddition.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Finalizeledger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CulminateLedger != nil {
		l = m.CulminateLedger.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Appendtrans) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppendTransfer != nil {
		l = m.AppendTransfer.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Solicit_Harvesttrans) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HarvestTrans != nil {
		l = m.HarvestTrans.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *SolicitReverberate) Extent() (n int) {
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

func (m *SolicitPurge) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SolicitDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Edition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.LedgerEdition != 0 {
		n += 1 + sovKinds(uint64(m.LedgerEdition))
	}
	if m.Peer2peerEdition != 0 {
		n += 1 + sovKinds(uint64(m.Peer2peerEdition))
	}
	l = len(m.IfaceEdition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SolicitInitializeSuccession) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.SuccessionUuid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.AgreementSettings != nil {
		l = m.AgreementSettings.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if len(m.Assessors) > 0 {
		for _, e := range m.Assessors {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.ApplicationStatusOctets)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.PrimaryAltitude != 0 {
		n += 1 + sovKinds(uint64(m.PrimaryAltitude))
	}
	return n
}

func (m *SolicitInquire) Extent() (n int) {
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
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Validate {
		n += 2
	}
	return n
}

func (m *SolicitInspectTransfer) Extent() (n int) {
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

func (m *SolicitAppendTransfer) Extent() (n int) {
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

func (m *SolicitHarvestTrans) Extent() (n int) {
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

func (m *SolicitEndorse) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SolicitCollectionImages) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SolicitExtendImage) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Image != nil {
		l = m.Image.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.PlatformDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SolicitFetchImageSegment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Layout != 0 {
		n += 1 + sovKinds(uint64(m.Layout))
	}
	if m.Segment != 0 {
		n += 1 + sovKinds(uint64(m.Segment))
	}
	return n
}

func (m *SolicitExecuteImageSegment) Extent() (n int) {
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
	l = len(m.Originator)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SolicitArrangeNomination) Extent() (n int) {
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
	l = m.RegionalFinalEndorse.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FollowingAssessorsDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.NominatorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SolicitHandleNomination) Extent() (n int) {
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
	l = m.ItemizedFinalEndorse.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FollowingAssessorsDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.NominatorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SolicitBroadenBallot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = m.ItemizedFinalEndorse.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.FollowingAssessorsDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.NominatorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SolicitValidateBallotAddition) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AssessorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = len(m.BallotAddition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SolicitCulminateLedger) Extent() (n int) {
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
	l = m.ResolvedFinalEndorse.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if len(m.Malpractice) > 0 {
		for _, e := range m.Malpractice {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.FollowingAssessorsDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.NominatorLocation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Reply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Datum != nil {
		n += m.Datum.Extent()
	}
	return n
}

func (m *Reply_Exemption) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exemption != nil {
		l = m.Exemption.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Reverberate) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Reverberate != nil {
		l = m.Reverberate.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Purge) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Purge != nil {
		l = m.Purge.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Details) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Details != nil {
		l = m.Details.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Initiatechain) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitializeSuccession != nil {
		l = m.InitializeSuccession.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Inquire) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Inquire != nil {
		l = m.Inquire.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Inspecttrans) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InspectTransfer != nil {
		l = m.InspectTransfer.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Endorse) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Endorse != nil {
		l = m.Endorse.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Catalogimages) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollectionImages != nil {
		l = m.CollectionImages.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Extendimage) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExtendImage != nil {
		l = m.ExtendImage.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Loadimagefragment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FetchImageSegment != nil {
		l = m.FetchImageSegment.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Executeimagefragment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecuteImageSegment != nil {
		l = m.ExecuteImageSegment.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Prepareitem) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ArrangeNomination != nil {
		l = m.ArrangeNomination.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Executeitem) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HandleNomination != nil {
		l = m.HandleNomination.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Extendballot) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BroadenBallot != nil {
		l = m.BroadenBallot.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Verifyballotaddition) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidateBallotAddition != nil {
		l = m.ValidateBallotAddition.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Finalizeledger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CulminateLedger != nil {
		l = m.CulminateLedger.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Appendtrans) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppendTransfer != nil {
		l = m.AppendTransfer.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Reply_Harvesttrans) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HarvestTrans != nil {
		l = m.HarvestTrans.Extent()
		n += 2 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *ReplyExemption) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Failure)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyReverberate) Extent() (n int) {
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

func (m *ReplyPurge) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ReplyDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Edition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.PlatformEdition != 0 {
		n += 1 + sovKinds(uint64(m.PlatformEdition))
	}
	if m.FinalLedgerAltitude != 0 {
		n += 1 + sovKinds(uint64(m.FinalLedgerAltitude))
	}
	l = len(m.FinalLedgerPlatformDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyInitializeSuccession) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AgreementSettings != nil {
		l = m.AgreementSettings.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if len(m.Assessors) > 0 {
		for _, e := range m.Assessors {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.PlatformDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyInquire) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cipher != 0 {
		n += 1 + sovKinds(uint64(m.Cipher))
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
	l = len(m.Datum)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.AttestationActions != nil {
		l = m.AttestationActions.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = len(m.Codeset)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyInspectTransfer) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cipher != 0 {
		n += 1 + sovKinds(uint64(m.Cipher))
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
	if m.FuelUtilized != 0 {
		n += 1 + sovKinds(uint64(m.FuelUtilized))
	}
	if len(m.Incidents) > 0 {
		for _, e := range m.Incidents {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Codeset)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ReplyAppendTransfer) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cipher != 0 {
		n += 1 + sovKinds(uint64(m.Cipher))
	}
	return n
}

func (m *ReplyHarvestTrans) Extent() (n int) {
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

func (m *ReplyEndorse) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PreserveAltitude != 0 {
		n += 1 + sovKinds(uint64(m.PreserveAltitude))
	}
	return n
}

func (m *ReplyCatalogImages) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Images) > 0 {
		for _, e := range m.Images {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ReplyExtendImage) Extent() (n int) {
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

func (m *ReplyFetchImageSegment) Extent() (n int) {
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

func (m *ReplyExecuteImageSegment) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Outcome != 0 {
		n += 1 + sovKinds(uint64(m.Outcome))
	}
	if len(m.RetrieveSegments) > 0 {
		l = 0
		for _, e := range m.RetrieveSegments {
			l += sovKinds(uint64(e))
		}
		n += 1 + sovKinds(uint64(l)) + l
	}
	if len(m.DeclineOriginators) > 0 {
		for _, s := range m.DeclineOriginators {
			l = len(s)
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ReplyArrangeNomination) Extent() (n int) {
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

func (m *ReplyHandleNomination) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Condition != 0 {
		n += 1 + sovKinds(uint64(m.Condition))
	}
	return n
}

func (m *ReplyBroadenBallot) Extent() (n int) {
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

func (m *ReplyValidateBallotAddition) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Condition != 0 {
		n += 1 + sovKinds(uint64(m.Condition))
	}
	return n
}

func (m *ReplyCulminateLedger) Extent() (n int) {
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
	if len(m.TransferOutcomes) > 0 {
		for _, e := range m.TransferOutcomes {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
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
	l = len(m.PlatformDigest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *EndorseDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	if len(m.Ballots) > 0 {
		for _, e := range m.Ballots {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *ExpandedEndorseDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Iteration != 0 {
		n += 1 + sovKinds(uint64(m.Iteration))
	}
	if len(m.Ballots) > 0 {
		for _, e := range m.Ballots {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *Incident) Extent() (n int) {
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
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *IncidentProperty) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Datum)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Ordinal {
		n += 2
	}
	return n
}

func (m *InvokeTransferOutcome) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cipher != 0 {
		n += 1 + sovKinds(uint64(m.Cipher))
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
	if m.FuelUtilized != 0 {
		n += 1 + sovKinds(uint64(m.FuelUtilized))
	}
	if len(m.Incidents) > 0 {
		for _, e := range m.Incidents {
			l = e.Extent()
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	l = len(m.Codeset)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *TransferOutcome) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = m.Outcome.Extent()
	n += 1 + l + sovKinds(uint64(l))
	return n
}

func (m *Assessor) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Potency != 0 {
		n += 1 + sovKinds(uint64(m.Potency))
	}
	return n
}

func (m *AssessorRevise) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PublicToken.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.Potency != 0 {
		n += 1 + sovKinds(uint64(m.Potency))
	}
	return n
}

func (m *BallotDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Assessor.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.LedgerUuidMarker != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUuidMarker))
	}
	return n
}

func (m *ExpandedBallotDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Assessor.Extent()
	n += 1 + l + sovKinds(uint64(l))
	l = len(m.BallotAddition)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.AdditionNotation)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.LedgerUuidMarker != 0 {
		n += 1 + sovKinds(uint64(m.LedgerUuidMarker))
	}
	return n
}

func (m *Malpractice) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovKinds(uint64(m.Kind))
	}
	l = m.Assessor.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovKinds(uint64(l))
	if m.SumBallotingPotency != 0 {
		n += 1 + sovKinds(uint64(m.SumBallotingPotency))
	}
	return n
}

func (m *Image) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
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
	l = len(m.Attributes)
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
func (m *Solicit) Decode(deltaLocatedAN []byte) error {
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
			v := &SolicitReverberate{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Reverberate{v}
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
			v := &SolicitPurge{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Purge{v}
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
			v := &SolicitDetails{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Details{v}
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
			v := &SolicitInitializeSuccession{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Initiatechain{v}
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
			v := &SolicitInquire{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Inquire{v}
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
			v := &SolicitInspectTransfer{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Inspecttrans{v}
			idxNdExc = submitOrdinal
		case 11:
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
			v := &SolicitEndorse{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Endorse{v}
			idxNdExc = submitOrdinal
		case 12:
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
			v := &SolicitCollectionImages{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Catalogimages{v}
			idxNdExc = submitOrdinal
		case 13:
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
			v := &SolicitExtendImage{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Extendimage{v}
			idxNdExc = submitOrdinal
		case 14:
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
			v := &SolicitFetchImageSegment{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Loadimagefragment{v}
			idxNdExc = submitOrdinal
		case 15:
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
			v := &SolicitExecuteImageSegment{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Executeimagefragment{v}
			idxNdExc = submitOrdinal
		case 16:
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
			v := &SolicitArrangeNomination{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Prepareitem{v}
			idxNdExc = submitOrdinal
		case 17:
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
			v := &SolicitHandleNomination{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Executeitem{v}
			idxNdExc = submitOrdinal
		case 18:
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
			v := &SolicitBroadenBallot{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Extendballot{v}
			idxNdExc = submitOrdinal
		case 19:
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
			v := &SolicitValidateBallotAddition{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Verifyballotaddition{v}
			idxNdExc = submitOrdinal
		case 20:
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
			v := &SolicitCulminateLedger{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Finalizeledger{v}
			idxNdExc = submitOrdinal
		case 21:
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
			v := &SolicitAppendTransfer{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Appendtrans{v}
			idxNdExc = submitOrdinal
		case 22:
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
			v := &SolicitHarvestTrans{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Solicit_Harvesttrans{v}
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
func (m *SolicitReverberate) Decode(deltaLocatedAN []byte) error {
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
			m.Signal = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *SolicitPurge) Decode(deltaLocatedAN []byte) error {
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
func (m *SolicitDetails) Decode(deltaLocatedAN []byte) error {
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
			m.Edition = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.LedgerEdition = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.LedgerEdition |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Peer2peerEdition = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Peer2peerEdition |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
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
			m.IfaceEdition = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *SolicitInitializeSuccession) Decode(deltaLocatedAN []byte) error {
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Moment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.SuccessionUuid = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
			if m.AgreementSettings == nil {
				m.AgreementSettings = &kinds1.AgreementSettings{}
			}
			if err := m.AgreementSettings.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Assessors = append(m.Assessors, AssessorRevise{})
			if err := m.Assessors[len(m.Assessors)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
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
			m.ApplicationStatusOctets = append(m.ApplicationStatusOctets[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.ApplicationStatusOctets == nil {
				m.ApplicationStatusOctets = []byte{}
			}
			idxNdExc = submitOrdinal
		case 6:
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
func (m *SolicitInquire) Decode(deltaLocatedAN []byte) error {
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
			m.Data = append(m.Data[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
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
			m.Route = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
			m.Validate = bool(v != 0)
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
func (m *SolicitInspectTransfer) Decode(deltaLocatedAN []byte) error {
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
			m.Tx = append(m.Tx[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			idxNdExc = submitOrdinal
		case 2:
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
				m.Kind |= InspectTransferKind(b&0x7F) << relocate
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
func (m *SolicitAppendTransfer) Decode(deltaLocatedAN []byte) error {
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
			m.Tx = append(m.Tx[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Tx == nil {
				m.Tx = []byte{}
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
func (m *SolicitHarvestTrans) Decode(deltaLocatedAN []byte) error {
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
			m.MaximumOctets = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MaximumOctets |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.MaximumFuel = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MaximumFuel |= uint64(b&0x7F) << relocate
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
func (m *SolicitEndorse) Decode(deltaLocatedAN []byte) error {
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
func (m *SolicitCollectionImages) Decode(deltaLocatedAN []byte) error {
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
func (m *SolicitExtendImage) Decode(deltaLocatedAN []byte) error {
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
			if m.Image == nil {
				m.Image = &Image{}
			}
			if err := m.Image.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
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
			m.PlatformDigest = append(m.PlatformDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.PlatformDigest == nil {
				m.PlatformDigest = []byte{}
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
func (m *SolicitFetchImageSegment) Decode(deltaLocatedAN []byte) error {
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
				m.Altitude |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Layout |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Segment = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Segment |= uint32(b&0x7F) << relocate
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
func (m *SolicitExecuteImageSegment) Decode(deltaLocatedAN []byte) error {
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
			m.Segment = append(m.Segment[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Segment == nil {
				m.Segment = []byte{}
			}
			idxNdExc = submitOrdinal
		case 3:
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
			m.Originator = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *SolicitArrangeNomination) Decode(deltaLocatedAN []byte) error {
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
			m.MaximumTransferOctets = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MaximumTransferOctets |= int64(b&0x7F) << relocate
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
			m.Txs = append(m.Txs, make([]byte, submitOrdinal-idxNdExc))
			copy(m.Txs[len(m.Txs)-1], deltaLocatedAN[idxNdExc:submitOrdinal])
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
			if err := m.RegionalFinalEndorse.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 5:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Moment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.FollowingAssessorsDigest = append(m.FollowingAssessorsDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FollowingAssessorsDigest == nil {
				m.FollowingAssessorsDigest = []byte{}
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
func (m *SolicitHandleNomination) Decode(deltaLocatedAN []byte) error {
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
			if err := m.ItemizedFinalEndorse.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Digest = append(m.Digest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 5:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Moment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.FollowingAssessorsDigest = append(m.FollowingAssessorsDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FollowingAssessorsDigest == nil {
				m.FollowingAssessorsDigest = []byte{}
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
func (m *SolicitBroadenBallot) Decode(deltaLocatedAN []byte) error {
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Moment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Txs = append(m.Txs, make([]byte, submitOrdinal-idxNdExc))
			copy(m.Txs[len(m.Txs)-1], deltaLocatedAN[idxNdExc:submitOrdinal])
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
			if err := m.ItemizedFinalEndorse.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.FollowingAssessorsDigest = append(m.FollowingAssessorsDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FollowingAssessorsDigest == nil {
				m.FollowingAssessorsDigest = []byte{}
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
func (m *SolicitValidateBallotAddition) Decode(deltaLocatedAN []byte) error {
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
			m.BallotAddition = append(m.BallotAddition[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.BallotAddition == nil {
				m.BallotAddition = []byte{}
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
func (m *SolicitCulminateLedger) Decode(deltaLocatedAN []byte) error {
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
			if err := m.ResolvedFinalEndorse.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Malpractice = append(m.Malpractice, Malpractice{})
			if err := m.Malpractice[len(m.Malpractice)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Digest = append(m.Digest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 5:
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Moment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.FollowingAssessorsDigest = append(m.FollowingAssessorsDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FollowingAssessorsDigest == nil {
				m.FollowingAssessorsDigest = []byte{}
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
func (m *Reply) Decode(deltaLocatedAN []byte) error {
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
			v := &ReplyExemption{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Exemption{v}
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
			v := &ReplyReverberate{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Reverberate{v}
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
			v := &ReplyPurge{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Purge{v}
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
			v := &ReplyDetails{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Details{v}
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
			v := &ReplyInitializeSuccession{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Initiatechain{v}
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
			v := &ReplyInquire{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Inquire{v}
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
			v := &ReplyInspectTransfer{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Inspecttrans{v}
			idxNdExc = submitOrdinal
		case 12:
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
			v := &ReplyEndorse{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Endorse{v}
			idxNdExc = submitOrdinal
		case 13:
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
			v := &ReplyCatalogImages{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Catalogimages{v}
			idxNdExc = submitOrdinal
		case 14:
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
			v := &ReplyExtendImage{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Extendimage{v}
			idxNdExc = submitOrdinal
		case 15:
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
			v := &ReplyFetchImageSegment{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Loadimagefragment{v}
			idxNdExc = submitOrdinal
		case 16:
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
			v := &ReplyExecuteImageSegment{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Executeimagefragment{v}
			idxNdExc = submitOrdinal
		case 17:
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
			v := &ReplyArrangeNomination{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Prepareitem{v}
			idxNdExc = submitOrdinal
		case 18:
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
			v := &ReplyHandleNomination{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Executeitem{v}
			idxNdExc = submitOrdinal
		case 19:
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
			v := &ReplyBroadenBallot{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Extendballot{v}
			idxNdExc = submitOrdinal
		case 20:
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
			v := &ReplyValidateBallotAddition{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Verifyballotaddition{v}
			idxNdExc = submitOrdinal
		case 21:
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
			v := &ReplyCulminateLedger{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Finalizeledger{v}
			idxNdExc = submitOrdinal
		case 22:
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
			v := &ReplyAppendTransfer{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Appendtrans{v}
			idxNdExc = submitOrdinal
		case 23:
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
			v := &ReplyHarvestTrans{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Datum = &Reply_Harvesttrans{v}
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
func (m *ReplyExemption) Decode(deltaLocatedAN []byte) error {
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
			m.Failure = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *ReplyReverberate) Decode(deltaLocatedAN []byte) error {
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
			m.Signal = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *ReplyPurge) Decode(deltaLocatedAN []byte) error {
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
func (m *ReplyDetails) Decode(deltaLocatedAN []byte) error {
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
			m.Data = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
			m.Edition = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PlatformEdition = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.PlatformEdition |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
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
			m.FinalLedgerPlatformDigest = append(m.FinalLedgerPlatformDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.FinalLedgerPlatformDigest == nil {
				m.FinalLedgerPlatformDigest = []byte{}
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
func (m *ReplyInitializeSuccession) Decode(deltaLocatedAN []byte) error {
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
			if m.AgreementSettings == nil {
				m.AgreementSettings = &kinds1.AgreementSettings{}
			}
			if err := m.AgreementSettings.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.Assessors = append(m.Assessors, AssessorRevise{})
			if err := m.Assessors[len(m.Assessors)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 3:
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
func (m *ReplyInquire) Decode(deltaLocatedAN []byte) error {
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
			m.Cipher = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Cipher |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
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
			m.Log = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 4:
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
			m.Details = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 5:
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
				m.Ordinal |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
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
			m.Key = append(m.Key[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Key == nil {
				m.Key = []byte{}
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
			m.Datum = append(m.Datum[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Datum == nil {
				m.Datum = []byte{}
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
			if m.AttestationActions == nil {
				m.AttestationActions = &security.AttestationActions{}
			}
			if err := m.AttestationActions.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 9:
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
		case 10:
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
			m.Codeset = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *ReplyInspectTransfer) Decode(deltaLocatedAN []byte) error {
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
			m.Cipher = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Cipher |= uint32(b&0x7F) << relocate
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
			m.Data = append(m.Data[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			idxNdExc = submitOrdinal
		case 3:
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
			m.Log = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 4:
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
			m.Details = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelDesired = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FuelDesired |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 6:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelUtilized = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FuelUtilized |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
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
			m.Incidents = append(m.Incidents, Incident{})
			if err := m.Incidents[len(m.Incidents)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 8:
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
			m.Codeset = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *ReplyAppendTransfer) Decode(deltaLocatedAN []byte) error {
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
			m.Cipher = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Cipher |= uint32(b&0x7F) << relocate
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
func (m *ReplyHarvestTrans) Decode(deltaLocatedAN []byte) error {
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
func (m *ReplyEndorse) Decode(deltaLocatedAN []byte) error {
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
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.PreserveAltitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.PreserveAltitude |= int64(b&0x7F) << relocate
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
func (m *ReplyCatalogImages) Decode(deltaLocatedAN []byte) error {
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
			m.Images = append(m.Images, &Image{})
			if err := m.Images[len(m.Images)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *ReplyExtendImage) Decode(deltaLocatedAN []byte) error {
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
			m.Outcome = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Outcome |= Replyextendimage_Outcome(b&0x7F) << relocate
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
func (m *ReplyFetchImageSegment) Decode(deltaLocatedAN []byte) error {
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
			m.Segment = append(m.Segment[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Segment == nil {
				m.Segment = []byte{}
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
func (m *ReplyExecuteImageSegment) Decode(deltaLocatedAN []byte) error {
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
			m.Outcome = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Outcome |= Replyapplyimagefragment_Outcome(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind == 0 {
				var v uint32
				for relocate := uint(0); ; relocate += 7 {
					if relocate >= 64 {
						return FaultIntegerOverrunKinds
					}
					if idxNdExc >= l {
						return io.ErrUnexpectedEOF
					}
					b := deltaLocatedAN[idxNdExc]
					idxNdExc++
					v |= uint32(b&0x7F) << relocate
					if b < 0x80 {
						break
					}
				}
				m.RetrieveSegments = append(m.RetrieveSegments, v)
			} else if cableKind == 2 {
				var compressedSize int
				for relocate := uint(0); ; relocate += 7 {
					if relocate >= 64 {
						return FaultIntegerOverrunKinds
					}
					if idxNdExc >= l {
						return io.ErrUnexpectedEOF
					}
					b := deltaLocatedAN[idxNdExc]
					idxNdExc++
					compressedSize |= int(b&0x7F) << relocate
					if b < 0x80 {
						break
					}
				}
				if compressedSize < 0 {
					return FaultUnfitMagnitudeKinds
				}
				submitOrdinal := idxNdExc + compressedSize
				if submitOrdinal < 0 {
					return FaultUnfitMagnitudeKinds
				}
				if submitOrdinal > l {
					return io.ErrUnexpectedEOF
				}
				var componentTotal int
				var tally int
				for _, number := range deltaLocatedAN[idxNdExc:submitOrdinal] {
					if number < 128 {
						tally++
					}
				}
				componentTotal = tally
				if componentTotal != 0 && len(m.RetrieveSegments) == 0 {
					m.RetrieveSegments = make([]uint32, 0, componentTotal)
				}
				for idxNdExc < submitOrdinal {
					var v uint32
					for relocate := uint(0); ; relocate += 7 {
						if relocate >= 64 {
							return FaultIntegerOverrunKinds
						}
						if idxNdExc >= l {
							return io.ErrUnexpectedEOF
						}
						b := deltaLocatedAN[idxNdExc]
						idxNdExc++
						v |= uint32(b&0x7F) << relocate
						if b < 0x80 {
							break
						}
					}
					m.RetrieveSegments = append(m.RetrieveSegments, v)
				}
			} else {
				return fmt.Errorf("REDACTED", cableKind)
			}
		case 3:
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
			m.DeclineOriginators = append(m.DeclineOriginators, string(deltaLocatedAN[idxNdExc:submitOrdinal]))
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
func (m *ReplyArrangeNomination) Decode(deltaLocatedAN []byte) error {
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
func (m *ReplyHandleNomination) Decode(deltaLocatedAN []byte) error {
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
			m.Condition = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Condition |= Responseexecuteitem_Itemstatus(b&0x7F) << relocate
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
func (m *ReplyBroadenBallot) Decode(deltaLocatedAN []byte) error {
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
			m.BallotAddition = append(m.BallotAddition[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.BallotAddition == nil {
				m.BallotAddition = []byte{}
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
func (m *ReplyValidateBallotAddition) Decode(deltaLocatedAN []byte) error {
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
			m.Condition = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Condition |= Responsecertifyballotaddition_Verifystatus(b&0x7F) << relocate
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
func (m *ReplyCulminateLedger) Decode(deltaLocatedAN []byte) error {
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
			m.Incidents = append(m.Incidents, Incident{})
			if err := m.Incidents[len(m.Incidents)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.TransferOutcomes = append(m.TransferOutcomes, &InvokeTransferOutcome{})
			if err := m.TransferOutcomes[len(m.TransferOutcomes)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			m.AssessorRevisions = append(m.AssessorRevisions, AssessorRevise{})
			if err := m.AssessorRevisions[len(m.AssessorRevisions)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.AgreementArgumentRevisions == nil {
				m.AgreementArgumentRevisions = &kinds1.AgreementSettings{}
			}
			if err := m.AgreementArgumentRevisions.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
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
			m.PlatformDigest = append(m.PlatformDigest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.PlatformDigest == nil {
				m.PlatformDigest = []byte{}
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
func (m *EndorseDetails) Decode(deltaLocatedAN []byte) error {
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
			m.Ballots = append(m.Ballots, BallotDetails{})
			if err := m.Ballots[len(m.Ballots)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *ExpandedEndorseDetails) Decode(deltaLocatedAN []byte) error {
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
			m.Ballots = append(m.Ballots, ExpandedBallotDetails{})
			if err := m.Ballots[len(m.Ballots)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *Incident) Decode(deltaLocatedAN []byte) error {
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
			m.Kind = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
			m.Properties = append(m.Properties, IncidentProperty{})
			if err := m.Properties[len(m.Properties)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *IncidentProperty) Decode(deltaLocatedAN []byte) error {
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
			m.Key = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
			m.Datum = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 3:
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
			m.Ordinal = bool(v != 0)
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
func (m *InvokeTransferOutcome) Decode(deltaLocatedAN []byte) error {
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
			m.Cipher = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Cipher |= uint32(b&0x7F) << relocate
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
			m.Data = append(m.Data[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			idxNdExc = submitOrdinal
		case 3:
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
			m.Log = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 4:
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
			m.Details = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelDesired = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FuelDesired |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 6:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.FuelUtilized = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.FuelUtilized |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
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
			m.Incidents = append(m.Incidents, Incident{})
			if err := m.Incidents[len(m.Incidents)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 8:
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
			m.Codeset = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *TransferOutcome) Decode(deltaLocatedAN []byte) error {
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
		case 3:
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
			m.Tx = append(m.Tx[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Tx == nil {
				m.Tx = []byte{}
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
			if err := m.Outcome.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *Assessor) Decode(deltaLocatedAN []byte) error {
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
			m.Location = append(m.Location[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Location == nil {
				m.Location = []byte{}
			}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Potency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Potency |= int64(b&0x7F) << relocate
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
func (m *AssessorRevise) Decode(deltaLocatedAN []byte) error {
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
			if err := m.PublicToken.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Potency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Potency |= int64(b&0x7F) << relocate
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
func (m *BallotDetails) Decode(deltaLocatedAN []byte) error {
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
			if err := m.Assessor.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 3:
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
				m.LedgerUuidMarker |= kinds1.LedgerUUIDMarker(b&0x7F) << relocate
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
func (m *ExpandedBallotDetails) Decode(deltaLocatedAN []byte) error {
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
			if err := m.Assessor.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 3:
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
			m.BallotAddition = append(m.BallotAddition[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.BallotAddition == nil {
				m.BallotAddition = []byte{}
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
			m.AdditionNotation = append(m.AdditionNotation[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.AdditionNotation == nil {
				m.AdditionNotation = []byte{}
			}
			idxNdExc = submitOrdinal
		case 5:
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
				m.LedgerUuidMarker |= kinds1.LedgerUUIDMarker(b&0x7F) << relocate
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
func (m *Malpractice) Decode(deltaLocatedAN []byte) error {
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
				m.Kind |= MalpracticeKind(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
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
			if err := m.Assessor.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
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
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.SumBallotingPotency = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
func (m *Image) Decode(deltaLocatedAN []byte) error {
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
				m.Altitude |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Layout |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Segments = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Segments |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
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
			m.Digest = append(m.Digest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
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
			m.Attributes = append(m.Attributes[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Attributes == nil {
				m.Attributes = []byte{}
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
