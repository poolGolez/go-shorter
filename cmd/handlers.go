package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenUrl(ctx *gin.Context) {
	// TODO: Stub for now
	ctx.Header("Location", "http://localhost:8080/changeM3now")
	ctx.Status(http.StatusCreated)
}

func RedirectToTarget(ctx *gin.Context) {
	// TODO: Stub for now
	ctx.Redirect(http.StatusTemporaryRedirect, "http://example.com")
}
