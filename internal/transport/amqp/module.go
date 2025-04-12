package amqp

import (
	"github.com/WayDBae/eWallet/internal/transport/amqp/consumers"
	"github.com/WayDBae/eWallet/internal/transport/amqp/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	consumers.Module,
	router.Module,
)
