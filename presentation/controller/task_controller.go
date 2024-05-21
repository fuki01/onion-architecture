package controller

import (
	"net/http"

	"github.com/fuki01/onion-architecture/presentation/request"
	"github.com/fuki01/onion-architecture/usecase"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskusecase usecase.TaskUsecase
}

func NewTaskController(taskusecase usecase.TaskUsecase) *TaskController {
	return &TaskController{
		taskusecase: taskusecase,
	}
}

// タスクを登録する
func (tc *TaskController) CreateTask(c *gin.Context) {
	var input request.CreateTaskRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskID, err := tc.taskusecase.CreateTask(input.Name, input.UserId, input.DueDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task_id": taskID})
}

// タスクの期限を延長する
func (tc *TaskController) ExtendDueDate(c *gin.Context) {
	var input request.ExtendDueDateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.taskusecase.ExtendDueDate(input.ID, input.DueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// タスクのステータスを変更する
func (tc *TaskController) ChangeStatus(c *gin.Context) {
	var input request.ChangeStatusRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.taskusecase.ChangeStatus(input.ID, input.NewStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
