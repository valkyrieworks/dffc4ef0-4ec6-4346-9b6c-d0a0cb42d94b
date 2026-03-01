//
//

package txpool

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

type Txs struct {
	Txs [][]byte `protobuf:"octets,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *Txs) Restore()         { *m = Txs{} }
func (m *Txs) Text() string { return proto.CompactTextString(m) }
func (*Txs) SchemaArtifact()    {}
func (*Txs) Definition() ([]byte, []int) {
	return filedescriptor_2af51926fdbcbc05, []int{0}
}
func (m *Txs) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Txs) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Trans.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Txs) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Trans.Merge(m, src)
}
func (m *Txs) XXX_Extent() int {
	return m.Extent()
}
func (m *Txs) XXX_Dropunfamiliar() {
	xxx_signaldetails_Trans.DiscardUnknown(m)
}

var xxx_signaldetails_Trans proto.InternalMessageInfo

func (m *Txs) ObtainTrans() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Signal struct {
	//
	//
	//
	Sum isnote_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) Text() string { return proto.CompactTextString(m) }
func (*Signal) SchemaArtifact()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedescriptor_2af51926fdbcbc05, []int{1}
}
func (m *Signal) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Signal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Artifact.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signal) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Artifact.Merge(m, src)
}
func (m *Signal) XXX_Extent() int {
	return m.Extent()
}
func (m *Signal) XXX_Dropunfamiliar() {
	xxx_signaldetails_Artifact.DiscardUnknown(m)
}

var xxx_signaldetails_Artifact proto.InternalMessageInfo

type isnote_Total interface {
	isnote_Total()
	SerializeToward([]byte) (int, error)
	Extent() int
}

type Artifact_Trans struct {
	Txs *Txs `protobuf:"octets,1,opt,name=txs,proto3,oneof" json:"txs,omitempty"`
}

func (*Artifact_Trans) isnote_Total() {}

func (m *Signal) ObtainTotal() isnote_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) ObtainTrans() *Txs {
	if x, ok := m.ObtainTotal().(*Artifact_Trans); ok {
		return x.Txs
	}
	return nil
}

//
func (*Signal) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Artifact_Trans)(nil),
	}
}

func initialize() {
	proto.RegisterType((*Txs)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_2af51926fdbcbc05) }

var filedescriptor_2af51926fdbcbc05 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0xcf, 0x4d, 0xcd, 0x2d, 0xc8, 0xcf, 0xcf, 0xd1, 0x2f,
	0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0xc8, 0xeb, 0x41,
	0xe5, 0x95, 0xc4, 0xb9, 0x98, 0x43, 0x2a, 0x8a, 0x85, 0x04, 0xb8, 0x98, 0x4b, 0x2a, 0x8a, 0x25,
	0x18, 0x15, 0x98, 0x35, 0x78, 0x82, 0x40, 0x4c, 0x25, 0x5b, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2,
	0xc4, 0xf4, 0x54, 0x21, 0x6d, 0x98, 0x24, 0xa3, 0x06, 0xb7, 0x91, 0xb8, 0x1e, 0xa6, 0x29, 0x7a,
	0x21, 0x15, 0xc5, 0x1e, 0x0c, 0x60, 0x7d, 0x4e, 0xac, 0x5c, 0xcc, 0xc5, 0xa5, 0xb9, 0x4e, 0xfe,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9a, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x9f, 0x9b, 0x5a, 0x92, 0x94, 0x56, 0x82, 0x60,
	0x80, 0x5d, 0xaa, 0x8f, 0xe9, 0x91, 0x24, 0x36, 0xb0, 0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x53, 0xc3, 0xc4, 0x0a, 0xe5, 0x00, 0x00, 0x00,
}

func (m *Txs) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Txs) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Txs) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for idxNdExc := len(m.Txs) - 1; idxNdExc >= 0; idxNdExc-- {
			i -= len(m.Txs[idxNdExc])
			copy(deltaLocatedAN[i:], m.Txs[idxNdExc])
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Txs[idxNdExc])))
			i--
			deltaLocatedAN[i] = 0xa
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Signal) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Signal) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Signal) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			extent := m.Sum.Extent()
			i -= extent
			if _, err := m.Sum.SerializeToward(deltaLocatedAN[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Artifact_Trans) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Trans) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Txs != nil {
		{
			extent, err := m.Txs.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Txs) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovKinds(uint64(l))
		}
	}
	return n
}

func (m *Signal) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Extent()
	}
	return n
}

func (m *Artifact_Trans) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Txs != nil {
		l = m.Txs.Extent()
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
func (m *Txs) Decode(deltaLocatedAN []byte) error {
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
			m.Txs = append(m.Txs, make([]byte, submitOrdinal-idxNdExc))
			copy(m.Txs[len(m.Txs)-1], deltaLocatedAN[idxNdExc:submitOrdinal])
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
func (m *Signal) Decode(deltaLocatedAN []byte) error {
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
			v := &Txs{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Trans{v}
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
