package main

import (
	"context"
	"log"
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	infrastructure "task-manager/Infrastructure"
	repositories "task-manager/Repositories"
	usecases "task-manager/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	tasksCollection := client.Database("taskmanagerdb").Collection("tasks")
	usersCollection := client.Database("taskmanagerdb").Collection("users")

	jwtService := infrastructure.NewJWTService()
	passwordService := infrastructure.NewPasswordService()
	authMiddleware := infrastructure.NewAuthMiddleware(jwtService)

	taskRepository := repositories.NewTaskRepository(tasksCollection)
	userRepository := repositories.NewUserRepository(usersCollection)

	taskUseCase := usecases.NewTaskUseCase(taskRepository)
	userUseCase := usecases.NewUserUseCase(userRepository, passwordService, jwtService)

	taskController := controllers.NewTaskController(taskUseCase)
	userController := controllers.NewUserController(userUseCase)

	router := routers.NewRouter(taskController, userController, authMiddleware)

	app := router.SetupRoutes()

	log.Println("Server starting on port 8080...")
	if err := app.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
