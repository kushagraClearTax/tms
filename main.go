package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kushagragupta/tms/controllers"
	"github.com/kushagragupta/tms/database"
	"github.com/kushagragupta/tms/routes"
	"github.com/kushagragupta/tms/services"
)

func main() {
	repo := database.NewInMemoryTaskRepository()
	service := services.NewTaskServiceImpl(repo)
	controller := controllers.NewTaskController(service)

	r := gin.Default()
	routes.RegisterTaskRoutes(r, controller)

	// Start server
	port := "8080"
	r.Run(":" + port)
}