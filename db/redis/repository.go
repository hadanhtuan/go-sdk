package redis

import (
	"context"
	"encoding/json"
	"time"
)

func Set(key string, value any, ttl time.Duration) error {
	c := GetConnection() //TODO: this is singleton connection

	valueByte, err := json.Marshal(value)

	if err != nil {
		return err
	}

	c.Client.Set(context.Background(), key, valueByte, ttl)

	return nil
}

func Get(key string, record any) error {
	c := GetConnection() //TODO: this is singleton connection

	b, err := c.Client.Get(context.Background(), key).Result()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(b), &record)

	return err
}

func Increase(key string, value int64) (bool, error) {
	c := GetConnection()

	record, _ := c.Client.Get(context.Background(), key).Result()

	if record == "" {
		Set(key, 1, 0)
		return true, nil
	}

	c.Client.IncrBy(context.Background(), key, value)

	return true, nil
}

func Decrease(key string, value int64) (bool, error) {
	c := GetConnection()

	c.Client.DecrBy(context.Background(), key, value).Result()

	return true, nil
}

func Delete(keys []string) (bool, error) {
	c := GetConnection()

	c.Client.Del(context.Background(), keys...).Result()

	return true, nil
}
