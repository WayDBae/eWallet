package wallet

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

func (p *provider) Create(data entities.Wallet, ctx context.Context) (wallet entities.Wallet, err error) {
	err = p.postgres.WithContext(ctx).Create(&data).Error
	if err != nil {
		p.logger.Error().Err(err).Interface("wallet", data).Msg("Failed to create wallet")
		err = response.ErrInternalServer
		return
	}

	return data, nil
}
