package svc

import (
	"gorm.io/gorm"

	"house-repair-api/pkg/db"
	"house-repair-api/rpc/repair/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db.Open(c.DataSourceName, "repair_"),
	}
}
