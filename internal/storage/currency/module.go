package currency

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(NewSCurrency)

type SCurrency interface {
	GetMany(filter entities.Currency, ctx context.Context) (currencies []entities.Currency, err error)
	Get(data entities.Currency, ctx context.Context) (currency entities.Currency, err error)
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

// NewSCurrency ...
func NewSCurrency(params Params) SCurrency {
	return &provider{
		// Logger
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
