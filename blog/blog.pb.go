// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package blog

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SaveRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{1}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaveRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaveRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Post struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{2}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "blog.GetRequest")
	proto.RegisterType((*SaveRequest)(nil), "blog.SaveRequest")
	proto.RegisterType((*Post)(nil), "blog.Post")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1) }

var fileDescriptor_6745b25902462fb1 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xc9, 0x4f,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x64, 0xb8, 0xb8, 0xdc, 0x53,
	0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x32, 0x53, 0x94, 0xdc, 0xb9, 0xb8, 0x83, 0x13, 0xcb, 0x52,
	0x71, 0x48, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0xcc, 0x60, 0x41,
	0x30, 0x5b, 0xc9, 0x81, 0x8b, 0x25, 0x20, 0x9f, 0x12, 0x13, 0x8c, 0xe2, 0xb8, 0x58, 0x9c, 0x72,
	0xf2, 0xd3, 0x85, 0x34, 0xb9, 0xd8, 0xdd, 0x53, 0x4b, 0xc0, 0x86, 0x09, 0xe8, 0x81, 0xbd, 0x83,
	0x70, 0xbf, 0x14, 0x17, 0x44, 0x04, 0x24, 0xab, 0xc4, 0x20, 0xa4, 0xcd, 0xc5, 0x01, 0x72, 0x3d,
	0x58, 0xad, 0x20, 0x44, 0x06, 0xc9, 0x37, 0xa8, 0x8a, 0x93, 0xd8, 0xc0, 0xa1, 0x62, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0xd5, 0xc2, 0x96, 0x23, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogClient interface {
	GetPost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Post, error)
	SavePost(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*Post, error)
}

type blogClient struct {
	cc *grpc.ClientConn
}

func NewBlogClient(cc *grpc.ClientConn) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) GetPost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) SavePost(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/SavePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServer is the server API for Blog service.
type BlogServer interface {
	GetPost(context.Context, *GetRequest) (*Post, error)
	SavePost(context.Context, *SaveRequest) (*Post, error)
}

// UnimplementedBlogServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (*UnimplementedBlogServer) GetPost(ctx context.Context, req *GetRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (*UnimplementedBlogServer) SavePost(ctx context.Context, req *SaveRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePost not implemented")
}

func RegisterBlogServer(s *grpc.Server, srv BlogServer) {
	s.RegisterService(&_Blog_serviceDesc, srv)
}

func _Blog_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetPost(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_SavePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).SavePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/SavePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).SavePost(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPost",
			Handler:    _Blog_GetPost_Handler,
		},
		{
			MethodName: "SavePost",
			Handler:    _Blog_SavePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog.proto",
}
