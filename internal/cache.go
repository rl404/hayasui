package internal

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cacher is cache interface.
type Cacher interface {
	Get(key string, data interface{}) error
	Set(key string, data interface{}) error
	Delete(key string) error
	Close() error
}

type cache struct {
	client      *redis.Client
	expiredTime time.Duration
	ctx         context.Context
}

// NewCache to create cache cache with default config.
func NewCache(address, password string, expiredTime time.Duration) (Cacher, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping test.
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &cache{
		client:      client,
		expiredTime: expiredTime,
		ctx:         context.Background(),
	}, nil
}

// Set to save data to cache,
func (c *cache) Set(key string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, key, d, c.expiredTime).Err()
}

// Get to get data from cache.
func (c *cache) Get(key string, data interface{}) error {
	d, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(d), &data)
}

// Delete to delete data from cache.
func (c *cache) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

// Close to close cache connection.
func (c *cache) Close() error {
	return c.client.Close()
}
