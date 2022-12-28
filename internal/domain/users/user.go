package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id,primary_key" json:"id"`
	UUID      uuid.UUID      `json:"uuid" gorm:"index:uuid_uniq,unique"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"-"`
	IsBlocked bool           `json:"is_blocked"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (u *User) BeforeCreate(*gorm.DB) error {
	u.UUID = uuid.New()
	return nil
}
