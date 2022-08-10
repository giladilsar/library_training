package server

import (
	"gin/activity"
	"gin/books"
	"gin/books_search"
	"gin/config"
	"gin/health"
	"gin/server/middleware"
	"gin/store"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", health.Ping)

	cachedRequests := router.Group("/")
	cachedRequests.Use(middleware.CacheRequest(config.RedisClient))
	{ // Books
		cachedRequests.GET("/books/:id", books.GetBookById)
		cachedRequests.PUT("/books", books.CreateBook)
		cachedRequests.POST("/books/:id", books.UpdateBookTitle)
		cachedRequests.DELETE("/books/:id", books.DeleteBook)
		// Search
		cachedRequests.GET("/search", books_search.SearchBook)
		// Store
		cachedRequests.GET("/store", store.GetStore)
	}
	router.GET("/activity/:username", activity.GetActivityByUsername)

	return router
}
