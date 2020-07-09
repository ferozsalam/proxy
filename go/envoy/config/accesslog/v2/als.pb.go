// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/accesslog/v2/als.proto

package envoy_config_accesslog_v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration for the built-in *envoy.access_loggers.http_grpc*
// :ref:`AccessLog <envoy_api_msg_config.filter.accesslog.v2.AccessLog>`. This configuration will
// populate :ref:`StreamAccessLogsMessage.http_logs
// <envoy_api_field_service.accesslog.v2.StreamAccessLogsMessage.http_logs>`.
// [#extension: envoy.access_loggers.http_grpc]
type HttpGrpcAccessLogConfig struct {
	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// Additional request headers to log in :ref:`HTTPRequestProperties.request_headers
	// <envoy_api_field_data.accesslog.v2.HTTPRequestProperties.request_headers>`.
	AdditionalRequestHeadersToLog []string `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	// Additional response headers to log in :ref:`HTTPResponseProperties.response_headers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_headers>`.
	AdditionalResponseHeadersToLog []string `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	// Additional response trailers to log in :ref:`HTTPResponseProperties.response_trailers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_trailers>`.
	AdditionalResponseTrailersToLog []string `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{0}
}

func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Size(m)
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// Configuration for the built-in *envoy.access_loggers.tcp_grpc* type. This configuration will
// populate *StreamAccessLogsMessage.tcp_logs*.
// [#extension: envoy.access_loggers.tcp_grpc]
type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{1}
}

func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Size(m)
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

// Common configuration for gRPC access logs.
// [#next-free-field: 6]
type CommonGrpcAccessLogConfig struct {
	// The friendly name of the access log to be returned in :ref:`StreamAccessLogsMessage.Identifier
	// <envoy_api_msg_service.accesslog.v2.StreamAccessLogsMessage.Identifier>`. This allows the
	// access log server to differentiate between different access logs coming from the same Envoy.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The gRPC service for the access log service.
	GrpcService *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// Interval for flushing access logs to the gRPC stream. Logger will flush requests every time
	// this interval is elapsed, or when batch size limit is hit, whichever comes first. Defaults to
	// 1 second.
	BufferFlushInterval *duration.Duration `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	// Soft size limit in bytes for access log entries buffer. Logger will buffer requests until
	// this limit it hit, or every time flush interval is elapsed, whichever comes first. Setting it
	// to zero effectively disables the batching. Defaults to 16384.
	BufferSizeBytes *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	// Additional filter state objects to log in :ref:`filter_state_objects
	// <envoy_api_field_data.accesslog.v2.AccessLogCommon.filter_state_objects>`.
	// Logger will call `FilterState::Object::serializeAsProto` to serialize the filter state object.
	FilterStateObjectsToLog []string `protobuf:"bytes,5,rep,name=filter_state_objects_to_log,json=filterStateObjectsToLog,proto3" json:"filter_state_objects_to_log,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{2}
}

func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Size(m)
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *duration.Duration {
	if m != nil {
		return m.BufferFlushInterval
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.BufferSizeBytes
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetFilterStateObjectsToLog() []string {
	if m != nil {
		return m.FilterStateObjectsToLog
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v2/als.proto", fileDescriptor_e7b431652a309a2e)
}

var fileDescriptor_e7b431652a309a2e = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0x4f, 0x4f, 0xd4, 0x4e,
	0x18, 0xc7, 0x7f, 0x2d, 0xff, 0x07, 0x7e, 0x11, 0x6b, 0x94, 0x05, 0x05, 0x71, 0x31, 0x01, 0x3d,
	0xb4, 0xc9, 0xe2, 0xd1, 0x0b, 0xc5, 0xe8, 0xa2, 0x44, 0x49, 0x59, 0xf5, 0xd8, 0xcc, 0xb6, 0x4f,
	0x87, 0x31, 0xb3, 0x9d, 0x3a, 0x33, 0xad, 0x2c, 0x89, 0x89, 0x67, 0x2f, 0x5e, 0x7d, 0x0d, 0xbe,
	0x04, 0x5f, 0x81, 0x57, 0xdf, 0x82, 0x77, 0x2f, 0x1e, 0x39, 0x18, 0x33, 0x33, 0x5d, 0x58, 0x5c,
	0xf6, 0xec, 0x6d, 0xb7, 0xcf, 0xf7, 0xfb, 0x99, 0xe7, 0x2f, 0xda, 0x80, 0xbc, 0xe2, 0xfd, 0x20,
	0xe1, 0x79, 0x46, 0x49, 0x80, 0x93, 0x04, 0xa4, 0x64, 0x9c, 0x04, 0x55, 0x2b, 0xc0, 0x4c, 0xfa,
	0x85, 0xe0, 0x8a, 0x7b, 0xcb, 0x46, 0xe4, 0x5b, 0x91, 0x7f, 0x26, 0xf2, 0xab, 0xd6, 0xca, 0x5d,
	0xeb, 0xc7, 0x05, 0xd5, 0x96, 0x84, 0x0b, 0x08, 0x88, 0x28, 0x92, 0x58, 0x82, 0xa8, 0x68, 0x02,
	0x16, 0xb0, 0xb2, 0x46, 0x38, 0x27, 0x0c, 0x02, 0xf3, 0xaf, 0x5b, 0x66, 0x41, 0x5a, 0x0a, 0xac,
	0x28, 0xcf, 0xc7, 0xc5, 0xdf, 0x09, 0x5c, 0x14, 0x20, 0xe4, 0x20, 0x5e, 0xa6, 0x05, 0x0e, 0x70,
	0x9e, 0x73, 0x65, 0x6c, 0x32, 0xe8, 0x51, 0x22, 0xb0, 0x1a, 0xf0, 0x57, 0x47, 0xe2, 0x52, 0x61,
	0x55, 0x0e, 0xec, 0x4b, 0x15, 0x66, 0x34, 0xc5, 0x0a, 0x82, 0xc1, 0x0f, 0x1b, 0x68, 0xfe, 0x70,
	0xd1, 0x52, 0x5b, 0xa9, 0xe2, 0x89, 0x28, 0x92, 0x1d, 0x53, 0xd6, 0x3e, 0x27, 0xbb, 0xa6, 0x4c,
	0x2f, 0x41, 0xff, 0x27, 0xbc, 0xd7, 0xe3, 0x79, 0x6c, 0xeb, 0x6e, 0x38, 0xeb, 0xce, 0xd6, 0x7c,
	0xeb, 0x81, 0x3f, 0xb6, 0x19, 0xfe, 0xae, 0xd1, 0x5f, 0x02, 0x0b, 0x67, 0x4f, 0xc3, 0xa9, 0x8f,
	0x8e, 0xbb, 0xe8, 0x44, 0x0b, 0x16, 0x5a, 0x3f, 0xd2, 0x46, 0x77, 0x70, 0x9a, 0x52, 0x9d, 0x33,
	0x66, 0xb1, 0x80, 0xb7, 0x25, 0x48, 0x15, 0x1f, 0x01, 0x4e, 0x41, 0xc8, 0x58, 0xf1, 0x98, 0x71,
	0xd2, 0x70, 0xd7, 0x27, 0xb6, 0xe6, 0xa2, 0xd5, 0x73, 0x61, 0x64, 0x75, 0x6d, 0x2b, 0xeb, 0xf0,
	0x7d, 0x4e, 0xbc, 0xa7, 0xa8, 0x79, 0x81, 0x24, 0x0b, 0x9e, 0x4b, 0xf8, 0x1b, 0x35, 0x61, 0x50,
	0x6b, 0xc3, 0x28, 0x2b, 0xbc, 0xc0, 0xda, 0x47, 0x1b, 0x97, 0xb1, 0x94, 0xc0, 0x94, 0x0d, 0xc1,
	0x26, 0x0d, 0xec, 0xf6, 0x28, 0xac, 0x53, 0x0b, 0x0d, 0xad, 0xf9, 0x1e, 0xdd, 0xe8, 0x24, 0xff,
	0xac, 0xc5, 0xcd, 0x9f, 0x2e, 0x5a, 0x1e, 0xeb, 0xf2, 0x9a, 0x68, 0x96, 0x71, 0x12, 0xe7, 0xb8,
	0x07, 0xe6, 0xf5, 0xb9, 0x70, 0xe6, 0x34, 0x9c, 0x14, 0xee, 0xba, 0x13, 0xcd, 0x30, 0x4e, 0x9e,
	0xe3, 0x1e, 0x78, 0xcf, 0xd0, 0xc2, 0xf0, 0x4e, 0x37, 0x5c, 0x93, 0xe5, 0x5a, 0x9d, 0x25, 0x2e,
	0xa8, 0x4e, 0x4c, 0xaf, 0xbe, 0xaf, 0x5f, 0x38, 0xb4, 0xaa, 0xa1, 0x7c, 0xe6, 0xc9, 0xf9, 0x67,
	0xef, 0x35, 0xba, 0xde, 0x2d, 0xb3, 0x0c, 0x44, 0x9c, 0xb1, 0x52, 0x1e, 0xc5, 0x34, 0x57, 0x20,
	0x2a, 0xcc, 0x1a, 0x13, 0x86, 0xba, 0xec, 0xdb, 0x53, 0xf0, 0x07, 0xa7, 0xe0, 0x3f, 0xaa, 0x4f,
	0xc5, 0x00, 0xbf, 0x38, 0xee, 0xfd, 0xff, 0xa2, 0x6b, 0x96, 0xf0, 0x58, 0x03, 0xf6, 0x6a, 0xbf,
	0xd7, 0x46, 0x57, 0x6b, 0xb0, 0xa4, 0x27, 0x10, 0x77, 0xfb, 0x0a, 0x64, 0x63, 0xd2, 0x40, 0x6f,
	0x8d, 0x40, 0x5f, 0xee, 0xe5, 0x6a, 0xbb, 0xf5, 0x0a, 0xb3, 0x12, 0xa2, 0x2b, 0xd6, 0x76, 0x48,
	0x4f, 0x20, 0xd4, 0x26, 0xef, 0x21, 0xba, 0x99, 0x51, 0xa6, 0x34, 0x49, 0x61, 0x05, 0x31, 0xef,
	0xbe, 0x81, 0x44, 0x9d, 0x8d, 0x7d, 0xca, 0x8c, 0x7d, 0xc9, 0x4a, 0x0e, 0xb5, 0xe2, 0x85, 0x15,
	0x98, 0x71, 0x87, 0xec, 0xd7, 0xe7, 0xdf, 0x9f, 0xa6, 0xee, 0x79, 0x9b, 0xb6, 0x3d, 0x70, 0xac,
	0x20, 0x97, 0xfa, 0x26, 0xeb, 0x41, 0x6a, 0x06, 0xd1, 0xa7, 0xad, 0xbb, 0xe2, 0x57, 0xdb, 0x5f,
	0x3f, 0x7c, 0xfb, 0x3e, 0xed, 0x2e, 0x3a, 0x68, 0x93, 0x72, 0xdb, 0xd2, 0x42, 0xf0, 0xe3, 0xfe,
	0xf8, 0x1d, 0x08, 0x67, 0x77, 0x98, 0x3c, 0xd0, 0x95, 0x1c, 0x38, 0xdd, 0x69, 0x53, 0xd2, 0xf6,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x55, 0x8b, 0x5a, 0xc8, 0x04, 0x00, 0x00,
}
