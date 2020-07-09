// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/memory.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
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

// Proto representation of the internal memory consumption of an Envoy instance. These represent
// values extracted from an internal TCMalloc instance. For more information, see the section of the
// docs entitled ["Generic Tcmalloc Status"](https://gperftools.github.io/gperftools/tcmalloc.html).
// [#next-free-field: 7]
type Memory struct {
	// The number of bytes allocated by the heap for Envoy. This is an alias for
	// `generic.current_allocated_bytes`.
	Allocated uint64 `protobuf:"varint,1,opt,name=allocated,proto3" json:"allocated,omitempty"`
	// The number of bytes reserved by the heap but not necessarily allocated. This is an alias for
	// `generic.heap_size`.
	HeapSize uint64 `protobuf:"varint,2,opt,name=heap_size,json=heapSize,proto3" json:"heap_size,omitempty"`
	// The number of bytes in free, unmapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and depending on the OS, typically do not count towards physical memory
	// usage. This is an alias for `tcmalloc.pageheap_unmapped_bytes`.
	PageheapUnmapped uint64 `protobuf:"varint,3,opt,name=pageheap_unmapped,json=pageheapUnmapped,proto3" json:"pageheap_unmapped,omitempty"`
	// The number of bytes in free, mapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and unless the underlying memory is swapped out by the OS, they also
	// count towards physical memory usage. This is an alias for `tcmalloc.pageheap_free_bytes`.
	PageheapFree uint64 `protobuf:"varint,4,opt,name=pageheap_free,json=pageheapFree,proto3" json:"pageheap_free,omitempty"`
	// The amount of memory used by the TCMalloc thread caches (for small objects). This is an alias
	// for `tcmalloc.current_total_thread_cache_bytes`.
	TotalThreadCache uint64 `protobuf:"varint,5,opt,name=total_thread_cache,json=totalThreadCache,proto3" json:"total_thread_cache,omitempty"`
	// The number of bytes of the physical memory usage by the allocator. This is an alias for
	// `generic.total_physical_bytes`.
	TotalPhysicalBytes   uint64   `protobuf:"varint,6,opt,name=total_physical_bytes,json=totalPhysicalBytes,proto3" json:"total_physical_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b7ba9ad7a02b7f, []int{0}
}

func (m *Memory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memory.Unmarshal(m, b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
}
func (m *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(m, src)
}
func (m *Memory) XXX_Size() int {
	return xxx_messageInfo_Memory.Size(m)
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetAllocated() uint64 {
	if m != nil {
		return m.Allocated
	}
	return 0
}

func (m *Memory) GetHeapSize() uint64 {
	if m != nil {
		return m.HeapSize
	}
	return 0
}

func (m *Memory) GetPageheapUnmapped() uint64 {
	if m != nil {
		return m.PageheapUnmapped
	}
	return 0
}

func (m *Memory) GetPageheapFree() uint64 {
	if m != nil {
		return m.PageheapFree
	}
	return 0
}

func (m *Memory) GetTotalThreadCache() uint64 {
	if m != nil {
		return m.TotalThreadCache
	}
	return 0
}

func (m *Memory) GetTotalPhysicalBytes() uint64 {
	if m != nil {
		return m.TotalPhysicalBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*Memory)(nil), "envoy.admin.v2alpha.Memory")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/memory.proto", fileDescriptor_51b7ba9ad7a02b7f) }

var fileDescriptor_51b7ba9ad7a02b7f = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0xc4, 0x40,
	0x14, 0x46, 0xc9, 0xba, 0x06, 0x77, 0x54, 0x58, 0x47, 0x8b, 0xc1, 0x1f, 0x58, 0xb5, 0x11, 0x94,
	0x44, 0xb4, 0xb3, 0x8c, 0x60, 0x27, 0x04, 0x7f, 0xea, 0x70, 0x37, 0xb9, 0x9a, 0x81, 0x64, 0x66,
	0x98, 0x99, 0x2c, 0x66, 0x2b, 0x9f, 0xcb, 0x27, 0xb0, 0xf5, 0x61, 0xec, 0x25, 0x37, 0x59, 0x6d,
	0x6c, 0xbf, 0x73, 0xee, 0x2d, 0x0e, 0x9b, 0xa1, 0x5a, 0xe8, 0x36, 0x86, 0xa2, 0x96, 0x2a, 0x5e,
	0x5c, 0x41, 0x65, 0x4a, 0x88, 0x6b, 0xac, 0xb5, 0x6d, 0x23, 0x63, 0xb5, 0xd7, 0x7c, 0x97, 0x8c,
	0x88, 0x8c, 0x68, 0x30, 0xf6, 0x8f, 0x9a, 0xc2, 0x40, 0x0c, 0x4a, 0x69, 0x0f, 0x5e, 0x6a, 0xe5,
	0x62, 0xe7, 0xc1, 0x37, 0xae, 0xbf, 0x39, 0xf9, 0x0e, 0x58, 0x78, 0x4f, 0x4f, 0xf8, 0x21, 0x9b,
	0x40, 0x55, 0xe9, 0x1c, 0x3c, 0x16, 0x22, 0x98, 0x05, 0x67, 0xe3, 0x87, 0xbf, 0x81, 0x1f, 0xb0,
	0x49, 0x89, 0x60, 0x32, 0x27, 0x97, 0x28, 0x46, 0x44, 0x37, 0xba, 0xe1, 0x51, 0x2e, 0x91, 0x9f,
	0xb3, 0x1d, 0x03, 0xaf, 0x48, 0x42, 0xa3, 0x6a, 0x30, 0x06, 0x0b, 0xb1, 0x46, 0xd2, 0x74, 0x05,
	0x9e, 0x87, 0x9d, 0x9f, 0xb2, 0xed, 0x5f, 0xf9, 0xc5, 0x22, 0x8a, 0x31, 0x89, 0x5b, 0xab, 0xf1,
	0xce, 0x22, 0xf2, 0x0b, 0xc6, 0xbd, 0xf6, 0x50, 0x65, 0xbe, 0xb4, 0x08, 0x45, 0x96, 0x43, 0x5e,
	0xa2, 0x58, 0xef, 0x5f, 0x12, 0x79, 0x22, 0x70, 0xdb, 0xed, 0xfc, 0x92, 0xed, 0xf5, 0xb6, 0x29,
	0x5b, 0x27, 0x73, 0xa8, 0xb2, 0x79, 0xeb, 0xd1, 0x89, 0x90, 0xfc, 0xfe, 0x53, 0x3a, 0xa0, 0xa4,
	0x23, 0xc9, 0xcd, 0xc7, 0xfb, 0xe7, 0x57, 0x38, 0x9a, 0x06, 0xec, 0x58, 0xea, 0x88, 0xc2, 0x19,
	0xab, 0xdf, 0xda, 0xe8, 0x9f, 0x86, 0xc9, 0x66, 0x5f, 0x28, 0xed, 0x8a, 0xa5, 0xc1, 0x3c, 0xa4,
	0x74, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x3d, 0xef, 0xb5, 0x92, 0x01, 0x00, 0x00,
}
