package task

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Category    string             `bson:"category" json:"category"`
	Location    string             `bson:"location" json:"location"`
	PostedBy    primitive.ObjectID `bson:"posted_by" json:"posted_by"`
	CreatedAt   int64              `bson:"created_at" json:"created_at"`
}
