// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("Service.proto", fileDescriptor_88420990119a24c7)
}

var fileDescriptor_88420990119a24c7 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x61, 0xa0, 0xcd, 0x65, 0x50, 0x5d, 0x60, 0x48, 0xdd, 0x84, 0x60, 0x4f, 0xf0, 0xe2,
	0xd1, 0x4e, 0x08, 0x24, 0x34, 0x89, 0xae, 0xab, 0x04, 0x12, 0x93, 0xa2, 0x86, 0xa7, 0xbd, 0x79,
	0xc9, 0x6d, 0x6a, 0x29, 0xb1, 0x3d, 0xe7, 0x26, 0x12, 0xbf, 0xc2, 0x7f, 0xf0, 0x7f, 0xc8, 0x8e,
	0x93, 0x8e, 0x52, 0x55, 0x7b, 0x3b, 0xf7, 0x5c, 0x9f, 0x73, 0xae, 0xed, 0xcb, 0x0e, 0x62, 0xb4,
	0xb5, 0x4c, 0x90, 0x1b, 0xab, 0x49, 0xc3, 0x83, 0x7a, 0x34, 0x3c, 0xca, 0xb4, 0xce, 0x72, 0x3c,
	0xf5, 0xcc, 0x4d, 0xb5, 0x38, 0xc5, 0xc2, 0xd0, 0xaf, 0xe6, 0xc0, 0x90, 0xc5, 0x46, 0xa8, 0x0e,
	0x93, 0xa0, 0x80, 0x07, 0x3f, 0x97, 0x16, 0x45, 0x7a, 0x59, 0x15, 0x26, 0x30, 0xfb, 0xd3, 0x22,
	0x6d, 0xe0, 0xf8, 0x2b, 0xdb, 0x75, 0x32, 0xf8, 0xcc, 0xf6, 0x62, 0x54, 0xa9, 0xc7, 0x03, 0x5e,
	0x8f, 0x78, 0xe4, 0xe0, 0x15, 0x96, 0xa5, 0xc8, 0x70, 0x78, 0xc8, 0x9b, 0x60, 0xde, 0x06, 0xf3,
	0x99, 0x0b, 0x3e, 0xd9, 0x79, 0xd7, 0x1b, 0x23, 0x7b, 0x34, 0xc9, 0x50, 0x11, 0x8c, 0xd8, 0x60,
	0x8e, 0xb7, 0x15, 0x96, 0xe4, 0xeb, 0xef, 0x6a, 0xa1, 0xe1, 0xa9, 0xb7, 0xea, 0xea, 0x61, 0xdf,
	0xd7, 0x73, 0x2c, 0xab, 0x9c, 0x4e, 0x76, 0xe0, 0x3d, 0xeb, 0x47, 0x52, 0x65, 0x31, 0x96, 0xa5,
	0xd4, 0x0a, 0xf6, 0x7d, 0xd7, 0x31, 0xc3, 0x15, 0x74, 0x21, 0x1f, 0x7a, 0xe3, 0x3f, 0x3d, 0xb6,
	0x77, 0x85, 0x24, 0x52, 0x41, 0x02, 0x3e, 0x32, 0x08, 0x51, 0xf1, 0x6d, 0xee, 0xd8, 0x4b, 0xc7,
	0x86, 0xb9, 0x57, 0xcc, 0x7a, 0xdc, 0x4a, 0x36, 0x31, 0x72, 0x4d, 0x76, 0x87, 0x59, 0x97, 0x7d,
	0x61, 0x2f, 0xdb, 0x34, 0xb2, 0x52, 0x65, 0x9d, 0xf2, 0x79, 0x13, 0xf8, 0x0f, 0xb9, 0x26, 0x1e,
	0xcf, 0xd8, 0xae, 0xfb, 0x0b, 0x38, 0x77, 0xff, 0xa9, 0x52, 0xff, 0x14, 0x9e, 0x08, 0xd3, 0x92,
	0xa0, 0xfb, 0xbc, 0xf2, 0xef, 0x87, 0xec, 0x30, 0xb2, 0x7a, 0x21, 0x73, 0xb4, 0x53, 0x5d, 0x14,
	0x42, 0xa5, 0x61, 0x3d, 0xe0, 0x13, 0x3b, 0xf8, 0x26, 0x54, 0x9a, 0x63, 0xe0, 0xe1, 0x99, 0x77,
	0x9e, 0x16, 0x69, 0x6b, 0xdc, 0x11, 0xe1, 0x1a, 0xcd, 0x93, 0xc2, 0x39, 0xeb, 0x07, 0xc9, 0x2c,
	0x59, 0x6a, 0x78, 0xd1, 0x9e, 0x72, 0xd5, 0x1c, 0x4b, 0xa3, 0x55, 0xb9, 0x65, 0x28, 0xb8, 0x66,
	0xaf, 0xdb, 0x49, 0xc8, 0xa2, 0x28, 0x26, 0x09, 0xc9, 0x1a, 0x9b, 0x65, 0x9b, 0xea, 0x4a, 0x11,
	0x1c, 0xb7, 0x8e, 0xff, 0xb5, 0xe6, 0x58, 0x6e, 0xbb, 0x2e, 0x44, 0xec, 0x55, 0xf0, 0xbe, 0x2b,
	0x75, 0x2b, 0x0c, 0x47, 0x9b, 0x4c, 0x5d, 0x67, 0xab, 0x27, 0x5c, 0xb3, 0xe3, 0x0d, 0x8e, 0x3f,
	0x64, 0xb6, 0x24, 0x6f, 0xfb, 0x66, 0x93, 0x6d, 0xd7, 0xde, 0xea, 0x7d, 0x71, 0xc6, 0xde, 0x26,
	0xba, 0xe0, 0x4a, 0xd4, 0x68, 0x13, 0x6d, 0x0d, 0x37, 0x52, 0x19, 0x2d, 0x15, 0xf1, 0xcc, 0x9a,
	0x84, 0x93, 0x15, 0x09, 0x5e, 0x3c, 0x09, 0xff, 0x15, 0x39, 0x75, 0xd4, 0xbb, 0x79, 0xec, 0x6d,
	0xce, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xaf, 0x53, 0x3a, 0xe7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpanClient is the client API for Span service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpanClient interface {
	SendSpan(ctx context.Context, opts ...grpc.CallOption) (Span_SendSpanClient, error)
}

