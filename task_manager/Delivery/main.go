package main

import (
	"context"
	"log"
	"task_manager/Delivery/controllers"
	routers "task_manager/Delivery/router"
	"task_manager/Repositories"
	"task_manager/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	db := client.Database("taskmanagerdb")
	taskRepo := &Repositories.TaskRepository{Collection: db.Collection("tasks")}
	userRepo := &Repositories.UserRepository{Collection: db.Collection("users")}

	taskUC := &Usecases.TaskUseCase{TaskRepo: taskRepo}
	userUC := &Usecases.UserUseCase{UserRepo: userRepo}

	ctr := &controllers.Controller{TaskUC: taskUC, UserUC: userUC}

	r := routers.SetupRouter(ctr)
	r.Run(":8080")
}
