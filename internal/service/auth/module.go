package auth

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/internal/helpers/jwt"
	"github.com/WayDBae/eWallet/internal/storage/rdb"
	"github.com/WayDBae/eWallet/internal/storage/user"
	"github.com/WayDBae/eWallet/pkg/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBAuth)

type BAuth interface {
	// Registration - Регистрация
	Registration(data entities.AuthRegistration, ctx context.Context) (code string, err error)
	// OTPVerify - Проверка на OTP
	OTPVerify(data entities.AuthOTPVerify, ctx context.Context) (accessToken, refreshToken string, err error)
	// Login - Вход
	Login(data entities.AuthLogin, ctx context.Context) (accessToken, refreshToken string, err error)
	// Refresh - Обновление токена
	Refresh(token string, ctx context.Context) (accessToken, refreshToken string, err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger

	// Config
	config *config.Config

	// Storages

	user user.SUser
	rdb  rdb.SRedis
	jwt  jwt.HJWT
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
	JWT  jwt.HJWT
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
		jwt:  params.JWT,
	}
}
