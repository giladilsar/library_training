package books

import (
	"encoding/json"
	"fmt"
	"gin/context_helper"
	"gin/models"
	"github.com/olivere/elastic/v7"
)

const IndexName = "gilad_books"

func getBook(es *elastic.Client, id string) (*models.Book, error) {
	ctx, cancel := context_helper.GetContext()
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

func deleteBookById(es *elastic.Client, id string) error {
	ctx, cancel := context_helper.GetContext()
	defer cancel()

	_, err := es.Delete().
		Index(IndexName).
		Id(id).
		Do(ctx)

	return err
}

func createBookFromPayload(es *elastic.Client, req *createBookRequest) (*createBookResponse, error) {
	ctx, cancel := context_helper.GetContext()
	defer cancel()

	bookToSave := models.Book{
		Title:          req.Title,
		Name:           req.Name,
		Price:          req.Price,
		EbookAvailable: req.EbookAvailable,
		PublishDate:    req.PublishDate,
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

func updateBook(es *elastic.Client, req *updateBookRequest) error {
	ctx, cancel := context_helper.GetContext()
	defer cancel()

	_, err := es.Update().
		Index(IndexName).
		Id(req.Id).
		Doc(updateBookTitleCommand{req.Title}).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
