package logic

import (
	"context"

	"house-repair-api/rpc/sys/internal/svc"
	"house-repair-api/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户管理
func (l *UserListLogic) UserList(in *sys.UserListReq) (*sys.UserListResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UserListResp{}, nil
}
