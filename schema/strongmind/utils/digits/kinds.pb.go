//
//

package digits

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

type DigitSeries struct {
	Digits  int64    `protobuf:"variableint,1,opt,name=bits,proto3" json:"digits,omitempty"`
	Components []uint64 `protobuf:"variableint,2,rep,packed,name=elems,proto3" json:"components,omitempty"`
}

func (m *DigitSeries) Restore()         { *m = DigitSeries{} }
func (m *DigitSeries) Text() string { return proto.CompactTextString(m) }
func (*DigitSeries) SchemaArtifact()    {}
func (*DigitSeries) Definition() ([]byte, []int) {
	return filedescriptor_e91ab2672920d7d4, []int{0}
}
func (m *DigitSeries) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *DigitSeries) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Digitsequence.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DigitSeries) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Digitsequence.Merge(m, src)
}
func (m *DigitSeries) XXX_Extent() int {
	return m.Extent()
}
func (m *DigitSeries) XXX_Dropunfamiliar() {
	xxx_signaldetails_Digitsequence.DiscardUnknown(m)
}

var xxx_signaldetails_Digitsequence proto.InternalMessageInfo

func (m *DigitSeries) ObtainDigits() int64 {
	if m != nil {
		return m.Digits
	}
	return 0
}

func (m *DigitSeries) ObtainComponents() []uint64 {
	if m != nil {
		return m.Components
	}
	return nil
}

func initialize() {
	proto.RegisterType((*DigitSeries)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_e91ab2672920d7d4) }

var filedescriptor_e91ab2672920d7d4 = []byte{
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

func (m *DigitSeries) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *DigitSeries) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *DigitSeries) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Components) > 0 {
		deltaLocatedA2 := make([]byte, len(m.Components)*10)
		var j1 int
		for _, num := range m.Components {
			for num >= 1<<7 {
				deltaLocatedA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			deltaLocatedA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(deltaLocatedAN[i:], deltaLocatedA2[:j1])
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(j1))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Digits != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Digits))
		i--
		deltaLocatedAN[i] = 0x8
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
func (m *DigitSeries) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Digits != 0 {
		n += 1 + sovKinds(uint64(m.Digits))
	}
	if len(m.Components) > 0 {
		l = 0
		for _, e := range m.Components {
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
func (m *DigitSeries) Decode(deltaLocatedAN []byte) error {
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
			m.Digits = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Digits |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind == 0 {
				var v uint64
				for relocate := uint(0); ; relocate += 7 {
					if relocate >= 64 {
						return FaultIntegerOverrunKinds
					}
					if idxNdExc >= l {
						return io.ErrUnexpectedEOF
					}
					b := deltaLocatedAN[idxNdExc]
					idxNdExc++
					v |= uint64(b&0x7F) << relocate
					if b < 0x80 {
						break
					}
				}
				m.Components = append(m.Components, v)
			} else if cableKind == 2 {
				var compressedSize int
				for relocate := uint(0); ; relocate += 7 {
					if relocate >= 64 {
						return FaultIntegerOverrunKinds
					}
					if idxNdExc >= l {
						return io.ErrUnexpectedEOF
					}
					b := deltaLocatedAN[idxNdExc]
					idxNdExc++
					compressedSize |= int(b&0x7F) << relocate
					if b < 0x80 {
						break
					}
				}
				if compressedSize < 0 {
					return FaultUnfitMagnitudeKinds
				}
				submitOrdinal := idxNdExc + compressedSize
				if submitOrdinal < 0 {
					return FaultUnfitMagnitudeKinds
				}
				if submitOrdinal > l {
					return io.ErrUnexpectedEOF
				}
				var componentTotal int
				var tally int
				for _, number := range deltaLocatedAN[idxNdExc:submitOrdinal] {
					if number < 128 {
						tally++
					}
				}
				componentTotal = tally
				if componentTotal != 0 && len(m.Components) == 0 {
					m.Components = make([]uint64, 0, componentTotal)
				}
				for idxNdExc < submitOrdinal {
					var v uint64
					for relocate := uint(0); ; relocate += 7 {
						if relocate >= 64 {
							return FaultIntegerOverrunKinds
						}
						if idxNdExc >= l {
							return io.ErrUnexpectedEOF
						}
						b := deltaLocatedAN[idxNdExc]
						idxNdExc++
						v |= uint64(b&0x7F) << relocate
						if b < 0x80 {
							break
						}
					}
					m.Components = append(m.Components, v)
				}
			} else {
				return fmt.Errorf("REDACTED", cableKind)
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
