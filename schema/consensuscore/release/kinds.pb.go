//
//

package release

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

//
//
//
type App struct {
	Protocol uint64 `protobuf:"variableint,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Software string `protobuf:"octets,2,opt,name=software,proto3" json:"solution,omitempty"`
}

func (m *App) Restore()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) SchemaSignal()    {}
func (*App) Definition() ([]byte, []int) {
	return filedefinition_hash4, []int{0}
}
func (m *App) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *App) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Application.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *App) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Application.Merge(m, src)
}
func (m *App) XXX_Volume() int {
	return m.Volume()
}
func (m *App) XXX_Omitunclear() {
	xxx_messagedata_Application.DiscardUnknown(m)
}

var xxx_messagedata_Application proto.InternalMessageInfo

func (m *App) FetchProtocol() uint64 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *App) FetchSolution() string {
	if m != nil {
		return m.Software
	}
	return "REDACTED"
}

//
//
//
type Agreement struct {
	Ledger uint64 `protobuf:"variableint,1,opt,name=block,proto3" json:"ledger,omitempty"`
	App   uint64 `protobuf:"variableint,2,opt,name=app,proto3" json:"app,omitempty"`
}

func (m *Agreement) Restore()         { *m = Agreement{} }
func (m *Agreement) String() string { return proto.CompactTextString(m) }
func (*Agreement) SchemaSignal()    {}
func (*Agreement) Definition() ([]byte, []int) {
	return filedefinition_hash4, []int{1}
}
func (m *Agreement) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Agreement) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Agreement.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Agreement) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Agreement.Merge(m, src)
}
func (m *Agreement) XXX_Volume() int {
	return m.Volume()
}
func (m *Agreement) XXX_Omitunclear() {
	xxx_messagedata_Agreement.DiscardUnknown(m)
}

var xxx_messagedata_Agreement proto.InternalMessageInfo

func (m *Agreement) FetchLedger() uint64 {
	if m != nil {
		return m.Ledger
	}
	return 0
}

func (m *Agreement) FetchApplication() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func init() {
	proto.RegisterType((*App)(nil), "REDACTED")
	proto.RegisterType((*Agreement)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash4) }

var filedefinition_hash4 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2f,
	0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0xc8, 0xeb, 0x41,
	0xe5, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xd2, 0xfa, 0x20, 0x16, 0x44, 0xa5, 0x92, 0x2d,
	0x17, 0xb3, 0x63, 0x41, 0x81, 0x90, 0x14, 0x17, 0x07, 0x98, 0x9f, 0x9c, 0x9f, 0x23, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x12, 0x04, 0xe7, 0x83, 0xe4, 0x8a, 0xf3, 0xd3, 0x4a, 0xca, 0x13, 0x8b, 0x52,
	0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x4b, 0x2e, 0x4e, 0xe7, 0xfc, 0xbc,
	0xe2, 0xd4, 0xbc, 0xe2, 0xd2, 0x62, 0x21, 0x11, 0x2e, 0xd6, 0xa4, 0x9c, 0xfc, 0xe4, 0x6c, 0xa8,
	0x09, 0x10, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x62, 0x41, 0x01, 0x58, 0x27, 0x4b, 0x10, 0x88, 0x69,
	0xc5, 0xf2, 0x62, 0x81, 0x3c, 0xa3, 0x93, 0xff, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0x99, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7,
	0xe7, 0xa6, 0x96, 0x24, 0xa5, 0x95, 0x20, 0x18, 0x10, 0x2f, 0x60, 0x06, 0x40, 0x12, 0x1b, 0x58,
	0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xc7, 0x18, 0x2b, 0x1d, 0x01, 0x00, 0x00,
}

func (this *Agreement) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Agreement)
	if !ok {
		that2, ok := that.(Agreement)
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
	if this.Ledger != that1.Ledger {
		return false
	}
	if this.App != that1.App {
		return false
	}
	return true
}
func (m *App) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *App) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *App) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Software) > 0 {
		i -= len(m.Software)
		copy(dAtA[i:], m.Software)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Software)))
		i--
		dAtA[i] = 0x12
	}
	if m.Protocol != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Protocol))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Agreement) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Agreement) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Agreement) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.App))
		i--
		dAtA[i] = 0x10
	}
	if m.Ledger != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ledger))
		i--
		dAtA[i] = 0x8
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
func (m *App) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Protocol != 0 {
		n += 1 + sovKinds(uint64(m.Protocol))
	}
	l = len(m.Software)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Agreement) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ledger != 0 {
		n += 1 + sovKinds(uint64(m.Ledger))
	}
	if m.App != 0 {
		n += 1 + sovKinds(uint64(m.App))
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *App) Unserialize(dAtA []byte) error {
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
			m.Protocol = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Protocol |= uint64(b&0x7F) << displace
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
			m.Software = string(dAtA[idxNdEx:submitOrdinal])
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
func (m *Agreement) Unserialize(dAtA []byte) error {
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
			m.Ledger = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Ledger |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.App = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
