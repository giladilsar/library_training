package activity

import (
	"fmt"
	"gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActivityByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("username must be provided"))
		return
	}

	res, err := getUserActivity(config.RedisClient, username)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
