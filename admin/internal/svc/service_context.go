package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"

	"house-repair-api/admin/internal/config"
	"house-repair-api/admin/internal/middleware"
	"house-repair-api/rpc/sys/sysclient"
)

type ServiceContext struct {
	Config         config.Config
	Sys            sysclient.Sys
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address)
	return &ServiceContext{
		Config:         c,
		Sys:            sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		AuthMiddleware: middleware.NewAuthMiddleware(newRedis).Handle,
	}
}
