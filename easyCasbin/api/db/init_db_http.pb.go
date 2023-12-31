// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.19.3
// source: db/init_db.proto

package InitDB

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

const OperationInitDBInitRoleDB = "/db.init.InitDB/InitRoleDB"
const OperationInitDBInitUserDB = "/db.init.InitDB/InitUserDB"

type InitDBHTTPServer interface {
	InitRoleDB(context.Context, *emptypb.Empty) (*InitRpl, error)
	InitUserDB(context.Context, *emptypb.Empty) (*InitRpl, error)
}

func RegisterInitDBHTTPServer(s *http.Server, srv InitDBHTTPServer) {
	r := s.Route("/")
	r.GET("/easyCasbin/initDb/user", _InitDB_InitUserDB0_HTTP_Handler(srv))
	r.GET("/easyCasbin/initDb/role", _InitDB_InitRoleDB0_HTTP_Handler(srv))
}

func _InitDB_InitUserDB0_HTTP_Handler(srv InitDBHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationInitDBInitUserDB)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InitUserDB(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InitRpl)
		return ctx.Result(200, reply)
	}
}

func _InitDB_InitRoleDB0_HTTP_Handler(srv InitDBHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationInitDBInitRoleDB)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InitRoleDB(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InitRpl)
		return ctx.Result(200, reply)
	}
}

type InitDBHTTPClient interface {
	InitRoleDB(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *InitRpl, err error)
	InitUserDB(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *InitRpl, err error)
}

type InitDBHTTPClientImpl struct {
	cc *http.Client
}

func NewInitDBHTTPClient(client *http.Client) InitDBHTTPClient {
	return &InitDBHTTPClientImpl{client}
}

func (c *InitDBHTTPClientImpl) InitRoleDB(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*InitRpl, error) {
	var out InitRpl
	pattern := "/easyCasbin/initDb/role"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationInitDBInitRoleDB))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *InitDBHTTPClientImpl) InitUserDB(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*InitRpl, error) {
	var out InitRpl
	pattern := "/easyCasbin/initDb/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationInitDBInitUserDB))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
