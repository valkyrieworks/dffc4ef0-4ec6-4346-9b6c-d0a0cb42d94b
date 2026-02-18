//
//

package depot

import (
	fmt "fmt"
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

type LedgerDepotStatus struct {
	Root   int64 `protobuf:"variableint,1,opt,name=base,proto3" json:"root,omitempty"`
	Level int64 `protobuf:"variableint,2,opt,name=height,proto3" json:"level,omitempty"`
}

func (m *LedgerDepotStatus) Restore()         { *m = LedgerDepotStatus{} }
func (m *LedgerDepotStatus) String() string { return proto.CompactTextString(m) }
func (*LedgerDepotStatus) SchemaSignal()    {}
func (*LedgerDepotStatus) Definition() ([]byte, []int) {
	return filedefinition_hash9, []int{0}
}
func (m *LedgerDepotStatus) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *LedgerDepotStatus) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Blockdepotstatus.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LedgerDepotStatus) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Blockdepotstatus.Merge(m, src)
}
func (m *LedgerDepotStatus) XXX_Volume() int {
	return m.Volume()
}
func (m *LedgerDepotStatus) XXX_Omitunclear() {
	xxx_messagedata_Blockdepotstatus.DiscardUnknown(m)
}

var xxx_messagedata_Blockdepotstatus proto.InternalMessageInfo

func (m *LedgerDepotStatus) FetchRoot() int64 {
	if m != nil {
		return m.Root
	}
	return 0
}

func (m *LedgerDepotStatus) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*LedgerDepotStatus)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash9) }

var filedefinition_hash9 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2f, 0xa9, 0x2c,
	0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0xc8, 0xea, 0x81, 0x65, 0x95,
	0x6c, 0xb9, 0xf8, 0x9d, 0x72, 0xf2, 0x93, 0xb3, 0x83, 0x41, 0xbc, 0xe0, 0x92, 0xc4, 0x92, 0x54,
	0x21, 0x21, 0x2e, 0x96, 0xa4, 0xc4, 0xe2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x30,
	0x5b, 0x48, 0x8c, 0x8b, 0x2d, 0x23, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x09, 0x2c, 0x0a, 0xe5,
	0x39, 0xf9, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x71, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x72, 0x7e, 0x6e, 0x6a, 0x49, 0x52, 0x5a,
	0x09, 0x82, 0x01, 0x76, 0x8e, 0x3e, 0xba, 0x5b, 0x93, 0xd8, 0xc0, 0xe2, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0x2b, 0x34, 0x2a, 0xc6, 0x00, 0x00, 0x00,
}

func (m *LedgerDepotStatus) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LedgerDepotStatus) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *LedgerDepotStatus) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	if m.Root != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Root))
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
func (m *LedgerDepotStatus) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Root != 0 {
		n += 1 + sovKinds(uint64(m.Root))
	}
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LedgerDepotStatus) Unserialize(dAtA []byte) error {
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
		case 2:
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
