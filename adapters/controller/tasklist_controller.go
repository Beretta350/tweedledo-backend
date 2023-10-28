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
	log.Printf("P=Controller M=GetTaskListById tasklistId=%v", c.Param("id"))
	tasklist, err := ctrl.tasklistService.GetTaskListById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, tasklist)
}

func (ctrl *TaskListController) GetAllTaskList(c *gin.Context) {
	log.Printf("P=Controller M=GetAllTaskList")
	tasklist, err := ctrl.tasklistService.GetAllTaskList()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, tasklist)
}

func (ctrl *TaskListController) CreateTaskList(c *gin.Context) {
	log.Printf("P=Controller M=CreateTaskList name=%v", c.Query("name"))
	tasklist, err := ctrl.tasklistService.CreateTaskList(c.Query("name"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, tasklist)
}
