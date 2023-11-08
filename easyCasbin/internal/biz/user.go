package biz

import (
	"context"
	v1 "easyCasbin/api/user/v1"
	"easyCasbin/internal/conf"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"

	"database/sql/driver"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type JSON []byte

func (c JSON) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSON) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type User struct {
	ID             int64      `gorm:"primarykey"`
	FirstName      string     `json:"first_name" gorm:"type:text"`
	LastName       string     `json:"last_name" gorm:"type:text"`
	Username       string     `json:"username" gorm:"type:text"`
	Password       string     `json:"password" gorm:"comment:密码"`
	Active         int        `json:"active" gorm:"type:tinyint(1);not null;comment:是否激活"`
	Email          string     `json:"email" gorm:"type:text;comment:邮箱"`
	LastLogin      *LocalTime `json:"last_login" gorm:"type:time;comment:last_login"`
	LoginCount     int        `json:"login_count" gorm:"size:11;comment:login_count"`
	FailLoginCount int        `json:"fail_login_count" gorm:"size:11;comment:fail_login_count"`
	Params         JSON       `json:"params" gorm:"type:json;serializer:json"`
	CreatedByFk    int        `json:"created_by_fk" gorm:"size:11;comment:创建人id"`
	ChangedByFk    int        `json:"changed_by_fk" gorm:"size:11;comment:修改人id"`
	RegisterFrom   string     `json:"register_from" gorm:"type:text"`
	NickName       string     `json:"nick_name" gorm:"size:64"`
	DepartmentPath string     `json:"department_path" gorm:"size:200"`
	Mobile         string     `json:"mobile" gorm:"size:20"`
	Gender         string     `json:"gender" gorm:"size:20"`
	Position       string     `json:"position" gorm:"size:100"`
	ThumbAvatar    string     `json:"thumb_avatar" gorm:"size:1000"`
	WxUsername     string     `json:"wx_username" gorm:"size:100"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (User) TableName() string {
	return "user"
}

//go:generate mockgen -destination=user_mock.go -package=biz . UserRepo
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	GetUserByName(ctx context.Context, name string) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
	sc   *conf.Server
}

func NewUserUsecase(repo UserRepo, logger log.Logger, sc *conf.Server) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger), sc: sc}
}

func (uc *UserUsecase) Login(ctx context.Context, loginReq *LoginRequest,
	encryptService EncryptService) (*LoginResponse, error) {
	// 获取用户信息
	user, err := uc.repo.GetUserByName(ctx, loginReq.username)
	if err != nil {
		return nil, err
	}
	// 校验密码
	check := encryptService.CheckPassword(ctx, loginReq.password, user.Password)
	if err != nil || !check {
		return nil, v1.ErrorPasswordErr("password error!")
	}
	// 颁发token
	token, err := encryptService.Token(ctx, user)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		User:      UserInfoResponse{user.ID, user.Mobile, user.Username},
		Token:     token,
		ExpiresAt: (time.Now().Unix() + uc.sc.Jwt.ExpiresTime) * 1000,
	}, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUsecase) ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error) {
	ctx = context.WithValue(ctx, "serverConfig", uc.sc)
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUsecase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, user)
}
