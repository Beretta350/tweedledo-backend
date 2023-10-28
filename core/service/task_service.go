package service

import (
	"log"

	"github.com/tweedledo/core/domain"
	"github.com/tweedledo/core/ports"
)

type TaskService struct {
	taskRepository  ports.TaskRepositoryInterface
	taskListService ports.TaskListServiceInterface
}

func NewTaskService(
	taskRepository ports.TaskRepositoryInterface,
	taskListService ports.TaskListServiceInterface,
) *TaskService {
	return &TaskService{
		taskRepository:  taskRepository,
		taskListService: taskListService,
	}
}

func (s *TaskService) CreateTask(name string, desc string, tasklistId string) (*domain.Task, error) {
	tasklist, err := s.taskListService.GetTaskListById(tasklistId)

	task, err := domain.NewTask(name, desc, tasklist)
	if err != nil {
		log.Fatalf("P=Service M=CreateTask step=domain name=%v error=%v", name, err.Error())
		return task, err
	}

	s.taskRepository.CreateTask(task)
	if err != nil {
		log.Fatalf("P=Service M=CreateTask step=repository name=%v error=%v", name, err.Error())
		return task, err
	}

	return task, err
}
