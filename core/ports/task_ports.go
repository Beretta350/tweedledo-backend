package ports

import "github.com/tweedledo/core/domain"

// Task's inbound ports
type TaskServiceInterface interface {
	GetTaskById(id string) (*domain.Task, error)
	CreateTask(name string, desc string, tasklistId string) (*domain.Task, error)
	UpdateTask(id, name, desc string) (*domain.Task, error)
	DeleteTaskById(id string) (string, error)
}

// Task's outbound ports
type TaskRepositoryInterface interface {
	CreateTask(task *domain.Task) (*domain.Task, error)
	UpdateTask(task *domain.Task) (*domain.Task, error)
	DeleteTaskById(id string) (int64, error)
	GetTaskById(id string) (*domain.Task, error)
	GetTasksInTaskList(tasklistId string) ([]*domain.Task, error)
}
