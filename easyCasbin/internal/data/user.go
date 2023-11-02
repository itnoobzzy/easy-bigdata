package data

import (
	"context"
	"crypto/sha512"
	"easyCasbin/utils"
	"errors"
	"fmt"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"easyCasbin/api/user/v1"
	"easyCasbin/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) encryptPwd(pwd string) string {
	options := &password.Options{
		SaltLen:      16,
		Iterations:   1000,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodepwd := password.Encode(pwd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodepwd)
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user biz.User
	result := r.data.db.Where(&biz.User{Mobile: u.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, v1.ErrorUserExist("user %s exist", u.Mobile)
	}

	user.Username = u.Username
	user.NickName = u.NickName
	user.Mobile = u.Mobile
	user.Password = utils.BcryptHash(u.Password)
	user.Active = u.Active
	res := r.data.db.Create(&user)
	if res.Error != nil {
		return nil, v1.ErrorInternalErr("create user failed: %v!", res.Error)
	}
	return &user, nil
}

func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (r *userRepo) ListUser(ctx context.Context, pageNum, pageSize int) ([]*biz.User, int, error) {
	var users []biz.User
	result := r.data.db.Find(&users)
	rv := make([]*biz.User, 0)
	total := 0
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return rv, total, nil
	}
	if result.Error != nil {
		return nil, 0, v1.ErrorInternalErr("FIND_USER_ERR: %v", result.Error)
	}
	total = int(result.RowsAffected)
	r.data.db.Scopes(paginate(pageNum, pageSize)).Find(&users)
	for _, u := range users {
		rv = append(rv, &biz.User{
			ID:       u.ID,
			Mobile:   u.Mobile,
			Password: u.Password,
			NickName: u.NickName,
		})
	}
	return rv, total, nil
}

func (r *userRepo) UserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	var user biz.User
	result := r.data.db.Where(&biz.User{Mobile: mobile}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, nil
	}
	if result.Error != nil {
		return nil, v1.ErrorInternalErr("FIND_USER_ERR: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return &user, nil
	}
	return &user, nil
}

func (r *userRepo) GetUserById(ctx context.Context, Id int64) (*biz.User, error) {
	var user biz.User
	if err := r.data.db.Where(&biz.User{ID: Id}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, v1.ErrorInternalErr("FIND_USER_ERR: %v", err.Error())
	}
	return &user, nil
}

func (r *userRepo) GetUserByName(ctx context.Context, name string) (*biz.User, error) {
	var user biz.User
	if err := r.data.db.Where(&biz.User{Username: name}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, v1.ErrorInternalErr("FIND_USER_ERR: %v", err.Error())
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (bool, error) {
	var userInfo biz.User
	result := r.data.db.Where(&biz.User{ID: user.ID}).First(&userInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, v1.ErrorUserNotFound("USER_NOT_FOUND")
	}

	if result.RowsAffected == 0 {
		return false, v1.ErrorInternalErr("FIND_USER_ERR: %v", result.Error)
	}

	userInfo.NickName = user.NickName

	if err := r.data.db.Save(&userInfo).Error; err != nil {
		return false, v1.ErrorInternalErr("UPDATE_USER_ERR: %v", result.Error)
	}
	return true, nil
}
