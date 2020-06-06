package main

import (
	"database/sql"
	"github.com/A-ndrey/tododo/cmd/service-api/handler"
	"github.com/A-ndrey/tododo/internal/list"
	"github.com/A-ndrey/tododo/internal/postgres"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
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

	r.Run(":3000")
}

func initLogger() {
	logger, err := zap.NewDevelopment()
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
