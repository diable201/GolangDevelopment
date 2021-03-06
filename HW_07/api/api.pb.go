// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

type Anime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TitleJapanese string  `protobuf:"bytes,3,opt,name=titleJapanese,proto3" json:"titleJapanese,omitempty"`
	Source        string  `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Episodes      int32   `protobuf:"varint,5,opt,name=episodes,proto3" json:"episodes,omitempty"`
	Kind          string  `protobuf:"bytes,6,opt,name=kind,proto3" json:"kind,omitempty"`
	Score         float64 `protobuf:"fixed64,7,opt,name=score,proto3" json:"score,omitempty"`
	Status        string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Synopsis      string  `protobuf:"bytes,9,opt,name=synopsis,proto3" json:"synopsis,omitempty"`
}

func (x *Anime) Reset() {
	*x = Anime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Anime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Anime) ProtoMessage() {}

func (x *Anime) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Anime.ProtoReflect.Descriptor instead.
func (*Anime) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *Anime) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Anime) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Anime) GetTitleJapanese() string {
	if x != nil {
		return x.TitleJapanese
	}
	return ""
}

func (x *Anime) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Anime) GetEpisodes() int32 {
	if x != nil {
		return x.Episodes
	}
	return 0
}

func (x *Anime) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Anime) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Anime) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Anime) GetSynopsis() string {
	if x != nil {
		return x.Synopsis
	}
	return ""
}

type AnimeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anime []*Anime `protobuf:"bytes,1,rep,name=anime,proto3" json:"anime,omitempty"`
}

func (x *AnimeList) Reset() {
	*x = AnimeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeList) ProtoMessage() {}

func (x *AnimeList) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeList.ProtoReflect.Descriptor instead.
func (*AnimeList) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *AnimeList) GetAnime() []*Anime {
	if x != nil {
		return x.Anime
	}
	return nil
}

type AnimeRequestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AnimeRequestId) Reset() {
	*x = AnimeRequestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeRequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeRequestId) ProtoMessage() {}

func (x *AnimeRequestId) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeRequestId.ProtoReflect.Descriptor instead.
func (*AnimeRequestId) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *AnimeRequestId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xe5, 0x01,
	0x0a, 0x05, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4a, 0x61, 0x70, 0x61, 0x6e, 0x65, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4a, 0x61, 0x70, 0x61, 0x6e,
	0x65, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e,
	0x6f, 0x70, 0x73, 0x69, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x6e,
	0x6f, 0x70, 0x73, 0x69, 0x73, 0x22, 0x2d, 0x0a, 0x09, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x05, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc6, 0x01, 0x0a, 0x0c, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x0a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x0b, 0x5a, 0x09, 0x48, 0x57, 0x5f, 0x30, 0x37, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_api_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: api.Empty
	(*Anime)(nil),          // 1: api.Anime
	(*AnimeList)(nil),      // 2: api.AnimeList
	(*AnimeRequestId)(nil), // 3: api.AnimeRequestId
}
var file_api_api_proto_depIdxs = []int32{
	1, // 0: api.AnimeList.anime:type_name -> api.Anime
	0, // 1: api.AnimeService.GetAll:input_type -> api.Empty
	3, // 2: api.AnimeService.Get:input_type -> api.AnimeRequestId
	1, // 3: api.AnimeService.Post:input_type -> api.Anime
	1, // 4: api.AnimeService.Put:input_type -> api.Anime
	3, // 5: api.AnimeService.Delete:input_type -> api.AnimeRequestId
	2, // 6: api.AnimeService.GetAll:output_type -> api.AnimeList
	1, // 7: api.AnimeService.Get:output_type -> api.Anime
	1, // 8: api.AnimeService.Post:output_type -> api.Anime
	1, // 9: api.AnimeService.Put:output_type -> api.Anime
	0, // 10: api.AnimeService.Delete:output_type -> api.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Anime); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeRequestId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
