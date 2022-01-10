package logic

import (
	"context"

	"house-repair-api/rpc/repair/internal/svc"
	"house-repair-api/rpc/repair/repair"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPartDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartDeleteLogic {
	return &PartDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PartDeleteLogic) PartDelete(in *repair.PartDeleteReq) (*repair.Response, error) {
	// todo: add your logic here and delete this line

	return &repair.Response{}, nil
}
