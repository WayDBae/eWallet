package auth

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"unicode/utf8"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/WayDBae/eWallet/pkg/utils"
)

func (p *provider) Registration(data entities.AuthRegistration, ctx context.Context) (code string, err error) {
	for _, value := range []string{data.Name, data.Surname, data.Patronymic} {
		err = validateName(value)
		if err != nil {
			return
		}
	}

	_, err = p.user.GetByPhone(data.PhoneNumber, ctx)
	switch {
	case err == nil:
		err = response.ErrPhoneNumberExists
		return
	case errors.Is(err, response.ErrDataNotFound):
		err = nil
	default:
		return
	}

	code = utils.GenerateOTPWithLocalRand()

	data.OTPCode = code

	marshaledData, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = p.rdb.Set(data.PhoneNumber, string(marshaledData), ctx)
	if err != nil {
		return
	}

	value, err := p.rdb.Get(data.PhoneNumber, ctx)
	if err != nil {
		return
	}

	log.Println(value)

	return
}

// validateName проверяет корректность имени, фамилии или отчества
func validateName(name string) error {
	if name == "" || utf8.RuneCountInString(name) <= 2 {
		return response.ErrSmallLenName
	}

	// Проверка, что имя содержит только буквы
	matched, err := regexp.MatchString(`^[a-zA-Zа-яА-Я]+$`, name)
	if err != nil {
		return response.ErrSomethingWentWrong
	}
	if !matched {
		return response.ErrIncorrectName
	}

	return nil
}
