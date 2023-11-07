package service

import (
	"fmt"
	"log"
	"time"

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

func (s *TaskService) GetTaskById(id string) (*domain.Task, error) {
	task, err := s.taskRepository.GetTaskById(id)
	if err != nil {
		log.Printf("P=Service M=GetTaskById id=%v error=%v", id, err.Error())
		return task, err
	}
	return task, err
}

func (s *TaskService) CreateTask(name string, desc string, tasklistId string) (*domain.Task, error) {
	tasklist, err := s.taskListService.GetTaskListById(tasklistId)
	if err != nil {
		log.Printf("P=Service M=CreateTask step=GetTaskListById name=%v error=%v", name, err.Error())
		return nil, err
	}

	task, err := domain.NewTask(name, desc, tasklist)
	if err != nil {
		log.Printf("P=Service M=CreateTask step=NewTask name=%v error=%v", name, err.Error())
		return task, err
	}

	task, err = s.taskRepository.CreateTask(task)
	if err != nil {
		log.Printf("P=Service M=CreateTask name=%v error=%v", name, err.Error())
		return task, err
	}

	return task, err
}

func (s *TaskService) UpdateTask(id, name, desc string) (*domain.Task, error) {
	task, err := s.taskRepository.GetTaskById(id)
	if err != nil {
		log.Printf("P=Service M=UpdateTask step=GetTaskById id=%v error=%v", id, err.Error())
		return task, err
	}

	task.Name = name
	task.Description = desc
	task.UpdatedAt = time.Now()

	task, err = s.taskRepository.UpdateTask(task)
	if err != nil {
		log.Printf("P=Service M=UpdateTask id=%v error=%v", id, err.Error())
		return task, err
	}

	return task, err
}

func (s *TaskService) DeleteTaskById(id string) (string, error) {
	rowsAffected, err := s.taskRepository.DeleteTaskById(id)
	if err != nil {
		log.Printf("P=Service M=DeleteTaskById id=%v error=%v", id, err.Error())
		return "fail", err
	}
	response := fmt.Sprintf("success with %v rows affected", rowsAffected)
	return response, err
}
