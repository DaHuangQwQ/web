package bff

import (
	"context"
	"main/internal/bff/web"
	"strings"
	"time"

	"github.com/DaHuangQwQ/gpkg/ginx"
	"github.com/DaHuangQwQ/gpkg/ginx/jwt"
	"github.com/DaHuangQwQ/gpkg/ginx/middleware/jwt_token"
	prometheusx "github.com/DaHuangQwQ/gpkg/ginx/middleware/prometheus"
	"github.com/DaHuangQwQ/gpkg/ginx/middleware/ratelimit"
	"github.com/DaHuangQwQ/gpkg/logger"
	ratelimiter "github.com/DaHuangQwQ/gpkg/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type App struct {
	Server *ginx.Server
}

func initServer() []web.Handler {
	return []web.Handler{
		web.NewUserHandler(),
	}
}

func initGinServer(
	redisClient redis.Cmdable,
	l logger.Logger,
	webHandler ...web.Handler,
) *ginx.Server {
	pb := &prometheusx.Builder{
		Namespace: "DaHuang",
		Subsystem: "webook",
		Name:      "gin_http",
		Help:      "统计 GIN 的HTTP接口数据",
	}
	engine := gin.Default()
	engine.Use(
		pb.BuildResponseTime(),
		pb.BuildActiveRequest(),
		corsHdl(),
		timeout(),
		otelgin.Middleware("webook"),
		ratelimit.NewBuilder(ratelimiter.NewRedisSlidingWindowLimiter(redisClient, time.Second, 10), l).Build(),
		jwt_token.NewBuilder(jwt.NewRedisJWTHandler([]byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm"),
			[]byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixA"), time.Hour*24, redisClient)).Build())

	for _, h := range webHandler {
		h.RegisterRoutes(engine)
	}

	ginx.InitCounter(prometheus.CounterOpts{
		Namespace: "daming_geektime",
		Subsystem: "webook_bff",
		Name:      "http",
	})
	ginx.NewWarpLogger(l)
	return &ginx.Server{
		Engine: engine,
	}
}

func timeout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, ok := ctx.Request.Context().Deadline()
		if !ok {
			// 强制给一个超时，省得我前端调试等得不耐烦
			newCtx, cancel := context.WithTimeout(ctx.Request.Context(), time.Second*10)
			defer cancel()
			ctx.Request = ctx.Request.Clone(newCtx)
		}
		ctx.Next()
	}
}

func corsHdl() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		//AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 你不加这个，前端是拿不到的
		ExposeHeaders: []string{"x-jwt-token", "x-refresh-token"},
		// 是否允许你带 cookie 之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 你的开发环境
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	})
}
