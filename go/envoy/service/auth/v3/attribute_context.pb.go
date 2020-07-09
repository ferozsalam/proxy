// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v3/attribute_context.proto

package envoy_service_auth_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// An attribute is a piece of metadata that describes an activity on a network.
// For example, the size of an HTTP request, or the status code of an HTTP response.
//
// Each attribute has a type and a name, which is logically defined as a proto message field
// of the `AttributeContext`. The `AttributeContext` is a collection of individual attributes
// supported by Envoy authorization system.
// [#comment: The following items are left out of this proto
// Request.Auth field for jwt tokens
// Request.Api for api management
// Origin peer that originated the request
// Caching Protocol
// request_context return values to inject back into the filter chain
// peer.claims -- from X.509 extensions
// Configuration
// - field mask to send
// - which return values from request_context are copied back
// - which return values are copied into request_headers]
// [#next-free-field: 12]
type AttributeContext struct {
	// The source of a network activity, such as starting a TCP connection.
	// In a multi hop network activity, the source represents the sender of the
	// last hop.
	Source *AttributeContext_Peer `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The destination of a network activity, such as accepting a TCP connection.
	// In a multi hop network activity, the destination represents the receiver of
	// the last hop.
	Destination *AttributeContext_Peer `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Represents a network request, such as an HTTP request.
	Request *AttributeContext_Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	// This is analogous to http_request.headers, however these contents will not be sent to the
	// upstream server. Context_extensions provide an extension mechanism for sending additional
	// information to the auth server without modifying the proto definition. It maps to the
	// internal opaque context in the filter chain.
	ContextExtensions map[string]string `protobuf:"bytes,10,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Dynamic metadata associated with the request.
	MetadataContext      *v3.Metadata `protobuf:"bytes,11,opt,name=metadata_context,json=metadataContext,proto3" json:"metadata_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AttributeContext) Reset()         { *m = AttributeContext{} }
func (m *AttributeContext) String() string { return proto.CompactTextString(m) }
func (*AttributeContext) ProtoMessage()    {}
func (*AttributeContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a621374f4ae6edf, []int{0}
}

func (m *AttributeContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext.Unmarshal(m, b)
}
func (m *AttributeContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext.Marshal(b, m, deterministic)
}
func (m *AttributeContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext.Merge(m, src)
}
func (m *AttributeContext) XXX_Size() int {
	return xxx_messageInfo_AttributeContext.Size(m)
}
func (m *AttributeContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext proto.InternalMessageInfo

func (m *AttributeContext) GetSource() *AttributeContext_Peer {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttributeContext) GetDestination() *AttributeContext_Peer {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *AttributeContext) GetRequest() *AttributeContext_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AttributeContext) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func (m *AttributeContext) GetMetadataContext() *v3.Metadata {
	if m != nil {
		return m.MetadataContext
	}
	return nil
}

// This message defines attributes for a node that handles a network request.
// The node can be either a service or an application that sends, forwards,
// or receives the request. Service peers should fill in the `service`,
// `principal`, and `labels` as appropriate.
// [#next-free-field: 6]
type AttributeContext_Peer struct {
	// The address of the peer, this is typically the IP address.
	// It can also be UDS path, or others.
	Address *v3.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The canonical service name of the peer.
	// It should be set to :ref:`the HTTP x-envoy-downstream-service-cluster
	// <config_http_conn_man_headers_downstream-service-cluster>`
	// If a more trusted source of the service name is available through mTLS/secure naming, it
	// should be used.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// The labels associated with the peer.
	// These could be pod labels for Kubernetes or tags for VMs.
	// The source of the labels could be an X.509 certificate or other configuration.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The authenticated identity of this peer.
	// For example, the identity associated with the workload such as a service account.
	// If an X.509 certificate is used to assert the identity this field should be sourced from
	// `URI Subject Alternative Names`, `DNS Subject Alternate Names` or `Subject` in that order.
	// The primary identity should be the principal. The principal format is issuer specific.
	//
	// Example:
	// *    SPIFFE format is `spiffe://trust-domain/path`
	// *    Google account format is `https://accounts.google.com/{userid}`
	Principal string `protobuf:"bytes,4,opt,name=principal,proto3" json:"principal,omitempty"`
	// The X.509 certificate used to authenticate the identify of this peer.
	// When present, the certificate contents are encoded in URL and PEM format.
	Certificate          string   `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttributeContext_Peer) Reset()         { *m = AttributeContext_Peer{} }
func (m *AttributeContext_Peer) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Peer) ProtoMessage()    {}
func (*AttributeContext_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a621374f4ae6edf, []int{0, 0}
}

func (m *AttributeContext_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Peer.Unmarshal(m, b)
}
func (m *AttributeContext_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Peer.Marshal(b, m, deterministic)
}
func (m *AttributeContext_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Peer.Merge(m, src)
}
func (m *AttributeContext_Peer) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Peer.Size(m)
}
func (m *AttributeContext_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Peer proto.InternalMessageInfo

func (m *AttributeContext_Peer) GetAddress() *v3.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AttributeContext_Peer) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AttributeContext_Peer) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *AttributeContext_Peer) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *AttributeContext_Peer) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

// Represents a network request, such as an HTTP request.
type AttributeContext_Request struct {
	// The timestamp when the proxy receives the first byte of the request.
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Represents an HTTP request or an HTTP-like request.
	Http                 *AttributeContext_HttpRequest `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AttributeContext_Request) Reset()         { *m = AttributeContext_Request{} }
func (m *AttributeContext_Request) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Request) ProtoMessage()    {}
func (*AttributeContext_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a621374f4ae6edf, []int{0, 1}
}

