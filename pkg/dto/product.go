package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person represents a person document in MongoDB
type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	CreatedAt   primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
