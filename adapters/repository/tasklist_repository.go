package repository

import (
	"log"

	"github.com/tweedledo/application/domain/model"
	"gorm.io/gorm"
)

type TaskListRepository struct {
	DB *gorm.DB
}

func (t *TaskListRepository) CreateTaskList(tasklist *model.TaskList) (*model.TaskList, error) {
	result := t.DB.Create(tasklist)
	return tasklist, result.Error
}
func (t *TaskListRepository) UpdateTaskList(tasklist *model.TaskList) (*model.TaskList, error) {
	result := t.DB.Save(tasklist)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=UpdateTaskList no rows affected")
	}
	return tasklist, result.Error
}
func (t *TaskListRepository) DeleteTaskList(tasklist *model.TaskList) (*model.TaskList, error) {
	result := t.DB.Delete(tasklist)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=UpdateTaskList no rows affected")
	}
	return tasklist, result.Error
}
func (t *TaskListRepository) GetTaskListById(id string) (*model.TaskList, error) {
	tasklist := model.TaskList{}
	result := t.DB.Where("id = ?", id).First(tasklist)
	return &tasklist, result.Error
}
func (t *TaskListRepository) GetAllTasksLists() ([]*model.TaskList, error) {
	tasklists := []*model.TaskList{}
	result := t.DB.Find(tasklists)
	return tasklists, result.Error
}
