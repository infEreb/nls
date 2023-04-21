package router

import (
	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/agent/internal"
	"github.com/infEreb/nls/netflow/server/agent/internal/handler"
	"github.com/infEreb/nls/netflow/server/agent/internal/repository"
	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
)

func httpHandler(c *agent.ConfigAgent) *handler.HandlerHTTP {
	repo := repository.NewArrayRepository()
	ctrl := internal.NewController(repo)
	return handler.NewHandlerHTTP(ctrl)
}

// func RouterGinInit(r *gin.Engine, c *agent.ConfigAgent) {
// 	h := handler(c)
	
// 	routes := r.Group(agent.ROOT)
// 	routes.GET("/", func(gctx *gin.Context){
// 		h.GetAgents(gctx.Writer, gctx.Request)
// 	})

// 	routes.POST("/", func(gctx *gin.Context){
// 		h.PostAgent(gctx.Writer, gctx.Request)
// 	})
// }

func RouterGinInit(r *gin.RouterGroup, c *agent.ConfigAgent) {
	h := httpHandler(c)
	
	routes := r.Group(agent.ROOT)


	routes.GET("/", h.GetAgents)
	routes.GET("/:id", h.GetAgent)

	routes.POST("/", h.PostAgent)

	routes.PUT("/:id", h.PutAgent)

	routes.DELETE("/:id", h.DeleteAgent)
}