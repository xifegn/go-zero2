package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero2/internal/model/usermodel"
	"go-zero2/service/userservice/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel usermodel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: usermodel.NewUserModel(conn, c.CacheRedis),
	}
}
