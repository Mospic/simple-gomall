// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1-devel
// 	protoc        v5.29.2
// source: service/proto/token_utils.proto

package service

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

type GetIdByTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserToken     string                 `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"` // token
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIdByTokenRequest) Reset() {
	*x = GetIdByTokenRequest{}
	mi := &file_service_proto_token_utils_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIdByTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdByTokenRequest) ProtoMessage() {}

func (x *GetIdByTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_token_utils_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdByTokenRequest.ProtoReflect.Descriptor instead.
func (*GetIdByTokenRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_token_utils_proto_rawDescGZIP(), []int{0}
}

func (x *GetIdByTokenRequest) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

type GetIdByTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIdByTokenResponse) Reset() {
	*x = GetIdByTokenResponse{}
	mi := &file_service_proto_token_utils_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIdByTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdByTokenResponse) ProtoMessage() {}

func (x *GetIdByTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_token_utils_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdByTokenResponse.ProtoReflect.Descriptor instead.
func (*GetIdByTokenResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_token_utils_proto_rawDescGZIP(), []int{1}
}

func (x *GetIdByTokenResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GenerateTokenByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenByIDRequest) Reset() {
	*x = GenerateTokenByIDRequest{}
	mi := &file_service_proto_token_utils_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenByIDRequest) ProtoMessage() {}

func (x *GenerateTokenByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_token_utils_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenByIDRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenByIDRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_token_utils_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateTokenByIDRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GenerateTokenByIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserToken     string                 `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"` // token
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenByIDResponse) Reset() {
	*x = GenerateTokenByIDResponse{}
	mi := &file_service_proto_token_utils_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenByIDResponse) ProtoMessage() {}

func (x *GenerateTokenByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_token_utils_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenByIDResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenByIDResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_token_utils_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateTokenByIDResponse) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

var File_service_proto_token_utils_proto protoreflect.FileDescriptor

var file_service_proto_token_utils_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x49, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x1a, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xbf, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49,
	0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x23,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x44,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_proto_token_utils_proto_rawDescOnce sync.Once
	file_service_proto_token_utils_proto_rawDescData = file_service_proto_token_utils_proto_rawDesc
)

func file_service_proto_token_utils_proto_rawDescGZIP() []byte {
	file_service_proto_token_utils_proto_rawDescOnce.Do(func() {
		file_service_proto_token_utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_token_utils_proto_rawDescData)
	})
	return file_service_proto_token_utils_proto_rawDescData
}

var file_service_proto_token_utils_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_proto_token_utils_proto_goTypes = []any{
	(*GetIdByTokenRequest)(nil),       // 0: services.GetIdByToken_request
	(*GetIdByTokenResponse)(nil),      // 1: services.GetIdByToken_response
	(*GenerateTokenByIDRequest)(nil),  // 2: services.GenerateTokenByID_request
	(*GenerateTokenByIDResponse)(nil), // 3: services.GenerateTokenByID_response
}
var file_service_proto_token_utils_proto_depIdxs = []int32{
	0, // 0: services.TokenService.GetIdByToken:input_type -> services.GetIdByToken_request
	2, // 1: services.TokenService.GenerateTokenByID:input_type -> services.GenerateTokenByID_request
	1, // 2: services.TokenService.GetIdByToken:output_type -> services.GetIdByToken_response
	3, // 3: services.TokenService.GenerateTokenByID:output_type -> services.GenerateTokenByID_response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_token_utils_proto_init() }
func file_service_proto_token_utils_proto_init() {
	if File_service_proto_token_utils_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_token_utils_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_token_utils_proto_goTypes,
		DependencyIndexes: file_service_proto_token_utils_proto_depIdxs,
		MessageInfos:      file_service_proto_token_utils_proto_msgTypes,
	}.Build()
	File_service_proto_token_utils_proto = out.File
	file_service_proto_token_utils_proto_rawDesc = nil
	file_service_proto_token_utils_proto_goTypes = nil
	file_service_proto_token_utils_proto_depIdxs = nil
}
