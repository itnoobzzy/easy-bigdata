package biz

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
	Id         int64
	Name       string
	Domain     string
	CreateTime int64
	UpdateTime int64
	DeleteTime int64
}
