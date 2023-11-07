package ports

import "github.com/tweedledo/core/domain"

// TaskList's inbound ports
type TaskListServiceInterface interface {
	GetTaskListById(tasklistId string) (*domain.TaskList, error)
	GetAllTaskList() ([]*domain.TaskList, error)
	CreateTaskList(name string) (*domain.TaskList, error)
	UpdateTaskList(id string, name string) (*domain.TaskList, error)
	DeleteTaskListById(id string) (string, error)
}

// TaskList's outbound ports
type TaskListRepositoryInterface interface {
	CreateTaskList(t *domain.TaskList) (*domain.TaskList, error)
	UpdateTaskList(t *domain.TaskList) (*domain.TaskList, error)
	DeleteTaskListById(id string) (int64, error)
	GetTaskListById(id string) (*domain.TaskList, error)
	GetAllTaskLists() ([]*domain.TaskList, error)
}
