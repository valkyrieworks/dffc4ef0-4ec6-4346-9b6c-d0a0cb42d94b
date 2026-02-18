//
//

package vault

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
type PublicKey struct {
	//
	//
	//
	//
	//
	Sum ispublickey_Total `protobuf_oneof:"sum"`
}

func (m *PublicKey) Restore()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) SchemaSignal()    {}
func (*PublicKey) Definition() ([]byte, []int) {
	return filedefinition_cb048658b234868c, []int{0}
}
func (m *PublicKey) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *PublicKey) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Publickey.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKey) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Publickey.Merge(m, src)
}
func (m *PublicKey) XXX_Volume() int {
	return m.Volume()
}
func (m *PublicKey) XXX_Omitunclear() {
	xxx_messagedata_Publickey.DiscardUnknown(m)
}

var xxx_messagedata_Publickey proto.InternalMessageInfo

type ispublickey_Total interface {
	ispublickey_Total()
	Equivalent(interface{}) bool
	SerializeTo([]byte) (int, error)
	Volume() int
	Contrast(interface{}) int
}

type Publickey_Ed25519 struct {
	Ed25519 []byte `protobuf:"octets,1,opt,name=ed25519,proto3,oneof" json:"ed25519,omitempty"`
}
type Publickey_Secp256k1 struct {
	Secp256k1 []byte `protobuf:"octets,2,opt,name=secp256k1,proto3,oneof" json:"secp256k1,omitempty"`
}
type Publickey_Bls12381 struct {
	Bls12381 []byte `protobuf:"octets,3,opt,name=bls12381,proto3,oneof" json:"bls12381,omitempty"`
}

func (*Publickey_Ed25519) ispublickey_Total()   {}
func (*Publickey_Secp256k1) ispublickey_Total() {}
func (*Publickey_Bls12381) ispublickey_Total()  {}

func (m *PublicKey) FetchTotal() ispublickey_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *PublicKey) FetchEd25519() []byte {
	if x, ok := m.FetchTotal().(*Publickey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) FetchSecp256k1() []byte {
	if x, ok := m.FetchTotal().(*Publickey_Secp256k1); ok {
		return x.Secp256k1
	}
	return nil
}

func (m *PublicKey) FetchBls12381() []byte {
	if x, ok := m.FetchTotal().(*Publickey_Bls12381); ok {
		return x.Bls12381
	}
	return nil
}

//
func (*PublicKey) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Publickey_Ed25519)(nil),
		(*Publickey_Secp256k1)(nil),
		(*Publickey_Bls12381)(nil),
	}
}

func init() {
	proto.RegisterType((*PublicKey)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_cb048658b234868c) }

