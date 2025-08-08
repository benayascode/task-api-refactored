package repositories

import (
	"context"
	"errors"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositoryImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewTaskRepository(collection *mongo.Collection) domain.TaskRepository {
	return &TaskRepositoryImpl{
		collection: collection,
		ctx:        context.TODO(),
	}
}

func (t *TaskRepositoryImpl) GetAllTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := t.collection.Find(t.ctx, bson.M{})
	if err != nil {
		return tasks, err
	}
	defer cursor.Close(t.ctx)

	for cursor.Next(t.ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err == nil {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) GetTaskByID(userID int) (domain.Task, error) {
	var task domain.Task
	err := t.collection.FindOne(t.ctx, bson.M{"user_id": userID}).Decode(&task)
	if err != nil {
		return task, errors.New("task not found")
	}
	return task, nil
}

func (t *TaskRepositoryImpl) CreateTask(task domain.Task) error {
	if task.ID.IsZero() {
		task.ID = primitive.NewObjectID()
	}

	var lastTask domain.Task
	opts := options.FindOne().SetSort(bson.D{{"user_id", -1}})
	err := t.collection.FindOne(t.ctx, bson.D{}, opts).Decode(&lastTask)
	if err == nil {
		task.UserID = lastTask.UserID + 1
	} else {
		task.UserID = 1
	}

	_, err = t.collection.InsertOne(t.ctx, task)
	return err
}

func (t *TaskRepositoryImpl) UpdateTask(userID int, newTask domain.Task) (domain.Task, error) {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"title":       newTask.Title,
			"description": newTask.Description,
			"due_date":    newTask.DueDate,
			"status":      newTask.Status,
		},
	}

	result, err := t.collection.UpdateOne(t.ctx, filter, update)
	if err != nil {
		return domain.Task{}, err
	}

	if result.MatchedCount == 0 {
		return domain.Task{}, errors.New("task not found")
	}

	return newTask, nil
}

func (t *TaskRepositoryImpl) DeleteTask(userID int) error {
	result, err := t.collection.DeleteOne(t.ctx, bson.M{"user_id": userID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
