package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	db "easyCasbin/api/db"
	"easyCasbin/internal/biz"
)

type DbIniterService struct {
	uc *biz.DbIniterUsecase

	log *log.Helper
}

func NewDbIniterService(uc *biz.DbIniterUsecase, logger log.Logger) *DbIniterService {
	return &DbIniterService{
		log: log.NewHelper(logger),
		uc:  uc,
	}
}

func (d *DbIniterService) InitUserDB(context.Context, *emptypb.Empty) (rsp *db.InitRpl, err error) {
	return d.uc.InitUserModel(&biz.User{})
}

func (d *DbIniterService) InitRoleDB(context.Context, *emptypb.Empty) (rsp *db.InitRpl, err error) {
	return d.uc.InitDomainRoleModel(&biz.Role{})
}
