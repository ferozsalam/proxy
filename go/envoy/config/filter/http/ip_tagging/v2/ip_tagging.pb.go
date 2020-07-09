// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package envoy_config_filter_http_ip_tagging_v2

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

// The type of requests the filter should apply to. The supported types
// are internal, external or both. The
// :ref:`x-forwarded-for<config_http_conn_man_headers_x-forwarded-for_internal_origin>` header is
// used to determine if a request is internal and will result in
// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>`
// being set. The filter defaults to both, and it will apply to all request types.
type IPTagging_RequestType int32

const (
	// Both external and internal requests will be tagged. This is the default value.
	IPTagging_BOTH IPTagging_RequestType = 0
	// Only internal requests will be tagged.
	IPTagging_INTERNAL IPTagging_RequestType = 1
	// Only external requests will be tagged.
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}

var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}

func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}

type IPTagging struct {
	// The type of request the filter should apply to.
	RequestType IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType" json:"request_type,omitempty"`
	// [#comment:TODO(ccaraman): Extend functionality to load IP tags from file system.
	// Tracked by issue https://github.com/envoyproxy/envoy/issues/2695]
	// The set of IP tags for the filter.
	IpTags               []*IPTagging_IPTag `protobuf:"bytes,4,rep,name=ip_tags,json=ipTags,proto3" json:"ip_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IPTagging) Reset()         { *m = IPTagging{} }
func (m *IPTagging) String() string { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()    {}
func (*IPTagging) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0}
}

func (m *IPTagging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging.Unmarshal(m, b)
}
func (m *IPTagging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging.Marshal(b, m, deterministic)
}
func (m *IPTagging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging.Merge(m, src)
}
func (m *IPTagging) XXX_Size() int {
	return xxx_messageInfo_IPTagging.Size(m)
}
func (m *IPTagging) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging proto.InternalMessageInfo

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

// Supplies the IP tag name and the IP address subnets.
type IPTagging_IPTag struct {
	// Specifies the IP tag name to apply.
	IpTagName string `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName,proto3" json:"ip_tag_name,omitempty"`
	// A list of IP address subnets that will be tagged with
	// ip_tag_name. Both IPv4 and IPv6 are supported.
	IpList               []*core.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList,proto3" json:"ip_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IPTagging_IPTag) Reset()         { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()    {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}

func (m *IPTagging_IPTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging_IPTag.Unmarshal(m, b)
}
func (m *IPTagging_IPTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging_IPTag.Marshal(b, m, deterministic)
}
func (m *IPTagging_IPTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging_IPTag.Merge(m, src)
}
func (m *IPTagging_IPTag) XXX_Size() int {
	return xxx_messageInfo_IPTagging_IPTag.Size(m)
}
func (m *IPTagging_IPTag) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging_IPTag.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging_IPTag proto.InternalMessageInfo

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*core.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
	proto.RegisterType((*IPTagging)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging.IPTag")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto", fileDescriptor_f4275d0b367744d2)
}

