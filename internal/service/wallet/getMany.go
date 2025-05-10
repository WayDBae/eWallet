package wallet

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/google/uuid"
)

func (p *provider) GetMany(ctx context.Context) (wallets []entities.Wallet, err error) {
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

	wallets, err = p.wallet.GetMany(entities.Wallet{
		UserID: user.ID,
	}, ctx)
	if err != nil && err != response.ErrDataNotFound {
		return
	}

	return
}
