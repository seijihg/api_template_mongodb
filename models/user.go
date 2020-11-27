package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model.
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Name      string             `json:"name" bson:"name"`
	Surname   string             `json:"surname" bson:"surname"`
	Password  string             `json:"password" bson:"password" validate:"required,password" `
	Dob       time.Time          `json:"dob" bson:"dob"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt, omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt, omitempty"`
}