var fileDescriptor_f4275d0b367744d2 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xcf, 0x8a, 0x13, 0x31,
	0x18, 0x37, 0xe3, 0xb6, 0xb6, 0xe9, 0xb2, 0x94, 0x5c, 0x2c, 0x45, 0xd7, 0x65, 0x0f, 0xb2, 0x20,
	0x26, 0x30, 0x55, 0xf6, 0xe4, 0xc1, 0x91, 0x05, 0x0b, 0x4b, 0x2d, 0xc3, 0x1c, 0xc4, 0x83, 0x25,
	0xee, 0x64, 0xc7, 0x48, 0x9b, 0xc4, 0xe4, 0xeb, 0xb0, 0x73, 0x13, 0x7d, 0x00, 0xc1, 0x93, 0xcf,
	0xe2, 0x13, 0x78, 0xf5, 0x2d, 0x3c, 0x7b, 0xf4, 0x20, 0x92, 0x64, 0x56, 0x0b, 0x7b, 0xe9, 0xde,
	0xf2, 0xe5, 0xcb, 0xef, 0x5f, 0x7e, 0xf8, 0x58, 0xa8, 0x5a, 0x37, 0xec, 0x4c, 0xab, 0x73, 0x59,
	0xb1, 0x73, 0xb9, 0x04, 0x61, 0xd9, 0x5b, 0x00, 0xc3, 0xa4, 0x59, 0x00, 0xaf, 0x2a, 0xa9, 0x2a,
	0x56, 0xa7, 0x1b, 0x13, 0x35, 0x56, 0x83, 0x26, 0xf7, 0x03, 0x90, 0x46, 0x20, 0x8d, 0x40, 0xea,
	0x81, 0x74, 0xe3, 0x69, 0x9d, 0x8e, 0xef, 0x45, 0x01, 0x6e, 0xa4, 0xa7, 0x39, 0xd3, 0x56, 0x30,
	0x5e, 0x96, 0x56, 0x38, 0x17, 0x89, 0xc6, 0xfb, 0xeb, 0xd2, 0x70, 0xc6, 0x95, 0xd2, 0xc0, 0x41,
	0x6a, 0xe5, 0xd8, 0x4a, 0x56, 0x96, 0x83, 0x68, 0xf7, 0x77, 0xaf, 0xec, 0x1d, 0x70, 0x58, 0x5f,
	0xc2, 0x6f, 0xd7, 0x7c, 0x29, 0x4b, 0x0e, 0x82, 0x5d, 0x1e, 0xe2, 0xe2, 0xf0, 0x67, 0x82, 0xfb,
	0xd3, 0x79, 0x11, 0x9d, 0x90, 0x77, 0x78, 0xd7, 0x8a, 0xf7, 0x6b, 0xe1, 0x60, 0x01, 0x8d, 0x11,
	0x23, 0x74, 0x80, 0x8e, 0xf6, 0xd2, 0x27, 0x74, 0xbb, 0x14, 0xf4, 0x1f, 0x11, 0xcd, 0x23, 0x4b,
	0xd1, 0x18, 0x91, 0xf5, 0x7e, 0x67, 0x9d, 0x8f, 0x28, 0x19, 0xa2, 0x7c, 0x60, 0xff, 0x5f, 0x93,
	0x57, 0xf8, 0x56, 0x44, 0xbb, 0xd1, 0xce, 0xc1, 0xcd, 0xa3, 0x41, 0x7a, 0x7c, 0x7d, 0x99, 0x70,
	0x0a, 0x02, 0x5f, 0x50, 0xd2, 0x43, 0x79, 0x57, 0x9a, 0x82, 0x57, 0x6e, 0xfc, 0x1a, 0x77, 0xc2,
	0x8a, 0xec, 0xe3, 0x41, 0xc4, 0x2e, 0x14, 0x5f, 0xc5, 0x3c, 0xfd, 0xbc, 0x1f, 0x5e, 0xcd, 0xf8,
	0x4a, 0x90, 0xc7, 0xc1, 0xc4, 0x52, 0x3a, 0x18, 0x25, 0xc1, 0xc4, 0x9d, 0xd6, 0x04, 0x37, 0xd2,
	0x4b, 0xf9, 0x26, 0xe8, 0x33, 0x59, 0xda, 0x9c, 0xab, 0x4a, 0x78, 0xfe, 0x53, 0xe9, 0xe0, 0x70,
	0x82, 0x07, 0x1b, 0x09, 0x49, 0x0f, 0xef, 0x64, 0x2f, 0x8a, 0xe7, 0xc3, 0x1b, 0x64, 0x17, 0xf7,
	0xa6, 0xb3, 0xe2, 0x24, 0x9f, 0x3d, 0x3d, 0x1d, 0x22, 0x3f, 0x9d, 0xbc, 0x6c, 0xa7, 0x24, 0xfb,
	0x84, 0x7e, 0x7d, 0xfd, 0xf3, 0xb9, 0xf3, 0x90, 0x3c, 0x88, 0x12, 0xe2, 0x02, 0x84, 0x72, 0xbe,
	0xab, 0x36, 0xab, 0xbb, 0x1a, 0x76, 0xf2, 0xed, 0xc3, 0xf7, 0x1f, 0xdd, 0x64, 0x88, 0xf0, 0x23,
	0xa9, 0xa3, 0x35, 0x63, 0xf5, 0x45, 0xb3, 0xe5, 0x57, 0x65, 0x7b, 0x53, 0xd3, 0xfe, 0xd5, 0xdc,
	0xd7, 0x3d, 0x47, 0x6f, 0xba, 0xa1, 0xf7, 0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x6b,
	0x45, 0x6b, 0xd3, 0x02, 0x00, 0x00,
}
