package logic

import (
	"context"

	"house-repair-api/admin/internal/svc"
	"house-repair-api/admin/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleLogic {
	return RoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleLogic) Role(req types.RoleListReq) (*types.RoleListResp, error) {
	// todo: add your logic here and delete this line

	return &types.RoleListResp{}, nil
}
