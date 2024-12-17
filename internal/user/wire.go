//go:build wireinject

package user

import (
	"github.com/DaHuangQwQ/web/internal/user/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		service.NewUserService,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
