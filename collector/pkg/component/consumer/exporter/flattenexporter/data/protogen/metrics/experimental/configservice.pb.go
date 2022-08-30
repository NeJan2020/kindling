// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/metrics/experimental/configservice.proto

package experimental

import (
	context "context"
	fmt "fmt"
	v1 "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/resource/v1"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MetricConfigRequest struct {
	// Required. The resource for which configuration should be returned.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// Optional. The value of MetricConfigResponse.fingerprint for the last
	// configuration that the caller received and successfully applied.
	LastKnownFingerprint []byte   `protobuf:"bytes,2,opt,name=last_known_fingerprint,json=lastKnownFingerprint,proto3" json:"last_known_fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricConfigRequest) Reset()         { *m = MetricConfigRequest{} }
func (m *MetricConfigRequest) String() string { return proto.CompactTextString(m) }
func (*MetricConfigRequest) ProtoMessage()    {}
func (*MetricConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{0}
}
func (m *MetricConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigRequest.Unmarshal(m, b)
}
func (m *MetricConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigRequest.Marshal(b, m, deterministic)
}
func (m *MetricConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigRequest.Merge(m, src)
}
func (m *MetricConfigRequest) XXX_Size() int {
	return xxx_messageInfo_MetricConfigRequest.Size(m)
}
func (m *MetricConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigRequest proto.InternalMessageInfo

func (m *MetricConfigRequest) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *MetricConfigRequest) GetLastKnownFingerprint() []byte {
	if m != nil {
		return m.LastKnownFingerprint
	}
	return nil
}

type MetricConfigResponse struct {
	// Optional. The fingerprint associated with this MetricConfigResponse. Each
	// change in configs yields a different fingerprint. The resource SHOULD copy
	// this value to MetricConfigRequest.last_known_fingerprint for the next
	// configuration request. If there are no changes between fingerprint and
	// MetricConfigRequest.last_known_fingerprint, then all other fields besides
	// fingerprint in the response are optional, or the same as the last update if
	// present.
	//
	// The exact mechanics of generating the fingerprint is up to the
	// implementation. However, a fingerprint must be deterministically determined
	// by the configurations -- the same configuration will generate the same
	// fingerprint on any instance of an implementation. Hence using a timestamp is
	// unacceptable, but a deterministic hash is fine.
	Fingerprint []byte `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// A single metric may match multiple schedules. In such cases, the schedule
	// that specifies the smallest period is applied.
	//
	// Note, for optimization purposes, it is recommended to use as few schedules
	// as possible to capture all required metric updates. Where you can be
	// conservative, do take full advantage of the inclusion/exclusion patterns to
	// capture as much of your targeted metrics.
	Schedules []*MetricConfigResponse_Schedule `protobuf:"bytes,2,rep,name=schedules,proto3" json:"schedules,omitempty"`
	// Optional. The client is suggested to wait this long (in seconds) before
	// pinging the configuration service again.
	SuggestedWaitTimeSec int32    `protobuf:"varint,3,opt,name=suggested_wait_time_sec,json=suggestedWaitTimeSec,proto3" json:"suggested_wait_time_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricConfigResponse) Reset()         { *m = MetricConfigResponse{} }
func (m *MetricConfigResponse) String() string { return proto.CompactTextString(m) }
func (*MetricConfigResponse) ProtoMessage()    {}
func (*MetricConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{1}
}
func (m *MetricConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigResponse.Unmarshal(m, b)
}
func (m *MetricConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigResponse.Marshal(b, m, deterministic)
}
func (m *MetricConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigResponse.Merge(m, src)
}
func (m *MetricConfigResponse) XXX_Size() int {
	return xxx_messageInfo_MetricConfigResponse.Size(m)
}
func (m *MetricConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigResponse proto.InternalMessageInfo

func (m *MetricConfigResponse) GetFingerprint() []byte {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *MetricConfigResponse) GetSchedules() []*MetricConfigResponse_Schedule {
	if m != nil {
		return m.Schedules
	}
	return nil
}

func (m *MetricConfigResponse) GetSuggestedWaitTimeSec() int32 {
	if m != nil {
		return m.SuggestedWaitTimeSec
	}
	return 0
}

// A Schedule is used to apply a particular scheduling configuration to
// a metric. If a metric name matches a schedule's patterns, then the metric
// adopts the configuration specified by the schedule.
type MetricConfigResponse_Schedule struct {
	// Metrics with names that match a rule in the inclusion_patterns are
	// targeted by this schedule. Metrics that match the exclusion_patterns
	// are not targeted for this schedule, even if they match an inclusion
	// pattern.
	ExclusionPatterns []*MetricConfigResponse_Schedule_Pattern `protobuf:"bytes,1,rep,name=exclusion_patterns,json=exclusionPatterns,proto3" json:"exclusion_patterns,omitempty"`
	InclusionPatterns []*MetricConfigResponse_Schedule_Pattern `protobuf:"bytes,2,rep,name=inclusion_patterns,json=inclusionPatterns,proto3" json:"inclusion_patterns,omitempty"`
	// Describes the collection period for each metric in seconds.
	// A period of 0 means to not export.
	PeriodSec            int32    `protobuf:"varint,3,opt,name=period_sec,json=periodSec,proto3" json:"period_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricConfigResponse_Schedule) Reset()         { *m = MetricConfigResponse_Schedule{} }
func (m *MetricConfigResponse_Schedule) String() string { return proto.CompactTextString(m) }
func (*MetricConfigResponse_Schedule) ProtoMessage()    {}
func (*MetricConfigResponse_Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{1, 0}
}
func (m *MetricConfigResponse_Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigResponse_Schedule.Unmarshal(m, b)
}
func (m *MetricConfigResponse_Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigResponse_Schedule.Marshal(b, m, deterministic)
}
func (m *MetricConfigResponse_Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigResponse_Schedule.Merge(m, src)
}
func (m *MetricConfigResponse_Schedule) XXX_Size() int {
	return xxx_messageInfo_MetricConfigResponse_Schedule.Size(m)
}
func (m *MetricConfigResponse_Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigResponse_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigResponse_Schedule proto.InternalMessageInfo

func (m *MetricConfigResponse_Schedule) GetExclusionPatterns() []*MetricConfigResponse_Schedule_Pattern {
	if m != nil {
		return m.ExclusionPatterns
	}
	return nil
}

func (m *MetricConfigResponse_Schedule) GetInclusionPatterns() []*MetricConfigResponse_Schedule_Pattern {
	if m != nil {
		return m.InclusionPatterns
	}
	return nil
}

func (m *MetricConfigResponse_Schedule) GetPeriodSec() int32 {
	if m != nil {
		return m.PeriodSec
	}
	return 0
}

// A light-weight pattern that can match 1 or more
// metrics, for which this schedule will apply. The string is used to
// match against metric names. It should not exceed 100k characters.
type MetricConfigResponse_Schedule_Pattern struct {
	// Types that are valid to be assigned to Match:
	//	*MetricConfigResponse_Schedule_Pattern_Equals
	//	*MetricConfigResponse_Schedule_Pattern_StartsWith
	Match                isMetricConfigResponse_Schedule_Pattern_Match `protobuf_oneof:"match"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *MetricConfigResponse_Schedule_Pattern) Reset()         { *m = MetricConfigResponse_Schedule_Pattern{} }
func (m *MetricConfigResponse_Schedule_Pattern) String() string { return proto.CompactTextString(m) }
func (*MetricConfigResponse_Schedule_Pattern) ProtoMessage()    {}
func (*MetricConfigResponse_Schedule_Pattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{1, 0, 0}
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Unmarshal(m, b)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Marshal(b, m, deterministic)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Merge(m, src)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Size() int {
	return xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Size(m)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigResponse_Schedule_Pattern proto.InternalMessageInfo

type isMetricConfigResponse_Schedule_Pattern_Match interface {
	isMetricConfigResponse_Schedule_Pattern_Match()
}

type MetricConfigResponse_Schedule_Pattern_Equals struct {
	Equals string `protobuf:"bytes,1,opt,name=equals,proto3,oneof" json:"equals,omitempty"`
}
type MetricConfigResponse_Schedule_Pattern_StartsWith struct {
	StartsWith string `protobuf:"bytes,2,opt,name=starts_with,json=startsWith,proto3,oneof" json:"starts_with,omitempty"`
}

func (*MetricConfigResponse_Schedule_Pattern_Equals) isMetricConfigResponse_Schedule_Pattern_Match() {
}
func (*MetricConfigResponse_Schedule_Pattern_StartsWith) isMetricConfigResponse_Schedule_Pattern_Match() {
}

func (m *MetricConfigResponse_Schedule_Pattern) GetMatch() isMetricConfigResponse_Schedule_Pattern_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *MetricConfigResponse_Schedule_Pattern) GetEquals() string {
	if x, ok := m.GetMatch().(*MetricConfigResponse_Schedule_Pattern_Equals); ok {
		return x.Equals
	}
	return ""
}

func (m *MetricConfigResponse_Schedule_Pattern) GetStartsWith() string {
	if x, ok := m.GetMatch().(*MetricConfigResponse_Schedule_Pattern_StartsWith); ok {
		return x.StartsWith
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetricConfigResponse_Schedule_Pattern) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetricConfigResponse_Schedule_Pattern_Equals)(nil),
		(*MetricConfigResponse_Schedule_Pattern_StartsWith)(nil),
	}
}

func init() {
	proto.RegisterType((*MetricConfigRequest)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigRequest")
	proto.RegisterType((*MetricConfigResponse)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigResponse")
	proto.RegisterType((*MetricConfigResponse_Schedule)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule")
	proto.RegisterType((*MetricConfigResponse_Schedule_Pattern)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule.Pattern")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/metrics/experimental/configservice.proto", fileDescriptor_79b5d4ea55caf90b)
}

var fileDescriptor_79b5d4ea55caf90b = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x71, 0x4a, 0x3f, 0xb2, 0xa9, 0x84, 0x30, 0x11, 0x58, 0x91, 0x90, 0x42, 0x4f, 0xe1,
	0xc0, 0x9a, 0x06, 0xb8, 0x01, 0x87, 0x20, 0x28, 0x52, 0x85, 0x1a, 0x39, 0x48, 0x95, 0xb8, 0x58,
	0x9b, 0xcd, 0xc4, 0x59, 0xba, 0xde, 0x75, 0x77, 0xc7, 0x49, 0xb9, 0xf0, 0x0c, 0x88, 0x37, 0xe0,
	0x59, 0x78, 0x04, 0xde, 0x86, 0x13, 0x5a, 0xdb, 0x75, 0x1d, 0x11, 0x21, 0xc4, 0xc7, 0x6d, 0xf7,
	0x3f, 0x33, 0xbf, 0xff, 0x64, 0x36, 0x1e, 0xf2, 0x54, 0x67, 0xa0, 0x10, 0x24, 0xa4, 0x80, 0xe6,
	0x43, 0x98, 0x19, 0x8d, 0x3a, 0x74, 0x67, 0xc1, 0x6d, 0x08, 0x17, 0x19, 0x18, 0x91, 0x82, 0x42,
	0x26, 0x43, 0xae, 0xd5, 0x5c, 0x24, 0x16, 0xcc, 0x52, 0x70, 0xa0, 0x45, 0xa2, 0x3f, 0x58, 0xab,
	0x2e, 0x45, 0x5a, 0x55, 0xd3, 0x66, 0x75, 0x8f, 0x6e, 0xf2, 0x31, 0x60, 0x75, 0x6e, 0x38, 0x84,
	0xcb, 0xc3, 0xfa, 0x5c, 0x42, 0x0e, 0x3e, 0x7b, 0xe4, 0xd6, 0x9b, 0x02, 0xf4, 0xa2, 0xf0, 0x8d,
	0xe0, 0x3c, 0x07, 0x8b, 0xfe, 0x4b, 0xb2, 0x77, 0x99, 0x19, 0x78, 0x7d, 0x6f, 0xd0, 0x19, 0xde,
	0xa7, 0x9b, 0x9a, 0xa8, 0x71, 0xcb, 0x43, 0x1a, 0x55, 0xe7, 0xa8, 0x2e, 0xf5, 0x1f, 0x93, 0xdb,
	0x92, 0x59, 0x8c, 0xcf, 0x94, 0x5e, 0xa9, 0x78, 0x2e, 0x54, 0x02, 0x26, 0x33, 0x42, 0x61, 0xd0,
	0xea, 0x7b, 0x83, 0xfd, 0xa8, 0xeb, 0xa2, 0xc7, 0x2e, 0xf8, 0xea, 0x2a, 0x76, 0xf0, 0xed, 0x3a,
	0xe9, 0xae, 0x37, 0x65, 0x33, 0xad, 0x2c, 0xf8, 0x7d, 0xd2, 0x69, 0x32, 0xbc, 0x82, 0xd1, 0x94,
	0x7c, 0x20, 0x6d, 0xcb, 0x17, 0x30, 0xcb, 0x25, 0xd8, 0xa0, 0xd5, 0xdf, 0x1a, 0x74, 0x86, 0x47,
	0xf4, 0x77, 0xa7, 0x47, 0x37, 0x99, 0xd2, 0x49, 0xc5, 0x8b, 0xae, 0xc8, 0xfe, 0x13, 0x72, 0xc7,
	0xe6, 0x49, 0x02, 0x16, 0x61, 0x16, 0xaf, 0x98, 0xc0, 0x18, 0x45, 0x0a, 0xb1, 0x05, 0x1e, 0x6c,
	0xf5, 0xbd, 0xc1, 0x76, 0xd4, 0xad, 0xc3, 0xa7, 0x4c, 0xe0, 0x5b, 0x91, 0xc2, 0x04, 0x78, 0xef,
	0x7b, 0x8b, 0xec, 0x5d, 0xe2, 0xfc, 0x8f, 0xc4, 0x87, 0x0b, 0x2e, 0x73, 0x2b, 0xb4, 0x8a, 0x33,
	0x86, 0x08, 0x46, 0xd9, 0xc0, 0x2b, 0x7a, 0x3e, 0xf9, 0x47, 0x3d, 0xd3, 0x71, 0xc9, 0x8d, 0x6e,
	0xd6, 0x56, 0x95, 0x62, 0x9d, 0xbf, 0x50, 0x3f, 0xf9, 0xb7, 0xfe, 0x93, 0x7f, 0x6d, 0x55, 0xfb,
	0xdf, 0x25, 0xc4, 0x61, 0xf4, 0xac, 0x31, 0xb6, 0x76, 0xa9, 0xb8, 0x59, 0x9d, 0x90, 0xdd, 0x2a,
	0xd5, 0x0f, 0xc8, 0x0e, 0x9c, 0xe7, 0x4c, 0xda, 0xe2, 0xc5, 0xdb, 0xaf, 0xaf, 0x45, 0xd5, 0xdd,
	0xbf, 0x47, 0x3a, 0x16, 0x99, 0x41, 0x1b, 0xaf, 0x04, 0x2e, 0x8a, 0x3f, 0x95, 0x0b, 0x93, 0x52,
	0x3c, 0x15, 0xb8, 0x18, 0xed, 0x92, 0xed, 0x94, 0x21, 0x5f, 0x0c, 0xbf, 0x78, 0x64, 0xbf, 0xd9,
	0xac, 0xff, 0xc9, 0x23, 0x37, 0x8e, 0x00, 0xd7, 0xb4, 0x67, 0x7f, 0xfa, 0xc3, 0x8b, 0xcf, 0xa6,
	0xf7, 0xfc, 0xef, 0xe6, 0x36, 0xfa, 0xea, 0x91, 0x87, 0xda, 0x24, 0x74, 0x29, 0x85, 0xa5, 0x2c,
	0x4b, 0x4b, 0xc0, 0x34, 0x9f, 0xff, 0x82, 0x34, 0x0a, 0x9a, 0xa8, 0x49, 0xb9, 0x38, 0xc6, 0x2e,
	0x7d, 0xec, 0xbd, 0xb3, 0x89, 0xc0, 0x45, 0x3e, 0xa5, 0x5c, 0xa7, 0xe1, 0xb1, 0x50, 0x33, 0x29,
	0x54, 0xf2, 0x20, 0x33, 0xfa, 0x3d, 0x70, 0x0c, 0xcf, 0x2a, 0x21, 0xe4, 0x5a, 0x4a, 0xe0, 0xa8,
	0x8d, 0xdb, 0x3e, 0x36, 0x4f, 0xc1, 0xb8, 0x9d, 0xa4, 0x0d, 0x82, 0x09, 0xe7, 0xd2, 0xbd, 0x82,
	0xaa, 0xef, 0x33, 0x86, 0xac, 0xdc, 0x2b, 0x09, 0xa8, 0x8d, 0x2b, 0x6c, 0xba, 0x53, 0x84, 0x1f,
	0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x94, 0xd3, 0x49, 0x67, 0xf5, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricConfigClient is the client API for MetricConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricConfigClient interface {
	GetMetricConfig(ctx context.Context, in *MetricConfigRequest, opts ...grpc.CallOption) (*MetricConfigResponse, error)
}

type metricConfigClient struct {
	cc *grpc.ClientConn
}

func NewMetricConfigClient(cc *grpc.ClientConn) MetricConfigClient {
	return &metricConfigClient{cc}
}

func (c *metricConfigClient) GetMetricConfig(ctx context.Context, in *MetricConfigRequest, opts ...grpc.CallOption) (*MetricConfigResponse, error) {
	out := new(MetricConfigResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.metrics.experimental.MetricConfig/GetMetricConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricConfigServer is the server API for MetricConfig service.
type MetricConfigServer interface {
	GetMetricConfig(context.Context, *MetricConfigRequest) (*MetricConfigResponse, error)
}

// UnimplementedMetricConfigServer can be embedded to have forward compatible implementations.
type UnimplementedMetricConfigServer struct {
}

func (*UnimplementedMetricConfigServer) GetMetricConfig(ctx context.Context, req *MetricConfigRequest) (*MetricConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricConfig not implemented")
}

func RegisterMetricConfigServer(s *grpc.Server, srv MetricConfigServer) {
	s.RegisterService(&_MetricConfig_serviceDesc, srv)
}

func _MetricConfig_GetMetricConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricConfigServer).GetMetricConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.metrics.experimental.MetricConfig/GetMetricConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricConfigServer).GetMetricConfig(ctx, req.(*MetricConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.metrics.experimental.MetricConfig",
	HandlerType: (*MetricConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetricConfig",
			Handler:    _MetricConfig_GetMetricConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/metrics/experimental/configservice.proto",
}