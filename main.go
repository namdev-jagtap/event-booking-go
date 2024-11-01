package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"event-booking/config"
	"event-booking/controllers"
	"event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = "mongodb://localhost:27017"
	}

	client, err := config.ConnectDB(mongoUri)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

	defer client.Disconnect(context.TODO())

	// Initialize collections for controllers
	controllers.SetUserCollection(client)
	controllers.SetEventCollection(client)
	controllers.SetBookingCollection(client)

	router := gin.Default()
	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(router.Run(":" + port))
}
