//
//

package privatevalue

import (
	fmt "fmt"
	security "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
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

type Faults int32

const (
	Faults_FAULTS_UNFAMILIAR             Faults = 0
	Faults_FAULTS_UNFORESEEN_REPLY Faults = 1
	Faults_FAULTS_NEGATIVE_LINKAGE       Faults = 2
	Faults_FAULTS_LINKAGE_DEADLINE  Faults = 3
	Faults_FAULTS_FETCH_DEADLINE        Faults = 4
	Faults_FAULTS_RECORD_DEADLINE       Faults = 5
)

var Faults_alias = map[int32]string{
	0: "REDACTED",
	1: "REDACTED",
	2: "REDACTED",
	3: "REDACTED",
	4: "REDACTED",
	5: "REDACTED",
}

var Faults_datum = map[string]int32{
	"REDACTED":             0,
	"REDACTED": 1,
	"REDACTED":       2,
	"REDACTED":  3,
	"REDACTED":        4,
	"REDACTED":       5,
}

func (x Faults) Text() string {
	return proto.EnumName(Faults_alias, int32(x))
}

func (Faults) EnumerationDefinition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{0}
}

type RemoteEndorserFailure struct {
	Cipher        int32  `protobuf:"variableint,1,opt,name=code,proto3" json:"cipher,omitempty"`
	Characterization string `protobuf:"octets,2,opt,name=description,proto3" json:"definition,omitempty"`
}

func (m *RemoteEndorserFailure) Restore()         { *m = RemoteEndorserFailure{} }
func (m *RemoteEndorserFailure) Text() string { return proto.CompactTextString(m) }
func (*RemoteEndorserFailure) SchemaArtifact()    {}
func (*RemoteEndorserFailure) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{0}
}
func (m *RemoteEndorserFailure) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *RemoteEndorserFailure) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Remoteendorserfault.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoteEndorserFailure) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Remoteendorserfault.Merge(m, src)
}
func (m *RemoteEndorserFailure) XXX_Extent() int {
	return m.Extent()
}
func (m *RemoteEndorserFailure) XXX_Dropunfamiliar() {
	xxx_signaldetails_Remoteendorserfault.DiscardUnknown(m)
}

var xxx_signaldetails_Remoteendorserfault proto.InternalMessageInfo

func (m *RemoteEndorserFailure) ObtainCipher() int32 {
	if m != nil {
		return m.Cipher
	}
	return 0
}

func (m *RemoteEndorserFailure) ObtainCharacterization() string {
	if m != nil {
		return m.Characterization
	}
	return "REDACTED"
}

//
type PublicTokenSolicit struct {
	SuccessionUuid string `protobuf:"octets,1,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
}

func (m *PublicTokenSolicit) Restore()         { *m = PublicTokenSolicit{} }
func (m *PublicTokenSolicit) Text() string { return proto.CompactTextString(m) }
func (*PublicTokenSolicit) SchemaArtifact()    {}
func (*PublicTokenSolicit) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{1}
}
func (m *PublicTokenSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PublicTokenSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Publictokensolicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicTokenSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Publictokensolicit.Merge(m, src)
}
func (m *PublicTokenSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *PublicTokenSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Publictokensolicit.DiscardUnknown(m)
}

var xxx_signaldetails_Publictokensolicit proto.InternalMessageInfo

func (m *PublicTokenSolicit) ObtainSuccessionUuid() string {
	if m != nil {
		return m.SuccessionUuid
	}
	return "REDACTED"
}

//
type PublicTokenReply struct {
	PublicToken security.CommonToken   `protobuf:"octets,1,opt,name=pub_key,json=pubKey,proto3" json:"public_token"`
	Failure  *RemoteEndorserFailure `protobuf:"octets,2,opt,name=error,proto3" json:"failure,omitempty"`
}

func (m *PublicTokenReply) Restore()         { *m = PublicTokenReply{} }
func (m *PublicTokenReply) Text() string { return proto.CompactTextString(m) }
func (*PublicTokenReply) SchemaArtifact()    {}
func (*PublicTokenReply) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{2}
}
func (m *PublicTokenReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PublicTokenReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Publictokendata.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicTokenReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Publictokendata.Merge(m, src)
}
func (m *PublicTokenReply) XXX_Extent() int {
	return m.Extent()
}
func (m *PublicTokenReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Publictokendata.DiscardUnknown(m)
}

var xxx_signaldetails_Publictokendata proto.InternalMessageInfo

func (m *PublicTokenReply) ObtainPublicToken() security.CommonToken {
	if m != nil {
		return m.PublicToken
	}
	return security.CommonToken{}
}

func (m *PublicTokenReply) ObtainFailure() *RemoteEndorserFailure {
	if m != nil {
		return m.Failure
	}
	return nil
}

//
type AttestBallotSolicit struct {
	Ballot    *kinds.Ballot `protobuf:"octets,1,opt,name=vote,proto3" json:"ballot,omitempty"`
	SuccessionUuid string      `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
}

