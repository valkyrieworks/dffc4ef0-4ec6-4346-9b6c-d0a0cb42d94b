//
//
//
//
//

package content

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	//
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	//
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
//
//
type Content struct {
	status         protoimpl.MessageState
	extentStash     protoimpl.SizeCache
	unfamiliarAreas protoimpl.UnknownFields

	Linkages uint64                 `protobuf:"variableint,1,opt,name=connections,proto3" json:"linkages,omitempty"`
	Frequency        uint64                 `protobuf:"variableint,2,opt,name=rate,proto3" json:"frequency,omitempty"`
	Extent        uint64                 `protobuf:"variableint,3,opt,name=size,proto3" json:"extent,omitempty"`
	Moment        *timestamppb.Timestamp `protobuf:"octets,4,opt,name=time,proto3" json:"moment,omitempty"`
	Id          []byte                 `protobuf:"octets,5,opt,name=id,proto3" json:"id,omitempty"`
	Filling     []byte                 `protobuf:"octets,6,opt,name=padding,proto3" json:"filling,omitempty"`
}

func (x *Content) Restore() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &record_content_content_schema_messagetypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) Text() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) SchemaArtifact() {}

func (x *Content) SchemaMirror() protoreflect.Message {
	mi := &record_content_content_schema_messagetypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

//
func (*Content) Definition() ([]byte, []int) {
	return record_content_content_schema_uncookedschemagzip(), []int{0}
}

func (x *Content) ObtainLinkages() uint64 {
	if x != nil {
		return x.Linkages
	}
	return 0
}

func (x *Content) ObtainFrequency() uint64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *Content) ObtainExtent() uint64 {
	if x != nil {
		return x.Extent
	}
	return 0
}

func (x *Content) ObtainMoment() *timestamppb.Timestamp {
	if x != nil {
		return x.Moment
	}
	return nil
}

func (x *Content) ObtainUuid() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Content) ObtainFilling() []byte {
	if x != nil {
		return x.Filling
	}
	return nil
}

var Record_content_content_schema protoreflect.FileDescriptor

var record_content_content_schema_uncookedschema = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	record_content_content_schema_uncookedschemaconce sync.Once
	record_content_content_schema_uncookedschemadata = record_content_content_schema_uncookedschema
)

func record_content_content_schema_uncookedschemagzip() []byte {
	record_content_content_schema_uncookedschemaconce.Do(func() {
		record_content_content_schema_uncookedschemadata = protoimpl.X.CompressGZIP(record_content_content_schema_uncookedschemadata)
	})
	return record_content_content_schema_uncookedschemadata
}

var record_content_content_schema_messagetypes = make([]protoimpl.MessageInfo, 1)
var record_content_content_schema_goformats = []any{
	(*Content)(nil),               //
	(*timestamppb.Timestamp)(nil), //
}
var record_content_content_schema_depindices = []int32{
	1, //
	1, //
	1, //
	1, //
	1, //
	0, //
}

func initialize() { record_content_content_schema_initialize() }
func record_content_content_schema_initialize() {
	if Record_content_content_schema != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		record_content_content_schema_messagetypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Content); i {
			case 0:
				return &v.status
			case 1:
				return &v.extentStash
			case 2:
				return &v.unfamiliarAreas
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: record_content_content_schema_uncookedschema,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           record_content_content_schema_goformats,
		DependencyIndexes: record_content_content_schema_depindices,
		MessageInfos:      record_content_content_schema_messagetypes,
	}.Build()
	Record_content_content_schema = out.File
	record_content_content_schema_uncookedschema = nil
	record_content_content_schema_goformats = nil
	record_content_content_schema_depindices = nil
}
