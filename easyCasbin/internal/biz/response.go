package biz

import "time"

type UserInfoResponse struct {
	Id       int64
	Mobile   string
	Username string
}

type LoginResponse struct {
	User      UserInfoResponse `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
}

/*
域角色 response
*/

type DomainRoleResponse struct {
	Id         uint
	Name       string
	Domain     string
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime time.Time
}
