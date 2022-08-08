package main

import (
	"gin/books"
	"gin/config"
	"gin/health"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Setup()

	router := gin.Default()
	router.GET("/ping", health.Ping)
	// Bool
	router.GET("/books", books.GetBookById)
	router.PUT("/books", books.CreateBook)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
