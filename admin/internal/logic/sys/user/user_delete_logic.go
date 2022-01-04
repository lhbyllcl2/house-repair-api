package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"house-repair-api/admin/internal/svc"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserDeleteLogic {
	return UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete() error {
	// todo: add your logic here and delete this line

	return nil
}
