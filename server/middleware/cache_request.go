package middleware

import (
	"gin/models"
	"gin/repository/activity_repository"
	"github.com/gin-gonic/gin"
)

func CacheRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := models.UserRequest{
			Method: c.Request.Method,
			Route:  c.Request.URL.Path,
		}

		username, ok := c.GetQuery("username")
		if !ok {
			return
		}

		err := activity_repository.GetActivityRepository().SetUserActivity(username, req)
		if err != nil {
			c.Error(err)
		}
	}
}
