package router

import (
	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/ethernet/internal"
	"github.com/infEreb/nls/netflow/server/ethernet/internal/gateway"
	"github.com/infEreb/nls/netflow/server/ethernet/internal/handler"
	"github.com/infEreb/nls/netflow/server/ethernet/internal/repository"
	"github.com/infEreb/nls/netflow/server/ethernet/pkg/ethernet"
)

func httpHandler(c *ethernet.ConfigEthernet, ag gateway.GatewayAgent) *handler.HandlerHTTP {
	repo := repository.NewRepositoryMemory()
	ctrl := internal.NewController(repo, ag)
	return handler.NewHandlerHTTP(ctrl)
}


func RouterGinInit(r *gin.RouterGroup, c *ethernet.ConfigEthernet, agentGateway gateway.GatewayAgent) {
	h := httpHandler(c, agentGateway)
	
	routes := r.Group(ethernet.ROOT)


	routes.GET("/", func(ctx *gin.Context){
		h.GetEthernets(ctx.Writer, ctx.Request)
	})
	routes.GET("/:id", func(ctx *gin.Context){
		h.GetEthernet(ctx.Writer, ctx.Request, ctx.Param("id"))
	})

	routes.POST("/", func(ctx *gin.Context){
		h.PostEthernet(ctx.Writer, ctx.Request)
	})

	routes.PUT("/:id", func(ctx *gin.Context){
		h.PutEthernet(ctx.Writer, ctx.Request, ctx.Param("id"))
	})

	routes.DELETE("/:id", func(ctx *gin.Context){
		h.DeleteEthernet(ctx.Writer, ctx.Request, ctx.Param("id"))
	})
}