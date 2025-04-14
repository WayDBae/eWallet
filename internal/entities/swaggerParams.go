package entities

// Authorization
//
// Строение роутов для авторизации

// swagger:parameters authLogin
type swaggerAuthLogin struct {
	// in:body
	Body struct {
		swaggerPassword
		swaggerPhoneNumber
	}
}

// swagger:parameters authRegistration
type swaggerAuthRegistration struct {
	// in:body
	Body struct {
		swaggerName
		swaggerSurname
		swaggerPatronymic
		swaggerPhoneNumber
		swaggerPassword
	}
}

// swagger:parameters authRefreshToken
type swaggerAuthRefresh struct {
	// in:body
	Body struct {
		swaggerRefreshToken
	}
}

// swagger:parameters authOTPVerify
type swaggerAuthOTPVerify struct {
	// in:body
	Body struct {
		swaggerPhoneNumber
		swaggerOtpCode
	}
}