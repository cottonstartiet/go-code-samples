package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAblumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
