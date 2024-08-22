package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tianmarillio/technical-test-sagala/src/dtos"
	"github.com/tianmarillio/technical-test-sagala/src/services"
)

type TaskController struct {
	service *services.TaskService
}

func NewTaskController(s *services.TaskService) *TaskController {
	return &TaskController{service: s}
}

// TODO: error handlers
// TODO: docker
// TODO: e2e testng
// FIXME: optional when create: description, due date, status
// FIXME: use uint for create, update, delete json?
// FIXME: status validation on app level?
// TODO:  utils time parser, use UTC

// TODO: README
// TODO: postman docs
// TODO: docs format due date, sort

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var createTaskDto dtos.CreateTaskDTO

	if err := ctx.ShouldBindJSON(&createTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := c.service.CreateTask(createTaskDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"task_id": task.ID,
	})
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	sort := ctx.Query("sort")

	queryParams := dtos.TaskQueryParams{
		Sort: sort,
	}

	tasks, err := c.service.GetTasks(queryParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) GetTask(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	task, err := c.service.GetTask(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var updateTaskDto dtos.UpdateTaskDTO

	if err := ctx.ShouldBindJSON(&updateTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.UpdateTask(uint(id), updateTaskDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task_id": id,
	})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.service.DeleteTask(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task_id": id,
	})
}
