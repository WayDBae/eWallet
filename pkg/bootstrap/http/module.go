package http

import (
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/middlewares"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/server"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	middlewares.Module,
	misc.Module,

	server.ServerModule,
)
