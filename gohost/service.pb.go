// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: gohost/service.proto

package service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type AliveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Alive    bool   `protobuf:"varint,2,opt,name=alive,proto3" json:"alive,omitempty"`
}

func (x *AliveRequest) Reset() {
	*x = AliveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gohost_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AliveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AliveRequest) ProtoMessage() {}

func (x *AliveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gohost_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AliveRequest.ProtoReflect.Descriptor instead.
func (*AliveRequest) Descriptor() ([]byte, []int) {
	return file_gohost_service_proto_rawDescGZIP(), []int{0}
}

func (x *AliveRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *AliveRequest) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

var File_gohost_service_proto protoreflect.FileDescriptor

var file_gohost_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6f, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0c, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x32, 0x41, 0x0a,
	0x06, 0x47, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x6c, 0x69,
	0x76, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gohost_service_proto_rawDescOnce sync.Once
	file_gohost_service_proto_rawDescData = file_gohost_service_proto_rawDesc
)

func file_gohost_service_proto_rawDescGZIP() []byte {
	file_gohost_service_proto_rawDescOnce.Do(func() {
		file_gohost_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gohost_service_proto_rawDescData)
	})
	return file_gohost_service_proto_rawDescData
}

var file_gohost_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gohost_service_proto_goTypes = []interface{}{
	(*AliveRequest)(nil), // 0: gohost.AliveRequest
	(*empty.Empty)(nil),  // 1: google.protobuf.Empty
}
var file_gohost_service_proto_depIdxs = []int32{
	0, // 0: gohost.GoHost.IsAlive:input_type -> gohost.AliveRequest
	1, // 1: gohost.GoHost.IsAlive:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gohost_service_proto_init() }
func file_gohost_service_proto_init() {
	if File_gohost_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gohost_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AliveRequest); i {
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
			RawDescriptor: file_gohost_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gohost_service_proto_goTypes,
		DependencyIndexes: file_gohost_service_proto_depIdxs,
		MessageInfos:      file_gohost_service_proto_msgTypes,
	}.Build()
	File_gohost_service_proto = out.File
	file_gohost_service_proto_rawDesc = nil
	file_gohost_service_proto_goTypes = nil
	file_gohost_service_proto_depIdxs = nil
}
