package book_repository

import (
	"gin/config"
	"gin/models"
	"gin/service/books"
	"gin/service/books_search"
	"gin/utils"
	"github.com/olivere/elastic/v7"
)

const IndexName = "gilad_books"

func (r ElasticBookRepository) GetById(id string) (*books.SearchResult, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	res, err := r.es.Get().
		Index(IndexName).
		Id(id).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return &books.SearchResult{Found: res.Found, RawData: res.Source}, nil
}

func (r ElasticBookRepository) DeleteById(id string) error {
	ctx, cancel := utils.GetContext()
	defer cancel()

	_, err := r.es.Delete().
		Index(IndexName).
		Id(id).
		Do(ctx)

	return err
}

func (r ElasticBookRepository) InsertBook(book models.Book) (*string, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	res, err := r.es.Index().
		Index(IndexName).
		BodyJson(book).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return &res.Id, err
}

func (r ElasticBookRepository) UpdateBook(req books.UpdateBookTitleCommand, id string) error {
	ctx, cancel := utils.GetContext()
	defer cancel()

	_, err := r.es.Update().
		Index(IndexName).
		Id(id).
		Doc(req).
		Do(ctx)

	return err
}

func (r ElasticBookRepository) SearchBook(query books_search.BookSearchQuery) (*elastic.SearchResult, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	return r.es.Search().
		Index(IndexName).
		Query(query).
		From(0).
		Size(100).
		Do(ctx)
}

func (r ElasticBookRepository) Count() (int64, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	return r.es.Count().Index(IndexName).Do(ctx)
}

func (r ElasticBookRepository) CountAuthors() (*float64, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	cardinalityAgg := elastic.NewCardinalityAggregation().Field("author_name.keyword")
	res, err := r.es.Search().
		Index(IndexName).
		Aggregation("authors", cardinalityAgg).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	cardinalityValue, _ := res.Aggregations.Cardinality("authors")
	return cardinalityValue.Value, nil
}

type BookRepository interface {
	GetById(id string) (*books.SearchResult, error)
	DeleteById(id string) error
	InsertBook(book models.Book) (*string, error)
	UpdateBook(req books.UpdateBookTitleCommand, id string) error
	SearchBook(query books_search.BookSearchQuery) (*elastic.SearchResult, error)
	Count() (int64, error)
	CountAuthors() (*float64, error)
}

type ElasticBookRepository struct {
	es *elastic.Client
}

func GetBookRepository() BookRepository {
	return ElasticBookRepository{config.ElasticClient}
}
