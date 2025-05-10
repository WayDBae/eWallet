package service

import (
	"github.com/WayDBae/eWallet/internal/service/auth"
	"github.com/WayDBae/eWallet/internal/service/wallet"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	auth.Module,
	wallet.Module,
)
