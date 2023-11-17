package data

import (
	"context"
	v1 "easyCasbin/api/role/v1"
	"easyCasbin/internal/biz"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
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
		return nil, v1.ErrorDomainRoleExist("domain role exist")
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

// GetDomainRoles 获取指定域下角色列表
func (r *roleRepo) GetDomainRoles(ctx context.Context, domain, roleName string, offset, limit int) ([]*biz.Role, int32, error) {
	var roles []*biz.Role

	query := r.data.db.Table("role").Where("domain=?", domain)
	if roleName != "" {
		query.Where("name LIKE ?", "%"+roleName+"%")
	}

	total := int64(0)
	query.Count(&total)
	if offset >= 0 && limit > 0 {
		query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&roles).Error; err != nil {
		return nil, 0, err
	}
	return roles, int32(total), nil
}

// CheckDomainRole 校验域角色是否存在
func (r *roleRepo) CheckDomainRole(ctx context.Context, domain, role string) (bool, error) {
	result := r.data.db.Where(&biz.Role{Domain: domain, Name: role}).First(&biz.Role{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

// GetAllDomains 获取所有域
func (r *roleRepo) GetAllDomains() ([]biz.Role, error) {
	var roleList []biz.Role
	if err := r.data.db.Distinct("id", "domain", "name").Where("domain=name").Find(&roleList).Error; err != nil {
		return nil, err
	}
	return roleList, nil
}

// GetDomainRoleId 获取指定域角色的ID
func (r *roleRepo) GetDomainRoleId(ctx context.Context, domain, role string) (int, error) {
	var domainRole biz.Role
	result := r.data.db.Where(&biz.Role{Domain: domain, Name: role}).First(&domainRole)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return -1, nil
	}
	return int(domainRole.ID), nil
}

// GetRoleIdNameMap 获取所有角色 id 与角色名的映射关系
func (r *roleRepo) GetRoleIdNameMap() (map[string]string, error) {
	var domainRole []biz.Role
	maps := make(map[string]string)
	result := r.data.db.Find(&domainRole)
	if err := result.Error; err != nil {
		return nil, err
	}
	for _, role := range domainRole {
		maps["role:"+strconv.Itoa(int(role.ID))] = role.Name
	}
	return maps, nil
}

// GetSubsInDomainRole 查询域角色下包含的鉴权主体：角色或者用户
func (r *roleRepo) GetSubsInDomainRole(ctx context.Context, domain, role string) (*[]biz.CasbinRule, error) {
	var rules []biz.CasbinRule

	// 查询出对应域本身的角色 id：域名和角色名相等
	domainID, err := r.GetDomainRoleId(ctx, domain, domain)
	// 查询出对应域角色的 id
	roleId, err := r.GetDomainRoleId(ctx, domain, role)

	result := r.data.db.Where(&biz.CasbinRule{V2: "role:" + strconv.Itoa(domainID), V1: "role:" + strconv.Itoa(roleId)}).Find(&rules)
	if result.Error != nil {
		return nil, err
	}

	return &rules, nil
}

func (r *roleRepo) CheckDomains(ctx context.Context, domains []string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) DeleteDomain(ctx context.Context, domain string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
