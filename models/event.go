// models/event.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Event represents an event
type Event struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Location    string             `json:"location" bson:"location"`
	StartDate   string             `json:"start_date" bson:"start_date"`
	EndDate     string             `json:"end_date" bson:"end_date"`
	Description string             `json:"description" bson:"description"`
	Capacity    int                `json:"capacity" bson:"capacity"`
}
