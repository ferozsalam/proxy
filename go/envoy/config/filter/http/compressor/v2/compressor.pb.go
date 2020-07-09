// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/compressor/v2/compressor.proto

package envoy_config_filter_http_compressor_v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
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
	RuntimeEnabled       *core.RuntimeFeatureFlag `protobuf:"bytes,5,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Compressor) Reset()         { *m = Compressor{} }
func (m *Compressor) String() string { return proto.CompactTextString(m) }
func (*Compressor) ProtoMessage()    {}
func (*Compressor) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef4714f5f18ab25, []int{0}
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

func (m *Compressor) GetRuntimeEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.RuntimeEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*Compressor)(nil), "envoy.config.filter.http.compressor.v2.Compressor")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/compressor/v2/compressor.proto", fileDescriptor_8ef4714f5f18ab25)
}

var fileDescriptor_8ef4714f5f18ab25 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0xc9, 0xd6, 0x16, 0x9d, 0x6a, 0x2b, 0x11, 0x64, 0x59, 0x5a, 0x59, 0x05, 0x65, 0x41,
	0x9c, 0x81, 0xac, 0xe0, 0xb9, 0x2d, 0x5b, 0x14, 0x44, 0x4b, 0x50, 0xaf, 0xe1, 0x6d, 0xf2, 0x36,
	0x3b, 0x90, 0x9d, 0x37, 0xcc, 0xbc, 0xc4, 0xee, 0xcd, 0x83, 0x77, 0xaf, 0x7e, 0x16, 0xf1, 0x03,
	0x78, 0xf5, 0xab, 0xf8, 0x01, 0x44, 0x92, 0xc9, 0xd6, 0xc2, 0x5e, 0x7a, 0x7c, 0xf9, 0xff, 0x7f,
	0xff, 0xf7, 0x92, 0x7f, 0xc4, 0x2b, 0x34, 0x0d, 0xad, 0x55, 0x4e, 0x66, 0xa1, 0x4b, 0xb5, 0xd0,
	0x15, 0xa3, 0x53, 0x4b, 0x66, 0xab, 0x72, 0x5a, 0x59, 0x87, 0xde, 0x93, 0x53, 0x4d, 0x72, 0x6d,
	0x92, 0xd6, 0x11, 0x53, 0xfc, 0xac, 0x03, 0x65, 0x00, 0x65, 0x00, 0x65, 0x0b, 0xca, 0x6b, 0xd6,
	0x26, 0x19, 0x1d, 0x85, 0x05, 0x60, 0x75, 0x88, 0x71, 0xa8, 0xe6, 0xe0, 0x31, 0xa4, 0x8c, 0x1e,
	0x95, 0x44, 0x65, 0x85, 0xaa, 0x9b, 0xe6, 0xf5, 0x42, 0x7d, 0x76, 0x60, 0x2d, 0x3a, 0xbf, 0xd1,
	0xeb, 0xc2, 0x82, 0x02, 0x63, 0x88, 0x81, 0x35, 0x19, 0xaf, 0x56, 0xba, 0x74, 0xc0, 0x1b, 0xfe,
	0x78, 0x4b, 0xf7, 0x0c, 0x5c, 0xf7, 0xf8, 0x93, 0x9f, 0x03, 0x21, 0xce, 0xae, 0xce, 0x89, 0xcf,
	0xc4, 0x41, 0x4e, 0x86, 0xd1, 0x70, 0x56, 0xa1, 0x29, 0x79, 0x39, 0x8c, 0xc6, 0xd1, 0x64, 0x3f,
	0x39, 0x92, 0xe1, 0x0c, 0xb9, 0x39, 0x43, 0x7e, 0x7c, 0x63, 0x78, 0x9a, 0x7c, 0x82, 0xaa, 0xc6,
	0xf4, 0x5e, 0xcf, 0xbc, 0xed, 0x90, 0xf8, 0xb1, 0xb8, 0xbb, 0x09, 0xe1, 0xb5, 0xc5, 0xe1, 0x60,
	0xbc, 0x33, 0xb9, 0x93, 0xee, 0xf7, 0xcf, 0x3e, 0xac, 0x2d, 0xc6, 0x53, 0xf1, 0xb0, 0xd0, 0x1e,
	0xe6, 0x15, 0x66, 0x64, 0x32, 0x64, 0x28, 0xb3, 0x25, 0x42, 0x81, 0x6e, 0xb8, 0x33, 0x8e, 0x26,
	0xb7, 0xd3, 0x07, 0xbd, 0xfa, 0xde, 0xcc, 0x18, 0xca, 0xd7, 0x9d, 0x14, 0x9f, 0x88, 0x63, 0x87,
	0x2b, 0x6a, 0x30, 0x83, 0x3c, 0x47, 0xcb, 0x19, 0x9a, 0x9c, 0x0a, 0x6d, 0xae, 0xd8, 0x5b, 0x1d,
	0x3b, 0x0a, 0xa6, 0x93, 0xce, 0x33, 0xeb, 0x2d, 0x7d, 0xc4, 0x3b, 0x71, 0xe8, 0x6a, 0xc3, 0x7a,
	0x85, 0x19, 0x9a, 0x76, 0x41, 0x31, 0xdc, 0xed, 0x5e, 0xf0, 0xa9, 0x0c, 0x6d, 0x81, 0xd5, 0xb2,
	0x49, 0x64, 0xdb, 0x82, 0x4c, 0x83, 0xf3, 0x1c, 0x81, 0x6b, 0x87, 0xe7, 0x15, 0x94, 0xe9, 0x41,
	0x4f, 0xcf, 0x02, 0x7c, 0xfa, 0x35, 0xfa, 0xf3, 0xfd, 0xef, 0xb7, 0xdd, 0x17, 0xf1, 0xf3, 0x80,
	0xe3, 0x25, 0xa3, 0xf1, 0xed, 0x67, 0xee, 0x0b, 0xf7, 0xdb, 0x8d, 0x4f, 0x7f, 0x7c, 0xf9, 0xf5,
	0x7b, 0x6f, 0x70, 0x3f, 0x12, 0x2f, 0x35, 0x85, 0xb5, 0xd6, 0xd1, 0xe5, 0x5a, 0xde, 0xec, 0x7f,
	0x39, 0x3d, 0xfc, 0xdf, 0xd7, 0x45, 0xdb, 0xc6, 0x45, 0x34, 0xdf, 0xeb, 0x6a, 0x99, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0x31, 0xb8, 0xb8, 0xac, 0x02, 0x00, 0x00,
}
