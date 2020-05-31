package main

import (
	"github.com/A-ndrey/tododo/cmd/service-api/handler"
	"github.com/A-ndrey/tododo/internal/list"
	"github.com/A-ndrey/tododo/internal/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db, err := postgres.Connect()
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("can't ping database: %v", err)
	}

	listRepo := postgres.NewListRepository(db)
	listService := list.NewService(listRepo)

	r := gin.Default()

	api := r.Group("/api/v1")

	listHandler := &handler.Handler{ListService: listService}
	handler.RouteList(api, listHandler)

	r.Run(":3000")
}
