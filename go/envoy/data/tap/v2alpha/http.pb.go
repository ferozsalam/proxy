// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/http.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
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

// A fully buffered HTTP trace message.
type HttpBufferedTrace struct {
	// Request message.
	Request *HttpBufferedTrace_Message `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Response message.
	Response             *HttpBufferedTrace_Message `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpBufferedTrace) Reset()         { *m = HttpBufferedTrace{} }
func (m *HttpBufferedTrace) String() string { return proto.CompactTextString(m) }
func (*HttpBufferedTrace) ProtoMessage()    {}
func (*HttpBufferedTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d8a92b44eb7244, []int{0}
}

func (m *HttpBufferedTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpBufferedTrace.Unmarshal(m, b)
}
func (m *HttpBufferedTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpBufferedTrace.Marshal(b, m, deterministic)
}
func (m *HttpBufferedTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBufferedTrace.Merge(m, src)
}
func (m *HttpBufferedTrace) XXX_Size() int {
	return xxx_messageInfo_HttpBufferedTrace.Size(m)
}
func (m *HttpBufferedTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBufferedTrace.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBufferedTrace proto.InternalMessageInfo

func (m *HttpBufferedTrace) GetRequest() *HttpBufferedTrace_Message {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *HttpBufferedTrace) GetResponse() *HttpBufferedTrace_Message {
	if m != nil {
		return m.Response
	}
	return nil
}

// HTTP message wrapper.
type HttpBufferedTrace_Message struct {
	// Message headers.
	Headers []*core.HeaderValue `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	// Message body.
	Body *Body `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// Message trailers.
	Trailers             []*core.HeaderValue `protobuf:"bytes,3,rep,name=trailers,proto3" json:"trailers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HttpBufferedTrace_Message) Reset()         { *m = HttpBufferedTrace_Message{} }
func (m *HttpBufferedTrace_Message) String() string { return proto.CompactTextString(m) }
func (*HttpBufferedTrace_Message) ProtoMessage()    {}
func (*HttpBufferedTrace_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d8a92b44eb7244, []int{0, 0}
}

func (m *HttpBufferedTrace_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpBufferedTrace_Message.Unmarshal(m, b)
}
func (m *HttpBufferedTrace_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpBufferedTrace_Message.Marshal(b, m, deterministic)
}
func (m *HttpBufferedTrace_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBufferedTrace_Message.Merge(m, src)
}
func (m *HttpBufferedTrace_Message) XXX_Size() int {
	return xxx_messageInfo_HttpBufferedTrace_Message.Size(m)
}
func (m *HttpBufferedTrace_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBufferedTrace_Message.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBufferedTrace_Message proto.InternalMessageInfo

func (m *HttpBufferedTrace_Message) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpBufferedTrace_Message) GetBody() *Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HttpBufferedTrace_Message) GetTrailers() []*core.HeaderValue {
	if m != nil {
		return m.Trailers
	}
	return nil
}

// A streamed HTTP trace segment. Multiple segments make up a full trace.
type HttpStreamedTraceSegment struct {
	// Trace ID unique to the originating Envoy only. Trace IDs can repeat and should not be used
	// for long term stable uniqueness.
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are valid to be assigned to MessagePiece:
	//	*HttpStreamedTraceSegment_RequestHeaders
	//	*HttpStreamedTraceSegment_RequestBodyChunk
	//	*HttpStreamedTraceSegment_RequestTrailers
	//	*HttpStreamedTraceSegment_ResponseHeaders
	//	*HttpStreamedTraceSegment_ResponseBodyChunk
	//	*HttpStreamedTraceSegment_ResponseTrailers
	MessagePiece         isHttpStreamedTraceSegment_MessagePiece `protobuf_oneof:"message_piece"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *HttpStreamedTraceSegment) Reset()         { *m = HttpStreamedTraceSegment{} }
func (m *HttpStreamedTraceSegment) String() string { return proto.CompactTextString(m) }
func (*HttpStreamedTraceSegment) ProtoMessage()    {}
func (*HttpStreamedTraceSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d8a92b44eb7244, []int{1}
}

func (m *HttpStreamedTraceSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpStreamedTraceSegment.Unmarshal(m, b)
}
func (m *HttpStreamedTraceSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpStreamedTraceSegment.Marshal(b, m, deterministic)
}
func (m *HttpStreamedTraceSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStreamedTraceSegment.Merge(m, src)
}
func (m *HttpStreamedTraceSegment) XXX_Size() int {
	return xxx_messageInfo_HttpStreamedTraceSegment.Size(m)
}
func (m *HttpStreamedTraceSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStreamedTraceSegment.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStreamedTraceSegment proto.InternalMessageInfo

func (m *HttpStreamedTraceSegment) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

type isHttpStreamedTraceSegment_MessagePiece interface {
	isHttpStreamedTraceSegment_MessagePiece()
}

type HttpStreamedTraceSegment_RequestHeaders struct {
	RequestHeaders *core.HeaderMap `protobuf:"bytes,2,opt,name=request_headers,json=requestHeaders,proto3,oneof"`
}

type HttpStreamedTraceSegment_RequestBodyChunk struct {
	RequestBodyChunk *Body `protobuf:"bytes,3,opt,name=request_body_chunk,json=requestBodyChunk,proto3,oneof"`
}

type HttpStreamedTraceSegment_RequestTrailers struct {
	RequestTrailers *core.HeaderMap `protobuf:"bytes,4,opt,name=request_trailers,json=requestTrailers,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseHeaders struct {
	ResponseHeaders *core.HeaderMap `protobuf:"bytes,5,opt,name=response_headers,json=responseHeaders,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseBodyChunk struct {
	ResponseBodyChunk *Body `protobuf:"bytes,6,opt,name=response_body_chunk,json=responseBodyChunk,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseTrailers struct {
	ResponseTrailers *core.HeaderMap `protobuf:"bytes,7,opt,name=response_trailers,json=responseTrailers,proto3,oneof"`
}

func (*HttpStreamedTraceSegment_RequestHeaders) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_RequestBodyChunk) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_RequestTrailers) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseHeaders) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseBodyChunk) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseTrailers) isHttpStreamedTraceSegment_MessagePiece() {}

func (m *HttpStreamedTraceSegment) GetMessagePiece() isHttpStreamedTraceSegment_MessagePiece {
	if m != nil {
		return m.MessagePiece
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestHeaders() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestHeaders); ok {
		return x.RequestHeaders
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestBodyChunk() *Body {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestBodyChunk); ok {
		return x.RequestBodyChunk
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestTrailers() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestTrailers); ok {
		return x.RequestTrailers
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseHeaders() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseHeaders); ok {
		return x.ResponseHeaders
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseBodyChunk() *Body {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseBodyChunk); ok {
		return x.ResponseBodyChunk
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseTrailers() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseTrailers); ok {
		return x.ResponseTrailers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpStreamedTraceSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpStreamedTraceSegment_RequestHeaders)(nil),
		(*HttpStreamedTraceSegment_RequestBodyChunk)(nil),
		(*HttpStreamedTraceSegment_RequestTrailers)(nil),
		(*HttpStreamedTraceSegment_ResponseHeaders)(nil),
		(*HttpStreamedTraceSegment_ResponseBodyChunk)(nil),
		(*HttpStreamedTraceSegment_ResponseTrailers)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpBufferedTrace)(nil), "envoy.data.tap.v2alpha.HttpBufferedTrace")
	proto.RegisterType((*HttpBufferedTrace_Message)(nil), "envoy.data.tap.v2alpha.HttpBufferedTrace.Message")
	proto.RegisterType((*HttpStreamedTraceSegment)(nil), "envoy.data.tap.v2alpha.HttpStreamedTraceSegment")
}

func init() { proto.RegisterFile("envoy/data/tap/v2alpha/http.proto", fileDescriptor_90d8a92b44eb7244) }

var fileDescriptor_90d8a92b44eb7244 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x9b, 0xee, 0xd2, 0x94, 0xa9, 0xa0, 0xd4, 0x48, 0x28, 0xac, 0x2a, 0x54, 0x0a, 0x87,
	0x9e, 0x1c, 0x48, 0x2f, 0x88, 0x63, 0x38, 0x90, 0xaa, 0x2c, 0xaa, 0xd2, 0x8a, 0x6b, 0x34, 0x9b,
	0x4c, 0xbb, 0x11, 0x9b, 0xd8, 0x38, 0xde, 0x15, 0xfb, 0x16, 0xbc, 0x0a, 0xcf, 0xc4, 0x8b, 0x20,
	0x27, 0x76, 0x0a, 0x2a, 0x2b, 0xa2, 0x1e, 0x63, 0xfd, 0xff, 0x3f, 0xff, 0x37, 0xb1, 0xe1, 0x25,
	0xd5, 0x2b, 0xb1, 0x0e, 0x0b, 0xd4, 0x18, 0x6a, 0x94, 0xe1, 0x2a, 0xc2, 0x85, 0x9c, 0x63, 0x38,
	0xd7, 0x5a, 0x72, 0xa9, 0x84, 0x16, 0xec, 0x59, 0x2b, 0xe1, 0x46, 0xc2, 0x35, 0x4a, 0x6e, 0x25,
	0x93, 0xc3, 0xce, 0x8a, 0xb2, 0x0c, 0x57, 0x51, 0x98, 0x0b, 0x45, 0xe1, 0x0c, 0x1b, 0xea, 0x5c,
	0x93, 0x57, 0x1b, 0x82, 0x73, 0x51, 0x55, 0xa2, 0xee, 0x44, 0xc7, 0xbf, 0xb6, 0xe1, 0x20, 0xd1,
	0x5a, 0xc6, 0xcb, 0xeb, 0x6b, 0x52, 0x54, 0x5c, 0x29, 0xcc, 0x89, 0x9d, 0x83, 0xaf, 0xe8, 0xdb,
	0x92, 0x1a, 0x1d, 0x78, 0x47, 0xde, 0xc9, 0x5e, 0xf4, 0x96, 0xff, 0xbb, 0x02, 0xbf, 0xe3, 0xe5,
	0x53, 0x6a, 0x1a, 0xbc, 0xa1, 0xd4, 0x25, 0xb0, 0x29, 0xec, 0x2a, 0x6a, 0xa4, 0xa8, 0x1b, 0x0a,
	0xb6, 0xef, 0x9b, 0xd6, 0x47, 0x4c, 0x7e, 0x7a, 0xe0, 0xdb, 0x53, 0xf6, 0x0e, 0xfc, 0x39, 0x61,
	0x41, 0xaa, 0x09, 0xbc, 0xa3, 0xd1, 0xc9, 0x5e, 0xf4, 0xc2, 0x26, 0xa3, 0x2c, 0xf9, 0x2a, 0xe2,
	0x66, 0x25, 0x3c, 0x69, 0x15, 0x5f, 0x70, 0xb1, 0xa4, 0xd4, 0xc9, 0xd9, 0x1b, 0x18, 0xcf, 0x44,
	0xb1, 0xb6, 0x85, 0x0e, 0x37, 0x15, 0x8a, 0x45, 0xb1, 0x4e, 0x5b, 0x25, 0x7b, 0x0f, 0xbb, 0x5a,
	0x61, 0xb9, 0x30, 0xc3, 0x46, 0x83, 0x86, 0xf5, 0xfa, 0xe3, 0x1f, 0x63, 0x08, 0x0c, 0xdb, 0xa5,
	0x56, 0x84, 0x95, 0x65, 0xbb, 0xa4, 0x9b, 0x8a, 0x6a, 0xcd, 0x9e, 0xb7, 0xc1, 0x39, 0x65, 0x65,
	0xd1, 0x6e, 0x7b, 0x9c, 0xfa, 0xed, 0xf7, 0x59, 0xc1, 0x3e, 0xc2, 0xbe, 0xdd, 0x62, 0xe6, 0x38,
	0xff, 0x2e, 0x7c, 0x77, 0xf4, 0x14, 0x65, 0xb2, 0x95, 0x3e, 0xb6, 0xb6, 0xc4, 0xe2, 0x7e, 0x02,
	0xe6, 0x82, 0x0c, 0x4c, 0x96, 0xcf, 0x97, 0xf5, 0xd7, 0x60, 0xf4, 0x7f, 0xf8, 0x64, 0x2b, 0x7d,
	0x62, 0x9d, 0xe6, 0xf3, 0x83, 0xf1, 0xb1, 0x33, 0x70, 0x67, 0x59, 0xbf, 0x92, 0xf1, 0xa0, 0x5e,
	0x0e, 0xe7, 0xca, 0xda, 0xba, 0xa8, 0xee, 0xcf, 0xf6, 0x88, 0x0f, 0x86, 0x46, 0x75, 0x3e, 0xc7,
	0xf8, 0x19, 0x9e, 0xf6, 0x51, 0x7f, 0x40, 0xee, 0x0c, 0x82, 0x3c, 0x70, 0xd6, 0x5b, 0xca, 0x73,
	0xe8, 0x0f, 0x6f, 0x31, 0xfd, 0x41, 0xdd, 0x7a, 0x26, 0xc7, 0x19, 0xef, 0xc3, 0xa3, 0xaa, 0xbb,
	0xb4, 0x99, 0x2c, 0x29, 0xa7, 0xf8, 0x14, 0x5e, 0x97, 0xa2, 0x8b, 0x91, 0x4a, 0x7c, 0x5f, 0x6f,
	0xe8, 0x17, 0x3f, 0x34, 0xf7, 0xe6, 0xc2, 0xbc, 0xd5, 0x0b, 0x6f, 0xb6, 0xd3, 0x3e, 0xda, 0xd3,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x93, 0xbf, 0xfa, 0x34, 0x04, 0x00, 0x00,
}