package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// set the router to the default provided by Gin
	router = gin.Default()

	// process the templates here so they don't need to be loaded from disk again
	// this makes serving HTML pages very fast
	router.LoadHTMLGlob("templates/*")

	// initialise the routes
	initRoutes()

	// start the server
	router.Run()
}

// render will output HTML, JSON or XML based on the 'Accept' header of the request
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
