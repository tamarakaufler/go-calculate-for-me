// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/gcd/v1/gcd.proto

package pb_gcd_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GCDRequest struct {
	A                    uint64   `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B                    uint64   `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCDRequest) Reset()         { *m = GCDRequest{} }
func (m *GCDRequest) String() string { return proto.CompactTextString(m) }
func (*GCDRequest) ProtoMessage()    {}
func (*GCDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gcd_b2c73431b408b218, []int{0}
}
func (m *GCDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCDRequest.Unmarshal(m, b)
}
func (m *GCDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCDRequest.Marshal(b, m, deterministic)
}
func (dst *GCDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCDRequest.Merge(dst, src)
}
func (m *GCDRequest) XXX_Size() int {
	return xxx_messageInfo_GCDRequest.Size(m)
}
func (m *GCDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GCDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GCDRequest proto.InternalMessageInfo

func (m *GCDRequest) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *GCDRequest) GetB() uint64 {
	if m != nil {
		return m.B
	}
	return 0
}

type GCDResponse struct {
	Result               uint64   `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCDResponse) Reset()         { *m = GCDResponse{} }
func (m *GCDResponse) String() string { return proto.CompactTextString(m) }
func (*GCDResponse) ProtoMessage()    {}
func (*GCDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gcd_b2c73431b408b218, []int{1}
}
func (m *GCDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCDResponse.Unmarshal(m, b)
}
func (m *GCDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCDResponse.Marshal(b, m, deterministic)
}
func (dst *GCDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCDResponse.Merge(dst, src)
}
func (m *GCDResponse) XXX_Size() int {
	return xxx_messageInfo_GCDResponse.Size(m)
}
func (m *GCDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GCDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GCDResponse proto.InternalMessageInfo

func (m *GCDResponse) GetResult() uint64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*GCDRequest)(nil), "pb.gcd.v1.GCDRequest")
	proto.RegisterType((*GCDResponse)(nil), "pb.gcd.v1.GCDResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GCDServiceClient is the client API for GCDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GCDServiceClient interface {
	Compute(ctx context.Context, in *GCDRequest, opts ...grpc.CallOption) (*GCDResponse, error)
}

type gCDServiceClient struct {
	cc *grpc.ClientConn
}

func NewGCDServiceClient(cc *grpc.ClientConn) GCDServiceClient {
	return &gCDServiceClient{cc}
}

func (c *gCDServiceClient) Compute(ctx context.Context, in *GCDRequest, opts ...grpc.CallOption) (*GCDResponse, error) {
	out := new(GCDResponse)
	err := c.cc.Invoke(ctx, "/pb.gcd.v1.GCDService/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GCDService service

type GCDServiceServer interface {
	Compute(context.Context, *GCDRequest) (*GCDResponse, error)
}

func RegisterGCDServiceServer(s *grpc.Server, srv GCDServiceServer) {
	s.RegisterService(&_GCDService_serviceDesc, srv)
}

func _GCDService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GCDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GCDServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gcd.v1.GCDService/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GCDServiceServer).Compute(ctx, req.(*GCDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GCDService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.gcd.v1.GCDService",
	HandlerType: (*GCDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _GCDService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/gcd/v1/gcd.proto",
}

func init() { proto.RegisterFile("pb/gcd/v1/gcd.proto", fileDescriptor_gcd_b2c73431b408b218) }

var fileDescriptor_gcd_b2c73431b408b218 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0xd2, 0x4f,
	0x4f, 0x4e, 0xd1, 0x2f, 0x33, 0x04, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x05,
	0x49, 0x7a, 0x20, 0x5e, 0x99, 0xa1, 0x92, 0x06, 0x17, 0x97, 0xbb, 0xb3, 0x4b, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x89, 0x10, 0x0f, 0x17, 0x63, 0xa2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10,
	0x63, 0x22, 0x88, 0x97, 0x24, 0xc1, 0x04, 0xe1, 0x25, 0x29, 0xa9, 0x72, 0x71, 0x83, 0x55, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x40,
	0xd5, 0x43, 0x79, 0x46, 0x1e, 0x60, 0x03, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xac,
	0xb8, 0xd8, 0x9d, 0xf3, 0x73, 0x0b, 0x4a, 0x4b, 0x52, 0x85, 0x44, 0xf5, 0xe0, 0xb6, 0xea, 0x21,
	0xac, 0x94, 0x12, 0x43, 0x17, 0x86, 0x98, 0xaf, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xac, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0xd5, 0xcf, 0x97, 0xc3, 0x00, 0x00, 0x00,
}
