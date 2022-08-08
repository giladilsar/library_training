package main

import (
	"gin/bookSearch"
	"gin/books"
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
	router.GET("/bookSearch", bookSearch.SearchBook)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
