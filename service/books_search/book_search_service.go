package books_search

import (
	"encoding/json"
	"gin/models"
	"gin/repository/book_repository"
)

func searchBooks(req *bookSearchRequest) ([]models.Book, error) {
	query := buildQuery(req)

	searchResult, err := book_repository.GetBookRepository().SearchBook(query)

	if err != nil {
		return nil, err
	}

	books := make([]models.Book, searchResult.Hits.TotalHits.Value)
	for i, hit := range searchResult.Hits.Hits {
		book := models.Book{Id: hit.Id}
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			return nil, err
		}
		books[i] = book
	}

	return books, nil
}
