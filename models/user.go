package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Name     string             `json:"name" bson:"name"`
	Surname  string             `json:"surname" bson:"surname"`
	Password string             `json:"password" bson:"password" validate:"required,password" `
	Dob      primitive.DateTime `json:"dob" bson:"dob"`
}
