package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tweedledo/core/ports"
)

type TaskController struct {
	taskService ports.TaskServiceInterface
}

func NewTaskController(
	taskService ports.TaskServiceInterface,
) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (ctrl *TaskController) GetTaskById(c *gin.Context) {
	log.Printf("P=Controller M=GetTaskById id=%v", c.Param("id"))

	task, err := ctrl.taskService.GetTaskById(c.Param("id"))
	if err != nil {
		log.Printf("P=Controller M=GetTaskById id=%v error=%v", c.Param("id"), err.Error())
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, task)
}

func (ctrl *TaskController) CreateTask(c *gin.Context) {
	requestJSON := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		TaskListId  string `json:"tasklist_id"`
	}{}
	if err := c.BindJSON(&requestJSON); err != nil {
		log.Printf("P=Controller M=CreateTask error=%v", err.Error())
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("P=Controller M=CreateTask name=%v tasklistId=%v", requestJSON.Name, requestJSON.TaskListId)

	task, err := ctrl.taskService.CreateTask(requestJSON.Name, requestJSON.Description, requestJSON.TaskListId)
	if err != nil {
		log.Printf("P=Controller M=CreateTask name=%v error=%v", requestJSON.Name, err.Error())
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, task)
}

func (ctrl *TaskController) UpdateTask(c *gin.Context) {
	requestJSON := struct {
		Name        string
		Description string
	}{}
	if err := c.BindJSON(&requestJSON); err != nil {
		log.Printf("P=Controller M=UpdateTask error=%v", err.Error())
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("P=Controller M=UpdateTask id=%v name=%v", c.Param("id"), requestJSON.Name)

	task, err := ctrl.taskService.UpdateTask(c.Param("id"), requestJSON.Name, requestJSON.Description)
	if err != nil {
		log.Printf("P=Controller M=UpdateTask name=%v error=%v", requestJSON.Name, err.Error())
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, task)
}

func (ctrl *TaskController) DeleteTask(c *gin.Context) {
	log.Printf("P=Controller M=DeleteTask id=%v", c.Param("id"))

	response, err := ctrl.taskService.DeleteTaskById(c.Param("id"))
	if err != nil {
		log.Printf("P=Controller M=DeleteTask id=%v error=%v", c.Param("id"), err.Error())
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": response})
}
