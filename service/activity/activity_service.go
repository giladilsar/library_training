package activity

import (
	"encoding/json"
	"gin/models"
	"gin/repository/activity_repository"
)

func getUserActivity(username string) ([]models.UserRequest, error) {
	response, err := activity_repository.GetActivityRepository().GetUserActivity(username)
	if err != nil {
		return nil, err
	}
	operations := make([]models.UserRequest, len(response))
	for i, request := range response {
		operations[i] = models.UserRequest{}
		err := json.Unmarshal([]byte(request), &operations[i])
		if err != nil {
			return nil, err
		}
	}

	return operations, nil
}
