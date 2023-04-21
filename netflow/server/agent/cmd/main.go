package main

import (
	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/agent/pkg/router"
	"github.com/infEreb/nls/netflow/server/pkg/netflow"
)

const ADDR = ":9000"

func main() {
	r := gin.Default()
	root := r.Group(netflow.ROOT)
	router.RouterGinInit(root, nil)

	r.Run(ADDR)
}