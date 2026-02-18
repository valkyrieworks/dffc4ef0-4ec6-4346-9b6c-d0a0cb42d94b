//
//

package chainconnect

import (
	fmt "fmt"
	kinds "github.com/valkyrieworks/schema/consensuscore/kinds"
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
type LedgerQuery struct {
	Level int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
}

func (m *LedgerQuery) Restore()         { *m = LedgerQuery{} }
func (m *LedgerQuery) String() string { return proto.CompactTextString(m) }
func (*LedgerQuery) SchemaSignal()    {}
func (*LedgerQuery) Definition() ([]byte, []int) {
	return filedefinition_19b397c236e0fa07, []int{0}
}
func (m *LedgerQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *LedgerQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Chainrequest.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Chainrequest.Merge(m, src)
}
func (m *LedgerQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *LedgerQuery) XXX_Omitunclear() {
	xxx_messagedata_Chainrequest.DiscardUnknown(m)
}

var xxx_messagedata_Chainrequest proto.InternalMessageInfo

func (m *LedgerQuery) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

//
type NoLedgerReply struct {
	Level int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
}

func (m *NoLedgerReply) Restore()         { *m = NoLedgerReply{} }
func (m *NoLedgerReply) String() string { return proto.CompactTextString(m) }
func (*NoLedgerReply) SchemaSignal()    {}
func (*NoLedgerReply) Definition() ([]byte, []int) {
	return filedefinition_19b397c236e0fa07, []int{1}
}
func (m *NoLedgerReply) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *NoLedgerReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Nonledgerreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoLedgerReply) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Nonledgerreply.Merge(m, src)
}
func (m *NoLedgerReply) XXX_Volume() int {
	return m.Volume()
}
func (m *NoLedgerReply) XXX_Omitunclear() {
	xxx_messagedata_Nonledgerreply.DiscardUnknown(m)
}

var xxx_messagedata_Nonledgerreply proto.InternalMessageInfo

func (m *NoLedgerReply) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

//
type LedgerReply struct {
	Ledger     *kinds.Ledger          `protobuf:"octets,1,opt,name=block,proto3" json:"ledger,omitempty"`
	ExtensionEndorse *kinds.ExpandedEndorse `protobuf:"octets,2,opt,name=ext_commit,json=extCommit,proto3" json:"extension_endorse,omitempty"`
}

func (m *LedgerReply) Restore()         { *m = LedgerReply{} }
func (m *LedgerReply) String() string { return proto.CompactTextString(m) }
func (*LedgerReply) SchemaSignal()    {}
func (*LedgerReply) Definition() ([]byte, []int) {
	return filedefinition_19b397c236e0fa07, []int{2}
}
func (m *LedgerReply) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *LedgerReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Chainreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerReply) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Chainreply.Merge(m, src)
}
func (m *LedgerReply) XXX_Volume() int {
	return m.Volume()
}
func (m *LedgerReply) XXX_Omitunclear() {
	xxx_messagedata_Chainreply.DiscardUnknown(m)
}

var xxx_messagedata_Chainreply proto.InternalMessageInfo

func (m *LedgerReply) FetchLedger() *kinds.Ledger {
	if m != nil {
		return m.Ledger
	}
	return nil
}

func (m *LedgerReply) FetchExtensionEndorse() *kinds.ExpandedEndorse {
	if m != nil {
		return m.ExtensionEndorse
	}
	return nil
}

//
type StatusQuery struct {
}

func (m *StatusQuery) Restore()         { *m = StatusQuery{} }
func (m *StatusQuery) String() string { return proto.CompactTextString(m) }
func (*StatusQuery) SchemaSignal()    {}
func (*StatusQuery) Definition() ([]byte, []int) {
	return filedefinition_19b397c236e0fa07, []int{3}
}
func (m *StatusQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StatusQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Staterequest.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Staterequest.Merge(m, src)
}
func (m *StatusQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *StatusQuery) XXX_Omitunclear() {
	xxx_messagedata_Staterequest.DiscardUnknown(m)
}

var xxx_messagedata_Staterequest proto.InternalMessageInfo

//
type StatusReply struct {
	Level int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Root   int64 `protobuf:"variableint,2,opt,name=base,proto3" json:"root,omitempty"`
}

func (m *StatusReply) Restore()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) SchemaSignal()    {}
func (*StatusReply) Definition() ([]byte, []int) {
	return filedefinition_19b397c236e0fa07, []int{4}
}
func (m *StatusReply) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *StatusReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Statereply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusReply) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Statereply.Merge(m, src)
}
func (m *StatusReply) XXX_Volume() int {
	return m.Volume()
}
func (m *StatusReply) XXX_Omitunclear() {
	xxx_messagedata_Statereply.DiscardUnknown(m)
}

var xxx_messagedata_Statereply proto.InternalMessageInfo

