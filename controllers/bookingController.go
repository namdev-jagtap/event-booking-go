// controllers/bookingController.go
package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"event-booking/models"
)

var bookingCollection *mongo.Collection

// SetBookingCollection initializes the booking collection
func SetBookingCollection(client *mongo.Client) {
	bookingCollection = client.Database("eventdb").Collection("bookings")
}

// BookEvent creates a new booking for an event
func BookEvent(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the event exists
	eventID := booking.EventID
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var event models.Event
	err := eventCollection.FindOne(ctx, bson.M{"_id": eventID}).Decode(&event)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Check if there is enough capacity
	if booking.Quantity > event.Capacity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough capacity for booking"})
		return
	}

	// Reduce capacity in the event (you may want to use transactions for a production setup)
	event.Capacity -= booking.Quantity
	_, err = eventCollection.UpdateOne(ctx, bson.M{"_id": eventID}, bson.M{"$set": bson.M{"capacity": event.Capacity}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event capacity"})
		return
	}

	// Create the booking
	booking.ID = primitive.NewObjectID()
	_, err = bookingCollection.InsertOne(ctx, booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

// GetBookings retrieves all bookings for a specific event
func GetBookings(c *gin.Context) {
	eventID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(eventID)
	fmt.Println(objID)

	var bookings []models.Booking
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := bookingCollection.Find(ctx, bson.M{"event_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var booking models.Booking
		if err := cursor.Decode(&booking); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding booking"})
			return
		}
		bookings = append(bookings, booking)
	}

	c.JSON(http.StatusOK, bookings)
}
