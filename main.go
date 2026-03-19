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

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	fmt.Println("Hello, World!")
	router.Run("localhost:8080")
}
