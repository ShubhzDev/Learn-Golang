package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Title1", Artist: "Artist1", Price: 56},
	{ID: "2", Title: "Title2", Artist: "Artist2", Price: 72},
	{ID: "3", Title: "Title3", Artist: "Artist3", Price: 90},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(context *gin.Context) {
	id := context.Param("id")

	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusFound, album)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}
