package books

import (
	"encoding/json"
	"fmt"
	"gin/models"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

const IndexName = "gilad_books"

func getBook(es *elastic.Client, id string) (*models.Book, error) {
	ctx, cancel := utils.GetContext()
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
	err = json.Unmarshal(searchResults.Source, &book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func deleteBookById(es *elastic.Client, id string) error {
	ctx, cancel := utils.GetContext()
	defer cancel()

	_, err := es.Delete().
		Index(IndexName).
		Id(id).
		Do(ctx)

	return err
}

func createBookFromPayload(es *elastic.Client, req *createBookRequest) (gin.H, error) {
	ctx, cancel := utils.GetContext()
	defer cancel()

	bookToSave := models.Book{
		Title:          req.Title,
		AuthorName:     req.AuthorName,
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

	return gin.H{"Id": res.Id}, nil
}

func updateBook(es *elastic.Client, req *updateBookRequest) error {
	ctx, cancel := utils.GetContext()
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
