//go:build wireinject

package bff

import (
	"github.com/DaHuangQwQ/gpkg/logger"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitBff(l logger.Logger, redisClient redis.Cmdable) *App {
	wire.Build(
		initGinServer,
		initServer,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
