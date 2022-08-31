package books

import (
	"fmt"
	"gin/service/books/dto"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := getBook(id)
	if err != nil {
		c.JSON(utils.GetErrorResponseStatus(err), gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Errorf("id must be provided\n")})
		return
	}
	err := deleteBook(id)
	if err != nil {
		c.JSON(utils.GetErrorResponseStatus(err), gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func CreateBook(c *gin.Context) {
	req := &dto.CreateBookRequest{}
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Errorf("failed to bind request - invalid request. %s", err.Error())})
		return
	}

	id, indexErr := createBookFromPayload(req)
	if indexErr != nil {
		c.JSON(utils.GetErrorResponseStatus(indexErr), gin.H{"Error": indexErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Id": id})
}

func UpdateBookTitle(c *gin.Context) {
	id := c.Param("id")
	req := &dto.UpdateBookRequest{Id: id}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Errorf("failed to bind request - invalid request. %s", err.Error())})
		return
	}

	err := updateBook(req)
	if err != nil {
		c.JSON(utils.GetErrorResponseStatus(err), gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "OK")
}
