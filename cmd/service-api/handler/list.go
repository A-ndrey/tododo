package handler

import (
	"github.com/A-ndrey/tododo/internal/list"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ListHandler struct {
	ListService list.Service
}

func RouteList(apiGroup *gin.RouterGroup, handler *ListHandler) {
	listApi := apiGroup.Group("/list")

	listApi.GET("/", handler.GetList)

	listApi.GET("/:id", handler.GetItem)

	listApi.POST("/", handler.CreateItem)

	listApi.PUT("/", handler.UpdateItem)

	listApi.DELETE("/:id", handler.DeleteItem)
}

func (h *ListHandler) GetList(ctx *gin.Context) {
	_, isCompleted := ctx.GetQuery("completed")

	actualList, err := h.ListService.GetList(isCompleted)
	if err != nil {
		zap.L().Error("GetList",
			zap.Bool("completed", isCompleted),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError) // todo error type
		return
	}

	ctx.JSON(http.StatusOK, actualList)
}

func (h *ListHandler) GetItem(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		zap.L().Warn("GetItem",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	i, err := h.ListService.GetItem(id)
	if err != nil {
		zap.L().Error("GetItem",
			zap.Uint64("id", id),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError) // todo error type
		return
	}

	ctx.JSON(http.StatusOK, i)
}

func (h *ListHandler) CreateItem(ctx *gin.Context) {
	var i list.Item

	if err := ctx.BindJSON(&i); err != nil {
		zap.L().Warn("CreateItem",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := h.ListService.AddNewItem(i)
	if err != nil {
		zap.L().Error("CreateItem",
			zap.Reflect("item", i),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *ListHandler) UpdateItem(ctx *gin.Context) {
	var i list.Item

	if err := ctx.BindJSON(&i); err != nil {
		zap.L().Warn("UpdateItem",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := h.ListService.UpdateItem(i)
	if err != nil {
		zap.L().Error("UpdateItem",
			zap.Reflect("item", i),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *ListHandler) DeleteItem(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		zap.L().Warn("DeleteItem",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = h.ListService.DeleteItem(id)
	if err != nil {
		zap.L().Error("DeleteItem",
			zap.Uint64("id", id),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError) // todo error type
		return
	}

	ctx.Status(http.StatusOK)
}
