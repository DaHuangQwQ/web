// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"github.com/DaHuangQwQ/web/internal/user/service"
)

// Injectors from wire.go:

func InitApp() *App {
	userService := service.NewUserService()
	app := &App{
		Service: userService,
	}
	return app
}
