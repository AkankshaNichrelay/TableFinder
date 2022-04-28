package cache

import (
	"context"
	"fmt"
	"time"
)

// CacheAccessor interface for caching
type CacheAccessor interface {
	Get(ctx context.Context, key string) (interface{}, bool, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
}

// Cache implementing the CacheAccessor interface
type CacheClient struct {
	cache CacheAccessor
}

// New ...
func New(cache CacheAccessor) *CacheClient {
	cc := CacheClient{cache: cache}
	return &cc
}

// Get ...
func (cc *CacheClient) Get(key string) (interface{}, error) {
	ctx := context.Background()
	item, ok, _ := cc.cache.Get(ctx, key)
	if !ok {
		return nil, fmt.Errorf("cache Get failed. key %s not found", key)
	}

	return item, nil
}

// Set ...
func (cc *CacheClient) Set(key string, item string, expiration time.Duration) error {
	ctx := context.Background()
	err := cc.cache.Set(ctx, key, item, expiration)
	if err != nil {
		return fmt.Errorf("cache Set failed. key %s err %v", key, err)
	}

	return nil
}
