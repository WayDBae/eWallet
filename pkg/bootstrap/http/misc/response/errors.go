package response

import "errors"

// Errors used in the project. Sorted alphabetically
var (
	// Default errors
	ErrAccessDenied        = errors.New("Отказано в доступе")
	ErrBindingClientToCard = errors.New("Ошибка при привязки клиента с картой")
	ErrCheckPhoneInOrzu    = errors.New("Необработанное исключение")
	ErrDataNotFound        = errors.New("Не найдено")

	ErrInternalServer               = errors.New("Внутренняя ошибка сервера")
	ErrLimitExceeded                = errors.New("Превышен лимит попыток. Повторите позже")
	ErrInvalidData                  = errors.New("Ошибка валидации данных")
	ErrNoContent                    = errors.New("Нет контента")
	ErrNotImplementation            = errors.New("Не реализовано")
	ErrPhoneNumberExists            = errors.New("По этому номеру телефону уже зарегистрирован аккаунт")
	ErrSomethingWentWrong           = errors.New("Что-то пошло не так")
	ErrSuccess                      = errors.New("Успешно")
	ErrUnauthorized                 = errors.New("Не авторизованный пользователь")
	ErrWaitingApplicationFromBasket = errors.New("В ожидании получения новой заявки")

	ErrBadRequest    = errors.New("Неверный запрос")
	ErrSmallLenName  = errors.New("Ваше имя слишком короткое")
	ErrIncorrectName = errors.New("Ваше имя содержит символы")
	ErrIncorrectOTP  = errors.New("Неверный OTP")
	ErrPasswordLen   = errors.New("Пароль должен содержать как минимум 8 символов")
)

// Error statuses
var errorCode = map[string]int{
	// Default errors
	ErrSuccess.Error():                      200,
	ErrNoContent.Error():                    201,
	ErrBadRequest.Error():                   400,
	ErrBindingClientToCard.Error():          400,
	ErrWaitingApplicationFromBasket.Error(): 400,
	ErrSmallLenName.Error():                 400,
	ErrIncorrectName.Error():                400,
	ErrIncorrectOTP.Error():                 400,
	ErrPasswordLen.Error():                  400,
	ErrUnauthorized.Error():                 401,
	ErrAccessDenied.Error():                 403,
	ErrDataNotFound.Error():                 404,
	ErrCheckPhoneInOrzu.Error():             406,
	ErrPhoneNumberExists.Error():            409,
	ErrLimitExceeded.Error():                429,
	ErrSomethingWentWrong.Error():           500,
	ErrInternalServer.Error():               500,
	ErrNotImplementation.Error():            501,
}