func (m *AttestBallotSolicit) Restore()         { *m = AttestBallotSolicit{} }
func (m *AttestBallotSolicit) Text() string { return proto.CompactTextString(m) }
func (*AttestBallotSolicit) SchemaArtifact()    {}
func (*AttestBallotSolicit) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{3}
}
func (m *AttestBallotSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AttestBallotSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Notateballotsolicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestBallotSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Notateballotsolicit.Merge(m, src)
}
func (m *AttestBallotSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *AttestBallotSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Notateballotsolicit.DiscardUnknown(m)
}

var xxx_signaldetails_Notateballotsolicit proto.InternalMessageInfo

func (m *AttestBallotSolicit) FetchBallot() *kinds.Ballot {
	if m != nil {
		return m.Ballot
	}
	return nil
}

func (m *AttestBallotSolicit) ObtainSuccessionUuid() string {
	if m != nil {
		return m.SuccessionUuid
	}
	return "REDACTED"
}

//
type NotatedBallotReply struct {
	Ballot  kinds.Ballot         `protobuf:"octets,1,opt,name=vote,proto3" json:"ballot"`
	Failure *RemoteEndorserFailure `protobuf:"octets,2,opt,name=error,proto3" json:"failure,omitempty"`
}

func (m *NotatedBallotReply) Restore()         { *m = NotatedBallotReply{} }
func (m *NotatedBallotReply) Text() string { return proto.CompactTextString(m) }
func (*NotatedBallotReply) SchemaArtifact()    {}
func (*NotatedBallotReply) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{4}
}
func (m *NotatedBallotReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *NotatedBallotReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Notatedballotreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotatedBallotReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Notatedballotreply.Merge(m, src)
}
func (m *NotatedBallotReply) XXX_Extent() int {
	return m.Extent()
}
func (m *NotatedBallotReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Notatedballotreply.DiscardUnknown(m)
}

var xxx_signaldetails_Notatedballotreply proto.InternalMessageInfo

func (m *NotatedBallotReply) FetchBallot() kinds.Ballot {
	if m != nil {
		return m.Ballot
	}
	return kinds.Ballot{}
}

func (m *NotatedBallotReply) ObtainFailure() *RemoteEndorserFailure {
	if m != nil {
		return m.Failure
	}
	return nil
}

//
type AttestNominationSolicit struct {
	Nomination *kinds.Nomination `protobuf:"octets,1,opt,name=proposal,proto3" json:"nomination,omitempty"`
	SuccessionUuid  string          `protobuf:"octets,2,opt,name=chain_id,json=chainId,proto3" json:"succession_uuid,omitempty"`
}

func (m *AttestNominationSolicit) Restore()         { *m = AttestNominationSolicit{} }
func (m *AttestNominationSolicit) Text() string { return proto.CompactTextString(m) }
func (*AttestNominationSolicit) SchemaArtifact()    {}
func (*AttestNominationSolicit) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{5}
}
func (m *AttestNominationSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AttestNominationSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Notateproposalsolicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestNominationSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Notateproposalsolicit.Merge(m, src)
}
func (m *AttestNominationSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *AttestNominationSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Notateproposalsolicit.DiscardUnknown(m)
}

var xxx_signaldetails_Notateproposalsolicit proto.InternalMessageInfo

func (m *AttestNominationSolicit) ObtainNomination() *kinds.Nomination {
	if m != nil {
		return m.Nomination
	}
	return nil
}

func (m *AttestNominationSolicit) ObtainSuccessionUuid() string {
	if m != nil {
		return m.SuccessionUuid
	}
	return "REDACTED"
}

//
type NotatedNominationReply struct {
	Nomination kinds.Nomination     `protobuf:"octets,1,opt,name=proposal,proto3" json:"nomination"`
	Failure    *RemoteEndorserFailure `protobuf:"octets,2,opt,name=error,proto3" json:"failure,omitempty"`
}

func (m *NotatedNominationReply) Restore()         { *m = NotatedNominationReply{} }
func (m *NotatedNominationReply) Text() string { return proto.CompactTextString(m) }
func (*NotatedNominationReply) SchemaArtifact()    {}
func (*NotatedNominationReply) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{6}
}
func (m *NotatedNominationReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *NotatedNominationReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Notatedproposalreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotatedNominationReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Notatedproposalreply.Merge(m, src)
}
func (m *NotatedNominationReply) XXX_Extent() int {
	return m.Extent()
}
func (m *NotatedNominationReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Notatedproposalreply.DiscardUnknown(m)
}

