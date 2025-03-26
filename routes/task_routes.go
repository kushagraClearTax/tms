package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kushagragupta/tms/controllers"
)

func RegisterTaskRoutes(r *gin.Engine, controller *controllers.TaskController) {
	taskRoutes := r.Group("/tasks")
	{
	    taskRoutes.GET("/:id", controller.GetTask)
		taskRoutes.GET("/", controller.GetTasks)
		taskRoutes.POST("/", controller.CreateTask)
		taskRoutes.PUT("/:id", controller.UpdateTask)
		taskRoutes.DELETE("/:id", controller.DeleteTask)

	}
}