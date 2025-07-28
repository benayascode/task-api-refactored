package main

import (
	"context"
	"log"
	"task_manager/Delivery/controllers"
	routers "task_manager/Delivery/router"
	"task_manager/Repositories"
	"task_manager/Usecases"
	"task_manager/Infrastructure"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	db := client.Database("taskmanagerdb")
	taskRepo := &Repositories.MongoTaskRepository{Collection: db.Collection("tasks")}
	userRepo := &Repositories.MongoUserRepository{Collection: db.Collection("users")}
	passwordService := Infrastructure.NewBcryptPasswordService()
	jwtService := Infrastructure.NewSimpleJWTService()

	taskUC := &Usecases.TaskUseCase{TaskRepo: taskRepo}
	userUC := Usecases.NewUserUseCase(userRepo, passwordService, jwtService)

	ctr := &controllers.Controller{TaskUC: taskUC, UserUC: userUC}

	r := routers.SetupRouter(ctr)
	r.Run(":8080")
}
