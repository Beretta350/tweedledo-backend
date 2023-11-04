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

var tasklistUpdatedSuccess = &domain.TaskList{
	Base: domain.Base{
		ID:        "TestIDUpdated",
		CreatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
	},
	Name:  "TestUpdated",
	Tasks: []*domain.Task{},
}

func initTaskListMocks(t *testing.T) (*gomock.Controller, *mock.MockTaskListRepositoryInterface) {
	mockCtrl := gomock.NewController(t)
	mockTaskRepository := mock.NewMockTaskListRepositoryInterface(mockCtrl)
	return mockCtrl, mockTaskRepository
}

func TestTaskList_ServiceGetTaskListByIdSuccess(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskListById(gomock.Any()).Return(tasklistSuccess, nil)

	service := service.NewTaskListService(repository)
	response, err := service.GetTaskListById("TEST")
	if err != nil {
		t.Fatalf("P=Service T=TestTaskList_ServiceGetTaskListByIdSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestID")
	assert.Equal(t, response.Name, "Test")
	assert.Equal(t, response.Tasks, []*domain.Task{})
}

func TestTaskList_ServiceGetTaskListByIdError(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()

	repository.EXPECT().GetTaskListById(gomock.Any()).Return(nil, errors.New("Test error"))

	service := service.NewTaskListService(repository)
	_, err := service.GetTaskListById("TEST")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}

func TestTaskList_ServiceGetAllTaskListSuccess(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()

	tasklists := []*domain.TaskList{}
	tasklists = append(tasklists, tasklistSuccess)

	repository.EXPECT().GetAllTaskLists().Return(tasklists, nil)

	service := service.NewTaskListService(repository)
	response, err := service.GetAllTaskList()
	if err != nil {
		t.Fatalf("P=Service T=TestTaskList_ServiceGetAllTaskListSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, len(response), 1)
	assert.Equal(t, response[0].ID, "TestID")
	assert.Equal(t, response[0].Name, "Test")
	assert.Equal(t, response[0].Tasks, []*domain.Task{})
}

func TestTaskList_ServiceGetAllTaskListError(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetAllTaskLists().Return(nil, errors.New("Test error"))

	service := service.NewTaskListService(repository)
	_, err := service.GetAllTaskList()

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}

func TestTaskList_ServiceCreateTaskListSuccess(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().CreateTaskList(gomock.Any()).Return(tasklistSuccess, nil)

	service := service.NewTaskListService(repository)
	response, err := service.CreateTaskList("Test")
	if err != nil {
		t.Fatalf("P=Service T=TestTaskList_ServiceCreateTaskListSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestID")
	assert.Equal(t, response.Name, "Test")
	assert.Equal(t, response.Tasks, []*domain.Task{})
}

func TestTaskList_ServiceCreateTaskListRepositoryError(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().CreateTaskList(gomock.Any()).Return(nil, errors.New("Test error"))

	service := service.NewTaskListService(repository)
	_, err := service.CreateTaskList("Test")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}

func TestTaskList_ServiceCreateTaskListNewTaskListError(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()

	service := service.NewTaskListService(repository)
	_, err := service.CreateTaskList("")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, "name: Missing required field", err.Error())
}

func TestTaskList_ServiceUpdateTaskList(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskListById(gomock.Any()).Return(tasklistSuccess, nil)
	repository.EXPECT().UpdateTaskList(gomock.Any()).Return(tasklistUpdatedSuccess, nil)

	service := service.NewTaskListService(repository)
	response, err := service.UpdateTaskList("TestIDUpdated", "TestUpdated")
	if err != nil {
		t.Fatalf("P=Service T=TestTask_ServiceUpdateTaskSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestIDUpdated")
	assert.Equal(t, response.Name, "TestUpdated")
}

func TestTaskList_ServiceUpdateTaskListGetTaskListError(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskListById(gomock.Any()).Return(nil, errors.New("Test Error"))

	service := service.NewTaskListService(repository)
	_, err := service.UpdateTaskList("TestIDUpdated", "TestUpdated")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test Error")
}

func TestTaskList_ServiceUpdateTaskListRepositoryError(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskListById(gomock.Any()).Return(tasklistSuccess, nil)
	repository.EXPECT().UpdateTaskList(gomock.Any()).Return(nil, errors.New("Test Error"))

	service := service.NewTaskListService(repository)
	_, err := service.UpdateTaskList("TestIDUpdated", "TestUpdated")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test Error")
}

func TestTaskList_ServiceDeleteTaskListSuccess(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().DeleteTaskListById(gomock.Any()).Return(int64(1), nil)

	service := service.NewTaskListService(repository)
	response, err := service.DeleteTaskListById("TestID")
	if err != nil {
		t.Fatalf("P=Service T=TestTask_ServiceDeleteTaskSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response, "success with 1 rows affected")
}

func TestTaskList_ServiceDeleteTaskListError(t *testing.T) {
	mockCtrl, repository := initTaskListMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().DeleteTaskListById(gomock.Any()).Return(int64(0), errors.New("Test Error"))

	service := service.NewTaskListService(repository)
	response, err := service.DeleteTaskListById("TestID")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, "Test Error", err.Error())
	assert.Equal(t, response, "fail")
}
