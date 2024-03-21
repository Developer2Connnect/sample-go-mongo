// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/developer2connect/sample-go-mongo/db"
	"github.com/developer2connect/sample-go-mongo/handlers"
	"github.com/developer2connect/sample-go-mongo/repository"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR environment variable is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Initialize MongoDB connection
	db.InitMongo(mongoURI)

	// Initialize Redis connection
	db.InitRedis(redisAddr)

	// Initialize repository
	productRepo := repository.NewProductRepository("products")

	// Initialize handlers with repository
	router := handlers.NewRouter(productRepo)

	// Start server
	fmt.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
