package books_search

import (
	"fmt"
	"gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateSearchRequest(req *bookSearchRequest) error {
	if req.Title == "" && req.Name == "" && !(req.containsPriceFilter()) {
		return fmt.Errorf("request must contain at least one filter")
	}

	return nil
}

func SearchBook(c *gin.Context) {
	req, err := searchRequestBuilder(c)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if validationError := validateSearchRequest(req); validationError != nil {
		c.AbortWithError(http.StatusBadRequest, validationError)
		return
	}

	books, err := searchBooks(config.ElasticClient, req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}
