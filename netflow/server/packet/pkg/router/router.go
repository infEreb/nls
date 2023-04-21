package router

import (
	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/packet/internal"
	"github.com/infEreb/nls/netflow/server/packet/internal/handler"
	"github.com/infEreb/nls/netflow/server/packet/internal/repository"
	"github.com/infEreb/nls/netflow/server/packet/pkg/packet"
)

func httpHandler(c *packet.ConfigPacket) *handler.HandlerHTTP {
	repo := repository.NewRepositoryMemory()
	ctrl := internal.NewController(repo)
	return handler.NewHandlerHTTP(ctrl)
}


func RouterGinInit(r *gin.RouterGroup, c *packet.ConfigPacket) {
	h := httpHandler(c)
	
	routes := r.Group(packet.ROOT)


	routes.GET("/", func(ctx *gin.Context){
		h.GetPackets(ctx.Writer, ctx.Request)
	})
	routes.GET("/:id", func(ctx *gin.Context){
		h.GetPacket(ctx.Writer, ctx.Request, ctx.Param("id"))
	})

	routes.POST("/", func(ctx *gin.Context){
		h.PostPacket(ctx.Writer, ctx.Request)
	})

	routes.PUT("/:id", func(ctx *gin.Context){
		h.PutPacket(ctx.Writer, ctx.Request, ctx.Param("id"))
	})

	routes.DELETE("/:id", func(ctx *gin.Context){
		h.DeletePacket(ctx.Writer, ctx.Request, ctx.Param("id"))
	})
}
