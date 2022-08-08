package models

import "time"

type Book struct {
	Id             string    `json:"_id,omitempty"`
	Title          string    `json:"title"`
	Price          float32   `json:"price"`
	Name           string    `json:"name"`
	EbookAvailable bool      `json:"ebook_available"`
	PublishDate    time.Time `json:"publish_date"`
}
