package category

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	ParentId  *primitive.ObjectID `bson:"parent_id" json:"parentId"`
	CreatedBy primitive.ObjectID  `bson:"created_by" json:"createdBy"`
	Title     string              `bson:"title" json:"title"`
}
