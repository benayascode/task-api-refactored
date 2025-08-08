package routers

import (
	"task-manager/Delivery/controllers"
	infrastructure "task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

type Router struct {
	taskController *controllers.TaskController
	userController *controllers.UserController
	authMiddleware *infrastructure.AuthMiddleware
}

func NewRouter(
	taskController *controllers.TaskController,
	userController *controllers.UserController,
	authMiddleware *infrastructure.AuthMiddleware,
) *Router {
	return &Router{
		taskController: taskController,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", r.userController.Register)
	router.POST("/login", r.userController.Login)

	// Protected routes
	tasks := router.Group("/tasks")
	tasks.Use(r.authMiddleware.JWTAuthMiddleware())
	{
		tasks.GET("", r.taskController.GetTasks)
		tasks.GET("/:id", r.taskController.GetTaskByID)
	}

	// Admin only routes
	admin := router.Group("/")
	admin.Use(r.authMiddleware.JWTAuthMiddleware())
	admin.Use(r.authMiddleware.AdminOnly())
	{
		admin.POST("/tasks", r.taskController.CreateTask)
		admin.PUT("/tasks/:id", r.taskController.UpdateTask)
		admin.DELETE("/tasks/:id", r.taskController.DeleteTask)
		admin.POST("/promote/:username", r.userController.PromoteUser)
	}

	return router
}
