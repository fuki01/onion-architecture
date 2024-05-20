package usecase

import (
	"github.com/fuki01/onion-architecture/domain/repository"
	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
)

type Taskusecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) *Taskusecase {
	return &Taskusecase{
		taskRepository: taskRepository,
	}
}

// タスクを登録する
func (tu *Taskusecase) CreateTask(name string, userId user.UserId, dueDate string, delayCount int) (task.TaskId, error) {
	task := task.NewTask(name, userId, dueDate, delayCount)

	if err := task.Validate(); err != nil {
		return 0, err
	}

	task_id, err := tu.taskRepository.Insert(task)
	if err != nil {
		return 0, err
	}

	task.Id = task_id

	return task.Id, nil
}

// タスクの期限を延長する
func (tu *Taskusecase) ExtendDueDate(
	id task.TaskId,
	dueDate string,
) error {
	task, err := tu.taskRepository.FindById(id)
	if err != nil {
		return err
	}
	task.DueDate = dueDate
	task.DelayCount += 1
	return tu.taskRepository.Update(task)
}

// タスクのステータスを変更する
func (tu *Taskusecase) ChangeStatus(
	id task.TaskId,
	newStatus task.TaskStatus,
) error {
	task, err := tu.taskRepository.FindById(id)
	if err != nil {
		return err
	}
	if err := task.SetStatus(newStatus); err != nil {
		return err
	}
	return tu.taskRepository.Update(task)
}
