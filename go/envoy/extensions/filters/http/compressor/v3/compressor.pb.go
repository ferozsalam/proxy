// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/compressor/v3/compressor.proto

package envoy_extensions_filters_http_compressor_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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

// [#next-free-field: 6]
type Compressor struct {
	// Minimum response length, in bytes, which will trigger compression. The default value is 30.
	ContentLength *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	// Set of strings that allows specifying which mime-types yield compression; e.g.,
	// application/json, text/html, etc. When this field is not defined, compression will be applied
	// to the following mime-types: "application/javascript", "application/json",
	// "application/xhtml+xml", "image/svg+xml", "text/css", "text/html", "text/plain", "text/xml"
	// and their synonyms.
	ContentType []string `protobuf:"bytes,2,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, disables compression when the response contains an etag header. When it is false, the
	// filter will preserve weak etags and remove the ones that require strong validation.
	DisableOnEtagHeader bool `protobuf:"varint,3,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	// If true, removes accept-encoding from the request headers before dispatching it to the upstream
	// so that responses do not get compressed before reaching the filter.
	// .. attention:
	//
	//    To avoid interfering with other compression filters in the same chain use this option in
	//    the filter closest to the upstream.
	RemoveAcceptEncodingHeader bool `protobuf:"varint,4,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	// Runtime flag that controls whether the filter is enabled or not. If set to false, the
	// filter will operate as a pass-through filter. If not specified, defaults to enabled.
	RuntimeEnabled       *v3.RuntimeFeatureFlag `protobuf:"bytes,5,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Compressor) Reset()         { *m = Compressor{} }
func (m *Compressor) String() string { return proto.CompactTextString(m) }
func (*Compressor) ProtoMessage()    {}
func (*Compressor) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d6daf60a76e0a2b, []int{0}
}

func (m *Compressor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Compressor.Unmarshal(m, b)
}
func (m *Compressor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Compressor.Marshal(b, m, deterministic)
}
func (m *Compressor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Compressor.Merge(m, src)
}
func (m *Compressor) XXX_Size() int {
	return xxx_messageInfo_Compressor.Size(m)
}
func (m *Compressor) XXX_DiscardUnknown() {
	xxx_messageInfo_Compressor.DiscardUnknown(m)
}

var xxx_messageInfo_Compressor proto.InternalMessageInfo

func (m *Compressor) GetContentLength() *wrappers.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Compressor) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Compressor) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Compressor) GetRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.RemoveAcceptEncodingHeader
	}
	return false
}

func (m *Compressor) GetRuntimeEnabled() *v3.RuntimeFeatureFlag {
	if m != nil {
		return m.RuntimeEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*Compressor)(nil), "envoy.extensions.filters.http.compressor.v3.Compressor")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/compressor/v3/compressor.proto", fileDescriptor_8d6daf60a76e0a2b)
}

