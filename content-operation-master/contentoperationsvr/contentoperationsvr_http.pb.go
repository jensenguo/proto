// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.5.0
// source: contentoperationsvr.proto

package contentoperationsvr

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

const OperationContentOperationSvrAnswerArticle = "/contentoperationsvr.ContentOperationSvr/AnswerArticle"
const OperationContentOperationSvrAnswerWoa = "/contentoperationsvr.ContentOperationSvr/AnswerWoa"
const OperationContentOperationSvrFollowWoa = "/contentoperationsvr.ContentOperationSvr/FollowWoa"
const OperationContentOperationSvrGetArticleReaders = "/contentoperationsvr.ContentOperationSvr/GetArticleReaders"
const OperationContentOperationSvrGetMyArticles = "/contentoperationsvr.ContentOperationSvr/GetMyArticles"
const OperationContentOperationSvrGetMyWoas = "/contentoperationsvr.ContentOperationSvr/GetMyWoas"
const OperationContentOperationSvrGetShowArticles = "/contentoperationsvr.ContentOperationSvr/GetShowArticles"
const OperationContentOperationSvrGetUserInfo = "/contentoperationsvr.ContentOperationSvr/GetUserInfo"
const OperationContentOperationSvrGetWoaFans = "/contentoperationsvr.ContentOperationSvr/GetWoaFans"
const OperationContentOperationSvrLogin = "/contentoperationsvr.ContentOperationSvr/Login"
const OperationContentOperationSvrPublishArticle = "/contentoperationsvr.ContentOperationSvr/PublishArticle"
const OperationContentOperationSvrPublishWoa = "/contentoperationsvr.ContentOperationSvr/PublishWoa"
const OperationContentOperationSvrReadArticle = "/contentoperationsvr.ContentOperationSvr/ReadArticle"
const OperationContentOperationSvrSetArticleStatus = "/contentoperationsvr.ContentOperationSvr/SetArticleStatus"
const OperationContentOperationSvrSetWoaStatus = "/contentoperationsvr.ContentOperationSvr/SetWoaStatus"

type ContentOperationSvrHTTPServer interface {
	// AnswerArticle AnswerArticle 回复文章阅读
	AnswerArticle(context.Context, *AnswerArticleReq) (*AnswerArticleRsp, error)
	// AnswerWoa AnswerWoa 回复公众号
	AnswerWoa(context.Context, *AnswerWoaReq) (*AnswerWoaRsp, error)
	// FollowWoa FollowWoa 关注公众号
	FollowWoa(context.Context, *FollowWoaReq) (*FollowWoaRsp, error)
	// GetArticleReaders GetArticleReaders 查询文章阅读列表
	GetArticleReaders(context.Context, *GetArticleReadersReq) (*GetArticleReadersRsp, error)
	// GetMyArticles GetMyArticles 查询我的文章列表
	GetMyArticles(context.Context, *GetMyArticlesReq) (*GetMyArticlesRsp, error)
	// GetMyWoas GetMyWoas 查询我的公众号列表
	GetMyWoas(context.Context, *GetMyWoasReq) (*GetMyWoasRsp, error)
	// GetShowArticles GetShowArticles 查询外显文章
	GetShowArticles(context.Context, *GetShowArticlesReq) (*GetShowArticlesRsp, error)
	// GetUserInfo GetUserInfo 查询用户信息
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error)
	// GetWoaFans GetWoaFans 查询公众号关注列表
	GetWoaFans(context.Context, *GetWoaFansReq) (*GetWoaFansRsp, error)
	// Login Login 用户授权登录
	Login(context.Context, *LoginReq) (*LoginRsp, error)
	// PublishArticle PublishArticle 发布文章
	PublishArticle(context.Context, *PublishArticleReq) (*PublishArticleRsp, error)
	// PublishWoa PublishWoa 发布公众号
	PublishWoa(context.Context, *PublishWoaReq) (*PublishWoaRsp, error)
	// ReadArticle ReadArticle 开始读文章
	ReadArticle(context.Context, *ReadArticleReq) (*ReadArticleRsp, error)
	// SetArticleStatus SetArticleStatus 设置文章状态
	SetArticleStatus(context.Context, *SetArticleStatusReq) (*SetArticleStatusRsp, error)
	// SetWoaStatus SetWoaStatus 设置文章状态
	SetWoaStatus(context.Context, *SetWoaStatusReq) (*SetWoaStatusRsp, error)
}

