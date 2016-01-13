package main

import (
	"log"

	"github.com/Depado/govue/database"
	"github.com/Depado/govue/models/entry"
	"github.com/gin-gonic/gin"
)

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

	entryr := r.Group("/entry")
	{
		entryr.POST("/", entry.Post)
		// entryr.GET("/", entry.List)
		// entryr.GET("/:id", entry.Get)
		// entryr.PATCH("/:id", entry.Patch)
		// entryr.PUT("/:id", entry.Put)
		// entryr.DELETE("/:id", entry.Delete)
	}

	r.Run(":8080")
}
