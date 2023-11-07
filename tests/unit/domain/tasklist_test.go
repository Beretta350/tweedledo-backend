package unit_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tweedledo/core/domain"
)

func TestTaskList_NewTaskListConstructorSuccess(t *testing.T) {
	name := "TestConstructor"
	task := []*domain.Task{}
	tasklist, err := domain.NewTaskList(name, task)
	if err != nil || tasklist == nil {
		t.Fatalf("P=Domain T=TestNewTaskListConstructorSuccess failed error=%v", err.Error())
		t.FailNow()
	}
}

func TestTaskList_NewTaskListConstructorError(t *testing.T) {
	task := []*domain.Task{}
	_, err := domain.NewTaskList("", task)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "name: Missing required field", err.Error())
}
