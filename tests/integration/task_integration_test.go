package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tweedledo/adapters/controller"
	"github.com/tweedledo/adapters/repository"
	"github.com/tweedledo/adapters/router"
	"github.com/tweedledo/core/domain"
	"github.com/tweedledo/core/service"
	"github.com/tweedledo/infrastructure/db"
)

var taskRepository *repository.TaskRepository

func setupTaskIntegrationTest() *gin.Engine {
	os.Setenv("env", "test")
	database := db.ConnectDB(os.Getenv("env"))
	taskListRepository := repository.NewTaskListRepository(database)
	tasklistService := service.NewTaskListService(taskListRepository)
	taskRepository = repository.NewTaskRepository(database)
	taskService := service.NewTaskService(taskRepository, tasklistService)
	taskController := controller.NewTaskController(taskService)
	tasklistController := controller.NewTaskListController(tasklistService)
	routes := gin.Default()
	routes = router.AddTasksRoutes(routes, taskController)
	routes.POST("/tasklist", tasklistController.CreateTaskList)
	return routes
}

func createJsonPayloadForCreateTask(name, description, tasklist_id string) *bytes.Buffer {
	jsonPayloadStr := fmt.Sprintf(`{"name": "%v", "description": "%v", "tasklist_id": "%v"}`, name, description, tasklist_id)
	jsonPayloadBytes := []byte(jsonPayloadStr)
	return bytes.NewBuffer(jsonPayloadBytes)
}

func createJsonPayloadForUpdateTask(name, description string) *bytes.Buffer {
	jsonPayloadStr := fmt.Sprintf(`{"name": "%v", "description": "%v"}`, name, description)
	jsonPayloadBytes := []byte(jsonPayloadStr)
	return bytes.NewBuffer(jsonPayloadBytes)
}

func executeCreateTaskPostRequest(t *testing.T, router *gin.Engine, name, description, tasklist_id string) domain.Task {
	jsonPayloadBuf := createJsonPayloadForCreateTask(name, description, tasklist_id)
	req, _ := http.NewRequest("POST", "/task", jsonPayloadBuf)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
		t.FailNow()
	}

	var response domain.Task
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("handler returned unexpected body")
		t.FailNow()
	}

	return response
}

func TestIntegrationTask_GetTaskByIdExists(t *testing.T) {
	router := setupTaskIntegrationTest()
	responseTaskList := executeCreateTaskListPostRequest(t, router, "IntegrationTaskListTest")
	responseTask := executeCreateTaskPostRequest(t, router, "IntegrationTest", "Test Description", responseTaskList.ID)
	getRequest := fmt.Sprintf(`/task/%v`, responseTask.ID)
	req, _ := http.NewRequest("GET", getRequest, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	responsePostByte, err := json.Marshal(responseTask)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(responsePostByte))
		t.FailNow()
	}

	assert.Equal(t, rr.Body.String(), string(responsePostByte))
}

func TestIntegrationTask_GetTaskByIdDontExists(t *testing.T) {
	router := setupTaskIntegrationTest()
	req, _ := http.NewRequest("GET", "/task/invalid-id", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusInternalServerError)
		t.FailNow()
	}

	expected := "{\"message\":\"record not found\"}"

	assert.Equal(t, rr.Body.String(), expected)
}

func TestIntegrationTask_CreateTaskSuccess(t *testing.T) {
	router := setupTaskIntegrationTest()
	responseTaskList := executeCreateTaskListPostRequest(t, router, "IntegrationTaskListTest")
	responseTask := executeCreateTaskPostRequest(t, router, "IntegrationTest", "Test Description", responseTaskList.ID)

	taskData, err := taskRepository.GetTaskById(responseTask.ID)
	if taskData == nil || err != nil {
		t.Errorf("Error to get data from database: got %v want %v", nil, responseTask)
		t.FailNow()
	}

	assert.Equal(t, responseTask.Name, "IntegrationTest")
	assert.Equal(t, responseTask.Description, "Test Description")
	assert.Equal(t, taskData.ID, responseTask.ID)
	assert.Equal(t, taskData.Name, responseTask.Name)
	assert.Equal(t, taskData.TaskListID, responseTaskList.ID)
}

func TestIntegrationTask_UpdateTaskSuccess(t *testing.T) {
	router := setupTaskIntegrationTest()
	responseTaskList := executeCreateTaskListPostRequest(t, router, "IntegrationTaskListTest")
	responseTask := executeCreateTaskPostRequest(t, router, "IntegrationTest", "Test Description", responseTaskList.ID)

	putRequest := fmt.Sprintf(`/task/%v`, responseTask.ID)
	jsonPayload := createJsonPayloadForUpdateTask("UpdatedIntegrationTest", "Updated description")
	req, _ := http.NewRequest("PUT", putRequest, jsonPayload)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	taskData, err := taskRepository.GetTaskById(responseTask.ID)
	if taskData == nil || err != nil {
		t.Errorf("Error to get data from database: got %v want %v", nil, responseTask)
		t.FailNow()
	}

	assert.Equal(t, taskData.Name, "UpdatedIntegrationTest")
	assert.Equal(t, taskData.Description, "Updated description")
	assert.Equal(t, taskData.ID, responseTask.ID)
	assert.NotEqual(t, taskData.Name, responseTask.Name)
	assert.NotEqual(t, responseTask.UpdatedAt, taskData.UpdatedAt)
}

func TestIntegrationTask_DeleteTask(t *testing.T) {
	router := setupTaskIntegrationTest()
	responseTaskList := executeCreateTaskListPostRequest(t, router, "IntegrationTaskListTest")
	responseTask := executeCreateTaskPostRequest(t, router, "IntegrationTest", "Test Description", responseTaskList.ID)

	putRequest := fmt.Sprintf(`/task/%v`, responseTask.ID)
	req, _ := http.NewRequest("DELETE", putRequest, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	deleteResponse := "{\"message\":\"success with 1 rows affected\"}"
	assert.Equal(t, rr.Body.String(), deleteResponse)

	_, err := taskRepository.GetTaskById(responseTask.ID)
	if err == nil {
		t.Errorf("Error to get data from database: got %v want %v", nil, responseTask)
		t.FailNow()
	}

	assert.Equal(t, err.Error(), "record not found")
}

func TestIntegrationTask_DeleteTaskAlreadyDeleted(t *testing.T) {
	router := setupTaskIntegrationTest()
	putRequest := fmt.Sprintf(`/task/%v`, "anyID")
	req, _ := http.NewRequest("DELETE", putRequest, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	deleteResponse := "{\"message\":\"success with 0 rows affected\"}"
	assert.Equal(t, rr.Body.String(), deleteResponse)

	_, err := taskRepository.GetTaskById("anyID")
	if err == nil {
		t.Errorf("Error to get data from database: got %v want %v", nil, "anyID")
		t.FailNow()
	}

	assert.Equal(t, err.Error(), "record not found")
}
