package data

import (
	"context"
	"easyCasbin/internal/conf"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"

	"easyCasbin/internal/biz"
)

type CasbinRuleRepo struct {
	data *Data
	log  *log.Helper

	enforcer *casbin.Enforcer
}

func NewCasbinRuleRepo(data *Data, casbinConf *conf.Casbin, logger log.Logger) biz.CasbinRuleRepo {

	a, _ := gormadapter.NewAdapterByDB(data.db)
	m, err := model.NewModelFromFile(casbinConf.RbacConfPath)
	if err != nil {
		log.NewHelper(logger).Error("closing the data resources")
	}
	enforcer, _ := casbin.NewEnforcer(m, a)
	_ = enforcer.LoadPolicy()

	return &CasbinRuleRepo{
		data:     data,
		log:      log.NewHelper(logger),
		enforcer: enforcer,
	}
}

func (c *CasbinRuleRepo) GetAllSubjects(ctx context.Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *CasbinRuleRepo) GetRolesForUserInDomain(ctx context.Context, username, domain string) (roles []string, err error) {
	//TODO implement me
	panic("implement me")
}
