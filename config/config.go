package config

import "github.com/olivere/elastic/v7"

var (
	ElasticClient *elastic.Client
)

func setupElastic(sourceIndexURL string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(sourceIndexURL), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return client
}

func Setup() {
	ElasticClient = setupElastic("http://es-search-7.fiverrdev.com:9200")
}
