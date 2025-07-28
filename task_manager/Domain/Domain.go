package Domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      int                `bson:"user_id" json:"user_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	Status      string             `bson:"status" json:"status"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName string             `bson:"user_name" json:"user_name"`
	Password string             `bson:"password,omitempty" json:"_"`
	Role     string             `bson:"role" json:"role"`
}

type TaskRepository interface {
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTaskByID(ctx context.Context, id int) (Task, error)
	CreateTask(ctx context.Context, task Task) error
	UpdateTask(ctx context.Context, id int, updated Task) error
	DeleteTask(ctx context.Context, id int) error
}
type UserRepository interface {
	Register(ctx context.Context, user User) error
	Authenticate(ctx context.Context, username string) (User, error)
	Promote(ctx context.Context, username string) error
}
type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(hash, password string) bool
}
type JWTService interface {
	GenerateToken(user User) (string, error)
}
