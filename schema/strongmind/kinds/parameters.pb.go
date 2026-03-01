//
//

package kinds

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "github.com/golang/protobuf/ptypes/duration"
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
type AgreementSettings struct {
	Ledger     *LedgerParameters     `protobuf:"octets,1,opt,name=block,proto3" json:"ledger,omitempty"`
	Proof  *ProofParameters  `protobuf:"octets,2,opt,name=evidence,proto3" json:"proof,omitempty"`
	Assessor *AssessorParameters `protobuf:"octets,3,opt,name=validator,proto3" json:"assessor,omitempty"`
	Edition   *EditionParameters   `protobuf:"octets,4,opt,name=version,proto3" json:"edition,omitempty"`
	Iface      *IfaceParameters      `protobuf:"octets,5,opt,name=abci,proto3" json:"iface,omitempty"`
}

func (m *AgreementSettings) Restore()         { *m = AgreementSettings{} }
func (m *AgreementSettings) Text() string { return proto.CompactTextString(m) }
func (*AgreementSettings) SchemaArtifact()    {}
func (*AgreementSettings) Definition() ([]byte, []int) {
	return filedescriptor_e12598271a686f57, []int{0}
}
func (m *AgreementSettings) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AgreementSettings) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Agreementparameters.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgreementSettings) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Agreementparameters.Merge(m, src)
}
func (m *AgreementSettings) XXX_Extent() int {
	return m.Extent()
}
func (m *AgreementSettings) XXX_Dropunfamiliar() {
	xxx_signaldetails_Agreementparameters.DiscardUnknown(m)
}

var xxx_signaldetails_Agreementparameters proto.InternalMessageInfo

func (m *AgreementSettings) ObtainLedger() *LedgerParameters {
	if m != nil {
		return m.Ledger
	}
	return nil
}

func (m *AgreementSettings) ObtainProof() *ProofParameters {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *AgreementSettings) ObtainAssessor() *AssessorParameters {
	if m != nil {
		return m.Assessor
	}
	return nil
}

func (m *AgreementSettings) ObtainEdition() *EditionParameters {
	if m != nil {
		return m.Edition
	}
	return nil
}

func (m *AgreementSettings) ObtainIface() *IfaceParameters {
	if m != nil {
		return m.Iface
	}
	return nil
}

//
type LedgerParameters struct {
	//
	//
	MaximumOctets int64 `protobuf:"variableint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"maximum_octets,omitempty"`
	//
	//
	MaximumFuel int64 `protobuf:"variableint,2,opt,name=max_gas,json=maxGas,proto3" json:"maximum_fuel,omitempty"`
}

func (m *LedgerParameters) Restore()         { *m = LedgerParameters{} }
func (m *LedgerParameters) Text() string { return proto.CompactTextString(m) }
func (*LedgerParameters) SchemaArtifact()    {}
func (*LedgerParameters) Definition() ([]byte, []int) {
	return filedescriptor_e12598271a686f57, []int{1}
}
func (m *LedgerParameters) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *LedgerParameters) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ledgerparameters.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerParameters) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ledgerparameters.Merge(m, src)
}
func (m *LedgerParameters) XXX_Extent() int {
	return m.Extent()
}
func (m *LedgerParameters) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ledgerparameters.DiscardUnknown(m)
}

var xxx_signaldetails_Ledgerparameters proto.InternalMessageInfo

func (m *LedgerParameters) ObtainMaximumOctets() int64 {
	if m != nil {
		return m.MaximumOctets
	}
	return 0
}

func (m *LedgerParameters) ObtainMaximumFuel() int64 {
	if m != nil {
		return m.MaximumFuel
	}
	return 0
}

//
type ProofParameters struct {
	//
	//
	//
	//
	MaximumLifespanCountLedgers int64 `protobuf:"variableint,1,opt,name=max_age_num_blocks,json=maxAgeNumBlocks,proto3" json:"maximum_lifespan_count_ledgers,omitempty"`
	//
	//
	//
	//
	//
	MaximumLifespanInterval time.Duration `protobuf:"octets,2,opt,name=max_age_duration,json=maxAgeDuration,proto3,stdduration" json:"maximum_lifespan_interval"`
	//
	//
	//
	MaximumOctets int64 `protobuf:"variableint,3,opt,name=max_bytes,json=maxBytes,proto3" json:"maximum_octets,omitempty"`
}

