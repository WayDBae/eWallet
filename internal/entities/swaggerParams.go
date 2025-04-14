package entities

// Authorization
//
// Строение роутов для авторизации

// swagger:parameters authLogin
type swaggerAuthLogin struct {
	// in:body
	Body struct {
		swaggerPassword
		swaggerPhone
	}
}

// swagger:parameters authRegistration
type swaggerAuthRegistration struct {
	// in:body
	Body struct {
		swaggerName
		swaggerSurname
		swaggerPatronymic
		swaggerPhone
		swaggerPassword
	}
}

// swagger:parameters authOTPVerify
type swaggerAuthOTPVerify struct {
	//in:body
	Body struct {
		swaggerPhone
		swaggerOtpCode
	}
}

// swagger:parameters authRefreshToken
type swaggerAuthRefresh struct {
	//in:body
	Body struct {
		swaggerRefreshToken
	}
}
