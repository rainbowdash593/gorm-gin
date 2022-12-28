package health

import (
	"github.com/gin-gonic/gin"
)

func InitHealthRoutes(router *gin.RouterGroup) {
	commonHandlers := NewHealthHandlers()

	group := router.Group("/health")
	group.GET("/", commonHandlers.PingHandler)
}
