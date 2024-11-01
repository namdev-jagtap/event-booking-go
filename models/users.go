// models/user.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user in the database
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	Username string             `json:"username" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
