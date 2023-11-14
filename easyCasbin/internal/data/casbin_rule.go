package data

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"

	"easyCasbin/internal/biz"
	"easyCasbin/internal/conf"
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
	enforcer.EnableAutoBuildRoleLinks(true)
	_ = enforcer.LoadPolicy()

	return &CasbinRuleRepo{
		data:     data,
		log:      log.NewHelper(logger),
		enforcer: enforcer,
	}
}

// GetAllSubjects 获取所有鉴权主体
func (c *CasbinRuleRepo) GetAllSubjects(ctx context.Context) ([]string, error) {
	subs := c.enforcer.GetAllSubjects()
	return subs, nil
}

// GetRolesForUserInDomain 查询用户在域上的角色
func (c *CasbinRuleRepo) GetRolesForUserInDomain(ctx context.Context, username, domain string) (roles []string, err error) {
	roles = c.enforcer.GetRolesForUserInDomain(username, domain)
	return roles, nil
}

// GetImplicitUsersForRole 查询域角色下的所有用户
func (c *CasbinRuleRepo) GetImplicitUsersForRole(ctx context.Context, domain, role string) (users []string, err error) {
	return c.enforcer.GetImplicitUsersForRole(role, domain)
}

// AddRoleForUserInDomain 为用户添加域角色或者为角色继承另一个角色权限
func (c *CasbinRuleRepo) AddRoleForUserInDomain(ctx context.Context, user, domain, role string) (ok bool, err error) {
	return c.enforcer.AddRoleForUserInDomain(user, role, domain)
}

// DeleteRoleForUserInDomain 移除指定域角色下的用户或者取消指定域角色的继承
func (c *CasbinRuleRepo) DeleteRoleForUserInDomain(ctx context.Context, user, domain, role string) (ok bool, err error) {
	return c.enforcer.DeleteRoleForUserInDomain(user, role, domain)
}

// GetImplicitPermissionsForUser 查询鉴权主体（用户或角色）在对应域上的权限
func (c *CasbinRuleRepo) GetImplicitPermissionsForUser(ctx context.Context, username, domain string) (permissions [][]string, err error) {
	permissions, _ = c.enforcer.GetImplicitPermissionsForUser(username, domain)
	return permissions, nil
}

// GetAllPolicyInDomain 获取指定域的所有权限
func (c *CasbinRuleRepo) GetAllPolicyInDomain(ctx context.Context, domain string) (policies [][]string, err error) {
	var rules []biz.CasbinRule
	result := c.data.db.Where(&biz.CasbinRule{V1: domain}).Find(&rules)
	if result.Error != nil {
		return policies, err
	}
	for _, r := range rules {
		p := make([]string, 0, 5)
		p = append(p, r.V0, r.V1, r.V2, r.V3, r.V4)
		policies = append(policies, p)
	}
	return policies, nil
}

// DeleteDomain 删除域
func (c *CasbinRuleRepo) DeleteDomain(ctx context.Context, domain string) (bool, error) {
	ok, err := c.enforcer.DeleteDomains(domain)
	return ok, err
}

// BatchEnforce 批量校验权限
func (c *CasbinRuleRepo) BatchEnforce(ctx context.Context, rules [][]interface{}) ([]bool, error) {
	result, err := c.enforcer.BatchEnforce(rules)
	return result, err
}

// AddPolicies 添加权限
func (c *CasbinRuleRepo) AddPolicies(ctx context.Context, rules [][]string) (bool, error) {
	_, err := c.enforcer.AddPolicies(rules)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemovePolicies 删除权限
func (c *CasbinRuleRepo) RemovePolicies(ctx context.Context, rules [][]string) (bool, error) {
	_, err := c.enforcer.RemovePolicies(rules)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdatePolicies 更新权限
func (c *CasbinRuleRepo) UpdatePolicies(ctx context.Context, oldRules [][]string, newRules [][]string) (bool, error) {
	_, err := c.enforcer.UpdatePolicies(oldRules, newRules)
	if err != nil {
		return false, err
	}
	return true, nil
}
