// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.19.3
// source: casbin_rule/v1/casbin_rule.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCasbinRuleGetAllSubjects = "/api.casbin_rule.v1.CasbinRule/GetAllSubjects"

type CasbinRuleHTTPServer interface {
	// GetAllSubjects 获取所有鉴权主体
	GetAllSubjects(context.Context, *emptypb.Empty) (*GetAllSubjectsRpl, error)
}

func RegisterCasbinRuleHTTPServer(s *http.Server, srv CasbinRuleHTTPServer) {
	r := s.Route("/")
	r.GET("/easyCasbin/api/casbin_rule/v1/subs", _CasbinRule_GetAllSubjects0_HTTP_Handler(srv))
}

func _CasbinRule_GetAllSubjects0_HTTP_Handler(srv CasbinRuleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCasbinRuleGetAllSubjects)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAllSubjects(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAllSubjectsRpl)
		return ctx.Result(200, reply)
	}
}

type CasbinRuleHTTPClient interface {
	GetAllSubjects(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetAllSubjectsRpl, err error)
}

type CasbinRuleHTTPClientImpl struct {
	cc *http.Client
}

func NewCasbinRuleHTTPClient(client *http.Client) CasbinRuleHTTPClient {
	return &CasbinRuleHTTPClientImpl{client}
}

func (c *CasbinRuleHTTPClientImpl) GetAllSubjects(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetAllSubjectsRpl, error) {
	var out GetAllSubjectsRpl
	pattern := "/easyCasbin/api/casbin_rule/v1/subs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCasbinRuleGetAllSubjects))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
