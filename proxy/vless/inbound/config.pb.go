package inbound

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	protocol "v2ray.com/core/common/protocol"
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

type DefaultConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level uint32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *DefaultConfig) Reset() {
	*x = DefaultConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultConfig) ProtoMessage() {}

func (x *DefaultConfig) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultConfig.ProtoReflect.Descriptor instead.
func (*DefaultConfig) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescGZIP(), []int{0}
}

func (x *DefaultConfig) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    []*protocol.User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Default *DefaultConfig   `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetUser() []*protocol.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Config) GetDefault() *DefaultConfig {
	if x != nil {
		return x.Default
	}
	return nil
}

var File_v2ray_com_core_proxy_vless_inbound_config_proto protoreflect.FileDescriptor

var file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x1a, 0x29, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x87, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x69,
	0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x50, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x50, 0x01, 0x5a, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0xaa, 0x02,
	0x1e, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x56, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescOnce sync.Once
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData = file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc
)

func file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescGZIP() []byte {
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescOnce.Do(func() {
		file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData)
	})
	return file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData
}

var file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v2ray_com_core_proxy_vless_inbound_config_proto_goTypes = []interface{}{
	(*DefaultConfig)(nil), // 0: v2ray.core.proxy.vless.inbound.DefaultConfig
	(*Config)(nil),        // 1: v2ray.core.proxy.vless.inbound.Config
	(*protocol.User)(nil), // 2: v2ray.core.common.protocol.User
}
var file_v2ray_com_core_proxy_vless_inbound_config_proto_depIdxs = []int32{
	2, // 0: v2ray.core.proxy.vless.inbound.Config.user:type_name -> v2ray.core.common.protocol.User
	0, // 1: v2ray.core.proxy.vless.inbound.Config.default:type_name -> v2ray.core.proxy.vless.inbound.DefaultConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v2ray_com_core_proxy_vless_inbound_config_proto_init() }
func file_v2ray_com_core_proxy_vless_inbound_config_proto_init() {
	if File_v2ray_com_core_proxy_vless_inbound_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultConfig); i {
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
		file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2ray_com_core_proxy_vless_inbound_config_proto_goTypes,
		DependencyIndexes: file_v2ray_com_core_proxy_vless_inbound_config_proto_depIdxs,
		MessageInfos:      file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes,
	}.Build()
	File_v2ray_com_core_proxy_vless_inbound_config_proto = out.File
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc = nil
	file_v2ray_com_core_proxy_vless_inbound_config_proto_goTypes = nil
	file_v2ray_com_core_proxy_vless_inbound_config_proto_depIdxs = nil
}
