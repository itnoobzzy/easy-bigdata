package biz

import (
	casbin_rule_v1 "easyCasbin/api/casbin_rule/v1"
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
			p.Domain = "default"
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
