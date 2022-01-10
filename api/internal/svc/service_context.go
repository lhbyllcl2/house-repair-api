package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"house-repair-api/api/internal/config"
	"house-repair-api/api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
