package biz

import (
	"context"
	v1 "easyCasbin/api/role/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDomainRoleUseCase_AddDomainRole(t *testing.T) {
	controller := gomock.NewController(t)
	repo := NewMockDomainRoleRepo(controller)

	useCase := NewDomainRoleUseCase(repo, log.DefaultLogger)

	data := []struct {
		name     string
		mockFunc func(domain, role string)
		err      error
		ctx      context.Context
		domain   string
		role     string
	}{
		{
			name: "normal",
			mockFunc: func(domain, role string) {
				domainRole := Role{Domain: domain, Name: role}
				repo.EXPECT().CheckDomainRole(gomock.Any(), domain, role).Return(false, nil).Times(1)
				repo.EXPECT().AddDomainRole(gomock.Any(), domain, role).Return(&domainRole, nil).Times(1)
			},
			err:    nil,
			ctx:    context.Background(),
			domain: "domain1",
			role:   "admin",
		}, {
			name: "domain role exist",
			mockFunc: func(domain, role string) {
				//domainRole := Role{Domain: domain, Name: role}
				repo.EXPECT().CheckDomainRole(gomock.Any(), domain, role).Return(true, nil).Times(1)
				//repo.EXPECT().AddDomainRole(gomock.Any(), domain, role).Return(&domainRole, nil).Times(1)
			},
			err:    v1.ErrorDomainRoleExist("exist"),
			ctx:    context.Background(),
			domain: "domain1",
			role:   "admin",
		},
	}
	for _, item := range data {
		t.Run(item.name, func(t *testing.T) {
			item.mockFunc(item.domain, item.role)
			got, err := useCase.AddDomainRole(item.ctx, item.domain, item.role)
			if err != nil {
				assert.ErrorIs(t, err, item.err)
			} else {
				assert.Equal(t, item.domain, got.Domain)
				assert.Equal(t, item.role, got.Name)
			}
		})
	}

}
