// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: product/api/category/v1/category.proto

package v1

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
	ProductCategoryService_ListProductCategory_FullMethodName   = "/product.api.category.v1.ProductCategoryService/ListProductCategory"
	ProductCategoryService_GetProductCategory_FullMethodName    = "/product.api.category.v1.ProductCategoryService/GetProductCategory"
	ProductCategoryService_CreateProductCategory_FullMethodName = "/product.api.category.v1.ProductCategoryService/CreateProductCategory"
	ProductCategoryService_UpdateProductCategory_FullMethodName = "/product.api.category.v1.ProductCategoryService/UpdateProductCategory"
	ProductCategoryService_DeleteProductCategory_FullMethodName = "/product.api.category.v1.ProductCategoryService/DeleteProductCategory"
)

// ProductCategoryServiceClient is the client API for ProductCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCategoryServiceClient interface {
	ListProductCategory(ctx context.Context, in *ListProductCategoryRequest, opts ...grpc.CallOption) (*ListProductCategoryReply, error)
	GetProductCategory(ctx context.Context, in *GetProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategory, error)
	CreateProductCategory(ctx context.Context, in *CreateProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategory, error)
	UpdateProductCategory(ctx context.Context, in *UpdateProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategory, error)
	DeleteProductCategory(ctx context.Context, in *DeleteProductCategoryRequest, opts ...grpc.CallOption) (*DeleteProductCategoryReply, error)
}

type productCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCategoryServiceClient(cc grpc.ClientConnInterface) ProductCategoryServiceClient {
	return &productCategoryServiceClient{cc}
}

func (c *productCategoryServiceClient) ListProductCategory(ctx context.Context, in *ListProductCategoryRequest, opts ...grpc.CallOption) (*ListProductCategoryReply, error) {
	out := new(ListProductCategoryReply)
	err := c.cc.Invoke(ctx, ProductCategoryService_ListProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) GetProductCategory(ctx context.Context, in *GetProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategory, error) {
	out := new(ProductCategory)
	err := c.cc.Invoke(ctx, ProductCategoryService_GetProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) CreateProductCategory(ctx context.Context, in *CreateProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategory, error) {
	out := new(ProductCategory)
	err := c.cc.Invoke(ctx, ProductCategoryService_CreateProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) UpdateProductCategory(ctx context.Context, in *UpdateProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategory, error) {
	out := new(ProductCategory)
	err := c.cc.Invoke(ctx, ProductCategoryService_UpdateProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCategoryServiceClient) DeleteProductCategory(ctx context.Context, in *DeleteProductCategoryRequest, opts ...grpc.CallOption) (*DeleteProductCategoryReply, error) {
	out := new(DeleteProductCategoryReply)
	err := c.cc.Invoke(ctx, ProductCategoryService_DeleteProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCategoryServiceServer is the server API for ProductCategoryService service.
// All implementations should embed UnimplementedProductCategoryServiceServer
// for forward compatibility
type ProductCategoryServiceServer interface {
	ListProductCategory(context.Context, *ListProductCategoryRequest) (*ListProductCategoryReply, error)
	GetProductCategory(context.Context, *GetProductCategoryRequest) (*ProductCategory, error)
	CreateProductCategory(context.Context, *CreateProductCategoryRequest) (*ProductCategory, error)
	UpdateProductCategory(context.Context, *UpdateProductCategoryRequest) (*ProductCategory, error)
	DeleteProductCategory(context.Context, *DeleteProductCategoryRequest) (*DeleteProductCategoryReply, error)
}

// UnimplementedProductCategoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductCategoryServiceServer struct {
}

func (UnimplementedProductCategoryServiceServer) ListProductCategory(context.Context, *ListProductCategoryRequest) (*ListProductCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductCategory not implemented")
}
func (UnimplementedProductCategoryServiceServer) GetProductCategory(context.Context, *GetProductCategoryRequest) (*ProductCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCategory not implemented")
}
func (UnimplementedProductCategoryServiceServer) CreateProductCategory(context.Context, *CreateProductCategoryRequest) (*ProductCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductCategory not implemented")
}
func (UnimplementedProductCategoryServiceServer) UpdateProductCategory(context.Context, *UpdateProductCategoryRequest) (*ProductCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductCategory not implemented")
}
func (UnimplementedProductCategoryServiceServer) DeleteProductCategory(context.Context, *DeleteProductCategoryRequest) (*DeleteProductCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductCategory not implemented")
}

// UnsafeProductCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCategoryServiceServer will
// result in compilation errors.
type UnsafeProductCategoryServiceServer interface {
	mustEmbedUnimplementedProductCategoryServiceServer()
}

func RegisterProductCategoryServiceServer(s grpc.ServiceRegistrar, srv ProductCategoryServiceServer) {
	s.RegisterService(&ProductCategoryService_ServiceDesc, srv)
}

func _ProductCategoryService_ListProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).ListProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_ListProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).ListProductCategory(ctx, req.(*ListProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_GetProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).GetProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_GetProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).GetProductCategory(ctx, req.(*GetProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_CreateProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).CreateProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_CreateProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).CreateProductCategory(ctx, req.(*CreateProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_UpdateProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).UpdateProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_UpdateProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).UpdateProductCategory(ctx, req.(*UpdateProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCategoryService_DeleteProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCategoryServiceServer).DeleteProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCategoryService_DeleteProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCategoryServiceServer).DeleteProductCategory(ctx, req.(*DeleteProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCategoryService_ServiceDesc is the grpc.ServiceDesc for ProductCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.api.category.v1.ProductCategoryService",
	HandlerType: (*ProductCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProductCategory",
			Handler:    _ProductCategoryService_ListProductCategory_Handler,
		},
		{
			MethodName: "GetProductCategory",
			Handler:    _ProductCategoryService_GetProductCategory_Handler,
		},
		{
			MethodName: "CreateProductCategory",
			Handler:    _ProductCategoryService_CreateProductCategory_Handler,
		},
		{
			MethodName: "UpdateProductCategory",
			Handler:    _ProductCategoryService_UpdateProductCategory_Handler,
		},
		{
			MethodName: "DeleteProductCategory",
			Handler:    _ProductCategoryService_DeleteProductCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/api/category/v1/category.proto",
}
