package handlers

import (
	"log"
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
	var resp response.Response
	ctx := r.Context()
	defer resp.WriterJSON(rw, ctx)
	log.Println("1")

	const layout = "02/01/2006" // Define the date format

	// Parse the string into time.Time
	resp.Message = response.ErrSuccess.Error()
	resp.Payload = "Clean arch pong!"
}

// HNotImplementation ...
func (h *Handler) HNotImplementation(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	ctx := r.Context()
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

	if strings.Contains(r.URL.String(), "yaml") {
		http.ServeFile(w, r, filepath.Join(swaggerPath, "swagger.yaml"))
		return
	}

	http.ServeFile(w, r, filepath.Join(swaggerPath, "swagger.json"))
}
