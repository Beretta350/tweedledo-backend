package unit_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tweedledo/core/domain"
)

func TestTask_NewTaskConstructorSuccess(t *testing.T) {
	name := "TestConstructor"
	description := "Test description"
	task, err := domain.NewTask(name, description, &domain.TaskList{})
	if err != nil || task == nil {
		t.Fatalf("P=Domain T=TestNewTaskListConstructorSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, name, task.Name)
	assert.Equal(t, description, task.Description)
}

func TestTask_NewTaskListConstructorError(t *testing.T) {
	description := "Test description"
	_, err := domain.NewTask("", description, &domain.TaskList{})
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "name: Missing required field", err.Error())
}
