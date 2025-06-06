package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	InitDb()

	server.POST("/urls", ShortenUrl)
	server.GET("/:code", RedirectToTarget)

	server.Run()
}
