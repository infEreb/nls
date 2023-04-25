package router

import (
	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/app/agent/internal"
	"github.com/infEreb/nls/netflow/server/app/agent/internal/handler"
	"github.com/infEreb/nls/netflow/server/app/agent/internal/repository"
	"github.com/infEreb/nls/netflow/server/app/agent/pkg/agent"
)

func httpHandler(c *agent.ConfigAgent) *handler.HandlerHTTP {
	repo := repository.NewRepositoryMemory()
	ctrl := internal.NewController(repo)
	return handler.NewHandlerHTTP(ctrl)
}


func RouterGinInit(r *gin.RouterGroup, c *agent.ConfigAgent) {
	h := httpHandler(c)
	
	routes := r.Group(agent.ROOT)


	routes.GET("/", func(ctx *gin.Context){
		h.GetAgents(ctx.Writer, ctx.Request)
	})
	routes.GET("/:id", func(ctx *gin.Context){
		h.GetAgent(ctx.Writer, ctx.Request, ctx.Param("id"))
	})

	routes.POST("/", func(ctx *gin.Context){
		h.PostAgent(ctx.Writer, ctx.Request)
	})

	routes.PUT("/:id", func(ctx *gin.Context){
		h.PutAgent(ctx.Writer, ctx.Request, ctx.Param("id"))
	})

	routes.DELETE("/:id", func(ctx *gin.Context){
		h.DeleteAgent(ctx.Writer, ctx.Request, ctx.Param("id"))
	})
}
