package repository

import (
	"log"

	"github.com/tweedledo/application/domain/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (t *TaskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	result := t.DB.Create(task)
	return task, result.Error
}
func (t *TaskRepository) UpdateTask(task *model.Task) (*model.Task, error) {
	result := t.DB.Save(task)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=UpdateTask no rows affected")
	}
	return task, result.Error
}
func (t *TaskRepository) DeleteTask(task *model.Task) (*model.Task, error) {
	result := t.DB.Delete(task)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=DeleteTask no rows affected")
	}
	return task, result.Error
}
func (t *TaskRepository) GetTasksInTaskList(tasklistId string) ([]*model.Task, error) {
	tasks := []*model.Task{}
	result := t.DB.Where("tasklist_id = ?", tasklistId).Find(tasks)
	return tasks, result.Error
}
