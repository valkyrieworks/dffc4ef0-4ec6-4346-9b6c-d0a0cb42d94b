//
//

package edition

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
	Scheme uint64 `protobuf:"variableint,1,opt,name=protocol,proto3" json:"scheme,omitempty"`
	Package string `protobuf:"octets,2,opt,name=software,proto3" json:"package,omitempty"`
}

func (m *App) Restore()         { *m = App{} }
func (m *App) Text() string { return proto.CompactTextString(m) }
func (*App) SchemaArtifact()    {}
func (*App) Definition() ([]byte, []int) {
	return filedescriptor_f9b42966edc5edad, []int{0}
}
func (m *App) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *App) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Application.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *App) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Application.Merge(m, src)
}
func (m *App) XXX_Extent() int {
	return m.Extent()
}
func (m *App) XXX_Dropunfamiliar() {
	xxx_signaldetails_Application.DiscardUnknown(m)
}

var xxx_signaldetails_Application proto.InternalMessageInfo

func (m *App) ObtainScheme() uint64 {
	if m != nil {
		return m.Scheme
	}
	return 0
}

func (m *App) ObtainPackage() string {
	if m != nil {
		return m.Package
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
func (m *Agreement) Text() string { return proto.CompactTextString(m) }
func (*Agreement) SchemaArtifact()    {}
func (*Agreement) Definition() ([]byte, []int) {
	return filedescriptor_f9b42966edc5edad, []int{1}
}
func (m *Agreement) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Agreement) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Agreement.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Agreement) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Agreement.Merge(m, src)
}
func (m *Agreement) XXX_Extent() int {
	return m.Extent()
}
func (m *Agreement) XXX_Dropunfamiliar() {
	xxx_signaldetails_Agreement.DiscardUnknown(m)
}

var xxx_signaldetails_Agreement proto.InternalMessageInfo

func (m *Agreement) ObtainLedger() uint64 {
	if m != nil {
		return m.Ledger
	}
	return 0
}

func (m *Agreement) ObtainApplication() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func initialize() {
	proto.RegisterType((*App)(nil), "REDACTED")
	proto.RegisterType((*Agreement)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_f9b42966edc5edad) }

var filedescriptor_f9b42966edc5edad = []byte{
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

func (that *Agreement) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*Agreement)
	if !ok {
		which2, ok := which.(Agreement)
		if ok {
			which1 = &which2
		} else {
			return false
		}
	}
	if which1 == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if that.Ledger != which1.Ledger {
		return false
	}
	if that.App != which1.App {
		return false
	}
	return true
}
func (m *App) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *App) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *App) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Package) > 0 {
		i -= len(m.Package)
		copy(deltaLocatedAN[i:], m.Package)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Package)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	if m.Scheme != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Scheme))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Agreement) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Agreement) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Agreement) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.App))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Ledger != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ledger))
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
func (m *App) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Scheme != 0 {
		n += 1 + sovKinds(uint64(m.Scheme))
	}
	l = len(m.Package)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *Agreement) Extent() (n int) {
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
func (m *App) Decode(deltaLocatedAN []byte) error {
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
			m.Scheme = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Scheme |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
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
				return FaultUnfitMagnitudeKinds
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Package = string(deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *Agreement) Decode(deltaLocatedAN []byte) error {
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
			m.Ledger = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Ledger |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.App = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.App |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
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
