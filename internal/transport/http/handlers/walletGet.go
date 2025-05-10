package handlers

import (
	"context"
	"net/http"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

func (h *Handler) HWalletGet(rw http.ResponseWriter, r *http.Request) {
	var (
		resp    response.Response
		ctx     context.Context = r.Context()
		wallets []entities.Wallet
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
