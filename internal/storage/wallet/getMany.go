package wallet

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"gorm.io/gorm"
)

func (p *provider) GetMany(filter entities.Wallet, ctx context.Context) (wallets []entities.Wallet, err error) {
	err = p.postgres.Where(filter).
		Find(&wallets).
		Error
	if err != nil {
		p.logger.Error().Err(err).Interface("filter", filter).Msg("failed to get wallets")
		if err == gorm.ErrRecordNotFound {
			err = response.ErrDataNotFound
		} else {
			err = response.ErrInternalServer
		}
	}

	return
}
