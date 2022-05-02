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
func Get(ctx context.Context, key string) (interface{}, bool, error) {
	return nil, false, nil
}

// Set noop
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

// Delete noop
func Delete(ctx context.Context, keys ...string) error {
	return nil
}
