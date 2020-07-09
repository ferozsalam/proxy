// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2alpha/xray.proto

package envoy_config_trace_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type XRayConfig struct {
	// The UDP endpoint of the X-Ray Daemon where the spans will be sent.
	// If this value is not set, the default value of 127.0.0.1:2000 will be used.
	DaemonEndpoint *core.SocketAddress `protobuf:"bytes,1,opt,name=daemon_endpoint,json=daemonEndpoint,proto3" json:"daemon_endpoint,omitempty"`
	// The name of the X-Ray segment.
	SegmentName string `protobuf:"bytes,2,opt,name=segment_name,json=segmentName,proto3" json:"segment_name,omitempty"`
	// The location of a local custom sampling rules JSON file.
	// For an example of the sampling rules see:
	// `X-Ray SDK documentation
	// <https://docs.aws.amazon.com/xray/latest/devguide/xray-sdk-go-configuration.html#xray-sdk-go-configuration-sampling>`_
	SamplingRuleManifest *core.DataSource `protobuf:"bytes,3,opt,name=sampling_rule_manifest,json=samplingRuleManifest,proto3" json:"sampling_rule_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *XRayConfig) Reset()         { *m = XRayConfig{} }
func (m *XRayConfig) String() string { return proto.CompactTextString(m) }
func (*XRayConfig) ProtoMessage()    {}
func (*XRayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c272343b96f21fb4, []int{0}
}

func (m *XRayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRayConfig.Unmarshal(m, b)
}
func (m *XRayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRayConfig.Marshal(b, m, deterministic)
}
func (m *XRayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRayConfig.Merge(m, src)
}
func (m *XRayConfig) XXX_Size() int {
	return xxx_messageInfo_XRayConfig.Size(m)
}
func (m *XRayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRayConfig proto.InternalMessageInfo

func (m *XRayConfig) GetDaemonEndpoint() *core.SocketAddress {
	if m != nil {
		return m.DaemonEndpoint
	}
	return nil
}

func (m *XRayConfig) GetSegmentName() string {
	if m != nil {
		return m.SegmentName
	}
	return ""
}

func (m *XRayConfig) GetSamplingRuleManifest() *core.DataSource {
	if m != nil {
		return m.SamplingRuleManifest
	}
	return nil
}

func init() {
	proto.RegisterType((*XRayConfig)(nil), "envoy.config.trace.v2alpha.XRayConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v2alpha/xray.proto", fileDescriptor_c272343b96f21fb4)
}

var fileDescriptor_c272343b96f21fb4 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x99, 0xde, 0x4b, 0xa5, 0x53, 0xd1, 0x12, 0x44, 0x4b, 0xb1, 0x58, 0x04, 0xa1, 0xb8,
	0x98, 0x81, 0xf8, 0x00, 0x62, 0xd4, 0x85, 0x0b, 0xa5, 0xa4, 0x9b, 0xee, 0xc2, 0x69, 0x72, 0x5a,
	0x07, 0x93, 0x99, 0x61, 0x66, 0x12, 0x9a, 0x9d, 0xcf, 0xe5, 0x13, 0xb8, 0x75, 0xe7, 0xb3, 0xb8,
	0x92, 0x66, 0xd2, 0x95, 0xba, 0x1b, 0x38, 0xdf, 0x77, 0xfe, 0x9f, 0x33, 0xf4, 0x02, 0x65, 0xa5,
	0x6a, 0x9e, 0x2a, 0xb9, 0x12, 0x6b, 0xee, 0x0c, 0xa4, 0xc8, 0xab, 0x10, 0x72, 0xfd, 0x0c, 0x7c,
	0x63, 0xa0, 0x66, 0xda, 0x28, 0xa7, 0x82, 0x51, 0x83, 0x31, 0x8f, 0xb1, 0x06, 0x63, 0x2d, 0x36,
	0x3a, 0xf3, 0x2b, 0x40, 0x0b, 0x5e, 0x85, 0x3c, 0x55, 0x06, 0x39, 0x64, 0x99, 0x41, 0x6b, 0xbd,
	0x3c, 0x3a, 0xfd, 0x09, 0x2c, 0xc1, 0x62, 0x3b, 0x1d, 0x97, 0x99, 0x06, 0x0e, 0x52, 0x2a, 0x07,
	0x4e, 0x28, 0x69, 0xb9, 0x75, 0xe0, 0xca, 0x9d, 0x7c, 0x52, 0x41, 0x2e, 0x32, 0x70, 0xc8, 0x77,
	0x0f, 0x3f, 0x38, 0xff, 0x24, 0x94, 0x2e, 0x62, 0xa8, 0x6f, 0x9b, 0x4e, 0xc1, 0x03, 0x3d, 0xcc,
	0x00, 0x0b, 0x25, 0x13, 0x94, 0x99, 0x56, 0x42, 0xba, 0x21, 0x99, 0x90, 0x69, 0x3f, 0x9c, 0x30,
	0xdf, 0x1d, 0xb4, 0x60, 0x55, 0xc8, 0xb6, 0xf1, 0x6c, 0xae, 0xd2, 0x17, 0x74, 0x37, 0xbe, 0x65,
	0x7c, 0xe0, 0xc5, 0xfb, 0xd6, 0x0b, 0x2e, 0xe9, 0xbe, 0xc5, 0x75, 0x81, 0xd2, 0x25, 0x12, 0x0a,
	0x1c, 0x76, 0x26, 0x64, 0xda, 0x8b, 0xf6, 0xbe, 0xa2, 0xff, 0xa6, 0x33, 0x20, 0x71, 0xbf, 0x1d,
	0x3e, 0x41, 0x81, 0xc1, 0x9c, 0x1e, 0x5b, 0x28, 0x74, 0x2e, 0xe4, 0x3a, 0x31, 0x65, 0x8e, 0x49,
	0x01, 0x52, 0xac, 0xd0, 0xba, 0xe1, 0xbf, 0x26, 0x7d, 0xfc, 0x4b, 0xfa, 0x1d, 0x38, 0x98, 0xab,
	0xd2, 0xa4, 0x18, 0x1f, 0xed, 0xe4, 0xb8, 0xcc, 0xf1, 0xb1, 0x55, 0xa3, 0xeb, 0xb7, 0xd7, 0xf7,
	0x8f, 0x6e, 0x67, 0x40, 0xe8, 0x54, 0x28, 0xbf, 0x40, 0x1b, 0xb5, 0xa9, 0xd9, 0xdf, 0xbf, 0x10,
	0xf5, 0x16, 0x06, 0xea, 0xd9, 0xf6, 0x32, 0x33, 0xb2, 0xec, 0x36, 0x27, 0xba, 0xfa, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xbb, 0x5a, 0x39, 0xd6, 0xde, 0x01, 0x00, 0x00,
}
