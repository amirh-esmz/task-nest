package task

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	TaskCollection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{TaskCollection: collection}
}

// Create a new task
func (s *TaskService) CreateTask(ctx context.Context, task Task) (*mongo.InsertOneResult, error) {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now().Unix()
	return s.TaskCollection.InsertOne(ctx, task)
}

// Get a task by ID
func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var task Task
	err = s.TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	return &task, err
}

// Update a task
func (s *TaskService) UpdateTask(ctx context.Context, id string, updatedTask Task) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedTask}
	return s.TaskCollection.UpdateOne(ctx, filter, update)
}

// Delete a task
func (s *TaskService) DeleteTask(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
}
