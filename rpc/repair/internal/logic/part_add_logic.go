package logic

import (
	"context"

	"house-repair-api/model/repairmodel"
	"house-repair-api/rpc/repair/internal/svc"
	"house-repair-api/rpc/repair/repair"

	"github.com/tal-tech/go-zero/core/logx"
)

type PartAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPartAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartAddLogic {
	return &PartAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PartAddLogic) PartAdd(in *repair.PartAddReq) (*repair.Response, error) {
	model := repairmodel.NewRepairPartModel(l.svcCtx.Db)
	_ = model.Insert(repairmodel.RepairPart{})
	return &repair.Response{}, nil
}
