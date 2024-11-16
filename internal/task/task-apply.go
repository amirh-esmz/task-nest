package task

import "go.mongodb.org/mongo-driver/bson/primitive"

type TaskApply struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TaskId    primitive.ObjectID `bson:"task_id" json:"taskId"`
	UserId    primitive.ObjectID `bson:"user_id" json:"userId"`
	CreatedAt int64              `bson:"created_at" json:"created_at"`
}
