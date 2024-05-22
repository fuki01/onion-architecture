package repository

import (
	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
)

type TaskRepository interface {
	FindById(id task.TaskId) (*task.Task, error)
	FindByUserId(userId user.UserId) ([]*task.Task, error)
	Insert(task *task.Task) (task.TaskId, error)
	Update(task *task.Task) error
	Delete(task *task.Task) error
}
