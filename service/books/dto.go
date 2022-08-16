package books

type createBookRequest struct {
	Title          string  `json:"title" binding:"required"`
	AuthorName     string  `json:"author_name" binding:"required"`
	Price          float32 `json:"price" binding:"required"`
	EbookAvailable bool    `json:"ebook_available"`
	PublishDate    string  `json:"publish_date" binding:"required"`
}

type updateBookRequest struct {
	Id    string
	Title string `json:"title" binding:"required"`
}

type UpdateBookTitleCommand struct {
	Title string `json:"title" binding:"required"`
}

type SearchResult struct {
	Found   bool
	RawData []byte
}
