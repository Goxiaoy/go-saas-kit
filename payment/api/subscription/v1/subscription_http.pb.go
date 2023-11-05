// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: payment/api/subscription/v1/subscription.proto

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

const OperationSubscriptionServiceCancelMySubscription = "/payment.api.subscription.v1.SubscriptionService/CancelMySubscription"
const OperationSubscriptionServiceCancelSubscription = "/payment.api.subscription.v1.SubscriptionService/CancelSubscription"
const OperationSubscriptionServiceCreateMySubscription = "/payment.api.subscription.v1.SubscriptionService/CreateMySubscription"
const OperationSubscriptionServiceCreateSubscription = "/payment.api.subscription.v1.SubscriptionService/CreateSubscription"
const OperationSubscriptionServiceGetMySubscription = "/payment.api.subscription.v1.SubscriptionService/GetMySubscription"
const OperationSubscriptionServiceGetSubscription = "/payment.api.subscription.v1.SubscriptionService/GetSubscription"
const OperationSubscriptionServiceListMySubscription = "/payment.api.subscription.v1.SubscriptionService/ListMySubscription"
const OperationSubscriptionServiceListSubscription = "/payment.api.subscription.v1.SubscriptionService/ListSubscription"
const OperationSubscriptionServiceUpdateMySubscription = "/payment.api.subscription.v1.SubscriptionService/UpdateMySubscription"
const OperationSubscriptionServiceUpdateSubscription = "/payment.api.subscription.v1.SubscriptionService/UpdateSubscription"

type SubscriptionServiceHTTPServer interface {
	CancelMySubscription(context.Context, *CancelSubscriptionRequest) (*Subscription, error)
	CancelSubscription(context.Context, *CancelSubscriptionRequest) (*Subscription, error)
	CreateMySubscription(context.Context, *CreateMySubscriptionRequest) (*Subscription, error)
	CreateSubscription(context.Context, *CreateSubscriptionRequest) (*Subscription, error)
	GetMySubscription(context.Context, *GetMySubscriptionRequest) (*Subscription, error)
	GetSubscription(context.Context, *GetSubscriptionRequest) (*Subscription, error)
	ListMySubscription(context.Context, *ListMySubscriptionRequest) (*ListMySubscriptionReply, error)
	ListSubscription(context.Context, *ListSubscriptionRequest) (*ListSubscriptionReply, error)
	UpdateMySubscription(context.Context, *UpdateMySubscriptionRequest) (*Subscription, error)
	UpdateSubscription(context.Context, *UpdateSubscriptionRequest) (*Subscription, error)
}

func RegisterSubscriptionServiceHTTPServer(s *http.Server, srv SubscriptionServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/subscription", _SubscriptionService_CreateSubscription0_HTTP_Handler(srv))
	r.PUT("/v1/subscription/{subscription.id}", _SubscriptionService_UpdateSubscription0_HTTP_Handler(srv))
	r.POST("/v1/subscription/list", _SubscriptionService_ListSubscription0_HTTP_Handler(srv))
	r.GET("/v1/subscriptions", _SubscriptionService_ListSubscription1_HTTP_Handler(srv))
	r.GET("/v1/subscription/{id}", _SubscriptionService_GetSubscription0_HTTP_Handler(srv))
	r.POST("/v1/subscription/{id}/cancel", _SubscriptionService_CancelSubscription0_HTTP_Handler(srv))
	r.POST("/v1/subscription/my", _SubscriptionService_CreateMySubscription0_HTTP_Handler(srv))
	r.POST("/v1/subscription/my/{id}", _SubscriptionService_GetMySubscription0_HTTP_Handler(srv))
	r.POST("/v1/subscription/my/{id}/cancel", _SubscriptionService_CancelMySubscription0_HTTP_Handler(srv))
	r.PUT("/v1/subscription/my/{subscription.id}", _SubscriptionService_UpdateMySubscription0_HTTP_Handler(srv))
	r.POST("/v1/subscription/my/list", _SubscriptionService_ListMySubscription0_HTTP_Handler(srv))
	r.GET("/v1/subscriptions/my", _SubscriptionService_ListMySubscription1_HTTP_Handler(srv))
}

