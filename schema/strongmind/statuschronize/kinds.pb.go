//
//

package statuschronize

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
	Sum isnote_Total `protobuf_oneof:"sum"`
}

func (m *Signal) Restore()         { *m = Signal{} }
func (m *Signal) Text() string { return proto.CompactTextString(m) }
func (*Signal) SchemaArtifact()    {}
func (*Signal) Definition() ([]byte, []int) {
	return filedescriptor_a1c2869546ca7914, []int{0}
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

type Artifact_Imagessolicit struct {
	ImagesSolicit *ImagesSolicit `protobuf:"octets,1,opt,name=snapshots_request,json=snapshotsRequest,proto3,oneof" json:"images_solicit,omitempty"`
}
type Artifact_Imagesreply struct {
	ImagesReply *ImagesReply `protobuf:"octets,2,opt,name=snapshots_response,json=snapshotsResponse,proto3,oneof" json:"images_reply,omitempty"`
}
type Artifact_Fragmentsolicit struct {
	SegmentSolicit *SegmentSolicit `protobuf:"octets,3,opt,name=chunk_request,json=chunkRequest,proto3,oneof" json:"segment_solicit,omitempty"`
}
type Artifact_Fragmentreply struct {
	SegmentReply *SegmentReply `protobuf:"octets,4,opt,name=chunk_response,json=chunkResponse,proto3,oneof" json:"segment_reply,omitempty"`
}

func (*Artifact_Imagessolicit) isnote_Total()  {}
func (*Artifact_Imagesreply) isnote_Total() {}
func (*Artifact_Fragmentsolicit) isnote_Total()      {}
func (*Artifact_Fragmentreply) isnote_Total()     {}

func (m *Signal) ObtainTotal() isnote_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Signal) ObtainImagesSolicit() *ImagesSolicit {
	if x, ok := m.ObtainTotal().(*Artifact_Imagessolicit); ok {
		return x.ImagesSolicit
	}
	return nil
}

func (m *Signal) ObtainImagesReply() *ImagesReply {
	if x, ok := m.ObtainTotal().(*Artifact_Imagesreply); ok {
		return x.ImagesReply
	}
	return nil
}

func (m *Signal) ObtainSegmentSolicit() *SegmentSolicit {
	if x, ok := m.ObtainTotal().(*Artifact_Fragmentsolicit); ok {
		return x.SegmentSolicit
	}
	return nil
}

func (m *Signal) ObtainSegmentReply() *SegmentReply {
	if x, ok := m.ObtainTotal().(*Artifact_Fragmentreply); ok {
		return x.SegmentReply
	}
	return nil
}

//
func (*Signal) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Artifact_Imagessolicit)(nil),
		(*Artifact_Imagesreply)(nil),
		(*Artifact_Fragmentsolicit)(nil),
		(*Artifact_Fragmentreply)(nil),
	}
}

type ImagesSolicit struct {
}

func (m *ImagesSolicit) Restore()         { *m = ImagesSolicit{} }
func (m *ImagesSolicit) Text() string { return proto.CompactTextString(m) }
func (*ImagesSolicit) SchemaArtifact()    {}
func (*ImagesSolicit) Definition() ([]byte, []int) {
	return filedescriptor_a1c2869546ca7914, []int{1}
}
func (m *ImagesSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ImagesSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Imagessolicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImagesSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Imagessolicit.Merge(m, src)
}
func (m *ImagesSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *ImagesSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Imagessolicit.DiscardUnknown(m)
}

var xxx_signaldetails_Imagessolicit proto.InternalMessageInfo

type ImagesReply struct {
	Altitude   uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Layout   uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Segments   uint32 `protobuf:"variableint,3,opt,name=chunks,proto3" json:"segments,omitempty"`
	Digest     []byte `protobuf:"octets,4,opt,name=hash,proto3" json:"digest,omitempty"`
	Attributes []byte `protobuf:"octets,5,opt,name=metadata,proto3" json:"attributes,omitempty"`
}

