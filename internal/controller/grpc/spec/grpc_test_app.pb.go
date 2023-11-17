// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: grpc_test_app.proto

package spec

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

type GetItnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Itn uint64 `protobuf:"varint,1,opt,name=itn,proto3" json:"itn,omitempty"`
}

func (x *GetItnRequest) Reset() {
	*x = GetItnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_test_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItnRequest) ProtoMessage() {}

func (x *GetItnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_test_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItnRequest.ProtoReflect.Descriptor instead.
func (*GetItnRequest) Descriptor() ([]byte, []int) {
	return file_grpc_test_app_proto_rawDescGZIP(), []int{0}
}

func (x *GetItnRequest) GetItn() uint64 {
	if x != nil {
		return x.Itn
	}
	return 0
}

type GetItnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Itn         uint64 `protobuf:"varint,1,opt,name=itn,proto3" json:"itn,omitempty"`
	Kpp         uint64 `protobuf:"varint,2,opt,name=kpp,proto3" json:"kpp,omitempty"`
	NameCompany string `protobuf:"bytes,3,opt,name=nameCompany,proto3" json:"nameCompany,omitempty"`
	NameLeader  string `protobuf:"bytes,4,opt,name=nameLeader,proto3" json:"nameLeader,omitempty"`
	Success     bool   `protobuf:"varint,5,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetItnResponse) Reset() {
	*x = GetItnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_test_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItnResponse) ProtoMessage() {}

func (x *GetItnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_test_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItnResponse.ProtoReflect.Descriptor instead.
func (*GetItnResponse) Descriptor() ([]byte, []int) {
	return file_grpc_test_app_proto_rawDescGZIP(), []int{1}
}

func (x *GetItnResponse) GetItn() uint64 {
	if x != nil {
		return x.Itn
	}
	return 0
}

func (x *GetItnResponse) GetKpp() uint64 {
	if x != nil {
		return x.Kpp
	}
	return 0
}

func (x *GetItnResponse) GetNameCompany() string {
	if x != nil {
		return x.NameCompany
	}
	return ""
}

func (x *GetItnResponse) GetNameLeader() string {
	if x != nil {
		return x.NameLeader
	}
	return ""
}

func (x *GetItnResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_grpc_test_app_proto protoreflect.FileDescriptor

var file_grpc_test_app_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72, 0x70, 0x63, 0x54, 0x65, 0x73, 0x74, 0x22,
	0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x74, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x74, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x69,
	0x74, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x74, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x69, 0x74, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x70, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x61, 0x6d, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_test_app_proto_rawDescOnce sync.Once
	file_grpc_test_app_proto_rawDescData = file_grpc_test_app_proto_rawDesc
)

func file_grpc_test_app_proto_rawDescGZIP() []byte {
	file_grpc_test_app_proto_rawDescOnce.Do(func() {
		file_grpc_test_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_test_app_proto_rawDescData)
	})
	return file_grpc_test_app_proto_rawDescData
}

var file_grpc_test_app_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_test_app_proto_goTypes = []interface{}{
	(*GetItnRequest)(nil),  // 0: grpcTest.GetItnRequest
	(*GetItnResponse)(nil), // 1: grpcTest.GetItnResponse
}
var file_grpc_test_app_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_test_app_proto_init() }
func file_grpc_test_app_proto_init() {
	if File_grpc_test_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_test_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItnRequest); i {
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
		file_grpc_test_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItnResponse); i {
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
			RawDescriptor: file_grpc_test_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_test_app_proto_goTypes,
		DependencyIndexes: file_grpc_test_app_proto_depIdxs,
		MessageInfos:      file_grpc_test_app_proto_msgTypes,
	}.Build()
	File_grpc_test_app_proto = out.File
	file_grpc_test_app_proto_rawDesc = nil
	file_grpc_test_app_proto_goTypes = nil
	file_grpc_test_app_proto_depIdxs = nil
}
