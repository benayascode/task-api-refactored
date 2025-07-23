package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctr *controllers.Controller) *gin.Engine {
	r := gin.Default()

	r.POST("/register", ctr.Register)
	r.POST("/login", ctr.Login)
	r.GET("/tasks", Infrastructure.JWTAuthMiddleware(), ctr.GetTasks)
	r.GET("/tasks/:id", Infrastructure.JWTAuthMiddleware(), ctr.GetTaskByID)
	r.POST("/tasks", Infrastructure.JWTAuthMiddleware(), Infrastructure.AdminOnly(), ctr.CreateTask)
	r.PUT("/tasks/:id", Infrastructure.JWTAuthMiddleware(), Infrastructure.AdminOnly(), ctr.UpdateTask)
	r.DELETE("/tasks/:id", Infrastructure.JWTAuthMiddleware(), Infrastructure.AdminOnly(), ctr.DeleteTask)
	r.POST("/promote/:username", Infrastructure.JWTAuthMiddleware(), Infrastructure.AdminOnly(), ctr.PromoteUser)

	return r
}
