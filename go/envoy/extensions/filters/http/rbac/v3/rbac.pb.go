// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/rbac/v3/rbac.proto

package envoy_extensions_filters_http_rbac_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/rbac/v3"
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

// RBAC filter config.
type RBAC struct {
	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v3.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter (i.e., returning a 403)
	// but will emit stats and logs and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules          *v3.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77e76eac62eed05, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v3.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v3.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	// Override the global configuration of the filter with this new config.
	// If absent, the global RBAC policy will be disabled for this route.
	Rbac                 *RBAC    `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACPerRoute) Reset()         { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()    {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77e76eac62eed05, []int{1}
}

func (m *RBACPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBACPerRoute.Unmarshal(m, b)
}
func (m *RBACPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBACPerRoute.Marshal(b, m, deterministic)
}
func (m *RBACPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACPerRoute.Merge(m, src)
}
func (m *RBACPerRoute) XXX_Size() int {
	return xxx_messageInfo_RBACPerRoute.Size(m)
}
func (m *RBACPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RBACPerRoute proto.InternalMessageInfo

func (m *RBACPerRoute) GetRbac() *RBAC {
	if m != nil {
		return m.Rbac
	}
	return nil
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.extensions.filters.http.rbac.v3.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.extensions.filters.http.rbac.v3.RBACPerRoute")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/rbac/v3/rbac.proto", fileDescriptor_b77e76eac62eed05)
}

var fileDescriptor_b77e76eac62eed05 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x49, 0x99, 0x43, 0xb3, 0x1d, 0x46, 0x2f, 0xca, 0xc0, 0x7f, 0x83, 0xa1, 0xa0, 0x26,
	0x63, 0xf5, 0x34, 0x10, 0xb1, 0xde, 0x76, 0x2a, 0xfd, 0x02, 0x92, 0xb6, 0xd9, 0x16, 0x28, 0x49,
	0x49, 0xd2, 0xba, 0xdd, 0x3c, 0x0a, 0x7e, 0x03, 0xef, 0x7e, 0x09, 0xef, 0x82, 0x57, 0xbf, 0x91,
	0x24, 0x69, 0x51, 0x51, 0xb0, 0xa7, 0x06, 0xde, 0xdf, 0xef, 0x7d, 0x1e, 0xde, 0xc2, 0x09, 0xe5,
	0x95, 0xd8, 0x60, 0xba, 0xd6, 0x94, 0x2b, 0x26, 0xb8, 0xc2, 0x0b, 0x96, 0x6b, 0x2a, 0x15, 0x5e,
	0x69, 0x5d, 0x60, 0x99, 0x90, 0x14, 0x57, 0x81, 0xfd, 0xa2, 0x42, 0x0a, 0x2d, 0xfc, 0xb1, 0x35,
	0xd0, 0x97, 0x81, 0x6a, 0x03, 0x19, 0x03, 0x59, 0xb2, 0x0a, 0x86, 0x87, 0x6e, 0x71, 0x2a, 0xf8,
	0x82, 0x2d, 0xff, 0xd8, 0x33, 0xdc, 0x2f, 0xb3, 0x82, 0x60, 0xc2, 0xb9, 0xd0, 0x44, 0xdb, 0x64,
	0xa5, 0x89, 0x2e, 0x55, 0x3d, 0x3e, 0xfe, 0x35, 0xae, 0xa8, 0x34, 0x79, 0x8c, 0x2f, 0x6b, 0x64,
	0xb7, 0x22, 0x39, 0xcb, 0x88, 0xa6, 0xb8, 0x79, 0xb8, 0xc1, 0xe8, 0x05, 0xc0, 0x4e, 0x1c, 0xde,
	0xdc, 0xfa, 0x13, 0xb8, 0x25, 0xcb, 0x9c, 0xaa, 0x3d, 0x70, 0x04, 0x4e, 0x7b, 0xd3, 0x21, 0x72,
	0xdd, 0x5d, 0xa9, 0xa6, 0x2a, 0x32, 0x68, 0xec, 0x40, 0xff, 0x0a, 0xf6, 0xd5, 0x8a, 0x64, 0xe2,
	0xfe, 0xce, 0x89, 0xde, 0xbf, 0x62, 0xcf, 0xf1, 0xb1, 0xc1, 0x67, 0xe7, 0xcf, 0x6f, 0x8f, 0x07,
	0x27, 0x70, 0xfc, 0x03, 0x77, 0xf7, 0xf9, 0x7e, 0x9e, 0xa9, 0x55, 0x47, 0x4f, 0x00, 0xf6, 0xcd,
	0x23, 0xa2, 0x32, 0x16, 0xa5, 0xa6, 0xfe, 0x35, 0xec, 0x18, 0xa0, 0x4e, 0x3d, 0x43, 0xad, 0x4e,
	0xed, 0x6a, 0x58, 0x71, 0x76, 0x69, 0xf2, 0x31, 0xbc, 0x68, 0x95, 0xdf, 0xc4, 0xce, 0x3b, 0xdb,
	0x60, 0xe0, 0x85, 0xf3, 0xd7, 0x87, 0xf7, 0x8f, 0xae, 0x37, 0xf0, 0x60, 0xc0, 0x84, 0x8b, 0x2e,
	0xa4, 0x58, 0x6f, 0xda, 0xb5, 0x08, 0x77, 0xe2, 0x84, 0xa4, 0x91, 0xb9, 0x7f, 0x04, 0x92, 0xae,
	0xfd, 0x11, 0xc1, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xb6, 0xb2, 0xb6, 0x5f, 0x02, 0x00,
	0x00,
}
