//
//

package privatekey

import (
	fmt "fmt"
	vault "github.com/valkyrieworks/schema/consensuscore/vault"
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

type Faults int32

const (
	Faults_FAULTS_UNCLEAR             Faults = 0
	Faults_FAULTS_UNFORESEEN_ANSWER Faults = 1
	Faults_FAULTS_NO_LINKAGE       Faults = 2
	Faults_FAULTS_LINKAGE_DEADLINE  Faults = 3
	Faults_FAULTS_FETCH_DEADLINE        Faults = 4
	Faults_FAULTS_RECORD_DEADLINE       Faults = 5
)

var Faults_label = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
	4: "REDACTED",
	5: "REDACTED",
}

var Faults_item = map[string]int32{
	"REDACTED":             0,
	"REDACTED": 1,
	"REDACTED":       2,
	"REDACTED":  3,
	"REDACTED":        4,
	"REDACTED":       5,
}

func (x Faults) String() string {
	return proto.EnumName(Faults_label, int32(x))
}

func (Faults) EnumerationDefinition() ([]byte, []int) {
	return filedefinition_hash10, []int{0}
}

type DistantNotaryFault struct {
	Code        int32  `protobuf:"variableint,1,opt,name=code,proto3" json:"code,omitempty"`
	Summary string `protobuf:"octets,2,opt,name=description,proto3" json:"summary,omitempty"`
}

func (m *DistantNotaryFault) Restore()         { *m = DistantNotaryFault{} }
func (m *DistantNotaryFault) String() string { return proto.CompactTextString(m) }
func (*DistantNotaryFault) SchemaSignal()    {}
func (*DistantNotaryFault) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{0}
}
func (m *DistantNotaryFault) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *DistantNotaryFault) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Remotenotaryfault.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistantNotaryFault) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Remotenotaryfault.Merge(m, src)
}
func (m *DistantNotaryFault) XXX_Volume() int {
	return m.Volume()
}
func (m *DistantNotaryFault) XXX_Omitunclear() {
	xxx_messagedata_Remotenotaryfault.DiscardUnknown(m)
}

var xxx_messagedata_Remotenotaryfault proto.InternalMessageInfo

func (m *DistantNotaryFault) FetchCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DistantNotaryFault) FetchSummary() string {
	if m != nil {
		return m.Summary
	}
	return "REDACTED"
}

//
type PublicKeyQuery struct {
	SeriesUid string `protobuf:"octets,1,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
}

func (m *PublicKeyQuery) Restore()         { *m = PublicKeyQuery{} }
func (m *PublicKeyQuery) String() string { return proto.CompactTextString(m) }
func (*PublicKeyQuery) SchemaSignal()    {}
func (*PublicKeyQuery) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{1}
}
func (m *PublicKeyQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PublicKeyQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Publickeyquery.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKeyQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Publickeyquery.Merge(m, src)
}
func (m *PublicKeyQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *PublicKeyQuery) XXX_Omitunclear() {
	xxx_messagedata_Publickeyquery.DiscardUnknown(m)
}

var xxx_messagedata_Publickeyquery proto.InternalMessageInfo

func (m *PublicKeyQuery) FetchSeriesUid() string {
	if m != nil {
		return m.SeriesUid
	}
	return "REDACTED"
}

//
type PublicKeyAnswer struct {
	PublicKey vault.PublicKey   `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_key"`
	Fault  *DistantNotaryFault `protobuf:"octets,2,opt,name=error,proto3" json:"fault,omitempty"`
}

func (m *PublicKeyAnswer) Restore()         { *m = PublicKeyAnswer{} }
func (m *PublicKeyAnswer) String() string { return proto.CompactTextString(m) }
func (*PublicKeyAnswer) SchemaSignal()    {}
func (*PublicKeyAnswer) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{2}
}
func (m *PublicKeyAnswer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PublicKeyAnswer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Publickeyoutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKeyAnswer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Publickeyoutcome.Merge(m, src)
}
func (m *PublicKeyAnswer) XXX_Volume() int {
	return m.Volume()
}
func (m *PublicKeyAnswer) XXX_Omitunclear() {
	xxx_messagedata_Publickeyoutcome.DiscardUnknown(m)
}

