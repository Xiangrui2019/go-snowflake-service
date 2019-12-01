// Code generated by protoc-gen-go. DO NOT EDIT.
// source: id-generator-service.proto

package pb

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

type GenerateIdRequest struct {
	NodeId               int64    `protobuf:"varint,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateIdRequest) Reset()         { *m = GenerateIdRequest{} }
func (m *GenerateIdRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateIdRequest) ProtoMessage()    {}
func (*GenerateIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_792af50a6fe16ef8, []int{0}
}

func (m *GenerateIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateIdRequest.Unmarshal(m, b)
}
func (m *GenerateIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateIdRequest.Marshal(b, m, deterministic)
}
func (m *GenerateIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateIdRequest.Merge(m, src)
}
func (m *GenerateIdRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateIdRequest.Size(m)
}
func (m *GenerateIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateIdRequest proto.InternalMessageInfo

func (m *GenerateIdRequest) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

type GenerateIdReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateIdReply) Reset()         { *m = GenerateIdReply{} }
func (m *GenerateIdReply) String() string { return proto.CompactTextString(m) }
func (*GenerateIdReply) ProtoMessage()    {}
func (*GenerateIdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_792af50a6fe16ef8, []int{1}
}

func (m *GenerateIdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateIdReply.Unmarshal(m, b)
}
func (m *GenerateIdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateIdReply.Marshal(b, m, deterministic)
}
func (m *GenerateIdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateIdReply.Merge(m, src)
}
func (m *GenerateIdReply) XXX_Size() int {
	return xxx_messageInfo_GenerateIdReply.Size(m)
}
func (m *GenerateIdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateIdReply.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateIdReply proto.InternalMessageInfo

func (m *GenerateIdReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GenerateIdReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateIdRequest)(nil), "pb.GenerateIdRequest")
	proto.RegisterType((*GenerateIdReply)(nil), "pb.GenerateIdReply")
}

func init() { proto.RegisterFile("id-generator-service.proto", fileDescriptor_792af50a6fe16ef8) }

var fileDescriptor_792af50a6fe16ef8 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4c, 0xd1, 0x4d,
	0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe6, 0x12, 0x74, 0x87,
	0x48, 0xa7, 0x7a, 0xa6, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x71, 0xb1, 0xf9,
	0xe5, 0xa7, 0xa4, 0x7a, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x41, 0x79, 0x4a, 0xe6,
	0x5c, 0xfc, 0xc8, 0x8a, 0x0b, 0x72, 0x2a, 0x85, 0xf8, 0xb8, 0x98, 0xe0, 0xca, 0x98, 0x3c, 0x53,
	0x84, 0x44, 0xb8, 0x58, 0x5d, 0x8b, 0x8a, 0xf2, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0x20, 0x1c, 0xa3, 0x00, 0x2e, 0x21, 0xcf, 0x14, 0x77, 0x98, 0x33, 0x82, 0x21, 0xae, 0x10, 0xb2,
	0xe2, 0xe2, 0x42, 0x18, 0x27, 0x24, 0xaa, 0x57, 0x90, 0xa4, 0x87, 0xe1, 0x16, 0x29, 0x61, 0x74,
	0xe1, 0x82, 0x9c, 0x4a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x17, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x71, 0x94, 0x2c, 0x02, 0xe0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdGeneratorServiceClient is the client API for IdGeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdGeneratorServiceClient interface {
	GenerateId(ctx context.Context, in *GenerateIdRequest, opts ...grpc.CallOption) (*GenerateIdReply, error)
}

type idGeneratorServiceClient struct {
	cc *grpc.ClientConn
}

func NewIdGeneratorServiceClient(cc *grpc.ClientConn) IdGeneratorServiceClient {
	return &idGeneratorServiceClient{cc}
}

func (c *idGeneratorServiceClient) GenerateId(ctx context.Context, in *GenerateIdRequest, opts ...grpc.CallOption) (*GenerateIdReply, error) {
	out := new(GenerateIdReply)
	err := c.cc.Invoke(ctx, "/pb.IdGeneratorService/GenerateId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdGeneratorServiceServer is the server API for IdGeneratorService service.
type IdGeneratorServiceServer interface {
	GenerateId(context.Context, *GenerateIdRequest) (*GenerateIdReply, error)
}

// UnimplementedIdGeneratorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIdGeneratorServiceServer struct {
}

func (*UnimplementedIdGeneratorServiceServer) GenerateId(ctx context.Context, req *GenerateIdRequest) (*GenerateIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateId not implemented")
}

func RegisterIdGeneratorServiceServer(s *grpc.Server, srv IdGeneratorServiceServer) {
	s.RegisterService(&_IdGeneratorService_serviceDesc, srv)
}

func _IdGeneratorService_GenerateId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdGeneratorServiceServer).GenerateId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IdGeneratorService/GenerateId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdGeneratorServiceServer).GenerateId(ctx, req.(*GenerateIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdGeneratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IdGeneratorService",
	HandlerType: (*IdGeneratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateId",
			Handler:    _IdGeneratorService_GenerateId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "id-generator-service.proto",
}