package service

import (
	"context"
	v1 "easyCasbin/api/role/v1"
	"easyCasbin/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type DomainRoleService struct {
	v1.UnimplementedDomainRoleServer

	uc  *biz.DomainRoleUseCase
	log *log.Helper
}

func NewDomainRoleService(uc *biz.DomainRoleUseCase, logger log.Logger) *DomainRoleService {
	return &DomainRoleService{uc: uc, log: log.NewHelper(logger)}
}

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

func (s *DomainRoleService) DeleteRole(ctx context.Context, req *v1.DeleteDomainRoleReq) (*v1.DeleteDomainRoleRpl, error) {
	ok, err := s.uc.DeleteDomainRole(ctx, req.DomainName, req.RoleName)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &v1.DeleteDomainRoleRpl{
			Code:    1,
			Message: "delete domain role failed!",
			Data:    &v1.DeleteDomainRoleRpl_Data{DeleteTime: time.Now().Unix()},
		}, nil
	}
	return &v1.DeleteDomainRoleRpl{
		Code:    0,
		Message: "delete domain role success!",
		Data:    &v1.DeleteDomainRoleRpl_Data{DeleteTime: time.Now().Unix()},
	}, nil
}

func (s *DomainRoleService) GetDomainRoles(ctx context.Context, req *v1.GetAllRolesReq) (*v1.GetAllRolesRpl, error) {
	var resData []*v1.GetAllRolesRpl_Data
	roles, err := s.uc.GetDomainRoles(ctx, req.DomainName)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		resData = append(resData, &v1.GetAllRolesRpl_Data{
			Id:         role.Id,
			CreateTime: role.CreateTime,
			UpdateTime: role.UpdateTime,
			Name:       role.Name,
			Domain:     role.Domain,
		})
	}
	return &v1.GetAllRolesRpl{
		Code:    0,
		Message: "ok!",
		Data:    resData,
	}, nil
}

func (s *DomainRoleService) GetDomainSubsForRole(ctx context.Context, req *v1.GetDomainSubsForRoleReq) (*v1.GetDomainSubsForRoleRpl, error) {
	//TODO implement me
	panic("implement me")
}

func (s *DomainRoleService) AddRoleForSubInDomain(ctx context.Context, req *v1.AddRoleForSubInDomainReq) (*v1.AddRoleForSubInDomainRpl, error) {
	//TODO implement me
	panic("implement me")
}

func (s *DomainRoleService) DeleteRoleForSubInDomain(ctx context.Context, req *v1.DeleteRoleForSubInDomainReq) (*v1.DeleteRoleForSubInDomainRpl, error) {
	//TODO implement me
	panic("implement me")
}
