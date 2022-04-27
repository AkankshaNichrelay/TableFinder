//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log"

	"github.com/AkankshaNichrelay/TableFinder/handler"
	"github.com/AkankshaNichrelay/TableFinder/restaurant"
	"github.com/google/wire"
)

func InitializeAndRun(ctx context.Context, logger *log.Logger) (*handler.Handler, error) {
	panic(
		wire.Build(
			restaurant.New,
			handler.New,
		),
	)
}
