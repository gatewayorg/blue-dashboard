// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: rbac/public_rule.proto

package rbac

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

type GetRuleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetRuleReq) Reset() {
	*x = GetRuleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_public_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleReq) ProtoMessage() {}

func (x *GetRuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_public_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleReq.ProtoReflect.Descriptor instead.
func (*GetRuleReq) Descriptor() ([]byte, []int) {
	return file_rbac_public_rule_proto_rawDescGZIP(), []int{0}
}

func (x *GetRuleReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetRuleReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetRuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*Rule `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetRuleResp) Reset() {
	*x = GetRuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_public_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleResp) ProtoMessage() {}

func (x *GetRuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_public_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleResp.ProtoReflect.Descriptor instead.
func (*GetRuleResp) Descriptor() ([]byte, []int) {
	return file_rbac_public_rule_proto_rawDescGZIP(), []int{1}
}

func (x *GetRuleResp) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetRuleResp) GetData() []*Rule {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *SetDetailReq) Reset() {
	*x = SetDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_public_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDetailReq) ProtoMessage() {}

func (x *SetDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_public_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDetailReq.ProtoReflect.Descriptor instead.
func (*SetDetailReq) Descriptor() ([]byte, []int) {
	return file_rbac_public_rule_proto_rawDescGZIP(), []int{2}
}

func (x *SetDetailReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetDetailReq) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

var File_rbac_public_rule_proto protoreflect.FileDescriptor

var file_rbac_public_rule_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x1a, 0x11, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x4d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x72, 0x62, 0x61, 0x63, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x4b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10,
	0x01, 0x18, 0xff, 0x01, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x32, 0xc9, 0x01, 0x0a,
	0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x5b, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x72,
	0x62, 0x61, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x32, 0x10, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x73, 0x65, 0x74, 0x2f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x6f, 0x72,
	0x67, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x62, 0x61, 0x63,
	0x3b, 0x72, 0x62, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rbac_public_rule_proto_rawDescOnce sync.Once
	file_rbac_public_rule_proto_rawDescData = file_rbac_public_rule_proto_rawDesc
)

func file_rbac_public_rule_proto_rawDescGZIP() []byte {
	file_rbac_public_rule_proto_rawDescOnce.Do(func() {
		file_rbac_public_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_rbac_public_rule_proto_rawDescData)
	})
	return file_rbac_public_rule_proto_rawDescData
}

var file_rbac_public_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rbac_public_rule_proto_goTypes = []interface{}{
	(*GetRuleReq)(nil),   // 0: dashboard.rbac.GetRuleReq
	(*GetRuleResp)(nil),  // 1: dashboard.rbac.GetRuleResp
	(*SetDetailReq)(nil), // 2: dashboard.rbac.SetDetailReq
	(*Rule)(nil),         // 3: dashboard.rbac.Rule
	(*empty.Empty)(nil),  // 4: google.protobuf.Empty
}
var file_rbac_public_rule_proto_depIdxs = []int32{
	3, // 0: dashboard.rbac.GetRuleResp.data:type_name -> dashboard.rbac.Rule
	0, // 1: dashboard.rbac.PublicRule.GetRule:input_type -> dashboard.rbac.GetRuleReq
	2, // 2: dashboard.rbac.PublicRule.SetDetail:input_type -> dashboard.rbac.SetDetailReq
	1, // 3: dashboard.rbac.PublicRule.GetRule:output_type -> dashboard.rbac.GetRuleResp
	4, // 4: dashboard.rbac.PublicRule.SetDetail:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rbac_public_rule_proto_init() }
func file_rbac_public_rule_proto_init() {
	if File_rbac_public_rule_proto != nil {
		return
	}
	file_rbac_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rbac_public_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleReq); i {
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
		file_rbac_public_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleResp); i {
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
		file_rbac_public_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDetailReq); i {
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
			RawDescriptor: file_rbac_public_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rbac_public_rule_proto_goTypes,
		DependencyIndexes: file_rbac_public_rule_proto_depIdxs,
		MessageInfos:      file_rbac_public_rule_proto_msgTypes,
	}.Build()
	File_rbac_public_rule_proto = out.File
	file_rbac_public_rule_proto_rawDesc = nil
	file_rbac_public_rule_proto_goTypes = nil
	file_rbac_public_rule_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PublicRuleClient is the client API for PublicRule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublicRuleClient interface {
	GetRule(ctx context.Context, in *GetRuleReq, opts ...grpc.CallOption) (*GetRuleResp, error)
	SetDetail(ctx context.Context, in *SetDetailReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type publicRuleClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicRuleClient(cc grpc.ClientConnInterface) PublicRuleClient {
	return &publicRuleClient{cc}
}

func (c *publicRuleClient) GetRule(ctx context.Context, in *GetRuleReq, opts ...grpc.CallOption) (*GetRuleResp, error) {
	out := new(GetRuleResp)
	err := c.cc.Invoke(ctx, "/dashboard.rbac.PublicRule/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicRuleClient) SetDetail(ctx context.Context, in *SetDetailReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dashboard.rbac.PublicRule/SetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicRuleServer is the server API for PublicRule service.
type PublicRuleServer interface {
	GetRule(context.Context, *GetRuleReq) (*GetRuleResp, error)
	SetDetail(context.Context, *SetDetailReq) (*empty.Empty, error)
}

// UnimplementedPublicRuleServer can be embedded to have forward compatible implementations.
type UnimplementedPublicRuleServer struct {
}

func (*UnimplementedPublicRuleServer) GetRule(context.Context, *GetRuleReq) (*GetRuleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (*UnimplementedPublicRuleServer) SetDetail(context.Context, *SetDetailReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDetail not implemented")
}

func RegisterPublicRuleServer(s *grpc.Server, srv PublicRuleServer) {
	s.RegisterService(&_PublicRule_serviceDesc, srv)
}

func _PublicRule_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicRuleServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.rbac.PublicRule/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicRuleServer).GetRule(ctx, req.(*GetRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicRule_SetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicRuleServer).SetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.rbac.PublicRule/SetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicRuleServer).SetDetail(ctx, req.(*SetDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PublicRule_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.rbac.PublicRule",
	HandlerType: (*PublicRuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRule",
			Handler:    _PublicRule_GetRule_Handler,
		},
		{
			MethodName: "SetDetail",
			Handler:    _PublicRule_SetDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rbac/public_rule.proto",
}
