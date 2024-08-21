package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheClient struct {
	Client *redis.Client
}

var (
	Cache *CacheClient
)

func ConnectRedis(host, port, password string, dbName int) *CacheClient {
	Cache = new(CacheClient)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       dbName,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	Cache.Client = client

	fmt.Println("[ 🚀 ] Connected Successfully to Redis")
	return Cache
}

func GetConnection() *CacheClient {
	if Cache != nil {
		return Cache
	}
	panic("Cannot connect to Redis")
}
