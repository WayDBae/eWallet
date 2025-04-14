package auth

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

func (p *provider) OTPVerify(data entities.AuthOTPVerify, ctx context.Context) (accessToken, refreshToken string, err error) {
	_, err = strconv.Atoi(data.PhoneNumber[len(data.PhoneNumber)-9:])
	if err != nil {
		err = response.ErrBadRequest
		return
	}

	_, err = strconv.Atoi(data.PhoneNumber)
	if err != nil {
		err = response.ErrBadRequest
		return
	}

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
