package svc

import (
	"anysock/internal/config"
	"anysock/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql("root:@tcp(localhost:3306)/anysock?charset=utf8mb4&parseTime=True&loc=Local")),
	}
}
