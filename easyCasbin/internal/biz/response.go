package biz

type UserInfoResponse struct {
	Id       int64
	Mobile   string
	NickName string
}

type LoginResponse struct {
	User      UserInfoResponse `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
}