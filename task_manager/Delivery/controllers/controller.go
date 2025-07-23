package controllers

import (
	"context"
	"net/http"
	"strconv"
	"task_manager/Domain"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TaskUC *Usecases.TaskUseCase
	UserUC *Usecases.UserUseCase
}

func (ctr *Controller) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	if err := ctr.UserUC.Register(context.Background(), req.Username, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ctr *Controller) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	token, err := ctr.UserUC.Login(context.Background(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctr *Controller) GetTasks(c *gin.Context) {
	tasks, err := ctr.TaskUC.GetAllTasks(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (ctr *Controller) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	task, err := ctr.TaskUC.GetTaskByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (ctr *Controller) CreateTask(c *gin.Context) {
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	if err := ctr.TaskUC.CreateTask(context.Background(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func (ctr *Controller) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	if err := ctr.TaskUC.UpdateTask(context.Background(), id, task); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (ctr *Controller) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	if err := ctr.TaskUC.DeleteTask(context.Background(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (ctr *Controller) PromoteUser(c *gin.Context) {
	username := c.Param("username")
	if err := ctr.UserUC.Promote(context.Background(), username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}
