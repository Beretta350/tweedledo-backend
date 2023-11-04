package tests

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

var taskListRepository *repository.TaskListRepository
var taskRepository *repository.TaskRepository

func setupTest() *gin.Engine {
	os.Setenv("env", "test")
	database := db.ConnectDB(os.Getenv("env"))
	taskListRepository = repository.NewTaskListRepository(database)
	tasklistService := service.NewTaskListService(taskListRepository)
	taskRepository = repository.NewTaskRepository(database)
	tasklistController := controller.NewTaskListController(tasklistService)
	routes := gin.Default()
	routes = router.AddTaskListsRoutes(routes, tasklistController)
	return routes
}

func createJsonPayloadWithName(name string) *bytes.Buffer {
	jsonPayloadStr := fmt.Sprintf(`{"name": "%v"}`, name)
	jsonPayloadBytes := []byte(jsonPayloadStr)
	return bytes.NewBuffer(jsonPayloadBytes)
}

func executeCreateTaskListPostRequest(t *testing.T, router *gin.Engine, name string) domain.TaskList {
	jsonPayloadBuf := createJsonPayloadWithName(name)
	req, _ := http.NewRequest("POST", "/tasklist", jsonPayloadBuf)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
		t.FailNow()
	}

	var response domain.TaskList
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("handler returned unexpected body")
		t.FailNow()
	}

	return response
}

func TestIntegrationTaskList_GetTaskListByIdDontExists(t *testing.T) {
	router := setupTest()
	req, _ := http.NewRequest("GET", "/tasklist/invalid-id", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	expected := "{\"message\":\"record not found\"}"

	assert.Equal(t, rr.Body.String(), expected)
}

func TestIntegrationTaskList_GetTaskListByIdExists(t *testing.T) {
	router := setupTest()
	responsePost := executeCreateTaskListPostRequest(t, router, "IntegrationTest")
	getRequest := fmt.Sprintf(`/tasklist/%v`, responsePost.ID)
	req, _ := http.NewRequest("GET", getRequest, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	responsePostByte, err := json.Marshal(responsePost)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(responsePostByte))
		t.FailNow()
	}

	assert.Equal(t, rr.Body.String(), string(responsePostByte))
}

func TestIntegrationTaskList_GetAllTaskListNoData(t *testing.T) {
	router := setupTest()
	req, _ := http.NewRequest("GET", "/tasklist", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	expected := "[]"

	assert.Equal(t, rr.Body.String(), expected)
}

func TestIntegrationTaskList_GetAllTaskListWithData(t *testing.T) {
	router := setupTest()
	responsePost := executeCreateTaskListPostRequest(t, router, "IntegrationTest")
	req, _ := http.NewRequest("GET", "/tasklist", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	var sliceTaskList []domain.TaskList
	sliceTaskList = append(sliceTaskList, responsePost)

	responsePostByte, err := json.Marshal(sliceTaskList)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(responsePostByte))
		t.FailNow()
	}

	assert.Equal(t, rr.Body.String(), string(responsePostByte))
}

func TestIntegrationTaskList_CreateTaskList(t *testing.T) {
	router := setupTest()
	response := executeCreateTaskListPostRequest(t, router, "IntegrationTest")

	tasklistData, err := taskListRepository.GetTaskListById(response.ID)
	if tasklistData == nil || err != nil {
		t.Errorf("Error to get data from database: got %v want %v", nil, response)
		t.FailNow()
	}

	assert.Equal(t, response.Name, "IntegrationTest")
	assert.Equal(t, tasklistData.ID, response.ID)
	assert.Equal(t, tasklistData.Name, response.Name)
}

func TestIntegrationTaskList_UpdateTaskList(t *testing.T) {
	router := setupTest()
	response := executeCreateTaskListPostRequest(t, router, "IntegrationTest")
	putRequest := fmt.Sprintf(`/tasklist/%v`, response.ID)
	jsonPayload := createJsonPayloadWithName("UpdatedIntegrationTest")
	req, _ := http.NewRequest("PUT", putRequest, jsonPayload)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	tasklistData, err := taskListRepository.GetTaskListById(response.ID)
	if tasklistData == nil || err != nil {
		t.Errorf("Error to get data from database: got %v want %v", nil, response)
		t.FailNow()
	}

	assert.Equal(t, tasklistData.ID, response.ID)
	assert.Equal(t, tasklistData.Name, "UpdatedIntegrationTest")
	assert.NotEqual(t, tasklistData.Name, response.Name)
	assert.NotEqual(t, response.UpdatedAt, tasklistData.UpdatedAt)
}

func TestIntegrationTaskList_DeleteTaskList(t *testing.T) {
	router := setupTest()
	response := executeCreateTaskListPostRequest(t, router, "IntegrationTest")
	putRequest := fmt.Sprintf(`/tasklist/%v`, response.ID)
	req, _ := http.NewRequest("DELETE", putRequest, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		t.FailNow()
	}

	_, err := taskListRepository.GetTaskListById(response.ID)
	if err == nil {
		t.Errorf("Error to get data from database: got %v want %v", nil, response)
		t.FailNow()
	}

	assert.Equal(t, err.Error(), "record not found")
}
