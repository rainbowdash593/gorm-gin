//go:build wireinject
// +build wireinject

package users

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(db *gorm.DB) *UserHandlers {
	wire.Build(ProvideUserRepository, ProvideUserService, ProvideUserHandlers)
	return &UserHandlers{}
}
