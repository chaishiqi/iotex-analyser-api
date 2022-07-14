// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// XRC721ServiceClient is the client API for XRC721Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XRC721ServiceClient interface {
	//XRC721ByAddress returns Xrc721 actions given the sender address or recipient address
	XRC721ByAddress(ctx context.Context, in *XRC721ByAddressRequest, opts ...grpc.CallOption) (*XRC721ByAddressResponse, error)
	//XRC721ByContractAddress returns Xrc721 actions given the Xrc721 contract address
	XRC721ByContractAddress(ctx context.Context, in *XRC721ByContractAddressRequest, opts ...grpc.CallOption) (*XRC721ByContractAddressResponse, error)
	//XRC721ByPage returns Xrc721 actions by pagination
	XRC721ByPage(ctx context.Context, in *XRC721ByPageRequest, opts ...grpc.CallOption) (*XRC721ByPageResponse, error)
	//XRC721Addresses returns Xrc721 contract addresses
	XRC721Addresses(ctx context.Context, in *XRC721AddressesRequest, opts ...grpc.CallOption) (*XRC721AddressesResponse, error)
	//XRC721TokenHolderAddresses returns Xrc721 token holder addresses given a Xrc721 contract address
	XRC721TokenHolderAddresses(ctx context.Context, in *XRC721TokenHolderAddressesRequest, opts ...grpc.CallOption) (*XRC721TokenHolderAddressesResponse, error)
}

type xRC721ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewXRC721ServiceClient(cc grpc.ClientConnInterface) XRC721ServiceClient {
	return &xRC721ServiceClient{cc}
}

