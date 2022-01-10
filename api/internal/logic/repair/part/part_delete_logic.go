package logic

import (
	"context"

	"house-repair-api/api/internal/svc"
	"house-repair-api/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) PartDeleteLogic {
	return PartDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartDeleteLogic) PartDelete(req types.PartDeleteReq) error {
	// todo: add your logic here and delete this line

	return nil
}
