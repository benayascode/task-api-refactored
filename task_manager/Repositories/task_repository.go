package Repositories

import (
	"context"
	"errors"

	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTaskRepository struct {
	Collection *mongo.Collection
}

func (tr *MongoTaskRepository) GetAllTasks(ctx context.Context) ([]Domain.Task, error) {
	var tasks []Domain.Task

	cursor, err := tr.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var t Domain.Task
		if err := cursor.Decode(&t); err == nil {
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

func (tr *MongoTaskRepository) GetTaskByID(ctx context.Context, id int) (Domain.Task, error) {
	var task Domain.Task
	err := tr.Collection.FindOne(ctx, bson.M{"user_id": id}).Decode(&task)
	if err != nil {
		return task, errors.New("task not found")
	}
	return task, nil
}

func (tr *MongoTaskRepository) CreateTask(ctx context.Context, task Domain.Task) error {
	if task.ID.IsZero() {
		task.ID = primitive.NewObjectID()
	}

	var lastTask Domain.Task
	opts := options.FindOne().SetSort(bson.D{{"user_id", -1}})
	err := tr.Collection.FindOne(ctx, bson.D{}, opts).Decode(&lastTask)
	if err == nil {
		task.UserID = lastTask.UserID + 1
	} else {
		task.UserID = 1
	}

	_, err = tr.Collection.InsertOne(ctx, task)
	return err
}

func (tr *MongoTaskRepository) UpdateTask(ctx context.Context, id int, updated Domain.Task) error {
	filter := bson.M{"user_id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"due_date":    updated.DueDate,
			"status":      updated.Status,
		},
	}

	res, err := tr.Collection.UpdateOne(ctx, filter, update)
	if err != nil || res.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (tr *MongoTaskRepository) DeleteTask(ctx context.Context, id int) error {
	res, err := tr.Collection.DeleteOne(ctx, bson.M{"user_id": id})
	if err != nil || res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
