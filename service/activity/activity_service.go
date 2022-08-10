package activity

import (
	"encoding/json"
	"gin/models"
	"gopkg.in/redis.v5"
)

func getUserActivity(redisClient *redis.Client, username string) ([]models.UserRequest, error) {
	response, err := redisClient.LRange(username, 0, 2).Result()
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
