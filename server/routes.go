package server

import (
	"gin/config"
	"gin/server/middleware"
	"gin/service/activity"
	"gin/service/books"
	"gin/service/books_search"
	"gin/service/health"
	"gin/service/store"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", health.Ping)

	cachedRequests := router.Group("/")
	cachedRequests.Use(middleware.CacheRequest(config.RedisClient))
	{
		cachedRequests.GET("/book/:id", books.GetBook)
		cachedRequests.PUT("/book", books.CreateBook)
		cachedRequests.POST("/book/:id", books.UpdateBookTitle)
		cachedRequests.DELETE("/book/:id", books.DeleteBook)

		cachedRequests.GET("/search", books_search.SearchBook)

		cachedRequests.GET("/store", store.GetStoreInfo)
	}
	router.GET("/activity/:username", activity.GetActivityByUsername)

	return router
}
