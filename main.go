package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create new type called album which is extended from the struct type, or if you did: var album struct - then that would create a new variable 'album' with it's type equal to anonymous struct.
//
//	   ID is an example of a name of a field within the struct declaration.
//		 `json:"id"` is an example of the 'tag' that is part of the struct field declaration.
//			https://go.dev/talks/2015/tricks.slide#6
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.97},
	{ID: "2", Title: "Jeru", Artist: "Geri Mulligan", Price: 17.99},
	{ID: "3", Title: "Songs of the Past", Artist: "Miles Davis", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	for _, albumRecord := range albums {
		if c.Param("id") == albumRecord.ID {
			c.JSON(http.StatusOK, albumRecord)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": 9})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8081")

}

// To invoke the API for example, localhost:8081/albums
