package book_repository

import (
	"gin/config"
	"gin/models"
	"gin/service/books/dto"
	"gin/utils"
	"github.com/olivere/elastic/v7"
)

type ElasticBookManager struct {
	es        *elastic.Client
	indexName string
}

func NewBookManager() BookManager {
	return ElasticBookManager{config.ElasticClient, "gilad_books"}
}

func (r ElasticBookManager) GetById(id string) (*dto.SearchResult, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	res, err := r.es.Get().
		Index(r.indexName).
		Id(id).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return &dto.SearchResult{Found: res.Found, RawData: res.Source}, nil
}

func (r ElasticBookManager) DeleteById(id string) error {
	ctx, cancel := utils.GetContext()
	defer cancel()

	_, err := r.es.Delete().
		Index(r.indexName).
		Id(id).
		Do(ctx)

	return err
}

func (r ElasticBookManager) InsertBook(book models.Book) (*string, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	res, err := r.es.Index().
		Index(r.indexName).
		BodyJson(book).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return &res.Id, err
}

func (r ElasticBookManager) UpdateBook(req dto.UpdateBookTitleCommand, id string) error {
	ctx, cancel := utils.GetContext()
	defer cancel()

	_, err := r.es.Update().
		Index(r.indexName).
		Id(id).
		Doc(req).
		Do(ctx)

	return err
}

func (r ElasticBookManager) SearchBook(query dto.BookSearchQuery) (*elastic.SearchResult, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	return r.es.Search().
		Index(r.indexName).
		Query(query).
		From(0).
		Size(100).
		Do(ctx)
}

func (r ElasticBookManager) Count() (int64, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	return r.es.Count().Index(r.indexName).Do(ctx)
}

func (r ElasticBookManager) CountAuthors() (*float64, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	cardinalityAgg := elastic.NewCardinalityAggregation().Field("author_name.keyword")
	res, err := r.es.Search().
		Index(r.indexName).
		Aggregation("authors", cardinalityAgg).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	cardinalityValue, _ := res.Aggregations.Cardinality("authors")
	return cardinalityValue.Value, nil
}
