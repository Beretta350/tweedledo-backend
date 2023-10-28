package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tweedledo/adapters/controller"
)

func Initialize(controller *controller.TaskListController) {
	router := gin.Default()
	router.GET("/tasklist", controller.GetTaskListById)
	router.Run(":8080")
}
