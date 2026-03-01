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

type IncidentDataIterationStatus struct {
	Altitude int64  `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration  int32  `protobuf:"variableint,2,opt,name=round,proto3" json:"iteration,omitempty"`
	Phase   string `protobuf:"octets,3,opt,name=step,proto3" json:"phase,omitempty"`
}

func (m *IncidentDataIterationStatus) Restore()         { *m = IncidentDataIterationStatus{} }
func (m *IncidentDataIterationStatus) Text() string { return proto.CompactTextString(m) }
func (*IncidentDataIterationStatus) SchemaArtifact()    {}
func (*IncidentDataIterationStatus) Definition() ([]byte, []int) {
	return filedescriptor_72cfafd446dedf7c, []int{0}
}
func (m *IncidentDataIterationStatus) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *IncidentDataIterationStatus) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Incidentiterationstate.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncidentDataIterationStatus) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Incidentiterationstate.Merge(m, src)
}
func (m *IncidentDataIterationStatus) XXX_Extent() int {
	return m.Extent()
}
func (m *IncidentDataIterationStatus) XXX_Dropunfamiliar() {
	xxx_signaldetails_Incidentiterationstate.DiscardUnknown(m)
}

var xxx_signaldetails_Incidentiterationstate proto.InternalMessageInfo

func (m *IncidentDataIterationStatus) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *IncidentDataIterationStatus) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *IncidentDataIterationStatus) ObtainPhase() string {
	if m != nil {
		return m.Phase
	}
	return "REDACTED"
}

func initialize() {
	proto.RegisterType((*IncidentDataIterationStatus)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_72cfafd446dedf7c) }

var filedescriptor_72cfafd446dedf7c = []byte{
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

func (m *IncidentDataIterationStatus) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *IncidentDataIterationStatus) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *IncidentDataIterationStatus) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Phase) > 0 {
		i -= len(m.Phase)
		copy(deltaLocatedAN[i:], m.Phase)
		i = encodeVariableintIncidents(deltaLocatedAN, i, uint64(len(m.Phase)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	if m.Iteration != 0 {
		i = encodeVariableintIncidents(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = encodeVariableintIncidents(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintIncidents(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovIncidents(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *IncidentDataIterationStatus) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovIncidents(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovIncidents(uint64(m.Iteration))
	}
	l = len(m.Phase)
	if l > 0 {
		n += 1 + l + sovIncidents(uint64(l))
	}
	return n
}

func sovIncidents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncidents(x uint64) (n int) {
	return sovIncidents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IncidentDataIterationStatus) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunIncidents
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
					return FaultIntegerOverrunIncidents
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
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunIncidents
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunIncidents
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeIncidents
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeIncidents
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Phase = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitIncidents(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeIncidents
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
func omitIncidents(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunIncidents
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
					return 0, FaultIntegerOverrunIncidents
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
					return 0, FaultIntegerOverrunIncidents
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
				return 0, FaultUnfitMagnitudeIncidents
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionIncidents
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeIncidents
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeIncidents        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunIncidents          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionIncidents = fmt.Errorf("REDACTED")
)
