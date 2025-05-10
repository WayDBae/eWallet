package router

import (
	"net/http"

	"github.com/WayDBae/eWallet/internal/transport/http/handlers"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/middlewares"
	transportHTTP "github.com/WayDBae/eWallet/pkg/bootstrap/http/router"
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

	router.GET("/ping", h.HPingPong, mw.CORS)

	// Auth
	router.POST("/auth/registration", h.HRegistration, mw.CORS)
	router.POST("/auth/otp-verify", h.HOTPVerify, mw.CORS)
	router.POST("/auth/login", h.HLogin, mw.CORS)
	router.POST("/auth/refresh", h.HRefreshToken, mw.CORS)
	
	router.GET("/wallet/getMany", h.HWalletGetMany, mw.JWT, mw.CORS)
	// router.GET("/metrics", AdaptHandler(promhttp.Handler()), mw.CORS)
	return
}
