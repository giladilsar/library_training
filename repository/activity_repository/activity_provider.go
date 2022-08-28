package activity_repository

import "gin/models"

type ActivityProvider interface {
	GetUserActivity(username string) ([]string, error)
	SetUserActivity(username string, request models.UserRequest) error
}
