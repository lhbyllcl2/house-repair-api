package logic

import (
	"context"

	"house-repair-api/api/internal/svc"
	"house-repair-api/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) PartUpdateLogic {
	return PartUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartUpdateLogic) PartUpdate(req types.PartUpdateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
