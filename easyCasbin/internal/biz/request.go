package biz

import (
	casbin_rule_v1 "easyCasbin/api/casbin_rule/v1"
	rv1 "easyCasbin/api/role/v1"
	v1 "easyCasbin/api/user/v1"
)

// LoginRequest 登录参数
type LoginRequest struct {
	username string
	password string
}

func NewLoginRequest(username, password string) (*LoginRequest, error) {
	if username == "" {
		return nil, v1.ErrorPasswordErr("username or password error!")
	}
	if password == "" {
		return nil, v1.ErrorPasswordErr("username or password error!")
	}
	return &LoginRequest{
		username: username,
		password: password,
	}, nil
}

// PolicyParams 单个权限操作所需参数
type PolicyParams struct {
	name     string
	domain   string
	resource string
	action   string
	eft      string
}

// PoliciesParams 批量操作权限所需参数
type PoliciesParams struct {
	Policies []*PolicyParams
}

// BatchEnforceParams 批量校验权限所需参数
type BatchEnforceParams struct {
	Rules [][]interface{}
}

// UpdatePermissionsParams 批量修改权限所需参数
type UpdatePermissionsParams struct {
	OldPolicies []*PolicyParams
	NewPolicies []*PolicyParams
}

// GetDomainAuthParams 查询域下权限规则列表所需参数
type GetDomainAuthParams struct {
	domain string
	search string
	limit  int
	offset int
}

// GetDomainRolesParams 查询域下角色列表
type GetDomainRolesParams struct {
	domain   string
	roleName string
	limit    int
	offset   int
}

func NewGetDomainRolesParams(req *rv1.GetDomainRolesReq) *GetDomainRolesParams {
	return &GetDomainRolesParams{
		domain:   req.DomainName,
		roleName: req.RoleName,
		limit:    int(req.PageSize),
		offset:   int((req.PageNum - 1) * req.PageSize),
	}
}

func NewGetDomainAuthParams(req *casbin_rule_v1.GetDomainAuthReq) (*GetDomainAuthParams, error) {
	return &GetDomainAuthParams{
		domain: req.Domain,
		search: req.Search,
		limit:  int(req.PageSize),
		offset: int((req.PageNum - 1) * req.PageSize),
	}, nil
}

func NewBatchEnforceParams(policies *casbin_rule_v1.BatchEnforceReq) (*BatchEnforceParams, error) {
	var rules [][]interface{}
	for _, p := range policies.Policies {
		rule := make([]interface{}, 0)
		rule = append(rule, p.Name, p.Domain, p.Resource, p.Action)
		rules = append(rules, rule)
	}
	return &BatchEnforceParams{Rules: rules}, nil
}

func NewAddPermissionsParams(policies *casbin_rule_v1.AddPermissionsReq) (*PoliciesParams, error) {
	var ps []*PolicyParams
	for _, p := range policies.Policies {

		// 默认添加至default 域，并且添加的权限是生效的
		if p.Domain == "" {
			p.Domain = "domain:default"
		}
		if p.Eft == "" {
			p.Eft = "allow"
		}

		ps = append(ps, &PolicyParams{
			name:     p.Name,
			domain:   "domain:" + p.Domain,
			resource: p.Resource,
			action:   p.Action,
			eft:      p.Eft,
		})
	}
	return &PoliciesParams{Policies: ps}, nil
}

func NewUpdatePermissionsParams(policies *casbin_rule_v1.UpdatePermissionsReq) (*UpdatePermissionsParams, error) {
	var oldPs []*PolicyParams
	var newPs []*PolicyParams
	for _, p := range policies.OldPolicies {

		// 默认添加至default 域，并且添加的权限是生效的
		if p.Domain == "" {
			p.Domain = "domain:default"
		}
		if p.Eft == "" {
			p.Eft = "allow"
		}

		oldPs = append(oldPs, &PolicyParams{
			name:     p.Name,
			domain:   "domain:" + p.Domain,
			resource: p.Resource,
			action:   p.Action,
			eft:      p.Eft,
		})
	}
	for _, p := range policies.NewPolicies {

		// 默认添加至default 域，并且添加的权限是生效的
		if p.Domain == "" {
			p.Domain = "domain:default"
		}
		if p.Eft == "" {
			p.Eft = "allow"
		}

		newPs = append(newPs, &PolicyParams{
			name:     p.Name,
			domain:   "domain:" + p.Domain,
			resource: p.Resource,
			action:   p.Action,
			eft:      p.Eft,
		})
	}

	if len(oldPs) != len(newPs) {
		return nil, casbin_rule_v1.ErrorInvalidArgs("The number of old rules must match the number of new rules")
	}

	return &UpdatePermissionsParams{OldPolicies: oldPs, NewPolicies: newPs}, nil
}

func NewDeletePermissionsParams(policies *casbin_rule_v1.DeletePermissionsReq) (*PoliciesParams, error) {
	var ps []*PolicyParams
	for _, p := range policies.Policies {

		// 默认添加至default 域，并且添加的权限是生效的
		if p.Domain == "" {
			p.Domain = "domain:default"
		}
		if p.Eft == "" {
			p.Eft = "allow"
		}

		ps = append(ps, &PolicyParams{
			name:     p.Name,
			domain:   "domain:" + p.Domain,
			resource: p.Resource,
			action:   p.Action,
			eft:      p.Eft,
		})
	}
	return &PoliciesParams{Policies: ps}, nil
}
