package book_repository

import (
	"gin/models"
	"gin/service/books/dto"
	"github.com/olivere/elastic/v7"
)

type BookManager interface {
	GetById(id string) (*dto.SearchResult, error)
	DeleteById(id string) error
	InsertBook(book models.Book) (*string, error)
	UpdateBook(req dto.UpdateBookTitleCommand, id string) error
	SearchBook(query dto.BookSearchQuery) (*elastic.SearchResult, error)
	Count() (int64, error)
	CountAuthors() (*float64, error)
}
