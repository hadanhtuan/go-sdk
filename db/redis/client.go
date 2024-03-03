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
	Client *redis.Client
}

var (
	Cache    *CacheClient
	cacheEnv sdk.CacheEnv
)

func ConnectRedis() {
	if Cache != nil {
		return
	}

	sdk.ParseENV(&cacheEnv)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cacheEnv.CacheHost, cacheEnv.CachePort),
		Password: cacheEnv.CachePass,
		DB:       cacheEnv.CacheDB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	Cache = new(CacheClient)
	Cache.Client = client
	log.Println("🗃️  Connected Successfully to the Redis")
}
