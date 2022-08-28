package books_search

import (
	"gin/service/books/dto"
	"github.com/olivere/elastic/v7"
)

func buildQuery(req *bookSearchRequest) dto.BookSearchQuery {
	query := elastic.NewBoolQuery()
	if req.containsPriceFilter() {
		priceRangeQuery := elastic.NewRangeQuery("price")
		if req.PriceRange.From != 0 {
			priceRangeQuery = priceRangeQuery.Gte(req.PriceRange.From)
		}
		if req.PriceRange.To != 0 {
			priceRangeQuery = priceRangeQuery.Lte(req.PriceRange.To)
		}
		query = query.Must(priceRangeQuery)
	}
	if req.Title != "" {
		query = query.Must(elastic.NewTermQuery("title", req.Title))
	}
	if req.AuthorName != "" {
		query = query.Must(elastic.NewMatchQuery("author_name", req.AuthorName))
	}

	return query
}
