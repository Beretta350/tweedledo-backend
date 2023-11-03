package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tweedledo/adapters/controller"
)

func Initialize(tasklistCtrl *controller.TaskListController, taskCtrl *controller.TaskController) {
	log.Printf("P=Router M=Initialize initializing routes...")

	router := gin.Default()
	router.GET("/tasklist", tasklistCtrl.GetAllTaskList)
	router.GET("/tasklist/:id", tasklistCtrl.GetTaskListById)
	router.POST("/tasklist", tasklistCtrl.CreateTaskList)
	router.GET("/task/:id", taskCtrl.GetTaskById)
	router.POST("/task", taskCtrl.CreateTask)
	router.PUT("/task/:id", taskCtrl.UpdateTask)
	router.DELETE("/task/:id", taskCtrl.DeleteTask)
	router.Run(":8080")
}
