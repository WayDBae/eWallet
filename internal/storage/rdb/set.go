package rdb

import (
	"context"
	"time"
)

func (p *provider) Set(phone, code string, ctx context.Context) (err error) {
	key := "user:" + phone

	err = p.client.Set(ctx, key, code, 5*time.Minute).Err()
	if err != nil {
		return
	}

	return
}
