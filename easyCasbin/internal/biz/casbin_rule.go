package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	GetAllSubjects(ctx context.Context) interface{}

	// GetRolesForUserInDomain 查询用户在域上的角色
	GetRolesForUserInDomain(ctx context.Context, username, domain string) (roles []string, err error)
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
