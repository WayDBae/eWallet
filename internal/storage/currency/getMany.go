package currency

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"gorm.io/gorm"
)

func (p *provider) GetMany(filter entities.Currency, ctx context.Context) (currencies []entities.Currency, err error) {
	err = p.postgres.Where(filter).
		Find(&currencies).
		Error
	if err != nil {
		p.logger.Error().Err(err).Interface("filter", filter).Msg("failed to get currencies")
		if err == gorm.ErrRecordNotFound {
			err = response.ErrDataNotFound
		} else {
			err = response.ErrInternalServer
		}
	}

	return
}
