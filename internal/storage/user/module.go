package user

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(NewSUser)

type SUser interface {
	Get(data entities.User, ctx context.Context) (user entities.User, err error)
	GetByPhone(phone string, ctx context.Context) (user entities.User, err error)
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

// NewSUser ...
func NewSUser(params Params) SUser {
	return &provider{
		// Logger
		logger: params.Logger,

		postgres: params.Postgres,
	}
}
