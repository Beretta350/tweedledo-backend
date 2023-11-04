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

var taskSuccess = &domain.Task{
	Base: domain.Base{
		ID:        "TestID",
		CreatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
	},
	Name:        "Test",
	Description: "Simple description",
	TaskList:    &domain.TaskList{},
}

var taskUpdatedSuccess = &domain.Task{
	Base: domain.Base{
		ID:        "TestIDUpdated",
		CreatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2023, 10, 29, 00, 00, 00, 00, time.UTC),
	},
	Name:        "TestUpdated",
	Description: "Simple updated description",
	TaskList:    &domain.TaskList{},
}

func initTaskMocks(t *testing.T) (*gomock.Controller, *mock.MockTaskRepositoryInterface, *mock.MockTaskListServiceInterface) {
	mockCtrl := gomock.NewController(t)
	mockTaskListService := mock.NewMockTaskListServiceInterface(mockCtrl)
	mockTaskRepository := mock.NewMockTaskRepositoryInterface(mockCtrl)
	return mockCtrl, mockTaskRepository, mockTaskListService
}

func TestTask_ServiceGetTaskByIdSuccess(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskById(gomock.Any()).Return(taskSuccess, nil)

	service := service.NewTaskService(repository, taskListService)
	response, err := service.GetTaskById("TestID")
	if err != nil {
		t.Fatalf("P=Service T=TestTask_ServiceGetTaskByIdSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestID")
	assert.Equal(t, response.Name, "Test")
	assert.Equal(t, response.TaskList, &domain.TaskList{})
}

func TestTask_ServiceGetTaskByIdError(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskById(gomock.Any()).Return(nil, errors.New("Test error"))

	service := service.NewTaskService(repository, taskListService)
	_, err := service.GetTaskById("TestID")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}

func TestTask_ServiceCreateTaskSuccess(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	taskListService.EXPECT().GetTaskListById(gomock.Any()).Return(tasklistSuccess, nil)
	repository.EXPECT().CreateTask(gomock.Any()).Return(taskSuccess, nil)

	service := service.NewTaskService(repository, taskListService)
	response, err := service.CreateTask("Test", "Simple description", "TestListID")
	if err != nil {
		t.Fatalf("P=Service T=TestTask_ServiceCreateTaskSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestID")
	assert.Equal(t, response.Name, "Test")
	assert.Equal(t, response.TaskList, &domain.TaskList{})
}

func TestTask_ServiceCreateTaskRepositoryError(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	taskListService.EXPECT().GetTaskListById(gomock.Any()).Return(tasklistSuccess, nil)
	repository.EXPECT().CreateTask(gomock.Any()).Return(nil, errors.New("Test error"))

	service := service.NewTaskService(repository, taskListService)
	_, err := service.CreateTask("Test", "Simple description", "TestListID")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "Test error")
}

func TestTask_ServiceCreateTaskNewTaskError(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	taskListService.EXPECT().GetTaskListById(gomock.Any()).Return(tasklistSuccess, nil)

	service := service.NewTaskService(repository, taskListService)
	_, err := service.CreateTask("", "Simple description", "TestListID")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, "name: Missing required field", err.Error())
}

func TestTask_ServiceCreateTaskGetTaskListError(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	taskListService.EXPECT().GetTaskListById(gomock.Any()).Return(nil, errors.New("Test Error"))

	service := service.NewTaskService(repository, taskListService)
	_, err := service.CreateTask("", "Simple description", "TestListID")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, "Test Error", err.Error())
}

func TestTask_ServiceUpdateTaskSuccess(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskById(gomock.Any()).Return(taskSuccess, nil)
	repository.EXPECT().UpdateTask(gomock.Any()).Return(taskUpdatedSuccess, nil)

	service := service.NewTaskService(repository, taskListService)
	response, err := service.UpdateTask("TestIDUpdated", "TestUpdated", "Simple updated description")
	if err != nil {
		t.Fatalf("P=Service T=TestTask_ServiceUpdateTaskSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response.ID, "TestIDUpdated")
	assert.Equal(t, response.Name, "TestUpdated")
	assert.Equal(t, response.Description, "Simple updated description")
	assert.Equal(t, response.TaskList, &domain.TaskList{})
}

func TestTask_ServiceUpdateTaskGetTaskError(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskById(gomock.Any()).Return(nil, errors.New("Test Error"))

	service := service.NewTaskService(repository, taskListService)
	_, err := service.UpdateTask("TestIDUpdated", "TestUpdated", "Simple updated description")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, "Test Error", err.Error())
}

func TestTask_ServiceUpdateTaskUpdateError(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().GetTaskById(gomock.Any()).Return(taskSuccess, nil)
	repository.EXPECT().UpdateTask(gomock.Any()).Return(nil, errors.New("Test Error"))

	service := service.NewTaskService(repository, taskListService)
	_, err := service.UpdateTask("TestIDUpdated", "TestUpdated", "Simple updated description")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, "Test Error", err.Error())
}

func TestTask_ServiceDeleteTaskSuccess(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().DeleteTaskById(gomock.Any()).Return(int64(1), nil)

	service := service.NewTaskService(repository, taskListService)
	response, err := service.DeleteTaskById("TestID")
	if err != nil {
		t.Fatalf("P=Service T=TestTask_ServiceDeleteTaskSuccess failed error=%v", err.Error())
		t.FailNow()
	}

	assert.Equal(t, response, "success with 1 rows affected")
}

func TestTask_ServiceDeleteTaskError(t *testing.T) {
	mockCtrl, repository, taskListService := initTaskMocks(t)
	defer mockCtrl.Finish()
	repository.EXPECT().DeleteTaskById(gomock.Any()).Return(int64(0), errors.New("Test Error"))

	service := service.NewTaskService(repository, taskListService)
	response, err := service.DeleteTaskById("TestID")

	assert.NotEqual(t, err, nil)
	assert.Equal(t, "Test Error", err.Error())
	assert.Equal(t, response, "fail")
}
