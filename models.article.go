package main

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 content"},
	{ID: 2, Title: "Article 2", Content: "Article 2 content"},
}

func getAllArticles() []Article {
	return articleList
}

func getArticleByID(id int) (*Article, error) {
	// loop through the list of articles, looking for an article with the given id
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
