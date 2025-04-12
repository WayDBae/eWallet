package rdb

import (
	"context"
	"time"
)

func (p *provider) Set(phone, data string, ctx context.Context) (err error) {
	key := "user:" + phone
	err = p.client.Set(ctx, key, data, 5*time.Minute).Err()
	p.logger.Debug().Err(err).Str("phone", phone).Str("data", data).Msg("Setting data")
	if err != nil {
		return
	}

	return
}
