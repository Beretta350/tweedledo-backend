package service

import (
	"log"

	"github.com/tweedledo/core/domain"
	"github.com/tweedledo/core/ports"
)

type TaskListService struct {
	TaskListRepository ports.TaskListRepositoryInterface
}

func (tl *TaskListService) GetTaskListById(tasklistId string) (*domain.TaskList, error) {
	tasklist, err := tl.TaskListRepository.GetTaskListById(tasklistId)
	if err != nil {
		log.Fatalf("P=Service M=GetTaskListById tasklistId=%v error=%v", tasklistId, err.Error())
		return tasklist, err
	}

	return tasklist, nil
}

func (tl *TaskListService) CreateTaskList(name string) (*domain.TaskList, error) {
	tasks := make([]*domain.Task, 0)
	tasklist, err := domain.NewTaskList(name, tasks)
	if err != nil {
		log.Fatalf("P=Service M=GetTaskListById tasklistName=%v error=%v", name, err.Error())
		return tasklist, err
	}

	tasklist, err = tl.TaskListRepository.CreateTaskList(tasklist)
	if err != nil {
		log.Fatalf("P=Service M=GetTaskListById step=repository tasklistName=%v error=%v", name, err.Error())
		return tasklist, err
	}

	return tasklist, nil
}
