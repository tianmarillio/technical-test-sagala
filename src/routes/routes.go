package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tianmarillio/technical-test-sagala/src/controllers"
	"github.com/tianmarillio/technical-test-sagala/src/repositories"
	"github.com/tianmarillio/technical-test-sagala/src/services"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Root
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Tasks
	taskRepository := repositories.NewGormTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	r.POST("/tasks", taskController.CreateTask)
	r.GET("/tasks", taskController.GetTasks)
	r.GET("/tasks/:id", taskController.GetTask)
	r.PATCH("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)
}
