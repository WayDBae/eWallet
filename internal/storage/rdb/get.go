package rdb

import (
	"context"

	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/go-redis/redis"
)

func (p *provider) Get(phone string, ctx context.Context) (value string, err error) {
	key := "user:" + phone

	value, err = p.client.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			p.logger.Warn().Err(err).Str("phone", phone).Ctx(ctx).Msg("An error occurred while trying to data get from redis")
			err = response.ErrDataNotFound
			return
		}
		p.logger.Error().Err(err).Str("phone", phone).Ctx(ctx).Msg("An error occurred while trying to data get from redis")
		err = response.ErrInternalServer
		return
	}

	return
}
