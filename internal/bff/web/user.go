package web

import (
	"github.com/gin-gonic/gin"
)

var _ Handler = (*UserHandler)(nil)

type UserHandler struct {
}

func NewUserHandler() Handler {
	return &UserHandler{}
}

func (u UserHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World")
	})
}
