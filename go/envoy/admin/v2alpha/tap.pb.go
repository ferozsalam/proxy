// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/tap.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	v2alpha "github.com/cilium/proxy/go/envoy/service/tap/v2alpha"
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

// The /tap admin request body that is used to configure an active tap session.
type TapRequest struct {
	// The opaque configuration ID used to match the configuration to a loaded extension.
	// A tap extension configures a similar opaque ID that is used to match.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	// The tap configuration to load.
	TapConfig            *v2alpha.TapConfig `protobuf:"bytes,2,opt,name=tap_config,json=tapConfig,proto3" json:"tap_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TapRequest) Reset()         { *m = TapRequest{} }
func (m *TapRequest) String() string { return proto.CompactTextString(m) }
func (*TapRequest) ProtoMessage()    {}
func (*TapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ef50f493e7b8aa, []int{0}
}

func (m *TapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapRequest.Unmarshal(m, b)
}
func (m *TapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapRequest.Marshal(b, m, deterministic)
}
func (m *TapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapRequest.Merge(m, src)
}
func (m *TapRequest) XXX_Size() int {
	return xxx_messageInfo_TapRequest.Size(m)
}
func (m *TapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TapRequest proto.InternalMessageInfo

func (m *TapRequest) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func (m *TapRequest) GetTapConfig() *v2alpha.TapConfig {
	if m != nil {
		return m.TapConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*TapRequest)(nil), "envoy.admin.v2alpha.TapRequest")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/tap.proto", fileDescriptor_69ef50f493e7b8aa) }

var fileDescriptor_69ef50f493e7b8aa = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0xe5, 0xe8, 0xfb, 0x4a, 0x62, 0x96, 0x2a, 0x0c, 0x54, 0x95, 0x2a, 0x05, 0x54, 0xa1,
	0x4e, 0xb6, 0x54, 0x06, 0xf6, 0x30, 0x31, 0x20, 0x55, 0x51, 0xf6, 0xea, 0x88, 0x0d, 0x58, 0x6a,
	0x7c, 0x26, 0xbe, 0x44, 0x74, 0x83, 0x95, 0x9f, 0xc4, 0x2f, 0x60, 0xe5, 0xef, 0x74, 0x42, 0x89,
	0x43, 0x27, 0x36, 0xdb, 0xcf, 0xbd, 0xd6, 0xf3, 0x1e, 0x5f, 0x68, 0xdb, 0xe1, 0x5e, 0x82, 0xaa,
	0x8d, 0x95, 0xdd, 0x1a, 0x76, 0xee, 0x19, 0x24, 0x81, 0x13, 0xae, 0x41, 0xc2, 0xf4, 0x6c, 0xc0,
	0x62, 0xc0, 0x62, 0xc4, 0xf3, 0xab, 0x90, 0xf1, 0xba, 0xe9, 0x4c, 0xa5, 0xfb, 0xe9, 0x63, 0xb2,
	0xc2, 0xba, 0x46, 0x1b, 0xc2, 0xf3, 0x45, 0xab, 0x1c, 0x48, 0xb0, 0x16, 0x09, 0xc8, 0xa0, 0xf5,
	0xd2, 0x13, 0x50, 0xeb, 0x47, 0x7c, 0xde, 0xc1, 0xce, 0x28, 0x20, 0x2d, 0x7f, 0x0f, 0x01, 0x5c,
	0xbe, 0x33, 0xce, 0x4b, 0x70, 0x85, 0x7e, 0x69, 0xb5, 0xa7, 0x74, 0xc9, 0x93, 0x0a, 0xed, 0xa3,
	0x79, 0xda, 0x1a, 0x35, 0x63, 0x19, 0x5b, 0x25, 0xf9, 0xc9, 0x21, 0xff, 0xd7, 0x44, 0x19, 0x2b,
	0xe2, 0x40, 0xee, 0x54, 0x7a, 0xcf, 0x39, 0x81, 0xdb, 0x86, 0xfb, 0x2c, 0xca, 0xd8, 0xea, 0x74,
	0xbd, 0x14, 0x41, 0x7f, 0x34, 0x15, 0x7d, 0xaf, 0xd1, 0x54, 0x94, 0xe0, 0x6e, 0x87, 0xd9, 0x3c,
	0x3e, 0xe4, 0xff, 0x3f, 0x58, 0x34, 0x65, 0x45, 0x42, 0xc7, 0xc7, 0x9b, 0xcf, 0xb7, 0xaf, 0xef,
	0x49, 0x34, 0x65, 0xfc, 0xc2, 0x60, 0xf8, 0xc6, 0x35, 0xf8, 0xba, 0x17, 0x7f, 0x2c, 0x24, 0x8f,
	0x4b, 0x70, 0x9b, 0x5e, 0x7d, 0xc3, 0x1e, 0x26, 0x43, 0x87, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0f, 0x93, 0xc7, 0xd5, 0x59, 0x01, 0x00, 0x00,
}
