package wallet

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/internal/storage/user"
	"github.com/WayDBae/eWallet/internal/storage/wallet"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBWallet)

type BWallet interface {
	GetMany(ctx context.Context) (wallets []entities.Wallet, err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger
	user   user.SUser
	wallet wallet.SWallet
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger
	User   user.SUser
	Wallet wallet.SWallet
}

// NewBWallet ...
func NewBWallet(params Params) BWallet {
	return &provider{
		// Logger
		logger: params.Logger,
		user:   params.User,
		wallet: params.Wallet,
	}
}
