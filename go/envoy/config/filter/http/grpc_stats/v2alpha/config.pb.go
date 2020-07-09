// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_stats/v2alpha/config.proto

package envoy_config_filter_http_grpc_stats_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
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

// gRPC statistics filter configuration
type FilterConfig struct {
	// If true, the filter maintains a filter state object with the request and response message
	// counts.
	EmitFilterState bool `protobuf:"varint,1,opt,name=emit_filter_state,json=emitFilterState,proto3" json:"emit_filter_state,omitempty"`
	// Types that are valid to be assigned to PerMethodStatSpecifier:
	//	*FilterConfig_IndividualMethodStatsAllowlist
	//	*FilterConfig_StatsForAllMethods
	PerMethodStatSpecifier isFilterConfig_PerMethodStatSpecifier `protobuf_oneof:"per_method_stat_specifier"`
	XXX_NoUnkeyedLiteral   struct{}                              `json:"-"`
	XXX_unrecognized       []byte                                `json:"-"`
	XXX_sizecache          int32                                 `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1419fab6b23f453d, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetEmitFilterState() bool {
	if m != nil {
		return m.EmitFilterState
	}
	return false
}

type isFilterConfig_PerMethodStatSpecifier interface {
	isFilterConfig_PerMethodStatSpecifier()
}

type FilterConfig_IndividualMethodStatsAllowlist struct {
	IndividualMethodStatsAllowlist *core.GrpcMethodList `protobuf:"bytes,2,opt,name=individual_method_stats_allowlist,json=individualMethodStatsAllowlist,proto3,oneof"`
}

type FilterConfig_StatsForAllMethods struct {
	StatsForAllMethods *wrappers.BoolValue `protobuf:"bytes,3,opt,name=stats_for_all_methods,json=statsForAllMethods,proto3,oneof"`
}

func (*FilterConfig_IndividualMethodStatsAllowlist) isFilterConfig_PerMethodStatSpecifier() {}

func (*FilterConfig_StatsForAllMethods) isFilterConfig_PerMethodStatSpecifier() {}

func (m *FilterConfig) GetPerMethodStatSpecifier() isFilterConfig_PerMethodStatSpecifier {
	if m != nil {
		return m.PerMethodStatSpecifier
	}
	return nil
}

func (m *FilterConfig) GetIndividualMethodStatsAllowlist() *core.GrpcMethodList {
	if x, ok := m.GetPerMethodStatSpecifier().(*FilterConfig_IndividualMethodStatsAllowlist); ok {
		return x.IndividualMethodStatsAllowlist
	}
	return nil
}

func (m *FilterConfig) GetStatsForAllMethods() *wrappers.BoolValue {
	if x, ok := m.GetPerMethodStatSpecifier().(*FilterConfig_StatsForAllMethods); ok {
		return x.StatsForAllMethods
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FilterConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FilterConfig_IndividualMethodStatsAllowlist)(nil),
		(*FilterConfig_StatsForAllMethods)(nil),
	}
}

// gRPC statistics filter state object in protobuf form.
type FilterObject struct {
	// Count of request messages in the request stream.
	RequestMessageCount uint64 `protobuf:"varint,1,opt,name=request_message_count,json=requestMessageCount,proto3" json:"request_message_count,omitempty"`
	// Count of response messages in the response stream.
	ResponseMessageCount uint64   `protobuf:"varint,2,opt,name=response_message_count,json=responseMessageCount,proto3" json:"response_message_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterObject) Reset()         { *m = FilterObject{} }
func (m *FilterObject) String() string { return proto.CompactTextString(m) }
func (*FilterObject) ProtoMessage()    {}
func (*FilterObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_1419fab6b23f453d, []int{1}
}

func (m *FilterObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterObject.Unmarshal(m, b)
}
func (m *FilterObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterObject.Marshal(b, m, deterministic)
}
func (m *FilterObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterObject.Merge(m, src)
}
func (m *FilterObject) XXX_Size() int {
	return xxx_messageInfo_FilterObject.Size(m)
}
func (m *FilterObject) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterObject.DiscardUnknown(m)
}

var xxx_messageInfo_FilterObject proto.InternalMessageInfo

func (m *FilterObject) GetRequestMessageCount() uint64 {
	if m != nil {
		return m.RequestMessageCount
	}
	return 0
}

func (m *FilterObject) GetResponseMessageCount() uint64 {
	if m != nil {
		return m.ResponseMessageCount
	}
	return 0
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.grpc_stats.v2alpha.FilterConfig")
	proto.RegisterType((*FilterObject)(nil), "envoy.config.filter.http.grpc_stats.v2alpha.FilterObject")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_stats/v2alpha/config.proto", fileDescriptor_1419fab6b23f453d)
}

