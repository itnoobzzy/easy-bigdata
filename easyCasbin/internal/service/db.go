package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	db "easyCasbin/api/db"
	"easyCasbin/internal/biz"
)

//var ProviderSet = wire.NewSet(NewDbIniter)

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

func (d *DbIniterService) InitUserDB(ctx context.Context, req *emptypb.Empty) (rsp *db.InitRpl, err error) {
	return d.uc.InitUserModel(ctx, &biz.User{})
}

func (d *DbIniterService) InitRoleDB(ctx context.Context, req *emptypb.Empty) (rsp *db.InitRpl, err error) {
	return &db.InitRpl{
		Code:    0,
		Message: "",
	}, nil
}
