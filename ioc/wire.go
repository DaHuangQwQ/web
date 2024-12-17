//go:build wireinject

package ioc

import (
	"main/internal/bff"
	"main/ioc/third"

	"github.com/google/wire"
)

var thirdSet = wire.NewSet(
	third.InitLogger,
	third.InitRedis,
)

func InitApp() *App {
	wire.Build(
		thirdSet,

		bff.InitBff,
		wire.FieldsOf(new(*bff.App), "Server"),

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
