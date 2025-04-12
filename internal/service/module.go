package service

import (
	"github.com/WayDBae/eWallet/internal/service/auth"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	auth.Module,
)
