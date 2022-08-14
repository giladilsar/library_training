package store

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

const indexName = "gilad_books"

func fetchStoreDate(es *elastic.Client) (gin.H, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	countResponse, err := es.Count().Index(indexName).Do(ctx)
	if err != nil {
		return nil, err
	}

	cardRes, err := es.Search().
		Index("gilad_books").
		Aggregation("authors",
			elastic.NewCardinalityAggregation().
				Field("author_name.keyword")).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	cardinalityValue, _ := cardRes.Aggregations.Cardinality("authors")
	return gin.H{"number_of_books": countResponse, "number_of_authors": *cardinalityValue.Value}, nil
}
