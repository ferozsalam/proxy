// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint/load_report.proto

package envoy_api_v2_endpoint

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// These are stats Envoy reports to GLB every so often. Report frequency is
// defined by
// :ref:`LoadStatsResponse.load_reporting_interval<envoy_api_field_service.load_stats.v2.LoadStatsResponse.load_reporting_interval>`.
// Stats per upstream region/zone and optionally per subzone.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
// [#next-free-field: 9]
type UpstreamLocalityStats struct {
	// Name of zone, region and optionally endpoint group these metrics were
	// collected from. Zone and region names could be empty if unknown.
	Locality *core.Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The total number of requests successfully completed by the endpoints in the
	// locality.
	TotalSuccessfulRequests uint64 `protobuf:"varint,2,opt,name=total_successful_requests,json=totalSuccessfulRequests,proto3" json:"total_successful_requests,omitempty"`
	// The total number of unfinished requests
	TotalRequestsInProgress uint64 `protobuf:"varint,3,opt,name=total_requests_in_progress,json=totalRequestsInProgress,proto3" json:"total_requests_in_progress,omitempty"`
	// The total number of requests that failed due to errors at the endpoint,
	// aggregated over all endpoints in the locality.
	TotalErrorRequests uint64 `protobuf:"varint,4,opt,name=total_error_requests,json=totalErrorRequests,proto3" json:"total_error_requests,omitempty"`
	// The total number of requests that were issued by this Envoy since
	// the last report. This information is aggregated over all the
	// upstream endpoints in the locality.
	TotalIssuedRequests uint64 `protobuf:"varint,8,opt,name=total_issued_requests,json=totalIssuedRequests,proto3" json:"total_issued_requests,omitempty"`
	// Stats for multi-dimensional load balancing.
	LoadMetricStats []*EndpointLoadMetricStats `protobuf:"bytes,5,rep,name=load_metric_stats,json=loadMetricStats,proto3" json:"load_metric_stats,omitempty"`
	// Endpoint granularity stats information for this locality. This information
	// is populated if the Server requests it by setting
	// :ref:`LoadStatsResponse.report_endpoint_granularity<envoy_api_field_service.load_stats.v2.LoadStatsResponse.report_endpoint_granularity>`.
	UpstreamEndpointStats []*UpstreamEndpointStats `protobuf:"bytes,7,rep,name=upstream_endpoint_stats,json=upstreamEndpointStats,proto3" json:"upstream_endpoint_stats,omitempty"`
	// [#not-implemented-hide:] The priority of the endpoint group these metrics
	// were collected from.
	Priority             uint32   `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamLocalityStats) Reset()         { *m = UpstreamLocalityStats{} }
func (m *UpstreamLocalityStats) String() string { return proto.CompactTextString(m) }
func (*UpstreamLocalityStats) ProtoMessage()    {}
func (*UpstreamLocalityStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5134f8f33d8f8d01, []int{0}
}

func (m *UpstreamLocalityStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamLocalityStats.Unmarshal(m, b)
}
func (m *UpstreamLocalityStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamLocalityStats.Marshal(b, m, deterministic)
}
func (m *UpstreamLocalityStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamLocalityStats.Merge(m, src)
}
func (m *UpstreamLocalityStats) XXX_Size() int {
	return xxx_messageInfo_UpstreamLocalityStats.Size(m)
}
func (m *UpstreamLocalityStats) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamLocalityStats.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamLocalityStats proto.InternalMessageInfo

func (m *UpstreamLocalityStats) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *UpstreamLocalityStats) GetTotalSuccessfulRequests() uint64 {
	if m != nil {
		return m.TotalSuccessfulRequests
	}
	return 0
}

func (m *UpstreamLocalityStats) GetTotalRequestsInProgress() uint64 {
	if m != nil {
		return m.TotalRequestsInProgress
	}
	return 0
}

func (m *UpstreamLocalityStats) GetTotalErrorRequests() uint64 {
	if m != nil {
		return m.TotalErrorRequests
	}
	return 0
}

func (m *UpstreamLocalityStats) GetTotalIssuedRequests() uint64 {
	if m != nil {
		return m.TotalIssuedRequests
	}
	return 0
}

func (m *UpstreamLocalityStats) GetLoadMetricStats() []*EndpointLoadMetricStats {
	if m != nil {
		return m.LoadMetricStats
	}
	return nil
}

func (m *UpstreamLocalityStats) GetUpstreamEndpointStats() []*UpstreamEndpointStats {
	if m != nil {
		return m.UpstreamEndpointStats
	}
	return nil
}

func (m *UpstreamLocalityStats) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
// [#next-free-field: 8]
type UpstreamEndpointStats struct {
	// Upstream host address.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Opaque and implementation dependent metadata of the
	// endpoint. Envoy will pass this directly to the management server.
	Metadata *_struct.Struct `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The total number of requests successfully completed by the endpoints in the
	// locality. These include non-5xx responses for HTTP, where errors
	// originate at the client and the endpoint responded successfully. For gRPC,
	// the grpc-status values are those not covered by total_error_requests below.
	TotalSuccessfulRequests uint64 `protobuf:"varint,2,opt,name=total_successful_requests,json=totalSuccessfulRequests,proto3" json:"total_successful_requests,omitempty"`
	// The total number of unfinished requests for this endpoint.
	TotalRequestsInProgress uint64 `protobuf:"varint,3,opt,name=total_requests_in_progress,json=totalRequestsInProgress,proto3" json:"total_requests_in_progress,omitempty"`
	// The total number of requests that failed due to errors at the endpoint.
	// For HTTP these are responses with 5xx status codes and for gRPC the
	// grpc-status values:
	//
	//   - DeadlineExceeded
	//   - Unimplemented
	//   - Internal
	//   - Unavailable
	//   - Unknown
	//   - DataLoss
	TotalErrorRequests uint64 `protobuf:"varint,4,opt,name=total_error_requests,json=totalErrorRequests,proto3" json:"total_error_requests,omitempty"`
	// The total number of requests that were issued to this endpoint
	// since the last report. A single TCP connection, HTTP or gRPC
	// request or stream is counted as one request.
	TotalIssuedRequests uint64 `protobuf:"varint,7,opt,name=total_issued_requests,json=totalIssuedRequests,proto3" json:"total_issued_requests,omitempty"`
	// Stats for multi-dimensional load balancing.
	LoadMetricStats      []*EndpointLoadMetricStats `protobuf:"bytes,5,rep,name=load_metric_stats,json=loadMetricStats,proto3" json:"load_metric_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UpstreamEndpointStats) Reset()         { *m = UpstreamEndpointStats{} }
func (m *UpstreamEndpointStats) String() string { return proto.CompactTextString(m) }
func (*UpstreamEndpointStats) ProtoMessage()    {}
func (*UpstreamEndpointStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5134f8f33d8f8d01, []int{1}
}

func (m *UpstreamEndpointStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamEndpointStats.Unmarshal(m, b)
}
func (m *UpstreamEndpointStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamEndpointStats.Marshal(b, m, deterministic)
}
func (m *UpstreamEndpointStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamEndpointStats.Merge(m, src)
}
func (m *UpstreamEndpointStats) XXX_Size() int {
	return xxx_messageInfo_UpstreamEndpointStats.Size(m)
}
func (m *UpstreamEndpointStats) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamEndpointStats.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamEndpointStats proto.InternalMessageInfo

func (m *UpstreamEndpointStats) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UpstreamEndpointStats) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UpstreamEndpointStats) GetTotalSuccessfulRequests() uint64 {
	if m != nil {
		return m.TotalSuccessfulRequests
	}
	return 0
}

func (m *UpstreamEndpointStats) GetTotalRequestsInProgress() uint64 {
	if m != nil {
		return m.TotalRequestsInProgress
	}
	return 0
}

func (m *UpstreamEndpointStats) GetTotalErrorRequests() uint64 {
	if m != nil {
		return m.TotalErrorRequests
	}
	return 0
}

func (m *UpstreamEndpointStats) GetTotalIssuedRequests() uint64 {
	if m != nil {
		return m.TotalIssuedRequests
	}
	return 0
}

func (m *UpstreamEndpointStats) GetLoadMetricStats() []*EndpointLoadMetricStats {
	if m != nil {
		return m.LoadMetricStats
	}
	return nil
}

// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type EndpointLoadMetricStats struct {
	// Name of the metric; may be empty.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// Number of calls that finished and included this metric.
	NumRequestsFinishedWithMetric uint64 `protobuf:"varint,2,opt,name=num_requests_finished_with_metric,json=numRequestsFinishedWithMetric,proto3" json:"num_requests_finished_with_metric,omitempty"`
	// Sum of metric values across all calls that finished with this metric for
	// load_reporting_interval.
	TotalMetricValue     float64  `protobuf:"fixed64,3,opt,name=total_metric_value,json=totalMetricValue,proto3" json:"total_metric_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointLoadMetricStats) Reset()         { *m = EndpointLoadMetricStats{} }
func (m *EndpointLoadMetricStats) String() string { return proto.CompactTextString(m) }
func (*EndpointLoadMetricStats) ProtoMessage()    {}
func (*EndpointLoadMetricStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5134f8f33d8f8d01, []int{2}
}

func (m *EndpointLoadMetricStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointLoadMetricStats.Unmarshal(m, b)
}
func (m *EndpointLoadMetricStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointLoadMetricStats.Marshal(b, m, deterministic)
}
func (m *EndpointLoadMetricStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointLoadMetricStats.Merge(m, src)
}
func (m *EndpointLoadMetricStats) XXX_Size() int {
	return xxx_messageInfo_EndpointLoadMetricStats.Size(m)
}
func (m *EndpointLoadMetricStats) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointLoadMetricStats.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointLoadMetricStats proto.InternalMessageInfo

func (m *EndpointLoadMetricStats) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *EndpointLoadMetricStats) GetNumRequestsFinishedWithMetric() uint64 {
	if m != nil {
		return m.NumRequestsFinishedWithMetric
	}
	return 0
}

