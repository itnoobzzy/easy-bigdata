package data

import (
	"context"
	v1 "easyCasbin/api/user/v1"
	"easyCasbin/internal/biz"
	"easyCasbin/internal/conf"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	tdb  *gorm.DB
	trdb *redis.Client
)

// TestMain 是在当前package下，最先运行的一个函数，常用于初始化
func TestMain(m *testing.M) {
	c := &conf.Data{
		Database: &conf.Data_Database{
			Driver: "mysql",
			Source: "dbadmin:hE4sqSfuCQeXEXwz@tcp(rm-3nsc58907o3epw2me.mysql.rds.aliyuncs.com:3306)/easyBigdata?charset=utf8mb4&parseTime=True&loc=Local",
			DbName: "easyBigdata",
		},
		Redis: &conf.Data_Redis{
			Addr: "",
			ReadTimeout: &duration.Duration{
				Seconds: 2,
				Nanos:   0,
			},
			WriteTimeout: &duration.Duration{
				Seconds: 2,
				Nanos:   0,
			},
		},
	}
	tdb = NewDB(c)
	trdb = NewRedis(c)
	os.Exit(m.Run())
}

func TestUserRepo_CreateUser(t *testing.T) {
	ctx := context.Background()
	repo := NewUserRepo(&Data{
		db:  tdb,
		rdb: trdb,
	}, log.DefaultLogger)

	cases := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
		ctx     context.Context
		user    *biz.User
	}{
		{
			name: "normal",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.NoError(t, err)
				return true
			},
			ctx: ctx,
			user: &biz.User{
				Username: "zzy",
				NickName: "志勇",
				Mobile:   "17720495379",
				Password: "123",
				Active:   0,
			},
		},
		{
			name: "user exists",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.ErrorIs(t, err, v1.ErrorUserExist("user exists"))
				return false
			},
			ctx: ctx,
			user: &biz.User{
				Username: "zzy",
				NickName: "志勇",
				Mobile:   "17720495379",
				Password: "123",
				Active:   0,
			},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			user, err := repo.CreateUser(item.ctx, item.user)
			if !item.wantErr(t, err) {
				return
			}
			assert.NotNil(t, user)
			assert.Equal(t, item.user.Mobile, user.Mobile)
		})
	}
}
