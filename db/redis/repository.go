package cache

import (
	"context"
	"encoding/json"
	"time"
)

func (c *CacheClient) Set(key string, value any, ttl time.Duration) error {
	valueByte, err := json.Marshal(value)

	if err != nil {
		return err
	}

	c.Client.Set(context.Background(), key, valueByte, ttl)

	return nil
}

func (c *CacheClient) Get(key string, record any) error {
	b, err := c.Client.Get(context.Background(), key).Result()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(b), &record)

	return err
}

func (c *CacheClient) Increase(key string, value int64) (bool, error) {
	record, _ := c.Client.Get(context.Background(), key).Result()

	if record == "" {
		c.Set(key, 1, 0)
		return true, nil
	}

	c.Client.IncrBy(context.Background(), key, value)

	return true, nil
}

func (c *CacheClient) Decrease(key string, value int64) (bool, error) {
	c.Client.DecrBy(context.Background(), key, value).Result()

	return true, nil
}

func (c *CacheClient) Delete(keys []string) (bool, error) {
	c.Client.Del(context.Background(), keys...).Result()

	return true, nil
}
