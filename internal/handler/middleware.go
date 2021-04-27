package handler

import (
	"github.com/A-ndrey/tododo/internal/config"
	"github.com/A-ndrey/tododo/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserIDKey = "userID"

func AuthMiddleware(userService services.UserService, authService services.AuthService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		email, err := authService.GetUserEmail(authHeader)
		if err != nil {
			authConfig := config.GetAuth()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, struct {
				AuthService string `json:"authService"`
				AuthHost    string `json:"authHost"`
			}{
				AuthService: authConfig.Service,
				AuthHost:    authConfig.Host,
			})
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

func CorsMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	}
}
