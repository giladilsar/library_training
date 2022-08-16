package store

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStoreInfo(c *gin.Context) {
	res, err := fetchStoreDate()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
