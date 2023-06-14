package config

import (
	"flag"
	"github.com/zeromicro/go-zero/rest"
)

var ConfigFile = flag.String("f", "etc/user.yaml", "the config file")

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

	RedisConf struct {
		Host string
	}
}
