//
//

package security

import (
	bytes "bytes"
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
type CommonToken struct {
	//
	//
	//
	//
	//
	Sum iscommonkey_Total `protobuf_oneof:"sum"`
}

func (m *CommonToken) Restore()         { *m = CommonToken{} }
func (m *CommonToken) Text() string { return proto.CompactTextString(m) }
func (*CommonToken) SchemaArtifact()    {}
func (*CommonToken) Definition() ([]byte, []int) {
	return filedescriptor_cb048658b234868c, []int{0}
}
func (m *CommonToken) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *CommonToken) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Commonkey.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonToken) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Commonkey.Merge(m, src)
}
func (m *CommonToken) XXX_Extent() int {
	return m.Extent()
}
func (m *CommonToken) XXX_Dropunfamiliar() {
	xxx_signaldetails_Commonkey.DiscardUnknown(m)
}

var xxx_signaldetails_Commonkey proto.InternalMessageInfo

type iscommonkey_Total interface {
	iscommonkey_Total()
	Equivalent(interface{}) bool
	SerializeToward([]byte) (int, error)
	Extent() int
	Contrast(interface{}) int
}

type Commonkey_Edwards25519 struct {
	Edwards25519 []byte `protobuf:"octets,1,opt,name=ed25519,proto3,oneof" json:"edwards25519,omitempty"`
}
type Commonkey_Ellipticp256 struct {
	Ellipticp256 []byte `protobuf:"octets,2,opt,name=secp256k1,proto3,oneof" json:"ellipticp256,omitempty"`
}
type Commonkey_Signature381 struct {
	Signature381 []byte `protobuf:"octets,3,opt,name=bls12381,proto3,oneof" json:"signature381,omitempty"`
}

func (*Commonkey_Edwards25519) iscommonkey_Total()   {}
func (*Commonkey_Ellipticp256) iscommonkey_Total() {}
func (*Commonkey_Signature381) iscommonkey_Total()  {}

func (m *CommonToken) ObtainTotal() iscommonkey_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *CommonToken) ObtainEdwards25519() []byte {
	if x, ok := m.ObtainTotal().(*Commonkey_Edwards25519); ok {
		return x.Edwards25519
	}
	return nil
}

func (m *CommonToken) ObtainEllipticp256() []byte {
	if x, ok := m.ObtainTotal().(*Commonkey_Ellipticp256); ok {
		return x.Ellipticp256
	}
	return nil
}

func (m *CommonToken) ObtainSignature381() []byte {
	if x, ok := m.ObtainTotal().(*Commonkey_Signature381); ok {
		return x.Signature381
	}
	return nil
}

//
func (*CommonToken) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Commonkey_Edwards25519)(nil),
		(*Commonkey_Ellipticp256)(nil),
		(*Commonkey_Signature381)(nil),
	}
}

func initialize() {
	proto.RegisterType((*CommonToken)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_cb048658b234868c) }

var filedescriptor_cb048658b234868c = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xd7, 0xcf, 0x4e,
	0xad, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x44, 0xc8, 0xea, 0x41, 0x64, 0xa5,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44, 0xa1, 0x52, 0x19, 0x17, 0x67,
	0x40, 0x69, 0x52, 0x4e, 0x66, 0xb2, 0x77, 0x6a, 0xa5, 0x90, 0x14, 0x17, 0x7b, 0x6a, 0x8a, 0x91,
	0xa9, 0xa9, 0xa1, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x07, 0x43, 0x10, 0x4c, 0x40, 0x48,
	0x8e, 0x8b, 0xb3, 0x38, 0x35, 0xb9, 0xc0, 0xc8, 0xd4, 0x2c, 0xdb, 0x50, 0x82, 0x09, 0x2a, 0x8b,
	0x10, 0x12, 0x92, 0xe1, 0xe2, 0x48, 0xca, 0x29, 0x36, 0x34, 0x32, 0xb6, 0x30, 0x94, 0x60, 0x86,
	0x4a, 0xc3, 0x45, 0xac, 0x38, 0x5e, 0x2c, 0x90, 0x67, 0x7c, 0xb1, 0x50, 0x9e, 0xd1, 0x89, 0x95,
	0x8b, 0xb9, 0xb8, 0x34, 0xd7, 0xc9, 0xef, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x4c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x73,
	0x53, 0x4b, 0x92, 0xd2, 0x4a, 0x10, 0x0c, 0x88, 0x07, 0x30, 0xfc, 0x9e, 0xc4, 0x06, 0x96, 0x30,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x32, 0x1d, 0x68, 0x17, 0x01, 0x00, 0x00,
}

