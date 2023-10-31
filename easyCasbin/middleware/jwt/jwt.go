package jwt

import (
	"context"
	"time"

	kjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"

	"easyCasbin/internal/conf"
)

type JWT struct {
	SigningKey []byte
	C          *conf.Server
}

type CustomClaims struct {
	ID       int64  `json:"id"`
	NickName string `json:"nickname"`
	jwt.RegisteredClaims
}

func (j *JWT) CreateClaims(id int64, nickname string) jwt.MapClaims {
	return jwt.MapClaims{
		"id":       id,
		"nickname": nickname,
	}
}

func (j *JWT) CreateToken(claims jwt.MapClaims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256,
		CustomClaims{
			claims["id"].(int64),
			claims["nickname"].(string),
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+j.C.Jwt.ExpiresTime, 0)),
				Issuer:    j.C.Jwt.Issuer,
				NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-1000, 0)),
			},
		}).SignedString([]byte(j.C.Jwt.SigningKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// FromContext 从 ctx 中解析token 内容， kratos 中间件会将token 内容解析至 ctx 中
func FromContext(ctx context.Context) (CustomClaims, error) {
	token, _ := kjwt.FromContext(ctx)
	c, _ := token.(jwt.MapClaims)
	return CustomClaims{
		ID:       int64(c["id"].(float64)),
		NickName: c["nickname"].(string),
	}, nil
}
