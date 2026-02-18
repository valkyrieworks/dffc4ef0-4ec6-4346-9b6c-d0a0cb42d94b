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
type AgreementOptions struct {
	Ledger     *LedgerOptions     `protobuf:"octets,1,opt,name=block,proto3" json:"ledger,omitempty"`
	Proof  *ProofOptions  `protobuf:"octets,2,opt,name=evidence,proto3" json:"proof,omitempty"`
	Ratifier *RatifierOptions `protobuf:"octets,3,opt,name=validator,proto3" json:"ratifier,omitempty"`
	Release   *ReleaseOptions   `protobuf:"octets,4,opt,name=version,proto3" json:"release,omitempty"`
	Iface      *IfaceOptions      `protobuf:"octets,5,opt,name=abci,proto3" json:"iface,omitempty"`
}

func (m *AgreementOptions) Restore()         { *m = AgreementOptions{} }
func (m *AgreementOptions) String() string { return proto.CompactTextString(m) }
func (*AgreementOptions) SchemaSignal()    {}
func (*AgreementOptions) Definition() ([]byte, []int) {
	return filedefinition_e12598271a686f57, []int{0}
}
func (m *AgreementOptions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AgreementOptions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Agreementparams.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgreementOptions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Agreementparams.Merge(m, src)
}
func (m *AgreementOptions) XXX_Volume() int {
	return m.Volume()
}
func (m *AgreementOptions) XXX_Omitunclear() {
	xxx_messagedata_Agreementparams.DiscardUnknown(m)
}

var xxx_messagedata_Agreementparams proto.InternalMessageInfo

func (m *AgreementOptions) FetchLedger() *LedgerOptions {
	if m != nil {
		return m.Ledger
	}
	return nil
}

func (m *AgreementOptions) FetchProof() *ProofOptions {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *AgreementOptions) FetchRatifier() *RatifierOptions {
	if m != nil {
		return m.Ratifier
	}
	return nil
}

func (m *AgreementOptions) FetchRelease() *ReleaseOptions {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *AgreementOptions) FetchIface() *IfaceOptions {
	if m != nil {
		return m.Iface
	}
	return nil
}

//
type LedgerOptions struct {
	//
	//
	MaximumOctets int64 `protobuf:"variableint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"maximum_octets,omitempty"`
	//
	//
	MaximumFuel int64 `protobuf:"variableint,2,opt,name=max_gas,json=maxGas,proto3" json:"maximum_fuel,omitempty"`
}

func (m *LedgerOptions) Restore()         { *m = LedgerOptions{} }
func (m *LedgerOptions) String() string { return proto.CompactTextString(m) }
func (*LedgerOptions) SchemaSignal()    {}
func (*LedgerOptions) Definition() ([]byte, []int) {
	return filedefinition_e12598271a686f57, []int{1}
}
func (m *LedgerOptions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *LedgerOptions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ledgerparams.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerOptions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ledgerparams.Merge(m, src)
}
func (m *LedgerOptions) XXX_Volume() int {
	return m.Volume()
}
func (m *LedgerOptions) XXX_Omitunclear() {
	xxx_messagedata_Ledgerparams.DiscardUnknown(m)
}

var xxx_messagedata_Ledgerparams proto.InternalMessageInfo

func (m *LedgerOptions) FetchMaximumOctets() int64 {
	if m != nil {
		return m.MaximumOctets
	}
	return 0
}

func (m *LedgerOptions) FetchMaximumFuel() int64 {
	if m != nil {
		return m.MaximumFuel
	}
	return 0
}

//
type ProofOptions struct {
	//
	//
	//
	//
	MaximumDurationCountLedgers int64 `protobuf:"variableint,1,opt,name=max_age_num_blocks,json=maxAgeNumBlocks,proto3" json:"maximum_duration_count_ledgers,omitempty"`
	//
	//
	//
	//
	//
	MaximumDurationPeriod time.Duration `protobuf:"octets,2,opt,name=max_age_duration,json=maxAgeDuration,proto3,stdduration" json:"maximum_duration_period"`
	//
	//
	//
	MaximumOctets int64 `protobuf:"variableint,3,opt,name=max_bytes,json=maxBytes,proto3" json:"maximum_octets,omitempty"`
}

