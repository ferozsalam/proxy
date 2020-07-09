// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3/clusters.proto

package envoy_admin_v3

import (
	fmt "fmt"
	v31 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v3 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Admin endpoint uses this wrapper for `/clusters` to display cluster status information.
// See :ref:`/clusters <operations_admin_interface_clusters>` for more information.
type Clusters struct {
	// Mapping from cluster name to each cluster's status.
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{0}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

// Details an individual cluster's current status.
// [#next-free-field: 6]
type ClusterStatus struct {
	// Name of the cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Denotes whether this cluster was added via API or configured statically.
	AddedViaApi bool `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	// The success rate threshold used in the last interval.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is *false*, all errors: externally and locally generated were used to calculate the threshold.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is *true*, only externally generated errors were used to calculate the threshold.
	// The threshold is used to eject hosts based on their success rate. See
	// :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for details.
	//
	// Note: this field may be omitted in any of the three following cases:
	//
	// 1. There were not enough hosts with enough request volume to proceed with success rate based
	//    outlier ejection.
	// 2. The threshold is computed to be < 0 because a negative value implies that there was no
	//    threshold for that interval.
	// 3. Outlier detection is not enabled for this cluster.
	SuccessRateEjectionThreshold *v3.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	// Mapping from host address to the host's current status.
	HostStatuses []*HostStatus `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	// The success rate threshold used in the last interval when only locally originated failures were
	// taken into account and externally originated errors were treated as success.
	// This field should be interpreted only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is *true*. The threshold is used to eject hosts based on their success rate.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	//
	// Note: this field may be omitted in any of the three following cases:
	//
	// 1. There were not enough hosts with enough request volume to proceed with success rate based
	//    outlier ejection.
	// 2. The threshold is computed to be < 0 because a negative value implies that there was no
	//    threshold for that interval.
	// 3. Outlier detection is not enabled for this cluster.
	LocalOriginSuccessRateEjectionThreshold *v3.Percent `protobuf:"bytes,5,opt,name=local_origin_success_rate_ejection_threshold,json=localOriginSuccessRateEjectionThreshold,proto3" json:"local_origin_success_rate_ejection_threshold,omitempty"`
	XXX_NoUnkeyedLiteral                    struct{}    `json:"-"`
	XXX_unrecognized                        []byte      `json:"-"`
	XXX_sizecache                           int32       `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{1}
}

func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStatus.Unmarshal(m, b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterStatus.Size(m)
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetAddedViaApi() bool {
	if m != nil {
		return m.AddedViaApi
	}
	return false
}

func (m *ClusterStatus) GetSuccessRateEjectionThreshold() *v3.Percent {
	if m != nil {
		return m.SuccessRateEjectionThreshold
	}
	return nil
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

func (m *ClusterStatus) GetLocalOriginSuccessRateEjectionThreshold() *v3.Percent {
	if m != nil {
		return m.LocalOriginSuccessRateEjectionThreshold
	}
	return nil
}

// Current state of a particular host.
// [#next-free-field: 10]
type HostStatus struct {
	// Address of this host.
	Address *v31.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// List of stats specific to this host.
	Stats []*SimpleMetric `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	// The host's current health status.
	HealthStatus *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	// Request success rate for this host over the last calculated interval.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is *false*, all errors: externally and locally generated were used in success rate
	// calculation. If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is *true*, only externally generated errors were used in success rate calculation.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	//
	// Note: the message will not be present if host did not have enough request volume to calculate
	// success rate or the cluster did not have enough hosts to run through success rate outlier
	// ejection.
	SuccessRate *v3.Percent `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	// The host's weight. If not configured, the value defaults to 1.
	Weight uint32 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// The hostname of the host, if applicable.
	Hostname string `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The host's priority. If not configured, the value defaults to 0 (highest priority).
	Priority uint32 `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	// Request success rate for this host over the last calculated
	// interval when only locally originated errors are taken into account and externally originated
	// errors were treated as success.
	// This field should be interpreted only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is *true*.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	//
	// Note: the message will not be present if host did not have enough request volume to calculate
	// success rate or the cluster did not have enough hosts to run through success rate outlier
	// ejection.
	LocalOriginSuccessRate *v3.Percent `protobuf:"bytes,8,opt,name=local_origin_success_rate,json=localOriginSuccessRate,proto3" json:"local_origin_success_rate,omitempty"`
	// locality of the host.
	Locality             *v31.Locality `protobuf:"bytes,9,opt,name=locality,proto3" json:"locality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{2}
}

func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatus.Unmarshal(m, b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return xxx_messageInfo_HostStatus.Size(m)
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() *v31.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *HostStatus) GetStats() []*SimpleMetric {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *HostStatus) GetHealthStatus() *HostHealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return nil
}

func (m *HostStatus) GetSuccessRate() *v3.Percent {
	if m != nil {
		return m.SuccessRate
	}
	return nil
}

func (m *HostStatus) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *HostStatus) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostStatus) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *HostStatus) GetLocalOriginSuccessRate() *v3.Percent {
	if m != nil {
		return m.LocalOriginSuccessRate
	}
	return nil
}

