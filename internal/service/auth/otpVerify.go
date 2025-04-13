package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/golang-jwt/jwt/v5"
)

func (p *provider) OTPVerify(data entities.AuthOTPVerify, ctx context.Context) (signedToken string, err error) {
	value, err := p.rdb.Get(data.PhoneNumber, ctx)
	if err != nil {
		return
	}

	var reg entities.AuthRegistration
	err = json.Unmarshal([]byte(value), &reg)
	if err != nil {
		return
	}

	if reg.OTPCode != data.OTPCode {
		err = response.ErrIncorrectOTP
		return
	}

	_, err = p.user.GetByPhone(data.PhoneNumber, ctx)
	if err != nil && err != response.ErrDataNotFound {
		return
	}

	if err == nil {
		err = p.rdb.Delete(data.PhoneNumber, ctx)
		if err != nil {
			return
		}
		err = response.ErrPhoneNumberExists
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

	err = p.rdb.Delete(user.PhoneNumber, ctx)
	if err != nil {
		return
	}

	// Секретный ключ для подписи JWT-токена
	// Необходимо хранить в безопасном месте
	var jwtSecretKey = []byte(p.config.Server.SecretKey)

	// Генерируем полезные данные, которые будут храниться в токене
	claims := entities.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString(jwtSecretKey)
	if err != nil {
		p.logger.Error().Err(err).Interface("user_id", user.ID).Msg("JWT token signing")
		return
	}

	return
}
