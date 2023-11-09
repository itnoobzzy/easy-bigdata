package data

import (
	"context"
	"easyCasbin/internal/conf"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCasbinRuleRepo_GetAllSubjects(t *testing.T) {
	ctx := context.Background()

	repo := NewCasbinRuleRepo(&Data{
		db:  tdb,
		rdb: trdb,
	}, &conf.Casbin{RbacConfPath: "../../casbin_rbac_domain.conf"}, log.DefaultLogger)

	cases := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
		ctx     context.Context
	}{
		{
			name: "normal",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.NoError(t, err)
				return true
			},
			ctx: ctx,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			subs, err := repo.GetAllSubjects(ctx)
			if !item.wantErr(t, err) {
				return
			}
			fmt.Println(subs)
			//assert.Equal(t, item.newRole, newRole.Name)
		})
	}
}
