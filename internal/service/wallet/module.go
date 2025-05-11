package wallet

import (
	"context"

	"github.com/WayDBae/eWallet/internal/storage/currency"
	"github.com/WayDBae/eWallet/internal/storage/user"
	"github.com/WayDBae/eWallet/internal/storage/wallet"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBWallet)

type BWallet interface {
	GetMany(ctx context.Context) (data []map[string]any, err error)
}

type provider struct {
	// Logger

	logger   zerolog.Logger
	user     user.SUser
	wallet   wallet.SWallet
	currency currency.SCurrency
}

type Params struct {
	fx.In

	// Logger
	Logger   zerolog.Logger
	User     user.SUser
	Wallet   wallet.SWallet
	Currency currency.SCurrency
}

// NewBWallet ...
func NewBWallet(params Params) BWallet {
	return &provider{
		// Logger
		logger:   params.Logger,
		user:     params.User,
		wallet:   params.Wallet,
		currency: params.Currency,
	}
}