func (m *ProofParameters) Restore()         { *m = ProofParameters{} }
func (m *ProofParameters) Text() string { return proto.CompactTextString(m) }
func (*ProofParameters) SchemaArtifact()    {}
func (*ProofParameters) Definition() ([]byte, []int) {
	return filedescriptor_e12598271a686f57, []int{2}
}
func (m *ProofParameters) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ProofParameters) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Prooffactors.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofParameters) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Prooffactors.Merge(m, src)
}
func (m *ProofParameters) XXX_Extent() int {
	return m.Extent()
}
func (m *ProofParameters) XXX_Dropunfamiliar() {
	xxx_signaldetails_Prooffactors.DiscardUnknown(m)
}

var xxx_signaldetails_Prooffactors proto.InternalMessageInfo

func (m *ProofParameters) ObtainMaximumLifespanCountLedgers() int64 {
	if m != nil {
		return m.MaximumLifespanCountLedgers
	}
	return 0
}

func (m *ProofParameters) ObtainMaximumLifespanInterval() time.Duration {
	if m != nil {
		return m.MaximumLifespanInterval
	}
	return 0
}

func (m *ProofParameters) ObtainMaximumOctets() int64 {
	if m != nil {
		return m.MaximumOctets
	}
	return 0
}

//
//
type AssessorParameters struct {
	PublicTokenKinds []string `protobuf:"octets,1,rep,name=pub_key_types,json=pubKeyTypes,proto3" json:"public_token_kinds,omitempty"`
}

func (m *AssessorParameters) Restore()         { *m = AssessorParameters{} }
func (m *AssessorParameters) Text() string { return proto.CompactTextString(m) }
func (*AssessorParameters) SchemaArtifact()    {}
func (*AssessorParameters) Definition() ([]byte, []int) {
	return filedescriptor_e12598271a686f57, []int{3}
}
func (m *AssessorParameters) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *AssessorParameters) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Assessorparameters.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssessorParameters) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Assessorparameters.Merge(m, src)
}
func (m *AssessorParameters) XXX_Extent() int {
	return m.Extent()
}
func (m *AssessorParameters) XXX_Dropunfamiliar() {
	xxx_signaldetails_Assessorparameters.DiscardUnknown(m)
}

var xxx_signaldetails_Assessorparameters proto.InternalMessageInfo

func (m *AssessorParameters) ObtainPublicTokenKinds() []string {
	if m != nil {
		return m.PublicTokenKinds
	}
	return nil
}