func (m *ImagesReply) Restore()         { *m = ImagesReply{} }
func (m *ImagesReply) Text() string { return proto.CompactTextString(m) }
func (*ImagesReply) SchemaArtifact()    {}
func (*ImagesReply) Definition() ([]byte, []int) {
	return filedescriptor_a1c2869546ca7914, []int{2}
}
func (m *ImagesReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ImagesReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Imagesreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImagesReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Imagesreply.Merge(m, src)
}
func (m *ImagesReply) XXX_Extent() int {
	return m.Extent()
}
func (m *ImagesReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Imagesreply.DiscardUnknown(m)
}

var xxx_signaldetails_Imagesreply proto.InternalMessageInfo

func (m *ImagesReply) ObtainAltitude() uint64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *ImagesReply) ObtainLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *ImagesReply) ObtainSegments() uint32 {
	if m != nil {
		return m.Segments
	}
	return 0
}

func (m *ImagesReply) ObtainDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *ImagesReply) ObtainAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type SegmentSolicit struct {
	Altitude uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Layout uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Ordinal  uint32 `protobuf:"variableint,3,opt,name=index,proto3" json:"ordinal,omitempty"`
}

func (m *SegmentSolicit) Restore()         { *m = SegmentSolicit{} }
func (m *SegmentSolicit) Text() string { return proto.CompactTextString(m) }
func (*SegmentSolicit) SchemaArtifact()    {}
func (*SegmentSolicit) Definition() ([]byte, []int) {
	return filedescriptor_a1c2869546ca7914, []int{3}
}
func (m *SegmentSolicit) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SegmentSolicit) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Fragmentsolicit.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SegmentSolicit) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Fragmentsolicit.Merge(m, src)
}
func (m *SegmentSolicit) XXX_Extent() int {
	return m.Extent()
}
func (m *SegmentSolicit) XXX_Dropunfamiliar() {
	xxx_signaldetails_Fragmentsolicit.DiscardUnknown(m)
}

var xxx_signaldetails_Fragmentsolicit proto.InternalMessageInfo

func (m *SegmentSolicit) ObtainAltitude() uint64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SegmentSolicit) ObtainLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *SegmentSolicit) ObtainOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type SegmentReply struct {
	Altitude  uint64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
	Layout  uint32 `protobuf:"variableint,2,opt,name=format,proto3" json:"layout,omitempty"`
	Ordinal   uint32 `protobuf:"variableint,3,opt,name=index,proto3" json:"ordinal,omitempty"`
	Segment   []byte `protobuf:"octets,4,opt,name=chunk,proto3" json:"segment,omitempty"`
	Absent bool   `protobuf:"variableint,5,opt,name=missing,proto3" json:"absent,omitempty"`
}

func (m *SegmentReply) Restore()         { *m = SegmentReply{} }
func (m *SegmentReply) Text() string { return proto.CompactTextString(m) }
func (*SegmentReply) SchemaArtifact()    {}
func (*SegmentReply) Definition() ([]byte, []int) {
	return filedescriptor_a1c2869546ca7914, []int{4}
}
func (m *SegmentReply) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SegmentReply) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Fragmentreply.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SegmentReply) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Fragmentreply.Merge(m, src)
}
func (m *SegmentReply) XXX_Extent() int {
	return m.Extent()
}
func (m *SegmentReply) XXX_Dropunfamiliar() {
	xxx_signaldetails_Fragmentreply.DiscardUnknown(m)
}

var xxx_signaldetails_Fragmentreply proto.InternalMessageInfo

func (m *SegmentReply) ObtainAltitude() uint64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *SegmentReply) ObtainLayout() uint32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *SegmentReply) ObtainOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *SegmentReply) ObtainSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *SegmentReply) ObtainAbsent() bool {
	if m != nil {
		return m.Absent
	}
	return false
}

func initialize() {
	proto.RegisterType((*Signal)(nil), "REDACTED")
	proto.RegisterType((*ImagesSolicit)(nil), "REDACTED")
	proto.RegisterType((*ImagesReply)(nil), "REDACTED")
	proto.RegisterType((*SegmentSolicit)(nil), "REDACTED")
	proto.RegisterType((*SegmentReply)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_a1c2869546ca7914) }

