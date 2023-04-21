package main

import (
	"github.com/gin-gonic/gin"
	agentR "github.com/infEreb/nls/netflow/server/agent/pkg/agent/router"
)

const (
	ADDR = ":9000"
	ROOT = "/netflow"
)

func main() {
	r := gin.Default()
	
	agentR.RouterGinInit(r.Group(ROOT), nil)

	r.Run(ADDR)
}