func (that *CommonToken) Contrast(which interface{}) int {
	if which == nil {
		if that == nil {
			return 0
		}
		return 1
	}

	which1, ok := which.(*CommonToken)
	if !ok {
		which2, ok := which.(CommonToken)
		if ok {
			which1 = &which2
		} else {
			return 1
		}
	}
	if which1 == nil {
		if that == nil {
			return 0
		}
		return 1
	} else if that == nil {
		return -1
	}
	if which1.Sum == nil {
		if that.Sum != nil {
			return 1
		}
	} else if that.Sum == nil {
		return -1
	} else {
		thatKind := -1
		switch that.Sum.(type) {
		case *Commonkey_Edwards25519:
			thatKind = 0
		case *Commonkey_Ellipticp256:
			thatKind = 1
		case *Commonkey_Signature381:
			thatKind = 2
		default:
			panic(fmt.Sprintf("REDACTED", that.Sum))
		}
		which1kind := -1
		switch which1.Sum.(type) {
		case *Commonkey_Edwards25519:
			which1kind = 0
		case *Commonkey_Ellipticp256:
			which1kind = 1
		case *Commonkey_Signature381:
			which1kind = 2
		default:
			panic(fmt.Sprintf("REDACTED", which1.Sum))
		}
		if thatKind == which1kind {
			if c := that.Sum.Contrast(which1.Sum); c != 0 {
				return c
			}
		} else if thatKind < which1kind {
			return -1
		} else if thatKind > which1kind {
			return 1
		}
	}
	return 0
}
func (that *Commonkey_Edwards25519) Contrast(which interface{}) int {
	if which == nil {
		if that == nil {
			return 0
		}
		return 1
	}

	which1, ok := which.(*Commonkey_Edwards25519)
	if !ok {
		which2, ok := which.(Commonkey_Edwards25519)
		if ok {
			which1 = &which2
		} else {
			return 1
		}
	}
	if which1 == nil {
		if that == nil {
			return 0
		}
		return 1
	} else if that == nil {
		return -1
	}
	if c := bytes.Compare(that.Edwards25519, which1.Edwards25519); c != 0 {
		return c
	}
	return 0
}
func (that *Commonkey_Ellipticp256) Contrast(which interface{}) int {
	if which == nil {
		if that == nil {
			return 0
		}
		return 1
	}

	which1, ok := which.(*Commonkey_Ellipticp256)
	if !ok {
		which2, ok := which.(Commonkey_Ellipticp256)
		if ok {
			which1 = &which2
		} else {
			return 1
		}
	}
	if which1 == nil {
		if that == nil {
			return 0
		}
		return 1
	} else if that == nil {
		return -1
	}
	if c := bytes.Compare(that.Ellipticp256, which1.Ellipticp256); c != 0 {
		return c
	}
	return 0
}
func (that *Commonkey_Signature381) Contrast(which interface{}) int {
	if which == nil {
		if that == nil {
			return 0
		}
		return 1
	}

	which1, ok := which.(*Commonkey_Signature381)
	if !ok {
		which2, ok := which.(Commonkey_Signature381)
		if ok {
			which1 = &which2
		} else {
			return 1
		}
	}
	if which1 == nil {
		if that == nil {
			return 0
		}
		return 1
	} else if that == nil {
		return -1
	}
	if c := bytes.Compare(that.Signature381, which1.Signature381); c != 0 {
		return c
	}
	return 0
}
func (that *CommonToken) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*CommonToken)
	if !ok {
		which2, ok := which.(CommonToken)
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
	if which1.Sum == nil {
		if that.Sum != nil {
			return false
		}
	} else if that.Sum == nil {
		return false
	} else if !that.Sum.Equivalent(which1.Sum) {
		return false
	}
	return true
}
func (that *Commonkey_Edwards25519) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*Commonkey_Edwards25519)
	if !ok {
		which2, ok := which.(Commonkey_Edwards25519)
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
	if !bytes.Equal(that.Edwards25519, which1.Edwards25519) {
		return false
	}
	return true
}
func (that *Commonkey_Ellipticp256) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*Commonkey_Ellipticp256)
	if !ok {
		which2, ok := which.(Commonkey_Ellipticp256)
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
	if !bytes.Equal(that.Ellipticp256, which1.Ellipticp256) {
		return false
	}
	return true
}
func (that *Commonkey_Signature381) Equivalent(which interface{}) bool {
	if which == nil {
		return that == nil
	}

	which1, ok := which.(*Commonkey_Signature381)
	if !ok {
		which2, ok := which.(Commonkey_Signature381)
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
	if !bytes.Equal(that.Signature381, which1.Signature381) {
		return false
	}
	return true
}
func (m *CommonToken) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *CommonToken) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *CommonToken) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
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

