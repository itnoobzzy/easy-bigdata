package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "easyCasbin/api/user/v1"
	"easyCasbin/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserInfo) (*v1.UserInfoResponse, error) {
	user, err := u.uc.CreateUser(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		NickName: user.NickName,
	}, nil
}

func (u *UserService) GetUserList(ctx context.Context, req *v1.PageInfo) (*v1.UserListResponse, error) {
	list, total, err := u.uc.ListUser(ctx, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &v1.UserListResponse{}
	rsp.Total = int32(total)

	for _, user := range list {
		rsp.Data = append(rsp.Data, &v1.UserInfoResponse{
			Id:       user.ID,
			Mobile:   user.Mobile,
			NickName: user.NickName,
		})
	}
	return rsp, nil
}

func (u *UserService) GetUserByMobile(ctx context.Context, req *v1.MobileRequest) (*v1.UserInfoResponse, error) {
	user, err := u.uc.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		NickName: user.NickName,
	}, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserInfo) (*emptypb.Empty, error) {
	user, err := u.uc.UpdateUser(ctx, &biz.User{
		ID:       req.Id,
		NickName: req.NickName,
	})

	if err != nil {
		return nil, err
	}

	if user == false {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (u *UserService) CheckPassword(ctx context.Context, req *v1.PasswordCheckInfo) (*v1.CheckResponse, error) {
	check, err := u.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &v1.CheckResponse{Success: check}, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *v1.IdRequest) (*v1.UserInfoResponse, error) {
	user, err := u.uc.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		NickName: user.NickName,
	}, nil
}

func (u *UserService) RegisterUser(ctx context.Context, req *v1.RegisterRequest) (*emptypb.Empty, error) {
	_, err := u.uc.CreateUser(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
		Active:   1,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (u *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginRpl, error) {
	res, err := u.uc.Login(ctx, req.NickName, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRpl{
		User: &v1.UserInfoResponse{
			Id:       res.User.Id,
			Mobile:   res.User.Mobile,
			NickName: res.User.NickName,
		},
		Token:     res.Token,
		ExpiresAt: res.ExpiresAt,
	}, nil
}
