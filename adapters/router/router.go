package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tweedledo/adapters/controller"
)

func Initialize(controller *controller.TaskListController) {
	log.Printf("P=Router M=Initialize initializing routes...")

	router := gin.Default()
	router.GET("/tasklist", controller.GetAllTaskList)
	router.GET("/tasklist/:id", controller.GetTaskListById)
	router.POST("/tasklist", controller.CreateTaskList)
	router.Run(":8080")
}
