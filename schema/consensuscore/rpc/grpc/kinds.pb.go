//
//

package coregrpc

import (
	context "context"
	fmt "fmt"
	kinds "github.com/valkyrieworks/iface/kinds"
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

type QueryPing struct {
}

func (m *QueryPing) Restore()         { *m = QueryPing{} }
func (m *QueryPing) String() string { return proto.CompactTextString(m) }
func (*QueryPing) SchemaSignal()    {}
func (*QueryPing) Definition() ([]byte, []int) {
	return filedefinition_hash6, []int{0}
}
func (m *QueryPing) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryPing) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Queryping.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPing) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Queryping.Merge(m, src)
}
func (m *QueryPing) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryPing) XXX_Omitunclear() {
	xxx_messagedata_Queryping.DiscardUnknown(m)
}

var xxx_messagedata_Queryping proto.InternalMessageInfo

type QueryMulticastTransfer struct {
	Tx []byte `protobuf:"octets,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *QueryMulticastTransfer) Restore()         { *m = QueryMulticastTransfer{} }
func (m *QueryMulticastTransfer) String() string { return proto.CompactTextString(m) }
func (*QueryMulticastTransfer) SchemaSignal()    {}
func (*QueryMulticastTransfer) Definition() ([]byte, []int) {
	return filedefinition_hash6, []int{1}
}
func (m *QueryMulticastTransfer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *QueryMulticastTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Querymulticasttransfer.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMulticastTransfer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Querymulticasttransfer.Merge(m, src)
}
func (m *QueryMulticastTransfer) XXX_Volume() int {
	return m.Volume()
}
func (m *QueryMulticastTransfer) XXX_Omitunclear() {
	xxx_messagedata_Querymulticasttransfer.DiscardUnknown(m)
}

var xxx_messagedata_Querymulticasttransfer proto.InternalMessageInfo

func (m *QueryMulticastTransfer) FetchTransfer() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type AnswerPing struct {
}

func (m *AnswerPing) Restore()         { *m = AnswerPing{} }
func (m *AnswerPing) String() string { return proto.CompactTextString(m) }
func (*AnswerPing) SchemaSignal()    {}
func (*AnswerPing) Definition() ([]byte, []int) {
	return filedefinition_hash6, []int{2}
}
func (m *AnswerPing) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AnswerPing) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Answerping.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnswerPing) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Answerping.Merge(m, src)
}
func (m *AnswerPing) XXX_Volume() int {
	return m.Volume()
}
func (m *AnswerPing) XXX_Omitunclear() {
	xxx_messagedata_Answerping.DiscardUnknown(m)
}

var xxx_messagedata_Answerping proto.InternalMessageInfo

type AnswerMulticastTransfer struct {
	InspectTransfer  *kinds.ReplyInspectTransfer `protobuf:"octets,1,opt,name=check_tx,json=checkTx,proto3" json:"inspect_transfer,omitempty"`
	TransOutcome *kinds.InvokeTransferOutcome    `protobuf:"octets,2,opt,name=tx_result,json=txResult,proto3" json:"transfer_outcome,omitempty"`
}

func (m *AnswerMulticastTransfer) Restore()         { *m = AnswerMulticastTransfer{} }
func (m *AnswerMulticastTransfer) String() string { return proto.CompactTextString(m) }
func (*AnswerMulticastTransfer) SchemaSignal()    {}
func (*AnswerMulticastTransfer) Definition() ([]byte, []int) {
	return filedefinition_hash6, []int{3}
}
func (m *AnswerMulticastTransfer) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *AnswerMulticastTransfer) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Answermulticasttransfer.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnswerMulticastTransfer) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Answermulticasttransfer.Merge(m, src)
}
func (m *AnswerMulticastTransfer) XXX_Volume() int {
	return m.Volume()
}
func (m *AnswerMulticastTransfer) XXX_Omitunclear() {
	xxx_messagedata_Answermulticasttransfer.DiscardUnknown(m)
}

var xxx_messagedata_Answermulticasttransfer proto.InternalMessageInfo

func (m *AnswerMulticastTransfer) FetchInspectTransfer() *kinds.ReplyInspectTransfer {
	if m != nil {
		return m.InspectTransfer
	}
	return nil
}

func (m *AnswerMulticastTransfer) FetchTransferOutcome() *kinds.InvokeTransferOutcome {
	if m != nil {
		return m.TransOutcome
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryPing)(nil), "REDACTED")
	proto.RegisterType((*QueryMulticastTransfer)(nil), "REDACTED")
	proto.RegisterType((*AnswerPing)(nil), "REDACTED")
	proto.RegisterType((*AnswerMulticastTransfer)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash6) }

var filedefinition_hash6 = []byte{
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
	Ping(ctx context.Context, in *QueryPing, opts ...grpc.CallOption) (*AnswerPing, error)
	MulticastTransfer(ctx context.Context, in *QueryMulticastTransfer, opts ...grpc.CallOption) (*AnswerMulticastTransfer, error)
}

type multicastAPICustomer struct {
	cc grpc1.ClientConn
}

func NewMulticastAPICustomer(cc grpc1.ClientConn) MulticastAPICustomer {
	return &multicastAPICustomer{cc}
}

func (c *multicastAPICustomer) Ping(ctx context.Context, in *QueryPing, opts ...grpc.CallOption) (*AnswerPing, error) {
	out := new(AnswerPing)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multicastAPICustomer) MulticastTransfer(ctx context.Context, in *QueryMulticastTransfer, opts ...grpc.CallOption) (*AnswerMulticastTransfer, error) {
	out := new(AnswerMulticastTransfer)
	err := c.cc.Invoke(ctx, "REDACTED", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//
type MulticastAPIHost interface {
	Ping(context.Context, *QueryPing) (*AnswerPing, error)
	MulticastTransfer(context.Context, *QueryMulticastTransfer) (*AnswerMulticastTransfer, error)
}

//
type UnexecutedMulticastAPIHost struct {
}

func (*UnexecutedMulticastAPIHost) Ping(ctx context.Context, req *QueryPing) (*AnswerPing, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}
func (*UnexecutedMulticastAPIHost) MulticastTransfer(ctx context.Context, req *QueryMulticastTransfer) (*AnswerMulticastTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "REDACTED")
}

func EnrollMulticastAPIHost(s grpc1.Server, srv MulticastAPIHost) {
	s.RegisterService(&_Multicastapi_servicedefinition, srv)
}

func _Multicastapi_Ping_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(MulticastAPIHost).Ping(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MulticastAPIHost).Ping(ctx, req.(*QueryPing))
	}
	return overseer(ctx, in, details, manager)
}

func _Multicastapi_Multicasttransfer_Manager(srv interface{}, ctx context.Context, dec func(interface{}) error, overseer grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMulticastTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if overseer == nil {
		return srv.(MulticastAPIHost).MulticastTransfer(ctx, in)
	}
	details := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "REDACTED",
	}
	manager := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MulticastAPIHost).MulticastTransfer(ctx, req.(*QueryMulticastTransfer))
	}
	return overseer(ctx, in, details, manager)
}

var Multicastapi_servicedefinition = _Multicastapi_servicedefinition
var _Multicastapi_servicedefinition = grpc.ServiceDesc{
	ServiceName: "REDACTED",
	HandlerType: (*MulticastAPIHost)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "REDACTED",
			Handler:    _Multicastapi_Ping_Manager,
		},
		{
			MethodName: "REDACTED",
			Handler:    _Multicastapi_Multicasttransfer_Manager,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "REDACTED",
}

func (m *QueryPing) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPing) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryPing) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMulticastTransfer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMulticastTransfer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *QueryMulticastTransfer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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

func (m *AnswerPing) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnswerPing) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AnswerPing) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *AnswerMulticastTransfer) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnswerMulticastTransfer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *AnswerMulticastTransfer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TransOutcome != nil {
		{
			volume, err := m.TransOutcome.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
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
		dAtA[i] = 0xa
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
func (m *QueryPing) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMulticastTransfer) Volume() (n int) {
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

func (m *AnswerPing) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *AnswerMulticastTransfer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InspectTransfer != nil {
		l = m.InspectTransfer.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.TransOutcome != nil {
		l = m.TransOutcome.Volume()
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
func (m *QueryPing) Unserialize(dAtA []byte) error {
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
func (m *QueryMulticastTransfer) Unserialize(dAtA []byte) error {
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
func (m *AnswerPing) Unserialize(dAtA []byte) error {
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
func (m *AnswerMulticastTransfer) Unserialize(dAtA []byte) error {
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
			if m.InspectTransfer == nil {
				m.InspectTransfer = &kinds.ReplyInspectTransfer{}
			}
			if err := m.InspectTransfer.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
			if m.TransOutcome == nil {
				m.TransOutcome = &kinds.InvokeTransferOutcome{}
			}
			if err := m.TransOutcome.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
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
