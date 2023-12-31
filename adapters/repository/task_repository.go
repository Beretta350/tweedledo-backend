package repository

import (
	"log"

	"github.com/tweedledo/core/domain"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(database *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: database,
	}
}

func (t *TaskRepository) CreateTask(task *domain.Task) (*domain.Task, error) {
	result := t.DB.Create(task)
	return task, result.Error
}
func (t *TaskRepository) UpdateTask(task *domain.Task) (*domain.Task, error) {
	result := t.DB.Save(task)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=UpdateTask no rows affected")
	}
	return task, result.Error
}
func (t *TaskRepository) DeleteTaskById(id string) (int64, error) {
	var task *domain.Task
	result := t.DB.Where("id = ?", id).Delete(&task)
	if result.RowsAffected == 0 {
		log.Print("P=Repository M=DeleteTask no rows affected")
	}
	return result.RowsAffected, result.Error
}
func (t *TaskRepository) GetTaskById(id string) (*domain.Task, error) {
	var task *domain.Task
	result := t.DB.Where("id = ?", id).First(&task)
	return task, result.Error
}
func (t *TaskRepository) GetTasksInTaskList(tasklistId string) ([]*domain.Task, error) {
	tasks := []*domain.Task{}
	result := t.DB.Where("tasklist_id = ?", tasklistId).Find(tasks)
	return tasks, result.Error
}
