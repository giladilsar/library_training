package utils

import (
	"context"
	"github.com/olivere/elastic/v7"
	"net/http"
	"time"
)

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 1*time.Second)
}

func GetErrorResponseStatus(err error) int {
	if e, ok := err.(*elastic.Error); ok == true {
		return e.Status
	} else {
		return http.StatusInternalServerError
	}
}
