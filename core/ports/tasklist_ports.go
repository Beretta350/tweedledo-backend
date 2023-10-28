package ports

import "github.com/tweedledo/core/domain"

// TaskList's inbound ports
type TaskListServiceInterface interface {
	GetTaskListById(tasklistId string) (*domain.TaskList, error)
	CreateTaskList(name string) (*domain.TaskList, error)
}

// TaskList's outbound ports
type TaskListRepositoryInterface interface {
	CreateTaskList(t *domain.TaskList) (*domain.TaskList, error)
	UpdateTaskList(t *domain.TaskList) (*domain.TaskList, error)
	DeleteTaskList(t *domain.TaskList) (*domain.TaskList, error)
	GetTaskListById(id string) (*domain.TaskList, error)
	GetAllTasksLists() ([]*domain.TaskList, error)
}
