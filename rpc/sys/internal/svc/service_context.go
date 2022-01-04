package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"house-repair-api/rpc/model/sysmodel"
	"house-repair-api/rpc/sys/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	RoleModel sysmodel.SysRoleModel
	UserModel sysmodel.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:    c,
		RoleModel: sysmodel.NewSysRoleModel(sqlConn),
		UserModel: sysmodel.NewSysUserModel(sqlConn),
	}
}