func _SubscriptionService_CreateSubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceCreateSubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSubscription(ctx, req.(*CreateSubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_UpdateSubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceUpdateSubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSubscription(ctx, req.(*UpdateSubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_ListSubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceListSubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSubscription(ctx, req.(*ListSubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSubscriptionReply)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_ListSubscription1_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSubscriptionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceListSubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSubscription(ctx, req.(*ListSubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSubscriptionReply)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_GetSubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSubscriptionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceGetSubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSubscription(ctx, req.(*GetSubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_CancelSubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CancelSubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceCancelSubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CancelSubscription(ctx, req.(*CancelSubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_CreateMySubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMySubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceCreateMySubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMySubscription(ctx, req.(*CreateMySubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_GetMySubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMySubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceGetMySubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMySubscription(ctx, req.(*GetMySubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_CancelMySubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CancelSubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceCancelMySubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CancelMySubscription(ctx, req.(*CancelSubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_UpdateMySubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMySubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceUpdateMySubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMySubscription(ctx, req.(*UpdateMySubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Subscription)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_ListMySubscription0_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMySubscriptionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceListMySubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMySubscription(ctx, req.(*ListMySubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMySubscriptionReply)
		return ctx.Result(200, reply)
	}
}

func _SubscriptionService_ListMySubscription1_HTTP_Handler(srv SubscriptionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMySubscriptionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubscriptionServiceListMySubscription)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMySubscription(ctx, req.(*ListMySubscriptionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMySubscriptionReply)
		return ctx.Result(200, reply)
	}
}

type SubscriptionServiceHTTPClient interface {
	CancelMySubscription(ctx context.Context, req *CancelSubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
	CancelSubscription(ctx context.Context, req *CancelSubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
	CreateMySubscription(ctx context.Context, req *CreateMySubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
	CreateSubscription(ctx context.Context, req *CreateSubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
	GetMySubscription(ctx context.Context, req *GetMySubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
	GetSubscription(ctx context.Context, req *GetSubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
	ListMySubscription(ctx context.Context, req *ListMySubscriptionRequest, opts ...http.CallOption) (rsp *ListMySubscriptionReply, err error)
	ListSubscription(ctx context.Context, req *ListSubscriptionRequest, opts ...http.CallOption) (rsp *ListSubscriptionReply, err error)
	UpdateMySubscription(ctx context.Context, req *UpdateMySubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
	UpdateSubscription(ctx context.Context, req *UpdateSubscriptionRequest, opts ...http.CallOption) (rsp *Subscription, err error)
}

type SubscriptionServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewSubscriptionServiceHTTPClient(client *http.Client) SubscriptionServiceHTTPClient {
	return &SubscriptionServiceHTTPClientImpl{client}
}

func (c *SubscriptionServiceHTTPClientImpl) CancelMySubscription(ctx context.Context, in *CancelSubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription/my/{id}/cancel"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubscriptionServiceCancelMySubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) CancelSubscription(ctx context.Context, in *CancelSubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription/{id}/cancel"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubscriptionServiceCancelSubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) CreateMySubscription(ctx context.Context, in *CreateMySubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription/my"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubscriptionServiceCreateMySubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubscriptionServiceCreateSubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) GetMySubscription(ctx context.Context, in *GetMySubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription/my/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubscriptionServiceGetMySubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSubscriptionServiceGetSubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) ListMySubscription(ctx context.Context, in *ListMySubscriptionRequest, opts ...http.CallOption) (*ListMySubscriptionReply, error) {
	var out ListMySubscriptionReply
	pattern := "/v1/subscriptions/my"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSubscriptionServiceListMySubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) ListSubscription(ctx context.Context, in *ListSubscriptionRequest, opts ...http.CallOption) (*ListSubscriptionReply, error) {
	var out ListSubscriptionReply
	pattern := "/v1/subscriptions"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSubscriptionServiceListSubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) UpdateMySubscription(ctx context.Context, in *UpdateMySubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription/my/{subscription.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubscriptionServiceUpdateMySubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SubscriptionServiceHTTPClientImpl) UpdateSubscription(ctx context.Context, in *UpdateSubscriptionRequest, opts ...http.CallOption) (*Subscription, error) {
	var out Subscription
	pattern := "/v1/subscription/{subscription.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubscriptionServiceUpdateSubscription))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
