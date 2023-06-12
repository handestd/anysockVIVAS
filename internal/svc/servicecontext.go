package svc

import (
	"anysock/internal/config"
	"anysock/internal/middleware"
	"anysock/internal/model"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"strconv"
)

type ServiceContext struct {
	Config              config.Config
	UserModel           model.UserModel
	UserAgentMiddleware rest.Middleware
	LogMiddleware       rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	port := strconv.Itoa(int(c.MySqlConf.Port))
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.MySqlConf.User, "", c.MySqlConf.Host, port, c.MySqlConf.DbName, c.MySqlConf.Parameter)

	return &ServiceContext{
		Config: c,
		//UserModel: model.NewUserModel(sqlx.NewMysql("root:@tcp(localhost:3306)/anysock?charset=utf8mb4&parseTime=True&loc=Local")),
		UserModel:           model.NewUserModel(sqlx.NewMysql(dataSource)),
		UserAgentMiddleware: middleware.NewUserAgentMiddleware().Handle,
		LogMiddleware:       middleware.NewLogMiddleware().Handle,
	}
}
