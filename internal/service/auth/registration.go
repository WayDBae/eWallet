package auth

import (
	"context"
	"log"
	"regexp"
	"unicode/utf8"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/WayDBae/eWallet/pkg/utils"
)

func (p *provider) Registration(data entities.Registration, ctx context.Context) (code string, err error) {
	for _, value := range []string{data.Name, data.Surname, data.Patronymic} {
		err = validateName(value)
		if err != nil {
			return
		}
	}

	_, err = p.user.GetByPhone(data.PhoneNumber, ctx)
	if err != nil && err != response.ErrDataNotFound {
		return
	}

	code = utils.GenerateOTPWithLocalRand()

	err = p.rdb.Set(data.PhoneNumber, code, ctx)
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
