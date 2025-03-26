package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/kushagragupta/tms/models"
	"github.com/kushagragupta/tms/services"
)

type TaskController struct {
	service services.TaskService
}

func NewTaskController(repo services.TaskService) *TaskController {
	return &TaskController{
		service: repo,
	}
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := c.service.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.service.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var updateTask map[string]interface{}
	err = ctx.BindJSON(&updateTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	if title, ok := updateTask["title"]; ok {
		task.Title = title.(string)
	}
	if completed, ok := updateTask["completed"]; ok {
		task.Completed = completed.(bool)
	}

	err = c.service.UpdateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func (c *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.service.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}