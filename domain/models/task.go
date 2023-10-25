package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Task struct {
	Base        `valid:"required"`
	Name        string    `gorm:"column:name;type:varchar(255);not null" valid:"notnull"`
	Description string    `gorm:"column:description;type:text" valid:"-"`
	TaskList    *TaskList `valid:"-"`
	TaskListID  string    `gorm:"column:tasklist_id;type:uuid;not null" valid:"-"`
}

func (t *Task) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	return nil
}

// constructor
func NewTask(name string, desc string, tasklist *TaskList) (*Task, error) {
	task := Task{
		Name:        name,
		Description: desc,
		TaskList:    tasklist,
	}

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err := task.isValid()
	if err != nil {
		return &task, err
	}

	return &task, nil
}
