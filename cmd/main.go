package main

import (
	"log"

	"example.com/go-shorter/internals/gorm"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	server := gin.Default()
	initDb()

	server.POST("/urls", ShortenUrl)
	server.GET("/:code", RedirectToTarget)

	server.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	} else {
		log.Println("A .env file was loaded")
	}
}

func initDb() {
	gorm.DbConnect()
	migrateDb()
}
