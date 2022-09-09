// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/collector/metrics/flatten/flatten_metrics_service.proto

package flatten

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	v1 "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/common/v1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type FlattenExportMetricsServiceRequest struct {
	// An array of ResourceMetrics.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	RequestMetrics   []byte `protobuf:"bytes,1,opt,name=requestMetrics,proto3" json:"requestMetrics,omitempty"`
	ConnectMetrics   []byte `protobuf:"bytes,2,opt,name=connectMetrics,proto3" json:"connectMetrics,omitempty"`
	InterfaceMetrics []byte `protobuf:"bytes,3,opt,name=InterfaceMetrics,proto3" json:"InterfaceMetrics,omitempty"`
	// 采集时间
	StartTimeUnixNano uint64 `protobuf:"fixed64,4,opt,name=start_time_unix_nano,json=startTimeUnixNano,proto3" json:"start_time_unix_nano,omitempty"`
	// 主机信息
	Service              *v1.Service `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FlattenExportMetricsServiceRequest) Reset()         { *m = FlattenExportMetricsServiceRequest{} }
func (m *FlattenExportMetricsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*FlattenExportMetricsServiceRequest) ProtoMessage()    {}
func (*FlattenExportMetricsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfb7d089b040c0da, []int{0}
}
func (m *FlattenExportMetricsServiceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FlattenExportMetricsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FlattenExportMetricsServiceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FlattenExportMetricsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlattenExportMetricsServiceRequest.Merge(m, src)
}
func (m *FlattenExportMetricsServiceRequest) XXX_Size() int {
	return m.Size()
}
func (m *FlattenExportMetricsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlattenExportMetricsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlattenExportMetricsServiceRequest proto.InternalMessageInfo

func (m *FlattenExportMetricsServiceRequest) GetRequestMetrics() []byte {
	if m != nil {
		return m.RequestMetrics
	}
	return nil
}

func (m *FlattenExportMetricsServiceRequest) GetConnectMetrics() []byte {
	if m != nil {
		return m.ConnectMetrics
	}
	return nil
}

func (m *FlattenExportMetricsServiceRequest) GetInterfaceMetrics() []byte {
	if m != nil {
		return m.InterfaceMetrics
	}
	return nil
}

func (m *FlattenExportMetricsServiceRequest) GetStartTimeUnixNano() uint64 {
	if m != nil {
		return m.StartTimeUnixNano
	}
	return 0
}

func (m *FlattenExportMetricsServiceRequest) GetService() *v1.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type FlattenExportMetricsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlattenExportMetricsServiceResponse) Reset()         { *m = FlattenExportMetricsServiceResponse{} }
func (m *FlattenExportMetricsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*FlattenExportMetricsServiceResponse) ProtoMessage()    {}
func (*FlattenExportMetricsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfb7d089b040c0da, []int{1}
}
func (m *FlattenExportMetricsServiceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FlattenExportMetricsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FlattenExportMetricsServiceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FlattenExportMetricsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlattenExportMetricsServiceResponse.Merge(m, src)
}
func (m *FlattenExportMetricsServiceResponse) XXX_Size() int {
	return m.Size()
}
func (m *FlattenExportMetricsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlattenExportMetricsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlattenExportMetricsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FlattenExportMetricsServiceRequest)(nil), "opentelemetry.proto.collector.metrics.flatten.FlattenExportMetricsServiceRequest")
	proto.RegisterType((*FlattenExportMetricsServiceResponse)(nil), "opentelemetry.proto.collector.metrics.flatten.FlattenExportMetricsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/metrics/flatten/flatten_metrics_service.proto", fileDescriptor_cfb7d089b040c0da)
}

var fileDescriptor_cfb7d089b040c0da = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0x36, 0xab, 0x56, 0x88, 0x22, 0x3a, 0x28, 0x94, 0x1e, 0x4a, 0xa9, 0xb8, 0x94, 0x85, 0x26,
	0xec, 0x8a, 0x77, 0x59, 0x50, 0x90, 0x45, 0x59, 0xab, 0x5e, 0xbc, 0x0c, 0x69, 0xf6, 0x75, 0x8c,
	0x4e, 0xde, 0x1b, 0x93, 0x4c, 0xa9, 0xbf, 0xc3, 0x5f, 0xe2, 0xbf, 0x10, 0xbc, 0x78, 0xf7, 0x22,
	0xfd, 0x25, 0x32, 0x93, 0x4c, 0xd9, 0x6a, 0x75, 0x11, 0x3c, 0xe5, 0xe5, 0xcb, 0xf7, 0xbd, 0x97,
	0x7c, 0x79, 0x8f, 0x9f, 0x50, 0x05, 0x18, 0xa0, 0x04, 0x0b, 0xc1, 0x7d, 0x94, 0x95, 0xa3, 0x40,
	0x52, 0x53, 0x59, 0x82, 0x0e, 0xe4, 0x64, 0x83, 0x1a, 0xed, 0xe5, 0xa2, 0x54, 0x21, 0x00, 0x76,
	0x6b, 0x9e, 0xf0, 0xdc, 0x83, 0x5b, 0x1a, 0x0d, 0xa2, 0xd5, 0x65, 0xd3, 0xad, 0x64, 0x11, 0x14,
	0x9b, 0x64, 0x22, 0x89, 0x44, 0x4a, 0x32, 0x38, 0xd8, 0x5d, 0xdb, 0x5a, 0x42, 0xb9, 0x3c, 0x4c,
	0x51, 0xcc, 0x32, 0x98, 0x16, 0x26, 0xbc, 0xad, 0xe7, 0x42, 0x93, 0x95, 0x05, 0x15, 0x14, 0xd9,
	0xf3, 0x7a, 0xd1, 0xee, 0xa2, 0xb4, 0x89, 0x22, 0x7d, 0xfc, 0x69, 0x8f, 0x8f, 0x9f, 0xc4, 0x32,
	0x8f, 0x57, 0x15, 0xb9, 0xf0, 0x2c, 0xd6, 0x7e, 0x19, 0xef, 0x3b, 0x83, 0x0f, 0x35, 0xf8, 0x90,
	0xed, 0xf3, 0x9b, 0x2e, 0x86, 0xe9, 0xbc, 0xcf, 0x46, 0x6c, 0x72, 0x63, 0xf6, 0x0b, 0xda, 0xf0,
	0x34, 0x21, 0x82, 0xde, 0xf0, 0xf6, 0x22, 0x6f, 0x1b, 0xcd, 0x0e, 0xf8, 0xad, 0xa7, 0x18, 0xc0,
	0x2d, 0x94, 0x86, 0x8e, 0x79, 0xb9, 0x65, 0xfe, 0x86, 0x67, 0x92, 0xdf, 0xf1, 0x41, 0xb9, 0x90,
	0x07, 0x63, 0x21, 0xaf, 0xd1, 0xac, 0x72, 0x54, 0x48, 0xfd, 0x2b, 0x23, 0x36, 0xe9, 0xcd, 0x6e,
	0xb7, 0x67, 0xaf, 0x8c, 0x85, 0xd7, 0x68, 0x56, 0xcf, 0x15, 0x52, 0xf6, 0x88, 0x5f, 0x4b, 0x76,
	0xf7, 0xaf, 0x8e, 0xd8, 0xe4, 0xfa, 0xd1, 0xbe, 0xd8, 0xed, 0x77, 0x6b, 0xdb, 0xf2, 0x50, 0x74,
	0x8f, 0xed, 0x64, 0xe3, 0xfb, 0xfc, 0xde, 0x5f, 0x4d, 0xf1, 0x15, 0xa1, 0x87, 0xa3, 0xaf, 0x8c,
	0xdf, 0x4d, 0xbc, 0x6d, 0x46, 0xf6, 0x99, 0xf1, 0x5e, 0x94, 0x66, 0x2f, 0xc4, 0x3f, 0x7d, 0xb6,
	0xb8, 0xf8, 0x37, 0x06, 0xb3, 0xff, 0x99, 0x32, 0xbe, 0x65, 0x7c, 0xe9, 0xf8, 0x3b, 0xfb, 0xb2,
	0x1e, 0xb2, 0x6f, 0xeb, 0x21, 0xfb, 0xb1, 0x1e, 0x32, 0xfe, 0x90, 0x5c, 0x21, 0x96, 0xa5, 0xf1,
	0x42, 0x55, 0x56, 0x74, 0x5d, 0x74, 0x51, 0xad, 0xe3, 0xc1, 0x4e, 0x43, 0x4e, 0x1b, 0xd5, 0x29,
	0x7b, 0xb3, 0x3a, 0xd7, 0x9d, 0x27, 0x06, 0xcf, 0x4a, 0x83, 0xc5, 0xb4, 0x72, 0xf4, 0x0e, 0x74,
	0x90, 0xef, 0x13, 0x70, 0x6e, 0xac, 0x34, 0xa1, 0xaf, 0x2d, 0x38, 0x09, 0xed, 0xdd, 0xc1, 0x75,
	0x83, 0xb5, 0xd9, 0x9f, 0xa9, 0xa0, 0x62, 0x8f, 0x17, 0x80, 0x7f, 0x1e, 0xc8, 0x79, 0xaf, 0xe5,
	0x3c, 0xf8, 0x19, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x68, 0x81, 0xec, 0xc8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlattenMetricsServiceClient is the client API for FlattenMetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlattenMetricsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *FlattenExportMetricsServiceRequest, opts ...grpc.CallOption) (*FlattenExportMetricsServiceResponse, error)
}

type flattenMetricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewFlattenMetricsServiceClient(cc *grpc.ClientConn) FlattenMetricsServiceClient {
	return &flattenMetricsServiceClient{cc}
}

func (c *flattenMetricsServiceClient) Export(ctx context.Context, in *FlattenExportMetricsServiceRequest, opts ...grpc.CallOption) (*FlattenExportMetricsServiceResponse, error) {
	out := new(FlattenExportMetricsServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.metrics.flatten.FlattenMetricsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlattenMetricsServiceServer is the server API for FlattenMetricsService service.
type FlattenMetricsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *FlattenExportMetricsServiceRequest) (*FlattenExportMetricsServiceResponse, error)
}

// UnimplementedFlattenMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFlattenMetricsServiceServer struct {
}

func (*UnimplementedFlattenMetricsServiceServer) Export(ctx context.Context, req *FlattenExportMetricsServiceRequest) (*FlattenExportMetricsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterFlattenMetricsServiceServer(s *grpc.Server, srv FlattenMetricsServiceServer) {
	s.RegisterService(&_FlattenMetricsService_serviceDesc, srv)
}

func _FlattenMetricsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlattenExportMetricsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlattenMetricsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.metrics.flatten.FlattenMetricsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlattenMetricsServiceServer).Export(ctx, req.(*FlattenExportMetricsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlattenMetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.metrics.flatten.FlattenMetricsService",
	HandlerType: (*FlattenMetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _FlattenMetricsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/metrics/flatten/flatten_metrics_service.proto",
}

func (m *FlattenExportMetricsServiceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlattenExportMetricsServiceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FlattenExportMetricsServiceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Service != nil {
		{
			size, err := m.Service.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlattenMetricsService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.StartTimeUnixNano != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.StartTimeUnixNano))
		i--
		dAtA[i] = 0x21
	}
	if len(m.InterfaceMetrics) > 0 {
		i -= len(m.InterfaceMetrics)
		copy(dAtA[i:], m.InterfaceMetrics)
		i = encodeVarintFlattenMetricsService(dAtA, i, uint64(len(m.InterfaceMetrics)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectMetrics) > 0 {
		i -= len(m.ConnectMetrics)
		copy(dAtA[i:], m.ConnectMetrics)
		i = encodeVarintFlattenMetricsService(dAtA, i, uint64(len(m.ConnectMetrics)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RequestMetrics) > 0 {
		i -= len(m.RequestMetrics)
		copy(dAtA[i:], m.RequestMetrics)
		i = encodeVarintFlattenMetricsService(dAtA, i, uint64(len(m.RequestMetrics)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FlattenExportMetricsServiceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlattenExportMetricsServiceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FlattenExportMetricsServiceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintFlattenMetricsService(dAtA []byte, offset int, v uint64) int {
	offset -= sovFlattenMetricsService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FlattenExportMetricsServiceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestMetrics)
	if l > 0 {
		n += 1 + l + sovFlattenMetricsService(uint64(l))
	}
	l = len(m.ConnectMetrics)
	if l > 0 {
		n += 1 + l + sovFlattenMetricsService(uint64(l))
	}
	l = len(m.InterfaceMetrics)
	if l > 0 {
		n += 1 + l + sovFlattenMetricsService(uint64(l))
	}
	if m.StartTimeUnixNano != 0 {
		n += 9
	}
	if m.Service != nil {
		l = m.Service.Size()
		n += 1 + l + sovFlattenMetricsService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FlattenExportMetricsServiceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFlattenMetricsService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFlattenMetricsService(x uint64) (n int) {
	return sovFlattenMetricsService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FlattenExportMetricsServiceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlattenMetricsService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlattenExportMetricsServiceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlattenExportMetricsServiceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMetrics", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlattenMetricsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestMetrics = append(m.RequestMetrics[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestMetrics == nil {
				m.RequestMetrics = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectMetrics", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlattenMetricsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectMetrics = append(m.ConnectMetrics[:0], dAtA[iNdEx:postIndex]...)
			if m.ConnectMetrics == nil {
				m.ConnectMetrics = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceMetrics", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlattenMetricsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceMetrics = append(m.InterfaceMetrics[:0], dAtA[iNdEx:postIndex]...)
			if m.InterfaceMetrics == nil {
				m.InterfaceMetrics = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimeUnixNano", wireType)
			}
			m.StartTimeUnixNano = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.StartTimeUnixNano = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlattenMetricsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Service == nil {
				m.Service = &v1.Service{}
			}
			if err := m.Service.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlattenMetricsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FlattenExportMetricsServiceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlattenMetricsService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlattenExportMetricsServiceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlattenExportMetricsServiceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFlattenMetricsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlattenMetricsService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFlattenMetricsService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlattenMetricsService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFlattenMetricsService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFlattenMetricsService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFlattenMetricsService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFlattenMetricsService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFlattenMetricsService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFlattenMetricsService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlattenMetricsService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFlattenMetricsService = fmt.Errorf("proto: unexpected end of group")
)