func (m *AttributeContext_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Request.Unmarshal(m, b)
}
func (m *AttributeContext_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Request.Marshal(b, m, deterministic)
}
func (m *AttributeContext_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Request.Merge(m, src)
}
func (m *AttributeContext_Request) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Request.Size(m)
}
func (m *AttributeContext_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Request.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Request proto.InternalMessageInfo

func (m *AttributeContext_Request) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *AttributeContext_Request) GetHttp() *AttributeContext_HttpRequest {
	if m != nil {
		return m.Http
	}
	return nil
}

// This message defines attributes for an HTTP request.
// HTTP/1.x, HTTP/2, gRPC are all considered as HTTP requests.
// [#next-free-field: 12]
type AttributeContext_HttpRequest struct {
	// The unique ID for a request, which can be propagated to downstream
	// systems. The ID should have low probability of collision
	// within a single day for a specific service.
	// For HTTP requests, it should be X-Request-ID or equivalent.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The HTTP request method, such as `GET`, `POST`.
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// The HTTP request headers. If multiple headers share the same key, they
	// must be merged according to the HTTP spec. All header keys must be
	// lower-cased, because HTTP header keys are case-insensitive.
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The request target, as it appears in the first line of the HTTP request. This includes
	// the URL path and query-string. No decoding is performed.
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// The HTTP request `Host` or 'Authority` header value.
	Host string `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	// The HTTP URL scheme, such as `http` and `https`.
	Scheme string `protobuf:"bytes,6,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// This field is always empty, and exists for compatibility reasons. The HTTP URL query is
	// included in `path` field.
	Query string `protobuf:"bytes,7,opt,name=query,proto3" json:"query,omitempty"`
	// This field is always empty, and exists for compatibility reasons. The URL fragment is
	// not submitted as part of HTTP requests; it is unknowable.
	Fragment string `protobuf:"bytes,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
	// The HTTP request size in bytes. If unknown, it must be -1.
	Size int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// The network protocol used with the request, such as "HTTP/1.0", "HTTP/1.1", or "HTTP/2".
	//
	// See :repo:`headers.h:ProtocolStrings <source/common/http/headers.h>` for a list of all
	// possible values.
	Protocol string `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The HTTP request body.
	Body                 string   `protobuf:"bytes,11,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttributeContext_HttpRequest) Reset()         { *m = AttributeContext_HttpRequest{} }
func (m *AttributeContext_HttpRequest) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_HttpRequest) ProtoMessage()    {}
func (*AttributeContext_HttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a621374f4ae6edf, []int{0, 2}
}

func (m *AttributeContext_HttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_HttpRequest.Unmarshal(m, b)
}
func (m *AttributeContext_HttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_HttpRequest.Marshal(b, m, deterministic)
}
func (m *AttributeContext_HttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_HttpRequest.Merge(m, src)
}
func (m *AttributeContext_HttpRequest) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_HttpRequest.Size(m)
}
func (m *AttributeContext_HttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_HttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_HttpRequest proto.InternalMessageInfo

func (m *AttributeContext_HttpRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AttributeContext_HttpRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetFragment() string {
	if m != nil {
		return m.Fragment
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AttributeContext_HttpRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*AttributeContext)(nil), "envoy.service.auth.v3.AttributeContext")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v3.AttributeContext.ContextExtensionsEntry")
	proto.RegisterType((*AttributeContext_Peer)(nil), "envoy.service.auth.v3.AttributeContext.Peer")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v3.AttributeContext.Peer.LabelsEntry")
	proto.RegisterType((*AttributeContext_Request)(nil), "envoy.service.auth.v3.AttributeContext.Request")
	proto.RegisterType((*AttributeContext_HttpRequest)(nil), "envoy.service.auth.v3.AttributeContext.HttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v3.AttributeContext.HttpRequest.HeadersEntry")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v3/attribute_context.proto", fileDescriptor_2a621374f4ae6edf)
}

var fileDescriptor_2a621374f4ae6edf = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x85, 0x3e, 0x2c, 0x59, 0xc3, 0xa2, 0x55, 0x17, 0xb5, 0x41, 0x10, 0xb5, 0xab, 0xba, 0x40,
	0x21, 0xa0, 0x35, 0x59, 0x48, 0x2d, 0x6a, 0xeb, 0x50, 0xd4, 0xae, 0xdd, 0xda, 0x40, 0x5b, 0x08,
	0x44, 0x4e, 0xb9, 0x18, 0x2b, 0x72, 0x24, 0x2d, 0x22, 0x71, 0xe9, 0xdd, 0xa5, 0x60, 0xe5, 0x94,
	0x63, 0x7e, 0x43, 0xfe, 0x49, 0x72, 0xc8, 0x2d, 0x40, 0xae, 0xb9, 0xe6, 0xb7, 0xe4, 0x10, 0xec,
	0x07, 0x1d, 0xc1, 0xd1, 0x41, 0xf6, 0x49, 0x3b, 0xcb, 0x37, 0x6f, 0x67, 0xde, 0x9b, 0x11, 0x1c,
	0x62, 0xb6, 0xe0, 0xcb, 0x48, 0xa2, 0x58, 0xb0, 0x04, 0x23, 0x5a, 0xa8, 0x69, 0xb4, 0xe8, 0x47,
	0x54, 0x29, 0xc1, 0x46, 0x85, 0xc2, 0xab, 0x84, 0x67, 0x0a, 0x6f, 0x54, 0x98, 0x0b, 0xae, 0x38,
	0xd9, 0x31, 0xf0, 0xd0, 0xc1, 0x43, 0x0d, 0x0f, 0x17, 0xfd, 0xe0, 0xc0, 0xb2, 0x24, 0x3c, 0x1b,
	0xb3, 0x49, 0x94, 0x70, 0x81, 0x86, 0x24, 0x4d, 0x05, 0x4a, 0x69, 0x53, 0x83, 0xef, 0xd6, 0x62,
	0x46, 0x54, 0x62, 0x09, 0x98, 0x70, 0x3e, 0x99, 0x61, 0x64, 0xa2, 0x51, 0x31, 0x8e, 0x14, 0x9b,
	0xa3, 0x54, 0x74, 0x9e, 0x3b, 0xc0, 0x5e, 0x91, 0xe6, 0x34, 0xa2, 0x59, 0xc6, 0x15, 0x55, 0x8c,
	0x67, 0x32, 0x92, 0x8a, 0xaa, 0xa2, 0x7c, 0xe0, 0xfb, 0xcf, 0x3e, 0x2f, 0x50, 0x48, 0xc6, 0x33,
	0x96, 0x4d, 0x2c, 0xe4, 0xe0, 0x83, 0x07, 0xed, 0x93, 0xb2, 0xb5, 0xbf, 0x6c, 0x67, 0xe4, 0x0c,
	0x1a, 0x92, 0x17, 0x22, 0x41, 0xbf, 0xd2, 0xa9, 0x74, 0xbd, 0xde, 0xcf, 0xe1, 0xda, 0x26, 0xc3,
	0xbb, 0x89, 0xe1, 0x10, 0x51, 0xc4, 0x2e, 0x97, 0xfc, 0x0f, 0x5e, 0x8a, 0x52, 0xb1, 0xcc, 0xbc,
	0xed, 0x57, 0x1f, 0x40, 0xb5, 0x4a, 0x40, 0x2e, 0xa1, 0x29, 0xf0, 0xba, 0x40, 0xa9, 0xfc, 0xba,
	0xe1, 0x8a, 0x36, 0xe5, 0x8a, 0x6d, 0x5a, 0x5c, 0xe6, 0x93, 0x39, 0x10, 0xe7, 0xe2, 0x15, 0xde,
	0x28, 0xcc, 0xb4, 0x26, 0xd2, 0x87, 0x4e, 0xad, 0xeb, 0xf5, 0xfe, 0xd8, 0x94, 0xd5, 0xfd, 0x9e,
	0xdf, 0x12, 0x9c, 0x67, 0x4a, 0x2c, 0xe3, 0xaf, 0x93, 0xbb, 0xf7, 0xe4, 0x12, 0xda, 0x73, 0x54,
	0x34, 0xa5, 0x8a, 0x96, 0xd3, 0xe3, 0x7b, 0xa6, 0x85, 0x7d, 0xf7, 0x98, 0x9d, 0x81, 0x50, 0xcf,
	0x80, 0x7e, 0xeb, 0x3f, 0x87, 0x8e, 0xbf, 0x2a, 0xf3, 0xdc, 0x63, 0xc1, 0xfb, 0x2a, 0xd4, 0xb5,
	0x34, 0xe4, 0x77, 0x68, 0xba, 0x69, 0x72, 0x26, 0xed, 0xad, 0xa7, 0x3a, 0xb1, 0xa0, 0xb8, 0x44,
	0x13, 0x1f, 0x9a, 0xae, 0x35, 0x63, 0x49, 0x2b, 0x2e, 0x43, 0x32, 0x84, 0xc6, 0x8c, 0x8e, 0x70,
	0x26, 0xfd, 0x9a, 0x51, 0xe2, 0xe8, 0x3e, 0x5e, 0x85, 0xff, 0x9a, 0x54, 0xab, 0x81, 0xe3, 0x21,
	0xdf, 0x42, 0x2b, 0x17, 0x2c, 0x4b, 0x58, 0x4e, 0x67, 0xc6, 0xb4, 0x56, 0xfc, 0xe9, 0x82, 0x74,
	0xc0, 0x4b, 0x50, 0x28, 0x36, 0x66, 0x09, 0x55, 0xe8, 0x6f, 0x99, 0xef, 0xab, 0x57, 0xc1, 0x31,
	0x78, 0x2b, 0xb4, 0xa4, 0x0d, 0xb5, 0x27, 0xb8, 0x34, 0xfd, 0xb6, 0x62, 0x7d, 0x24, 0xdf, 0xc0,
	0xd6, 0x82, 0xce, 0x8a, 0xb2, 0x15, 0x1b, 0x0c, 0xaa, 0x47, 0x95, 0x41, 0xef, 0xc5, 0x9b, 0xe7,
	0xfb, 0x87, 0xf0, 0xd3, 0xba, 0x16, 0x7a, 0xeb, 0x5b, 0x08, 0x5e, 0x56, 0xa0, 0xe9, 0x66, 0x85,
	0x84, 0x50, 0xd7, 0xdb, 0xe6, 0xc4, 0x0d, 0x42, 0xbb, 0x8a, 0x61, 0xb9, 0x8a, 0xe1, 0xa3, 0x72,
	0x15, 0x63, 0x83, 0x23, 0xff, 0x40, 0x7d, 0xaa, 0x54, 0xee, 0xc6, 0xbc, 0xbf, 0xa9, 0x74, 0x17,
	0x4a, 0xe5, 0xe5, 0x78, 0x1a, 0x82, 0xc1, 0x6f, 0xba, 0xf0, 0x5f, 0x20, 0xdc, 0xb0, 0x70, 0x97,
	0x1c, 0xbc, 0xae, 0x81, 0xb7, 0x42, 0x46, 0xbe, 0x84, 0x2a, 0x4b, 0x9d, 0x54, 0x55, 0x96, 0x92,
	0x5d, 0x68, 0xcc, 0x51, 0x4d, 0x79, 0xea, 0xa4, 0x72, 0x11, 0x79, 0x0c, 0xcd, 0x29, 0xd2, 0x14,
	0x45, 0xe9, 0xfa, 0x9f, 0x0f, 0x28, 0x3d, 0xbc, 0xb0, 0x14, 0xd6, 0xfd, 0x92, 0x90, 0x10, 0xa8,
	0xe7, 0x54, 0x4d, 0x9d, 0xf3, 0xe6, 0xac, 0xef, 0xa6, 0x5c, 0x2a, 0xe7, 0xb6, 0x39, 0xeb, 0xda,
	0x64, 0x32, 0xc5, 0x39, 0xfa, 0x0d, 0x5b, 0x9b, 0x8d, 0xb4, 0xbb, 0xd7, 0x05, 0x8a, 0xa5, 0xdf,
	0xb4, 0xee, 0x9a, 0x80, 0x04, 0xb0, 0x3d, 0x16, 0x74, 0x32, 0xc7, 0x4c, 0xf9, 0xdb, 0xe6, 0xc3,
	0x6d, 0xac, 0xd9, 0x25, 0x7b, 0x8a, 0x7e, 0xab, 0x53, 0xe9, 0xd6, 0x62, 0x73, 0xd6, 0x78, 0xe3,
	0x5a, 0xc2, 0x67, 0x3e, 0x58, 0x7c, 0x19, 0x6b, 0xfc, 0x88, 0xa7, 0x4b, 0xb3, 0x8d, 0xad, 0xd8,
	0x9c, 0x83, 0x01, 0x7c, 0xb1, 0xda, 0xce, 0xbd, 0xa6, 0xee, 0x58, 0x9b, 0xf7, 0x2b, 0xf4, 0x36,
	0x34, 0x6f, 0x45, 0xc2, 0xe0, 0x0c, 0x76, 0xd7, 0xff, 0xa3, 0xdc, 0xab, 0x80, 0x43, 0x5d, 0x40,
	0x17, 0x7e, 0xdc, 0xac, 0x80, 0xd3, 0xbf, 0x5f, 0x3d, 0x7b, 0xfb, 0xae, 0x51, 0x6d, 0x57, 0xe1,
	0x07, 0xc6, 0xad, 0xf1, 0xb9, 0xe0, 0x37, 0xcb, 0xf5, 0x33, 0x70, 0xba, 0x73, 0x97, 0x60, 0xa8,
	0x85, 0x1c, 0x56, 0x46, 0x0d, 0xa3, 0x68, 0xff, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x58,
	0xa5, 0xbc, 0x3d, 0x07, 0x00, 0x00,
}
