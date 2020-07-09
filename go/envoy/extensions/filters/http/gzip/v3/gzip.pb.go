// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/gzip/v3/gzip.proto

package envoy_extensions_filters_http_gzip_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/extensions/filters/http/compressor/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}

var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}

func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}

var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}

func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0, 0}
}

// [#next-free-field: 11]
type Gzip struct {
	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. The default value is 5.
	MemoryLevel *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	// A value used for selecting the zlib compression level. This setting will affect speed and
	// amount of compression applied to the content. "BEST" provides higher compression at the cost of
	// higher latency, "SPEED" provides lower compression with minimum impact on response time.
	// "DEFAULT" provides an optimal result between speed and compression. This field will be set to
	// "DEFAULT" if not specified.
	CompressionLevel Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	// A value used for selecting the zlib compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT" will be the best choice, though
	// there are situations which changing this parameter might produce better results. For example,
	// run-length encoding (RLE) is typically used when the content is known for having sequences
	// which same data occurs many consecutive times. For more information about each strategy, please
	// refer to zlib manual.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Set of configuration parameters common for all compression filters. If this field is set then
	// the fields `content_length`, `content_type`, `disable_on_etag_header` and
	// `remove_accept_encoding_header` are ignored.
	Compressor           *v3.Compressor `protobuf:"bytes,10,opt,name=compressor,proto3" json:"compressor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0}
}

func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (m *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(m, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetMemoryLevel() *wrappers.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

func (m *Gzip) GetCompressor() *v3.Compressor {
	if m != nil {
		return m.Compressor
	}
	return nil
}

type Gzip_CompressionLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gzip_CompressionLevel) Reset()         { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()    {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0}
}

func (m *Gzip_CompressionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip_CompressionLevel.Unmarshal(m, b)
}
func (m *Gzip_CompressionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip_CompressionLevel.Marshal(b, m, deterministic)
}
func (m *Gzip_CompressionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip_CompressionLevel.Merge(m, src)
}
func (m *Gzip_CompressionLevel) XXX_Size() int {
	return xxx_messageInfo_Gzip_CompressionLevel.Size(m)
}
func (m *Gzip_CompressionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip_CompressionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip_CompressionLevel proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.extensions.filters.http.gzip.v3.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.extensions.filters.http.gzip.v3.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
	proto.RegisterType((*Gzip)(nil), "envoy.extensions.filters.http.gzip.v3.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.extensions.filters.http.gzip.v3.Gzip.CompressionLevel")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/gzip/v3/gzip.proto", fileDescriptor_75f828f0702c619c)
}

var fileDescriptor_75f828f0702c619c = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x5f, 0x6b, 0xd3, 0x5e,
	0x18, 0xfe, 0xa5, 0xcd, 0xda, 0xe4, 0x74, 0xec, 0x77, 0xcc, 0x44, 0xcb, 0x70, 0x63, 0x0e, 0x86,
	0x45, 0xe4, 0x44, 0x5a, 0x50, 0x19, 0x8a, 0x2c, 0x2e, 0x75, 0x2b, 0x55, 0x46, 0xb6, 0xe9, 0x65,
	0x38, 0x4d, 0xdf, 0x65, 0x07, 0xd2, 0x73, 0xc2, 0xc9, 0x69, 0xb6, 0x4e, 0x2f, 0xc4, 0x2b, 0x3f,
	0x83, 0x1f, 0xc5, 0x7b, 0xc1, 0x1b, 0x2f, 0xfc, 0x3a, 0xbb, 0x92, 0x24, 0xad, 0xad, 0xdb, 0xd0,
	0xe2, 0x55, 0xfb, 0xf2, 0xbc, 0xcf, 0x9f, 0x3c, 0xe7, 0x45, 0x0f, 0x81, 0xa7, 0x62, 0x64, 0xc3,
	0x99, 0x02, 0x9e, 0x30, 0xc1, 0x13, 0xfb, 0x98, 0x45, 0x0a, 0x64, 0x62, 0x9f, 0x28, 0x15, 0xdb,
	0xe1, 0x39, 0x8b, 0xed, 0xb4, 0x95, 0xff, 0x92, 0x58, 0x0a, 0x25, 0xac, 0xcd, 0x9c, 0x41, 0xa6,
	0x0c, 0x32, 0x66, 0x90, 0x8c, 0x41, 0xf2, 0xcd, 0xb4, 0xb5, 0xf2, 0xf4, 0xcf, 0xc2, 0x81, 0x18,
	0xc4, 0x12, 0x92, 0x44, 0xc8, 0x4c, 0x7e, 0x3a, 0x15, 0x26, 0x2b, 0x6b, 0xa1, 0x10, 0x61, 0x04,
	0x76, 0x3e, 0xf5, 0x86, 0xc7, 0xf6, 0xa9, 0xa4, 0x71, 0x9c, 0x99, 0x14, 0xf8, 0xea, 0xb0, 0x1f,
	0x53, 0x9b, 0x72, 0x2e, 0x14, 0x55, 0xb9, 0x7a, 0xa2, 0xa8, 0x1a, 0x4e, 0xe0, 0xbb, 0x57, 0xe0,
	0x14, 0x64, 0x96, 0x82, 0xf1, 0x70, 0xbc, 0x72, 0x3b, 0xa5, 0x11, 0xeb, 0x53, 0x05, 0xf6, 0xe4,
	0x4f, 0x01, 0x6c, 0x7c, 0xaf, 0x20, 0xfd, 0xe5, 0x39, 0x8b, 0xad, 0x0e, 0x5a, 0x1c, 0xc0, 0x40,
	0xc8, 0x91, 0x1f, 0x41, 0x0a, 0x51, 0x5d, 0x5b, 0xd7, 0x1a, 0xb5, 0xe6, 0x1d, 0x52, 0x44, 0x23,
	0x93, 0x68, 0xe4, 0x68, 0x8f, 0xab, 0x56, 0xf3, 0x0d, 0x8d, 0x86, 0xe0, 0x98, 0x17, 0x4e, 0xe5,
	0xbe, 0x5e, 0x37, 0x1b, 0x9a, 0x57, 0x2b, 0xc8, 0xdd, 0x8c, 0x6b, 0x9d, 0xa1, 0x1b, 0x93, 0x6f,
	0x64, 0x82, 0x8f, 0x05, 0xcb, 0xeb, 0x5a, 0x63, 0xa9, 0xb9, 0x4d, 0xe6, 0x2a, 0x94, 0x64, 0x99,
	0xc8, 0x8b, 0xa9, 0x48, 0x2e, 0x4c, 0x5c, 0x3e, 0x1c, 0x38, 0xc6, 0x85, 0xb3, 0xf0, 0x51, 0x2b,
	0x61, 0xcd, 0xc3, 0xc1, 0xa5, 0x05, 0xeb, 0x3d, 0xba, 0x39, 0xeb, 0x9c, 0x28, 0x49, 0x15, 0x84,
	0xa3, 0xba, 0x9e, 0x9b, 0x3f, 0xff, 0x47, 0xf3, 0x83, 0xb1, 0xcc, 0x8c, 0xf5, 0x72, 0x70, 0x15,
	0xb6, 0x76, 0x51, 0xed, 0x94, 0xf1, 0xbe, 0x38, 0xf5, 0x7b, 0x4c, 0x25, 0x75, 0x73, 0xfe, 0x0a,
	0xff, 0x6f, 0x98, 0x1e, 0x2a, 0xb8, 0x0e, 0x53, 0x89, 0xf5, 0x16, 0xa1, 0xe9, 0x95, 0xd4, 0x51,
	0x2e, 0xf4, 0xf8, 0x2f, 0xe9, 0x67, 0xce, 0x2a, 0x6d, 0xfd, 0x8a, 0x2f, 0xa4, 0x37, 0x23, 0xb5,
	0xf2, 0x0e, 0xe1, 0xcb, 0xad, 0x6e, 0x34, 0x90, 0x9e, 0x15, 0x6b, 0xd5, 0x50, 0x75, 0xc7, 0x6d,
	0x6f, 0x1f, 0x75, 0x0f, 0xf1, 0x7f, 0x96, 0x81, 0x74, 0xc7, 0x3d, 0x38, 0xc4, 0x9a, 0x65, 0xa2,
	0x85, 0x83, 0x7d, 0xd7, 0xdd, 0xc1, 0xa5, 0xad, 0x67, 0x9f, 0xbf, 0x7e, 0x5a, 0x7b, 0x82, 0x1e,
	0x15, 0x41, 0x02, 0xc1, 0x8f, 0x59, 0x38, 0x0e, 0x31, 0xdb, 0x60, 0xf3, 0xfa, 0xe7, 0xdb, 0x68,
	0xa3, 0xe5, 0x6b, 0x5a, 0xfd, 0xdd, 0x77, 0x11, 0x19, 0xed, 0xbd, 0xee, 0xa1, 0xeb, 0xb9, 0x3b,
	0x58, 0xcb, 0xa0, 0xdd, 0xa3, 0x76, 0xfb, 0xd5, 0xf6, 0x6b, 0x5c, 0xb2, 0xaa, 0xa8, 0xec, 0x75,
	0x5d, 0x5c, 0xde, 0x7a, 0x90, 0xc5, 0xb8, 0x87, 0x36, 0xe7, 0x8a, 0xd1, 0xd1, 0x8d, 0x12, 0x2e,
	0x77, 0x74, 0xa3, 0x82, 0xab, 0x1d, 0xdd, 0xa8, 0x62, 0xa3, 0xa3, 0x1b, 0x06, 0x36, 0xbd, 0xa5,
	0x40, 0x70, 0x05, 0x5c, 0xf9, 0x11, 0xf0, 0x50, 0x9d, 0x78, 0x8b, 0x93, 0x59, 0x8d, 0x62, 0xf0,
	0x6e, 0xf5, 0x59, 0x42, 0x7b, 0x11, 0xf8, 0x82, 0xfb, 0xa0, 0x68, 0xe8, 0x9f, 0x00, 0xed, 0x83,
	0xf4, 0x56, 0x25, 0x0c, 0x44, 0x0a, 0x3e, 0x0d, 0x02, 0x88, 0x95, 0x0f, 0x3c, 0x10, 0x7d, 0xc6,
	0x27, 0xb0, 0xd3, 0xf9, 0xf2, 0xe1, 0xdb, 0x8f, 0x4a, 0x09, 0x97, 0x50, 0x8b, 0x89, 0xe2, 0xc1,
	0x62, 0x29, 0xce, 0x46, 0xf3, 0x5d, 0x9e, 0x63, 0x66, 0x89, 0xf7, 0xb3, 0x43, 0xd9, 0xd7, 0x7a,
	0x95, 0xfc, 0x62, 0x5a, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x19, 0x0f, 0x07, 0xb6, 0x04,
	0x00, 0x00,
}
