package logic

import (
	"context"

	"house-repair-api/rpc/sys/internal/svc"
	"house-repair-api/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *sys.UserDeleteReq) (*sys.Response, error) {
	// todo: add your logic here and delete this line

	return &sys.Response{}, nil
}
