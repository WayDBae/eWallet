package helpers

import (
	"github.com/WayDBae/eWallet/internal/helpers/jwt"
	"go.uber.org/fx"
)

var Module = fx.Options(jwt.Module)
