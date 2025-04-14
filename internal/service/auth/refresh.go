package auth

import (
	"context"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/google/uuid"
)

func (p *provider) Refresh(token string, ctx context.Context) (accessToken, refreshToken string, err error) {
	claims, err := p.jwt.ParseToken(token, ctx)
	if err != nil {
		return
	}

	userID := uuid.MustParse(claims.Subject)

	user, err := p.user.Get(entities.User{
		BaseGorm: entities.BaseGorm{
			ID: userID,
		},
	}, ctx)
	if err != nil {
		return
	}

	accessToken, err = p.jwt.GenerateAccessToken(user, time.Minute*15, ctx)
	if err != nil {
		return
	}

	refreshToken, err = p.jwt.GenerateRefreshToken(user, time.Hour*24*7, ctx)
	if err != nil {
		return
	}

	return
}
