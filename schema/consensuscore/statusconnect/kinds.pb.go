//
//

package statusconnect

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

type Signal struct {
	//
	//
	//
	//
	//
	//
	Sum ismessage_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) SchemaSignal()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedefinition_a1c2869546ca7914, []int{0}
}
func (m *Signal) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *Signal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Signal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signal) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Signal.Merge(m, src)
}
func (m *Signal) XXX_Volume() int {
	return m.Volume()
}
func (m *Signal) XXX_Omitunclear() {
	xxx_messagedata_Signal.DiscardUnknown(m)
}

var xxx_messagedata_Signal proto.InternalMessageInfo

type ismessage_Total interface {
	ismessage_Total()
	SerializeTo([]byte) (int, error)
	Volume() int
}

type Signal_Mirrorsrequest struct {
	MirrorsQuery *MirrorsQuery `protobuf:"octets,1,opt,name=snapshots_request,json=snapshotsRequest,proto3,oneof" json:"mirrors_query,omitempty"`
}
type Signal_Mirrorsreply struct {
	MirrorsReply *MirrorsReply `protobuf:"octets,2,opt,name=snapshots_response,json=snapshotsResponse,proto3,oneof" json:"mirrors_reply,omitempty"`
}
type Signal_Segmentrequest struct {
	SegmentQuery *SegmentQuery `protobuf:"octets,3,opt,name=chunk_request,json=chunkRequest,proto3,oneof" json:"segment_query,omitempty"`
}
type Signal_Segmentreply struct {
	SegmentReply *SegmentReply `protobuf:"octets,4,opt,name=chunk_response,json=chunkResponse,proto3,oneof" json:"segment_reply,omitempty"`
}

func (*Signal_Mirrorsrequest) ismessage_Total()  {}
func (*Signal_Mirrorsreply) ismessage_Total() {}
func (*Signal_Segmentrequest) ismessage_Total()      {}
func (*Signal_Segmentreply) ismessage_Total()     {}

func (m *Signal) FetchTotal() ismessage_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) FetchMirrorsQuery() *MirrorsQuery {
	if x, ok := m.FetchTotal().(*Signal_Mirrorsrequest); ok {
		return x.MirrorsQuery
	}
	return nil
}

func (m *Signal) FetchMirrorsReply() *MirrorsReply {
	if x, ok := m.FetchTotal().(*Signal_Mirrorsreply); ok {
		return x.MirrorsReply
	}
	return nil
}

func (m *Signal) FetchSegmentQuery() *SegmentQuery {
	if x, ok := m.FetchTotal().(*Signal_Segmentrequest); ok {
		return x.SegmentQuery
	}
	return nil
}

func (m *Signal) FetchSegmentReply() *SegmentReply {
	if x, ok := m.FetchTotal().(*Signal_Segmentreply); ok {
		return x.SegmentReply
	}
	return nil
}

//
func (*Signal) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Signal_Mirrorsrequest)(nil),
		(*Signal_Mirrorsreply)(nil),
		(*Signal_Segmentrequest)(nil),
		(*Signal_Segmentreply)(nil),
	}
}

type MirrorsQuery struct {
}

func (m *MirrorsQuery) Restore()         { *m = MirrorsQuery{} }
func (m *MirrorsQuery) String() string { return proto.CompactTextString(m) }
func (*MirrorsQuery) SchemaSignal()    {}
func (*MirrorsQuery) Definition() ([]byte, []int) {
	return filedefinition_a1c2869546ca7914, []int{1}
}
func (m *MirrorsQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *MirrorsQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Mirrorsrequest.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MirrorsQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Mirrorsrequest.Merge(m, src)
}
func (m *MirrorsQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *MirrorsQuery) XXX_Omitunclear() {
	xxx_messagedata_Mirrorsrequest.DiscardUnknown(m)
}

var xxx_messagedata_Mirrorsrequest proto.InternalMessageInfo

