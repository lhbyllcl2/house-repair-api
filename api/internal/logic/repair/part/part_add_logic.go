package logic

import (
	"context"

	"house-repair-api/api/internal/svc"
	"house-repair-api/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) PartAddLogic {
	return PartAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartAddLogic) PartAdd(req types.PartAddReq) error {
	// todo: add your logic here and delete this line

	return nil
}
