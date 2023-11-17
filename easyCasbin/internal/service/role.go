package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	v1 "easyCasbin/api/role/v1"
	"easyCasbin/internal/biz"
)

type DomainRoleService struct {
	v1.UnimplementedDomainRoleServer

	uc  *biz.DomainRoleUseCase
	log *log.Helper

	// 依赖casbin_rule服务对权限信息的操作
	cuc *biz.CasbinRuleUseCase
}

func NewDomainRoleService(uc *biz.DomainRoleUseCase, logger log.Logger, cuc *biz.CasbinRuleUseCase) *DomainRoleService {
	return &DomainRoleService{uc: uc, log: log.NewHelper(logger), cuc: cuc}
}

// GetAllDomains 获取所有域
func (s *DomainRoleService) GetAllDomains(context.Context, *emptypb.Empty) (*v1.GetAllDomainsRpl, error) {
	var resData []*v1.GetAllDomainsRpl_Data
	domains, err := s.uc.GetAllDomains()
	if err != nil {
		return nil, err
	}
	for _, d := range domains {
		resData = append(resData, &v1.GetAllDomainsRpl_Data{
			Id:     d.Id,
			Name:   d.Name,
			Domain: d.Domain,
		})
	}
	return &v1.GetAllDomainsRpl{
		Status:  0,
		Message: "ok!",
		Data:    resData,
	}, nil
}

// AddDomainRole 添加域角色
func (s *DomainRoleService) AddDomainRole(ctx context.Context, req *v1.AddDomainRoleReq) (*v1.AddDomainRoleRpl, error) {
	role, err := s.uc.AddDomainRole(ctx, req.DomainName, req.RoleName)
	if err != nil {
		return nil, err
	}
	return &v1.AddDomainRoleRpl{
		Code:    0,
		Message: "add domain role success!",
		Data: &v1.AddDomainRoleRpl_Data{
			Id:        role.Id,
			Domain:    role.Domain,
			Name:      role.Name,
			CreatTime: role.CreateTime,
		},
	}, nil
}

// UpdateRoleInfo 更新域角色信息
func (s *DomainRoleService) UpdateRoleInfo(ctx context.Context, req *v1.UpdateDomainRoleReq) (*v1.UpdateDomainRoleRpl, error) {
	role, err := s.uc.UpdateDomainRoleInfo(ctx, req.DomainName, req.RoleName, req.NewRoleName)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDomainRoleRpl{
		Code:    0,
		Message: "update domain role success!",
		Data: &v1.UpdateDomainRoleRpl_Data{
			Id:         role.Id,
			Domain:     role.Domain,
			Name:       role.Name,
			UpdateTime: role.UpdateTime,
		},
	}, nil
}

// DeleteRole 删除域角色，删除角色的同时将删除对应的权限规则
func (s *DomainRoleService) DeleteRole(ctx context.Context, req *v1.DeleteDomainRoleReq) (*v1.DeleteDomainRoleRpl, error) {
	ok, err := s.uc.DeleteDomainRole(ctx, req.DomainName, req.RoleName)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &v1.DeleteDomainRoleRpl{
			Code:    1,
			Message: "delete domain role failed!",
			Data:    &v1.DeleteDomainRoleRpl_Data{DeleteTime: int32(time.Now().Unix())},
		}, nil
	}
	return &v1.DeleteDomainRoleRpl{
		Code:    0,
		Message: "delete domain role success!",
		Data:    &v1.DeleteDomainRoleRpl_Data{DeleteTime: int32(time.Now().Unix())},
	}, nil
}

// GetDomainRoles 获取指定域下所有角色列表
func (s *DomainRoleService) GetDomainRoles(ctx context.Context, req *v1.GetDomainRolesReq) (*v1.GetDomainRolesRpl, error) {
	params := biz.NewGetDomainRolesParams(req)
	roles, total, err := s.uc.GetDomainRoles(ctx, params)
	if err != nil {
		return nil, err
	}

	var resRoles []*v1.GetDomainRolesRpl_Role
	for _, role := range roles {
		resRoles = append(resRoles, &v1.GetDomainRolesRpl_Role{
			Id:     role.Id,
			Name:   role.Name,
			Domain: role.Domain,
		})
	}

	return &v1.GetDomainRolesRpl{
		Status:  0,
		Message: "ok!",
		Data: &v1.GetDomainRolesRpl_Data{
			Roles: resRoles,
			Total: total,
		},
	}, nil
}

// GetSubsInDomainRole 查询域角色下所有鉴权主体，用户或者角色
func (s *DomainRoleService) GetSubsInDomainRole(ctx context.Context, req *v1.GetSubsInDomainRoleReq) (*v1.GetSubsInDomainRoleRpl, error) {
	subs, err := s.uc.GetSubsInDomainRole(ctx, req.DomainName, req.RoleName)
	if err != nil {
		return nil, err
	}
	var resData []*v1.GetSubsInDomainRoleRpl_Data
	for _, s := range subs {
		resData = append(resData, &v1.GetSubsInDomainRoleRpl_Data{
			Id:   s["id"],
			Name: s["name"],
		})
	}
	return &v1.GetSubsInDomainRoleRpl{
		Code:    0,
		Message: "ok!",
		Data:    resData,
	}, nil
}

// AddRoleForSubInDomain 为用户添加域角色或者为角色继承另一个角色权限
func (s *DomainRoleService) AddRoleForSubInDomain(ctx context.Context, req *v1.AddRoleForSubInDomainReq) (*v1.AddRoleForSubInDomainRpl, error) {
	//TODO implement me
	panic("implement me")
}

func (s *DomainRoleService) DeleteRoleForSubInDomain(ctx context.Context, req *v1.DeleteRoleForSubInDomainReq) (*v1.DeleteRoleForSubInDomainRpl, error) {
	//TODO implement me
	panic("implement me")
}
