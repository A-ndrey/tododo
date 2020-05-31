package handler

import (
	"github.com/A-ndrey/tododo/internal/list"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	ListService list.Service
}

func RouteList(apiGroup *gin.RouterGroup, handler *Handler) {
	listApi := apiGroup.Group("/list")

	listApi.GET("/", handler.GetList)

	listApi.GET("/:id", handler.GetItem)

	listApi.POST("/", handler.CreateItem)

	listApi.PUT("/", handler.UpdateItem)

	listApi.DELETE("/:id", handler.DeleteItem)
}

func (h *Handler) GetList(context *gin.Context) {
	isCompleted, err := strconv.ParseBool(context.DefaultQuery("completed", "false"))
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	actualList, err := h.ListService.GetList(isCompleted)
	if err != nil {
		context.Status(http.StatusInternalServerError) // todo error type
		return
	}

	context.JSON(http.StatusOK, actualList)
}

func (h *Handler) GetItem(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	i, err := h.ListService.GetItem(id)
	if err != nil {
		context.Status(http.StatusInternalServerError) // todo error type
		return
	}

	context.JSON(http.StatusOK, i)
}

func (h *Handler) CreateItem(context *gin.Context) {
	var i list.Item

	if err := context.BindJSON(&i); err != nil {
		log.Println(err)
		return
	}

	err := h.ListService.AddNewItem(i)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusCreated)
}

func (h *Handler) UpdateItem(context *gin.Context) {
	var i list.Item

	if err := context.BindJSON(&i); err != nil {
		log.Println(err)
		return
	}

	err := h.ListService.UpdateItem(i)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}

func (h *Handler) DeleteItem(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err = h.ListService.DeleteItem(id)
	if err != nil {
		context.Status(http.StatusInternalServerError) // todo error type
		return
	}

	context.Status(http.StatusOK)
}
