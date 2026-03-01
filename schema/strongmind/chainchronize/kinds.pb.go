//
//

package chainchronize

import (
	fmt "fmt"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
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

//
type LedgerSolicit struct {
	Altitude int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
}

func (m *LedgerSolicit) Restore()         { *m = LedgerSolicit{} }
func (m *LedgerSolicit) Text() string { return proto.CompactTextString(m) }
func (*LedgerSolicit) SchemaArtifact()    {}
func (*LedgerSolicit) Definition() ([]byte, []int) {
	return filedescriptor_19b397c236e0fa07, []int{0}
}
func (m *LedgerSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *LedgerSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ledgerrequest.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ledgerrequest.Merge(m, src)
}
func (m *LedgerSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *LedgerSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ledgerrequest.DiscardUnknown(m)
}

var xxx_signaldetails_Ledgerrequest proto.InternalMessageInfo

func (m *LedgerSolicit) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

//
type NegativeLedgerReply struct {
	Altitude int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
}

func (m *NegativeLedgerReply) Restore()         { *m = NegativeLedgerReply{} }
func (m *NegativeLedgerReply) Text() string { return proto.CompactTextString(m) }
func (*NegativeLedgerReply) SchemaArtifact()    {}
func (*NegativeLedgerReply) Definition() ([]byte, []int) {
	return filedescriptor_19b397c236e0fa07, []int{1}
}
func (m *NegativeLedgerReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *NegativeLedgerReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Noledgerreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NegativeLedgerReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Noledgerreply.Merge(m, src)
}
func (m *NegativeLedgerReply) XXX_Extent() int {
	return m.Extent()
}
func (m *NegativeLedgerReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Noledgerreply.DiscardUnknown(m)
}

var xxx_signaldetails_Noledgerreply proto.InternalMessageInfo

func (m *NegativeLedgerReply) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

//
type LedgerReply struct {
	Ledger     *kinds.Ledger          `protobuf:"octets,1,opt,name=block,proto3" json:"ledger,omitempty"`
	AddnEndorse *kinds.ExpandedEndorse `protobuf:"octets,2,opt,name=ext_commit,json=extCommit,proto3" json:"addn_endorse,omitempty"`
}

func (m *LedgerReply) Restore()         { *m = LedgerReply{} }
func (m *LedgerReply) Text() string { return proto.CompactTextString(m) }
func (*LedgerReply) SchemaArtifact()    {}
func (*LedgerReply) Definition() ([]byte, []int) {
	return filedescriptor_19b397c236e0fa07, []int{2}
}
func (m *LedgerReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *LedgerReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ledgerreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ledgerreply.Merge(m, src)
}
func (m *LedgerReply) XXX_Extent() int {
	return m.Extent()
}
func (m *LedgerReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ledgerreply.DiscardUnknown(m)
}

var xxx_signaldetails_Ledgerreply proto.InternalMessageInfo

func (m *LedgerReply) ObtainLedger() *kinds.Ledger {
	if m != nil {
		return m.Ledger
	}
	return nil
}

func (m *LedgerReply) ObtainAddnEndorse() *kinds.ExpandedEndorse {
	if m != nil {
		return m.AddnEndorse
	}
	return nil
}

//
type ConditionSolicit struct {
}

func (m *ConditionSolicit) Restore()         { *m = ConditionSolicit{} }
func (m *ConditionSolicit) Text() string { return proto.CompactTextString(m) }
func (*ConditionSolicit) SchemaArtifact()    {}
func (*ConditionSolicit) Definition() ([]byte, []int) {
	return filedescriptor_19b397c236e0fa07, []int{3}
}
func (m *ConditionSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ConditionSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Conditionrequest.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConditionSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Conditionrequest.Merge(m, src)
}
func (m *ConditionSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *ConditionSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Conditionrequest.DiscardUnknown(m)
}

var xxx_signaldetails_Conditionrequest proto.InternalMessageInfo

//
type ConditionReply struct {
	Altitude int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Foundation   int64 `protobuf:"variableint,2,opt,name=base,proto3" json:"foundation,omitempty"`
}

func (m *ConditionReply) Restore()         { *m = ConditionReply{} }
func (m *ConditionReply) Text() string { return proto.CompactTextString(m) }
func (*ConditionReply) SchemaArtifact()    {}
func (*ConditionReply) Definition() ([]byte, []int) {
	return filedescriptor_19b397c236e0fa07, []int{4}
}
func (m *ConditionReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ConditionReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Conditionreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConditionReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Conditionreply.Merge(m, src)
}
func (m *ConditionReply) XXX_Extent() int {
	return m.Extent()
}
func (m *ConditionReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Conditionreply.DiscardUnknown(m)
}

var xxx_signaldetails_Conditionreply proto.InternalMessageInfo

func (m *ConditionReply) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *ConditionReply) ObtainFoundation() int64 {
	if m != nil {
		return m.Foundation
	}
	return 0
}