func (m *StatusReply) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *StatusReply) FetchRoot() int64 {
	if m != nil {
		return m.Root
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
	Sum ismessage_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) SchemaSignal()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedefinition_19b397c236e0fa07, []int{5}
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

type Signal_Chainrequest struct {
	LedgerQuery *LedgerQuery `protobuf:"octets,1,opt,name=block_request,json=blockRequest,proto3,oneof" json:"ledger_query,omitempty"`
}
type Signal_Nonledgerreply struct {
	NoLedgerReply *NoLedgerReply `protobuf:"octets,2,opt,name=no_block_response,json=noBlockResponse,proto3,oneof" json:"no_ledger_reply,omitempty"`
}
type Signal_Chainreply struct {
	LedgerReply *LedgerReply `protobuf:"octets,3,opt,name=block_response,json=blockResponse,proto3,oneof" json:"ledger_reply,omitempty"`
}
type Signal_Staterequest struct {
	StatusQuery *StatusQuery `protobuf:"octets,4,opt,name=status_request,json=statusRequest,proto3,oneof" json:"state_query,omitempty"`
}
type Signal_Statereply struct {
	StatusReply *StatusReply `protobuf:"octets,5,opt,name=status_response,json=statusResponse,proto3,oneof" json:"state_reply,omitempty"`
}

func (*Signal_Chainrequest) ismessage_Total()    {}
func (*Signal_Nonledgerreply) ismessage_Total() {}
func (*Signal_Chainreply) ismessage_Total()   {}
func (*Signal_Staterequest) ismessage_Total()   {}
func (*Signal_Statereply) ismessage_Total()  {}

func (m *Signal) FetchTotal() ismessage_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) FetchLedgerQuery() *LedgerQuery {
	if x, ok := m.FetchTotal().(*Signal_Chainrequest); ok {
		return x.LedgerQuery
	}
	return nil
}

func (m *Signal) FetchNoLedgerReply() *NoLedgerReply {
	if x, ok := m.FetchTotal().(*Signal_Nonledgerreply); ok {
		return x.NoLedgerReply
	}
	return nil
}

func (m *Signal) FetchLedgerReply() *LedgerReply {
	if x, ok := m.FetchTotal().(*Signal_Chainreply); ok {
		return x.LedgerReply
	}
	return nil
}

func (m *Signal) FetchStateQuery() *StatusQuery {
	if x, ok := m.FetchTotal().(*Signal_Staterequest); ok {
		return x.StatusQuery
	}
	return nil
}

func (m *Signal) FetchStateReply() *StatusReply {
	if x, ok := m.FetchTotal().(*Signal_Statereply); ok {
		return x.StatusReply
	}
	return nil
}

//
func (*Signal) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Signal_Chainrequest)(nil),
		(*Signal_Nonledgerreply)(nil),
		(*Signal_Chainreply)(nil),
		(*Signal_Staterequest)(nil),
		(*Signal_Statereply)(nil),
	}
}

func init() {
	proto.RegisterType((*LedgerQuery)(nil), "REDACTED")
	proto.RegisterType((*NoLedgerReply)(nil), "REDACTED")
	proto.RegisterType((*LedgerReply)(nil), "REDACTED")
	proto.RegisterType((*StatusQuery)(nil), "REDACTED")
	proto.RegisterType((*StatusReply)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_19b397c236e0fa07) }

var filedefinition_19b397c236e0fa07 = []byte{
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

func (m *LedgerQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LedgerQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *LedgerQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NoLedgerReply) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoLedgerReply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *NoLedgerReply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LedgerReply) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LedgerReply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *LedgerReply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExtensionEndorse != nil {
		{
			volume, err := m.ExtensionEndorse.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
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
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatusQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StatusQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *StatusReply) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusReply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *StatusReply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Root != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Root))
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

func (m *Signal_Chainrequest) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Chainrequest) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LedgerQuery != nil {
		{
			volume, err := m.LedgerQuery.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Nonledgerreply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Nonledgerreply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoLedgerReply != nil {
		{
			volume, err := m.NoLedgerReply.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Chainreply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Chainreply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LedgerReply != nil {
		{
			volume, err := m.LedgerReply.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Staterequest) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Staterequest) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StatusQuery != nil {
		{
			volume, err := m.StatusQuery.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *Signal_Statereply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Statereply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StatusReply != nil {
		{
			volume, err := m.StatusReply.SerializeToDimensionedBuffer(dAtA[:i])
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
func (m *LedgerQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	return n
}

func (m *NoLedgerReply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	return n
}

func (m *LedgerReply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ledger != nil {
		l = m.Ledger.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.ExtensionEndorse != nil {
		l = m.ExtensionEndorse.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *StatusQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *StatusReply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Root != 0 {
		n += 1 + sovKinds(uint64(m.Root))
	}
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

func (m *Signal_Chainrequest) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerQuery != nil {
		l = m.LedgerQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Nonledgerreply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoLedgerReply != nil {
		l = m.NoLedgerReply.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Chainreply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LedgerReply != nil {
		l = m.LedgerReply.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Staterequest) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StatusQuery != nil {
		l = m.StatusQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Statereply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StatusReply != nil {
		l = m.StatusReply.Volume()
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
func (m *LedgerQuery) Unserialize(dAtA []byte) error {
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
func (m *NoLedgerReply) Unserialize(dAtA []byte) error {
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
func (m *LedgerReply) Unserialize(dAtA []byte) error {
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
			if m.Ledger == nil {
				m.Ledger = &kinds.Ledger{}
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
			if m.ExtensionEndorse == nil {
				m.ExtensionEndorse = &kinds.ExpandedEndorse{}
			}
			if err := m.ExtensionEndorse.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
func (m *StatusQuery) Unserialize(dAtA []byte) error {
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
func (m *StatusReply) Unserialize(dAtA []byte) error {
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
			m.Root = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Root |= int64(b&0x7F) << displace
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
			v := &LedgerQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Chainrequest{v}
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
			v := &NoLedgerReply{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Nonledgerreply{v}
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
			v := &LedgerReply{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Chainreply{v}
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
			v := &StatusQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Staterequest{v}
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
			v := &StatusReply{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Statereply{v}
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
