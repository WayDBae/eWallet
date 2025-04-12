package auth

import (
	"context"
	"log"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
)

func (p *provider) OTPVerify(data entities.OTPVerify, ctx context.Context) (err error) {
	_, err = p.user.GetByPhone(data.PhoneNumber, ctx)
	if err != nil && err != response.ErrDataNotFound {
		return
	}

	value, err := p.rdb.Get(data.PhoneNumber, ctx)
	if err != nil {
		return
	}

	if value != data.OtpCode {
		err = response.ErrIncorrectOTP
		return
	}

	log.Println(value)

	return
}