func (m *Commonkey_Edwards25519) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Commonkey_Edwards25519) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Edwards25519 != nil {
		i -= len(m.Edwards25519)
		copy(deltaLocatedAN[i:], m.Edwards25519)
		i = encodeVariableintTokens(deltaLocatedAN, i, uint64(len(m.Edwards25519)))
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Commonkey_Ellipticp256) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Commonkey_Ellipticp256) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Ellipticp256 != nil {
		i -= len(m.Ellipticp256)
		copy(deltaLocatedAN[i:], m.Ellipticp256)
		i = encodeVariableintTokens(deltaLocatedAN, i, uint64(len(m.Ellipticp256)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Commonkey_Signature381) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Commonkey_Signature381) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.Signature381 != nil {
		i -= len(m.Signature381)
		copy(deltaLocatedAN[i:], m.Signature381)
		i = encodeVariableintTokens(deltaLocatedAN, i, uint64(len(m.Signature381)))
		i--
		deltaLocatedAN[i] = 0x1a
	}
	return len(deltaLocatedAN) - i, nil
}
func encodeVariableintTokens(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovTokens(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *CommonToken) Extent() (n int) {
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

func (m *Commonkey_Edwards25519) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Edwards25519 != nil {
		l = len(m.Edwards25519)
		n += 1 + l + sovTokens(uint64(l))
	}
	return n
}
func (m *Commonkey_Ellipticp256) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ellipticp256 != nil {
		l = len(m.Ellipticp256)
		n += 1 + l + sovTokens(uint64(l))
	}
	return n
}
func (m *Commonkey_Signature381) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Signature381 != nil {
		l = len(m.Signature381)
		n += 1 + l + sovTokens(uint64(l))
	}
	return n
}

func sovTokens(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokens(x uint64) (n int) {
	return sovTokens(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonToken) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunTokens
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
					return FaultIntegerOverrunTokens
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
				return FaultUnfitMagnitudeTokens
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeTokens
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, submitOrdinal-idxNdExc)
			copy(v, deltaLocatedAN[idxNdExc:submitOrdinal])
			m.Sum = &Commonkey_Edwards25519{v}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunTokens
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
				return FaultUnfitMagnitudeTokens
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeTokens
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, submitOrdinal-idxNdExc)
			copy(v, deltaLocatedAN[idxNdExc:submitOrdinal])
			m.Sum = &Commonkey_Ellipticp256{v}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunTokens
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
				return FaultUnfitMagnitudeTokens
			}
			submitOrdinal := idxNdExc + octetSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeTokens
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, submitOrdinal-idxNdExc)
			copy(v, deltaLocatedAN[idxNdExc:submitOrdinal])
			m.Sum = &Commonkey_Signature381{v}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitTokens(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeTokens
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
func omitTokens(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunTokens
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
					return 0, FaultIntegerOverrunTokens
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
					return 0, FaultIntegerOverrunTokens
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
				return 0, FaultUnfitMagnitudeTokens
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionTokens
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeTokens
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeTokens        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunTokens          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionTokens = fmt.Errorf("REDACTED")
)
