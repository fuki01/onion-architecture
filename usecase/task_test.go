package usecase_test

import (
	"errors"
	"testing"

	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
	"github.com/fuki01/onion-architecture/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Insert(task *task.Task) (task.TaskId, error) {
	args := m.Called(task)
	return 1, args.Error(1)
}

func (m *MockTaskRepository) FindById(id task.TaskId) (*task.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(task *task.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(task *task.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

// タスクを作成する
func TestCreateTask(t *testing.T) {
	createMock := func(returnId task.TaskId, returnErr error) *MockTaskRepository {
		mockRepo := new(MockTaskRepository)
		mockRepo.On("Insert", mock.AnythingOfType("*task.Task")).Return(returnId, returnErr)
		return mockRepo
	}

	createUsecase := func(mockRepo *MockTaskRepository) *usecase.Taskusecase {
		return usecase.NewTaskUsecase(mockRepo)
	}

	t.Run("create", func(t *testing.T) {
		mockRepo := createMock(task.TaskId(1), nil)
		usecase := createUsecase(mockRepo)

		taskId, err := usecase.CreateTask("test", user.UserId(1), "2024-01-01", 1)

		assert.NoError(t, err)
		assert.Equal(t, task.TaskId(1), taskId)
		mockRepo.AssertExpectations(t)
	})

	t.Run("validate", func(t *testing.T) {
		mockRepo := createMock(task.TaskId(1), nil)
		usecase := createUsecase(mockRepo)

		taskId, err := usecase.CreateTask("", user.UserId(1), "2024-01-01", 1)
		assert.Error(t, err)
		assert.Equal(t, task.TaskId(0), taskId)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := createMock(task.TaskId(0), errors.New("repository error"))
		usecase := createUsecase(mockRepo)

		taskId, err := usecase.CreateTask("test", user.UserId(1), "2024-01-01", 1)

		assert.Error(t, err)
		assert.Equal(t, task.TaskId(0), taskId)
		assert.Contains(t, err.Error(), "repository error")
		mockRepo.AssertExpectations(t)
	})
}

func TestExtendDueDate(t *testing.T) {
	createMock := func(task *task.Task, findErr, updateErr error) *MockTaskRepository {
		mockRepo := new(MockTaskRepository)
		mockRepo.On("FindById", task.Id).Return(task, findErr)
		mockRepo.On("Update", task).Return(updateErr)
		return mockRepo
	}

	createUsecase := func(mock *MockTaskRepository) *usecase.Taskusecase {
		return usecase.NewTaskUsecase(mock)
	}

	t.Run("success", func(t *testing.T) {
		// 初期値の設定
		existingTask := task.NewTask("test", user.UserId(1), "2024-01-01", 1)
		existingTask.Id = task.TaskId(1)

		// モック作成
		mockRepo := createMock(existingTask, nil, nil)
		usecase := createUsecase(mockRepo)

		// 締切の延長
		err := usecase.ExtendDueDate(task.TaskId(1), "2024-01-02")

		// 検証
		assert.NoError(t, err)
		assert.Equal(t, "2024-01-02", existingTask.DueDate)
		assert.Equal(t, 2, existingTask.DelayCount)
		mockRepo.AssertExpectations(t)
	})

	t.Run("find error", func(t *testing.T) {
		// 初期値の設定
		existingTask := task.NewTask("test", user.UserId(1), "2024-01-01", 1)
		existingTask.Id = task.TaskId(1)

		// モック作成
		mockRepo := createMock(existingTask, errors.New("find error"), nil)
		usecase := createUsecase(mockRepo)

		// 検証
		assert.Error(t, usecase.ExtendDueDate(task.TaskId(1), "2024-01-02"))
	})

	t.Run("update error", func(t *testing.T) {
		// 初期値の設定
		existingTask := task.NewTask("test", user.UserId(1), "2024-01-01", 1)
		existingTask.Id = task.TaskId(1)

		// モック作成
		mockRepo := createMock(existingTask, nil, errors.New("update error"))
		usecase := createUsecase(mockRepo)

		// 検証
		assert.Error(t, usecase.ExtendDueDate(task.TaskId(1), "2024-01-02"))
	})
}

func TestChangeStatus(t *testing.T) {
	createMock := func(task *task.Task, findErr, updateErr error) *MockTaskRepository {
		mockRepo := new(MockTaskRepository)
		mockRepo.On("FindById", task.Id).Return(task, findErr)
		mockRepo.On("Update", task).Return(updateErr)
		return mockRepo
	}

	createUsecase := func(mock *MockTaskRepository) *usecase.Taskusecase {
		return usecase.NewTaskUsecase(mock)
	}

	t.Run("success", func(t *testing.T) {
		// 初期値の設定
		existingTask := task.NewTask("test", user.UserId(1), "2024-01-01", 1)
		existingTask.Id = task.TaskId(1)

		// モック作成
		mockRepo := createMock(existingTask, nil, nil)
		usecase := createUsecase(mockRepo)

		// ステータスの変更
		err := usecase.ChangeStatus(task.TaskId(1), "完了")

		// 検証
		assert.NoError(t, err)
		assert.Equal(t, task.TaskStatus("完了"), existingTask.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("find error", func(t *testing.T) {
		// 初期値の設定
		existingTask := task.NewTask("test", user.UserId(1), "2024-01-01", 1)
		existingTask.Id = task.TaskId(1)

		// モック作成
		mockRepo := createMock(existingTask, errors.New("find error"), nil)
		usecase := createUsecase(mockRepo)

		// 検証
		assert.Error(t, usecase.ChangeStatus(task.TaskId(1), "完了"))
	})

	t.Run("update error", func(t *testing.T) {
		// 初期値の設定
		existingTask := task.NewTask("test", user.UserId(1), "2024-01-01", 1)
		existingTask.Id = task.TaskId(1)

		// モック作成
		mockRepo := createMock(existingTask, nil, errors.New("update error"))
		usecase := createUsecase(mockRepo)

		// 検証
		assert.Error(t, usecase.ChangeStatus(task.TaskId(1), "完了"))
	})

	t.Run("invalid status", func(t *testing.T) {
		// 初期値の設定
		existingTask := task.NewTask("test", user.UserId(1), "2024-01-01", 1)
		existingTask.Id = task.TaskId(1)

		// モック作成
		mockRepo := createMock(existingTask, nil, nil)
		usecase := createUsecase(mockRepo)
		usecase.ChangeStatus(task.TaskId(1), "完了")

		err := usecase.ChangeStatus(task.TaskId(1), "未完了")

		// 検証
		assert.Error(t, err)
	})
}
