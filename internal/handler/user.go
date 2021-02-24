package handler

import (
	"github.com/A-ndrey/tododo/internal/domains"
	"github.com/A-ndrey/tododo/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	UserService services.UserService
}

func RouteUser(apiGroup *gin.RouterGroup, handler *UserHandler) {
	userGroup := apiGroup.Group("/user")

	userGroup.GET("/username", handler.getUsername)
}

func (h *UserHandler) getUsername(ctx *gin.Context) {
	userID, ok := getUserId(ctx)
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	username, err := h.UserService.GetUsernameByID(userID)
	if err != nil {
		zap.L().Error("GetUsername",
			zap.Uint64("user ID", userID),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError) // todo error type
		return
	}

	ctx.JSON(http.StatusOK, domains.User{Username: username})
}
