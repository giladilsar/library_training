package books

import (
	"encoding/json"
	"fmt"
	"gin/models"
	"gin/repository/book_repository"
)

func getBook(id string) (*models.Book, error) {
	searchResults, err := book_repository.GetBookRepository().GetById(id)

	if err != nil {
		return nil, err
	}

	if !searchResults.Found {
		return nil, fmt.Errorf("books with id %s could no be found", id)
	}

	book := models.Book{Id: id}
	err = json.Unmarshal(searchResults.RawData, &book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func deleteBook(id string) error {
	return book_repository.GetBookRepository().DeleteById(id)
}

func createBookFromPayload(req *createBookRequest) (*string, error) {
	bookToSave := models.Book{
		Title:          req.Title,
		AuthorName:     req.AuthorName,
		Price:          req.Price,
		EbookAvailable: req.EbookAvailable,
		PublishDate:    req.PublishDate,
	}

	return book_repository.GetBookRepository().InsertBook(bookToSave)
}

func updateBook(req *updateBookRequest) error {
	return book_repository.GetBookRepository().UpdateBook(UpdateBookTitleCommand{req.Title}, req.Id)
}