func (m *EndpointLoadMetricStats) GetTotalMetricValue() float64 {
	if m != nil {
		return m.TotalMetricValue
	}
	return 0
}

// Per cluster load stats. Envoy reports these stats a management server in a
// :ref:`LoadStatsRequest<envoy_api_msg_service.load_stats.v2.LoadStatsRequest>`
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
// Next ID: 7
// [#next-free-field: 7]
type ClusterStats struct {
	// The name of the cluster.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// The eds_cluster_config service_name of the cluster.
	// It's possible that two clusters send the same service_name to EDS,
	// in that case, the management server is supposed to do aggregation on the load reports.
	ClusterServiceName string `protobuf:"bytes,6,opt,name=cluster_service_name,json=clusterServiceName,proto3" json:"cluster_service_name,omitempty"`
	// Need at least one.
	UpstreamLocalityStats []*UpstreamLocalityStats `protobuf:"bytes,2,rep,name=upstream_locality_stats,json=upstreamLocalityStats,proto3" json:"upstream_locality_stats,omitempty"`
	// Cluster-level stats such as total_successful_requests may be computed by
	// summing upstream_locality_stats. In addition, below there are additional
	// cluster-wide stats.
	//
	// The total number of dropped requests. This covers requests
	// deliberately dropped by the drop_overload policy and circuit breaking.
	TotalDroppedRequests uint64 `protobuf:"varint,3,opt,name=total_dropped_requests,json=totalDroppedRequests,proto3" json:"total_dropped_requests,omitempty"`
	// Information about deliberately dropped requests for each category specified
	// in the DropOverload policy.
	DroppedRequests []*ClusterStats_DroppedRequests `protobuf:"bytes,5,rep,name=dropped_requests,json=droppedRequests,proto3" json:"dropped_requests,omitempty"`
	// Period over which the actual load report occurred. This will be guaranteed to include every
	// request reported. Due to system load and delays between the *LoadStatsRequest* sent from Envoy
	// and the *LoadStatsResponse* message sent from the management server, this may be longer than
	// the requested load reporting interval in the *LoadStatsResponse*.
	LoadReportInterval   *duration.Duration `protobuf:"bytes,4,opt,name=load_report_interval,json=loadReportInterval,proto3" json:"load_report_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClusterStats) Reset()         { *m = ClusterStats{} }
func (m *ClusterStats) String() string { return proto.CompactTextString(m) }
func (*ClusterStats) ProtoMessage()    {}
func (*ClusterStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5134f8f33d8f8d01, []int{3}
}

func (m *ClusterStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStats.Unmarshal(m, b)
}
func (m *ClusterStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStats.Marshal(b, m, deterministic)
}
func (m *ClusterStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStats.Merge(m, src)
}
func (m *ClusterStats) XXX_Size() int {
	return xxx_messageInfo_ClusterStats.Size(m)
}
func (m *ClusterStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStats.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStats proto.InternalMessageInfo

func (m *ClusterStats) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterStats) GetClusterServiceName() string {
	if m != nil {
		return m.ClusterServiceName
	}
	return ""
}

func (m *ClusterStats) GetUpstreamLocalityStats() []*UpstreamLocalityStats {
	if m != nil {
		return m.UpstreamLocalityStats
	}
	return nil
}

func (m *ClusterStats) GetTotalDroppedRequests() uint64 {
	if m != nil {
		return m.TotalDroppedRequests
	}
	return 0
}

func (m *ClusterStats) GetDroppedRequests() []*ClusterStats_DroppedRequests {
	if m != nil {
		return m.DroppedRequests
	}
	return nil
}

func (m *ClusterStats) GetLoadReportInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportInterval
	}
	return nil
}

type ClusterStats_DroppedRequests struct {
	// Identifier for the policy specifying the drop.
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// Total number of deliberately dropped requests for the category.
	DroppedCount         uint64   `protobuf:"varint,2,opt,name=dropped_count,json=droppedCount,proto3" json:"dropped_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterStats_DroppedRequests) Reset()         { *m = ClusterStats_DroppedRequests{} }
func (m *ClusterStats_DroppedRequests) String() string { return proto.CompactTextString(m) }
func (*ClusterStats_DroppedRequests) ProtoMessage()    {}
func (*ClusterStats_DroppedRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_5134f8f33d8f8d01, []int{3, 0}
}

func (m *ClusterStats_DroppedRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStats_DroppedRequests.Unmarshal(m, b)
}
func (m *ClusterStats_DroppedRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStats_DroppedRequests.Marshal(b, m, deterministic)
}
func (m *ClusterStats_DroppedRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStats_DroppedRequests.Merge(m, src)
}
func (m *ClusterStats_DroppedRequests) XXX_Size() int {
	return xxx_messageInfo_ClusterStats_DroppedRequests.Size(m)
}
func (m *ClusterStats_DroppedRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStats_DroppedRequests.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStats_DroppedRequests proto.InternalMessageInfo

func (m *ClusterStats_DroppedRequests) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterStats_DroppedRequests) GetDroppedCount() uint64 {
	if m != nil {
		return m.DroppedCount
	}
	return 0
}

func init() {
	proto.RegisterType((*UpstreamLocalityStats)(nil), "envoy.api.v2.endpoint.UpstreamLocalityStats")
	proto.RegisterType((*UpstreamEndpointStats)(nil), "envoy.api.v2.endpoint.UpstreamEndpointStats")
	proto.RegisterType((*EndpointLoadMetricStats)(nil), "envoy.api.v2.endpoint.EndpointLoadMetricStats")
	proto.RegisterType((*ClusterStats)(nil), "envoy.api.v2.endpoint.ClusterStats")
	proto.RegisterType((*ClusterStats_DroppedRequests)(nil), "envoy.api.v2.endpoint.ClusterStats.DroppedRequests")
}

func init() {
	proto.RegisterFile("envoy/api/v2/endpoint/load_report.proto", fileDescriptor_5134f8f33d8f8d01)
}

var fileDescriptor_5134f8f33d8f8d01 = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0x4d, 0x8b, 0x23, 0x45,
	0x18, 0xa6, 0x33, 0x99, 0x24, 0x56, 0x66, 0xc9, 0x58, 0x4e, 0x4c, 0x36, 0xba, 0x6e, 0xcc, 0x1c,
	0x0c, 0xb2, 0x74, 0x2f, 0xc9, 0x82, 0xa0, 0x27, 0xb3, 0x3b, 0xe2, 0xe0, 0x28, 0x43, 0x07, 0x15,
	0x14, 0x6c, 0x6b, 0xba, 0x2b, 0x99, 0x82, 0xee, 0xaa, 0xb6, 0x3e, 0x5a, 0x73, 0xf3, 0x0f, 0x88,
	0xe0, 0xc9, 0x83, 0xbf, 0xc2, 0xa3, 0xe0, 0xdd, 0xab, 0x7f, 0xc5, 0xe3, 0x1c, 0x44, 0xba, 0x3e,
	0x3a, 0x9f, 0x03, 0xde, 0x04, 0x6f, 0xdd, 0xef, 0xf3, 0x3c, 0xf5, 0x7e, 0x3d, 0xd5, 0x0d, 0xde,
	0xc2, 0xb4, 0x60, 0xab, 0x00, 0xe5, 0x24, 0x28, 0x26, 0x01, 0xa6, 0x49, 0xce, 0x08, 0x95, 0x41,
	0xca, 0x50, 0x12, 0x71, 0x9c, 0x33, 0x2e, 0xfd, 0x9c, 0x33, 0xc9, 0x60, 0x57, 0x13, 0x7d, 0x94,
	0x13, 0xbf, 0x98, 0xf8, 0x8e, 0x38, 0x78, 0xbc, 0xa5, 0x8f, 0x19, 0xc7, 0x01, 0x4a, 0x12, 0x8e,
	0x85, 0x30, 0xba, 0xc1, 0xeb, 0xfb, 0x84, 0x1b, 0x24, 0xb0, 0x45, 0xdf, 0x58, 0x32, 0xb6, 0x4c,
	0x71, 0xa0, 0xdf, 0x6e, 0xd4, 0x22, 0x48, 0x14, 0x47, 0x92, 0x30, 0xea, 0xd4, 0xbb, 0xb8, 0x90,
	0x5c, 0xc5, 0xd2, 0xa9, 0x55, 0x92, 0xa3, 0x00, 0x51, 0xca, 0xa4, 0x16, 0x89, 0x20, 0x23, 0x4b,
	0x8e, 0xa4, 0x3b, 0xfd, 0xd1, 0x1e, 0x2e, 0x24, 0x92, 0xca, 0x95, 0xd6, 0x2b, 0x50, 0x4a, 0x12,
	0x24, 0x71, 0xe0, 0x1e, 0x0c, 0x30, 0xfa, 0xa1, 0x0e, 0xba, 0x9f, 0xe6, 0x42, 0x72, 0x8c, 0xb2,
	0x2b, 0x16, 0xa3, 0x94, 0xc8, 0xd5, 0x5c, 0x22, 0x29, 0xe0, 0x3b, 0xa0, 0x95, 0xda, 0x40, 0xdf,
	0x1b, 0x7a, 0xe3, 0xf6, 0xe4, 0x35, 0x7f, 0x6b, 0x30, 0x65, 0x83, 0xbe, 0xd3, 0x84, 0x15, 0x19,
	0xbe, 0x0b, 0x1e, 0x4a, 0x26, 0x51, 0x1a, 0x09, 0x15, 0xc7, 0x58, 0x88, 0x85, 0x4a, 0x23, 0x8e,
	0xbf, 0x51, 0x58, 0x48, 0xd1, 0xaf, 0x0d, 0xbd, 0x71, 0x3d, 0xec, 0x69, 0xc2, 0xbc, 0xc2, 0x43,
	0x0b, 0xc3, 0xf7, 0xc0, 0xc0, 0x68, 0x9d, 0x20, 0x22, 0x34, 0xca, 0x39, 0x5b, 0x96, 0x63, 0xee,
	0x1f, 0x6d, 0x88, 0x9d, 0xe4, 0x92, 0x5e, 0x5b, 0x18, 0x3e, 0x05, 0x67, 0x46, 0x8c, 0x39, 0x67,
	0x7c, 0x9d, 0xb3, 0xae, 0x65, 0x50, 0x63, 0x17, 0x25, 0x54, 0xa5, 0x9b, 0x80, 0xae, 0x51, 0x10,
	0x21, 0x14, 0x4e, 0xd6, 0x92, 0x96, 0x96, 0xbc, 0xa2, 0xc1, 0x4b, 0x8d, 0x55, 0x9a, 0x2f, 0xc0,
	0xcb, 0xda, 0x32, 0x19, 0x96, 0x9c, 0xc4, 0x51, 0x39, 0x66, 0xd1, 0x3f, 0x1e, 0x1e, 0x8d, 0xdb,
	0x13, 0xdf, 0x3f, 0xe8, 0x1c, 0xff, 0xc2, 0x3e, 0x5c, 0x31, 0x94, 0x7c, 0xac, 0x65, 0x7a, 0xc4,
	0x61, 0x27, 0xdd, 0x0e, 0xc0, 0x04, 0xf4, 0x94, 0x5d, 0x46, 0xe4, 0xd4, 0x36, 0x43, 0x53, 0x67,
	0x78, 0x72, 0x4f, 0x06, 0xb7, 0x42, 0x97, 0xc9, 0x9c, 0xdf, 0x55, 0x87, 0xc2, 0x70, 0x00, 0x5a,
	0x39, 0x27, 0x8c, 0x97, 0x9b, 0x6d, 0x0c, 0xbd, 0xf1, 0x83, 0xb0, 0x7a, 0x1f, 0xfd, 0x7e, 0xb4,
	0xf6, 0xc3, 0xb6, 0xea, 0x19, 0x68, 0x5a, 0xbb, 0x5b, 0x3b, 0x0c, 0x0e, 0xd8, 0xe1, 0x7d, 0xc3,
	0x08, 0x1d, 0x15, 0x4e, 0x41, 0x2b, 0xc3, 0x12, 0x25, 0x48, 0x22, 0x9d, 0xab, 0x3d, 0xe9, 0xf9,
	0xc6, 0xe8, 0xbe, 0x33, 0xba, 0x3f, 0xd7, 0x46, 0x0f, 0x2b, 0xe2, 0xff, 0xc2, 0x41, 0xcd, 0xff,
	0xc4, 0x41, 0xa3, 0x5f, 0x3d, 0xd0, 0xbb, 0x87, 0x0c, 0x1f, 0x83, 0xb6, 0x4d, 0x49, 0x51, 0x86,
	0xf5, 0x16, 0x5f, 0x0a, 0x81, 0x09, 0x7d, 0x82, 0x32, 0x0c, 0x3f, 0x04, 0x6f, 0x52, 0x95, 0xad,
	0x27, 0xb7, 0x20, 0x94, 0x88, 0x5b, 0x9c, 0x44, 0xdf, 0x12, 0x79, 0x6b, 0xcb, 0xb5, 0xf3, 0x7f,
	0x44, 0x55, 0xe6, 0x1a, 0xfa, 0xc0, 0xd2, 0x3e, 0x27, 0xf2, 0xd6, 0xe4, 0x83, 0x4f, 0x80, 0x19,
	0x96, 0xeb, 0xb1, 0x40, 0xa9, 0xc2, 0x7a, 0xfa, 0x5e, 0x78, 0xaa, 0x11, 0x43, 0xfc, 0xac, 0x8c,
	0x8f, 0x7e, 0xa9, 0x83, 0x93, 0xe7, 0xa9, 0x12, 0x12, 0x73, 0x53, 0xe9, 0xdb, 0xe0, 0x24, 0x36,
	0xef, 0x1b, 0xa5, 0xce, 0x9a, 0x77, 0xb3, 0x3a, 0xaf, 0x0d, 0xbd, 0xb0, 0x6d, 0x41, 0x5d, 0xf4,
	0x53, 0x70, 0xe6, 0xb8, 0x02, 0xf3, 0x82, 0xc4, 0xd8, 0x68, 0x1a, 0xba, 0x3d, 0x68, 0xb1, 0xb9,
	0x81, 0xb4, 0x82, 0x6e, 0xdc, 0x32, 0xf7, 0xd5, 0xb2, 0x5b, 0xa8, 0xfd, 0xab, 0x5b, 0xb6, 0xf5,
	0xa1, 0x9c, 0xb5, 0xee, 0x66, 0xc7, 0x3f, 0x79, 0xb5, 0x96, 0xb7, 0xbe, 0x6f, 0xdb, 0x5f, 0xd2,
	0x67, 0xe0, 0x55, 0x33, 0x8c, 0x84, 0xb3, 0x3c, 0xdf, 0x34, 0x89, 0xb1, 0xa3, 0xf1, 0xdc, 0x0b,
	0x03, 0x56, 0x2e, 0xf9, 0x0a, 0x9c, 0xee, 0xf1, 0x8d, 0x49, 0xa6, 0xf7, 0x94, 0xb7, 0x39, 0x42,
	0x7f, 0xe7, 0xb8, 0xb0, 0x93, 0xec, 0x9c, 0xff, 0x11, 0x38, 0xdb, 0xf8, 0xf5, 0x45, 0x84, 0x4a,
	0xcc, 0x0b, 0x94, 0x6a, 0xaf, 0xb7, 0x27, 0x0f, 0xf7, 0x6e, 0xe9, 0x0b, 0xfb, 0xbb, 0x0a, 0x61,
	0x29, 0x0b, 0xb5, 0xea, 0xd2, 0x8a, 0x06, 0x5f, 0x82, 0xce, 0x6e, 0xfd, 0xe7, 0xa0, 0x15, 0x23,
	0x89, 0x97, 0x8c, 0xaf, 0x76, 0xf7, 0x57, 0x01, 0xf0, 0x1c, 0x3c, 0x70, 0x4d, 0xc6, 0x4c, 0x51,
	0x69, 0xdd, 0x75, 0x62, 0x83, 0xcf, 0xcb, 0xd8, 0xec, 0xeb, 0xbf, 0x7e, 0xfe, 0xfb, 0xc7, 0xe3,
	0x01, 0xec, 0x9b, 0xb6, 0x63, 0x46, 0x17, 0x64, 0xb9, 0x6e, 0xbb, 0x98, 0xfe, 0xf6, 0xfd, 0x1f,
	0x7f, 0x36, 0x6a, 0xa7, 0x1e, 0x38, 0x27, 0xcc, 0xcc, 0x26, 0xe7, 0xec, 0xbb, 0xd5, 0xe1, 0x31,
	0xcd, 0x3a, 0x57, 0x55, 0xfd, 0xd7, 0x65, 0x73, 0xd7, 0xde, 0x4d, 0x43, 0x77, 0x39, 0xfd, 0x27,
	0x00, 0x00, 0xff, 0xff, 0xd3, 0xcb, 0x4d, 0x04, 0x23, 0x08, 0x00, 0x00,
}
