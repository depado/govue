package main

import (
	"log"

	"github.com/Depado/govue/database"
	"github.com/Depado/govue/models/entry"
	"github.com/gin-gonic/gin"
)

const currentAPIVersion = "1"

func main() {
	var err error

	// Database initialization
	if err = database.Main.Open(); err != nil {
		log.Fatal(err)
	}
	defer database.Main.Close()

	r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	// r.Static("/static", "./assets")

	currentAPI := r.Group("/api/v" + currentAPIVersion)
	entryEndpoint := currentAPI.Group("/entry")
	{
		entryEndpoint.POST("/", entry.Post)
		// entryr.GET("/", entry.List)
		entryEndpoint.GET("/:id", entry.Get)
		entryEndpoint.PATCH("/:id", entry.Patch)
		entryEndpoint.DELETE("/:id", entry.Delete)
	}

	r.Run(":8080")
}
