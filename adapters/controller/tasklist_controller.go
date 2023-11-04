package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tweedledo/core/ports"
)

type TaskListController struct {
	tasklistService ports.TaskListServiceInterface
}

func NewTaskListController(tasklistService ports.TaskListServiceInterface) *TaskListController {
	return &TaskListController{
		tasklistService: tasklistService,
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
	tasklistJSON := struct{ Name string }{}
	if err := c.BindJSON(&tasklistJSON); err != nil {
		log.Printf("P=Controller M=CreateTaskList error=%v", err.Error())
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("P=Controller M=CreateTaskList name=%v", tasklistJSON.Name)

	tasklist, err := ctrl.tasklistService.CreateTaskList(tasklistJSON.Name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, tasklist)
}

func (ctrl *TaskListController) UpdateTaskList(c *gin.Context) {
	requestJSON := struct{ Name string }{}
	if err := c.BindJSON(&requestJSON); err != nil {
		log.Printf("P=Controller M=UpdateTaskList error=%v", err.Error())
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("P=Controller M=UpdateTaskList id=%v name=%v", c.Param("id"), requestJSON.Name)

	task, err := ctrl.tasklistService.UpdateTaskList(c.Param("id"), requestJSON.Name)
	if err != nil {
		log.Printf("P=Controller M=UpdateTaskList name=%v error=%v", requestJSON.Name, err.Error())
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, task)
}

func (ctrl *TaskListController) DeleteTaskListById(c *gin.Context) {
	log.Printf("P=Controller M=DeleteTaskListById id=%v", c.Param("id"))

	response, err := ctrl.tasklistService.DeleteTaskListById(c.Param("id"))
	if err != nil {
		log.Printf("P=Controller M=DeleteTaskListById id=%v error=%v", c.Param("id"), err.Error())
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": response})
}
