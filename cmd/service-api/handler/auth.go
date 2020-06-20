package handler

import (
	"github.com/A-ndrey/tododo/internal/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AuthHandler struct {
	JWTSecret   []byte
	UserService user.Service
}

type UserClaims struct {
	UserId uint64
	jwt.StandardClaims
}

func RouteAuth(apiGroup *gin.RouterGroup, handler *AuthHandler) {
	apiGroup.POST("/signup", handler.SignUp)
	apiGroup.POST("/signin", handler.SignIn)
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var newUser user.User

	if err := ctx.BindJSON(&newUser); err != nil {
		zap.L().Warn("SignUp",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := h.UserService.Register(newUser)
	if err != nil {
		zap.L().Error("SignUp",
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var loginUser user.User

	if err := ctx.BindJSON(&loginUser); err != nil {
		zap.L().Warn("SignIn",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	userId, err := h.UserService.Login(loginUser)
	if err != nil {
		zap.L().Error("SignIn",
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	claims := UserClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(7 * 24 * time.Hour),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(h.JWTSecret)
	if err != nil {
		zap.L().Error("SignIn",
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Token string `json:"token"`
	}{tokenString})
}
