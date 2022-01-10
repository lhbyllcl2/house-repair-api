package logic

import (
	"context"

	"house-repair-api/api/internal/svc"
	"house-repair-api/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) PartListLogic {
	return PartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartListLogic) PartList(req types.PartListReq) (*types.PartListResp, error) {
	// todo: add your logic here and delete this line

	return &types.PartListResp{}, nil
}
