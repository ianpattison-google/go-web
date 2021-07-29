package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// store the article list in memory - ultimately this will be in a database
var articleList = []article{
	{ID: 1, Title: "T-shirt", Content: "Made from sustainable cotton"},
	{ID: 2, Title: "Hoodie", Content: "Available in blue or grey"},
}

// return a list of all articles
func getAllArticles() []article {
	return articleList
}

// getArticleByID returns the article with the supplied ID
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}