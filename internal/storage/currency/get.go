package currency

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"gorm.io/gorm"
)

func (p *provider) Get(data entities.Currency, ctx context.Context) (currency entities.Currency, err error) {
	err = p.postgres.Where(data).
		First(&currency).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			p.logger.Error().Err(err).Interface("id", data.ID).Msg("An error occurred while trying to get currency")
			err = response.ErrDataNotFound
			return
		}

		err = response.ErrInternalServer
		return
	}

	return
}