type spanClient struct {
	cc grpc.ClientConnInterface
}

func NewSpanClient(cc grpc.ClientConnInterface) SpanClient {
	return &spanClient{cc}
}

func (c *spanClient) SendSpan(ctx context.Context, opts ...grpc.CallOption) (Span_SendSpanClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Span_serviceDesc.Streams[0], "/v1.Span/SendSpan", opts...)
	if err != nil {
		return nil, err
	}
	x := &spanSendSpanClient{stream}
	return x, nil
}

type Span_SendSpanClient interface {
	Send(*PSpanMessage) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type spanSendSpanClient struct {
	grpc.ClientStream
}

func (x *spanSendSpanClient) Send(m *PSpanMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *spanSendSpanClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpanServer is the server API for Span service.
type SpanServer interface {
	SendSpan(Span_SendSpanServer) error
}

// UnimplementedSpanServer can be embedded to have forward compatible implementations.
type UnimplementedSpanServer struct {
}

func (*UnimplementedSpanServer) SendSpan(srv Span_SendSpanServer) error {
	return status.Errorf(codes.Unimplemented, "method SendSpan not implemented")
}

func RegisterSpanServer(s *grpc.Server, srv SpanServer) {
	s.RegisterService(&_Span_serviceDesc, srv)
}

func _Span_SendSpan_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SpanServer).SendSpan(&spanSendSpanServer{stream})
}

type Span_SendSpanServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PSpanMessage, error)
	grpc.ServerStream
}

type spanSendSpanServer struct {
	grpc.ServerStream
}

