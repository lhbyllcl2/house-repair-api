package logic

import (
	"context"

	"house-repair-api/admin/internal/svc"
	"house-repair-api/admin/internal/types"
	"house-repair-api/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleAddLogic {
	return RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req types.RoleAddReq) error {
	_, err := l.svcCtx.Sys.RoleAdd(l.ctx, &sysclient.RoleAddReq{
		Name:   req.Name,
		Remark: req.Remark,
		Status: req.Status,
	})
	if err != nil {
		return err
	}
	return nil
}
