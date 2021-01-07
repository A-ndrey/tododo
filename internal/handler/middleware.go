package handler

import (
	"github.com/A-ndrey/tododo/internal/auth"
	"github.com/A-ndrey/tododo/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserIDKey = "userID"

func AuthMiddleware(userService user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		email, err := auth.GetUserEmail(authHeader)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, err := userService.GetIDByEmail(email)
		if err != nil {
			userID, err = userService.SaveNewUser(email)
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError) // todo
				return
			}
		}

		ctx.Set(UserIDKey, userID)
	}
}