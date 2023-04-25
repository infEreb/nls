package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/app/agent/pkg/router"
	"github.com/infEreb/nls/netflow/server/pkg/api"
	"github.com/infEreb/nls/netflow/server/pkg/netflow"
)

func main() {
	flags, err := api.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	conf, err := api.ParseConfig(flags.Config)
	if err != nil {
		log.Fatalln(err)
	}

	// logger config...

	r := gin.Default()
	root := r.Group(netflow.ROOT)
	router.RouterGinInit(root, nil)

	r.Run(fmt.Sprintf(":%d", conf.Api.Port))
}