var xxx_signaldetails_Notatedproposalreply proto.InternalMessageInfo

func (m *NotatedNominationReply) ObtainNomination() kinds.Nomination {
	if m != nil {
		return m.Nomination
	}
	return kinds.Nomination{}
}

func (m *NotatedNominationReply) ObtainFailure() *RemoteEndorserFailure {
	if m != nil {
		return m.Failure
	}
	return nil
}

//
type PingSolicit struct {
}

func (m *PingSolicit) Restore()         { *m = PingSolicit{} }
func (m *PingSolicit) Text() string { return proto.CompactTextString(m) }
func (*PingSolicit) SchemaArtifact()    {}
func (*PingSolicit) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{7}
}
func (m *PingSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PingSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Pingsolicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Pingsolicit.Merge(m, src)
}
func (m *PingSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *PingSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Pingsolicit.DiscardUnknown(m)
}

var xxx_signaldetails_Pingsolicit proto.InternalMessageInfo

//
type PingReply struct {
}

func (m *PingReply) Restore()         { *m = PingReply{} }
func (m *PingReply) Text() string { return proto.CompactTextString(m) }
func (*PingReply) SchemaArtifact()    {}
func (*PingReply) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{8}
}
func (m *PingReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PingReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Pingreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Pingreply.Merge(m, src)
}
func (m *PingReply) XXX_Extent() int {
	return m.Extent()
}
func (m *PingReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Pingreply.DiscardUnknown(m)
}

var xxx_signaldetails_Pingreply proto.InternalMessageInfo

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
	Sum isnote_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) Text() string { return proto.CompactTextString(m) }
func (*Signal) SchemaArtifact()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedescriptor_cb4e437a5328cf9c, []int{9}
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

type Artifact_Publictokensolicit struct {
	PublicTokenSolicit *PublicTokenSolicit `protobuf:"octets,1,opt,name=pub_key_request,json=pubKeyRequest,proto3,oneof" json:"public_token_solicit,omitempty"`
}
type Artifact_Publictokendata struct {
	PublicTokenReply *PublicTokenReply `protobuf:"octets,2,opt,name=pub_key_response,json=pubKeyResponse,proto3,oneof" json:"public_token_reply,omitempty"`
}
type Artifact_Notateballotsolicit struct {
	AttestBallotSolicit *AttestBallotSolicit `protobuf:"octets,3,opt,name=sign_vote_request,json=signVoteRequest,proto3,oneof" json:"attest_ballot_solicit,omitempty"`
}
type Artifact_Notatedballotreply struct {
	NotatedBallotReply *NotatedBallotReply `protobuf:"octets,4,opt,name=signed_vote_response,json=signedVoteResponse,proto3,oneof" json:"notated_ballot_reply,omitempty"`
}
type Artifact_Notateproposalsolicit struct {
	AttestNominationSolicit *AttestNominationSolicit `protobuf:"octets,5,opt,name=sign_proposal_request,json=signProposalRequest,proto3,oneof" json:"attest_nomination_solicit,omitempty"`
}
type Artifact_Notatedproposalreply struct {
	NotatedNominationReply *NotatedNominationReply `protobuf:"octets,6,opt,name=signed_proposal_response,json=signedProposalResponse,proto3,oneof" json:"notated_nomination_reply,omitempty"`
}
type Artifact_Pingsolicit struct {
	PingSolicit *PingSolicit `protobuf:"octets,7,opt,name=ping_request,json=pingRequest,proto3,oneof" json:"ping_solicit,omitempty"`
}
type Artifact_Pingreply struct {
	PingReply *PingReply `protobuf:"octets,8,opt,name=ping_response,json=pingResponse,proto3,oneof" json:"ping_reply,omitempty"`
}

func (*Artifact_Publictokensolicit) isnote_Total()          {}
func (*Artifact_Publictokendata) isnote_Total()         {}
func (*Artifact_Notateballotsolicit) isnote_Total()        {}
func (*Artifact_Notatedballotreply) isnote_Total()     {}
func (*Artifact_Notateproposalsolicit) isnote_Total()    {}
func (*Artifact_Notatedproposalreply) isnote_Total() {}
func (*Artifact_Pingsolicit) isnote_Total()            {}
func (*Artifact_Pingreply) isnote_Total()           {}

