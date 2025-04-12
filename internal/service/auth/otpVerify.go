package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/golang-jwt/jwt/v5"
)

func (p *provider) OTPVerify(data entities.OTPVerify, ctx context.Context) (signedToken string, err error) {
	_, err = p.user.GetByPhone(data.PhoneNumber, ctx)
	if err != nil && err != response.ErrDataNotFound {
		return
	}

	value, err := p.rdb.Get(data.PhoneNumber, ctx)
	if err != nil {
		return
	}

	var reg entities.Registration
	err = json.Unmarshal([]byte(value), &reg)
	if err != nil {
		return
	}

	if reg.OTPCode != data.OTPCode {
		err = response.ErrIncorrectOTP
		return
	}

	user, err := p.user.Create(entities.User{
		Name:        reg.Name,
		Surname:     reg.Surname,
		Patronymic:  reg.Patronymic,
		Password:    reg.Password,
		PhoneNumber: reg.PhoneNumber,
	}, ctx)

	if err != nil {
		return
	}

	// Секретный ключ для подписи JWT-токена
	// Необходимо хранить в безопасном месте
	var jwtSecretKey = []byte(p.config.Server.SecretKey)
	// Генерируем полезные данные, которые будут храниться в токене
	payload := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err = token.SignedString(jwtSecretKey)
	if err != nil {
		p.logger.Error().Err(err).Interface("user_id", user.ID).Msg("JWT token signing")
		return
	}

	return
}
