// Lógica de conexión a DB 
package db

import (
	"fmt"
	"log"
	"os"

	"go-microservices/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL") // e.g. "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.AutoMigrate(&models.Organization{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}

	DB = db
	fmt.Println("Database connected and migrated")
}
