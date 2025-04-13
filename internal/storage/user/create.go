package user

import (
	"context"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/google/uuid"
)

func (p *provider) Create(data entities.User, ctx context.Context) (user entities.User, err error) {
	// Создание нового пользователя
	data.ID = uuid.New()

	err = p.postgres.WithContext(ctx).Create(&data).Error
	if err != nil {
		p.logger.Error().Err(err).Interface("user", data).Msg("Failed to create user")
		err = response.ErrInternalServer
		return
	}

	user = data
	return
}
