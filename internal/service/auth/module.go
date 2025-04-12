package auth

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/internal/storage/rdb"
	"github.com/WayDBae/eWallet/internal/storage/user"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBAuth)

type BAuth interface {
	Registration(data entities.Registration, ctx context.Context) (code string, err error)
	// Проверка на OTP
	OTPVerify(data entities.OTPVerify, ctx context.Context) (err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger

	// Storages
	user user.SUser
	rdb  rdb.SRedis
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger

	// Storages
	User user.SUser
	Rdb  rdb.SRedis
}

// NewBAuth ...
func NewBAuth(params Params) BAuth {
	return &provider{
		// Logger
		logger: params.Logger,

		// Storages
		user: params.User,
		rdb:  params.Rdb,
	}
}
