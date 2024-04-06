// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/mesh/v1alpha1/zoneegress.proto

package v1alpha1

import (
	_ "github.com/apache/dubbo-kubernetes/api/mesh"
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

type ZoneEgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Zone field contains Zone name where egress is serving, field will be
	// automatically set by Global Dubbo CP
	Zone string `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty"`
	// Networking defines the address and port of the Egress to listen on.
	Networking *ZoneEgress_Networking `protobuf:"bytes,2,opt,name=networking,proto3" json:"networking,omitempty"`
}

func (x *ZoneEgress) Reset() {
	*x = ZoneEgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_zoneegress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneEgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneEgress) ProtoMessage() {}

func (x *ZoneEgress) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_zoneegress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneEgress.ProtoReflect.Descriptor instead.
func (*ZoneEgress) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_zoneegress_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneEgress) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *ZoneEgress) GetNetworking() *ZoneEgress_Networking {
	if x != nil {
		return x.Networking
	}
	return nil
}

type ZoneEgress_Networking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address on which inbound listener will be exposed
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Port of the inbound interface that will forward requests to the service.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Admin contains configuration related to Envoy Admin API
	Admin *EnvoyAdmin `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *ZoneEgress_Networking) Reset() {
	*x = ZoneEgress_Networking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_zoneegress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneEgress_Networking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneEgress_Networking) ProtoMessage() {}

func (x *ZoneEgress_Networking) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_zoneegress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneEgress_Networking.ProtoReflect.Descriptor instead.
func (*ZoneEgress_Networking) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_zoneegress_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ZoneEgress_Networking) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ZoneEgress_Networking) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ZoneEgress_Networking) GetAdmin() *EnvoyAdmin {
	if x != nil {
		return x.Admin
	}
	return nil
}

var File_api_mesh_v1alpha1_zoneegress_proto protoreflect.FileDescriptor

var file_api_mesh_v1alpha1_zoneegress_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x0a, 0x5a, 0x6f, 0x6e, 0x65, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x71, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x35, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x7c, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x14,
	0x0a, 0x12, 0x5a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x0c, 0x12, 0x0a, 0x5a, 0x6f, 0x6e, 0x65,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x06, 0x22, 0x04, 0x6d, 0x65,
	0x73, 0x68, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x04, 0x52, 0x02, 0x10, 0x01, 0xaa, 0x8c, 0x89, 0xa6,
	0x01, 0x0e, 0x3a, 0x0c, 0x0a, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73,
	0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x10, 0x3a, 0x0e, 0x12, 0x0c, 0x7a, 0x6f, 0x6e, 0x65, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x68, 0x01, 0xaa, 0x8c,
	0x89, 0xa6, 0x01, 0x02, 0x58, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_mesh_v1alpha1_zoneegress_proto_rawDescOnce sync.Once
	file_api_mesh_v1alpha1_zoneegress_proto_rawDescData = file_api_mesh_v1alpha1_zoneegress_proto_rawDesc
)

func file_api_mesh_v1alpha1_zoneegress_proto_rawDescGZIP() []byte {
	file_api_mesh_v1alpha1_zoneegress_proto_rawDescOnce.Do(func() {
		file_api_mesh_v1alpha1_zoneegress_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mesh_v1alpha1_zoneegress_proto_rawDescData)
	})
	return file_api_mesh_v1alpha1_zoneegress_proto_rawDescData
}

var file_api_mesh_v1alpha1_zoneegress_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_mesh_v1alpha1_zoneegress_proto_goTypes = []interface{}{
	(*ZoneEgress)(nil),            // 0: dubbo.mesh.v1alpha1.ZoneEgress
	(*ZoneEgress_Networking)(nil), // 1: dubbo.mesh.v1alpha1.ZoneEgress.Networking
	(*EnvoyAdmin)(nil),            // 2: dubbo.mesh.v1alpha1.EnvoyAdmin
}
var file_api_mesh_v1alpha1_zoneegress_proto_depIdxs = []int32{
	1, // 0: dubbo.mesh.v1alpha1.ZoneEgress.networking:type_name -> dubbo.mesh.v1alpha1.ZoneEgress.Networking
	2, // 1: dubbo.mesh.v1alpha1.ZoneEgress.Networking.admin:type_name -> dubbo.mesh.v1alpha1.EnvoyAdmin
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_mesh_v1alpha1_zoneegress_proto_init() }
func file_api_mesh_v1alpha1_zoneegress_proto_init() {
	if File_api_mesh_v1alpha1_zoneegress_proto != nil {
		return
	}
	file_api_mesh_v1alpha1_envoy_admin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_mesh_v1alpha1_zoneegress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneEgress); i {
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
		file_api_mesh_v1alpha1_zoneegress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneEgress_Networking); i {
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
			RawDescriptor: file_api_mesh_v1alpha1_zoneegress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_mesh_v1alpha1_zoneegress_proto_goTypes,
		DependencyIndexes: file_api_mesh_v1alpha1_zoneegress_proto_depIdxs,
		MessageInfos:      file_api_mesh_v1alpha1_zoneegress_proto_msgTypes,
	}.Build()
	File_api_mesh_v1alpha1_zoneegress_proto = out.File
	file_api_mesh_v1alpha1_zoneegress_proto_rawDesc = nil
	file_api_mesh_v1alpha1_zoneegress_proto_goTypes = nil
	file_api_mesh_v1alpha1_zoneegress_proto_depIdxs = nil
}
