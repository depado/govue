package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Depado/govue/conf"
	"github.com/Depado/govue/database"
	"github.com/Depado/govue/models/entry"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "templates/index.html")
}

func main() {
	var err error

	if err = conf.Load("conf.yml"); err != nil {
		log.Fatal(err)
	}

	// Database initialization
	if err = database.Main.Open(); err != nil {
		log.Fatal(err)
	}
	defer database.Main.Close()

	r := gin.Default()
	if !conf.C.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Static("/static", "./assets")
	r.Static("/vendor", "./node_modules")

	r.GET("/", index)

	currentAPI := r.Group("/api/v" + strconv.Itoa(conf.C.APIVersion))
	entryEndpoint := currentAPI.Group("/entry")
	{
		entryEndpoint.POST("/", entry.Post)
		entryEndpoint.GET("/", entry.List)
		entryEndpoint.GET("/:id", entry.Get)
		entryEndpoint.PATCH("/:id", entry.Patch)
		entryEndpoint.DELETE("/:id", entry.Delete)
	}

	r.Run(fmt.Sprintf(":%s", strconv.Itoa(conf.C.Port)))
}