func (m *Signal) ObtainTotal() isnote_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) ObtainPublicTokenSolicit() *PublicTokenSolicit {
	if x, ok := m.ObtainTotal().(*Artifact_Publictokensolicit); ok {
		return x.PublicTokenSolicit
	}
	return nil
}

func (m *Signal) ObtainPublicTokenReply() *PublicTokenReply {
	if x, ok := m.ObtainTotal().(*Artifact_Publictokendata); ok {
		return x.PublicTokenReply
	}
	return nil
}

func (m *Signal) ObtainAttestBallotSolicit() *AttestBallotSolicit {
	if x, ok := m.ObtainTotal().(*Artifact_Notateballotsolicit); ok {
		return x.AttestBallotSolicit
	}
	return nil
}

func (m *Signal) ObtainNotatedBallotReply() *NotatedBallotReply {
	if x, ok := m.ObtainTotal().(*Artifact_Notatedballotreply); ok {
		return x.NotatedBallotReply
	}
	return nil
}

func (m *Signal) ObtainAttestNominationSolicit() *AttestNominationSolicit {
	if x, ok := m.ObtainTotal().(*Artifact_Notateproposalsolicit); ok {
		return x.AttestNominationSolicit
	}
	return nil
}

func (m *Signal) ObtainNotatedNominationReply() *NotatedNominationReply {
	if x, ok := m.ObtainTotal().(*Artifact_Notatedproposalreply); ok {
		return x.NotatedNominationReply
	}
	return nil
}

func (m *Signal) ObtainPingSolicit() *PingSolicit {
	if x, ok := m.ObtainTotal().(*Artifact_Pingsolicit); ok {
		return x.PingSolicit
	}
	return nil
}

func (m *Signal) ObtainPingReply() *PingReply {
	if x, ok := m.ObtainTotal().(*Artifact_Pingreply); ok {
		return x.PingReply
	}
	return nil
}

//
func (*Signal) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Artifact_Publictokensolicit)(nil),
		(*Artifact_Publictokendata)(nil),
		(*Artifact_Notateballotsolicit)(nil),
		(*Artifact_Notatedballotreply)(nil),
		(*Artifact_Notateproposalsolicit)(nil),
		(*Artifact_Notatedproposalreply)(nil),
		(*Artifact_Pingsolicit)(nil),
		(*Artifact_Pingreply)(nil),
	}
}

func initialize() {
	proto.RegisterEnum("REDACTED", Faults_alias, Faults_datum)
	proto.RegisterType((*RemoteEndorserFailure)(nil), "REDACTED")
	proto.RegisterType((*PublicTokenSolicit)(nil), "REDACTED")
	proto.RegisterType((*PublicTokenReply)(nil), "REDACTED")
	proto.RegisterType((*AttestBallotSolicit)(nil), "REDACTED")
	proto.RegisterType((*NotatedBallotReply)(nil), "REDACTED")
	proto.RegisterType((*AttestNominationSolicit)(nil), "REDACTED")
	proto.RegisterType((*NotatedNominationReply)(nil), "REDACTED")
	proto.RegisterType((*PingSolicit)(nil), "REDACTED")
	proto.RegisterType((*PingReply)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_cb4e437a5328cf9c) }

var filedescriptor_cb4e437a5328cf9c = []byte{
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

func (m *RemoteEndorserFailure) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *RemoteEndorserFailure) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *RemoteEndorserFailure) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Characterization) > 0 {
		i -= len(m.Characterization)
		copy(deltaLocatedAN[i:], m.Characterization)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Characterization)))
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

func (m *PublicTokenSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PublicTokenSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PublicTokenSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.SuccessionUuid) > 0 {
		i -= len(m.SuccessionUuid)
		copy(deltaLocatedAN[i:], m.SuccessionUuid)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.SuccessionUuid)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *PublicTokenReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PublicTokenReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PublicTokenReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Failure != nil {
		{
			extent, err := m.Failure.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
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

func (m *AttestBallotSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AttestBallotSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AttestBallotSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.SuccessionUuid) > 0 {
		i -= len(m.SuccessionUuid)
		copy(deltaLocatedAN[i:], m.SuccessionUuid)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.SuccessionUuid)))
		i--
		deltaLocatedAN[i] = 0x12
	}
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

