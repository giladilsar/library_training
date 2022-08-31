package models

type Book struct {
	Id             string  `json:"_id,omitempty"`
	Title          string  `json:"title"`
	Price          float32 `json:"price"`
	AuthorName     string  `json:"author_name"`
	EbookAvailable bool    `json:"ebook_available"`
	PublishDate    string  `json:"publish_date"`
}
