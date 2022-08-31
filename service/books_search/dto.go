package books_search

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type priceRange struct {
	From int
	To   int
}

type bookSearchRequest struct {
	Title      string
	AuthorName string
	PriceRange priceRange
	Username   string
}

func (req *bookSearchRequest) containsPriceFilter() bool {
	return req.PriceRange.To > 0 || req.PriceRange.From > 0
}

func searchRequestBuilder(c *gin.Context) (*bookSearchRequest, error) {
	fromPriceStr := c.DefaultQuery("from_price", "0")
	fromPrice, err := strconv.Atoi(fromPriceStr)
	if err != nil {
		return nil, err
	}

	toPriceStr := c.DefaultQuery("to_price", "0")
	toPrice, err := strconv.Atoi(toPriceStr)
	if err != nil {
		return nil, err
	}

	req := bookSearchRequest{
		Title:      c.DefaultQuery("title", ""),
		AuthorName: c.DefaultQuery("author_name", ""),
		PriceRange: priceRange{
			From: fromPrice,
			To:   toPrice,
		},
	}

	return &req, nil
}
