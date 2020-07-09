// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package envoy_config_filter_http_transcoder_v2

import (
	fmt "fmt"
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

// [#next-free-field: 10]
type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	// Whether to route methods without the ``google.api.http`` option.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     package bookstore;
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {}
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The client could ``post`` a json body ``{"shelf": 1234}`` with the path of
	// ``/bookstore.Bookstore/GetShelfRequest`` to call ``GetShelfRequest``.
	AutoMapping bool `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"auto_mapping,omitempty"`
	// Whether to ignore query parameters that cannot be mapped to a corresponding
	// protobuf field. Use this if you cannot control the query parameters and do
	// not know them beforehand. Otherwise use ``ignored_query_parameters``.
	// Defaults to false.
	IgnoreUnknownQueryParameters bool `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignore_unknown_query_parameters,omitempty"`
	// Whether to convert gRPC status headers to JSON.
	// When trailer indicates a gRPC error and there was no HTTP body, take ``google.rpc.Status``
	// from the ``grpc-status-details-bin`` header and use it as JSON body.
	// If there was no such header, make ``google.rpc.Status`` out of the ``grpc-status`` and
	// ``grpc-message`` headers.
	// The error details types must be present in the ``proto_descriptor``.
	//
	// For example, if an upstream server replies with headers:
	//
	// .. code-block:: none
	//
	//     grpc-status: 5
	//     grpc-status-details-bin:
	//         CAUaMwoqdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLlJlcXVlc3RJbmZvEgUKA3ItMQ
	//
	// The ``grpc-status-details-bin`` header contains a base64-encoded protobuf message
	// ``google.rpc.Status``. It will be transcoded into:
	//
	// .. code-block:: none
	//
	//     HTTP/1.1 404 Not Found
	//     content-type: application/json
	//
	//     {"code":5,"details":[{"@type":"type.googleapis.com/google.rpc.RequestInfo","requestId":"r-1"}]}
	//
	//  In order to transcode the message, the ``google.rpc.RequestInfo`` type from
	//  the ``google/rpc/error_details.proto`` should be included in the configured
	//  :ref:`proto descriptor set <config_grpc_json_generate_proto_descriptor_set>`.
	ConvertGrpcStatus    bool     `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convert_grpc_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0}
}

func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder.Size(m)
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof"`
}

type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet() {}

func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetAutoMapping() bool {
	if m != nil {
		return m.AutoMapping
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoreUnknownQueryParameters() bool {
	if m != nil {
		return m.IgnoreUnknownQueryParameters
	}
	return false
}

func (m *GrpcJsonTranscoder) GetConvertGrpcStatus() bool {
	if m != nil {
		return m.ConvertGrpcStatus
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0, 0}
}

func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Size(m)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptor_540f2f6de8b0585c)
}

