package service

import (
	"fmt"
	"log"
	"time"

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

func (s *TaskListService) GetTaskListById(tasklistId string) (*domain.TaskList, error) {
	tasklist, err := s.taskListRepository.GetTaskListById(tasklistId)
	if err != nil {
		log.Printf("P=Service M=GetTaskListById tasklistId=%v error=%v", tasklistId, err)
		return nil, err
	}

	return tasklist, nil
}

func (s *TaskListService) GetAllTaskList() ([]*domain.TaskList, error) {
	tasklist, err := s.taskListRepository.GetAllTaskLists()
	if err != nil {
		log.Printf("P=Service M=GetAllTaskList error=%v", err)
		return nil, err
	}

	return tasklist, nil
}

func (s *TaskListService) CreateTaskList(name string) (*domain.TaskList, error) {
	tasklist, err := domain.NewTaskList(name, []*domain.Task{})
	if err != nil {
		log.Printf("P=Service M=GetTaskListById tasklistName=%v error=%v", name, err.Error())
		return tasklist, err
	}

	tasklist, err = s.taskListRepository.CreateTaskList(tasklist)
	if err != nil {
		log.Printf("P=Service M=GetTaskListById step=repository tasklistName=%v error=%v", name, err.Error())
		return tasklist, err
	}

	return tasklist, nil
}

func (s *TaskListService) UpdateTaskList(id string, name string) (*domain.TaskList, error) {
	tasklist, err := s.taskListRepository.GetTaskListById(id)
	if err != nil {
		log.Printf("P=Service M=UpdateTaskList step=GetTaskListById id=%v error=%v", id, err.Error())
		return tasklist, err
	}

	tasklist.Name = name
	tasklist.UpdatedAt = time.Now()

	tasklist, err = s.taskListRepository.UpdateTaskList(tasklist)
	if err != nil {
		log.Printf("P=Service M=UpdateTaskList id=%v error=%v", id, err.Error())
		return tasklist, err
	}

	return tasklist, err
}

func (s *TaskListService) DeleteTaskListById(id string) (string, error) {
	rowsAffected, err := s.taskListRepository.DeleteTaskListById(id)
	if err != nil {
		log.Printf("P=Service M=DeleteTaskById id=%v error=%v", id, err.Error())
		return "fail", err
	}
	response := fmt.Sprintf("success with %v rows affected", rowsAffected)
	return response, err
}
