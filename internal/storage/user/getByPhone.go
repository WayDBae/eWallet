package user

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"gorm.io/gorm"
)

func (p *provider) GetByPhone(phone string, ctx context.Context) (user entities.User, err error) {
	err = p.postgres.Where("phone_number = ?", phone).
		First(&user).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			p.logger.Error().Err(err).Str("phone", phone).Msg("An error occurred while trying to get user by phone number")
			err = response.ErrDataNotFound
			return
		}

		err = response.ErrInternalServer
		return
	}

	return
}