var filedescriptor_a1c2869546ca7914 = []byte{
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

func (m *Artifact_Imagessolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Imagessolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ImagesSolicit != nil {
		{
			extent, err := m.ImagesSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
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
func (m *Artifact_Imagesreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Imagesreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.ImagesReply != nil {
		{
			extent, err := m.ImagesReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Artifact_Fragmentsolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Fragmentsolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.SegmentSolicit != nil {
		{
			extent, err := m.SegmentSolicit.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Artifact_Fragmentreply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Artifact_Fragmentreply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.SegmentReply != nil {
		{
			extent, err := m.SegmentReply.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = serializeVariableintKinds(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x22
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *ImagesSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ImagesSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ImagesSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	return len(deltaLocatedAN) - i, nil
}

func (m *ImagesReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ImagesReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ImagesReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		i -= len(m.Attributes)
		copy(deltaLocatedAN[i:], m.Attributes)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Attributes)))
		i--
		deltaLocatedAN[i] = 0x2a
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(deltaLocatedAN[i:], m.Digest)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Digest)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Segments != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Segments))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Layout != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Layout))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SegmentSolicit) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SegmentSolicit) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SegmentSolicit) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Ordinal != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ordinal))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Layout != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Layout))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *SegmentReply) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SegmentReply) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SegmentReply) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Absent {
		i--
		if m.Absent {
			deltaLocatedAN[i] = 1
		} else {
			deltaLocatedAN[i] = 0
		}
		i--
		deltaLocatedAN[i] = 0x28
	}
	if len(m.Segment) > 0 {
		i -= len(m.Segment)
		copy(deltaLocatedAN[i:], m.Segment)
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(len(m.Segment)))
		i--
		deltaLocatedAN[i] = 0x22
	}
	if m.Ordinal != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Ordinal))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Layout != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Layout))
		i--
		deltaLocatedAN[i] = 0x10
	}
	if m.Altitude != 0 {
		i = serializeVariableintKinds(deltaLocatedAN, i, uint64(m.Altitude))
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

func (m *Artifact_Imagessolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ImagesSolicit != nil {
		l = m.ImagesSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Imagesreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ImagesReply != nil {
		l = m.ImagesReply.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Fragmentsolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SegmentSolicit != nil {
		l = m.SegmentSolicit.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *Artifact_Fragmentreply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SegmentReply != nil {
		l = m.SegmentReply.Extent()
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}
func (m *ImagesSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ImagesReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
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
	l = len(m.Attributes)
	if l > 0 {
		n += 1 + l + sovKinds(uint64(l))
	}
	return n
}

func (m *SegmentSolicit) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
	}
	if m.Layout != 0 {
		n += 1 + sovKinds(uint64(m.Layout))
	}
	if m.Ordinal != 0 {
		n += 1 + sovKinds(uint64(m.Ordinal))
	}
	return n
}

func (m *SegmentReply) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovKinds(uint64(m.Altitude))
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
			v := &ImagesSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Imagessolicit{v}
			idxNdExc = submitOrdinal
		case 2:
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
			v := &ImagesReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Imagesreply{v}
			idxNdExc = submitOrdinal
		case 3:
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
			v := &SegmentSolicit{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Fragmentsolicit{v}
			idxNdExc = submitOrdinal
		case 4:
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
			v := &SegmentReply{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Artifact_Fragmentreply{v}
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
func (m *ImagesSolicit) Decode(deltaLocatedAN []byte) error {
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
func (m *ImagesReply) Decode(deltaLocatedAN []byte) error {
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Layout |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Segments = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Segments |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
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
			m.Digest = append(m.Digest[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			idxNdExc = submitOrdinal
		case 5:
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
			m.Attributes = append(m.Attributes[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Attributes == nil {
				m.Attributes = []byte{}
			}
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
func (m *SegmentSolicit) Decode(deltaLocatedAN []byte) error {
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Layout |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Ordinal |= uint32(b&0x7F) << relocate
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
func (m *SegmentReply) Decode(deltaLocatedAN []byte) error {
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
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Layout = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Layout |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Ordinal = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Ordinal |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
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
			m.Segment = append(m.Segment[:0], deltaLocatedAN[idxNdExc:submitOrdinal]...)
			if m.Segment == nil {
				m.Segment = []byte{}
			}
			idxNdExc = submitOrdinal
		case 5:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var v int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunKinds
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				v |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			m.Absent = bool(v != 0)
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
