// Package deckvredis provides Redis storage backend for Deckv.
package deckvredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Config holds the configuration for Redis connection.
type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// Storage implements the deckv.DataStorage interface using Redis as backend.
type Storage struct {
	c *redis.Client
}

// New creates a new Redis storage instance with the provided configuration.
func New(ctx context.Context, cfg Config) *Storage {
	return &Storage{
		c: redis.NewClient(&redis.Options{
			Addr:     cfg.Host + ":" + cfg.Port,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
	}
}

// Save stores the blocklist data in Redis.
func (s *Storage) Save(ctx context.Context, data map[string]uint8) error {
	d := make(map[string]interface{})
	for k, v := range data {
		d[k] = v
	}
	return s.c.HSet(ctx, "deckv", d).Err()
}

// Check verifies if a key exists in the Redis storage.
func (s *Storage) Check(ctx context.Context, key string) (bool, error) {
	v, err := s.c.HGet(ctx, "deckv", key).Int()
	if err != nil {
		return false, err
	}
	return v == 1, nil
}
