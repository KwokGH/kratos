// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.17.3
// source: account/account.proto

package account

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

const OperationAccountLogin = "/api.account.Account/Login"
const OperationAccountRegister = "/api.account.Account/Register"

type AccountHTTPServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.POST("/account/login", _Account_Login1_HTTP_Handler(srv))
	r.POST("/account/register", _Account_Register0_HTTP_Handler(srv))
}

func _Account_Login1_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResp)
		return ctx.Result(200, reply)
	}
}

func _Account_Register0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResp)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginResp, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterResp, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginResp, error) {
	var out LoginResp
	pattern := "/account/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterResp, error) {
	var out RegisterResp
	pattern := "/account/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
