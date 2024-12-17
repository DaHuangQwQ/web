package web

import (
	"github.com/DaHuangQwQ/web/internal/user"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*UserHandler)(nil)

type UserHandler struct {
	userService *user.App
}

func NewUserHandler(userService *user.App) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, h.userService.Service.GetId())
	})
}
