// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/AkankshaNichrelay/TableFinder/internal/cache"
	"github.com/AkankshaNichrelay/TableFinder/internal/handler"
	"github.com/AkankshaNichrelay/TableFinder/internal/redis"
	"github.com/AkankshaNichrelay/TableFinder/internal/restaurant"
	"github.com/google/wire"
	"log"
)

// Injectors from wire.go:

func InitializeAndRun(ctx context.Context, logger *log.Logger) (*handler.Handler, error) {
	config := redis.NewRedisConfig()
	client := redis.New(logger, config)
	cacheClient := cache.New(client)
	restaurants := restaurant.New(cacheClient)
	handlerHandler := handler.New(logger, restaurants)
	return handlerHandler, nil
}

// wire.go:

// cacheSet associates redis client struct with cacheAccessor interface
var cacheSet = wire.NewSet(redis.New, wire.Bind(new(cache.CacheAccessor), new(*redis.Client)))
