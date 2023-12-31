// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.19.3
// source: user/v1/user.proto

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

const OperationUserCreateUser = "/api.user.v1.User/CreateUser"
const OperationUserGetUserById = "/api.user.v1.User/GetUserById"
const OperationUserGetUserByMobile = "/api.user.v1.User/GetUserByMobile"
const OperationUserGetUserList = "/api.user.v1.User/GetUserList"
const OperationUserLogin = "/api.user.v1.User/Login"
const OperationUserRegisterUser = "/api.user.v1.User/RegisterUser"
const OperationUserUpdateUser = "/api.user.v1.User/UpdateUser"

type UserHTTPServer interface {
	// CreateUser 创建用户
	CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
	// GetUserById 通过 Id 查询用户
	GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
	// GetUserByMobile 通过 mobile 查询用户
	GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
	// GetUserList 获取用户列表
	GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
	// Login 登录
	Login(context.Context, *LoginRequest) (*LoginRpl, error)
	// RegisterUser 注册
	RegisterUser(context.Context, *RegisterRequest) (*emptypb.Empty, error)
	// UpdateUser 更新用户
	UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/easyCasbin/api/v1/users", _User_GetUserList0_HTTP_Handler(srv))
	r.GET("/easyCasbin/api/v1/user/{mobile}", _User_GetUserByMobile0_HTTP_Handler(srv))
	r.GET("/easyCasbin/api/v1/user/{id}", _User_GetUserById0_HTTP_Handler(srv))
	r.POST("/easyCasbin/api/v1/user", _User_CreateUser0_HTTP_Handler(srv))
	r.PUT("/easyCasbin/api/v1/user", _User_UpdateUser0_HTTP_Handler(srv))
	r.POST("/easyCasbin/api/v1/user/register", _User_RegisterUser0_HTTP_Handler(srv))
	r.POST("/easyCasbin/api/v1/user/login", _User_Login0_HTTP_Handler(srv))
}

func _User_GetUserList0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageInfo
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserList(ctx, req.(*PageInfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserListResponse)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserByMobile0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MobileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserByMobile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserByMobile(ctx, req.(*MobileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserById0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserById(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserInfo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserInfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserInfo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserInfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_RegisterUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegisterUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterUser(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginRpl)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserInfo, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	GetUserById(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	GetUserByMobile(ctx context.Context, req *MobileRequest, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	GetUserList(ctx context.Context, req *PageInfo, opts ...http.CallOption) (rsp *UserListResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginRpl, err error)
	RegisterUser(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdateUser(ctx context.Context, req *UpdateUserInfo, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/easyCasbin/api/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserById(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/easyCasbin/api/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/easyCasbin/api/v1/user/{mobile}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserByMobile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserList(ctx context.Context, in *PageInfo, opts ...http.CallOption) (*UserListResponse, error) {
	var out UserListResponse
	pattern := "/easyCasbin/api/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginRpl, error) {
	var out LoginRpl
	pattern := "/easyCasbin/api/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/easyCasbin/api/v1/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegisterUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/easyCasbin/api/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
