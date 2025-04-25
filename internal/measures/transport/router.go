package transport

import (
	"testovoe/internal/measures/transport/http/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *handlers.Handler
}

func NewRouter(handler *handlers.Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/measures")
	{
		group.GET("", r.handler.GetAll)
		group.GET("/:id", r.handler.GetByID)
		group.POST("", r.handler.Create)
		group.PUT("/:id", r.handler.Update)
		group.DELETE("/:id", r.handler.Delete)
	}
}
