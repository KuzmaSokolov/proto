// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LoggerAPI_LogError_FullMethodName = "/proto.LoggerAPI/LogError"
)

// LoggerAPIClient is the client API for LoggerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerAPIClient interface {
	LogError(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type loggerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerAPIClient(cc grpc.ClientConnInterface) LoggerAPIClient {
	return &loggerAPIClient{cc}
}

func (c *loggerAPIClient) LogError(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, LoggerAPI_LogError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerAPIServer is the server API for LoggerAPI service.
// All implementations must embed UnimplementedLoggerAPIServer
// for forward compatibility
type LoggerAPIServer interface {
	LogError(context.Context, *LogRequest) (*LogResponse, error)
	mustEmbedUnimplementedLoggerAPIServer()
}

// UnimplementedLoggerAPIServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerAPIServer struct {
}

func (UnimplementedLoggerAPIServer) LogError(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogError not implemented")
}
func (UnimplementedLoggerAPIServer) mustEmbedUnimplementedLoggerAPIServer() {}

// UnsafeLoggerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerAPIServer will
// result in compilation errors.
type UnsafeLoggerAPIServer interface {
	mustEmbedUnimplementedLoggerAPIServer()
}

func RegisterLoggerAPIServer(s grpc.ServiceRegistrar, srv LoggerAPIServer) {
	s.RegisterService(&LoggerAPI_ServiceDesc, srv)
}

func _LoggerAPI_LogError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerAPIServer).LogError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoggerAPI_LogError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerAPIServer).LogError(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoggerAPI_ServiceDesc is the grpc.ServiceDesc for LoggerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LoggerAPI",
	HandlerType: (*LoggerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogError",
			Handler:    _LoggerAPI_LogError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
