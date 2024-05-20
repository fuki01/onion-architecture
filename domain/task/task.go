package task

import (
	"errors"

	"github.com/fuki01/onion-architecture/domain/user"
)

type TaskStatus string

const (
	StatusIncomplete TaskStatus = "未完了"
	StatusComplete   TaskStatus = "完了"
)

type Task struct {
	Id         TaskId      `json:"id"`
	Name       string      `json:"name"`
	UserId     user.UserId `json:"user_id"`
	Status     TaskStatus  `json:"status"`
	DueDate    string      `json:"due_date"`
	DelayCount int         `json:"delay_count"`
}

func NewTask(name string, userId user.UserId, dueDate string, delayCount int) *Task {
	return &Task{
		Name:       name,
		UserId:     userId,
		Status:     "未完了",
		DueDate:    dueDate,
		DelayCount: delayCount,
	}
}

func (t *Task) Validate() error {
	if t.Name == "" {
		return errors.New("invalid task name")
	}

	if t.UserId == 0 {
		return errors.New("invalid user id")
	}

	if t.DueDate == "" {
		return errors.New("invalid due date")
	}

	if t.DelayCount < 0 {
		return errors.New("invalid delay count")
	}

	return nil
}

func (t *Task) SetStatus(newStatus TaskStatus) error {
	if t.Status == StatusComplete && newStatus == StatusComplete {
		return errors.New("already completed")
	} else if t.Status == StatusComplete && newStatus == StatusIncomplete {
		return errors.New("cannot revert to incomplete")
	} else {
		t.Status = newStatus
	}
	return nil
}


