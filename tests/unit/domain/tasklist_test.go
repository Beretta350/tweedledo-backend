package tests

import (
	"testing"

	"github.com/tweedledo/core/domain"
)

func TestNewTaskListConstructorSuccess(t *testing.T) {
	name := "TestConstructor"
	task := []*domain.Task{}
	tasklist, err := domain.NewTaskList(name, task)
	if err != nil || tasklist == nil {
		t.Fatalf("P=Domain T=TestNewTaskListConstructorSuccess failed error=%v", err.Error())
	}
}
