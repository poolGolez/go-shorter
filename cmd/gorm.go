package main

import (
	"fmt"
	"log"
	"os"

	shorturl "example.com/go-shorter/internals/short_url"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println("=====-----------------------------------------------------------------------------------=====")
	fmt.Println(dsn)
	fmt.Println("=====-----------------------------------------------------------------------------------=====")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connection successful!!")

	fmt.Println("Migrating database tables...")
	MigrateModels(db)
	fmt.Println("Migration successful...")
}

func MigrateModels(db *gorm.DB) {
	models := []interface{}{
		&shorturl.ShortUrl{},
	}

	db.AutoMigrate(models...)
}
