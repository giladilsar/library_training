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

type bookSearchRequest struct {
	Title      string
	Name       string
	PriceRange priceRange
	Username   string
}

func (req *bookSearchRequest) containsPriceFilter() bool {
	return req.PriceRange.To < math.MaxInt || req.PriceRange.From > 0
}

func searchRequestBuilder(c *gin.Context) (*bookSearchRequest, error) {
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

	req := bookSearchRequest{
		Title: c.DefaultQuery("title", ""),
		Name:  c.DefaultQuery("author_name", ""),
		PriceRange: priceRange{
			From: fromPrice,
			To:   toPrice,
		},
	}

	return &req, nil
}
