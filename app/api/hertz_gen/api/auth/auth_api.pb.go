// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: auth_api.proto

package auth

import (
	_ "Gomall/app/api/hertz_gen/api"
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

type DeliverTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" form:"user_id" json:"user_id,omitempty"`
}

func (x *DeliverTokenReq) Reset() {
	*x = DeliverTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverTokenReq) ProtoMessage() {}

func (x *DeliverTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverTokenReq.ProtoReflect.Descriptor instead.
func (*DeliverTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{0}
}

func (x *DeliverTokenReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type VerifyTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" form:"token" json:"token,omitempty"`
}

func (x *VerifyTokenReq) Reset() {
	*x = VerifyTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenReq) ProtoMessage() {}

func (x *VerifyTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenReq.ProtoReflect.Descriptor instead.
func (*VerifyTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeliveryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`
}

func (x *DeliveryResp) Reset() {
	*x = DeliveryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryResp) ProtoMessage() {}

func (x *DeliveryResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryResp.ProtoReflect.Descriptor instead.
func (*DeliveryResp) Descriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty" form:"res" query:"res"`
}

func (x *VerifyResp) Reset() {
	*x = VerifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResp) ProtoMessage() {}

func (x *VerifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResp.ProtoReflect.Descriptor instead.
func (*VerifyResp) Descriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyResp) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

var File_auth_api_proto protoreflect.FileDescriptor

var file_auth_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31,
	0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x1f, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xca, 0xbb, 0x18, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x24, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1e, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0xb4, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x12, 0x4f, 0x0a,
	0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10, 0xd2, 0xc1,
	0x18, 0x0c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x23,
	0x5a, 0x21, 0x47, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_api_proto_rawDescOnce sync.Once
	file_auth_api_proto_rawDescData = file_auth_api_proto_rawDesc
)

func file_auth_api_proto_rawDescGZIP() []byte {
	file_auth_api_proto_rawDescOnce.Do(func() {
		file_auth_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_api_proto_rawDescData)
	})
	return file_auth_api_proto_rawDescData
}

var file_auth_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_api_proto_goTypes = []interface{}{
	(*DeliverTokenReq)(nil), // 0: api.auth.DeliverTokenReq
	(*VerifyTokenReq)(nil),  // 1: api.auth.VerifyTokenReq
	(*DeliveryResp)(nil),    // 2: api.auth.DeliveryResp
	(*VerifyResp)(nil),      // 3: api.auth.VerifyResp
}
var file_auth_api_proto_depIdxs = []int32{
	0, // 0: api.auth.AuthService.DeliverToken:input_type -> api.auth.DeliverTokenReq
	1, // 1: api.auth.AuthService.VerifyToken:input_type -> api.auth.VerifyTokenReq
	2, // 2: api.auth.AuthService.DeliverToken:output_type -> api.auth.DeliveryResp
	3, // 3: api.auth.AuthService.VerifyToken:output_type -> api.auth.VerifyResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_api_proto_init() }
func file_auth_api_proto_init() {
	if File_auth_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverTokenReq); i {
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
		file_auth_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenReq); i {
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
		file_auth_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryResp); i {
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
		file_auth_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResp); i {
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
			RawDescriptor: file_auth_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_api_proto_goTypes,
		DependencyIndexes: file_auth_api_proto_depIdxs,
		MessageInfos:      file_auth_api_proto_msgTypes,
	}.Build()
	File_auth_api_proto = out.File
	file_auth_api_proto_rawDesc = nil
	file_auth_api_proto_goTypes = nil
	file_auth_api_proto_depIdxs = nil
}
