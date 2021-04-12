package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cacher is cache interface.
type Cacher interface {
	Get(data interface{}, keys ...interface{}) error
	Set(data interface{}, keys ...interface{}) error
	Delete(keys ...interface{}) error
	Close() error
}

type cache struct {
	keyPrefix   string
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
		keyPrefix:   "hayasui",
		client:      client,
		expiredTime: expiredTime,
		ctx:         context.Background(),
	}, nil
}

func (c *cache) getKey(keys []interface{}) string {
	keyArr := []string{c.keyPrefix}
	for _, k := range keys {
		keyArr = append(keyArr, fmt.Sprintf("%v", k))
	}
	return strings.Join(keyArr, ":")
}

// Set to save data to cache,
func (c *cache) Set(data interface{}, keys ...interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, c.getKey(keys), d, c.expiredTime).Err()
}

// Get to get data from cache.
func (c *cache) Get(data interface{}, keys ...interface{}) error {
	d, err := c.client.Get(c.ctx, c.getKey(keys)).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(d), &data)
}

// Delete to delete data from cache.
func (c *cache) Delete(keys ...interface{}) error {
	return c.client.Del(c.ctx, c.getKey(keys)).Err()
}

// Close to close cache connection.
func (c *cache) Close() error {
	return c.client.Close()
}
