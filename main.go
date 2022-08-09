package main

import (
	"gin/books"
	"gin/books_search"
	"gin/config"
	"gin/health"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Setup()

	router := gin.Default()
	router.GET("/ping", health.Ping)
	// Books
	router.GET("/books", books.GetBookById)
	router.PUT("/books", books.CreateBook)
	router.POST("/books", books.UpdateBookTitle)
	// Search
	router.GET("/search", books_search.SearchBook)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
