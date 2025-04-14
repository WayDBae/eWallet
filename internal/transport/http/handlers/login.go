package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

// HLogin - Вход
func (h *Handler) HLogin(rw http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  context.Context = r.Context()
		// Extracting data from a request
		data entities.AuthLogin
	)

	defer resp.WriterJSON(rw, ctx)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Сheck the integrity of the received data
	err := decoder.Decode(&data)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	accessToken, refreshToken, err := h.auth.Login(data, ctx)
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

// swagger:operation POST /auth/otp-verify Authorization authOTPVerify
//
// Подтверждение номера телефона пользователя по OTP после регистрации
//
// ## Роут предназначен для подтверждения номера телефона пользователя после регистрации.
//
// Используемый <b>OTP Verification Flow</b> подразумевает, что пользователь вводит полученный OTP-код, который был отправлен на его номер телефона.
// Время жизни OTP-кода составляет <code>1 минуту</code>. Если код подтверждения правильный, пользователь может продолжить процесс регистрации или авторизации.

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
//         <li>Не верный OTP код</li>
//         <li>Неправильный логин или пароль</li>
//         <li>Одно или несколько полей пустые</li>
//       </ul>
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
