package logic

import (
	"context"

	"house-repair-api/rpc/sys/internal/svc"
	"house-repair-api/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 角色管理
func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	// todo: add your logic here and delete this line

	return &sys.RoleListResp{}, nil
}
