package jwt

import (
	"context"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/golang-jwt/jwt/v5"
)

func (p *provider) GenerateRefreshToken(user entities.User, t time.Duration, ctx context.Context) (accessToken string, err error) {
	// Секретный ключ для подписи JWT-токена
	// Необходимо хранить в безопасном месте
	var refreshSecretKey = []byte(p.config.Server.RefreshSecretKey)

	// Генерируем полезные данные, которые будут храниться в токене
	claims := entities.CustomClaims{
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(refreshSecretKey)
}
