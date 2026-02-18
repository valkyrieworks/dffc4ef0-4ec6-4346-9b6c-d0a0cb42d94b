//
//

package kinds

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

type Ledger struct {
	Heading     Heading       `protobuf:"octets,1,opt,name=header,proto3" json:"heading"`
	Data       Data         `protobuf:"octets,2,opt,name=data,proto3" json:"data"`
	Proof   ProofCatalog `protobuf:"octets,3,opt,name=evidence,proto3" json:"proof"`
	FinalEndorse *Endorse      `protobuf:"octets,4,opt,name=last_commit,json=lastCommit,proto3" json:"final_endorse,omitempty"`
}

func (m *Ledger) Restore()         { *m = Ledger{} }
func (m *Ledger) String() string { return proto.CompactTextString(m) }
func (*Ledger) SchemaSignal()    {}
func (*Ledger) Definition() ([]byte, []int) {
	return filedefinition_70840e82f4357ab1, []int{0}
}
func (m *Ledger) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Ledger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Ledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ledger) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Ledger.Merge(m, src)
}
func (m *Ledger) XXX_Volume() int {
	return m.Volume()
}
func (m *Ledger) XXX_Omitunclear() {
	xxx_messagedata_Ledger.DiscardUnknown(m)
}

var xxx_messagedata_Ledger proto.InternalMessageInfo

func (m *Ledger) FetchHeading() Heading {
	if m != nil {
		return m.Heading
	}
	return Heading{}
}

func (m *Ledger) FetchData() Data {
	if m != nil {
		return m.Data
	}
	return Data{}
}

func (m *Ledger) FetchProof() ProofCatalog {
	if m != nil {
		return m.Proof
	}
	return ProofCatalog{}
}

func (m *Ledger) FetchFinalEndorse() *Endorse {
	if m != nil {
		return m.FinalEndorse
	}
	return nil
}

func init() {
	proto.RegisterType((*Ledger)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_70840e82f4357ab1) }

var filedefinition_70840e82f4357ab1 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0xca, 0xc9,
	0x4f, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0xc8, 0xea, 0x81, 0x65, 0xa5,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44, 0x9d, 0x14, 0xa6, 0x29, 0x60,
	0x12, 0x2a, 0x2b, 0x8f, 0x21, 0x9b, 0x5a, 0x96, 0x99, 0x92, 0x9a, 0x97, 0x9c, 0x0a, 0x51, 0xa0,
	0xf4, 0x8e, 0x91, 0x8b, 0xd5, 0x09, 0x64, 0xad, 0x90, 0x19, 0x17, 0x5b, 0x46, 0x6a, 0x62, 0x4a,
	0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x84, 0x1e, 0xba, 0x0b, 0xf4, 0x3c, 0xc0,
	0xf2, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0x41, 0x55, 0x0b, 0x19, 0x70, 0xb1, 0xa4, 0x24,
	0x96, 0x24, 0x4a, 0x30, 0x81, 0x75, 0x89, 0x61, 0xea, 0x72, 0x49, 0x2c, 0x49, 0x84, 0xea, 0x01,
	0xab, 0x14, 0x72, 0xe0, 0xe2, 0x80, 0xb9, 0x42, 0x82, 0x19, 0xac, 0x4b, 0x0e, 0x53, 0x97, 0x2b,
	0x54, 0x85, 0x4f, 0x66, 0x71, 0x09, 0x54, 0x37, 0x5c, 0x97, 0x90, 0x25, 0x17, 0x77, 0x4e, 0x62,
	0x71, 0x49, 0x7c, 0x72, 0x7e, 0x6e, 0x6e, 0x66, 0x89, 0x04, 0x0b, 0x2e, 0x07, 0x3b, 0x83, 0xe5,
	0x83, 0xb8, 0x40, 0x8a, 0x21, 0x6c, 0x27, 0xdf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce,
	0xcf, 0x4d, 0x2d, 0x49, 0x4a, 0x2b, 0x41, 0x30, 0x20, 0x01, 0x8f, 0x1e, 0x9c, 0x49, 0x6c, 0x60,
	0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0xdf, 0xde, 0x0a, 0xcd, 0x01, 0x00, 0x00,
}

func (m *Ledger) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ledger) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Ledger) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalEndorse != nil {
		{
			volume, err := m.FinalEndorse.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintLedger(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		volume, err := m.Proof.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = encodeVariableintLedger(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x1a
	{
		volume, err := m.Data.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = encodeVariableintLedger(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0x12
	{
		volume, err := m.Heading.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = encodeVariableintLedger(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVariableintLedger(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovLedger(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *Ledger) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Heading.Volume()
	n += 1 + l + sovLedger(uint64(l))
	l = m.Data.Volume()
	n += 1 + l + sovLedger(uint64(l))
	l = m.Proof.Volume()
	n += 1 + l + sovLedger(uint64(l))
	if m.FinalEndorse != nil {
		l = m.FinalEndorse.Volume()
		n += 1 + l + sovLedger(uint64(l))
	}
	return n
}

func sovLedger(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLedger(x uint64) (n int) {
	return sovLedger(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ledger) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadLedger
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
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLedger
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentLedger
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Heading.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLedger
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentLedger
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLedger
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentLedger
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proof.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadLedger
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentLedger
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.FinalEndorse == nil {
				m.FinalEndorse = &Endorse{}
			}
			if err := m.FinalEndorse.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitLedger(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentLedger
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
func omitLedger(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadLedger
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
					return 0, ErrIntegerOverloadLedger
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
					return 0, ErrIntegerOverloadLedger
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
				return 0, ErrCorruptExtentLedger
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterLedger
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentLedger
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentLedger        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadLedger          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterLedger = fmt.Errorf("REDACTED")
)
