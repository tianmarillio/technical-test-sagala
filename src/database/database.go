package database

import (
	"fmt"
	"log"

	"github.com/tianmarillio/technical-test-sagala/src/config"
	"github.com/tianmarillio/technical-test-sagala/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (db *gorm.DB) {
	DB, err := gorm.Open(postgres.Open(cfg.DatabaseConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	fmt.Println("Connected to database")

	// Create enums
	DB.Exec("CREATE TYPE task_status AS ENUM ('waiting_list', 'in_progress', 'done')")

	// Database migration
	// Note: used auto migrate for simplicity purpose
	// In production or real cases, it's recommended to use SQL or migration tools
	// 	like golang-migrate or goose
	//  to avoid data loss and increase overall migration control
	if err = DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("Migration error")
	}

	fmt.Println("Database migration success")

	return DB
}
