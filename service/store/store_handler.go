package store

import (
	"gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStore(c *gin.Context) {
	res, err := fetchStoreDate(config.ElasticClient)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
