package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/WayDBae/eWallet/pkg/bootstrap/http/router"
	"github.com/WayDBae/eWallet/pkg/config"
	"go.uber.org/fx"
)

// ServerModule ...
var ServerModule = fx.Provide(NewServer)

// Dependecies ...
type Dependecies struct {
	fx.In

	Config *config.Config
	Router *router.HTTPRouter
}

// NewServer ...
func NewServer(params Dependecies) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("[Server] PORT env is empty, using config...")
		port = fmt.Sprint(params.Config.Server.Port)
	} else {
		fmt.Println("[Server] Got PORT from env:", port)
	}
	url := net.JoinHostPort("", port)

	return &http.Server{
		MaxHeaderBytes: 32 << 20, // 32 Mb
		Addr:           url,
		Handler:        params.Router,
	}
}
