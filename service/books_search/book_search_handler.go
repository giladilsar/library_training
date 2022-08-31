package books_search

import (
	"fmt"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateSearchRequest(req *bookSearchRequest) error {
	if req.Title == "" && req.AuthorName == "" && !(req.containsPriceFilter()) {
		return fmt.Errorf("request must contain at least one filter")
	}

	return nil
}

func SearchBook(c *gin.Context) {
	req, err := searchRequestBuilder(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if validationError := validateSearchRequest(req); validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": validationError.Error()})
		return
	}

	books, err := searchBooks(req)
	if err != nil {
		c.JSON(utils.GetErrorResponseStatus(err), gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
