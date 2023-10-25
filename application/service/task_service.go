package service

import "github.com/tweedledo/application/domain/model"

type TaskService struct {
	TaskRepository model.TaskRepositoryInterface
}

func (s *TaskService) CreateTask(name string, desc string) {

}
