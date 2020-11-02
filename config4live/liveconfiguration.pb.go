// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: config4live/liveconfiguration.proto

package config4live

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

type ConfigResponse_Format int32

const (
	ConfigResponse_text   ConfigResponse_Format = 0
	ConfigResponse_number ConfigResponse_Format = 1
	ConfigResponse_bool   ConfigResponse_Format = 2
	ConfigResponse_json   ConfigResponse_Format = 3
)

// Enum value maps for ConfigResponse_Format.
var (
	ConfigResponse_Format_name = map[int32]string{
		0: "text",
		1: "number",
		2: "bool",
		3: "json",
	}
	ConfigResponse_Format_value = map[string]int32{
		"text":   0,
		"number": 1,
		"bool":   2,
		"json":   3,
	}
)

func (x ConfigResponse_Format) Enum() *ConfigResponse_Format {
	p := new(ConfigResponse_Format)
	*p = x
	return p
}

func (x ConfigResponse_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigResponse_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_config4live_liveconfiguration_proto_enumTypes[0].Descriptor()
}

func (ConfigResponse_Format) Type() protoreflect.EnumType {
	return &file_config4live_liveconfiguration_proto_enumTypes[0]
}

func (x ConfigResponse_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigResponse_Format.Descriptor instead.
func (ConfigResponse_Format) EnumDescriptor() ([]byte, []int) {
	return file_config4live_liveconfiguration_proto_rawDescGZIP(), []int{1, 0}
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config4live_liveconfiguration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config4live_liveconfiguration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config4live_liveconfiguration_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value       string                `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Description string                `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Format      ConfigResponse_Format `protobuf:"varint,5,opt,name=format,proto3,enum=config4live.ConfigResponse_Format" json:"format,omitempty"`
	Owner       string                `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config4live_liveconfiguration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config4live_liveconfiguration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_config4live_liveconfiguration_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConfigResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ConfigResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ConfigResponse) GetFormat() ConfigResponse_Format {
	if x != nil {
		return x.Format
	}
	return ConfigResponse_text
}

func (x *ConfigResponse) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

var File_config4live_liveconfiguration_proto protoreflect.FileDescriptor

var file_config4live_liveconfiguration_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x34, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x6c, 0x69,
	0x76, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x34, 0x6c, 0x69,
	0x76, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x34,
	0x6c, 0x69, 0x76, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x10, 0x03, 0x32, 0x5c, 0x0a, 0x11,
	0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x47, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x34, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x34, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x0a, 0x18, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x61, 0x79, 0x75, 0x72, 0x62, 0x6f, 0x78, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x34, 0x6c, 0x69, 0x76, 0x65, 0x42, 0x16, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0xa2, 0x02, 0x03, 0x48, 0x4c, 0x57, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config4live_liveconfiguration_proto_rawDescOnce sync.Once
	file_config4live_liveconfiguration_proto_rawDescData = file_config4live_liveconfiguration_proto_rawDesc
)

func file_config4live_liveconfiguration_proto_rawDescGZIP() []byte {
	file_config4live_liveconfiguration_proto_rawDescOnce.Do(func() {
		file_config4live_liveconfiguration_proto_rawDescData = protoimpl.X.CompressGZIP(file_config4live_liveconfiguration_proto_rawDescData)
	})
	return file_config4live_liveconfiguration_proto_rawDescData
}

var file_config4live_liveconfiguration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config4live_liveconfiguration_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config4live_liveconfiguration_proto_goTypes = []interface{}{
	(ConfigResponse_Format)(0), // 0: config4live.ConfigResponse.Format
	(*ConfigRequest)(nil),      // 1: config4live.ConfigRequest
	(*ConfigResponse)(nil),     // 2: config4live.ConfigResponse
}
var file_config4live_liveconfiguration_proto_depIdxs = []int32{
	0, // 0: config4live.ConfigResponse.format:type_name -> config4live.ConfigResponse.Format
	1, // 1: config4live.LiveConfiguration.FindConfig:input_type -> config4live.ConfigRequest
	2, // 2: config4live.LiveConfiguration.FindConfig:output_type -> config4live.ConfigResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config4live_liveconfiguration_proto_init() }
func file_config4live_liveconfiguration_proto_init() {
	if File_config4live_liveconfiguration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config4live_liveconfiguration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_config4live_liveconfiguration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
			RawDescriptor: file_config4live_liveconfiguration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config4live_liveconfiguration_proto_goTypes,
		DependencyIndexes: file_config4live_liveconfiguration_proto_depIdxs,
		EnumInfos:         file_config4live_liveconfiguration_proto_enumTypes,
		MessageInfos:      file_config4live_liveconfiguration_proto_msgTypes,
	}.Build()
	File_config4live_liveconfiguration_proto = out.File
	file_config4live_liveconfiguration_proto_rawDesc = nil
	file_config4live_liveconfiguration_proto_goTypes = nil
	file_config4live_liveconfiguration_proto_depIdxs = nil
}
