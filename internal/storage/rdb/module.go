package rdb

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewSRedis)

type SRedis interface {
	Get(phone string, ctx context.Context) (data string, err error)
	Set(phone, data string, ctx context.Context) (err error)
}

type provider struct {
	client *redis.Client
	// Logger

	logger zerolog.Logger
}

type Params struct {
	fx.In
	Client *redis.Client

	// Logger
	Logger zerolog.Logger
}

// NewSRedis ...
func NewSRedis(params Params) SRedis {

	return &provider{
		// Logger
		client: params.Client,
		logger: params.Logger,
	}
}
