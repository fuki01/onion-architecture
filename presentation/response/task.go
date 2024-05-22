package response

import "github.com/fuki01/onion-architecture/domain/task"

type CreateTaskResponse struct {
	TaskID task.TaskId `json:"task_id"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}


type GetTaskResponse struct {
	TaskID   task.TaskId `json:"task_id"`
	Name     string      `json:"name"`
	DueDate  string      `json:"due_date"`
	Status   string      `json:"status"`
	DelayCnt int         `json:"delay_cnt"`
}
