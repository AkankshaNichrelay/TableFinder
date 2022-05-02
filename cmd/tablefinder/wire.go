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
var cacheSet = wire.NewSet(redis.New, wire.Bind(new(cache.Accessor), new(*redis.Client)))

var dbSet = wire.NewSet(database.New, wire.Bind(new(database.DBAccessor), new(*database.MySQL)))

func InitializeAndRun(ctx context.Context, logger *log.Logger) (*handler.Handler, error) {
	panic(
		wire.Build(
			redis.NewRedisConfig,
			cacheSet,
			cache.New,
			dbSet,
			restaurant.New,
			handler.New,
		),
	)
}
