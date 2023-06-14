package main

import (
	"anysock/internal/config"
	"anysock/internal/handler"
	"anysock/internal/svc"
	"anysock/pkg/cache"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func main() {

	cache.Connect()

	flag.Parse()

	var c config.Config
	conf.MustLoad(*config.ConfigFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	//my code

}
