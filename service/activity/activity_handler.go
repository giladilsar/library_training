package activity

import (
	"fmt"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActivityByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Errorf("username must be provided")})
		return
	}

	res, err := getUserActivity(username)
	if err != nil {
		c.JSON(utils.GetErrorResponseStatus(err), gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, res)
}
