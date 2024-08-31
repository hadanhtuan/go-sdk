package es

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log/slog"
	"time"
)

type ESClient struct {
	Client *elasticsearch.TypedClient
}

var (
	ES *ESClient
)

func ConnectElasticSearch(host, port, username, password string) *ESClient {
	ES = new(ESClient)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	esUri := fmt.Sprintf(
		"http://%s:%s", host, port,
	)

	esCnf := elasticsearch.Config{
		Addresses: []string{esUri},
		Username:  username,
		Password:  password,
	}

	client, err := elasticsearch.NewTypedClient(esCnf)

	if err != nil {
		slog.Info(err.Error())
	}

	ES.Client = client
	fmt.Println("[ 🚀 ] Connected Successfully to Elasticsearch")

	return ES
}

func GetConnection() *ESClient {
	if ES != nil {
		return ES
	}
	panic("Cannot connect to Elasticsearch")
}
