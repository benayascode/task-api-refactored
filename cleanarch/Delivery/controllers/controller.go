package controllers

import (
	"net/http"
	"strconv"
	domain "task-manager/Domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase domain.TaskUseCase
}

func NewTaskController(taskUseCase domain.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}

func (t *TaskController) GetTasks(c *gin.Context) {
	tasks, err := t.taskUseCase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := t.taskUseCase.GetTaskByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := t.taskUseCase.CreateTask(newTask)
	if err != nil {
		switch err {
		case domain.ErrInvalidTaskTitle:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task title cannot be empty"})
		case domain.ErrInvalidTaskDescription:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task description cannot be empty"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	task, err := t.taskUseCase.UpdateTask(userID, updatedTask)
	if err != nil {
		switch err {
		case domain.ErrInvalidTaskTitle:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task title cannot be empty"})
		case domain.ErrInvalidTaskDescription:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task description cannot be empty"})
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": task})
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = t.taskUseCase.DeleteTask(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

type UserController struct {
	userUseCase domain.UserUseCase
}

func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := u.userUseCase.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (u *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	token, err := u.userUseCase.LoginUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserController) PromoteUser(c *gin.Context) {
	username := c.Param("username")

	err := u.userUseCase.PromoteUser(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}
