package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{"title": "Home Page",},
		)

	})
	router.GET("/articles", func(c *gin.Context) {
		c.HTML(http.StatusOK, "articles.html", gin.H{"title": "Articles",},)
	})
	router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gin.H{"title": "Projects",},)
	})
	router.GET("/gettingstarted", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gettingstarted.html", gin.H{"title": "Getting Started",},)
	})
	router.GET("/documentation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "documentation.html", gin.H{"title": "Documentation",},)
	})
	router.GET("/pricing", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pricing.html", gin.H{"title": "Pricing",},)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login",},)
	})
	// Start service the application
	router.Run()
}