func (x *spanSendSpanServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *spanSendSpanServer) Recv() (*PSpanMessage, error) {
	m := new(PSpanMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Span_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Span",
	HandlerType: (*SpanServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendSpan",
			Handler:       _Span_SendSpan_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Service.proto",
}

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	RequestAgentInfo(ctx context.Context, in *PAgentInfo, opts ...grpc.CallOption) (*PResult, error)
	PingSession(ctx context.Context, opts ...grpc.CallOption) (Agent_PingSessionClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) RequestAgentInfo(ctx context.Context, in *PAgentInfo, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Agent/RequestAgentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PingSession(ctx context.Context, opts ...grpc.CallOption) (Agent_PingSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Agent_serviceDesc.Streams[0], "/v1.Agent/PingSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentPingSessionClient{stream}
	return x, nil
}

type Agent_PingSessionClient interface {
	Send(*PPing) error
	Recv() (*PPing, error)
	grpc.ClientStream
}

type agentPingSessionClient struct {
	grpc.ClientStream
}

func (x *agentPingSessionClient) Send(m *PPing) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentPingSessionClient) Recv() (*PPing, error) {
	m := new(PPing)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	RequestAgentInfo(context.Context, *PAgentInfo) (*PResult, error)
	PingSession(Agent_PingSessionServer) error
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) RequestAgentInfo(ctx context.Context, req *PAgentInfo) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAgentInfo not implemented")
}
func (*UnimplementedAgentServer) PingSession(srv Agent_PingSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method PingSession not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_RequestAgentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PAgentInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RequestAgentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Agent/RequestAgentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RequestAgentInfo(ctx, req.(*PAgentInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_PingSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).PingSession(&agentPingSessionServer{stream})
}

type Agent_PingSessionServer interface {
	Send(*PPing) error
	Recv() (*PPing, error)
	grpc.ServerStream
}

type agentPingSessionServer struct {
	grpc.ServerStream
}

func (x *agentPingSessionServer) Send(m *PPing) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentPingSessionServer) Recv() (*PPing, error) {
	m := new(PPing)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAgentInfo",
			Handler:    _Agent_RequestAgentInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingSession",
			Handler:       _Agent_PingSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Service.proto",
}

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataClient interface {
	RequestSqlMetaData(ctx context.Context, in *PSqlMetaData, opts ...grpc.CallOption) (*PResult, error)
	RequestApiMetaData(ctx context.Context, in *PApiMetaData, opts ...grpc.CallOption) (*PResult, error)
	RequestStringMetaData(ctx context.Context, in *PStringMetaData, opts ...grpc.CallOption) (*PResult, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) RequestSqlMetaData(ctx context.Context, in *PSqlMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestSqlMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) RequestApiMetaData(ctx context.Context, in *PApiMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestApiMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) RequestStringMetaData(ctx context.Context, in *PStringMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestStringMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
type MetadataServer interface {
	RequestSqlMetaData(context.Context, *PSqlMetaData) (*PResult, error)
	RequestApiMetaData(context.Context, *PApiMetaData) (*PResult, error)
	RequestStringMetaData(context.Context, *PStringMetaData) (*PResult, error)
}

// UnimplementedMetadataServer can be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (*UnimplementedMetadataServer) RequestSqlMetaData(ctx context.Context, req *PSqlMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSqlMetaData not implemented")
}
func (*UnimplementedMetadataServer) RequestApiMetaData(ctx context.Context, req *PApiMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestApiMetaData not implemented")
}
func (*UnimplementedMetadataServer) RequestStringMetaData(ctx context.Context, req *PStringMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestStringMetaData not implemented")
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_RequestSqlMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PSqlMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestSqlMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestSqlMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestSqlMetaData(ctx, req.(*PSqlMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_RequestApiMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PApiMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestApiMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestApiMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestApiMetaData(ctx, req.(*PApiMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_RequestStringMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PStringMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestStringMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestStringMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestStringMetaData(ctx, req.(*PStringMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestSqlMetaData",
			Handler:    _Metadata_RequestSqlMetaData_Handler,
		},
		{
			MethodName: "RequestApiMetaData",
			Handler:    _Metadata_RequestApiMetaData_Handler,
		},
		{
			MethodName: "RequestStringMetaData",
			Handler:    _Metadata_RequestStringMetaData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Service.proto",
}

// StatClient is the client API for Stat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatClient interface {
	SendAgentStat(ctx context.Context, opts ...grpc.CallOption) (Stat_SendAgentStatClient, error)
}

type statClient struct {
	cc grpc.ClientConnInterface
}

func NewStatClient(cc grpc.ClientConnInterface) StatClient {
	return &statClient{cc}
}

func (c *statClient) SendAgentStat(ctx context.Context, opts ...grpc.CallOption) (Stat_SendAgentStatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stat_serviceDesc.Streams[0], "/v1.Stat/SendAgentStat", opts...)
	if err != nil {
		return nil, err
	}
	x := &statSendAgentStatClient{stream}
	return x, nil
}

type Stat_SendAgentStatClient interface {
	Send(*PStatMessage) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type statSendAgentStatClient struct {
	grpc.ClientStream
}

func (x *statSendAgentStatClient) Send(m *PStatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *statSendAgentStatClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatServer is the server API for Stat service.
type StatServer interface {
	SendAgentStat(Stat_SendAgentStatServer) error
}

// UnimplementedStatServer can be embedded to have forward compatible implementations.
type UnimplementedStatServer struct {
}

func (*UnimplementedStatServer) SendAgentStat(srv Stat_SendAgentStatServer) error {
	return status.Errorf(codes.Unimplemented, "method SendAgentStat not implemented")
}

func RegisterStatServer(s *grpc.Server, srv StatServer) {
	s.RegisterService(&_Stat_serviceDesc, srv)
}

func _Stat_SendAgentStat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StatServer).SendAgentStat(&statSendAgentStatServer{stream})
}

type Stat_SendAgentStatServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PStatMessage, error)
	grpc.ServerStream
}

type statSendAgentStatServer struct {
	grpc.ServerStream
}

func (x *statSendAgentStatServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *statSendAgentStatServer) Recv() (*PStatMessage, error) {
	m := new(PStatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Stat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Stat",
	HandlerType: (*StatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendAgentStat",
			Handler:       _Stat_SendAgentStat_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Service.proto",
}

// ProfilerCommandServiceClient is the client API for ProfilerCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfilerCommandServiceClient interface {
	HandleCommand(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandClient, error)
	CommandEcho(ctx context.Context, in *PCmdEchoResponse, opts ...grpc.CallOption) (*empty.Empty, error)
	CommandStreamActiveThreadCount(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_CommandStreamActiveThreadCountClient, error)
	CommandActiveThreadDump(ctx context.Context, in *PCmdActiveThreadDumpRes, opts ...grpc.CallOption) (*empty.Empty, error)
	CommandActiveThreadLightDump(ctx context.Context, in *PCmdActiveThreadLightDumpRes, opts ...grpc.CallOption) (*empty.Empty, error)
}

type profilerCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilerCommandServiceClient(cc grpc.ClientConnInterface) ProfilerCommandServiceClient {
	return &profilerCommandServiceClient{cc}
}

func (c *profilerCommandServiceClient) HandleCommand(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfilerCommandService_serviceDesc.Streams[0], "/v1.ProfilerCommandService/HandleCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilerCommandServiceHandleCommandClient{stream}
	return x, nil
}

type ProfilerCommandService_HandleCommandClient interface {
	Send(*PCmdMessage) error
	Recv() (*PCmdRequest, error)
	grpc.ClientStream
}

type profilerCommandServiceHandleCommandClient struct {
	grpc.ClientStream
}

func (x *profilerCommandServiceHandleCommandClient) Send(m *PCmdMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandClient) Recv() (*PCmdRequest, error) {
	m := new(PCmdRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profilerCommandServiceClient) CommandEcho(ctx context.Context, in *PCmdEchoResponse, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerCommandServiceClient) CommandStreamActiveThreadCount(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_CommandStreamActiveThreadCountClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfilerCommandService_serviceDesc.Streams[1], "/v1.ProfilerCommandService/CommandStreamActiveThreadCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilerCommandServiceCommandStreamActiveThreadCountClient{stream}
	return x, nil
}

type ProfilerCommandService_CommandStreamActiveThreadCountClient interface {
	Send(*PCmdActiveThreadCountRes) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type profilerCommandServiceCommandStreamActiveThreadCountClient struct {
	grpc.ClientStream
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountClient) Send(m *PCmdActiveThreadCountRes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profilerCommandServiceClient) CommandActiveThreadDump(ctx context.Context, in *PCmdActiveThreadDumpRes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandActiveThreadDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerCommandServiceClient) CommandActiveThreadLightDump(ctx context.Context, in *PCmdActiveThreadLightDumpRes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandActiveThreadLightDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilerCommandServiceServer is the server API for ProfilerCommandService service.
type ProfilerCommandServiceServer interface {
	HandleCommand(ProfilerCommandService_HandleCommandServer) error
	CommandEcho(context.Context, *PCmdEchoResponse) (*empty.Empty, error)
	CommandStreamActiveThreadCount(ProfilerCommandService_CommandStreamActiveThreadCountServer) error
	CommandActiveThreadDump(context.Context, *PCmdActiveThreadDumpRes) (*empty.Empty, error)
	CommandActiveThreadLightDump(context.Context, *PCmdActiveThreadLightDumpRes) (*empty.Empty, error)
}

// UnimplementedProfilerCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfilerCommandServiceServer struct {
}

func (*UnimplementedProfilerCommandServiceServer) HandleCommand(srv ProfilerCommandService_HandleCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleCommand not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandEcho(ctx context.Context, req *PCmdEchoResponse) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandEcho not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandStreamActiveThreadCount(srv ProfilerCommandService_CommandStreamActiveThreadCountServer) error {
	return status.Errorf(codes.Unimplemented, "method CommandStreamActiveThreadCount not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandActiveThreadDump(ctx context.Context, req *PCmdActiveThreadDumpRes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandActiveThreadDump not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandActiveThreadLightDump(ctx context.Context, req *PCmdActiveThreadLightDumpRes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandActiveThreadLightDump not implemented")
}

func RegisterProfilerCommandServiceServer(s *grpc.Server, srv ProfilerCommandServiceServer) {
	s.RegisterService(&_ProfilerCommandService_serviceDesc, srv)
}

func _ProfilerCommandService_HandleCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilerCommandServiceServer).HandleCommand(&profilerCommandServiceHandleCommandServer{stream})
}

type ProfilerCommandService_HandleCommandServer interface {
	Send(*PCmdRequest) error
	Recv() (*PCmdMessage, error)
	grpc.ServerStream
}

type profilerCommandServiceHandleCommandServer struct {
	grpc.ServerStream
}

func (x *profilerCommandServiceHandleCommandServer) Send(m *PCmdRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandServer) Recv() (*PCmdMessage, error) {
	m := new(PCmdMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfilerCommandService_CommandEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdEchoResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandEcho(ctx, req.(*PCmdEchoResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerCommandService_CommandStreamActiveThreadCount_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilerCommandServiceServer).CommandStreamActiveThreadCount(&profilerCommandServiceCommandStreamActiveThreadCountServer{stream})
}

type ProfilerCommandService_CommandStreamActiveThreadCountServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PCmdActiveThreadCountRes, error)
	grpc.ServerStream
}

type profilerCommandServiceCommandStreamActiveThreadCountServer struct {
	grpc.ServerStream
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountServer) Recv() (*PCmdActiveThreadCountRes, error) {
	m := new(PCmdActiveThreadCountRes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfilerCommandService_CommandActiveThreadDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdActiveThreadDumpRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandActiveThreadDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadDump(ctx, req.(*PCmdActiveThreadDumpRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerCommandService_CommandActiveThreadLightDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdActiveThreadLightDumpRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadLightDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandActiveThreadLightDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadLightDump(ctx, req.(*PCmdActiveThreadLightDumpRes))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfilerCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProfilerCommandService",
	HandlerType: (*ProfilerCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommandEcho",
			Handler:    _ProfilerCommandService_CommandEcho_Handler,
		},
		{
			MethodName: "CommandActiveThreadDump",
			Handler:    _ProfilerCommandService_CommandActiveThreadDump_Handler,
		},
		{
			MethodName: "CommandActiveThreadLightDump",
			Handler:    _ProfilerCommandService_CommandActiveThreadLightDump_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HandleCommand",
			Handler:       _ProfilerCommandService_HandleCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CommandStreamActiveThreadCount",
			Handler:       _ProfilerCommandService_CommandStreamActiveThreadCount_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Service.proto",
}
