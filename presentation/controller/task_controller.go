package controller

import (
	"net/http"

	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
	"github.com/fuki01/onion-architecture/usecase"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskusecase *usecase.Taskusecase
}

func NewTaskController(taskusecase *usecase.Taskusecase) *TaskController {
	return &TaskController{
		taskusecase: taskusecase,
	}
}


// タスクを登録する
func (tc *TaskController) CreateTask(c *gin.Context) {
	var input struct {
		Name       string      `json:"name" binding:"required"`
		UserID     user.UserId `json:"user_id" binding:"required"`
		DueDate    string      `json:"due_date" binding:"required"`
		DelayCount int         `json:"delay_count" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskID, err := tc.taskusecase.CreateTask(input.Name, input.UserID, input.DueDate, input.DelayCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task_id": taskID})
}

// タスクの期限を延長する
func (tc *TaskController) ExtendDueDate(c *gin.Context) {
	var input struct {
		ID      task.TaskId `json:"id"`
		DueDate string      `json:"due_date"`
	}

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
	var input struct {
		ID        task.TaskId     `json:"id"`
		NewStatus task.TaskStatus `json:"new_status"`
	}

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
