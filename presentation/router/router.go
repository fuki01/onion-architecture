package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fuki01/onion-architecture/presentation/controller"
)

func SetupRouter(taskController *controller.TaskController) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskController.CreateTask)
			tasks.GET("/:id", taskController.GetTask)
			tasks.PUT("/:id/extend", taskController.ExtendDueDate)
			tasks.PUT("/:id/status", taskController.ChangeStatus)
		}
	}

	return router
}
