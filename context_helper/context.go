package context_helper

import (
	"context"
	"time"
)

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 1000*time.Second)
}
