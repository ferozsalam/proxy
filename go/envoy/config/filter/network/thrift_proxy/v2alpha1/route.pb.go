// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto

package envoy_config_filter_network_thrift_proxy_v2alpha1

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
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

type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	// Inverts whatever matching is done in the :ref:`method_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.method_name>` or
	// :ref:`service_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.service_name>` fields.
	// Cannot be combined with wildcard matching as that would result in routes never being matched.
	//
	// .. note::
	//
	//   This does not invert matching done as part of the :ref:`headers field
	//   <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.headers>` field. To
	//   invert header matching, see :ref:`invert_match
	//   <envoy_api_field_route.HeaderMatcher.invert_match>`.
	Invert bool `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config). Note that this only applies for Thrift transports and/or
	// protocols that support headers.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

// [#next-free-field: 7]
type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	//	*RouteAction_ClusterHeader
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field will be considered.
	// Note that this will be merged with what's provided in :ref:`WeightedCluster.metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight.metadata_match>`,
	// with values there taking precedence. Keys and values should be provided under the "envoy.lb"
	// metadata key.
	MetadataMatch *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// Specifies a set of rate limit configurations that could be applied to the route.
	// N.B. Thrift service or method name matching can be achieved by specifying a RequestHeaders
	// action with the header name ":method-name".
	RateLimits []*route.RateLimit `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Strip the service prefix from the method name, if there's a prefix. For
	// example, the method call Service:method would end up being just method.
	StripServiceName     bool     `protobuf:"varint,5,opt,name=strip_service_name,json=stripServiceName,proto3" json:"strip_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

type RouteAction_ClusterHeader struct {
	ClusterHeader string `protobuf:"bytes,6,opt,name=cluster_header,json=clusterHeader,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_ClusterHeader) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetClusterHeader() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_ClusterHeader); ok {
		return x.ClusterHeader
	}
	return ""
}

func (m *RouteAction) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*route.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *RouteAction) GetStripServiceName() bool {
	if m != nil {
		return m.StripServiceName
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
		(*RouteAction_ClusterHeader)(nil),
	}
}

// Allows for specification of multiple upstream clusters along with weights that indicate the
// percentage of traffic to be forwarded to each cluster. The router selects an upstream cluster
// based on these weights.
type WeightedCluster struct {
	// Specifies one or more upstream clusters associated with the route.
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is determined by its
	// weight. The sum of weights across all entries in the clusters array determines the total
	// weight.
	Weight *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field, combined with what's
	// provided in :ref:`RouteAction's metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteAction.metadata_match>`,
	// will be considered. Values here will take precedence. Keys and values should be provided
	// under the "envoy.lb" metadata key.
	MetadataMatch        *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto", fileDescriptor_3de6ac0eae6369f5)
}

