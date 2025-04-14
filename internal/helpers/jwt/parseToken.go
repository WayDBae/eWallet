package jwt

import (
	"context"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/golang-jwt/jwt/v5"
)

func (p *provider) ParseToken(tokenStr string, ctx context.Context) (claims *entities.CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &entities.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(p.config.Server.RefreshSecretKey), nil
	})
	if err != nil {
		err = response.ErrInvalidToken
		return
	}

	claims, ok := token.Claims.(*entities.CustomClaims)
	if !ok || !token.Valid {
		err = response.ErrInvalidToken
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		err = response.ErrExpiredToken
		return
	}

	return
}
