package main

import (
	"log"
	"net/http"

	"github.com/Depado/govue/database"
	"github.com/Depado/govue/models/entry"
	"github.com/gin-gonic/gin"
)

const currentAPIVersion = "1"

func index(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "templates/index.html")
}

func main() {
	var err error

	// Database initialization
	if err = database.Main.Open(); err != nil {
		log.Fatal(err)
	}
	defer database.Main.Close()

	r := gin.Default()
	r.Static("/static", "./assets")
	r.Static("/vendor", "./node_modules")

	r.GET("/", index)

	currentAPI := r.Group("/api/v" + currentAPIVersion)
	entryEndpoint := currentAPI.Group("/entry")
	{
		entryEndpoint.POST("/", entry.Post)
		entryEndpoint.GET("/", entry.List)
		entryEndpoint.GET("/:id", entry.Get)
		entryEndpoint.PATCH("/:id", entry.Patch)
		entryEndpoint.DELETE("/:id", entry.Delete)
	}

	r.Run(":8080")
}
