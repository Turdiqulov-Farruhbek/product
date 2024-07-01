// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: basket.proto

package genproto

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

// BasketServiceClient is the client API for BasketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasketServiceClient interface {
	CreateBasket(ctx context.Context, in *CreatedBasket, opts ...grpc.CallOption) (*B_Void, error)
	GetByID(ctx context.Context, in *ByIdBasket, opts ...grpc.CallOption) (*Basket, error)
	GetAll(ctx context.Context, in *FilterBasket, opts ...grpc.CallOption) (*GetAllBasketRes, error)
	UpdateBasket(ctx context.Context, in *ByIdBasket, opts ...grpc.CallOption) (*B_Void, error)
	DeleteBasket(ctx context.Context, in *ByIdBasket, opts ...grpc.CallOption) (*B_Void, error)
}

type basketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasketServiceClient(cc grpc.ClientConnInterface) BasketServiceClient {
	return &basketServiceClient{cc}
}

func (c *basketServiceClient) CreateBasket(ctx context.Context, in *CreatedBasket, opts ...grpc.CallOption) (*B_Void, error) {
	out := new(B_Void)
	err := c.cc.Invoke(ctx, "/product.BasketService/CreateBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) GetByID(ctx context.Context, in *ByIdBasket, opts ...grpc.CallOption) (*Basket, error) {
	out := new(Basket)
	err := c.cc.Invoke(ctx, "/product.BasketService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) GetAll(ctx context.Context, in *FilterBasket, opts ...grpc.CallOption) (*GetAllBasketRes, error) {
	out := new(GetAllBasketRes)
	err := c.cc.Invoke(ctx, "/product.BasketService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) UpdateBasket(ctx context.Context, in *ByIdBasket, opts ...grpc.CallOption) (*B_Void, error) {
	out := new(B_Void)
	err := c.cc.Invoke(ctx, "/product.BasketService/UpdateBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) DeleteBasket(ctx context.Context, in *ByIdBasket, opts ...grpc.CallOption) (*B_Void, error) {
	out := new(B_Void)
	err := c.cc.Invoke(ctx, "/product.BasketService/DeleteBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasketServiceServer is the server API for BasketService service.
// All implementations must embed UnimplementedBasketServiceServer
// for forward compatibility
type BasketServiceServer interface {
	CreateBasket(context.Context, *CreatedBasket) (*B_Void, error)
	GetByID(context.Context, *ByIdBasket) (*Basket, error)
	GetAll(context.Context, *FilterBasket) (*GetAllBasketRes, error)
	UpdateBasket(context.Context, *ByIdBasket) (*B_Void, error)
	DeleteBasket(context.Context, *ByIdBasket) (*B_Void, error)
	mustEmbedUnimplementedBasketServiceServer()
}

// UnimplementedBasketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBasketServiceServer struct {
}

func (UnimplementedBasketServiceServer) CreateBasket(context.Context, *CreatedBasket) (*B_Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBasket not implemented")
}
func (UnimplementedBasketServiceServer) GetByID(context.Context, *ByIdBasket) (*Basket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedBasketServiceServer) GetAll(context.Context, *FilterBasket) (*GetAllBasketRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBasketServiceServer) UpdateBasket(context.Context, *ByIdBasket) (*B_Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBasket not implemented")
}
func (UnimplementedBasketServiceServer) DeleteBasket(context.Context, *ByIdBasket) (*B_Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBasket not implemented")
}
func (UnimplementedBasketServiceServer) mustEmbedUnimplementedBasketServiceServer() {}

// UnsafeBasketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasketServiceServer will
// result in compilation errors.
type UnsafeBasketServiceServer interface {
	mustEmbedUnimplementedBasketServiceServer()
}

func RegisterBasketServiceServer(s grpc.ServiceRegistrar, srv BasketServiceServer) {
	s.RegisterService(&BasketService_ServiceDesc, srv)
}

func _BasketService_CreateBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).CreateBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.BasketService/CreateBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).CreateBasket(ctx, req.(*CreatedBasket))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.BasketService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).GetByID(ctx, req.(*ByIdBasket))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.BasketService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).GetAll(ctx, req.(*FilterBasket))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_UpdateBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).UpdateBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.BasketService/UpdateBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).UpdateBasket(ctx, req.(*ByIdBasket))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_DeleteBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdBasket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).DeleteBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.BasketService/DeleteBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).DeleteBasket(ctx, req.(*ByIdBasket))
	}
	return interceptor(ctx, in, info, handler)
}

// BasketService_ServiceDesc is the grpc.ServiceDesc for BasketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.BasketService",
	HandlerType: (*BasketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBasket",
			Handler:    _BasketService_CreateBasket_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _BasketService_GetByID_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BasketService_GetAll_Handler,
		},
		{
			MethodName: "UpdateBasket",
			Handler:    _BasketService_UpdateBasket_Handler,
		},
		{
			MethodName: "DeleteBasket",
			Handler:    _BasketService_DeleteBasket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basket.proto",
}