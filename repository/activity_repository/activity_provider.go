package activity_repository

import "gin/models"

type ActivityManager interface {
	GetUserActivity(username string) ([]string, error)
	SetUserActivity(username string, request models.UserRequest) error
}
