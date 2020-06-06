package main

import (
	"database/sql"
	"fmt"
	"github.com/A-ndrey/tododo/cmd/service-api/handler"
	"github.com/A-ndrey/tododo/internal/config"
	_ "github.com/A-ndrey/tododo/internal/config"
	"github.com/A-ndrey/tododo/internal/list"
	"github.com/A-ndrey/tododo/internal/postgres"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	initLogger()

	db := initDB()
	defer func() {
		err := db.Close()
		zap.S().Fatalf("can't close database connection: %v", err)
	}()

	listRepo := postgres.NewListRepository(db)
	listService := list.NewService(listRepo)

	r := gin.Default()

	api := r.Group("/api/v1")

	listHandler := &handler.ListHandler{ListService: listService}
	handler.RouteList(api, listHandler)

	serverConf := config.GetServer()
	addr := fmt.Sprintf("%s:%d", serverConf.Host, serverConf.Port)
	http.ListenAndServe(addr, r)
}

func initLogger() {
	var createLogger func(options ...zap.Option) (*zap.Logger, error)

	switch config.GetEnvironment() {
	case "prod":
		createLogger = zap.NewProduction
	case "dev":
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

func initDB() *sql.DB {
	zap.S().Info("connecting to database...")

	db, err := postgres.Connect()
	if err != nil {
		zap.S().Fatalf("can't connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		zap.S().Fatalf("can't ping database: %v", err)
	}

	zap.S().Info("successful connection to the database")

	return db
}
