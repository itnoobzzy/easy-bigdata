package biz

import (
	v1 "easyCasbin/api/user/v1"
)

type LoginRequest struct {
	username string
	password string
}

// NewLoginRequest 校验登录参数，返回通过校验后的登录请求结构体
func NewLoginRequest(username, password string) (*LoginRequest, error) {
	if username == "" {
		return nil, v1.ErrorPasswordErr("username or password error!")
	}
	if password == "" {
		return nil, v1.ErrorPasswordErr("username or password error!")
	}
	return &LoginRequest{
		username: username,
		password: password,
	}, nil
}
