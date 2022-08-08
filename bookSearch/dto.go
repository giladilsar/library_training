package bookSearch

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

type priceRange struct {
	From int
	To   int
}

type searchRequest struct {
	Title      string
	Name       string
	PriceRange priceRange
	Username   string
}

func searchRequestBuilder(c *gin.Context) (*searchRequest, error) {
	fromPriceStr := c.DefaultQuery("from_price", "0")
	fromPrice, err := strconv.Atoi(fromPriceStr)
	if err != nil {
		return nil, err
	}

	toPriceStr := c.DefaultQuery("to_price", strconv.Itoa(math.MaxInt))
	toPrice, err := strconv.Atoi(toPriceStr)
	if err != nil {
		return nil, err
	}

	req := searchRequest{
		Title: c.DefaultQuery("title", ""),
		Name:  c.DefaultQuery("author_name", ""),
		PriceRange: priceRange{
			From: fromPrice,
			To:   toPrice,
		},
	}

	return &req, nil
}
