// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: yoyoinfo.proto

package protoc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type YoyoInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	YoyoCode           int64  `protobuf:"varint,2,opt,name=YoyoCode,proto3" json:"YoyoCode,omitempty"`
	YoyoVersionCode    string `protobuf:"bytes,3,opt,name=YoyoVersionCode,proto3" json:"YoyoVersionCode,omitempty"`
	ImageUrls          string `protobuf:"bytes,4,opt,name=ImageUrls,proto3" json:"ImageUrls,omitempty"`
	NickName           string `protobuf:"bytes,5,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Introduce          string `protobuf:"bytes,6,opt,name=Introduce,proto3" json:"Introduce,omitempty"`
	Composition        string `protobuf:"bytes,7,opt,name=Composition,proto3" json:"Composition,omitempty"`
	Effect             string `protobuf:"bytes,8,opt,name=Effect,proto3" json:"Effect,omitempty"`
	Collocation        string `protobuf:"bytes,9,opt,name=Collocation,proto3" json:"Collocation,omitempty"`
	MatchColors        string `protobuf:"bytes,10,opt,name=MatchColors,proto3" json:"MatchColors,omitempty"`
	FirstCategoryCode  int32  `protobuf:"varint,11,opt,name=FirstCategoryCode,proto3" json:"FirstCategoryCode,omitempty"`
	FirstCategoryName  string `protobuf:"bytes,12,opt,name=FirstCategoryName,proto3" json:"FirstCategoryName,omitempty"`
	SecondCategoryCode int32  `protobuf:"varint,13,opt,name=SecondCategoryCode,proto3" json:"SecondCategoryCode,omitempty"`
	SecondCategoryName string `protobuf:"bytes,14,opt,name=SecondCategoryName,proto3" json:"SecondCategoryName,omitempty"`
	ThirdCategoryCode  int32  `protobuf:"varint,15,opt,name=ThirdCategoryCode,proto3" json:"ThirdCategoryCode,omitempty"`
	ThirdCategoryName  string `protobuf:"bytes,16,opt,name=ThirdCategoryName,proto3" json:"ThirdCategoryName,omitempty"`
	ForthCategoryCode  int32  `protobuf:"varint,17,opt,name=ForthCategoryCode,proto3" json:"ForthCategoryCode,omitempty"`
	ForthCategoryName  string `protobuf:"bytes,18,opt,name=ForthCategoryName,proto3" json:"ForthCategoryName,omitempty"`
}

func (x *YoyoInfo) Reset() {
	*x = YoyoInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yoyoinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YoyoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YoyoInfo) ProtoMessage() {}

func (x *YoyoInfo) ProtoReflect() protoreflect.Message {
	mi := &file_yoyoinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YoyoInfo.ProtoReflect.Descriptor instead.
func (*YoyoInfo) Descriptor() ([]byte, []int) {
	return file_yoyoinfo_proto_rawDescGZIP(), []int{0}
}

func (x *YoyoInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *YoyoInfo) GetYoyoCode() int64 {
	if x != nil {
		return x.YoyoCode
	}
	return 0
}

func (x *YoyoInfo) GetYoyoVersionCode() string {
	if x != nil {
		return x.YoyoVersionCode
	}
	return ""
}

func (x *YoyoInfo) GetImageUrls() string {
	if x != nil {
		return x.ImageUrls
	}
	return ""
}

func (x *YoyoInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *YoyoInfo) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

func (x *YoyoInfo) GetComposition() string {
	if x != nil {
		return x.Composition
	}
	return ""
}

func (x *YoyoInfo) GetEffect() string {
	if x != nil {
		return x.Effect
	}
	return ""
}

func (x *YoyoInfo) GetCollocation() string {
	if x != nil {
		return x.Collocation
	}
	return ""
}

func (x *YoyoInfo) GetMatchColors() string {
	if x != nil {
		return x.MatchColors
	}
	return ""
}

func (x *YoyoInfo) GetFirstCategoryCode() int32 {
	if x != nil {
		return x.FirstCategoryCode
	}
	return 0
}

func (x *YoyoInfo) GetFirstCategoryName() string {
	if x != nil {
		return x.FirstCategoryName
	}
	return ""
}

func (x *YoyoInfo) GetSecondCategoryCode() int32 {
	if x != nil {
		return x.SecondCategoryCode
	}
	return 0
}

func (x *YoyoInfo) GetSecondCategoryName() string {
	if x != nil {
		return x.SecondCategoryName
	}
	return ""
}

func (x *YoyoInfo) GetThirdCategoryCode() int32 {
	if x != nil {
		return x.ThirdCategoryCode
	}
	return 0
}

func (x *YoyoInfo) GetThirdCategoryName() string {
	if x != nil {
		return x.ThirdCategoryName
	}
	return ""
}

func (x *YoyoInfo) GetForthCategoryCode() int32 {
	if x != nil {
		return x.ForthCategoryCode
	}
	return 0
}

func (x *YoyoInfo) GetForthCategoryName() string {
	if x != nil {
		return x.ForthCategoryName
	}
	return ""
}

type YoyoInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YoyoInfo []*YoyoInfo `protobuf:"bytes,1,rep,name=YoyoInfo,proto3" json:"YoyoInfo,omitempty"`
}

func (x *YoyoInfoList) Reset() {
	*x = YoyoInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yoyoinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YoyoInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YoyoInfoList) ProtoMessage() {}

func (x *YoyoInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_yoyoinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YoyoInfoList.ProtoReflect.Descriptor instead.
func (*YoyoInfoList) Descriptor() ([]byte, []int) {
	return file_yoyoinfo_proto_rawDescGZIP(), []int{1}
}

func (x *YoyoInfoList) GetYoyoInfo() []*YoyoInfo {
	if x != nil {
		return x.YoyoInfo
	}
	return nil
}

var File_yoyoinfo_proto protoreflect.FileDescriptor

var file_yoyoinfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x79, 0x6f, 0x79, 0x6f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x22, 0xae, 0x05, 0x0a, 0x08, 0x59, 0x6f, 0x79,
	0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x59, 0x6f, 0x79,
	0x6f, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x59, 0x6f, 0x79,
	0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x59, 0x6f, 0x79, 0x6f, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x59, 0x6f, 0x79, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x68, 0x69, 0x72, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x54, 0x68,
	0x69, 0x72, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x54, 0x68, 0x69, 0x72, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x68, 0x69, 0x72,
	0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x46, 0x6f, 0x72, 0x74, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x46, 0x6f, 0x72, 0x74, 0x68, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x46,
	0x6f, 0x72, 0x74, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x46, 0x6f, 0x72, 0x74, 0x68, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x59, 0x6f, 0x79,
	0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x59, 0x6f, 0x79,
	0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x59, 0x6f, 0x79, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x59,
	0x6f, 0x79, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yoyoinfo_proto_rawDescOnce sync.Once
	file_yoyoinfo_proto_rawDescData = file_yoyoinfo_proto_rawDesc
)

func file_yoyoinfo_proto_rawDescGZIP() []byte {
	file_yoyoinfo_proto_rawDescOnce.Do(func() {
		file_yoyoinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_yoyoinfo_proto_rawDescData)
	})
	return file_yoyoinfo_proto_rawDescData
}

var file_yoyoinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yoyoinfo_proto_goTypes = []interface{}{
	(*YoyoInfo)(nil),     // 0: protoc.YoyoInfo
	(*YoyoInfoList)(nil), // 1: protoc.YoyoInfoList
}
var file_yoyoinfo_proto_depIdxs = []int32{
	0, // 0: protoc.YoyoInfoList.YoyoInfo:type_name -> protoc.YoyoInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yoyoinfo_proto_init() }
func file_yoyoinfo_proto_init() {
	if File_yoyoinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yoyoinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YoyoInfo); i {
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
		file_yoyoinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YoyoInfoList); i {
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
			RawDescriptor: file_yoyoinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yoyoinfo_proto_goTypes,
		DependencyIndexes: file_yoyoinfo_proto_depIdxs,
		MessageInfos:      file_yoyoinfo_proto_msgTypes,
	}.Build()
	File_yoyoinfo_proto = out.File
	file_yoyoinfo_proto_rawDesc = nil
	file_yoyoinfo_proto_goTypes = nil
	file_yoyoinfo_proto_depIdxs = nil
}
