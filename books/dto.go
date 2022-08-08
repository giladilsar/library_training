package books

type createBookRequest struct {
	Title          string  `json:"title" binding:"required"`
	Name           string  `json:"author_name" binding:"required"`
	Price          float32 `json:"price" binding:"required"`
	EbookAvailable bool    `json:"ebook_available"`
	PublishDate    string  `json:"publish_date" binding:"required"`
	Username       string  `json:"username" binding:"required"`
}

type createBookResponse struct {
	Id string `json:"id"`
}

type updateBookRequest struct {
	Id       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type updateBookTitleCommand struct {
	Title string `json:"title" binding:"required"`
}
