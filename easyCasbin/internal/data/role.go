package data

import (
	"context"
	"errors"

	v1 "easyCasbin/api/role/v1"
	"easyCasbin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.DomainRoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// AddDomainRole 添加域角色，如果存在数据库不会新增，幂等性
func (r *roleRepo) AddDomainRole(ctx context.Context, domain, role string) (*biz.Role, error) {
	var domainRole biz.Role
	domainRole.Name = role
	domainRole.Domain = domain
	domainRole.IsDelete = 0
	if err := r.data.db.Create(&domainRole).Error; err != nil {
		return nil, err
	}
	return &domainRole, nil
}

// UpdateDomainRole 更新域角色，如果域角色不存在报错
func (r *roleRepo) UpdateDomainRole(ctx context.Context, domain, oldRole, newRole string) (*biz.Role, error) {
	var domainRole biz.Role
	result := r.data.db.Where(&biz.Role{Domain: domain, Name: oldRole}).First(&domainRole)
	domainRole.Domain = domain
	domainRole.Name = newRole
	if err := r.data.db.Save(&domainRole).Error; err != nil {
		return nil, v1.ErrorInternalErr("update domain role err: %v", result.Error)
	}
	return &domainRole, nil
}

// DeleteDomainRole 删除域角色
func (r *roleRepo) DeleteDomainRole(ctx context.Context, domain, role string) (bool, error) {
	var domainRole biz.Role
	result := r.data.db.Where(&domainRole).First(&domainRole)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, v1.ErrorDomainRoleNotFound("domain role not found!")
	}
	domainRole.IsDelete = 1
	if err := r.data.db.Save(&domainRole).Delete(&domainRole).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *roleRepo) GetDomainRoles(ctx context.Context, domain string) ([]biz.Role, error) {
	//TODO implement me
	panic("implement me")
}

// CheckDomainRole 校验域角色是否存在
func (r *roleRepo) CheckDomainRole(ctx context.Context, domain, role string) (bool, error) {
	result := r.data.db.Where(&biz.Role{Domain: domain, Name: role}).First(&biz.Role{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

func (r *roleRepo) GetAllDomains(ctx context.Context) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) CheckDomains(ctx context.Context, domains []string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) DeleteDomain(ctx context.Context, domain string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
