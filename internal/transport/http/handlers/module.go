package handlers

import (
	"github.com/WayDBae/eWallet/internal/service/auth"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHandlerProvider)

// HandlerDependencies ...
type HandlerDependencies struct {
	fx.In
	Logger zerolog.Logger

	// Services
	Auth auth.BAuth
}

// Handler ...
type Handler struct {
	logger zerolog.Logger

	// Services
	auth auth.BAuth
}

// NewHandlerProvider ...
func NewHandlerProvider(params HandlerDependencies) *Handler {
	return &Handler{
		logger: params.Logger,

		// Services
		auth: params.Auth,
	}
}
