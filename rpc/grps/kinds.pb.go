//
//

package coregrpc

import (
	context "context"
	fmt "fmt"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SolicitPing struct {
}

func (m *SolicitPing) Restore()         { *m = SolicitPing{} }
func (m *SolicitPing) Text() string { return proto.CompactTextString(m) }
func (*SolicitPing) SchemaArtifact()    {}
func (*SolicitPing) Definition() ([]byte, []int) {
	return filedescriptor_0ffff5682c662b95, []int{0}
}
func (m *SolicitPing) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitPing) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Solicitping.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitPing) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Solicitping.Merge(m, src)
}
func (m *SolicitPing) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitPing) XXX_Dropunfamiliar() {
	xxx_signaldetails_Solicitping.DiscardUnknown(m)
}

var xxx_signaldetails_Solicitping proto.InternalMessageInfo

type SolicitMulticastTransfer struct {
	Tx []byte `protobuf:"octets,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *SolicitMulticastTransfer) Restore()         { *m = SolicitMulticastTransfer{} }
func (m *SolicitMulticastTransfer) Text() string { return proto.CompactTextString(m) }
func (*SolicitMulticastTransfer) SchemaArtifact()    {}
func (*SolicitMulticastTransfer) Definition() ([]byte, []int) {
	return filedescriptor_0ffff5682c662b95, []int{1}
}
func (m *SolicitMulticastTransfer) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SolicitMulticastTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Solicitmulticasttrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolicitMulticastTransfer) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Solicitmulticasttrans.Merge(m, src)
}
func (m *SolicitMulticastTransfer) XXX_Extent() int {
	return m.Extent()
}
func (m *SolicitMulticastTransfer) XXX_Dropunfamiliar() {
	xxx_signaldetails_Solicitmulticasttrans.DiscardUnknown(m)
}

var xxx_signaldetails_Solicitmulticasttrans proto.InternalMessageInfo

func (m *SolicitMulticastTransfer) ObtainTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type ReplyPing struct {
}

func (m *ReplyPing) Restore()         { *m = ReplyPing{} }
func (m *ReplyPing) Text() string { return proto.CompactTextString(m) }
func (*ReplyPing) SchemaArtifact()    {}
func (*ReplyPing) Definition() ([]byte, []int) {
	return filedescriptor_0ffff5682c662b95, []int{2}
}
func (m *ReplyPing) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyPing) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replyping.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyPing) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replyping.Merge(m, src)
}
func (m *ReplyPing) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyPing) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replyping.DiscardUnknown(m)
}

var xxx_signaldetails_Replyping proto.InternalMessageInfo

type ReplyMulticastTransfer struct {
	InspectTransfer  *kinds.ReplyInspectTransfer `protobuf:"octets,1,opt,name=check_tx,json=checkTx,proto3" json:"inspect_transfer,omitempty"`
	TransferOutcome *kinds.InvokeTransferOutcome    `protobuf:"octets,2,opt,name=tx_result,json=txResult,proto3" json:"transfer_outcome,omitempty"`
}

func (m *ReplyMulticastTransfer) Restore()         { *m = ReplyMulticastTransfer{} }
func (m *ReplyMulticastTransfer) Text() string { return proto.CompactTextString(m) }
func (*ReplyMulticastTransfer) SchemaArtifact()    {}
func (*ReplyMulticastTransfer) Definition() ([]byte, []int) {
	return filedescriptor_0ffff5682c662b95, []int{3}
}
func (m *ReplyMulticastTransfer) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ReplyMulticastTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Replymulticasttrans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyMulticastTransfer) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Replymulticasttrans.Merge(m, src)
}
func (m *ReplyMulticastTransfer) XXX_Extent() int {
	return m.Extent()
}
func (m *ReplyMulticastTransfer) XXX_Dropunfamiliar() {
	xxx_signaldetails_Replymulticasttrans.DiscardUnknown(m)
}

var xxx_signaldetails_Replymulticasttrans proto.InternalMessageInfo

func (m *ReplyMulticastTransfer) ObtainInspectTransfer() *kinds.ReplyInspectTransfer {
	if m != nil {
		return m.InspectTransfer
	}
	return nil
}

func (m *ReplyMulticastTransfer) ObtainTransferOutcome() *kinds.InvokeTransferOutcome {
	if m != nil {
		return m.TransferOutcome
	}
	return nil
}

func initialize() {
	proto.RegisterType((*SolicitPing)(nil), "REDACTED")
	proto.RegisterType((*SolicitMulticastTransfer)(nil), "REDACTED")
	proto.RegisterType((*ReplyPing)(nil), "REDACTED")
	proto.RegisterType((*ReplyMulticastTransfer)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_0ffff5682c662b95) }

var filedescriptor_0ffff5682c662b95 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x31, 0x4f, 0x02, 0x31,
	0x14, 0xc7, 0x29, 0x31, 0x8a, 0x05, 0x19, 0xca, 0x42, 0x30, 0x9e, 0x48, 0x4c, 0x64, 0x2a, 0x09,
	0x6e, 0x32, 0x89, 0x31, 0xd1, 0xb8, 0x90, 0x86, 0xc9, 0x05, 0xb9, 0xf2, 0x84, 0x8b, 0x72, 0x3d,
	0xdb, 0x47, 0x52, 0xbf, 0x84, 0xf1, 0x0b, 0xb9, 0x3b, 0x32, 0x3a, 0x1a, 0xf8, 0x22, 0xa6, 0x27,
	0x27, 0x35, 0x46, 0x96, 0xe6, 0xdf, 0xe6, 0xff, 0x7b, 0xfd, 0xbf, 0xd7, 0xd2, 0x43, 0x84, 0x78,
	0x04, 0x7a, 0x1a, 0xc5, 0xd8, 0xd2, 0x89, 0x6c, 0x8d, 0xdd, 0x82, 0xcf, 0x09, 0x18, 0x9e, 0x68,
	0x85, 0x8a, 0x55, 0xd6, 0x06, 0xae, 0x13, 0xc9, 0x9d, 0xa1, 0xb6, 0xef, 0x51, 0xc3, 0x50, 0x46,
	0x3e, 0xd1, 0xd8, 0xa3, 0x45, 0x01, 0x4f, 0x33, 0x30, 0xd8, 0x8b, 0xe2, 0x71, 0xe3, 0x98, 0xb2,
	0xd5, 0xb6, 0xab, 0xd5, 0x70, 0x24, 0x87, 0x06, 0xfb, 0x96, 0x95, 0x69, 0x1e, 0x6d, 0x95, 0xd4,
	0x49, 0xb3, 0x24, 0xf2, 0x68, 0x1b, 0x65, 0x5a, 0x12, 0x60, 0x12, 0x15, 0x1b, 0x48, 0xa9, 0x17,
	0x42, 0x2b, 0xd9, 0x81, 0xcf, 0x75, 0x68, 0x41, 0x4e, 0x40, 0x3e, 0x0c, 0x56, 0x74, 0xb1, 0x5d,
	0xe7, 0x5e, 0x42, 0x17, 0x86, 0x67, 0xdc, 0x85, 0x33, 0xf6, 0xad, 0xd8, 0x91, 0xdf, 0x82, 0x9d,
	0xd1, 0x5d, 0xb4, 0x03, 0x0d, 0x66, 0xf6, 0x88, 0xd5, 0x7c, 0x4a, 0x1f, 0xfc, 0xa1, 0x2f, 0x2d,
	0xc8, 0xbe, 0x15, 0xa9, 0x49, 0x14, 0x70, 0xa5, 0xda, 0x6f, 0x84, 0x96, 0x7e, 0x82, 0x9c, 0xf7,
	0xae, 0xd9, 0x0d, 0xdd, 0x72, 0x49, 0xd9, 0xaf, 0xfb, 0xb3, 0x09, 0x71, 0x6f, 0x02, 0xb5, 0xa3,
	0x7f, 0x1c, 0xeb, 0x76, 0xd9, 0x1d, 0x2d, 0xfa, 0x5d, 0x9e, 0x6c, 0xaa, 0xe9, 0x19, 0x6b, 0xcd,
	0x8d, 0xa5, 0x3d, 0x67, 0xf7, 0xea, 0x7d, 0x11, 0x90, 0xf9, 0x22, 0x20, 0x9f, 0x8b, 0x80, 0xbc,
	0x2e, 0x83, 0xdc, 0x7c, 0x19, 0xe4, 0x3e, 0x96, 0x41, 0xee, 0x96, 0x8f, 0x23, 0x9c, 0xcc, 0x42,
	0x2e, 0xd5, 0xb4, 0x25, 0xd5, 0x14, 0x30, 0xbc, 0xc7, 0xb5, 0xc8, 0x3e, 0x45, 0x47, 0x2a, 0x0d,
	0x4e, 0x84, 0xdb, 0xe9, 0x33, 0x9f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xca, 0xdb, 0xe7,
	0x3b, 0x02, 0x00, 0x00,
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
type MulticastAPICustomer interface {
	Ping(ctx context.Context, in *SolicitPing, choices ...grpc.CallOption) (*ReplyPing, error)
	MulticastTransfer(ctx context.Context, in *SolicitMulticastTransfer, choices ...grpc.CallOption) (*ReplyMulticastTransfer, error)
}

type multicastAPICustomer struct {
	cc grpc1.ClientConn
}

func FreshMulticastAPICustomer(cc grpc1.ClientConn) MulticastAPICustomer {
	return &multicastAPICustomer{cc}
}

func (c *multicastAPICustomer) Ping(ctx context.Context, in *SolicitPing, choices ...grpc.CallOption) (*ReplyPing, error) {
	out := new(ReplyPing)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multicastAPICustomer) MulticastTransfer(ctx context.Context, in *SolicitMulticastTransfer, choices ...grpc.CallOption) (*ReplyMulticastTransfer, error) {
	out := new(ReplyMulticastTransfer)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, choices...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//
type MulticastAPIDaemon interface {
	Ping(context.Context, *SolicitPing) (*ReplyPing, error)
	MulticastTransfer(context.Context, *SolicitMulticastTransfer) (*ReplyMulticastTransfer, error)
}

//
type UndevelopedMulticastAPIDaemon struct {
}

func (*UndevelopedMulticastAPIDaemon) Ping(ctx context.Context, req *SolicitPing) (*ReplyPing, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UndevelopedMulticastAPIDaemon) MulticastTransfer(ctx context.Context, req *SolicitMulticastTransfer) (*ReplyMulticastTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}

func EnrollMulticastAPIDaemon(s grpc1.Server, srv MulticastAPIDaemon) {
	s.RegisterService(&_Multicastapi_servicedetails, srv)
}

func _Multicastapi_Ping_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(MulticastAPIDaemon).Ping(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MulticastAPIDaemon).Ping(ctx, req.(*SolicitPing))
	}
	return overseer(ctx, in, details, processor)
}

func _Multicastapi_Multicasttrans_Processor(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitMulticastTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(MulticastAPIDaemon).MulticastTransfer(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	processor := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MulticastAPIDaemon).MulticastTransfer(ctx, req.(*SolicitMulticastTransfer))
	}
	return overseer(ctx, in, details, processor)
}

var Multicastapi_servicedetails = _Multicastapi_servicedetails
var _Multicastapi_servicedetails = grpc.ServiceDesc{
	ServiceName: "REDACTED",
	HandlerType: (*MulticastAPIDaemon)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "REDACTED",
			Handler:    _Multicastapi_Ping_Processor,
		},
		{
			MethodName: "REDACTED",
			Handler:    _Multicastapi_Multicasttrans_Processor,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "REDACTED",
}

func (m *SolicitPing) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitPing) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitPing) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *SolicitMulticastTransfer) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SolicitMulticastTransfer) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SolicitMulticastTransfer) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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

func (m *ReplyPing) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyPing) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyPing) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *ReplyMulticastTransfer) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ReplyMulticastTransfer) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ReplyMulticastTransfer) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.TransferOutcome != nil {
		{
			extent, err := m.TransferOutcome.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
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
		deltaLocatedAN[i] = 0xa
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
func (m *SolicitPing) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SolicitMulticastTransfer) Extent() (n int) {
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

func (m *ReplyPing) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ReplyMulticastTransfer) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InspectTransfer != nil {
		l = m.InspectTransfer.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.TransferOutcome != nil {
		l = m.TransferOutcome.Extent()
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
func (m *SolicitPing) Decode(deltaLocatedAN []byte) error {
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
func (m *SolicitMulticastTransfer) Decode(deltaLocatedAN []byte) error {
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
func (m *ReplyPing) Decode(deltaLocatedAN []byte) error {
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
func (m *ReplyMulticastTransfer) Decode(deltaLocatedAN []byte) error {
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
			if m.InspectTransfer == nil {
				m.InspectTransfer = &kinds.ReplyInspectTransfer{}
			}
			if err := m.InspectTransfer.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
			if m.TransferOutcome == nil {
				m.TransferOutcome = &kinds.InvokeTransferOutcome{}
			}
			if err := m.TransferOutcome.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
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
