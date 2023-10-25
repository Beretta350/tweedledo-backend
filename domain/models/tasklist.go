package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type TaskList struct {
	Base  `valid:"required"`
	Name  string  `gorm:"column:name;type:varchar(255);not null" valid:"notnull"`
	Tasks []*Task `gorm:"ForeignKey:TaskListID" valid:"-"`
}

func (t *TaskList) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	return nil
}

func NewTaskList(name string, tasks []*Task) (*TaskList, error) {
	taskList := TaskList{
		Name:  name,
		Tasks: tasks,
	}

	taskList.ID = uuid.New().String()
	taskList.CreatedAt = time.Now()
	taskList.UpdatedAt = time.Now()

	err := taskList.isValid()
	if err != nil {
		return &taskList, err
	}

	return &taskList, nil
}
