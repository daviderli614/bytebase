// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/vcs_connector_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VCSConnectorService_CreateVCSConnector_FullMethodName = "/bytebase.v1.VCSConnectorService/CreateVCSConnector"
	VCSConnectorService_GetVCSConnector_FullMethodName    = "/bytebase.v1.VCSConnectorService/GetVCSConnector"
	VCSConnectorService_ListVCSConnectors_FullMethodName  = "/bytebase.v1.VCSConnectorService/ListVCSConnectors"
	VCSConnectorService_UpdateVCSConnector_FullMethodName = "/bytebase.v1.VCSConnectorService/UpdateVCSConnector"
	VCSConnectorService_DeleteVCSConnector_FullMethodName = "/bytebase.v1.VCSConnectorService/DeleteVCSConnector"
)

// VCSConnectorServiceClient is the client API for VCSConnectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VCSConnectorServiceClient interface {
	CreateVCSConnector(ctx context.Context, in *CreateVCSConnectorRequest, opts ...grpc.CallOption) (*VCSConnector, error)
	GetVCSConnector(ctx context.Context, in *GetVCSConnectorRequest, opts ...grpc.CallOption) (*VCSConnector, error)
	ListVCSConnectors(ctx context.Context, in *ListVCSConnectorsRequest, opts ...grpc.CallOption) (*ListVCSConnectorsResponse, error)
	UpdateVCSConnector(ctx context.Context, in *UpdateVCSConnectorRequest, opts ...grpc.CallOption) (*VCSConnector, error)
	DeleteVCSConnector(ctx context.Context, in *DeleteVCSConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type vCSConnectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVCSConnectorServiceClient(cc grpc.ClientConnInterface) VCSConnectorServiceClient {
	return &vCSConnectorServiceClient{cc}
}

func (c *vCSConnectorServiceClient) CreateVCSConnector(ctx context.Context, in *CreateVCSConnectorRequest, opts ...grpc.CallOption) (*VCSConnector, error) {
	out := new(VCSConnector)
	err := c.cc.Invoke(ctx, VCSConnectorService_CreateVCSConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectorServiceClient) GetVCSConnector(ctx context.Context, in *GetVCSConnectorRequest, opts ...grpc.CallOption) (*VCSConnector, error) {
	out := new(VCSConnector)
	err := c.cc.Invoke(ctx, VCSConnectorService_GetVCSConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectorServiceClient) ListVCSConnectors(ctx context.Context, in *ListVCSConnectorsRequest, opts ...grpc.CallOption) (*ListVCSConnectorsResponse, error) {
	out := new(ListVCSConnectorsResponse)
	err := c.cc.Invoke(ctx, VCSConnectorService_ListVCSConnectors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectorServiceClient) UpdateVCSConnector(ctx context.Context, in *UpdateVCSConnectorRequest, opts ...grpc.CallOption) (*VCSConnector, error) {
	out := new(VCSConnector)
	err := c.cc.Invoke(ctx, VCSConnectorService_UpdateVCSConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectorServiceClient) DeleteVCSConnector(ctx context.Context, in *DeleteVCSConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VCSConnectorService_DeleteVCSConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VCSConnectorServiceServer is the server API for VCSConnectorService service.
// All implementations must embed UnimplementedVCSConnectorServiceServer
// for forward compatibility
type VCSConnectorServiceServer interface {
	CreateVCSConnector(context.Context, *CreateVCSConnectorRequest) (*VCSConnector, error)
	GetVCSConnector(context.Context, *GetVCSConnectorRequest) (*VCSConnector, error)
	ListVCSConnectors(context.Context, *ListVCSConnectorsRequest) (*ListVCSConnectorsResponse, error)
	UpdateVCSConnector(context.Context, *UpdateVCSConnectorRequest) (*VCSConnector, error)
	DeleteVCSConnector(context.Context, *DeleteVCSConnectorRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedVCSConnectorServiceServer()
}

// UnimplementedVCSConnectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVCSConnectorServiceServer struct {
}

func (UnimplementedVCSConnectorServiceServer) CreateVCSConnector(context.Context, *CreateVCSConnectorRequest) (*VCSConnector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVCSConnector not implemented")
}
func (UnimplementedVCSConnectorServiceServer) GetVCSConnector(context.Context, *GetVCSConnectorRequest) (*VCSConnector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVCSConnector not implemented")
}
func (UnimplementedVCSConnectorServiceServer) ListVCSConnectors(context.Context, *ListVCSConnectorsRequest) (*ListVCSConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVCSConnectors not implemented")
}
func (UnimplementedVCSConnectorServiceServer) UpdateVCSConnector(context.Context, *UpdateVCSConnectorRequest) (*VCSConnector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVCSConnector not implemented")
}
func (UnimplementedVCSConnectorServiceServer) DeleteVCSConnector(context.Context, *DeleteVCSConnectorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVCSConnector not implemented")
}
func (UnimplementedVCSConnectorServiceServer) mustEmbedUnimplementedVCSConnectorServiceServer() {}

// UnsafeVCSConnectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VCSConnectorServiceServer will
// result in compilation errors.
type UnsafeVCSConnectorServiceServer interface {
	mustEmbedUnimplementedVCSConnectorServiceServer()
}

func RegisterVCSConnectorServiceServer(s grpc.ServiceRegistrar, srv VCSConnectorServiceServer) {
	s.RegisterService(&VCSConnectorService_ServiceDesc, srv)
}

func _VCSConnectorService_CreateVCSConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVCSConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectorServiceServer).CreateVCSConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VCSConnectorService_CreateVCSConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectorServiceServer).CreateVCSConnector(ctx, req.(*CreateVCSConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnectorService_GetVCSConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVCSConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectorServiceServer).GetVCSConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VCSConnectorService_GetVCSConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectorServiceServer).GetVCSConnector(ctx, req.(*GetVCSConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnectorService_ListVCSConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVCSConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectorServiceServer).ListVCSConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VCSConnectorService_ListVCSConnectors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectorServiceServer).ListVCSConnectors(ctx, req.(*ListVCSConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnectorService_UpdateVCSConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVCSConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectorServiceServer).UpdateVCSConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VCSConnectorService_UpdateVCSConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectorServiceServer).UpdateVCSConnector(ctx, req.(*UpdateVCSConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnectorService_DeleteVCSConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVCSConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectorServiceServer).DeleteVCSConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VCSConnectorService_DeleteVCSConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectorServiceServer).DeleteVCSConnector(ctx, req.(*DeleteVCSConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VCSConnectorService_ServiceDesc is the grpc.ServiceDesc for VCSConnectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VCSConnectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.VCSConnectorService",
	HandlerType: (*VCSConnectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVCSConnector",
			Handler:    _VCSConnectorService_CreateVCSConnector_Handler,
		},
		{
			MethodName: "GetVCSConnector",
			Handler:    _VCSConnectorService_GetVCSConnector_Handler,
		},
		{
			MethodName: "ListVCSConnectors",
			Handler:    _VCSConnectorService_ListVCSConnectors_Handler,
		},
		{
			MethodName: "UpdateVCSConnector",
			Handler:    _VCSConnectorService_UpdateVCSConnector_Handler,
		},
		{
			MethodName: "DeleteVCSConnector",
			Handler:    _VCSConnectorService_DeleteVCSConnector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/vcs_connector_service.proto",
}