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
	ctx, delCtx := context.WithTimeout(context.Background(), 5*time.Second)
	// guarantee that the connection to MongoDB
	// will be properly unplugged.
	defer mgH.Disconnect(ctx)
	defer delCtx()

	// define a router
	router := gin.Default()
	// define a route
	router.GET("/", index)
	router.GET("/pingdb", pingDB)
	router.GET("/allpodcasts", getAllPodcasts)
	router.GET("/sherlock", getSherlockStuff)

	router.POST("/addPodcast", addPodcast)

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

// callback for the GET '/allpodcasts' route
func getAllPodcasts(c *gin.Context) {
	thePodCasts, err := dbGetAllPodcasts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error. Sorry."})
	}

	c.JSON(http.StatusOK, thePodCasts)
}

func addPodcast(c *gin.Context) {
	var newPodcast Podcast

	err := c.BindJSON(&newPodcast)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not handle request body!"})
		return
	}

	newPodcastID, err := dbAddPodcast(&newPodcast)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not write to the database."})
	}

	newPodcast.ID = newPodcastID

	c.JSON(http.StatusOK,
		gin.H{
			"message": "successfully inserted.",
			"body":    newPodcast,
		})
}

func getSherlockStuff(c *gin.Context) {
	sherlockEpisodes, err := dbGetSherlockEpisodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not execute this"})
	}
	if len(sherlockEpisodes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No items found!"})
	}
	c.JSON(http.StatusOK, sherlockEpisodes)
}