func (c *xRC721ServiceClient) XRC721ByAddress(ctx context.Context, in *XRC721ByAddressRequest, opts ...grpc.CallOption) (*XRC721ByAddressResponse, error) {
	out := new(XRC721ByAddressResponse)
	err := c.cc.Invoke(ctx, "/api.XRC721Service/XRC721ByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xRC721ServiceClient) XRC721ByContractAddress(ctx context.Context, in *XRC721ByContractAddressRequest, opts ...grpc.CallOption) (*XRC721ByContractAddressResponse, error) {
	out := new(XRC721ByContractAddressResponse)
	err := c.cc.Invoke(ctx, "/api.XRC721Service/XRC721ByContractAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xRC721ServiceClient) XRC721ByPage(ctx context.Context, in *XRC721ByPageRequest, opts ...grpc.CallOption) (*XRC721ByPageResponse, error) {
	out := new(XRC721ByPageResponse)
	err := c.cc.Invoke(ctx, "/api.XRC721Service/XRC721ByPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xRC721ServiceClient) XRC721Addresses(ctx context.Context, in *XRC721AddressesRequest, opts ...grpc.CallOption) (*XRC721AddressesResponse, error) {
	out := new(XRC721AddressesResponse)
	err := c.cc.Invoke(ctx, "/api.XRC721Service/XRC721Addresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xRC721ServiceClient) XRC721TokenHolderAddresses(ctx context.Context, in *XRC721TokenHolderAddressesRequest, opts ...grpc.CallOption) (*XRC721TokenHolderAddressesResponse, error) {
	out := new(XRC721TokenHolderAddressesResponse)
	err := c.cc.Invoke(ctx, "/api.XRC721Service/XRC721TokenHolderAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XRC721ServiceServer is the server API for XRC721Service service.
// All implementations must embed UnimplementedXRC721ServiceServer
// for forward compatibility
type XRC721ServiceServer interface {
	//XRC721ByAddress returns Xrc721 actions given the sender address or recipient address
	XRC721ByAddress(context.Context, *XRC721ByAddressRequest) (*XRC721ByAddressResponse, error)
	//XRC721ByContractAddress returns Xrc721 actions given the Xrc721 contract address
	XRC721ByContractAddress(context.Context, *XRC721ByContractAddressRequest) (*XRC721ByContractAddressResponse, error)
	//XRC721ByPage returns Xrc721 actions by pagination
	XRC721ByPage(context.Context, *XRC721ByPageRequest) (*XRC721ByPageResponse, error)
	//XRC721Addresses returns Xrc721 contract addresses
	XRC721Addresses(context.Context, *XRC721AddressesRequest) (*XRC721AddressesResponse, error)
	//XRC721TokenHolderAddresses returns Xrc721 token holder addresses given a Xrc721 contract address
	XRC721TokenHolderAddresses(context.Context, *XRC721TokenHolderAddressesRequest) (*XRC721TokenHolderAddressesResponse, error)
	mustEmbedUnimplementedXRC721ServiceServer()
}

// UnimplementedXRC721ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedXRC721ServiceServer struct {
}

func (UnimplementedXRC721ServiceServer) XRC721ByAddress(context.Context, *XRC721ByAddressRequest) (*XRC721ByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XRC721ByAddress not implemented")
}
func (UnimplementedXRC721ServiceServer) XRC721ByContractAddress(context.Context, *XRC721ByContractAddressRequest) (*XRC721ByContractAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XRC721ByContractAddress not implemented")
}
func (UnimplementedXRC721ServiceServer) XRC721ByPage(context.Context, *XRC721ByPageRequest) (*XRC721ByPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XRC721ByPage not implemented")
}
func (UnimplementedXRC721ServiceServer) XRC721Addresses(context.Context, *XRC721AddressesRequest) (*XRC721AddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XRC721Addresses not implemented")
}
func (UnimplementedXRC721ServiceServer) XRC721TokenHolderAddresses(context.Context, *XRC721TokenHolderAddressesRequest) (*XRC721TokenHolderAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XRC721TokenHolderAddresses not implemented")
}
func (UnimplementedXRC721ServiceServer) mustEmbedUnimplementedXRC721ServiceServer() {}

// UnsafeXRC721ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XRC721ServiceServer will
// result in compilation errors.
type UnsafeXRC721ServiceServer interface {
	mustEmbedUnimplementedXRC721ServiceServer()
}

func RegisterXRC721ServiceServer(s grpc.ServiceRegistrar, srv XRC721ServiceServer) {
	s.RegisterService(&XRC721Service_ServiceDesc, srv)
}

func _XRC721Service_XRC721ByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XRC721ByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XRC721ServiceServer).XRC721ByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.XRC721Service/XRC721ByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XRC721ServiceServer).XRC721ByAddress(ctx, req.(*XRC721ByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XRC721Service_XRC721ByContractAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XRC721ByContractAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XRC721ServiceServer).XRC721ByContractAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.XRC721Service/XRC721ByContractAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XRC721ServiceServer).XRC721ByContractAddress(ctx, req.(*XRC721ByContractAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XRC721Service_XRC721ByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XRC721ByPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XRC721ServiceServer).XRC721ByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.XRC721Service/XRC721ByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XRC721ServiceServer).XRC721ByPage(ctx, req.(*XRC721ByPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XRC721Service_XRC721Addresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XRC721AddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XRC721ServiceServer).XRC721Addresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.XRC721Service/XRC721Addresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XRC721ServiceServer).XRC721Addresses(ctx, req.(*XRC721AddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XRC721Service_XRC721TokenHolderAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XRC721TokenHolderAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XRC721ServiceServer).XRC721TokenHolderAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.XRC721Service/XRC721TokenHolderAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XRC721ServiceServer).XRC721TokenHolderAddresses(ctx, req.(*XRC721TokenHolderAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// XRC721Service_ServiceDesc is the grpc.ServiceDesc for XRC721Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XRC721Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.XRC721Service",
	HandlerType: (*XRC721ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "XRC721ByAddress",
			Handler:    _XRC721Service_XRC721ByAddress_Handler,
		},
		{
			MethodName: "XRC721ByContractAddress",
			Handler:    _XRC721Service_XRC721ByContractAddress_Handler,
		},
		{
			MethodName: "XRC721ByPage",
			Handler:    _XRC721Service_XRC721ByPage_Handler,
		},
		{
			MethodName: "XRC721Addresses",
			Handler:    _XRC721Service_XRC721Addresses_Handler,
		},
		{
			MethodName: "XRC721TokenHolderAddresses",
			Handler:    _XRC721Service_XRC721TokenHolderAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_xrc721.proto",
}