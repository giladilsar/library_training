package activity

import (
	"gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActivityByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.String(http.StatusBadRequest, "username must be provided.")
		return
	}

	res, err := getUserActivity(config.RedisClient, username)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
