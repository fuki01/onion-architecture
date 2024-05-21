package request

import (
	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
)

// request.go
type CreateTaskRequest struct {
	Name    string      `json:"name" binding:"required"`
	UserId  user.UserId `json:"user_id" binding:"required"`
	DueDate string      `json:"due_date" binding:"required"`
}

type ExtendDueDateRequest struct {
	ID      task.TaskId `json:"id" binding:"required"`
	DueDate string      `json:"due_date" binding:"required"`
}

type ChangeStatusRequest struct {
	ID        task.TaskId     `json:"id" binding:"required"`
	NewStatus task.TaskStatus `json:"new_status" binding:"required"`
}
