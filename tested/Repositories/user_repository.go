package repositories

import (
	"context"
	"errors"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserRepository(collection *mongo.Collection) domain.UserRepository {
	return &UserRepositoryImpl{
		collection: collection,
		ctx:        context.TODO(),
	}
}

func (u *UserRepositoryImpl) RegisterUser(username, password string) error {

	count, err := u.collection.CountDocuments(u.ctx, bson.M{"user_name": username})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username is taken")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	user := domain.User{
		UserName: username,
		Password: password,
		Role:     "user",
	}

	_, err = u.collection.InsertOne(u.ctx, user)
	return err
}

func (u *UserRepositoryImpl) AuthenticateUser(username, password string) (domain.User, error) {
	var user domain.User
	err := u.collection.FindOne(u.ctx, bson.M{"user_name": username}).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (u *UserRepositoryImpl) PromoteUser(username string) error {

	var user domain.User
	err := u.collection.FindOne(u.ctx, bson.M{"user_name": username}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	_, err = u.collection.UpdateOne(
		u.ctx,
		bson.M{"user_name": username},
		bson.M{"$set": bson.M{"role": "Admin"}},
	)

	return err
}
