package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero2/restful/user/internal/config"
	"go-zero2/service/userservice/user"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
