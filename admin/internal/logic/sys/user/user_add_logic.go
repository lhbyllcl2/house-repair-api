package logic

import (
	"context"
	"house-repair-api/rpc/sys/sysclient"

	"house-repair-api/admin/internal/svc"
	"house-repair-api/admin/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserAddLogic {
	return UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req types.AddUserReq) error {
	_, err := l.svcCtx.Sys.UserAdd(l.ctx, &sysclient.UserAddReq{
		Email:    req.Email,
		Mobile:   req.Mobile,
		Name:     req.Name,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		RoleId:   req.RoleId,
		JobId:    req.JobId,
	})
	if err != nil {
		return err
	}
	return nil
}
