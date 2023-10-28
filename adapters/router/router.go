package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tweedledo/adapters/controller"
)

func Initialize(controller *controller.TaskListController) {
	log.Printf("P=Router M=Initialize initializing routes...")

	router := gin.Default()
	router.GET("/tasklist", controller.GetTaskListById)
	router.Run(":8080")
}