type Signal struct {
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
	return filedescriptor_19b397c236e0fa07, []int{5}
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

type Signal_Ledgerrequest struct {
	LedgerSolicit *LedgerSolicit `protobuf:"octets,1,opt,name=block_request,json=blockRequest,proto3,oneof" json:"ledger_solicit,omitempty"`
}
type Signal_Noledgerreply struct {
	NegativeLedgerReply *NegativeLedgerReply `protobuf:"octets,2,opt,name=no_block_response,json=noBlockResponse,proto3,oneof" json:"negative_ledger_reply,omitempty"`
}
type Signal_Ledgerreply struct {
	LedgerReply *LedgerReply `protobuf:"octets,3,opt,name=block_response,json=blockResponse,proto3,oneof" json:"ledger_reply,omitempty"`
}
type Signal_Conditionrequest struct {
	ConditionSolicit *ConditionSolicit `protobuf:"octets,4,opt,name=status_request,json=statusRequest,proto3,oneof" json:"condition_solicit,omitempty"`
}
type Signal_Conditionreply struct {
	ConditionReply *ConditionReply `protobuf:"octets,5,opt,name=status_response,json=statusResponse,proto3,oneof" json:"condition_reply,omitempty"`
}

func (*Signal_Ledgerrequest) isnote_Total()    {}
func (*Signal_Noledgerreply) isnote_Total() {}
func (*Signal_Ledgerreply) isnote_Total()   {}
func (*Signal_Conditionrequest) isnote_Total()   {}
func (*Signal_Conditionreply) isnote_Total()  {}

func (m *Signal) ObtainTotal() isnote_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) ObtainLedgerSolicit() *LedgerSolicit {
	if x, ok := m.ObtainTotal().(*Signal_Ledgerrequest); ok {
		return x.LedgerSolicit
	}
	return nil
}

func (m *Signal) ObtainNegativeLedgerReply() *NegativeLedgerReply {
	if x, ok := m.ObtainTotal().(*Signal_Noledgerreply); ok {
		return x.NegativeLedgerReply
	}
	return nil
}

func (m *Signal) ObtainLedgerReply() *LedgerReply {
	if x, ok := m.ObtainTotal().(*Signal_Ledgerreply); ok {
		return x.LedgerReply
	}
	return nil
}

func (m *Signal) ObtainConditionSolicit() *ConditionSolicit {
	if x, ok := m.ObtainTotal().(*Signal_Conditionrequest); ok {
		return x.ConditionSolicit
	}
	return nil
}

func (m *Signal) ObtainConditionReply() *ConditionReply {
	if x, ok := m.ObtainTotal().(*Signal_Conditionreply); ok {
		return x.ConditionReply
	}
	return nil
}

//
func (*Signal) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Signal_Ledgerrequest)(nil),
		(*Signal_Noledgerreply)(nil),
		(*Signal_Ledgerreply)(nil),
		(*Signal_Conditionrequest)(nil),
		(*Signal_Conditionreply)(nil),
	}
}

func initialize() {
	proto.RegisterType((*LedgerSolicit)(nil), "REDACTED")
	proto.RegisterType((*NegativeLedgerReply)(nil), "REDACTED")
	proto.RegisterType((*LedgerReply)(nil), "REDACTED")
	proto.RegisterType((*ConditionSolicit)(nil), "REDACTED")
	proto.RegisterType((*ConditionReply)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_19b397c236e0fa07) }

