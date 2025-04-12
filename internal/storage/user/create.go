package user

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"gorm.io/gorm"
)

func (p *provider) Create(data entities.User, ctx context.Context) (user entities.User, err error) {
	err = p.postgres.Where(data).
		First(&user).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			p.logger.Error().Err(err).Interface("phone", data.ID).Msg("An error occurred while trying to get user")
			err = response.ErrDataNotFound
			return
		}

		err = response.ErrInternalServer
		return
	}

	return
}
