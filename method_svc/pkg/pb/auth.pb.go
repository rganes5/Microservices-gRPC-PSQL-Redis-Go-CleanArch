// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pkg/proto/auth.proto

package pb

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

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag int32 `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllRequest) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int32         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Response string        `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	Error    string        `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Count    int32         `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Users    []*MethodUser `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetAllResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *GetAllResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetAllResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllResponse) GetUsers() []*MethodUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type MethodUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
}

func (x *MethodUser) Reset() {
	*x = MethodUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodUser) ProtoMessage() {}

func (x *MethodUser) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodUser.ProtoReflect.Descriptor instead.
func (*MethodUser) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *MethodUser) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

var File_pkg_proto_auth_proto protoreflect.FileDescriptor

var file_pkg_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22,
	0x96, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x2b, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x40, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_auth_proto_rawDescOnce sync.Once
	file_pkg_proto_auth_proto_rawDescData = file_pkg_proto_auth_proto_rawDesc
)

func file_pkg_proto_auth_proto_rawDescGZIP() []byte {
	file_pkg_proto_auth_proto_rawDescOnce.Do(func() {
		file_pkg_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_auth_proto_rawDescData)
	})
	return file_pkg_proto_auth_proto_rawDescData
}

var file_pkg_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_auth_proto_goTypes = []interface{}{
	(*GetAllRequest)(nil),  // 0: pb.GetAllRequest
	(*GetAllResponse)(nil), // 1: pb.GetAllResponse
	(*MethodUser)(nil),     // 2: pb.MethodUser
}
var file_pkg_proto_auth_proto_depIdxs = []int32{
	2, // 0: pb.GetAllResponse.users:type_name -> pb.MethodUser
	0, // 1: pb.AuthService.GetAll:input_type -> pb.GetAllRequest
	1, // 2: pb.AuthService.GetAll:output_type -> pb.GetAllResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_auth_proto_init() }
func file_pkg_proto_auth_proto_init() {
	if File_pkg_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_pkg_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_pkg_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodUser); i {
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
			RawDescriptor: file_pkg_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_auth_proto_goTypes,
		DependencyIndexes: file_pkg_proto_auth_proto_depIdxs,
		MessageInfos:      file_pkg_proto_auth_proto_msgTypes,
	}.Build()
	File_pkg_proto_auth_proto = out.File
	file_pkg_proto_auth_proto_rawDesc = nil
	file_pkg_proto_auth_proto_goTypes = nil
	file_pkg_proto_auth_proto_depIdxs = nil
}
