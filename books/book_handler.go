package books

import (
	"gin/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func getBookValidator(id string, username string) bool {
	return id != "" && username != ""
}

func GetBookById(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	username := c.DefaultQuery("username", "")
	if !getBookValidator(id, username) {
		c.String(http.StatusBadRequest, "id and username must be provided.")
		return
	}
	book, err := getBook(config.ElasticClient, id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	req := &createBookRequest{}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.String(http.StatusBadRequest, "failed to bind request - invalid request. %s", err.Error())
		return
	}

	response, indexErr := createBookFromPayload(config.ElasticClient, req)
	if indexErr != nil {
		c.String(http.StatusInternalServerError, indexErr.Error())
		return
	}

	c.JSON(http.StatusOK, response)

}
