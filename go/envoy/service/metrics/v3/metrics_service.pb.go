// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/metrics/v3/metrics_service.proto

package envoy_service_metrics_v3

import (
	context "context"
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_go "github.com/prometheus/client_model/go"
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

type StreamMetricsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMetricsResponse) Reset()         { *m = StreamMetricsResponse{} }
func (m *StreamMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsResponse) ProtoMessage()    {}
func (*StreamMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1642cfdc06c3c66, []int{0}
}

func (m *StreamMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsResponse.Unmarshal(m, b)
}
func (m *StreamMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsResponse.Marshal(b, m, deterministic)
}
func (m *StreamMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsResponse.Merge(m, src)
}
func (m *StreamMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsResponse.Size(m)
}
func (m *StreamMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsResponse proto.InternalMessageInfo

type StreamMetricsMessage struct {
	// Identifier data effectively is a structured metadata. As a performance optimization this will
	// only be sent in the first message on the stream.
	Identifier *StreamMetricsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// A list of metric entries
	EnvoyMetrics         []*_go.MetricFamily `protobuf:"bytes,2,rep,name=envoy_metrics,json=envoyMetrics,proto3" json:"envoy_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StreamMetricsMessage) Reset()         { *m = StreamMetricsMessage{} }
func (m *StreamMetricsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage) ProtoMessage()    {}
func (*StreamMetricsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1642cfdc06c3c66, []int{1}
}

func (m *StreamMetricsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsMessage.Unmarshal(m, b)
}
func (m *StreamMetricsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsMessage.Marshal(b, m, deterministic)
}
func (m *StreamMetricsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage.Merge(m, src)
}
func (m *StreamMetricsMessage) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsMessage.Size(m)
}
func (m *StreamMetricsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage proto.InternalMessageInfo

func (m *StreamMetricsMessage) GetIdentifier() *StreamMetricsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamMetricsMessage) GetEnvoyMetrics() []*_go.MetricFamily {
	if m != nil {
		return m.EnvoyMetrics
	}
	return nil
}

type StreamMetricsMessage_Identifier struct {
	// The node sending metrics over the stream.
	Node                 *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMetricsMessage_Identifier) Reset()         { *m = StreamMetricsMessage_Identifier{} }
func (m *StreamMetricsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage_Identifier) ProtoMessage()    {}
func (*StreamMetricsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1642cfdc06c3c66, []int{1, 0}
}

func (m *StreamMetricsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamMetricsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamMetricsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage_Identifier.Merge(m, src)
}
func (m *StreamMetricsMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Size(m)
}
func (m *StreamMetricsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage_Identifier proto.InternalMessageInfo

func (m *StreamMetricsMessage_Identifier) GetNode() *v3.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamMetricsResponse)(nil), "envoy.service.metrics.v3.StreamMetricsResponse")
	proto.RegisterType((*StreamMetricsMessage)(nil), "envoy.service.metrics.v3.StreamMetricsMessage")
	proto.RegisterType((*StreamMetricsMessage_Identifier)(nil), "envoy.service.metrics.v3.StreamMetricsMessage.Identifier")
}

func init() {
	proto.RegisterFile("envoy/service/metrics/v3/metrics_service.proto", fileDescriptor_f1642cfdc06c3c66)
}

var fileDescriptor_f1642cfdc06c3c66 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0xcd, 0xa8, 0x8b, 0x64, 0xad, 0x2c, 0xa3, 0x62, 0x19, 0x50, 0xd7, 0x1e, 0xa4, 0x17,
	0x13, 0x99, 0x51, 0x58, 0xe7, 0x22, 0xcc, 0x41, 0x51, 0xd8, 0xa5, 0x4c, 0x4f, 0x9e, 0x4a, 0x3a,
	0xf3, 0xb6, 0x06, 0x3a, 0xc9, 0x90, 0xa4, 0x83, 0xbd, 0x79, 0xd2, 0xe2, 0x47, 0xf0, 0xa3, 0x78,
	0x17, 0xbc, 0xfa, 0x75, 0x7a, 0x92, 0x99, 0x64, 0xea, 0x96, 0xb6, 0xd0, 0xde, 0x92, 0xbc, 0xcf,
	0xf3, 0x7b, 0xff, 0xe4, 0xc5, 0x04, 0x44, 0x25, 0x17, 0x54, 0x83, 0xaa, 0x78, 0x06, 0xb4, 0x00,
	0xa3, 0x78, 0xa6, 0x69, 0x15, 0xb5, 0xc7, 0x91, 0x0b, 0x91, 0x52, 0x49, 0x23, 0xfd, 0x6e, 0xa3,
	0x27, 0xed, 0xa3, 0x13, 0x91, 0x2a, 0x0a, 0x9e, 0x5a, 0x52, 0x26, 0xc5, 0x84, 0x4f, 0x69, 0x26,
	0x15, 0xd4, 0x94, 0x31, 0xd3, 0xce, 0x1a, 0x74, 0x5a, 0xb1, 0xbd, 0x3e, 0x9e, 0xe7, 0x25, 0xa3,
	0x4c, 0x08, 0x69, 0x98, 0xe1, 0x52, 0x68, 0xaa, 0x0d, 0x33, 0xf3, 0x36, 0xfc, 0x6c, 0x2b, 0x5c,
	0x81, 0xd2, 0x5c, 0x0a, 0x2e, 0xa6, 0x4e, 0xf2, 0xa8, 0x62, 0x33, 0x9e, 0x33, 0x03, 0xb4, 0x3d,
	0xd8, 0x40, 0xef, 0x0a, 0x3f, 0x1c, 0x1a, 0x05, 0xac, 0xb8, 0xb4, 0x19, 0x53, 0xd0, 0xa5, 0x14,
	0x1a, 0xe2, 0xd7, 0x3f, 0x7f, 0x2f, 0x9f, 0xbc, 0x74, 0x4d, 0x6f, 0x37, 0x11, 0x92, 0x9d, 0xb6,
	0xde, 0xca, 0xc3, 0x0f, 0x36, 0x22, 0x97, 0xa0, 0x35, 0x9b, 0x82, 0xff, 0x09, 0x63, 0x9e, 0x83,
	0x30, 0x7c, 0xc2, 0x41, 0x75, 0xd1, 0x39, 0xea, 0x9f, 0x86, 0x6f, 0xf6, 0xd1, 0x23, 0xb2, 0x8b,
	0x41, 0x3e, 0xac, 0x01, 0xe9, 0x35, 0x98, 0xff, 0x1e, 0x77, 0x1a, 0xce, 0xc8, 0xf9, 0xbb, 0xde,
	0xf9, 0xcd, 0xfe, 0x69, 0xd8, 0x23, 0x5c, 0xd6, 0x5d, 0x16, 0x60, 0x3e, 0xc3, 0x5c, 0x93, 0x6c,
	0xc6, 0x41, 0x18, 0x62, 0x99, 0xef, 0x58, 0xc1, 0x67, 0x8b, 0xf4, 0x6e, 0x63, 0x74, 0x69, 0x82,
	0xef, 0x08, 0xe3, 0xff, 0x39, 0xfc, 0x0b, 0x7c, 0x4b, 0xc8, 0x1c, 0x5c, 0xb1, 0x81, 0x2b, 0xd6,
	0xfe, 0x1a, 0xa9, 0x7f, 0xad, 0x2e, 0xf4, 0x4a, 0xe6, 0x90, 0xdc, 0x59, 0x25, 0xb7, 0x7f, 0x20,
	0xef, 0x0c, 0xa5, 0x8d, 0x23, 0x7e, 0x5b, 0x0f, 0x2f, 0xc6, 0x17, 0x87, 0x0d, 0x6f, 0xbb, 0xbd,
	0xf8, 0x55, 0x0d, 0xa0, 0xf8, 0xc5, 0x51, 0x80, 0xf0, 0x1b, 0xc2, 0xf7, 0xdc, 0xd3, 0xd0, 0x5a,
	0x7c, 0x83, 0x3b, 0x1b, 0x52, 0x9f, 0x1c, 0x37, 0xf3, 0x80, 0x1e, 0xa8, 0x5f, 0x6f, 0xc0, 0x8d,
	0x3e, 0x4a, 0x3e, 0xfe, 0xfa, 0xfa, 0xe7, 0xef, 0x89, 0x77, 0xe6, 0xe1, 0xe7, 0x5c, 0x5a, 0x40,
	0xa9, 0xe4, 0x97, 0xc5, 0x5e, 0x56, 0x72, 0x7f, 0xb3, 0xee, 0x41, 0xbd, 0x9c, 0x03, 0xb4, 0x44,
	0x68, 0x7c, 0xd2, 0x2c, 0x6a, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xf5, 0x35, 0xf0, 0x7f,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetricsService_serviceDesc.Streams[0], "/envoy.service.metrics.v3.MetricsService/StreamMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceStreamMetricsClient{stream}
	return x, nil
}

type MetricsService_StreamMetricsClient interface {
	Send(*StreamMetricsMessage) error
	CloseAndRecv() (*StreamMetricsResponse, error)
	grpc.ClientStream
}

type metricsServiceStreamMetricsClient struct {
	grpc.ClientStream
}

func (x *metricsServiceStreamMetricsClient) Send(m *StreamMetricsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsClient) CloseAndRecv() (*StreamMetricsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamMetricsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(MetricsService_StreamMetricsServer) error
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) StreamMetrics(srv MetricsService_StreamMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMetrics not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_StreamMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).StreamMetrics(&metricsServiceStreamMetricsServer{stream})
}

type MetricsService_StreamMetricsServer interface {
	SendAndClose(*StreamMetricsResponse) error
	Recv() (*StreamMetricsMessage, error)
	grpc.ServerStream
}

type metricsServiceStreamMetricsServer struct {
	grpc.ServerStream
}

func (x *metricsServiceStreamMetricsServer) SendAndClose(m *StreamMetricsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsServer) Recv() (*StreamMetricsMessage, error) {
	m := new(StreamMetricsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.metrics.v3.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMetrics",
			Handler:       _MetricsService_StreamMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/metrics/v3/metrics_service.proto",
}