func (m *HostStatus) GetLocality() *v31.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

// Health status for a host.
// [#next-free-field: 7]
type HostHealthStatus struct {
	// The host is currently failing active health checks.
	FailedActiveHealthCheck bool `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	// The host is currently considered an outlier and has been ejected.
	FailedOutlierCheck bool `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	// The host is currently being marked as degraded through active health checking.
	FailedActiveDegradedCheck bool `protobuf:"varint,4,opt,name=failed_active_degraded_check,json=failedActiveDegradedCheck,proto3" json:"failed_active_degraded_check,omitempty"`
	// The host has been removed from service discovery, but is being stabilized due to active
	// health checking.
	PendingDynamicRemoval bool `protobuf:"varint,5,opt,name=pending_dynamic_removal,json=pendingDynamicRemoval,proto3" json:"pending_dynamic_removal,omitempty"`
	// The host has not yet been health checked.
	PendingActiveHc bool `protobuf:"varint,6,opt,name=pending_active_hc,json=pendingActiveHc,proto3" json:"pending_active_hc,omitempty"`
	// Health status as reported by EDS. Note: only HEALTHY and UNHEALTHY are currently supported
	// here.
	// [#comment:TODO(mrice32): pipe through remaining EDS health status possibilities.]
	EdsHealthStatus      v31.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.config.core.v3.HealthStatus" json:"eds_health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HostHealthStatus) Reset()         { *m = HostHealthStatus{} }
func (m *HostHealthStatus) String() string { return proto.CompactTextString(m) }
func (*HostHealthStatus) ProtoMessage()    {}
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{3}
}

