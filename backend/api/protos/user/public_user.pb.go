// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: user/public_user.proto

package user

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Passwd   string `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_public_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_public_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_user_public_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

type SelectRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RoleId uint64 `protobuf:"varint,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *SelectRoleReq) Reset() {
	*x = SelectRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_public_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectRoleReq) ProtoMessage() {}

func (x *SelectRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_public_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectRoleReq.ProtoReflect.Descriptor instead.
func (*SelectRoleReq) Descriptor() ([]byte, []int) {
	return file_user_public_user_proto_rawDescGZIP(), []int{1}
}

func (x *SelectRoleReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SelectRoleReq) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type GetListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*User `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetListResp) Reset() {
	*x = GetListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_public_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResp) ProtoMessage() {}

func (x *GetListResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_public_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResp.ProtoReflect.Descriptor instead.
func (*GetListResp) Descriptor() ([]byte, []int) {
	return file_user_public_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetListResp) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_user_public_user_proto protoreflect.FileDescriptor

var file_user_public_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3e, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x22,
	0x4a, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x00, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x96, 0x02, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x12, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x3b, 0x5a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_public_user_proto_rawDescOnce sync.Once
	file_user_public_user_proto_rawDescData = file_user_public_user_proto_rawDesc
)

func file_user_public_user_proto_rawDescGZIP() []byte {
	file_user_public_user_proto_rawDescOnce.Do(func() {
		file_user_public_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_public_user_proto_rawDescData)
	})
	return file_user_public_user_proto_rawDescData
}

var file_user_public_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_public_user_proto_goTypes = []interface{}{
	(*LoginReq)(nil),      // 0: dashboard.user.LoginReq
	(*SelectRoleReq)(nil), // 1: dashboard.user.SelectRoleReq
	(*GetListResp)(nil),   // 2: dashboard.user.GetListResp
	(*User)(nil),          // 3: dashboard.user.user
	(*empty.Empty)(nil),   // 4: google.protobuf.Empty
}
var file_user_public_user_proto_depIdxs = []int32{
	3, // 0: dashboard.user.GetListResp.data:type_name -> dashboard.user.user
	0, // 1: dashboard.user.PublicUser.Login:input_type -> dashboard.user.LoginReq
	1, // 2: dashboard.user.PublicUser.SelectRole:input_type -> dashboard.user.SelectRoleReq
	4, // 3: dashboard.user.PublicUser.GetList:input_type -> google.protobuf.Empty
	4, // 4: dashboard.user.PublicUser.Login:output_type -> google.protobuf.Empty
	4, // 5: dashboard.user.PublicUser.SelectRole:output_type -> google.protobuf.Empty
	2, // 6: dashboard.user.PublicUser.GetList:output_type -> dashboard.user.GetListResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_public_user_proto_init() }
func file_user_public_user_proto_init() {
	if File_user_public_user_proto != nil {
		return
	}
	file_user_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_public_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_user_public_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectRoleReq); i {
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
		file_user_public_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResp); i {
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
			RawDescriptor: file_user_public_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_public_user_proto_goTypes,
		DependencyIndexes: file_user_public_user_proto_depIdxs,
		MessageInfos:      file_user_public_user_proto_msgTypes,
	}.Build()
	File_user_public_user_proto = out.File
	file_user_public_user_proto_rawDesc = nil
	file_user_public_user_proto_goTypes = nil
	file_user_public_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PublicUserClient is the client API for PublicUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublicUserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*empty.Empty, error)
	SelectRole(ctx context.Context, in *SelectRoleReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetListResp, error)
}

type publicUserClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicUserClient(cc grpc.ClientConnInterface) PublicUserClient {
	return &publicUserClient{cc}
}

func (c *publicUserClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dashboard.user.PublicUser/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicUserClient) SelectRole(ctx context.Context, in *SelectRoleReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dashboard.user.PublicUser/SelectRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicUserClient) GetList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetListResp, error) {
	out := new(GetListResp)
	err := c.cc.Invoke(ctx, "/dashboard.user.PublicUser/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicUserServer is the server API for PublicUser service.
type PublicUserServer interface {
	Login(context.Context, *LoginReq) (*empty.Empty, error)
	SelectRole(context.Context, *SelectRoleReq) (*empty.Empty, error)
	GetList(context.Context, *empty.Empty) (*GetListResp, error)
}

// UnimplementedPublicUserServer can be embedded to have forward compatible implementations.
type UnimplementedPublicUserServer struct {
}

func (*UnimplementedPublicUserServer) Login(context.Context, *LoginReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedPublicUserServer) SelectRole(context.Context, *SelectRoleReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectRole not implemented")
}
func (*UnimplementedPublicUserServer) GetList(context.Context, *empty.Empty) (*GetListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}

func RegisterPublicUserServer(s *grpc.Server, srv PublicUserServer) {
	s.RegisterService(&_PublicUser_serviceDesc, srv)
}

func _PublicUser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicUserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.user.PublicUser/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicUserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicUser_SelectRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicUserServer).SelectRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.user.PublicUser/SelectRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicUserServer).SelectRole(ctx, req.(*SelectRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicUser_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicUserServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.user.PublicUser/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicUserServer).GetList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PublicUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.user.PublicUser",
	HandlerType: (*PublicUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _PublicUser_Login_Handler,
		},
		{
			MethodName: "SelectRole",
			Handler:    _PublicUser_SelectRole_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _PublicUser_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/public_user.proto",
}
