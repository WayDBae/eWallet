package logger

import (
	"github.com/WayDBae/eWallet/pkg/config"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(InitLogger)

// Dependencies ...
type Dependencies struct {
	fx.In

	Config *config.Config
}
