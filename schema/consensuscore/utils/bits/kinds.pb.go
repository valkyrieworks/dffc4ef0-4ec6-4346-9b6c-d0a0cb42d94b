//
//

package bits

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

type BitList struct {
	Bits  int64    `protobuf:"variableint,1,opt,name=bits,proto3" json:"bits,omitempty"`
	Elements []uint64 `protobuf:"variableint,2,rep,packed,name=elems,proto3" json:"elements,omitempty"`
}

func (m *BitList) Restore()         { *m = BitList{} }
func (m *BitList) String() string { return proto.CompactTextString(m) }
func (*BitList) SchemaSignal()    {}
func (*BitList) Definition() ([]byte, []int) {
	return filedefinition_hash5, []int{0}
}
func (m *BitList) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *BitList) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Bitfield.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BitList) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Bitfield.Merge(m, src)
}
func (m *BitList) XXX_Volume() int {
	return m.Volume()
}
func (m *BitList) XXX_Omitunclear() {
	xxx_messagedata_Bitfield.DiscardUnknown(m)
}

var xxx_messagedata_Bitfield proto.InternalMessageInfo

func (m *BitList) FetchBits() int64 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *BitList) FetchElements() []uint64 {
	if m != nil {
		return m.Elements
	}
	return nil
}

func init() {
	proto.RegisterType((*BitList)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_hash5) }

var filedefinition_hash5 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0xcf, 0xc9, 0x4c, 0x2a, 0xd6, 0x4f, 0xca, 0x2c, 0x29,
	0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x41, 0xa8,
	0xd0, 0x03, 0xa9, 0xd0, 0x03, 0xa9, 0x50, 0x32, 0xe1, 0xe2, 0x70, 0xca, 0x2c, 0x71, 0x2c, 0x2a,
	0x4a, 0xac, 0x14, 0x12, 0xe2, 0x62, 0x01, 0x89, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x81,
	0xd9, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0x39, 0xa9, 0xb9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0x2c,
	0x41, 0x10, 0x8e, 0x53, 0xe0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99,
	0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0xe7, 0xa6, 0x96,
	0x24, 0xa5, 0x95, 0x20, 0x18, 0x60, 0x97, 0xe8, 0x63, 0x73, 0x6a, 0x12, 0x1b, 0x58, 0xce, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x09, 0xfd, 0x78, 0xed, 0xc9, 0x00, 0x00, 0x00,
}

func (m *BitList) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BitList) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *BitList) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Elements) > 0 {
		dAtA2 := make([]byte, len(m.Elements)*10)
		var j1 int
		for _, num := range m.Elements {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = formatVariableintKinds(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Bits != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Bits))
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
func (m *BitList) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bits != 0 {
		n += 1 + sovKinds(uint64(m.Bits))
	}
	if len(m.Elements) > 0 {
		l = 0
		for _, e := range m.Elements {
			l += sovKinds(uint64(e))
		}
		n += 1 + sovKinds(uint64(l)) + l
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BitList) Unserialize(dAtA []byte) error {
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
			m.Bits = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Bits |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind == 0 {
				var v uint64
				for displace := uint(0); ; displace += 7 {
					if displace >= 64 {
						return ErrIntegerOverloadKinds
					}
					if idxNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[idxNdEx]
					idxNdEx++
					v |= uint64(b&0x7F) << displace
					if b < 0x80 {
						break
					}
				}
				m.Elements = append(m.Elements, v)
			} else if cableKind == 2 {
				var compressedSize int
				for displace := uint(0); ; displace += 7 {
					if displace >= 64 {
						return ErrIntegerOverloadKinds
					}
					if idxNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[idxNdEx]
					idxNdEx++
					compressedSize |= int(b&0x7F) << displace
					if b < 0x80 {
						break
					}
				}
				if compressedSize < 0 {
					return ErrCorruptExtentKinds
				}
				submitOrdinal := idxNdEx + compressedSize
				if submitOrdinal < 0 {
					return ErrCorruptExtentKinds
				}
				if submitOrdinal > l {
					return io.ErrUnexpectedEOF
				}
				var componentCount int
				var tally int
				for _, integer := range dAtA[idxNdEx:submitOrdinal] {
					if integer < 128 {
						tally++
					}
				}
				componentCount = tally
				if componentCount != 0 && len(m.Elements) == 0 {
					m.Elements = make([]uint64, 0, componentCount)
				}
				for idxNdEx < submitOrdinal {
					var v uint64
					for displace := uint(0); ; displace += 7 {
						if displace >= 64 {
							return ErrIntegerOverloadKinds
						}
						if idxNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[idxNdEx]
						idxNdEx++
						v |= uint64(b&0x7F) << displace
						if b < 0x80 {
							break
						}
					}
					m.Elements = append(m.Elements, v)
				}
			} else {
				return fmt.Errorf("REDACTED", cableKind)
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
