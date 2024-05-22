package infrastructure

// task_repositoryの実装

import (
	"gorm.io/gorm"

	"github.com/fuki01/onion-architecture/domain/repository"
	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
)

type taskPersistence struct {
	db *gorm.DB
}

func NewArticlePersistence(db *gorm.DB) repository.TaskRepository {
	return &taskPersistence{
		db: db,
	}
}

// FindById は指定したIDのタスクを取得する
func (tr *taskPersistence) FindById(id task.TaskId) (*task.Task, error) {
	var t task.Task
	if err := tr.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

// FindByUserId は指定したユーザーIDのタスクを取得する
func (tr *taskPersistence) FindByUserId(userId user.UserId) ([]*task.Task, error) {
	var tasks []*task.Task
	if err := tr.db.Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Insert はタスクを登録する
func (tr *taskPersistence) Insert(t *task.Task) (task.TaskId, error) {
	if err := tr.db.Create(t).Error; err != nil {
		return 0, err
	}
	return t.Id, nil
}

// Update はタスクを更新する
func (tr *taskPersistence) Update(t *task.Task) error {
	return tr.db.Save(t).Error
}

// Delete はタスクを削除する
func (tr *taskPersistence) Delete(t *task.Task) error {
	return tr.db.Delete(t).Error
}
