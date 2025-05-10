package wallet

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(NewSWallet)

type SWallet interface {
	GetMany(filter entities.Wallet, ctx context.Context) (wallets []entities.Wallet, err error)
	Create(data entities.Wallet, ctx context.Context) (wallet entities.Wallet, err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger

	postgres *gorm.DB
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger

	Postgres *gorm.DB
}

// NewSWallet ...
func NewSWallet(params Params) SWallet {
	return &provider{
		// Logger
		logger: params.Logger,

		postgres: params.Postgres,
	}
}
