package main

import (
	"github.com/A-ndrey/tododo/internal/list"
	"github.com/A-ndrey/tododo/internal/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

	listApi := api.Group("/list")

	listApi.GET("/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		i, err := listService.GetItem(id)
		if err != nil {
			context.Status(http.StatusInternalServerError) // todo error type
			return
		}

		context.JSON(http.StatusOK, i)
	})

	listApi.POST("/", func(context *gin.Context) {
		var i list.Item

		if err := context.BindJSON(&i); err != nil {
			log.Println(err)
			return
		}

		err := listService.AddNewItem(i)
		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}

		context.Status(http.StatusCreated)
	})

	listApi.PUT("/", func(context *gin.Context) {
		var i list.Item

		if err := context.BindJSON(&i); err != nil {
			log.Println(err)
			return
		}

		err := listService.UpdateItem(i)
		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}

		context.Status(http.StatusOK)
	})

	listApi.DELETE("/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		err = listService.DeleteItem(id)
		if err != nil {
			context.Status(http.StatusInternalServerError) // todo error type
			return
		}

		context.Status(http.StatusOK)
	})

	r.Run(":3000")
}
