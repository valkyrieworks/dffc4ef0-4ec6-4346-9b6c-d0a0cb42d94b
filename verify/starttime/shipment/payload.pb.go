//
//
//
//
//

package shipment

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
type Shipment struct {
	status         protoimpl.MessageState
	volumeRepository     protoimpl.SizeCache
	unclearAttributes protoimpl.UnknownFields

	Linkages uint64                 `protobuf:"variableint,1,opt,name=connections,proto3" json:"linkages,omitempty"`
	Ratio        uint64                 `protobuf:"variableint,2,opt,name=rate,proto3" json:"ratio,omitempty"`
	Volume        uint64                 `protobuf:"variableint,3,opt,name=size,proto3" json:"volume,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"octets,4,opt,name=time,proto3" json:"moment,omitempty"`
	Id          []byte                 `protobuf:"octets,5,opt,name=id,proto3" json:"id,omitempty"`
	Stuffing     []byte                 `protobuf:"octets,6,opt,name=padding,proto3" json:"stuffing,omitempty"`
}

func (x *Shipment) Restore() {
	*x = Shipment{}
	if protoimpl.UnsafeEnabled {
		mi := &entry_shipment_shipment_schema_messagetypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipment) SchemaSignal() {}

func (x *Shipment) SchemaMirror() protoreflect.Message {
	mi := &entry_shipment_shipment_schema_messagetypes[0]
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
func (*Shipment) Definition() ([]byte, []int) {
	return entry_shipment_shipment_schema_rawdescgzip(), []int{0}
}

func (x *Shipment) FetchLinkages() uint64 {
	if x != nil {
		return x.Linkages
	}
	return 0
}

func (x *Shipment) FetchRatio() uint64 {
	if x != nil {
		return x.Ratio
	}
	return 0
}

func (x *Shipment) FetchVolume() uint64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Shipment) FetchTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Shipment) FetchUid() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Shipment) FetchStuffing() []byte {
	if x != nil {
		return x.Stuffing
	}
	return nil
}

var Entry_shipment_shipment_schema protoreflect.FileDescriptor

var entry_shipment_shipment_schema_rawdesc = []byte{
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
	entry_shipment_shipment_schema_rawdesconce sync.Once
	entry_shipment_shipment_schema_rawdescdata = entry_shipment_shipment_schema_rawdesc
)

func entry_shipment_shipment_schema_rawdescgzip() []byte {
	entry_shipment_shipment_schema_rawdesconce.Do(func() {
		entry_shipment_shipment_schema_rawdescdata = protoimpl.X.CompressGZIP(entry_shipment_shipment_schema_rawdescdata)
	})
	return entry_shipment_shipment_schema_rawdescdata
}

var entry_shipment_shipment_schema_messagetypes = make([]protoimpl.MessageInfo, 1)
var entry_shipment_shipment_schema_gotypes = []any{
	(*Shipment)(nil),               //
	(*timestamppb.Timestamp)(nil), //
}
var entry_shipment_shipment_schema_depindexes = []int32{
	1, //
	1, //
	1, //
	1, //
	1, //
	0, //
}

func init() { entry_shipment_shipment_schema_init() }
func entry_shipment_shipment_schema_init() {
	if Entry_shipment_shipment_schema != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		entry_shipment_shipment_schema_messagetypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Shipment); i {
			case 0:
				return &v.status
			case 1:
				return &v.volumeRepository
			case 2:
				return &v.unclearAttributes
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: entry_shipment_shipment_schema_rawdesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           entry_shipment_shipment_schema_gotypes,
		DependencyIndexes: entry_shipment_shipment_schema_depindexes,
		MessageInfos:      entry_shipment_shipment_schema_messagetypes,
	}.Build()
	Entry_shipment_shipment_schema = out.File
	entry_shipment_shipment_schema_rawdesc = nil
	entry_shipment_shipment_schema_gotypes = nil
	entry_shipment_shipment_schema_depindexes = nil
}
