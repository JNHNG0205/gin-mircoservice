package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
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

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
