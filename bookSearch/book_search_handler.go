package bookSearch

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func validateSearchRequest(req *searchRequest) error {
	if req.Title == "" && req.Name == "" && req.PriceRange.From == 0 && req.PriceRange.To == math.MaxInt {
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

	c.JSON(http.StatusOK, "OK")
}