type MirrorsReply struct {
	Level   uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Layout   uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Segments   uint32 `protobuf:"variableint,3,opt,name=chunks,proto3" json:"segments,omitempty"`
	Digest     []byte `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Metainfo []byte `protobuf:"octets,5,opt,name=metadata,proto3" json:"metainfo,omitempty"`
}

func (m *MirrorsReply) Restore()         { *m = MirrorsReply{} }
func (m *MirrorsReply) String() string { return proto.CompactTextString(m) }
func (*MirrorsReply) SchemaSignal()    {}
func (*MirrorsReply) Definition() ([]byte, []int) {
	return filedefinition_a1c2869546ca7914, []int{2}
}
func (m *MirrorsReply) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *MirrorsReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Mirrorsreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MirrorsReply) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Mirrorsreply.Merge(m, src)
}
func (m *MirrorsReply) XXX_Volume() int {
	return m.Volume()
}
func (m *MirrorsReply) XXX_Omitunclear() {
	xxx_messagedata_Mirrorsreply.DiscardUnknown(m)
}

var xxx_messagedata_Mirrorsreply proto.InternalMessageInfo

func (m *MirrorsReply) FetchLevel() uint64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *MirrorsReply) FetchLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *MirrorsReply) FetchSegments() uint32 {
	if m != nil {
		return m.Segments
	}
	return 0
}

func (m *MirrorsReply) FetchDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *MirrorsReply) FetchMetainfo() []byte {
	if m != nil {
		return m.Metainfo
	}
	return nil
}

type SegmentQuery struct {
	Level uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Layout uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Ordinal  uint32 `protobuf:"variableint,3,opt,name=index,proto3" json:"ordinal,omitempty"`
}

func (m *SegmentQuery) Restore()         { *m = SegmentQuery{} }
func (m *SegmentQuery) String() string { return proto.CompactTextString(m) }
func (*SegmentQuery) SchemaSignal()    {}
func (*SegmentQuery) Definition() ([]byte, []int) {
	return filedefinition_a1c2869546ca7914, []int{3}
}
func (m *SegmentQuery) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *SegmentQuery) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Segmentrequest.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SegmentQuery) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Segmentrequest.Merge(m, src)
}
func (m *SegmentQuery) XXX_Volume() int {
	return m.Volume()
}
func (m *SegmentQuery) XXX_Omitunclear() {
	xxx_messagedata_Segmentrequest.DiscardUnknown(m)
}

var xxx_messagedata_Segmentrequest proto.InternalMessageInfo

func (m *SegmentQuery) FetchLevel() uint64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SegmentQuery) FetchLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *SegmentQuery) FetchOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type SegmentReply struct {
	Level  uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
	Layout  uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Ordinal   uint32 `protobuf:"variableint,3,opt,name=index,proto3" json:"ordinal,omitempty"`
	Segment   []byte `protobuf:"octets,4,opt,name=chunk,proto3" json:"segment,omitempty"`
	Absent bool   `protobuf:"variableint,5,opt,name=missing,proto3" json:"absent,omitempty"`
}

func (m *SegmentReply) Restore()         { *m = SegmentReply{} }
func (m *SegmentReply) String() string { return proto.CompactTextString(m) }
func (*SegmentReply) SchemaSignal()    {}
func (*SegmentReply) Definition() ([]byte, []int) {
	return filedefinition_a1c2869546ca7914, []int{4}
}
func (m *SegmentReply) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *SegmentReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Segmentreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SegmentReply) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Segmentreply.Merge(m, src)
}
func (m *SegmentReply) XXX_Volume() int {
	return m.Volume()
}
func (m *SegmentReply) XXX_Omitunclear() {
	xxx_messagedata_Segmentreply.DiscardUnknown(m)
}

var xxx_messagedata_Segmentreply proto.InternalMessageInfo

func (m *SegmentReply) FetchLevel() uint64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SegmentReply) FetchLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *SegmentReply) FetchOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *SegmentReply) FetchSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *SegmentReply) FetchAbsent() bool {
	if m != nil {
		return m.Absent
	}
	return false
}

func init() {
	proto.RegisterType((*Signal)(nil), "REDACTED")
	proto.RegisterType((*MirrorsQuery)(nil), "REDACTED")
	proto.RegisterType((*MirrorsReply)(nil), "REDACTED")
	proto.RegisterType((*SegmentQuery)(nil), "REDACTED")
	proto.RegisterType((*SegmentReply)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_a1c2869546ca7914) }

var filedefinition_a1c2869546ca7914 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x3f, 0x6b, 0xdb, 0x40,
	0x1c, 0x95, 0xfc, 0x9f, 0x5f, 0xad, 0x62, 0x1f, 0xa6, 0x88, 0x0e, 0xc2, 0xa8, 0xd0, 0x76, 0x92,
	0xa0, 0x1d, 0xba, 0xbb, 0x8b, 0x0b, 0xed, 0xd0, 0x6b, 0x03, 0x21, 0x4b, 0x38, 0xcb, 0x67, 0x49,
	0x04, 0x9d, 0x14, 0xfd, 0x4e, 0x10, 0x7f, 0x80, 0x4c, 0x59, 0xf2, 0xb1, 0x32, 0x7a, 0x0c, 0x99,
	0x82, 0xfd, 0x45, 0x82, 0x4e, 0xb2, 0xac, 0x38, 0x26, 0x21, 0x90, 0xed, 0xde, 0xd3, 0xd3, 0xbb,
	0xf7, 0x1e, 0x1c, 0x8c, 0x25, 0x17, 0x73, 0x9e, 0x46, 0xa1, 0x90, 0x2e, 0x4a, 0x26, 0x39, 0x2e,
	0x85, 0xe7, 0xca, 0x65, 0xc2, 0xd1, 0x49, 0xd2, 0x58, 0xc6, 0x64, 0xb4, 0x53, 0x38, 0x95, 0xc2,
	0xbe, 0x6b, 0x40, 0xf7, 0x0f, 0x47, 0x64, 0x3e, 0x27, 0x47, 0x30, 0x44, 0xc1, 0x12, 0x0c, 0x62,
	0x89, 0xa7, 0x29, 0x3f, 0xcf, 0x38, 0x4a, 0x53, 0x1f, 0xeb, 0x5f, 0xdf, 0x7d, 0xfb, 0xec, 0x1c,
	0xfa, 0xdb, 0xf9, 0xb7, 0x95, 0xd3, 0x42, 0x3d, 0xd5, 0xe8, 0x00, 0xf7, 0x38, 0x72, 0x0c, 0xa4,
	0x6e, 0x8b, 0x49, 0x2c, 0x90, 0x9b, 0x0d, 0xe5, 0xfb, 0xe5, 0x45, 0xdf, 0x42, 0x3e, 0xd5, 0xe8,
	0x10, 0xf7, 0x49, 0xf2, 0x0b, 0x0c, 0x2f, 0xc8, 0xc4, 0x59, 0x15, 0xb6, 0xa9, 0x4c, 0xed, 0xc3,
	0xa6, 0x3f, 0x73, 0xe9, 0x2e, 0x68, 0xdf, 0xab, 0x61, 0xf2, 0x1b, 0xde, 0x6f, 0xad, 0xca, 0x80,
	0x2d, 0xe5, 0xf5, 0xe9, 0x59, 0xaf, 0x2a, 0x9c, 0xe1, 0xd5, 0x89, 0x49, 0x1b, 0x9a, 0x98, 0x45,
	0x36, 0x81, 0xc1, 0xfe, 0x42, 0xf6, 0x95, 0x0e, 0xc3, 0x27, 0xf5, 0xc8, 0x07, 0xe8, 0x04, 0x3c,
	0xf4, 0x83, 0x62, 0xef, 0x16, 0x2d, 0x51, 0xce, 0x2f, 0xe2, 0x34, 0x62, 0x52, 0xed, 0x65, 0xd0,
	0x12, 0xe5, 0xbc, 0xba, 0x11, 0x55, 0x65, 0x83, 0x96, 0x88, 0x10, 0x68, 0x05, 0x0c, 0x03, 0x15,
	0xbe, 0x4f, 0xd5, 0x99, 0x7c, 0x84, 0x5e, 0xc4, 0x25, 0x9b, 0x33, 0xc9, 0xcc, 0xb6, 0xe2, 0x2b,
	0x6c, 0xff, 0x87, 0x7e, 0x7d, 0x96, 0x57, 0xe7, 0x18, 0x41, 0x3b, 0x14, 0x73, 0x7e, 0x51, 0xc6,
	0x28, 0x80, 0x7d, 0xa9, 0x83, 0xf1, 0x68, 0xa1, 0xb7, 0xf1, 0xcd, 0x59, 0xd5, 0xb3, 0xac, 0x57,
	0x00, 0x62, 0x42, 0x37, 0x0a, 0x11, 0x43, 0xe1, 0xab, 0x7a, 0x3d, 0xba, 0x85, 0x93, 0xbf, 0x37,
	0x6b, 0x4b, 0x5f, 0xad, 0x2d, 0xfd, 0x7e, 0x6d, 0xe9, 0xd7, 0x1b, 0x4b, 0x5b, 0x6d, 0x2c, 0xed,
	0x76, 0x63, 0x69, 0x27, 0x3f, 0xfc, 0x50, 0x06, 0xd9, 0xcc, 0xf1, 0xe2, 0xc8, 0xf5, 0xe2, 0x88,
	0xcb, 0xd9, 0x42, 0xee, 0x0e, 0xea, 0xc1, 0xb8, 0x87, 0x5e, 0xd4, 0xac, 0xa3, 0xbe, 0x7d, 0x7f,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x04, 0xbe, 0xb0, 0x90, 0x70, 0x03, 0x00, 0x00,
}

func (m *Signal) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signal) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
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

func (m *Signal_Mirrorsrequest) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Mirrorsrequest) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MirrorsQuery != nil {
		{
			volume, err := m.MirrorsQuery.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Mirrorsreply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Mirrorsreply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MirrorsReply != nil {
		{
			volume, err := m.MirrorsReply.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Segmentrequest) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Segmentrequest) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SegmentQuery != nil {
		{
			volume, err := m.SegmentQuery.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Signal_Segmentreply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Signal_Segmentreply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SegmentReply != nil {
		{
			volume, err := m.SegmentReply.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = formatVariableintKinds(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *MirrorsQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MirrorsQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *MirrorsQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MirrorsReply) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MirrorsReply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *MirrorsReply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metainfo) > 0 {
		i -= len(m.Metainfo)
		copy(dAtA[i:], m.Metainfo)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Metainfo)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x22
	}
	if m.Segments != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Segments))
		i--
		dAtA[i] = 0x18
	}
	if m.Layout != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Layout))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SegmentQuery) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SegmentQuery) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *SegmentQuery) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ordinal != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ordinal))
		i--
		dAtA[i] = 0x18
	}
	if m.Layout != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Layout))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SegmentReply) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SegmentReply) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *SegmentReply) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Absent {
		i--
		if m.Absent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Segment) > 0 {
		i -= len(m.Segment)
		copy(dAtA[i:], m.Segment)
		i = formatVariableintKinds(dAtA, i, uint64(len(m.Segment)))
		i--
		dAtA[i] = 0x22
	}
	if m.Ordinal != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Ordinal))
		i--
		dAtA[i] = 0x18
	}
	if m.Layout != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Layout))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = formatVariableintKinds(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func formatVariableintKinds(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovKinds(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *Signal) Volume() (n int) {
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

func (m *Signal_Mirrorsrequest) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MirrorsQuery != nil {
		l = m.MirrorsQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Mirrorsreply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MirrorsReply != nil {
		l = m.MirrorsReply.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Segmentrequest) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SegmentQuery != nil {
		l = m.SegmentQuery.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Signal_Segmentreply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SegmentReply != nil {
		l = m.SegmentReply.Volume()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *MirrorsQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MirrorsReply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Layout != 0 {
		n += 1 + sovKinds(uint64(m.Layout))
	}
	if m.Segments != 0 {
		n += 1 + sovKinds(uint64(m.Segments))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	l = len(m.Metainfo)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SegmentQuery) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Layout != 0 {
		n += 1 + sovKinds(uint64(m.Layout))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	return n
}

func (m *SegmentReply) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovKinds(uint64(m.Level))
	}
	if m.Layout != 0 {
		n += 1 + sovKinds(uint64(m.Layout))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	l = len(m.Segment)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	if m.Absent {
		n += 2
	}
	return n
}

func sovKinds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKinds(x uint64) (n int) {
	return sovKinds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Signal) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &MirrorsQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Mirrorsrequest{v}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &MirrorsReply{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Mirrorsreply{v}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &SegmentQuery{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Segmentrequest{v}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &SegmentReply{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Signal_Segmentreply{v}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *MirrorsQuery) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *MirrorsReply) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Layout |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Segments = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Segments |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = append(m.Digest[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Metainfo = append(m.Metainfo[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Metainfo == nil {
				m.Metainfo = []byte{}
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *SegmentQuery) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Layout |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Ordinal |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func (m *SegmentReply) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadKinds
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
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Layout |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Ordinal |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var octetSize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
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
				return ErrCorruptExtentKinds
			}
			submitOrdinal := idxNdEx + octetSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentKinds
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.Segment = append(m.Segment[:0], dAtA[idxNdEx:submitOrdinal]...)
			if m.Segment == nil {
				m.Segment = []byte{}
			}
			idxNdEx = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var v int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadKinds
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				v |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			m.Absent = bool(v != 0)
		default:
			idxNdEx = preOrdinal
			skippy, err := omitKinds(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentKinds
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
func omitKinds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadKinds
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
					return 0, ErrIntegerOverloadKinds
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
					return 0, ErrIntegerOverloadKinds
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
				return 0, ErrCorruptExtentKinds
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterKinds
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentKinds
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentKinds        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadKinds          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterKinds = fmt.Errorf("REDACTED")
)
