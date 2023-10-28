package ports

import "github.com/tweedledo/core/domain"

// Task's inbound ports
type TaskServiceInterface interface {
	CreateTask(name string, desc string, tasklistId string) (*domain.Task, error)
}

// Task's outbound ports
type TaskRepositoryInterface interface {
	CreateTask(task *domain.Task) (*domain.Task, error)
	UpdateTask(task *domain.Task) (*domain.Task, error)
	DeleteTask(task *domain.Task) (*domain.Task, error)
	GetTasksInTaskList(tasklistId string) ([]*domain.Task, error)
}