var filedescriptor_19b397c236e0fa07 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x13, 0xd3, 0x56, 0xbc, 0x36, 0x0d, 0x06, 0xd1, 0x22, 0x12, 0x4a, 0xfc, 0x41, 0x17,
	0x26, 0xa0, 0x0b, 0x37, 0x82, 0x50, 0x11, 0x2a, 0xf8, 0x83, 0xe9, 0xce, 0x4d, 0xe9, 0xa4, 0x63,
	0x1b, 0x34, 0x99, 0xda, 0x99, 0x40, 0xbb, 0xf2, 0x15, 0x7c, 0x01, 0xdf, 0xc7, 0x65, 0x97, 0x2e,
	0xa5, 0x7d, 0x11, 0xe9, 0x4c, 0x9a, 0xa6, 0x31, 0x66, 0x37, 0xb9, 0x73, 0xee, 0x97, 0x73, 0xee,
	0x65, 0xa0, 0xc6, 0x70, 0xd0, 0xc1, 0x03, 0xdf, 0x0b, 0x98, 0x8d, 0x5e, 0x89, 0xfb, 0x42, 0x47,
	0x81, 0x6b, 0xb3, 0x51, 0x1f, 0x53, 0xab, 0x3f, 0x20, 0x8c, 0xe8, 0x9b, 0x0b, 0x85, 0x15, 0x2b,
	0x76, 0x76, 0x13, 0x7d, 0x5c, 0x2d, 0xba, 0x45, 0x4f, 0xc6, 0x6d, 0x82, 0x68, 0x1e, 0x42, 0xb9,
	0x3e, 0x13, 0x3b, 0xf8, 0x2d, 0xc4, 0x94, 0xe9, 0x5b, 0x50, 0xea, 0x61, 0xaf, 0xdb, 0x63, 0x55,
	0xb9, 0x26, 0x1f, 0x29, 0x4e, 0xf4, 0x65, 0x1e, 0x83, 0x76, 0x4f, 0x22, 0x25, 0xed, 0x93, 0x80,
	0xe2, 0x7f, 0xa5, 0xef, 0xa0, 0x2e, 0x0b, 0x4f, 0xa0, 0xc8, 0x0d, 0x71, 0xdd, 0xfa, 0xe9, 0xb6,
	0x95, 0x48, 0x21, 0xbc, 0x08, 0xbd, 0x50, 0xe9, 0x97, 0x00, 0x78, 0xc8, 0x5a, 0x2e, 0xf1, 0x7d,
	0x8f, 0x55, 0x57, 0x78, 0x4f, 0xed, 0x6f, 0xcf, 0xf5, 0x90, 0x97, 0x3a, 0x57, 0x5c, 0xe7, 0xac,
	0xe1, 0x21, 0x13, 0x47, 0x53, 0x03, 0xb5, 0xc9, 0xda, 0x2c, 0xa4, 0x51, 0x28, 0xf3, 0x02, 0x2a,
	0xf3, 0x42, 0xbe, 0x77, 0x5d, 0x87, 0x02, 0x6a, 0x53, 0xcc, 0xff, 0xaa, 0x38, 0xfc, 0x6c, 0x7e,
	0x2a, 0xb0, 0x7a, 0x87, 0x29, 0x6d, 0x77, 0xb1, 0x7e, 0x03, 0x2a, 0x37, 0xd9, 0x1a, 0x08, 0x74,
	0x14, 0xc9, 0xb4, 0xb2, 0x16, 0x63, 0x25, 0x27, 0xdb, 0x90, 0x9c, 0x32, 0x4a, 0x4e, 0xba, 0x09,
	0x1b, 0x01, 0x69, 0xcd, 0x69, 0xc2, 0x57, 0x94, 0xf6, 0x20, 0x1b, 0x97, 0x5a, 0x40, 0x43, 0x72,
	0xb4, 0x20, 0xb5, 0x93, 0x5b, 0xa8, 0xa4, 0x88, 0x0a, 0x27, 0xee, 0xe5, 0x1a, 0x8c, 0x79, 0x2a,
	0x4a, 0xd3, 0x28, 0x9f, 0x5b, 0x1c, 0xb7, 0x90, 0x47, 0x5b, 0x1a, 0xfa, 0x8c, 0x46, 0x93, 0x05,
	0xfd, 0x01, 0xb4, 0x98, 0x16, 0x99, 0x2b, 0x72, 0xdc, 0x7e, 0x3e, 0x2e, 0x76, 0x57, 0xa1, 0x4b,
	0x95, 0x7a, 0x11, 0x14, 0x1a, 0xfa, 0xf5, 0xc7, 0xaf, 0x89, 0x21, 0x8f, 0x27, 0x86, 0xfc, 0x33,
	0x31, 0xe4, 0x8f, 0xa9, 0x21, 0x8d, 0xa7, 0x86, 0xf4, 0x3d, 0x35, 0xa4, 0xa7, 0xf3, 0xae, 0xc7,
	0x7a, 0x21, 0xb2, 0x5c, 0xe2, 0xdb, 0x2e, 0xf1, 0x31, 0x43, 0xcf, 0x6c, 0x71, 0xe0, 0x0f, 0xc0,
	0xce, 0x7a, 0x73, 0xa8, 0xc4, 0xef, 0xce, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x56, 0x8a, 0x71,
	0xcf, 0x92, 0x03, 0x00, 0x00,
}

