//go:build wireinject

package bff

import (
	"github.com/DaHuangQwQ/gpkg/logger"
	"github.com/DaHuangQwQ/web/internal/bff/web"
	"github.com/DaHuangQwQ/web/internal/user"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitApp(l logger.Logger, redisClient redis.Cmdable,
	userService *user.App,
) *App {
	wire.Build(
		initGinServer,

		web.NewUserHandler,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
