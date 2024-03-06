package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiErrorResponse struct {
	message string
}

type Album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

var albums = []Album{
	{Id: "id1", Title: "Blue Train", Artist: "John Coltrane", Name: "album 1"},
	{Id: "id2", Title: "Jeru", Artist: "Gerry Mulligan", Name: "album 2"},
	{Id: "id3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Name: "album 3"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func getAblumById(c *gin.Context) {
	id := c.Param("id")

	for _, alb := range albums {
		if alb.Id == id {
			c.JSON(http.StatusOK, alb)
			return
		}
	}

	c.JSON(http.StatusNotFound, ApiErrorResponse{
		message: "not found",
	})
}
