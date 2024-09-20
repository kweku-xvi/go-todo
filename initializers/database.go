package initializers

import (
	"log"
	"os"

	"github.com/kweku-xvi/todolist-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database")
	}

	// Running migrations on DB
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("failed to run database migrations")
	}
}
