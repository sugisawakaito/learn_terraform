// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messenger.proto

package messenger

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b99aba0cbf4e4b91, []int{0}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MessageResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b99aba0cbf4e4b91, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "messenger.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "messenger.MessageResponse")
}

func init() {
	proto.RegisterFile("messenger.proto", fileDescriptor_b99aba0cbf4e4b91)
}

var fileDescriptor_b99aba0cbf4e4b91 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0xcd, 0x4b, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x48,
	0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x25, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73,
	0x0b, 0x4a, 0x2a, 0x21, 0xea, 0x94, 0xb4, 0xb8, 0xf8, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x21, 0x22, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x36, 0x17, 0x3f, 0x5c, 0x6d, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0x2a, 0x6e, 0xc5, 0x46, 0x73, 0x18, 0xb9, 0x38, 0x7d, 0x61, 0x6e, 0x10, 0x72, 0xe5,
	0xe2, 0x76, 0x4f, 0x2d, 0x81, 0xea, 0x2e, 0x16, 0x12, 0xd3, 0x83, 0xb8, 0x49, 0x0f, 0xe6, 0x26,
	0x3d, 0x57, 0x90, 0x9b, 0xa4, 0xa4, 0xf4, 0x10, 0xfe, 0x40, 0xb3, 0x4a, 0x89, 0xc1, 0x80, 0x51,
	0xc8, 0x83, 0x8b, 0xd7, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x15, 0x2a, 0x29, 0x24, 0x89, 0x4d, 0x03,
	0xd8, 0x1f, 0xf8, 0xcd, 0x4a, 0x62, 0x03, 0xdb, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x34, 0x94, 0x0a, 0x39, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessengerClient interface {
	GetMessages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Messenger_GetMessagesClient, error)
	CreateMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) GetMessages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Messenger_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messenger_serviceDesc.Streams[0], "/messenger.Messenger/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messengerGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messenger_GetMessagesClient interface {
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type messengerGetMessagesClient struct {
	grpc.ClientStream
}

func (x *messengerGetMessagesClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messengerClient) CreateMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/messenger.Messenger/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServer is the server API for Messenger service.
type MessengerServer interface {
	GetMessages(*empty.Empty, Messenger_GetMessagesServer) error
	CreateMessage(context.Context, *MessageRequest) (*MessageResponse, error)
}

// UnimplementedMessengerServer can be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (*UnimplementedMessengerServer) GetMessages(req *empty.Empty, srv Messenger_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (*UnimplementedMessengerServer) CreateMessage(ctx context.Context, req *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}

func RegisterMessengerServer(s *grpc.Server, srv MessengerServer) {
	s.RegisterService(&_Messenger_serviceDesc, srv)
}

func _Messenger_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessengerServer).GetMessages(m, &messengerGetMessagesServer{stream})
}

type Messenger_GetMessagesServer interface {
	Send(*MessageResponse) error
	grpc.ServerStream
}

type messengerGetMessagesServer struct {
	grpc.ServerStream
}

func (x *messengerGetMessagesServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Messenger_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.Messenger/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).CreateMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messenger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _Messenger_CreateMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMessages",
			Handler:       _Messenger_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messenger.proto",
}
