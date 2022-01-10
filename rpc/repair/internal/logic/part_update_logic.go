package logic

import (
	"context"

	"house-repair-api/rpc/repair/internal/svc"
	"house-repair-api/rpc/repair/repair"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartUpdateLogic {
	return &PartUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PartUpdateLogic) PartUpdate(in *repair.PartUpdateReq) (*repair.Response, error) {
	// todo: add your logic here and delete this line

	return &repair.Response{}, nil
}
