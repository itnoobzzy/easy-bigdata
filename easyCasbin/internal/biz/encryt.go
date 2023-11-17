package biz

import (
	"context"
	"easyCasbin/internal/conf"
	"easyCasbin/middleware/jwt"
	"easyCasbin/utils"
)

//go:generate mockgen -destination=encryt_mock.go -package=biz . EncryptService
type EncryptService interface {
	// CheckPassword 校验密码
	CheckPassword(ctx context.Context, pwd, encryptedPassword string) bool
	// Token 签发token
	Token(ctx context.Context, user *User) (string, error)
}

type encryptServiceImpl struct {
	sc *conf.Server
}

func NewEncryptService(serverConfig *conf.Server) EncryptService {
	return &encryptServiceImpl{
		sc: serverConfig,
	}
}

func (e *encryptServiceImpl) CheckPassword(ctx context.Context, pwd, encryptedPassword string) bool {
	return utils.BcryptCheck(pwd, encryptedPassword)
}

func (e *encryptServiceImpl) Token(ctx context.Context, user *User) (string, error) {
	j := jwt.JWT{C: e.sc}
	claims := j.CreateClaims(int64(user.ID), user.Username)
	token, _ := j.CreateToken(claims)
	return token, nil
}
