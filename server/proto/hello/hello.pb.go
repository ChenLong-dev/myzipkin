// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: hello.proto

package hello

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HelloReq) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type HelloResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *HelloResp) Reset() {
	*x = HelloResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResp) ProtoMessage() {}

func (x *HelloResp) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResp.ProtoReflect.Descriptor instead.
func (*HelloResp) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HelloResp) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x22, 0x2f, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x76, 0x61, 0x6c, 0x32, 0x3a, 0x0a, 0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x31, 0x12,
	0x30, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x31, 0x12, 0x0f, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x32, 0x3a, 0x0a, 0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0x12, 0x30, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0x12, 0x0f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hello_proto_goTypes = []interface{}{
	(*HelloReq)(nil),  // 0: hello.HelloReq
	(*HelloResp)(nil), // 1: hello.HelloResp
}
var file_hello_proto_depIdxs = []int32{
	0, // 0: hello.Hello1.GetHello1:input_type -> hello.HelloReq
	0, // 1: hello.Hello2.GetHello2:input_type -> hello.HelloReq
	1, // 2: hello.Hello1.GetHello1:output_type -> hello.HelloResp
	1, // 3: hello.Hello2.GetHello2:output_type -> hello.HelloResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
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
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResp); i {
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
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Hello1Client is the client API for Hello1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Hello1Client interface {
	GetHello1(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type hello1Client struct {
	cc grpc.ClientConnInterface
}

func NewHello1Client(cc grpc.ClientConnInterface) Hello1Client {
	return &hello1Client{cc}
}

func (c *hello1Client) GetHello1(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.cc.Invoke(ctx, "/hello.Hello1/GetHello1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Hello1Server is the server API for Hello1 service.
type Hello1Server interface {
	GetHello1(context.Context, *HelloReq) (*HelloResp, error)
}

// UnimplementedHello1Server can be embedded to have forward compatible implementations.
type UnimplementedHello1Server struct {
}

func (*UnimplementedHello1Server) GetHello1(context.Context, *HelloReq) (*HelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHello1 not implemented")
}

func RegisterHello1Server(s *grpc.Server, srv Hello1Server) {
	s.RegisterService(&_Hello1_serviceDesc, srv)
}

func _Hello1_GetHello1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Hello1Server).GetHello1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello1/GetHello1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Hello1Server).GetHello1(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello1",
	HandlerType: (*Hello1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHello1",
			Handler:    _Hello1_GetHello1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

// Hello2Client is the client API for Hello2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Hello2Client interface {
	GetHello2(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type hello2Client struct {
	cc grpc.ClientConnInterface
}

func NewHello2Client(cc grpc.ClientConnInterface) Hello2Client {
	return &hello2Client{cc}
}

func (c *hello2Client) GetHello2(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.cc.Invoke(ctx, "/hello.Hello2/GetHello2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Hello2Server is the server API for Hello2 service.
type Hello2Server interface {
	GetHello2(context.Context, *HelloReq) (*HelloResp, error)
}

// UnimplementedHello2Server can be embedded to have forward compatible implementations.
type UnimplementedHello2Server struct {
}

func (*UnimplementedHello2Server) GetHello2(context.Context, *HelloReq) (*HelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHello2 not implemented")
}

func RegisterHello2Server(s *grpc.Server, srv Hello2Server) {
	s.RegisterService(&_Hello2_serviceDesc, srv)
}

func _Hello2_GetHello2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Hello2Server).GetHello2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello2/GetHello2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Hello2Server).GetHello2(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello2",
	HandlerType: (*Hello2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHello2",
			Handler:    _Hello2_GetHello2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
