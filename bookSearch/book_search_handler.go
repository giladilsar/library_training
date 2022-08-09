package bookSearch

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
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if validationError := validateSearchRequest(req); validationError != nil {
		c.String(http.StatusBadRequest, validationError.Error())
		return
	}

	books, err := searchBooks(config.ElasticClient, req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}