var fileDescriptor_540f2f6de8b0585c = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0x13, 0x3d,
	0x14, 0xc5, 0x3b, 0xed, 0xd7, 0x7c, 0x53, 0x37, 0xfd, 0x83, 0x81, 0x76, 0x88, 0x5a, 0x08, 0x08,
	0x50, 0x24, 0xa4, 0x19, 0x29, 0x05, 0x15, 0xc1, 0x02, 0x35, 0xa2, 0xb4, 0x45, 0x02, 0xc2, 0x00,
	0x62, 0x69, 0xb9, 0x33, 0xb7, 0xa9, 0x21, 0x63, 0xbb, 0xb6, 0x27, 0x6d, 0x76, 0x3c, 0x01, 0x48,
	0xac, 0x78, 0x05, 0x5e, 0x81, 0x27, 0x60, 0xcb, 0xa3, 0xc0, 0x0a, 0xb1, 0x40, 0xc8, 0x76, 0x1a,
	0x86, 0x76, 0xd3, 0xdd, 0xd8, 0xe7, 0x9c, 0x7b, 0xe7, 0x5e, 0xff, 0xd0, 0x3a, 0xf0, 0x81, 0x18,
	0x26, 0x99, 0xe0, 0x7b, 0xac, 0x97, 0xec, 0xb1, 0xbe, 0x01, 0x95, 0xec, 0x1b, 0x23, 0x13, 0xa3,
	0x28, 0xd7, 0x99, 0xc8, 0x41, 0x25, 0x83, 0x76, 0xe5, 0x14, 0x4b, 0x25, 0x8c, 0xc0, 0x37, 0x5d,
	0x30, 0xf6, 0xc1, 0xd8, 0x07, 0x63, 0x1b, 0x8c, 0x2b, 0xd6, 0x41, 0xbb, 0x71, 0xb9, 0xcc, 0x25,
	0x4d, 0x28, 0xe7, 0xc2, 0x50, 0xc3, 0x04, 0xd7, 0x49, 0xc1, 0x7a, 0x8a, 0x1a, 0xf0, 0x75, 0x1a,
	0xab, 0xa7, 0x74, 0x6d, 0xa8, 0x29, 0xf5, 0x48, 0x5e, 0x1e, 0xd0, 0x3e, 0xcb, 0xa9, 0x81, 0xe4,
	0xf8, 0xc3, 0x0b, 0xd7, 0x3e, 0xd7, 0x10, 0xde, 0x52, 0x32, 0x7b, 0xac, 0x05, 0x7f, 0x39, 0xee,
	0x88, 0x6f, 0xa1, 0x45, 0xa7, 0x93, 0x1c, 0x74, 0xa6, 0x98, 0x34, 0x42, 0x45, 0x41, 0x33, 0x68,
	0xcd, 0x6c, 0x4f, 0xa4, 0x0b, 0x4e, 0x79, 0x38, 0x16, 0x70, 0x1b, 0x5d, 0x38, 0x69, 0x26, 0xbb,
	0x8c, 0x47, 0xff, 0x35, 0x83, 0x56, 0x7d, 0x7b, 0x22, 0xc5, 0x27, 0x02, 0x1d, 0xc6, 0xf1, 0x75,
	0x14, 0x6a, 0x50, 0x03, 0x96, 0x81, 0x8e, 0x26, 0x9b, 0x53, 0xad, 0x99, 0x4e, 0xf8, 0xab, 0x33,
	0xfd, 0x31, 0x98, 0x0c, 0x83, 0x74, 0xac, 0xe0, 0x3e, 0x9a, 0x93, 0x8a, 0x71, 0x43, 0x84, 0x74,
	0x43, 0x45, 0x53, 0xcd, 0xa0, 0x35, 0xdb, 0xde, 0x8a, 0xcf, 0xb6, 0xb5, 0xf8, 0xf4, 0x64, 0x71,
	0xd7, 0xd6, 0x7b, 0xe6, 0xcb, 0xa5, 0x75, 0x59, 0x39, 0xe1, 0x07, 0x68, 0xa5, 0xa0, 0x26, 0xdb,
	0x27, 0x8c, 0x67, 0xa2, 0x60, 0xbc, 0x47, 0x14, 0x1c, 0x94, 0xa0, 0x0d, 0x51, 0xa2, 0x34, 0x10,
	0x4d, 0x37, 0x83, 0x56, 0x98, 0x5e, 0x72, 0x9e, 0x9d, 0x91, 0x25, 0xf5, 0x8e, 0xd4, 0x1a, 0xf0,
	0x5d, 0x14, 0xb1, 0x1e, 0x17, 0x0a, 0x72, 0x72, 0x50, 0x82, 0x1a, 0x12, 0x49, 0x15, 0x2d, 0xc0,
	0x80, 0xd2, 0x51, 0xcd, 0x0e, 0x99, 0x2e, 0x8d, 0xf4, 0xe7, 0x56, 0xee, 0x8e, 0x55, 0x7c, 0x15,
	0xd5, 0x69, 0x69, 0x04, 0x29, 0xa8, 0x94, 0x8c, 0xf7, 0xa2, 0xff, 0x5d, 0xab, 0x59, 0x7b, 0xf7,
	0xc4, 0x5f, 0xe1, 0x4d, 0x74, 0xc5, 0x87, 0x49, 0xc9, 0xdf, 0x72, 0x71, 0xc8, 0x4f, 0xf7, 0x08,
	0x5d, 0x6a, 0xc5, 0xdb, 0x5e, 0x79, 0xd7, 0xc9, 0x4e, 0x31, 0x3a, 0x9f, 0x09, 0x3e, 0x00, 0x65,
	0x48, 0x4f, 0xc9, 0x8c, 0x78, 0x4c, 0xa2, 0x19, 0x17, 0x3d, 0x37, 0x92, 0xec, 0xde, 0x5e, 0x38,
	0xa1, 0xf1, 0x3d, 0x40, 0xf5, 0xea, 0xce, 0xf0, 0x0d, 0x34, 0x4f, 0xf3, 0x9c, 0x1c, 0xee, 0x33,
	0x03, 0x5a, 0xd2, 0x0c, 0x1c, 0x18, 0x61, 0x3a, 0x47, 0xf3, 0xfc, 0xf5, 0xf8, 0x12, 0x6f, 0xa0,
	0x55, 0xda, 0x3f, 0xa4, 0x43, 0x4d, 0xfc, 0x0b, 0x4a, 0xc5, 0x0a, 0x66, 0xd8, 0x00, 0xc8, 0x1e,
	0x83, 0x7e, 0x6e, 0x5f, 0xdd, 0xa6, 0x1a, 0xde, 0xe4, 0x3a, 0x74, 0x8f, 0x2d, 0x8f, 0x9c, 0x03,
	0xdf, 0x43, 0x8d, 0x7f, 0x4a, 0x00, 0x2f, 0x0b, 0x4d, 0xa8, 0x26, 0x8c, 0x1b, 0x8f, 0x42, 0x98,
	0x2e, 0x55, 0xf2, 0x9b, 0x56, 0xdf, 0xd0, 0x3b, 0xdc, 0x68, 0x7c, 0x1f, 0x35, 0xa4, 0x02, 0x0b,
	0x12, 0x10, 0x0f, 0xa7, 0x6b, 0x4b, 0x38, 0x2d, 0x40, 0x3b, 0x32, 0xc3, 0x74, 0xf9, 0xd8, 0xd1,
	0xb5, 0x06, 0xd7, 0xf4, 0xa9, 0x95, 0x3b, 0x17, 0xd1, 0x7c, 0x05, 0x65, 0x0d, 0x06, 0x4f, 0xfd,
	0xec, 0x04, 0x9d, 0xf7, 0xc1, 0x8f, 0x4f, 0xbf, 0x3f, 0x4c, 0xaf, 0xe3, 0x3b, 0x1e, 0x3f, 0x38,
	0x32, 0xc0, 0xb5, 0x5d, 0xca, 0x08, 0x41, 0xed, 0x19, 0x74, 0x7b, 0x7d, 0xa3, 0x05, 0x27, 0x55,
	0x1a, 0xd7, 0xbe, 0xbc, 0xfb, 0xfa, 0xad, 0x36, 0xb9, 0x18, 0xa0, 0xdb, 0x4c, 0x78, 0x80, 0xa5,
	0x12, 0x47, 0xc3, 0x33, 0xb2, 0xdc, 0x59, 0xf8, 0x0b, 0xb1, 0xfb, 0xdf, 0x6e, 0xb0, 0x5b, 0x73,
	0x93, 0xad, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x8b, 0xae, 0xf9, 0x7e, 0x04, 0x00, 0x00,
}
