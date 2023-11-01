package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tweedledo/core/domain"
	"github.com/tweedledo/core/ports/mock"
	"github.com/tweedledo/core/service"
)

var tasklistSuccess = &domain.TaskList{
	Base: domain.Base{
		ID:        "TestID",
		CreatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
	},
	Name:  "Test",
	Tasks: []*domain.Task{},
}

func TestServiceGetTaskListByIdSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockTaskListRepositoryInterface(mockCtrl)
	mockObj.EXPECT().GetTaskListById(gomock.Any()).Return(tasklistSuccess, nil)

	service := service.NewTaskListService(mockObj)
	response, err := service.GetTaskListById("TEST")
	if err != nil {
		t.Fatalf("P=Service T=TestServiceGetTaskListByIdSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestID")
	assert.Equal(t, response.Name, "Test")
	assert.Equal(t, response.Tasks, []*domain.Task{})
}

func TestServiceGetTaskListByIdError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockTaskListRepositoryInterface(mockCtrl)
	mockObj.EXPECT().GetTaskListById(gomock.Any()).Return(nil, errors.New("Test error"))

	service := service.NewTaskListService(mockObj)
	_, err := service.GetTaskListById("TEST")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}

func TestServiceGetAllTaskListSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tasklists := []*domain.TaskList{}
	tasklists = append(tasklists, tasklistSuccess)

	mockObj := mock.NewMockTaskListRepositoryInterface(mockCtrl)
	mockObj.EXPECT().GetAllTaskLists().Return(tasklists, nil)

	service := service.NewTaskListService(mockObj)
	response, err := service.GetAllTaskList()
	if err != nil {
		t.Fatalf("P=Service T=TestServiceGetAllTaskListSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, len(response), 1)
	assert.Equal(t, response[0].ID, "TestID")
	assert.Equal(t, response[0].Name, "Test")
	assert.Equal(t, response[0].Tasks, []*domain.Task{})
}

func TestServiceGetAllTaskListError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockTaskListRepositoryInterface(mockCtrl)
	mockObj.EXPECT().GetAllTaskLists().Return(nil, errors.New("Test error"))

	service := service.NewTaskListService(mockObj)
	_, err := service.GetAllTaskList()

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}

func TestServiceCreateTaskListSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockTaskListRepositoryInterface(mockCtrl)
	mockObj.EXPECT().CreateTaskList(gomock.Any()).Return(tasklistSuccess, nil)

	service := service.NewTaskListService(mockObj)
	response, err := service.CreateTaskList("Test")
	if err != nil {
		t.Fatalf("P=Service T=TestServiceCreateTaskListSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestID")
	assert.Equal(t, response.Name, "Test")
	assert.Equal(t, response.Tasks, []*domain.Task{})
}

func TestServiceCreateTaskListError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockTaskListRepositoryInterface(mockCtrl)
	mockObj.EXPECT().CreateTaskList(gomock.Any()).Return(nil, errors.New("Test error"))

	service := service.NewTaskListService(mockObj)
	_, err := service.CreateTaskList("Test")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}
