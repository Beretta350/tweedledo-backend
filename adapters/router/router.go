package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tweedledo/adapters/controller"
)

func Initialize(tasklistCtrl *controller.TaskListController, taskCtrl *controller.TaskController) {
	log.Printf("P=Router M=Initialize initializing routes...")

	router := gin.Default()
	//tasklists
	router.GET("/tasklist", tasklistCtrl.GetAllTaskList)
	router.GET("/tasklist/:id", tasklistCtrl.GetTaskListById)
	router.POST("/tasklist", tasklistCtrl.CreateTaskList)
	router.PUT("/tasklist/:id", tasklistCtrl.UpdateTaskList)
	router.DELETE("/tasklist/:id", tasklistCtrl.DeleteTaskListById)

	//tasks
	router.GET("/task/:id", taskCtrl.GetTaskById)
	router.POST("/task", taskCtrl.CreateTask)
	router.PUT("/task/:id", taskCtrl.UpdateTask)
	router.DELETE("/task/:id", taskCtrl.DeleteTask)

	//run
	router.Run(":8080")
}