func (m *ProofOptions) Restore()         { *m = ProofOptions{} }
func (m *ProofOptions) String() string { return proto.CompactTextString(m) }
func (*ProofOptions) SchemaSignal()    {}
func (*ProofOptions) Definition() ([]byte, []int) {
	return filedefinition_e12598271a686f57, []int{2}
}
func (m *ProofOptions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ProofOptions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Evidenceparams.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOptions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Evidenceparams.Merge(m, src)
}
func (m *ProofOptions) XXX_Volume() int {
	return m.Volume()
}
func (m *ProofOptions) XXX_Omitunclear() {
	xxx_messagedata_Evidenceparams.DiscardUnknown(m)
}

var xxx_messagedata_Evidenceparams proto.InternalMessageInfo

func (m *ProofOptions) FetchMaximumEraCountLedgers() int64 {
	if m != nil {
		return m.MaximumDurationCountLedgers
	}
	return 0
}

func (m *ProofOptions) FetchMaximumEraPeriod() time.Duration {
	if m != nil {
		return m.MaximumDurationPeriod
	}
	return 0
}

func (m *ProofOptions) FetchMaximumOctets() int64 {
	if m != nil {
		return m.MaximumOctets
	}
	return 0
}

//
//
type RatifierOptions struct {
	PublicKeyKinds []string `protobuf:"octets,1,rep,name=pub_key_types,json=pubKeyTypes,proto3" json:"public_key_kinds,omitempty"`
}

func (m *RatifierOptions) Restore()         { *m = RatifierOptions{} }
func (m *RatifierOptions) String() string { return proto.CompactTextString(m) }
func (*RatifierOptions) SchemaSignal()    {}
func (*RatifierOptions) Definition() ([]byte, []int) {
	return filedefinition_e12598271a686f57, []int{3}
}
func (m *RatifierOptions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *RatifierOptions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ballotparams.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RatifierOptions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ballotparams.Merge(m, src)
}
func (m *RatifierOptions) XXX_Volume() int {
	return m.Volume()
}
func (m *RatifierOptions) XXX_Omitunclear() {
	xxx_messagedata_Ballotparams.DiscardUnknown(m)
}

var xxx_messagedata_Ballotparams proto.InternalMessageInfo

func (m *RatifierOptions) FetchPublicKeyKinds() []string {
	if m != nil {
		return m.PublicKeyKinds
	}
	return nil
}

