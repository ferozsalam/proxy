// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v3/tapds.proto

package envoy_service_tap_v3

import (
	context "context"
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/tap/v3"
	v31 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// [#not-implemented-hide:] A tap resource is essentially a tap configuration with a name
// The filter TapDS config references this name.
type TapResource struct {
	// The name of the tap configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Tap config to apply
	Config               *v3.TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TapResource) Reset()         { *m = TapResource{} }
func (m *TapResource) String() string { return proto.CompactTextString(m) }
func (*TapResource) ProtoMessage()    {}
func (*TapResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_d118f9afd670f788, []int{0}
}

func (m *TapResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapResource.Unmarshal(m, b)
}
func (m *TapResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapResource.Marshal(b, m, deterministic)
}
func (m *TapResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapResource.Merge(m, src)
}
func (m *TapResource) XXX_Size() int {
	return xxx_messageInfo_TapResource.Size(m)
}
func (m *TapResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TapResource.DiscardUnknown(m)
}

var xxx_messageInfo_TapResource proto.InternalMessageInfo

func (m *TapResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TapResource) GetConfig() *v3.TapConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*TapResource)(nil), "envoy.service.tap.v3.TapResource")
}

func init() { proto.RegisterFile("envoy/service/tap/v3/tapds.proto", fileDescriptor_d118f9afd670f788) }

