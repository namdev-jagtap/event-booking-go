// models/event.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Event represents an event
type Event struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required" bson:"name"`
	Location    string             `json:"location" validate:"required" bson:"location"`
	StartDate   string             `json:"start_date" validate:"required" bson:"start_date"`
	EndDate     string             `json:"end_date" validate:"required" bson:"end_date"`
	Description string             `json:"description" validate:"required" bson:"description"`
	Capacity    int                `json:"capacity" validate:"required" bson:"capacity"`
}
