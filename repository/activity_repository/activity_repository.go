package activity_repository

import (
	"encoding/json"
	"gin/config"
	"gin/models"
	"gopkg.in/redis.v5"
)

type RedisActivityManager struct {
	client *redis.Client
}

func NewActivityManager() ActivityManager {
	return RedisActivityManager{config.RedisClient}
}

func (rs RedisActivityManager) GetUserActivity(username string) ([]string, error) {
	return rs.client.LRange(username, 0, 2).Result()
}

func (rs RedisActivityManager) SetUserActivity(username string, request models.UserRequest) error {
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
