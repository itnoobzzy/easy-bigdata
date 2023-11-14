package biz

import (
	"context"
	"easyCasbin/api/casbin_rule/v1"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type CasbinRule struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `json:"ptype" gorm:"size:100"`
	V0    string `json:"v0" gorm:"size:100"`
	V1    string `json:"v1" gorm:"size:100"`
	V2    string `json:"v2" gorm:"size:100"`
	V3    string `json:"v3" gorm:"size:100"`
	V4    string `json:"v4" gorm:"size:100"`
	V5    string `json:"v5" gorm:"size:100"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}

type CasbinRuleRepo interface {
	// GetAllSubjects 获取所有鉴权主体
	GetAllSubjects(ctx context.Context) ([]string, error)

	// GetRolesForUserInDomain 查询用户在域上的角色
	GetRolesForUserInDomain(ctx context.Context, username, domain string) (roles []string, err error)

	// GetImplicitPermissionsForUser 查询鉴权主体（用户或角色）在对应域上的权限
	GetImplicitPermissionsForUser(ctx context.Context, username, domain string) (permissions [][]string, err error)

	// GetImplicitUsersForRole 查询域角色下的所有用户
	GetImplicitUsersForRole(ctx context.Context, domain, role string) (users []string, err error)

	// AddRoleForUserInDomain 为用户添加域角色或者为角色继承另一个角色权限
	AddRoleForUserInDomain(ctx context.Context, user, domain, role string) (ok bool, err error)

	// DeleteRoleForUserInDomain 移除指定域角色下的用户或者取消指定域角色的继承
	DeleteRoleForUserInDomain(ctx context.Context, user, domain, role string) (ok bool, err error)

	// DeleteDomain 删除域上的所有规则
	DeleteDomain(ctx context.Context, domain string) (bool, error)

	// BatchEnforce 批量校验权限
	BatchEnforce(ctx context.Context, rules [][]interface{}) ([]bool, error)

	// AddPolicies 添加权限
	AddPolicies(ctx context.Context, rules [][]string) (bool, error)

	// RemovePolicies 删除权限
	RemovePolicies(ctx context.Context, rules [][]string) (bool, error)

	// UpdatePolicies 更新权限
	UpdatePolicies(ctx context.Context, oldRules [][]string, newRules [][]string) (bool, error)
}

type CasbinRuleUseCase struct {
	repo CasbinRuleRepo
	log  *log.Helper
}

func NewCasbinRuleUseCase(repo CasbinRuleRepo, logger log.Logger) *CasbinRuleUseCase {
	return &CasbinRuleUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// GetAllSubjects 获取所有鉴权主体
func (uc *CasbinRuleUseCase) GetAllSubjects(ctx context.Context) ([]string, error) {
	return uc.repo.GetAllSubjects(ctx)
}

// DeleteDomainRules 删除域规则
func (uc *CasbinRuleUseCase) DeleteDomainRules(ctx context.Context, domain string) (bool, error) {
	domain = "domain:" + domain
	return uc.repo.DeleteDomain(ctx, domain)
}

// BatchEnforce 批量校验规则
func (uc *CasbinRuleUseCase) BatchEnforce(ctx context.Context, rules [][]interface{}) ([]bool, error) {
	return uc.repo.BatchEnforce(ctx, rules)
}

// CheckDomains 校验规则中的域是否存在
// 不存在返回错误，正常返回规则列表
func (uc *CasbinRuleUseCase) CheckDomains(ctx context.Context, domainRoleUseCase *DomainRoleUseCase,
	policies []*PolicyParams) (rules [][]string, err error) {
	allDomains, _ := domainRoleUseCase.repo.GetAllDomains(ctx)
	sort.Strings(allDomains)
	for _, v := range policies {
		rule := make([]string, 0)
		rule = append(rule, v.name, v.domain, v.resource, v.action, v.eft)
		rules = append(rules, rule)
		domain := strings.Split(v.domain, ":")[1]
		idx := sort.SearchStrings(allDomains, domain)
		if in := idx < len(allDomains) && allDomains[idx] == domain; !in {
			return rules, v1.ErrorDomainNotFound("domain: %s not found", domain)
		}
	}
	return rules, nil
}

// AddPermissionsForSubInDomain 为鉴权主体批量添加权限
func (uc *CasbinRuleUseCase) AddPermissionsForSubInDomain(ctx context.Context, rules [][]string) (bool, error) {
	// 先删除这些权限，否则如果添加的权限中有任何已经存在的，将不会执行添加动作
	_, err := uc.repo.RemovePolicies(ctx, rules)
	_, err = uc.repo.AddPolicies(ctx, rules)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetPermissions 获取指定域上鉴权主体的所属权限，包括继承的权限
func (uc *CasbinRuleUseCase) GetPermissions(ctx context.Context, domain, sub string) (permissions []PermissionsResponse, err error) {
	ps, err := uc.repo.GetImplicitPermissionsForUser(ctx, sub, domain)
	if err != nil {
		return nil, err
	}

	newSet := mapset.NewSet[PermissionsResponse]()
	for _, p := range ps {
		newSet.Add(PermissionsResponse{
			Resource: p[2],
			Action:   p[3],
		})
	}
	all := newSet.ToSlice()
	for _, v := range all {
		permissions = append(permissions, PermissionsResponse{
			Resource: v.Resource,
			Action:   v.Action,
		})
	}
	return permissions, nil
}

// UpdatePermissions 批量更新权限
func (uc *CasbinRuleUseCase) UpdatePermissions(ctx context.Context, oldRules, newRules [][]string) (bool, error) {
	return uc.repo.UpdatePolicies(ctx, oldRules, newRules)
}

// DeletePermissions 批量删除权限
func (uc *CasbinRuleUseCase) DeletePermissions(ctx context.Context, rules [][]string) (bool, error) {
	return uc.repo.RemovePolicies(ctx, rules)
}
