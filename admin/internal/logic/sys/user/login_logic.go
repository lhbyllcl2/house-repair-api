package logic

import (
	"context"
	"encoding/json"
	"errors"
	"house-repair-api/admin/internal/svc"
	"house-repair-api/admin/internal/types"
	"house-repair-api/rpc/sys/sysclient"
	"strings"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.UserReply, error) {
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("用户名或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, errors.New("用户名或密码不能为空")
	}
	resp, err := l.svcCtx.Sys.Login(l.ctx, &sysclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据用户名: %s和密码: %s查询用户异常:%s", req.UserName, req.Password, err.Error())
		return nil, errors.New("登录失败")
	}
	//保存登录日志
	return &types.UserReply{
		Uid:      resp.Id,
		Username: resp.UserName,
		Avator:   "",
		JwtToken: types.JwtToken{
			AccessToken:  resp.AccessToken,
			AccessExpire: resp.AccessExpire,
			RefreshAfter: resp.RefreshAfter,
		},
	}, nil
}