var filedefinition_cb048658b234868c = []byte{
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

func (this *PublicKey) Contrast(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*PublicKey)
	if !ok {
		that2, ok := that.(PublicKey)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if that1.Sum == nil {
		if this.Sum != nil {
			return 1
		}
	} else if this.Sum == nil {
		return -1
	} else {
		thisKind := -1
		switch this.Sum.(type) {
		case *Publickey_Ed25519:
			thisKind = 0
		case *Publickey_Secp256k1:
			thisKind = 1
		case *Publickey_Bls12381:
			thisKind = 2
		default:
			panic(fmt.Sprintf("REDACTED", this.Sum))
		}
		that1kind := -1
		switch that1.Sum.(type) {
		case *Publickey_Ed25519:
			that1kind = 0
		case *Publickey_Secp256k1:
			that1kind = 1
		case *Publickey_Bls12381:
			that1kind = 2
		default:
			panic(fmt.Sprintf("REDACTED", that1.Sum))
		}
		if thisKind == that1kind {
			if c := this.Sum.Contrast(that1.Sum); c != 0 {
				return c
			}
		} else if thisKind < that1kind {
			return -1
		} else if thisKind > that1kind {
			return 1
		}
	}
	return 0
}
func (this *Publickey_Ed25519) Contrast(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Publickey_Ed25519)
	if !ok {
		that2, ok := that.(Publickey_Ed25519)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := bytes.Compare(this.Ed25519, that1.Ed25519); c != 0 {
		return c
	}
	return 0
}
func (this *Publickey_Secp256k1) Contrast(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Publickey_Secp256k1)
	if !ok {
		that2, ok := that.(Publickey_Secp256k1)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := bytes.Compare(this.Secp256k1, that1.Secp256k1); c != 0 {
		return c
	}
	return 0
}
func (this *Publickey_Bls12381) Contrast(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Publickey_Bls12381)
	if !ok {
		that2, ok := that.(Publickey_Bls12381)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := bytes.Compare(this.Bls12381, that1.Bls12381); c != 0 {
		return c
	}
	return 0
}
func (this *PublicKey) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PublicKey)
	if !ok {
		that2, ok := that.(PublicKey)
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
	if that1.Sum == nil {
		if this.Sum != nil {
			return false
		}
	} else if this.Sum == nil {
		return false
	} else if !this.Sum.Equivalent(that1.Sum) {
		return false
	}
	return true
}
func (this *Publickey_Ed25519) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Publickey_Ed25519)
	if !ok {
		that2, ok := that.(Publickey_Ed25519)
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
	if !bytes.Equal(this.Ed25519, that1.Ed25519) {
		return false
	}
	return true
}
func (this *Publickey_Secp256k1) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Publickey_Secp256k1)
	if !ok {
		that2, ok := that.(Publickey_Secp256k1)
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
	if !bytes.Equal(this.Secp256k1, that1.Secp256k1) {
		return false
	}
	return true
}
func (this *Publickey_Bls12381) Equivalent(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Publickey_Bls12381)
	if !ok {
		that2, ok := that.(Publickey_Bls12381)
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
	if !bytes.Equal(this.Bls12381, that1.Bls12381) {
		return false
	}
	return true
}
func (m *PublicKey) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKey) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *PublicKey) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			volume := m.Sum.Volume()
			i -= volume
			if _, err := m.Sum.SerializeTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Publickey_Ed25519) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Publickey_Ed25519) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ed25519 != nil {
		i -= len(m.Ed25519)
		copy(dAtA[i:], m.Ed25519)
		i = encodeVariableintKeys(dAtA, i, uint64(len(m.Ed25519)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Publickey_Secp256k1) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Publickey_Secp256k1) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Secp256k1 != nil {
		i -= len(m.Secp256k1)
		copy(dAtA[i:], m.Secp256k1)
		i = encodeVariableintKeys(dAtA, i, uint64(len(m.Secp256k1)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Publickey_Bls12381) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Publickey_Bls12381) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Bls12381 != nil {
		i -= len(m.Bls12381)
		copy(dAtA[i:], m.Bls12381)
		i = encodeVariableintKeys(dAtA, i, uint64(len(m.Bls12381)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVariableintKeys(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovKeys(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *PublicKey) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Volume()
	}
	return n
}

func (m *Publickey_Ed25519) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ed25519 != nil {
		l = len(m.Ed25519)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *Publickey_Secp256k1) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Secp256k1 != nil {
		l = len(m.Secp256k1)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func (m *Publickey_Bls12381) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bls12381 != nil {
		l = len(m.Bls12381)
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeys(x uint64) (n int) {
	return sovKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublicKey) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKeys
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
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKeys
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				octetSize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if octetSize < 0 {
				return ErrCorruptExtentKeys
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKeys
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, submitOrdinal-idxNdEx)
			copy(v, dAtA[idxNdEx:submitOrdinal])
			m.Sum = &Publickey_Ed25519{v}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKeys
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				octetSize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if octetSize < 0 {
				return ErrCorruptExtentKeys
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKeys
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, submitOrdinal-idxNdEx)
			copy(v, dAtA[idxNdEx:submitOrdinal])
			m.Sum = &Publickey_Secp256k1{v}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKeys
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				octetSize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if octetSize < 0 {
				return ErrCorruptExtentKeys
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKeys
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, submitOrdinal-idxNdEx)
			copy(v, dAtA[idxNdEx:submitOrdinal])
			m.Sum = &Publickey_Bls12381{v}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKeys(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKeys
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
func omitKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadKeys
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
					return 0, ErrIntegerOverloadKeys
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
					return 0, ErrIntegerOverloadKeys
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
				return 0, ErrCorruptExtentKeys
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterKeys
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentKeys
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentKeys        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadKeys          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterKeys = fmt.Errorf("REDACTED")
)
