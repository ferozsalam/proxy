// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/tap/v4alpha/tap.proto

package envoy_extensions_filters_http_tap_v4alpha

import (
	fmt "fmt"
	v4alpha "github.com/cilium/proxy/go/envoy/extensions/common/tap/v4alpha"
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

// Top level configuration for the tap filter.
type Tap struct {
	// Common configuration for the HTTP tap filter.
	CommonConfig         *v4alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_041331360e2659a1, []int{0}
}

func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v4alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.extensions.filters.http.tap.v4alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/tap/v4alpha/tap.proto", fileDescriptor_041331360e2659a1)
}

var fileDescriptor_041331360e2659a1 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2f, 0x49, 0x2c, 0xd0, 0x2f, 0x33,
	0x49, 0xcc, 0x29, 0xc8, 0x48, 0x04, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x34, 0xc1,
	0x9a, 0xf4, 0x10, 0x9a, 0xf4, 0xa0, 0x9a, 0xf4, 0x40, 0x9a, 0xf4, 0x40, 0x0a, 0xa1, 0x9a, 0xa4,
	0x0c, 0x30, 0xcc, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x43, 0x31, 0x19, 0x22, 0x04, 0x31, 0x5c,
	0x4a, 0xb6, 0x34, 0xa5, 0x20, 0x51, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x04, 0xac, 0xa3,
	0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x18, 0x2a, 0xad, 0x88, 0x21, 0x5d, 0x96, 0x5a, 0x04, 0x32, 0x39,
	0x33, 0x2f, 0x1d, 0xaa, 0x44, 0xbc, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6,
	0x80, 0x48, 0x28, 0x2d, 0x64, 0xe4, 0x62, 0x0e, 0x49, 0x2c, 0x10, 0xca, 0xe4, 0xe2, 0x85, 0x58,
	0x19, 0x9f, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x64, 0xa5,
	0x87, 0xe1, 0x2f, 0xa8, 0xcb, 0x90, 0x7c, 0xa4, 0xe7, 0x0c, 0x16, 0x72, 0x85, 0xa9, 0x71, 0x06,
	0x9b, 0xe0, 0xc4, 0xf1, 0xcb, 0x89, 0xb5, 0x8b, 0x91, 0x49, 0x80, 0x31, 0x88, 0x07, 0xa2, 0x07,
	0x22, 0x6e, 0xa5, 0x3f, 0xeb, 0x68, 0x87, 0x9c, 0x16, 0x97, 0x06, 0x11, 0x21, 0x66, 0xac, 0x17,
	0x92, 0x58, 0xe0, 0xe4, 0xbb, 0xab, 0xe1, 0xc4, 0x45, 0x36, 0x26, 0x01, 0x66, 0x2e, 0xf3, 0xcc,
	0x7c, 0x88, 0x83, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x88, 0x0e, 0x73, 0x27, 0x8e, 0x90, 0xc4,
	0x82, 0x00, 0x90, 0x87, 0x03, 0x18, 0x93, 0xd8, 0xc0, 0x3e, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x89, 0x89, 0xf7, 0x2c, 0xe8, 0x01, 0x00, 0x00,
}
