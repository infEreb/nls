package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/app/ethernet/internal/gateway"
	"github.com/infEreb/nls/netflow/server/app/ethernet/pkg/router"
	"github.com/infEreb/nls/netflow/server/pkg/api"
	"github.com/infEreb/nls/netflow/server/pkg/discovery"
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

	reg := discovery.NewRegistryK8s(flags.Type)

	// logger config...

	r := gin.Default()
	root := r.Group(netflow.ROOT)
	router.RouterGinInit(root, nil, gateway.NewAgentGatewayHTTP(reg))

	r.Run(fmt.Sprintf(":%d", conf.Api.Port))
}