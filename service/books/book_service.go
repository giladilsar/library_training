package books

import (
	"encoding/json"
	"fmt"
	"gin/models"
	"gin/repository/book_repository"
	"gin/service/books/dto"
)

func getBook(id string) (*models.Book, error) {
	searchResults, err := book_repository.NewBookProvider().GetById(id)

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
	return book_repository.NewBookProvider().DeleteById(id)
}

func createBookFromPayload(req *dto.CreateBookRequest) (*string, error) {
	bookToSave := models.Book{
		Title:          req.Title,
		AuthorName:     req.AuthorName,
		Price:          req.Price,
		EbookAvailable: req.EbookAvailable,
		PublishDate:    req.PublishDate,
	}

	return book_repository.NewBookProvider().InsertBook(bookToSave)
}

func updateBook(req *dto.UpdateBookRequest) error {
	return book_repository.NewBookProvider().UpdateBook(dto.UpdateBookTitleCommand{req.Title}, req.Id)
}