func (m *LedgerSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *LedgerSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *LedgerSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *NegativeLedgerReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *NegativeLedgerReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *NegativeLedgerReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *LedgerReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *LedgerReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *LedgerReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.AddnEndorse != nil {
		{
			extent, err := m.AddnEndorse.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
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
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *ConditionSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ConditionSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ConditionSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *ConditionReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ConditionReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ConditionReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Foundation != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Foundation))
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

func (m *Signal_Ledgerrequest) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Ledgerrequest) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.LedgerSolicit != nil {
		{
			extent, err := m.LedgerSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Signal_Noledgerreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Noledgerreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.NegativeLedgerReply != nil {
		{
			extent, err := m.NegativeLedgerReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Signal_Ledgerreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Ledgerreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.LedgerReply != nil {
		{
			extent, err := m.LedgerReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Signal_Conditionrequest) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Conditionrequest) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ConditionSolicit != nil {
		{
			extent, err := m.ConditionSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Signal_Conditionreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal_Conditionreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ConditionReply != nil {
		{
			extent, err := m.ConditionReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *LedgerSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	return n
}

func (m *NegativeLedgerReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	return n
}

func (m *LedgerReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ledger != nil {
		l = m.Ledger.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.AddnEndorse != nil {
		l = m.AddnEndorse.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *ConditionSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ConditionReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Foundation != 0 {
		n += 1 + sovKinds(uint64(m.Foundation))
	}
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

func (m *Signal_Ledgerrequest) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerSolicit != nil {
		l = m.LedgerSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Noledgerreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NegativeLedgerReply != nil {
		l = m.NegativeLedgerReply.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Ledgerreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerReply != nil {
		l = m.LedgerReply.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Conditionrequest) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConditionSolicit != nil {
		l = m.ConditionSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Conditionreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConditionReply != nil {
		l = m.ConditionReply.Extent()
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
func (m *LedgerSolicit) Decode(deltaLocatedAN []byte) error {
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
func (m *NegativeLedgerReply) Decode(deltaLocatedAN []byte) error {
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
func (m *LedgerReply) Decode(deltaLocatedAN []byte) error {
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
			if m.Ledger == nil {
				m.Ledger = &kinds.Ledger{}
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
			if m.AddnEndorse == nil {
				m.AddnEndorse = &kinds.ExpandedEndorse{}
			}
			if err := m.AddnEndorse.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
func (m *ConditionSolicit) Decode(deltaLocatedAN []byte) error {
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
func (m *ConditionReply) Decode(deltaLocatedAN []byte) error {
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
			m.Foundation = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Foundation |= int64(b&0x7F) << relocate
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
			v := &LedgerSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ledgerrequest{v}
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
			v := &NegativeLedgerReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Noledgerreply{v}
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
			v := &LedgerReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Ledgerreply{v}
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
			v := &ConditionSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Conditionrequest{v}
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
			v := &ConditionReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Conditionreply{v}
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
