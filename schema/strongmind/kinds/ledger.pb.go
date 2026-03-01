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
func (m *Ledger) Text() string { return proto.CompactTextString(m) }
func (*Ledger) SchemaArtifact()    {}
func (*Ledger) Definition() ([]byte, []int) {
	return filedescriptor_70840e82f4357ab1, []int{0}
}
func (m *Ledger) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *Ledger) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Ledger.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ledger) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Ledger.Merge(m, src)
}
func (m *Ledger) XXX_Extent() int {
	return m.Extent()
}
func (m *Ledger) XXX_Dropunfamiliar() {
	xxx_signaldetails_Ledger.DiscardUnknown(m)
}

var xxx_signaldetails_Ledger proto.InternalMessageInfo

func (m *Ledger) ObtainHeadline() Heading {
	if m != nil {
		return m.Heading
	}
	return Heading{}
}

func (m *Ledger) ObtainData() Data {
	if m != nil {
		return m.Data
	}
	return Data{}
}

func (m *Ledger) ObtainProof() ProofCatalog {
	if m != nil {
		return m.Proof
	}
	return ProofCatalog{}
}

func (m *Ledger) ObtainFinalEndorse() *Endorse {
	if m != nil {
		return m.FinalEndorse
	}
	return nil
}

func initialize() {
	proto.RegisterType((*Ledger)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_70840e82f4357ab1) }

var filedescriptor_70840e82f4357ab1 = []byte{
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

func (m *Ledger) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *Ledger) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Ledger) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.FinalEndorse != nil {
		{
			extent, err := m.FinalEndorse.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintLedger(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x22
	}
	{
		extent, err := m.Proof.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = encodeVariableintLedger(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x1a
	{
		extent, err := m.Data.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = encodeVariableintLedger(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0x12
	{
		extent, err := m.Heading.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = encodeVariableintLedger(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintLedger(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovLedger(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *Ledger) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Heading.Extent()
	n += 1 + l + sovLedger(uint64(l))
	l = m.Data.Extent()
	n += 1 + l + sovLedger(uint64(l))
	l = m.Proof.Extent()
	n += 1 + l + sovLedger(uint64(l))
	if m.FinalEndorse != nil {
		l = m.FinalEndorse.Extent()
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
func (m *Ledger) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunLedger
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
					return FaultIntegerOverrunLedger
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
				return FaultUnfitMagnitudeLedger
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Heading.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLedger
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
				return FaultUnfitMagnitudeLedger
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLedger
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
				return FaultUnfitMagnitudeLedger
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proof.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunLedger
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
				return FaultUnfitMagnitudeLedger
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeLedger
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.FinalEndorse == nil {
				m.FinalEndorse = &Endorse{}
			}
			if err := m.FinalEndorse.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitLedger(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeLedger
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
func omitLedger(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunLedger
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
					return 0, FaultIntegerOverrunLedger
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
					return 0, FaultIntegerOverrunLedger
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
				return 0, FaultUnfitMagnitudeLedger
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionLedger
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeLedger
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeLedger        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunLedger          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionLedger = fmt.Errorf("REDACTED")
)
