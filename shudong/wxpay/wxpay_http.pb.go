// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.5.0
// source: wxpay.proto

package wxpay

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

const OperationWexinPaySvrPayCallback = "/wxpay.WexinPaySvr/PayCallback"
const OperationWexinPaySvrTransactions = "/wxpay.WexinPaySvr/Transactions"

type WexinPaySvrHTTPServer interface {
	// PayCallback PayCallback 交易结果回调
	PayCallback(context.Context, *PayCallbackReq) (*PayCallbackRsp, error)
	// Transactions Transactions 交易接口
	Transactions(context.Context, *TransactionsReq) (*TransactionsRsp, error)
}

func RegisterWexinPaySvrHTTPServer(s *http.Server, srv WexinPaySvrHTTPServer) {
	r := s.Route("/")
	r.POST("/shudong/wxpay/weixin_pay_svr/transactions", _WexinPaySvr_Transactions0_HTTP_Handler(srv))
	r.POST("/shudong/wxpay/weixin_pay_svr/pay_callback", _WexinPaySvr_PayCallback0_HTTP_Handler(srv))
}

func _WexinPaySvr_Transactions0_HTTP_Handler(srv WexinPaySvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TransactionsReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWexinPaySvrTransactions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Transactions(ctx, req.(*TransactionsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TransactionsRsp)
		return ctx.Result(200, reply)
	}
}

func _WexinPaySvr_PayCallback0_HTTP_Handler(srv WexinPaySvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PayCallbackReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWexinPaySvrPayCallback)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PayCallback(ctx, req.(*PayCallbackReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PayCallbackRsp)
		return ctx.Result(200, reply)
	}
}

type WexinPaySvrHTTPClient interface {
	PayCallback(ctx context.Context, req *PayCallbackReq, opts ...http.CallOption) (rsp *PayCallbackRsp, err error)
	Transactions(ctx context.Context, req *TransactionsReq, opts ...http.CallOption) (rsp *TransactionsRsp, err error)
}

type WexinPaySvrHTTPClientImpl struct {
	cc *http.Client
}

func NewWexinPaySvrHTTPClient(client *http.Client) WexinPaySvrHTTPClient {
	return &WexinPaySvrHTTPClientImpl{client}
}

func (c *WexinPaySvrHTTPClientImpl) PayCallback(ctx context.Context, in *PayCallbackReq, opts ...http.CallOption) (*PayCallbackRsp, error) {
	var out PayCallbackRsp
	pattern := "/shudong/wxpay/weixin_pay_svr/pay_callback"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWexinPaySvrPayCallback))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WexinPaySvrHTTPClientImpl) Transactions(ctx context.Context, in *TransactionsReq, opts ...http.CallOption) (*TransactionsRsp, error) {
	var out TransactionsRsp
	pattern := "/shudong/wxpay/weixin_pay_svr/transactions"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWexinPaySvrTransactions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
