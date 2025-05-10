// Entrypoint 
package main

import (
	"fmt"
	"go-microservices/db"
	"go-microservices/routes"
	"os"

	"github.com/gin-gonic/gin"
)

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
