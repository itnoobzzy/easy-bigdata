package service

import (
	"context"
	v1 "easyCasbin/api/casbin_rule/v1"
	"easyCasbin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CasbinRuleService struct {
	v1.UnimplementedCasbinRuleServer

	uc  *biz.CasbinRuleUseCase
	log *log.Helper

	// 依赖domain_role服务对域角色信息的操作
	duc *biz.DomainRoleUseCase
}

func NewCasbinRuleService(uc *biz.CasbinRuleUseCase, logger log.Logger, duc *biz.DomainRoleUseCase) *CasbinRuleService {
	return &CasbinRuleService{uc: uc, log: log.NewHelper(logger), duc: duc}
}

// GetAllSubjects 获取所有鉴权主体
func (s *CasbinRuleService) GetAllSubjects(ctx context.Context, e *emptypb.Empty) (*v1.GetAllSubjectsRpl, error) {

	var resData []*v1.GetAllSubjectsRpl_Data

	subs, err := s.uc.GetAllSubjects(ctx)
	if err != nil {
		return nil, err
	}

	for _, sub := range subs {
		resData = append(resData, &v1.GetAllSubjectsRpl_Data{
			Sub: sub,
		})
	}

	return &v1.GetAllSubjectsRpl{
		Code:    0,
		Message: "ok",
		Data:    resData,
	}, nil
}

// DeleteDomain 删除对应域下所有规则
func (s *CasbinRuleService) DeleteDomain(ctx context.Context, req *v1.DeleteDomainReq) (*v1.DeleteDomainRpl, error) {
	ok, err := s.uc.DeleteDomainRules(ctx, req.Domain)
	if err != nil || !ok {
		return nil, err
	}
	return &v1.DeleteDomainRpl{
		Code:    0,
		Message: "ok",
	}, nil
}

// BatchEnforce 批量校验所有规则
func (s *CasbinRuleService) BatchEnforce(ctx context.Context, req *v1.BatchEnforceReq) (*v1.BatchEnforceRpl, error) {
	params, _ := biz.NewBatchEnforceParams(req)
	results, err := s.uc.BatchEnforce(ctx, params.Rules)
	if err != nil {
		return nil, err
	}
	return &v1.BatchEnforceRpl{
		Code:    0,
		Message: "ok",
		Data:    results,
	}, nil
}

// AddPermissions 为鉴权主体批量添加权限
func (s *CasbinRuleService) AddPermissions(ctx context.Context, req *v1.AddPermissionsReq) (*v1.AddPermissionsRpl, error) {

	params, _ := biz.NewAddPermissionsParams(req)

	rules, err := s.uc.CheckDomains(ctx, s.duc, params.Policies)
	if err != nil {
		return nil, err
	}
	ok, err := s.uc.AddPermissionsForSubInDomain(ctx, rules)
	if err != nil || !ok {
		return nil, err
	}
	return &v1.AddPermissionsRpl{
		Code:    0,
		Message: "ok",
	}, nil
}

// GetPermissions 获取鉴权主体所有权限
func (s *CasbinRuleService) GetPermissions(ctx context.Context, req *v1.GetPermissionsReq) (*v1.GetPermissionsRpl, error) {
	ps, err := s.uc.GetPermissions(ctx, "domain:"+req.Domain, req.Sub)
	if err != nil {
		return nil, err
	}
	var resData []*v1.GetPermissionsRpl_Data
	for _, p := range ps {
		resData = append(resData, &v1.GetPermissionsRpl_Data{
			Resource: p.Resource,
			Action:   p.Action,
		})
	}

	return &v1.GetPermissionsRpl{
		Code:    0,
		Message: "ok",
		Data:    resData,
	}, nil
}

// UpdatePermissions 为鉴权主体批量更新权限
func (s *CasbinRuleService) UpdatePermissions(ctx context.Context, req *v1.UpdatePermissionsReq) (*v1.UpdatePermissionsRpl, error) {
	params, err := biz.NewUpdatePermissionsParams(req)
	if err != nil {
		return nil, err
	}

	oldRules, err := s.uc.CheckDomains(ctx, s.duc, params.OldPolicies)
	newRules, err := s.uc.CheckDomains(ctx, s.duc, params.NewPolicies)

	ok, err := s.uc.UpdatePermissions(ctx, oldRules, newRules)
	if !ok || err != nil {
		return nil, err
	}
	return &v1.UpdatePermissionsRpl{
		Code:    0,
		Message: "ok",
	}, nil

}

// DeletePermissions 为鉴权主体批量删除权限
func (s *CasbinRuleService) DeletePermissions(ctx context.Context, req *v1.DeletePermissionsReq) (*v1.DeletePermissionsRpl, error) {
	params, _ := biz.NewDeletePermissionsParams(req)

	rules, err := s.uc.CheckDomains(ctx, s.duc, params.Policies)
	if err != nil {
		return nil, err
	}
	ok, err := s.uc.DeletePermissions(ctx, rules)
	if !ok || err != nil {
		return nil, err
	}
	return &v1.DeletePermissionsRpl{
		Code:    0,
		Message: "ok",
	}, nil
}
