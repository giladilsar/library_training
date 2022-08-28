package store

import (
	"gin/repository/book_repository"
)

func fetchStoreDate() (*StoreInfo, error) {

	countResponse, err := book_repository.NewBookProvider().Count()
	if err != nil {
		return nil, err
	}

	numOfAuthors, err := book_repository.NewBookProvider().CountAuthors()
	if err != nil {
		return nil, err
	}

	return &StoreInfo{NumberOfBooks: countResponse, NumberOfAuthors: *numOfAuthors}, nil
}