var fileDescriptor_d118f9afd670f788 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x9d, 0x6d, 0x89, 0x38, 0x3d, 0xb4, 0xac, 0x82, 0x35, 0xd5, 0x12, 0x03, 0x62, 0x08,
	0x75, 0x56, 0x13, 0xf0, 0x10, 0xf0, 0x12, 0x8b, 0xe7, 0x92, 0xe6, 0x2e, 0xcf, 0xcd, 0x33, 0x1d,
	0xc8, 0xce, 0x1b, 0x67, 0x66, 0x07, 0x73, 0x11, 0xf1, 0xd4, 0xa3, 0xe0, 0xad, 0xdf, 0xc0, 0xaf,
	0xe0, 0x5d, 0xf0, 0x2a, 0x7e, 0x05, 0x3f, 0x85, 0x27, 0xd9, 0x9d, 0xdd, 0x98, 0xa5, 0x41, 0xf4,
	0xd2, 0xd3, 0xee, 0xbe, 0xf7, 0x7b, 0xef, 0x3f, 0xf3, 0xdf, 0xf7, 0x78, 0x07, 0x95, 0xa7, 0x65,
	0x62, 0xd1, 0x78, 0x99, 0x62, 0xe2, 0x40, 0x27, 0x7e, 0x58, 0x3c, 0x66, 0x56, 0x68, 0x43, 0x8e,
	0xe2, 0x5b, 0x25, 0x21, 0x2a, 0x42, 0x38, 0xd0, 0xc2, 0x0f, 0xdb, 0x55, 0x5d, 0x4a, 0xea, 0xb5,
	0x9c, 0xd7, 0x65, 0x29, 0x65, 0x19, 0xa9, 0x50, 0xd7, 0xee, 0x37, 0x3b, 0xcf, 0xa4, 0x4d, 0xc9,
	0xa3, 0x59, 0x16, 0xe0, 0xea, 0xa3, 0x62, 0xef, 0xce, 0x89, 0xe6, 0x0b, 0x4c, 0x40, 0xcb, 0x04,
	0x94, 0x22, 0x07, 0x4e, 0x92, 0xaa, 0x4e, 0xd0, 0xbe, 0x97, 0xcf, 0x34, 0xac, 0xc7, 0x13, 0xeb,
	0xc0, 0xe5, 0x75, 0xfa, 0xfe, 0xa5, 0xb4, 0x47, 0x63, 0x25, 0x29, 0xa9, 0xe6, 0x15, 0x72, 0xdb,
	0xc3, 0x42, 0xce, 0xc0, 0x61, 0x52, 0xbf, 0x84, 0x44, 0xf7, 0x23, 0xe3, 0x3b, 0x53, 0xd0, 0x13,
	0xb4, 0x94, 0x9b, 0x14, 0xe3, 0x03, 0xbe, 0xad, 0x20, 0xc3, 0x7d, 0xd6, 0x61, 0xbd, 0x1b, 0xe3,
	0xeb, 0xbf, 0xc6, 0xdb, 0x26, 0xea, 0xb0, 0x49, 0x19, 0x8c, 0x9f, 0xf2, 0x56, 0xb8, 0xef, 0x7e,
	0xd4, 0x61, 0xbd, 0x9d, 0xc1, 0xa1, 0x08, 0xd6, 0x84, 0x60, 0xe5, 0x8c, 0x98, 0x82, 0x7e, 0x5e,
	0x06, 0x26, 0x15, 0x3d, 0x3a, 0xba, 0xf8, 0x7a, 0x7e, 0xf8, 0x90, 0x3f, 0xd8, 0x60, 0xe4, 0x00,
	0x16, 0xfa, 0x0c, 0xc4, 0xda, 0x11, 0x06, 0x9f, 0xb7, 0xf8, 0xcd, 0x29, 0xe8, 0xe3, 0xda, 0xa2,
	0xd3, 0xc0, 0xc7, 0x39, 0xdf, 0x3b, 0x75, 0x06, 0x21, 0x5b, 0x09, 0xd8, 0xf8, 0x48, 0x34, 0x7b,
	0xfe, 0xf1, 0xd5, 0x0f, 0xc5, 0xaa, 0xc3, 0x04, 0xdf, 0xe4, 0x68, 0x5d, 0xfb, 0xd1, 0x3f, 0xd2,
	0x56, 0x93, 0xb2, 0xd8, 0xbd, 0xd6, 0x63, 0x8f, 0x59, 0xfc, 0x8e, 0xef, 0x1e, 0xe3, 0xc2, 0xc1,
	0x9a, 0xea, 0x93, 0xbf, 0xf6, 0x29, 0xe0, 0x4b, 0xd2, 0x83, 0xff, 0x29, 0x69, 0xe8, 0x5f, 0x30,
	0xbe, 0xfb, 0x02, 0x5d, 0x7a, 0x76, 0x55, 0xd7, 0xee, 0x7f, 0xf8, 0xf1, 0xf3, 0x53, 0x74, 0xd0,
	0xbd, 0xd3, 0x18, 0xd4, 0x91, 0x03, 0xfd, 0x32, 0xfc, 0x4f, 0x5b, 0x02, 0x5b, 0x23, 0xd6, 0x1f,
	0x3f, 0xfb, 0xf2, 0xfe, 0xdb, 0xf7, 0x56, 0xb4, 0x17, 0xf1, 0xae, 0xa4, 0x20, 0xa3, 0x0d, 0xbd,
	0x5d, 0x8a, 0x4d, 0x3b, 0x33, 0xe6, 0xd3, 0x62, 0xad, 0x4e, 0x8a, 0xc1, 0x3b, 0x61, 0xe7, 0x8c,
	0xbd, 0x6a, 0x95, 0x43, 0x38, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x93, 0x88, 0x67, 0x85,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v31.DiscoveryRequest, opts ...grpc.CallOption) (*v31.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapDiscoveryServiceClient(cc *grpc.ClientConn) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v3.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v31.DiscoveryRequest) error
	Recv() (*v31.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v31.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v31.DiscoveryResponse, error) {
	m := new(v31.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v3.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v31.DeltaDiscoveryRequest) error
	Recv() (*v31.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v31.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v31.DeltaDiscoveryResponse, error) {
	m := new(v31.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v31.DiscoveryRequest, opts ...grpc.CallOption) (*v31.DiscoveryResponse, error) {
	out := new(v31.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v3.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v31.DiscoveryRequest) (*v31.DiscoveryResponse, error)
}

// UnimplementedTapDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapDiscoveryServiceServer struct {
}

func (*UnimplementedTapDiscoveryServiceServer) StreamTapConfigs(srv TapDiscoveryService_StreamTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) DeltaTapConfigs(srv TapDiscoveryService_DeltaTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) FetchTapConfigs(ctx context.Context, req *v31.DiscoveryRequest) (*v31.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTapConfigs not implemented")
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v31.DiscoveryResponse) error
	Recv() (*v31.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v31.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v31.DiscoveryRequest, error) {
	m := new(v31.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v31.DeltaDiscoveryResponse) error
	Recv() (*v31.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v31.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v31.DeltaDiscoveryRequest, error) {
	m := new(v31.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v31.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v3.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v31.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v3.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v3/tapds.proto",
}