func (m *NotatedBallotReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *NotatedBallotReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *NotatedBallotReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Failure != nil {
		{
			extent, err := m.Failure.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
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
	return len(deltaLocatedAN) - i, nil
}

func (m *AttestNominationSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AttestNominationSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AttestNominationSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.SuccessionUuid) > 0 {
		i -= len(m.SuccessionUuid)
		copy(deltaLocatedAN[i:], m.SuccessionUuid)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.SuccessionUuid)))
		i--
		deltaLocatedAN[i] = 0x12
	}
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
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *NotatedNominationReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *NotatedNominationReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *NotatedNominationReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Failure != nil {
		{
			extent, err := m.Failure.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
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

func (m *PingSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PingSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PingSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *PingReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PingReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PingReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
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

func (m *Artifact_Publictokensolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Publictokensolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PublicTokenSolicit != nil {
		{
			extent, err := m.PublicTokenSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Publictokendata) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Publictokendata) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PublicTokenReply != nil {
		{
			extent, err := m.PublicTokenReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Notateballotsolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Notateballotsolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.AttestBallotSolicit != nil {
		{
			extent, err := m.AttestBallotSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Notatedballotreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Notatedballotreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.NotatedBallotReply != nil {
		{
			extent, err := m.NotatedBallotReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Notateproposalsolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Notateproposalsolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.AttestNominationSolicit != nil {
		{
			extent, err := m.AttestNominationSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Notatedproposalreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Notatedproposalreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.NotatedNominationReply != nil {
		{
			extent, err := m.NotatedNominationReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Pingsolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Pingsolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PingSolicit != nil {
		{
			extent, err := m.PingSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Pingreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Pingreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PingReply != nil {
		{
			extent, err := m.PingReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *RemoteEndorserFailure) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cipher != 0 {
		n += 1 + sovKinds(uint64(m.Cipher))
	}
	l = len(m.Characterization)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *PublicTokenSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SuccessionUuid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *PublicTokenReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PublicToken.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.Failure != nil {
		l = m.Failure.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AttestBallotSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ballot != nil {
		l = m.Ballot.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.SuccessionUuid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *NotatedBallotReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Ballot.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.Failure != nil {
		l = m.Failure.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *AttestNominationSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nomination != nil {
		l = m.Nomination.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.SuccessionUuid)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *NotatedNominationReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Nomination.Extent()
	n += 1 + l + sovKinds(uint64(l))
	if m.Failure != nil {
		l = m.Failure.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *PingSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PingReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
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

func (m *Artifact_Publictokensolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicTokenSolicit != nil {
		l = m.PublicTokenSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Publictokendata) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicTokenReply != nil {
		l = m.PublicTokenReply.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Notateballotsolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestBallotSolicit != nil {
		l = m.AttestBallotSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Notatedballotreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotatedBallotReply != nil {
		l = m.NotatedBallotReply.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Notateproposalsolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestNominationSolicit != nil {
		l = m.AttestNominationSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Notatedproposalreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotatedNominationReply != nil {
		l = m.NotatedNominationReply.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Pingsolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingSolicit != nil {
		l = m.PingSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Pingreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingReply != nil {
		l = m.PingReply.Extent()
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
func (m *RemoteEndorserFailure) Decode(deltaLocatedAN []byte) error {
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
				m.Cipher |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
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
			m.Characterization = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *PublicTokenSolicit) Decode(deltaLocatedAN []byte) error {
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
			m.SuccessionUuid = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *PublicTokenReply) Decode(deltaLocatedAN []byte) error {
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
			if m.Failure == nil {
				m.Failure = &RemoteEndorserFailure{}
			}
			if err := m.Failure.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *AttestBallotSolicit) Decode(deltaLocatedAN []byte) error {
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
func (m *NotatedBallotReply) Decode(deltaLocatedAN []byte) error {
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
			if err := m.Ballot.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.Failure == nil {
				m.Failure = &RemoteEndorserFailure{}
			}
			if err := m.Failure.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *AttestNominationSolicit) Decode(deltaLocatedAN []byte) error {
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
			if m.Nomination == nil {
				m.Nomination = &kinds.Nomination{}
			}
			if err := m.Nomination.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *NotatedNominationReply) Decode(deltaLocatedAN []byte) error {
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
			if m.Failure == nil {
				m.Failure = &RemoteEndorserFailure{}
			}
			if err := m.Failure.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *PingSolicit) Decode(deltaLocatedAN []byte) error {
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
func (m *PingReply) Decode(deltaLocatedAN []byte) error {
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
			v := &PublicTokenSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Publictokensolicit{v}
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
			v := &PublicTokenReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Publictokendata{v}
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
			v := &AttestBallotSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Notateballotsolicit{v}
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
			v := &NotatedBallotReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Notatedballotreply{v}
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
			v := &AttestNominationSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Notateproposalsolicit{v}
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
			v := &NotatedNominationReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Notatedproposalreply{v}
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
			v := &PingSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Pingsolicit{v}
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
			v := &PingReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Pingreply{v}
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
