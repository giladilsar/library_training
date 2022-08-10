package store

import (
	"gin/context_helper"
	"github.com/olivere/elastic/v7"
)

func fetchStoreDate(es *elastic.Client) (*fetchStoreResponse, error) {
	ctx, cancel := context_helper.GetContext()
	defer cancel()

	countResponse, err := es.Count().Index("gilad_books").Do(ctx)
	if err != nil {
		return nil, err
	}

	cardRes, err := es.Search().
		Index("gilad_books").
		Aggregation("authors",
			elastic.NewCardinalityAggregation().
				Field("name.keyword")).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	cardinalityValue, _ := cardRes.Aggregations.Cardinality("authors")
	return &fetchStoreResponse{NumOfBooks: countResponse, NumOfAuthors: *cardinalityValue.Value}, nil
}
