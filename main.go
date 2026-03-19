package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Calloway", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Essence", Artist: "Wizkid ft Tems", Price: 19.99},
	{ID: "5", Title: "Fall", Artist: "Davido", Price: 19.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func createAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"album": a, "status": http.StatusOK})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found", "status": http.StatusNotFound})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbum)
	router.GET("/albums/:id", getAlbumByID)
	fmt.Println("Hello, World!")
	router.Run("localhost:8080")
}
