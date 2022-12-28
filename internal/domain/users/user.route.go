package users

import (
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(route *gin.RouterGroup, handlers *UserHandlers) {
	group := route.Group("/users")
	group.GET("/:id", handlers.Find)
	group.PUT("/", handlers.Create)
	group.POST("/:id", handlers.Update)
	group.DELETE("/:id", handlers.Delete)
}
