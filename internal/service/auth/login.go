package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

func (p *provider) Login(data entities.AuthLogin, ctx context.Context) (accessToken, refreshToken string, err error) {
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

	user, err := p.user.GetByPhone(data.PhoneNumber, ctx)
	if err != nil && err != response.ErrDataNotFound {
		return
	}

	if user.Password != data.Password {
		err = response.ErrWrongPassword
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
