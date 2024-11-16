package category

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService struct {
	CategoryCollection *mongo.Collection
}

func NewCategoryService(collection *mongo.Collection) *CategoryService {
	return &CategoryService{CategoryCollection: collection}
}

// Create a new task
func (s *CategoryService) Create(ctx context.Context, category Category) (*mongo.InsertOneResult, error) {
	category.Id = primitive.NewObjectID()
	return s.CategoryCollection.InsertOne(ctx, category)
}

// Get a task by ID
func (s *CategoryService) GetById(ctx context.Context, id string) (*Category, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var category Category
	err = s.CategoryCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&category)
	return &category, err
}
