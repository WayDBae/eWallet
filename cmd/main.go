package main

import (
	"context"

	net_http "net/http"

	"github.com/WayDBae/eWallet/internal/service"
	"github.com/WayDBae/eWallet/internal/storage"
	"github.com/WayDBae/eWallet/internal/transport/http/handlers"
	"github.com/WayDBae/eWallet/internal/transport/http/router"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http"
	"github.com/WayDBae/eWallet/pkg/config"
	"github.com/WayDBae/eWallet/pkg/databases"
	"github.com/WayDBae/eWallet/pkg/gateways"
	"github.com/WayDBae/eWallet/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		config.Module,
		logger.Module,
		databases.Module,
		service.Module,
		storage.Module,
		handlers.Module,
		router.Module,
		gateways.Module,
		ModuleLifecycleHooks,
	)
	app.Run()
}

var ModuleLifecycleHooks = fx.Invoke(RegisterHooks)

// RegisterHooks ...
func RegisterHooks(lifecycle fx.Lifecycle, server *net_http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
