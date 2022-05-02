//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log"

	"github.com/AkankshaNichrelay/TableFinder/internal/cache"
	"github.com/AkankshaNichrelay/TableFinder/internal/database"
	"github.com/AkankshaNichrelay/TableFinder/internal/handler"
	"github.com/AkankshaNichrelay/TableFinder/internal/redis"
	"github.com/AkankshaNichrelay/TableFinder/internal/restaurant"
	"github.com/google/wire"
)

// cacheSet associates redis client struct with cacheAccessor interface
var cacheSet = wire.NewSet(redis.New, wire.Bind(new(cache.CacheAccessor), new(*redis.Client)))

func InitializeAndRun(ctx context.Context, logger *log.Logger) (*handler.Handler, error) {
	panic(
		wire.Build(
			redis.NewRedisConfig,
			cacheSet,
			cache.New,
			database.New,
			restaurant.New,
			handler.New,
		),
	)
}
