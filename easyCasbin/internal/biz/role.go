package biz

import (
	"context"
	v1 "easyCasbin/api/role/v1"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;size:64;uniqueIndex:domain_role"`
	Domain   string `json:"domain" gorm:"size:127;uniqueIndex:domain_role"`
	IsDelete int32  `json:"is_delete" gorm:"size:127;uniqueIndex:domain_role"`
}

func (Role) TableName() string {
	return "role"
}

type DomainRoleRepo interface {
	// AddDomainRole 添加域角色
	AddDomainRole(ctx context.Context, domain, role string) (*Role, error)
	// UpdateDomainRole 更新域角色
	UpdateDomainRole(ctx context.Context, domain, oldRole, newRole string) (*Role, error)
	// DeleteDomainRole 删除域角色
	DeleteDomainRole(ctx context.Context, domain, role string) (bool, error)
	// GetDomainRoles 获取指定域下所有角色
	GetDomainRoles(ctx context.Context, domain string) ([]Role, error)
	// CheckDomainRole 校验域角色是否存在
	CheckDomainRole(ctx context.Context, domain, role string) (bool, error)

	// GetAllDomains 获取所有域
	GetAllDomains(ctx context.Context) ([]string, error)
	// CheckDomains 校验是否存在相关域信息
	CheckDomains(ctx context.Context, domains []string) (bool, error)
	// DeleteDomain 删除指定域的所有相关信息
	DeleteDomain(ctx context.Context, domain string) (bool, error)
}

// DomainRoleUseCase 域角色用例
type DomainRoleUseCase struct {
	repo DomainRoleRepo
	log  *log.Helper
}

func NewDomainRepoUseCase(repo DomainRoleRepo, logger log.Logger) *DomainRoleUseCase {
	return &DomainRoleUseCase{repo: repo, log: log.NewHelper(logger)}
}

// AddDomainRole 添加域角色，如果域角色已经存在直接返回
func (uc *DomainRoleUseCase) AddDomainRole(ctx context.Context, domain, role string) (*DomainRoleResponse, error) {

	exist, _ := uc.repo.CheckDomainRole(ctx, domain, role)
	if exist {
		return nil, v1.ErrorDomainRoleExist("domain role already exist!")
	}

	domainRole, err := uc.repo.AddDomainRole(ctx, domain, role)
	if err != nil {
		return nil, err
	}
	return &DomainRoleResponse{Id: domainRole.ID, Name: domainRole.Name,
		Domain: domainRole.Domain, CreateTime: domainRole.CreatedAt}, nil
}

// UpdateDomainRoleInfo UpdateRoleInfo 更新域角色，同一个域下角色名不能重复
func (uc *DomainRoleUseCase) UpdateDomainRoleInfo(ctx context.Context, domain, oldRoleName, newRoleName string) (*DomainRoleResponse, error) {
	oldRole, _ := uc.repo.CheckDomainRole(ctx, domain, oldRoleName)
	if !oldRole {
		return nil, v1.ErrorDomainRoleNotFound("domain role not found!")
	}
	newRole, _ := uc.repo.CheckDomainRole(ctx, domain, newRoleName)
	if newRole {
		return nil, v1.ErrorDomainRoleExist("domain role name already exist!")
	}
	domainRole, err := uc.repo.UpdateDomainRole(ctx, domain, oldRoleName, newRoleName)
	if err != nil {
		return nil, err
	}
	return &DomainRoleResponse{Id: domainRole.ID, Name: domainRole.Name,
		Domain: domainRole.Domain, UpdateTime: domainRole.UpdatedAt}, nil
}

// DeleteDomainRole 删除域角色，同时需要删除域角色对应的所有权限信息
func (uc *DomainRoleUseCase) DeleteDomainRole(ctx context.Context, domain, role string) (bool, error) {
	return uc.repo.DeleteDomainRole(ctx, domain, role)
}
