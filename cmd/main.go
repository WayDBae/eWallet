package main

import (
	"context"
	"fmt"

	net_http "net/http"

	"github.com/WayDBae/eWallet/internal/helpers"
	"github.com/WayDBae/eWallet/internal/service"
	"github.com/WayDBae/eWallet/internal/storage"
	"github.com/WayDBae/eWallet/internal/transport/http/handlers"
	"github.com/WayDBae/eWallet/internal/transport/http/router"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http"
	"github.com/WayDBae/eWallet/pkg/config"
	"github.com/WayDBae/eWallet/pkg/databases"
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
		helpers.Module,
		// gateways.Module,
		ModuleLifecycleHooks,
	)
	app.Run()
}

var ModuleLifecycleHooks = fx.Invoke(RegisterHooks)

// RegisterHooks ...
func RegisterHooks(lifecycle fx.Lifecycle, server *net_http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("[Server] Starting HTTP server on", server.Addr)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != net_http.ErrServerClosed {
					fmt.Println("[Server] ListenAndServe error:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("[Server] Shutting down HTTP server...")
			return server.Shutdown(ctx)
		},
	})
}
