package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// showIndexPage renders the index page, passing the page title and returning status code 200
// NB gin.H{} is a shortcut for map[string]interface{} and is used to pass data to the template
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// render the index.html template
	render(c, gin.H {
		"title": "Home Page",
		"payload": articles,
	},
	"index.html")
}

// getArticle renders the individual article page
func getArticle(c *gin.Context) {
	// check that the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {

		// check if the article exists
		if article, err := getArticleByID(articleID); err == nil {

			// render the page
			render(c, gin.H {
				"title": article.Title,
				"payload": article,
			},
			"article.html")

		} else {
			// article not found
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// invalid article ID
		c.AbortWithStatus(http.StatusNotFound)
	}
}