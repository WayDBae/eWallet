package jwt

import (
	"context"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewHJWT)

type HJWT interface {
	GenerateAccessToken(user entities.User, t time.Duration, ctx context.Context) (accessToken string, err error)
	GenerateRefreshToken(user entities.User, t time.Duration, ctx context.Context) (accessToken string, err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger

	// Config
	config *config.Config
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger

	// Config
	Config *config.Config
}

// NewHJWT ...
func NewHJWT(params Params) HJWT {
	return &provider{
		// Logger
		logger: params.Logger,

		// Config
		config: params.Config,
	}
}
