package task_test

import (
	"testing"

	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	task := task.NewTask("test", user.UserId(1), "2024-01-01")

	assert.Equal(t, "test", task.Name)
	assert.Equal(t, user.UserId(1), task.UserId)
	assert.Equal(t, "2024-01-01", task.DueDate)
	assert.Equal(t, 0, task.DelayCount)
}

func TestValidate(t *testing.T) {
	testCases := []struct {
			name        string
			task        *task.Task
			expectedErr string
	}{
			{
					name:        "Valid task",
					task:        task.NewTask("task1", user.UserId(1), "2024-01-01"),
					expectedErr: "",
			},
			{
					name:        "Invalid task name",
					task:        task.NewTask("", user.UserId(1), "2024-01-01"),
					expectedErr: "invalid task name",
			},
			{
					name:        "Invalid user ID",
					task:        task.NewTask("task1", user.UserId(0), "2024-01-01"),
					expectedErr: "invalid user id",
			},
			{
					name:        "Invalid due date",
					task:        task.NewTask("task1", user.UserId(1), ""),
					expectedErr: "invalid due date",
			},
	}

	for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
					err := tc.task.Validate()
					if tc.expectedErr == "" {
							assert.NoError(t, err)
					} else {
							assert.EqualError(t, err, tc.expectedErr)
					}
			})
	}
}


func TestSetStatus(t *testing.T) {
	task := task.NewTask("test", user.UserId(1), "2024-01-01")
	assert.NoError(t, task.SetStatus("完了"))
	assert.EqualError(t, task.SetStatus("未完了"), "cannot revert to incomplete")
	assert.EqualError(t, task.SetStatus("完了"), "already completed")
}
