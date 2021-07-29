package main

func initRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}