var xxx_messagedata_Publickeyoutcome proto.InternalMessageInfo

func (m *PublicKeyAnswer) FetchPublicKey() vault.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return vault.PublicKey{}
}

func (m *PublicKeyAnswer) FetchFault() *DistantNotaryFault {
	if m != nil {
		return m.Fault
	}
	return nil
}

//
type AttestBallotQuery struct {
	Ballot    *kinds.Ballot `protobuf:"octets,1,opt,name=vote,proto3" json:"ballot,omitempty"`
	SeriesUid string      `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
}

func (m *AttestBallotQuery) Restore()         { *m = AttestBallotQuery{} }
func (m *AttestBallotQuery) String() string { return proto.CompactTextString(m) }
func (*AttestBallotQuery) SchemaSignal()    {}
func (*AttestBallotQuery) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{3}
}
func (m *AttestBallotQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AttestBallotQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Attestballotquery.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestBallotQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Attestballotquery.Merge(m, src)
}
func (m *AttestBallotQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *AttestBallotQuery) XXX_Omitunclear() {
	xxx_messagedata_Attestballotquery.DiscardUnknown(m)
}

var xxx_messagedata_Attestballotquery proto.InternalMessageInfo

func (m *AttestBallotQuery) FetchBallot() *kinds.Ballot {
	if m != nil {
		return m.Ballot
	}
	return nil
}

func (m *AttestBallotQuery) FetchSeriesUid() string {
	if m != nil {
		return m.SeriesUid
	}
	return "REDACTED"
}

//
type AttestedBallotAnswer struct {
	Ballot  kinds.Ballot         `protobuf:"octets,1,opt,name=vote,proto3" json:"ballot"`
	Fault *DistantNotaryFault `protobuf:"octets,2,opt,name=error,proto3" json:"fault,omitempty"`
}

func (m *AttestedBallotAnswer) Restore()         { *m = AttestedBallotAnswer{} }
func (m *AttestedBallotAnswer) String() string { return proto.CompactTextString(m) }
func (*AttestedBallotAnswer) SchemaSignal()    {}
func (*AttestedBallotAnswer) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{4}
}
func (m *AttestedBallotAnswer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AttestedBallotAnswer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Attestedballotoutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestedBallotAnswer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Attestedballotoutcome.Merge(m, src)
}
func (m *AttestedBallotAnswer) XXX_Volume() int {
	return m.Volume()
}
func (m *AttestedBallotAnswer) XXX_Omitunclear() {
	xxx_messagedata_Attestedballotoutcome.DiscardUnknown(m)
}

var xxx_messagedata_Attestedballotoutcome proto.InternalMessageInfo

func (m *AttestedBallotAnswer) FetchBallot() kinds.Ballot {
	if m != nil {
		return m.Ballot
	}
	return kinds.Ballot{}
}

func (m *AttestedBallotAnswer) FetchFault() *DistantNotaryFault {
	if m != nil {
		return m.Fault
	}
	return nil
}

//
type AttestNominationQuery struct {
	Nomination *kinds.Nomination `protobuf:"octets,1,opt,name=proposal,proto3" json:"nomination,omitempty"`
	SeriesUid  string          `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"series_uid,omitempty"`
}

func (m *AttestNominationQuery) Restore()         { *m = AttestNominationQuery{} }
func (m *AttestNominationQuery) String() string { return proto.CompactTextString(m) }
func (*AttestNominationQuery) SchemaSignal()    {}
func (*AttestNominationQuery) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{5}
}
func (m *AttestNominationQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AttestNominationQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Attestproposalquery.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestNominationQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Attestproposalquery.Merge(m, src)
}
func (m *AttestNominationQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *AttestNominationQuery) XXX_Omitunclear() {
	xxx_messagedata_Attestproposalquery.DiscardUnknown(m)
}

var xxx_messagedata_Attestproposalquery proto.InternalMessageInfo

func (m *AttestNominationQuery) FetchNomination() *kinds.Nomination {
	if m != nil {
		return m.Nomination
	}
	return nil
}

func (m *AttestNominationQuery) FetchSeriesUid() string {
	if m != nil {
		return m.SeriesUid
	}
	return "REDACTED"
}

//
type AttestedNominationAnswer struct {
	Nomination kinds.Nomination     `protobuf:"octets,1,opt,name=proposal,proto3" json:"nomination"`
	Fault    *DistantNotaryFault `protobuf:"octets,2,opt,name=error,proto3" json:"fault,omitempty"`
}

func (m *AttestedNominationAnswer) Restore()         { *m = AttestedNominationAnswer{} }
func (m *AttestedNominationAnswer) String() string { return proto.CompactTextString(m) }
func (*AttestedNominationAnswer) SchemaSignal()    {}
func (*AttestedNominationAnswer) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{6}
}
func (m *AttestedNominationAnswer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AttestedNominationAnswer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Attestedproposaloutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestedNominationAnswer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Attestedproposaloutcome.Merge(m, src)
}
func (m *AttestedNominationAnswer) XXX_Volume() int {
	return m.Volume()
}
func (m *AttestedNominationAnswer) XXX_Omitunclear() {
	xxx_messagedata_Attestedproposaloutcome.DiscardUnknown(m)
}

var xxx_messagedata_Attestedproposaloutcome proto.InternalMessageInfo

func (m *AttestedNominationAnswer) FetchNomination() kinds.Nomination {
	if m != nil {
		return m.Nomination
	}
	return kinds.Nomination{}
}

func (m *AttestedNominationAnswer) FetchFault() *DistantNotaryFault {
	if m != nil {
		return m.Fault
	}
	return nil
}

//
type PingQuery struct {
}

func (m *PingQuery) Restore()         { *m = PingQuery{} }
func (m *PingQuery) String() string { return proto.CompactTextString(m) }
func (*PingQuery) SchemaSignal()    {}
func (*PingQuery) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{7}
}
func (m *PingQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PingQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Pingquery.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Pingquery.Merge(m, src)
}
func (m *PingQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *PingQuery) XXX_Omitunclear() {
	xxx_messagedata_Pingquery.DiscardUnknown(m)
}

var xxx_messagedata_Pingquery proto.InternalMessageInfo

//
type PingAnswer struct {
}

func (m *PingAnswer) Restore()         { *m = PingAnswer{} }
func (m *PingAnswer) String() string { return proto.CompactTextString(m) }
func (*PingAnswer) SchemaSignal()    {}
func (*PingAnswer) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{8}
}
func (m *PingAnswer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PingAnswer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Pingoutcome.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingAnswer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Pingoutcome.Merge(m, src)
}
func (m *PingAnswer) XXX_Volume() int {
	return m.Volume()
}
func (m *PingAnswer) XXX_Omitunclear() {
	xxx_messagedata_Pingoutcome.DiscardUnknown(m)
}

var xxx_messagedata_Pingoutcome proto.InternalMessageInfo

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
	Sum ismessage_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) SchemaSignal()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedefinition_hash10, []int{9}
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

type Signal_Publickeyquery struct {
	PublicKeyQuery *PublicKeyQuery `protobuf:"octets,1,opt,name=pub_key_request,json=pubKeyRequest,proto3,oneof" json:"public_key_query,omitempty"`
}
type Signal_Publickeyoutcome struct {
	PublicKeyAnswer *PublicKeyAnswer `protobuf:"octets,2,opt,name=pub_key_response,json=pubKeyResponse,proto3,oneof" json:"public_key_reply,omitempty"`
}
type Signal_Attestballotquery struct {
	AttestBallotQuery *AttestBallotQuery `protobuf:"octets,3,opt,name=sign_vote_request,json=signVoteRequest,proto3,oneof" json:"attest_ballot_query,omitempty"`
}
type Signal_Attestedballotoutcome struct {
	AttestedBallotAnswer *AttestedBallotAnswer `protobuf:"octets,4,opt,name=signed_vote_response,json=signedVoteResponse,proto3,oneof" json:"attested_ballot_reply,omitempty"`
}
type Signal_Attestproposalquery struct {
	AttestNominationQuery *AttestNominationQuery `protobuf:"octets,5,opt,name=sign_proposal_request,json=signProposalRequest,proto3,oneof" json:"attest_nomination_query,omitempty"`
}
type Signal_Attestedproposaloutcome struct {
	AttestedNominationAnswer *AttestedNominationAnswer `protobuf:"octets,6,opt,name=signed_proposal_response,json=signedProposalResponse,proto3,oneof" json:"attested_nomination_reply,omitempty"`
}
type Signal_Pingquery struct {
	PingQuery *PingQuery `protobuf:"octets,7,opt,name=ping_request,json=pingRequest,proto3,oneof" json:"ping_query,omitempty"`
}
type Signal_Pingoutcome struct {
	PingAnswer *PingAnswer `protobuf:"octets,8,opt,name=ping_response,json=pingResponse,proto3,oneof" json:"ping_reply,omitempty"`
}

func (*Signal_Publickeyquery) ismessage_Total()          {}
func (*Signal_Publickeyoutcome) ismessage_Total()         {}
func (*Signal_Attestballotquery) ismessage_Total()        {}
func (*Signal_Attestedballotoutcome) ismessage_Total()     {}
func (*Signal_Attestproposalquery) ismessage_Total()    {}
func (*Signal_Attestedproposaloutcome) ismessage_Total() {}
func (*Signal_Pingquery) ismessage_Total()            {}
func (*Signal_Pingoutcome) ismessage_Total()           {}

func (m *Signal) FetchTotal() ismessage_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) FetchPublicKeyQuery() *PublicKeyQuery {
	if x, ok := m.FetchTotal().(*Signal_Publickeyquery); ok {
		return x.PublicKeyQuery
	}
	return nil
}

func (m *Signal) FetchPublicKeyAnswer() *PublicKeyAnswer {
	if x, ok := m.FetchTotal().(*Signal_Publickeyoutcome); ok {
		return x.PublicKeyAnswer
	}
	return nil
}

func (m *Signal) FetchAttestBallotQuery() *AttestBallotQuery {
	if x, ok := m.FetchTotal().(*Signal_Attestballotquery); ok {
		return x.AttestBallotQuery
	}
	return nil
}

func (m *Signal) FetchAttestedBallotAnswer() *AttestedBallotAnswer {
	if x, ok := m.FetchTotal().(*Signal_Attestedballotoutcome); ok {
		return x.AttestedBallotAnswer
	}
	return nil
}

func (m *Signal) FetchAttestNominationQuery() *AttestNominationQuery {
	if x, ok := m.FetchTotal().(*Signal_Attestproposalquery); ok {
		return x.AttestNominationQuery
	}
	return nil
}

func (m *Signal) FetchAttestedNominationAnswer() *AttestedNominationAnswer {
	if x, ok := m.FetchTotal().(*Signal_Attestedproposaloutcome); ok {
		return x.AttestedNominationAnswer
	}
	return nil
}

func (m *Signal) FetchPingQuery() *PingQuery {
	if x, ok := m.FetchTotal().(*Signal_Pingquery); ok {
		return x.PingQuery
	}
	return nil
}

func (m *Signal) FetchPingAnswer() *PingAnswer {
	if x, ok := m.FetchTotal().(*Signal_Pingoutcome); ok {
		return x.PingAnswer
	}
	return nil
}

//
func (*Signal) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Signal_Publickeyquery)(nil),
		(*Signal_Publickeyoutcome)(nil),
		(*Signal_Attestballotquery)(nil),
		(*Signal_Attestedballotoutcome)(nil),
		(*Signal_Attestproposalquery)(nil),
		(*Signal_Attestedproposaloutcome)(nil),
		(*Signal_Pingquery)(nil),
		(*Signal_Pingoutcome)(nil),
	}
}

func init() {
	proto.RegisterEnum("REDACTED", Faults_label, Faults_item)
	proto.RegisterType((*DistantNotaryFault)(nil), "REDACTED")
	proto.RegisterType((*PublicKeyQuery)(nil), "REDACTED")
	proto.RegisterType((*PublicKeyAnswer)(nil), "REDACTED")
	proto.RegisterType((*AttestBallotQuery)(nil), "REDACTED")
	proto.RegisterType((*AttestedBallotAnswer)(nil), "REDACTED")
	proto.RegisterType((*AttestNominationQuery)(nil), "REDACTED")
	proto.RegisterType((*AttestedNominationAnswer)(nil), "REDACTED")
	proto.RegisterType((*PingQuery)(nil), "REDACTED")
	proto.RegisterType((*PingAnswer)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash10) }

var filedefinition_hash10 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x4f, 0xe3, 0x46,
	0x18, 0xb6, 0x21, 0x1f, 0xf0, 0x86, 0x84, 0x30, 0x50, 0x1a, 0x22, 0x6a, 0xd2, 0x54, 0x6d, 0x51,
	0x0e, 0x49, 0x45, 0xd5, 0x5e, 0xe8, 0xa5, 0x80, 0xd5, 0x44, 0x11, 0x76, 0x3a, 0x09, 0x05, 0x21,
	0x55, 0x56, 0x3e, 0x06, 0x63, 0x41, 0x3c, 0x5e, 0x8f, 0x83, 0x94, 0xf3, 0xde, 0xf6, 0xb4, 0xd2,
	0xfe, 0x89, 0x3d, 0xef, 0xaf, 0xe0, 0xc8, 0x71, 0x4f, 0xab, 0x15, 0xfc, 0x91, 0x55, 0xc6, 0x13,
	0xdb, 0xf9, 0x42, 0xbb, 0xe2, 0x36, 0xf3, 0xbe, 0xef, 0x3c, 0x1f, 0x33, 0x8f, 0x65, 0x50, 0x3c,
	0x62, 0xf7, 0x88, 0xdb, 0xb7, 0x6c, 0xaf, 0xe2, 0xb8, 0xd6, 0xdd, 0x5d, 0xfb, 0xb6, 0xe2, 0x0d,
	0x1d, 0xc2, 0xca, 0x8e, 0x4b, 0x3d, 0x8a, 0x50, 0xd8, 0x2f, 0x8b, 0x7e, 0x7e, 0x37, 0x72, 0xa6,
	0xeb, 0x0e, 0x1d, 0x8f, 0x56, 0x6e, 0xc8, 0x50, 0x9c, 0x98, 0xe8, 0x72, 0xa4, 0x28, 0x5e, 0x7e,
	0xcb, 0xa4, 0x26, 0xe5, 0xcb, 0xca, 0x68, 0xe5, 0x57, 0x8b, 0x35, 0xd8, 0xc0, 0xa4, 0x4f, 0x3d,
	0xd2, 0xb4, 0x4c, 0x9b, 0xb8, 0xaa, 0xeb, 0x52, 0x17, 0x21, 0x88, 0x75, 0x69, 0x8f, 0xe4, 0xe4,
	0x82, 0xbc, 0x1f, 0xc7, 0x7c, 0x8d, 0x0a, 0x90, 0xea, 0x11, 0xd6, 0x75, 0x2d, 0xc7, 0xb3, 0xa8,
	0x9d, 0x5b, 0x2a, 0xc8, 0xfb, 0xab, 0x38, 0x5a, 0x2a, 0x96, 0x20, 0xdd, 0x18, 0x74, 0xea, 0x64,
	0x88, 0xc9, 0xab, 0x01, 0x61, 0x1e, 0xda, 0x81, 0x95, 0xee, 0x75, 0xdb, 0xb2, 0x0d, 0xab, 0xc7,
	0xa1, 0x56, 0x71, 0x92, 0xef, 0x6b, 0xbd, 0xe2, 0x1b, 0x19, 0x32, 0xe3, 0x61, 0xe6, 0x50, 0x9b,
	0x11, 0x74, 0x08, 0x49, 0x67, 0xd0, 0x31, 0x6e, 0xc8, 0x90, 0x0f, 0xa7, 0x0e, 0x76, 0xcb, 0x91,
	0x1b, 0xf0, 0xdd, 0x96, 0x1b, 0x83, 0xce, 0xad, 0xd5, 0xad, 0x93, 0xe1, 0x51, 0xec, 0xfe, 0xd3,
	0x9e, 0x84, 0x13, 0x0e, 0x07, 0x41, 0x87, 0x10, 0x27, 0x23, 0xe9, 0x5c, 0x57, 0xea, 0xe0, 0xe7,
	0xf2, 0xec, 0xe5, 0x95, 0x67, 0x7c, 0x62, 0xff, 0x4c, 0xf1, 0x02, 0xd6, 0x47, 0xd5, 0xff, 0xa8,
	0x47, 0xc6, 0xd2, 0x4b, 0x10, 0xbb, 0xa3, 0x1e, 0x11, 0x4a, 0xb6, 0xa3, 0x70, 0xfe, 0x9d, 0xf2,
	0x61, 0x3e, 0x33, 0x61, 0x73, 0x69, 0xd2, 0xe6, 0x6b, 0x19, 0x10, 0x27, 0xec, 0xf9, 0xe0, 0xc2,
	0xea, 0x6f, 0x5f, 0x83, 0x2e, 0x1c, 0xfa, 0x1c, 0x2f, 0xf2, 0x77, 0x0d, 0x9b, 0xa3, 0x6a, 0xc3,
	0xa5, 0x0e, 0x65, 0xed, 0xdb, 0xb1, 0xc7, 0x3f, 0x61, 0xc5, 0x11, 0x25, 0xa1, 0x24, 0x3f, 0xab,
	0x24, 0x38, 0x14, 0xcc, 0x3e, 0xe7, 0xf7, 0x9d, 0x0c, 0xdb, 0xbe, 0xdf, 0x90, 0x4c, 0x78, 0xfe,
	0xeb, 0x5b, 0xd8, 0x84, 0xf7, 0x90, 0xf3, 0x45, 0xfe, 0xd3, 0x90, 0x6a, 0x58, 0xb6, 0x29, 0x7c,
	0x17, 0x33, 0xb0, 0xe6, 0x6f, 0x7d, 0x65, 0xc5, 0x0f, 0x71, 0x48, 0x9e, 0x12, 0xc6, 0xda, 0x26,
	0x41, 0x75, 0x58, 0x17, 0x21, 0x34, 0x5c, 0x7f, 0x5c, 0x88, 0xfd, 0x71, 0x1e, 0xe3, 0x44, 0xdc,
	0xab, 0x12, 0x4e, 0x3b, 0x13, 0xf9, 0xd7, 0x20, 0x1b, 0x82, 0xf9, 0x64, 0x42, 0x7f, 0xf1, 0x39,
	0x34, 0x7f, 0xb2, 0x2a, 0xe1, 0x8c, 0x33, 0xf9, 0x85, 0xfc, 0x0b, 0x1b, 0xcc, 0x32, 0x6d, 0x63,
	0x94, 0x88, 0x40, 0xde, 0x32, 0x07, 0xfc, 0x69, 0x1e, 0xe0, 0x54, 0xa8, 0xab, 0x12, 0x5e, 0x67,
	0x53, 0x39, 0xbf, 0x84, 0x2d, 0xc6, 0xdf, 0x6b, 0x0c, 0x2a, 0x64, 0xc6, 0x38, 0xea, 0x2f, 0x8b,
	0x50, 0x27, 0xf3, 0x5c, 0x95, 0x30, 0x62, 0xb3, 0x29, 0xff, 0x1f, 0xbe, 0xe3, 0x72, 0xc7, 0x8f,
	0x18, 0x48, 0x8e, 0x73, 0xf0, 0x5f, 0x17, 0x81, 0x4f, 0xe5, 0xb4, 0x2a, 0xe1, 0x4d, 0x36, 0x27,
	0xbe, 0x57, 0x90, 0x13, 0xd2, 0x23, 0x04, 0x42, 0x7e, 0x82, 0x33, 0x94, 0x16, 0xcb, 0x9f, 0x8e,
	0x67, 0x55, 0xc2, 0xdb, 0x6c, 0x7e, 0x70, 0x4f, 0x60, 0xcd, 0xb1, 0x6c, 0x33, 0x50, 0x9f, 0xe4,
	0xd8, 0x7b, 0x73, 0x5f, 0x30, 0x4c, 0x59, 0x55, 0xc2, 0x29, 0x27, 0xdc, 0xa2, 0x7f, 0x20, 0x2d,
	0x50, 0x84, 0xc4, 0x15, 0x0e, 0x53, 0x58, 0x0c, 0x13, 0x08, 0x5b, 0x73, 0x22, 0xfb, 0xa3, 0x38,
	0x2c, 0xb3, 0x41, 0xbf, 0xf4, 0x5e, 0x86, 0x04, 0x0f, 0x39, 0x43, 0x08, 0x32, 0x2a, 0xc6, 0x3a,
	0x6e, 0x1a, 0x67, 0x5a, 0x5d, 0xd3, 0xcf, 0xb5, 0xac, 0x84, 0x14, 0xc8, 0x07, 0x35, 0xf5, 0xa2,
	0xa1, 0x1e, 0xb7, 0xd4, 0x13, 0x03, 0xab, 0xcd, 0x86, 0xae, 0x35, 0xd5, 0xac, 0x8c, 0x72, 0xb0,
	0x25, 0xfa, 0x9a, 0x6e, 0x1c, 0xeb, 0x9a, 0xa6, 0x1e, 0xb7, 0x6a, 0xba, 0x96, 0x5d, 0x42, 0x3f,
	0xc0, 0x8e, 0xe8, 0x84, 0x65, 0xa3, 0x55, 0x3b, 0x55, 0xf5, 0xb3, 0x56, 0x76, 0x19, 0x7d, 0x0f,
	0x9b, 0xa2, 0x8d, 0xd5, 0xbf, 0x4f, 0x82, 0x46, 0x2c, 0x82, 0x78, 0x8e, 0x6b, 0x2d, 0x35, 0xe8,
	0xc4, 0x8f, 0xf4, 0xfb, 0x47, 0x45, 0x7e, 0x78, 0x54, 0xe4, 0xcf, 0x8f, 0x8a, 0xfc, 0xf6, 0x49,
	0x91, 0x1e, 0x9e, 0x14, 0xe9, 0xe3, 0x93, 0x22, 0x5d, 0xfe, 0x61, 0x5a, 0xde, 0xf5, 0xa0, 0x53,
	0xee, 0xd2, 0x7e, 0xa5, 0x4b, 0xfb, 0xc4, 0xeb, 0x5c, 0x79, 0xe1, 0xc2, 0xff, 0x57, 0xcd, 0xfe,
	0x25, 0x3b, 0x09, 0xde, 0xf9, 0xfd, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x9f, 0x99, 0x3e,
	0x42, 0x07, 0x00, 0x00,
}

func (m *DistantNotaryFault) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistantNotaryFault) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *DistantNotaryFault) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Summary) > 0 {
		i -= len(m.Summary)
		copy(dAtA[i:], m.Summary)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Summary)))
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

func (m *PublicKeyQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKeyQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PublicKeyQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SeriesUid) > 0 {
		i -= len(m.SeriesUid)
		copy(dAtA[i:], m.SeriesUid)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.SeriesUid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PublicKeyAnswer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKeyAnswer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PublicKeyAnswer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fault != nil {
		{
			volume, err := m.Fault.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
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

func (m *AttestBallotQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestBallotQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AttestBallotQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SeriesUid) > 0 {
		i -= len(m.SeriesUid)
		copy(dAtA[i:], m.SeriesUid)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.SeriesUid)))
		i--
		dAtA[i] = 0x12
	}
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

func (m *AttestedBallotAnswer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestedBallotAnswer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AttestedBallotAnswer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fault != nil {
		{
			volume, err := m.Fault.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
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
	return len(dAtA) - i, nil
}

func (m *AttestNominationQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestNominationQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AttestNominationQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SeriesUid) > 0 {
		i -= len(m.SeriesUid)
		copy(dAtA[i:], m.SeriesUid)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.SeriesUid)))
		i--
		dAtA[i] = 0x12
	}
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AttestedNominationAnswer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestedNominationAnswer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AttestedNominationAnswer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fault != nil {
		{
			volume, err := m.Fault.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
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

func (m *PingQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PingQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PingAnswer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingAnswer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PingAnswer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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

func (m *Signal_Publickeyquery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Publickeyquery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PublicKeyQuery != nil {
		{
			volume, err := m.PublicKeyQuery.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Publickeyoutcome) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Publickeyoutcome) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PublicKeyAnswer != nil {
		{
			volume, err := m.PublicKeyAnswer.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Attestballotquery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Attestballotquery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AttestBallotQuery != nil {
		{
			volume, err := m.AttestBallotQuery.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Attestedballotoutcome) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Attestedballotoutcome) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AttestedBallotAnswer != nil {
		{
			volume, err := m.AttestedBallotAnswer.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Attestproposalquery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Attestproposalquery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AttestNominationQuery != nil {
		{
			volume, err := m.AttestNominationQuery.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Attestedproposaloutcome) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Attestedproposaloutcome) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AttestedNominationAnswer != nil {
		{
			volume, err := m.AttestedNominationAnswer.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Pingquery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Pingquery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PingQuery != nil {
		{
			volume, err := m.PingQuery.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Pingoutcome) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Pingoutcome) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PingAnswer != nil {
		{
			volume, err := m.PingAnswer.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *DistantNotaryFault) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovKinds(uint64(m.Code))
	}
	l = len(m.Summary)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *PublicKeyQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SeriesUid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *PublicKeyAnswer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PublicKey.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.Fault != nil {
		l = m.Fault.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AttestBallotQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ballot != nil {
		l = m.Ballot.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.SeriesUid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AttestedBallotAnswer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Ballot.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.Fault != nil {
		l = m.Fault.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AttestNominationQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nomination != nil {
		l = m.Nomination.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.SeriesUid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AttestedNominationAnswer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Nomination.Volume()
	n += 1 + l + sovKinds(uint64(l))
	if m.Fault != nil {
		l = m.Fault.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *PingQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PingAnswer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
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

func (m *Signal_Publickeyquery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKeyQuery != nil {
		l = m.PublicKeyQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Publickeyoutcome) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKeyAnswer != nil {
		l = m.PublicKeyAnswer.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Attestballotquery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestBallotQuery != nil {
		l = m.AttestBallotQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Attestedballotoutcome) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestedBallotAnswer != nil {
		l = m.AttestedBallotAnswer.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Attestproposalquery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestNominationQuery != nil {
		l = m.AttestNominationQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Attestedproposaloutcome) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestedNominationAnswer != nil {
		l = m.AttestedNominationAnswer.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Pingquery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingQuery != nil {
		l = m.PingQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Pingoutcome) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingAnswer != nil {
		l = m.PingAnswer.Volume()
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
func (m *DistantNotaryFault) Unserialize(dAtA []byte) error {
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
				m.Code |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
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
			m.Summary = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *PublicKeyQuery) Unserialize(dAtA []byte) error {
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
			m.SeriesUid = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *PublicKeyAnswer) Unserialize(dAtA []byte) error {
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
			if m.Fault == nil {
				m.Fault = &DistantNotaryFault{}
			}
			if err := m.Fault.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *AttestBallotQuery) Unserialize(dAtA []byte) error {
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
func (m *AttestedBallotAnswer) Unserialize(dAtA []byte) error {
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
			if err := m.Ballot.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.Fault == nil {
				m.Fault = &DistantNotaryFault{}
			}
			if err := m.Fault.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *AttestNominationQuery) Unserialize(dAtA []byte) error {
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
			if m.Nomination == nil {
				m.Nomination = &kinds.Nomination{}
			}
			if err := m.Nomination.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *AttestedNominationAnswer) Unserialize(dAtA []byte) error {
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
			if m.Fault == nil {
				m.Fault = &DistantNotaryFault{}
			}
			if err := m.Fault.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *PingQuery) Unserialize(dAtA []byte) error {
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
func (m *PingAnswer) Unserialize(dAtA []byte) error {
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
			v := &PublicKeyQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Publickeyquery{v}
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
			v := &PublicKeyAnswer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Publickeyoutcome{v}
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
			v := &AttestBallotQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Attestballotquery{v}
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
			v := &AttestedBallotAnswer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Attestedballotoutcome{v}
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
			v := &AttestNominationQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Attestproposalquery{v}
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
			v := &AttestedNominationAnswer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Attestedproposaloutcome{v}
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
			v := &PingQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Pingquery{v}
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
			v := &PingAnswer{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Pingoutcome{v}
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
