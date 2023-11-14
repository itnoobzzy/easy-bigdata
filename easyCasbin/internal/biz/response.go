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

/*
casbin rule response
*/

type PermissionsResponse struct {
	Sub      string // 鉴权对象，用户或者角色
	Domain   string // 域
	Resource string // 资源
	Action   string // 操作
	Eft      string // 是否允许
}
