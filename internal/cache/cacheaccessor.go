package cache

import (
	"context"
	"time"
)

// Accessor interface for caching
type Accessor interface {
	Get(ctx context.Context, key string) (interface{}, bool, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
}

// CacheNoop used for testing implements CacheAccessor interface
type CacheNoop struct{}

// Get noop
func (c *CacheNoop) Get(ctx context.Context, key string) (interface{}, bool, error) {
	return nil, false, nil
}

// Set noop
func (c *CacheNoop) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

// Delete noop
func (c *CacheNoop) Delete(ctx context.Context, keys ...string) error {
	return nil
}
