package data

import (
	"context"
	v1 "easyCasbin/api/role/v1"
	"easyCasbin/internal/biz"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

func TestRoleRepo_AddDomainRole(t *testing.T) {
	ctx := context.Background()
	now := time.Now().Unix()
	roleName := fmt.Sprintf("admin_%v", now)
	domain := fmt.Sprintf("domain_%v", now)

	repo := NewRoleRepo(&Data{
		db:  tdb,
		rdb: trdb,
	}, log.DefaultLogger)

	cases := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
		ctx     context.Context
		role    *biz.Role
	}{
		{
			name: "normal",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.NoError(t, err)
				return true
			},
			ctx: ctx,
			role: &biz.Role{
				Name:   roleName,
				Domain: domain,
			},
		}, {
			name: "domain role exist",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.ErrorIs(t, err, v1.ErrorDomainRoleExist("domain role exist"))
				return false
			},
			ctx: ctx,
			role: &biz.Role{
				Name:   roleName,
				Domain: domain,
			},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			role, err := repo.AddDomainRole(item.ctx, item.role.Domain, item.role.Name)
			if !item.wantErr(t, err) {
				return
			}
			assert.NotNil(t, role.CreatedAt)
			assert.Equal(t, item.role.Domain, role.Domain)
		})
	}
	repo.DeleteDomainRole(ctx, domain, roleName)
}

func TestRoleRepo_CheckDomainRole(t *testing.T) {
	ctx := context.Background()
	now := time.Now().Unix()
	roleName := fmt.Sprintf("admin_%v", now)

	repo := NewRoleRepo(&Data{
		db:  tdb,
		rdb: trdb,
	}, log.DefaultLogger)

	cases := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
		ctx     context.Context
		role    *biz.Role
	}{
		{
			name: "normal",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.NoError(t, err)
				return true
			},
			ctx: ctx,
			role: &biz.Role{
				Name:   roleName,
				Domain: "test_domain1",
			},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			exist, err := repo.CheckDomainRole(item.ctx, item.role.Domain, item.role.Name)
			if !item.wantErr(t, err) {
				return
			}
			if reflect.TypeOf(exist).Kind() != reflect.Bool {
				t.Errorf("Expected a bool, but got %T", exist)
			}
		})
	}
}

func TestRoleRepo_UpdateDomainRole(t *testing.T) {
	ctx := context.Background()
	now := time.Now().Unix()
	domain := fmt.Sprintf("test_domain_%v", now)
	oldRole := fmt.Sprintf("test_old_role_%v", now)
	newRole := fmt.Sprintf("test_new_role_%v", now)

	repo := NewRoleRepo(&Data{
		db:  tdb,
		rdb: trdb,
	}, log.DefaultLogger)

	repo.AddDomainRole(ctx, domain, oldRole)

	cases := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
		ctx     context.Context
		domain  string
		oldRole string
		newRole string
	}{
		{
			name: "normal",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.NoError(t, err)
				return true
			},
			ctx:     ctx,
			domain:  domain,
			oldRole: oldRole,
			newRole: newRole,
		}, {
			name: "new role name exist",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.ErrorIs(t, err, v1.ErrorInternalErr("new role name exist"))
				return false
			},
			ctx:     ctx,
			domain:  domain,
			oldRole: oldRole,
			newRole: newRole,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {

			if item.name == "new role name exist" {
				repo.AddDomainRole(ctx, domain, newRole)
			}

			newRole, err := repo.UpdateDomainRole(ctx, item.domain, item.oldRole, item.newRole)
			if !item.wantErr(t, err) {
				return
			}
			assert.Equal(t, item.newRole, newRole.Name)
		})
	}
	repo.DeleteDomainRole(ctx, domain, oldRole)
	repo.DeleteDomainRole(ctx, domain, newRole)
}
