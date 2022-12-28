package main

import (
	"bridge/users-service/internal/config"
	"bridge/users-service/internal/database"
	"bridge/users-service/internal/domain/core/health"
	"bridge/users-service/internal/domain/users"
	"bridge/users-service/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.GetLogger()

	db := database.InitDB()

	router := gin.Default()
	api := router.Group("/api/v1")

	health.InitHealthRoutes(api)

	userHandlers := users.Wire(db)
	users.InitUserRoutes(api, userHandlers)

	logger.Info(fmt.Sprintf("app started on: %s:%d", cfg.HTTP.Host, cfg.HTTP.Port))
	log.Fatal(router.Run(fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)))
}
