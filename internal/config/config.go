package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct { // Key and expiration time configuration required for JWT authentication
		AccessSecret string
		AccessExpire int64
	}
}
