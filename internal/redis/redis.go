package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// Config contains Redis client configuration
type Config struct {
	Host     string
	Port     int
	Password string
}

// Client is a wrapper for Redis Client
type Client struct {
	Client *redis.Client
	Config *Config
	logger *log.Logger
}

// New creates a new redis client - TODO. This is just for wire.go
func New(log *log.Logger, config *Config) *Client {
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	rd := Client{
		logger: log,
		Config: config,
		Client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.Password,
		}),
	}

	return &rd
}

// Get TODO
func (rd *Client) Get(ctx context.Context, key string) (interface{}, bool, error) {
	return nil, false, nil
}

// Set TODO
func (rd *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

// Delete TODO
func (rd *Client) Delete(ctx context.Context, keys ...string) error {
	return nil
}
