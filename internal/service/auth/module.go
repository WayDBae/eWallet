package auth

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/internal/storage/rdb"
	"github.com/WayDBae/eWallet/internal/storage/user"
	"github.com/WayDBae/eWallet/pkg/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBAuth)

type BAuth interface {
	Registration(data entities.Registration, ctx context.Context) (code string, err error)
	// Проверка на OTP
	OTPVerify(data entities.OTPVerify, ctx context.Context) (signedToken string, err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger

	// Config
	config *config.Config

	// Storages

	user user.SUser
	rdb  rdb.SRedis
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger

	// Config
	Config *config.Config

	// Storages
	User user.SUser
	Rdb  rdb.SRedis
}

// NewBAuth ...
func NewBAuth(params Params) BAuth {
	return &provider{
		// Logger
		logger: params.Logger,

		// Config
		config: params.Config,

		// Storages
		user: params.User,
		rdb:  params.Rdb,
	}
}
