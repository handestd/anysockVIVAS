package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	MySqlConf struct {
		User      string
		Pass      string
		Host      string
		Port      int64
		DbName    string
		Parameter string
	}
}
