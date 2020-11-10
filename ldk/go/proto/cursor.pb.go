// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: cursor.proto

package proto

import (
	context "context"
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

type CursorPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *CursorPositionRequest) Reset() {
	*x = CursorPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cursor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CursorPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CursorPositionRequest) ProtoMessage() {}

func (x *CursorPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cursor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CursorPositionRequest.ProtoReflect.Descriptor instead.
func (*CursorPositionRequest) Descriptor() ([]byte, []int) {
	return file_cursor_proto_rawDescGZIP(), []int{0}
}

func (x *CursorPositionRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type CursorPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X      int32  `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32  `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Screen uint32 `protobuf:"varint,3,opt,name=screen,proto3" json:"screen,omitempty"`
}

func (x *CursorPositionResponse) Reset() {
	*x = CursorPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cursor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CursorPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CursorPositionResponse) ProtoMessage() {}

func (x *CursorPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cursor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CursorPositionResponse.ProtoReflect.Descriptor instead.
func (*CursorPositionResponse) Descriptor() ([]byte, []int) {
	return file_cursor_proto_rawDescGZIP(), []int{1}
}

func (x *CursorPositionResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *CursorPositionResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *CursorPositionResponse) GetScreen() uint32 {
	if x != nil {
		return x.Screen
	}
	return 0
}

type CursorPositionStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *CursorPositionStreamRequest) Reset() {
	*x = CursorPositionStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cursor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CursorPositionStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CursorPositionStreamRequest) ProtoMessage() {}

func (x *CursorPositionStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cursor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CursorPositionStreamRequest.ProtoReflect.Descriptor instead.
func (*CursorPositionStreamRequest) Descriptor() ([]byte, []int) {
	return file_cursor_proto_rawDescGZIP(), []int{2}
}

func (x *CursorPositionStreamRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type CursorPositionStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X      int32  `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32  `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Screen uint32 `protobuf:"varint,3,opt,name=screen,proto3" json:"screen,omitempty"`
	Error  string `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CursorPositionStreamResponse) Reset() {
	*x = CursorPositionStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cursor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CursorPositionStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CursorPositionStreamResponse) ProtoMessage() {}

func (x *CursorPositionStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cursor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CursorPositionStreamResponse.ProtoReflect.Descriptor instead.
func (*CursorPositionStreamResponse) Descriptor() ([]byte, []int) {
	return file_cursor_proto_rawDescGZIP(), []int{3}
}

func (x *CursorPositionStreamResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *CursorPositionStreamResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *CursorPositionStreamResponse) GetScreen() uint32 {
	if x != nil {
		return x.Screen
	}
	return 0
}

func (x *CursorPositionStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_cursor_proto protoreflect.FileDescriptor

var file_cursor_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x16, 0x43, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x1b, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x68,
	0x0a, 0x1c, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xba, 0x01, 0x0a, 0x06, 0x43, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x61, 0x0a, 0x14, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cursor_proto_rawDescOnce sync.Once
	file_cursor_proto_rawDescData = file_cursor_proto_rawDesc
)

func file_cursor_proto_rawDescGZIP() []byte {
	file_cursor_proto_rawDescOnce.Do(func() {
		file_cursor_proto_rawDescData = protoimpl.X.CompressGZIP(file_cursor_proto_rawDescData)
	})
	return file_cursor_proto_rawDescData
}

var file_cursor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cursor_proto_goTypes = []interface{}{
	(*CursorPositionRequest)(nil),        // 0: proto.CursorPositionRequest
	(*CursorPositionResponse)(nil),       // 1: proto.CursorPositionResponse
	(*CursorPositionStreamRequest)(nil),  // 2: proto.CursorPositionStreamRequest
	(*CursorPositionStreamResponse)(nil), // 3: proto.CursorPositionStreamResponse
	(*Session)(nil),                      // 4: proto.Session
}
var file_cursor_proto_depIdxs = []int32{
	4, // 0: proto.CursorPositionRequest.session:type_name -> proto.Session
	4, // 1: proto.CursorPositionStreamRequest.session:type_name -> proto.Session
	0, // 2: proto.Cursor.CursorPosition:input_type -> proto.CursorPositionRequest
	2, // 3: proto.Cursor.CursorPositionStream:input_type -> proto.CursorPositionStreamRequest
	1, // 4: proto.Cursor.CursorPosition:output_type -> proto.CursorPositionResponse
	3, // 5: proto.Cursor.CursorPositionStream:output_type -> proto.CursorPositionStreamResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cursor_proto_init() }
func file_cursor_proto_init() {
	if File_cursor_proto != nil {
		return
	}
	file_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cursor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CursorPositionRequest); i {
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
		file_cursor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CursorPositionResponse); i {
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
		file_cursor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CursorPositionStreamRequest); i {
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
		file_cursor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CursorPositionStreamResponse); i {
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
			RawDescriptor: file_cursor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cursor_proto_goTypes,
		DependencyIndexes: file_cursor_proto_depIdxs,
		MessageInfos:      file_cursor_proto_msgTypes,
	}.Build()
	File_cursor_proto = out.File
	file_cursor_proto_rawDesc = nil
	file_cursor_proto_goTypes = nil
	file_cursor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CursorClient is the client API for Cursor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CursorClient interface {
	// get the current position of the cursor
	CursorPosition(ctx context.Context, in *CursorPositionRequest, opts ...grpc.CallOption) (*CursorPositionResponse, error)
	// get the position of the cursor every time it changes
	CursorPositionStream(ctx context.Context, in *CursorPositionStreamRequest, opts ...grpc.CallOption) (Cursor_CursorPositionStreamClient, error)
}

type cursorClient struct {
	cc grpc.ClientConnInterface
}

func NewCursorClient(cc grpc.ClientConnInterface) CursorClient {
	return &cursorClient{cc}
}

func (c *cursorClient) CursorPosition(ctx context.Context, in *CursorPositionRequest, opts ...grpc.CallOption) (*CursorPositionResponse, error) {
	out := new(CursorPositionResponse)
	err := c.cc.Invoke(ctx, "/proto.Cursor/CursorPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cursorClient) CursorPositionStream(ctx context.Context, in *CursorPositionStreamRequest, opts ...grpc.CallOption) (Cursor_CursorPositionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Cursor_serviceDesc.Streams[0], "/proto.Cursor/CursorPositionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &cursorCursorPositionStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cursor_CursorPositionStreamClient interface {
	Recv() (*CursorPositionStreamResponse, error)
	grpc.ClientStream
}

type cursorCursorPositionStreamClient struct {
	grpc.ClientStream
}

func (x *cursorCursorPositionStreamClient) Recv() (*CursorPositionStreamResponse, error) {
	m := new(CursorPositionStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CursorServer is the server API for Cursor service.
type CursorServer interface {
	// get the current position of the cursor
	CursorPosition(context.Context, *CursorPositionRequest) (*CursorPositionResponse, error)
	// get the position of the cursor every time it changes
	CursorPositionStream(*CursorPositionStreamRequest, Cursor_CursorPositionStreamServer) error
}

// UnimplementedCursorServer can be embedded to have forward compatible implementations.
type UnimplementedCursorServer struct {
}

func (*UnimplementedCursorServer) CursorPosition(context.Context, *CursorPositionRequest) (*CursorPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CursorPosition not implemented")
}
func (*UnimplementedCursorServer) CursorPositionStream(*CursorPositionStreamRequest, Cursor_CursorPositionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CursorPositionStream not implemented")
}

func RegisterCursorServer(s *grpc.Server, srv CursorServer) {
	s.RegisterService(&_Cursor_serviceDesc, srv)
}

func _Cursor_CursorPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CursorPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CursorServer).CursorPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cursor/CursorPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CursorServer).CursorPosition(ctx, req.(*CursorPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cursor_CursorPositionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CursorPositionStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CursorServer).CursorPositionStream(m, &cursorCursorPositionStreamServer{stream})
}

type Cursor_CursorPositionStreamServer interface {
	Send(*CursorPositionStreamResponse) error
	grpc.ServerStream
}

type cursorCursorPositionStreamServer struct {
	grpc.ServerStream
}

func (x *cursorCursorPositionStreamServer) Send(m *CursorPositionStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Cursor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Cursor",
	HandlerType: (*CursorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CursorPosition",
			Handler:    _Cursor_CursorPosition_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CursorPositionStream",
			Handler:       _Cursor_CursorPositionStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cursor.proto",
}
