// controllers/userController.go
package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"event-booking/models"
	"event-booking/utils"
)

var userCollection *mongo.Collection

// SetUserCollection initializes the user collection
func SetUserCollection(client *mongo.Client) {
	userCollection = client.Database("eventdb").Collection("users")
}

// RegisterUser registers a new user
func RegisterUser(c *gin.Context) {
	// Retrieve the validated data from the context
	user, err := utils.RetrieveValidatedData[models.User](c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert the user into the collection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err = userCollection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}).Decode(&existingUser)

	if err == nil {
		// User with the same username or email already exists
		c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
		return
	}
	if err != mongo.ErrNoDocuments {
		// Some other error occurred
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Return a success message upon successful registration
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser authenticates the user and returns a JWT token
func LoginUser(c *gin.Context) {
	// Retrieve the validated data from context (plain text credentials)
	credentials, err := utils.RetrieveValidatedData[models.User](c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user from the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch the user from the database using the username
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Compare the plain text password from the request with the hashed password from the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// Return the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
