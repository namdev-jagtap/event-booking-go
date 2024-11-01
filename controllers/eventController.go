// controllers/eventController.go
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
	"event-booking/utils"
)

var eventCollection *mongo.Collection

// SetEventCollection initializes the event collection
func SetEventCollection(client *mongo.Client) {
	eventCollection = client.Database("eventdb").Collection("events")
}

// CreateEvent creates a new event
func CreateEvent(c *gin.Context) {

	event, err := utils.RetrieveValidatedData[models.Event](c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	event.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = eventCollection.InsertOne(ctx, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

// GetEvent retrieves a specific event by ID
func GetEvent(c *gin.Context) {
	eventID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(eventID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var event models.Event
	err := eventCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&event)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// GetEvents retrieves all events
func GetEvents(c *gin.Context) {
	var events []models.Event

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := eventCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var event models.Event
		if err := cursor.Decode(&event); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding event"})
			return
		}
		events = append(events, event)
	}
	fmt.Println("Events", events)
	c.JSON(http.StatusOK, events)
}

// UpdateEvent updates an event by ID
func UpdateEvent(c *gin.Context) {
	eventID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(eventID)

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"$set": event}
	_, err := eventCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

// DeleteEvent deletes an event by ID
func DeleteEvent(c *gin.Context) {
	eventID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(eventID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := eventCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
