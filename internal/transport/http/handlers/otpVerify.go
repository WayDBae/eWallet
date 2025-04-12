package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

// HOTPVerify ...
func (h *Handler) HOTPVerify(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	ctx := r.Context()
	defer resp.WriterJSON(rw, ctx)

	// Extracting data from a request
	var data entities.OTPVerify

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Сheck the integrity of the received data
	err := decoder.Decode(&data)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	_, err = strconv.Atoi(data.PhoneNumber[len(data.PhoneNumber)-9:])
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	_, err = strconv.Atoi(data.PhoneNumber)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	_, err = strconv.Atoi(data.OtpCode)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	if data.OtpCode == "" || len(data.OtpCode) != 4 {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	// Execution of business logic
	err = h.auth.OTPVerify(data, ctx)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	var test map[string]any = map[string]any{
		"access_token": "Типо access_token",
	}
	// Sending a response
	resp.Message = response.ErrSuccess.Error()
	resp.Payload = test
}

// swagger:operation POST /auth/otp-verify Authorization authOTPVerify
//
// Подтверждение номера пользователя по OTP после регистрации
//
// ## Роут предназначен для авторизации ранее зарегистрированного пользователя
// Используемый <b>Authorization Flow</b> подразумевает сначала проверка введенных пользователем <b>credential</b> (phone, password)
// После успешной проверки, пользователю высылается на номер телефона OTP код с временем жизни в <code>1 минуту</code>
//
// ---
//
// responses:
//   200:
//     description: |-
//       ## Успешная авторизация
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
