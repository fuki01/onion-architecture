package request_test

import (
	"testing"

	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
	"github.com/fuki01/onion-architecture/presentation/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateTaskRequest(t *testing.T) {
	t.Run("Valid request", func(t *testing.T) {
		req := request.CreateTaskRequest{
			Name:    "Task 1",
			UserId:  user.UserId(1),
			DueDate: "2023-05-31",
		}
		assert.Equal(t, "Task 1", req.Name)
		assert.Equal(t, user.UserId(1), req.UserId)
		assert.Equal(t, "2023-05-31", req.DueDate)
	})

	t.Run("Missing required fields", func(t *testing.T) {
		req := request.CreateTaskRequest{}
		assert.Empty(t, req.Name)
		assert.Equal(t, user.UserId(0), req.UserId)
		assert.Empty(t, req.DueDate)
	})
}

func TestExtendDueDateRequest(t *testing.T) {
	t.Run("Valid request", func(t *testing.T) {
		req := request.ExtendDueDateRequest{
			ID:      task.TaskId(1),
			DueDate: "2023-05-31",
		}
		assert.Equal(t, task.TaskId(1), req.ID)
		assert.Equal(t, "2023-05-31", req.DueDate)
	})

	t.Run("Missing required fields", func(t *testing.T) {
		req := request.ExtendDueDateRequest{}
		assert.Equal(t, task.TaskId(0), req.ID)
		assert.Empty(t, req.DueDate)
	})
}

func TestChangeStatusRequest(t *testing.T) {
	t.Run("Valid request", func(t *testing.T) {
		req := request.ChangeStatusRequest{
			ID:        task.TaskId(1),
			NewStatus: task.StatusComplete,
		}
		assert.Equal(t, task.TaskId(1), req.ID)
		assert.Equal(t, task.StatusComplete, req.NewStatus)
	})

	t.Run("Missing required fields", func(t *testing.T) {
		req := request.ChangeStatusRequest{}
		assert.Equal(t, task.TaskId(0), req.ID)
		assert.Equal(t, task.TaskStatus(""), req.NewStatus)
	})
}
