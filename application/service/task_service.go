package service

import (
	"log"

	"github.com/tweedledo/application/domain/model"
)

type TaskService struct {
	TaskRepository  model.TaskRepositoryInterface
	TaskListService TaskListService
}

func (s *TaskService) CreateTask(name string, desc string, tasklistId string) (*model.Task, error) {
	tasklist, err := s.TaskListService.GetTaskListById(tasklistId)

	task, err := model.NewTask(name, desc, tasklist)
	if err != nil {
		log.Fatalf("P=Service M=CreateTask step=model name=%v error=%v", name, err.Error())
		return task, err
	}

	s.TaskRepository.CreateTask(task)
	if err != nil {
		log.Fatalf("P=Service M=CreateTask step=repository name=%v error=%v", name, err.Error())
		return task, err
	}

	return task, err
}