var fileDescriptor_8d6daf60a76e0a2b = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6b, 0xd4, 0x40,
	0x18, 0xc5, 0xd9, 0xad, 0x16, 0x9d, 0x6a, 0x2b, 0x11, 0x64, 0x59, 0x6c, 0xdd, 0x7a, 0x5a, 0x10,
	0x66, 0x70, 0x73, 0x51, 0xf1, 0xd2, 0x96, 0x2d, 0x0a, 0x82, 0x35, 0x68, 0xaf, 0x61, 0x36, 0xf9,
	0x36, 0x1b, 0xc8, 0x7e, 0xdf, 0x30, 0xf3, 0x25, 0x76, 0x6f, 0x1e, 0xfd, 0x1b, 0xfc, 0x53, 0xbc,
	0x0b, 0x5e, 0xfd, 0x73, 0xbc, 0x49, 0x66, 0xb2, 0xdd, 0xca, 0x9e, 0x7a, 0x9c, 0xbc, 0xf7, 0x7b,
	0xf3, 0x92, 0x17, 0xf1, 0x16, 0xb0, 0xa1, 0x95, 0x82, 0x2b, 0x06, 0x74, 0x25, 0xa1, 0x53, 0xf3,
	0xb2, 0x62, 0xb0, 0x4e, 0x2d, 0x98, 0x8d, 0xca, 0x68, 0x69, 0x2c, 0x38, 0x47, 0x56, 0x35, 0xf1,
	0x8d, 0x93, 0x34, 0x96, 0x98, 0xa2, 0x17, 0x9e, 0x96, 0x1b, 0x5a, 0x76, 0xb4, 0x6c, 0x69, 0x79,
	0xc3, 0xdf, 0xc4, 0xc3, 0x67, 0xe1, 0xaa, 0x8c, 0x70, 0x5e, 0x16, 0x2a, 0x23, 0x0b, 0x6d, 0xe6,
	0x4c, 0x3b, 0x08, 0x69, 0xc3, 0xa3, 0x82, 0xa8, 0xa8, 0x40, 0xf9, 0xd3, 0xac, 0x9e, 0xab, 0xaf,
	0x56, 0x1b, 0xd3, 0xa6, 0x05, 0xfd, 0xb0, 0xce, 0x8d, 0x56, 0x1a, 0x91, 0x58, 0xb3, 0xef, 0xea,
	0x58, 0x73, 0xbd, 0x96, 0x8f, 0xb7, 0xe4, 0x06, 0x6c, 0xdb, 0xaa, 0xc4, 0x22, 0x58, 0x9e, 0xff,
	0xed, 0x0b, 0x71, 0x76, 0x5d, 0x2a, 0x3a, 0x13, 0xfb, 0x19, 0x21, 0x03, 0x72, 0x5a, 0x01, 0x16,
	0xbc, 0x18, 0xf4, 0x46, 0xbd, 0xf1, 0xde, 0xe4, 0xa9, 0x0c, 0x4d, 0xe4, 0xba, 0x89, 0xfc, 0xf2,
	0x1e, 0x39, 0x9e, 0x5c, 0xea, 0xaa, 0x86, 0xe4, 0x61, 0xc7, 0x7c, 0xf0, 0x48, 0x74, 0x2c, 0x1e,
	0xac, 0x43, 0x78, 0x65, 0x60, 0xd0, 0x1f, 0xed, 0x8c, 0xef, 0x27, 0x7b, 0xdd, 0xb3, 0xcf, 0x2b,
	0x03, 0x51, 0x2c, 0x9e, 0xe4, 0xa5, 0xd3, 0xb3, 0x0a, 0x52, 0xc2, 0x14, 0x58, 0x17, 0xe9, 0x02,
	0x74, 0x0e, 0x76, 0xb0, 0x33, 0xea, 0x8d, 0xef, 0x25, 0x8f, 0x3b, 0xf5, 0x23, 0x4e, 0x59, 0x17,
	0xef, 0xbc, 0x14, 0x9d, 0x88, 0x43, 0x0b, 0x4b, 0x6a, 0x20, 0xd5, 0x59, 0x06, 0x86, 0x53, 0xc0,
	0x8c, 0xf2, 0x12, 0xaf, 0xd9, 0x3b, 0x9e, 0x1d, 0x06, 0xd3, 0x89, 0xf7, 0x4c, 0x3b, 0x4b, 0x17,
	0xf1, 0x49, 0x1c, 0xd8, 0x1a, 0xb9, 0x5c, 0x42, 0x0a, 0xd8, 0x5e, 0x90, 0x0f, 0xee, 0xfa, 0x17,
	0x1c, 0xcb, 0x30, 0x5c, 0xd8, 0x42, 0xb6, 0x5b, 0xc8, 0x26, 0x96, 0x49, 0x30, 0x9f, 0x83, 0xe6,
	0xda, 0xc2, 0x79, 0xa5, 0x8b, 0x64, 0xbf, 0x0b, 0x98, 0x06, 0xfe, 0xcd, 0xab, 0x1f, 0xbf, 0xbe,
	0x1f, 0xc5, 0xe2, 0xe5, 0x7f, 0x7c, 0x18, 0x7d, 0x7b, 0xf3, 0x89, 0xdc, 0x7c, 0xec, 0xd3, 0xcb,
	0x9f, 0xdf, 0x7e, 0xff, 0xd9, 0xed, 0x3f, 0xea, 0x8b, 0xd7, 0x25, 0x85, 0xfb, 0x8d, 0xa5, 0xab,
	0x95, 0xbc, 0xc5, 0x3f, 0x74, 0x7a, 0xb0, 0x09, 0xbc, 0x68, 0xb7, 0xb9, 0xe8, 0xcd, 0x76, 0xfd,
	0x48, 0xf1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x92, 0x65, 0x2a, 0xca, 0x02, 0x00, 0x00,
}
