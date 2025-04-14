package entities

// SwaggerInn ...
type SwaggerInn struct {
	// ИНН клиента
	//
	// example: 123456789
	// required: true
	Inn string `json:"inn"`
}

type swaggerOtpCode struct {
	// Одноразовый код
	//
	// required: true
	// max length: 4
	// default: 1234
	Code string `json:"otp_code"`
}

type swaggerPassword struct {
	// Пароль
	//
	// required: true
	// min length: 8
	// max length: 100
	// default: !qwerty123
	Password string `json:"password"`
}

type swaggerPhone struct {
	// Телефонный номер
	//
	// required: true
	// default: 992911170603
	// pattern: 992\d{9}$
	Phone string `json:"phone_number"`
}

type swaggerName struct {
	// Имя пользователя
	//
	// required: true
	// default: Далер
	Name string `json:"name"`
}

type swaggerSurname struct {
	// Фамилия пользователя
	//
	// required: true
	// default: Хайраков
	Surname string `json:"surname"`
}

type swaggerPatronymic struct {
	// Отчество пользователя
	//
	// required: true
	// default: Химатджонович
	Patronymic string `json:"patronymic"`
}

type swaggerRefreshToken struct {
	// Рефреш токен
	//
	// required: true
	RefreshToken string `json:"refresh_token"`
}
