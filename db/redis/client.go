package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hadanhtuan/go-sdk/config"
	"github.com/redis/go-redis/v9"
)

type CacheClient struct {
	Client *redis.Client
}

var (
	Cache *CacheClient
)

func ConnectRedis() *CacheClient {
	Cache = new(CacheClient)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.AppConfig.Cache.CacheHost, config.AppConfig.Cache.CachePort),
		Password: config.AppConfig.Cache.CachePass,
		DB:       config.AppConfig.Cache.CacheDB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	Cache.Client = client
	
	log.Println("🚀 Connected Successfully to Redis")
	return Cache
}

func GetConnection() *CacheClient {
	if Cache != nil {
		return Cache
	}
	return ConnectRedis()
}
