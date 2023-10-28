package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tweedledo/core/ports"
)

type TaskListController struct {
	tasklistService ports.TaskListServiceInterface
	taskService     ports.TaskServiceInterface
}

func NewTaskListController(
	tasklistService ports.TaskListServiceInterface,
	taskService ports.TaskServiceInterface,
) *TaskListController {
	return &TaskListController{
		tasklistService: tasklistService,
		taskService:     taskService,
	}
}

func (ctrl *TaskListController) GetTaskListById(c *gin.Context) {
	log.Printf("P=Controller M=GetTaskListById tasklistId=%v", c.Query("id"))
	tasklist, err := ctrl.tasklistService.GetTaskListById(c.Query("id"))
	if err != nil {
		log.Printf("P=Controller M=GetTaskListById error=%v", err.Error())
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, tasklist)
}
