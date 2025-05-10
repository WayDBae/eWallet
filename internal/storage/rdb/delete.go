package rdb

import (
	"context"

	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

func (p *provider) Delete(phone string, ctx context.Context) (err error) {
	key := "user:" + phone

	err = p.client.Del(ctx, key).Err()
	if err != nil {
		p.logger.Error().Err(err).Str("phone", phone).Ctx(ctx).Msg("An error occurred while trying to deleting data from redis")
		err = response.ErrInternalServer
		return
	}

	return
}
