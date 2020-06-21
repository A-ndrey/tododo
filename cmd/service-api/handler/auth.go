package handler

import (
	"fmt"
	"github.com/A-ndrey/tododo/internal/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

const tokenDuration = 7 * 24 * time.Hour

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
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
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

func (h *AuthHandler) AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	prefix := "Bearer "

	if !strings.HasPrefix(authHeader, prefix) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len(prefix):]

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return h.JWTSecret, nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", claims.UserId)
}
