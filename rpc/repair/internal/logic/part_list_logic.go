package logic

import (
	"context"

	"house-repair-api/rpc/repair/internal/svc"
	"house-repair-api/rpc/repair/repair"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartListLogic {
	return &PartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// part管理
func (l *PartListLogic) PartList(in *repair.PartListReq) (*repair.PartListResp, error) {
	// todo: add your logic here and delete this line

	return &repair.PartListResp{}, nil
}
