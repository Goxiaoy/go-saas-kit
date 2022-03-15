// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user/api/auth/v1/web.proto

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

// AuthWebClient is the client API for AuthWeb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthWebClient interface {
	GetWebLogin(ctx context.Context, in *GetLoginRequest, opts ...grpc.CallOption) (*GetLoginResponse, error)
	WebLogin(ctx context.Context, in *WebLoginAuthRequest, opts ...grpc.CallOption) (*WebLoginAuthReply, error)
	GetWebLogout(ctx context.Context, in *GetLogoutRequest, opts ...grpc.CallOption) (*GetLogoutResponse, error)
	WebLogout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GetConsent(ctx context.Context, in *GetConsentRequest, opts ...grpc.CallOption) (*GetConsentResponse, error)
	GrantConsent(ctx context.Context, in *GrantConsentRequest, opts ...grpc.CallOption) (*GrantConsentResponse, error)
}

type authWebClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthWebClient(cc grpc.ClientConnInterface) AuthWebClient {
	return &authWebClient{cc}
}

func (c *authWebClient) GetWebLogin(ctx context.Context, in *GetLoginRequest, opts ...grpc.CallOption) (*GetLoginResponse, error) {
	out := new(GetLoginResponse)
	err := c.cc.Invoke(ctx, "/user.api.auth.v1.AuthWeb/GetWebLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authWebClient) WebLogin(ctx context.Context, in *WebLoginAuthRequest, opts ...grpc.CallOption) (*WebLoginAuthReply, error) {
	out := new(WebLoginAuthReply)
	err := c.cc.Invoke(ctx, "/user.api.auth.v1.AuthWeb/WebLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authWebClient) GetWebLogout(ctx context.Context, in *GetLogoutRequest, opts ...grpc.CallOption) (*GetLogoutResponse, error) {
	out := new(GetLogoutResponse)
	err := c.cc.Invoke(ctx, "/user.api.auth.v1.AuthWeb/GetWebLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authWebClient) WebLogout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/user.api.auth.v1.AuthWeb/WebLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authWebClient) GetConsent(ctx context.Context, in *GetConsentRequest, opts ...grpc.CallOption) (*GetConsentResponse, error) {
	out := new(GetConsentResponse)
	err := c.cc.Invoke(ctx, "/user.api.auth.v1.AuthWeb/GetConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authWebClient) GrantConsent(ctx context.Context, in *GrantConsentRequest, opts ...grpc.CallOption) (*GrantConsentResponse, error) {
	out := new(GrantConsentResponse)
	err := c.cc.Invoke(ctx, "/user.api.auth.v1.AuthWeb/GrantConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthWebServer is the server API for AuthWeb service.
// All implementations must embed UnimplementedAuthWebServer
// for forward compatibility
type AuthWebServer interface {
	GetWebLogin(context.Context, *GetLoginRequest) (*GetLoginResponse, error)
	WebLogin(context.Context, *WebLoginAuthRequest) (*WebLoginAuthReply, error)
	GetWebLogout(context.Context, *GetLogoutRequest) (*GetLogoutResponse, error)
	WebLogout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GetConsent(context.Context, *GetConsentRequest) (*GetConsentResponse, error)
	GrantConsent(context.Context, *GrantConsentRequest) (*GrantConsentResponse, error)
	mustEmbedUnimplementedAuthWebServer()
}

// UnimplementedAuthWebServer must be embedded to have forward compatible implementations.
type UnimplementedAuthWebServer struct {
}

func (UnimplementedAuthWebServer) GetWebLogin(context.Context, *GetLoginRequest) (*GetLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebLogin not implemented")
}
func (UnimplementedAuthWebServer) WebLogin(context.Context, *WebLoginAuthRequest) (*WebLoginAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebLogin not implemented")
}
func (UnimplementedAuthWebServer) GetWebLogout(context.Context, *GetLogoutRequest) (*GetLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebLogout not implemented")
}
func (UnimplementedAuthWebServer) WebLogout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebLogout not implemented")
}
func (UnimplementedAuthWebServer) GetConsent(context.Context, *GetConsentRequest) (*GetConsentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsent not implemented")
}
func (UnimplementedAuthWebServer) GrantConsent(context.Context, *GrantConsentRequest) (*GrantConsentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantConsent not implemented")
}
func (UnimplementedAuthWebServer) mustEmbedUnimplementedAuthWebServer() {}

// UnsafeAuthWebServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthWebServer will
// result in compilation errors.
type UnsafeAuthWebServer interface {
	mustEmbedUnimplementedAuthWebServer()
}

func RegisterAuthWebServer(s grpc.ServiceRegistrar, srv AuthWebServer) {
	s.RegisterService(&AuthWeb_ServiceDesc, srv)
}

func _AuthWeb_GetWebLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthWebServer).GetWebLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.auth.v1.AuthWeb/GetWebLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthWebServer).GetWebLogin(ctx, req.(*GetLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthWeb_WebLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebLoginAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthWebServer).WebLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.auth.v1.AuthWeb/WebLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthWebServer).WebLogin(ctx, req.(*WebLoginAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthWeb_GetWebLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthWebServer).GetWebLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.auth.v1.AuthWeb/GetWebLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthWebServer).GetWebLogout(ctx, req.(*GetLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthWeb_WebLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthWebServer).WebLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.auth.v1.AuthWeb/WebLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthWebServer).WebLogout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthWeb_GetConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthWebServer).GetConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.auth.v1.AuthWeb/GetConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthWebServer).GetConsent(ctx, req.(*GetConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthWeb_GrantConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthWebServer).GrantConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.auth.v1.AuthWeb/GrantConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthWebServer).GrantConsent(ctx, req.(*GrantConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthWeb_ServiceDesc is the grpc.ServiceDesc for AuthWeb service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthWeb_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.auth.v1.AuthWeb",
	HandlerType: (*AuthWebServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWebLogin",
			Handler:    _AuthWeb_GetWebLogin_Handler,
		},
		{
			MethodName: "WebLogin",
			Handler:    _AuthWeb_WebLogin_Handler,
		},
		{
			MethodName: "GetWebLogout",
			Handler:    _AuthWeb_GetWebLogout_Handler,
		},
		{
			MethodName: "WebLogout",
			Handler:    _AuthWeb_WebLogout_Handler,
		},
		{
			MethodName: "GetConsent",
			Handler:    _AuthWeb_GetConsent_Handler,
		},
		{
			MethodName: "GrantConsent",
			Handler:    _AuthWeb_GrantConsent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/api/auth/v1/web.proto",
}