func RegisterContentOperationSvrHTTPServer(s *http.Server, srv ContentOperationSvrHTTPServer) {
	r := s.Route("/")
	r.POST("/content-operation-proxy/content-operation-svr/login", _ContentOperationSvr_Login0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/get-user-info", _ContentOperationSvr_GetUserInfo0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/publish-article", _ContentOperationSvr_PublishArticle0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/set-article-status", _ContentOperationSvr_SetArticleStatus0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/get-my-articles", _ContentOperationSvr_GetMyArticles0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/get-article-readers", _ContentOperationSvr_GetArticleReaders0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/read-article", _ContentOperationSvr_ReadArticle0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/answer-article", _ContentOperationSvr_AnswerArticle0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/get-show-article", _ContentOperationSvr_GetShowArticles0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/publish-woa", _ContentOperationSvr_PublishWoa0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/set-woa-status", _ContentOperationSvr_SetWoaStatus0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/get-my-woas", _ContentOperationSvr_GetMyWoas0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/get-woa-fans", _ContentOperationSvr_GetWoaFans0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/read-article", _ContentOperationSvr_FollowWoa0_HTTP_Handler(srv))
	r.POST("/content-operation-proxy/content-operation-svr/answer-article", _ContentOperationSvr_AnswerWoa0_HTTP_Handler(srv))
}

func _ContentOperationSvr_Login0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_GetUserInfo0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_PublishArticle0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PublishArticleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrPublishArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PublishArticle(ctx, req.(*PublishArticleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PublishArticleRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_SetArticleStatus0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetArticleStatusReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrSetArticleStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetArticleStatus(ctx, req.(*SetArticleStatusReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetArticleStatusRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_GetMyArticles0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMyArticlesReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrGetMyArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyArticles(ctx, req.(*GetMyArticlesReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMyArticlesRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_GetArticleReaders0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleReadersReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrGetArticleReaders)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticleReaders(ctx, req.(*GetArticleReadersReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArticleReadersRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_ReadArticle0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReadArticleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrReadArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReadArticle(ctx, req.(*ReadArticleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReadArticleRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_AnswerArticle0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnswerArticleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrAnswerArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AnswerArticle(ctx, req.(*AnswerArticleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnswerArticleRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_GetShowArticles0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetShowArticlesReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrGetShowArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetShowArticles(ctx, req.(*GetShowArticlesReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetShowArticlesRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_PublishWoa0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PublishWoaReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrPublishWoa)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PublishWoa(ctx, req.(*PublishWoaReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PublishWoaRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_SetWoaStatus0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetWoaStatusReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrSetWoaStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetWoaStatus(ctx, req.(*SetWoaStatusReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetWoaStatusRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_GetMyWoas0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMyWoasReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrGetMyWoas)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyWoas(ctx, req.(*GetMyWoasReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMyWoasRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_GetWoaFans0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWoaFansReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrGetWoaFans)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWoaFans(ctx, req.(*GetWoaFansReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWoaFansRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_FollowWoa0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FollowWoaReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrFollowWoa)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FollowWoa(ctx, req.(*FollowWoaReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FollowWoaRsp)
		return ctx.Result(200, reply)
	}
}

func _ContentOperationSvr_AnswerWoa0_HTTP_Handler(srv ContentOperationSvrHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnswerWoaReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContentOperationSvrAnswerWoa)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AnswerWoa(ctx, req.(*AnswerWoaReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnswerWoaRsp)
		return ctx.Result(200, reply)
	}
}

type ContentOperationSvrHTTPClient interface {
	AnswerArticle(ctx context.Context, req *AnswerArticleReq, opts ...http.CallOption) (rsp *AnswerArticleRsp, err error)
	AnswerWoa(ctx context.Context, req *AnswerWoaReq, opts ...http.CallOption) (rsp *AnswerWoaRsp, err error)
	FollowWoa(ctx context.Context, req *FollowWoaReq, opts ...http.CallOption) (rsp *FollowWoaRsp, err error)
	GetArticleReaders(ctx context.Context, req *GetArticleReadersReq, opts ...http.CallOption) (rsp *GetArticleReadersRsp, err error)
	GetMyArticles(ctx context.Context, req *GetMyArticlesReq, opts ...http.CallOption) (rsp *GetMyArticlesRsp, err error)
	GetMyWoas(ctx context.Context, req *GetMyWoasReq, opts ...http.CallOption) (rsp *GetMyWoasRsp, err error)
	GetShowArticles(ctx context.Context, req *GetShowArticlesReq, opts ...http.CallOption) (rsp *GetShowArticlesRsp, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoReq, opts ...http.CallOption) (rsp *GetUserInfoRsp, err error)
	GetWoaFans(ctx context.Context, req *GetWoaFansReq, opts ...http.CallOption) (rsp *GetWoaFansRsp, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginRsp, err error)
	PublishArticle(ctx context.Context, req *PublishArticleReq, opts ...http.CallOption) (rsp *PublishArticleRsp, err error)
	PublishWoa(ctx context.Context, req *PublishWoaReq, opts ...http.CallOption) (rsp *PublishWoaRsp, err error)
	ReadArticle(ctx context.Context, req *ReadArticleReq, opts ...http.CallOption) (rsp *ReadArticleRsp, err error)
	SetArticleStatus(ctx context.Context, req *SetArticleStatusReq, opts ...http.CallOption) (rsp *SetArticleStatusRsp, err error)
	SetWoaStatus(ctx context.Context, req *SetWoaStatusReq, opts ...http.CallOption) (rsp *SetWoaStatusRsp, err error)
}

type ContentOperationSvrHTTPClientImpl struct {
	cc *http.Client
}

func NewContentOperationSvrHTTPClient(client *http.Client) ContentOperationSvrHTTPClient {
	return &ContentOperationSvrHTTPClientImpl{client}
}

func (c *ContentOperationSvrHTTPClientImpl) AnswerArticle(ctx context.Context, in *AnswerArticleReq, opts ...http.CallOption) (*AnswerArticleRsp, error) {
	var out AnswerArticleRsp
	pattern := "/content-operation-proxy/content-operation-svr/answer-article"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrAnswerArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) AnswerWoa(ctx context.Context, in *AnswerWoaReq, opts ...http.CallOption) (*AnswerWoaRsp, error) {
	var out AnswerWoaRsp
	pattern := "/content-operation-proxy/content-operation-svr/answer-article"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrAnswerWoa))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) FollowWoa(ctx context.Context, in *FollowWoaReq, opts ...http.CallOption) (*FollowWoaRsp, error) {
	var out FollowWoaRsp
	pattern := "/content-operation-proxy/content-operation-svr/read-article"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrFollowWoa))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) GetArticleReaders(ctx context.Context, in *GetArticleReadersReq, opts ...http.CallOption) (*GetArticleReadersRsp, error) {
	var out GetArticleReadersRsp
	pattern := "/content-operation-proxy/content-operation-svr/get-article-readers"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrGetArticleReaders))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) GetMyArticles(ctx context.Context, in *GetMyArticlesReq, opts ...http.CallOption) (*GetMyArticlesRsp, error) {
	var out GetMyArticlesRsp
	pattern := "/content-operation-proxy/content-operation-svr/get-my-articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrGetMyArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) GetMyWoas(ctx context.Context, in *GetMyWoasReq, opts ...http.CallOption) (*GetMyWoasRsp, error) {
	var out GetMyWoasRsp
	pattern := "/content-operation-proxy/content-operation-svr/get-my-woas"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrGetMyWoas))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) GetShowArticles(ctx context.Context, in *GetShowArticlesReq, opts ...http.CallOption) (*GetShowArticlesRsp, error) {
	var out GetShowArticlesRsp
	pattern := "/content-operation-proxy/content-operation-svr/get-show-article"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrGetShowArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...http.CallOption) (*GetUserInfoRsp, error) {
	var out GetUserInfoRsp
	pattern := "/content-operation-proxy/content-operation-svr/get-user-info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) GetWoaFans(ctx context.Context, in *GetWoaFansReq, opts ...http.CallOption) (*GetWoaFansRsp, error) {
	var out GetWoaFansRsp
	pattern := "/content-operation-proxy/content-operation-svr/get-woa-fans"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrGetWoaFans))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginRsp, error) {
	var out LoginRsp
	pattern := "/content-operation-proxy/content-operation-svr/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) PublishArticle(ctx context.Context, in *PublishArticleReq, opts ...http.CallOption) (*PublishArticleRsp, error) {
	var out PublishArticleRsp
	pattern := "/content-operation-proxy/content-operation-svr/publish-article"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrPublishArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) PublishWoa(ctx context.Context, in *PublishWoaReq, opts ...http.CallOption) (*PublishWoaRsp, error) {
	var out PublishWoaRsp
	pattern := "/content-operation-proxy/content-operation-svr/publish-woa"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrPublishWoa))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) ReadArticle(ctx context.Context, in *ReadArticleReq, opts ...http.CallOption) (*ReadArticleRsp, error) {
	var out ReadArticleRsp
	pattern := "/content-operation-proxy/content-operation-svr/read-article"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrReadArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) SetArticleStatus(ctx context.Context, in *SetArticleStatusReq, opts ...http.CallOption) (*SetArticleStatusRsp, error) {
	var out SetArticleStatusRsp
	pattern := "/content-operation-proxy/content-operation-svr/set-article-status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrSetArticleStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ContentOperationSvrHTTPClientImpl) SetWoaStatus(ctx context.Context, in *SetWoaStatusReq, opts ...http.CallOption) (*SetWoaStatusRsp, error) {
	var out SetWoaStatusRsp
	pattern := "/content-operation-proxy/content-operation-svr/set-woa-status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContentOperationSvrSetWoaStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}