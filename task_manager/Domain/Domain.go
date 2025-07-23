package Domain

import (
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
