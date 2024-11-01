// models/booking.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Booking represents a booking for an event
type Booking struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" validate:"required" bson:"user_id"`
	EventID  primitive.ObjectID `json:"event_id" validate:"required" bson:"event_id"`
	Quantity int                `json:"quantity" validate:"required" bson:"quantity"`
	Date     string             `json:"date" validate:"required" bson:"date"`
}
