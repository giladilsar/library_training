package books

import (
	"context"
	"encoding/json"
	"fmt"
	"gin/models"
	"github.com/olivere/elastic/v7"
	"time"
)

const IndexName = "gilad_books"

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 1000*time.Second)
}

func getBook(es *elastic.Client, id string) (*models.Book, error) {
	ctx, cancel := getContext()
	defer cancel()

	searchResults, err := es.Get().
		Index(IndexName).
		Id(id).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	if !searchResults.Found {
		return nil, fmt.Errorf("books with id %s could no be found", id)
	}

	book := models.Book{Id: id}
	jsonError := json.Unmarshal(searchResults.Source, &book)
	if jsonError != nil {
		return nil, err
	}

	return &book, nil
}

func createBookFromPayload(es *elastic.Client, req *createBookRequest) (*createBookResponse, error) {
	ctx, cancel := getContext()
	defer cancel()

	publishDate, parseError := time.Parse("2006-01-02", req.PublishDate)
	if parseError != nil {
		return nil, parseError
	}

	bookToSave := models.Book{
		Title:          req.Title,
		Name:           req.Name,
		Price:          req.Price,
		EbookAvailable: req.EbookAvailable,
		PublishDate:    publishDate,
	}
	res, err := es.Index().
		Index(IndexName).
		BodyJson(bookToSave).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return &createBookResponse{res.Id}, nil
}
