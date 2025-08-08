package domain

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidTaskTitle       = errors.New("task title cannot be empty")
	ErrInvalidTaskDescription = errors.New("task description cannot be empty")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrUserNotFound           = errors.New("user not found")
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      int                `bson:"user_id" json:"user_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	Status      string             `bson:"status" json:"status"`
}

func (t Task) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return ErrInvalidTaskTitle
	}
	if strings.TrimSpace(t.Description) == "" {
		return ErrInvalidTaskDescription
	}
	return nil
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName string             `bson:"user_name" json:"user_name"`
	Password string             `bson:"password,omitempty" json:"_"`
	Role     string             `bson:"role" json:"role"`
}

func (u User) IsAdmin() bool {
	return u.Role == "Admin"
}

type TaskRepository interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(userID int) (Task, error)
	CreateTask(task Task) error
	UpdateTask(userID int, task Task) (Task, error)
	DeleteTask(userID int) error
}

type UserRepository interface {
	RegisterUser(username, password string) error
	AuthenticateUser(username, password string) (User, error)
	PromoteUser(username string) error
}

type TaskUseCase interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(userID int) (Task, error)
	CreateTask(task Task) error
	UpdateTask(userID int, task Task) (Task, error)
	DeleteTask(userID int) error
}

type UserUseCase interface {
	RegisterUser(username, password string) error
	LoginUser(username, password string) (string, error)
	PromoteUser(username string) error
}
