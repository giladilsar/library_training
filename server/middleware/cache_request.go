package middleware

import (
	"encoding/json"
	"gin/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v5"
)

func CacheRequest(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := models.UserRequest{
			Method: c.Request.Method,
			Route:  c.Request.URL.Path,
		}

		username, ok := c.GetQuery("username")
		if !ok {
			return
		}

		reqJson, err := json.Marshal(req)
		if err != nil {
			return
		}

		cmd := redisClient.LPush(username, reqJson)
		if cmd.Err() != nil {
			return
		}

		err = redisClient.LTrim(username, 0, 2).Err()
		if err != nil {
			return
		}
	}
}
