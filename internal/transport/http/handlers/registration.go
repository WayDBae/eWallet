package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

// HLogin - Вход
func (h *Handler) HRegistration(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	ctx := r.Context()
	defer resp.WriterJSON(rw, ctx)

	// Extracting data from a request
	var data entities.AuthRegistration

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

	if len(data.Password) < 8 {
		resp.Message = response.ErrPasswordLen.Error()
		return
	}

	// Execution of business logic
	code, err := h.auth.Registration(data, ctx)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	// Sending a response
	resp.Message = response.ErrSuccess.Error()
	resp.Payload = code
}

// swagger:operation POST /auth/registration Authorization authRegistration
//
// Регистрация нового пользователя
//
// ## Роут предназначен для регистрации нового пользователя в системе
// Используемый <b>Registration Flow</b> подразумевает отправку номера телефона и пароля
// После получения данных, на указанный номер телефона отправляется OTP-код для подтверждения личности
// Время жизни OTP-кода составляет <code>5 минут</code>
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
//         <li>Ваше имя слишком короткое</li>
//         <li>Ваше имя содержит символы</li>
//         <li>Пароль должен содержать как минимум 8 символов</li>
//       </ul>
//   500:
//     description: |-
//       Internal Server Error or Something went wrong
//     schema:
//       $ref: "#/responses/internalServer/schema"
