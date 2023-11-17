package biz

type UserInfoResponse struct {
	Id       int32
	Mobile   string
	Username string
}

type LoginResponse struct {
	User      UserInfoResponse `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int32            `json:"expiresAt"`
}

/*
域角色 response
*/

type DomainRoleResponse struct {
	Id         int32
	Name       string
	Domain     string
	CreateTime int32
	UpdateTime int32
	DeleteTime int32
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

type DomainAuthResponse struct {
	Rules [][]string        `json:"rules"`
	Roles map[string]string `json:"roles"`
	Total int32             `json:"total"`
}
