package main

import (
	"context"
	"fmt"
	"github.com/A-ndrey/tododo/cmd/service-api/handler"
	"github.com/A-ndrey/tododo/internal/config"
	_ "github.com/A-ndrey/tododo/internal/config"
	"github.com/A-ndrey/tododo/internal/postgres"
	"github.com/A-ndrey/tododo/internal/task"
	"github.com/A-ndrey/tododo/internal/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	initLogger()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		zap.S().Fatal("environment variable JWT_SECRET not set")
	}

	db := initDB()

	listRepo := task.NewRepository(db)
	listService := task.NewService(listRepo)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)

	r := gin.Default()

	api := r.Group("/api/v1")

	listHandler := &handler.ListHandler{ListService: listService}

	authHandler := &handler.AuthHandler{
		JWTSecret:   []byte(jwtSecret),
		UserService: userService,
	}

	handler.RouteAuth(api, authHandler)
	handler.RouteList(api, listHandler, authHandler.AuthMiddleware)

	serverConf := config.GetServer()
	addr := fmt.Sprintf("%s:%d", serverConf.Host, serverConf.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server forced to shutdown:", err)
	}
}

func initLogger() {
	var createLogger func(options ...zap.Option) (*zap.Logger, error)

	switch config.GetEnvironment() {
	case config.ENV_PROD:
		createLogger = zap.NewProduction
	case config.ENV_DEV:
		createLogger = zap.NewDevelopment
	default:
		log.Fatal("unknown environment")
	}

	logger, err := createLogger()
	if err != nil {
		log.Fatal("can't initialize logger")
	}

	zap.ReplaceGlobals(logger)
}

func initDB() *gorm.DB {
	zap.S().Info("connecting to database...")

	db, err := postgres.Connect()
	if err != nil {
		zap.S().Fatalf("can't connect to database: %v", err)
	}

	zap.S().Info("successful connection to the database")

	return db
}
