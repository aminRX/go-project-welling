// Entrypoint
package main

import (
	"fmt"
	"go-microservices/db"
	"go-microservices/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found or failed to load")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	fmt.Println("Server running on port", port)
	err := r.Run(":" + port)
	if err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
