package initialize

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/emptypb"

	db "easyCasbin/api/db"
	"easyCasbin/internal/biz"
	"easyCasbin/internal/data"
)

var ProviderSet = wire.NewSet(NewDbIniter)

type dbIniter struct {
	userModel *biz.User

	log  *log.Helper
	data *data.Data
}

func NewDbIniter(um *biz.User, data *data.Data, logger log.Logger) *dbIniter {
	return &dbIniter{
		userModel: um,
		log:       log.NewHelper(logger),
		data:      data,
	}
}

func (d *dbIniter) InitUserDB(ctx context.Context, req *emptypb.Empty) (rsp *db.InitRpl, err error) {
	if err = d.data.Db.AutoMigrate(&d.userModel); err != nil {
		return &db.InitRpl{
			Code:    1,
			Message: err.Error(),
		}, err
	}
	return &db.InitRpl{
		Code:    0,
		Message: "init user db success!",
	}, nil
}

func (d *dbIniter) InitRoleDB(ctx context.Context, req *emptypb.Empty) (rsp *db.InitRpl, err error) {
	return &db.InitRpl{
		Code:    0,
		Message: "",
	}, nil
}
