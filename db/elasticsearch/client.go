package es

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/hadanhtuan/go-sdk/config"
)

type ESClient struct {
	Client *elasticsearch.TypedClient
}

var (
	ES *ESClient
)

func ConnectElasticSearch() *ESClient {
	ES = new(ESClient)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	esUri := fmt.Sprintf(
		"http://%s:%d",
		config.AppConfig.ES.Host,
		config.AppConfig.ES.Port,
	)

	esCnf := elasticsearch.Config{
		Addresses: []string{esUri},
		Username:  config.AppConfig.ES.Username,
		Password:  config.AppConfig.ES.Password,
	}

	client, err := elasticsearch.NewTypedClient(esCnf)

	if err != nil {
		slog.Info(err.Error())
	}

	ES.Client = client

	return ES
}

func GetConnection() *ESClient {
	if ES != nil {
		return ES
	}
	return ConnectElasticSearch()
}
