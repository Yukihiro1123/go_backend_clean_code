package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


const (
	CollectionTask = "taks"
)

type Task struct {
	ID primitive.ObjectID `bson:"_id" json:"-"`
	Title string `bson:"title" form:"title" binding:"required" json:"title"`
	UserID primitive.ObjectID `bson:"UserID" json:"-"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}