package handlers

import (
	"context"
	"net/http"

	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

func (h *Handler) HWalletGetMany(rw http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  context.Context = r.Context()
	)

	defer resp.WriterJSON(rw, ctx)

	// Execution of business logic
	wallets, err := h.wallet.GetMany(ctx)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	// Sending a response
	resp.Message = response.ErrSuccess.Error()
	resp.Payload = wallets
}

// swagger:operation GET /wallet/getMany Wallet walletGetMany
//
// Кошельки пользователя
//
// ## Роут предназначен для получения всех кошельков пользователя
//
// ---
//
// security:
//   - Bearer: []
//
// responses:
//   200:
//     description: |-
//       ## Успешно
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
//         <li>Неправильный логин или пароль</li>
//         <li>Одно или несколько полей пустые</li>
//       </ul>
//   500:
//     description: |-
//       Internal Server Error or Something went wrong
//     schema:
//       $ref: "#/responses/internalServer/schema"
