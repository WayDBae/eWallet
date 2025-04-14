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
	resp.Payload = "Clean arch pong! 🥎"
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
		return
	}

	// Формируем путь к файлам
	// swaggerPath := filepath.Join(pwd, "../pkg/docs/")

	swaggerPath := "../pkg/docs/"

	h.logger.Debug().Str("pwd", pwd).Str("swaggerPath", swaggerPath).Msg("Debugging swagger path")

	// Если все файлы существуют, отдаём их
	if strings.Contains(r.URL.String(), "yaml") {
		// Проверяем, существует ли файл swagger.yaml
		yamlPath := filepath.Join(swaggerPath, "swagger.yaml")
		if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
			h.logger.Error().Err(err).Str("file", yamlPath).Msg("swagger.yaml not found")
			return
		}
		http.ServeFile(w, r, yamlPath)
	} else {
		// Проверяем, существует ли файл swagger.json
		jsonPath := filepath.Join(swaggerPath, "swagger.json")
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			h.logger.Error().Err(err).Str("file", jsonPath).Msg("swagger.json not found")
			return
		}
		http.ServeFile(w, r, jsonPath)
	}
}
