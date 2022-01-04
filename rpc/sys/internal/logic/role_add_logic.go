package logic

import (
	"context"
	"errors"
	"house-repair-api/rpc/model/sysmodel"

	"house-repair-api/rpc/sys/internal/svc"
	"house-repair-api/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

var ErrorRoleIsExist = errors.New("角色名已存在")

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.Response, error) {
	if isExist := l.svcCtx.RoleModel.NameIsExist(in.Name); isExist {
		return &sys.Response{}, ErrorRoleIsExist
	}
	_, err := l.svcCtx.RoleModel.Insert(sysmodel.SysRole{
		Name:   in.Name,
		Remark: in.Remark,
		Status: in.Status,
	})
	if err != nil {
		return &sys.Response{}, err
	}
	return &sys.Response{}, nil
}
