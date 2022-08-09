package searchHelper

import (
	"context"
	"time"
)

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 1000*time.Second)
}
