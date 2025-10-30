package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// call the HTML method of the Context to render a template
	c.HTML(
		// set the status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// pass the data that the page needs
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			c.HTML(
				// set the status to 200 (OK)
				http.StatusOK,
				// Use the article.html template
				"article.html",
				// pass the data that the page needs
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// If the article is not found, abort with the error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If the article ID is not valid, abort with the error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
