package service

import (
	"log"

	"github.com/tweedledo/core/domain"
	"github.com/tweedledo/core/ports"
)

type TaskListService struct {
	taskListRepository ports.TaskListRepositoryInterface
}

func NewTaskListService(taskListRepository ports.TaskListRepositoryInterface) *TaskListService {
	return &TaskListService{
		taskListRepository: taskListRepository,
	}
}

func (tl *TaskListService) GetTaskListById(tasklistId string) (*domain.TaskList, error) {
	tasklist, err := tl.taskListRepository.GetTaskListById(tasklistId)
	if err != nil {
		log.Printf("P=Service M=GetTaskListById tasklistId=%v error=%v", tasklistId, err)
		return nil, err
	}

	return tasklist, nil
}

func (tl *TaskListService) GetAllTaskList() ([]*domain.TaskList, error) {
	tasklist, err := tl.taskListRepository.GetAllTasksLists()
	if err != nil {
		log.Printf("P=Service M=GetAllTaskList error=%v", err)
		return nil, err
	}

	return tasklist, nil
}

func (tl *TaskListService) CreateTaskList(name string) (*domain.TaskList, error) {
	tasks := make([]*domain.Task, 0)
	tasklist, err := domain.NewTaskList(name, tasks)
	if err != nil {
		log.Printf("P=Service M=GetTaskListById tasklistName=%v error=%v", name, err.Error())
		return tasklist, err
	}

	tasklist, err = tl.taskListRepository.CreateTaskList(tasklist)
	if err != nil {
		log.Printf("P=Service M=GetTaskListById step=repository tasklistName=%v error=%v", name, err.Error())
		return tasklist, err
	}

	return tasklist, nil
}
