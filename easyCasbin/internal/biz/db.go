package biz

import (
	"github.com/go-kratos/kratos/v2/log"

	db "easyCasbin/api/db"
)

type DbIniterRepo interface {
	InitUserModel(*User) (rsp *db.InitRpl, err error)
	InitDomainRoleModel(*Role) (rsp *db.InitRpl, err error)
}

type DbIniterUsecase struct {
	repo DbIniterRepo
	log  *log.Helper
}

func NewDbiniterUsecase(repo DbIniterRepo, logger log.Logger) *DbIniterUsecase {
	return &DbIniterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DbIniterUsecase) InitUserModel(u *User) (rsp *db.InitRpl, err error) {
	return uc.repo.InitUserModel(u)
}

func (uc *DbIniterUsecase) InitDomainRoleModel(u *Role) (rsp *db.InitRpl, err error) {
	return uc.repo.InitDomainRoleModel(u)
}
