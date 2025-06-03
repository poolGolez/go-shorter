package main

import (
	"net/http"

	shorturl "example.com/go-shorter/internals/short_url"
	"github.com/gin-gonic/gin"
)

func ShortenUrl(ctx *gin.Context) {
	var requestBody ShortenUrlRequestBody
	ctx.BindJSON(&requestBody)

	shortUrl, err := shorturl.CreateShortenedUrl(requestBody.Url)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Header("Location", "http://localhost:8080/urls/"+shortUrl.Code)
	ctx.Status(http.StatusCreated)
}

type ShortenUrlRequestBody struct {
	Url string `binding: "required"`
}

func RedirectToTarget(ctx *gin.Context) {
	code := ctx.Param("code")
	shortUrl := shorturl.FindByCode(code)
	ctx.Redirect(http.StatusTemporaryRedirect, shortUrl.TargetUrl)
}
