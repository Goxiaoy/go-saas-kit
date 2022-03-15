// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type AuthWebHTTPServer interface {
	GetConsent(context.Context, *GetConsentRequest) (*GetConsentResponse, error)
	GetWebLogin(context.Context, *GetLoginRequest) (*GetLoginResponse, error)
	GetWebLogout(context.Context, *GetLogoutRequest) (*GetLogoutResponse, error)
	GrantConsent(context.Context, *GrantConsentRequest) (*GrantConsentResponse, error)
	WebLogin(context.Context, *WebLoginAuthRequest) (*WebLoginAuthReply, error)
	WebLogout(context.Context, *LogoutRequest) (*LogoutResponse, error)
}

func RegisterAuthWebHTTPServer(s *http.Server, srv AuthWebHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/auth/web/login", _AuthWeb_GetWebLogin0_HTTP_Handler(srv))
	r.POST("/v1/auth/web/login", _AuthWeb_WebLogin0_HTTP_Handler(srv))
	r.GET("/v1/auth/web/logout", _AuthWeb_GetWebLogout0_HTTP_Handler(srv))
	r.POST("/v1/auth/web/logout", _AuthWeb_WebLogout0_HTTP_Handler(srv))
	r.GET("/v1/auth/web/consent", _AuthWeb_GetConsent0_HTTP_Handler(srv))
	r.GET("/v1/auth/web/consent", _AuthWeb_GrantConsent0_HTTP_Handler(srv))
}

func _AuthWeb_GetWebLogin0_HTTP_Handler(srv AuthWebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLoginRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.AuthWeb/GetWebLogin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWebLogin(ctx, req.(*GetLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLoginResponse)
		return ctx.Result(200, reply)
	}
}

func _AuthWeb_WebLogin0_HTTP_Handler(srv AuthWebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WebLoginAuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.AuthWeb/WebLogin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WebLogin(ctx, req.(*WebLoginAuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WebLoginAuthReply)
		return ctx.Result(200, reply)
	}
}

func _AuthWeb_GetWebLogout0_HTTP_Handler(srv AuthWebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLogoutRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.AuthWeb/GetWebLogout")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWebLogout(ctx, req.(*GetLogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLogoutResponse)
		return ctx.Result(200, reply)
	}
}

func _AuthWeb_WebLogout0_HTTP_Handler(srv AuthWebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.AuthWeb/WebLogout")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WebLogout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutResponse)
		return ctx.Result(200, reply)
	}
}

func _AuthWeb_GetConsent0_HTTP_Handler(srv AuthWebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetConsentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.AuthWeb/GetConsent")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetConsent(ctx, req.(*GetConsentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetConsentResponse)
		return ctx.Result(200, reply)
	}
}

func _AuthWeb_GrantConsent0_HTTP_Handler(srv AuthWebHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GrantConsentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.AuthWeb/GrantConsent")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GrantConsent(ctx, req.(*GrantConsentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GrantConsentResponse)
		return ctx.Result(200, reply)
	}
}

type AuthWebHTTPClient interface {
	GetConsent(ctx context.Context, req *GetConsentRequest, opts ...http.CallOption) (rsp *GetConsentResponse, err error)
	GetWebLogin(ctx context.Context, req *GetLoginRequest, opts ...http.CallOption) (rsp *GetLoginResponse, err error)
	GetWebLogout(ctx context.Context, req *GetLogoutRequest, opts ...http.CallOption) (rsp *GetLogoutResponse, err error)
	GrantConsent(ctx context.Context, req *GrantConsentRequest, opts ...http.CallOption) (rsp *GrantConsentResponse, err error)
	WebLogin(ctx context.Context, req *WebLoginAuthRequest, opts ...http.CallOption) (rsp *WebLoginAuthReply, err error)
	WebLogout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutResponse, err error)
}

type AuthWebHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthWebHTTPClient(client *http.Client) AuthWebHTTPClient {
	return &AuthWebHTTPClientImpl{client}
}

func (c *AuthWebHTTPClientImpl) GetConsent(ctx context.Context, in *GetConsentRequest, opts ...http.CallOption) (*GetConsentResponse, error) {
	var out GetConsentResponse
	pattern := "/v1/auth/web/consent"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.auth.v1.AuthWeb/GetConsent"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthWebHTTPClientImpl) GetWebLogin(ctx context.Context, in *GetLoginRequest, opts ...http.CallOption) (*GetLoginResponse, error) {
	var out GetLoginResponse
	pattern := "/v1/auth/web/login"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.auth.v1.AuthWeb/GetWebLogin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthWebHTTPClientImpl) GetWebLogout(ctx context.Context, in *GetLogoutRequest, opts ...http.CallOption) (*GetLogoutResponse, error) {
	var out GetLogoutResponse
	pattern := "/v1/auth/web/logout"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.auth.v1.AuthWeb/GetWebLogout"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthWebHTTPClientImpl) GrantConsent(ctx context.Context, in *GrantConsentRequest, opts ...http.CallOption) (*GrantConsentResponse, error) {
	var out GrantConsentResponse
	pattern := "/v1/auth/web/consent"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.auth.v1.AuthWeb/GrantConsent"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthWebHTTPClientImpl) WebLogin(ctx context.Context, in *WebLoginAuthRequest, opts ...http.CallOption) (*WebLoginAuthReply, error) {
	var out WebLoginAuthReply
	pattern := "/v1/auth/web/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.AuthWeb/WebLogin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthWebHTTPClientImpl) WebLogout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutResponse, error) {
	var out LogoutResponse
	pattern := "/v1/auth/web/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.AuthWeb/WebLogout"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
