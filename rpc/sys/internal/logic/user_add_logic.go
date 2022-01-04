package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stringx"
	"house-repair-api/rpc/model/sysmodel"
	"house-repair-api/rpc/pkg/password"
	"house-repair-api/rpc/sys/internal/svc"
	"house-repair-api/rpc/sys/sys"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (*sys.Response, error) {
	//check whether the user exists
	user, err := l.svcCtx.UserModel.UserInfoWithFields(in.Name, in.Email, in.Mobile)
	if err != nil {
		return &sys.Response{}, err
	}
	if in.Name == user.Name {
		return &sys.Response{}, errors.New("用户名已被注册")
	}
	if in.Email == user.Email {
		return &sys.Response{}, errors.New("邮箱已被注册")
	}
	if in.Mobile == user.Mobile {
		return &sys.Response{}, errors.New("手机号已被注册")
	}
	salt := stringx.Randn(6)
	_, err = l.svcCtx.UserModel.Insert(sysmodel.SysUser{
		NickName: in.Name,
		Avatar:   in.Avatar,
		Password: password.GeneratePassword("123456", salt),
		Salt:     salt,
		Email:    in.Email,
		Mobile:   in.Mobile,
		Status:   1,
		CreateBy: "admin",
		JobId:    0,
		Name:     in.Name,
	})
	if err != nil {
		return &sys.Response{}, err
	}
	return &sys.Response{}, nil
}
