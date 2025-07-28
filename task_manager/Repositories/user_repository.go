package Repositories

import (
	"context"
	"errors"
	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func (ur *MongoUserRepository) Register(ctx context.Context, user Domain.User) error {
	count, _ := ur.Collection.CountDocuments(ctx, bson.M{"user_name": user.UserName})
	if count > 0 {
		return errors.New("username is already taken")
	}

	_, err := ur.Collection.InsertOne(ctx, user)
	return err
}

func (ur *MongoUserRepository) Authenticate(ctx context.Context, username string) (Domain.User, error) {
	var user Domain.User
	err := ur.Collection.FindOne(ctx, bson.M{"user_name": username}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (ur *MongoUserRepository) Promote(ctx context.Context, username string) error {
	res, err := ur.Collection.UpdateOne(ctx, bson.M{"user_name": username}, bson.M{"$set": bson.M{"role": "Admin"}})
	if err != nil || res.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