var fileDescriptor_3de6ac0eae6369f5 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xb1, 0x6f, 0x13, 0x3d,
	0x14, 0xaf, 0x2f, 0x4d, 0x9a, 0xcf, 0xf9, 0xda, 0xa6, 0x1e, 0xfa, 0x45, 0x69, 0xfb, 0x29, 0x4d,
	0x97, 0x80, 0x90, 0xaf, 0x4d, 0x17, 0x24, 0xd4, 0x22, 0xae, 0x4b, 0x90, 0x28, 0x8a, 0x0e, 0x01,
	0x13, 0x9c, 0xdc, 0x8b, 0x93, 0x58, 0x24, 0xe7, 0xc3, 0x76, 0x2e, 0x2d, 0x13, 0x33, 0x4b, 0x25,
	0x26, 0x76, 0xfe, 0x00, 0x76, 0x98, 0x58, 0x10, 0x2b, 0xff, 0x0a, 0x13, 0xea, 0x80, 0xd0, 0xd9,
	0xbe, 0x34, 0x29, 0x30, 0xb4, 0xb0, 0x24, 0xbe, 0xf7, 0xde, 0xef, 0xf7, 0xde, 0xfb, 0x3d, 0xfb,
	0xc1, 0x3d, 0x1a, 0x25, 0xfc, 0xc4, 0x0d, 0x79, 0xd4, 0x65, 0x3d, 0xb7, 0xcb, 0x06, 0x8a, 0x0a,
	0x37, 0xa2, 0x6a, 0xcc, 0xc5, 0x33, 0x57, 0xf5, 0x05, 0xeb, 0xaa, 0x20, 0x16, 0xfc, 0xf8, 0xc4,
	0x4d, 0x9a, 0x64, 0x10, 0xf7, 0xc9, 0x8e, 0x2b, 0xf8, 0x48, 0x51, 0x1c, 0x0b, 0xae, 0x38, 0xda,
	0xd1, 0x70, 0x6c, 0xe0, 0xd8, 0xc0, 0xb1, 0x85, 0xe3, 0x69, 0x38, 0xce, 0xe0, 0xd5, 0x75, 0x93,
	0x91, 0xc4, 0xcc, 0x4d, 0x9a, 0x6e, 0xc8, 0x05, 0x75, 0x8f, 0x88, 0xb4, 0x84, 0xd5, 0x6b, 0x33,
	0x5e, 0x9d, 0xca, 0xfc, 0x06, 0x21, 0x1f, 0xc6, 0x3c, 0xa2, 0x91, 0x92, 0x36, 0xf4, 0xff, 0x1e,
	0xe7, 0xbd, 0x01, 0x75, 0xf5, 0xd7, 0xd1, 0xa8, 0xeb, 0x8e, 0x05, 0x89, 0x63, 0x2a, 0x26, 0xfe,
	0x51, 0x27, 0x26, 0x2e, 0x89, 0x22, 0xae, 0x88, 0x62, 0x3c, 0x92, 0xee, 0x90, 0xf5, 0x04, 0xc9,
	0x6a, 0xaf, 0x6e, 0xfc, 0xe4, 0x97, 0x8a, 0xa8, 0x51, 0x06, 0xff, 0x2f, 0x21, 0x03, 0xd6, 0x21,
	0x8a, 0xba, 0xd9, 0xc1, 0x38, 0xea, 0x2f, 0x20, 0xf2, 0xd3, 0x8a, 0x0e, 0x74, 0xd3, 0x23, 0xa1,
	0xc1, 0x08, 0xc1, 0xf9, 0x88, 0x0c, 0x69, 0x05, 0xd4, 0x40, 0xe3, 0x1f, 0x5f, 0x9f, 0x51, 0x1b,
	0x16, 0x74, 0xed, 0xb2, 0xe2, 0xd4, 0x72, 0x8d, 0x52, 0xf3, 0x26, 0xbe, 0xb4, 0x5c, 0x58, 0xa7,
	0xf2, 0x2d, 0x4f, 0xfd, 0x13, 0x80, 0x79, 0x6d, 0x41, 0x4f, 0x60, 0x7e, 0x48, 0x54, 0xd8, 0xd7,
	0x09, 0x4b, 0xcd, 0xbd, 0xab, 0x52, 0x1f, 0xa6, 0x24, 0x5e, 0xf1, 0xcc, 0xcb, 0xbf, 0x02, 0x4e,
	0x19, 0xf8, 0x86, 0x15, 0x3d, 0x85, 0x79, 0x9d, 0xb2, 0xe2, 0x68, 0xfa, 0xfd, 0xab, 0xd2, 0xdf,
	0x09, 0x53, 0x75, 0xa6, 0xf9, 0x35, 0x6d, 0xfd, 0x23, 0x80, 0xf0, 0x3c, 0x3f, 0xda, 0x84, 0xa5,
	0x21, 0x55, 0x7d, 0xde, 0x09, 0xce, 0x45, 0x6c, 0xcd, 0xf9, 0xd0, 0x18, 0xef, 0xa7, 0x62, 0x6e,
	0xc1, 0x7f, 0x25, 0x15, 0x09, 0x0b, 0xa9, 0x89, 0x71, 0x6c, 0x4c, 0xc9, 0x5a, 0x75, 0xd0, 0x2a,
	0x2c, 0xb0, 0x28, 0xa1, 0x42, 0x55, 0x72, 0x35, 0xd0, 0x28, 0xfa, 0xf6, 0x0b, 0xdd, 0x82, 0x0b,
	0x7d, 0x4a, 0x3a, 0x54, 0xc8, 0xca, 0xbc, 0x1e, 0xc5, 0xa6, 0x6d, 0x88, 0xc4, 0x0c, 0x27, 0x4d,
	0x6c, 0xee, 0x74, 0x4b, 0x87, 0xe8, 0x8a, 0xa8, 0xf0, 0x33, 0x84, 0xb7, 0x0a, 0x97, 0xb5, 0x28,
	0x81, 0x8c, 0x69, 0xc8, 0xba, 0x8c, 0x0a, 0x94, 0xfb, 0xe6, 0x81, 0xfa, 0xdb, 0x1c, 0x2c, 0x4d,
	0x35, 0x89, 0xb6, 0xe0, 0x42, 0x38, 0x18, 0x49, 0x45, 0x85, 0x69, 0xc0, 0x5b, 0x38, 0xf3, 0xe6,
	0x85, 0x53, 0x03, 0xad, 0x39, 0x3f, 0xf3, 0xa0, 0xe7, 0x70, 0x65, 0x4c, 0x59, 0xaf, 0xaf, 0x68,
	0x27, 0xb0, 0x36, 0x69, 0x45, 0xf6, 0xae, 0x20, 0xf2, 0x63, 0xcb, 0x75, 0x60, 0xa8, 0x5a, 0x73,
	0x7e, 0x79, 0x3c, 0x6b, 0x92, 0x68, 0x1b, 0x2e, 0xd9, 0x4c, 0x81, 0x69, 0xa9, 0x52, 0xb8, 0x58,
	0xde, 0xa2, 0x0d, 0x30, 0x0a, 0x20, 0x0f, 0x2e, 0x0d, 0xa9, 0x22, 0x1d, 0xa2, 0x48, 0x60, 0x6e,
	0x59, 0x4e, 0x57, 0xb8, 0x36, 0xab, 0x5a, 0xfa, 0x78, 0xf1, 0xa1, 0x0d, 0xf4, 0x17, 0x33, 0x88,
	0x19, 0xe9, 0x3e, 0x2c, 0xa5, 0x8f, 0x2d, 0x18, 0xb0, 0x21, 0x53, 0x99, 0xec, 0x1b, 0xbf, 0x92,
	0xdd, 0x27, 0x8a, 0xde, 0x4b, 0xa3, 0x7c, 0x28, 0xb2, 0xa3, 0x44, 0x37, 0x20, 0x92, 0x4a, 0xb0,
	0x38, 0x98, 0x99, 0x7a, 0x5e, 0x8f, 0xb5, 0xac, 0x3d, 0x0f, 0xce, 0x07, 0xef, 0x55, 0xe0, 0x4a,
	0xd6, 0xe3, 0x85, 0x29, 0x7d, 0x70, 0xe0, 0xf2, 0x05, 0x95, 0x50, 0x02, 0x8b, 0x13, 0xed, 0x81,
	0x2e, 0xac, 0xfd, 0xe7, 0xda, 0x63, 0xfb, 0x6f, 0xcc, 0xfa, 0xca, 0xbf, 0x06, 0x4e, 0x11, 0xf8,
	0x93, 0x5c, 0xd5, 0x77, 0x00, 0x2e, 0xce, 0x44, 0xa1, 0xb5, 0xe9, 0xb5, 0x31, 0x99, 0x88, 0xdd,
	0x1f, 0x7b, 0xb0, 0x60, 0x86, 0x69, 0x2f, 0xc8, 0x3a, 0x36, 0x2b, 0x0f, 0x67, 0x2b, 0x0f, 0x3f,
	0xbc, 0x1b, 0xa9, 0xdd, 0xe6, 0x23, 0x32, 0x18, 0x51, 0x0d, 0xbe, 0xee, 0x34, 0x80, 0x6f, 0x41,
	0x7f, 0x63, 0x8a, 0xde, 0x29, 0xf8, 0xfa, 0xe6, 0xfb, 0x69, 0xbe, 0x89, 0xb6, 0x0d, 0x86, 0x1e,
	0x2b, 0x1a, 0xc9, 0x74, 0x5b, 0x5a, 0x8d, 0xe4, 0x6f, 0x44, 0xda, 0x7d, 0xff, 0xf2, 0xf3, 0x97,
	0x82, 0x53, 0x06, 0xf0, 0x36, 0xe3, 0x26, 0xa1, 0xf1, 0x5c, 0x5a, 0x67, 0xcf, 0x2c, 0x8a, 0x76,
	0xda, 0x73, 0x1b, 0x1c, 0x15, 0x74, 0xf3, 0xbb, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x80, 0xcc,
	0xb3, 0xa5, 0xba, 0x06, 0x00, 0x00,
}
