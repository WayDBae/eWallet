package wallet

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/google/uuid"
)

func (p *provider) GetMany(ctx context.Context) (data []map[string]any, err error) {
	userID, ok := ctx.Value(entities.ContextUserIDKey).(uuid.UUID)
	if !ok {
		err = response.ErrAccessDenied
		return
	}

	var user entities.User

	user, err = p.user.Get(entities.User{
		BaseGorm: entities.BaseGorm{
			ID: userID,
		},
	}, ctx)
	if err != nil {
		return
	}

	var wallets []entities.Wallet
	wallets, err = p.wallet.GetMany(entities.Wallet{
		UserID: user.ID,
	}, ctx)
	if err != nil && err != response.ErrDataNotFound {
		return
	}

	data = make([]map[string]any, 0, len(wallets))

	for _, wallet := range wallets {
		var currency entities.Currency

		currency, err = p.currency.Get(entities.Currency{
			BaseGorm: entities.BaseGorm{
				ID: wallet.CurrencyID,
			},
		}, ctx)
		if err != nil {
			return
		}

		data = append(data, map[string]any{
			"currency_code":      currency.Code,
			"currency_character": currency.Character,
			"balance":            wallet.Balance.StringFixed(2),
			"updated_at":         wallet.UpdatedAt,
		})
	}

	return
}
