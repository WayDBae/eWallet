package router

import (
	"net/http"

	"github.com/WayDBae/eWallet/internal/transport/http/handlers"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/middlewares"
	transportHTTP "github.com/WayDBae/eWallet/pkg/bootstrap/http/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func AdaptHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

// NewRouter ..
func NewRouter(h *handlers.Handler, mw middlewares.Middleware) (router *transportHTTP.HTTPRouter) {
	router = transportHTTP.NewRouter()
	router.ConnectSwagger(h.ServeSwaggerFiles)

	router.POST("/auth/registration", h.HRegistration, mw.RequestLog, mw.Metrics, mw.CORS)
	router.POST("/auth/otp-verify", h.HOTPVerify, mw.RequestLog, mw.Metrics, mw.CORS)

	router.POST("/auth/login", h.HLogin, mw.RequestLog, mw.Metrics, mw.CORS)
	router.GET("/ping", h.HPingPong, mw.RequestLog, mw.Metrics, mw.CORS)
	router.GET("/metrics", AdaptHandler(promhttp.Handler()), mw.RequestLog, mw.Metrics, mw.CORS)
	return
}
