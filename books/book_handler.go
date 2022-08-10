package books

import (
	"fmt"
	"gin/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("id must be provided"))
		return
	}
	book, err := getBook(config.ElasticClient, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("id must be provided"))
		return
	}
	err := deleteBookById(config.ElasticClient, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func CreateBook(c *gin.Context) {
	req := &createBookRequest{}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to bind request - invalid request. %s", err.Error()))
		return
	}

	response, indexErr := createBookFromPayload(config.ElasticClient, req)
	if indexErr != nil {
		c.AbortWithError(http.StatusInternalServerError, indexErr)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateBookTitle(c *gin.Context) {
	id := c.Param("id")
	req := &updateBookRequest{Id: id}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to bind request - invalid request. %s", err.Error()))
		return
	}

	err := updateBook(config.ElasticClient, req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "OK")
}
