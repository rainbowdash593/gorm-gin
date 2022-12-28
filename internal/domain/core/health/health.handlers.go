package health

import (
	"bridge/users-service/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandlers struct{}

func (h *HealthHandlers) PingHandler(ctx *gin.Context) {
	logger := logging.GetLogger()
	logger.Info("Ping route")
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func NewHealthHandlers() *HealthHandlers {
	return &HealthHandlers{}
}
