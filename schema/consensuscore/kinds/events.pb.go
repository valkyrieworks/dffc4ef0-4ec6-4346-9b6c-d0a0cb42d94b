//
//

package kinds

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

type EventDataDurationStatus struct {
	Level int64  `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Cycle  int32  `protobuf:"variableint,2,opt,name=round,proto3" json:"epoch,omitempty"`
	Phase   string `protobuf:"octets,3,opt,name=step,proto3" json:"phase,omitempty"`
}

func (m *EventDataDurationStatus) Restore()         { *m = EventDataDurationStatus{} }
func (m *EventDataDurationStatus) String() string { return proto.CompactTextString(m) }
func (*EventDataDurationStatus) SchemaSignal()    {}
func (*EventDataDurationStatus) Definition() ([]byte, []int) {
	return filedefinition_72cfafd446dedf7c, []int{0}
}
func (m *EventDataDurationStatus) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *EventDataDurationStatus) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Signaldatadurationstate.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDataDurationStatus) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Signaldatadurationstate.Merge(m, src)
}
func (m *EventDataDurationStatus) XXX_Volume() int {
	return m.Volume()
}
func (m *EventDataDurationStatus) XXX_Omitunclear() {
	xxx_messagedata_Signaldatadurationstate.DiscardUnknown(m)
}

var xxx_messagedata_Signaldatadurationstate proto.InternalMessageInfo

func (m *EventDataDurationStatus) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *EventDataDurationStatus) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *EventDataDurationStatus) FetchPhase() string {
	if m != nil {
		return m.Phase
	}
	return "REDACTED"
}

func init() {
	proto.RegisterType((*EventDataDurationStatus)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_72cfafd446dedf7c) }

var filedefinition_72cfafd446dedf7c = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x2d, 0x4b,
	0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0x48, 0xeb, 0x81, 0xa5,
	0x95, 0xc2, 0xb9, 0x84, 0x5d, 0x41, 0x2a, 0x5c, 0x12, 0x4b, 0x12, 0x83, 0xf2, 0x4b, 0xf3, 0x52,
	0x82, 0x4b, 0x12, 0x4b, 0x52, 0x85, 0xc4, 0xb8, 0xd8, 0x32, 0x52, 0x33, 0xd3, 0x33, 0x4a, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xa0, 0x3c, 0x21, 0x11, 0x2e, 0xd6, 0x22, 0x90, 0x2a, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x48, 0x88, 0x8b, 0xa5, 0xb8, 0x24, 0xb5, 0x40,
	0x82, 0x59, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x76, 0xf2, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0xe3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xe4, 0xfc, 0xdc, 0xd4, 0x92, 0xa4, 0xb4, 0x12, 0x04, 0x03, 0xec, 0x50, 0x7d, 0x74, 0x6f,
	0x24, 0xb1, 0x81, 0xc5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xb3, 0x7f, 0x37, 0xe1,
	0x00, 0x00, 0x00,
}

func (m *EventDataDurationStatus) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDataDurationStatus) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *EventDataDurationStatus) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Phase) > 0 {
		i -= len(m.Phase)
		copy(dAtA[i:], m.Phase)
		i = encodeVariableintEvents(dAtA, i, uint64(len(m.Phase)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Cycle != 0 {
		i = encodeVariableintEvents(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = encodeVariableintEvents(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVariableintEvents(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovEvents(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *EventDataDurationStatus) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovEvents(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovEvents(uint64(m.Cycle))
	}
	l = len(m.Phase)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventDataDurationStatus) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadEvents
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
					return ErrIntegerOverloadEvents
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
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvents
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadEvents
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
				return ErrCorruptExtentEvents
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentEvents
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Phase = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitEvents(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentEvents
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
func omitEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadEvents
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
					return 0, ErrIntegerOverloadEvents
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
					return 0, ErrIntegerOverloadEvents
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
				return 0, ErrCorruptExtentEvents
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterEvents
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentEvents
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentEvents        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadEvents          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterEvents = fmt.Errorf("REDACTED")
)
