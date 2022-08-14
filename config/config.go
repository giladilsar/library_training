package config

import (
	"github.com/olivere/elastic/v7"
	"gopkg.in/redis.v5"
)

var (
	ElasticClient *elastic.Client
	RedisClient   *redis.Client
)

const (
	elasticConnectionUrl = "http://es-search-7.fiverrdev.com:9200"
	redisConnectionUrl   = "redis.default.fiverrdev.com:6382"
)

func setupElastic(sourceIndexURL string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(sourceIndexURL), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return client
}

func setupRedis(host string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: host,
	})

	pingError := client.Ping().Err()
	if pingError != nil {
		panic(pingError)
	}

	return client
}

func Setup() {
	ElasticClient = setupElastic(elasticConnectionUrl)
	RedisClient = setupRedis(redisConnectionUrl)
}
