package database

import (
	"log"
	"sagara-msib-test/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(driver, url string) {
	var err error
	switch driver {
	case "postgres":
		DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
		}
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(url), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to SQLite database: %v", err)
		}
	default:
		log.Fatalf("Unsupported database driver: %s", driver)
	}

	// Automatically migrate the schema
	err = DB.AutoMigrate(&models.Cloth{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}
}
