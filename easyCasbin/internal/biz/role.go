package biz

import (
	"context"
	v1 "easyCasbin/api/role/v1"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strings"
)

/*
Role 业务逻辑：
 1. 用户组管理和域管理混合
 2. name==domain时，该条记录是域，反之为组
 3. _* 是所有域，_all是所有默认组，_root是系统管理员组
 4. (_*,_*) 标示所有域，（_all,_*）中的_* 无特殊含义，占位标示
*/
type Role struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;size:64;uniqueIndex:domain_role"`
	Domain   string `json:"domain" gorm:"size:127;uniqueIndex:domain_role"`
	IsDelete int32  `json:"is_delete" gorm:"size:127;uniqueIndex:domain_role"`
}

func (Role) TableName() string {
	return "role"
}

//go:generate mockgen -destination=role_mock.go -package=biz . DomainRoleRepo
type DomainRoleRepo interface {
	// AddDomainRole 添加域角色
	AddDomainRole(ctx context.Context, domain, role string) (*Role, error)
	// UpdateDomainRole 更新域角色
	UpdateDomainRole(ctx context.Context, domain, oldRole, newRole string) (*Role, error)
	// DeleteDomainRole 删除域角色
	DeleteDomainRole(ctx context.Context, domain, role string) (bool, error)
	// GetDomainRoles 获取指定域下角色列表
	GetDomainRoles(ctx context.Context, domain, roleName string, offset, limit int) ([]*Role, int32, error)
	// CheckDomainRole 校验域角色是否存在
	CheckDomainRole(ctx context.Context, domain, role string) (bool, error)
	// GetDomainRoleId 获取指定域角色的ID
	GetDomainRoleId(ctx context.Context, domain, role string) (int, error)
	// GetRoleIdNameMap 获取所有角色 id 与角色名的映射关系
	GetRoleIdNameMap() (maps map[string]string, err error)
	// GetAllDomains 获取所有域
	GetAllDomains() ([]Role, error)
	// CheckDomains 校验是否存在相关域信息
	CheckDomains(ctx context.Context, domains []string) (bool, error)
	// DeleteDomain 删除指定域的所有相关信息
	DeleteDomain(ctx context.Context, domain string) (bool, error)
	// GetSubsInDomainRole 查询域角色下包含的鉴权主体：角色或者用户
	GetSubsInDomainRole(ctx context.Context, domain, role string) (*[]CasbinRule, error)
}

// DomainRoleUseCase 域角色用例
type DomainRoleUseCase struct {
	repo DomainRoleRepo
	log  *log.Helper
}

func NewDomainRoleUseCase(repo DomainRoleRepo, logger log.Logger) *DomainRoleUseCase {
	return &DomainRoleUseCase{repo: repo, log: log.NewHelper(logger)}
}

// GetAllDomains 获取所有域
func (uc *DomainRoleUseCase) GetAllDomains() ([]*DomainRoleResponse, error) {
	domains, err := uc.repo.GetAllDomains()
	if err != nil {
		return nil, err
	}
	var domainRoles []*DomainRoleResponse
	for _, role := range domains {
		domainRoles = append(domainRoles, &DomainRoleResponse{
			Id:     int32(role.ID),
			Name:   role.Name,
			Domain: role.Domain,
		})
	}
	return domainRoles, nil
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
	return &DomainRoleResponse{Id: int32(domainRole.ID), Name: domainRole.Name,
		Domain: domainRole.Domain, CreateTime: int32(domainRole.CreatedAt.Unix())}, nil
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
	return &DomainRoleResponse{Id: int32(domainRole.ID), Name: domainRole.Name,
		Domain: domainRole.Domain, UpdateTime: int32(domainRole.UpdatedAt.Unix())}, nil
}

// DeleteDomainRole 删除域角色，同时需要删除域角色对应的所有权限信息
func (uc *DomainRoleUseCase) DeleteDomainRole(ctx context.Context, domain, role string) (bool, error) {
	return uc.repo.DeleteDomainRole(ctx, domain, role)
}

// GetDomainRoles 查询域下所有角色
func (uc *DomainRoleUseCase) GetDomainRoles(ctx context.Context, params *GetDomainRolesParams) ([]*DomainRoleResponse, int32, error) {
	roles, total, err := uc.repo.GetDomainRoles(ctx, params.domain, params.roleName, params.offset, params.limit)
	if err != nil {
		return nil, 0, err
	}
	var domainRoles []*DomainRoleResponse
	for _, role := range roles {
		domainRoles = append(domainRoles, &DomainRoleResponse{
			Id:     int32(role.ID),
			Name:   role.Name,
			Domain: role.Domain,
		})
	}
	return domainRoles, total, nil
}

// GetSubsInDomainRole 查询域角色下所有鉴权主体，包括用户和角色
func (uc *DomainRoleUseCase) GetSubsInDomainRole(ctx context.Context, domain, role string) ([]map[string]string, error) {
	rules, err := uc.repo.GetSubsInDomainRole(ctx, domain, role)
	if err != nil {
		return nil, err
	}

	nameMap, _ := uc.repo.GetRoleIdNameMap()

	var subs []map[string]string
	for _, r := range *rules {
		id := r.V0
		// sub 鉴权主体可以是角色或者用户, eg: role:1 或 user:zzy
		// subType 鉴权主体类型 subName 鉴权主体名称
		subType := strings.Split(id, ":")[0]
		subName := strings.Split(id, ":")[1]
		// 如果鉴权主体类型是角色的话，需要查询出对应的角色名
		if subType == "role" {
			roleName, ok := nameMap[subName]
			if ok {
				subName = roleName
			}
		}
		sub := map[string]string{}
		sub["id"] = id
		sub["name"] = subName
		subs = append(subs, sub)
	}

	return subs, nil

}
