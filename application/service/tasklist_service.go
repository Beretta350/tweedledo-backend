package service

import (
	"log"

	"github.com/tweedledo/application/domain/model"
)

type TaskListService struct {
	TaskListRepository model.TaskListRepositoryInterface
}

func (tl *TaskListService) GetTaskListById(tasklistId string) (*model.TaskList, error) {
	tasklist, err := tl.TaskListRepository.GetTaskListById(tasklistId)
	if err != nil {
		log.Fatalf("P=Service M=GetTaskListById tasklistId=%v error=%v", tasklistId, err.Error())
		return tasklist, err
	}

	return tasklist, nil
}

func (tl *TaskListService) CreateTaskList(name string) (*model.TaskList, error) {
	tasks := make([]*model.Task, 0)
	tasklist, err := model.NewTaskList(name, tasks)
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
