package gateways

import (
	"github.com/WayDBae/eWallet/pkg/gateways/example"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	example.Module,
)
