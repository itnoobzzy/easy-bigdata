package biz

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"easyCasbin/internal/conf"
)

func TestUserUseCase_login(t *testing.T) {
	controller := gomock.NewController(t)
	repo := NewMockUserRepo(controller)
	encryptService := NewMockEncryptService(controller)

	userUseCase := NewUserUsecase(repo, log.DefaultLogger, &conf.Server{
		Jwt: &conf.Server_JWT{
			SigningKey:  "d5cfc646-3692-4c98-98b3-ca7b8553d289",
			ExpiresTime: 604800,
			BufferTime:  86400,
			Issuer:      "easyCasbin",
		},
	})

	data := []struct {
		name      string
		mockFunc  func()
		wantErr   assert.ErrorAssertionFunc
		wantToken string
		ctx       context.Context
		req       func() *LoginRequest
	}{
		{
			name: "normal",
			mockFunc: func() {
				user := &User{
					ID:       123,
					Username: "zzy",
					Password: "123",
				}
				repo.EXPECT().GetUserByName(gomock.Any(), "zzy").Return(user, nil).Times(1)
				encryptService.EXPECT().CheckPassword(gomock.Any(), "123", "123").Return(true).Times(1)
				encryptService.EXPECT().Token(gomock.Any(), user).Return("token string", nil).Times(1)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.NoError(t, err)
				return true
			},
			wantToken: "token string",
			ctx:       context.Background(),
			req: func() *LoginRequest {
				r, _ := NewLoginRequest("zzy", "123")
				return r
			},
		},
	}
	for _, item := range data {
		t.Run(item.name, func(t *testing.T) {
			item.mockFunc()
			loginRequest := item.req()
			got, err := userUseCase.Login(item.ctx, loginRequest, encryptService)
			if !item.wantErr(t, err) {
				return
			}
			assert.Equal(t, item.wantToken, got.Token)
		})
	}

}