//
type ReleaseOptions struct {
	App uint64 `protobuf:"variableint,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (m *ReleaseOptions) Restore()         { *m = ReleaseOptions{} }
func (m *ReleaseOptions) String() string { return proto.CompactTextString(m) }
func (*ReleaseOptions) SchemaSignal()    {}
func (*ReleaseOptions) Definition() ([]byte, []int) {
	return filedefinition_e12598271a686f57, []int{4}
}
func (m *ReleaseOptions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ReleaseOptions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Releaseparams.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReleaseOptions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Releaseparams.Merge(m, src)
}
func (m *ReleaseOptions) XXX_Volume() int {
	return m.Volume()
}
func (m *ReleaseOptions) XXX_Omitunclear() {
	xxx_messagedata_Releaseparams.DiscardUnknown(m)
}

var xxx_messagedata_Releaseparams proto.InternalMessageInfo

func (m *ReleaseOptions) FetchApplication() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

//
//
//
type DigestedOptions struct {
	LedgerMaximumOctets int64 `protobuf:"variableint,1,opt,name=block_max_bytes,json=blockMaxBytes,proto3" json:"ledger_maximum_octets,omitempty"`
	LedgerMaximumFuel   int64 `protobuf:"variableint,2,opt,name=block_max_gas,json=blockMaxGas,proto3" json:"ledger_maximum_fuel,omitempty"`
}

func (m *DigestedOptions) Restore()         { *m = DigestedOptions{} }
func (m *DigestedOptions) String() string { return proto.CompactTextString(m) }
func (*DigestedOptions) SchemaSignal()    {}
func (*DigestedOptions) Definition() ([]byte, []int) {
	return filedefinition_e12598271a686f57, []int{5}
}
func (m *DigestedOptions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *DigestedOptions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Digestparams.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DigestedOptions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Digestparams.Merge(m, src)
}
func (m *DigestedOptions) XXX_Volume() int {
	return m.Volume()
}
func (m *DigestedOptions) XXX_Omitunclear() {
	xxx_messagedata_Digestparams.DiscardUnknown(m)
}

var xxx_messagedata_Digestparams proto.InternalMessageInfo

func (m *DigestedOptions) FetchLedgerMaximumOctets() int64 {
	if m != nil {
		return m.LedgerMaximumOctets
	}
	return 0
}

func (m *DigestedOptions) FetchLedgerMaximumFuel() int64 {
	if m != nil {
		return m.LedgerMaximumFuel
	}
	return 0
}

//
type IfaceOptions struct {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	BallotPluginsActivateLevel int64 `protobuf:"variableint,1,opt,name=vote_extensions_enable_height,json=voteExtensionsEnableHeight,proto3" json:"ballot_plugins_activate_level,omitempty"`
}

func (m *IfaceOptions) Restore()         { *m = IfaceOptions{} }
func (m *IfaceOptions) String() string { return proto.CompactTextString(m) }
func (*IfaceOptions) SchemaSignal()    {}
func (*IfaceOptions) Definition() ([]byte, []int) {
	return filedefinition_e12598271a686f57, []int{6}
}
func (m *IfaceOptions) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *IfaceOptions) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Abciparams.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IfaceOptions) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Abciparams.Merge(m, src)
}
func (m *IfaceOptions) XXX_Volume() int {
	return m.Volume()
}
func (m *IfaceOptions) XXX_Omitunclear() {
	xxx_messagedata_Abciparams.DiscardUnknown(m)
}

var xxx_messagedata_Abciparams proto.InternalMessageInfo

func (m *IfaceOptions) FetchBallotPluginsActivateLevel() int64 {
	if m != nil {
		return m.BallotPluginsActivateLevel
	}
	return 0
}

func init() {
	proto.RegisterType((*AgreementOptions)(nil), "REDACTED")
	proto.RegisterType((*LedgerOptions)(nil), "REDACTED")
	proto.RegisterType((*ProofOptions)(nil), "REDACTED")
	proto.RegisterType((*RatifierOptions)(nil), "REDACTED")
	proto.RegisterType((*ReleaseOptions)(nil), "REDACTED")
	proto.RegisterType((*DigestedOptions)(nil), "REDACTED")
	proto.RegisterType((*IfaceOptions)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_e12598271a686f57) }

var filedefinition_e12598271a686f57 = []byte{
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

func (this *AgreementOptions) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AgreementOptions)
	if !ok {
		that2, ok := that.(AgreementOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ledger.Equivalent(that1.Ledger) {
		return false
	}
	if !this.Proof.Equivalent(that1.Proof) {
		return false
	}
	if !this.Ratifier.Equivalent(that1.Ratifier) {
		return false
	}
	if !this.Release.Equivalent(that1.Release) {
		return false
	}
	if !this.Iface.Equivalent(that1.Iface) {
		return false
	}
	return true
}
func (this *LedgerOptions) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LedgerOptions)
	if !ok {
		that2, ok := that.(LedgerOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaximumOctets != that1.MaximumOctets {
		return false
	}
	if this.MaximumFuel != that1.MaximumFuel {
		return false
	}
	return true
}
func (this *ProofOptions) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProofOptions)
	if !ok {
		that2, ok := that.(ProofOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaximumDurationCountLedgers != that1.MaximumDurationCountLedgers {
		return false
	}
	if this.MaximumDurationPeriod != that1.MaximumDurationPeriod {
		return false
	}
	if this.MaximumOctets != that1.MaximumOctets {
		return false
	}
	return true
}
func (this *RatifierOptions) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RatifierOptions)
	if !ok {
		that2, ok := that.(RatifierOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.PublicKeyKinds) != len(that1.PublicKeyKinds) {
		return false
	}
	for i := range this.PublicKeyKinds {
		if this.PublicKeyKinds[i] != that1.PublicKeyKinds[i] {
			return false
		}
	}
	return true
}
func (this *ReleaseOptions) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReleaseOptions)
	if !ok {
		that2, ok := that.(ReleaseOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.App != that1.App {
		return false
	}
	return true
}
func (this *DigestedOptions) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DigestedOptions)
	if !ok {
		that2, ok := that.(DigestedOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.LedgerMaximumOctets != that1.LedgerMaximumOctets {
		return false
	}
	if this.LedgerMaximumFuel != that1.LedgerMaximumFuel {
		return false
	}
	return true
}
func (this *IfaceOptions) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IfaceOptions)
	if !ok {
		that2, ok := that.(IfaceOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BallotPluginsActivateLevel != that1.BallotPluginsActivateLevel {
		return false
	}
	return true
}
func (m *AgreementOptions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgreementOptions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AgreementOptions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Iface != nil {
		{
			volume, err := m.Iface.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintOptions(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Release != nil {
		{
			volume, err := m.Release.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintOptions(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Ratifier != nil {
		{
			volume, err := m.Ratifier.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintOptions(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Proof != nil {
		{
			volume, err := m.Proof.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintOptions(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Ledger != nil {
		{
			volume, err := m.Ledger.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintOptions(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LedgerOptions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LedgerOptions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *LedgerOptions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaximumFuel != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.MaximumFuel))
		i--
		dAtA[i] = 0x10
	}
	if m.MaximumOctets != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.MaximumOctets))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProofOptions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOptions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ProofOptions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaximumOctets != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.MaximumOctets))
		i--
		dAtA[i] = 0x18
	}
	n6, err6 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaximumDurationPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaximumDurationPeriod):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVariableintOptions(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x12
	if m.MaximumDurationCountLedgers != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.MaximumDurationCountLedgers))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RatifierOptions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RatifierOptions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *RatifierOptions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicKeyKinds) > 0 {
		for idxNdEx := len(m.PublicKeyKinds) - 1; idxNdEx >= 0; idxNdEx-- {
			i -= len(m.PublicKeyKinds[idxNdEx])
			copy(dAtA[i:], m.PublicKeyKinds[idxNdEx])
			i = encodeVariableintOptions(dAtA, i, uint64(len(m.PublicKeyKinds[idxNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReleaseOptions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReleaseOptions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ReleaseOptions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.App))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DigestedOptions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DigestedOptions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *DigestedOptions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LedgerMaximumFuel != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.LedgerMaximumFuel))
		i--
		dAtA[i] = 0x10
	}
	if m.LedgerMaximumOctets != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.LedgerMaximumOctets))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IfaceOptions) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IfaceOptions) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *IfaceOptions) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BallotPluginsActivateLevel != 0 {
		i = encodeVariableintOptions(dAtA, i, uint64(m.BallotPluginsActivateLevel))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVariableintOptions(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovOptions(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func NewFilledRatifierOptions(r randyOptions, simple bool) *RatifierOptions {
	this := &RatifierOptions{}
	v1 := r.Intn(10)
	this.PublicKeyKinds = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.PublicKeyKinds[i] = string(randomStringOptions(r))
	}
	if !simple && r.Intn(10) != 0 {
	}
	return this
}

func NewFilledReleaseOptions(r randyOptions, simple bool) *ReleaseOptions {
	this := &ReleaseOptions{}
	this.App = uint64(uint64(r.Uint32()))
	if !simple && r.Intn(10) != 0 {
	}
	return this
}

type randyOptions interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randomUtf8runeOptions(r randyOptions) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randomStringOptions(r randyOptions) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randomUtf8runeOptions(r)
	}
	return string(tmps)
}
func randomUnknownOptions(r randyOptions, maximumFieldCount int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		cable := r.Intn(4)
		if cable == 3 {
			cable = 5
		}
		fieldCount := maximumFieldCount + r.Intn(100)
		dAtA = randomFieldOptions(dAtA, r, fieldCount, cable)
	}
	return dAtA
}
func randomFieldOptions(dAtA []byte, r randyOptions, fieldCount int, cable int) []byte {
	key := uint32(fieldCount)<<3 | uint32(cable)
	switch cable {
	case 0:
		dAtA = encodeVariableintFillOptions(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVariableintFillOptions(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVariableintFillOptions(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVariableintFillOptions(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVariableintFillOptions(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVariableintFillOptions(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVariableintFillOptions(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AgreementOptions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ledger != nil {
		l = m.Ledger.Volume()
		n += 1 + l + sovOptions(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Volume()
		n += 1 + l + sovOptions(uint64(l))
	}
	if m.Ratifier != nil {
		l = m.Ratifier.Volume()
		n += 1 + l + sovOptions(uint64(l))
	}
	if m.Release != nil {
		l = m.Release.Volume()
		n += 1 + l + sovOptions(uint64(l))
	}
	if m.Iface != nil {
		l = m.Iface.Volume()
		n += 1 + l + sovOptions(uint64(l))
	}
	return n
}

func (m *LedgerOptions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumOctets != 0 {
		n += 1 + sovOptions(uint64(m.MaximumOctets))
	}
	if m.MaximumFuel != 0 {
		n += 1 + sovOptions(uint64(m.MaximumFuel))
	}
	return n
}

func (m *ProofOptions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumDurationCountLedgers != 0 {
		n += 1 + sovOptions(uint64(m.MaximumDurationCountLedgers))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaximumDurationPeriod)
	n += 1 + l + sovOptions(uint64(l))
	if m.MaximumOctets != 0 {
		n += 1 + sovOptions(uint64(m.MaximumOctets))
	}
	return n
}

func (m *RatifierOptions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PublicKeyKinds) > 0 {
		for _, s := range m.PublicKeyKinds {
			l = len(s)
			n += 1 + l + sovOptions(uint64(l))
		}
	}
	return n
}

func (m *ReleaseOptions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.App != 0 {
		n += 1 + sovOptions(uint64(m.App))
	}
	return n
}

func (m *DigestedOptions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerMaximumOctets != 0 {
		n += 1 + sovOptions(uint64(m.LedgerMaximumOctets))
	}
	if m.LedgerMaximumFuel != 0 {
		n += 1 + sovOptions(uint64(m.LedgerMaximumFuel))
	}
	return n
}

func (m *IfaceOptions) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BallotPluginsActivateLevel != 0 {
		n += 1 + sovOptions(uint64(m.BallotPluginsActivateLevel))
	}
	return n
}

func sovOptions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOptions(x uint64) (n int) {
	return sovOptions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AgreementOptions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadOptions
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
					return ErrIntegerOverloadOptions
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
				return ErrCorruptExtentOptions
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentOptions
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ledger == nil {
				m.Ledger = &LedgerOptions{}
			}
			if err := m.Ledger.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
					return ErrIntegerOverloadOptions
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
				return ErrCorruptExtentOptions
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentOptions
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &ProofOptions{}
			}
			if err := m.Proof.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
					return ErrIntegerOverloadOptions
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
				return ErrCorruptExtentOptions
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentOptions
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ratifier == nil {
				m.Ratifier = &RatifierOptions{}
			}
			if err := m.Ratifier.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
					return ErrIntegerOverloadOptions
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
				return ErrCorruptExtentOptions
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentOptions
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Release == nil {
				m.Release = &ReleaseOptions{}
			}
			if err := m.Release.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
					return ErrIntegerOverloadOptions
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
				return ErrCorruptExtentOptions
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentOptions
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Iface == nil {
				m.Iface = &IfaceOptions{}
			}
			if err := m.Iface.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitOptions(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentOptions
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
func (m *LedgerOptions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadOptions
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
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MaximumOctets |= int64(b&0x7F) << displace
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
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MaximumFuel |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitOptions(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentOptions
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
func (m *ProofOptions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadOptions
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
			m.MaximumDurationCountLedgers = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MaximumDurationCountLedgers |= int64(b&0x7F) << displace
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
					return ErrIntegerOverloadOptions
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
				return ErrCorruptExtentOptions
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentOptions
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaximumDurationPeriod, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.MaximumOctets = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.MaximumOctets |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitOptions(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentOptions
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
func (m *RatifierOptions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadOptions
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
					return ErrIntegerOverloadOptions
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
				return ErrCorruptExtentOptions
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentOptions
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeyKinds = append(m.PublicKeyKinds, string(dAtA[idxNdEx:submitOrdinal]))
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitOptions(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentOptions
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
func (m *ReleaseOptions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadOptions
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
			m.App = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.App |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitOptions(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentOptions
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
func (m *DigestedOptions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadOptions
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
			m.LedgerMaximumOctets = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.LedgerMaximumOctets |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.LedgerMaximumFuel = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.LedgerMaximumFuel |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitOptions(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentOptions
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
func (m *IfaceOptions) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadOptions
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
			m.BallotPluginsActivateLevel = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadOptions
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.BallotPluginsActivateLevel |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitOptions(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentOptions
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
func omitOptions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadOptions
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
					return 0, ErrIntegerOverloadOptions
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
					return 0, ErrIntegerOverloadOptions
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
				return 0, ErrCorruptExtentOptions
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterOptions
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentOptions
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentOptions        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadOptions          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterOptions = fmt.Errorf("REDACTED")
)
