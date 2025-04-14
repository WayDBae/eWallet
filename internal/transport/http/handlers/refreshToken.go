package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

type RefreshToken struct {
	Token string `json:"refresh_token"`
}

// HRefreshToken - Обновление токена
func (h *Handler) HRefreshToken(rw http.ResponseWriter, r *http.Request) {
	var (
		resp            response.Response
		ctx             context.Context = r.Context()
		oldRefreshToken RefreshToken
		// Extracting data from a request
	)

	defer resp.WriterJSON(rw, ctx)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Сheck the integrity of the received data
	err := decoder.Decode(&oldRefreshToken)
	if err != nil || oldRefreshToken.Token == "" {
		resp.Message = response.ErrInvalidToken.Error()
		return
	}

	accessToken, refreshToken, err := h.auth.Refresh(oldRefreshToken.Token, ctx)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	// Sending a response
	resp.Message = response.ErrSuccess.Error()
	resp.Payload = map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
}

// swagger:operation POST /auth/refresh Authorization authRefreshToken
//
// Обновление токенов
//
// ## Роут предназначен для получения новой пары токенов (access и refresh) по действующему refresh-токену.
//
// Пользователь должен передать refresh токен в теле запроса.
// Access токен не требуется.
//
// ---
//
// responses:
//   200:
//     description: |-
//       ## Успешная авторизаиця
//     schema:
//       $ref: "#/responses/success/schema"
//   400:
//     schema:
//       $ref: "#/responses/badRequest/schema"
//     description: |-
//       ## Неверный запрос
//       Возможно, вам стоит перепроверить введенные данные
//
//       Все возможные сообщения об ошибках в полезной нагрузке (payload):
//       <ul>
//         <li>Отсутствует refresh токен</li>
//         <li>Невалидный формат токена</li>
//       </ul>
//   401:
//     description: |-
//       ## Недействительный или просроченный refresh токен
//     schema:
//       $ref: "#/responses/unauthorized/schema"
//   429:
//     description: |-
//       ## Retry Limit Exceeded
//       Количество попыток: <code>4</code>
//       Сбросить лимит после <code>5 минуты</code>
//     schema:
//       $ref: "#/responses/retryLimitExceeded/schema"
//   500:
//     description: |-
//       Internal Server Error or Something went wrong
//     schema:
//       $ref: "#/responses/internalServer/schema"
