package handler

import (
	"github.com/A-ndrey/tododo/internal/task"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ListHandler struct {
	ListService task.Service
}

func RouteTasks(apiGroup *gin.RouterGroup, handler *ListHandler) {
	listApi := apiGroup.Group("/tasks")

	listApi.GET("/", handler.GetTasks)

	listApi.GET("/:id", handler.GetTask)

	listApi.POST("/", handler.CreateTask)

	listApi.PUT("/", handler.UpdateTask)

	listApi.DELETE("/:id", handler.DeleteTask)
}

func (h *ListHandler) GetTasks(ctx *gin.Context) {
	_, isCompleted := ctx.GetQuery("completed")

	userId, ok := getUserId(ctx)
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	actualList, err := h.ListService.GetList(userId, isCompleted)
	if err != nil {
		zap.L().Error("GetTasks",
			zap.Bool("completed", isCompleted),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError) // todo error type
		return
	}

	ctx.JSON(http.StatusOK, actualList)
}

func (h *ListHandler) GetTask(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		zap.L().Warn("GetTask",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	userId, ok := getUserId(ctx)
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	i, err := h.ListService.GetTask(userId, id)
	if err != nil {
		zap.L().Error("GetTask",
			zap.Uint64("id", id),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError) // todo error type
		return
	}

	ctx.JSON(http.StatusOK, i)
}

func (h *ListHandler) CreateTask(ctx *gin.Context) {
	var t task.Task

	if err := ctx.BindJSON(&t); err != nil {
		zap.L().Warn("CreateTask",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	userId, ok := getUserId(ctx)
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	t.UserID = userId

	t, err := h.ListService.AddNewTask(t)
	if err != nil {
		zap.L().Error("CreateTask",
			zap.Reflect("task", t),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, t)
}

func (h *ListHandler) UpdateTask(ctx *gin.Context) {
	var t task.Task

	if err := ctx.BindJSON(&t); err != nil {
		zap.L().Warn("UpdateTask",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	userId, ok := getUserId(ctx)
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	t.UserID = userId

	err := h.ListService.UpdateTask(t)
	if err != nil {
		zap.L().Error("UpdateTask",
			zap.Reflect("task", t),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *ListHandler) DeleteTask(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		zap.L().Warn("DeleteTask",
			zap.Error(err),
		)
		ctx.Status(http.StatusBadRequest)
		return
	}

	userId, ok := getUserId(ctx)
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = h.ListService.DeleteTask(userId, id)
	if err != nil {
		zap.L().Error("DeleteTask",
			zap.Uint64("id", id),
			zap.Error(err),
		)
		ctx.Status(http.StatusInternalServerError) // todo error type
		return
	}

	ctx.Status(http.StatusOK)
}

func getUserId(ctx *gin.Context) (uint64, bool) {
	userIdRaw, ok := ctx.Get(UserIDKey)
	if !ok {
		zap.L().Error("User ID not specified",
			zap.String("path", ctx.FullPath()),
		)
		return 0, false
	}

	return userIdRaw.(uint64), true
}
