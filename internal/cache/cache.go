package cache

import (
	"context"
	"fmt"
	"time"
)

// Cache implementing the Accessor interface
type Client struct {
	cache Accessor
}

// New returns a new Cache Client instance
func New(cache Accessor) *Client {
	cc := Client{cache: cache}
	return &cc
}

// Get ...
func (cc *Client) Get(key string) (interface{}, error) {
	ctx := context.Background()
	item, ok, _ := cc.cache.Get(ctx, key)
	if !ok {
		return nil, fmt.Errorf("cache Get failed. key %s not found", key)
	}

	return item, nil
}

// Set ...
func (cc *Client) Set(key string, item string, expiration time.Duration) error {
	ctx := context.Background()
	err := cc.cache.Set(ctx, key, item, expiration)
	if err != nil {
		return fmt.Errorf("cache Set failed. key %s err %v", key, err)
	}

	return nil
}
