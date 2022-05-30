package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	// guarantee that the connection to MongoDB 
	// will be properly unplugged.
	defer mgH.Disconnect(ctx)

	// define a router
	router := gin.Default()
	// define a route
	router.GET("/", index)
	router.GET("/pingdb", pingDB)

	// run the server
	router.Run("localhost:8080")
}

// callback for the GET '/' route
func index(c *gin.Context) {
	// return JSON to the client
	c.JSON(http.StatusOK, gin.H{"message": "You made it!"})
}

// callback for the GET '/pingdb' route
func pingDB(c *gin.Context) {
	ctx, cancelTO := context.WithTimeout(context.Background(), 5*time.Second)
	// test if we have a connection to the MongoDB service
	err := mgH.Ping(ctx, readpref.Primary())
	if err != nil {
		cancelTO()
		log.Fatal(err)
	}
	cancelTO()
	// We made it this far, so ... Woohooo!
	c.JSON(http.StatusOK, gin.H{"message": "We could ping to MongoDB!"})
}
