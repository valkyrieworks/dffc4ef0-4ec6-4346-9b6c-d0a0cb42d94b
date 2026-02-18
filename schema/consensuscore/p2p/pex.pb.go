//
//

package p2p

import (
	fmt "fmt"
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

type PexQuery struct {
}

func (m *PexQuery) Restore()         { *m = PexQuery{} }
func (m *PexQuery) String() string { return proto.CompactTextString(m) }
func (*PexQuery) SchemaSignal()    {}
func (*PexQuery) Definition() ([]byte, []int) {
	return filedefinition_hash2, []int{0}
}
func (m *PexQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PexQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Pexquery.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Pexquery.Merge(m, src)
}
func (m *PexQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *PexQuery) XXX_Omitunclear() {
	xxx_messagedata_Pexquery.DiscardUnknown(m)
}

var xxx_messagedata_Pexquery proto.InternalMessageInfo

type PexLocations struct {
	Locations []NetLocation `protobuf:"octets,1,rep,name=addrs,proto3" json:"locations"`
}

func (m *PexLocations) Restore()         { *m = PexLocations{} }
func (m *PexLocations) String() string { return proto.CompactTextString(m) }
func (*PexLocations) SchemaSignal()    {}
func (*PexLocations) Definition() ([]byte, []int) {
	return filedefinition_hash2, []int{1}
}
func (m *PexLocations) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PexLocations) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Pexlocations.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexLocations) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Pexlocations.Merge(m, src)
}
func (m *PexLocations) XXX_Volume() int {
	return m.Volume()
}
func (m *PexLocations) XXX_Omitunclear() {
	xxx_messagedata_Pexlocations.DiscardUnknown(m)
}

var xxx_messagedata_Pexlocations proto.InternalMessageInfo

func (m *PexLocations) FetchLocations() []NetLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

type Signal struct {
	//
	//
	//
	Sum ismessage_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) SchemaSignal()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedefinition_hash2, []int{2}
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

type Signal_Pexquery struct {
	PexQuery *PexQuery `protobuf:"octets,1,opt,name=pex_request,json=pexRequest,proto3,oneof" json:"pex_query,omitempty"`
}
type Signal_Pexlocations struct {
	PexLocations *PexLocations `protobuf:"octets,2,opt,name=pex_addrs,json=pexAddrs,proto3,oneof" json:"pex_locations,omitempty"`
}

func (*Signal_Pexquery) ismessage_Total() {}
func (*Signal_Pexlocations) ismessage_Total()   {}

func (m *Signal) FetchTotal() ismessage_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) FetchPexQuery() *PexQuery {
	if x, ok := m.FetchTotal().(*Signal_Pexquery); ok {
		return x.PexQuery
	}
	return nil
}

func (m *Signal) FetchPexLocations() *PexLocations {
	if x, ok := m.FetchTotal().(*Signal_Pexlocations); ok {
		return x.PexLocations
	}
	return nil
}

//
func (*Signal) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Signal_Pexquery)(nil),
		(*Signal_Pexlocations)(nil),
	}
}

func init() {
	proto.RegisterType((*PexQuery)(nil), "REDACTED")
	proto.RegisterType((*PexLocations)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash2) }

var filedefinition_hash2 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x30, 0x2a, 0xd0, 0x2f, 0x48, 0xad, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x43, 0xc8, 0xe8, 0x15, 0x18, 0x15, 0x48, 0x49, 0xa1, 0xa9,
	0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0xa8, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5,
	0x41, 0x2c, 0x88, 0xa8, 0x12, 0x0f, 0x17, 0x57, 0x40, 0x6a, 0x45, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x89, 0x92, 0x13, 0x17, 0x47, 0x40, 0x6a, 0x85, 0x63, 0x4a, 0x4a, 0x51, 0xb1, 0x90, 0x19,
	0x17, 0x6b, 0x22, 0x88, 0x21, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa5, 0x87, 0x6a, 0x97,
	0x9e, 0x5f, 0x6a, 0x09, 0x48, 0x61, 0x6a, 0x71, 0xb1, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41,
	0x10, 0xe5, 0x4a, 0x1d, 0x8c, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0xb6,
	0x5c, 0xdc, 0x05, 0xa9, 0x15, 0xf1, 0x45, 0x10, 0xe3, 0x25, 0x18, 0x15, 0x18, 0xb1, 0x99, 0x84,
	0x70, 0x80, 0x07, 0x43, 0x10, 0x57, 0x01, 0x9c, 0x27, 0x64, 0xce, 0xc5, 0x09, 0xd2, 0x0e, 0x71,
	0x06, 0x13, 0x58, 0xb3, 0x04, 0x16, 0xcd, 0x60, 0xf7, 0x7a, 0x30, 0x04, 0x71, 0x14, 0x40, 0xd9,
	0x4e, 0xac, 0x5c, 0xcc, 0xc5, 0xa5, 0xb9, 0x4e, 0xde, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0x65, 0x98, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f,
	0x9c, 0x9f, 0x9b, 0x5a, 0x92, 0x94, 0x56, 0x82, 0x60, 0x40, 0x42, 0x09, 0x35, 0x2c, 0x93, 0xd8,
	0xc0, 0xa2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xad, 0x52, 0xe1, 0x8e, 0x01, 0x00,
	0x00,
}

func (m *PexQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PexQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PexLocations) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexLocations) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PexLocations) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Locations) > 0 {
		for idxNdEx := len(m.Locations) - 1; idxNdEx >= 0; idxNdEx-- {
			{
				volume, err := m.Locations[idxNdEx].SerializeToDimensionedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= volume
				i = encodeVariableintPex(dAtA, i, uint64(volume))
			}
			i--
			dAtA[i] = 0xa
		}
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

func (m *Signal_Pexquery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Pexquery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexQuery != nil {
		{
			volume, err := m.PexQuery.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintPex(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Pexlocations) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Pexlocations) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexLocations != nil {
		{
			volume, err := m.PexLocations.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintPex(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVariableintPex(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovPex(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *PexQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PexLocations) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Locations) > 0 {
		for _, e := range m.Locations {
			l = e.Volume()
			n += 1 + l + sovPex(uint64(l))
		}
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

func (m *Signal_Pexquery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexQuery != nil {
		l = m.PexQuery.Volume()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}
func (m *Signal_Pexlocations) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexLocations != nil {
		l = m.PexLocations.Volume()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}

func sovPex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPex(x uint64) (n int) {
	return sovPex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PexQuery) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadPex
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
			skippy, err := omitPex(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentPex
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
func (m *PexLocations) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadPex
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
					return ErrIntegerOverloadPex
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
				return ErrCorruptExtentPex
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentPex
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Locations = append(m.Locations, NetLocation{})
			if err := m.Locations[len(m.Locations)-1].Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitPex(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentPex
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
				return ErrIntegerOverloadPex
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
					return ErrIntegerOverloadPex
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
				return ErrCorruptExtentPex
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentPex
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Pexquery{v}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadPex
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
				return ErrCorruptExtentPex
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentPex
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexLocations{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Pexlocations{v}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitPex(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentPex
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
func omitPex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadPex
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
					return 0, ErrIntegerOverloadPex
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
					return 0, ErrIntegerOverloadPex
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
				return 0, ErrCorruptExtentPex
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterPex
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentPex
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentPex        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadPex          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterPex = fmt.Errorf("REDACTED")
)
