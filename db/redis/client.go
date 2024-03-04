package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hadanhtuan/go-sdk"
	"github.com/redis/go-redis/v9"
)

type CacheClient struct {
	Client   *redis.Client
	CacheEnv sdk.CacheEnv
}

var (
	Cache *CacheClient
)

func ConnectRedis() *CacheClient {
	if Cache != nil {
		return Cache
	}
	Cache = new(CacheClient)

	sdk.ParseENV(&Cache.CacheEnv)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", Cache.CacheEnv.CacheHost, Cache.CacheEnv.CachePort),
		Password: Cache.CacheEnv.CachePass,
		DB:       Cache.CacheEnv.CacheDB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	Cache.Client = client

	return Cache
}