func (m *HostHealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostHealthStatus.Unmarshal(m, b)
}
func (m *HostHealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostHealthStatus.Marshal(b, m, deterministic)
}
func (m *HostHealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostHealthStatus.Merge(m, src)
}
func (m *HostHealthStatus) XXX_Size() int {
	return xxx_messageInfo_HostHealthStatus.Size(m)
}
func (m *HostHealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostHealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostHealthStatus proto.InternalMessageInfo

func (m *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if m != nil {
		return m.FailedActiveHealthCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedOutlierCheck() bool {
	if m != nil {
		return m.FailedOutlierCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedActiveDegradedCheck() bool {
	if m != nil {
		return m.FailedActiveDegradedCheck
	}
	return false
}

func (m *HostHealthStatus) GetPendingDynamicRemoval() bool {
	if m != nil {
		return m.PendingDynamicRemoval
	}
	return false
}

func (m *HostHealthStatus) GetPendingActiveHc() bool {
	if m != nil {
		return m.PendingActiveHc
	}
	return false
}

func (m *HostHealthStatus) GetEdsHealthStatus() v31.HealthStatus {
	if m != nil {
		return m.EdsHealthStatus
	}
	return v31.HealthStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*Clusters)(nil), "envoy.admin.v3.Clusters")
	proto.RegisterType((*ClusterStatus)(nil), "envoy.admin.v3.ClusterStatus")
	proto.RegisterType((*HostStatus)(nil), "envoy.admin.v3.HostStatus")
	proto.RegisterType((*HostHealthStatus)(nil), "envoy.admin.v3.HostHealthStatus")
}

func init() { proto.RegisterFile("envoy/admin/v3/clusters.proto", fileDescriptor_703c0abdc0a96402) }

var fileDescriptor_703c0abdc0a96402 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0x23, 0x35,
	0x18, 0x56, 0x3e, 0x36, 0x3b, 0xeb, 0x6c, 0xb6, 0x5d, 0x0b, 0xba, 0xb3, 0x21, 0xcd, 0x66, 0xb3,
	0x40, 0x23, 0x40, 0x09, 0x4a, 0x25, 0x2a, 0xc2, 0xa1, 0xea, 0x97, 0xd4, 0x03, 0xd0, 0x32, 0x45,
	0xdc, 0xd0, 0xc8, 0xf5, 0xbc, 0xcd, 0x18, 0x26, 0xe3, 0x91, 0xed, 0x04, 0x72, 0x83, 0x5b, 0x7f,
	0x03, 0x57, 0xfe, 0x05, 0x77, 0x24, 0xae, 0xfc, 0x23, 0xe4, 0x8f, 0x49, 0x26, 0x69, 0x1a, 0x6e,
	0xe3, 0x79, 0x9e, 0xe7, 0xf5, 0xeb, 0xe7, 0x79, 0x6d, 0xb4, 0x0f, 0xe9, 0x8c, 0xcf, 0x07, 0x24,
	0x9a, 0xb0, 0x74, 0x30, 0x3b, 0x1c, 0xd0, 0x64, 0x2a, 0x15, 0x08, 0xd9, 0xcf, 0x04, 0x57, 0x1c,
	0xbf, 0x30, 0x70, 0xdf, 0xc0, 0xfd, 0xd9, 0x61, 0xb3, 0xb5, 0x46, 0x9f, 0x80, 0x12, 0x8c, 0x3a,
	0x76, 0xb3, 0x6b, 0x51, 0xca, 0xd3, 0x3b, 0x36, 0x1e, 0x50, 0x2e, 0x40, 0x73, 0x48, 0x14, 0x09,
	0x90, 0x39, 0xe7, 0xcd, 0x46, 0xce, 0x2d, 0x91, 0xe0, 0x08, 0x07, 0x1b, 0x09, 0x31, 0x90, 0x44,
	0xc5, 0x21, 0x8d, 0x81, 0xfe, 0xec, 0x88, 0x1f, 0x58, 0xa2, 0x9a, 0x67, 0x86, 0x91, 0x81, 0xa0,
	0x90, 0x2a, 0x07, 0xee, 0x4f, 0xa3, 0x8c, 0x0c, 0x48, 0x9a, 0x72, 0x45, 0x14, 0xe3, 0xa9, 0x1c,
	0x48, 0x45, 0xd4, 0x34, 0xef, 0xe2, 0xed, 0x03, 0x78, 0x06, 0x42, 0x32, 0x9e, 0xb2, 0x74, 0x6c,
	0x29, 0xdd, 0x39, 0xf2, 0xce, 0x9c, 0x19, 0xf8, 0x12, 0xed, 0x3a, 0x63, 0x42, 0x5b, 0x06, 0xa4,
	0x5f, 0xea, 0x54, 0x7a, 0xf5, 0xe1, 0x7e, 0x7f, 0xd5, 0xa1, 0xbe, 0xd3, 0xdc, 0x18, 0x5a, 0xb0,
	0x43, 0x8b, 0x4b, 0x90, 0xa3, 0x77, 0x7f, 0xfc, 0x7d, 0xdf, 0x6e, 0xa3, 0xd6, 0x8a, 0x6a, 0x48,
	0x92, 0x2c, 0x26, 0xb9, 0x54, 0x76, 0xef, 0x2b, 0xa8, 0xb1, 0x52, 0x07, 0x63, 0x54, 0x4d, 0xc9,
	0x04, 0xfc, 0x52, 0xa7, 0xd4, 0x7b, 0x16, 0x98, 0x6f, 0xdc, 0x45, 0x0d, 0x12, 0x45, 0x10, 0x85,
	0x33, 0x46, 0x42, 0x92, 0x31, 0xbf, 0xdc, 0x29, 0xf5, 0xbc, 0xa0, 0x6e, 0x7e, 0xfe, 0xc0, 0xc8,
	0x49, 0xc6, 0xf0, 0x8f, 0xe8, 0x8d, 0x9c, 0x52, 0x0a, 0x52, 0x86, 0x82, 0x28, 0x08, 0xe1, 0x27,
	0xa0, 0xfa, 0xbc, 0xa1, 0x8a, 0x05, 0xc8, 0x98, 0x27, 0x91, 0x5f, 0xe9, 0x94, 0x7a, 0xf5, 0xe1,
	0x9e, 0x3b, 0x87, 0x76, 0x53, 0x1f, 0xe3, 0xda, 0xba, 0x19, 0xb4, 0x9c, 0x3c, 0x20, 0x0a, 0x2e,
	0x9c, 0xf8, 0xfb, 0x5c, 0x8b, 0x8f, 0x51, 0x23, 0xe6, 0x52, 0x2d, 0x4d, 0xa9, 0x1a, 0x53, 0x9a,
	0xeb, 0xa6, 0x5c, 0x72, 0xa9, 0x9c, 0x23, 0xcf, 0xe3, 0xc5, 0x37, 0x48, 0x3c, 0x45, 0x9f, 0x25,
	0x9c, 0x92, 0x24, 0xe4, 0x82, 0x8d, 0x59, 0x1a, 0xfe, 0x5f, 0xb3, 0x4f, 0xb6, 0x36, 0x7b, 0x60,
	0x6a, 0x5d, 0x99, 0x52, 0x37, 0x5b, 0xfa, 0x1e, 0xf5, 0x74, 0x0a, 0xef, 0xd0, 0xdb, 0x2d, 0x29,
	0xd8, 0x16, 0xbb, 0xbf, 0x57, 0x11, 0x5a, 0x76, 0x8f, 0x8f, 0xd0, 0x53, 0x37, 0xce, 0x26, 0x8a,
	0x65, 0xfe, 0x76, 0x5c, 0xfb, 0x7a, 0x5c, 0x75, 0x47, 0x27, 0x96, 0x14, 0xe4, 0x6c, 0x3c, 0x44,
	0x4f, 0xb4, 0x49, 0xd2, 0x2f, 0x1b, 0x87, 0x5a, 0xeb, 0x0e, 0xdd, 0xb0, 0x49, 0x96, 0xc0, 0x37,
	0xe6, 0x3a, 0x05, 0x96, 0x8a, 0x2f, 0x50, 0xc3, 0x8d, 0xbd, 0xf5, 0xd7, 0x45, 0xd5, 0xd9, 0xe4,
	0xee, 0xa5, 0x21, 0x2e, 0x3c, 0x2e, 0xac, 0xf0, 0x97, 0xe8, 0x79, 0xd1, 0x56, 0xbf, 0xba, 0xd5,
	0xc3, 0x7a, 0x21, 0x70, 0xbc, 0x87, 0x6a, 0xbf, 0x00, 0x1b, 0xc7, 0xca, 0x18, 0xdf, 0x08, 0xdc,
	0x0a, 0x37, 0x91, 0xa7, 0x63, 0x34, 0x23, 0x59, 0x33, 0x23, 0xb9, 0x58, 0x6b, 0x2c, 0x13, 0x8c,
	0x0b, 0xa6, 0xe6, 0xfe, 0x53, 0xa3, 0x5a, 0xac, 0xf1, 0x77, 0xe8, 0xf5, 0xa3, 0x71, 0xfb, 0xde,
	0xd6, 0xbe, 0xf6, 0x36, 0x67, 0x8b, 0x47, 0xc8, 0x33, 0x88, 0xde, 0xee, 0x99, 0xa9, 0xd0, 0xde,
	0x1c, 0xc9, 0xd7, 0x8e, 0x15, 0x2c, 0xf8, 0xa3, 0x8f, 0xf4, 0x18, 0x74, 0x50, 0x7b, 0xd3, 0x18,
	0x2c, 0x43, 0xef, 0xfe, 0x59, 0x41, 0xbb, 0xeb, 0x1e, 0xe3, 0xaf, 0x50, 0xf3, 0x8e, 0xb0, 0x04,
	0xa2, 0x90, 0x50, 0xc5, 0x66, 0x10, 0x16, 0x5f, 0x28, 0x33, 0x1c, 0x5e, 0xf0, 0xca, 0x32, 0x4e,
	0x0c, 0xc1, 0xaa, 0xcf, 0x34, 0x8c, 0x3f, 0x47, 0xef, 0x39, 0x31, 0x9f, 0xaa, 0x84, 0x81, 0x70,
	0x32, 0x7b, 0x83, 0xb1, 0xc5, 0xae, 0x2c, 0x64, 0x15, 0xc7, 0xa8, 0xb5, 0xba, 0x5d, 0x04, 0x63,
	0x41, 0xf4, 0xed, 0xb7, 0xca, 0xaa, 0x51, 0xbe, 0x2e, 0x6e, 0x78, 0xee, 0x18, 0xb6, 0xc0, 0x17,
	0xe8, 0x55, 0x06, 0x69, 0xc4, 0xd2, 0x71, 0x18, 0xcd, 0x53, 0x32, 0x61, 0x34, 0x14, 0x30, 0xe1,
	0x33, 0x92, 0x98, 0x6c, 0xbd, 0xe0, 0x7d, 0x07, 0x9f, 0x5b, 0x34, 0xb0, 0x20, 0xfe, 0x04, 0xbd,
	0xcc, 0x75, 0xf9, 0x41, 0xa9, 0xc9, 0xdc, 0x0b, 0x76, 0x1c, 0xe0, 0xce, 0x47, 0xf1, 0xb7, 0xe8,
	0x25, 0x44, 0x32, 0x7c, 0x38, 0xb4, 0x2f, 0x86, 0xdd, 0xcd, 0xa1, 0xac, 0x8c, 0xed, 0x0e, 0x44,
	0xb2, 0xf8, 0x63, 0xf4, 0xa9, 0xce, 0xe7, 0x63, 0xf4, 0xe1, 0x63, 0xf9, 0x14, 0xc9, 0xa7, 0x47,
	0x7f, 0xfd, 0xf6, 0xcf, 0xbf, 0xb5, 0xf2, 0x6e, 0x19, 0xb5, 0x18, 0xb7, 0xbb, 0x65, 0x82, 0xff,
	0x3a, 0x5f, 0xbb, 0x2d, 0xa7, 0xf9, 0xcb, 0x2a, 0xaf, 0xf5, 0x33, 0x7f, 0x5d, 0xba, 0xad, 0x99,
	0xf7, 0xfe, 0xf0, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x11, 0xa5, 0x0c, 0x0b, 0x07, 0x00,
	0x00,
}
