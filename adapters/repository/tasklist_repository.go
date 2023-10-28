package repository

import (
	"log"

	"github.com/tweedledo/core/domain"
	"gorm.io/gorm"
)

type TaskListRepository struct {
	DB *gorm.DB
}

func (t *TaskListRepository) CreateTaskList(tasklist *domain.TaskList) (*domain.TaskList, error) {
	result := t.DB.Create(tasklist)
	return tasklist, result.Error
}
func (t *TaskListRepository) UpdateTaskList(tasklist *domain.TaskList) (*domain.TaskList, error) {
	result := t.DB.Save(tasklist)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=UpdateTaskList no rows affected")
	}
	return tasklist, result.Error
}
func (t *TaskListRepository) DeleteTaskList(tasklist *domain.TaskList) (*domain.TaskList, error) {
	result := t.DB.Delete(tasklist)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=UpdateTaskList no rows affected")
	}
	return tasklist, result.Error
}
func (t *TaskListRepository) GetTaskListById(id string) (*domain.TaskList, error) {
	tasklist := domain.TaskList{}
	result := t.DB.Where("id = ?", id).First(tasklist)
	return &tasklist, result.Error
}
func (t *TaskListRepository) GetAllTasksLists() ([]*domain.TaskList, error) {
	tasklists := []*domain.TaskList{}
	result := t.DB.Find(tasklists)
	return tasklists, result.Error
}
