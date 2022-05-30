package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// define a router
	router := gin.Default()
	// define a route
	router.GET("/", index)

	// run the server
	router.Run("localhost:8080")
}

// callback for the '/' route
func index(c *gin.Context) {
	// return JSON to the client
	c.JSON(http.StatusOK, gin.H{"message": "You made it!"})
}
