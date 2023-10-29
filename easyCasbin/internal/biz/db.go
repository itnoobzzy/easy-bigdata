package biz

import (
	"context"
	db "easyCasbin/api/db"
	"github.com/go-kratos/kratos/v2/log"
)

type DbIniterRepo interface {
	InitUserModel(context.Context, *User) (rsp *db.InitRpl, err error)
}

type DbIniterUsecase struct {
	repo DbIniterRepo
	log  *log.Helper
}

func NewDbiniterUsecase(repo DbIniterRepo, logger log.Logger) *DbIniterUsecase {
	return &DbIniterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DbIniterUsecase) InitUserModel(ctx context.Context, u *User) (rsp *db.InitRpl, err error) {
	return uc.repo.InitUserModel(ctx, u)
}