//
type EditionParameters struct {
	App uint64 `protobuf:"variableint,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (m *EditionParameters) Restore()         { *m = EditionParameters{} }
func (m *EditionParameters) Text() string { return proto.CompactTextString(m) }
func (*EditionParameters) SchemaArtifact()    {}
func (*EditionParameters) Definition() ([]byte, []int) {
	return filedescriptor_e12598271a686f57, []int{4}
}
func (m *EditionParameters) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *EditionParameters) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Editionparameters.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditionParameters) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Editionparameters.Merge(m, src)
}
func (m *EditionParameters) XXX_Extent() int {
	return m.Extent()
}
func (m *EditionParameters) XXX_Dropunfamiliar() {
	xxx_signaldetails_Editionparameters.DiscardUnknown(m)
}

var xxx_signaldetails_Editionparameters proto.InternalMessageInfo

func (m *EditionParameters) ObtainApplication() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

//
//
//
type DigestedParameters struct {
	LedgerMaximumOctets int64 `protobuf:"variableint,1,opt,name=block_max_bytes,json=blockMaxBytes,proto3" json:"ledger_maximum_octets,omitempty"`
	LedgerMaximumFuel   int64 `protobuf:"variableint,2,opt,name=block_max_gas,json=blockMaxGas,proto3" json:"ledger_maximum_fuel,omitempty"`
}

func (m *DigestedParameters) Restore()         { *m = DigestedParameters{} }
func (m *DigestedParameters) Text() string { return proto.CompactTextString(m) }
func (*DigestedParameters) SchemaArtifact()    {}
func (*DigestedParameters) Definition() ([]byte, []int) {
	return filedescriptor_e12598271a686f57, []int{5}
}
func (m *DigestedParameters) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *DigestedParameters) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Digestedparameters.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DigestedParameters) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Digestedparameters.Merge(m, src)
}
func (m *DigestedParameters) XXX_Extent() int {
	return m.Extent()
}
func (m *DigestedParameters) XXX_Dropunfamiliar() {
	xxx_signaldetails_Digestedparameters.DiscardUnknown(m)
}

var xxx_signaldetails_Digestedparameters proto.InternalMessageInfo

func (m *DigestedParameters) ObtainLedgerMaximumOctets() int64 {
	if m != nil {
		return m.LedgerMaximumOctets
	}
	return 0
}

func (m *DigestedParameters) ObtainLedgerMaximumFuel() int64 {
	if m != nil {
		return m.LedgerMaximumFuel
	}
	return 0
}

//
type IfaceParameters struct {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	BallotAdditionsActivateAltitude int64 `protobuf:"variableint,1,opt,name=vote_extensions_enable_height,json=voteExtensionsEnableHeight,proto3" json:"ballot_additions_activate_altitude,omitempty"`
}

func (m *IfaceParameters) Restore()         { *m = IfaceParameters{} }
func (m *IfaceParameters) Text() string { return proto.CompactTextString(m) }
func (*IfaceParameters) SchemaArtifact()    {}
func (*IfaceParameters) Definition() ([]byte, []int) {
	return filedescriptor_e12598271a686f57, []int{6}
}
func (m *IfaceParameters) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *IfaceParameters) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ifaceparameters.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IfaceParameters) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ifaceparameters.Merge(m, src)
}
func (m *IfaceParameters) XXX_Extent() int {
	return m.Extent()
}
func (m *IfaceParameters) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ifaceparameters.DiscardUnknown(m)
}

var xxx_signaldetails_Ifaceparameters proto.InternalMessageInfo

func (m *IfaceParameters) ObtainBallotAdditionsActivateAltitude() int64 {
	if m != nil {
		return m.BallotAdditionsActivateAltitude
	}
	return 0
}

func initialize() {
	proto.RegisterType((*AgreementSettings)(nil), "REDACTED")
	proto.RegisterType((*LedgerParameters)(nil), "REDACTED")
	proto.RegisterType((*ProofParameters)(nil), "REDACTED")
	proto.RegisterType((*AssessorParameters)(nil), "REDACTED")
	proto.RegisterType((*EditionParameters)(nil), "REDACTED")
	proto.RegisterType((*DigestedParameters)(nil), "REDACTED")
	proto.RegisterType((*IfaceParameters)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_e12598271a686f57) }

var filedescriptor_e12598271a686f57 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0x73, 0x75, 0xda, 0xa6, 0x4f, 0x7e, 0x69, 0xa2, 0xd3, 0x4f, 0xc2, 0x14, 0xe2, 0x04,
	0x0f, 0xa8, 0x52, 0x25, 0x1b, 0x91, 0x09, 0x84, 0x54, 0x25, 0x25, 0x6a, 0x0b, 0x2a, 0x7f, 0x2c,
	0xc4, 0xd0, 0xc5, 0x3a, 0x27, 0x57, 0xc7, 0x6a, 0xec, 0xb3, 0x7c, 0xe7, 0x28, 0x79, 0x17, 0x8c,
	0x8c, 0x1d, 0x61, 0x65, 0xe2, 0x25, 0x74, 0xec, 0xc8, 0x04, 0x28, 0x59, 0x78, 0x19, 0xc8, 0x67,
	0xbb, 0x6e, 0x12, 0xb6, 0xbb, 0x7b, 0x3e, 0x9f, 0xfb, 0xf3, 0x7d, 0x74, 0xd0, 0x14, 0x34, 0x18,
	0xd2, 0xc8, 0xf7, 0x02, 0x61, 0x8a, 0x59, 0x48, 0xb9, 0x19, 0x92, 0x88, 0xf8, 0xdc, 0x08, 0x23,
	0x26, 0x18, 0x6e, 0x14, 0x65, 0x43, 0x96, 0xf7, 0xfe, 0x77, 0x99, 0xcb, 0x64, 0xd1, 0x4c, 0x46,
	0x29, 0xb7, 0xa7, 0xb9, 0x8c, 0xb9, 0x63, 0x6a, 0xca, 0x99, 0x13, 0x5f, 0x98, 0xc3, 0x38, 0x22,
	0xc2, 0x63, 0x41, 0x5a, 0xd7, 0xbf, 0x6d, 0x40, 0xfd, 0x88, 0x05, 0x9c, 0x06, 0x3c, 0xe6, 0xef,
	0xe4, 0x09, 0xb8, 0x03, 0x9b, 0xce, 0x98, 0x0d, 0x2e, 0x55, 0xd4, 0x46, 0xfb, 0xd5, 0xa7, 0x4d,
	0x63, 0xf5, 0x2c, 0xa3, 0x97, 0x94, 0x53, 0xda, 0x4a, 0x59, 0xfc, 0x02, 0x2a, 0x74, 0xe2, 0x0d,
	0x69, 0x30, 0xa0, 0xea, 0x86, 0xf4, 0xda, 0xeb, 0x5e, 0x3f, 0x23, 0x32, 0xf5, 0xd6, 0xc0, 0x87,
	0xb0, 0x33, 0x21, 0x63, 0x6f, 0x48, 0x04, 0x8b, 0x54, 0x45, 0xea, 0x8f, 0xd6, 0xf5, 0x8f, 0x39,
	0x92, 0xf9, 0x85, 0x83, 0x9f, 0xc1, 0xf6, 0x84, 0x46, 0xdc, 0x63, 0x81, 0x5a, 0x96, 0x7a, 0xeb,
	0x1f, 0x7a, 0x0a, 0x64, 0x72, 0xce, 0xe3, 0x27, 0x50, 0x26, 0xce, 0xc0, 0x53, 0x37, 0xa5, 0xf7,
	0x70, 0xdd, 0xeb, 0xf6, 0x8e, 0x4e, 0x33, 0x49, 0x92, 0xfa, 0x29, 0x54, 0xef, 0x24, 0x80, 0x1f,
	0xc0, 0x8e, 0x4f, 0xa6, 0xb6, 0x33, 0x13, 0x94, 0xcb, 0xcc, 0x14, 0xab, 0xe2, 0x93, 0x69, 0x2f,
	0x99, 0xe3, 0x7b, 0xb0, 0x9d, 0x14, 0x5d, 0xc2, 0x65, 0x2c, 0x8a, 0xb5, 0xe5, 0x93, 0xe9, 0x31,
	0xe1, 0xaf, 0xca, 0x15, 0xa5, 0x51, 0xd6, 0xbf, 0x22, 0xd8, 0x5d, 0x4e, 0x05, 0x1f, 0x00, 0x4e,
	0x0c, 0xe2, 0x52, 0x3b, 0x88, 0x7d, 0x5b, 0xc6, 0x9b, 0xef, 0x5b, 0xf7, 0xc9, 0xb4, 0xeb, 0xd2,
	0x37, 0xb1, 0x2f, 0x2f, 0xc0, 0xf1, 0x19, 0x34, 0x72, 0x38, 0xef, 0x6c, 0x16, 0xff, 0x7d, 0x23,
	0x6d, 0xbd, 0x91, 0xb7, 0xde, 0x78, 0x99, 0x01, 0xbd, 0xca, 0xf5, 0xcf, 0x56, 0xe9, 0xf3, 0xaf,
	0x16, 0xb2, 0x76, 0xd3, 0xfd, 0xf2, 0xca, 0xf2, 0x53, 0x94, 0xe5, 0xa7, 0xe8, 0x87, 0x50, 0x5f,
	0xe9, 0x00, 0xd6, 0xa1, 0x16, 0xc6, 0x8e, 0x7d, 0x49, 0x67, 0xb6, 0xcc, 0x4a, 0x45, 0x6d, 0x65,
	0x7f, 0xc7, 0xaa, 0x86, 0xb1, 0xf3, 0x9a, 0xce, 0x3e, 0x24, 0x4b, 0xcf, 0x2b, 0xdf, 0xaf, 0x5a,
	0xe8, 0xcf, 0x55, 0x0b, 0xe9, 0x07, 0x50, 0x5b, 0xea, 0x01, 0x6e, 0x80, 0x42, 0xc2, 0x50, 0xbe,
	0xad, 0x6c, 0x25, 0xc3, 0x3b, 0xf0, 0x39, 0xfc, 0x77, 0x42, 0xf8, 0x88, 0x0e, 0x33, 0xf6, 0x31,
	0xd4, 0x65, 0x14, 0xf6, 0x6a, 0xd6, 0x35, 0xb9, 0x7c, 0x96, 0x07, 0xae, 0x43, 0xad, 0xe0, 0x8a,
	0xd8, 0xab, 0x39, 0x75, 0x4c, 0xb8, 0xfe, 0x16, 0xa0, 0x68, 0x2a, 0xee, 0x42, 0x73, 0xc2, 0x04,
	0xb5, 0xe9, 0x54, 0xd0, 0x20, 0xb9, 0x1d, 0xb7, 0x69, 0x40, 0x9c, 0x31, 0xb5, 0x47, 0xd4, 0x73,
	0x47, 0x22, 0x3b, 0x67, 0x2f, 0x81, 0xfa, 0xb7, 0x4c, 0x5f, 0x22, 0x27, 0x92, 0xe8, 0xbd, 0xff,
	0x32, 0xd7, 0xd0, 0xf5, 0x5c, 0x43, 0x37, 0x73, 0x0d, 0xfd, 0x9e, 0x6b, 0xe8, 0xd3, 0x42, 0x2b,
	0xdd, 0x2c, 0xb4, 0xd2, 0x8f, 0x85, 0x56, 0x3a, 0xef, 0xb8, 0x9e, 0x18, 0xc5, 0x8e, 0x31, 0x60,
	0xbe, 0x39, 0x60, 0x3e, 0x15, 0xce, 0x85, 0x28, 0x06, 0xe9, 0x9f, 0x5d, 0xfd, 0xee, 0xce, 0x96,
	0x5c, 0xef, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xe8, 0xce, 0x9a, 0x09, 0x04, 0x00, 0x00,
}

func (that *AgreementSettings) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*AgreementSettings)
	if !ok {
		which2, ok := which.(AgreementSettings)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if !that.Ledger.Equivalent(which1.Ledger) {
		return false
	}
	if !that.Proof.Equivalent(which1.Proof) {
		return false
	}
	if !that.Assessor.Equivalent(which1.Assessor) {
		return false
	}
	if !that.Edition.Equivalent(which1.Edition) {
		return false
	}
	if !that.Iface.Equivalent(which1.Iface) {
		return false
	}
	return true
}
func (that *LedgerParameters) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*LedgerParameters)
	if !ok {
		which2, ok := which.(LedgerParameters)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if that.MaximumOctets != which1.MaximumOctets {
		return false
	}
	if that.MaximumFuel != which1.MaximumFuel {
		return false
	}
	return true
}
func (that *ProofParameters) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*ProofParameters)
	if !ok {
		which2, ok := which.(ProofParameters)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if that.MaximumLifespanCountLedgers != which1.MaximumLifespanCountLedgers {
		return false
	}
	if that.MaximumLifespanInterval != which1.MaximumLifespanInterval {
		return false
	}
	if that.MaximumOctets != which1.MaximumOctets {
		return false
	}
	return true
}
func (that *AssessorParameters) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*AssessorParameters)
	if !ok {
		which2, ok := which.(AssessorParameters)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if len(that.PublicTokenKinds) != len(which1.PublicTokenKinds) {
		return false
	}
	for i := range that.PublicTokenKinds {
		if that.PublicTokenKinds[i] != which1.PublicTokenKinds[i] {
			return false
		}
	}
	return true
}
func (that *EditionParameters) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*EditionParameters)
	if !ok {
		which2, ok := which.(EditionParameters)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if that.App != which1.App {
		return false
	}
	return true
}
func (that *DigestedParameters) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*DigestedParameters)
	if !ok {
		which2, ok := which.(DigestedParameters)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if that.LedgerMaximumOctets != which1.LedgerMaximumOctets {
		return false
	}
	if that.LedgerMaximumFuel != which1.LedgerMaximumFuel {
		return false
	}
	return true
}
func (that *IfaceParameters) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*IfaceParameters)
	if !ok {
		which2, ok := which.(IfaceParameters)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if that.BallotAdditionsActivateAltitude != which1.BallotAdditionsActivateAltitude {
		return false
	}
	return true
}
func (m *AgreementSettings) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AgreementSettings) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AgreementSettings) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Iface != nil {
		{
			extent, err := m.Iface.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintParameters(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if m.Edition != nil {
		{
			extent, err := m.Edition.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintParameters(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Assessor != nil {
		{
			extent, err := m.Assessor.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintParameters(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.Proof != nil {
		{
			extent, err := m.Proof.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintParameters(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Ledger != nil {
		{
			extent, err := m.Ledger.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintParameters(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *LedgerParameters) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *LedgerParameters) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *LedgerParameters) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.MaximumFuel != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.MaximumFuel))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.MaximumOctets != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.MaximumOctets))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ProofParameters) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ProofParameters) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ProofParameters) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.MaximumOctets != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.MaximumOctets))
		i--
		deltaLocatedAN[i] = 0x18
	}
	n6, fault6 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaximumLifespanInterval, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaximumLifespanInterval):])
	if fault6 != nil {
		return 0, fault6
	}
	i -= n6
	i = encodeVariableintParameters(deltaLocatedAN, i, uint64(n6))
	i--
	deltaLocatedAN[i] = 0x12
	if m.MaximumLifespanCountLedgers != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.MaximumLifespanCountLedgers))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *AssessorParameters) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *AssessorParameters) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *AssessorParameters) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.PublicTokenKinds) > 0 {
		for idxNdExc := len(m.PublicTokenKinds) - 1; idxNdExc >= 0; idxNdExc-- {
			i -= len(m.PublicTokenKinds[idxNdExc])
			copy(deltaLocatedAN[i:], m.PublicTokenKinds[idxNdExc])
			i = encodeVariableintParameters(deltaLocatedAN, i, uint64(len(m.PublicTokenKinds[idxNdExc])))
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *EditionParameters) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *EditionParameters) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *EditionParameters) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.App))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *DigestedParameters) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *DigestedParameters) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *DigestedParameters) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.LedgerMaximumFuel != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.LedgerMaximumFuel))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.LedgerMaximumOctets != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.LedgerMaximumOctets))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *IfaceParameters) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *IfaceParameters) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *IfaceParameters) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.BallotAdditionsActivateAltitude != 0 {
		i = encodeVariableintParameters(deltaLocatedAN, i, uint64(m.BallotAdditionsActivateAltitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintParameters(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovParameters(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func FreshInhabitedAssessorParameters(r arbitraryyParameters, easy bool) *AssessorParameters {
	that := &AssessorParameters{}
	v1 := r.Integern(10)
	that.PublicTokenKinds = make([]string, v1)
	for i := 0; i < v1; i++ {
		that.PublicTokenKinds[i] = string(arbitraryTextParameters(r))
	}
	if !easy && r.Integern(10) != 0 {
	}
	return that
}

func FreshInhabitedEditionParameters(r arbitraryyParameters, easy bool) *EditionParameters {
	that := &EditionParameters{}
	that.App = uint64(uint64(r.Uint32n()))
	if !easy && r.Integern(10) != 0 {
	}
	return that
}

type arbitraryyParameters interface {
	Float32() float32
	Float64() float64
	Int63n() int64
	Int31n() int32
	Uint32n() uint32
	Integern(n int) int
}

func arbitraryUtf8runeParameters(r arbitraryyParameters) rune {
	ru := r.Integern(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func arbitraryTextParameters(r arbitraryyParameters) string {
	v2 := r.Integern(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = arbitraryUtf8runeParameters(r)
	}
	return string(tmps)
}
func arbitraryUnidentifiedParameters(r arbitraryyParameters, maximumAttributeNumeral int) (deltaLocatedAN []byte) {
	l := r.Integern(5)
	for i := 0; i < l; i++ {
		cable := r.Integern(4)
		if cable == 3 {
			cable = 5
		}
		attributeNumeral := maximumAttributeNumeral + r.Integern(100)
		deltaLocatedAN = arbitraryAttributeParameters(deltaLocatedAN, r, attributeNumeral, cable)
	}
	return deltaLocatedAN
}
func arbitraryAttributeParameters(deltaLocatedAN []byte, r arbitraryyParameters, attributeNumeral int, cable int) []byte {
	key := uint32(attributeNumeral)<<3 | uint32(cable)
	switch cable {
	case 0:
		deltaLocatedAN = encodeVariableintInhabitParameters(deltaLocatedAN, uint64(key))
		v3 := r.Int63n()
		if r.Integern(2) == 0 {
			v3 *= -1
		}
		deltaLocatedAN = encodeVariableintInhabitParameters(deltaLocatedAN, uint64(v3))
	case 1:
		deltaLocatedAN = encodeVariableintInhabitParameters(deltaLocatedAN, uint64(key))
		deltaLocatedAN = append(deltaLocatedAN, byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)))
	case 2:
		deltaLocatedAN = encodeVariableintInhabitParameters(deltaLocatedAN, uint64(key))
		ll := r.Integern(100)
		deltaLocatedAN = encodeVariableintInhabitParameters(deltaLocatedAN, uint64(ll))
		for j := 0; j < ll; j++ {
			deltaLocatedAN = append(deltaLocatedAN, byte(r.Integern(256)))
		}
	default:
		deltaLocatedAN = encodeVariableintInhabitParameters(deltaLocatedAN, uint64(key))
		deltaLocatedAN = append(deltaLocatedAN, byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)), byte(r.Integern(256)))
	}
	return deltaLocatedAN
}
func encodeVariableintInhabitParameters(deltaLocatedAN []byte, v uint64) []byte {
	for v >= 1<<7 {
		deltaLocatedAN = append(deltaLocatedAN, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	deltaLocatedAN = append(deltaLocatedAN, uint8(v))
	return deltaLocatedAN
}
func (m *AgreementSettings) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ledger != nil {
		l = m.Ledger.Extent()
		n += 1 + l + sovParameters(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Extent()
		n += 1 + l + sovParameters(uint64(l))
	}
	if m.Assessor != nil {
		l = m.Assessor.Extent()
		n += 1 + l + sovParameters(uint64(l))
	}
	if m.Edition != nil {
		l = m.Edition.Extent()
		n += 1 + l + sovParameters(uint64(l))
	}
	if m.Iface != nil {
		l = m.Iface.Extent()
		n += 1 + l + sovParameters(uint64(l))
	}
	return n
}

func (m *LedgerParameters) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumOctets != 0 {
		n += 1 + sovParameters(uint64(m.MaximumOctets))
	}
	if m.MaximumFuel != 0 {
		n += 1 + sovParameters(uint64(m.MaximumFuel))
	}
	return n
}

func (m *ProofParameters) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumLifespanCountLedgers != 0 {
		n += 1 + sovParameters(uint64(m.MaximumLifespanCountLedgers))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaximumLifespanInterval)
	n += 1 + l + sovParameters(uint64(l))
	if m.MaximumOctets != 0 {
		n += 1 + sovParameters(uint64(m.MaximumOctets))
	}
	return n
}

func (m *AssessorParameters) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PublicTokenKinds) > 0 {
		for _, s := range m.PublicTokenKinds {
			l = len(s)
			n += 1 + l + sovParameters(uint64(l))
		}
	}
	return n
}

func (m *EditionParameters) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.App != 0 {
		n += 1 + sovParameters(uint64(m.App))
	}
	return n
}

func (m *DigestedParameters) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerMaximumOctets != 0 {
		n += 1 + sovParameters(uint64(m.LedgerMaximumOctets))
	}
	if m.LedgerMaximumFuel != 0 {
		n += 1 + sovParameters(uint64(m.LedgerMaximumFuel))
	}
	return n
}

func (m *IfaceParameters) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotAdditionsActivateAltitude != 0 {
		n += 1 + sovParameters(uint64(m.BallotAdditionsActivateAltitude))
	}
	return n
}

func sovParameters(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParameters(x uint64) (n int) {
	return sovParameters(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AgreementSettings) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunParameters
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
					return FaultIntegerOverrunParameters
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
				return FaultUnfitMagnitudeParameters
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeParameters
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ledger == nil {
				m.Ledger = &LedgerParameters{}
			}
			if err := m.Ledger.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
					return FaultIntegerOverrunParameters
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
				return FaultUnfitMagnitudeParameters
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeParameters
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &ProofParameters{}
			}
			if err := m.Proof.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
					return FaultIntegerOverrunParameters
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
				return FaultUnfitMagnitudeParameters
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeParameters
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Assessor == nil {
				m.Assessor = &AssessorParameters{}
			}
			if err := m.Assessor.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
					return FaultIntegerOverrunParameters
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
				return FaultUnfitMagnitudeParameters
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeParameters
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Edition == nil {
				m.Edition = &EditionParameters{}
			}
			if err := m.Edition.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
					return FaultIntegerOverrunParameters
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
				return FaultUnfitMagnitudeParameters
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeParameters
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Iface == nil {
				m.Iface = &IfaceParameters{}
			}
			if err := m.Iface.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitParameters(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeParameters
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
func (m *LedgerParameters) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunParameters
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
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MaximumOctets |= int64(b&0x7F) << relocate
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
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MaximumFuel |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitParameters(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeParameters
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
func (m *ProofParameters) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunParameters
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
			m.MaximumLifespanCountLedgers = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MaximumLifespanCountLedgers |= int64(b&0x7F) << relocate
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
					return FaultIntegerOverrunParameters
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
				return FaultUnfitMagnitudeParameters
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeParameters
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaximumLifespanInterval, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.MaximumOctets = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.MaximumOctets |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitParameters(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeParameters
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
func (m *AssessorParameters) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunParameters
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
					return FaultIntegerOverrunParameters
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
				return FaultUnfitMagnitudeParameters
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeParameters
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicTokenKinds = append(m.PublicTokenKinds, string(deltaLocatedAN[idxNdExc:submitOrdinal]))
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitParameters(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeParameters
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
func (m *EditionParameters) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunParameters
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
			m.App = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.App |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitParameters(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeParameters
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
func (m *DigestedParameters) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunParameters
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
			m.LedgerMaximumOctets = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.LedgerMaximumOctets |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.LedgerMaximumFuel = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.LedgerMaximumFuel |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitParameters(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeParameters
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
func (m *IfaceParameters) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunParameters
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
			m.BallotAdditionsActivateAltitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunParameters
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.BallotAdditionsActivateAltitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitParameters(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeParameters
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
func omitParameters(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunParameters
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
					return 0, FaultIntegerOverrunParameters
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
					return 0, FaultIntegerOverrunParameters
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
				return 0, FaultUnfitMagnitudeParameters
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionParameters
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeParameters
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeParameters        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunParameters          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionParameters = fmt.Errorf("REDACTED")
)
