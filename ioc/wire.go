//go:build wireinject

package ioc

import (
	"github.com/DaHuangQwQ/web/internal/bff"
	"github.com/DaHuangQwQ/web/internal/user"
	"github.com/DaHuangQwQ/web/ioc/third"

	"github.com/google/wire"
)

var thirdSet = wire.NewSet(
	third.InitLogger,
	third.InitRedis,
)

func InitApp() *App {
	wire.Build(
		thirdSet,

		bff.InitApp,
		wire.FieldsOf(new(*bff.App), "Server"),

		user.InitApp,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
