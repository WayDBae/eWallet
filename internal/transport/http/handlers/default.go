package handlers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

// HPingPong — simple router to check server status
// swagger:operation GET /ping Ping ping
//
// Ping Pong!
//
// ## A simple router to check server status
//
// ---
//
//	responses:
//		200:
//			description: |-
//				## Pong ! 🥎
//			schema:
//				$ref: "#/responses/ping/schema"

func (h *Handler) HPingPong(rw http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  context.Context = r.Context()
	)

	defer resp.WriterJSON(rw, ctx)

	// Parse the string into time.Time
	resp.Message = response.ErrSuccess.Error()
	resp.Payload = "Clean arch pong!"
}

// HNotImplementation ...
func (h *Handler) HNotImplementation(rw http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  context.Context = r.Context()
	)

	defer resp.WriterJSON(rw, ctx)

	resp.Message = response.ErrNotImplementation.Error()
}

// ServeSwaggerFiles ...
func (h *Handler) ServeSwaggerFiles(w http.ResponseWriter, r *http.Request) {
	pwd, err := os.Getwd()
	if err != nil {
		h.logger.Error().Err(err).Msg("An error occurred in ServeSwaggerFiles")
	}

	// Используем filepath.Join для корректного формирования пути
	swaggerPath := filepath.Join(pwd, "../pkg/docs")
	h.logger.Debug().Err(err).Str("pwd", pwd).Str("swaggerPath", swaggerPath).Msg("Debugging swagger path")

	if strings.Contains(r.URL.String(), "yaml") {
		http.ServeFile(w, r, filepath.Join(swaggerPath, "swagger.yaml"))
		return
	}

	http.ServeFile(w, r, filepath.Join(swaggerPath, "swagger.json"))
}
