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

type PeerxSolicit struct {
}

func (m *PeerxSolicit) Restore()         { *m = PeerxSolicit{} }
func (m *PeerxSolicit) Text() string { return proto.CompactTextString(m) }
func (*PeerxSolicit) SchemaArtifact()    {}
func (*PeerxSolicit) Definition() ([]byte, []int) {
	return filedescriptor_81c2f011fd13be57, []int{0}
}
func (m *PeerxSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PeerxSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Peerxsolicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerxSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Peerxsolicit.Merge(m, src)
}
func (m *PeerxSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *PeerxSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Peerxsolicit.DiscardUnknown(m)
}

var xxx_signaldetails_Peerxsolicit proto.InternalMessageInfo

type PeerxLocations struct {
	Locations []NetworkLocator `protobuf:"octets,1,rep,name=addrs,proto3" json:"locations"`
}

func (m *PeerxLocations) Restore()         { *m = PeerxLocations{} }
func (m *PeerxLocations) Text() string { return proto.CompactTextString(m) }
func (*PeerxLocations) SchemaArtifact()    {}
func (*PeerxLocations) Definition() ([]byte, []int) {
	return filedescriptor_81c2f011fd13be57, []int{1}
}
func (m *PeerxLocations) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *PeerxLocations) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Peerxlocations.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerxLocations) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Peerxlocations.Merge(m, src)
}
func (m *PeerxLocations) XXX_Extent() int {
	return m.Extent()
}
func (m *PeerxLocations) XXX_Dropunfamiliar() {
	xxx_signaldetails_Peerxlocations.DiscardUnknown(m)
}

var xxx_signaldetails_Peerxlocations proto.InternalMessageInfo

func (m *PeerxLocations) ObtainLocations() []NetworkLocator {
	if m != nil {
		return m.Locations
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
	return filedescriptor_81c2f011fd13be57, []int{2}
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

type Artifact_Peerxsolicit struct {
	PeerxSolicit *PeerxSolicit `protobuf:"octets,1,opt,name=pex_request,json=pexRequest,proto3,oneof" json:"peerx_solicit,omitempty"`
}
type Artifact_Peerxlocations struct {
	PeerxLocations *PeerxLocations `protobuf:"octets,2,opt,name=pex_addrs,json=pexAddrs,proto3,oneof" json:"peerx_locations,omitempty"`
}

func (*Artifact_Peerxsolicit) isnote_Total() {}
func (*Artifact_Peerxlocations) isnote_Total()   {}

func (m *Signal) ObtainTotal() isnote_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) ObtainPeerxSolicit() *PeerxSolicit {
	if x, ok := m.ObtainTotal().(*Artifact_Peerxsolicit); ok {
		return x.PeerxSolicit
	}
	return nil
}

func (m *Signal) ObtainPeerxLocations() *PeerxLocations {
	if x, ok := m.ObtainTotal().(*Artifact_Peerxlocations); ok {
		return x.PeerxLocations
	}
	return nil
}

//
func (*Signal) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Artifact_Peerxsolicit)(nil),
		(*Artifact_Peerxlocations)(nil),
	}
}

func initialize() {
	proto.RegisterType((*PeerxSolicit)(nil), "REDACTED")
	proto.RegisterType((*PeerxLocations)(nil), "REDACTED")
	proto.RegisterType((*Signal)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_81c2f011fd13be57) }

var filedescriptor_81c2f011fd13be57 = []byte{
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

func (m *PeerxSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PeerxSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PeerxSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *PeerxLocations) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *PeerxLocations) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *PeerxLocations) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Locations) > 0 {
		for idxNdExc := len(m.Locations) - 1; idxNdExc >= 0; idxNdExc-- {
			{
				extent, err := m.Locations[idxNdExc].SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
				if err != nil {
					return 0, err
				}
				i -= extent
				i = encodeVariableintPeerx(deltaLocatedAN, i, uint64(extent))
			}
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

func (m *Artifact_Peerxsolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Peerxsolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PeerxSolicit != nil {
		{
			extent, err := m.PeerxSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintPeerx(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Artifact_Peerxlocations) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Peerxlocations) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.PeerxLocations != nil {
		{
			extent, err := m.PeerxLocations.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintPeerx(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	return len(deltaLocatedAN) - i, nil
}
func encodeVariableintPeerx(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovPeerx(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *PeerxSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PeerxLocations) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Locations) > 0 {
		for _, e := range m.Locations {
			l = e.Extent()
			n += 1 + l + sovPeerx(uint64(l))
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

func (m *Artifact_Peerxsolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeerxSolicit != nil {
		l = m.PeerxSolicit.Extent()
		n += 1 + l + sovPeerx(uint64(l))
	}
	return n
}
func (m *Artifact_Peerxlocations) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeerxLocations != nil {
		l = m.PeerxLocations.Extent()
		n += 1 + l + sovPeerx(uint64(l))
	}
	return n
}

func sovPeerx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPeerx(x uint64) (n int) {
	return sovPeerx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PeerxSolicit) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunPeerx
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
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitPeerx(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudePeerx
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
func (m *PeerxLocations) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunPeerx
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
					return FaultIntegerOverrunPeerx
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
				return FaultUnfitMagnitudePeerx
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudePeerx
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Locations = append(m.Locations, NetworkLocator{})
			if err := m.Locations[len(m.Locations)-1].Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitPeerx(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudePeerx
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
				return FaultIntegerOverrunPeerx
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
					return FaultIntegerOverrunPeerx
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
				return FaultUnfitMagnitudePeerx
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudePeerx
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PeerxSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Peerxsolicit{v}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunPeerx
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
				return FaultUnfitMagnitudePeerx
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudePeerx
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &PeerxLocations{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Peerxlocations{v}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitPeerx(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudePeerx
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
func omitPeerx(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunPeerx
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
					return 0, FaultIntegerOverrunPeerx
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
					return 0, FaultIntegerOverrunPeerx
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
				return 0, FaultUnfitMagnitudePeerx
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionPeerx
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudePeerx
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudePeerx        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunPeerx          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionPeerx = fmt.Errorf("REDACTED")
)
