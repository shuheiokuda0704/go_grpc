// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/file.proto

package pb

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

// FIleServiceClient is the client API for FIleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FIleServiceClient interface {
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFIlesResponse, error)
}

type fIleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFIleServiceClient(cc grpc.ClientConnInterface) FIleServiceClient {
	return &fIleServiceClient{cc}
}

func (c *fIleServiceClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFIlesResponse, error) {
	out := new(ListFIlesResponse)
	err := c.cc.Invoke(ctx, "/file.FIleService/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FIleServiceServer is the server API for FIleService service.
// All implementations must embed UnimplementedFIleServiceServer
// for forward compatibility
type FIleServiceServer interface {
	ListFiles(context.Context, *ListFilesRequest) (*ListFIlesResponse, error)
	mustEmbedUnimplementedFIleServiceServer()
}

// UnimplementedFIleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFIleServiceServer struct {
}

func (UnimplementedFIleServiceServer) ListFiles(context.Context, *ListFilesRequest) (*ListFIlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFIleServiceServer) mustEmbedUnimplementedFIleServiceServer() {}

// UnsafeFIleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FIleServiceServer will
// result in compilation errors.
type UnsafeFIleServiceServer interface {
	mustEmbedUnimplementedFIleServiceServer()
}

func RegisterFIleServiceServer(s grpc.ServiceRegistrar, srv FIleServiceServer) {
	s.RegisterService(&FIleService_ServiceDesc, srv)
}

func _FIleService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FIleServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FIleService/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FIleServiceServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FIleService_ServiceDesc is the grpc.ServiceDesc for FIleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FIleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FIleService",
	HandlerType: (*FIleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FIleService_ListFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/file.proto",
}
