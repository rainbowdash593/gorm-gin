package database

import (
	"bridge/users-service/internal/domain/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(users.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}
