package brokers

import (
	"context"

	"github.com/WayDBae/eWallet/pkg/brokers/rabbitmq"
	"github.com/WayDBae/eWallet/pkg/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// AMQPModule ...
// var Module = fx.Provide(NewAMQPConn)

// Dependencies ...
type Dependencies struct {
	fx.In

	Logger zerolog.Logger
	Config *config.Config
}

// NewAMQPConn ...
func NewAMQPConn(params Dependencies) rabbitmq.Client {
	return NewRabbitMQConn(params, context.Background())
}
