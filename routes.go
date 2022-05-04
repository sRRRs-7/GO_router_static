package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/:article_id", getArticle)
}
