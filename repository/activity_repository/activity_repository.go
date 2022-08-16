package activity_repository

import (
	"encoding/json"
	"gin/config"
	"gin/models"
	"gopkg.in/redis.v5"
)

func (rs RedisActivityRepository) GetUserActivity(username string) ([]string, error) {
	return rs.client.LRange(username, 0, 2).Result()
}

func (rs RedisActivityRepository) SetUserActivity(username string, request models.UserRequest) error {
	reqJson, err := json.Marshal(request)
	if err != nil {
		return err
	}

	cmd := rs.client.LPush(username, reqJson)
	if err := cmd.Err(); err != nil {
		return err
	}

	err = rs.client.LTrim(username, 0, 2).Err()
	return err
}

type ActivityService interface {
	GetUserActivity(username string) ([]string, error)
	SetUserActivity(username string, request models.UserRequest) error
}

type RedisActivityRepository struct {
	client *redis.Client
}

func GetActivityRepository() ActivityService {
	return RedisActivityRepository{config.RedisClient}
}