var fileDescriptor_1419fab6b23f453d = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x9b, 0x05, 0x2a, 0x94, 0x22, 0x01, 0x81, 0x42, 0x59, 0x44, 0xd5, 0xf6, 0xb4, 0x02,
	0x61, 0x4b, 0x29, 0x07, 0x38, 0x76, 0x2b, 0x95, 0x1e, 0xa8, 0x5a, 0x2d, 0x12, 0xd7, 0xc8, 0x9b,
	0x4c, 0xb2, 0x46, 0x5e, 0x8f, 0xb1, 0x27, 0xe9, 0xf6, 0xc6, 0x85, 0x33, 0xd7, 0x3e, 0x0b, 0x4f,
	0xc0, 0x95, 0x57, 0xe1, 0x01, 0x10, 0xb2, 0x9d, 0xa8, 0x15, 0x3d, 0x71, 0xdb, 0xf5, 0xef, 0xef,
	0x1f, 0xff, 0x7f, 0x26, 0x7d, 0x0b, 0xba, 0xc3, 0x0b, 0x5e, 0xa2, 0xae, 0x65, 0xc3, 0x6b, 0xa9,
	0x08, 0x2c, 0x5f, 0x10, 0x19, 0xde, 0x58, 0x53, 0x16, 0x8e, 0x04, 0x39, 0xde, 0xe5, 0x42, 0x99,
	0x85, 0xe8, 0x6f, 0x31, 0x63, 0x91, 0x30, 0x7b, 0x15, 0x48, 0xd6, 0x9f, 0x45, 0x92, 0x79, 0x92,
	0x5d, 0x91, 0xac, 0x27, 0xc7, 0x93, 0x38, 0x46, 0x18, 0xc9, 0xbb, 0x9c, 0x97, 0x68, 0x21, 0xfa,
	0x2f, 0x81, 0x16, 0x58, 0x15, 0x4a, 0x3a, 0x8a, 0xb6, 0xe3, 0xed, 0x06, 0xb1, 0x51, 0xc0, 0xc3,
	0xbf, 0x79, 0x5b, 0xf3, 0x73, 0x2b, 0x8c, 0x01, 0xeb, 0x06, 0xbd, 0xad, 0x8c, 0xe0, 0x42, 0x6b,
	0x24, 0x41, 0x12, 0xb5, 0xe3, 0x4b, 0xd9, 0x58, 0x41, 0xd0, 0xeb, 0x2f, 0x6e, 0xe8, 0xfe, 0x25,
	0xed, 0x80, 0x3f, 0xed, 0x84, 0x92, 0x95, 0x20, 0xe0, 0xc3, 0x8f, 0x28, 0xec, 0x5d, 0x8e, 0xd2,
	0x7b, 0x47, 0x21, 0xc4, 0x61, 0x48, 0x94, 0xbd, 0x4c, 0x1f, 0xc2, 0x52, 0x52, 0x11, 0x93, 0x85,
	0x3c, 0xb0, 0x95, 0xec, 0x24, 0x93, 0xbb, 0xb3, 0xfb, 0x5e, 0x88, 0x97, 0x3f, 0xfa, 0xe3, 0x4c,
	0xa7, 0xbb, 0x52, 0x57, 0xb2, 0x93, 0x55, 0x2b, 0xd4, 0x10, 0x2a, 0x34, 0x50, 0x08, 0xa5, 0xf0,
	0xdc, 0xe7, 0xdb, 0x1a, 0xed, 0x24, 0x93, 0x8d, 0x7c, 0x97, 0xc5, 0xde, 0x84, 0x91, 0xac, 0xcb,
	0x99, 0xaf, 0x82, 0xbd, 0xb7, 0xa6, 0x3c, 0x09, 0xd0, 0x07, 0xe9, 0xe8, 0x78, 0x6d, 0xb6, 0x7d,
	0xe5, 0x16, 0xcf, 0xfd, 0x18, 0x77, 0x30, 0x58, 0x65, 0xa7, 0xe9, 0x66, 0x74, 0xaf, 0xd1, 0xfa,
	0x09, 0xfd, 0x48, 0xb7, 0x75, 0x2b, 0xcc, 0x18, 0xb3, 0x58, 0x22, 0x1b, 0x4a, 0x64, 0x53, 0x44,
	0xf5, 0x49, 0xa8, 0x16, 0x8e, 0xd7, 0x66, 0x59, 0x40, 0x8f, 0xd0, 0x1e, 0xa8, 0xde, 0xdd, 0x4d,
	0x9f, 0xa7, 0xcf, 0x0c, 0xd8, 0xeb, 0x2f, 0x2f, 0x9c, 0x81, 0x52, 0xd6, 0x12, 0xec, 0xde, 0x6a,
	0x68, 0xe6, 0x74, 0xfe, 0x19, 0x4a, 0xca, 0xf2, 0x74, 0xd3, 0xc2, 0x97, 0x16, 0x1c, 0x15, 0x4b,
	0x70, 0x4e, 0x34, 0x50, 0x94, 0xd8, 0x6a, 0x0a, 0xed, 0xdc, 0x9e, 0x3d, 0xea, 0xc5, 0x93, 0xa8,
	0x1d, 0x7a, 0x29, 0x7b, 0x93, 0x3e, 0xb1, 0xe0, 0x0c, 0x6a, 0x07, 0xff, 0x40, 0xa3, 0x00, 0x3d,
	0x1e, 0xd4, 0xeb, 0xd4, 0xf4, 0x5b, 0xf2, 0xfb, 0xf2, 0xcf, 0xf7, 0x3b, 0xaf, 0x87, 0x65, 0x83,
	0x15, 0x81, 0x76, 0xfe, 0xab, 0xf6, 0x0b, 0xe7, 0x6e, 0x6e, 0xdc, 0xfe, 0x8f, 0xaf, 0x3f, 0x7f,
	0xad, 0x8f, 0x1e, 0x24, 0xe9, 0x3b, 0x89, 0xb1, 0x6c, 0x63, 0x71, 0x75, 0xc1, 0xfe, 0x63, 0x5f,
	0xa7, 0x1b, 0x71, 0x09, 0xce, 0x7c, 0x7f, 0x67, 0xc9, 0x7c, 0x3d, 0x14, 0xb9, 0xff, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0x64, 0xa9, 0x6d, 0x2e, 0x03, 0x00, 0x00,
}
