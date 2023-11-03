package main

import (
	"os"

	"github.com/tweedledo/adapters/controller"
	"github.com/tweedledo/adapters/repository"
	"github.com/tweedledo/adapters/router"
	"github.com/tweedledo/core/service"
	"github.com/tweedledo/infrastructure/db"
)

func main() {
	database := db.ConnectDB(os.Getenv("env"))
	taskListRepository := repository.NewTaskListRepository(database)
	tasklistService := service.NewTaskListService(taskListRepository)
	taskRepository := repository.NewTaskRepository(database)
	taskService := service.NewTaskService(taskRepository, tasklistService)
	tasklistController := controller.NewTaskListController(tasklistService, taskService)
	taskController := controller.NewTaskController(taskService)
	router.Initialize(tasklistController, taskController)
}
