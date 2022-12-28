package main

import (
	"bridge/users-service/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.InitDB()
	database.Migrate(db)
}
