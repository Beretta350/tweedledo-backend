package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Task struct {
	Base        `valid:"required"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(255);not null" valid:"notnull"`
	Description string    `json:"description" gorm:"column:description;type:text" valid:"-"`
	TaskList    *TaskList `json:"-" valid:"-"`
	TaskListID  string    `json:"tasklist_id" gorm:"column:tasklist_id;type:uuid;not null" valid:"-"`
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
