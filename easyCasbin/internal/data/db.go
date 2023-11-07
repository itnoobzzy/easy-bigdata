package data

import (
	db "easyCasbin/api/db"
	"easyCasbin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type dbIniterRepo struct {
	data *Data
	log  *log.Helper
}

func NewDbIniterRepo(data *Data, logger log.Logger) biz.DbIniterRepo {
	return &dbIniterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d *dbIniterRepo) InitUserModel(user *biz.User) (rsp *db.InitRpl, err error) {
	if err = d.data.db.AutoMigrate(&user); err != nil {
		return &db.InitRpl{
			Code:    1,
			Message: err.Error(),
		}, err
	}
	d.log.Info("init user table success!")
	return &db.InitRpl{
		Code:    0,
		Message: "init user db success!",
	}, nil
}

func (d *dbIniterRepo) InitDomainRoleModel(role *biz.Role) (rsp *db.InitRpl, err error) {
	if err = d.data.db.AutoMigrate(&role); err != nil {
		return &db.InitRpl{
			Code:    1,
			Message: err.Error(),
		}, err
	}
	d.log.Info("init role table success!")
	return &db.InitRpl{
		Code:    0,
		Message: "init role db success!",
	}, nil
}
