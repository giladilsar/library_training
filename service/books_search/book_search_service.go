package books_search

import (
	"encoding/json"
	"gin/models"
	"gin/utils"
	"github.com/olivere/elastic/v7"
)

const IndexName = "gilad_books"

func searchBooks(es *elastic.Client, req *bookSearchRequest) ([]models.Book, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	query := elastic.NewBoolQuery()
	if req.containsPriceFilter() {
		priceRangeQuery := elastic.NewRangeQuery("price").Gte(req.PriceRange.From).Lte(req.PriceRange.To)
		query = query.Must(priceRangeQuery)
	}
	if req.Title != "" {
		query = query.Must(elastic.NewTermQuery("title", req.Title))
	}
	if req.AuthorName != "" {
		query = query.Must(elastic.NewMatchQuery("name", req.AuthorName))
	}

	searchResult, err := es.Search().
		Index(IndexName).
		Query(query).
		From(0).
		Size(100).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	books := make([]models.Book, searchResult.Hits.TotalHits.Value)
	for i, hit := range searchResult.Hits.Hits {
		book := models.Book{Id: hit.Id}
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			return nil, err
		}
		books[i] = book
	}

	return books, nil
}
